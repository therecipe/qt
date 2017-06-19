// +build !qml

package view

import (
	"github.com/therecipe/qt/widgets"

	"github.com/therecipe/qt/internal/examples/sql/masterdetail_qml/controller"

	"github.com/therecipe/qt/internal/examples/sql/masterdetail_qml/view/album"
	"github.com/therecipe/qt/internal/examples/sql/masterdetail_qml/view/artist"
	"github.com/therecipe/qt/internal/examples/sql/masterdetail_qml/view/detail"
	"github.com/therecipe/qt/internal/examples/sql/masterdetail_qml/view/dialog"
)

var ViewControllerInstance *viewController

type viewController struct {
	widgets.QMainWindow

	_ func() `constructor:"init"`

	//->controller
	_ func() `signal:"aboutQt"`
	_ func() `signal:"addAlbumShowRequest"`
	_ func() `signal:"deleteAlbumRequest"`
}

func (v *viewController) init() {
	ViewControllerInstance = v

	layout := widgets.NewQGridLayout2()
	layout.AddWidget(artist.NewArtistController2("Artist", nil), 0, 0, 0)
	layout.AddWidget(album.NewAlbumController2("Album", nil), 1, 0, 0)
	layout.AddWidget3(detail.NewDetailController2("Detail", nil), 0, 1, 2, 1, 0)
	layout.SetColumnStretch(1, 1)
	layout.SetColumnMinimumWidth(0, 500)

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(layout)
	v.SetCentralWidget(widget)

	v.Resize2(850, 400)
	v.SetWindowTitle("Music Archive")

	dialog.NewAddDialogController(nil, 0)
	dialog.NewDeleteDialogController(nil)
	NewMenuBarController(nil)

	//->controller
	v.ConnectAboutQt(controller.Instance.AboutQt)
	v.ConnectAddAlbumShowRequest(controller.Instance.AddAlbumShowRequest)
	v.ConnectDeleteAlbumRequest(controller.Instance.DeleteAlbumRequest)
}
