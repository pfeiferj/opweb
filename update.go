package main

import (
	_ "embed"
	"github.com/cbroglie/mustache"
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
	"os/exec"
)

func routeUpdate(r *chi.Mux) {
	r.Get("/updater", func(w http.ResponseWriter, r *http.Request) {
		update := generateUpdate()
		_, err := w.Write([]byte(update.Render()))
		check(err)
	})

	r.Post("/updater/check", func(w http.ResponseWriter, r *http.Request) {
		cmd := exec.Command("pkill", "-SIGUSR1", "-f", "selfdrive.updated")
		err := cmd.Run()
		check(err)

		_, err = w.Write([]byte("triggered"))
		check(err)
	})

	r.Post("/updater/download", func(w http.ResponseWriter, r *http.Request) {
		cmd := exec.Command("pkill", "-SIGHUP", "-f", "selfdrive.updated")
		err := cmd.Run()
		check(err)

		_, err = w.Write([]byte("triggered"))
		check(err)
	})

	r.Post("/updater/install", func(w http.ResponseWriter, r *http.Request) {
		exists, err := Exists("/data/openpilot/prebuilt")
		check(err)
		if exists {
			err = os.Remove("/data/openpilot/prebuilt")
			check(err)
		}

		err = PutParam(ParamPath("DoReboot", false), []byte{'1'})
		check(err)

		_, err = w.Write([]byte("triggered"))
		check(err)
	})
}

func generateUpdate() Update {
	buttonText := "Check For Update"
	buttonLink := "/updater/check"
	updaterFetchAvailable, err := GetParam(ParamPath("UpdaterFetchAvailable", false))
	ignore(err)
	if len(updaterFetchAvailable) > 0 && updaterFetchAvailable[0] == '1' {
		buttonText = "Download Update"
		buttonLink = "/updater/download"
	}

	downloaded, err := GetParam(ParamPath("UpdateAvailable", false))
	ignore(err)

	state, err := GetParam(ParamPath("UpdaterState", false))
	ignore(err)

	version, err := GetParam(ParamPath("UpdaterNewDescription", false))
	ignore(err)

	updateBranch, err := GetParam(ParamPath("UpdaterTargetBranch", false))
	ignore(err)

	lastCheck, err := GetParam(ParamPath("LastUpdateTime", false))
	ignore(err)

	updateNotes, err := GetParam(ParamPath("UpdaterNewReleaseNotes", false))
	ignore(err)

	updateFailedCount, err := GetParam(ParamPath("UpdateFailedCount", false))
	ignore(err)

	currentVersion, err := GetParam(ParamPath("UpdaterCurrentDescription", false))
	ignore(err)

	currentRepo, err := GetParam(ParamPath("GitRemote", false))
	ignore(err)

	currentBranch, err := GetParam(ParamPath("GitBranch", false))
	ignore(err)

	currentInstallDate, err := GetParam(ParamPath("InstallDate", false))
	ignore(err)

	currentReleaseNotes, err := GetParam(ParamPath("UpdaterCurrentReleaseNotes", false))
	ignore(err)

	currentCommit, err := GetParam(ParamPath("GitCommit", false))
	ignore(err)

	return Update{
		ButtonText:          buttonText,
		ButtonLink:          buttonLink,
		IdleButton:          Equal(state, []byte("idle")),
		InstallButton:       len(downloaded) > 0 && downloaded[0] == '1',
		UpdateState:         string(state),
		UpdateVersion:       string(version),
		UpdateBranch:        string(updateBranch),
		UpdateLastCheck:     string(lastCheck),
		UpdateReleaseNotes:  string(updateNotes),
		UpdateFailureCount:  string(updateFailedCount),
		CurrentVersion:      string(currentVersion),
		CurrentRepo:         string(currentRepo),
		CurrentBranch:       string(currentBranch),
		CurrentInstallDate:  string(currentInstallDate),
		CurrentReleaseNotes: string(currentReleaseNotes),
		CurrentCommit:       string(currentCommit),
		Title:               "Updater",
		Links:               PageLinks,
	}
}

//go:embed templates/update.html
var updateTemplate string

type Update struct {
	IdleButton          bool
	ButtonText          string
	ButtonLink          string
	InstallButton       bool
	UpdateVersion       string
	UpdateBranch        string
	UpdateState         string
	UpdateLastCheck     string
	UpdateFailureCount  string
	UpdateReleaseNotes  string
	CurrentVersion      string
	CurrentInstallDate  string
	CurrentRepo         string
	CurrentBranch       string
	CurrentCommit       string
	CurrentReleaseNotes string
	Title               string
	Links               []Links
}

func (u Update) Render() string {
	page, err := mustache.RenderInLayoutPartials(updateTemplate, basePartial, PartialProvider, u)
	check(err)
	return page
}
