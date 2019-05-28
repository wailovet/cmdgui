package app

import (
	"github.com/wailovet/cmdgui/helper"
	"github.com/wailovet/osmanthuswine/src/core"
	"runtime"
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
	var pty *helper.Pty
	if runtime.GOOS == "windows" {
		pty = helper.NewPty("powershell.exe")
		_ = pty.Start(func(data []byte) {
			ws := GetInstanceConsole("All")
			if ws != nil {
				ws.GetWebSocket().Broadcast(data)
			}
			print(string(data))
		})
		_, _ = pty.Write([]byte(cmd + "\r\nexit\r\n"))
	} else {
		pty = helper.NewPty("sh")
		_ = pty.Start(func(data []byte) {
			ws := GetInstanceConsole("All")
			if ws != nil {
				ws.GetWebSocket().Broadcast(data)
			}
			print(string(data))
		})
		_, _ = pty.Write([]byte(cmd + "\nexit\n"))
	}

}
