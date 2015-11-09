package widgets

//#include "qlineedit.h"
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
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QLineEdit_" + qt.RandomIdentifier())
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
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QLineEdit_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLineEdit) CursorMoveStyle() core.Qt__CursorMoveStyle {
	if ptr.Pointer() != nil {
		return core.Qt__CursorMoveStyle(C.QLineEdit_CursorMoveStyle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLineEdit) CursorPosition() int {
	if ptr.Pointer() != nil {
		return int(C.QLineEdit_CursorPosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLineEdit) DisplayText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QLineEdit_DisplayText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLineEdit) DragEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QLineEdit_DragEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLineEdit) EchoMode() QLineEdit__EchoMode {
	if ptr.Pointer() != nil {
		return QLineEdit__EchoMode(C.QLineEdit_EchoMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLineEdit) HasAcceptableInput() bool {
	if ptr.Pointer() != nil {
		return C.QLineEdit_HasAcceptableInput(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLineEdit) HasFrame() bool {
	if ptr.Pointer() != nil {
		return C.QLineEdit_HasFrame(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLineEdit) HasSelectedText() bool {
	if ptr.Pointer() != nil {
		return C.QLineEdit_HasSelectedText(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLineEdit) InputMask() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QLineEdit_InputMask(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLineEdit) IsClearButtonEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QLineEdit_IsClearButtonEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLineEdit) IsModified() bool {
	if ptr.Pointer() != nil {
		return C.QLineEdit_IsModified(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLineEdit) IsReadOnly() bool {
	if ptr.Pointer() != nil {
		return C.QLineEdit_IsReadOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLineEdit) IsRedoAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QLineEdit_IsRedoAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLineEdit) IsUndoAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QLineEdit_IsUndoAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLineEdit) MaxLength() int {
	if ptr.Pointer() != nil {
		return int(C.QLineEdit_MaxLength(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLineEdit) PlaceholderText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QLineEdit_PlaceholderText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLineEdit) SelectedText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QLineEdit_SelectedText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLineEdit) SetAlignment(flag core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetAlignment(ptr.Pointer(), C.int(flag))
	}
}

func (ptr *QLineEdit) SetClearButtonEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetClearButtonEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QLineEdit) SetCursorMoveStyle(style core.Qt__CursorMoveStyle) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetCursorMoveStyle(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QLineEdit) SetCursorPosition(v int) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetCursorPosition(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLineEdit) SetDragEnabled(b bool) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetDragEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QLineEdit) SetEchoMode(v QLineEdit__EchoMode) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetEchoMode(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLineEdit) SetFrame(v bool) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetFrame(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QLineEdit) SetInputMask(inputMask string) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetInputMask(ptr.Pointer(), C.CString(inputMask))
	}
}

func (ptr *QLineEdit) SetMaxLength(v int) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetMaxLength(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLineEdit) SetModified(v bool) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetModified(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QLineEdit) SetPlaceholderText(v string) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetPlaceholderText(ptr.Pointer(), C.CString(v))
	}
}

func (ptr *QLineEdit) SetReadOnly(v bool) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetReadOnly(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QLineEdit) SetText(v string) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetText(ptr.Pointer(), C.CString(v))
	}
}

func (ptr *QLineEdit) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QLineEdit_Text(ptr.Pointer()))
	}
	return ""
}

func NewQLineEdit(parent QWidget_ITF) *QLineEdit {
	return NewQLineEditFromPointer(C.QLineEdit_NewQLineEdit(PointerFromQWidget(parent)))
}

func NewQLineEdit2(contents string, parent QWidget_ITF) *QLineEdit {
	return NewQLineEditFromPointer(C.QLineEdit_NewQLineEdit2(C.CString(contents), PointerFromQWidget(parent)))
}

func (ptr *QLineEdit) AddAction2(icon gui.QIcon_ITF, position QLineEdit__ActionPosition) *QAction {
	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QLineEdit_AddAction2(ptr.Pointer(), gui.PointerFromQIcon(icon), C.int(position)))
	}
	return nil
}

func (ptr *QLineEdit) AddAction(action QAction_ITF, position QLineEdit__ActionPosition) {
	if ptr.Pointer() != nil {
		C.QLineEdit_AddAction(ptr.Pointer(), PointerFromQAction(action), C.int(position))
	}
}

func (ptr *QLineEdit) Backspace() {
	if ptr.Pointer() != nil {
		C.QLineEdit_Backspace(ptr.Pointer())
	}
}

