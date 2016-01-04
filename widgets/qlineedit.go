package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QLineEdit struct {
	QWidget
}

type QLineEdit_ITF interface {
	QWidget_ITF
	QLineEdit_PTR() *QLineEdit
}

func PointerFromQLineEdit(ptr QLineEdit_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLineEdit_PTR().Pointer()
	}
	return nil
}

func NewQLineEditFromPointer(ptr unsafe.Pointer) *QLineEdit {
	var n = new(QLineEdit)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QLineEdit_") {
		n.SetObjectName("QLineEdit_" + qt.Identifier())
	}
	return n
}

func (ptr *QLineEdit) QLineEdit_PTR() *QLineEdit {
	return ptr
}

//QLineEdit::ActionPosition
type QLineEdit__ActionPosition int64

const (
	QLineEdit__LeadingPosition  = QLineEdit__ActionPosition(0)
	QLineEdit__TrailingPosition = QLineEdit__ActionPosition(1)
)

//QLineEdit::EchoMode
type QLineEdit__EchoMode int64

const (
	QLineEdit__Normal             = QLineEdit__EchoMode(0)
	QLineEdit__NoEcho             = QLineEdit__EchoMode(1)
	QLineEdit__Password           = QLineEdit__EchoMode(2)
	QLineEdit__PasswordEchoOnEdit = QLineEdit__EchoMode(3)
)

