package main

import (
	_ "embed"
	"github.com/cbroglie/mustache"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
)

type DataSegment struct {
	Name          string
	Time          float64
	FormattedTime string
}

type RouteData struct {
	Segments []DataSegment
}

func generateRouteData(route string) RouteData {
	segment_dirs := Segments(route)

	playlist := RouteData{Segments: make([]DataSegment, len(segment_dirs))}

	durationChannels := make([]chan float64, len(segment_dirs))
	for i, dir := range segment_dirs {
		durationChannels[i] = make(chan float64)
		go getDuration(durationChannels[i], BasePath+dir+"/qcamera.ts")
	}
	for i, dir := range segment_dirs {
		playlist.Segments[i].Name = dir
		duration := <-durationChannels[i]
		baseTime := float64(0)
		if i > 0 {
			baseTime = playlist.Segments[i-1].Time
		}
		playlist.Segments[i].Time = baseTime + duration
		var t time.Time
		t = t.Add(time.Duration(int64(baseTime+duration)) * time.Second)
		playlist.Segments[i].FormattedTime = t.Format("15:04:05")
	}

	return playlist
}

type SegmentData struct {
	Name string
}

func (s SegmentData) Render() string {
	page, err := mustache.Render(segmentDataTemplate, s)
	check(err)
	return page
}

//go:embed templates/segment_data.html
var segmentDataTemplate string

func routeData(r *chi.Mux) {
	r.Get("/data/{segment}", func(w http.ResponseWriter, r *http.Request) {
		segment := chi.URLParam(r, "segment")
		segmentData := SegmentData{Name: segment}
		_, err := w.Write([]byte(segmentData.Render()))
		check(err)
	})
}
