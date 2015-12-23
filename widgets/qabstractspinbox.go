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

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "clear"); signal != nil {
		signal.(func())()
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

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "editingFinished"); signal != nil {
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

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QAbstractSpinBox::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QAbstractSpinBox_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
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

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) SizeHint() *core.QSize {
	defer qt.Recovering("QAbstractSpinBox::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QAbstractSpinBox_SizeHint(ptr.Pointer()))
	}
	return nil
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

	if signal := qt.GetSignal(C.GoString(ptrName), "stepBy"); signal != nil {
		signal.(func(int))(int(steps))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
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

func (ptr *QAbstractSpinBox) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QAbstractSpinBox::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QAbstractSpinBox::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQAbstractSpinBoxActionEvent
func callbackQAbstractSpinBoxActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSpinBox::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QAbstractSpinBox::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QAbstractSpinBox::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQAbstractSpinBoxDragEnterEvent
func callbackQAbstractSpinBoxDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSpinBox::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QAbstractSpinBox::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QAbstractSpinBox::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQAbstractSpinBoxDragLeaveEvent
func callbackQAbstractSpinBoxDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSpinBox::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QAbstractSpinBox::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QAbstractSpinBox::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQAbstractSpinBoxDragMoveEvent
func callbackQAbstractSpinBoxDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSpinBox::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QAbstractSpinBox::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QAbstractSpinBox::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQAbstractSpinBoxDropEvent
func callbackQAbstractSpinBoxDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSpinBox::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractSpinBox::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QAbstractSpinBox::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQAbstractSpinBoxEnterEvent
func callbackQAbstractSpinBoxEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSpinBox::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractSpinBox::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QAbstractSpinBox::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQAbstractSpinBoxLeaveEvent
func callbackQAbstractSpinBoxLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSpinBox::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QAbstractSpinBox::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QAbstractSpinBox::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQAbstractSpinBoxMoveEvent
func callbackQAbstractSpinBoxMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSpinBox::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QAbstractSpinBox::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QAbstractSpinBox::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQAbstractSpinBoxSetVisible
func callbackQAbstractSpinBoxSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QAbstractSpinBox::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QAbstractSpinBox::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QAbstractSpinBox::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQAbstractSpinBoxInitPainter
func callbackQAbstractSpinBoxInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSpinBox::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QAbstractSpinBox::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QAbstractSpinBox::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQAbstractSpinBoxInputMethodEvent
func callbackQAbstractSpinBoxInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSpinBox::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QAbstractSpinBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QAbstractSpinBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQAbstractSpinBoxMouseDoubleClickEvent
func callbackQAbstractSpinBoxMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSpinBox::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QAbstractSpinBox::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QAbstractSpinBox::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQAbstractSpinBoxTabletEvent
func callbackQAbstractSpinBoxTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSpinBox::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAbstractSpinBox::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAbstractSpinBox::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAbstractSpinBoxChildEvent
func callbackQAbstractSpinBoxChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSpinBox::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractSpinBox::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAbstractSpinBox::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAbstractSpinBoxCustomEvent
func callbackQAbstractSpinBoxCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSpinBox::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
