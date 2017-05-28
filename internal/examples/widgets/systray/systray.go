package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

type QSystemTrayIconWithCustomSlot struct {
	widgets.QSystemTrayIcon

	_ func() `slot:"triggerSlot"`
}

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	var (
		widget       = widgets.NewQWidget(nil, 0)
		widgetLayout = widgets.NewQVBoxLayout()
	)
	widget.SetLayout(widgetLayout)

	var systray = NewQSystemTrayIconWithCustomSlot(nil)
	systray.SetIcon(widget.Style().StandardIcon(widgets.QStyle__SP_MessageBoxCritical, nil, nil))

	var systrayMenu = widgets.NewQMenu(nil)
	systrayMenu.AddAction("first action")
	systrayMenu.AddAction("second action")
	systray.SetContextMenu(systrayMenu)

	systray.Show()

	//works
	var buttonMain = widgets.NewQPushButton2("call from main thread", nil)
	buttonMain.ConnectClicked(func(bool) {
		systray.ShowMessage("title", "main thread message", widgets.QSystemTrayIcon__Information, 5000)
	})
	widgetLayout.AddWidget(buttonMain, 0, 0)

	//fails
	var buttonOther = widgets.NewQPushButton2("call from other thread", nil)
	buttonOther.ConnectClicked(func(bool) {
		go func() {
			systray.ShowMessage("title", "other thread message", widgets.QSystemTrayIcon__Information, 5000)
		}()
	})
	widgetLayout.AddWidget(buttonOther, 0, 0)

	//works
	var buttonSlot = widgets.NewQPushButton2("call from other thread with slot", nil)
	systray.ConnectTriggerSlot(func() {
		systray.ShowMessage("title", "other thread message with slot", widgets.QSystemTrayIcon__Information, 5000)
	})
	buttonSlot.ConnectClicked(func(bool) {
		go func() {
			systray.TriggerSlot()
		}()
	})
	widgetLayout.AddWidget(buttonSlot, 0, 0)

	widget.Show()

	widgets.QApplication_Exec()
}
