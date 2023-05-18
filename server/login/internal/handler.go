package internal

import (
	"reflect"
	"server/user"
	"time"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"

	"server/msg"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handler(&msg.Login{}, handlerLogin)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handlerLogin(args []interface{}) {
	m := args[0].(*msg.Login)
	a := args[1].(gate.Agent)

	log.Debug("[%s] Login Name:%s Pwd:%s", a.RemoteAddr().String(), m.Username, m.Password)
	player := user.NewPlayer(m.Username, m.Password, a)
	a.WriteMsg(&msg.Login{
		MsgId:     "Login",
		Username:  player.Username,
		Password:  player.Password,
		TokenText: player.TokenText,
		Time:      int(time.Now().Unix()),
	})
}
