package files

import (
	"github.com/therecipe/qt/quick"

	"github.com/therecipe/qt/internal/examples/showcases/sia/files/controller"
)

func init() { actionButtonTemplate_QmlRegisterType2("FilesTemplate", 1, 0, "ActionButtonTemplate") }

type actionButtonTemplate struct { //TODO: fix name clash
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ func(string) `signal:"showDownload"`
	_ func(string) `signal:"deleteRequest"`
}

func (t *actionButtonTemplate) init() {
	c := cfiles.ActionButtonController
	if c == nil {
		c = cfiles.NewActionButtonController(nil)
	}

	t.ConnectShowDownload(c.ShowDownload)
	t.ConnectDeleteRequest(c.DeleteRequest)
}
