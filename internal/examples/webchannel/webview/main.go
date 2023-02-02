package main

import (
	"os"

	"github.com/bluszcz/cutego/core"
	"github.com/bluszcz/cutego/gui"
	"github.com/bluszcz/cutego/network"
	"github.com/bluszcz/cutego/qml"
	"github.com/bluszcz/cutego/webchannel"
	"github.com/bluszcz/cutego/websockets"
	"github.com/bluszcz/cutego/webview"

	"github.com/bluszcz/cutego/internal/examples/webchannel/shared"
)

type Dialog struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(text string) `signal:"sendText"`

	_ func(text string) `slot:"receiveText"`
}

func (d *Dialog) init() {
	d.ConnectReceiveText(d.receiveText)
}

func (d *Dialog) receiveText(text string) {
	println("received (from js):", text)

	answer := "hello"
	println("    send (from go):", answer)
	d.SendText(answer)
}

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)
	webview.QtWebView_Initialize()

	server := websockets.NewQWebSocketServer2("QWebChannel Standalone Example Server", websockets.QWebSocketServer__NonSecureMode, nil)
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

	//

	jsFileInfo := core.NewQFileInfo3(core.QDir_TempPath() + "/qwebchannel.js")

	if !jsFileInfo.Exists() {
		core.QFile_Copy2(":/qwebchannel.js", jsFileInfo.AbsoluteFilePath())
	}

	htmlFileInfo := core.NewQFileInfo3(core.QDir_TempPath() + "/index.html")

	if !htmlFileInfo.Exists() {
		core.QFile_Copy2(":/index.html", htmlFileInfo.AbsoluteFilePath())
	}

	url := core.QUrl_FromLocalFile(htmlFileInfo.AbsoluteFilePath())

	var view = qml.NewQQmlApplicationEngine(nil)
	view.RootContext().SetContextProperty2("pathToIndex", core.NewQVariant1(url.ToString(0)))
	view.Load(core.NewQUrl3("qrc:/main.qml", 0))

	gui.QGuiApplication_Exec()
}
