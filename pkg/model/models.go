package model

const (
	MethodGetMe      = "getMe"
	MethodGetUpdates = "getUpdates"
)

type User struct {
	ID                      int    `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name,omitempty"`
	Username                string `json:"username,omitempty"`
	LanguageCode            string `json:"language_code,omitempty"`
	IsPremium               bool   `json:"is_premium,omitempty"`
	AddedToAttachmentMenu   bool   `json:"added_to_attachment_menu,omitempty"`
	CanJoinGroups           bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages,omitempty"`
	SupportInlineQueries    bool   `json:"support_inline_queries,omitempty"`
}

type ResponseUser struct {
	OK     bool  `json:"ok"`
	Result *User `json:"result"`
}

type Update struct {
	UpdateID           int                 `json:"update_id"`
	Message            *Message            `json:"message,omitempty"`
	EditedMessage      *Message            `json:"edited_message,omitempty"`
	ChannelPost        *Message            `json:"channel_post,omitempty"`
	EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`
	InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`
	ShippingQuery      *ShippingQuery      `json:"shipping_query,omitempty"`
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query,omitempty"`
	Poll               *Poll               `json:"poll,omitempty"`
	PollAnswer         *PollAnswer         `json:"poll_answer,omitempty"`
	MyChatMember       *ChatMemberUpdated  `json:"my_chat_member,omitempty"`
	ChatMember         *ChatMemberUpdated  `json:"chat_member,omitempty"`
	ChatJoinRequest    *ChatJoinRequest    `json:"chat_join_request,omitempty"`
}

type Message struct {
	MessageID                     int                            `json:"message_id"`
	From                          *User                          `json:"from,omitempty"`
	SenderChat                    *Chat                          `json:"sender_chat,omitempty"`
	Date                          int                            `json:"date"`
	Chat                          *Chat                          `json:"chat"`
	ForwardFrom                   *User                          `json:"forward_from,omitempty"`
	ForwardFromChat               *Chat                          `json:"forward_from_chat,omitempty"`
	ForwardFromMessageID          int                            `json:"forward_from_message_id,omitempty"`
	ForwardSignature              string                         `json:"forward_signature,omitempty"`
	ForwardSenderName             string                         `json:"forward_sender_name,omitempty"`
	ForwardDate                   int                            `json:"forward_date,omitempty"`
	IsAutomaticForward            bool                           `json:"is_automatic_forward,omitempty"`
	ReplyToMessage                *Message                       `json:"reply_to_message,omitempty"`
	ViaBot                        *User                          `json:"via_bot,omitempty"`
	EditDate                      int                            `json:"edit_date,omitempty"`
	HasProtectedContent           bool                           `json:"has_protected_content,omitempty"`
	MediaGroupID                  string                         `json:"media_group_id,omitempty"`
	AuthorSignature               string                         `json:"author_signature,omitempty"`
	Text                          string                         `json:"text,omitempty"`
	Entities                      []*MessageEntity               `json:"entities,omitempty"`
	Animation                     *Animation                     `json:"animation,omitempty"`
	Audio                         *Audio                         `json:"audio,omitempty"`
	Document                      *Document                      `json:"document,omitempty"`
	Photo                         []*PhotoSize                   `json:"photo,omitempty"`
	Sticker                       *Sticker                       `json:"sticker,omitempty"`
	Video                         *Video                         `json:"video,omitempty"`
	VideoNote                     *VideoNote                     `json:"video_note,omitempty"`
	Voice                         *Voice                         `json:"voice,omitempty"`
	Caption                       string                         `json:"caption,omitempty"`
	CaptionEntities               []*MessageEntity               `json:"caption_entities,omitempty"`
	Contact                       *Contact                       `json:"contact,omitempty"`
	Dice                          *Dice                          `json:"dice,omitempty"`
	Game                          *Game                          `json:"game,omitempty"`
	Poll                          *Poll                          `json:"poll,omitempty"`
	Venue                         *Venue                         `json:"venue,omitempty"`
	Location                      *Location                      `json:"location,omitempty"`
	NewChatMembers                []*User                        `json:"new_chat_members,omitempty"`
	LeftChatMember                *User                          `json:"left_chat_member,omitempty"`
	NewChatTitle                  string                         `json:"new_chat_title,omitempty"`
	NewChatPhoto                  []*PhotoSize                   `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto               bool                           `json:"delete_chat_photo,omitempty"`
	GroupChatCreated              bool                           `json:"group_chat_created,omitempty"`
	SupergroupChatCreated         bool                           `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated            bool                           `json:"channel_chat_created,omitempty"`
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	MigrateToChatID               int                            `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatID             int                            `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage                 *Message                       `json:"pinned_message,omitempty"`
	Invoice                       *Invoice                       `json:"invoice,omitempty"`
	SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment,omitempty"`
	ConnectedWebsite              string                         `json:"connected_website,omitempty"`
	PassportData                  *PassportData                  `json:"passport_data,omitempty"`
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered,omitempty"`
	VideoChatScheduled            *VideoChatScheduled            `json:"video_chat_scheduled,omitempty"`
	VideoChatStarted              *VideoChatStarted              `json:"video_chat_started,omitempty"`
	VideoChatEnded                *VideoChatEnded                `json:"video_chat_ended,omitempty"`
	VideoChatParticipantsInvited  *VideoChatParticipantsInvited  `json:"video_chat_participants_invited,omitempty"`
	WebAppData                    *WebAppData                    `json:"web_app_data,omitempty"`
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup,omitempty"`
}

