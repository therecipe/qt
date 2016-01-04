package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QCheckBox struct {
	QAbstractButton
}

type QCheckBox_ITF interface {
	QAbstractButton_ITF
	QCheckBox_PTR() *QCheckBox
}

func PointerFromQCheckBox(ptr QCheckBox_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCheckBox_PTR().Pointer()
	}
	return nil
}

func NewQCheckBoxFromPointer(ptr unsafe.Pointer) *QCheckBox {
	var n = new(QCheckBox)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCheckBox_") {
		n.SetObjectName("QCheckBox_" + qt.Identifier())
	}
	return n
}

func (ptr *QCheckBox) QCheckBox_PTR() *QCheckBox {
	return ptr
}

func (ptr *QCheckBox) IsTristate() bool {
	defer qt.Recovering("QCheckBox::isTristate")

	if ptr.Pointer() != nil {
		return C.QCheckBox_IsTristate(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCheckBox) SetTristate(y bool) {
	defer qt.Recovering("QCheckBox::setTristate")

	if ptr.Pointer() != nil {
		C.QCheckBox_SetTristate(ptr.Pointer(), C.int(qt.GoBoolToInt(y)))
	}
}

func NewQCheckBox(parent QWidget_ITF) *QCheckBox {
	defer qt.Recovering("QCheckBox::QCheckBox")

	return NewQCheckBoxFromPointer(C.QCheckBox_NewQCheckBox(PointerFromQWidget(parent)))
}

func NewQCheckBox2(text string, parent QWidget_ITF) *QCheckBox {
	defer qt.Recovering("QCheckBox::QCheckBox")

	return NewQCheckBoxFromPointer(C.QCheckBox_NewQCheckBox2(C.CString(text), PointerFromQWidget(parent)))
}

func (ptr *QCheckBox) CheckState() core.Qt__CheckState {
	defer qt.Recovering("QCheckBox::checkState")

	if ptr.Pointer() != nil {
		return core.Qt__CheckState(C.QCheckBox_CheckState(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCheckBox) ConnectCheckStateSet(f func()) {
	defer qt.Recovering("connect QCheckBox::checkStateSet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "checkStateSet", f)
	}
}

func (ptr *QCheckBox) DisconnectCheckStateSet() {
	defer qt.Recovering("disconnect QCheckBox::checkStateSet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "checkStateSet")
	}
}

//export callbackQCheckBoxCheckStateSet
func callbackQCheckBoxCheckStateSet(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QCheckBox::checkStateSet")

	if signal := qt.GetSignal(C.GoString(ptrName), "checkStateSet"); signal != nil {
		signal.(func())()
	} else {
		NewQCheckBoxFromPointer(ptr).CheckStateSetDefault()
	}
}

func (ptr *QCheckBox) CheckStateSet() {
	defer qt.Recovering("QCheckBox::checkStateSet")

	if ptr.Pointer() != nil {
		C.QCheckBox_CheckStateSet(ptr.Pointer())
	}
}

func (ptr *QCheckBox) CheckStateSetDefault() {
	defer qt.Recovering("QCheckBox::checkStateSet")

	if ptr.Pointer() != nil {
		C.QCheckBox_CheckStateSetDefault(ptr.Pointer())
	}
}

func (ptr *QCheckBox) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QCheckBox::event")

	if ptr.Pointer() != nil {
		return C.QCheckBox_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QCheckBox) HitButton(pos core.QPoint_ITF) bool {
	defer qt.Recovering("QCheckBox::hitButton")

	if ptr.Pointer() != nil {
		return C.QCheckBox_HitButton(ptr.Pointer(), core.PointerFromQPoint(pos)) != 0
	}
	return false
}

func (ptr *QCheckBox) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QCheckBox::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QCheckBox_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCheckBox) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QCheckBox::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QCheckBox::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQCheckBoxMouseMoveEvent
func callbackQCheckBoxMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QCheckBox::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQCheckBoxFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QCheckBox) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCheckBox::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QCheckBox) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCheckBox::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QCheckBox) ConnectNextCheckState(f func()) {
	defer qt.Recovering("connect QCheckBox::nextCheckState")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "nextCheckState", f)
	}
}

