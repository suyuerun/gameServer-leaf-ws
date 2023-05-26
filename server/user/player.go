package user

import "github.com/name5566/leaf/gate"

type Player struct {
	Username  string
	Password  string
	TokenText string
	Agent     gate.Agent
}
