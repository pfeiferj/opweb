package main

import (
	"embed"
	"github.com/go-chi/chi/v5"
	"net/http"
	"path"
	"path/filepath"
	"strconv"
)

//go:embed static/*
var f embed.FS

//go:embed static/favicon.ico
var favicon []byte

func routeFiles(r *chi.Mux) {
	r.Handle("/static/*", http.FileServer(http.FS(f)))
	r.Get("/files/*", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "*")
		path := filepath.Join(BasePath, filepath.FromSlash(path.Clean("/"+name)))
		http.ServeFile(w, r, path)
	})

	r.Get("/download/*", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "*")
		path := filepath.Join(BasePath, filepath.FromSlash(path.Clean("/"+name)))

		w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(filepath.Base(filepath.Dir(path))+"-"+filepath.Base(path)))
		w.Header().Set("Content-Type", "application/octet-stream")

		http.ServeFile(w, r, path)
	})
	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(favicon)
		check(err)
	})
}