func (ptr *QCheckBox) DisconnectNextCheckState() {
	defer qt.Recovering("disconnect QCheckBox::nextCheckState")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "nextCheckState")
	}
}

//export callbackQCheckBoxNextCheckState
func callbackQCheckBoxNextCheckState(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QCheckBox::nextCheckState")

	if signal := qt.GetSignal(C.GoString(ptrName), "nextCheckState"); signal != nil {
		signal.(func())()
	} else {
		NewQCheckBoxFromPointer(ptr).NextCheckStateDefault()
	}
}

func (ptr *QCheckBox) NextCheckState() {
	defer qt.Recovering("QCheckBox::nextCheckState")

	if ptr.Pointer() != nil {
		C.QCheckBox_NextCheckState(ptr.Pointer())
	}
}

func (ptr *QCheckBox) NextCheckStateDefault() {
	defer qt.Recovering("QCheckBox::nextCheckState")

	if ptr.Pointer() != nil {
		C.QCheckBox_NextCheckStateDefault(ptr.Pointer())
	}
}

func (ptr *QCheckBox) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QCheckBox::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QCheckBox::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQCheckBoxPaintEvent
func callbackQCheckBoxPaintEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QCheckBox::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
	} else {
		NewQCheckBoxFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(v))
	}
}

func (ptr *QCheckBox) PaintEvent(v gui.QPaintEvent_ITF) {
	defer qt.Recovering("QCheckBox::paintEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(v))
	}
}

func (ptr *QCheckBox) PaintEventDefault(v gui.QPaintEvent_ITF) {
	defer qt.Recovering("QCheckBox::paintEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(v))
	}
}

func (ptr *QCheckBox) SetCheckState(state core.Qt__CheckState) {
	defer qt.Recovering("QCheckBox::setCheckState")

	if ptr.Pointer() != nil {
		C.QCheckBox_SetCheckState(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QCheckBox) SizeHint() *core.QSize {
	defer qt.Recovering("QCheckBox::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QCheckBox_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCheckBox) ConnectStateChanged(f func(state int)) {
	defer qt.Recovering("connect QCheckBox::stateChanged")

	if ptr.Pointer() != nil {
		C.QCheckBox_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QCheckBox) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QCheckBox::stateChanged")

	if ptr.Pointer() != nil {
		C.QCheckBox_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQCheckBoxStateChanged
func callbackQCheckBoxStateChanged(ptr unsafe.Pointer, ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QCheckBox::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(int))(int(state))
	}

}

func (ptr *QCheckBox) StateChanged(state int) {
	defer qt.Recovering("QCheckBox::stateChanged")

	if ptr.Pointer() != nil {
		C.QCheckBox_StateChanged(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QCheckBox) DestroyQCheckBox() {
	defer qt.Recovering("QCheckBox::~QCheckBox")

	if ptr.Pointer() != nil {
		C.QCheckBox_DestroyQCheckBox(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCheckBox) ConnectChangeEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QCheckBox::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QCheckBox::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQCheckBoxChangeEvent
func callbackQCheckBoxChangeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QCheckBox::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
	} else {
		NewQCheckBoxFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(e))
	}
}

func (ptr *QCheckBox) ChangeEvent(e core.QEvent_ITF) {
	defer qt.Recovering("QCheckBox::changeEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

func (ptr *QCheckBox) ChangeEventDefault(e core.QEvent_ITF) {
	defer qt.Recovering("QCheckBox::changeEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

func (ptr *QCheckBox) ConnectFocusInEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QCheckBox::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QCheckBox::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQCheckBoxFocusInEvent
func callbackQCheckBoxFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QCheckBox::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
	} else {
		NewQCheckBoxFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(e))
	}
}

func (ptr *QCheckBox) FocusInEvent(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QCheckBox::focusInEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QCheckBox) FocusInEventDefault(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QCheckBox::focusInEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QCheckBox) ConnectFocusOutEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QCheckBox::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QCheckBox::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQCheckBoxFocusOutEvent
func callbackQCheckBoxFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QCheckBox::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
	} else {
		NewQCheckBoxFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(e))
	}
}

func (ptr *QCheckBox) FocusOutEvent(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QCheckBox::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QCheckBox) FocusOutEventDefault(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QCheckBox::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QCheckBox) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QCheckBox::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QCheckBox::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQCheckBoxKeyPressEvent
func callbackQCheckBoxKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QCheckBox::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQCheckBoxFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QCheckBox) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QCheckBox::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QCheckBox) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QCheckBox::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QCheckBox) ConnectKeyReleaseEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QCheckBox::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QCheckBox::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQCheckBoxKeyReleaseEvent
func callbackQCheckBoxKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QCheckBox::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQCheckBoxFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QCheckBox) KeyReleaseEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QCheckBox::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QCheckBox) KeyReleaseEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QCheckBox::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QCheckBox) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QCheckBox::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QCheckBox::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQCheckBoxMousePressEvent
func callbackQCheckBoxMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QCheckBox::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQCheckBoxFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QCheckBox) MousePressEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCheckBox::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QCheckBox) MousePressEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCheckBox::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QCheckBox) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QCheckBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QCheckBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQCheckBoxMouseReleaseEvent
func callbackQCheckBoxMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QCheckBox::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQCheckBoxFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QCheckBox) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCheckBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QCheckBox) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCheckBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QCheckBox) ConnectTimerEvent(f func(e *core.QTimerEvent)) {
	defer qt.Recovering("connect QCheckBox::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCheckBox::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCheckBoxTimerEvent
func callbackQCheckBoxTimerEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QCheckBox::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
	} else {
		NewQCheckBoxFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(e))
	}
}

