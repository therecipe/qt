package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QToolButton struct {
	QAbstractButton
}

type QToolButton_ITF interface {
	QAbstractButton_ITF
	QToolButton_PTR() *QToolButton
}

func PointerFromQToolButton(ptr QToolButton_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QToolButton_PTR().Pointer()
	}
	return nil
}

func NewQToolButtonFromPointer(ptr unsafe.Pointer) *QToolButton {
	var n = new(QToolButton)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QToolButton_") {
		n.SetObjectName("QToolButton_" + qt.Identifier())
	}
	return n
}

func (ptr *QToolButton) QToolButton_PTR() *QToolButton {
	return ptr
}

//QToolButton::ToolButtonPopupMode
type QToolButton__ToolButtonPopupMode int64

const (
	QToolButton__DelayedPopup    = QToolButton__ToolButtonPopupMode(0)
	QToolButton__MenuButtonPopup = QToolButton__ToolButtonPopupMode(1)
	QToolButton__InstantPopup    = QToolButton__ToolButtonPopupMode(2)
)

func (ptr *QToolButton) ArrowType() core.Qt__ArrowType {
	defer qt.Recovering("QToolButton::arrowType")

	if ptr.Pointer() != nil {
		return core.Qt__ArrowType(C.QToolButton_ArrowType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QToolButton) AutoRaise() bool {
	defer qt.Recovering("QToolButton::autoRaise")

	if ptr.Pointer() != nil {
		return C.QToolButton_AutoRaise(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QToolButton) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QToolButton::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QToolButton) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QToolButton::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQToolButtonPaintEvent
func callbackQToolButtonPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::paintEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "paintEvent")
	if signal != nil {
		defer signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QToolButton) PopupMode() QToolButton__ToolButtonPopupMode {
	defer qt.Recovering("QToolButton::popupMode")

	if ptr.Pointer() != nil {
		return QToolButton__ToolButtonPopupMode(C.QToolButton_PopupMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QToolButton) SetArrowType(ty core.Qt__ArrowType) {
	defer qt.Recovering("QToolButton::setArrowType")

	if ptr.Pointer() != nil {
		C.QToolButton_SetArrowType(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QToolButton) SetAutoRaise(enable bool) {
	defer qt.Recovering("QToolButton::setAutoRaise")

	if ptr.Pointer() != nil {
		C.QToolButton_SetAutoRaise(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QToolButton) SetPopupMode(mode QToolButton__ToolButtonPopupMode) {
	defer qt.Recovering("QToolButton::setPopupMode")

	if ptr.Pointer() != nil {
		C.QToolButton_SetPopupMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QToolButton) SetToolButtonStyle(style core.Qt__ToolButtonStyle) {
	defer qt.Recovering("QToolButton::setToolButtonStyle")

	if ptr.Pointer() != nil {
		C.QToolButton_SetToolButtonStyle(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QToolButton) ToolButtonStyle() core.Qt__ToolButtonStyle {
	defer qt.Recovering("QToolButton::toolButtonStyle")

	if ptr.Pointer() != nil {
		return core.Qt__ToolButtonStyle(C.QToolButton_ToolButtonStyle(ptr.Pointer()))
	}
	return 0
}

func NewQToolButton(parent QWidget_ITF) *QToolButton {
	defer qt.Recovering("QToolButton::QToolButton")

	return NewQToolButtonFromPointer(C.QToolButton_NewQToolButton(PointerFromQWidget(parent)))
}

func (ptr *QToolButton) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QToolButton::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QToolButton) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QToolButton::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQToolButtonActionEvent
func callbackQToolButtonActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::actionEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "actionEvent")
	if signal != nil {
		defer signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectChangeEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QToolButton::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QToolButton) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QToolButton::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQToolButtonChangeEvent
func callbackQToolButtonChangeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::changeEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "changeEvent")
	if signal != nil {
		defer signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectEnterEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QToolButton::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QToolButton) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QToolButton::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQToolButtonEnterEvent
func callbackQToolButtonEnterEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::enterEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "enterEvent")
	if signal != nil {
		defer signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectLeaveEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QToolButton::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QToolButton) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QToolButton::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQToolButtonLeaveEvent
func callbackQToolButtonLeaveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::leaveEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "leaveEvent")
	if signal != nil {
		defer signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QToolButton) Menu() *QMenu {
	defer qt.Recovering("QToolButton::menu")

	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QToolButton_Menu(ptr.Pointer()))
	}
	return nil
}

func (ptr *QToolButton) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QToolButton::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QToolButton) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QToolButton::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQToolButtonMousePressEvent
func callbackQToolButtonMousePressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::mousePressEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mousePressEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QToolButton::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QToolButton) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QToolButton::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQToolButtonMouseReleaseEvent
func callbackQToolButtonMouseReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::mouseReleaseEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectNextCheckState(f func()) {
	defer qt.Recovering("connect QToolButton::nextCheckState")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "nextCheckState", f)
	}
}

func (ptr *QToolButton) DisconnectNextCheckState() {
	defer qt.Recovering("disconnect QToolButton::nextCheckState")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "nextCheckState")
	}
}

//export callbackQToolButtonNextCheckState
func callbackQToolButtonNextCheckState(ptrName *C.char) bool {
	defer qt.Recovering("callback QToolButton::nextCheckState")

	var signal = qt.GetSignal(C.GoString(ptrName), "nextCheckState")
	if signal != nil {
		defer signal.(func())()
		return true
	}
	return false

}

func (ptr *QToolButton) SetMenu(menu QMenu_ITF) {
	defer qt.Recovering("QToolButton::setMenu")

	if ptr.Pointer() != nil {
		C.QToolButton_SetMenu(ptr.Pointer(), PointerFromQMenu(menu))
	}
}

func (ptr *QToolButton) ShowMenu() {
	defer qt.Recovering("QToolButton::showMenu")

	if ptr.Pointer() != nil {
		C.QToolButton_ShowMenu(ptr.Pointer())
	}
}

func (ptr *QToolButton) ConnectTimerEvent(f func(e *core.QTimerEvent)) {
	defer qt.Recovering("connect QToolButton::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QToolButton) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QToolButton::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQToolButtonTimerEvent
func callbackQToolButtonTimerEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::timerEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "timerEvent")
	if signal != nil {
		defer signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectTriggered(f func(action *QAction)) {
	defer qt.Recovering("connect QToolButton::triggered")

	if ptr.Pointer() != nil {
		C.QToolButton_ConnectTriggered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "triggered", f)
	}
}

func (ptr *QToolButton) DisconnectTriggered() {
	defer qt.Recovering("disconnect QToolButton::triggered")

	if ptr.Pointer() != nil {
		C.QToolButton_DisconnectTriggered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "triggered")
	}
}

//export callbackQToolButtonTriggered
func callbackQToolButtonTriggered(ptrName *C.char, action unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::triggered")

	var signal = qt.GetSignal(C.GoString(ptrName), "triggered")
	if signal != nil {
		signal.(func(*QAction))(NewQActionFromPointer(action))
	}

}

func (ptr *QToolButton) DestroyQToolButton() {
	defer qt.Recovering("QToolButton::~QToolButton")

	if ptr.Pointer() != nil {
		C.QToolButton_DestroyQToolButton(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
