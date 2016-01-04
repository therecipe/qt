package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QCommandLinkButton struct {
	QPushButton
}

type QCommandLinkButton_ITF interface {
	QPushButton_ITF
	QCommandLinkButton_PTR() *QCommandLinkButton
}

func PointerFromQCommandLinkButton(ptr QCommandLinkButton_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCommandLinkButton_PTR().Pointer()
	}
	return nil
}

func NewQCommandLinkButtonFromPointer(ptr unsafe.Pointer) *QCommandLinkButton {
	var n = new(QCommandLinkButton)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCommandLinkButton_") {
		n.SetObjectName("QCommandLinkButton_" + qt.Identifier())
	}
	return n
}

func (ptr *QCommandLinkButton) QCommandLinkButton_PTR() *QCommandLinkButton {
	return ptr
}

func (ptr *QCommandLinkButton) Description() string {
	defer qt.Recovering("QCommandLinkButton::description")

	if ptr.Pointer() != nil {
		return C.GoString(C.QCommandLinkButton_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QCommandLinkButton) SetDescription(description string) {
	defer qt.Recovering("QCommandLinkButton::setDescription")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_SetDescription(ptr.Pointer(), C.CString(description))
	}
}

func NewQCommandLinkButton(parent QWidget_ITF) *QCommandLinkButton {
	defer qt.Recovering("QCommandLinkButton::QCommandLinkButton")

	return NewQCommandLinkButtonFromPointer(C.QCommandLinkButton_NewQCommandLinkButton(PointerFromQWidget(parent)))
}

func NewQCommandLinkButton2(text string, parent QWidget_ITF) *QCommandLinkButton {
	defer qt.Recovering("QCommandLinkButton::QCommandLinkButton")

	return NewQCommandLinkButtonFromPointer(C.QCommandLinkButton_NewQCommandLinkButton2(C.CString(text), PointerFromQWidget(parent)))
}

func NewQCommandLinkButton3(text string, description string, parent QWidget_ITF) *QCommandLinkButton {
	defer qt.Recovering("QCommandLinkButton::QCommandLinkButton")

	return NewQCommandLinkButtonFromPointer(C.QCommandLinkButton_NewQCommandLinkButton3(C.CString(text), C.CString(description), PointerFromQWidget(parent)))
}

func (ptr *QCommandLinkButton) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QCommandLinkButton::event")

	if ptr.Pointer() != nil {
		return C.QCommandLinkButton_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QCommandLinkButton) HeightForWidth(width int) int {
	defer qt.Recovering("QCommandLinkButton::heightForWidth")

	if ptr.Pointer() != nil {
		return int(C.QCommandLinkButton_HeightForWidth(ptr.Pointer(), C.int(width)))
	}
	return 0
}

func (ptr *QCommandLinkButton) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QCommandLinkButton::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QCommandLinkButton_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCommandLinkButton) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQCommandLinkButtonPaintEvent
func callbackQCommandLinkButtonPaintEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QCommandLinkButton::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
	} else {
		NewQCommandLinkButtonFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(v))
	}
}

func (ptr *QCommandLinkButton) PaintEvent(v gui.QPaintEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::paintEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(v))
	}
}

func (ptr *QCommandLinkButton) PaintEventDefault(v gui.QPaintEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::paintEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(v))
	}
}

func (ptr *QCommandLinkButton) SizeHint() *core.QSize {
	defer qt.Recovering("QCommandLinkButton::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QCommandLinkButton_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCommandLinkButton) DestroyQCommandLinkButton() {
	defer qt.Recovering("QCommandLinkButton::~QCommandLinkButton")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_DestroyQCommandLinkButton(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCommandLinkButton) ConnectFocusInEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQCommandLinkButtonFocusInEvent
func callbackQCommandLinkButtonFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QCommandLinkButton::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
	} else {
		NewQCommandLinkButtonFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(e))
	}
}

func (ptr *QCommandLinkButton) FocusInEvent(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::focusInEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QCommandLinkButton) FocusInEventDefault(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::focusInEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QCommandLinkButton) ConnectFocusOutEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQCommandLinkButtonFocusOutEvent
func callbackQCommandLinkButtonFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QCommandLinkButton::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
	} else {
		NewQCommandLinkButtonFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(e))
	}
}

func (ptr *QCommandLinkButton) FocusOutEvent(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QCommandLinkButton) FocusOutEventDefault(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QCommandLinkButton) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQCommandLinkButtonKeyPressEvent
func callbackQCommandLinkButtonKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QCommandLinkButton::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQCommandLinkButtonFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QCommandLinkButton) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QCommandLinkButton) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QCommandLinkButton) ConnectChangeEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQCommandLinkButtonChangeEvent
func callbackQCommandLinkButtonChangeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QCommandLinkButton::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
	} else {
		NewQCommandLinkButtonFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(e))
	}
}

