package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)

func init() { Embedded_QmlRegisterType2("CustomModule", 1, 0, "Embedded") }

type Embedded struct {
	quick.QQuickPaintedItem

	_ func() `constructor:"init"`
}

func (e *Embedded) init() {
	e.SetAcceptedMouseButtons(core.Qt__AllButtons)

	input := widgets.NewQLineEdit(nil)
	input.SetPlaceholderText("Write something ...")

	e.SetImplicitHeight(float64(input.SizeHint().Height()))
	e.SetImplicitWidth(float64(input.SizeHint().Width()))

	e.ConnectEventFilter(func(watched *core.QObject, event *core.QEvent) bool {
		if watched.Pointer() == input.Pointer() {
			switch event.Type() {
			case core.QEvent__Paint, core.QEvent__UpdateRequest:
				e.UpdateDefault()
			}
		}
		return e.EventFilterDefault(watched, event)
	})

	e.ConnectPaint(func(painter *gui.QPainter) {
		input.Render2(painter, core.NewQPoint(), gui.NewQRegion(), widgets.QWidget__DrawWindowBackground|widgets.QWidget__DrawChildren)
	})

	e.ConnectEvent(func(event *core.QEvent) bool {
		switch event.Type() {
		case core.QEvent__MouseButtonPress,
			core.QEvent__MouseButtonRelease,
			core.QEvent__MouseButtonDblClick,
			core.QEvent__KeyPress,
			core.QEvent__KeyRelease,
			core.QEvent__FocusOut,
			core.QEvent__InputMethod,
			core.QEvent__InputMethodQuery:

			core.QCoreApplication_SendEvent(input, event)
			return true

		case core.QEvent__FocusIn:
			e.ForceActiveFocus()
		}

		return e.EventDefault(event)
	})

	input.InstallEventFilter(e)
}

func main() {
	os.Setenv("QT_IM_MODULE", "qtvirtualkeyboard")
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	widgets.NewQApplication(len(os.Args), os.Args)

	engine := qml.NewQQmlApplicationEngine(nil)
	engine.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))

	widgets.QApplication_Exec()
}
