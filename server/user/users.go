package user

import (
	"fmt"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"server/conf"
	"server/msg"
	"sync"
	"time"
)

type Users struct {
	UserMap     map[string]*Player
	UserRWMutex sync.RWMutex
}

var UsersData *Users

// 初始化生成 玩家 全局变量
func init() {
	UsersData = &Users{}
	UsersData.UserMap = make(map[string]*Player)
}

// 初始化单个玩家, 加入全局玩家map, 并打印所有当前玩家
func NewPlayer(Username string, Password string, Agent gate.Agent) *Player {
	player := &Player{
		Username:  Username,
		Password:  Password,
		TokenText: conf.Server.TokenKey + "@" + Username,
		Agent:     Agent,
	}
	UsersData.UserRWMutex.Lock()
	defer UsersData.UserRWMutex.Unlock()
	UsersData.UserMap[player.TokenText] = player
	log.Debug("funcName:NewPlayer UserMap:%v", &UsersData.UserMap)
	UsersData.BroadcastToAll(&msg.Broadcast{
		MsgId:   "Broadcast",
		Message: fmt.Sprintf("{\"player\":\"%s\",\"opt\":\"%s\"}", player.Username, "online"),
		Time:    int(time.Now().Unix()),
	})
	UsersData.ShowAllPlayers()
	return player
}

// 打印所有当前玩家
func (u *Users) ShowAllPlayers() {
	for key, value := range UsersData.UserMap {
		log.Debug("ShowAllPlayers-player.key: %v, player.data: %v", key, value)
	}
	log.Debug("ShowAllPlayers-\u001B[33m show all Finished! \u001B[0m")
}

// 验证玩家是否在全局map里不在返回nil, 在则返回player
func (u *Users) VerifyPlayer(tokenText string) *Player {
	UsersData.UserRWMutex.RLock()
	defer UsersData.UserRWMutex.RUnlock()
	if UsersData.UserMap[tokenText] == nil {
		return nil
	}
	return UsersData.UserMap[tokenText]
}

// 广播给所有玩家
func (u *Users) BroadcastToAll(BroadcastMsg *msg.Broadcast) {
	for _, player := range u.UserMap {
		player.Agent.WriteMsg(BroadcastMsg)
	}

}
