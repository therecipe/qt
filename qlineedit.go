package qt

//#include "qlineedit.h"
import "C"

type qlineedit struct {
	qwidget
}

type QLineEdit interface {
	QWidget
	Alignment() AlignmentFlag
	Backspace()
	CreateStandardContextMenu() QMenu
	CursorBackward_Bool_Int(mark bool, steps int)
	CursorForward_Bool_Int(mark bool, steps int)
	CursorMoveStyle() CursorMoveStyle
	CursorPosition() int
	CursorWordBackward_Bool(mark bool)
	CursorWordForward_Bool(mark bool)
	Del()
	Deselect()
	DisplayText() string
	DragEnabled() bool
	EchoMode() EchoMode
	End_Bool(mark bool)
	HasAcceptableInput() bool
	HasFrame() bool
	HasSelectedText() bool
	Home_Bool(mark bool)
	InputMask() string
	Insert_String(newText string)
	IsClearButtonEnabled() bool
	IsModified() bool
	IsReadOnly() bool
	IsRedoAvailable() bool
	IsUndoAvailable() bool
	MaxLength() int
	PlaceholderText() string
	SelectedText() string
	SelectionStart() int
	SetAlignment_AlignmentFlag(flag AlignmentFlag)
	SetClearButtonEnabled_Bool(enable bool)
	SetCursorMoveStyle_CursorMoveStyle(style CursorMoveStyle)
	SetCursorPosition_Int(cursorPosition int)
	SetDragEnabled_Bool(dragEnabled bool)
	SetEchoMode_EchoMode(Ec EchoMode)
	SetFrame_Bool(frame bool)
	SetInputMask_String(inputMask string)
	SetMaxLength_Int(maxLength int)
	SetModified_Bool(modified bool)
	SetPlaceholderText_String(placeholderText string)
	SetReadOnly_Bool(readOnly bool)
	SetSelection_Int_Int(start int, length int)
	SetTextMargins_Int_Int_Int_Int(left int, top int, right int, bottom int)
	SetValidator_QValidator(v QValidator)
	Text() string
	ConnectSlotClear()
	DisconnectSlotClear()
	SlotClear()
	ConnectSlotCopy()
	DisconnectSlotCopy()
	SlotCopy()
	ConnectSlotCut()
	DisconnectSlotCut()
	SlotCut()
	ConnectSlotPaste()
	DisconnectSlotPaste()
	SlotPaste()
	ConnectSlotRedo()
	DisconnectSlotRedo()
	SlotRedo()
	ConnectSlotSelectAll()
	DisconnectSlotSelectAll()
	SlotSelectAll()
	ConnectSlotSetText()
	DisconnectSlotSetText()
	SlotSetText_String(text string)
	ConnectSlotUndo()
	DisconnectSlotUndo()
	SlotUndo()
	ConnectSignalCursorPositionChanged(f func())
	DisconnectSignalCursorPositionChanged()
	SignalCursorPositionChanged() func()
	ConnectSignalEditingFinished(f func())
	DisconnectSignalEditingFinished()
	SignalEditingFinished() func()
	ConnectSignalReturnPressed(f func())
	DisconnectSignalReturnPressed()
	SignalReturnPressed() func()
	ConnectSignalSelectionChanged(f func())
	DisconnectSignalSelectionChanged()
	SignalSelectionChanged() func()
	ConnectSignalTextChanged(f func())
	DisconnectSignalTextChanged()
	SignalTextChanged() func()
	ConnectSignalTextEdited(f func())
	DisconnectSignalTextEdited()
	SignalTextEdited() func()
}

func (p *qlineedit) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qlineedit) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

//ActionPosition
type ActionPosition int

var (
	LEADINGPOSITION  = ActionPosition(C.QLineEdit_LeadingPosition())
	TRAILINGPOSITION = ActionPosition(C.QLineEdit_TrailingPosition())
)

//EchoMode
type EchoMode int

var (
	NORMAL             = EchoMode(C.QLineEdit_Normal())
	NOECHO             = EchoMode(C.QLineEdit_NoEcho())
	PASSWORD           = EchoMode(C.QLineEdit_Password())
	PASSWORDECHOONEDIT = EchoMode(C.QLineEdit_PasswordEchoOnEdit())
)

