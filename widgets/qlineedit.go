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
func callbackQLineEditChangeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QLineEdit::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
		return true
	}
	return false

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
func callbackQLineEditContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLineEdit::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

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
func callbackQLineEditCursorPositionChanged(ptrName *C.char, old C.int, n C.int) {
	defer qt.Recovering("callback QLineEdit::cursorPositionChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "cursorPositionChanged"); signal != nil {
		signal.(func(int, int))(int(old), int(n))
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
func callbackQLineEditDragEnterEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QLineEdit::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(e))
		return true
	}
	return false

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
func callbackQLineEditDragLeaveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QLineEdit::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(e))
		return true
	}
	return false

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
func callbackQLineEditDragMoveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QLineEdit::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(e))
		return true
	}
	return false

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
func callbackQLineEditDropEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QLineEdit::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(e))
		return true
	}
	return false

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
func callbackQLineEditEditingFinished(ptrName *C.char) {
	defer qt.Recovering("callback QLineEdit::editingFinished")

	if signal := qt.GetSignal(C.GoString(ptrName), "editingFinished"); signal != nil {
		signal.(func())()
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
func callbackQLineEditFocusInEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QLineEdit::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
		return true
	}
	return false

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
func callbackQLineEditFocusOutEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QLineEdit::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
		return true
	}
	return false

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
func callbackQLineEditInputMethodEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QLineEdit::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(e))
		return true
	}
	return false

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
func callbackQLineEditKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLineEdit::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

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
func callbackQLineEditMouseDoubleClickEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QLineEdit::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

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
func callbackQLineEditMouseMoveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QLineEdit::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

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
func callbackQLineEditMousePressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QLineEdit::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

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
func callbackQLineEditMouseReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QLineEdit::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

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
func callbackQLineEditPaintEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QLineEdit::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
		return true
	}
	return false

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
func callbackQLineEditReturnPressed(ptrName *C.char) {
	defer qt.Recovering("callback QLineEdit::returnPressed")

	if signal := qt.GetSignal(C.GoString(ptrName), "returnPressed"); signal != nil {
		signal.(func())()
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
func callbackQLineEditSelectionChanged(ptrName *C.char) {
	defer qt.Recovering("callback QLineEdit::selectionChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectionChanged"); signal != nil {
		signal.(func())()
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
func callbackQLineEditTextChanged(ptrName *C.char, text *C.char) {
	defer qt.Recovering("callback QLineEdit::textChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "textChanged"); signal != nil {
		signal.(func(string))(C.GoString(text))
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
func callbackQLineEditTextEdited(ptrName *C.char, text *C.char) {
	defer qt.Recovering("callback QLineEdit::textEdited")

	if signal := qt.GetSignal(C.GoString(ptrName), "textEdited"); signal != nil {
		signal.(func(string))(C.GoString(text))
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
