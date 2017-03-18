package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/network"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/webchannel"
	"github.com/therecipe/qt/websockets"
	"github.com/therecipe/qt/widgets"

	"github.com/therecipe/qt/internal/examples/webchannel/shared"
)

type Dialog struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(text string) `signal:"sendText"`

	_ func(text string) `slot:"receiveText"`
	_ func()            `slot:"clicked"`

	ui        *widgets.QDialog
	ui_send   *widgets.QPushButton
	ui_output *widgets.QPlainTextEdit
	ui_input  *widgets.QLineEdit
}

func (d *Dialog) init() {

	loader := uitools.NewQUiLoader(nil)
	file := core.NewQFile2(":/dialog.ui")

	file.Open(core.QIODevice__ReadOnly)
	d.ui = widgets.NewQDialogFromPointer(loader.Load(file, nil).Pointer())
	file.Close()

	d.ui_send = widgets.NewQPushButtonFromPointer(d.ui.FindChild("send", core.Qt__FindChildrenRecursively).Pointer())
	d.ui_output = widgets.NewQPlainTextEditFromPointer(d.ui.FindChild("output", core.Qt__FindChildrenRecursively).Pointer())
	d.ui_input = widgets.NewQLineEditFromPointer(d.ui.FindChild("input", core.Qt__FindChildrenRecursively).Pointer())

	d.ui.Show()

	d.ConnectReceiveText(d.receiveText)
	d.ui_send.ConnectClicked(func(checked bool) { d.clicked() })
}

func (d *Dialog) displayMessage(message string) {
	d.ui_output.AppendPlainText(message)
}

func (d *Dialog) receiveText(text string) {
	d.displayMessage("Received message: " + text)
}

func (d *Dialog) clicked() {

	text := d.ui_input.Text()
	if text == "" {
		return
	}

	d.SendText(text)
	d.displayMessage("Sent message: " + text)

	d.ui_input.Clear()
}

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	jsFileInfo := core.NewQFileInfo3(core.QDir_TempPath() + "/qwebchannel.js")

	if !jsFileInfo.Exists() {
		core.QFile_Copy2(":/qwebchannel.js", jsFileInfo.AbsoluteFilePath())
	}

	server := websockets.NewQWebSocketServer("QWebChannel Standalone Example Server", websockets.QWebSocketServer__NonSecureMode, nil)
	if !server.Listen(network.NewQHostAddress9(network.QHostAddress__LocalHost), 12345) {
		panic("Failed to open web socket server")
	}

	clientWrapper := shared.NewWebSocketClientWrapper(nil)
	clientWrapper.Init(server)

	channel := webchannel.NewQWebChannel(nil)
	clientWrapper.ConnectClientConnected(func(client *shared.WebSocketTransport) {
		channel.ConnectTo(client.QWebChannelAbstractTransport_PTR())
	})

	dialog := NewDialog(nil)
	channel.RegisterObject("dialog", dialog)

	htmlFileInfo := core.NewQFileInfo3(core.QDir_TempPath() + "/index.html")

	if !htmlFileInfo.Exists() {
		core.QFile_Copy2(":/index.html", htmlFileInfo.AbsoluteFilePath())
	}

	url := core.QUrl_FromLocalFile(htmlFileInfo.AbsoluteFilePath())
	gui.QDesktopServices_OpenUrl(url)

	dialog.displayMessage("Initialization complete, opening browser at " + url.ToDisplayString(0))

	widgets.QApplication_Exec()
}