func NewQLineEdit_QWidget(parent QWidget) QLineEdit {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qlineedit = new(qlineedit)
	qlineedit.SetPointer(C.QLineEdit_New_QWidget(parentPtr))
	qlineedit.SetObjectName_String("QLineEdit_" + randomIdentifier())
	return qlineedit
}

func NewQLineEdit_String_QWidget(contents string, parent QWidget) QLineEdit {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qlineedit = new(qlineedit)
	qlineedit.SetPointer(C.QLineEdit_New_String_QWidget(C.CString(contents), parentPtr))
	qlineedit.SetObjectName_String("QLineEdit_" + randomIdentifier())
	return qlineedit
}

func (p *qlineedit) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QLineEdit_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qlineedit) Alignment() AlignmentFlag {
	if p.Pointer() == nil {
		return 0
	} else {
		return AlignmentFlag(C.QLineEdit_Alignment(p.Pointer()))
	}
}

func (p *qlineedit) Backspace() {
	if p.Pointer() != nil {
		C.QLineEdit_Backspace(p.Pointer())
	}
}

func (p *qlineedit) CreateStandardContextMenu() QMenu {
	if p.Pointer() == nil {
		return nil
	} else {
		var qmenu = new(qmenu)
		qmenu.SetPointer(C.QLineEdit_CreateStandardContextMenu(p.Pointer()))
		if qmenu.ObjectName() == "" {
			qmenu.SetObjectName_String("QMenu_" + randomIdentifier())
		}
		return qmenu
	}
}

func (p *qlineedit) CursorBackward_Bool_Int(mark bool, steps int) {
	if p.Pointer() != nil {
		C.QLineEdit_CursorBackward_Bool_Int(p.Pointer(), goBoolToCInt(mark), C.int(steps))
	}
}

func (p *qlineedit) CursorForward_Bool_Int(mark bool, steps int) {
	if p.Pointer() != nil {
		C.QLineEdit_CursorForward_Bool_Int(p.Pointer(), goBoolToCInt(mark), C.int(steps))
	}
}

func (p *qlineedit) CursorMoveStyle() CursorMoveStyle {
	if p.Pointer() == nil {
		return 0
	} else {
		return CursorMoveStyle(C.QLineEdit_CursorMoveStyle(p.Pointer()))
	}
}

func (p *qlineedit) CursorPosition() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QLineEdit_CursorPosition(p.Pointer()))
	}
}

func (p *qlineedit) CursorWordBackward_Bool(mark bool) {
	if p.Pointer() != nil {
		C.QLineEdit_CursorWordBackward_Bool(p.Pointer(), goBoolToCInt(mark))
	}
}

func (p *qlineedit) CursorWordForward_Bool(mark bool) {
	if p.Pointer() != nil {
		C.QLineEdit_CursorWordForward_Bool(p.Pointer(), goBoolToCInt(mark))
	}
}

func (p *qlineedit) Del() {
	if p.Pointer() != nil {
		C.QLineEdit_Del(p.Pointer())
	}
}

func (p *qlineedit) Deselect() {
	if p.Pointer() != nil {
		C.QLineEdit_Deselect(p.Pointer())
	}
}

func (p *qlineedit) DisplayText() string {
	if p.Pointer() == nil {
		return ""
	} else {
		return C.GoString(C.QLineEdit_DisplayText(p.Pointer()))
	}
}

func (p *qlineedit) DragEnabled() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QLineEdit_DragEnabled(p.Pointer()) != 0
	}
}

func (p *qlineedit) EchoMode() EchoMode {
	if p.Pointer() == nil {
		return 0
	} else {
		return EchoMode(C.QLineEdit_EchoMode(p.Pointer()))
	}
}

func (p *qlineedit) End_Bool(mark bool) {
	if p.Pointer() != nil {
		C.QLineEdit_End_Bool(p.Pointer(), goBoolToCInt(mark))
	}
}

func (p *qlineedit) HasAcceptableInput() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QLineEdit_HasAcceptableInput(p.Pointer()) != 0
	}
}

func (p *qlineedit) HasFrame() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QLineEdit_HasFrame(p.Pointer()) != 0
	}
}

func (p *qlineedit) HasSelectedText() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QLineEdit_HasSelectedText(p.Pointer()) != 0
	}
}

