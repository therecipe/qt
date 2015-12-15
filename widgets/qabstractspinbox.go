package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QAbstractSpinBox struct {
	QWidget
}

type QAbstractSpinBox_ITF interface {
	QWidget_ITF
	QAbstractSpinBox_PTR() *QAbstractSpinBox
}

func PointerFromQAbstractSpinBox(ptr QAbstractSpinBox_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractSpinBox_PTR().Pointer()
	}
	return nil
}

func NewQAbstractSpinBoxFromPointer(ptr unsafe.Pointer) *QAbstractSpinBox {
	var n = new(QAbstractSpinBox)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractSpinBox_") {
		n.SetObjectName("QAbstractSpinBox_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractSpinBox) QAbstractSpinBox_PTR() *QAbstractSpinBox {
	return ptr
}

//QAbstractSpinBox::ButtonSymbols
type QAbstractSpinBox__ButtonSymbols int64

const (
	QAbstractSpinBox__UpDownArrows = QAbstractSpinBox__ButtonSymbols(0)
	QAbstractSpinBox__PlusMinus    = QAbstractSpinBox__ButtonSymbols(1)
	QAbstractSpinBox__NoButtons    = QAbstractSpinBox__ButtonSymbols(2)
)

//QAbstractSpinBox::CorrectionMode
type QAbstractSpinBox__CorrectionMode int64

const (
	QAbstractSpinBox__CorrectToPreviousValue = QAbstractSpinBox__CorrectionMode(0)
	QAbstractSpinBox__CorrectToNearestValue  = QAbstractSpinBox__CorrectionMode(1)
)

//QAbstractSpinBox::StepEnabledFlag
type QAbstractSpinBox__StepEnabledFlag int64

const (
	QAbstractSpinBox__StepNone        = QAbstractSpinBox__StepEnabledFlag(0x00)
	QAbstractSpinBox__StepUpEnabled   = QAbstractSpinBox__StepEnabledFlag(0x01)
	QAbstractSpinBox__StepDownEnabled = QAbstractSpinBox__StepEnabledFlag(0x02)
)

func (ptr *QAbstractSpinBox) Alignment() core.Qt__AlignmentFlag {
	defer qt.Recovering("QAbstractSpinBox::alignment")

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QAbstractSpinBox_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSpinBox) ButtonSymbols() QAbstractSpinBox__ButtonSymbols {
	defer qt.Recovering("QAbstractSpinBox::buttonSymbols")

	if ptr.Pointer() != nil {
		return QAbstractSpinBox__ButtonSymbols(C.QAbstractSpinBox_ButtonSymbols(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSpinBox) CorrectionMode() QAbstractSpinBox__CorrectionMode {
	defer qt.Recovering("QAbstractSpinBox::correctionMode")

	if ptr.Pointer() != nil {
		return QAbstractSpinBox__CorrectionMode(C.QAbstractSpinBox_CorrectionMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSpinBox) HasAcceptableInput() bool {
	defer qt.Recovering("QAbstractSpinBox::hasAcceptableInput")

	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_HasAcceptableInput(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSpinBox) HasFrame() bool {
	defer qt.Recovering("QAbstractSpinBox::hasFrame")

	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_HasFrame(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSpinBox) IsAccelerated() bool {
	defer qt.Recovering("QAbstractSpinBox::isAccelerated")

	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_IsAccelerated(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSpinBox) IsGroupSeparatorShown() bool {
	defer qt.Recovering("QAbstractSpinBox::isGroupSeparatorShown")

	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_IsGroupSeparatorShown(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSpinBox) IsReadOnly() bool {
	defer qt.Recovering("QAbstractSpinBox::isReadOnly")

	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_IsReadOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSpinBox) KeyboardTracking() bool {
	defer qt.Recovering("QAbstractSpinBox::keyboardTracking")

	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_KeyboardTracking(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSpinBox) SetAccelerated(on bool) {
	defer qt.Recovering("QAbstractSpinBox::setAccelerated")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetAccelerated(ptr.Pointer(), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QAbstractSpinBox) SetAlignment(flag core.Qt__AlignmentFlag) {
	defer qt.Recovering("QAbstractSpinBox::setAlignment")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetAlignment(ptr.Pointer(), C.int(flag))
	}
}

func (ptr *QAbstractSpinBox) SetButtonSymbols(bs QAbstractSpinBox__ButtonSymbols) {
	defer qt.Recovering("QAbstractSpinBox::setButtonSymbols")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetButtonSymbols(ptr.Pointer(), C.int(bs))
	}
}

func (ptr *QAbstractSpinBox) SetCorrectionMode(cm QAbstractSpinBox__CorrectionMode) {
	defer qt.Recovering("QAbstractSpinBox::setCorrectionMode")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetCorrectionMode(ptr.Pointer(), C.int(cm))
	}
}

func (ptr *QAbstractSpinBox) SetFrame(v bool) {
	defer qt.Recovering("QAbstractSpinBox::setFrame")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetFrame(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractSpinBox) SetGroupSeparatorShown(shown bool) {
	defer qt.Recovering("QAbstractSpinBox::setGroupSeparatorShown")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetGroupSeparatorShown(ptr.Pointer(), C.int(qt.GoBoolToInt(shown)))
	}
}

func (ptr *QAbstractSpinBox) SetKeyboardTracking(kt bool) {
	defer qt.Recovering("QAbstractSpinBox::setKeyboardTracking")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetKeyboardTracking(ptr.Pointer(), C.int(qt.GoBoolToInt(kt)))
	}
}

func (ptr *QAbstractSpinBox) SetReadOnly(r bool) {
	defer qt.Recovering("QAbstractSpinBox::setReadOnly")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetReadOnly(ptr.Pointer(), C.int(qt.GoBoolToInt(r)))
	}
}

func (ptr *QAbstractSpinBox) SetSpecialValueText(txt string) {
	defer qt.Recovering("QAbstractSpinBox::setSpecialValueText")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetSpecialValueText(ptr.Pointer(), C.CString(txt))
	}
}

func (ptr *QAbstractSpinBox) SetWrapping(w bool) {
	defer qt.Recovering("QAbstractSpinBox::setWrapping")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetWrapping(ptr.Pointer(), C.int(qt.GoBoolToInt(w)))
	}
}

func (ptr *QAbstractSpinBox) SpecialValueText() string {
	defer qt.Recovering("QAbstractSpinBox::specialValueText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractSpinBox_SpecialValueText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstractSpinBox) Text() string {
	defer qt.Recovering("QAbstractSpinBox::text")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractSpinBox_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstractSpinBox) Wrapping() bool {
	defer qt.Recovering("QAbstractSpinBox::wrapping")

	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_Wrapping(ptr.Pointer()) != 0
	}
	return false
}

func NewQAbstractSpinBox(parent QWidget_ITF) *QAbstractSpinBox {
	defer qt.Recovering("QAbstractSpinBox::QAbstractSpinBox")

	return NewQAbstractSpinBoxFromPointer(C.QAbstractSpinBox_NewQAbstractSpinBox(PointerFromQWidget(parent)))
}

func (ptr *QAbstractSpinBox) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractSpinBox::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QAbstractSpinBox::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQAbstractSpinBoxChangeEvent
func callbackQAbstractSpinBoxChangeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSpinBox::changeEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "changeEvent")
	if signal != nil {
		defer signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) ConnectClear(f func()) {
	defer qt.Recovering("connect QAbstractSpinBox::clear")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "clear", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectClear() {
	defer qt.Recovering("disconnect QAbstractSpinBox::clear")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "clear")
	}
}

//export callbackQAbstractSpinBoxClear
func callbackQAbstractSpinBoxClear(ptrName *C.char) bool {
	defer qt.Recovering("callback QAbstractSpinBox::clear")

	var signal = qt.GetSignal(C.GoString(ptrName), "clear")
	if signal != nil {
		defer signal.(func())()
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QAbstractSpinBox::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QAbstractSpinBox::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQAbstractSpinBoxCloseEvent
func callbackQAbstractSpinBoxCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSpinBox::closeEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "closeEvent")
	if signal != nil {
		defer signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QAbstractSpinBox::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QAbstractSpinBox::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQAbstractSpinBoxContextMenuEvent
func callbackQAbstractSpinBoxContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSpinBox::contextMenuEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "contextMenuEvent")
	if signal != nil {
		defer signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) ConnectEditingFinished(f func()) {
	defer qt.Recovering("connect QAbstractSpinBox::editingFinished")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_ConnectEditingFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "editingFinished", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectEditingFinished() {
	defer qt.Recovering("disconnect QAbstractSpinBox::editingFinished")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_DisconnectEditingFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "editingFinished")
	}
}

//export callbackQAbstractSpinBoxEditingFinished
func callbackQAbstractSpinBoxEditingFinished(ptrName *C.char) {
	defer qt.Recovering("callback QAbstractSpinBox::editingFinished")

	var signal = qt.GetSignal(C.GoString(ptrName), "editingFinished")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractSpinBox) Event(event core.QEvent_ITF) bool {
	defer qt.Recovering("QAbstractSpinBox::event")

	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QAbstractSpinBox) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QAbstractSpinBox::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QAbstractSpinBox::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQAbstractSpinBoxFocusInEvent
func callbackQAbstractSpinBoxFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSpinBox::focusInEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "focusInEvent")
	if signal != nil {
		defer signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QAbstractSpinBox::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QAbstractSpinBox::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQAbstractSpinBoxFocusOutEvent
func callbackQAbstractSpinBoxFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSpinBox::focusOutEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "focusOutEvent")
	if signal != nil {
		defer signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QAbstractSpinBox::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QAbstractSpinBox::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQAbstractSpinBoxHideEvent
func callbackQAbstractSpinBoxHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSpinBox::hideEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "hideEvent")
	if signal != nil {
		defer signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QAbstractSpinBox::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QAbstractSpinBox_InputMethodQuery(ptr.Pointer(), C.int(query)))
	}
	return nil
}

