package internal

import (
	"reflect"
	"server/conf/code"
	"server/msg/game"
	"server/msg/test"
	"server/user"
	"time"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	handler(&test.Hello{}, handlerHello)
	handler(&game.Add{}, handlerAdd)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handlerHello(args []interface{}) {
	m := args[0].(*test.Hello)
	a := args[1].(gate.Agent)

	log.Debug("ClientRemoteAddr:[%s] MsgId: Hello ,Name: %v, Text:%v, timestamp:%v", a.RemoteAddr().String(), m.Name, m.Text, m.Time)
	a.WriteMsg(&test.Hello{
		MsgId: "Hello",
		IncId: m.IncId,
		Name:  m.Name,
		Text:  "serverToClient==>handlerHello",
		Time:  int(time.Now().Unix()),
	})
}

func handlerAdd(args []interface{}) {
	m := args[0].(*game.Add)
	a := args[1].(gate.Agent)
	returnWriteMsg := &game.Add{
		MsgId: "Add",
		IncId: m.IncId,
		Code:  code.OK,
	}
	verifyResult := user.UsersData.VerifyPlayer(m.TokenText)
	if verifyResult == nil {
		returnWriteMsg.Code = code.PlayerIsNotLogin
		log.Debug("remoteAddr: [%s] error code:%v", a.RemoteAddr(), returnWriteMsg.Code)
	} else {
		log.Debug("remoteAddr: [%s] Add %d + %d = %d", a.RemoteAddr(), m.A, m.B, m.A+m.B)
		returnWriteMsg.A = m.A
		returnWriteMsg.B = m.B
		returnWriteMsg.Result = m.A + m.B
		returnWriteMsg.Time = int(time.Now().Unix())
	}

	a.WriteMsg(returnWriteMsg)

}
