package main

import (
	"fmt"
	"os"
	"time"

	"github.com/therecipe/qt/androidextras"
	"github.com/therecipe/qt/widgets"
)

func main() {
	if os.Args[len(os.Args)-1] == "-service" {
		fmt.Println("ARGS SERVICE:", os.Args)

		s := androidextras.NewQAndroidService(len(os.Args), os.Args)

		go func() {
			t := time.Now()
			for range time.Tick(1 * time.Second) {
				println("serving for", time.Since(t).String())
			}
		}()

		s.Exec()
		return
	}

	fmt.Println("ARGS ACTIVITY:", os.Args)

	widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)

	pb := widgets.NewQPushButton2("start service", nil)
	pb.ConnectClicked(func(bool) {
		androidextras.QAndroidJniObject_CallStaticMethodVoid2("org/qtproject/example/MyService", "startMyService", "(Landroid/content/Context;)V", androidextras.QtAndroid_AndroidActivity().Object())
	})
	window.SetCentralWidget(pb)

	window.ShowMaximized()

	widgets.QApplication_Exec()
}
