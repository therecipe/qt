//source: https://doc.qt.io/qt-5/qtcharts-audio-example.html

package main

import (
	"os"

	"github.com/akiyosi/qt/widgets"
)

func main() {
	var app = widgets.NewQApplication(len(os.Args), os.Args)

	var w = NewWidget(nil, 0)
	w.Show()

	app.Exec()
}
