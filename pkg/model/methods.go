package model

const (
	MethodGetMe      = "getMe"
	MethodGetUpdates = "getUpdates"
	MethodGetFile    = "getFile"

	MethodSendMessage      = "sendMessage"
	MethodSendPhoto        = "sendPhoto"
	MethodSendDocument     = "sendDocument"
	MethodSendAnimation    = "sendAnimation"
	MethodSendLocation     = "sendLocation"
	MethodSendDice         = "sendDice"
	MethodSendChatAction   = "sendChatAction"
	MethodSendPoll         = "sendPoll"
	MethodSetMyCommands    = "setMyCommands"
	MethodDeleteMyCommands = "deleteMyCommands"
	MethodGetMyCommands    = "getMyCommands"
)

const (
	DiceOneToSix = iota
	DiceOneToFive
	DiceOneToSixtyFour
)

type ChatAction interface {
	chatActionPrivate()
}

type chatAction string

func (ca chatAction) chatActionPrivate() {
}

var (
	ActionTyping          ChatAction = chatAction("typing")
	ActionFindLocation    ChatAction = chatAction("find_location")
	ActionChooseSticker   ChatAction = chatAction("choose_sticker")
	ActionUploadPhoto     ChatAction = chatAction("upload_photo")
	ActionUploadDocument  ChatAction = chatAction("upload_document")
	ActionUploadVideo     ChatAction = chatAction("upload_video")
	ActionUploadVideoNote ChatAction = chatAction("upload_video_note")
	ActionUploadVoice     ChatAction = chatAction("upload_voice")
)

type PollType interface {
	pollTypePrivate()
}
type pollType string

func (ca pollType) pollTypePrivate() {
}

var (
	PollTypeQuiz    PollType = pollType("quiz")
	PollTypeRegular PollType = pollType("regular")
)

type CommandScope interface {
	commandScopePrivate()
}
type commandScope string

func (cs commandScope) commandScopePrivate() {
}

var (
	BotCommandScopeDefault               CommandScope = commandScope("default")
	BotCommandScopeAllPrivateChats       CommandScope = commandScope("all_private_chats")
	BotCommandScopeAllGroupChats         CommandScope = commandScope("all_group_chats")
	BotCommandScopeAllChatAdministrators CommandScope = commandScope("all_chat_administrators")
	BotCommandScopeChat                  CommandScope = commandScope("chat")
	BotCommandScopeChatAdministrators    CommandScope = commandScope("chat_administrators")
	BotCommandScopeChatMember            CommandScope = commandScope("chat_member")
)
