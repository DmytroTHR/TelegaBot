package telega

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/DmytroTHR/telegabot/pkg/helpers"
	"github.com/DmytroTHR/telegabot/pkg/model"
	"github.com/pquerna/ffjson/ffjson"
)

func (b *Bot) GetFile(ctx context.Context, fileID string) (*model.File, error) {
	methodStr := fmt.Sprintf("method <%s>", model.MethodGetFile)
	body, err := ffjson.Marshal(struct {
		FileID string `json:"file_id"`
	}{FileID: fileID})
	if err != nil {
		return nil, helpers.WrapError(methodStr, err)
	}
	data, err := b.GetAPIResponse(ctx, model.MethodGetFile, http.MethodGet,
		bytes.NewReader(body), helpers.DefaultHeader())
	if err != nil {
		return nil, helpers.WrapError(methodStr, err)
	}

	response := &struct {
		OK     bool
		Result *model.File
	}{}
	err = ffjson.Unmarshal(data, response)
	if err != nil {
		return nil, helpers.WrapError(methodStr, helpers.WrapError("unmarshal result", err))
	}
	if !response.OK {
		return nil, helpers.WrapError(methodStr, helpers.Error("false API request result"))
	}

	return response.Result, nil
}

func (b *Bot) CompleteFileLink(linkPath string) (string, error) {
	url, err := b.Config.FullAPIFilePath(linkPath)
	if err != nil {
		return "", helpers.WrapError(fmt.Sprintf("complete link for %s", linkPath), err)
	}

	return url.String(), nil
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
