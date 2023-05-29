package internal

import (
	"fmt"
	"reflect"
	"server/conf/code"
	"server/msg/chat"
	"server/msg/status"
	"server/user"
	"time"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	handler(&status.PlayerStatus{}, handlerStatus)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handlerStatus(args []interface{}) {
	m := args[0].(*status.PlayerStatus)
	a := args[1].(gate.Agent)
	returnWriteMsg := &status.PlayerStatus{
		MsgId: "Status",
		IncId: m.IncId,
		Code:  code.OK,
	}
	for true {
		verifyResult := user.UsersData.VerifyPlayer(m.TokenText)
		if verifyResult == nil {
			returnWriteMsg.Code = code.PlayerIsNotLogin
			log.Debug("remoteAddr: [%s] error code:%v", a.RemoteAddr(), returnWriteMsg.Code)
			break
		}

		returnWriteMsg.Result = make([]*status.Player, 0)
		if len(m.PlayerNames) == 0 { // 则返回所有
			for _, player := range user.UsersData.UserMap {
				playerStatus := &status.Player{
					Username:  player.Username,
					TokenText: player.TokenText,
				}
				returnWriteMsg.Result = append(returnWriteMsg.Result, playerStatus)
			}
			break
		}
		// 否则返回 参数内的人 支持多个
		log.Debug("remoteAddr: [%s] get player status names %v", a.RemoteAddr(), m.PlayerNames)
		for _, name := range m.PlayerNames {
			if user.UsersData.UserMap[name] == nil {
				continue
			}

		}

	}
	returnWriteMsg.Time = int(time.Now().Unix())
	a.WriteMsg(returnWriteMsg)
}
func chatToAll(content string, player *user.Player) {
	user.UsersData.BroadcastToAll(&chat.Broadcast{
		MsgId:   "Broadcast",
		Message: fmt.Sprintf("{\"player\":\"%s\",\"opt\":\"%s\",\"content\":\"%s\"}", player.Username, "chatToAll", content),
		Time:    int(time.Now().Unix()),
	})
}