func (ptr *QAbstractSpinBox) InterpretText() {
	defer qt.Recovering("QAbstractSpinBox::interpretText")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_InterpretText(ptr.Pointer())
	}
}

func (ptr *QAbstractSpinBox) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QAbstractSpinBox::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QAbstractSpinBox::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQAbstractSpinBoxKeyPressEvent
func callbackQAbstractSpinBoxKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSpinBox::keyPressEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "keyPressEvent")
	if signal != nil {
		defer signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QAbstractSpinBox::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QAbstractSpinBox::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQAbstractSpinBoxKeyReleaseEvent
func callbackQAbstractSpinBoxKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSpinBox::keyReleaseEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent")
	if signal != nil {
		defer signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QAbstractSpinBox::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QAbstractSpinBox::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQAbstractSpinBoxMouseMoveEvent
func callbackQAbstractSpinBoxMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSpinBox::mouseMoveEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QAbstractSpinBox::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QAbstractSpinBox::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQAbstractSpinBoxMousePressEvent
func callbackQAbstractSpinBoxMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSpinBox::mousePressEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mousePressEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QAbstractSpinBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QAbstractSpinBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQAbstractSpinBoxMouseReleaseEvent
func callbackQAbstractSpinBoxMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSpinBox::mouseReleaseEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QAbstractSpinBox::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QAbstractSpinBox::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQAbstractSpinBoxPaintEvent
func callbackQAbstractSpinBoxPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSpinBox::paintEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "paintEvent")
	if signal != nil {
		defer signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QAbstractSpinBox::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QAbstractSpinBox::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQAbstractSpinBoxResizeEvent
func callbackQAbstractSpinBoxResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSpinBox::resizeEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "resizeEvent")
	if signal != nil {
		defer signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) SelectAll() {
	defer qt.Recovering("QAbstractSpinBox::selectAll")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SelectAll(ptr.Pointer())
	}
}

