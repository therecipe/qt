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
func callbackQComboBoxActivated2(ptr unsafe.Pointer, ptrName *C.char, text *C.char) {
	defer qt.Recovering("callback QComboBox::activated")

	if signal := qt.GetSignal(C.GoString(ptrName), "activated2"); signal != nil {
		signal.(func(string))(C.GoString(text))
	}

}

func (ptr *QComboBox) Activated2(text string) {
	defer qt.Recovering("QComboBox::activated")

	if ptr.Pointer() != nil {
		C.QComboBox_Activated2(ptr.Pointer(), C.CString(text))
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
func callbackQComboBoxActivated(ptr unsafe.Pointer, ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QComboBox::activated")

	if signal := qt.GetSignal(C.GoString(ptrName), "activated"); signal != nil {
		signal.(func(int))(int(index))
	}

}

func (ptr *QComboBox) Activated(index int) {
	defer qt.Recovering("QComboBox::activated")

	if ptr.Pointer() != nil {
		C.QComboBox_Activated(ptr.Pointer(), C.int(index))
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
func callbackQComboBoxChangeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QComboBox::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
	} else {
		NewQComboBoxFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(e))
	}
}

func (ptr *QComboBox) ChangeEvent(e core.QEvent_ITF) {
	defer qt.Recovering("QComboBox::changeEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

func (ptr *QComboBox) ChangeEventDefault(e core.QEvent_ITF) {
	defer qt.Recovering("QComboBox::changeEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(e))
	}
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
func callbackQComboBoxContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QComboBox::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQComboBoxFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QComboBox) ContextMenuEvent(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QComboBox::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QComboBox) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QComboBox::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
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
func callbackQComboBoxCurrentIndexChanged2(ptr unsafe.Pointer, ptrName *C.char, text *C.char) {
	defer qt.Recovering("callback QComboBox::currentIndexChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentIndexChanged2"); signal != nil {
		signal.(func(string))(C.GoString(text))
	}

}

func (ptr *QComboBox) CurrentIndexChanged2(text string) {
	defer qt.Recovering("QComboBox::currentIndexChanged")

	if ptr.Pointer() != nil {
		C.QComboBox_CurrentIndexChanged2(ptr.Pointer(), C.CString(text))
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
func callbackQComboBoxCurrentIndexChanged(ptr unsafe.Pointer, ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QComboBox::currentIndexChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentIndexChanged"); signal != nil {
		signal.(func(int))(int(index))
	}

}

func (ptr *QComboBox) CurrentIndexChanged(index int) {
	defer qt.Recovering("QComboBox::currentIndexChanged")

	if ptr.Pointer() != nil {
		C.QComboBox_CurrentIndexChanged(ptr.Pointer(), C.int(index))
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
func callbackQComboBoxCurrentTextChanged(ptr unsafe.Pointer, ptrName *C.char, text *C.char) {
	defer qt.Recovering("callback QComboBox::currentTextChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentTextChanged"); signal != nil {
		signal.(func(string))(C.GoString(text))
	}

}

func (ptr *QComboBox) CurrentTextChanged(text string) {
	defer qt.Recovering("QComboBox::currentTextChanged")

	if ptr.Pointer() != nil {
		C.QComboBox_CurrentTextChanged(ptr.Pointer(), C.CString(text))
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
func callbackQComboBoxEditTextChanged(ptr unsafe.Pointer, ptrName *C.char, text *C.char) {
	defer qt.Recovering("callback QComboBox::editTextChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "editTextChanged"); signal != nil {
		signal.(func(string))(C.GoString(text))
	}

}

func (ptr *QComboBox) EditTextChanged(text string) {
	defer qt.Recovering("QComboBox::editTextChanged")

	if ptr.Pointer() != nil {
		C.QComboBox_EditTextChanged(ptr.Pointer(), C.CString(text))
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
func callbackQComboBoxFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QComboBox::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
	} else {
		NewQComboBoxFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(e))
	}
}

func (ptr *QComboBox) FocusInEvent(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QComboBox::focusInEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QComboBox) FocusInEventDefault(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QComboBox::focusInEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
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
func callbackQComboBoxFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QComboBox::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
	} else {
		NewQComboBoxFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(e))
	}
}

func (ptr *QComboBox) FocusOutEvent(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QComboBox::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QComboBox) FocusOutEventDefault(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QComboBox::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
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
func callbackQComboBoxHideEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QComboBox::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(e))
	} else {
		NewQComboBoxFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(e))
	}
}

func (ptr *QComboBox) HideEvent(e gui.QHideEvent_ITF) {
	defer qt.Recovering("QComboBox::hideEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(e))
	}
}

func (ptr *QComboBox) HideEventDefault(e gui.QHideEvent_ITF) {
	defer qt.Recovering("QComboBox::hideEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(e))
	}
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
func callbackQComboBoxHidePopup(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QComboBox::hidePopup")

	if signal := qt.GetSignal(C.GoString(ptrName), "hidePopup"); signal != nil {
		signal.(func())()
	} else {
		NewQComboBoxFromPointer(ptr).HidePopupDefault()
	}
}

func (ptr *QComboBox) HidePopup() {
	defer qt.Recovering("QComboBox::hidePopup")

	if ptr.Pointer() != nil {
		C.QComboBox_HidePopup(ptr.Pointer())
	}
}

func (ptr *QComboBox) HidePopupDefault() {
	defer qt.Recovering("QComboBox::hidePopup")

	if ptr.Pointer() != nil {
		C.QComboBox_HidePopupDefault(ptr.Pointer())
	}
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
func callbackQComboBoxHighlighted2(ptr unsafe.Pointer, ptrName *C.char, text *C.char) {
	defer qt.Recovering("callback QComboBox::highlighted")

	if signal := qt.GetSignal(C.GoString(ptrName), "highlighted2"); signal != nil {
		signal.(func(string))(C.GoString(text))
	}

}

func (ptr *QComboBox) Highlighted2(text string) {
	defer qt.Recovering("QComboBox::highlighted")

	if ptr.Pointer() != nil {
		C.QComboBox_Highlighted2(ptr.Pointer(), C.CString(text))
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
func callbackQComboBoxHighlighted(ptr unsafe.Pointer, ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QComboBox::highlighted")

	if signal := qt.GetSignal(C.GoString(ptrName), "highlighted"); signal != nil {
		signal.(func(int))(int(index))
	}

}

func (ptr *QComboBox) Highlighted(index int) {
	defer qt.Recovering("QComboBox::highlighted")

	if ptr.Pointer() != nil {
		C.QComboBox_Highlighted(ptr.Pointer(), C.int(index))
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
func callbackQComboBoxInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QComboBox::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(e))
	} else {
		NewQComboBoxFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(e))
	}
}

func (ptr *QComboBox) InputMethodEvent(e gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QComboBox::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(e))
	}
}

func (ptr *QComboBox) InputMethodEventDefault(e gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QComboBox::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(e))
	}
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
func callbackQComboBoxKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QComboBox::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQComboBoxFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QComboBox) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QComboBox::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QComboBox) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QComboBox::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
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
func callbackQComboBoxKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QComboBox::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQComboBoxFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QComboBox) KeyReleaseEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QComboBox::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QComboBox) KeyReleaseEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QComboBox::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
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
func callbackQComboBoxMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QComboBox::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQComboBoxFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QComboBox) MousePressEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QComboBox::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QComboBox) MousePressEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QComboBox::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
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
func callbackQComboBoxMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QComboBox::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQComboBoxFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QComboBox) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QComboBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QComboBox) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QComboBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
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
func callbackQComboBoxPaintEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QComboBox::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(e))
	} else {
		NewQComboBoxFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(e))
	}
}

