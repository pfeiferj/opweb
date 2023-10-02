package main

import (
	"capnproto.org/go/capnp/v3"
	_ "embed"
	"github.com/cbroglie/mustache"
	"github.com/go-chi/chi/v5"
	"github.com/zsefvlol/timezonemapper"
	"math"
	"net/http"
	"os"
	"pfeifer.dev/opwebd/proto"
	"time"
)

type GpxPoint struct {
	Latitude  float64
	Longitude float64
	Elevation float64
	Bearing   float64
	Timestamp string
}

type Gpx struct {
	Points []GpxPoint
}

//go:embed templates/track.gpx
var gpxTemplate string

func (g Gpx) Render() string {
	page, err := mustache.Render(gpxTemplate, g)
	check(err)
	return page
}

func routeGpx(r *chi.Mux) {
	r.Get("/gpx/rlog/{route}.gpx", func(w http.ResponseWriter, r *http.Request) {
		route := chi.URLParam(r, "route")
		gpx := generateGpx(route, "rlog", false)

		w.Header().Set("Content-Disposition", "attachment; filename=\""+route+"-route-rlog.gpx\"")
		_, err := w.Write([]byte(gpx.Render()))
		check(err)
	})
	r.Get("/gpx/qlog/{route}.gpx", func(w http.ResponseWriter, r *http.Request) {
		route := chi.URLParam(r, "route")
		gpx := generateGpx(route, "qlog", false)
		w.Header().Set("Content-Disposition", "attachment; filename=\""+route+"-route-qlog.gpx\"")
		_, err := w.Write([]byte(gpx.Render()))
		check(err)
	})
	r.Get("/gpx/rlog/segment/{route}.gpx", func(w http.ResponseWriter, r *http.Request) {
		route := chi.URLParam(r, "route")
		gpx := generateGpx(route, "rlog", true)
		w.Header().Set("Content-Disposition", "attachment; filename=\""+route+"-segment-qlog.gpx\"")
		_, err := w.Write([]byte(gpx.Render()))
		check(err)
	})
	r.Get("/gpx/qlog/segment/{route}.gpx", func(w http.ResponseWriter, r *http.Request) {
		route := chi.URLParam(r, "route")
		gpx := generateGpx(route, "qlog", true)
		w.Header().Set("Content-Disposition", "attachment; filename=\""+route+"-segment-qlog.gpx\"")
		_, err := w.Write([]byte(gpx.Render()))
		check(err)
	})
}

func generateGpx(route string, typ string, isSegment bool) Gpx {
	var segments []string
	if isSegment {
		segments = []string{route}
	} else {
		segments = Segments(route)
	}
	points := []GpxPoint{}

	for _, segment := range segments {
		logPath := BasePath + segment + "/" + typ
		logDat, err := os.ReadFile(logPath)
		check(err)
		monoOffset := MonoOffset(logDat)
		offset := uint64(0)
		for offset < uint64(len(logDat)) {
			msg, err := capnp.Unmarshal(logDat[offset:])
			check(err)

			size, err := msg.TotalSize()
			check(err)
			offset += size

			evt, err := proto.ReadRootEvent(msg)
			if err != nil {
				continue
			}
			if evt.Which() != proto.Event_Which_liveLocationKalman {
				continue
			}

			llk, err := evt.LiveLocationKalman()
			if err != nil {
				continue
			}

			if llk.Status() != proto.LiveLocationKalman_Status_valid {
				continue
			}

			geoPos, err := llk.PositionGeodetic()
			if err != nil {
				continue
			}

			if !geoPos.Valid() {
				continue
			}

			geoPosVal, err := geoPos.Value()
			if err != nil {
				continue
			}

			orientation, err := llk.CalibratedOrientationNED()
			if err != nil {
				continue
			}

			orientationVal, err := orientation.Value()
			if err != nil {
				continue
			}

			lat := geoPosVal.At(0)
			long := geoPosVal.At(1)
			timezone := timezonemapper.LatLngToTimezoneString(lat, long)
			point := GpxPoint{
				Latitude:  lat,
				Longitude: long,
				Elevation: geoPosVal.At(2),
				Bearing:   orientationVal.At(2) * (180 / math.Pi),
				Timestamp: MonoDatetime(evt, monoOffset, timezone).Format(time.RFC3339Nano),
			}
			points = append(points, point)
		}
	}

	return Gpx{Points: points}
}
