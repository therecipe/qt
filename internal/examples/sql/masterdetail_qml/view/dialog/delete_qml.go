// +build qml

package dialog

import (
	"github.com/therecipe/qt/quick"

	"github.com/therecipe/qt/internal/examples/sql/masterdetail_qml/controller"
)

func init() {
	deleteDialogController_QmlRegisterType2("Dialog", 1, 0, "DeleteDialogController")
}

type deleteDialogController struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	//qml<-controller
	_ func(title, artist string) `signal:"deleteAlbumShowRequest"`

	//qml->controller
	_ func() `signal:"deleteAlbumCommand"`
}

func (d *deleteDialogController) init() {

	//qml<-controller
	controller.Instance.ConnectDeleteAlbumShowRequest(d.DeleteAlbumShowRequest)

	//qml->controller
	d.ConnectDeleteAlbumCommand(controller.Instance.DeleteAlbumCommand)
}
