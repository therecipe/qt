package main

import (
	"os"

	"github.com/StarAurryon/qt/widgets"
)

type QSystemTrayIconWithCustomSlot struct {
	widgets.QSystemTrayIcon

	_ func(f func()) `slot:"triggerSlot,auto"` //create a slot that takes a function and automatically connect it
}

func (tray *QSystemTrayIconWithCustomSlot) triggerSlot(f func()) { f() } //the slot just needs to call the passed function to execute it inside the main thread

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	var (
		widget       = widgets.NewQWidget(nil, 0)
		widgetLayout = widgets.NewQVBoxLayout2(widget)
	)

	systray := NewQSystemTrayIconWithCustomSlot(nil)
	systray.SetIcon(widget.Style().StandardIcon(widgets.QStyle__SP_MessageBoxCritical, nil, nil))

	systrayMenu := widgets.NewQMenu(nil)
	systrayMenu.AddAction("first empty action")
	systrayMenu.AddAction("second empty action")
	systray.SetContextMenu(systrayMenu)

	systray.Show()

	//WORKS because ShowMessage is called from the main thread
	buttonMain := widgets.NewQPushButton2("call from main thread", nil)
	buttonMain.ConnectClicked(func(bool) {
		//in the main thread
		systray.ShowMessage("title", "main thread message", widgets.QSystemTrayIcon__Information, 5000)
	})
	widgetLayout.AddWidget(buttonMain, 0, 0)

	//FAILS because ShowMessage is called from the different thread
	buttonOther := widgets.NewQPushButton2("call from other thread", nil)
	buttonOther.ConnectClicked(func(bool) {
		//in the main thread
		go func() {
			//in a different thread
			systray.ShowMessage("this won't work and", "this message will never show up", widgets.QSystemTrayIcon__Information, 5000)
		}()
	})
	widgetLayout.AddWidget(buttonOther, 0, 0)

	//WORKS because ShowMessage is called from the main thread with the help of a custom slot
	buttonSlot := widgets.NewQPushButton2("call from other thread with slot", nil)
	buttonSlot.ConnectClicked(func(bool) {
		//in the main thread
		go func() {
			//in a different thread
			systray.TriggerSlot(func() {
				//in the main thread again
				systray.ShowMessage("title", "other thread message with slot", widgets.QSystemTrayIcon__Information, 5000)
			})
		}()
	})
	widgetLayout.AddWidget(buttonSlot, 0, 0)

	widget.Show()

	widgets.QApplication_Exec()
}
