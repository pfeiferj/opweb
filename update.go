package main

import (
	"net/http"
	"os"
	"os/exec"

	"github.com/go-chi/chi/v5"
	"pfeifer.dev/opwebd/components"
)

func routeUpdate(r *chi.Mux) {
	r.Get("/updater", func(w http.ResponseWriter, r *http.Request) {
		update := generateUpdate()
		routes := Routes()
		err := components.Layout("Updater", routes, components.Update(update.Update, update.Current, update.IdleButton, update.InstallButton)).Render(r.Context(), w)
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
		IdleButton: components.IdleButton{
			Show: Equal(state, []byte("idle")),
			Text: buttonText,
			Link: buttonLink,
		},
		InstallButton: len(downloaded) > 0 && downloaded[0] == '1',
		Update: components.UpdateData{
			State:        string(state),
			Version:      string(version),
			Branch:       string(updateBranch),
			LastCheck:    string(lastCheck),
			ReleaseNotes: string(updateNotes),
			FailureCount: string(updateFailedCount),
		},
		Current: components.CurrentData{
			Version:      string(currentVersion),
			Repo:         string(currentRepo),
			Branch:       string(currentBranch),
			InstallDate:  string(currentInstallDate),
			ReleaseNotes: string(currentReleaseNotes),
			Commit:       string(currentCommit),
		},
	}
}

type Update struct {
	IdleButton    components.IdleButton
	InstallButton bool
	Update        components.UpdateData
	Current       components.CurrentData
}
