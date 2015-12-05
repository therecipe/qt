package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::alignment")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QLineEdit_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLineEdit) CursorMoveStyle() core.Qt__CursorMoveStyle {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::cursorMoveStyle")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__CursorMoveStyle(C.QLineEdit_CursorMoveStyle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLineEdit) CursorPosition() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::cursorPosition")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QLineEdit_CursorPosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLineEdit) DisplayText() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::displayText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLineEdit_DisplayText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLineEdit) DragEnabled() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::dragEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLineEdit_DragEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLineEdit) EchoMode() QLineEdit__EchoMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::echoMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QLineEdit__EchoMode(C.QLineEdit_EchoMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLineEdit) HasAcceptableInput() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::hasAcceptableInput")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLineEdit_HasAcceptableInput(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLineEdit) HasFrame() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::hasFrame")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLineEdit_HasFrame(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLineEdit) HasSelectedText() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::hasSelectedText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLineEdit_HasSelectedText(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLineEdit) InputMask() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::inputMask")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLineEdit_InputMask(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLineEdit) IsClearButtonEnabled() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::isClearButtonEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLineEdit_IsClearButtonEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLineEdit) IsModified() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::isModified")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLineEdit_IsModified(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLineEdit) IsReadOnly() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::isReadOnly")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLineEdit_IsReadOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLineEdit) IsRedoAvailable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::isRedoAvailable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLineEdit_IsRedoAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLineEdit) IsUndoAvailable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::isUndoAvailable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLineEdit_IsUndoAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLineEdit) MaxLength() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::maxLength")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QLineEdit_MaxLength(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLineEdit) PlaceholderText() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::placeholderText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLineEdit_PlaceholderText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLineEdit) SelectedText() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::selectedText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLineEdit_SelectedText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLineEdit) SetAlignment(flag core.Qt__AlignmentFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::setAlignment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_SetAlignment(ptr.Pointer(), C.int(flag))
	}
}

func (ptr *QLineEdit) SetClearButtonEnabled(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::setClearButtonEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_SetClearButtonEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QLineEdit) SetCursorMoveStyle(style core.Qt__CursorMoveStyle) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::setCursorMoveStyle")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_SetCursorMoveStyle(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QLineEdit) SetCursorPosition(v int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::setCursorPosition")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_SetCursorPosition(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLineEdit) SetDragEnabled(b bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::setDragEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_SetDragEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QLineEdit) SetEchoMode(v QLineEdit__EchoMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::setEchoMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_SetEchoMode(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLineEdit) SetFrame(v bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::setFrame")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_SetFrame(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QLineEdit) SetInputMask(inputMask string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::setInputMask")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_SetInputMask(ptr.Pointer(), C.CString(inputMask))
	}
}

func (ptr *QLineEdit) SetMaxLength(v int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::setMaxLength")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_SetMaxLength(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLineEdit) SetModified(v bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::setModified")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_SetModified(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QLineEdit) SetPlaceholderText(v string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::setPlaceholderText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_SetPlaceholderText(ptr.Pointer(), C.CString(v))
	}
}

func (ptr *QLineEdit) SetReadOnly(v bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::setReadOnly")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_SetReadOnly(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QLineEdit) SetText(v string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::setText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_SetText(ptr.Pointer(), C.CString(v))
	}
}

func (ptr *QLineEdit) Text() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::text")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLineEdit_Text(ptr.Pointer()))
	}
	return ""
}

func NewQLineEdit(parent QWidget_ITF) *QLineEdit {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::QLineEdit")
		}
	}()

	return NewQLineEditFromPointer(C.QLineEdit_NewQLineEdit(PointerFromQWidget(parent)))
}

func NewQLineEdit2(contents string, parent QWidget_ITF) *QLineEdit {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::QLineEdit")
		}
	}()

	return NewQLineEditFromPointer(C.QLineEdit_NewQLineEdit2(C.CString(contents), PointerFromQWidget(parent)))
}

