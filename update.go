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
			err = os.Remove("GeeksforGeeks.txt")
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
	check(err)
	if updaterFetchAvailable[0] == '1' {
		buttonText = "Download Update"
		buttonLink = "/updater/download"
	}

	downloaded, err := GetParam(ParamPath("UpdateAvailable", false))
	check(err)

	state, err := GetParam(ParamPath("UpdaterState", false))
	check(err)

	version, err := GetParam(ParamPath("UpdaterNewDescription", false))
	check(err)

	updateBranch, err := GetParam(ParamPath("UpdaterTargetBranch", false))
	check(err)

	lastCheck, err := GetParam(ParamPath("LastUpdateTime", false))
	check(err)

	updateNotes, err := GetParam(ParamPath("UpdaterNewReleaseNotes", false))
	check(err)

	updateFailedCount, err := GetParam(ParamPath("UpdateFailedCount", false))
	check(err)

	currentVersion, err := GetParam(ParamPath("UpdaterCurrentDescription", false))
	check(err)

	currentRepo, err := GetParam(ParamPath("GitRemote", false))
	check(err)

	currentBranch, err := GetParam(ParamPath("GitBranch", false))
	check(err)

	currentInstallDate, err := GetParam(ParamPath("InstallDate", false))
	check(err)

	currentReleaseNotes, err := GetParam(ParamPath("UpdaterCurrentReleaseNotes", false))
	check(err)

	currentCommit, err := GetParam(ParamPath("GitCommit", false))
	check(err)

	return Update{
		ButtonText:          buttonText,
		ButtonLink:          buttonLink,
		IdleButton:          Equal(state, []byte("idle")),
		InstallButton:       downloaded[0] == '1',
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
