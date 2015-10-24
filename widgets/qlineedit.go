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

type QLineEditITF interface {
	QWidgetITF
	QLineEditPTR() *QLineEdit
}

func PointerFromQLineEdit(ptr QLineEditITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLineEditPTR().Pointer()
	}
	return nil
}

func QLineEditFromPointer(ptr unsafe.Pointer) *QLineEdit {
	var n = new(QLineEdit)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QLineEdit_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QLineEdit) QLineEditPTR() *QLineEdit {
	return ptr
}

//QLineEdit::ActionPosition
type QLineEdit__ActionPosition int

var (
	QLineEdit__LeadingPosition  = QLineEdit__ActionPosition(0)
	QLineEdit__TrailingPosition = QLineEdit__ActionPosition(1)
)

//QLineEdit::EchoMode
type QLineEdit__EchoMode int

var (
	QLineEdit__Normal             = QLineEdit__EchoMode(0)
	QLineEdit__NoEcho             = QLineEdit__EchoMode(1)
	QLineEdit__Password           = QLineEdit__EchoMode(2)
	QLineEdit__PasswordEchoOnEdit = QLineEdit__EchoMode(3)
)

func (ptr *QLineEdit) Alignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QLineEdit_Alignment(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLineEdit) CursorMoveStyle() core.Qt__CursorMoveStyle {
	if ptr.Pointer() != nil {
		return core.Qt__CursorMoveStyle(C.QLineEdit_CursorMoveStyle(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLineEdit) CursorPosition() int {
	if ptr.Pointer() != nil {
		return int(C.QLineEdit_CursorPosition(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLineEdit) DisplayText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QLineEdit_DisplayText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QLineEdit) DragEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QLineEdit_DragEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLineEdit) EchoMode() QLineEdit__EchoMode {
	if ptr.Pointer() != nil {
		return QLineEdit__EchoMode(C.QLineEdit_EchoMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLineEdit) HasAcceptableInput() bool {
	if ptr.Pointer() != nil {
		return C.QLineEdit_HasAcceptableInput(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLineEdit) HasFrame() bool {
	if ptr.Pointer() != nil {
		return C.QLineEdit_HasFrame(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLineEdit) HasSelectedText() bool {
	if ptr.Pointer() != nil {
		return C.QLineEdit_HasSelectedText(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLineEdit) InputMask() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QLineEdit_InputMask(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QLineEdit) IsClearButtonEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QLineEdit_IsClearButtonEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLineEdit) IsModified() bool {
	if ptr.Pointer() != nil {
		return C.QLineEdit_IsModified(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLineEdit) IsReadOnly() bool {
	if ptr.Pointer() != nil {
		return C.QLineEdit_IsReadOnly(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLineEdit) IsRedoAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QLineEdit_IsRedoAvailable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLineEdit) IsUndoAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QLineEdit_IsUndoAvailable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLineEdit) MaxLength() int {
	if ptr.Pointer() != nil {
		return int(C.QLineEdit_MaxLength(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLineEdit) PlaceholderText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QLineEdit_PlaceholderText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QLineEdit) SelectedText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QLineEdit_SelectedText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QLineEdit) SetAlignment(flag core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetAlignment(C.QtObjectPtr(ptr.Pointer()), C.int(flag))
	}
}

func (ptr *QLineEdit) SetClearButtonEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetClearButtonEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QLineEdit) SetCursorMoveStyle(style core.Qt__CursorMoveStyle) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetCursorMoveStyle(C.QtObjectPtr(ptr.Pointer()), C.int(style))
	}
}

func (ptr *QLineEdit) SetCursorPosition(v int) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetCursorPosition(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QLineEdit) SetDragEnabled(b bool) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetDragEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QLineEdit) SetEchoMode(v QLineEdit__EchoMode) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetEchoMode(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QLineEdit) SetFrame(v bool) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetFrame(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QLineEdit) SetInputMask(inputMask string) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetInputMask(C.QtObjectPtr(ptr.Pointer()), C.CString(inputMask))
	}
}

func (ptr *QLineEdit) SetMaxLength(v int) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetMaxLength(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QLineEdit) SetModified(v bool) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetModified(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QLineEdit) SetPlaceholderText(v string) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetPlaceholderText(C.QtObjectPtr(ptr.Pointer()), C.CString(v))
	}
}

func (ptr *QLineEdit) SetReadOnly(v bool) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetReadOnly(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QLineEdit) SetText(v string) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetText(C.QtObjectPtr(ptr.Pointer()), C.CString(v))
	}
}

func (ptr *QLineEdit) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QLineEdit_Text(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func NewQLineEdit(parent QWidgetITF) *QLineEdit {
	return QLineEditFromPointer(unsafe.Pointer(C.QLineEdit_NewQLineEdit(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func NewQLineEdit2(contents string, parent QWidgetITF) *QLineEdit {
	return QLineEditFromPointer(unsafe.Pointer(C.QLineEdit_NewQLineEdit2(C.CString(contents), C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QLineEdit) AddAction2(icon gui.QIconITF, position QLineEdit__ActionPosition) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QLineEdit_AddAction2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQIcon(icon)), C.int(position))))
	}
	return nil
}

func (ptr *QLineEdit) AddAction(action QActionITF, position QLineEdit__ActionPosition) {
	if ptr.Pointer() != nil {
		C.QLineEdit_AddAction(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAction(action)), C.int(position))
	}
}

func (ptr *QLineEdit) Backspace() {
	if ptr.Pointer() != nil {
		C.QLineEdit_Backspace(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QLineEdit) Clear() {
	if ptr.Pointer() != nil {
		C.QLineEdit_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QLineEdit) Completer() *QCompleter {
	if ptr.Pointer() != nil {
		return QCompleterFromPointer(unsafe.Pointer(C.QLineEdit_Completer(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QLineEdit) Copy() {
	if ptr.Pointer() != nil {
		C.QLineEdit_Copy(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QLineEdit) CreateStandardContextMenu() *QMenu {
	if ptr.Pointer() != nil {
		return QMenuFromPointer(unsafe.Pointer(C.QLineEdit_CreateStandardContextMenu(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QLineEdit) CursorBackward(mark bool, steps int) {
	if ptr.Pointer() != nil {
		C.QLineEdit_CursorBackward(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(mark)), C.int(steps))
	}
}

func (ptr *QLineEdit) CursorForward(mark bool, steps int) {
	if ptr.Pointer() != nil {
		C.QLineEdit_CursorForward(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(mark)), C.int(steps))
	}
}

func (ptr *QLineEdit) CursorPositionAt(pos core.QPointITF) int {
	if ptr.Pointer() != nil {
		return int(C.QLineEdit_CursorPositionAt(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(pos))))
	}
	return 0
}

func (ptr *QLineEdit) ConnectCursorPositionChanged(f func(old int, n int)) {
	if ptr.Pointer() != nil {
		C.QLineEdit_ConnectCursorPositionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "cursorPositionChanged", f)
	}
}

func (ptr *QLineEdit) DisconnectCursorPositionChanged() {
	if ptr.Pointer() != nil {
		C.QLineEdit_DisconnectCursorPositionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "cursorPositionChanged")
	}
}

//export callbackQLineEditCursorPositionChanged
func callbackQLineEditCursorPositionChanged(ptrName *C.char, old C.int, n C.int) {
	qt.GetSignal(C.GoString(ptrName), "cursorPositionChanged").(func(int, int))(int(old), int(n))
}

func (ptr *QLineEdit) CursorWordBackward(mark bool) {
	if ptr.Pointer() != nil {
		C.QLineEdit_CursorWordBackward(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(mark)))
	}
}

func (ptr *QLineEdit) CursorWordForward(mark bool) {
	if ptr.Pointer() != nil {
		C.QLineEdit_CursorWordForward(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(mark)))
	}
}

func (ptr *QLineEdit) Cut() {
	if ptr.Pointer() != nil {
		C.QLineEdit_Cut(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QLineEdit) Del() {
	if ptr.Pointer() != nil {
		C.QLineEdit_Del(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QLineEdit) Deselect() {
	if ptr.Pointer() != nil {
		C.QLineEdit_Deselect(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QLineEdit) ConnectEditingFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QLineEdit_ConnectEditingFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "editingFinished", f)
	}
}

func (ptr *QLineEdit) DisconnectEditingFinished() {
	if ptr.Pointer() != nil {
		C.QLineEdit_DisconnectEditingFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "editingFinished")
	}
}

//export callbackQLineEditEditingFinished
func callbackQLineEditEditingFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "editingFinished").(func())()
}

func (ptr *QLineEdit) End(mark bool) {
	if ptr.Pointer() != nil {
		C.QLineEdit_End(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(mark)))
	}
}

func (ptr *QLineEdit) Event(e core.QEventITF) bool {
	if ptr.Pointer() != nil {
		return C.QLineEdit_Event(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQEvent(e))) != 0
	}
	return false
}

func (ptr *QLineEdit) GetTextMargins(left int, top int, right int, bottom int) {
	if ptr.Pointer() != nil {
		C.QLineEdit_GetTextMargins(C.QtObjectPtr(ptr.Pointer()), C.int(left), C.int(top), C.int(right), C.int(bottom))
	}
}

func (ptr *QLineEdit) Home(mark bool) {
	if ptr.Pointer() != nil {
		C.QLineEdit_Home(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(mark)))
	}
}

func (ptr *QLineEdit) InputMethodQuery(property core.Qt__InputMethodQuery) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QLineEdit_InputMethodQuery(C.QtObjectPtr(ptr.Pointer()), C.int(property)))
	}
	return ""
}

func (ptr *QLineEdit) Paste() {
	if ptr.Pointer() != nil {
		C.QLineEdit_Paste(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QLineEdit) Redo() {
	if ptr.Pointer() != nil {
		C.QLineEdit_Redo(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QLineEdit) ConnectReturnPressed(f func()) {
	if ptr.Pointer() != nil {
		C.QLineEdit_ConnectReturnPressed(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "returnPressed", f)
	}
}

func (ptr *QLineEdit) DisconnectReturnPressed() {
	if ptr.Pointer() != nil {
		C.QLineEdit_DisconnectReturnPressed(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "returnPressed")
	}
}

//export callbackQLineEditReturnPressed
func callbackQLineEditReturnPressed(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "returnPressed").(func())()
}

func (ptr *QLineEdit) SelectAll() {
	if ptr.Pointer() != nil {
		C.QLineEdit_SelectAll(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QLineEdit) ConnectSelectionChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QLineEdit_ConnectSelectionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "selectionChanged", f)
	}
}

func (ptr *QLineEdit) DisconnectSelectionChanged() {
	if ptr.Pointer() != nil {
		C.QLineEdit_DisconnectSelectionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "selectionChanged")
	}
}

//export callbackQLineEditSelectionChanged
func callbackQLineEditSelectionChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "selectionChanged").(func())()
}

func (ptr *QLineEdit) SelectionStart() int {
	if ptr.Pointer() != nil {
		return int(C.QLineEdit_SelectionStart(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLineEdit) SetCompleter(c QCompleterITF) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetCompleter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQCompleter(c)))
	}
}

func (ptr *QLineEdit) SetSelection(start int, length int) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetSelection(C.QtObjectPtr(ptr.Pointer()), C.int(start), C.int(length))
	}
}

