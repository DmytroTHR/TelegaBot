package telega

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/DmytroTHR/telegabot/pkg/helpers"
	"github.com/DmytroTHR/telegabot/pkg/model"
	"github.com/pquerna/ffjson/ffjson"
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

func (b *Bot) UpdateReceiver(ctx context.Context) <-chan UpdateResponse {
	methodStr := fmt.Sprintf("method <%s>", model.MethodGetUpdates)
	answerCh := make(chan UpdateResponse, b.Config.UpdateMsgLimit)
	bodyJSON := &model.UpdateMessageRequest{
		Offset:         0,
		Limit:          b.Config.UpdateMsgLimit,
		Timeout:        b.Config.UpdateTimeout,
		AllowedUpdates: []string{typeMessage, typeEditedChan, typeCallback},
	}

	go func() {
		for {
			body, err := ffjson.Marshal(bodyJSON)
			if err != nil {
				answerCh <- NewUpdateResponse(nil,
					helpers.WrapError(methodStr, helpers.WrapError("marshal request", err)))
				continue
			}
			response, err := b.GetAPIResponse(ctx, model.MethodGetUpdates, http.MethodPost,
				bytes.NewReader(body), helpers.DefaultHeader())
			if err != nil {
				answerCh <- NewUpdateResponse(nil, helpers.WrapError(methodStr, err))
				continue
			}

			result := &model.ResponseUpdate{}
			err = ffjson.Unmarshal(response, result)
			if err != nil {
				answerCh <- NewUpdateResponse(nil, helpers.WrapError(methodStr,
					helpers.WrapError("unmarshal result", err)))
				continue
			}
			if !result.OK {
				answerCh <- NewUpdateResponse(nil, helpers.WrapError(methodStr,
					helpers.Error(fmt.Sprintf("request API result: %s", string(response)))))
				continue
			}

			for _, upd := range result.Result {
				answerCh <- NewUpdateResponse(upd, nil)
				bodyJSON.Offset = upd.UpdateID + 1
			}

			select {
			case <-ctx.Done():
				close(answerCh)
				log.Println("Stop receiving updates for:", b.ID)
				return
			default:
			}
		}
	}()

	return answerCh
}

func (b *Bot) SendMessage(ctx context.Context, message *model.SendMessageRequest) (*model.Message, error) {
	methodStr := fmt.Sprintf("method <%s>", model.MethodSendMessage)

	body, err := ffjson.Marshal(message)
	if err != nil {
		return nil, helpers.WrapError(methodStr, helpers.WrapError("marshal request", err))
	}
	response, err := b.GetAPIResponse(ctx, model.MethodSendMessage, http.MethodPost,
		bytes.NewReader(body), helpers.DefaultHeader())
	if err != nil {
		return nil, helpers.WrapError(methodStr, err)
	}

	result := &model.ResponseMessage{}
	err = ffjson.Unmarshal(response, result)
	if err != nil {
		return nil, helpers.WrapError(methodStr, helpers.WrapError("unmarshal result", err))
	}
	if !result.OK {
		return nil, helpers.WrapError(methodStr,
			helpers.Error(fmt.Sprintf("request API result: %s", string(response))))
	}

	return result.Result, nil
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

func (b *Bot) SendPhoto(ctx context.Context, photoRequest *model.SendPhotoRequest) (*model.Message, error) {
	methodStr := fmt.Sprintf("method <%s>", model.MethodSendPhoto)

	body, err := ffjson.Marshal(photoRequest)
	if err != nil {
		return nil, helpers.WrapError(methodStr, helpers.WrapError("marshal request", err))
	}
	conv := map[string]json.RawMessage{}
	err = json.Unmarshal(body, &conv)
	if err != nil {
		return nil, helpers.WrapError(methodStr, helpers.WrapError("unmarshal request to options", err))
	}
	opts := map[string]string{}
	for k := range conv {
		opts[k] = string(conv[k])
	}
	delete(opts, "photo")

	preparedData, contentType, err := prepareFileToUpload(photoRequest.Photo, opts)
	if err != nil {
		return nil, helpers.WrapError(methodStr, err)
	}

	header := map[string]string{
		"Content-Type": contentType,
	}
	response, err := b.GetAPIResponse(ctx, model.MethodSendPhoto, http.MethodPost, preparedData, header)
	if err != nil {
		return nil, helpers.WrapError(methodStr, err)
	}
	result := &model.ResponseMessage{}
	err = ffjson.Unmarshal(response, result)
	if err != nil {
		return nil, helpers.WrapError(methodStr, helpers.WrapError("unmarshal result", err))
	}
	if !result.OK {
		return nil, helpers.WrapError(methodStr,
			helpers.Error(fmt.Sprintf("request API result: %s", string(response))))
	}

	return result.Result, nil
}

func prepareFileToUpload(filePath string, options map[string]string) (*bytes.Buffer, string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, "", helpers.WrapError("open file for upload", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for key, val := range options {
		err = writer.WriteField(key, val)
		if err != nil {
			return nil, "", helpers.WrapError("set mime fields", err)
		}
	}
	part, err := writer.CreateFormFile("photo", filePath)
	if err != nil {
		return nil, "", helpers.WrapError("create form file", err)
	}
	n, err := io.Copy(part, file)
	writer.Close()
	if err != nil {
		return nil, "", helpers.WrapError("copy file to part", err)
	}
	if n == 0 {
		return nil, "", helpers.Error("no data copied from file")
	}

	return body, writer.FormDataContentType(), nil
}