func (p *qlineedit) Home_Bool(mark bool) {
	if p.Pointer() != nil {
		C.QLineEdit_Home_Bool(p.Pointer(), goBoolToCInt(mark))
	}
}

func (p *qlineedit) InputMask() string {
	if p.Pointer() == nil {
		return ""
	} else {
		return C.GoString(C.QLineEdit_InputMask(p.Pointer()))
	}
}

func (p *qlineedit) Insert_String(newText string) {
	if p.Pointer() != nil {
		C.QLineEdit_Insert_String(p.Pointer(), C.CString(newText))
	}
}

func (p *qlineedit) IsClearButtonEnabled() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QLineEdit_IsClearButtonEnabled(p.Pointer()) != 0
	}
}

func (p *qlineedit) IsModified() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QLineEdit_IsModified(p.Pointer()) != 0
	}
}

func (p *qlineedit) IsReadOnly() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QLineEdit_IsReadOnly(p.Pointer()) != 0
	}
}

func (p *qlineedit) IsRedoAvailable() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QLineEdit_IsRedoAvailable(p.Pointer()) != 0
	}
}

func (p *qlineedit) IsUndoAvailable() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QLineEdit_IsUndoAvailable(p.Pointer()) != 0
	}
}

func (p *qlineedit) MaxLength() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QLineEdit_MaxLength(p.Pointer()))
	}
}

func (p *qlineedit) PlaceholderText() string {
	if p.Pointer() == nil {
		return ""
	} else {
		return C.GoString(C.QLineEdit_PlaceholderText(p.Pointer()))
	}
}

func (p *qlineedit) SelectedText() string {
	if p.Pointer() == nil {
		return ""
	} else {
		return C.GoString(C.QLineEdit_SelectedText(p.Pointer()))
	}
}

func (p *qlineedit) SelectionStart() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QLineEdit_SelectionStart(p.Pointer()))
	}
}

func (p *qlineedit) SetAlignment_AlignmentFlag(flag AlignmentFlag) {
	if p.Pointer() != nil {
		C.QLineEdit_SetAlignment_AlignmentFlag(p.Pointer(), C.int(flag))
	}
}

func (p *qlineedit) SetClearButtonEnabled_Bool(enable bool) {
	if p.Pointer() != nil {
		C.QLineEdit_SetClearButtonEnabled_Bool(p.Pointer(), goBoolToCInt(enable))
	}
}

func (p *qlineedit) SetCursorMoveStyle_CursorMoveStyle(style CursorMoveStyle) {
	if p.Pointer() != nil {
		C.QLineEdit_SetCursorMoveStyle_CursorMoveStyle(p.Pointer(), C.int(style))
	}
}

func (p *qlineedit) SetCursorPosition_Int(cursorPosition int) {
	if p.Pointer() != nil {
		C.QLineEdit_SetCursorPosition_Int(p.Pointer(), C.int(cursorPosition))
	}
}

func (p *qlineedit) SetDragEnabled_Bool(dragEnabled bool) {
	if p.Pointer() != nil {
		C.QLineEdit_SetDragEnabled_Bool(p.Pointer(), goBoolToCInt(dragEnabled))
	}
}

func (p *qlineedit) SetEchoMode_EchoMode(Ec EchoMode) {
	if p.Pointer() != nil {
		C.QLineEdit_SetEchoMode_EchoMode(p.Pointer(), C.int(Ec))
	}
}

func (p *qlineedit) SetFrame_Bool(frame bool) {
	if p.Pointer() != nil {
		C.QLineEdit_SetFrame_Bool(p.Pointer(), goBoolToCInt(frame))
	}
}

func (p *qlineedit) SetInputMask_String(inputMask string) {
	if p.Pointer() != nil {
		C.QLineEdit_SetInputMask_String(p.Pointer(), C.CString(inputMask))
	}
}

func (p *qlineedit) SetMaxLength_Int(maxLength int) {
	if p.Pointer() != nil {
		C.QLineEdit_SetMaxLength_Int(p.Pointer(), C.int(maxLength))
	}
}

func (p *qlineedit) SetModified_Bool(modified bool) {
	if p.Pointer() != nil {
		C.QLineEdit_SetModified_Bool(p.Pointer(), goBoolToCInt(modified))
	}
}

