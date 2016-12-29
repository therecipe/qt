//source: https://github.com/ekke/drawer_nav_x

package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	os.Setenv("QT_QUICK_CONTROLS_STYLE", "material")
	gui.NewQGuiApplication(len(os.Args), os.Args)

	var translator = core.NewQTranslator(nil)
	if translator.Load2(core.NewQLocale(), "drawer_nav_x", "_", ":/translations", ".qm") {
		core.QCoreApplication_InstallTranslator(translator)
	} else {
		println("cannot load translator", core.QLocale_System().Name(), "check content of translations.qrc")
	}

	var appui = initApplicationUI()
	var engine = qml.NewQQmlApplicationEngine(nil)

	// from QML we have access to ApplicationUI as myApp
	var context = engine.RootContext()
	context.SetContextProperty("myApp", appui)

	engine.Load(core.NewQUrl3("qrc:/main.qml", 0))

	gui.QGuiApplication_Exec()
}