func (ptr *QCommandLinkButton) ChangeEvent(e core.QEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::changeEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

func (ptr *QCommandLinkButton) ChangeEventDefault(e core.QEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::changeEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

func (ptr *QCommandLinkButton) ConnectCheckStateSet(f func()) {
	defer qt.Recovering("connect QCommandLinkButton::checkStateSet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "checkStateSet", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectCheckStateSet() {
	defer qt.Recovering("disconnect QCommandLinkButton::checkStateSet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "checkStateSet")
	}
}

//export callbackQCommandLinkButtonCheckStateSet
func callbackQCommandLinkButtonCheckStateSet(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QCommandLinkButton::checkStateSet")

	if signal := qt.GetSignal(C.GoString(ptrName), "checkStateSet"); signal != nil {
		signal.(func())()
	} else {
		NewQCommandLinkButtonFromPointer(ptr).CheckStateSetDefault()
	}
}

func (ptr *QCommandLinkButton) CheckStateSet() {
	defer qt.Recovering("QCommandLinkButton::checkStateSet")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_CheckStateSet(ptr.Pointer())
	}
}

func (ptr *QCommandLinkButton) CheckStateSetDefault() {
	defer qt.Recovering("QCommandLinkButton::checkStateSet")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_CheckStateSetDefault(ptr.Pointer())
	}
}

func (ptr *QCommandLinkButton) ConnectKeyReleaseEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQCommandLinkButtonKeyReleaseEvent
func callbackQCommandLinkButtonKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QCommandLinkButton::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQCommandLinkButtonFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QCommandLinkButton) KeyReleaseEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QCommandLinkButton) KeyReleaseEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QCommandLinkButton) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQCommandLinkButtonMouseMoveEvent
func callbackQCommandLinkButtonMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QCommandLinkButton::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQCommandLinkButtonFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QCommandLinkButton) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QCommandLinkButton) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QCommandLinkButton) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQCommandLinkButtonMousePressEvent
func callbackQCommandLinkButtonMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QCommandLinkButton::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQCommandLinkButtonFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QCommandLinkButton) MousePressEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QCommandLinkButton) MousePressEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QCommandLinkButton) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQCommandLinkButtonMouseReleaseEvent
func callbackQCommandLinkButtonMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QCommandLinkButton::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQCommandLinkButtonFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QCommandLinkButton) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QCommandLinkButton) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QCommandLinkButton) ConnectNextCheckState(f func()) {
	defer qt.Recovering("connect QCommandLinkButton::nextCheckState")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "nextCheckState", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectNextCheckState() {
	defer qt.Recovering("disconnect QCommandLinkButton::nextCheckState")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "nextCheckState")
	}
}

//export callbackQCommandLinkButtonNextCheckState
func callbackQCommandLinkButtonNextCheckState(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QCommandLinkButton::nextCheckState")

	if signal := qt.GetSignal(C.GoString(ptrName), "nextCheckState"); signal != nil {
		signal.(func())()
	} else {
		NewQCommandLinkButtonFromPointer(ptr).NextCheckStateDefault()
	}
}

func (ptr *QCommandLinkButton) NextCheckState() {
	defer qt.Recovering("QCommandLinkButton::nextCheckState")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_NextCheckState(ptr.Pointer())
	}
}

func (ptr *QCommandLinkButton) NextCheckStateDefault() {
	defer qt.Recovering("QCommandLinkButton::nextCheckState")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_NextCheckStateDefault(ptr.Pointer())
	}
}

func (ptr *QCommandLinkButton) ConnectTimerEvent(f func(e *core.QTimerEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCommandLinkButtonTimerEvent
func callbackQCommandLinkButtonTimerEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QCommandLinkButton::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
	} else {
		NewQCommandLinkButtonFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(e))
	}
}

