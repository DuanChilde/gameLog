package controllers

import (
	"bytes"
	"gameLog/models"
	"gameLog/protos"
	"net/http"
	"time"

	"github.com/astaxie/beego"
	"github.com/golang/protobuf/proto"
)

// Operations about log
type LogController struct {
	beego.Controller
}

// @Title Record log
// @Description record log
// @Failure 403 body is empty
// @router / [post]
func (o *LogController) Post() {
	//解码
	log := &protos.LogMsg{}
	err := proto.Unmarshal(o.Ctx.Input.RequestBody, log)
	if err != nil {
		beego.Debug(err)
	}
	//请求与当前时间相隔超过时间限制，则不记录

	models.Add(log)
	o.Data["json"] = res.Success("success", map[string]interface{}{})
	o.ServeJSON()
}

// @Title create msg
// @Description get
// @Success 200 {string}
// @router / [get]
func (o *LogController) Get() {
	log := &protos.LogMsg{
		GameId: proto.String("pk365"),
		UserId: proto.String("17"),
		Event:  proto.String("login"),
		Ctime:  proto.String(time.Now().Format("2006-01-02 15:04:05")),
	}
	//beego.Debug(time.Now().Format("2006-01-02 15:04:05"))
	//编码
	data, err := proto.Marshal(log)
	if err != nil {
		beego.Debug(err)
	}

	body := bytes.NewReader(data)
	request, _ := http.NewRequest("POST", "http://192.168.1.150:8082/api/log", body)
	request.Header.Set("Connection", "Keep-Alive")
	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		beego.Debug(err)
	}
	defer resp.Body.Close()
}
