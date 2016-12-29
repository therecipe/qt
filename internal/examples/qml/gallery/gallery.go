package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
)

func main() {

	core.QCoreApplication_SetApplicationName("Gallery")
	core.QCoreApplication_SetOrganizationName("QtProject")
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	/*
			TODO:
			var settings = core.NewQSettings5(nil)
			QString style = QQuickStyle::name();
			if (!style.isEmpty())
		    settings.setValue("style", style);
			else
		    QQuickStyle::setStyle(settings.value("style").toString());
	*/

	var engine = qml.NewQQmlApplicationEngine(nil)
	engine.Load(core.NewQUrl3("qrc:/gallery.qml", 0))
	if len(engine.RootObjects()) == 0 {
		return
	}

	gui.QGuiApplication_Exec()
}
