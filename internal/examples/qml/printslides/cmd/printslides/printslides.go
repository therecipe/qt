//source: https://github.com/qt-labs/qml-presentation-system

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
)

func main() {
	app := gui.NewQGuiApplication(len(os.Args), os.Args)
	app.SetOrganizationDomain("com")
	app.SetOrganizationName("ics")
	app.SetApplicationName("printslides")
	app.SetApplicationVersion("0.2")
	app.SetApplicationDisplayName("QML Presentation Slide Printer")
	mainView := NewSlideView(nil)
	if len(os.Args) != 2 || strings.HasSuffix(os.Args[1], "?") {
		progName := app.ApplicationName()
		fmt.Printf("%[1]s usage:\t%[1]s path/to/presentation.qml\n", progName)
		os.Exit(2)
	}
	mainView.SetSource(core.QUrl_FromLocalFile(os.Args[1]))
	mainView.Show()
	app.Exec()
}
