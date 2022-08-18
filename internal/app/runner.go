package app

import (
	"github.com/DmytroTHR/telegabot/pkg/telega"
)

type Application struct {
	host, token string
}

func NewApplication(host, token string) *Application {
	return &Application{
		host:  host,
		token: token,
	}
}

func (app *Application) Run() error {
	_, err := telega.NewBot(app.host, app.token)

	return err
}
