package telega

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/DmytroTHR/telegabot/config"
	"github.com/DmytroTHR/telegabot/pkg/helpers"
	"github.com/DmytroTHR/telegabot/pkg/model"
	"github.com/pquerna/ffjson/ffjson"
)

var log = helpers.Logger()

type Bot struct {
	Config   *config.Config
	ID       int64
	Username string
}

func NewBot(host, token string, timeout, limit int) (*Bot, error) {
	conf, err := config.NewConfig(host, token, timeout, limit)
	if err != nil {
		return nil, helpers.WrapError("create Bot", err)
	}
	theBot := &Bot{
		Config: conf,
	}
	err = theBot.fillInfo()
	if err != nil {
		return nil, err
	}

	log.Println("New Bot connected:", theBot.ID)

	return theBot, nil
}

func (b *Bot) GetAPIResponse(ctx context.Context, methodCalled, httpMethod string,
	reqBody io.Reader, headers map[string]string) ([]byte, error) {
	addr, err := b.Config.FullAPIPath(methodCalled)
	if err != nil {
		return nil, helpers.WrapError("make API path", err)
	}

	req, err := http.NewRequestWithContext(ctx, httpMethod, addr.String(), reqBody)
	if err != nil {
		return nil, helpers.WrapError("prepare API request", err)
	}
	for key, val := range headers {
		req.Header.Add(key, val)
	}

	data, err := helpers.ReadFromRequest(req)
	log.Debugln("Made request on:", methodCalled)
	if err != nil {
		return nil, helpers.WrapError("execute API request", err)
	}

	return data, nil
}

func (b *Bot) fillInfo() error {
	methodStr := fmt.Sprintf("method <%s>", model.MethodGetMe)
	data, err := b.GetAPIResponse(context.Background(), model.MethodGetMe, http.MethodGet,
		nil, helpers.DefaultHeader())
	if err != nil {
		return helpers.WrapError(methodStr, err)
	}

	response := &model.ResponseUser{}
	err = ffjson.Unmarshal(data, response)
	if err != nil {
		return helpers.WrapError(methodStr, helpers.WrapError("unmarshal result", err))
	}
	if !response.OK {
		return helpers.WrapError(methodStr, helpers.Error("false API request result"))
	}
	b.ID = response.Result.ID
	b.Username = response.Result.Username

	return nil
}

func (b *Bot) GetFile(fileID string) (*model.File, error) {
	methodStr := fmt.Sprintf("method <%s>", model.MethodGetFile)
	body, err := ffjson.Marshal(struct {
		FileID string `json:"file_id"`
	}{FileID: fileID})
	if err != nil {
		return nil, helpers.WrapError(methodStr, err)
	}
	data, err := b.GetAPIResponse(context.Background(), model.MethodGetFile, http.MethodGet,
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