func (ptr *QLineEdit) Alignment() core.Qt__AlignmentFlag {
	defer qt.Recovering("QLineEdit::alignment")

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QLineEdit_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLineEdit) CursorMoveStyle() core.Qt__CursorMoveStyle {
	defer qt.Recovering("QLineEdit::cursorMoveStyle")

	if ptr.Pointer() != nil {
		return core.Qt__CursorMoveStyle(C.QLineEdit_CursorMoveStyle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLineEdit) CursorPosition() int {
	defer qt.Recovering("QLineEdit::cursorPosition")

	if ptr.Pointer() != nil {
		return int(C.QLineEdit_CursorPosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLineEdit) DisplayText() string {
	defer qt.Recovering("QLineEdit::displayText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLineEdit_DisplayText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLineEdit) DragEnabled() bool {
	defer qt.Recovering("QLineEdit::dragEnabled")

	if ptr.Pointer() != nil {
		return C.QLineEdit_DragEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLineEdit) EchoMode() QLineEdit__EchoMode {
	defer qt.Recovering("QLineEdit::echoMode")

	if ptr.Pointer() != nil {
		return QLineEdit__EchoMode(C.QLineEdit_EchoMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLineEdit) HasAcceptableInput() bool {
	defer qt.Recovering("QLineEdit::hasAcceptableInput")

	if ptr.Pointer() != nil {
		return C.QLineEdit_HasAcceptableInput(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLineEdit) HasFrame() bool {
	defer qt.Recovering("QLineEdit::hasFrame")

	if ptr.Pointer() != nil {
		return C.QLineEdit_HasFrame(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLineEdit) HasSelectedText() bool {
	defer qt.Recovering("QLineEdit::hasSelectedText")

	if ptr.Pointer() != nil {
		return C.QLineEdit_HasSelectedText(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLineEdit) InputMask() string {
	defer qt.Recovering("QLineEdit::inputMask")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLineEdit_InputMask(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLineEdit) IsClearButtonEnabled() bool {
	defer qt.Recovering("QLineEdit::isClearButtonEnabled")

	if ptr.Pointer() != nil {
		return C.QLineEdit_IsClearButtonEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLineEdit) IsModified() bool {
	defer qt.Recovering("QLineEdit::isModified")

	if ptr.Pointer() != nil {
		return C.QLineEdit_IsModified(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLineEdit) IsReadOnly() bool {
	defer qt.Recovering("QLineEdit::isReadOnly")

	if ptr.Pointer() != nil {
		return C.QLineEdit_IsReadOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLineEdit) IsRedoAvailable() bool {
	defer qt.Recovering("QLineEdit::isRedoAvailable")

	if ptr.Pointer() != nil {
		return C.QLineEdit_IsRedoAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLineEdit) IsUndoAvailable() bool {
	defer qt.Recovering("QLineEdit::isUndoAvailable")

	if ptr.Pointer() != nil {
		return C.QLineEdit_IsUndoAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLineEdit) MaxLength() int {
	defer qt.Recovering("QLineEdit::maxLength")

	if ptr.Pointer() != nil {
		return int(C.QLineEdit_MaxLength(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLineEdit) PlaceholderText() string {
	defer qt.Recovering("QLineEdit::placeholderText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLineEdit_PlaceholderText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLineEdit) SelectedText() string {
	defer qt.Recovering("QLineEdit::selectedText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLineEdit_SelectedText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLineEdit) SetAlignment(flag core.Qt__AlignmentFlag) {
	defer qt.Recovering("QLineEdit::setAlignment")

	if ptr.Pointer() != nil {
		C.QLineEdit_SetAlignment(ptr.Pointer(), C.int(flag))
	}
}

func (ptr *QLineEdit) SetClearButtonEnabled(enable bool) {
	defer qt.Recovering("QLineEdit::setClearButtonEnabled")

	if ptr.Pointer() != nil {
		C.QLineEdit_SetClearButtonEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QLineEdit) SetCursorMoveStyle(style core.Qt__CursorMoveStyle) {
	defer qt.Recovering("QLineEdit::setCursorMoveStyle")

	if ptr.Pointer() != nil {
		C.QLineEdit_SetCursorMoveStyle(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QLineEdit) SetCursorPosition(v int) {
	defer qt.Recovering("QLineEdit::setCursorPosition")

	if ptr.Pointer() != nil {
		C.QLineEdit_SetCursorPosition(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLineEdit) SetDragEnabled(b bool) {
	defer qt.Recovering("QLineEdit::setDragEnabled")

	if ptr.Pointer() != nil {
		C.QLineEdit_SetDragEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QLineEdit) SetEchoMode(v QLineEdit__EchoMode) {
	defer qt.Recovering("QLineEdit::setEchoMode")

	if ptr.Pointer() != nil {
		C.QLineEdit_SetEchoMode(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLineEdit) SetFrame(v bool) {
	defer qt.Recovering("QLineEdit::setFrame")

	if ptr.Pointer() != nil {
		C.QLineEdit_SetFrame(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QLineEdit) SetInputMask(inputMask string) {
	defer qt.Recovering("QLineEdit::setInputMask")

	if ptr.Pointer() != nil {
		C.QLineEdit_SetInputMask(ptr.Pointer(), C.CString(inputMask))
	}
}

func (ptr *QLineEdit) SetMaxLength(v int) {
	defer qt.Recovering("QLineEdit::setMaxLength")

	if ptr.Pointer() != nil {
		C.QLineEdit_SetMaxLength(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLineEdit) SetModified(v bool) {
	defer qt.Recovering("QLineEdit::setModified")

	if ptr.Pointer() != nil {
		C.QLineEdit_SetModified(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QLineEdit) SetPlaceholderText(v string) {
	defer qt.Recovering("QLineEdit::setPlaceholderText")

	if ptr.Pointer() != nil {
		C.QLineEdit_SetPlaceholderText(ptr.Pointer(), C.CString(v))
	}
}

func (ptr *QLineEdit) SetReadOnly(v bool) {
	defer qt.Recovering("QLineEdit::setReadOnly")

	if ptr.Pointer() != nil {
		C.QLineEdit_SetReadOnly(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QLineEdit) SetText(v string) {
	defer qt.Recovering("QLineEdit::setText")

	if ptr.Pointer() != nil {
		C.QLineEdit_SetText(ptr.Pointer(), C.CString(v))
	}
}

func (ptr *QLineEdit) Text() string {
	defer qt.Recovering("QLineEdit::text")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLineEdit_Text(ptr.Pointer()))
	}
	return ""
}

func NewQLineEdit(parent QWidget_ITF) *QLineEdit {
	defer qt.Recovering("QLineEdit::QLineEdit")

	return NewQLineEditFromPointer(C.QLineEdit_NewQLineEdit(PointerFromQWidget(parent)))
}

func NewQLineEdit2(contents string, parent QWidget_ITF) *QLineEdit {
	defer qt.Recovering("QLineEdit::QLineEdit")

	return NewQLineEditFromPointer(C.QLineEdit_NewQLineEdit2(C.CString(contents), PointerFromQWidget(parent)))
}

func (ptr *QLineEdit) AddAction2(icon gui.QIcon_ITF, position QLineEdit__ActionPosition) *QAction {
	defer qt.Recovering("QLineEdit::addAction")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QLineEdit_AddAction2(ptr.Pointer(), gui.PointerFromQIcon(icon), C.int(position)))
	}
	return nil
}

func (ptr *QLineEdit) AddAction(action QAction_ITF, position QLineEdit__ActionPosition) {
	defer qt.Recovering("QLineEdit::addAction")

	if ptr.Pointer() != nil {
		C.QLineEdit_AddAction(ptr.Pointer(), PointerFromQAction(action), C.int(position))
	}
}

func (ptr *QLineEdit) Backspace() {
	defer qt.Recovering("QLineEdit::backspace")

	if ptr.Pointer() != nil {
		C.QLineEdit_Backspace(ptr.Pointer())
	}
}

func (ptr *QLineEdit) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QLineEdit::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QLineEdit) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QLineEdit::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQLineEditChangeEvent
func callbackQLineEditChangeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QLineEdit::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQLineEditFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QLineEdit) ChangeEvent(ev core.QEvent_ITF) {
	defer qt.Recovering("QLineEdit::changeEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QLineEdit) ChangeEventDefault(ev core.QEvent_ITF) {
	defer qt.Recovering("QLineEdit::changeEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QLineEdit) Clear() {
	defer qt.Recovering("QLineEdit::clear")

	if ptr.Pointer() != nil {
		C.QLineEdit_Clear(ptr.Pointer())
	}
}

func (ptr *QLineEdit) Completer() *QCompleter {
	defer qt.Recovering("QLineEdit::completer")

	if ptr.Pointer() != nil {
		return NewQCompleterFromPointer(C.QLineEdit_Completer(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLineEdit) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QLineEdit::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QLineEdit) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QLineEdit::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQLineEditContextMenuEvent
func callbackQLineEditContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLineEdit::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQLineEditFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QLineEdit) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QLineEdit::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QLineEdit) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QLineEdit::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QLineEdit) Copy() {
	defer qt.Recovering("QLineEdit::copy")

	if ptr.Pointer() != nil {
		C.QLineEdit_Copy(ptr.Pointer())
	}
}

func (ptr *QLineEdit) CreateStandardContextMenu() *QMenu {
	defer qt.Recovering("QLineEdit::createStandardContextMenu")

	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QLineEdit_CreateStandardContextMenu(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLineEdit) CursorBackward(mark bool, steps int) {
	defer qt.Recovering("QLineEdit::cursorBackward")

	if ptr.Pointer() != nil {
		C.QLineEdit_CursorBackward(ptr.Pointer(), C.int(qt.GoBoolToInt(mark)), C.int(steps))
	}
}

func (ptr *QLineEdit) CursorForward(mark bool, steps int) {
	defer qt.Recovering("QLineEdit::cursorForward")

	if ptr.Pointer() != nil {
		C.QLineEdit_CursorForward(ptr.Pointer(), C.int(qt.GoBoolToInt(mark)), C.int(steps))
	}
}

func (ptr *QLineEdit) CursorPositionAt(pos core.QPoint_ITF) int {
	defer qt.Recovering("QLineEdit::cursorPositionAt")

	if ptr.Pointer() != nil {
		return int(C.QLineEdit_CursorPositionAt(ptr.Pointer(), core.PointerFromQPoint(pos)))
	}
	return 0
}

func (ptr *QLineEdit) ConnectCursorPositionChanged(f func(old int, n int)) {
	defer qt.Recovering("connect QLineEdit::cursorPositionChanged")

	if ptr.Pointer() != nil {
		C.QLineEdit_ConnectCursorPositionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cursorPositionChanged", f)
	}
}

func (ptr *QLineEdit) DisconnectCursorPositionChanged() {
	defer qt.Recovering("disconnect QLineEdit::cursorPositionChanged")

	if ptr.Pointer() != nil {
		C.QLineEdit_DisconnectCursorPositionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cursorPositionChanged")
	}
}

//export callbackQLineEditCursorPositionChanged
func callbackQLineEditCursorPositionChanged(ptr unsafe.Pointer, ptrName *C.char, old C.int, n C.int) {
	defer qt.Recovering("callback QLineEdit::cursorPositionChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "cursorPositionChanged"); signal != nil {
		signal.(func(int, int))(int(old), int(n))
	}

}

func (ptr *QLineEdit) CursorPositionChanged(old int, n int) {
	defer qt.Recovering("QLineEdit::cursorPositionChanged")

	if ptr.Pointer() != nil {
		C.QLineEdit_CursorPositionChanged(ptr.Pointer(), C.int(old), C.int(n))
	}
}

func (ptr *QLineEdit) CursorWordBackward(mark bool) {
	defer qt.Recovering("QLineEdit::cursorWordBackward")

	if ptr.Pointer() != nil {
		C.QLineEdit_CursorWordBackward(ptr.Pointer(), C.int(qt.GoBoolToInt(mark)))
	}
}

func (ptr *QLineEdit) CursorWordForward(mark bool) {
	defer qt.Recovering("QLineEdit::cursorWordForward")

	if ptr.Pointer() != nil {
		C.QLineEdit_CursorWordForward(ptr.Pointer(), C.int(qt.GoBoolToInt(mark)))
	}
}

func (ptr *QLineEdit) Cut() {
	defer qt.Recovering("QLineEdit::cut")

	if ptr.Pointer() != nil {
		C.QLineEdit_Cut(ptr.Pointer())
	}
}

func (ptr *QLineEdit) Del() {
	defer qt.Recovering("QLineEdit::del")

	if ptr.Pointer() != nil {
		C.QLineEdit_Del(ptr.Pointer())
	}
}

func (ptr *QLineEdit) Deselect() {
	defer qt.Recovering("QLineEdit::deselect")

	if ptr.Pointer() != nil {
		C.QLineEdit_Deselect(ptr.Pointer())
	}
}

func (ptr *QLineEdit) ConnectDragEnterEvent(f func(e *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QLineEdit::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QLineEdit) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QLineEdit::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQLineEditDragEnterEvent
func callbackQLineEditDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QLineEdit::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(e))
	} else {
		NewQLineEditFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(e))
	}
}

func (ptr *QLineEdit) DragEnterEvent(e gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QLineEdit::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(e))
	}
}

func (ptr *QLineEdit) DragEnterEventDefault(e gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QLineEdit::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(e))
	}
}

func (ptr *QLineEdit) ConnectDragLeaveEvent(f func(e *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QLineEdit::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QLineEdit) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QLineEdit::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQLineEditDragLeaveEvent
func callbackQLineEditDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QLineEdit::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(e))
	} else {
		NewQLineEditFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(e))
	}
}

func (ptr *QLineEdit) DragLeaveEvent(e gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QLineEdit::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

func (ptr *QLineEdit) DragLeaveEventDefault(e gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QLineEdit::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

func (ptr *QLineEdit) ConnectDragMoveEvent(f func(e *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QLineEdit::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QLineEdit) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QLineEdit::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQLineEditDragMoveEvent
func callbackQLineEditDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QLineEdit::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(e))
	} else {
		NewQLineEditFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(e))
	}
}

func (ptr *QLineEdit) DragMoveEvent(e gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QLineEdit::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

func (ptr *QLineEdit) DragMoveEventDefault(e gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QLineEdit::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

func (ptr *QLineEdit) ConnectDropEvent(f func(e *gui.QDropEvent)) {
	defer qt.Recovering("connect QLineEdit::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QLineEdit) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QLineEdit::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQLineEditDropEvent
func callbackQLineEditDropEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QLineEdit::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(e))
	} else {
		NewQLineEditFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(e))
	}
}

func (ptr *QLineEdit) DropEvent(e gui.QDropEvent_ITF) {
	defer qt.Recovering("QLineEdit::dropEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
}

func (ptr *QLineEdit) DropEventDefault(e gui.QDropEvent_ITF) {
	defer qt.Recovering("QLineEdit::dropEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
}

func (ptr *QLineEdit) ConnectEditingFinished(f func()) {
	defer qt.Recovering("connect QLineEdit::editingFinished")

	if ptr.Pointer() != nil {
		C.QLineEdit_ConnectEditingFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "editingFinished", f)
	}
}

func (ptr *QLineEdit) DisconnectEditingFinished() {
	defer qt.Recovering("disconnect QLineEdit::editingFinished")

	if ptr.Pointer() != nil {
		C.QLineEdit_DisconnectEditingFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "editingFinished")
	}
}

//export callbackQLineEditEditingFinished
func callbackQLineEditEditingFinished(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QLineEdit::editingFinished")

	if signal := qt.GetSignal(C.GoString(ptrName), "editingFinished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLineEdit) EditingFinished() {
	defer qt.Recovering("QLineEdit::editingFinished")

	if ptr.Pointer() != nil {
		C.QLineEdit_EditingFinished(ptr.Pointer())
	}
}

func (ptr *QLineEdit) End(mark bool) {
	defer qt.Recovering("QLineEdit::end")

	if ptr.Pointer() != nil {
		C.QLineEdit_End(ptr.Pointer(), C.int(qt.GoBoolToInt(mark)))
	}
}

func (ptr *QLineEdit) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QLineEdit::event")

	if ptr.Pointer() != nil {
		return C.QLineEdit_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QLineEdit) ConnectFocusInEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QLineEdit::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QLineEdit) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QLineEdit::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQLineEditFocusInEvent
func callbackQLineEditFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QLineEdit::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
	} else {
		NewQLineEditFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(e))
	}
}

func (ptr *QLineEdit) FocusInEvent(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QLineEdit::focusInEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QLineEdit) FocusInEventDefault(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QLineEdit::focusInEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QLineEdit) ConnectFocusOutEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QLineEdit::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QLineEdit) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QLineEdit::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQLineEditFocusOutEvent
func callbackQLineEditFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QLineEdit::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
	} else {
		NewQLineEditFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(e))
	}
}

func (ptr *QLineEdit) FocusOutEvent(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QLineEdit::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QLineEdit) FocusOutEventDefault(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QLineEdit::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QLineEdit) GetTextMargins(left int, top int, right int, bottom int) {
	defer qt.Recovering("QLineEdit::getTextMargins")

	if ptr.Pointer() != nil {
		C.QLineEdit_GetTextMargins(ptr.Pointer(), C.int(left), C.int(top), C.int(right), C.int(bottom))
	}
}

func (ptr *QLineEdit) Home(mark bool) {
	defer qt.Recovering("QLineEdit::home")

	if ptr.Pointer() != nil {
		C.QLineEdit_Home(ptr.Pointer(), C.int(qt.GoBoolToInt(mark)))
	}
}

func (ptr *QLineEdit) ConnectInputMethodEvent(f func(e *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QLineEdit::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QLineEdit) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QLineEdit::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQLineEditInputMethodEvent
func callbackQLineEditInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QLineEdit::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(e))
	} else {
		NewQLineEditFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(e))
	}
}

func (ptr *QLineEdit) InputMethodEvent(e gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QLineEdit::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(e))
	}
}

func (ptr *QLineEdit) InputMethodEventDefault(e gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QLineEdit::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(e))
	}
}

func (ptr *QLineEdit) InputMethodQuery(property core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QLineEdit::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QLineEdit_InputMethodQuery(ptr.Pointer(), C.int(property)))
	}
	return nil
}

func (ptr *QLineEdit) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QLineEdit::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QLineEdit) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QLineEdit::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQLineEditKeyPressEvent
func callbackQLineEditKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLineEdit::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQLineEditFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QLineEdit) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QLineEdit::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QLineEdit) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QLineEdit::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QLineEdit) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QLineEdit::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QLineEdit_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLineEdit) ConnectMouseDoubleClickEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QLineEdit::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QLineEdit) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QLineEdit::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQLineEditMouseDoubleClickEvent
func callbackQLineEditMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QLineEdit::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQLineEditFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QLineEdit) MouseDoubleClickEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QLineEdit::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QLineEdit) MouseDoubleClickEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QLineEdit::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QLineEdit) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QLineEdit::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QLineEdit) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QLineEdit::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQLineEditMouseMoveEvent
func callbackQLineEditMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QLineEdit::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQLineEditFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QLineEdit) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QLineEdit::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QLineEdit) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QLineEdit::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QLineEdit) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QLineEdit::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QLineEdit) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QLineEdit::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQLineEditMousePressEvent
func callbackQLineEditMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QLineEdit::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQLineEditFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QLineEdit) MousePressEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QLineEdit::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QLineEdit) MousePressEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QLineEdit::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QLineEdit) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QLineEdit::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QLineEdit) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QLineEdit::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQLineEditMouseReleaseEvent
func callbackQLineEditMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QLineEdit::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQLineEditFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QLineEdit) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QLineEdit::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QLineEdit) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QLineEdit::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QLineEdit) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QLineEdit::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QLineEdit) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QLineEdit::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQLineEditPaintEvent
func callbackQLineEditPaintEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QLineEdit::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
	} else {
		NewQLineEditFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(v))
	}
}

