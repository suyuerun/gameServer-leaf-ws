package main

import (
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	"server/chat"
	"server/conf"
	"server/game"
	"server/gate"
	"server/login"
	"server/status"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogConsole = conf.Server.LogConsole
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	leaf.Run(
		game.Module,
		chat.Module,
		status.Module,
		gate.Module,
		login.Module,
	)
}

// The system cannot find the path specified
