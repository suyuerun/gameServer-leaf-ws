package msg

import "server/conf/code"

type Hello struct {
	MsgId string
	IncId string
	Name  string
	Text  string
	Time  int
	Code  code.Code
}
