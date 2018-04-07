package controller

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/therecipe/qt/core"

	maincontroller "github.com/therecipe/qt/internal/examples/showcases/sia/controller"
	_ "github.com/therecipe/qt/internal/examples/showcases/sia/view/controller"
)

var Controller *dialogController

type dialogController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(cident string) `signal:"show"`
	_ func(string)        `signal:"showDownload"`
	_ func(bool)          `signal:"blur,->(controller.Controller)"`

	_ func([]string)       `signal:"uploadFiles,auto"`
	_ func(string)         `signal:"uploadFolder,auto"`
	_ func(string, string) `signal:"download,auto"`

	_ func() bool `slot:"isLocked,->(maincontroller.Controller)"`
}

func (c *dialogController) init() {
	Controller = c
}

func (c *dialogController) uploadFiles(files []string) {
	for _, f := range files {
		go maincontroller.Client.RenterUploadPost("/renter/upload/"+filepath.Base(f), "source="+f, 1, 1) //TODO: dataPieces, parityPieces ?
	}
}

func (c *dialogController) uploadFolder(source string) {
	var files []string
	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			println("Warning: skipping file:", err.Error())
			return nil
		}
		if info.IsDir() {
			return nil
		}
		if strings.HasPrefix(info.Name(), ".") {
			return nil
		}
		files = append(files, path)
		return nil
	})
	c.uploadFiles(files)
}

func (c *dialogController) download(name, path string) {
	go func() {
		err := maincontroller.Client.RenterDownloadGet(name, path, 0, 1, true) //TODO: offset, lenght, async ?
		if err != nil {
			println("Could not download file:", err.Error())
		}
	}()
}
