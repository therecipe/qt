package controller

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

var colorDialog *widgets.QMainWindow

func initColorDialog() {
	colorDialog = widgets.NewQMainWindow(nil, 0)
	colorDialog.SetWindowTitle("Dev Color Dialog")

	centralWidget := widgets.NewQWidget(nil, 0)
	formLayout := widgets.NewQFormLayout(nil)

	for _, n := range []string{
		"accent",
		"background", "darkBackground",
		"walletTableHeader", "walletTableAlternate", "walletTableHighlight",
		"inputFieldBackground", "font", "fontHighlight",
	} {
		formLayout.AddRow3(n, newColorPicker(n))
	}

	centralWidget.SetLayout(formLayout)
	colorDialog.SetCentralWidget(centralWidget)
}

func newColorPicker(n string) *widgets.QPushButton {

	stdcolor := gui.NewQColor6(getCurrentColorFor(n))

	pixmap := gui.NewQPixmap2(core.NewQSize2(36, 36))
	pixmap.Fill(stdcolor)

	button := widgets.NewQPushButton(nil)
	button.SetIcon(gui.NewQIcon2(pixmap))

	dialog := widgets.NewQColorDialog2(stdcolor, nil)
	dialog.SetOption(widgets.QColorDialog__NoButtons, true)
	button.ConnectClicked(func(bool) {
		Controller.Hide()
		dialog.SetCurrentColor(gui.NewQColor6(getCurrentColorFor(n)))
		dialog.Show()
	})

	dialog.ConnectCurrentColorChanged(func(c *gui.QColor) {
		pixmap := gui.NewQPixmap2(core.NewQSize2(36, 36))
		pixmap.Fill(c)
		button.SetIcon(gui.NewQIcon2(pixmap))

		color := c.Name2(gui.QColor__HexRgb)
		switch n {
		case "accent":
			Controller.SetAccent(color)

		case "nextAccent":
			Controller.SetNextAccent(color)

		case "background":
			Controller.SetBackground(color)

		case "darkBackground":
			Controller.SetDarkBackground(color)

		case "walletTableHeader":
			Controller.SetWalletTableHeader(color)

		case "walletTableAlternate":
			Controller.SetWalletTableAlternate(color)

		case "walletTableHighlight":
			Controller.SetWalletTableHighlight(color)

		case "inputFieldBackground":
			Controller.SetInputFieldBackground(color)

		case "font":
			Controller.SetFont(color)

		case "fontHighlight":
			Controller.SetFontHighlight(color)
		}
	})

	Controller.ConnectHide(dialog.Hide)

	return button
}

func getCurrentColorFor(n string) string {
	switch n {
	case "accent":
		return Controller.Accent()

	case "nextAccent":
		return Controller.NextAccent()

	case "background":
		return Controller.Background()

	case "darkBackground":
		return Controller.DarkBackground()

	case "walletTableHeader":
		return Controller.WalletTableHeader()

	case "walletTableAlternate":
		return Controller.WalletTableAlternate()

	case "walletTableHighlight":
		return Controller.WalletTableHighlight()

	case "inputFieldBackground":
		return Controller.InputFieldBackground()

	case "font":
		return Controller.Font()

	case "fontHighlight":
		return Controller.FontHighlight()

	default:
		return ""
	}
}
