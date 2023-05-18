package internal

import (
	"fmt"
	"reflect"
	"server/msg"
	"time"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	handler(&msg.Hello{}, handlerHello)
	handler(&msg.Add{}, handlerAdd)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handlerHello(args []interface{}) {
	m := args[0].(*msg.Hello)
	a := args[1].(gate.Agent)

	log.Debug("ClientRemoteAddr:[%s] MsgId: Hello ,Name: %v, Text:%v, timestamp:%v", a.RemoteAddr().String(), m.Name, m.Text, m.Time)
	a.WriteMsg(&msg.Hello{
		MsgId: "Hello",
		Name:  m.Name,
		Text:  "serverToClient==>handlerHello",
		Time:  int(time.Now().Unix()),
	})
}

func handlerAdd(args []interface{}) {
	m := args[0].(*msg.Add)
	a := args[1].(gate.Agent)

	log.Debug("[%s] Add %d + %d = %d", a.RemoteAddr(), m.A, m.B, m.A+m.B)
	a.WriteMsg(&msg.Hello{
		Name: fmt.Sprintf("handlerAdd %d + %d = %d", m.A, m.B, m.A+m.B),
	})
}
