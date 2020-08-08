package main

import (
	"os"

	"github.com/StarAurryon/qt/core"
	"github.com/StarAurryon/qt/gui"
	"github.com/StarAurryon/qt/quick"
)

type HelloClientRPC struct {
	core.QObject

	_ func(string) []*core.QVariant `slot:"sayHello"`
	_ func() error                  `slot:"shutdown"`
}

type HelloClientFactory struct {
	core.QObject

	_ func(addr string) *HelloClientRPC `slot:"newClient"`
}

func main() {
	go serve()

	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

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
		if err != nil {
			return nil
		}

		ret.ConnectSayHello(func(s string) []*core.QVariant {
			var out, err = client.SayHello(s)
			if err != nil {
				return []*core.QVariant{core.NewQVariant1(out), core.NewQVariant1(err.Error())}
			}
			return []*core.QVariant{core.NewQVariant1(out)}
		})

		ret.ConnectShutdown(func() error {
			var err = client.Shutdown()
			if err != nil {
				return err
			}
			return nil
		})

		return ret
	})

	view.RootContext().SetContextProperty("HelloClientFactory", helloClientFactory)
	view.SetSource(core.NewQUrl3("qrc:/qml/hello_world.qml", 0))

	view.Show()

	gui.QGuiApplication_Exec()
}
