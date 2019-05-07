package dialog

import (
	"github.com/therecipe/qt/widgets"

	_ "github.com/therecipe/qt/internal/examples/showcases/wallet/files/dialog/controller"
)

type folderUploadTemplate struct {
	filesDialogTemplate

	_ func(string) `signal:"uploadFolder,->(controller.Controller)"`

	_ func(string, string) `signal:"download,->(controller.Controller)"`
	_ func(string)         `signal:"showDownload,<-(controller.Controller)"`
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
