package main

import (
	_ "github.com/miriko-channel/NekoBox/routers"
	"github.com/miriko-channel/NekoBox/template"
	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.ServerName = "NekoBox"
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "nekoboxSession"
    //fix linter warning 'Error return value of `beego.AddFuncMap` is not checked (errcheck)'
    _ = beego.AddFuncMap("answerFormat", template.AnswerFormat)
	beego.Run()
}