type MessageEntity struct {
	Type          string `json:"type"`
	Offset        int    `json:"offset"`
	Length        int    `json:"length"`
	URL           string `json:"url,omitempty"`
	User          *User  `json:"user,omitempty"`
	Language      string `json:"language,omitempty"`
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
}

type Chat struct {
	ID                                 int              `json:"id"`
	Type                               string           `json:"type"`
	Title                              string           `json:"title,omitempty"`
	Username                           string           `json:"username,omitempty"`
	FirstName                          string           `json:"first_name,omitempty"`
	LastName                           string           `json:"last_name,omitempty"`
	Photo                              *ChatPhoto       `json:"photo,omitempty"`
	Bio                                string           `json:"bio,omitempty"`
	HasPrivateForwards                 bool             `json:"has_private_forwards,omitempty"`
	HasRestrictedVoiceAndVideoMessages bool             `json:"has_restricted_voice_and_video_messages,omitempty"`
	JoinToSendMessages                 bool             `json:"join_to_send_messages,omitempty"`
	JoinByRequest                      bool             `json:"join_by_request,omitempty"`
	Description                        string           `json:"description,omitempty"`
	InviteLink                         string           `json:"invite_link,omitempty"`
	PinnedMessage                      *Message         `json:"pinned_message,omitempty"`
	Permissions                        *ChatPermissions `json:"permissions,omitempty"`
	SlowModeDelay                      int              `json:"slow_mode_delay,omitempty"`
	MessageAutoDeleteTime              int              `json:"message_auto_delete_time,omitempty"`
	HasProtectedContent                bool             `json:"has_protected_content,omitempty"`
	StickerSetName                     string           `json:"sticker_set_name,omitempty"`
	CanSetStickerSet                   bool             `json:"can_set_sticker_set,omitempty"`
	LinkedChatID                       int              `json:"linked_chat_id,omitempty"`
	Location                           *ChatLocation    `json:"location,omitempty"`
}

//TODO - fill structs below from https://core.telegram.org/bots/api#available-types

type InlineQuery struct {
}
type ChosenInlineResult struct {
}
type CallbackQuery struct {
}
type ShippingQuery struct {
}
type PreCheckoutQuery struct {
}
type Poll struct {
}
type PollAnswer struct {
}
type ChatPhoto struct {
}
type ChatPermissions struct {
}
type ChatLocation struct {
}
type ChatMemberUpdated struct {
}
type ChatJoinRequest struct {
}
type Animation struct {
}
type Audio struct {
}
type Document struct {
}
type PhotoSize struct {
}
type Sticker struct {
}
type Video struct {
}
type VideoNote struct {
}
type Voice struct {
}
type Contact struct {
}
type Dice struct {
}
type Game struct {
}
type Venue struct {
}
type Location struct {
}
type MessageAutoDeleteTimerChanged struct {
}
type Invoice struct {
}
type SuccessfulPayment struct {
}
type PassportData struct {
}
type ProximityAlertTriggered struct {
}
type VideoChatScheduled struct {
}
type VideoChatStarted struct {
}
type VideoChatEnded struct {
}
type VideoChatParticipantsInvited struct {
}
type WebAppData struct {
}
type InlineKeyboardMarkup struct {
}
