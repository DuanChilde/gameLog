package models

import (
	"gameLog/protos"
	"os"

	"github.com/astaxie/beego"
)

var (
	Logs map[string]*Log
)

type Log struct {
	GameId string
	UserId string
	Event  string
	Ctime  string
}

func init() {

}

func Add(log *protos.LogMsg) {
	//log.Ctime = "123"
	f, err := os.OpenFile("eventLog.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		beego.Debug("dir not exists")
	}
	defer f.Close()
	str := log.GetGameId() + "|" + log.GetUserId() + "|" + log.GetEvent() + "|" + log.GetCtime() + "|\n"
	buf := []byte(str)
	f.Write(buf)

}

func AddOne(log Log) {
	//log.Ctime = "123"
	f, err := os.OpenFile("eventLog.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		beego.Debug("dir not exists")
	}
	defer f.Close()
	str := log.GameId + "|" + log.UserId + "|" + log.Event + "|" + log.Ctime + "|\n"
	buf := []byte(str)
	f.Write(buf)
}
