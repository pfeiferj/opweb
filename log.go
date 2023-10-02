package main

import (
	"capnproto.org/go/capnp/v3"
	_ "embed"
	"encoding/json"
	"github.com/cbroglie/mustache"
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
	"strconv"
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
	r.Get("/logs/view/rlog/{key}/{segment}/{page}", func(w http.ResponseWriter, r *http.Request) {
		segment := chi.URLParam(r, "segment")
		key := chi.URLParam(r, "key")
		page := chi.URLParam(r, "page")
		parsedPage, err := strconv.Atoi(page)
		ignore(err)
		data := generateLogView(segment, "rlog", key, uint32(parsedPage))
		_, err = w.Write([]byte(data.Render()))
		check(err)
	})
	r.Get("/logs/view/qlog/{key}/{segment}/{page}", func(w http.ResponseWriter, r *http.Request) {
		segment := chi.URLParam(r, "segment")
		key := chi.URLParam(r, "key")
		page := chi.URLParam(r, "page")
		parsedPage, err := strconv.Atoi(page)
		ignore(err)
		data := generateLogView(segment, "qlog", key, uint32(parsedPage))
		_, err = w.Write([]byte(data.Render()))
		check(err)
	})
}

func getEventTypes() (map[uint16]string, error) {
	eventTypes := map[uint16]string{}
	schemas := getSchemas()
	logFile := schemas["log.capnp"]
	eventSchema, err := logFile.Nested("Event")
	if err != nil {
		return nil, err
	}

	eventStruct := eventSchema.AsStruct()
	unionFields := eventStruct.UnionFields()

	for _, field := range unionFields {
		p, err := field.Proto()
		if err != nil {
			return nil, err
		}
		eventTypes[p.DiscriminantValue()] = p.Name()
	}
	return eventTypes, nil
}

//go:embed templates/view_log.html
var viewLogTemplate string

func generateLogView(segment string, typ string, key string, page uint32) LogView {
	logPath := BasePath + segment + "/" + typ

	schemas := getSchemas()

	logFile := schemas["log.capnp"]
	eventSchema, err := logFile.Nested("Event")
	if err != nil {
		panic(err)
	}

	eventStruct := eventSchema.AsStruct()

	logDat, err := os.ReadFile(logPath)
	check(err)

	eventTypes, err := getEventTypes()
	check(err)

	offset := uint64(0)
	count := uint32(0)
	start := (page - 1) * 10

	logView := LogView{Logs: make([]LogViewPart, 10)}
	for offset < uint64(len(logDat)) {
		msg, err := capnp.Unmarshal(logDat[offset:])
		check(err)

		root, err := msg.Root()
		check(err)
		rootStruct := root.Struct()
		which := rootStruct.Uint16(8)
		name := eventTypes[which]

		if name == key {
			if count < start {
				count += 1
			} else {
				data, err := eventStruct.Decode(logDat[offset:])
				check(err)

				var obj map[string]interface{}
				err = json.Unmarshal(data, &obj)
				check(err)

				pretty, err := json.MarshalIndent(obj, "", "  ")
				check(err)
				logView.Logs[count-start].Log = string(pretty)

				count += 1
				if count-start == 10 {
					break
				}
			}
		}

		size, err := msg.TotalSize()
		check(err)
		offset += size
	}
	return logView
}

type LogViewPart struct {
	Log string
}

type LogView struct {
	Logs []LogViewPart
}

func (l LogView) Render() string {
	page, err := mustache.RenderInLayoutPartials(viewLogTemplate, basePartial, PartialProvider, l)
	check(err)
	return page
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
