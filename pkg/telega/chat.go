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

func ChatIDFrom(chatID any) (model.IntOrStr, error) {
	var id model.IntOrStr
	switch got := chatID.(type) {
	case int64:
		id = model.Int(got)
	case string:
		id = model.Str(got)
	default:
		return nil, helpers.Error("send message wrong ID type provided")
	}
	return id, nil
}

func (b *Bot) SendChatAction(ctx context.Context, actionRequest *model.SendChatActionRequest) (bool, error) {
	methodStr := fmt.Sprintf("method <%s>", model.MethodSendChatAction)

	body, err := ffjson.Marshal(actionRequest)
	if err != nil {
		return false, helpers.WrapError(methodStr, helpers.WrapError("marshal request", err))
	}

	response, err := b.GetAPIResponse(ctx, model.MethodSendChatAction, http.MethodPost,
		bytes.NewReader(body), helpers.DefaultHeader())
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