func (ptr *QAbstractSpinBox) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QAbstractSpinBox::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QAbstractSpinBox::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQAbstractSpinBoxShowEvent
func callbackQAbstractSpinBoxShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSpinBox::showEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "showEvent")
	if signal != nil {
		defer signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) ConnectStepBy(f func(steps int)) {
	defer qt.Recovering("connect QAbstractSpinBox::stepBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "stepBy", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectStepBy() {
	defer qt.Recovering("disconnect QAbstractSpinBox::stepBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "stepBy")
	}
}

//export callbackQAbstractSpinBoxStepBy
func callbackQAbstractSpinBoxStepBy(ptrName *C.char, steps C.int) bool {
	defer qt.Recovering("callback QAbstractSpinBox::stepBy")

	var signal = qt.GetSignal(C.GoString(ptrName), "stepBy")
	if signal != nil {
		defer signal.(func(int))(int(steps))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) StepDown() {
	defer qt.Recovering("QAbstractSpinBox::stepDown")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_StepDown(ptr.Pointer())
	}
}

func (ptr *QAbstractSpinBox) StepUp() {
	defer qt.Recovering("QAbstractSpinBox::stepUp")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_StepUp(ptr.Pointer())
	}
}

func (ptr *QAbstractSpinBox) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAbstractSpinBox::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAbstractSpinBox::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAbstractSpinBoxTimerEvent
func callbackQAbstractSpinBoxTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSpinBox::timerEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "timerEvent")
	if signal != nil {
		defer signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QAbstractSpinBox::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QAbstractSpinBox::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQAbstractSpinBoxWheelEvent
func callbackQAbstractSpinBoxWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSpinBox::wheelEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "wheelEvent")
	if signal != nil {
		defer signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) DestroyQAbstractSpinBox() {
	defer qt.Recovering("QAbstractSpinBox::~QAbstractSpinBox")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_DestroyQAbstractSpinBox(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
