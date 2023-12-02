package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/devsisters/go-dyncapnp"
)

func getCerealPath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	exPath := filepath.Dir(ex)
	if strings.HasPrefix(exPath, "/tmp") {
		exPath, _ = os.Getwd()
	}

	path := "/data/openpilot/cereal"
	exists, _ := Exists(path)
	if exists {
		return path, nil
	}

	path = filepath.Join(exPath, "cereal")
	exists, _ = Exists(path)
	if exists {
		return path, nil
	}

	path = filepath.Join(exPath, "../cereal")
	exists, _ = Exists(path)
	if exists {
		return path, nil
	}

	path = filepath.Join(exPath, "../../cereal")
	exists, _ = Exists(path)
	if exists {
		return path, nil
	}

	path = filepath.Join(exPath, "../openpilot/cereal")
	exists, _ = Exists(path)
	if exists {
		return path, nil
	}

	return "", errors.New("Could not find cereal")
}

func getSchemas() map[string]*dyncapnp.ParsedSchema {
	cereal_path, err := getCerealPath()
	check(err)

	logSchemaPath := filepath.Join(cereal_path, "log.capnp")
	carSchemaPath := filepath.Join(cereal_path, "car.capnp")
	legacySchemaPath := filepath.Join(cereal_path, "legacy.capnp")
	customSchemaPath := filepath.Join(cereal_path, "custom.capnp")
	cppSchemaPath := filepath.Join(cereal_path, "include/c++.capnp")

	logDat, err := os.ReadFile(logSchemaPath)
	check(err)

	carDat, err := os.ReadFile(carSchemaPath)
	check(err)

	legacyDat, err := os.ReadFile(legacySchemaPath)
	check(err)

	customDat, err := os.ReadFile(customSchemaPath)
	check(err)

	cppDat, err := os.ReadFile(cppSchemaPath)
	check(err)

	schemaFiles := map[string][]byte{
		"log.capnp":         logDat,
		"car.capnp":         carDat,
		"legacy.capnp":      legacyDat,
		"custom.capnp":      customDat,
		"include/c++.capnp": cppDat,
	}

	schemaImports := map[string][]byte{}

	schemaPaths := []string{"log.capnp", "car.capnp", "legacy.capnp", "custom.capnp"}

	schemas, err := dyncapnp.ParseFromFiles(schemaFiles, schemaImports, schemaPaths)
	check(err)

	return schemas
}
