package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"strings"
	"unsafe"
)

type QComboBox struct {
	QWidget
}

type QComboBox_ITF interface {
	QWidget_ITF
	QComboBox_PTR() *QComboBox
}

func PointerFromQComboBox(ptr QComboBox_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QComboBox_PTR().Pointer()
	}
	return nil
}

func NewQComboBoxFromPointer(ptr unsafe.Pointer) *QComboBox {
	var n = new(QComboBox)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QComboBox_") {
		n.SetObjectName("QComboBox_" + qt.Identifier())
	}
	return n
}

func (ptr *QComboBox) QComboBox_PTR() *QComboBox {
	return ptr
}

//QComboBox::InsertPolicy
type QComboBox__InsertPolicy int64

const (
	QComboBox__NoInsert             = QComboBox__InsertPolicy(0)
	QComboBox__InsertAtTop          = QComboBox__InsertPolicy(1)
	QComboBox__InsertAtCurrent      = QComboBox__InsertPolicy(2)
	QComboBox__InsertAtBottom       = QComboBox__InsertPolicy(3)
	QComboBox__InsertAfterCurrent   = QComboBox__InsertPolicy(4)
	QComboBox__InsertBeforeCurrent  = QComboBox__InsertPolicy(5)
	QComboBox__InsertAlphabetically = QComboBox__InsertPolicy(6)
)

//QComboBox::SizeAdjustPolicy
type QComboBox__SizeAdjustPolicy int64

const (
	QComboBox__AdjustToContents                      = QComboBox__SizeAdjustPolicy(0)
	QComboBox__AdjustToContentsOnFirstShow           = QComboBox__SizeAdjustPolicy(1)
	QComboBox__AdjustToMinimumContentsLength         = QComboBox__SizeAdjustPolicy(2)
	QComboBox__AdjustToMinimumContentsLengthWithIcon = QComboBox__SizeAdjustPolicy(3)
)

func (ptr *QComboBox) Count() int {
	defer qt.Recovering("QComboBox::count")

	if ptr.Pointer() != nil {
		return int(C.QComboBox_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QComboBox) CurrentData(role int) *core.QVariant {
	defer qt.Recovering("QComboBox::currentData")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QComboBox_CurrentData(ptr.Pointer(), C.int(role)))
	}
	return nil
}

