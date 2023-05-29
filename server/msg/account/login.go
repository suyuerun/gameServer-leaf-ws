package account

import "server/conf/code"

type Login struct {
	MsgId     string
	IncId     string
	Username  string
	Password  string
	TokenText string
	Time      int
	Code      code.Code
}
