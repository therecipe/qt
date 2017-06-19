// +build qml

package album

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"

	"github.com/therecipe/qt/internal/examples/sql/masterdetail_qml/controller"
)

func init() {
	albumController_QmlRegisterType2("Album", 1, 0, "AlbumController")
}

type albumController struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ *core.QAbstractItemModel `property:"viewModel"`

	//qml<-controller
	_ func() `signal:"albumAdded"`
	_ func() `signal:"deleteAlbumRequest"`
	_ func() `signal:"deleteAlbumCommand"`

	//qml->controller
	_ func()                                     `signal:"showImageLabel"`
	_ func(index *core.QModelIndex)              `signal:"deleteAlbum"`
	_ func(column int, order core.Qt__SortOrder) `signal:"sortTableView"`
	_ func(index *core.QModelIndex)              `signal:"showAlbumDetails"`
	_ func(title, artist string)                 `signal:"deleteAlbumShowRequest"`
}

func (a *albumController) init() {

	a.SetViewModel(controller.Instance.AlbumModel())

	//<-controller
	controller.Instance.ConnectAddAlbum(func(string, string, int, string) { a.AlbumAdded() })
	controller.Instance.ConnectDeleteAlbumRequest(a.DeleteAlbumRequest)
	controller.Instance.ConnectDeleteAlbumCommand(a.DeleteAlbumCommand)

	//->controller
	a.ConnectShowImageLabel(controller.Instance.ShowImageLabel)
	a.ConnectDeleteAlbum(func(index *core.QModelIndex) { controller.Instance.DeleteAlbum(index) })
	a.ConnectSortTableView(controller.Instance.SortTableView)
	a.ConnectShowAlbumDetails(func(index *core.QModelIndex) { controller.Instance.ShowAlbumDetails(index) })
	a.ConnectDeleteAlbumShowRequest(controller.Instance.DeleteAlbumShowRequest)
}
