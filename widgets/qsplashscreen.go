package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QSplashScreen struct {
	QWidget
}

type QSplashScreen_ITF interface {
	QWidget_ITF
	QSplashScreen_PTR() *QSplashScreen
}

func PointerFromQSplashScreen(ptr QSplashScreen_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSplashScreen_PTR().Pointer()
	}
	return nil
}

func NewQSplashScreenFromPointer(ptr unsafe.Pointer) *QSplashScreen {
	var n = new(QSplashScreen)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSplashScreen_") {
		n.SetObjectName("QSplashScreen_" + qt.Identifier())
	}
	return n
}

func (ptr *QSplashScreen) QSplashScreen_PTR() *QSplashScreen {
	return ptr
}

func NewQSplashScreen2(parent QWidget_ITF, pixmap gui.QPixmap_ITF, f core.Qt__WindowType) *QSplashScreen {
	defer qt.Recovering("QSplashScreen::QSplashScreen")

	return NewQSplashScreenFromPointer(C.QSplashScreen_NewQSplashScreen2(PointerFromQWidget(parent), gui.PointerFromQPixmap(pixmap), C.int(f)))
}

func NewQSplashScreen(pixmap gui.QPixmap_ITF, f core.Qt__WindowType) *QSplashScreen {
	defer qt.Recovering("QSplashScreen::QSplashScreen")

	return NewQSplashScreenFromPointer(C.QSplashScreen_NewQSplashScreen(gui.PointerFromQPixmap(pixmap), C.int(f)))
}

func (ptr *QSplashScreen) ClearMessage() {
	defer qt.Recovering("QSplashScreen::clearMessage")

	if ptr.Pointer() != nil {
		C.QSplashScreen_ClearMessage(ptr.Pointer())
	}
}

func (ptr *QSplashScreen) ConnectDrawContents(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QSplashScreen::drawContents")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "drawContents", f)
	}
}

func (ptr *QSplashScreen) DisconnectDrawContents() {
	defer qt.Recovering("disconnect QSplashScreen::drawContents")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "drawContents")
	}
}

//export callbackQSplashScreenDrawContents
func callbackQSplashScreenDrawContents(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplashScreen::drawContents")

	if signal := qt.GetSignal(C.GoString(ptrName), "drawContents"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QSplashScreen) Finish(mainWin QWidget_ITF) {
	defer qt.Recovering("QSplashScreen::finish")

	if ptr.Pointer() != nil {
		C.QSplashScreen_Finish(ptr.Pointer(), PointerFromQWidget(mainWin))
	}
}

func (ptr *QSplashScreen) Message() string {
	defer qt.Recovering("QSplashScreen::message")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSplashScreen_Message(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSplashScreen) ConnectMessageChanged(f func(message string)) {
	defer qt.Recovering("connect QSplashScreen::messageChanged")

	if ptr.Pointer() != nil {
		C.QSplashScreen_ConnectMessageChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "messageChanged", f)
	}
}

func (ptr *QSplashScreen) DisconnectMessageChanged() {
	defer qt.Recovering("disconnect QSplashScreen::messageChanged")

	if ptr.Pointer() != nil {
		C.QSplashScreen_DisconnectMessageChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "messageChanged")
	}
}

//export callbackQSplashScreenMessageChanged
func callbackQSplashScreenMessageChanged(ptrName *C.char, message *C.char) {
	defer qt.Recovering("callback QSplashScreen::messageChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "messageChanged"); signal != nil {
		signal.(func(string))(C.GoString(message))
	}

}

func (ptr *QSplashScreen) ConnectMousePressEvent(f func(v *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSplashScreen::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QSplashScreen) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QSplashScreen::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQSplashScreenMousePressEvent
func callbackQSplashScreenMousePressEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplashScreen::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QSplashScreen) Repaint() {
	defer qt.Recovering("QSplashScreen::repaint")

	if ptr.Pointer() != nil {
		C.QSplashScreen_Repaint(ptr.Pointer())
	}
}

func (ptr *QSplashScreen) SetPixmap(pixmap gui.QPixmap_ITF) {
	defer qt.Recovering("QSplashScreen::setPixmap")

	if ptr.Pointer() != nil {
		C.QSplashScreen_SetPixmap(ptr.Pointer(), gui.PointerFromQPixmap(pixmap))
	}
}

func (ptr *QSplashScreen) ShowMessage(message string, alignment int, color gui.QColor_ITF) {
	defer qt.Recovering("QSplashScreen::showMessage")

	if ptr.Pointer() != nil {
		C.QSplashScreen_ShowMessage(ptr.Pointer(), C.CString(message), C.int(alignment), gui.PointerFromQColor(color))
	}
}

func (ptr *QSplashScreen) DestroyQSplashScreen() {
	defer qt.Recovering("QSplashScreen::~QSplashScreen")

	if ptr.Pointer() != nil {
		C.QSplashScreen_DestroyQSplashScreen(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
