package helper

import (
	"github.com/sonyarouje/simdb/db"
	"github.com/wailovet/osmanthuswine/src/helper"
)

type CmdData struct {
	CmdDataID string `json:"cmd_data_id"`
	Name      string `json:"name"`
	Cmd       string `json:"cmd"`
}

func (that CmdData) ID() (jsonField string, value interface{}) {
	value = that.CmdDataID
	jsonField = "cmd_data_id"
	return
}

func GetCmdData() (data []CmdData) {

	driver, err := db.New("cmdgui")
	if err != nil {
		return data
	}
	_ = driver.Open(CmdData{}).Get().AsEntity(&data)
	return data
}
func DeleteCmdData(id string) error {
	driver, err := db.New("cmdgui")
	if err != nil {
		return err
	}
	return driver.Delete(CmdData{
		CmdDataID: id,
	})
}

func SaveCmdData(cd CmdData) error {
	driver, err := db.New("cmdgui")
	if err != nil {
		return err
	}
	if cd.CmdDataID != "" {
		return driver.Update(cd)
	} else {
		cd.CmdDataID = helper.CreateUUID()
		return driver.Insert(cd)
	}

}
