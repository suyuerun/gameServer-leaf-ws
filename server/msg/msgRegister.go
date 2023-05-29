package msg

import (
	"github.com/name5566/leaf/network/json"
	"server/msg/account"
	"server/msg/chat"
	"server/msg/game"
	"server/msg/status"
	"server/msg/test"
)

var Processor = json.NewProcessor()

func init() {
	Processor.Register(&test.Hello{})
	Processor.Register(&account.Login{})
	Processor.Register(&chat.Broadcast{})
	Processor.Register(&chat.Chat{})
	Processor.Register(&game.Add{})
	Processor.Register(&status.PlayerStatus{})
}
