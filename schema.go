package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/devsisters/go-dyncapnp"
)

func getSchemas() map[string]*dyncapnp.ParsedSchema {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	if strings.HasPrefix(exPath, "/tmp") {
		exPath, _ = os.Getwd()
	}

	logSchemaPath := filepath.Join(exPath, "../../cereal/log.capnp")
	carSchemaPath := filepath.Join(exPath, "../../cereal/car.capnp")
	legacySchemaPath := filepath.Join(exPath, "../../cereal/legacy.capnp")
	customSchemaPath := filepath.Join(exPath, "../../cereal/custom.capnp")
	cppSchemaPath := filepath.Join(exPath, "../../cereal/include/c++.capnp")

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
