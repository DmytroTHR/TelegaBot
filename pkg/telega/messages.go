package telega

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DmytroTHR/telegabot/pkg/helpers"
	"github.com/DmytroTHR/telegabot/pkg/model"
	"github.com/pquerna/ffjson/ffjson"
)

const (
	updateTimeout = 10
	updateLimit   = 10
)
const (
	typeMessage    = "message"
	typeEditedChan = "edited_channel_post"
	typeCallback   = "callback_query"
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

func (b *Bot) GetUpdates(ctx context.Context) <-chan UpdateResponse {
	methodStr := fmt.Sprintf("method <%s>", model.MethodGetUpdates)
	response := make(chan UpdateResponse, updateLimit)
	bodyJSON := &struct {
		Offset         int      `json:"offset"`
		Limit          int      `json:"limit"`
		Timeout        int      `json:"timeout"`
		AllowedUpdates []string `json:"allowed_updates"`
	}{
		Offset:         0,
		Limit:          updateLimit,
		Timeout:        updateTimeout,
		AllowedUpdates: []string{typeMessage, typeEditedChan, typeCallback},
	}

	go func() {
		for {
			body, err := json.Marshal(bodyJSON)
			if err != nil {
				response <- NewUpdateResponse(nil,
					helpers.WrapError(methodStr, helpers.WrapError("marshalling update request", err)))
				continue
			}
			data, err := b.GetAPIResponse(model.MethodGetUpdates, http.MethodPost, string(body))
			if err != nil {
				response <- NewUpdateResponse(nil, helpers.WrapError(methodStr, err))
				continue
			}

			result := &model.ResponseUpdate{}
			err = ffjson.Unmarshal(data, result)
			if err != nil {
				response <- NewUpdateResponse(nil, helpers.WrapError(methodStr,
					helpers.WrapError("unmarshal result", err)))
				continue
			}
			if !result.OK {
				response <- NewUpdateResponse(nil, helpers.WrapError(methodStr,
					helpers.WrapError("false API request result", err)))
				continue
			}

			for _, upd := range result.Result {
				response <- NewUpdateResponse(upd, nil)
				bodyJSON.Offset = upd.UpdateID + 1
			}

			select {
			case <-ctx.Done():
				close(response)
				return
			default:
			}
		}
	}()

	return response
}

func (b *Bot) SendHTMLMessageToChat(chatID int, text string, silent bool) (*model.Message, error) {
	methodStr := fmt.Sprintf("method <%s>", model.MethodSendMessage)
	msgJSON := &struct {
		ChatID              int    `json:"chat_id"`
		Text                string `json:"text"`
		ParseMode           string `json:"parse_mode"`
		DisableNotification bool   `json:"disable_notification"`
	}{
		ChatID:              chatID,
		Text:                text,
		ParseMode:           "HTML",
		DisableNotification: silent,
	}
	body, err := json.Marshal(msgJSON)
	if err != nil {
		return nil, helpers.WrapError(methodStr, helpers.WrapError("marshalling update request", err))
	}
	data, err := b.GetAPIResponse(model.MethodSendMessage, http.MethodPost, string(body))
	if err != nil {
		return nil, helpers.WrapError(methodStr, err)
	}

	log.Warnln(string(data))
	result := &model.ResponseMessage{}
	err = ffjson.Unmarshal(data, result)
	if err != nil {
		return nil, helpers.WrapError(methodStr, helpers.WrapError("unmarshal result", err))
	}
	if !result.OK {
		return nil, helpers.WrapError(methodStr, helpers.WrapError("false API request result", err))
	}

	return result.Result, nil
}
