package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

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
	routeUpdate(r)
	routeParams(r)

	err := http.ListenAndServe(":3000", r)
	check(err)
}
