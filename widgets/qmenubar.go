package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QMenuBar struct {
	QWidget
}

type QMenuBar_ITF interface {
	QWidget_ITF
	QMenuBar_PTR() *QMenuBar
}

func PointerFromQMenuBar(ptr QMenuBar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMenuBar_PTR().Pointer()
	}
	return nil
}

func NewQMenuBarFromPointer(ptr unsafe.Pointer) *QMenuBar {
	var n = new(QMenuBar)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMenuBar_") {
		n.SetObjectName("QMenuBar_" + qt.Identifier())
	}
	return n
}

func (ptr *QMenuBar) QMenuBar_PTR() *QMenuBar {
	return ptr
}

func (ptr *QMenuBar) IsDefaultUp() bool {
	defer qt.Recovering("QMenuBar::isDefaultUp")

	if ptr.Pointer() != nil {
		return C.QMenuBar_IsDefaultUp(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMenuBar) IsNativeMenuBar() bool {
	defer qt.Recovering("QMenuBar::isNativeMenuBar")

	if ptr.Pointer() != nil {
		return C.QMenuBar_IsNativeMenuBar(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMenuBar) SetCornerWidget(widget QWidget_ITF, corner core.Qt__Corner) {
	defer qt.Recovering("QMenuBar::setCornerWidget")

	if ptr.Pointer() != nil {
		C.QMenuBar_SetCornerWidget(ptr.Pointer(), PointerFromQWidget(widget), C.int(corner))
	}
}

func (ptr *QMenuBar) SetDefaultUp(v bool) {
	defer qt.Recovering("QMenuBar::setDefaultUp")

	if ptr.Pointer() != nil {
		C.QMenuBar_SetDefaultUp(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QMenuBar) SetNativeMenuBar(nativeMenuBar bool) {
	defer qt.Recovering("QMenuBar::setNativeMenuBar")

	if ptr.Pointer() != nil {
		C.QMenuBar_SetNativeMenuBar(ptr.Pointer(), C.int(qt.GoBoolToInt(nativeMenuBar)))
	}
}

func NewQMenuBar(parent QWidget_ITF) *QMenuBar {
	defer qt.Recovering("QMenuBar::QMenuBar")

	return NewQMenuBarFromPointer(C.QMenuBar_NewQMenuBar(PointerFromQWidget(parent)))
}

func (ptr *QMenuBar) ActionAt(pt core.QPoint_ITF) *QAction {
	defer qt.Recovering("QMenuBar::actionAt")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenuBar_ActionAt(ptr.Pointer(), core.PointerFromQPoint(pt)))
	}
	return nil
}

func (ptr *QMenuBar) ConnectActionEvent(f func(e *gui.QActionEvent)) {
	defer qt.Recovering("connect QMenuBar::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QMenuBar::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQMenuBarActionEvent
func callbackQMenuBarActionEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenuBar::actionEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "actionEvent")
	if signal != nil {
		defer signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QMenuBar) ActionGeometry(act QAction_ITF) *core.QRect {
	defer qt.Recovering("QMenuBar::actionGeometry")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QMenuBar_ActionGeometry(ptr.Pointer(), PointerFromQAction(act)))
	}
	return nil
}

func (ptr *QMenuBar) ActiveAction() *QAction {
	defer qt.Recovering("QMenuBar::activeAction")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenuBar_ActiveAction(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMenuBar) AddAction(text string) *QAction {
	defer qt.Recovering("QMenuBar::addAction")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenuBar_AddAction(ptr.Pointer(), C.CString(text)))
	}
	return nil
}

func (ptr *QMenuBar) AddAction2(text string, receiver core.QObject_ITF, member string) *QAction {
	defer qt.Recovering("QMenuBar::addAction")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenuBar_AddAction2(ptr.Pointer(), C.CString(text), core.PointerFromQObject(receiver), C.CString(member)))
	}
	return nil
}

