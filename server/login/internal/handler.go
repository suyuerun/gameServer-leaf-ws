package internal

import (
	"fmt"
	"reflect"
	"server/conf"
	"server/msg/account"
	"server/msg/chat"
	"server/user"
	"time"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handler(&account.Login{}, handlerLogin)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handlerLogin(args []interface{}) {
	m := args[0].(*account.Login)
	a := args[1].(gate.Agent)
	verifyResult := user.UsersData.VerifyPlayer(conf.Server.TokenKey + "@" + m.Username)
	// 踢掉之前登录的人
	if verifyResult != nil {
		verifyResult.Agent.WriteMsg(&chat.Broadcast{
			MsgId:   "Kicked",
			Message: fmt.Sprintf("{\"player\":\"%s\",\"opt\":\"%s\"}", verifyResult.Username, "PlayerMultiLogin"),
			Time:    int(time.Now().Unix()),
		})
		log.Debug("remoteAddr: [%s] error PlayerMultiLogin kicked:%v", a.RemoteAddr())
	}
	log.Debug("[%s] Login Name:%s Pwd:%s", a.RemoteAddr().String(), m.Username, m.Password)
	player := user.NewPlayer(m.Username, m.Password, a)
	a.WriteMsg(&account.Login{
		MsgId:     "Login",
		IncId:     m.IncId,
		Username:  player.Username,
		Password:  player.Password,
		TokenText: player.TokenText,
		Time:      int(time.Now().Unix()),
	})
}
