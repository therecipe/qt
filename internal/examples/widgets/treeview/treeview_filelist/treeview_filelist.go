//author: https://github.com/5k3105

//  This is a filesystem treeview with a listview for individual files on the right.
//  Approx translation from: http://www.bogotobogo.com/Qt/Qt5_QTreeView_QFileSystemModel_ModelView_MVC.php

package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

var (
	FsDirModel *widgets.QFileSystemModel
	FsTreeView *widgets.QTreeView

	FsFileModel *widgets.QFileSystemModel
	FsListView  *widgets.QListView
)

type Application struct {
	*widgets.QApplication
	Window    *widgets.QMainWindow
	Statusbar *widgets.QStatusBar
}

func main() {
	var ap = &Application{}
	ap.QApplication = widgets.NewQApplication(len(os.Args), os.Args) // Application

	ap.Window = widgets.NewQMainWindow(nil, 0)
	ap.Window.SetWindowTitle("Fs Example")

	ap.Statusbar = widgets.NewQStatusBar(ap.Window) // Statusbar
	ap.Window.SetStatusBar(ap.Statusbar)

	var sp = widgets.NewQSplitter(nil)

	var FsTree = NewFsTree()

	var FsList = NewFsList()

	sp.AddWidget(FsTree)
	sp.AddWidget(FsList)

	ap.Window.SetCentralWidget(sp)

	ap.Statusbar.ShowMessage("Application Path: "+core.QCoreApplication_ApplicationDirPath(), 0)

	widgets.QApplication_SetStyle2("fusion")
	ap.Window.Show()
	widgets.QApplication_Exec()
}

func NewFsList() *widgets.QListView {

	var drive string = "C:/"

	FsFileModel = widgets.NewQFileSystemModel(nil)

	FsFileModel.SetFilter(core.QDir__NoDotAndDotDot |
		core.QDir__Files)

	FsFileModel.SetRootPath(drive)

	FsListView = widgets.NewQListView(nil)
	FsListView.SetModel(FsFileModel)

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

	//FsTreeView.HideColumn(1)
	//FsTreeView.HideColumn(2)
	//FsTreeView.HideColumn(3)

	FsTreeView.ConnectCurrentChanged(treeViewCurrentChanged)

	return FsTreeView
}

func treeViewCurrentChanged(current *core.QModelIndex, previous *core.QModelIndex) {

	FilePath := FsDirModel.FilePath(current)

	FsListView.SetRootIndex(FsFileModel.SetRootPath(FilePath))

	FsTreeView.ScrollTo(current, 0) // 0 = EnsureVisible, 3 = PositionAtCenter

}
