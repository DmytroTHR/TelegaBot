package telega

import (
	"fmt"
	"net/http"

	"github.com/DmytroTHR/telegabot/config"
	"github.com/DmytroTHR/telegabot/pkg/helpers"
	"github.com/DmytroTHR/telegabot/pkg/model"
	"github.com/pquerna/ffjson/ffjson"
)

var log = helpers.Logger()

type Bot struct {
	Config   *config.Config
	ID       int
	Username string
}

func NewBot(host, token string) (*Bot, error) {
	conf, err := config.NewConfig(host, token)
	if err != nil {
		return nil, helpers.WrapError("create Bot", err)
	}
	theBot := &Bot{Config: conf}
	err = theBot.fillInfo()
	if err != nil {
		return nil, err
	}

	log.Println("New Bot connected:", theBot.ID)

	return theBot, nil
}

func (b *Bot) GetAPIResponse(methodCalled, httpMethod, reqBody string) ([]byte, error) {
	addr, err := b.Config.FullAPIPath(methodCalled)
	if err != nil {
		return nil, helpers.WrapError("make API path", err)
	}
	data, err := helpers.ReadFromRequest(addr, httpMethod, reqBody)
	if err != nil {
		return nil, helpers.WrapError("execute API request", err)
	}

	return data, nil
}

func (b *Bot) fillInfo() error {
	methodStr := fmt.Sprintf("method <%s>", model.MethodGetMe)
	data, err := b.GetAPIResponse(model.MethodGetMe, http.MethodGet, "")
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
