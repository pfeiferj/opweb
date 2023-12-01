package main

import (
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/go-chi/chi/v5"

	"pfeifer.dev/opwebd/components"
)

func routeRoute(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		routes := Routes()
		route := routes[0]
		segments := segmentData(route)
		components.Layout("Route", routes, components.Route(route, segments)).Render(r.Context(), w)
	})

	r.Get("/{route}", func(w http.ResponseWriter, r *http.Request) {
		route := chi.URLParam(r, "route")
		routes := Routes()
		segments := segmentData(route)
		components.Layout("Route", routes, components.Route(route, segments)).Render(r.Context(), w)
	})

	r.Get("/routes/{state}", func(w http.ResponseWriter, r *http.Request) {
		state := chi.URLParam(r, "state")
		routes := Routes()
		components.RoutesDropdown(routes, state == "open").Render(r.Context(), w)
	})
}

func Routes() []string {
	files, err := os.ReadDir(BasePath)
	check(err)

	route_dirs := make(map[string]bool)
	for _, file := range files {
		name := file.Name()

		split := strings.SplitN(name, "--", 4)
		if len(split) == 3 {
			route_dirs[split[0]+"--"+split[1]] = true
		}
	}

	route_dirs_sorted := make([]string, len(route_dirs))

	i := 0
	for k := range route_dirs {
		route_dirs_sorted[i] = k
		i++
	}

	slices.Sort(route_dirs_sorted)
	slices.Reverse(route_dirs_sorted)

	return route_dirs_sorted
}
