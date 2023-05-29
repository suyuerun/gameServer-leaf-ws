package status

import (
	"server/conf/code"
)

type Player struct {
	Username  string
	TokenText string
}
type PlayerStatus struct {
	MsgId       string
	IncId       string
	TokenText   string
	PlayerNames []string
	Result      []*Player
	Time        int
	Code        code.Code
}
