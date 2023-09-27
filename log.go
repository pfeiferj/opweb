package main

import (
	"capnproto.org/go/capnp/v3"
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
)

func routeLogs(r *chi.Mux) {
	r.Get("/logs/rlog/{segment}.json", func(w http.ResponseWriter, r *http.Request) {
		segment := chi.URLParam(r, "segment")
		writeLogs(w, segment, "rlog")
	})
	r.Get("/logs/qlog/{segment}.json", func(w http.ResponseWriter, r *http.Request) {
		segment := chi.URLParam(r, "segment")
		writeLogs(w, segment, "qlog")
	})
}

func writeLogs(w http.ResponseWriter, segment string, typ string) {
	schemas := getSchemas()

	logFile := schemas["log.capnp"]
	eventSchema, err := logFile.Nested("Event")
	if err != nil {
		panic(err)
	}

	eventStruct := eventSchema.AsStruct()

	logPath := BasePath + segment + "/" + typ
	logDat, err := os.ReadFile(logPath)
	check(err)

	w.Header().Set("Content-Disposition", "attachment; filename=\""+segment+"-"+typ+".json\"")

	offset := uint64(0)
	for offset < uint64(len(logDat)) {
		msg, err := capnp.Unmarshal(logDat[offset:])
		check(err)

		data, err := eventStruct.Decode(logDat[offset:])
		check(err)

		_, err = w.Write([]byte(data))
		check(err)

		_, err = w.Write([]byte("\n"))
		check(err)

		size, err := msg.TotalSize()
		check(err)
		offset += size
	}
}
