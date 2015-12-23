package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QSpinBox struct {
	QAbstractSpinBox
}

type QSpinBox_ITF interface {
	QAbstractSpinBox_ITF
	QSpinBox_PTR() *QSpinBox
}

func PointerFromQSpinBox(ptr QSpinBox_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSpinBox_PTR().Pointer()
	}
	return nil
}

func NewQSpinBoxFromPointer(ptr unsafe.Pointer) *QSpinBox {
	var n = new(QSpinBox)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSpinBox_") {
		n.SetObjectName("QSpinBox_" + qt.Identifier())
	}
	return n
}

func (ptr *QSpinBox) QSpinBox_PTR() *QSpinBox {
	return ptr
}

func (ptr *QSpinBox) CleanText() string {
	defer qt.Recovering("QSpinBox::cleanText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSpinBox_CleanText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSpinBox) DisplayIntegerBase() int {
	defer qt.Recovering("QSpinBox::displayIntegerBase")

	if ptr.Pointer() != nil {
		return int(C.QSpinBox_DisplayIntegerBase(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSpinBox) Maximum() int {
	defer qt.Recovering("QSpinBox::maximum")

	if ptr.Pointer() != nil {
		return int(C.QSpinBox_Maximum(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSpinBox) Minimum() int {
	defer qt.Recovering("QSpinBox::minimum")

	if ptr.Pointer() != nil {
		return int(C.QSpinBox_Minimum(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSpinBox) Prefix() string {
	defer qt.Recovering("QSpinBox::prefix")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSpinBox_Prefix(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSpinBox) SetDisplayIntegerBase(base int) {
	defer qt.Recovering("QSpinBox::setDisplayIntegerBase")

	if ptr.Pointer() != nil {
		C.QSpinBox_SetDisplayIntegerBase(ptr.Pointer(), C.int(base))
	}
}

func (ptr *QSpinBox) SetMaximum(max int) {
	defer qt.Recovering("QSpinBox::setMaximum")

	if ptr.Pointer() != nil {
		C.QSpinBox_SetMaximum(ptr.Pointer(), C.int(max))
	}
}

func (ptr *QSpinBox) SetMinimum(min int) {
	defer qt.Recovering("QSpinBox::setMinimum")

	if ptr.Pointer() != nil {
		C.QSpinBox_SetMinimum(ptr.Pointer(), C.int(min))
	}
}

func (ptr *QSpinBox) SetPrefix(prefix string) {
	defer qt.Recovering("QSpinBox::setPrefix")

	if ptr.Pointer() != nil {
		C.QSpinBox_SetPrefix(ptr.Pointer(), C.CString(prefix))
	}
}

func (ptr *QSpinBox) SetSingleStep(val int) {
	defer qt.Recovering("QSpinBox::setSingleStep")

	if ptr.Pointer() != nil {
		C.QSpinBox_SetSingleStep(ptr.Pointer(), C.int(val))
	}
}

func (ptr *QSpinBox) SetSuffix(suffix string) {
	defer qt.Recovering("QSpinBox::setSuffix")

	if ptr.Pointer() != nil {
		C.QSpinBox_SetSuffix(ptr.Pointer(), C.CString(suffix))
	}
}

func (ptr *QSpinBox) SetValue(val int) {
	defer qt.Recovering("QSpinBox::setValue")

	if ptr.Pointer() != nil {
		C.QSpinBox_SetValue(ptr.Pointer(), C.int(val))
	}
}

func (ptr *QSpinBox) SingleStep() int {
	defer qt.Recovering("QSpinBox::singleStep")

	if ptr.Pointer() != nil {
		return int(C.QSpinBox_SingleStep(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSpinBox) Suffix() string {
	defer qt.Recovering("QSpinBox::suffix")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSpinBox_Suffix(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSpinBox) Value() int {
	defer qt.Recovering("QSpinBox::value")

	if ptr.Pointer() != nil {
		return int(C.QSpinBox_Value(ptr.Pointer()))
	}
	return 0
}

func NewQSpinBox(parent QWidget_ITF) *QSpinBox {
	defer qt.Recovering("QSpinBox::QSpinBox")

	return NewQSpinBoxFromPointer(C.QSpinBox_NewQSpinBox(PointerFromQWidget(parent)))
}

func (ptr *QSpinBox) SetRange(minimum int, maximum int) {
	defer qt.Recovering("QSpinBox::setRange")

	if ptr.Pointer() != nil {
		C.QSpinBox_SetRange(ptr.Pointer(), C.int(minimum), C.int(maximum))
	}
}

func (ptr *QSpinBox) ConnectValueChanged2(f func(text string)) {
	defer qt.Recovering("connect QSpinBox::valueChanged")

	if ptr.Pointer() != nil {
		C.QSpinBox_ConnectValueChanged2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "valueChanged2", f)
	}
}

func (ptr *QSpinBox) DisconnectValueChanged2() {
	defer qt.Recovering("disconnect QSpinBox::valueChanged")

	if ptr.Pointer() != nil {
		C.QSpinBox_DisconnectValueChanged2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "valueChanged2")
	}
}

//export callbackQSpinBoxValueChanged2
func callbackQSpinBoxValueChanged2(ptrName *C.char, text *C.char) {
	defer qt.Recovering("callback QSpinBox::valueChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "valueChanged2"); signal != nil {
		signal.(func(string))(C.GoString(text))
	}

}

func (ptr *QSpinBox) ConnectValueChanged(f func(i int)) {
	defer qt.Recovering("connect QSpinBox::valueChanged")

	if ptr.Pointer() != nil {
		C.QSpinBox_ConnectValueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "valueChanged", f)
	}
}

func (ptr *QSpinBox) DisconnectValueChanged() {
	defer qt.Recovering("disconnect QSpinBox::valueChanged")

	if ptr.Pointer() != nil {
		C.QSpinBox_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "valueChanged")
	}
}

//export callbackQSpinBoxValueChanged
func callbackQSpinBoxValueChanged(ptrName *C.char, i C.int) {
	defer qt.Recovering("callback QSpinBox::valueChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "valueChanged"); signal != nil {
		signal.(func(int))(int(i))
	}

}

func (ptr *QSpinBox) DestroyQSpinBox() {
	defer qt.Recovering("QSpinBox::~QSpinBox")

	if ptr.Pointer() != nil {
		C.QSpinBox_DestroyQSpinBox(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSpinBox) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSpinBox::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QSpinBox) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QSpinBox::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQSpinBoxChangeEvent
func callbackQSpinBoxChangeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSpinBox::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectClear(f func()) {
	defer qt.Recovering("connect QSpinBox::clear")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "clear", f)
	}
}

func (ptr *QSpinBox) DisconnectClear() {
	defer qt.Recovering("disconnect QSpinBox::clear")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "clear")
	}
}

//export callbackQSpinBoxClear
func callbackQSpinBoxClear(ptrName *C.char) bool {
	defer qt.Recovering("callback QSpinBox::clear")

	if signal := qt.GetSignal(C.GoString(ptrName), "clear"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QSpinBox::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QSpinBox) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QSpinBox::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQSpinBoxCloseEvent
func callbackQSpinBoxCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSpinBox::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QSpinBox::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QSpinBox) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QSpinBox::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQSpinBoxContextMenuEvent
func callbackQSpinBoxContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSpinBox::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QSpinBox::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QSpinBox) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QSpinBox::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQSpinBoxFocusInEvent
func callbackQSpinBoxFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSpinBox::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QSpinBox::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QSpinBox) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QSpinBox::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQSpinBoxFocusOutEvent
func callbackQSpinBoxFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSpinBox::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QSpinBox::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QSpinBox) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QSpinBox::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQSpinBoxHideEvent
func callbackQSpinBoxHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSpinBox::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QSpinBox::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QSpinBox) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QSpinBox::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQSpinBoxKeyPressEvent
func callbackQSpinBoxKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSpinBox::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QSpinBox::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QSpinBox) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QSpinBox::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQSpinBoxKeyReleaseEvent
func callbackQSpinBoxKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSpinBox::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSpinBox::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QSpinBox) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QSpinBox::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQSpinBoxMouseMoveEvent
func callbackQSpinBoxMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSpinBox::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSpinBox::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QSpinBox) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QSpinBox::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQSpinBoxMousePressEvent
func callbackQSpinBoxMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSpinBox::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSpinBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QSpinBox) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QSpinBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQSpinBoxMouseReleaseEvent
func callbackQSpinBoxMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSpinBox::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QSpinBox::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QSpinBox) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QSpinBox::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQSpinBoxPaintEvent
func callbackQSpinBoxPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSpinBox::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QSpinBox::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QSpinBox) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QSpinBox::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQSpinBoxResizeEvent
func callbackQSpinBoxResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSpinBox::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QSpinBox::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QSpinBox) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QSpinBox::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQSpinBoxShowEvent
func callbackQSpinBoxShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSpinBox::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectStepBy(f func(steps int)) {
	defer qt.Recovering("connect QSpinBox::stepBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "stepBy", f)
	}
}

