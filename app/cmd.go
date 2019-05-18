package app

import (
	"github.com/wailovet/cmdgui/helper"
	"github.com/wailovet/osmanthuswine/src/core"
)

type Cmd struct {
	core.Controller
}

func (that *Cmd) Delete() {
	cd := helper.CmdData{}
	err := that.RequestToStruct(&cd)
	that.CheckErrDisplayByError(err)
	err = helper.DeleteCmdData(cd.CmdDataID)
	that.CheckErrDisplayByError(err)
	that.DisplayBySuccess("成功")
}

func (that *Cmd) Save() {
	cd := helper.CmdData{}
	err := that.RequestToStruct(&cd)
	that.CheckErrDisplayByError(err)
	err = helper.SaveCmdData(cd)
	that.CheckErrDisplayByError(err)
	that.DisplayBySuccess("成功")
}

func (that *Cmd) Get() {
	data := helper.GetCmdData()
	that.DisplayByData(data)
}

func (that *Cmd) Start() {
	cmd := that.Request.REQUEST["cmd"]

	pty := helper.NewPty("sh")
	pty.Start(func(data []byte) {
		println(string(data))
	})
	pty.Write([]byte(cmd + "\nexit\n"))

}
