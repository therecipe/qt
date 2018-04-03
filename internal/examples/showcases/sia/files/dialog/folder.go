package dialog

import (
	"github.com/therecipe/qt/widgets"

	"github.com/therecipe/qt/internal/examples/showcases/sia/files/dialog/controller"
)

type folderUploadTemplate struct {
	filesDialogTemplate

	_ func() `constructor:"init"`

	_ func(string)         `signal:"uploadFolder"`
	_ func(string, string) `signal:"download"`
}

func (t *folderUploadTemplate) init() {
	t.ConnectUploadFolder(controller.Controller.UploadFolder)
	t.ConnectShow(t.show)

	t.ConnectDownload(controller.Controller.Download)
	controller.Controller.ConnectShowDownload(t.showDownload)
}

func (t *folderUploadTemplate) show(cident string) {
	if cident == "folder" {
		t.Blur(true)
		t.UploadFolder(widgets.QFileDialog_GetExistingDirectory(nil, "Select a folder to upload", "", widgets.QFileDialog__ShowDirsOnly|widgets.QFileDialog__DontResolveSymlinks))
		t.Blur(false)
	}
}

func (t *folderUploadTemplate) showDownload(name string) {
	t.Blur(true)
	t.Download(name, widgets.QFileDialog_GetExistingDirectory(nil, "Select a folder to download "+name, "", widgets.QFileDialog__ShowDirsOnly|widgets.QFileDialog__DontResolveSymlinks))
	t.Blur(false)
}
