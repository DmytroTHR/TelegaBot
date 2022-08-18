package app

import (
	"context"
	"fmt"

	"github.com/DmytroTHR/telegabot/pkg/helpers"
	"github.com/DmytroTHR/telegabot/pkg/telega"
)

var log = helpers.Logger()

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
	bot, err := telega.NewBot(app.host, app.token)
	if err != nil {
		return err
	}
	updates := bot.GetUpdates(context.TODO())
	for upd := range updates {
		if upd.Err != nil {
			return upd.Err
		}
		log.Println(upd.Data.Message)

		msg := upd.Data.Message
		user := msg.From

		textResponse := fmt.Sprintf(
			"<a href=\"tg://user?id=%d\"><b>%s</b></a> just posted:\n <span class=\"tg-spoiler\">%s</span>",
			user.ID, user.Username, msg.Text)
		sentMsg, err := bot.SendHTMLMessageToChat(msg.Chat.ID, textResponse, false)
		if err != nil {
			return err
		}
		log.Printf("%#v\n", sentMsg)
	}

	return nil
}
