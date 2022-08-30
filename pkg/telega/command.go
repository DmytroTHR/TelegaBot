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

func (b *Bot) SetMyCommands(ctx context.Context, commandList *model.SetMyCommands) (bool, error) {
	return b.doWithCommands(ctx, model.MethodSetMyCommands, commandList)
}

func (b *Bot) DeleteMyCommands(ctx context.Context, commandParams *model.MyCommands) (bool, error) {
	return b.doWithCommands(ctx, model.MethodDeleteMyCommands, commandParams)
}

func (b *Bot) doWithCommands(ctx context.Context, method string, params any) (bool, error) {
	var commandParams any
	switch params := params.(type) {
	case *model.MyCommands:
		commandParams = params
	case *model.SetMyCommands:
		commandParams = params
	default:
		return false, helpers.Error(
			fmt.Sprintf("wrong data type to work with commands: %v of type %T", params, params))
	}

	methodStr := fmt.Sprintf("method <%s>", method)

	body, err := ffjson.Marshal(commandParams)
	if err != nil {
		return false, helpers.WrapError(methodStr, helpers.WrapError("marshal request", err))
	}

	response, err := b.GetAPIResponse(ctx, method, http.MethodPost, bytes.NewReader(body), helpers.DefaultHeader())
	if err != nil {
		return false, helpers.WrapError(methodStr, err)
	}
	result := &struct {
		OK     bool
		Result bool
	}{}
	err = ffjson.Unmarshal(response, result)
	if err != nil {
		return false, helpers.WrapError(methodStr, helpers.WrapError("unmarshal result", err))
	}
	if !result.OK {
		return false, helpers.WrapError(methodStr,
			helpers.Error(fmt.Sprintf("request API result: %s", string(response))))
	}

	return result.Result, nil
}

func (b *Bot) GetMyCommands(ctx context.Context, commandParams *model.MyCommands) ([]*model.BotCommand, error) {
	methodStr := fmt.Sprintf("method <%s>", model.MethodGetMyCommands)

	body, err := ffjson.Marshal(commandParams)
	if err != nil {
		return nil, helpers.WrapError(methodStr, helpers.WrapError("marshal request", err))
	}

	response, err := b.GetAPIResponse(ctx, model.MethodGetMyCommands, http.MethodPost,
		bytes.NewReader(body), helpers.DefaultHeader())
	if err != nil {
		return nil, helpers.WrapError(methodStr, err)
	}
	result := &struct {
		OK     bool
		Result []*model.BotCommand
	}{}
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
