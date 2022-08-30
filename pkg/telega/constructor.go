package telega

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/DmytroTHR/telegabot/pkg/helpers"
	"github.com/DmytroTHR/telegabot/pkg/model"
)

type UpdateResponse struct {
	Data *model.Update
	Err  error
}

func NewUpdateResponse(data *model.Update, err error) UpdateResponse {
	return UpdateResponse{
		Data: data,
		Err:  err,
	}
}

func NewMessageSimple(chatID any, text string) (*model.SendMessageRequest, error) {
	id, err := ChatIDFrom(chatID)
	if err != nil {
		return nil, err
	}

	return &model.SendMessageRequest{
		ChatID:    id,
		Text:      text,
		ParseMode: "HTML",
	}, nil
}

func NewMessageReplySimple(chatID any, text string, replyMessageID int) (*model.SendMessageRequest, error) {
	simpleMessage, err := NewMessageSimple(chatID, text)
	if err != nil {
		return nil, err
	}
	simpleMessage.ReplyToMessageID = replyMessageID
	simpleMessage.AllowSendingWithoutReply = true

	return simpleMessage, nil
}

func NewLocationRequest(chatID any, latitude, longitude float64) (*model.SendLocationRequest, error) {
	id, err := ChatIDFrom(chatID)
	if err != nil {
		return nil, err
	}

	return &model.SendLocationRequest{
		ChatID:    id,
		Latitude:  latitude,
		Longitude: longitude,
	}, nil
}

func NewDiceRequest(chatID any, diceEmojiType uint) (*model.SendDiceRequest, error) {
	id, err := ChatIDFrom(chatID)
	if err != nil {
		return nil, err
	}

	var emoji string
	switch diceEmojiType {
	case model.DiceOneToSix:
		possibleEmojis := []string{`🎲`, `🎯`, `🎳`}
		emoji = possibleEmojis[helpers.RandomBetween(0, len(possibleEmojis)-1)]
	case model.DiceOneToFive:
		possibleEmojis := []string{`🏀`, `⚽`}
		emoji = possibleEmojis[helpers.RandomBetween(0, len(possibleEmojis)-1)]
	case model.DiceOneToSixtyFour:
		emoji = `🎰`
	default:
		return nil, helpers.Error("unknown dice type selected")
	}

	return &model.SendDiceRequest{
		ChatID: id,
		Emoji:  emoji,
	}, nil
}

func NewPollRequest(chatID any, question string, options []string) (*model.SendPollRequest, error) {
	id, err := ChatIDFrom(chatID)
	if err != nil {
		return nil, err
	}

	return &model.SendPollRequest{
		ChatID:   id,
		Question: question,
		Options:  options,
	}, nil
}

func NewPhotoRequest(chatID any, pathToFile string) (*model.SendPhotoRequest, error) {
	id, err := ChatIDFrom(chatID)
	if err != nil {
		return nil, err
	}

	return &model.SendPhotoRequest{
		ChatID: id,
		Photo:  pathToFile,
	}, nil
}

func NewPhotoRequestWithCaption(chatID any, pathToFile, caption string) (*model.SendPhotoRequest, error) {
	simpleRequest, err := NewPhotoRequest(chatID, pathToFile)
	if err != nil {
		return nil, err
	}
	simpleRequest.Caption = caption

	return simpleRequest, nil
}

func NewDocumentRequest(chatID any, pathToFile string) (*model.SendDocumentRequest, error) {
	id, err := ChatIDFrom(chatID)
	if err != nil {
		return nil, err
	}

	return &model.SendDocumentRequest{
		ChatID:   id,
		Document: pathToFile,
	}, nil
}

func NewDocumentRequestWithCaption(chatID any, pathToFile, caption string) (*model.SendDocumentRequest, error) {
	simpleRequest, err := NewDocumentRequest(chatID, pathToFile)
	if err != nil {
		return nil, err
	}
	simpleRequest.Caption = caption

	return simpleRequest, nil
}

func NewAnimationRequest(chatID any, pathToFile string) (*model.SendAnimationRequest, error) {
	id, err := ChatIDFrom(chatID)
	if err != nil {
		return nil, err
	}

	return &model.SendAnimationRequest{
		ChatID:    id,
		Animation: pathToFile,
	}, nil
}

func NewAnimationRequestWithCaption(chatID any, pathToFile, caption string) (*model.SendAnimationRequest, error) {
	simpleRequest, err := NewAnimationRequest(chatID, pathToFile)
	if err != nil {
		return nil, err
	}
	simpleRequest.Caption = caption

	return simpleRequest, nil
}

func NewChatActionRequest(chatID any, action model.ChatAction) (*model.SendChatActionRequest, error) {
	id, err := ChatIDFrom(chatID)
	if err != nil {
		return nil, err
	}

	return &model.SendChatActionRequest{
		ChatID: id,
		Action: action,
	}, nil
}

func NewBotCommandScope(scope model.CommandScope) *model.BotCommandScope {
	return &model.BotCommandScope{Type: scope}
}

func NewMyCommands(scope model.CommandScope) *model.MyCommands {
	return &model.MyCommands{Scope: NewBotCommandScope(scope)}
}

func NewSetMyCommands(commands map[string]string) (*model.SetMyCommands, error) {
	result := &model.SetMyCommands{}
	for com, des := range commands {
		botCom, err := NewBotCommand(com, des)
		if err != nil {
			return nil, helpers.WrapError("set commands", err)
		}
		result.Commands = append(result.Commands, botCom)
	}
	result.Scope = NewBotCommandScope(model.BotCommandScopeDefault)

	return result, nil
}

func NewBotCommand(command, description string) (*model.BotCommand, error) {
	if !botCommandVerified(command) {
		return nil, fmt.Errorf("%s cannot be a bot command (1-32 chars, lowercase, digits, _ )", command)
	}
	if len(description) == 0 || utf8.RuneCountInString(description) > 256 {
		return nil, fmt.Errorf("bot command description must be 1-256 chars long")
	}

	return &model.BotCommand{
		Command:     command,
		Description: description,
	}, nil
}

func botCommandVerified(command string) bool {
	allowedSymbols := "abcdefjhijklmnopqrstuvwxyz1234567890_"
	if len(command) == 0 || len(command) > 32 {
		return false
	}
	for _, sym := range command {
		if !strings.ContainsRune(allowedSymbols, sym) {
			return false
		}
	}

	return true
}
