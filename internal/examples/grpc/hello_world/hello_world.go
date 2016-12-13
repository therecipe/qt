package main

import (
	"encoding/json"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
)

//go:generate qtmoc
type HelloClientRPC struct {
	core.QObject

	_ func(string) string `slot:"sayHello"` //TODO: should be func(string) []*core.QVariant, once generic containers/lists are supported
	_ func() string       `slot:"shutdown"` //TODO: should be func() error, once "error" is mapped to QString
}

type HelloClientFactory struct {
	HelloClientRPC //TODO: only needed until a bug is fixed (should be core.QObject)

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
		if err != nil {
			return nil
		}

		ret.ConnectSayHello(func(s string) string {
			var out, err = client.SayHello(s)
			if err != nil {
				var rs, _ = json.Marshal(struct{ Error string }{err.Error()})
				return string(rs)
			}

			var rs, _ = json.Marshal(struct{ Data string }{out})
			return string(rs)
		})

		ret.ConnectShutdown(func() string {
			var err = client.Shutdown()
			if err != nil {
				return err.Error()
			}
			return ""
		})

		return ret
	})

	view.RootContext().SetContextProperty("HelloClientFactory", helloClientFactory)
	view.SetSource(core.NewQUrl3("qrc:///qml/hello_world.qml", 0))

	view.Show()

	gui.QGuiApplication_Exec()
}
