// +build !qml

package album

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"

	"github.com/therecipe/qt/internal/examples/sql/masterdetail_qml/controller"
)

type albumController struct {
	widgets.QGroupBox

	_ func() `constructor:"init"`

	_ *core.QAbstractItemModel `property:"viewModel"`

	//<-controller
	_ func() `signal:"albumAdded"`
	_ func() `signal:"deleteAlbumRequest"`
	_ func() `signal:"deleteAlbumCommand"`

	//->controller
	_ func()                                     `signal:"showImageLabel"`
	_ func(index *core.QModelIndex)              `signal:"deleteAlbum"`
	_ func(column int, order core.Qt__SortOrder) `signal:"sortTableView"`
	_ func(index *core.QModelIndex)              `signal:"showAlbumDetails"`
	_ func(title, artist string)                 `signal:"deleteAlbumShowRequest"`

	//

	albumView *widgets.QTableView
}

func (a *albumController) init() {

	a.albumView = widgets.NewQTableView(nil)
	a.albumView.SetModel(controller.Instance.AlbumModel())
	a.albumView.SetEditTriggers(widgets.QAbstractItemView__NoEditTriggers)
	a.albumView.SetSortingEnabled(true)
	a.albumView.SetSelectionBehavior(widgets.QAbstractItemView__SelectRows)
	a.albumView.SetSelectionMode(widgets.QAbstractItemView__SingleSelection)
	a.albumView.SetShowGrid(false)
	a.albumView.VerticalHeader().Hide()
	a.albumView.SetAlternatingRowColors(true)
	a.adjustHeader()

	locale := a.albumView.Locale()
	locale.SetNumberOptions(core.QLocale__OmitGroupSeparator)
	a.albumView.SetLocale(locale)

	a.albumView.ConnectClicked(func(index *core.QModelIndex) { a.ShowAlbumDetails(index) })
	a.albumView.ConnectActivated(func(index *core.QModelIndex) { a.ShowAlbumDetails(index) })

	layout := widgets.NewQVBoxLayout()
	layout.AddWidget(a.albumView, 0, 0)
	a.SetLayout(layout)

	//

	//<-controller
	controller.Instance.ConnectAddAlbum(func(string, string, int, string) { a.albumAdded() })
	controller.Instance.ConnectDeleteAlbumRequest(a.deleteAlbumRequest)
	controller.Instance.ConnectDeleteAlbumCommand(a.deleteAlbumCommand)

	//->controller
	a.ConnectShowImageLabel(controller.Instance.ShowImageLabel)
	a.ConnectDeleteAlbum(func(index *core.QModelIndex) { controller.Instance.DeleteAlbum(index) })
	a.ConnectSortTableView(controller.Instance.SortTableView)
	a.ConnectShowAlbumDetails(func(index *core.QModelIndex) { controller.Instance.ShowAlbumDetails(index) })
	a.ConnectDeleteAlbumShowRequest(controller.Instance.DeleteAlbumShowRequest)
}

func (a *albumController) albumAdded() {
	a.albumView.SelectionModel().Clear()
	a.albumView.SelectRow(a.albumView.Model().RowCount(core.NewQModelIndex()) - 1)
	a.ShowAlbumDetails(a.albumView.Model().Index(a.albumView.Model().RowCount(core.NewQModelIndex())-1, 0, core.NewQModelIndex()))
	a.adjustHeader()
}

func (a *albumController) deleteAlbumRequest() {
	index := a.albumView.CurrentIndex()
	title := index.Data(int(core.Qt__UserRole) + 2).ToString()
	artist := index.Data(int(core.Qt__UserRole) + 3).ToString()
	a.DeleteAlbumShowRequest(title, artist)
}

func (a *albumController) deleteAlbumCommand() {
	a.DeleteAlbum(a.albumView.CurrentIndex())
	a.ShowImageLabel()
	a.adjustHeader()
}

func (a *albumController) adjustHeader() {
	a.albumView.HideColumn(0)
	a.albumView.HorizontalHeader().SetSectionResizeMode2(1, widgets.QHeaderView__Stretch)
	a.albumView.ResizeColumnToContents(2)
	a.albumView.ResizeColumnToContents(3)
}
