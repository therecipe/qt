package qt

//#include "qtextedit.h"
import "C"

type qtextedit struct {
	qabstractscrollarea
}

type QTextEdit interface {
	QAbstractScrollArea
	AcceptRichText() bool
	Alignment() AlignmentFlag
	CanPaste() bool
	CreateStandardContextMenu() QMenu
	CursorWidth() int
	DocumentTitle() string
	EnsureCursorVisible()
	FontFamily() string
	FontItalic() bool
	FontUnderline() bool
	FontWeight() int
	IsReadOnly() bool
	IsUndoRedoEnabled() bool
	LineWrapColumnOrWidth() int
	OverwriteMode() bool
	PlaceholderText() string
	SetAcceptRichText_Bool(accept bool)
	SetCursorWidth_Int(width int)
	SetDocumentTitle_String(title string)
	SetLineWrapColumnOrWidth_Int(w int)
	SetOverwriteMode_Bool(overwrite bool)
	SetPlaceholderText_String(placeholderText string)
	SetReadOnly_Bool(ro bool)
	SetTabChangesFocus_Bool(b bool)
	SetTabStopWidth_Int(width int)
	SetTextInteractionFlags_TextInteractionFlag(flags TextInteractionFlag)
	SetUndoRedoEnabled_Bool(enable bool)
	TabChangesFocus() bool
	TabStopWidth() int
	TextInteractionFlags() TextInteractionFlag
	ToHtml() string
	ToPlainText() string
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
	ConnectSlotSetFontItalic()
	DisconnectSlotSetFontItalic()
	SlotSetFontItalic_Bool(italic bool)
	ConnectSlotSetFontUnderline()
	DisconnectSlotSetFontUnderline()
	SlotSetFontUnderline_Bool(underline bool)
	ConnectSlotSetFontWeight()
	DisconnectSlotSetFontWeight()
	SlotSetFontWeight_Int(weight int)
	ConnectSlotUndo()
	DisconnectSlotUndo()
	SlotUndo()
	ConnectSlotZoomIn()
	DisconnectSlotZoomIn()
	SlotZoomIn_Int(rang int)
	ConnectSlotZoomOut()
	DisconnectSlotZoomOut()
	SlotZoomOut_Int(rang int)
	ConnectSignalCopyAvailable(f func())
	DisconnectSignalCopyAvailable()
	SignalCopyAvailable() func()
	ConnectSignalCursorPositionChanged(f func())
	DisconnectSignalCursorPositionChanged()
	SignalCursorPositionChanged() func()
	ConnectSignalRedoAvailable(f func())
	DisconnectSignalRedoAvailable()
	SignalRedoAvailable() func()
	ConnectSignalSelectionChanged(f func())
	DisconnectSignalSelectionChanged()
	SignalSelectionChanged() func()
	ConnectSignalTextChanged(f func())
	DisconnectSignalTextChanged()
	SignalTextChanged() func()
	ConnectSignalUndoAvailable(f func())
	DisconnectSignalUndoAvailable()
	SignalUndoAvailable() func()
}

func (p *qtextedit) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qtextedit) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQTextEdit_QWidget(parent QWidget) QTextEdit {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qtextedit = new(qtextedit)
	qtextedit.SetPointer(C.QTextEdit_New_QWidget(parentPtr))
	qtextedit.SetObjectName_String("QTextEdit_" + randomIdentifier())
	return qtextedit
}

func NewQTextEdit_String_QWidget(text string, parent QWidget) QTextEdit {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qtextedit = new(qtextedit)
	qtextedit.SetPointer(C.QTextEdit_New_String_QWidget(C.CString(text), parentPtr))
	qtextedit.SetObjectName_String("QTextEdit_" + randomIdentifier())
	return qtextedit
}

func (p *qtextedit) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QTextEdit_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qtextedit) AcceptRichText() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTextEdit_AcceptRichText(p.Pointer()) != 0
	}
}

