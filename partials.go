package main

import (
	"embed"
)

//go:embed partials/*
var partials embed.FS

type Partials struct{}

func (_ Partials) Get(p string) (string, error) {
	data, err := partials.ReadFile("partials/" + p)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

var PartialProvider = Partials{}
