package dialog

import (
	"github.com/bluszcz/cutego/widgets"

	_ "github.com/bluszcz/cutego/internal/examples/showcases/wallet/files/dialog/controller"
)

type filesUploadTemplate struct {
	filesDialogTemplate

	_ func([]string) `slot:"uploadFiles,->(controller.Controller)"`
}

func (t *filesUploadTemplate) show(cident string) {
	if cident == "files" {
		t.Blur(true)
		t.UploadFiles(widgets.QFileDialog_GetOpenFileNames(nil, "Select one or more files to upload", "", "", "", 0))
		t.Blur(false)
	}
}
