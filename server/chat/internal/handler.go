package internal

import (
	"fmt"
	"reflect"
	"server/conf/code"
	"server/msg/chat"
	"server/user"
	"time"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	handler(&chat.Chat{}, handlerChat)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handlerChat(args []interface{}) {
	m := args[0].(*chat.Chat)
	a := args[1].(gate.Agent)
	returnWriteMsg := &chat.Chat{
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
		} else {
			chatToSecret(m.To, m.Content, verifyResult)
		}
	}

	a.WriteMsg(returnWriteMsg)

}
func chatToAll(content string, player *user.Player) {
	user.UsersData.BroadcastToAll(&chat.Broadcast{
		MsgId:   "Broadcast",
		Message: fmt.Sprintf("{\"player\":\"%s\",\"opt\":\"%s\",\"content\":\"%s\"}", player.Username, "chatToAll", content),
		Time:    int(time.Now().Unix()),
	})
}
func chatToSecret(toPlayerKey string, content string, player *user.Player) {
	toPlayer := user.UsersData.VerifyPlayer(toPlayerKey)
	if toPlayer == nil {
		log.Release("remoteAddr: [%s] Chat to %s  content: %s failed ! player not found", player.Agent.RemoteAddr(), toPlayerKey, content)
		return
	}
	user.UsersData.ChatToSecrete(toPlayer, &chat.ChatToSecret{
		MsgId:   "ChatToSecret",
		From:    player.Username,
		To:      toPlayer.Username,
		Message: fmt.Sprintf("{\"player\":\"%s\",\"opt\":\"%s\",\"content\":\"%s\"}", player.Username, "chatToSecret", content),
		Time:    int(time.Now().Unix()),
	})
}
