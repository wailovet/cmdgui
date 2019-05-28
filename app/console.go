package app

import (
	"github.com/wailovet/osmanthuswine/src/core"
	"gopkg.in/olahol/melody.v1"
	"log"
	"sync"
)

type Console struct {
	core.WebSocket
}

var instanceMapFunName sync.Map

func GetInstanceConsole(funName string) *Console {
	result, ok := instanceMapFunName.Load(funName)
	if ok {
		return result.(*Console)
	}

	return nil
}

func (that *Console) HandleConnect(*melody.Session) {
	//panic("implement me")
	log.Println("HandleConnect:", that.FunName)
	instanceMapFunName.Delete(that.FunName)
	instanceMapFunName.Store(that.FunName, that)
}

func (that *Console) HandlePong(*melody.Session) {
	//panic("implement me")
}

func (that *Console) HandleMessage(*melody.Session, []byte) {
	//panic("implement me")
}

func (that *Console) HandleMessageBinary(*melody.Session, []byte) {
	//panic("implement me")
}

func (that *Console) HandleSentMessage(*melody.Session, []byte) {
	//panic("implement me")
}

func (that *Console) HandleSentMessageBinary(*melody.Session, []byte) {
	//panic("implement me")
}

func (that *Console) HandleDisconnect(*melody.Session) {
	//panic("implement me")
}

func (that *Console) HandleError(*melody.Session, error) {
	//panic("implement me")
}
