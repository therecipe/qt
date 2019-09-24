//source: https://doc.qt.io/qt-5/qtcharts-modeldata-example.html

package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	NewTableWidget().Show()

	widgets.QApplication_Exec()
}