func (ptr *QLineEdit) SetTextMargins2(margins core.QMarginsITF) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetTextMargins2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQMargins(margins)))
	}
}

func (ptr *QLineEdit) SetTextMargins(left int, top int, right int, bottom int) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetTextMargins(C.QtObjectPtr(ptr.Pointer()), C.int(left), C.int(top), C.int(right), C.int(bottom))
	}
}

func (ptr *QLineEdit) SetValidator(v gui.QValidatorITF) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetValidator(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQValidator(v)))
	}
}

func (ptr *QLineEdit) ConnectTextChanged(f func(text string)) {
	if ptr.Pointer() != nil {
		C.QLineEdit_ConnectTextChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "textChanged", f)
	}
}

func (ptr *QLineEdit) DisconnectTextChanged() {
	if ptr.Pointer() != nil {
		C.QLineEdit_DisconnectTextChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "textChanged")
	}
}

//export callbackQLineEditTextChanged
func callbackQLineEditTextChanged(ptrName *C.char, text *C.char) {
	qt.GetSignal(C.GoString(ptrName), "textChanged").(func(string))(C.GoString(text))
}

func (ptr *QLineEdit) ConnectTextEdited(f func(text string)) {
	if ptr.Pointer() != nil {
		C.QLineEdit_ConnectTextEdited(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "textEdited", f)
	}
}

func (ptr *QLineEdit) DisconnectTextEdited() {
	if ptr.Pointer() != nil {
		C.QLineEdit_DisconnectTextEdited(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "textEdited")
	}
}

//export callbackQLineEditTextEdited
func callbackQLineEditTextEdited(ptrName *C.char, text *C.char) {
	qt.GetSignal(C.GoString(ptrName), "textEdited").(func(string))(C.GoString(text))
}

func (ptr *QLineEdit) Undo() {
	if ptr.Pointer() != nil {
		C.QLineEdit_Undo(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QLineEdit) Validator() *gui.QValidator {
	if ptr.Pointer() != nil {
		return gui.QValidatorFromPointer(unsafe.Pointer(C.QLineEdit_Validator(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QLineEdit) DestroyQLineEdit() {
	if ptr.Pointer() != nil {
		C.QLineEdit_DestroyQLineEdit(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
