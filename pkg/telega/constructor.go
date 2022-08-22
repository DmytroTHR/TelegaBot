package telega

import (
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

func NewChatActionRequest(chatID any, action model.ChatActioner) (*model.SendChatActionRequest, error) {
	id, err := ChatIDFrom(chatID)
	if err != nil {
		return nil, err
	}

	return &model.SendChatActionRequest{
		ChatID: id,
		Action: action,
	}, nil
}