func (ptr *QCommandLinkButton) TimerEvent(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::timerEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

func (ptr *QCommandLinkButton) TimerEventDefault(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::timerEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

func (ptr *QCommandLinkButton) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQCommandLinkButtonActionEvent
func callbackQCommandLinkButtonActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCommandLinkButton::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQCommandLinkButtonFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QCommandLinkButton) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::actionEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QCommandLinkButton) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::actionEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QCommandLinkButton) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQCommandLinkButtonDragEnterEvent
func callbackQCommandLinkButtonDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCommandLinkButton::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQCommandLinkButtonFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QCommandLinkButton) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QCommandLinkButton) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QCommandLinkButton) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQCommandLinkButtonDragLeaveEvent
func callbackQCommandLinkButtonDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCommandLinkButton::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQCommandLinkButtonFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QCommandLinkButton) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QCommandLinkButton) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QCommandLinkButton) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQCommandLinkButtonDragMoveEvent
func callbackQCommandLinkButtonDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCommandLinkButton::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQCommandLinkButtonFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QCommandLinkButton) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QCommandLinkButton) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QCommandLinkButton) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQCommandLinkButtonDropEvent
func callbackQCommandLinkButtonDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCommandLinkButton::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQCommandLinkButtonFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QCommandLinkButton) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::dropEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QCommandLinkButton) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::dropEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QCommandLinkButton) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQCommandLinkButtonEnterEvent
func callbackQCommandLinkButtonEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCommandLinkButton::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCommandLinkButtonFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCommandLinkButton) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::enterEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCommandLinkButton) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::enterEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCommandLinkButton) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQCommandLinkButtonHideEvent
func callbackQCommandLinkButtonHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCommandLinkButton::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQCommandLinkButtonFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QCommandLinkButton) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::hideEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QCommandLinkButton) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::hideEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QCommandLinkButton) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQCommandLinkButtonLeaveEvent
func callbackQCommandLinkButtonLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCommandLinkButton::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCommandLinkButtonFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCommandLinkButton) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::leaveEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCommandLinkButton) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::leaveEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCommandLinkButton) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQCommandLinkButtonMoveEvent
func callbackQCommandLinkButtonMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCommandLinkButton::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQCommandLinkButtonFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QCommandLinkButton) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::moveEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QCommandLinkButton) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::moveEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QCommandLinkButton) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QCommandLinkButton::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QCommandLinkButton::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQCommandLinkButtonSetVisible
func callbackQCommandLinkButtonSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QCommandLinkButton::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) SetVisible(visible bool) {
	defer qt.Recovering("QCommandLinkButton::setVisible")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QCommandLinkButton) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QCommandLinkButton::setVisible")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QCommandLinkButton) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQCommandLinkButtonShowEvent
func callbackQCommandLinkButtonShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCommandLinkButton::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQCommandLinkButtonFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QCommandLinkButton) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::showEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QCommandLinkButton) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::showEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QCommandLinkButton) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQCommandLinkButtonCloseEvent
func callbackQCommandLinkButtonCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCommandLinkButton::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQCommandLinkButtonFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QCommandLinkButton) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::closeEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QCommandLinkButton) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::closeEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QCommandLinkButton) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQCommandLinkButtonContextMenuEvent
func callbackQCommandLinkButtonContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCommandLinkButton::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQCommandLinkButtonFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QCommandLinkButton) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QCommandLinkButton) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QCommandLinkButton) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QCommandLinkButton::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QCommandLinkButton::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQCommandLinkButtonInitPainter
func callbackQCommandLinkButtonInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QCommandLinkButton::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQCommandLinkButtonFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QCommandLinkButton) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QCommandLinkButton::initPainter")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QCommandLinkButton) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QCommandLinkButton::initPainter")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QCommandLinkButton) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQCommandLinkButtonInputMethodEvent
func callbackQCommandLinkButtonInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCommandLinkButton::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQCommandLinkButtonFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QCommandLinkButton) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QCommandLinkButton) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QCommandLinkButton) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQCommandLinkButtonMouseDoubleClickEvent
func callbackQCommandLinkButtonMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCommandLinkButton::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQCommandLinkButtonFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QCommandLinkButton) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QCommandLinkButton) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QCommandLinkButton) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQCommandLinkButtonResizeEvent
func callbackQCommandLinkButtonResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCommandLinkButton::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQCommandLinkButtonFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QCommandLinkButton) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::resizeEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QCommandLinkButton) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::resizeEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QCommandLinkButton) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQCommandLinkButtonTabletEvent
func callbackQCommandLinkButtonTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCommandLinkButton::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQCommandLinkButtonFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QCommandLinkButton) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::tabletEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QCommandLinkButton) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::tabletEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QCommandLinkButton) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQCommandLinkButtonWheelEvent
func callbackQCommandLinkButtonWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCommandLinkButton::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQCommandLinkButtonFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QCommandLinkButton) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::wheelEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QCommandLinkButton) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::wheelEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QCommandLinkButton) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCommandLinkButtonChildEvent
func callbackQCommandLinkButtonChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCommandLinkButton::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCommandLinkButtonFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCommandLinkButton) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::childEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCommandLinkButton) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::childEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCommandLinkButton) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCommandLinkButtonCustomEvent
func callbackQCommandLinkButtonCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCommandLinkButton::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCommandLinkButtonFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCommandLinkButton) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::customEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCommandLinkButton) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCommandLinkButton::customEvent")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
