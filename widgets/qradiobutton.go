package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QRadioButton struct {
	QAbstractButton
}

type QRadioButton_ITF interface {
	QAbstractButton_ITF
	QRadioButton_PTR() *QRadioButton
}

func PointerFromQRadioButton(ptr QRadioButton_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRadioButton_PTR().Pointer()
	}
	return nil
}

func NewQRadioButtonFromPointer(ptr unsafe.Pointer) *QRadioButton {
	var n = new(QRadioButton)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QRadioButton_") {
		n.SetObjectName("QRadioButton_" + qt.Identifier())
	}
	return n
}

func (ptr *QRadioButton) QRadioButton_PTR() *QRadioButton {
	return ptr
}

func NewQRadioButton(parent QWidget_ITF) *QRadioButton {
	defer qt.Recovering("QRadioButton::QRadioButton")

	return NewQRadioButtonFromPointer(C.QRadioButton_NewQRadioButton(PointerFromQWidget(parent)))
}

func NewQRadioButton2(text string, parent QWidget_ITF) *QRadioButton {
	defer qt.Recovering("QRadioButton::QRadioButton")

	return NewQRadioButtonFromPointer(C.QRadioButton_NewQRadioButton2(C.CString(text), PointerFromQWidget(parent)))
}

func (ptr *QRadioButton) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QRadioButton::event")

	if ptr.Pointer() != nil {
		return C.QRadioButton_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QRadioButton) HitButton(pos core.QPoint_ITF) bool {
	defer qt.Recovering("QRadioButton::hitButton")

	if ptr.Pointer() != nil {
		return C.QRadioButton_HitButton(ptr.Pointer(), core.PointerFromQPoint(pos)) != 0
	}
	return false
}

func (ptr *QRadioButton) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QRadioButton::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QRadioButton_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QRadioButton) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QRadioButton::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QRadioButton::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQRadioButtonMouseMoveEvent
func callbackQRadioButtonMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QRadioButton::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQRadioButtonFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QRadioButton) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QRadioButton::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QRadioButton) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QRadioButton::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QRadioButton) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QRadioButton::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QRadioButton::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQRadioButtonPaintEvent
func callbackQRadioButtonPaintEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QRadioButton::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
	} else {
		NewQRadioButtonFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(v))
	}
}

func (ptr *QRadioButton) PaintEvent(v gui.QPaintEvent_ITF) {
	defer qt.Recovering("QRadioButton::paintEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(v))
	}
}

func (ptr *QRadioButton) PaintEventDefault(v gui.QPaintEvent_ITF) {
	defer qt.Recovering("QRadioButton::paintEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(v))
	}
}

func (ptr *QRadioButton) SizeHint() *core.QSize {
	defer qt.Recovering("QRadioButton::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QRadioButton_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QRadioButton) DestroyQRadioButton() {
	defer qt.Recovering("QRadioButton::~QRadioButton")

	if ptr.Pointer() != nil {
		C.QRadioButton_DestroyQRadioButton(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QRadioButton) ConnectChangeEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QRadioButton::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QRadioButton::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQRadioButtonChangeEvent
func callbackQRadioButtonChangeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QRadioButton::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
	} else {
		NewQRadioButtonFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(e))
	}
}

func (ptr *QRadioButton) ChangeEvent(e core.QEvent_ITF) {
	defer qt.Recovering("QRadioButton::changeEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

func (ptr *QRadioButton) ChangeEventDefault(e core.QEvent_ITF) {
	defer qt.Recovering("QRadioButton::changeEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

func (ptr *QRadioButton) ConnectCheckStateSet(f func()) {
	defer qt.Recovering("connect QRadioButton::checkStateSet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "checkStateSet", f)
	}
}

func (ptr *QRadioButton) DisconnectCheckStateSet() {
	defer qt.Recovering("disconnect QRadioButton::checkStateSet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "checkStateSet")
	}
}

//export callbackQRadioButtonCheckStateSet
func callbackQRadioButtonCheckStateSet(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QRadioButton::checkStateSet")

	if signal := qt.GetSignal(C.GoString(ptrName), "checkStateSet"); signal != nil {
		signal.(func())()
	} else {
		NewQRadioButtonFromPointer(ptr).CheckStateSetDefault()
	}
}

func (ptr *QRadioButton) CheckStateSet() {
	defer qt.Recovering("QRadioButton::checkStateSet")

	if ptr.Pointer() != nil {
		C.QRadioButton_CheckStateSet(ptr.Pointer())
	}
}

func (ptr *QRadioButton) CheckStateSetDefault() {
	defer qt.Recovering("QRadioButton::checkStateSet")

	if ptr.Pointer() != nil {
		C.QRadioButton_CheckStateSetDefault(ptr.Pointer())
	}
}

func (ptr *QRadioButton) ConnectFocusInEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QRadioButton::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QRadioButton::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQRadioButtonFocusInEvent
func callbackQRadioButtonFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QRadioButton::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
	} else {
		NewQRadioButtonFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(e))
	}
}

func (ptr *QRadioButton) FocusInEvent(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QRadioButton::focusInEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QRadioButton) FocusInEventDefault(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QRadioButton::focusInEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QRadioButton) ConnectFocusOutEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QRadioButton::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QRadioButton::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQRadioButtonFocusOutEvent
func callbackQRadioButtonFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QRadioButton::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
	} else {
		NewQRadioButtonFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(e))
	}
}

