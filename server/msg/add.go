package msg

import "server/conf/code"

type Add struct {
	MsgId     string
	IncId     string
	TokenText string
	A         int
	B         int
	Result    int
	Time      int
	Code      code.Code
}
