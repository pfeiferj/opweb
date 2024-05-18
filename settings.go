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

	r.Post("/settings", func(w http.ResponseWriter, r *http.Request) {
    // get all form values
    err := r.ParseForm()
    check(err)
    settings, err := GetSettings()
    check(err)
    for i, group := range settings {
      for j, setting := range group.Settings {
        value := r.FormValue(setting.Key)
        settings[i].Settings[j].Param.Value = value
        if settings[i].Settings[j].Type == "toggle" && value == "" {
          settings[i].Settings[j].Param.Value = "0"
        }
      }
    }
    err = putSettingsParams(settings)
    check(err)
	})

  // Set default settings
  settings, err := GetSettings()
  check(err)
  err = putSettingsParams(settings)
  check(err)
}

func GetSettings() ([]components.SettingGroup, error) {
	exists, err := Exists("/data/openpilot")
  if err != nil {
    return nil, err
  }
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

  settings = getSettingsCurrentState(settings)

	return settings, err
}

func getSettingsCurrentState(settings []components.SettingGroup) []components.SettingGroup {
  for i, group := range settings {
    for j, setting := range group.Settings {
      value, err := GetParam(ParamPath(setting.Key, false))
      if err != nil {
        settings[i].Settings[j].Param.Value = settings[i].Settings[j].DefaultValue
      } else {
        settings[i].Settings[j].Param.Value = string(value)
      }
    }
  }
  return settings
}

func putSettingsParams(settings []components.SettingGroup) error {
  for _, group := range settings {
    for _, setting := range group.Settings {
      err := PutParam(ParamPath(setting.Key, false), []byte(setting.Param.Value))
      if err != nil {
        return err
      }
    }
  }
  return nil
}