func (ptr *QLineEdit) PaintEvent(v gui.QPaintEvent_ITF) {
	defer qt.Recovering("QLineEdit::paintEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(v))
	}
}

func (ptr *QLineEdit) PaintEventDefault(v gui.QPaintEvent_ITF) {
	defer qt.Recovering("QLineEdit::paintEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(v))
	}
}

func (ptr *QLineEdit) Paste() {
	defer qt.Recovering("QLineEdit::paste")

	if ptr.Pointer() != nil {
		C.QLineEdit_Paste(ptr.Pointer())
	}
}

func (ptr *QLineEdit) Redo() {
	defer qt.Recovering("QLineEdit::redo")

	if ptr.Pointer() != nil {
		C.QLineEdit_Redo(ptr.Pointer())
	}
}

func (ptr *QLineEdit) ConnectReturnPressed(f func()) {
	defer qt.Recovering("connect QLineEdit::returnPressed")

	if ptr.Pointer() != nil {
		C.QLineEdit_ConnectReturnPressed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "returnPressed", f)
	}
}

func (ptr *QLineEdit) DisconnectReturnPressed() {
	defer qt.Recovering("disconnect QLineEdit::returnPressed")

	if ptr.Pointer() != nil {
		C.QLineEdit_DisconnectReturnPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "returnPressed")
	}
}

