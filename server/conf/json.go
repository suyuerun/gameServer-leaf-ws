package conf

import (
	"encoding/json"
	"io/ioutil"

	"github.com/name5566/leaf/log"
)

var Server struct {
	LogLevel    string
	LogPath     string
	LogConsole  bool
	WSAddr      string
	CertFile    string
	KeyFile     string
	TCPAddr     string
	TokenKey    string
	MaxConnNum  int
	ConsolePort int
	ProfilePath string
}

func init() {
	data, err := ioutil.ReadFile("../bin/conf/server.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
}
