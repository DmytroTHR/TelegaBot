package main

import (
	"os"

	"github.com/DmytroTHR/telegabot/internal/app"
	"github.com/DmytroTHR/telegabot/pkg/helpers"
)

var log = helpers.Logger()

func main() {
	token := os.Getenv("TELEGA_TOKEN")
	host := os.Getenv("TELEGA_HOST")
	application := app.NewApplication(host, token)
	err := application.Run()
	if err != nil {
		log.Panicln("Application crashed on -", err)
	}
}
