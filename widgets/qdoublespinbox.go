package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QDoubleSpinBox struct {
	QAbstractSpinBox
}

type QDoubleSpinBox_ITF interface {
	QAbstractSpinBox_ITF
	QDoubleSpinBox_PTR() *QDoubleSpinBox
}

func PointerFromQDoubleSpinBox(ptr QDoubleSpinBox_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDoubleSpinBox_PTR().Pointer()
	}
	return nil
}

func NewQDoubleSpinBoxFromPointer(ptr unsafe.Pointer) *QDoubleSpinBox {
	var n = new(QDoubleSpinBox)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDoubleSpinBox_") {
		n.SetObjectName("QDoubleSpinBox_" + qt.Identifier())
	}
	return n
}

func (ptr *QDoubleSpinBox) QDoubleSpinBox_PTR() *QDoubleSpinBox {
	return ptr
}

func (ptr *QDoubleSpinBox) CleanText() string {
	defer qt.Recovering("QDoubleSpinBox::cleanText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDoubleSpinBox_CleanText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDoubleSpinBox) Decimals() int {
	defer qt.Recovering("QDoubleSpinBox::decimals")

	if ptr.Pointer() != nil {
		return int(C.QDoubleSpinBox_Decimals(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDoubleSpinBox) Prefix() string {
	defer qt.Recovering("QDoubleSpinBox::prefix")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDoubleSpinBox_Prefix(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDoubleSpinBox) SetDecimals(prec int) {
	defer qt.Recovering("QDoubleSpinBox::setDecimals")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_SetDecimals(ptr.Pointer(), C.int(prec))
	}
}

func (ptr *QDoubleSpinBox) SetPrefix(prefix string) {
	defer qt.Recovering("QDoubleSpinBox::setPrefix")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_SetPrefix(ptr.Pointer(), C.CString(prefix))
	}
}

func (ptr *QDoubleSpinBox) SetSuffix(suffix string) {
	defer qt.Recovering("QDoubleSpinBox::setSuffix")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_SetSuffix(ptr.Pointer(), C.CString(suffix))
	}
}

func (ptr *QDoubleSpinBox) Suffix() string {
	defer qt.Recovering("QDoubleSpinBox::suffix")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDoubleSpinBox_Suffix(ptr.Pointer()))
	}
	return ""
}

func NewQDoubleSpinBox(parent QWidget_ITF) *QDoubleSpinBox {
	defer qt.Recovering("QDoubleSpinBox::QDoubleSpinBox")

	return NewQDoubleSpinBoxFromPointer(C.QDoubleSpinBox_NewQDoubleSpinBox(PointerFromQWidget(parent)))
}

func (ptr *QDoubleSpinBox) ConnectValueChanged2(f func(text string)) {
	defer qt.Recovering("connect QDoubleSpinBox::valueChanged")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_ConnectValueChanged2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "valueChanged2", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectValueChanged2() {
	defer qt.Recovering("disconnect QDoubleSpinBox::valueChanged")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_DisconnectValueChanged2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "valueChanged2")
	}
}

//export callbackQDoubleSpinBoxValueChanged2
func callbackQDoubleSpinBoxValueChanged2(ptrName *C.char, text *C.char) {
	defer qt.Recovering("callback QDoubleSpinBox::valueChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "valueChanged2"); signal != nil {
		signal.(func(string))(C.GoString(text))
	}

}

func (ptr *QDoubleSpinBox) DestroyQDoubleSpinBox() {
	defer qt.Recovering("QDoubleSpinBox::~QDoubleSpinBox")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_DestroyQDoubleSpinBox(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDoubleSpinBox) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDoubleSpinBox::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QDoubleSpinBox::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQDoubleSpinBoxChangeEvent
func callbackQDoubleSpinBoxChangeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleSpinBox::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectClear(f func()) {
	defer qt.Recovering("connect QDoubleSpinBox::clear")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "clear", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectClear() {
	defer qt.Recovering("disconnect QDoubleSpinBox::clear")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "clear")
	}
}

