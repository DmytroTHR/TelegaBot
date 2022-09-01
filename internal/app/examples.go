package app

import (
	"context"
	"fmt"
	"time"

	"github.com/DmytroTHR/telegabot/pkg/helpers"
	"github.com/DmytroTHR/telegabot/pkg/model"
	"github.com/DmytroTHR/telegabot/pkg/telega"
)

func SendMessageExample(ctx context.Context, bot *telega.Bot, msg *model.Message) error {
	textResponse := fmt.Sprintf("Hello\t<b>Telega</b>!\n<i>Whats up))</i>")

	msgJSON, err := telega.NewMessageSimple(msg.Chat.ID, textResponse)
	if err != nil {
		return err
	}
	_, err = bot.SendMessage(ctx, msgJSON)

	return err
}

func SendReplyExample(ctx context.Context, bot *telega.Bot, msg *model.Message) error {
	user := msg.From
	textResponse := fmt.Sprintf(
		"<a href=\"tg://user?id=%d\"><b>%s</b></a> just posted:\n <i>%s</i>",
		user.ID, user.Username, msg.Text)

	msgJSON, err := telega.NewMessageReplySimple(msg.Chat.ID, textResponse, msg.MessageID)
	if err != nil {
		return err
	}
	_, err = bot.SendMessage(ctx, msgJSON)

	return err
}

func SendPhotoExample(ctx context.Context, bot *telega.Bot, msg *model.Message) error {
	pathToPhoto := "./testdata/photo.jpeg"

	msgPhoto, err := telega.NewPhotoRequestWithCaption(msg.Chat.ID, pathToPhoto, "Elephant")
	if err != nil {
		return err
	}
	_, err = bot.SendPhoto(ctx, msgPhoto)

	return err
}

func SendDocumentExample(ctx context.Context, bot *telega.Bot, msg *model.Message) error {
	pathToDoc := "./testdata/constitution.pdf"

	msgDoc, err := telega.NewDocumentRequestWithCaption(msg.Chat.ID, pathToDoc, "Constitution")
	if err != nil {
		return err
	}
	_, err = bot.SendDocument(ctx, msgDoc)

	return err
}

func SendAnimationExample(ctx context.Context, bot *telega.Bot, msg *model.Message) error {
	pathToAnim := "./testdata/ASAP.gif"

	msgAnim, err := telega.NewAnimationRequestWithCaption(msg.Chat.ID, pathToAnim, "Do it quickly")
	if err != nil {
		return err
	}
	_, err = bot.SendAnimation(ctx, msgAnim)

	return err
}

func SendLocationExample(ctx context.Context, bot *telega.Bot, msg *model.Message) error {
	msgLoc, err := telega.NewLocationRequest(msg.Chat.ID, 46.45, 30.77)
	if err != nil {
		return err
	}
	_, err = bot.SendLocation(ctx, msgLoc)

	return err
}

func SendContactExample(ctx context.Context, bot *telega.Bot, msg *model.Message) error {
	msgContact, err := telega.NewContactRequest(msg.Chat.ID, "+123456789", "Volan", "Demort")
	if err != nil {
		return err
	}
	_, err = bot.SendContact(ctx, msgContact)

	return err
}

func SendChatActionExample(ctx context.Context, bot *telega.Bot, msg *model.Message) error {
	msgAct, err := telega.NewChatActionRequest(msg.Chat.ID, model.ActionTyping)
	if err != nil {
		return err
	}
	_, err = bot.SendChatAction(ctx, msgAct)
	if err != nil {
		return err
	}
	time.Sleep(3 * time.Second)

	return nil
}

func SendDiceExample(ctx context.Context, bot *telega.Bot, msg *model.Message) error {
	msgDice, err := telega.NewDiceRequest(msg.Chat.ID, uint(helpers.RandomBetween(0, 2)))
	//can be one of - model.DiceOneToSixtyFour (2), model.DiceOneToFive (1), model.DiceOneToSix (0)
	if err != nil {
		return err
	}
	got, err := bot.SendDice(ctx, msgDice)
	if err != nil {
		return err
	}

	//GET BACK DICE RESULT
	response := fmt.Sprintf(`Your result is <span class="tg-spoiler">%d</span> points`, got.Dice.Value)
	msgResp, err := telega.NewMessageSimple(msg.Chat.ID, response)
	if err != nil {
		return err
	}
	_, err = bot.SendMessage(ctx, msgResp)

	return err
}

func SendPollExample(ctx context.Context, bot *telega.Bot, msg *model.Message) error {
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

	return err
}

func GetFileExample(ctx context.Context, bot *telega.Bot, msg *model.Message) error {
	if msg.Document != nil {
		gotFile, err := bot.GetFile(ctx, msg.Document.FileID)
		if err != nil {
			return err
		}

		pathToSave := "./testdata/" + gotFile.FilePath
		linkToDownload, err := bot.CompleteFileLink(gotFile.FilePath)
		if err != nil {
			return err
		}

		downlFile, err := helpers.DownloadFile(ctx, linkToDownload, pathToSave)
		if err != nil {
			return err
		}
		log.Println(downlFile)
	}

	return nil
}

func WorkWithCommandsExample(ctx context.Context, bot *telega.Bot, msg *model.Message) error {
	commands := map[string]string{
		"hi":   "say hi",
		"bye":  "say bye-bye",
		"help": "show help",
	}
	//SET COMMANDS
	commandSetter, err := telega.NewSetMyCommands(commands)
	if err != nil {
		return err
	}
	_, err = bot.SetMyCommands(ctx, commandSetter)
	if err != nil {
		return err
	}
	log.Println("commands were set")

	//GET CURRENT COMMANDS
	comResult, err := bot.GetMyCommands(ctx, telega.NewMyCommands(model.BotCommandScopeDefault))
	if err != nil {
		return err
	}
	for _, v := range comResult {
		log.Println(*v)
	}

	//DELETE CURRENT COMMANDS
	_, err = bot.DeleteMyCommands(ctx, telega.NewMyCommands(model.BotCommandScopeDefault))
	if err != nil {
		return err
	}
	log.Println("commands were deleted")

	return nil
}
