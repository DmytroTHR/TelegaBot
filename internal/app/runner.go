package app

import (
	"context"

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

		////SEND MESSAGE
		//err = SendMessageExample(ctx, bot, msg)
		//if err != nil {
		//	return err
		//}
		//
		//SEND REPLY
		err = SendReplyExample(ctx, bot, msg)
		if err != nil {
			return err
		}
		//
		////SEND PHOTO
		//err = SendPhotoExample(ctx, bot, msg)
		//if err != nil {
		//	return err
		//}
		//
		////SEND DOCUMENT
		//err = SendDocumentExample(ctx, bot, msg)
		//if err != nil {
		//	return err
		//}
		//
		////SEND ANIMATION
		//err = SendAnimationExample(ctx, bot, msg)
		//if err != nil {
		//	return err
		//}
		//
		////SEND LOCATION
		//err = SendLocationExample(ctx, bot, msg)
		//if err != nil {
		//	return err
		//}
		//
		////SEND CONTACT
		//err = SendContactExample(ctx, bot, msg)
		//if err != nil {
		//	return err
		//}
		//
		////SEND DICE
		//err = SendDiceExample(ctx, bot, msg)
		//if err != nil {
		//	return err
		//}
		//
		////SEND CHAT ACTION
		//err = SendChatActionExample(ctx, bot, msg)
		//if err != nil {
		//	return err
		//}
		//
		////SEND POLL
		//err = SendPollExample(ctx, bot, msg)
		//if err != nil {
		//	return err
		//}
		//
		////GET & DOWNLOAD FILE
		//err = GetFileExample(ctx, bot, msg)
		//if err != nil {
		//	return err
		//}
		//
		////MANIPULATING WITH COMMANDS
		//err = WorkWithCommandsExample(ctx, bot, msg)
		//if err != nil {
		//	return err
		//}
	}

	return nil
}