//export callbackQLineEditReturnPressed
func callbackQLineEditReturnPressed(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QLineEdit::returnPressed")

	if signal := qt.GetSignal(C.GoString(ptrName), "returnPressed"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLineEdit) ReturnPressed() {
	defer qt.Recovering("QLineEdit::returnPressed")

	if ptr.Pointer() != nil {
		C.QLineEdit_ReturnPressed(ptr.Pointer())
	}
}

func (ptr *QLineEdit) SelectAll() {
	defer qt.Recovering("QLineEdit::selectAll")

	if ptr.Pointer() != nil {
		C.QLineEdit_SelectAll(ptr.Pointer())
	}
}

func (ptr *QLineEdit) ConnectSelectionChanged(f func()) {
	defer qt.Recovering("connect QLineEdit::selectionChanged")

	if ptr.Pointer() != nil {
		C.QLineEdit_ConnectSelectionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "selectionChanged", f)
	}
}

func (ptr *QLineEdit) DisconnectSelectionChanged() {
	defer qt.Recovering("disconnect QLineEdit::selectionChanged")

	if ptr.Pointer() != nil {
		C.QLineEdit_DisconnectSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "selectionChanged")
	}
}

//export callbackQLineEditSelectionChanged
func callbackQLineEditSelectionChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QLineEdit::selectionChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectionChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLineEdit) SelectionChanged() {
	defer qt.Recovering("QLineEdit::selectionChanged")

	if ptr.Pointer() != nil {
		C.QLineEdit_SelectionChanged(ptr.Pointer())
	}
}