func (ptr *QComboBox) CurrentIndex() int {
	defer qt.Recovering("QComboBox::currentIndex")

	if ptr.Pointer() != nil {
		return int(C.QComboBox_CurrentIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QComboBox) CurrentText() string {
	defer qt.Recovering("QComboBox::currentText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QComboBox_CurrentText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QComboBox) DuplicatesEnabled() bool {
	defer qt.Recovering("QComboBox::duplicatesEnabled")

	if ptr.Pointer() != nil {
		return C.QComboBox_DuplicatesEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QComboBox) HasFrame() bool {
	defer qt.Recovering("QComboBox::hasFrame")

	if ptr.Pointer() != nil {
		return C.QComboBox_HasFrame(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QComboBox) IconSize() *core.QSize {
	defer qt.Recovering("QComboBox::iconSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QComboBox_IconSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QComboBox) InsertPolicy() QComboBox__InsertPolicy {
	defer qt.Recovering("QComboBox::insertPolicy")

	if ptr.Pointer() != nil {
		return QComboBox__InsertPolicy(C.QComboBox_InsertPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QComboBox) IsEditable() bool {
	defer qt.Recovering("QComboBox::isEditable")

	if ptr.Pointer() != nil {
		return C.QComboBox_IsEditable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QComboBox) MaxCount() int {
	defer qt.Recovering("QComboBox::maxCount")

	if ptr.Pointer() != nil {
		return int(C.QComboBox_MaxCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QComboBox) MaxVisibleItems() int {
	defer qt.Recovering("QComboBox::maxVisibleItems")

	if ptr.Pointer() != nil {
		return int(C.QComboBox_MaxVisibleItems(ptr.Pointer()))
	}
	return 0
}

func (ptr *QComboBox) MinimumContentsLength() int {
	defer qt.Recovering("QComboBox::minimumContentsLength")

	if ptr.Pointer() != nil {
		return int(C.QComboBox_MinimumContentsLength(ptr.Pointer()))
	}
	return 0
}

func (ptr *QComboBox) ModelColumn() int {
	defer qt.Recovering("QComboBox::modelColumn")

	if ptr.Pointer() != nil {
		return int(C.QComboBox_ModelColumn(ptr.Pointer()))
	}
	return 0
}

func (ptr *QComboBox) SetCompleter(completer QCompleter_ITF) {
	defer qt.Recovering("QComboBox::setCompleter")

	if ptr.Pointer() != nil {
		C.QComboBox_SetCompleter(ptr.Pointer(), PointerFromQCompleter(completer))
	}
}

func (ptr *QComboBox) SetCurrentIndex(index int) {
	defer qt.Recovering("QComboBox::setCurrentIndex")

	if ptr.Pointer() != nil {
		C.QComboBox_SetCurrentIndex(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QComboBox) SetCurrentText(text string) {
	defer qt.Recovering("QComboBox::setCurrentText")

	if ptr.Pointer() != nil {
		C.QComboBox_SetCurrentText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QComboBox) SetDuplicatesEnabled(enable bool) {
	defer qt.Recovering("QComboBox::setDuplicatesEnabled")

	if ptr.Pointer() != nil {
		C.QComboBox_SetDuplicatesEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QComboBox) SetEditable(editable bool) {
	defer qt.Recovering("QComboBox::setEditable")

	if ptr.Pointer() != nil {
		C.QComboBox_SetEditable(ptr.Pointer(), C.int(qt.GoBoolToInt(editable)))
	}
}

func (ptr *QComboBox) SetFrame(v bool) {
	defer qt.Recovering("QComboBox::setFrame")

	if ptr.Pointer() != nil {
		C.QComboBox_SetFrame(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QComboBox) SetIconSize(size core.QSize_ITF) {
	defer qt.Recovering("QComboBox::setIconSize")

	if ptr.Pointer() != nil {
		C.QComboBox_SetIconSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QComboBox) SetInsertPolicy(policy QComboBox__InsertPolicy) {
	defer qt.Recovering("QComboBox::setInsertPolicy")

	if ptr.Pointer() != nil {
		C.QComboBox_SetInsertPolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QComboBox) SetMaxCount(max int) {
	defer qt.Recovering("QComboBox::setMaxCount")

	if ptr.Pointer() != nil {
		C.QComboBox_SetMaxCount(ptr.Pointer(), C.int(max))
	}
}

func (ptr *QComboBox) SetMaxVisibleItems(maxItems int) {
	defer qt.Recovering("QComboBox::setMaxVisibleItems")

	if ptr.Pointer() != nil {
		C.QComboBox_SetMaxVisibleItems(ptr.Pointer(), C.int(maxItems))
	}
}

func (ptr *QComboBox) SetMinimumContentsLength(characters int) {
	defer qt.Recovering("QComboBox::setMinimumContentsLength")

	if ptr.Pointer() != nil {
		C.QComboBox_SetMinimumContentsLength(ptr.Pointer(), C.int(characters))
	}
}

func (ptr *QComboBox) SetModelColumn(visibleColumn int) {
	defer qt.Recovering("QComboBox::setModelColumn")

	if ptr.Pointer() != nil {
		C.QComboBox_SetModelColumn(ptr.Pointer(), C.int(visibleColumn))
	}
}

func (ptr *QComboBox) SetSizeAdjustPolicy(policy QComboBox__SizeAdjustPolicy) {
	defer qt.Recovering("QComboBox::setSizeAdjustPolicy")

	if ptr.Pointer() != nil {
		C.QComboBox_SetSizeAdjustPolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QComboBox) SetValidator(validator gui.QValidator_ITF) {
	defer qt.Recovering("QComboBox::setValidator")

	if ptr.Pointer() != nil {
		C.QComboBox_SetValidator(ptr.Pointer(), gui.PointerFromQValidator(validator))
	}
}

func (ptr *QComboBox) SizeAdjustPolicy() QComboBox__SizeAdjustPolicy {
	defer qt.Recovering("QComboBox::sizeAdjustPolicy")

	if ptr.Pointer() != nil {
		return QComboBox__SizeAdjustPolicy(C.QComboBox_SizeAdjustPolicy(ptr.Pointer()))
	}
	return 0
}

func NewQComboBox(parent QWidget_ITF) *QComboBox {
	defer qt.Recovering("QComboBox::QComboBox")

	return NewQComboBoxFromPointer(C.QComboBox_NewQComboBox(PointerFromQWidget(parent)))
}

func (ptr *QComboBox) ConnectActivated2(f func(text string)) {
	defer qt.Recovering("connect QComboBox::activated")

	if ptr.Pointer() != nil {
		C.QComboBox_ConnectActivated2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activated2", f)
	}
}

func (ptr *QComboBox) DisconnectActivated2() {
	defer qt.Recovering("disconnect QComboBox::activated")

	if ptr.Pointer() != nil {
		C.QComboBox_DisconnectActivated2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activated2")
	}
}

//export callbackQComboBoxActivated2
func callbackQComboBoxActivated2(ptrName *C.char, text *C.char) {
	defer qt.Recovering("callback QComboBox::activated")

	var signal = qt.GetSignal(C.GoString(ptrName), "activated2")
	if signal != nil {
		signal.(func(string))(C.GoString(text))
	}

}

func (ptr *QComboBox) ConnectActivated(f func(index int)) {
	defer qt.Recovering("connect QComboBox::activated")

	if ptr.Pointer() != nil {
		C.QComboBox_ConnectActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activated", f)
	}
}

func (ptr *QComboBox) DisconnectActivated() {
	defer qt.Recovering("disconnect QComboBox::activated")

	if ptr.Pointer() != nil {
		C.QComboBox_DisconnectActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activated")
	}
}

//export callbackQComboBoxActivated
func callbackQComboBoxActivated(ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QComboBox::activated")

	var signal = qt.GetSignal(C.GoString(ptrName), "activated")
	if signal != nil {
		signal.(func(int))(int(index))
	}

}

func (ptr *QComboBox) AddItem2(icon gui.QIcon_ITF, text string, userData core.QVariant_ITF) {
	defer qt.Recovering("QComboBox::addItem")

	if ptr.Pointer() != nil {
		C.QComboBox_AddItem2(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(text), core.PointerFromQVariant(userData))
	}
}

func (ptr *QComboBox) AddItem(text string, userData core.QVariant_ITF) {
	defer qt.Recovering("QComboBox::addItem")

	if ptr.Pointer() != nil {
		C.QComboBox_AddItem(ptr.Pointer(), C.CString(text), core.PointerFromQVariant(userData))
	}
}

func (ptr *QComboBox) AddItems(texts []string) {
	defer qt.Recovering("QComboBox::addItems")

	if ptr.Pointer() != nil {
		C.QComboBox_AddItems(ptr.Pointer(), C.CString(strings.Join(texts, ",,,")))
	}
}

func (ptr *QComboBox) ConnectChangeEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QComboBox::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QComboBox) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QComboBox::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQComboBoxChangeEvent
func callbackQComboBoxChangeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QComboBox::changeEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "changeEvent")
	if signal != nil {
		defer signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QComboBox) Clear() {
	defer qt.Recovering("QComboBox::clear")

	if ptr.Pointer() != nil {
		C.QComboBox_Clear(ptr.Pointer())
	}
}

func (ptr *QComboBox) ClearEditText() {
	defer qt.Recovering("QComboBox::clearEditText")

	if ptr.Pointer() != nil {
		C.QComboBox_ClearEditText(ptr.Pointer())
	}
}

func (ptr *QComboBox) Completer() *QCompleter {
	defer qt.Recovering("QComboBox::completer")

	if ptr.Pointer() != nil {
		return NewQCompleterFromPointer(C.QComboBox_Completer(ptr.Pointer()))
	}
	return nil
}

func (ptr *QComboBox) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QComboBox::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QComboBox) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QComboBox::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQComboBoxContextMenuEvent
func callbackQComboBoxContextMenuEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QComboBox::contextMenuEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "contextMenuEvent")
	if signal != nil {
		defer signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QComboBox) ConnectCurrentIndexChanged2(f func(text string)) {
	defer qt.Recovering("connect QComboBox::currentIndexChanged")

	if ptr.Pointer() != nil {
		C.QComboBox_ConnectCurrentIndexChanged2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentIndexChanged2", f)
	}
}

func (ptr *QComboBox) DisconnectCurrentIndexChanged2() {
	defer qt.Recovering("disconnect QComboBox::currentIndexChanged")

	if ptr.Pointer() != nil {
		C.QComboBox_DisconnectCurrentIndexChanged2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentIndexChanged2")
	}
}

//export callbackQComboBoxCurrentIndexChanged2
func callbackQComboBoxCurrentIndexChanged2(ptrName *C.char, text *C.char) {
	defer qt.Recovering("callback QComboBox::currentIndexChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "currentIndexChanged2")
	if signal != nil {
		signal.(func(string))(C.GoString(text))
	}

}

func (ptr *QComboBox) ConnectCurrentIndexChanged(f func(index int)) {
	defer qt.Recovering("connect QComboBox::currentIndexChanged")

	if ptr.Pointer() != nil {
		C.QComboBox_ConnectCurrentIndexChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentIndexChanged", f)
	}
}

func (ptr *QComboBox) DisconnectCurrentIndexChanged() {
	defer qt.Recovering("disconnect QComboBox::currentIndexChanged")

	if ptr.Pointer() != nil {
		C.QComboBox_DisconnectCurrentIndexChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentIndexChanged")
	}
}

//export callbackQComboBoxCurrentIndexChanged
func callbackQComboBoxCurrentIndexChanged(ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QComboBox::currentIndexChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "currentIndexChanged")
	if signal != nil {
		signal.(func(int))(int(index))
	}

}

func (ptr *QComboBox) ConnectCurrentTextChanged(f func(text string)) {
	defer qt.Recovering("connect QComboBox::currentTextChanged")

	if ptr.Pointer() != nil {
		C.QComboBox_ConnectCurrentTextChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentTextChanged", f)
	}
}

func (ptr *QComboBox) DisconnectCurrentTextChanged() {
	defer qt.Recovering("disconnect QComboBox::currentTextChanged")

	if ptr.Pointer() != nil {
		C.QComboBox_DisconnectCurrentTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentTextChanged")
	}
}

//export callbackQComboBoxCurrentTextChanged
func callbackQComboBoxCurrentTextChanged(ptrName *C.char, text *C.char) {
	defer qt.Recovering("callback QComboBox::currentTextChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "currentTextChanged")
	if signal != nil {
		signal.(func(string))(C.GoString(text))
	}

}

func (ptr *QComboBox) ConnectEditTextChanged(f func(text string)) {
	defer qt.Recovering("connect QComboBox::editTextChanged")

	if ptr.Pointer() != nil {
		C.QComboBox_ConnectEditTextChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "editTextChanged", f)
	}
}

func (ptr *QComboBox) DisconnectEditTextChanged() {
	defer qt.Recovering("disconnect QComboBox::editTextChanged")

	if ptr.Pointer() != nil {
		C.QComboBox_DisconnectEditTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "editTextChanged")
	}
}

//export callbackQComboBoxEditTextChanged
func callbackQComboBoxEditTextChanged(ptrName *C.char, text *C.char) {
	defer qt.Recovering("callback QComboBox::editTextChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "editTextChanged")
	if signal != nil {
		signal.(func(string))(C.GoString(text))
	}

}

func (ptr *QComboBox) Event(event core.QEvent_ITF) bool {
	defer qt.Recovering("QComboBox::event")

	if ptr.Pointer() != nil {
		return C.QComboBox_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QComboBox) FindData(data core.QVariant_ITF, role int, flags core.Qt__MatchFlag) int {
	defer qt.Recovering("QComboBox::findData")

	if ptr.Pointer() != nil {
		return int(C.QComboBox_FindData(ptr.Pointer(), core.PointerFromQVariant(data), C.int(role), C.int(flags)))
	}
	return 0
}

func (ptr *QComboBox) FindText(text string, flags core.Qt__MatchFlag) int {
	defer qt.Recovering("QComboBox::findText")

	if ptr.Pointer() != nil {
		return int(C.QComboBox_FindText(ptr.Pointer(), C.CString(text), C.int(flags)))
	}
	return 0
}

func (ptr *QComboBox) ConnectFocusInEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QComboBox::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QComboBox) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QComboBox::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQComboBoxFocusInEvent
func callbackQComboBoxFocusInEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QComboBox::focusInEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "focusInEvent")
	if signal != nil {
		defer signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QComboBox) ConnectFocusOutEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QComboBox::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QComboBox) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QComboBox::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQComboBoxFocusOutEvent
func callbackQComboBoxFocusOutEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QComboBox::focusOutEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "focusOutEvent")
	if signal != nil {
		defer signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QComboBox) ConnectHideEvent(f func(e *gui.QHideEvent)) {
	defer qt.Recovering("connect QComboBox::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QComboBox) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QComboBox::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQComboBoxHideEvent
func callbackQComboBoxHideEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QComboBox::hideEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "hideEvent")
	if signal != nil {
		defer signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QComboBox) ConnectHidePopup(f func()) {
	defer qt.Recovering("connect QComboBox::hidePopup")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hidePopup", f)
	}
}

func (ptr *QComboBox) DisconnectHidePopup() {
	defer qt.Recovering("disconnect QComboBox::hidePopup")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hidePopup")
	}
}

//export callbackQComboBoxHidePopup
func callbackQComboBoxHidePopup(ptrName *C.char) bool {
	defer qt.Recovering("callback QComboBox::hidePopup")

	var signal = qt.GetSignal(C.GoString(ptrName), "hidePopup")
	if signal != nil {
		defer signal.(func())()
		return true
	}
	return false

}

func (ptr *QComboBox) ConnectHighlighted2(f func(text string)) {
	defer qt.Recovering("connect QComboBox::highlighted")

	if ptr.Pointer() != nil {
		C.QComboBox_ConnectHighlighted2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "highlighted2", f)
	}
}

func (ptr *QComboBox) DisconnectHighlighted2() {
	defer qt.Recovering("disconnect QComboBox::highlighted")

	if ptr.Pointer() != nil {
		C.QComboBox_DisconnectHighlighted2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "highlighted2")
	}
}

//export callbackQComboBoxHighlighted2
func callbackQComboBoxHighlighted2(ptrName *C.char, text *C.char) {
	defer qt.Recovering("callback QComboBox::highlighted")

	var signal = qt.GetSignal(C.GoString(ptrName), "highlighted2")
	if signal != nil {
		signal.(func(string))(C.GoString(text))
	}

}

func (ptr *QComboBox) ConnectHighlighted(f func(index int)) {
	defer qt.Recovering("connect QComboBox::highlighted")

	if ptr.Pointer() != nil {
		C.QComboBox_ConnectHighlighted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "highlighted", f)
	}
}

func (ptr *QComboBox) DisconnectHighlighted() {
	defer qt.Recovering("disconnect QComboBox::highlighted")

	if ptr.Pointer() != nil {
		C.QComboBox_DisconnectHighlighted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "highlighted")
	}
}

//export callbackQComboBoxHighlighted
func callbackQComboBoxHighlighted(ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QComboBox::highlighted")

	var signal = qt.GetSignal(C.GoString(ptrName), "highlighted")
	if signal != nil {
		signal.(func(int))(int(index))
	}

}

func (ptr *QComboBox) ConnectInputMethodEvent(f func(e *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QComboBox::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QComboBox) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QComboBox::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQComboBoxInputMethodEvent
func callbackQComboBoxInputMethodEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QComboBox::inputMethodEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "inputMethodEvent")
	if signal != nil {
		defer signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QComboBox) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QComboBox::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QComboBox_InputMethodQuery(ptr.Pointer(), C.int(query)))
	}
	return nil
}

func (ptr *QComboBox) InsertItem2(index int, icon gui.QIcon_ITF, text string, userData core.QVariant_ITF) {
	defer qt.Recovering("QComboBox::insertItem")

	if ptr.Pointer() != nil {
		C.QComboBox_InsertItem2(ptr.Pointer(), C.int(index), gui.PointerFromQIcon(icon), C.CString(text), core.PointerFromQVariant(userData))
	}
}

func (ptr *QComboBox) InsertItem(index int, text string, userData core.QVariant_ITF) {
	defer qt.Recovering("QComboBox::insertItem")

	if ptr.Pointer() != nil {
		C.QComboBox_InsertItem(ptr.Pointer(), C.int(index), C.CString(text), core.PointerFromQVariant(userData))
	}
}

func (ptr *QComboBox) InsertItems(index int, list []string) {
	defer qt.Recovering("QComboBox::insertItems")

	if ptr.Pointer() != nil {
		C.QComboBox_InsertItems(ptr.Pointer(), C.int(index), C.CString(strings.Join(list, ",,,")))
	}
}

func (ptr *QComboBox) InsertSeparator(index int) {
	defer qt.Recovering("QComboBox::insertSeparator")

	if ptr.Pointer() != nil {
		C.QComboBox_InsertSeparator(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QComboBox) ItemData(index int, role int) *core.QVariant {
	defer qt.Recovering("QComboBox::itemData")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QComboBox_ItemData(ptr.Pointer(), C.int(index), C.int(role)))
	}
	return nil
}

func (ptr *QComboBox) ItemDelegate() *QAbstractItemDelegate {
	defer qt.Recovering("QComboBox::itemDelegate")

	if ptr.Pointer() != nil {
		return NewQAbstractItemDelegateFromPointer(C.QComboBox_ItemDelegate(ptr.Pointer()))
	}
	return nil
}

func (ptr *QComboBox) ItemIcon(index int) *gui.QIcon {
	defer qt.Recovering("QComboBox::itemIcon")

	if ptr.Pointer() != nil {
		return gui.NewQIconFromPointer(C.QComboBox_ItemIcon(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QComboBox) ItemText(index int) string {
	defer qt.Recovering("QComboBox::itemText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QComboBox_ItemText(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QComboBox) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QComboBox::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QComboBox) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QComboBox::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQComboBoxKeyPressEvent
func callbackQComboBoxKeyPressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QComboBox::keyPressEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "keyPressEvent")
	if signal != nil {
		defer signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QComboBox) ConnectKeyReleaseEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QComboBox::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QComboBox) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QComboBox::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQComboBoxKeyReleaseEvent
func callbackQComboBoxKeyReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QComboBox::keyReleaseEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent")
	if signal != nil {
		defer signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QComboBox) LineEdit() *QLineEdit {
	defer qt.Recovering("QComboBox::lineEdit")

	if ptr.Pointer() != nil {
		return NewQLineEditFromPointer(C.QComboBox_LineEdit(ptr.Pointer()))
	}
	return nil
}

func (ptr *QComboBox) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QComboBox::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QComboBox_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QComboBox) Model() *core.QAbstractItemModel {
	defer qt.Recovering("QComboBox::model")

	if ptr.Pointer() != nil {
		return core.NewQAbstractItemModelFromPointer(C.QComboBox_Model(ptr.Pointer()))
	}
	return nil
}

func (ptr *QComboBox) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QComboBox::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QComboBox) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QComboBox::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQComboBoxMousePressEvent
func callbackQComboBoxMousePressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QComboBox::mousePressEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mousePressEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QComboBox) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QComboBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QComboBox) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QComboBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQComboBoxMouseReleaseEvent
func callbackQComboBoxMouseReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QComboBox::mouseReleaseEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QComboBox) ConnectPaintEvent(f func(e *gui.QPaintEvent)) {
	defer qt.Recovering("connect QComboBox::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QComboBox) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QComboBox::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQComboBoxPaintEvent
func callbackQComboBoxPaintEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QComboBox::paintEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "paintEvent")
	if signal != nil {
		defer signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QComboBox) RemoveItem(index int) {
	defer qt.Recovering("QComboBox::removeItem")

	if ptr.Pointer() != nil {
		C.QComboBox_RemoveItem(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QComboBox) ConnectResizeEvent(f func(e *gui.QResizeEvent)) {
	defer qt.Recovering("connect QComboBox::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QComboBox) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QComboBox::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQComboBoxResizeEvent
func callbackQComboBoxResizeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QComboBox::resizeEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "resizeEvent")
	if signal != nil {
		defer signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QComboBox) RootModelIndex() *core.QModelIndex {
	defer qt.Recovering("QComboBox::rootModelIndex")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QComboBox_RootModelIndex(ptr.Pointer()))
	}
	return nil
}

func (ptr *QComboBox) SetEditText(text string) {
	defer qt.Recovering("QComboBox::setEditText")

	if ptr.Pointer() != nil {
		C.QComboBox_SetEditText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QComboBox) SetItemData(index int, value core.QVariant_ITF, role int) {
	defer qt.Recovering("QComboBox::setItemData")

	if ptr.Pointer() != nil {
		C.QComboBox_SetItemData(ptr.Pointer(), C.int(index), core.PointerFromQVariant(value), C.int(role))
	}
}

func (ptr *QComboBox) SetItemDelegate(delegate QAbstractItemDelegate_ITF) {
	defer qt.Recovering("QComboBox::setItemDelegate")

	if ptr.Pointer() != nil {
		C.QComboBox_SetItemDelegate(ptr.Pointer(), PointerFromQAbstractItemDelegate(delegate))
	}
}

func (ptr *QComboBox) SetItemIcon(index int, icon gui.QIcon_ITF) {
	defer qt.Recovering("QComboBox::setItemIcon")

	if ptr.Pointer() != nil {
		C.QComboBox_SetItemIcon(ptr.Pointer(), C.int(index), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QComboBox) SetItemText(index int, text string) {
	defer qt.Recovering("QComboBox::setItemText")

	if ptr.Pointer() != nil {
		C.QComboBox_SetItemText(ptr.Pointer(), C.int(index), C.CString(text))
	}
}

func (ptr *QComboBox) SetLineEdit(edit QLineEdit_ITF) {
	defer qt.Recovering("QComboBox::setLineEdit")

	if ptr.Pointer() != nil {
		C.QComboBox_SetLineEdit(ptr.Pointer(), PointerFromQLineEdit(edit))
	}
}

func (ptr *QComboBox) SetModel(model core.QAbstractItemModel_ITF) {
	defer qt.Recovering("QComboBox::setModel")

	if ptr.Pointer() != nil {
		C.QComboBox_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QComboBox) SetRootModelIndex(index core.QModelIndex_ITF) {
	defer qt.Recovering("QComboBox::setRootModelIndex")

	if ptr.Pointer() != nil {
		C.QComboBox_SetRootModelIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QComboBox) SetView(itemView QAbstractItemView_ITF) {
	defer qt.Recovering("QComboBox::setView")

	if ptr.Pointer() != nil {
		C.QComboBox_SetView(ptr.Pointer(), PointerFromQAbstractItemView(itemView))
	}
}

func (ptr *QComboBox) ConnectShowEvent(f func(e *gui.QShowEvent)) {
	defer qt.Recovering("connect QComboBox::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QComboBox) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QComboBox::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQComboBoxShowEvent
func callbackQComboBoxShowEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QComboBox::showEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "showEvent")
	if signal != nil {
		defer signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QComboBox) ConnectShowPopup(f func()) {
	defer qt.Recovering("connect QComboBox::showPopup")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showPopup", f)
	}
}

func (ptr *QComboBox) DisconnectShowPopup() {
	defer qt.Recovering("disconnect QComboBox::showPopup")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showPopup")
	}
}

//export callbackQComboBoxShowPopup
func callbackQComboBoxShowPopup(ptrName *C.char) bool {
	defer qt.Recovering("callback QComboBox::showPopup")

	var signal = qt.GetSignal(C.GoString(ptrName), "showPopup")
	if signal != nil {
		defer signal.(func())()
		return true
	}
	return false

}

func (ptr *QComboBox) SizeHint() *core.QSize {
	defer qt.Recovering("QComboBox::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QComboBox_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QComboBox) Validator() *gui.QValidator {
	defer qt.Recovering("QComboBox::validator")

	if ptr.Pointer() != nil {
		return gui.NewQValidatorFromPointer(C.QComboBox_Validator(ptr.Pointer()))
	}
	return nil
}

func (ptr *QComboBox) View() *QAbstractItemView {
	defer qt.Recovering("QComboBox::view")

	if ptr.Pointer() != nil {
		return NewQAbstractItemViewFromPointer(C.QComboBox_View(ptr.Pointer()))
	}
	return nil
}

func (ptr *QComboBox) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QComboBox::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QComboBox) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QComboBox::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQComboBoxWheelEvent
func callbackQComboBoxWheelEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QComboBox::wheelEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "wheelEvent")
	if signal != nil {
		defer signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QComboBox) DestroyQComboBox() {
	defer qt.Recovering("QComboBox::~QComboBox")

	if ptr.Pointer() != nil {
		C.QComboBox_DestroyQComboBox(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
