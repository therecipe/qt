//go:build qml
// +build qml

package view

import (
	"github.com/akiyosi/qt/core"

	"github.com/akiyosi/qt/internal/examples/sql/masterdetail_qml/controller"

	_ "github.com/akiyosi/qt/internal/examples/sql/masterdetail_qml/view/album"
	_ "github.com/akiyosi/qt/internal/examples/sql/masterdetail_qml/view/artist"
	_ "github.com/akiyosi/qt/internal/examples/sql/masterdetail_qml/view/detail"
	_ "github.com/akiyosi/qt/internal/examples/sql/masterdetail_qml/view/dialog"
)

func init() {
	viewController_QmlRegisterType2("View", 1, 0, "ViewController")
}

type viewController struct {
	core.QObject

	_ func() `constructor:"init"`

	//qml->controller
	_ func() `signal:"aboutQt"`
	_ func() `signal:"addAlbumShowRequest"`
	_ func() `signal:"deleteAlbumRequest"`
}

func (v *viewController) init() {

	//qml->controller
	v.ConnectAboutQt(controller.Instance.AboutQt)
	v.ConnectAddAlbumShowRequest(controller.Instance.AddAlbumShowRequest)
	v.ConnectDeleteAlbumRequest(controller.Instance.DeleteAlbumRequest)
}