func (ptr *QComboBox) PaintEvent(e gui.QPaintEvent_ITF) {
	defer qt.Recovering("QComboBox::paintEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

func (ptr *QComboBox) PaintEventDefault(e gui.QPaintEvent_ITF) {
	defer qt.Recovering("QComboBox::paintEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
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
func callbackQComboBoxResizeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QComboBox::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
	} else {
		NewQComboBoxFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(e))
	}
}

func (ptr *QComboBox) ResizeEvent(e gui.QResizeEvent_ITF) {
	defer qt.Recovering("QComboBox::resizeEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
}

func (ptr *QComboBox) ResizeEventDefault(e gui.QResizeEvent_ITF) {
	defer qt.Recovering("QComboBox::resizeEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
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
func callbackQComboBoxShowEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QComboBox::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(e))
	} else {
		NewQComboBoxFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(e))
	}
}

func (ptr *QComboBox) ShowEvent(e gui.QShowEvent_ITF) {
	defer qt.Recovering("QComboBox::showEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(e))
	}
}

func (ptr *QComboBox) ShowEventDefault(e gui.QShowEvent_ITF) {
	defer qt.Recovering("QComboBox::showEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(e))
	}
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
func callbackQComboBoxShowPopup(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QComboBox::showPopup")

	if signal := qt.GetSignal(C.GoString(ptrName), "showPopup"); signal != nil {
		signal.(func())()
	} else {
		NewQComboBoxFromPointer(ptr).ShowPopupDefault()
	}
}

func (ptr *QComboBox) ShowPopup() {
	defer qt.Recovering("QComboBox::showPopup")

	if ptr.Pointer() != nil {
		C.QComboBox_ShowPopup(ptr.Pointer())
	}
}

func (ptr *QComboBox) ShowPopupDefault() {
	defer qt.Recovering("QComboBox::showPopup")

	if ptr.Pointer() != nil {
		C.QComboBox_ShowPopupDefault(ptr.Pointer())
	}
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
func callbackQComboBoxWheelEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QComboBox::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQComboBoxFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QComboBox) WheelEvent(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QComboBox::wheelEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QComboBox) WheelEventDefault(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QComboBox::wheelEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QComboBox) DestroyQComboBox() {
	defer qt.Recovering("QComboBox::~QComboBox")

	if ptr.Pointer() != nil {
		C.QComboBox_DestroyQComboBox(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QComboBox) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QComboBox::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QComboBox) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QComboBox::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQComboBoxActionEvent
func callbackQComboBoxActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QComboBox::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQComboBoxFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QComboBox) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QComboBox::actionEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QComboBox) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QComboBox::actionEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QComboBox) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QComboBox::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QComboBox) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QComboBox::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQComboBoxDragEnterEvent
func callbackQComboBoxDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QComboBox::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQComboBoxFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QComboBox) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QComboBox::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QComboBox) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QComboBox::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QComboBox) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QComboBox::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QComboBox) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QComboBox::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQComboBoxDragLeaveEvent
func callbackQComboBoxDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QComboBox::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQComboBoxFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QComboBox) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QComboBox::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QComboBox) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QComboBox::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QComboBox) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QComboBox::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QComboBox) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QComboBox::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQComboBoxDragMoveEvent
func callbackQComboBoxDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QComboBox::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQComboBoxFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QComboBox) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QComboBox::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QComboBox) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QComboBox::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QComboBox) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QComboBox::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QComboBox) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QComboBox::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQComboBoxDropEvent
func callbackQComboBoxDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QComboBox::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQComboBoxFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QComboBox) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QComboBox::dropEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QComboBox) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QComboBox::dropEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QComboBox) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QComboBox::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QComboBox) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QComboBox::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQComboBoxEnterEvent
func callbackQComboBoxEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QComboBox::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQComboBoxFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QComboBox) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QComboBox::enterEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QComboBox) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QComboBox::enterEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QComboBox) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QComboBox::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QComboBox) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QComboBox::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQComboBoxLeaveEvent
func callbackQComboBoxLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QComboBox::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQComboBoxFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QComboBox) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QComboBox::leaveEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QComboBox) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QComboBox::leaveEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QComboBox) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QComboBox::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QComboBox) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QComboBox::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQComboBoxMoveEvent
func callbackQComboBoxMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QComboBox::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQComboBoxFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QComboBox) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QComboBox::moveEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QComboBox) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QComboBox::moveEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QComboBox) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QComboBox::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QComboBox) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QComboBox::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQComboBoxSetVisible
func callbackQComboBoxSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QComboBox::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QComboBox) SetVisible(visible bool) {
	defer qt.Recovering("QComboBox::setVisible")

	if ptr.Pointer() != nil {
		C.QComboBox_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QComboBox) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QComboBox::setVisible")

	if ptr.Pointer() != nil {
		C.QComboBox_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QComboBox) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QComboBox::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QComboBox) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QComboBox::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQComboBoxCloseEvent
func callbackQComboBoxCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QComboBox::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQComboBoxFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QComboBox) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QComboBox::closeEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QComboBox) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QComboBox::closeEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QComboBox) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QComboBox::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QComboBox) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QComboBox::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQComboBoxInitPainter
func callbackQComboBoxInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QComboBox::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQComboBoxFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QComboBox) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QComboBox::initPainter")

	if ptr.Pointer() != nil {
		C.QComboBox_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QComboBox) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QComboBox::initPainter")

	if ptr.Pointer() != nil {
		C.QComboBox_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QComboBox) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QComboBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QComboBox) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QComboBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQComboBoxMouseDoubleClickEvent
func callbackQComboBoxMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QComboBox::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQComboBoxFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QComboBox) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QComboBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QComboBox) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QComboBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QComboBox) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QComboBox::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QComboBox) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QComboBox::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQComboBoxMouseMoveEvent
func callbackQComboBoxMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QComboBox::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQComboBoxFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QComboBox) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QComboBox::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QComboBox) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QComboBox::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QComboBox) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QComboBox::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QComboBox) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QComboBox::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQComboBoxTabletEvent
func callbackQComboBoxTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QComboBox::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQComboBoxFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QComboBox) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QComboBox::tabletEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QComboBox) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QComboBox::tabletEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QComboBox) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QComboBox::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QComboBox) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QComboBox::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQComboBoxTimerEvent
func callbackQComboBoxTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QComboBox::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQComboBoxFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QComboBox) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QComboBox::timerEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QComboBox) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QComboBox::timerEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QComboBox) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QComboBox::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QComboBox) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QComboBox::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQComboBoxChildEvent
func callbackQComboBoxChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QComboBox::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQComboBoxFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QComboBox) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QComboBox::childEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QComboBox) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QComboBox::childEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QComboBox) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QComboBox::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QComboBox) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QComboBox::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQComboBoxCustomEvent
func callbackQComboBoxCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QComboBox::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQComboBoxFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QComboBox) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QComboBox::customEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QComboBox) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QComboBox::customEvent")

	if ptr.Pointer() != nil {
		C.QComboBox_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
