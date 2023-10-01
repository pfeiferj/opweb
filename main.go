package main

import (
	"net/http"

	_ "embed"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//go:embed partials/base.html
var basePartial string

func main() {
	EnsureParamDirectories()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	routeRoute(r)
	routePlaylist(r)
	routeFiles(r)
	routeGpx(r)
	routeLogs(r)
	routeData(r)

	err := http.ListenAndServe(":3000", r)
	check(err)
}
