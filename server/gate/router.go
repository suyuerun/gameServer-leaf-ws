package gate

import (
	"server/chat"
	"server/game"
	"server/login"
	"server/msg"
	"server/msg/account"
	chat2 "server/msg/chat"
	game2 "server/msg/game"
	status2 "server/msg/status"
	"server/msg/test"
	"server/status"
)

func init() {
	msg.Processor.SetRouter(&test.Hello{}, game.ChanRPC)
	msg.Processor.SetRouter(&game2.Add{}, game.ChanRPC)
	msg.Processor.SetRouter(&chat2.Chat{}, chat.ChanRPC)
	msg.Processor.SetRouter(&status2.PlayerStatus{}, status.ChanRPC)
	msg.Processor.SetRouter(&account.Login{}, login.ChanRPC)
}