func (ptr *QLineEdit) Clear() {
	if ptr.Pointer() != nil {
		C.QLineEdit_Clear(ptr.Pointer())
	}
}

func (ptr *QLineEdit) Completer() *QCompleter {
	if ptr.Pointer() != nil {
		return NewQCompleterFromPointer(C.QLineEdit_Completer(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLineEdit) Copy() {
	if ptr.Pointer() != nil {
		C.QLineEdit_Copy(ptr.Pointer())
	}
}

func (ptr *QLineEdit) CreateStandardContextMenu() *QMenu {
	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QLineEdit_CreateStandardContextMenu(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLineEdit) CursorBackward(mark bool, steps int) {
	if ptr.Pointer() != nil {
		C.QLineEdit_CursorBackward(ptr.Pointer(), C.int(qt.GoBoolToInt(mark)), C.int(steps))
	}
}

func (ptr *QLineEdit) CursorForward(mark bool, steps int) {
	if ptr.Pointer() != nil {
		C.QLineEdit_CursorForward(ptr.Pointer(), C.int(qt.GoBoolToInt(mark)), C.int(steps))
	}
}

func (ptr *QLineEdit) CursorPositionAt(pos core.QPoint_ITF) int {
	if ptr.Pointer() != nil {
		return int(C.QLineEdit_CursorPositionAt(ptr.Pointer(), core.PointerFromQPoint(pos)))
	}
	return 0
}

func (ptr *QLineEdit) ConnectCursorPositionChanged(f func(old int, n int)) {
	if ptr.Pointer() != nil {
		C.QLineEdit_ConnectCursorPositionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cursorPositionChanged", f)
	}
}

func (ptr *QLineEdit) DisconnectCursorPositionChanged() {
	if ptr.Pointer() != nil {
		C.QLineEdit_DisconnectCursorPositionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cursorPositionChanged")
	}
}

//export callbackQLineEditCursorPositionChanged
func callbackQLineEditCursorPositionChanged(ptrName *C.char, old C.int, n C.int) {
	qt.GetSignal(C.GoString(ptrName), "cursorPositionChanged").(func(int, int))(int(old), int(n))
}

func (ptr *QLineEdit) CursorWordBackward(mark bool) {
	if ptr.Pointer() != nil {
		C.QLineEdit_CursorWordBackward(ptr.Pointer(), C.int(qt.GoBoolToInt(mark)))
	}
}

func (ptr *QLineEdit) CursorWordForward(mark bool) {
	if ptr.Pointer() != nil {
		C.QLineEdit_CursorWordForward(ptr.Pointer(), C.int(qt.GoBoolToInt(mark)))
	}
}

func (ptr *QLineEdit) Cut() {
	if ptr.Pointer() != nil {
		C.QLineEdit_Cut(ptr.Pointer())
	}
}

func (ptr *QLineEdit) Del() {
	if ptr.Pointer() != nil {
		C.QLineEdit_Del(ptr.Pointer())
	}
}

func (ptr *QLineEdit) Deselect() {
	if ptr.Pointer() != nil {
		C.QLineEdit_Deselect(ptr.Pointer())
	}
}

func (ptr *QLineEdit) ConnectEditingFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QLineEdit_ConnectEditingFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "editingFinished", f)
	}
}

func (ptr *QLineEdit) DisconnectEditingFinished() {
	if ptr.Pointer() != nil {
		C.QLineEdit_DisconnectEditingFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "editingFinished")
	}
}

//export callbackQLineEditEditingFinished
func callbackQLineEditEditingFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "editingFinished").(func())()
}

func (ptr *QLineEdit) End(mark bool) {
	if ptr.Pointer() != nil {
		C.QLineEdit_End(ptr.Pointer(), C.int(qt.GoBoolToInt(mark)))
	}
}

func (ptr *QLineEdit) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLineEdit_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QLineEdit) GetTextMargins(left int, top int, right int, bottom int) {
	if ptr.Pointer() != nil {
		C.QLineEdit_GetTextMargins(ptr.Pointer(), C.int(left), C.int(top), C.int(right), C.int(bottom))
	}
}

func (ptr *QLineEdit) Home(mark bool) {
	if ptr.Pointer() != nil {
		C.QLineEdit_Home(ptr.Pointer(), C.int(qt.GoBoolToInt(mark)))
	}
}