func (p *qlineedit) SetPlaceholderText_String(placeholderText string) {
	if p.Pointer() != nil {
		C.QLineEdit_SetPlaceholderText_String(p.Pointer(), C.CString(placeholderText))
	}
}

func (p *qlineedit) SetReadOnly_Bool(readOnly bool) {
	if p.Pointer() != nil {
		C.QLineEdit_SetReadOnly_Bool(p.Pointer(), goBoolToCInt(readOnly))
	}
}

func (p *qlineedit) SetSelection_Int_Int(start int, length int) {
	if p.Pointer() != nil {
		C.QLineEdit_SetSelection_Int_Int(p.Pointer(), C.int(start), C.int(length))
	}
}

func (p *qlineedit) SetTextMargins_Int_Int_Int_Int(left int, top int, right int, bottom int) {
	if p.Pointer() != nil {
		C.QLineEdit_SetTextMargins_Int_Int_Int_Int(p.Pointer(), C.int(left), C.int(top), C.int(right), C.int(bottom))
	}
}

func (p *qlineedit) SetValidator_QValidator(v QValidator) {
	if p.Pointer() == nil {
	} else {
		var vPtr C.QtObjectPtr = nil
		if v != nil {
			vPtr = v.Pointer()
		}
		C.QLineEdit_SetValidator_QValidator(p.Pointer(), vPtr)
	}
}

func (p *qlineedit) Text() string {
	if p.Pointer() == nil {
		return ""
	} else {
		return C.GoString(C.QLineEdit_Text(p.Pointer()))
	}
}

func (p *qlineedit) ConnectSlotClear() {
	C.QLineEdit_ConnectSlotClear(p.Pointer())
}

func (p *qlineedit) DisconnectSlotClear() {
	C.QLineEdit_DisconnectSlotClear(p.Pointer())
}

func (p *qlineedit) SlotClear() {
	if p.Pointer() != nil {
		C.QLineEdit_Clear(p.Pointer())
	}
}

func (p *qlineedit) ConnectSlotCopy() {
	C.QLineEdit_ConnectSlotCopy(p.Pointer())
}

func (p *qlineedit) DisconnectSlotCopy() {
	C.QLineEdit_DisconnectSlotCopy(p.Pointer())
}

func (p *qlineedit) SlotCopy() {
	if p.Pointer() != nil {
		C.QLineEdit_Copy(p.Pointer())
	}
}

func (p *qlineedit) ConnectSlotCut() {
	C.QLineEdit_ConnectSlotCut(p.Pointer())
}

func (p *qlineedit) DisconnectSlotCut() {
	C.QLineEdit_DisconnectSlotCut(p.Pointer())
}

func (p *qlineedit) SlotCut() {
	if p.Pointer() != nil {
		C.QLineEdit_Cut(p.Pointer())
	}
}

func (p *qlineedit) ConnectSlotPaste() {
	C.QLineEdit_ConnectSlotPaste(p.Pointer())
}

func (p *qlineedit) DisconnectSlotPaste() {
	C.QLineEdit_DisconnectSlotPaste(p.Pointer())
}

func (p *qlineedit) SlotPaste() {
	if p.Pointer() != nil {
		C.QLineEdit_Paste(p.Pointer())
	}
}

func (p *qlineedit) ConnectSlotRedo() {
	C.QLineEdit_ConnectSlotRedo(p.Pointer())
}

func (p *qlineedit) DisconnectSlotRedo() {
	C.QLineEdit_DisconnectSlotRedo(p.Pointer())
}

func (p *qlineedit) SlotRedo() {
	if p.Pointer() != nil {
		C.QLineEdit_Redo(p.Pointer())
	}
}

func (p *qlineedit) ConnectSlotSelectAll() {
	C.QLineEdit_ConnectSlotSelectAll(p.Pointer())
}

func (p *qlineedit) DisconnectSlotSelectAll() {
	C.QLineEdit_DisconnectSlotSelectAll(p.Pointer())
}

func (p *qlineedit) SlotSelectAll() {
	if p.Pointer() != nil {
		C.QLineEdit_SelectAll(p.Pointer())
	}
}

func (p *qlineedit) ConnectSlotSetText() {
	C.QLineEdit_ConnectSlotSetText(p.Pointer())
}

