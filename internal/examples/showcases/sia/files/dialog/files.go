package dialog

import (
	"github.com/therecipe/qt/widgets"

	"github.com/therecipe/qt/internal/examples/showcases/sia/files/dialog/controller"
)

type filesUploadTemplate struct {
	filesDialogTemplate

	_ func() `constructor:"init"`

	_ func([]string) `slot:"uploadFiles"`
}

func (t *filesUploadTemplate) init() {
	t.ConnectUploadFiles(cdialog.Controller.UploadFiles)
	t.ConnectShow(t.show)
}

func (t *filesUploadTemplate) show(cident string) {
	if cident == "files" {
		t.Blur(true)
		t.UploadFiles(widgets.QFileDialog_GetOpenFileNames(nil, "Select one or more files to upload", "", "", "", 0))
		t.Blur(false)
	}
}
