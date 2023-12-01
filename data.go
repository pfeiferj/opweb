package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"pfeifer.dev/opwebd/components"
)

type DataSegment struct {
	Name          string
	Time          float64
	FormattedTime string
}

type RouteData struct {
	Segments []DataSegment
}

func segmentData(route string) []components.Segment {
	segment_dirs := Segments(route)

	segments := make([]components.Segment, len(segment_dirs))

	durationChannels := make([]chan float64, len(segment_dirs))
	for i, dir := range segment_dirs {
		durationChannels[i] = make(chan float64)
		go getDuration(durationChannels[i], BasePath+dir+"/qcamera.ts")
	}
	for i, dir := range segment_dirs {
		segments[i].Name = dir
		duration := <-durationChannels[i]
		baseTime := float64(0)
		if i > 0 {
			baseTime = segments[i-1].Time
		}
		segments[i].Time = baseTime + duration
		var t time.Time
		t = t.Add(time.Duration(int64(baseTime+duration)) * time.Second)
		segments[i].FormattedTime = t.Format("15:04:05")
	}

	return segments
}

func routeData(r *chi.Mux) {
	r.Get("/data/{segment}", func(w http.ResponseWriter, r *http.Request) {
		segment := chi.URLParam(r, "segment")
		components.SegmentData(segment).Render(r.Context(), w)
	})
}
