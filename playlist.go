package main

import (
	_ "embed"
	"fmt"
	"github.com/cbroglie/mustache"
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func routePlaylist(r *chi.Mux) {
	r.Get("/playlist/{route}.m3u8", func(w http.ResponseWriter, r *http.Request) {
		route := chi.URLParam(r, "route")
		playlist := generatePlaylist(route)
		_, err := w.Write([]byte(playlist.Render()))
		check(err)
	})
}

type PlaylistSegment struct {
	Index    string
	Duration string
	Link     string
}

type Playlist struct {
	Segments []PlaylistSegment
}

//go:embed templates/playlist.m3u8
var playlistTemplate string

func (p Playlist) Render() string {
	page, err := mustache.Render(playlistTemplate, p)
	check(err)
	return page
}

func generatePlaylist(route string) Playlist {
	files, err := os.ReadDir(BasePath)
	check(err)

	segment_dirs := []string{}
	for _, file := range files {
		name := file.Name()

		if strings.HasPrefix(name, route) {
			segment_dirs = append(segment_dirs, name)
		}
	}

	playlist := Playlist{Segments: make([]PlaylistSegment, len(segment_dirs))}

	durationChannels := make([]chan float64, len(segment_dirs))
	for i, dir := range segment_dirs {
		durationChannels[i] = make(chan float64)
		go getDuration(durationChannels[i], BasePath+dir+"/qcamera.ts")
	}
	for i, dir := range segment_dirs {
		playlist.Segments[i].Index = fmt.Sprint(i)
		playlist.Segments[i].Link = "/files/" + dir + "/qcamera.ts"
		duration := <-durationChannels[i]
		playlist.Segments[i].Duration = fmt.Sprintf("%.3f", duration)
	}

	return playlist
}

func getDuration(out chan float64, video string) {
	cmd := exec.Command("bash", "-c", "ffprobe -i "+video+" -show_entries format=duration -v quiet -of csv='p=0'")
	output, _ := cmd.Output()
	duration, _ := strconv.ParseFloat(strings.TrimSpace(string(output)), 64)
	out <- duration
}
