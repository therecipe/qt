package controller

import (
	"encoding/json"
	"time"

	"github.com/therecipe/qt/core"

	"github.com/therecipe/qt/internal/examples/showcases/wallet/files/model"
)

var FilesController *filesController

type filesController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ *model.FilesModel `property:"model"`
}

func (c *filesController) init() {
	FilesController = c

	c.SetModel(model.NewFilesModel(nil))

	go c.loop()
}

func (c *filesController) loop() {
	for range time.NewTicker(1 * time.Second).C {
		if DEMO {
			var df []model.File
			json.Unmarshal([]byte(DEMO_FILES), &df)
			c.Model().UpdateWith(df)
		}
	}
}
