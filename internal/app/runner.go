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
			"<a href=\"tg://user?id=%d\"><b>%s</b></a> just posted:\n <i>%s</i>",
			user.ID, user.Username, msg.Text)
		msgJSON, err := telega.NewMessageReplySimple(msg.Chat.ID, textResponse, msg.MessageID)
		if err != nil {
			return err
		}
		_, err = bot.SendMessage(ctx, msgJSON)
		if err != nil {
			return err
		}

		////SEND PHOTO
		//pathToPhoto := "./testdata/photo.jpeg"
		//msgPhoto, err := telega.NewPhotoRequestWithCaption(msg.Chat.ID, pathToPhoto, "Elephant")
		//if err != nil {
		//	return err
		//}
		//_, err = bot.SendPhoto(ctx, msgPhoto)

		////SEND DOCUMENT
		//pathToDoc := "./testdata/constitution.pdf"
		//msgDoc, err := telega.NewDocumentRequestWithCaption(msg.Chat.ID, pathToDoc, "Constitution")
		//if err != nil {
		//	return err
		//}
		//_, err = bot.SendDocument(ctx, msgDoc)

		////SEND ANIMATION
		//pathToAnim := "./testdata/ASAP.gif"
		//msgAnim, err := telega.NewAnimationRequestWithCaption(msg.Chat.ID, pathToAnim, "Do it quickly")
		//if err != nil {
		//	return err
		//}
		//_, err = bot.SendAnimation(ctx, msgAnim)

		////SEND LOCATION
		//msgLoc, err := telega.NewLocationRequest(msg.Chat.ID, 15.0, 43.34)
		//if err != nil {
		//	return err
		//}
		//_, err = bot.SendLocation(ctx, msgLoc)

		////SEND DICE
		//msgDice, err := telega.NewDiceRequest(msg.Chat.ID, uint(helpers.RandomBetween(0, 2)))
		////can be one of - model.DiceOneToSixtyFour (2), model.DiceOneToFive (1), model.DiceOneToSix (0)
		//if err != nil {
		//	return err
		//}
		//got, err := bot.SendDice(ctx, msgDice)
		//if err != nil {
		//	return err
		//}
		////SEND TYPING ACTION TO CHAT
		//msgAct, err := telega.NewChatActionRequest(msg.Chat.ID, model.ActionTyping)
		//if err != nil {
		//	return err
		//}
		//_, err = bot.SendChatAction(ctx, msgAct)
		//if err != nil {
		//	return err
		//}
		//time.Sleep(3 * time.Second)
		////GET BACK DICE RESULT
		//response := fmt.Sprintf(`Your result is <span class="tg-spoiler">%d</span> points`, got.Dice.Value)
		//msgResp, err := telega.NewMessageSimple(msg.Chat.ID, response)
		//if err != nil {
		//	return err
		//}
		//_, err = bot.SendMessage(ctx, msgResp)

		////GET & DOWNLOAD FILE
		//if msg.Document != nil {
		//	msgAct, _ := telega.NewChatActionRequest(msg.Chat.ID, model.ActionUploadDocument)
		//	bot.SendChatAction(ctx, msgAct)
		//
		//	gotFile, err := bot.GetFile(ctx, msg.Document.FileID)
		//	if err != nil {
		//		return err
		//	}
		//
		//	pathToSave := "./testdata/" + gotFile.FilePath
		//	linkToDownload, err := bot.CompleteFileLink(gotFile.FilePath)
		//	if err != nil {
		//		return err
		//	}
		//	downlFile, err := helpers.DownloadFile(ctx, linkToDownload, pathToSave)
		//	if err != nil {
		//		return err
		//	}
		//	log.Println(downlFile)
		//}

		//SEND POLL
		question := "What color do you like the most?"
		answers := []string{
			"🍓 red",
			"🍏 green",
			"🫐 blue",
		}
		msgPoll, err := telega.NewPollRequest(msg.Chat.ID, question, answers)
		if err != nil {
			return err
		}
		_, err = bot.SendPoll(ctx, msgPoll)

		if err != nil {
			return err
		}
	}

	return nil
}