func (ptr *QLineEdit) AddAction2(icon gui.QIcon_ITF, position QLineEdit__ActionPosition) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::addAction")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QLineEdit_AddAction2(ptr.Pointer(), gui.PointerFromQIcon(icon), C.int(position)))
	}
	return nil
}

func (ptr *QLineEdit) AddAction(action QAction_ITF, position QLineEdit__ActionPosition) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::addAction")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_AddAction(ptr.Pointer(), PointerFromQAction(action), C.int(position))
	}
}

func (ptr *QLineEdit) Backspace() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::backspace")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_Backspace(ptr.Pointer())
	}
}

func (ptr *QLineEdit) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_Clear(ptr.Pointer())
	}
}

func (ptr *QLineEdit) Completer() *QCompleter {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::completer")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQCompleterFromPointer(C.QLineEdit_Completer(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLineEdit) Copy() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::copy")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_Copy(ptr.Pointer())
	}
}

func (ptr *QLineEdit) CreateStandardContextMenu() *QMenu {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::createStandardContextMenu")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QLineEdit_CreateStandardContextMenu(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLineEdit) CursorBackward(mark bool, steps int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::cursorBackward")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_CursorBackward(ptr.Pointer(), C.int(qt.GoBoolToInt(mark)), C.int(steps))
	}
}

func (ptr *QLineEdit) CursorForward(mark bool, steps int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::cursorForward")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_CursorForward(ptr.Pointer(), C.int(qt.GoBoolToInt(mark)), C.int(steps))
	}
}

func (ptr *QLineEdit) CursorPositionAt(pos core.QPoint_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::cursorPositionAt")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QLineEdit_CursorPositionAt(ptr.Pointer(), core.PointerFromQPoint(pos)))
	}
	return 0
}

func (ptr *QLineEdit) ConnectCursorPositionChanged(f func(old int, n int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::cursorPositionChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_ConnectCursorPositionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cursorPositionChanged", f)
	}
}

func (ptr *QLineEdit) DisconnectCursorPositionChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::cursorPositionChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_DisconnectCursorPositionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cursorPositionChanged")
	}
}

//export callbackQLineEditCursorPositionChanged
func callbackQLineEditCursorPositionChanged(ptrName *C.char, old C.int, n C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::cursorPositionChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "cursorPositionChanged").(func(int, int))(int(old), int(n))
}

func (ptr *QLineEdit) CursorWordBackward(mark bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::cursorWordBackward")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_CursorWordBackward(ptr.Pointer(), C.int(qt.GoBoolToInt(mark)))
	}
}

func (ptr *QLineEdit) CursorWordForward(mark bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::cursorWordForward")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_CursorWordForward(ptr.Pointer(), C.int(qt.GoBoolToInt(mark)))
	}
}

func (ptr *QLineEdit) Cut() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::cut")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_Cut(ptr.Pointer())
	}
}

func (ptr *QLineEdit) Del() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::del")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_Del(ptr.Pointer())
	}
}

func (ptr *QLineEdit) Deselect() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::deselect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_Deselect(ptr.Pointer())
	}
}

func (ptr *QLineEdit) ConnectEditingFinished(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::editingFinished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_ConnectEditingFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "editingFinished", f)
	}
}

func (ptr *QLineEdit) DisconnectEditingFinished() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::editingFinished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_DisconnectEditingFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "editingFinished")
	}
}

//export callbackQLineEditEditingFinished
func callbackQLineEditEditingFinished(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::editingFinished")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "editingFinished").(func())()
}

func (ptr *QLineEdit) End(mark bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::end")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_End(ptr.Pointer(), C.int(qt.GoBoolToInt(mark)))
	}
}

func (ptr *QLineEdit) Event(e core.QEvent_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::event")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLineEdit_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QLineEdit) GetTextMargins(left int, top int, right int, bottom int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::getTextMargins")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_GetTextMargins(ptr.Pointer(), C.int(left), C.int(top), C.int(right), C.int(bottom))
	}
}