func (ptr *QLineEdit) InputMethodQuery(property core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QLineEdit_InputMethodQuery(ptr.Pointer(), C.int(property)))
	}
	return nil
}

func (ptr *QLineEdit) Paste() {
	if ptr.Pointer() != nil {
		C.QLineEdit_Paste(ptr.Pointer())
	}
}

func (ptr *QLineEdit) Redo() {
	if ptr.Pointer() != nil {
		C.QLineEdit_Redo(ptr.Pointer())
	}
}

func (ptr *QLineEdit) ConnectReturnPressed(f func()) {
	if ptr.Pointer() != nil {
		C.QLineEdit_ConnectReturnPressed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "returnPressed", f)
	}
}

func (ptr *QLineEdit) DisconnectReturnPressed() {
	if ptr.Pointer() != nil {
		C.QLineEdit_DisconnectReturnPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "returnPressed")
	}
}

//export callbackQLineEditReturnPressed
func callbackQLineEditReturnPressed(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "returnPressed").(func())()
}

func (ptr *QLineEdit) SelectAll() {
	if ptr.Pointer() != nil {
		C.QLineEdit_SelectAll(ptr.Pointer())
	}
}

func (ptr *QLineEdit) ConnectSelectionChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QLineEdit_ConnectSelectionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "selectionChanged", f)
	}
}

func (ptr *QLineEdit) DisconnectSelectionChanged() {
	if ptr.Pointer() != nil {
		C.QLineEdit_DisconnectSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "selectionChanged")
	}
}

//export callbackQLineEditSelectionChanged
func callbackQLineEditSelectionChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "selectionChanged").(func())()
}

func (ptr *QLineEdit) SelectionStart() int {
	if ptr.Pointer() != nil {
		return int(C.QLineEdit_SelectionStart(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLineEdit) SetCompleter(c QCompleter_ITF) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetCompleter(ptr.Pointer(), PointerFromQCompleter(c))
	}
}

func (ptr *QLineEdit) SetSelection(start int, length int) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetSelection(ptr.Pointer(), C.int(start), C.int(length))
	}
}

func (ptr *QLineEdit) SetTextMargins2(margins core.QMargins_ITF) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetTextMargins2(ptr.Pointer(), core.PointerFromQMargins(margins))
	}
}

func (ptr *QLineEdit) SetTextMargins(left int, top int, right int, bottom int) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetTextMargins(ptr.Pointer(), C.int(left), C.int(top), C.int(right), C.int(bottom))
	}
}

func (ptr *QLineEdit) SetValidator(v gui.QValidator_ITF) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetValidator(ptr.Pointer(), gui.PointerFromQValidator(v))
	}
}

func (ptr *QLineEdit) ConnectTextChanged(f func(text string)) {
	if ptr.Pointer() != nil {
		C.QLineEdit_ConnectTextChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "textChanged", f)
	}
}

func (ptr *QLineEdit) DisconnectTextChanged() {
	if ptr.Pointer() != nil {
		C.QLineEdit_DisconnectTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "textChanged")
	}
}

//export callbackQLineEditTextChanged
func callbackQLineEditTextChanged(ptrName *C.char, text *C.char) {
	qt.GetSignal(C.GoString(ptrName), "textChanged").(func(string))(C.GoString(text))
}

func (ptr *QLineEdit) ConnectTextEdited(f func(text string)) {
	if ptr.Pointer() != nil {
		C.QLineEdit_ConnectTextEdited(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "textEdited", f)
	}
}

func (ptr *QLineEdit) DisconnectTextEdited() {
	if ptr.Pointer() != nil {
		C.QLineEdit_DisconnectTextEdited(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "textEdited")
	}
}

//export callbackQLineEditTextEdited
func callbackQLineEditTextEdited(ptrName *C.char, text *C.char) {
	qt.GetSignal(C.GoString(ptrName), "textEdited").(func(string))(C.GoString(text))
}

func (ptr *QLineEdit) Undo() {
	if ptr.Pointer() != nil {
		C.QLineEdit_Undo(ptr.Pointer())
	}
}

func (ptr *QLineEdit) Validator() *gui.QValidator {
	if ptr.Pointer() != nil {
		return gui.NewQValidatorFromPointer(C.QLineEdit_Validator(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLineEdit) DestroyQLineEdit() {
	if ptr.Pointer() != nil {
		C.QLineEdit_DestroyQLineEdit(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
