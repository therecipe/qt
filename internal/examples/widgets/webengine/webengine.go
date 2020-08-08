package main

import (
	"os"

	"github.com/StarAurryon/qt/core"
	"github.com/StarAurryon/qt/webengine"
	"github.com/StarAurryon/qt/widgets"
)

const htmlData = `<!DOCTYPE html>
<html>
<body>

<div id="div1">
<p id="p1">This is a paragraph.</p>
<p id="p2">This is another paragraph.</p>
</div>

</body>
</html>`

const jsData = `var para = document.createElement("p");
var node = document.createTextNode("This is new.");
para.appendChild(node);

var element = document.getElementById("div1");
element.appendChild(para);`

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	var window = widgets.NewQMainWindow(nil, 0)

	var centralWidget = widgets.NewQWidget(nil, 0)
	centralWidget.SetLayout(widgets.NewQVBoxLayout())

	var view = webengine.NewQWebEngineView(nil)
	view.SetHtml(htmlData, core.NewQUrl())
	centralWidget.Layout().AddWidget(view)

	var button = widgets.NewQPushButton2("click me", nil)
	button.ConnectClicked(func(checked bool) {
		view.Page().RunJavaScript(jsData)
	})
	centralWidget.Layout().AddWidget(button)

	window.SetCentralWidget(centralWidget)
	window.Show()

	widgets.QApplication_Exec()
}
