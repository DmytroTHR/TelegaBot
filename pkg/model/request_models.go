//go:generate ffjson -force-regenerate $GOFILE
package model

type Int int64
type Str string

func (t Int) IntOrStr() {
}
func (t Str) IntOrStr() {
}

type IntOrStr interface {
	IntOrStr()
}

type SendMessageRequest struct {
	ChatID                   IntOrStr         `json:"chat_id"`
	Text                     string           `json:"text"`
	ParseMode                string           `json:"parse_mode,omitempty"`
	Entities                 []*MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview    bool             `json:"disable_web_page_preview,omitempty"`
	DisableNotification      bool             `json:"disable_notification,omitempty"`
	ProtectContent           bool             `json:"protect_content,omitempty"`
	ReplyToMessageID         int              `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              *ReplyMarkup     `json:"reply_markup,omitempty"`
}

type SendPhotoRequest struct {
	ChatID                   IntOrStr         `json:"chat_id"`
	Photo                    string           `json:"photo"`
	Caption                  string           `json:"caption,omitempty"`
	ParseMode                string           `json:"parse_mode,omitempty"`
	CaptionEntities          []*MessageEntity `json:"caption_entities,omitempty"`
	DisableNotification      bool             `json:"disable_notification,omitempty"`
	ProtectContent           bool             `json:"protect_content,omitempty"`
	ReplyToMessageID         int              `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              *ReplyMarkup     `json:"reply_markup,omitempty"`
}

type UpdateMessageRequest struct {
	Offset         int      `json:"offset,omitempty"`
	Limit          int      `json:"limit,omitempty"`
	Timeout        int      `json:"timeout,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

type ReplyMarkup interface {
	ReplyType() string
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}

func (kb *InlineKeyboardMarkup) ReplyType() string {
	return "InlineKeyboardMarkup"
}

type ReplyKeyboardMarkup struct {
	Keyboard              [][]*KeyboardButton `json:"keyboard"`
	ResizeKeyboard        bool                `json:"resize_keyboard,omitempty"`
	OneTimeKeyBoard       bool                `json:"one_time_key_board,omitempty"`
	InputFieldPlaceholder string              `json:"input_field_placeholder,omitempty"`
	Selective             bool                `json:"selective,omitempty"`
}

func (kb *ReplyKeyboardMarkup) ReplyType() string {
	return "ReplyKeyboardMarkup"
}

type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective,omitempty"`
}

func (kb *ReplyKeyboardRemove) ReplyType() string {
	return "ReplyKeyboardRemove"
}

type ForceReply struct {
	ForceReply            bool   `json:"force_reply"`
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`
	Selective             bool   `json:"selective,omitempty"`
}

func (kb *ForceReply) ReplyType() string {
	return "ForceReply"
}

type InlineKeyboardButton struct {
	Text                         string        `json:"text"`
	URL                          string        `json:"url,omitempty"`
	CallbackData                 string        `json:"callback_data,omitempty"`
	WebApp                       *WebAppInfo   `json:"web_app,omitempty"`
	LoginURL                     *LoginURL     `json:"login_url,omitempty"`
	SwitchInlineQuery            string        `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string        `json:"switch_inline_query_current_chat,omitempty"`
	CallbackGame                 *CallbackGame `json:"callback_game,omitempty"`
	Pay                          bool          `json:"pay,omitempty"`
}

type KeyboardButtonPollType struct {
	Type string `json:"type"`
}

type KeyboardButton struct {
	Text            string                  `json:"text"`
	RequestContact  bool                    `json:"request_contact,omitempty"`
	RequestLocation bool                    `json:"request_location,omitempty"`
	RequestPoll     *KeyboardButtonPollType `json:"request_poll,omitempty"`
	WebApp          *WebAppInfo             `json:"web_app,omitempty"`
}

type InlineQuery struct {
	ID       string    `json:"id"`
	From     *User     `json:"from"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
	ChatType string    `json:"chat_type,omitempty"`
	Location *Location `json:"location,omitempty"`
}