func (p *qtextedit) Alignment() AlignmentFlag {
	if p.Pointer() == nil {
		return 0
	} else {
		return AlignmentFlag(C.QTextEdit_Alignment(p.Pointer()))
	}
}

func (p *qtextedit) CanPaste() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTextEdit_CanPaste(p.Pointer()) != 0
	}
}

func (p *qtextedit) CreateStandardContextMenu() QMenu {
	if p.Pointer() == nil {
		return nil
	} else {
		var qmenu = new(qmenu)
		qmenu.SetPointer(C.QTextEdit_CreateStandardContextMenu(p.Pointer()))
		if qmenu.ObjectName() == "" {
			qmenu.SetObjectName_String("QMenu_" + randomIdentifier())
		}
		return qmenu
	}
}

func (p *qtextedit) CursorWidth() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QTextEdit_CursorWidth(p.Pointer()))
	}
}

func (p *qtextedit) DocumentTitle() string {
	if p.Pointer() == nil {
		return ""
	} else {
		return C.GoString(C.QTextEdit_DocumentTitle(p.Pointer()))
	}
}

func (p *qtextedit) EnsureCursorVisible() {
	if p.Pointer() != nil {
		C.QTextEdit_EnsureCursorVisible(p.Pointer())
	}
}

func (p *qtextedit) FontFamily() string {
	if p.Pointer() == nil {
		return ""
	} else {
		return C.GoString(C.QTextEdit_FontFamily(p.Pointer()))
	}
}

func (p *qtextedit) FontItalic() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTextEdit_FontItalic(p.Pointer()) != 0
	}
}

func (p *qtextedit) FontUnderline() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTextEdit_FontUnderline(p.Pointer()) != 0
	}
}

func (p *qtextedit) FontWeight() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QTextEdit_FontWeight(p.Pointer()))
	}
}

func (p *qtextedit) IsReadOnly() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTextEdit_IsReadOnly(p.Pointer()) != 0
	}
}

func (p *qtextedit) IsUndoRedoEnabled() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTextEdit_IsUndoRedoEnabled(p.Pointer()) != 0
	}
}

func (p *qtextedit) LineWrapColumnOrWidth() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QTextEdit_LineWrapColumnOrWidth(p.Pointer()))
	}
}

func (p *qtextedit) OverwriteMode() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTextEdit_OverwriteMode(p.Pointer()) != 0
	}
}

func (p *qtextedit) PlaceholderText() string {
	if p.Pointer() == nil {
		return ""
	} else {
		return C.GoString(C.QTextEdit_PlaceholderText(p.Pointer()))
	}
}

func (p *qtextedit) SetAcceptRichText_Bool(accept bool) {
	if p.Pointer() != nil {
		C.QTextEdit_SetAcceptRichText_Bool(p.Pointer(), goBoolToCInt(accept))
	}
}

func (p *qtextedit) SetCursorWidth_Int(width int) {
	if p.Pointer() != nil {
		C.QTextEdit_SetCursorWidth_Int(p.Pointer(), C.int(width))
	}
}

func (p *qtextedit) SetDocumentTitle_String(title string) {
	if p.Pointer() != nil {
		C.QTextEdit_SetDocumentTitle_String(p.Pointer(), C.CString(title))
	}
}

func (p *qtextedit) SetLineWrapColumnOrWidth_Int(w int) {
	if p.Pointer() != nil {
		C.QTextEdit_SetLineWrapColumnOrWidth_Int(p.Pointer(), C.int(w))
	}
}

func (p *qtextedit) SetOverwriteMode_Bool(overwrite bool) {
	if p.Pointer() != nil {
		C.QTextEdit_SetOverwriteMode_Bool(p.Pointer(), goBoolToCInt(overwrite))
	}
}

