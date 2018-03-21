package cdialog

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/therecipe/qt/core"

	"github.com/therecipe/qt/internal/examples/showcases/sia/controller"
	"github.com/therecipe/qt/internal/examples/showcases/sia/view/controller"
)

var Controller *filesDialogController

type filesDialogController struct { //TODO: fix name clash
	core.QObject

	_ func() `constructor:"init"`

	_ func(cident string) `signal:"show"`
	_ func(string)        `signal:"showDownload"`
	_ func(bool)          `signal:"blur"`

	_ func([]string)       `signal:"uploadFiles"`
	_ func(string)         `signal:"uploadFolder"`
	_ func(string, string) `signal:"download"`

	_ func() bool `slot:"isLocked"`
}

func (c *filesDialogController) init() {
	Controller = c

	c.ConnectBlur(cview.Controller.Blur)

	c.ConnectUploadFiles(c.uploadFiles)
	c.ConnectUploadFolder(c.uploadFolder)
	c.ConnectDownload(c.download)

	c.ConnectIsLocked(controller.Controller.IsLocked)
}

func (c *filesDialogController) uploadFiles(files []string) {
	for _, f := range files {
		go controller.Client.RenterUploadPost("/renter/upload/"+filepath.Base(f), "source="+f, 1, 1) //TODO: dataPieces, parityPieces ?
	}
}

func (c *filesDialogController) uploadFolder(source string) {
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

func (c *filesDialogController) download(name, path string) {
	go func() {
		err := controller.Client.RenterDownloadGet(name, path, 0, 1, true) //TODO: offset, lenght, async ?
		if err != nil {
			println("Could not download file:", err.Error())
		}
	}()
}