//export callbackQDoubleSpinBoxClear
func callbackQDoubleSpinBoxClear(ptrName *C.char) bool {
	defer qt.Recovering("callback QDoubleSpinBox::clear")

	if signal := qt.GetSignal(C.GoString(ptrName), "clear"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QDoubleSpinBox::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QDoubleSpinBox::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQDoubleSpinBoxCloseEvent
func callbackQDoubleSpinBoxCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleSpinBox::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QDoubleSpinBox::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QDoubleSpinBox::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQDoubleSpinBoxContextMenuEvent
func callbackQDoubleSpinBoxContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleSpinBox::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QDoubleSpinBox::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QDoubleSpinBox::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQDoubleSpinBoxFocusInEvent
func callbackQDoubleSpinBoxFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleSpinBox::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QDoubleSpinBox::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QDoubleSpinBox::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQDoubleSpinBoxFocusOutEvent
func callbackQDoubleSpinBoxFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleSpinBox::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QDoubleSpinBox::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QDoubleSpinBox::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQDoubleSpinBoxHideEvent
func callbackQDoubleSpinBoxHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleSpinBox::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QDoubleSpinBox::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QDoubleSpinBox::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQDoubleSpinBoxKeyPressEvent
func callbackQDoubleSpinBoxKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleSpinBox::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QDoubleSpinBox::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QDoubleSpinBox::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQDoubleSpinBoxKeyReleaseEvent
func callbackQDoubleSpinBoxKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleSpinBox::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDoubleSpinBox::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QDoubleSpinBox::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQDoubleSpinBoxMouseMoveEvent
func callbackQDoubleSpinBoxMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleSpinBox::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDoubleSpinBox::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QDoubleSpinBox::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQDoubleSpinBoxMousePressEvent
func callbackQDoubleSpinBoxMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleSpinBox::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDoubleSpinBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QDoubleSpinBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQDoubleSpinBoxMouseReleaseEvent
func callbackQDoubleSpinBoxMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleSpinBox::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QDoubleSpinBox::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QDoubleSpinBox::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQDoubleSpinBoxPaintEvent
func callbackQDoubleSpinBoxPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleSpinBox::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QDoubleSpinBox::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QDoubleSpinBox::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQDoubleSpinBoxResizeEvent
func callbackQDoubleSpinBoxResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleSpinBox::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QDoubleSpinBox::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QDoubleSpinBox::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQDoubleSpinBoxShowEvent
func callbackQDoubleSpinBoxShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleSpinBox::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectStepBy(f func(steps int)) {
	defer qt.Recovering("connect QDoubleSpinBox::stepBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "stepBy", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectStepBy() {
	defer qt.Recovering("disconnect QDoubleSpinBox::stepBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "stepBy")
	}
}

//export callbackQDoubleSpinBoxStepBy
func callbackQDoubleSpinBoxStepBy(ptrName *C.char, steps C.int) bool {
	defer qt.Recovering("callback QDoubleSpinBox::stepBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "stepBy"); signal != nil {
		signal.(func(int))(int(steps))
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDoubleSpinBox::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDoubleSpinBox::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDoubleSpinBoxTimerEvent
func callbackQDoubleSpinBoxTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleSpinBox::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QDoubleSpinBox::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QDoubleSpinBox::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQDoubleSpinBoxWheelEvent
func callbackQDoubleSpinBoxWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleSpinBox::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QDoubleSpinBox::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QDoubleSpinBox::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQDoubleSpinBoxActionEvent
func callbackQDoubleSpinBoxActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleSpinBox::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QDoubleSpinBox::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QDoubleSpinBox::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQDoubleSpinBoxDragEnterEvent
func callbackQDoubleSpinBoxDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleSpinBox::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QDoubleSpinBox::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QDoubleSpinBox::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQDoubleSpinBoxDragLeaveEvent
func callbackQDoubleSpinBoxDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleSpinBox::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QDoubleSpinBox::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QDoubleSpinBox::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQDoubleSpinBoxDragMoveEvent
func callbackQDoubleSpinBoxDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleSpinBox::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QDoubleSpinBox::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QDoubleSpinBox::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQDoubleSpinBoxDropEvent
func callbackQDoubleSpinBoxDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleSpinBox::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDoubleSpinBox::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QDoubleSpinBox::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQDoubleSpinBoxEnterEvent
func callbackQDoubleSpinBoxEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleSpinBox::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDoubleSpinBox::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QDoubleSpinBox::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQDoubleSpinBoxLeaveEvent
func callbackQDoubleSpinBoxLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleSpinBox::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QDoubleSpinBox::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QDoubleSpinBox::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQDoubleSpinBoxMoveEvent
func callbackQDoubleSpinBoxMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleSpinBox::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QDoubleSpinBox::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QDoubleSpinBox::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQDoubleSpinBoxSetVisible
func callbackQDoubleSpinBoxSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QDoubleSpinBox::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QDoubleSpinBox::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QDoubleSpinBox::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQDoubleSpinBoxInitPainter
func callbackQDoubleSpinBoxInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleSpinBox::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QDoubleSpinBox::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QDoubleSpinBox::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQDoubleSpinBoxInputMethodEvent
func callbackQDoubleSpinBoxInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleSpinBox::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDoubleSpinBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QDoubleSpinBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQDoubleSpinBoxMouseDoubleClickEvent
func callbackQDoubleSpinBoxMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleSpinBox::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QDoubleSpinBox::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QDoubleSpinBox::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQDoubleSpinBoxTabletEvent
func callbackQDoubleSpinBoxTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleSpinBox::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDoubleSpinBox::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDoubleSpinBox::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDoubleSpinBoxChildEvent
func callbackQDoubleSpinBoxChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleSpinBox::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDoubleSpinBox::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDoubleSpinBox) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDoubleSpinBox::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDoubleSpinBoxCustomEvent
func callbackQDoubleSpinBoxCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDoubleSpinBox::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