func (ptr *QMenuBar) AddMenu(menu QMenu_ITF) *QAction {
	defer qt.Recovering("QMenuBar::addMenu")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenuBar_AddMenu(ptr.Pointer(), PointerFromQMenu(menu)))
	}
	return nil
}

func (ptr *QMenuBar) AddMenu3(icon gui.QIcon_ITF, title string) *QMenu {
	defer qt.Recovering("QMenuBar::addMenu")

	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QMenuBar_AddMenu3(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(title)))
	}
	return nil
}

func (ptr *QMenuBar) AddMenu2(title string) *QMenu {
	defer qt.Recovering("QMenuBar::addMenu")

	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QMenuBar_AddMenu2(ptr.Pointer(), C.CString(title)))
	}
	return nil
}

func (ptr *QMenuBar) AddSeparator() *QAction {
	defer qt.Recovering("QMenuBar::addSeparator")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenuBar_AddSeparator(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMenuBar) ConnectChangeEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QMenuBar::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QMenuBar::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQMenuBarChangeEvent
func callbackQMenuBarChangeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenuBar::changeEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "changeEvent")
	if signal != nil {
		defer signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QMenuBar) Clear() {
	defer qt.Recovering("QMenuBar::clear")

	if ptr.Pointer() != nil {
		C.QMenuBar_Clear(ptr.Pointer())
	}
}

func (ptr *QMenuBar) CornerWidget(corner core.Qt__Corner) *QWidget {
	defer qt.Recovering("QMenuBar::cornerWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QMenuBar_CornerWidget(ptr.Pointer(), C.int(corner)))
	}
	return nil
}

func (ptr *QMenuBar) ConnectFocusInEvent(f func(v *gui.QFocusEvent)) {
	defer qt.Recovering("connect QMenuBar::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QMenuBar::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQMenuBarFocusInEvent
func callbackQMenuBarFocusInEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenuBar::focusInEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "focusInEvent")
	if signal != nil {
		defer signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QMenuBar) ConnectFocusOutEvent(f func(v *gui.QFocusEvent)) {
	defer qt.Recovering("connect QMenuBar::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QMenuBar::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQMenuBarFocusOutEvent
func callbackQMenuBarFocusOutEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenuBar::focusOutEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "focusOutEvent")
	if signal != nil {
		defer signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QMenuBar) HeightForWidth(v int) int {
	defer qt.Recovering("QMenuBar::heightForWidth")

	if ptr.Pointer() != nil {
		return int(C.QMenuBar_HeightForWidth(ptr.Pointer(), C.int(v)))
	}
	return 0
}

func (ptr *QMenuBar) ConnectHovered(f func(action *QAction)) {
	defer qt.Recovering("connect QMenuBar::hovered")

	if ptr.Pointer() != nil {
		C.QMenuBar_ConnectHovered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "hovered", f)
	}
}

func (ptr *QMenuBar) DisconnectHovered() {
	defer qt.Recovering("disconnect QMenuBar::hovered")

	if ptr.Pointer() != nil {
		C.QMenuBar_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "hovered")
	}
}

//export callbackQMenuBarHovered
func callbackQMenuBarHovered(ptrName *C.char, action unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::hovered")

	var signal = qt.GetSignal(C.GoString(ptrName), "hovered")
	if signal != nil {
		signal.(func(*QAction))(NewQActionFromPointer(action))
	}

}

func (ptr *QMenuBar) InsertMenu(before QAction_ITF, menu QMenu_ITF) *QAction {
	defer qt.Recovering("QMenuBar::insertMenu")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenuBar_InsertMenu(ptr.Pointer(), PointerFromQAction(before), PointerFromQMenu(menu)))
	}
	return nil
}

func (ptr *QMenuBar) InsertSeparator(before QAction_ITF) *QAction {
	defer qt.Recovering("QMenuBar::insertSeparator")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenuBar_InsertSeparator(ptr.Pointer(), PointerFromQAction(before)))
	}
	return nil
}