func (p *qlineedit) DisconnectSlotSetText() {
	C.QLineEdit_DisconnectSlotSetText(p.Pointer())
}

func (p *qlineedit) SlotSetText_String(text string) {
	if p.Pointer() != nil {
		C.QLineEdit_SetText_String(p.Pointer(), C.CString(text))
	}
}

func (p *qlineedit) ConnectSlotUndo() {
	C.QLineEdit_ConnectSlotUndo(p.Pointer())
}

func (p *qlineedit) DisconnectSlotUndo() {
	C.QLineEdit_DisconnectSlotUndo(p.Pointer())
}

func (p *qlineedit) SlotUndo() {
	if p.Pointer() != nil {
		C.QLineEdit_Undo(p.Pointer())
	}
}

func (p *qlineedit) ConnectSignalCursorPositionChanged(f func()) {
	C.QLineEdit_ConnectSignalCursorPositionChanged(p.Pointer())
	connectSignal(p.ObjectName(), "cursorPositionChanged", f)
}

func (p *qlineedit) DisconnectSignalCursorPositionChanged() {
	C.QLineEdit_DisconnectSignalCursorPositionChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "cursorPositionChanged")
}

func (p *qlineedit) SignalCursorPositionChanged() func() {
	return getSignal(p.ObjectName(), "cursorPositionChanged")
}

func (p *qlineedit) ConnectSignalEditingFinished(f func()) {
	C.QLineEdit_ConnectSignalEditingFinished(p.Pointer())
	connectSignal(p.ObjectName(), "editingFinished", f)
}

func (p *qlineedit) DisconnectSignalEditingFinished() {
	C.QLineEdit_DisconnectSignalEditingFinished(p.Pointer())
	disconnectSignal(p.ObjectName(), "editingFinished")
}

func (p *qlineedit) SignalEditingFinished() func() {
	return getSignal(p.ObjectName(), "editingFinished")
}

func (p *qlineedit) ConnectSignalReturnPressed(f func()) {
	C.QLineEdit_ConnectSignalReturnPressed(p.Pointer())
	connectSignal(p.ObjectName(), "returnPressed", f)
}

func (p *qlineedit) DisconnectSignalReturnPressed() {
	C.QLineEdit_DisconnectSignalReturnPressed(p.Pointer())
	disconnectSignal(p.ObjectName(), "returnPressed")
}

func (p *qlineedit) SignalReturnPressed() func() {
	return getSignal(p.ObjectName(), "returnPressed")
}

func (p *qlineedit) ConnectSignalSelectionChanged(f func()) {
	C.QLineEdit_ConnectSignalSelectionChanged(p.Pointer())
	connectSignal(p.ObjectName(), "selectionChanged", f)
}

func (p *qlineedit) DisconnectSignalSelectionChanged() {
	C.QLineEdit_DisconnectSignalSelectionChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "selectionChanged")
}

func (p *qlineedit) SignalSelectionChanged() func() {
	return getSignal(p.ObjectName(), "selectionChanged")
}

func (p *qlineedit) ConnectSignalTextChanged(f func()) {
	C.QLineEdit_ConnectSignalTextChanged(p.Pointer())
	connectSignal(p.ObjectName(), "textChanged", f)
}

func (p *qlineedit) DisconnectSignalTextChanged() {
	C.QLineEdit_DisconnectSignalTextChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "textChanged")
}

func (p *qlineedit) SignalTextChanged() func() {
	return getSignal(p.ObjectName(), "textChanged")
}

func (p *qlineedit) ConnectSignalTextEdited(f func()) {
	C.QLineEdit_ConnectSignalTextEdited(p.Pointer())
	connectSignal(p.ObjectName(), "textEdited", f)
}

func (p *qlineedit) DisconnectSignalTextEdited() {
	C.QLineEdit_DisconnectSignalTextEdited(p.Pointer())
	disconnectSignal(p.ObjectName(), "textEdited")
}

func (p *qlineedit) SignalTextEdited() func() {
	return getSignal(p.ObjectName(), "textEdited")
}

//export callbackQLineEdit
func callbackQLineEdit(ptr C.QtObjectPtr, msg *C.char) {
	var qlineedit = new(qlineedit)
	qlineedit.SetPointer(ptr)
	getSignal(qlineedit.ObjectName(), C.GoString(msg))()
}
