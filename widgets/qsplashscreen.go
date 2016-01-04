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
func callbackQSplashScreenDrawContents(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::drawContents")

	if signal := qt.GetSignal(C.GoString(ptrName), "drawContents"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQSplashScreenFromPointer(ptr).DrawContentsDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QSplashScreen) DrawContents(painter gui.QPainter_ITF) {
	defer qt.Recovering("QSplashScreen::drawContents")

	if ptr.Pointer() != nil {
		C.QSplashScreen_DrawContents(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QSplashScreen) DrawContentsDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QSplashScreen::drawContents")

	if ptr.Pointer() != nil {
		C.QSplashScreen_DrawContentsDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QSplashScreen) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSplashScreen::event")

	if ptr.Pointer() != nil {
		return C.QSplashScreen_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
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
func callbackQSplashScreenMessageChanged(ptr unsafe.Pointer, ptrName *C.char, message *C.char) {
	defer qt.Recovering("callback QSplashScreen::messageChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "messageChanged"); signal != nil {
		signal.(func(string))(C.GoString(message))
	}

}

func (ptr *QSplashScreen) MessageChanged(message string) {
	defer qt.Recovering("QSplashScreen::messageChanged")

	if ptr.Pointer() != nil {
		C.QSplashScreen_MessageChanged(ptr.Pointer(), C.CString(message))
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
func callbackQSplashScreenMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(v))
	} else {
		NewQSplashScreenFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(v))
	}
}

func (ptr *QSplashScreen) MousePressEvent(v gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSplashScreen::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(v))
	}
}

func (ptr *QSplashScreen) MousePressEventDefault(v gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSplashScreen::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(v))
	}
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

func (ptr *QSplashScreen) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QSplashScreen::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QSplashScreen) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QSplashScreen::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQSplashScreenActionEvent
func callbackQSplashScreenActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQSplashScreenFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QSplashScreen) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QSplashScreen::actionEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QSplashScreen) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QSplashScreen::actionEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QSplashScreen) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QSplashScreen::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QSplashScreen) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QSplashScreen::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQSplashScreenDragEnterEvent
func callbackQSplashScreenDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQSplashScreenFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QSplashScreen) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QSplashScreen::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QSplashScreen) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QSplashScreen::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QSplashScreen) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QSplashScreen::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QSplashScreen) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QSplashScreen::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQSplashScreenDragLeaveEvent
func callbackQSplashScreenDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQSplashScreenFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QSplashScreen) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QSplashScreen::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QSplashScreen) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QSplashScreen::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QSplashScreen) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QSplashScreen::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QSplashScreen) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QSplashScreen::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQSplashScreenDragMoveEvent
func callbackQSplashScreenDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQSplashScreenFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QSplashScreen) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QSplashScreen::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QSplashScreen) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QSplashScreen::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QSplashScreen) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QSplashScreen::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QSplashScreen) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QSplashScreen::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQSplashScreenDropEvent
func callbackQSplashScreenDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQSplashScreenFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QSplashScreen) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QSplashScreen::dropEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QSplashScreen) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QSplashScreen::dropEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QSplashScreen) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSplashScreen::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QSplashScreen) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QSplashScreen::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQSplashScreenEnterEvent
func callbackQSplashScreenEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSplashScreenFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSplashScreen) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSplashScreen::enterEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSplashScreen) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSplashScreen::enterEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSplashScreen) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QSplashScreen::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QSplashScreen) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QSplashScreen::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQSplashScreenFocusInEvent
func callbackQSplashScreenFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQSplashScreenFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QSplashScreen) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSplashScreen::focusInEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QSplashScreen) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSplashScreen::focusInEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QSplashScreen) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QSplashScreen::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QSplashScreen) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QSplashScreen::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQSplashScreenFocusOutEvent
func callbackQSplashScreenFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQSplashScreenFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QSplashScreen) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSplashScreen::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QSplashScreen) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSplashScreen::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QSplashScreen) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QSplashScreen::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QSplashScreen) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QSplashScreen::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQSplashScreenHideEvent
func callbackQSplashScreenHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQSplashScreenFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QSplashScreen) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QSplashScreen::hideEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QSplashScreen) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QSplashScreen::hideEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QSplashScreen) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSplashScreen::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QSplashScreen) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QSplashScreen::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQSplashScreenLeaveEvent
func callbackQSplashScreenLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSplashScreenFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSplashScreen) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSplashScreen::leaveEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSplashScreen) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSplashScreen::leaveEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSplashScreen) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QSplashScreen::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QSplashScreen) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QSplashScreen::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQSplashScreenMoveEvent
func callbackQSplashScreenMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQSplashScreenFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QSplashScreen) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QSplashScreen::moveEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QSplashScreen) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QSplashScreen::moveEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QSplashScreen) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QSplashScreen::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QSplashScreen) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QSplashScreen::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQSplashScreenPaintEvent
func callbackQSplashScreenPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQSplashScreenFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QSplashScreen) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QSplashScreen::paintEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QSplashScreen) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QSplashScreen::paintEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QSplashScreen) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QSplashScreen::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QSplashScreen) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QSplashScreen::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQSplashScreenSetVisible
func callbackQSplashScreenSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QSplashScreen::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QSplashScreen) SetVisible(visible bool) {
	defer qt.Recovering("QSplashScreen::setVisible")

	if ptr.Pointer() != nil {
		C.QSplashScreen_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QSplashScreen) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QSplashScreen::setVisible")

	if ptr.Pointer() != nil {
		C.QSplashScreen_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QSplashScreen) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QSplashScreen::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QSplashScreen) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QSplashScreen::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQSplashScreenShowEvent
func callbackQSplashScreenShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQSplashScreenFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QSplashScreen) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QSplashScreen::showEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QSplashScreen) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QSplashScreen::showEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QSplashScreen) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSplashScreen::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QSplashScreen) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QSplashScreen::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQSplashScreenChangeEvent
func callbackQSplashScreenChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSplashScreenFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSplashScreen) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSplashScreen::changeEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSplashScreen) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSplashScreen::changeEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSplashScreen) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QSplashScreen::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QSplashScreen) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QSplashScreen::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQSplashScreenCloseEvent
func callbackQSplashScreenCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQSplashScreenFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QSplashScreen) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QSplashScreen::closeEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QSplashScreen) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QSplashScreen::closeEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QSplashScreen) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QSplashScreen::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QSplashScreen) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QSplashScreen::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQSplashScreenContextMenuEvent
func callbackQSplashScreenContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQSplashScreenFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QSplashScreen) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QSplashScreen::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QSplashScreen) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QSplashScreen::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QSplashScreen) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QSplashScreen::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QSplashScreen) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QSplashScreen::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQSplashScreenInitPainter
func callbackQSplashScreenInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQSplashScreenFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QSplashScreen) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QSplashScreen::initPainter")

	if ptr.Pointer() != nil {
		C.QSplashScreen_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QSplashScreen) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QSplashScreen::initPainter")

	if ptr.Pointer() != nil {
		C.QSplashScreen_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QSplashScreen) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QSplashScreen::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QSplashScreen) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QSplashScreen::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQSplashScreenInputMethodEvent
func callbackQSplashScreenInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQSplashScreenFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QSplashScreen) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QSplashScreen::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QSplashScreen) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QSplashScreen::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QSplashScreen) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QSplashScreen::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QSplashScreen) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QSplashScreen::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQSplashScreenKeyPressEvent
func callbackQSplashScreenKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQSplashScreenFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QSplashScreen) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSplashScreen::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QSplashScreen) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSplashScreen::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QSplashScreen) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QSplashScreen::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QSplashScreen) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QSplashScreen::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQSplashScreenKeyReleaseEvent
func callbackQSplashScreenKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQSplashScreenFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QSplashScreen) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSplashScreen::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QSplashScreen) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSplashScreen::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QSplashScreen) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSplashScreen::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QSplashScreen) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QSplashScreen::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQSplashScreenMouseDoubleClickEvent
func callbackQSplashScreenMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQSplashScreenFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QSplashScreen) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSplashScreen::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSplashScreen) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSplashScreen::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSplashScreen) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSplashScreen::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QSplashScreen) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QSplashScreen::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQSplashScreenMouseMoveEvent
func callbackQSplashScreenMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQSplashScreenFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QSplashScreen) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSplashScreen::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSplashScreen) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSplashScreen::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSplashScreen) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSplashScreen::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QSplashScreen) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QSplashScreen::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQSplashScreenMouseReleaseEvent
func callbackQSplashScreenMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQSplashScreenFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QSplashScreen) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSplashScreen::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSplashScreen) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSplashScreen::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSplashScreen) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QSplashScreen::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QSplashScreen) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QSplashScreen::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQSplashScreenResizeEvent
func callbackQSplashScreenResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQSplashScreenFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QSplashScreen) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QSplashScreen::resizeEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QSplashScreen) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QSplashScreen::resizeEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QSplashScreen) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QSplashScreen::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QSplashScreen) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QSplashScreen::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQSplashScreenTabletEvent
func callbackQSplashScreenTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQSplashScreenFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QSplashScreen) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QSplashScreen::tabletEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QSplashScreen) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QSplashScreen::tabletEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QSplashScreen) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QSplashScreen::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QSplashScreen) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QSplashScreen::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQSplashScreenWheelEvent
func callbackQSplashScreenWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQSplashScreenFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QSplashScreen) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QSplashScreen::wheelEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QSplashScreen) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QSplashScreen::wheelEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QSplashScreen) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSplashScreen::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSplashScreen) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSplashScreen::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSplashScreenTimerEvent
func callbackQSplashScreenTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSplashScreenFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSplashScreen) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSplashScreen::timerEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSplashScreen) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSplashScreen::timerEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSplashScreen) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSplashScreen::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSplashScreen) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSplashScreen::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSplashScreenChildEvent
func callbackQSplashScreenChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSplashScreenFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSplashScreen) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSplashScreen::childEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSplashScreen) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSplashScreen::childEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSplashScreen) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSplashScreen::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSplashScreen) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSplashScreen::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSplashScreenCustomEvent
func callbackQSplashScreenCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplashScreen::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSplashScreenFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSplashScreen) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSplashScreen::customEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSplashScreen) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSplashScreen::customEvent")

	if ptr.Pointer() != nil {
		C.QSplashScreen_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
