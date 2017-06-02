package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
)

type HelloClientRPC struct {
	core.QObject

	_ func(string) `slot:"sayHello"`
	_ func()       `slot:"shutdown"`
}

type HelloClientFactory struct {
	core.QObject

	_ func(msg string)      `signal:"info"`
	_ func(err, msg string) `signal:"error"`
	_ func(msg string)      `signal:"success"`

	_ func(addr string) *HelloClientRPC `slot:"newClient"`
}

func main() {
	go serve()

	gui.NewQGuiApplication(len(os.Args), os.Args)

	var view = quick.NewQQuickView(nil)
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)

	var helloClientFactory = NewHelloClientFactory(nil)
	HelloClientRPC_QRegisterMetaType() //needed to register HelloClientRPC* as a return struct for QML
	helloClientFactory.ConnectNewClient(func(addr string) *HelloClientRPC {
		var (
			client, err = New(addr)
			ret         = NewHelloClientRPC(nil)
		)
		helloClientFactory.Info("Trying " + addr + "\n")
		if err != nil {
			helloClientFactory.Error("create client", err.Error())
			return nil
		}

		ret.ConnectSayHello(func(s string) {
			var out, err = client.SayHello(s)
			if err != nil {
				helloClientFactory.Error("receive response", err.Error())
				return
			}
			helloClientFactory.Success(string(out))
		})

		ret.ConnectShutdown(func() {
			var err = client.Shutdown()
			if err != nil {
				helloClientFactory.Error("shutdown client", err.Error())
			}
		})

		return ret
	})

	view.RootContext().SetContextProperty("HelloClientFactory", helloClientFactory)
	view.SetSource(core.NewQUrl3("qrc:/qml/hello_world2.qml", 0))

	view.Show()

	gui.QGuiApplication_Exec()
}
