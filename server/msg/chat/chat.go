package chat

import "server/conf/code"

type Chat struct {
	MsgId     string
	To        string
	IncId     string
	TokenText string
	Content   string
	Time      int
	Code      code.Code
}
