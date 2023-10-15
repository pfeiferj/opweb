package main

import (
	_ "embed"
	"encoding/hex"
	"net/http"
	"os"
	"path/filepath"
	"sort"

	"github.com/cbroglie/mustache"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/flock"
)

var ParamsPath string = "/data/params/d"
var MemParamsPath string = "/dev/shm/params/d"

type Params struct {
	Title  string
	Links  []Links
	Params []Param
}

type Param struct {
	Key   string
	Value string
}

func routeParams(r *chi.Mux) {
	r.Get("/params", func(w http.ResponseWriter, r *http.Request) {
		params := generateParams()
		_, err := w.Write([]byte(params.Render()))
		check(err)
	})
}

func generateParams() Params {
	keys, err := GetParams(true)
	check(err)
	params := make([]Param, len(keys))
	for i, k := range keys {
		params[i].Key = k
		val, err := GetParam(ParamPath(k, true))
		check(err)
		if IsString(val) {
			params[i].Value = string(val)
		} else {
			params[i].Value = hex.EncodeToString(val)
		}
	}
	return Params{
		Title:  "Updater",
		Links:  PageLinks,
		Params: params,
	}
}

//go:embed templates/params.html
var paramsTemplate string

func (p Params) Render() string {
	page, err := mustache.RenderInLayoutPartials(paramsTemplate, basePartial, PartialProvider, p)
	check(err)
	return page
}

func EnsureParamDirectories() {
	os.MkdirAll(ParamsPath, 0775)
	os.MkdirAll(MemParamsPath, 0775)
}

func IsString(data []byte) bool {
	for _, b := range data {
		if (b < 32 || b > 126) && !(b == 9 || b == 13 || b == 10) {
			return false
		}
	}
	return true
}

func GetParams(isMem bool) ([]string, error) {
	basePath := ParamsPath
	if isMem {
		basePath = MemParamsPath
	}

	files, err := os.ReadDir(basePath)
	if err != nil {
		return nil, err
	}

	paramFiles := []string{}
	for _, file := range files {
		name := file.Name()
		if file.Type().IsRegular() && name[0] != '.' {
			paramFiles = append(paramFiles, name)
		}
	}
	sort.Strings(paramFiles)

	return paramFiles, nil
}

func HasMemParams() (bool, error) {
	files, err := os.ReadDir(BasePath)
	if err != nil {
		return false, err
	}
	return len(files) > 1, nil
}

func ParamPath(name string, isMem bool) string {
	var basePath string
	if isMem {
		basePath = MemParamsPath
	} else {
		basePath = ParamsPath
	}
	return filepath.Join(basePath, name)
}

func GetParam(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func PutParam(path string, data []byte) error {
	dir := filepath.Dir(path)
	file, err := os.CreateTemp(dir, ".tmp_value_"+filepath.Base(path))
	if err != nil {
		return err
	}
	tmpName := file.Name()
	defer os.Remove(tmpName)

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	err = file.Sync()
	if err != nil {
		return err
	}

	fileLock := flock.New(filepath.Join(dir, ".lock"))

	err = fileLock.Lock()
	if err != nil {
		return err
	}
	defer fileLock.Unlock()

	err = os.Rename(tmpName, path)
	if err != nil {
		return err
	}

	directory, err := os.Open(dir)
	if err != nil {
		return err
	}

	err = directory.Sync()
	if err != nil {
		return err
	}

	return nil
}

func RemoveParam(path string) error {
	dir := filepath.Dir(path)
	fileLock := flock.New(filepath.Join(dir, ".lock"))

	err := fileLock.Lock()
	if err != nil {
		return err
	}
	defer fileLock.Unlock()

	os.Remove(path)

	directory, err := os.Open(dir)
	if err != nil {
		return err
	}

	err = directory.Sync()
	if err != nil {
		return err
	}

	return nil
}
