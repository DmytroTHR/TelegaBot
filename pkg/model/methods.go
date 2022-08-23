package model

import "fmt"

const (
	MethodGetMe      = "getMe"
	MethodGetUpdates = "getUpdates"
	MethodGetFile    = "getFile"

	MethodSendMessage    = "sendMessage"
	MethodSendPhoto      = "sendPhoto"
	MethodSendDocument   = "sendDocument"
	MethodSendAnimation  = "sendAnimation"
	MethodSendLocation   = "sendLocation"
	MethodSendDice       = "sendDice"
	MethodSendChatAction = "sendChatAction"
	MethodSendPoll       = "sendPoll"
)

const (
	DiceOneToSix = iota
	DiceOneToFive
	DiceOneToSixtyFour
)

type ChatActioner interface {
	fmt.Stringer
	privateUnimplemented()
}
type chatAction string

func (ca chatAction) privateUnimplemented() {
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