func (p *qtextedit) SetPlaceholderText_String(placeholderText string) {
	if p.Pointer() != nil {
		C.QTextEdit_SetPlaceholderText_String(p.Pointer(), C.CString(placeholderText))
	}
}

func (p *qtextedit) SetReadOnly_Bool(ro bool) {
	if p.Pointer() != nil {
		C.QTextEdit_SetReadOnly_Bool(p.Pointer(), goBoolToCInt(ro))
	}
}

func (p *qtextedit) SetTabChangesFocus_Bool(b bool) {
	if p.Pointer() != nil {
		C.QTextEdit_SetTabChangesFocus_Bool(p.Pointer(), goBoolToCInt(b))
	}
}

func (p *qtextedit) SetTabStopWidth_Int(width int) {
	if p.Pointer() != nil {
		C.QTextEdit_SetTabStopWidth_Int(p.Pointer(), C.int(width))
	}
}

func (p *qtextedit) SetTextInteractionFlags_TextInteractionFlag(flags TextInteractionFlag) {
	if p.Pointer() != nil {
		C.QTextEdit_SetTextInteractionFlags_TextInteractionFlag(p.Pointer(), C.int(flags))
	}
}

func (p *qtextedit) SetUndoRedoEnabled_Bool(enable bool) {
	if p.Pointer() != nil {
		C.QTextEdit_SetUndoRedoEnabled_Bool(p.Pointer(), goBoolToCInt(enable))
	}
}

func (p *qtextedit) TabChangesFocus() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTextEdit_TabChangesFocus(p.Pointer()) != 0
	}
}

func (p *qtextedit) TabStopWidth() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QTextEdit_TabStopWidth(p.Pointer()))
	}
}

func (p *qtextedit) TextInteractionFlags() TextInteractionFlag {
	if p.Pointer() == nil {
		return 0
	} else {
		return TextInteractionFlag(C.QTextEdit_TextInteractionFlags(p.Pointer()))
	}
}

func (p *qtextedit) ToHtml() string {
	if p.Pointer() == nil {
		return ""
	} else {
		return C.GoString(C.QTextEdit_ToHtml(p.Pointer()))
	}
}

func (p *qtextedit) ToPlainText() string {
	if p.Pointer() == nil {
		return ""
	} else {
		return C.GoString(C.QTextEdit_ToPlainText(p.Pointer()))
	}
}

func (p *qtextedit) ConnectSlotClear() {
	C.QTextEdit_ConnectSlotClear(p.Pointer())
}

func (p *qtextedit) DisconnectSlotClear() {
	C.QTextEdit_DisconnectSlotClear(p.Pointer())
}

func (p *qtextedit) SlotClear() {
	if p.Pointer() != nil {
		C.QTextEdit_Clear(p.Pointer())
	}
}

func (p *qtextedit) ConnectSlotCopy() {
	C.QTextEdit_ConnectSlotCopy(p.Pointer())
}

func (p *qtextedit) DisconnectSlotCopy() {
	C.QTextEdit_DisconnectSlotCopy(p.Pointer())
}

func (p *qtextedit) SlotCopy() {
	if p.Pointer() != nil {
		C.QTextEdit_Copy(p.Pointer())
	}
}

func (p *qtextedit) ConnectSlotCut() {
	C.QTextEdit_ConnectSlotCut(p.Pointer())
}

func (p *qtextedit) DisconnectSlotCut() {
	C.QTextEdit_DisconnectSlotCut(p.Pointer())
}

func (p *qtextedit) SlotCut() {
	if p.Pointer() != nil {
		C.QTextEdit_Cut(p.Pointer())
	}
}

func (p *qtextedit) ConnectSlotPaste() {
	C.QTextEdit_ConnectSlotPaste(p.Pointer())
}

func (p *qtextedit) DisconnectSlotPaste() {
	C.QTextEdit_DisconnectSlotPaste(p.Pointer())
}

