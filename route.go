package main

import (
	_ "embed"
	"github.com/cbroglie/mustache"
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
	"sort"
	"strings"
)

type Links struct {
	Path  string
	Label string
}

type RoutesData struct {
	Title string
	Body  string
	Links []Links
	Route string
}

//go:embed templates/routes.html
var routesTemplate string

type RouteLinks struct {
	Routes []RouteLink
	Open   bool
}

type RouteLink struct {
	Route string
}

//go:embed templates/routes_dropdown.html
var routeLinksTemplate string

func routeRoute(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		routes := Routes()
		ctx := RoutesData{
			Title: "Route",
			Links: []Links{},
			Route: routes[len(routes)-1].Route,
		}
		result, err := mustache.RenderInLayoutPartials(routesTemplate, basePartial, PartialProvider, ctx)
		check(err)
		_, err = w.Write([]byte(result))
		check(err)
	})

	r.Get("/{route}", func(w http.ResponseWriter, r *http.Request) {
		route := chi.URLParam(r, "route")
		ctx := RoutesData{
			Title: "Route",
			Links: []Links{},
			Route: route,
		}
		result, err := mustache.RenderInLayoutPartials(routesTemplate, basePartial, PartialProvider, ctx)
		check(err)
		_, err = w.Write([]byte(result))
		check(err)
	})

	r.Get("/routes/{state}", func(w http.ResponseWriter, r *http.Request) {
		state := chi.URLParam(r, "state")
		routes := Routes()
		ctx := RouteLinks{Routes: routes, Open: state == "open"}
		result, err := mustache.RenderPartials(routeLinksTemplate, PartialProvider, ctx)
		check(err)
		_, err = w.Write([]byte(result))
		check(err)
	})
}

func Routes() []RouteLink {
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

	sort.Strings(route_dirs_sorted)

	route_dirs_list := make([]RouteLink, len(route_dirs))

	for i, v := range route_dirs_sorted {
		route_dirs_list[i].Route = v
	}

	return route_dirs_list
}
