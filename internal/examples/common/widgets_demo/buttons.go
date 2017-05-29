package main

import (
	"fmt"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func buttons() {

	//Push Button
	pushButton := widgets.NewQPushButton2("Push Button", nil)
	pushButton.SetWindowTitle("Push Button")

	pushButton.ConnectClicked(func(checked bool) {
		println("You Clicked The Push Button")
	})
	addWidget(pushButton)

	//Tool Button (+ ToolBar)
	toolWindow := widgets.NewQMainWindow(nil, 0)
	toolWindow.SetWindowTitle("Tool Button")

	toolBar := toolWindow.AddToolBar3("Just Some ToolBar")
	toolBar.SetAllowedAreas(core.Qt__AllToolBarAreas)
	toolBar.SetMovable(true)
	toolBar.SetFloatable(true)

	toolButton := widgets.NewQToolButton(nil)
	toolButton.SetIcon(toolButton.Style().StandardIcon(widgets.QStyle__SP_DialogHelpButton, nil, nil))
	toolButton.ConnectClicked(func(checked bool) {
		println("You Clicked The Tool Button")
		toolBar.AddSeparator()
	})
	toolBar.AddWidget(toolButton)
	addWidget(toolWindow)

	//Radio Button
	radioButtonGroup := widgets.NewQWidget(nil, 0)
	radioButtonGroup.SetWindowTitle("Radio Button")
	radioButtonGroupLayout := widgets.NewQVBoxLayout2(radioButtonGroup)

	for i := 0; i < 3; i++ {
		radioButton := widgets.NewQRadioButton2(fmt.Sprintf("Radio Button %v", i), nil)
		radioButton.ConnectClicked(func(checked bool) {
			println("You Clicked The " + radioButton.Text())
		})
		radioButtonGroupLayout.AddWidget(radioButton, 0, 0)
	}
	addWidget(radioButtonGroup)

	//Check Box
	checkBox := widgets.NewQCheckBox2("Check Box", nil)
	checkBox.SetWindowTitle("Check Box")

	checkBox.ConnectClicked(func(checked bool) {
		println("You Clicked The Check Box :", checked)
	})
	addWidget(checkBox)

	//Command Link Button
	commandLinkButton := widgets.NewQCommandLinkButton2("Command Link Button", nil)
	commandLinkButton.SetWindowTitle("Command Link Button")

	commandLinkButton.ConnectClicked(func(checked bool) {
		println("You Clicked The Command Link Button")
	})
	addWidget(commandLinkButton)

	//Dialog Button Box
	dialogButtonBox := widgets.NewQDialogButtonBox3(widgets.QDialogButtonBox__Cancel|widgets.QDialogButtonBox__Ok, nil)
	dialogButtonBox.SetWindowTitle("Dialog Button Box")

	dialogButtonBox.ConnectAccepted(func() {
		println("You Clicked The Accept Button")
	})
	dialogButtonBox.ConnectRejected(func() {
		println("You Clicked The Reject Button")
	})
	addWidget(dialogButtonBox)

}
