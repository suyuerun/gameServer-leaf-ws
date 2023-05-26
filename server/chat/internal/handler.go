package internal

import (
	"fmt"
	"reflect"
	"server/conf/code"
	"server/msg"
	"server/user"
	"time"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	handler(&msg.Chat{}, handlerChat)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handlerChat(args []interface{}) {
	m := args[0].(*msg.Chat)
	a := args[1].(gate.Agent)
	returnWriteMsg := &msg.Chat{
		MsgId: "Chat",
		IncId: m.IncId,
		Code:  code.OK,
	}
	verifyResult := user.UsersData.VerifyPlayer(m.TokenText)
	if verifyResult == nil {
		returnWriteMsg.Code = code.PlayerIsNotLogin
		log.Debug("remoteAddr: [%s] error code:%v", a.RemoteAddr(), returnWriteMsg.Code)
	} else {
		log.Debug("remoteAddr: [%s] Chat to %s  content: %s", a.RemoteAddr(), m.To, m.Content)
		returnWriteMsg.To = m.To
		returnWriteMsg.Content = m.Content
		returnWriteMsg.Time = int(time.Now().Unix())
		if m.To == "All" {
			chatToAll(m.Content, verifyResult)
		}
	}

	a.WriteMsg(returnWriteMsg)

}
func chatToAll(content string, player *user.Player) {
	user.UsersData.BroadcastToAll(&msg.Broadcast{
		MsgId:   "Broadcast",
		Message: fmt.Sprintf("{\"player\":\"%s\",\"opt\":\"%s\",\"content\":\"%s\"}", player.Username, "chatToAll", content),
		Time:    int(time.Now().Unix()),
	})
}