func (ptr *QLineEdit) SelectionStart() int {
	defer qt.Recovering("QLineEdit::selectionStart")

	if ptr.Pointer() != nil {
		return int(C.QLineEdit_SelectionStart(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLineEdit) SetCompleter(c QCompleter_ITF) {
	defer qt.Recovering("QLineEdit::setCompleter")

	if ptr.Pointer() != nil {
		C.QLineEdit_SetCompleter(ptr.Pointer(), PointerFromQCompleter(c))
	}
}

func (ptr *QLineEdit) SetSelection(start int, length int) {
	defer qt.Recovering("QLineEdit::setSelection")

	if ptr.Pointer() != nil {
		C.QLineEdit_SetSelection(ptr.Pointer(), C.int(start), C.int(length))
	}
}

func (ptr *QLineEdit) SetTextMargins2(margins core.QMargins_ITF) {
	defer qt.Recovering("QLineEdit::setTextMargins")

	if ptr.Pointer() != nil {
		C.QLineEdit_SetTextMargins2(ptr.Pointer(), core.PointerFromQMargins(margins))
	}
}

func (ptr *QLineEdit) SetTextMargins(left int, top int, right int, bottom int) {
	defer qt.Recovering("QLineEdit::setTextMargins")

	if ptr.Pointer() != nil {
		C.QLineEdit_SetTextMargins(ptr.Pointer(), C.int(left), C.int(top), C.int(right), C.int(bottom))
	}
}

func (ptr *QLineEdit) SetValidator(v gui.QValidator_ITF) {
	defer qt.Recovering("QLineEdit::setValidator")

	if ptr.Pointer() != nil {
		C.QLineEdit_SetValidator(ptr.Pointer(), gui.PointerFromQValidator(v))
	}
}

func (ptr *QLineEdit) SizeHint() *core.QSize {
	defer qt.Recovering("QLineEdit::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QLineEdit_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLineEdit) ConnectTextChanged(f func(text string)) {
	defer qt.Recovering("connect QLineEdit::textChanged")

	if ptr.Pointer() != nil {
		C.QLineEdit_ConnectTextChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "textChanged", f)
	}
}

func (ptr *QLineEdit) DisconnectTextChanged() {
	defer qt.Recovering("disconnect QLineEdit::textChanged")

	if ptr.Pointer() != nil {
		C.QLineEdit_DisconnectTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "textChanged")
	}
}

//export callbackQLineEditTextChanged
func callbackQLineEditTextChanged(ptr unsafe.Pointer, ptrName *C.char, text *C.char) {
	defer qt.Recovering("callback QLineEdit::textChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "textChanged"); signal != nil {
		signal.(func(string))(C.GoString(text))
	}

}

func (ptr *QLineEdit) TextChanged(text string) {
	defer qt.Recovering("QLineEdit::textChanged")

	if ptr.Pointer() != nil {
		C.QLineEdit_TextChanged(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QLineEdit) ConnectTextEdited(f func(text string)) {
	defer qt.Recovering("connect QLineEdit::textEdited")

	if ptr.Pointer() != nil {
		C.QLineEdit_ConnectTextEdited(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "textEdited", f)
	}
}

func (ptr *QLineEdit) DisconnectTextEdited() {
	defer qt.Recovering("disconnect QLineEdit::textEdited")

	if ptr.Pointer() != nil {
		C.QLineEdit_DisconnectTextEdited(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "textEdited")
	}
}

//export callbackQLineEditTextEdited
func callbackQLineEditTextEdited(ptr unsafe.Pointer, ptrName *C.char, text *C.char) {
	defer qt.Recovering("callback QLineEdit::textEdited")

	if signal := qt.GetSignal(C.GoString(ptrName), "textEdited"); signal != nil {
		signal.(func(string))(C.GoString(text))
	}

}

func (ptr *QLineEdit) TextEdited(text string) {
	defer qt.Recovering("QLineEdit::textEdited")

	if ptr.Pointer() != nil {
		C.QLineEdit_TextEdited(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QLineEdit) Undo() {
	defer qt.Recovering("QLineEdit::undo")

	if ptr.Pointer() != nil {
		C.QLineEdit_Undo(ptr.Pointer())
	}
}

func (ptr *QLineEdit) Validator() *gui.QValidator {
	defer qt.Recovering("QLineEdit::validator")

	if ptr.Pointer() != nil {
		return gui.NewQValidatorFromPointer(C.QLineEdit_Validator(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLineEdit) DestroyQLineEdit() {
	defer qt.Recovering("QLineEdit::~QLineEdit")

	if ptr.Pointer() != nil {
		C.QLineEdit_DestroyQLineEdit(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLineEdit) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QLineEdit::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QLineEdit) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QLineEdit::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQLineEditActionEvent
func callbackQLineEditActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLineEdit::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQLineEditFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QLineEdit) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QLineEdit::actionEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QLineEdit) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QLineEdit::actionEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QLineEdit) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QLineEdit::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QLineEdit) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QLineEdit::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQLineEditEnterEvent
func callbackQLineEditEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLineEdit::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQLineEditFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLineEdit) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QLineEdit::enterEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLineEdit) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QLineEdit::enterEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLineEdit) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QLineEdit::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QLineEdit) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QLineEdit::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQLineEditHideEvent
func callbackQLineEditHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLineEdit::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQLineEditFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QLineEdit) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QLineEdit::hideEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QLineEdit) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QLineEdit::hideEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QLineEdit) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QLineEdit::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QLineEdit) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QLineEdit::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQLineEditLeaveEvent
func callbackQLineEditLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLineEdit::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQLineEditFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLineEdit) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QLineEdit::leaveEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLineEdit) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QLineEdit::leaveEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLineEdit) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QLineEdit::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QLineEdit) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QLineEdit::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQLineEditMoveEvent
func callbackQLineEditMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLineEdit::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQLineEditFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QLineEdit) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QLineEdit::moveEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QLineEdit) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QLineEdit::moveEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QLineEdit) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QLineEdit::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QLineEdit) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QLineEdit::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQLineEditSetVisible
func callbackQLineEditSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QLineEdit::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QLineEdit) SetVisible(visible bool) {
	defer qt.Recovering("QLineEdit::setVisible")

	if ptr.Pointer() != nil {
		C.QLineEdit_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QLineEdit) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QLineEdit::setVisible")

	if ptr.Pointer() != nil {
		C.QLineEdit_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QLineEdit) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QLineEdit::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QLineEdit) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QLineEdit::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQLineEditShowEvent
func callbackQLineEditShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLineEdit::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQLineEditFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QLineEdit) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QLineEdit::showEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QLineEdit) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QLineEdit::showEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QLineEdit) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QLineEdit::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QLineEdit) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QLineEdit::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQLineEditCloseEvent
func callbackQLineEditCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLineEdit::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQLineEditFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QLineEdit) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QLineEdit::closeEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QLineEdit) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QLineEdit::closeEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QLineEdit) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QLineEdit::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QLineEdit) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QLineEdit::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQLineEditInitPainter
func callbackQLineEditInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QLineEdit::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQLineEditFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QLineEdit) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QLineEdit::initPainter")

	if ptr.Pointer() != nil {
		C.QLineEdit_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QLineEdit) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QLineEdit::initPainter")

	if ptr.Pointer() != nil {
		C.QLineEdit_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QLineEdit) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QLineEdit::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QLineEdit) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QLineEdit::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQLineEditKeyReleaseEvent
func callbackQLineEditKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLineEdit::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQLineEditFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QLineEdit) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QLineEdit::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QLineEdit) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QLineEdit::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QLineEdit) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QLineEdit::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QLineEdit) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QLineEdit::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQLineEditResizeEvent
func callbackQLineEditResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLineEdit::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQLineEditFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QLineEdit) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QLineEdit::resizeEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QLineEdit) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QLineEdit::resizeEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QLineEdit) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QLineEdit::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QLineEdit) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QLineEdit::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQLineEditTabletEvent
func callbackQLineEditTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLineEdit::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQLineEditFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QLineEdit) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QLineEdit::tabletEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QLineEdit) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QLineEdit::tabletEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QLineEdit) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QLineEdit::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QLineEdit) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QLineEdit::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQLineEditWheelEvent
func callbackQLineEditWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLineEdit::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQLineEditFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QLineEdit) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QLineEdit::wheelEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QLineEdit) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QLineEdit::wheelEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QLineEdit) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QLineEdit::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QLineEdit) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QLineEdit::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQLineEditTimerEvent
func callbackQLineEditTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLineEdit::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLineEditFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLineEdit) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLineEdit::timerEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLineEdit) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLineEdit::timerEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLineEdit) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QLineEdit::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QLineEdit) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QLineEdit::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQLineEditChildEvent
func callbackQLineEditChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLineEdit::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQLineEditFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QLineEdit) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QLineEdit::childEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QLineEdit) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QLineEdit::childEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QLineEdit) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QLineEdit::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QLineEdit) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QLineEdit::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQLineEditCustomEvent
func callbackQLineEditCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLineEdit::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQLineEditFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLineEdit) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QLineEdit::customEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLineEdit) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QLineEdit::customEvent")

	if ptr.Pointer() != nil {
		C.QLineEdit_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
