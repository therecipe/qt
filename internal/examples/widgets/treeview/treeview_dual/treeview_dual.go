//author: https://github.com/5k3105

//  Extended example for two treeviews and changed the listview to a treeview that only shows files.
//  This way you can now see file details, unlike previous example.

package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

var (
	FsDirModel *widgets.QFileSystemModel
	FsTreeView *widgets.QTreeView

	FsFileModel *widgets.QFileSystemModel
	FsListView  *widgets.QTreeView

	DbDirModel *widgets.QFileSystemModel
	DbTreeView *widgets.QTreeView

	DbFileModel *widgets.QFileSystemModel
	DbListView  *widgets.QTreeView

	ap *Application
)

type Application struct {
	*widgets.QApplication
	Window    *widgets.QMainWindow
	Statusbar *widgets.QStatusBar
}

func main() {
	ap = &Application{}
	ap.QApplication = widgets.NewQApplication(len(os.Args), os.Args) // Application

	ap.Window = widgets.NewQMainWindow(nil, 0)
	ap.Window.SetWindowTitle("Treeview Example 2")

	SetupUi()

	widgets.QApplication_SetStyle2("fusion")
	ap.Window.Show()
	widgets.QApplication_Exec()
}

func SetupUi() {

	ap.Statusbar = widgets.NewQStatusBar(ap.Window) // Statusbar
	ap.Window.SetStatusBar(ap.Statusbar)

	var hlayout = widgets.NewQHBoxLayout()

	var lblLocal = widgets.NewQLabel2(" Local File System: C ", nil, 0)
	lblLocal.SetSizePolicy2(widgets.QSizePolicy__Fixed, widgets.QSizePolicy__Fixed)
	lblLocal.SetFont(gui.NewQFont2("verdana", 13, 1, false))

	var sp1 = widgets.NewQSplitter(nil)

	var FsTree = NewFsTree()
	var FsList = NewFsList()

	sp1.AddWidget(FsTree)
	sp1.AddWidget(FsList)

	var vlayout1 = widgets.NewQVBoxLayout()

	vlayout1.AddWidget(lblLocal, 0, 0)
	vlayout1.AddWidget(sp1, 0, 0)

	var lblRemote = widgets.NewQLabel2(" Remote File System: X", nil, 0)
	lblRemote.SetSizePolicy2(widgets.QSizePolicy__Fixed, widgets.QSizePolicy__Fixed)
	lblRemote.SetFont(gui.NewQFont2("verdana", 13, 1, false))

	var sp2 = widgets.NewQSplitter(nil)

	var DbTree = NewDbTree()
	var DbList = NewDbList()

	sp2.AddWidget(DbTree)
	sp2.AddWidget(DbList)

	var vlayout2 = widgets.NewQVBoxLayout()

	vlayout2.AddWidget(lblRemote, 0, 0)
	vlayout2.AddWidget(sp2, 0, 0)

	var btnMoveR = widgets.NewQPushButton(nil)
	btnMoveR.SetText("Move ->")
	var btnCopyR = widgets.NewQPushButton(nil)
	btnCopyR.SetText("Copy ->")
	var btnCopyL = widgets.NewQPushButton(nil)
	btnCopyL.SetText("<- Copy")
	var btnDeleteR = widgets.NewQPushButton(nil)
	btnDeleteR.SetText("Delete ->")
	var btnDeleteL = widgets.NewQPushButton(nil)
	btnDeleteL.SetText("<- Delete")

	var vlayout = widgets.NewQVBoxLayout()
	vlayout.AddWidget(btnMoveR, 0, 0)
	vlayout.AddWidget(btnCopyR, 0, 0)
	vlayout.AddWidget(btnCopyL, 0, 0)
	vlayout.AddWidget(btnDeleteR, 0, 0)
	vlayout.AddWidget(btnDeleteL, 0, 0)
	var spacer = widgets.NewQSpacerItem(0, 0, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	vlayout.AddItem(spacer)
	var groupBoxBtn = widgets.NewQGroupBox(nil)
	groupBoxBtn.SetLayout(vlayout)

	hlayout.AddLayout(vlayout1, 0)
	hlayout.AddWidget(groupBoxBtn, 0, 0)
	hlayout.AddLayout(vlayout2, 0)

	var groupBox = widgets.NewQGroupBox(nil)

	groupBox.SetLayout(hlayout)

	ap.Window.SetCentralWidget(groupBox)

	ap.Statusbar.ShowMessage("Application Path: "+core.QCoreApplication_ApplicationDirPath(), 0)

}

func NewFsList() *widgets.QTreeView {

	var drive string = "C:/"

	FsFileModel = widgets.NewQFileSystemModel(nil)

	FsFileModel.SetFilter(core.QDir__NoDotAndDotDot |
		core.QDir__Files)

	FsFileModel.SetRootPath(drive)

	FsListView = widgets.NewQTreeView(nil)
	FsListView.SetModel(FsFileModel)

	FsListView.Header().SetSectionResizeMode(widgets.QHeaderView__ResizeToContents)

	FsListView.SetSelectionMode(3)
	FsListView.SetIndentation(3)

	return FsListView
}

func NewFsTree() *widgets.QTreeView {

	var drive string = "C:/"

	FsDirModel = widgets.NewQFileSystemModel(nil)

	FsDirModel.SetFilter(core.QDir__NoDotAndDotDot |
		core.QDir__AllDirs)

	FsDirModel.SetRootPath(drive)

	FsTreeView = widgets.NewQTreeView(nil)
	FsTreeView.SetModel(FsDirModel)
	FsTreeView.SetRootIndex(FsDirModel.Index2(drive, 0))

	FsTreeView.HideColumn(1)
	FsTreeView.HideColumn(2)
	//FsTreeView.HideColumn(3)

	FsTreeView.Header().SetSectionResizeMode(widgets.QHeaderView__ResizeToContents)

	FsTreeView.ConnectCurrentChanged(treeViewCurrentChanged)

	return FsTreeView
}

func treeViewCurrentChanged(current *core.QModelIndex, previous *core.QModelIndex) {

	FilePath := FsDirModel.FilePath(current)

	FsListView.SetRootIndex(FsFileModel.SetRootPath(FilePath))

	FsTreeView.ScrollTo(current, 0) // 0 = EnsureVisible, 3 = PositionAtCenter

}

// ---------------------------------------

func NewDbList() *widgets.QTreeView {

	var drive string = "C:/"

	DbFileModel = widgets.NewQFileSystemModel(nil)

	DbFileModel.SetFilter(core.QDir__NoDotAndDotDot |
		core.QDir__Files)

	DbFileModel.SetRootPath(drive)

	DbListView = widgets.NewQTreeView(nil)
	DbListView.SetModel(DbFileModel)

	DbListView.Header().SetSectionResizeMode(widgets.QHeaderView__ResizeToContents)
	DbListView.SetIndentation(3)
	return DbListView
}

func NewDbTree() *widgets.QTreeView {

	var drive string = "C:/"

	DbDirModel = widgets.NewQFileSystemModel(nil)

	DbDirModel.SetFilter(core.QDir__NoDotAndDotDot |
		core.QDir__AllDirs)

	DbDirModel.SetRootPath(drive)

	DbTreeView = widgets.NewQTreeView(nil)
	DbTreeView.SetModel(DbDirModel)
	DbTreeView.SetRootIndex(DbDirModel.Index2(drive, 0))

	DbTreeView.HideColumn(1)
	DbTreeView.HideColumn(2)
	//FsTreeView.HideColumn(3)

	DbTreeView.Header().SetSectionResizeMode(widgets.QHeaderView__ResizeToContents)

	DbTreeView.ConnectCurrentChanged(treeDbViewCurrentChanged)

	return DbTreeView
}

func treeDbViewCurrentChanged(current *core.QModelIndex, previous *core.QModelIndex) {

	FilePath := DbDirModel.FilePath(current)

	DbListView.SetRootIndex(DbFileModel.SetRootPath(FilePath))

	DbTreeView.ScrollTo(current, 0) // 0 = EnsureVisible, 3 = PositionAtCenter

}
