package user

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"server/conf"
	"sync"
)

type Player struct {
	Username  string
	Password  string
	TokenText string
	Agent     gate.Agent
}
type Users struct {
	UserMap     map[string]*Player
	UserRWMutex sync.RWMutex
}

var UsersData *Users

func init() {
	UsersData = &Users{}
	UsersData.UserMap = make(map[string]*Player)
}
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
	ShowAllPlayers()
	return player
}
func ShowAllPlayers() {
	for key, value := range UsersData.UserMap {
		log.Debug("ShowAllPlayers-player.key: %v, player.data: %v", key, value)
	}
	log.Debug("ShowAllPlayers-\u001B[33m show all finished \u001B[0m")
}