func (p *qtextedit) SlotPaste() {
	if p.Pointer() != nil {
		C.QTextEdit_Paste(p.Pointer())
	}
}

func (p *qtextedit) ConnectSlotRedo() {
	C.QTextEdit_ConnectSlotRedo(p.Pointer())
}

func (p *qtextedit) DisconnectSlotRedo() {
	C.QTextEdit_DisconnectSlotRedo(p.Pointer())
}

func (p *qtextedit) SlotRedo() {
	if p.Pointer() != nil {
		C.QTextEdit_Redo(p.Pointer())
	}
}

func (p *qtextedit) ConnectSlotSelectAll() {
	C.QTextEdit_ConnectSlotSelectAll(p.Pointer())
}

func (p *qtextedit) DisconnectSlotSelectAll() {
	C.QTextEdit_DisconnectSlotSelectAll(p.Pointer())
}

func (p *qtextedit) SlotSelectAll() {
	if p.Pointer() != nil {
		C.QTextEdit_SelectAll(p.Pointer())
	}
}

func (p *qtextedit) ConnectSlotSetFontItalic() {
	C.QTextEdit_ConnectSlotSetFontItalic(p.Pointer())
}

func (p *qtextedit) DisconnectSlotSetFontItalic() {
	C.QTextEdit_DisconnectSlotSetFontItalic(p.Pointer())
}

func (p *qtextedit) SlotSetFontItalic_Bool(italic bool) {
	if p.Pointer() != nil {
		C.QTextEdit_SetFontItalic_Bool(p.Pointer(), goBoolToCInt(italic))
	}
}

func (p *qtextedit) ConnectSlotSetFontUnderline() {
	C.QTextEdit_ConnectSlotSetFontUnderline(p.Pointer())
}

func (p *qtextedit) DisconnectSlotSetFontUnderline() {
	C.QTextEdit_DisconnectSlotSetFontUnderline(p.Pointer())
}

func (p *qtextedit) SlotSetFontUnderline_Bool(underline bool) {
	if p.Pointer() != nil {
		C.QTextEdit_SetFontUnderline_Bool(p.Pointer(), goBoolToCInt(underline))
	}
}

func (p *qtextedit) ConnectSlotSetFontWeight() {
	C.QTextEdit_ConnectSlotSetFontWeight(p.Pointer())
}

func (p *qtextedit) DisconnectSlotSetFontWeight() {
	C.QTextEdit_DisconnectSlotSetFontWeight(p.Pointer())
}

func (p *qtextedit) SlotSetFontWeight_Int(weight int) {
	if p.Pointer() != nil {
		C.QTextEdit_SetFontWeight_Int(p.Pointer(), C.int(weight))
	}
}

func (p *qtextedit) ConnectSlotUndo() {
	C.QTextEdit_ConnectSlotUndo(p.Pointer())
}

func (p *qtextedit) DisconnectSlotUndo() {
	C.QTextEdit_DisconnectSlotUndo(p.Pointer())
}

func (p *qtextedit) SlotUndo() {
	if p.Pointer() != nil {
		C.QTextEdit_Undo(p.Pointer())
	}
}

func (p *qtextedit) ConnectSlotZoomIn() {
	C.QTextEdit_ConnectSlotZoomIn(p.Pointer())
}

func (p *qtextedit) DisconnectSlotZoomIn() {
	C.QTextEdit_DisconnectSlotZoomIn(p.Pointer())
}

func (p *qtextedit) SlotZoomIn_Int(rang int) {
	if p.Pointer() != nil {
		C.QTextEdit_ZoomIn_Int(p.Pointer(), C.int(rang))
	}
}

func (p *qtextedit) ConnectSlotZoomOut() {
	C.QTextEdit_ConnectSlotZoomOut(p.Pointer())
}

func (p *qtextedit) DisconnectSlotZoomOut() {
	C.QTextEdit_DisconnectSlotZoomOut(p.Pointer())
}