func (ptr *QSpinBox) DisconnectStepBy() {
	defer qt.Recovering("disconnect QSpinBox::stepBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "stepBy")
	}
}

//export callbackQSpinBoxStepBy
func callbackQSpinBoxStepBy(ptrName *C.char, steps C.int) bool {
	defer qt.Recovering("callback QSpinBox::stepBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "stepBy"); signal != nil {
		signal.(func(int))(int(steps))
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSpinBox::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSpinBox) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSpinBox::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSpinBoxTimerEvent
func callbackQSpinBoxTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSpinBox::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QSpinBox::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QSpinBox) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QSpinBox::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQSpinBoxWheelEvent
func callbackQSpinBoxWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSpinBox::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QSpinBox::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QSpinBox) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QSpinBox::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQSpinBoxActionEvent
func callbackQSpinBoxActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSpinBox::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QSpinBox::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QSpinBox) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QSpinBox::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQSpinBoxDragEnterEvent
func callbackQSpinBoxDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSpinBox::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QSpinBox::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QSpinBox) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QSpinBox::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQSpinBoxDragLeaveEvent
func callbackQSpinBoxDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSpinBox::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QSpinBox::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QSpinBox) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QSpinBox::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQSpinBoxDragMoveEvent
func callbackQSpinBoxDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSpinBox::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QSpinBox::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QSpinBox) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QSpinBox::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQSpinBoxDropEvent
func callbackQSpinBoxDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSpinBox::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSpinBox::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QSpinBox) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QSpinBox::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQSpinBoxEnterEvent
func callbackQSpinBoxEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSpinBox::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSpinBox::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QSpinBox) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QSpinBox::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQSpinBoxLeaveEvent
func callbackQSpinBoxLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSpinBox::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QSpinBox::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QSpinBox) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QSpinBox::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQSpinBoxMoveEvent
func callbackQSpinBoxMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSpinBox::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QSpinBox::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QSpinBox) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QSpinBox::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQSpinBoxSetVisible
func callbackQSpinBoxSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QSpinBox::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QSpinBox::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QSpinBox) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QSpinBox::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQSpinBoxInitPainter
func callbackQSpinBoxInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QSpinBox::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QSpinBox::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QSpinBox) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QSpinBox::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQSpinBoxInputMethodEvent
func callbackQSpinBoxInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSpinBox::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSpinBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QSpinBox) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QSpinBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQSpinBoxMouseDoubleClickEvent
func callbackQSpinBoxMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSpinBox::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QSpinBox::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QSpinBox) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QSpinBox::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQSpinBoxTabletEvent
func callbackQSpinBoxTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSpinBox::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSpinBox::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSpinBox) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSpinBox::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSpinBoxChildEvent
func callbackQSpinBoxChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSpinBox::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSpinBox) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSpinBox::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSpinBox) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSpinBox::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSpinBoxCustomEvent
func callbackQSpinBoxCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSpinBox::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
