package main

import (
	"os"
	"path/filepath"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	_ "github.com/therecipe/qt/svg" //QSvgWidget <- TODO: this comment is needed for js
	"github.com/therecipe/qt/widgets"

	_ "github.com/therecipe/qt/internal/examples/showcases/sia/assets"

	_ "github.com/therecipe/qt/internal/examples/showcases/sia/dashboard"
	_ "github.com/therecipe/qt/internal/examples/showcases/sia/files"
	fcontroller "github.com/therecipe/qt/internal/examples/showcases/sia/files/controller"
	_ "github.com/therecipe/qt/internal/examples/showcases/sia/hosting"
	_ "github.com/therecipe/qt/internal/examples/showcases/sia/terminal"
	tcontroller "github.com/therecipe/qt/internal/examples/showcases/sia/terminal/controller"
	_ "github.com/therecipe/qt/internal/examples/showcases/sia/theme"
	_ "github.com/therecipe/qt/internal/examples/showcases/sia/view"
	_ "github.com/therecipe/qt/internal/examples/showcases/sia/view/left"
	_ "github.com/therecipe/qt/internal/examples/showcases/sia/view/top"
	_ "github.com/therecipe/qt/internal/examples/showcases/sia/wallet"
	wcontroller "github.com/therecipe/qt/internal/examples/showcases/sia/wallet/controller"
	_ "github.com/therecipe/qt/internal/examples/showcases/sia/wallet/dialog"

	"github.com/therecipe/qt/internal/examples/showcases/sia/controller"
)

var (
	pathToProject = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "therecipe", "qt", "internal", "examples", "showcases", "sia")
	pathToSiac    = filepath.Join(os.Getenv("GOPATH"), "bin", "sia", "siac")

	PRODUCTION = true //set to 'true' to use qrc: instead of qml files
	DEMO       = true //set to 'true' to use demo data for the wallet and files table
)

func init() {
	tcontroller.PathToSiac = pathToSiac

	if !PRODUCTION {
		os.Setenv("QML_DISABLE_DISK_CACHE", "true")
	}

	controller.DEMO = DEMO
	wcontroller.DEMO = DEMO
	fcontroller.DEMO = DEMO
}

func main() {
	path := filepath.Join(pathToProject, "view", "qml", "View.qml")

	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	widgets.NewQApplication(len(os.Args), os.Args)

	controller.NewController(nil)

	view := quick.NewQQuickView(nil)
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.SetMinimumSize(core.NewQSize2(1024, 384))
	view.Resize(core.NewQSize2(1024, 768))

	if PRODUCTION {
		view.Engine().AddImportPath("qrc:/qml/")
		view.SetSource(core.NewQUrl3("qrc:/qml/View.qml", 0))
	} else {
		view.Engine().AddImportPath(filepath.Join(pathToProject, "theme", "qml"))

		view.Engine().AddImportPath(filepath.Join(pathToProject, "view", "qml"))
		view.Engine().AddImportPath(filepath.Join(pathToProject, "view", "top", "qml"))
		view.Engine().AddImportPath(filepath.Join(pathToProject, "view", "left", "qml"))

		view.Engine().AddImportPath(filepath.Join(pathToProject, "dashboard", "qml"))
		view.Engine().AddImportPath(filepath.Join(pathToProject, "files", "qml"))
		view.Engine().AddImportPath(filepath.Join(pathToProject, "hosting", "qml"))

		view.Engine().AddImportPath(filepath.Join(pathToProject, "wallet", "qml"))
		view.Engine().AddImportPath(filepath.Join(pathToProject, "wallet", "dialog", "qml"))

		view.Engine().AddImportPath(filepath.Join(pathToProject, "terminal", "qml"))
		view.SetSource(core.QUrl_FromLocalFile(path))
	}

	view.Show()

	widgets.QApplication_Exec()
}
