package main

import (
	"log"

	"github.com/mjibson/esc/embed"
)

func main() {
	conf := &embed.Config{
		ModTime:    "0",
		OutputFile: "pkg/js/static.go",
		Package:    "js",
		Prefix:     "pkg/js",
		Private:    true,
		Files:      []string{`pkg/js/helpers.js`},
	}
	embed.Run(conf)

	if err := generateFeatureMatrix(); err != nil {
		log.Fatal(err)
	}
}
