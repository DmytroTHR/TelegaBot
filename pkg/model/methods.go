package model

import "fmt"

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

type ChatActioner interface {
	fmt.Stringer
	chatActionerPrivate()
}
type chatAction string

func (ca chatAction) chatActionerPrivate() {
}

func (ca chatAction) String() string {
	return string(ca)
}

var (
	ActionTyping          ChatActioner = chatAction("typing")
	ActionFindLocation    ChatActioner = chatAction("find_location")
	ActionChooseSticker   ChatActioner = chatAction("choose_sticker")
	ActionUploadPhoto     ChatActioner = chatAction("upload_photo")
	ActionUploadDocument  ChatActioner = chatAction("upload_document")
	ActionUploadVideo     ChatActioner = chatAction("upload_video")
	ActionUploadVideoNote ChatActioner = chatAction("upload_video_note")
	ActionUploadVoice     ChatActioner = chatAction("upload_voice")
)

type PollTyper interface {
	fmt.Stringer
	pollTyperPrivate()
}
type pollType string

func (ca pollType) pollTyperPrivate() {
}

func (ca pollType) String() string {
	return string(ca)
}

var (
	PollTypeQuiz    PollTyper = pollType("quiz")
	PollTypeRegular PollTyper = pollType("regular")
)
