package telega

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/DmytroTHR/telegabot/pkg/helpers"
	"github.com/DmytroTHR/telegabot/pkg/model"
	"github.com/pquerna/ffjson/ffjson"
)

const (
	typeMessage    = "message"
	typeEditedChan = "edited_channel_post"
	typeCallback   = "callback_query"
)

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

	return b.messageResultFor(ctx, model.MethodSendMessage, bytes.NewReader(body), helpers.DefaultHeader())
}

func (b *Bot) SendPhoto(ctx context.Context, photoRequest *model.SendPhotoRequest) (*model.Message, error) {
	return b.sendData(ctx, photoRequest, model.MethodSendPhoto)
}

func (b *Bot) SendDocument(ctx context.Context, docRequest *model.SendDocumentRequest) (*model.Message, error) {
	return b.sendData(ctx, docRequest, model.MethodSendDocument)
}

func (b *Bot) SendAnimation(ctx context.Context, animRequest *model.SendAnimationRequest) (*model.Message, error) {
	return b.sendData(ctx, animRequest, model.MethodSendAnimation)
}

func (b *Bot) SendLocation(ctx context.Context, locRequest *model.SendLocationRequest) (*model.Message, error) {
	methodStr := fmt.Sprintf("method <%s>", model.MethodSendLocation)

	body, err := ffjson.Marshal(locRequest)
	if err != nil {
		return nil, helpers.WrapError(methodStr, helpers.WrapError("marshal request", err))
	}

	return b.messageResultFor(ctx, model.MethodSendLocation, bytes.NewReader(body), helpers.DefaultHeader())
}

func (b *Bot) SendDice(ctx context.Context, diceRequest *model.SendDiceRequest) (*model.Message, error) {
	methodStr := fmt.Sprintf("method <%s>", model.MethodSendDice)

	body, err := ffjson.Marshal(diceRequest)
	if err != nil {
		return nil, helpers.WrapError(methodStr, helpers.WrapError("marshal request", err))
	}

	return b.messageResultFor(ctx, model.MethodSendDice, bytes.NewReader(body), helpers.DefaultHeader())
}

func (b *Bot) sendData(ctx context.Context, request model.DataSender, method string) (*model.Message, error) {
	methodStr := fmt.Sprintf("method <%s>", method)

	body, err := ffjson.Marshal(request)
	if err != nil {
		return nil, helpers.WrapError(methodStr, helpers.WrapError("marshal request", err))
	}
	opts := helpers.UnmarshalToKeyValueString(body)
	if len(opts) == 0 {
		return nil, helpers.WrapError(methodStr, fmt.Errorf("unmarshal to map string"))
	}
	files := map[string]string{}

	switch method {
	case model.MethodSendDocument:
		delete(opts, "document")
		delete(opts, "thumb")
		request, ok := request.(*model.SendDocumentRequest)
		if !ok {
			return nil, helpers.WrapError(methodStr, fmt.Errorf(""))
		}
		files["document"] = request.Document
		files["thumb"] = request.Thumb
	case model.MethodSendPhoto:
		delete(opts, "photo")
		request, ok := request.(*model.SendPhotoRequest)
		if !ok {
			return nil, helpers.WrapError(methodStr, fmt.Errorf(""))
		}
		files["photo"] = request.Photo
	case model.MethodSendAnimation:
		delete(opts, "animation")
		request, ok := request.(*model.SendAnimationRequest)
		if !ok {
			return nil, helpers.WrapError(methodStr, fmt.Errorf(""))
		}
		files["animation"] = request.Animation
	default:
		return nil, helpers.WrapError(methodStr, fmt.Errorf("send data with the method"))
	}

	preparedData, contentType, err := helpers.MultipartDataUpload(files, opts)
	if err != nil {
		return nil, helpers.WrapError(methodStr, err)
	}

	header := map[string]string{
		"Content-Type": contentType,
	}

	return b.messageResultFor(ctx, method, preparedData, header)
}

func (b *Bot) messageResultFor(ctx context.Context, method string,
	data io.Reader, headers map[string]string) (*model.Message, error) {
	methodStr := fmt.Sprintf("method <%s>", method)

	response, err := b.GetAPIResponse(ctx, method, http.MethodPost, data, headers)
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