func (p *qtextedit) SlotZoomOut_Int(rang int) {
	if p.Pointer() != nil {
		C.QTextEdit_ZoomOut_Int(p.Pointer(), C.int(rang))
	}
}

func (p *qtextedit) ConnectSignalCopyAvailable(f func()) {
	C.QTextEdit_ConnectSignalCopyAvailable(p.Pointer())
	connectSignal(p.ObjectName(), "copyAvailable", f)
}

func (p *qtextedit) DisconnectSignalCopyAvailable() {
	C.QTextEdit_DisconnectSignalCopyAvailable(p.Pointer())
	disconnectSignal(p.ObjectName(), "copyAvailable")
}

func (p *qtextedit) SignalCopyAvailable() func() {
	return getSignal(p.ObjectName(), "copyAvailable")
}

func (p *qtextedit) ConnectSignalCursorPositionChanged(f func()) {
	C.QTextEdit_ConnectSignalCursorPositionChanged(p.Pointer())
	connectSignal(p.ObjectName(), "cursorPositionChanged", f)
}

func (p *qtextedit) DisconnectSignalCursorPositionChanged() {
	C.QTextEdit_DisconnectSignalCursorPositionChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "cursorPositionChanged")
}

func (p *qtextedit) SignalCursorPositionChanged() func() {
	return getSignal(p.ObjectName(), "cursorPositionChanged")
}

func (p *qtextedit) ConnectSignalRedoAvailable(f func()) {
	C.QTextEdit_ConnectSignalRedoAvailable(p.Pointer())
	connectSignal(p.ObjectName(), "redoAvailable", f)
}

func (p *qtextedit) DisconnectSignalRedoAvailable() {
	C.QTextEdit_DisconnectSignalRedoAvailable(p.Pointer())
	disconnectSignal(p.ObjectName(), "redoAvailable")
}

func (p *qtextedit) SignalRedoAvailable() func() {
	return getSignal(p.ObjectName(), "redoAvailable")
}

func (p *qtextedit) ConnectSignalSelectionChanged(f func()) {
	C.QTextEdit_ConnectSignalSelectionChanged(p.Pointer())
	connectSignal(p.ObjectName(), "selectionChanged", f)
}

func (p *qtextedit) DisconnectSignalSelectionChanged() {
	C.QTextEdit_DisconnectSignalSelectionChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "selectionChanged")
}

func (p *qtextedit) SignalSelectionChanged() func() {
	return getSignal(p.ObjectName(), "selectionChanged")
}

func (p *qtextedit) ConnectSignalTextChanged(f func()) {
	C.QTextEdit_ConnectSignalTextChanged(p.Pointer())
	connectSignal(p.ObjectName(), "textChanged", f)
}

func (p *qtextedit) DisconnectSignalTextChanged() {
	C.QTextEdit_DisconnectSignalTextChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "textChanged")
}

func (p *qtextedit) SignalTextChanged() func() {
	return getSignal(p.ObjectName(), "textChanged")
}

func (p *qtextedit) ConnectSignalUndoAvailable(f func()) {
	C.QTextEdit_ConnectSignalUndoAvailable(p.Pointer())
	connectSignal(p.ObjectName(), "undoAvailable", f)
}

func (p *qtextedit) DisconnectSignalUndoAvailable() {
	C.QTextEdit_DisconnectSignalUndoAvailable(p.Pointer())
	disconnectSignal(p.ObjectName(), "undoAvailable")
}

func (p *qtextedit) SignalUndoAvailable() func() {
	return getSignal(p.ObjectName(), "undoAvailable")
}

//export callbackQTextEdit
func callbackQTextEdit(ptr C.QtObjectPtr, msg *C.char) {
	var qtextedit = new(qtextedit)
	qtextedit.SetPointer(ptr)
	getSignal(qtextedit.ObjectName(), C.GoString(msg))()
}
