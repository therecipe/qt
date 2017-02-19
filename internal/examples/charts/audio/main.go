//source: http://doc.qt.io/qt-5/qtcharts-audio-example.html
//TODO: leaks memory, probably *_newList

package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	var app = widgets.NewQApplication(len(os.Args), os.Args)

	var w = NewWidget(nil, 0)
	w.Show()

	app.Exec()
}