func (ptr *QMenuBar) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QMenuBar::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QMenuBar::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQMenuBarKeyPressEvent
func callbackQMenuBarKeyPressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenuBar::keyPressEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "keyPressEvent")
	if signal != nil {
		defer signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QMenuBar) ConnectLeaveEvent(f func(v *core.QEvent)) {
	defer qt.Recovering("connect QMenuBar::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QMenuBar::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQMenuBarLeaveEvent
func callbackQMenuBarLeaveEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenuBar::leaveEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "leaveEvent")
	if signal != nil {
		defer signal.(func(*core.QEvent))(core.NewQEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QMenuBar) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QMenuBar::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QMenuBar_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMenuBar) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMenuBar::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QMenuBar::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQMenuBarMouseMoveEvent
func callbackQMenuBarMouseMoveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenuBar::mouseMoveEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QMenuBar) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMenuBar::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QMenuBar::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQMenuBarMousePressEvent
func callbackQMenuBarMousePressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenuBar::mousePressEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mousePressEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QMenuBar) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMenuBar::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QMenuBar::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQMenuBarMouseReleaseEvent
func callbackQMenuBarMouseReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenuBar::mouseReleaseEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QMenuBar) ConnectPaintEvent(f func(e *gui.QPaintEvent)) {
	defer qt.Recovering("connect QMenuBar::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QMenuBar::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQMenuBarPaintEvent
func callbackQMenuBarPaintEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenuBar::paintEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "paintEvent")
	if signal != nil {
		defer signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QMenuBar) ConnectResizeEvent(f func(v *gui.QResizeEvent)) {
	defer qt.Recovering("connect QMenuBar::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QMenuBar::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQMenuBarResizeEvent
func callbackQMenuBarResizeEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenuBar::resizeEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "resizeEvent")
	if signal != nil {
		defer signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QMenuBar) SetActiveAction(act QAction_ITF) {
	defer qt.Recovering("QMenuBar::setActiveAction")

	if ptr.Pointer() != nil {
		C.QMenuBar_SetActiveAction(ptr.Pointer(), PointerFromQAction(act))
	}
}

func (ptr *QMenuBar) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QMenuBar::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QMenuBar) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QMenuBar::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQMenuBarSetVisible
func callbackQMenuBarSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QMenuBar::setVisible")

	var signal = qt.GetSignal(C.GoString(ptrName), "setVisible")
	if signal != nil {
		defer signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QMenuBar) SizeHint() *core.QSize {
	defer qt.Recovering("QMenuBar::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QMenuBar_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMenuBar) ConnectTimerEvent(f func(e *core.QTimerEvent)) {
	defer qt.Recovering("connect QMenuBar::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMenuBar::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMenuBarTimerEvent
func callbackQMenuBarTimerEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenuBar::timerEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "timerEvent")
	if signal != nil {
		defer signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QMenuBar) ConnectTriggered(f func(action *QAction)) {
	defer qt.Recovering("connect QMenuBar::triggered")

	if ptr.Pointer() != nil {
		C.QMenuBar_ConnectTriggered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "triggered", f)
	}
}

func (ptr *QMenuBar) DisconnectTriggered() {
	defer qt.Recovering("disconnect QMenuBar::triggered")

	if ptr.Pointer() != nil {
		C.QMenuBar_DisconnectTriggered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "triggered")
	}
}

//export callbackQMenuBarTriggered
func callbackQMenuBarTriggered(ptrName *C.char, action unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::triggered")

	var signal = qt.GetSignal(C.GoString(ptrName), "triggered")
	if signal != nil {
		signal.(func(*QAction))(NewQActionFromPointer(action))
	}

}

func (ptr *QMenuBar) DestroyQMenuBar() {
	defer qt.Recovering("QMenuBar::~QMenuBar")

	if ptr.Pointer() != nil {
		C.QMenuBar_DestroyQMenuBar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
