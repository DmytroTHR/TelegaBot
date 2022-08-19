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
	bot, err := telega.NewBot(app.host, app.token, 0, 0)
	if err != nil {
		return err
	}
	ctx := context.TODO()
	updates := bot.UpdateReceiver(ctx)
	for upd := range updates {
		if upd.Err != nil {
			return upd.Err
		}

		msg := upd.Data.Message
		user := msg.From

		textResponse := fmt.Sprintf(
			"<a href=\"tg://user?id=%d\"><b>%s</b></a> just posted:\n <span class=\"tg-spoiler\">%s</span>",
			user.ID, user.Username, msg.Text)
		msgJSON, err := telega.NewMessageReplySimple(msg.Chat.ID, textResponse, msg.MessageID)
		if err != nil {
			return err
		}
		_, err = bot.SendMessage(ctx, msgJSON)
		if err != nil {
			return err
		}

		pathToPhoto := "./testdata/photo.jpeg"
		msgPhoto, err := telega.NewPhotoRequestWithCaption(msg.Chat.ID, pathToPhoto, "Elephant")
		if err != nil {
			return err
		}
		_, err = bot.SendPhoto(ctx, msgPhoto)
		if err != nil {
			return err
		}
	}

	return nil
}