func (ptr *QRadioButton) FocusOutEvent(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QRadioButton::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QRadioButton) FocusOutEventDefault(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QRadioButton::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QRadioButton) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QRadioButton::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QRadioButton::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQRadioButtonKeyPressEvent
func callbackQRadioButtonKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QRadioButton::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQRadioButtonFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QRadioButton) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QRadioButton::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QRadioButton) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QRadioButton::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QRadioButton) ConnectKeyReleaseEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QRadioButton::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QRadioButton::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQRadioButtonKeyReleaseEvent
func callbackQRadioButtonKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QRadioButton::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQRadioButtonFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QRadioButton) KeyReleaseEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QRadioButton::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QRadioButton) KeyReleaseEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QRadioButton::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QRadioButton) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QRadioButton::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QRadioButton::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQRadioButtonMousePressEvent
func callbackQRadioButtonMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QRadioButton::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQRadioButtonFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QRadioButton) MousePressEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QRadioButton::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QRadioButton) MousePressEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QRadioButton::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QRadioButton) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QRadioButton::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QRadioButton::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQRadioButtonMouseReleaseEvent
func callbackQRadioButtonMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QRadioButton::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQRadioButtonFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QRadioButton) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QRadioButton::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QRadioButton) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QRadioButton::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QRadioButton) ConnectNextCheckState(f func()) {
	defer qt.Recovering("connect QRadioButton::nextCheckState")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "nextCheckState", f)
	}
}

func (ptr *QRadioButton) DisconnectNextCheckState() {
	defer qt.Recovering("disconnect QRadioButton::nextCheckState")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "nextCheckState")
	}
}

//export callbackQRadioButtonNextCheckState
func callbackQRadioButtonNextCheckState(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QRadioButton::nextCheckState")

	if signal := qt.GetSignal(C.GoString(ptrName), "nextCheckState"); signal != nil {
		signal.(func())()
	} else {
		NewQRadioButtonFromPointer(ptr).NextCheckStateDefault()
	}
}

func (ptr *QRadioButton) NextCheckState() {
	defer qt.Recovering("QRadioButton::nextCheckState")

	if ptr.Pointer() != nil {
		C.QRadioButton_NextCheckState(ptr.Pointer())
	}
}

func (ptr *QRadioButton) NextCheckStateDefault() {
	defer qt.Recovering("QRadioButton::nextCheckState")

	if ptr.Pointer() != nil {
		C.QRadioButton_NextCheckStateDefault(ptr.Pointer())
	}
}

func (ptr *QRadioButton) ConnectTimerEvent(f func(e *core.QTimerEvent)) {
	defer qt.Recovering("connect QRadioButton::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QRadioButton::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQRadioButtonTimerEvent
func callbackQRadioButtonTimerEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QRadioButton::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
	} else {
		NewQRadioButtonFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(e))
	}
}

