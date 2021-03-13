package main

import (
	"log"

	"github.com/dogument/elenktis"
	"github.com/dogument/elenktis/internal/app"
)

func main() {
	cfg := &elenktis.Config{}
	elenktisApp := app.New()
	log.Println("Starting Elenktis app")
	log.Fatal(elenktisApp.Run(cfg))
}
