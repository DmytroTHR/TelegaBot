package telega

import (
	"github.com/DmytroTHR/telegabot/pkg/helpers"
	"github.com/DmytroTHR/telegabot/pkg/model"
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