func (ptr *QRadioButton) TimerEvent(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QRadioButton::timerEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

func (ptr *QRadioButton) TimerEventDefault(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QRadioButton::timerEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

func (ptr *QRadioButton) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QRadioButton::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QRadioButton::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQRadioButtonActionEvent
func callbackQRadioButtonActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioButton::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQRadioButtonFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QRadioButton) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QRadioButton::actionEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QRadioButton) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QRadioButton::actionEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QRadioButton) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QRadioButton::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QRadioButton::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQRadioButtonDragEnterEvent
func callbackQRadioButtonDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioButton::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQRadioButtonFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QRadioButton) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QRadioButton::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QRadioButton) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QRadioButton::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QRadioButton) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QRadioButton::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QRadioButton::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQRadioButtonDragLeaveEvent
func callbackQRadioButtonDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioButton::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQRadioButtonFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QRadioButton) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QRadioButton::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QRadioButton) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QRadioButton::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QRadioButton) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QRadioButton::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QRadioButton::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQRadioButtonDragMoveEvent
func callbackQRadioButtonDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioButton::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQRadioButtonFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QRadioButton) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QRadioButton::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QRadioButton) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QRadioButton::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QRadioButton) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QRadioButton::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QRadioButton::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQRadioButtonDropEvent
func callbackQRadioButtonDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioButton::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQRadioButtonFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QRadioButton) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QRadioButton::dropEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QRadioButton) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QRadioButton::dropEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QRadioButton) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QRadioButton::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QRadioButton::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQRadioButtonEnterEvent
func callbackQRadioButtonEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioButton::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQRadioButtonFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QRadioButton) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QRadioButton::enterEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QRadioButton) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QRadioButton::enterEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QRadioButton) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QRadioButton::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QRadioButton::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQRadioButtonHideEvent
func callbackQRadioButtonHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioButton::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQRadioButtonFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QRadioButton) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QRadioButton::hideEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QRadioButton) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QRadioButton::hideEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QRadioButton) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QRadioButton::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QRadioButton::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQRadioButtonLeaveEvent
func callbackQRadioButtonLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioButton::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQRadioButtonFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QRadioButton) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QRadioButton::leaveEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QRadioButton) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QRadioButton::leaveEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QRadioButton) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QRadioButton::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QRadioButton::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQRadioButtonMoveEvent
func callbackQRadioButtonMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioButton::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQRadioButtonFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QRadioButton) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QRadioButton::moveEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QRadioButton) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QRadioButton::moveEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QRadioButton) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QRadioButton::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QRadioButton) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QRadioButton::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQRadioButtonSetVisible
func callbackQRadioButtonSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QRadioButton::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QRadioButton) SetVisible(visible bool) {
	defer qt.Recovering("QRadioButton::setVisible")

	if ptr.Pointer() != nil {
		C.QRadioButton_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QRadioButton) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QRadioButton::setVisible")

	if ptr.Pointer() != nil {
		C.QRadioButton_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QRadioButton) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QRadioButton::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QRadioButton::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQRadioButtonShowEvent
func callbackQRadioButtonShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioButton::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQRadioButtonFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QRadioButton) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QRadioButton::showEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QRadioButton) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QRadioButton::showEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QRadioButton) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QRadioButton::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QRadioButton::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQRadioButtonCloseEvent
func callbackQRadioButtonCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioButton::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQRadioButtonFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QRadioButton) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QRadioButton::closeEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QRadioButton) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QRadioButton::closeEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QRadioButton) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QRadioButton::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QRadioButton::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQRadioButtonContextMenuEvent
func callbackQRadioButtonContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioButton::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQRadioButtonFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QRadioButton) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QRadioButton::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QRadioButton) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QRadioButton::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QRadioButton) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QRadioButton::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QRadioButton) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QRadioButton::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQRadioButtonInitPainter
func callbackQRadioButtonInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QRadioButton::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQRadioButtonFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QRadioButton) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QRadioButton::initPainter")

	if ptr.Pointer() != nil {
		C.QRadioButton_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QRadioButton) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QRadioButton::initPainter")

	if ptr.Pointer() != nil {
		C.QRadioButton_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QRadioButton) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QRadioButton::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QRadioButton::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQRadioButtonInputMethodEvent
func callbackQRadioButtonInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioButton::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQRadioButtonFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QRadioButton) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QRadioButton::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QRadioButton) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QRadioButton::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QRadioButton) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QRadioButton::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QRadioButton::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQRadioButtonMouseDoubleClickEvent
func callbackQRadioButtonMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioButton::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQRadioButtonFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QRadioButton) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QRadioButton::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QRadioButton) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QRadioButton::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QRadioButton) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QRadioButton::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QRadioButton::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQRadioButtonResizeEvent
func callbackQRadioButtonResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioButton::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQRadioButtonFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QRadioButton) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QRadioButton::resizeEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QRadioButton) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QRadioButton::resizeEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QRadioButton) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QRadioButton::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QRadioButton::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQRadioButtonTabletEvent
func callbackQRadioButtonTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioButton::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQRadioButtonFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QRadioButton) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QRadioButton::tabletEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QRadioButton) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QRadioButton::tabletEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QRadioButton) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QRadioButton::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QRadioButton::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQRadioButtonWheelEvent
func callbackQRadioButtonWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioButton::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQRadioButtonFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QRadioButton) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QRadioButton::wheelEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QRadioButton) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QRadioButton::wheelEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QRadioButton) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QRadioButton::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QRadioButton::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQRadioButtonChildEvent
func callbackQRadioButtonChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioButton::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQRadioButtonFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QRadioButton) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QRadioButton::childEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QRadioButton) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QRadioButton::childEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QRadioButton) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QRadioButton::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QRadioButton::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQRadioButtonCustomEvent
func callbackQRadioButtonCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioButton::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQRadioButtonFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QRadioButton) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QRadioButton::customEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QRadioButton) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QRadioButton::customEvent")

	if ptr.Pointer() != nil {
		C.QRadioButton_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
