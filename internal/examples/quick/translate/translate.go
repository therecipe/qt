package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
)

func main() {

	gui.NewQGuiApplication(len(os.Args), os.Args)

	var translator = core.NewQTranslator(nil)
	translator.Load2(core.NewQLocale3(core.QLocale__French, core.QLocale__France), "translation", "_", ":/qml/i18n", ".qm")
	core.QCoreApplication_InstallTranslator(translator)

	var view = quick.NewQQuickView(nil)
	view.SetSource(core.NewQUrl3("qrc:/qml/translate.qml", 0))
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.Show()

	gui.QGuiApplication_Exec()
}
