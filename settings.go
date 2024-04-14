package main

import (
	"net/http"
	"os"
  "encoding/json"

	"github.com/go-chi/chi/v5"

	"pfeifer.dev/opwebd/components"
)

func routeSettings(r *chi.Mux) {
	r.Get("/settings", func(w http.ResponseWriter, r *http.Request) {
		settings, err := GetSettings()
    check(err)
		routes := Routes()
		err = components.Layout("Settings", routes, components.Settings(settings)).Render(r.Context(), w)
		check(err)
	})

	r.Post("/settings/{key}", func(w http.ResponseWriter, r *http.Request) {
		key := chi.URLParam(r, "key")
		value := r.FormValue("Value")
		err := PutParam(ParamPath(key, false), []byte(value))
		check(err)
	})
}

func GetSettings() ([]components.SettingGroup, error) {
	exists, err := Exists("/data/openpilot")
	settingsPath := "./settings.json"
	if exists {
		settingsPath = "/data/openpilot/settings.json"
	}

  var settings []components.SettingGroup
  settingsData, err := os.ReadFile(settingsPath)
  if err != nil {
    return settings, err
  }
  err = json.Unmarshal(settingsData, &settings)

	return settings, err
}
