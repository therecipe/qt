// +build qml

package dialog

import (
	"github.com/therecipe/qt/quick"

	"github.com/therecipe/qt/internal/examples/sql/masterdetail_qml/controller"
)

func init() {
	addDialogController_QmlRegisterType2("Dialog", 1, 0, "AddDialogController")
}

type addDialogController struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	//qml<-controller
	_ func() `signal:"addAlbumShowRequest"`

	//qml->controller
	_ func(artist string, title string, year int, tracks string) `signal:"addAlbum"`
}

func (a *addDialogController) init() {

	//qml<-controller
	controller.Instance.ConnectAddAlbumShowRequest(a.AddAlbumShowRequest)

	//qml->controller
	a.ConnectAddAlbum(controller.Instance.AddAlbum)
}