func (ptr *QLineEdit) Home(mark bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::home")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_Home(ptr.Pointer(), C.int(qt.GoBoolToInt(mark)))
	}
}

func (ptr *QLineEdit) InputMethodQuery(property core.Qt__InputMethodQuery) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::inputMethodQuery")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QLineEdit_InputMethodQuery(ptr.Pointer(), C.int(property)))
	}
	return nil
}

func (ptr *QLineEdit) Paste() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::paste")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_Paste(ptr.Pointer())
	}
}

func (ptr *QLineEdit) Redo() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::redo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_Redo(ptr.Pointer())
	}
}

func (ptr *QLineEdit) ConnectReturnPressed(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::returnPressed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_ConnectReturnPressed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "returnPressed", f)
	}
}

func (ptr *QLineEdit) DisconnectReturnPressed() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::returnPressed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_DisconnectReturnPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "returnPressed")
	}
}

//export callbackQLineEditReturnPressed
func callbackQLineEditReturnPressed(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::returnPressed")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "returnPressed").(func())()
}

func (ptr *QLineEdit) SelectAll() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::selectAll")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_SelectAll(ptr.Pointer())
	}
}

func (ptr *QLineEdit) ConnectSelectionChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::selectionChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_ConnectSelectionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "selectionChanged", f)
	}
}

func (ptr *QLineEdit) DisconnectSelectionChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::selectionChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_DisconnectSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "selectionChanged")
	}
}

//export callbackQLineEditSelectionChanged
func callbackQLineEditSelectionChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::selectionChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "selectionChanged").(func())()
}

func (ptr *QLineEdit) SelectionStart() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::selectionStart")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QLineEdit_SelectionStart(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLineEdit) SetCompleter(c QCompleter_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::setCompleter")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_SetCompleter(ptr.Pointer(), PointerFromQCompleter(c))
	}
}

func (ptr *QLineEdit) SetSelection(start int, length int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::setSelection")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_SetSelection(ptr.Pointer(), C.int(start), C.int(length))
	}
}

func (ptr *QLineEdit) SetTextMargins2(margins core.QMargins_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::setTextMargins")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_SetTextMargins2(ptr.Pointer(), core.PointerFromQMargins(margins))
	}
}

func (ptr *QLineEdit) SetTextMargins(left int, top int, right int, bottom int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::setTextMargins")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_SetTextMargins(ptr.Pointer(), C.int(left), C.int(top), C.int(right), C.int(bottom))
	}
}

func (ptr *QLineEdit) SetValidator(v gui.QValidator_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::setValidator")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_SetValidator(ptr.Pointer(), gui.PointerFromQValidator(v))
	}
}

func (ptr *QLineEdit) ConnectTextChanged(f func(text string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::textChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_ConnectTextChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "textChanged", f)
	}
}

func (ptr *QLineEdit) DisconnectTextChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::textChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_DisconnectTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "textChanged")
	}
}

//export callbackQLineEditTextChanged
func callbackQLineEditTextChanged(ptrName *C.char, text *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::textChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "textChanged").(func(string))(C.GoString(text))
}

func (ptr *QLineEdit) ConnectTextEdited(f func(text string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::textEdited")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_ConnectTextEdited(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "textEdited", f)
	}
}

func (ptr *QLineEdit) DisconnectTextEdited() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::textEdited")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_DisconnectTextEdited(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "textEdited")
	}
}

//export callbackQLineEditTextEdited
func callbackQLineEditTextEdited(ptrName *C.char, text *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::textEdited")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "textEdited").(func(string))(C.GoString(text))
}

func (ptr *QLineEdit) Undo() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::undo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_Undo(ptr.Pointer())
	}
}

func (ptr *QLineEdit) Validator() *gui.QValidator {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::validator")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQValidatorFromPointer(C.QLineEdit_Validator(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLineEdit) DestroyQLineEdit() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLineEdit::~QLineEdit")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLineEdit_DestroyQLineEdit(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