func (ptr *QCheckBox) TimerEvent(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QCheckBox::timerEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

func (ptr *QCheckBox) TimerEventDefault(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QCheckBox::timerEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

func (ptr *QCheckBox) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QCheckBox::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QCheckBox::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQCheckBoxActionEvent
func callbackQCheckBoxActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCheckBox::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQCheckBoxFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QCheckBox) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QCheckBox::actionEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QCheckBox) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QCheckBox::actionEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QCheckBox) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QCheckBox::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QCheckBox::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQCheckBoxDragEnterEvent
func callbackQCheckBoxDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCheckBox::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQCheckBoxFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QCheckBox) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QCheckBox::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QCheckBox) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QCheckBox::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QCheckBox) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QCheckBox::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QCheckBox::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQCheckBoxDragLeaveEvent
func callbackQCheckBoxDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCheckBox::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQCheckBoxFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QCheckBox) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QCheckBox::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QCheckBox) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QCheckBox::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QCheckBox) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QCheckBox::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QCheckBox::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQCheckBoxDragMoveEvent
func callbackQCheckBoxDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCheckBox::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQCheckBoxFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QCheckBox) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QCheckBox::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QCheckBox) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QCheckBox::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QCheckBox) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QCheckBox::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QCheckBox::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQCheckBoxDropEvent
func callbackQCheckBoxDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCheckBox::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQCheckBoxFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QCheckBox) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QCheckBox::dropEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QCheckBox) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QCheckBox::dropEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QCheckBox) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCheckBox::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QCheckBox::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQCheckBoxEnterEvent
func callbackQCheckBoxEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCheckBox::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCheckBoxFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCheckBox) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCheckBox::enterEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCheckBox) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCheckBox::enterEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCheckBox) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QCheckBox::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QCheckBox::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQCheckBoxHideEvent
func callbackQCheckBoxHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCheckBox::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQCheckBoxFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QCheckBox) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QCheckBox::hideEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QCheckBox) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QCheckBox::hideEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QCheckBox) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCheckBox::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QCheckBox::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQCheckBoxLeaveEvent
func callbackQCheckBoxLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCheckBox::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCheckBoxFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCheckBox) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCheckBox::leaveEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCheckBox) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCheckBox::leaveEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCheckBox) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QCheckBox::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QCheckBox::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQCheckBoxMoveEvent
func callbackQCheckBoxMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCheckBox::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQCheckBoxFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QCheckBox) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QCheckBox::moveEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QCheckBox) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QCheckBox::moveEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QCheckBox) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QCheckBox::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QCheckBox) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QCheckBox::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQCheckBoxSetVisible
func callbackQCheckBoxSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QCheckBox::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QCheckBox) SetVisible(visible bool) {
	defer qt.Recovering("QCheckBox::setVisible")

	if ptr.Pointer() != nil {
		C.QCheckBox_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QCheckBox) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QCheckBox::setVisible")

	if ptr.Pointer() != nil {
		C.QCheckBox_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QCheckBox) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QCheckBox::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QCheckBox::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQCheckBoxShowEvent
func callbackQCheckBoxShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCheckBox::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQCheckBoxFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QCheckBox) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QCheckBox::showEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QCheckBox) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QCheckBox::showEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QCheckBox) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QCheckBox::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QCheckBox::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQCheckBoxCloseEvent
func callbackQCheckBoxCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCheckBox::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQCheckBoxFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QCheckBox) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QCheckBox::closeEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QCheckBox) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QCheckBox::closeEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QCheckBox) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QCheckBox::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QCheckBox::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQCheckBoxContextMenuEvent
func callbackQCheckBoxContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCheckBox::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQCheckBoxFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QCheckBox) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QCheckBox::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QCheckBox) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QCheckBox::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QCheckBox) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QCheckBox::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QCheckBox) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QCheckBox::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQCheckBoxInitPainter
func callbackQCheckBoxInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QCheckBox::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQCheckBoxFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QCheckBox) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QCheckBox::initPainter")

	if ptr.Pointer() != nil {
		C.QCheckBox_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QCheckBox) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QCheckBox::initPainter")

	if ptr.Pointer() != nil {
		C.QCheckBox_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QCheckBox) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QCheckBox::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QCheckBox::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQCheckBoxInputMethodEvent
func callbackQCheckBoxInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCheckBox::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQCheckBoxFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QCheckBox) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QCheckBox::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QCheckBox) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QCheckBox::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QCheckBox) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QCheckBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QCheckBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQCheckBoxMouseDoubleClickEvent
func callbackQCheckBoxMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCheckBox::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQCheckBoxFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QCheckBox) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCheckBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QCheckBox) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCheckBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QCheckBox) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QCheckBox::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QCheckBox::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQCheckBoxResizeEvent
func callbackQCheckBoxResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCheckBox::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQCheckBoxFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QCheckBox) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QCheckBox::resizeEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QCheckBox) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QCheckBox::resizeEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QCheckBox) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QCheckBox::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QCheckBox::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQCheckBoxTabletEvent
func callbackQCheckBoxTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCheckBox::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQCheckBoxFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QCheckBox) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QCheckBox::tabletEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QCheckBox) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QCheckBox::tabletEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QCheckBox) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QCheckBox::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QCheckBox::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQCheckBoxWheelEvent
func callbackQCheckBoxWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCheckBox::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQCheckBoxFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QCheckBox) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QCheckBox::wheelEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QCheckBox) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QCheckBox::wheelEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QCheckBox) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCheckBox::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCheckBox::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCheckBoxChildEvent
func callbackQCheckBoxChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCheckBox::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCheckBoxFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCheckBox) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCheckBox::childEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCheckBox) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCheckBox::childEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCheckBox) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCheckBox::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCheckBox::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCheckBoxCustomEvent
func callbackQCheckBoxCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCheckBox::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCheckBoxFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCheckBox) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCheckBox::customEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCheckBox) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCheckBox::customEvent")

	if ptr.Pointer() != nil {
		C.QCheckBox_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
