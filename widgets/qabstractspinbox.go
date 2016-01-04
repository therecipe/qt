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
func callbackQAbstractSpinBoxChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSpinBox::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractSpinBox) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::changeEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractSpinBox) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::changeEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
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
func callbackQAbstractSpinBoxClear(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QAbstractSpinBox::clear")

	if signal := qt.GetSignal(C.GoString(ptrName), "clear"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) Clear() {
	defer qt.Recovering("QAbstractSpinBox::clear")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_Clear(ptr.Pointer())
	}
}

func (ptr *QAbstractSpinBox) ClearDefault() {
	defer qt.Recovering("QAbstractSpinBox::clear")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_ClearDefault(ptr.Pointer())
	}
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
func callbackQAbstractSpinBoxCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSpinBox::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QAbstractSpinBox) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::closeEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QAbstractSpinBox) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::closeEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
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
func callbackQAbstractSpinBoxContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSpinBox::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QAbstractSpinBox) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QAbstractSpinBox) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
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
func callbackQAbstractSpinBoxEditingFinished(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractSpinBox::editingFinished")

	if signal := qt.GetSignal(C.GoString(ptrName), "editingFinished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractSpinBox) EditingFinished() {
	defer qt.Recovering("QAbstractSpinBox::editingFinished")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_EditingFinished(ptr.Pointer())
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
func callbackQAbstractSpinBoxFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSpinBox::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QAbstractSpinBox) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::focusInEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QAbstractSpinBox) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::focusInEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
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
func callbackQAbstractSpinBoxFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSpinBox::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QAbstractSpinBox) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QAbstractSpinBox) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
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
func callbackQAbstractSpinBoxHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSpinBox::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QAbstractSpinBox) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::hideEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QAbstractSpinBox) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::hideEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
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
func callbackQAbstractSpinBoxKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSpinBox::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QAbstractSpinBox) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QAbstractSpinBox) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
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
func callbackQAbstractSpinBoxKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSpinBox::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QAbstractSpinBox) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QAbstractSpinBox) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
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
func callbackQAbstractSpinBoxMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSpinBox::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QAbstractSpinBox) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QAbstractSpinBox) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQAbstractSpinBoxMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSpinBox::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QAbstractSpinBox) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QAbstractSpinBox) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQAbstractSpinBoxMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSpinBox::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QAbstractSpinBox) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QAbstractSpinBox) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQAbstractSpinBoxPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSpinBox::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QAbstractSpinBox) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::paintEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QAbstractSpinBox) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::paintEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
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
func callbackQAbstractSpinBoxResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSpinBox::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QAbstractSpinBox) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::resizeEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QAbstractSpinBox) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::resizeEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
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
func callbackQAbstractSpinBoxShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSpinBox::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QAbstractSpinBox) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::showEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QAbstractSpinBox) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::showEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
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
func callbackQAbstractSpinBoxStepBy(ptr unsafe.Pointer, ptrName *C.char, steps C.int) {
	defer qt.Recovering("callback QAbstractSpinBox::stepBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "stepBy"); signal != nil {
		signal.(func(int))(int(steps))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).StepByDefault(int(steps))
	}
}

func (ptr *QAbstractSpinBox) StepBy(steps int) {
	defer qt.Recovering("QAbstractSpinBox::stepBy")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_StepBy(ptr.Pointer(), C.int(steps))
	}
}

func (ptr *QAbstractSpinBox) StepByDefault(steps int) {
	defer qt.Recovering("QAbstractSpinBox::stepBy")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_StepByDefault(ptr.Pointer(), C.int(steps))
	}
}

func (ptr *QAbstractSpinBox) StepDown() {
	defer qt.Recovering("QAbstractSpinBox::stepDown")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_StepDown(ptr.Pointer())
	}
}

func (ptr *QAbstractSpinBox) StepEnabled() QAbstractSpinBox__StepEnabledFlag {
	defer qt.Recovering("QAbstractSpinBox::stepEnabled")

	if ptr.Pointer() != nil {
		return QAbstractSpinBox__StepEnabledFlag(C.QAbstractSpinBox_StepEnabled(ptr.Pointer()))
	}
	return 0
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
func callbackQAbstractSpinBoxTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSpinBox::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractSpinBox) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractSpinBox) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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
func callbackQAbstractSpinBoxWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSpinBox::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QAbstractSpinBox) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::wheelEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QAbstractSpinBox) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::wheelEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
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
func callbackQAbstractSpinBoxActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSpinBox::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QAbstractSpinBox) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::actionEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QAbstractSpinBox) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::actionEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
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
func callbackQAbstractSpinBoxDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSpinBox::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QAbstractSpinBox) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QAbstractSpinBox) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
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
func callbackQAbstractSpinBoxDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSpinBox::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QAbstractSpinBox) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QAbstractSpinBox) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
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
func callbackQAbstractSpinBoxDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSpinBox::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QAbstractSpinBox) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QAbstractSpinBox) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
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
func callbackQAbstractSpinBoxDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSpinBox::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QAbstractSpinBox) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::dropEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QAbstractSpinBox) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::dropEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
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
func callbackQAbstractSpinBoxEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSpinBox::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractSpinBox) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::enterEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractSpinBox) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::enterEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
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
func callbackQAbstractSpinBoxLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSpinBox::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractSpinBox) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::leaveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractSpinBox) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::leaveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
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
func callbackQAbstractSpinBoxMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSpinBox::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QAbstractSpinBox) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::moveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QAbstractSpinBox) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::moveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
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
func callbackQAbstractSpinBoxSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QAbstractSpinBox::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QAbstractSpinBox) SetVisible(visible bool) {
	defer qt.Recovering("QAbstractSpinBox::setVisible")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QAbstractSpinBox) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QAbstractSpinBox::setVisible")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
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
func callbackQAbstractSpinBoxInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSpinBox::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QAbstractSpinBox) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QAbstractSpinBox::initPainter")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QAbstractSpinBox) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QAbstractSpinBox::initPainter")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
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
func callbackQAbstractSpinBoxInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSpinBox::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QAbstractSpinBox) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QAbstractSpinBox) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
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
func callbackQAbstractSpinBoxMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSpinBox::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QAbstractSpinBox) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QAbstractSpinBox) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQAbstractSpinBoxTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSpinBox::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QAbstractSpinBox) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::tabletEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QAbstractSpinBox) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::tabletEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
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
func callbackQAbstractSpinBoxChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSpinBox::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractSpinBox) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractSpinBox) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
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
func callbackQAbstractSpinBoxCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSpinBox::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractSpinBoxFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractSpinBox) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractSpinBox) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractSpinBox::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
