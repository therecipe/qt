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
func callbackQDoubleSpinBoxValueChanged2(ptr unsafe.Pointer, ptrName *C.char, text *C.char) {
	defer qt.Recovering("callback QDoubleSpinBox::valueChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "valueChanged2"); signal != nil {
		signal.(func(string))(C.GoString(text))
	}

}

func (ptr *QDoubleSpinBox) ValueChanged2(text string) {
	defer qt.Recovering("QDoubleSpinBox::valueChanged")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_ValueChanged2(ptr.Pointer(), C.CString(text))
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
func callbackQDoubleSpinBoxChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDoubleSpinBox::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDoubleSpinBox) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::changeEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDoubleSpinBox) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::changeEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
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
func callbackQDoubleSpinBoxClear(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QDoubleSpinBox::clear")

	if signal := qt.GetSignal(C.GoString(ptrName), "clear"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) Clear() {
	defer qt.Recovering("QDoubleSpinBox::clear")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_Clear(ptr.Pointer())
	}
}

func (ptr *QDoubleSpinBox) ClearDefault() {
	defer qt.Recovering("QDoubleSpinBox::clear")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_ClearDefault(ptr.Pointer())
	}
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
func callbackQDoubleSpinBoxCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDoubleSpinBox::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QDoubleSpinBox) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::closeEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QDoubleSpinBox) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::closeEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
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
func callbackQDoubleSpinBoxContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDoubleSpinBox::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QDoubleSpinBox) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QDoubleSpinBox) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
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
func callbackQDoubleSpinBoxFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDoubleSpinBox::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDoubleSpinBox) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::focusInEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDoubleSpinBox) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::focusInEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
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
func callbackQDoubleSpinBoxFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDoubleSpinBox::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDoubleSpinBox) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDoubleSpinBox) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
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
func callbackQDoubleSpinBoxHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDoubleSpinBox::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QDoubleSpinBox) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::hideEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QDoubleSpinBox) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::hideEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
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
func callbackQDoubleSpinBoxKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDoubleSpinBox::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QDoubleSpinBox) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QDoubleSpinBox) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
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
func callbackQDoubleSpinBoxKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDoubleSpinBox::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QDoubleSpinBox) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QDoubleSpinBox) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
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
func callbackQDoubleSpinBoxMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDoubleSpinBox::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDoubleSpinBox) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDoubleSpinBox) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQDoubleSpinBoxMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDoubleSpinBox::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDoubleSpinBox) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDoubleSpinBox) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQDoubleSpinBoxMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDoubleSpinBox::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDoubleSpinBox) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDoubleSpinBox) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQDoubleSpinBoxPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDoubleSpinBox::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QDoubleSpinBox) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::paintEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QDoubleSpinBox) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::paintEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
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
func callbackQDoubleSpinBoxResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDoubleSpinBox::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QDoubleSpinBox) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::resizeEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QDoubleSpinBox) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::resizeEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
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
func callbackQDoubleSpinBoxShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDoubleSpinBox::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QDoubleSpinBox) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::showEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QDoubleSpinBox) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::showEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
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
func callbackQDoubleSpinBoxStepBy(ptr unsafe.Pointer, ptrName *C.char, steps C.int) {
	defer qt.Recovering("callback QDoubleSpinBox::stepBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "stepBy"); signal != nil {
		signal.(func(int))(int(steps))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).StepByDefault(int(steps))
	}
}

func (ptr *QDoubleSpinBox) StepBy(steps int) {
	defer qt.Recovering("QDoubleSpinBox::stepBy")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_StepBy(ptr.Pointer(), C.int(steps))
	}
}

func (ptr *QDoubleSpinBox) StepByDefault(steps int) {
	defer qt.Recovering("QDoubleSpinBox::stepBy")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_StepByDefault(ptr.Pointer(), C.int(steps))
	}
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
func callbackQDoubleSpinBoxTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDoubleSpinBox::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDoubleSpinBox) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::timerEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDoubleSpinBox) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::timerEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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
func callbackQDoubleSpinBoxWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDoubleSpinBox::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QDoubleSpinBox) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::wheelEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QDoubleSpinBox) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::wheelEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
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
func callbackQDoubleSpinBoxActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDoubleSpinBox::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QDoubleSpinBox) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::actionEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QDoubleSpinBox) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::actionEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
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
func callbackQDoubleSpinBoxDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDoubleSpinBox::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QDoubleSpinBox) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QDoubleSpinBox) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
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
func callbackQDoubleSpinBoxDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDoubleSpinBox::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QDoubleSpinBox) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QDoubleSpinBox) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
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
func callbackQDoubleSpinBoxDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDoubleSpinBox::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QDoubleSpinBox) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QDoubleSpinBox) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
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
func callbackQDoubleSpinBoxDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDoubleSpinBox::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QDoubleSpinBox) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::dropEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QDoubleSpinBox) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::dropEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
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
func callbackQDoubleSpinBoxEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDoubleSpinBox::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDoubleSpinBox) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::enterEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDoubleSpinBox) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::enterEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
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
func callbackQDoubleSpinBoxLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDoubleSpinBox::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDoubleSpinBox) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::leaveEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDoubleSpinBox) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::leaveEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
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
func callbackQDoubleSpinBoxMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDoubleSpinBox::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QDoubleSpinBox) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::moveEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QDoubleSpinBox) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::moveEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
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
func callbackQDoubleSpinBoxSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QDoubleSpinBox::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QDoubleSpinBox) SetVisible(visible bool) {
	defer qt.Recovering("QDoubleSpinBox::setVisible")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QDoubleSpinBox) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QDoubleSpinBox::setVisible")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
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
func callbackQDoubleSpinBoxInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QDoubleSpinBox::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QDoubleSpinBox) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QDoubleSpinBox::initPainter")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QDoubleSpinBox) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QDoubleSpinBox::initPainter")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
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
func callbackQDoubleSpinBoxInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDoubleSpinBox::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QDoubleSpinBox) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QDoubleSpinBox) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
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
func callbackQDoubleSpinBoxMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDoubleSpinBox::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDoubleSpinBox) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDoubleSpinBox) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQDoubleSpinBoxTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDoubleSpinBox::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QDoubleSpinBox) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::tabletEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QDoubleSpinBox) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::tabletEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
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
func callbackQDoubleSpinBoxChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDoubleSpinBox::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDoubleSpinBox) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::childEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDoubleSpinBox) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::childEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
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
func callbackQDoubleSpinBoxCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDoubleSpinBox::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDoubleSpinBoxFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDoubleSpinBox) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::customEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDoubleSpinBox) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDoubleSpinBox::customEvent")

	if ptr.Pointer() != nil {
		C.QDoubleSpinBox_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
