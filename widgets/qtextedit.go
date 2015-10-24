package widgets

//#include "qtextedit.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QTextEdit struct {
	QAbstractScrollArea
}

type QTextEditITF interface {
	QAbstractScrollAreaITF
	QTextEditPTR() *QTextEdit
}

func PointerFromQTextEdit(ptr QTextEditITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextEditPTR().Pointer()
	}
	return nil
}

func QTextEditFromPointer(ptr unsafe.Pointer) *QTextEdit {
	var n = new(QTextEdit)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTextEdit_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTextEdit) QTextEditPTR() *QTextEdit {
	return ptr
}

//QTextEdit::AutoFormattingFlag
type QTextEdit__AutoFormattingFlag int

var (
	QTextEdit__AutoNone       = QTextEdit__AutoFormattingFlag(0)
	QTextEdit__AutoBulletList = QTextEdit__AutoFormattingFlag(0x00000001)
	QTextEdit__AutoAll        = QTextEdit__AutoFormattingFlag(0xffffffff)
)

//QTextEdit::LineWrapMode
type QTextEdit__LineWrapMode int

var (
	QTextEdit__NoWrap           = QTextEdit__LineWrapMode(0)
	QTextEdit__WidgetWidth      = QTextEdit__LineWrapMode(1)
	QTextEdit__FixedPixelWidth  = QTextEdit__LineWrapMode(2)
	QTextEdit__FixedColumnWidth = QTextEdit__LineWrapMode(3)
)

func (ptr *QTextEdit) AcceptRichText() bool {
	if ptr.Pointer() != nil {
		return C.QTextEdit_AcceptRichText(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextEdit) AutoFormatting() QTextEdit__AutoFormattingFlag {
	if ptr.Pointer() != nil {
		return QTextEdit__AutoFormattingFlag(C.QTextEdit_AutoFormatting(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextEdit) CursorWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QTextEdit_CursorWidth(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextEdit) Document() *gui.QTextDocument {
	if ptr.Pointer() != nil {
		return gui.QTextDocumentFromPointer(unsafe.Pointer(C.QTextEdit_Document(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTextEdit) IsReadOnly() bool {
	if ptr.Pointer() != nil {
		return C.QTextEdit_IsReadOnly(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextEdit) LineWrapColumnOrWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QTextEdit_LineWrapColumnOrWidth(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextEdit) LineWrapMode() QTextEdit__LineWrapMode {
	if ptr.Pointer() != nil {
		return QTextEdit__LineWrapMode(C.QTextEdit_LineWrapMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextEdit) OverwriteMode() bool {
	if ptr.Pointer() != nil {
		return C.QTextEdit_OverwriteMode(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextEdit) PlaceholderText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextEdit_PlaceholderText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTextEdit) Redo() {
	if ptr.Pointer() != nil {
		C.QTextEdit_Redo(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextEdit) SetAcceptRichText(accept bool) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetAcceptRichText(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(accept)))
	}
}

func (ptr *QTextEdit) SetAutoFormatting(features QTextEdit__AutoFormattingFlag) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetAutoFormatting(C.QtObjectPtr(ptr.Pointer()), C.int(features))
	}
}

func (ptr *QTextEdit) SetCursorWidth(width int) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetCursorWidth(C.QtObjectPtr(ptr.Pointer()), C.int(width))
	}
}

func (ptr *QTextEdit) SetDocument(document gui.QTextDocumentITF) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetDocument(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQTextDocument(document)))
	}
}

func (ptr *QTextEdit) SetFontWeight(weight int) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetFontWeight(C.QtObjectPtr(ptr.Pointer()), C.int(weight))
	}
}

func (ptr *QTextEdit) SetHtml(text string) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetHtml(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QTextEdit) SetLineWrapColumnOrWidth(w int) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetLineWrapColumnOrWidth(C.QtObjectPtr(ptr.Pointer()), C.int(w))
	}
}

func (ptr *QTextEdit) SetLineWrapMode(mode QTextEdit__LineWrapMode) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetLineWrapMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QTextEdit) SetOverwriteMode(overwrite bool) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetOverwriteMode(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(overwrite)))
	}
}

func (ptr *QTextEdit) SetPlaceholderText(placeholderText string) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetPlaceholderText(C.QtObjectPtr(ptr.Pointer()), C.CString(placeholderText))
	}
}

func (ptr *QTextEdit) SetReadOnly(ro bool) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetReadOnly(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(ro)))
	}
}

func (ptr *QTextEdit) SetTabChangesFocus(b bool) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetTabChangesFocus(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QTextEdit) SetTabStopWidth(width int) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetTabStopWidth(C.QtObjectPtr(ptr.Pointer()), C.int(width))
	}
}

func (ptr *QTextEdit) SetTextInteractionFlags(flags core.Qt__TextInteractionFlag) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetTextInteractionFlags(C.QtObjectPtr(ptr.Pointer()), C.int(flags))
	}
}

func (ptr *QTextEdit) SetWordWrapMode(policy gui.QTextOption__WrapMode) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetWordWrapMode(C.QtObjectPtr(ptr.Pointer()), C.int(policy))
	}
}

func (ptr *QTextEdit) TabChangesFocus() bool {
	if ptr.Pointer() != nil {
		return C.QTextEdit_TabChangesFocus(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextEdit) TabStopWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QTextEdit_TabStopWidth(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextEdit) TextInteractionFlags() core.Qt__TextInteractionFlag {
	if ptr.Pointer() != nil {
		return core.Qt__TextInteractionFlag(C.QTextEdit_TextInteractionFlags(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextEdit) ToHtml() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextEdit_ToHtml(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTextEdit) WordWrapMode() gui.QTextOption__WrapMode {
	if ptr.Pointer() != nil {
		return gui.QTextOption__WrapMode(C.QTextEdit_WordWrapMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextEdit) ZoomIn(ran int) {
	if ptr.Pointer() != nil {
		C.QTextEdit_ZoomIn(C.QtObjectPtr(ptr.Pointer()), C.int(ran))
	}
}

func (ptr *QTextEdit) ZoomOut(ran int) {
	if ptr.Pointer() != nil {
		C.QTextEdit_ZoomOut(C.QtObjectPtr(ptr.Pointer()), C.int(ran))
	}
}

func NewQTextEdit(parent QWidgetITF) *QTextEdit {
	return QTextEditFromPointer(unsafe.Pointer(C.QTextEdit_NewQTextEdit(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func NewQTextEdit2(text string, parent QWidgetITF) *QTextEdit {
	return QTextEditFromPointer(unsafe.Pointer(C.QTextEdit_NewQTextEdit2(C.CString(text), C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QTextEdit) Alignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QTextEdit_Alignment(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextEdit) AnchorAt(pos core.QPointITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextEdit_AnchorAt(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(pos))))
	}
	return ""
}

func (ptr *QTextEdit) Append(text string) {
	if ptr.Pointer() != nil {
		C.QTextEdit_Append(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QTextEdit) CanPaste() bool {
	if ptr.Pointer() != nil {
		return C.QTextEdit_CanPaste(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextEdit) Clear() {
	if ptr.Pointer() != nil {
		C.QTextEdit_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextEdit) Copy() {
	if ptr.Pointer() != nil {
		C.QTextEdit_Copy(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextEdit) ConnectCopyAvailable(f func(yes bool)) {
	if ptr.Pointer() != nil {
		C.QTextEdit_ConnectCopyAvailable(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "copyAvailable", f)
	}
}

func (ptr *QTextEdit) DisconnectCopyAvailable() {
	if ptr.Pointer() != nil {
		C.QTextEdit_DisconnectCopyAvailable(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "copyAvailable")
	}
}

//export callbackQTextEditCopyAvailable
func callbackQTextEditCopyAvailable(ptrName *C.char, yes C.int) {
	qt.GetSignal(C.GoString(ptrName), "copyAvailable").(func(bool))(int(yes) != 0)
}

func (ptr *QTextEdit) CreateStandardContextMenu() *QMenu {
	if ptr.Pointer() != nil {
		return QMenuFromPointer(unsafe.Pointer(C.QTextEdit_CreateStandardContextMenu(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTextEdit) CreateStandardContextMenu2(position core.QPointITF) *QMenu {
	if ptr.Pointer() != nil {
		return QMenuFromPointer(unsafe.Pointer(C.QTextEdit_CreateStandardContextMenu2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(position)))))
	}
	return nil
}

func (ptr *QTextEdit) ConnectCursorPositionChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QTextEdit_ConnectCursorPositionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "cursorPositionChanged", f)
	}
}

func (ptr *QTextEdit) DisconnectCursorPositionChanged() {
	if ptr.Pointer() != nil {
		C.QTextEdit_DisconnectCursorPositionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "cursorPositionChanged")
	}
}

//export callbackQTextEditCursorPositionChanged
func callbackQTextEditCursorPositionChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "cursorPositionChanged").(func())()
}

func (ptr *QTextEdit) Cut() {
	if ptr.Pointer() != nil {
		C.QTextEdit_Cut(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextEdit) DocumentTitle() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextEdit_DocumentTitle(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTextEdit) EnsureCursorVisible() {
	if ptr.Pointer() != nil {
		C.QTextEdit_EnsureCursorVisible(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextEdit) FontFamily() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextEdit_FontFamily(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTextEdit) FontItalic() bool {
	if ptr.Pointer() != nil {
		return C.QTextEdit_FontItalic(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextEdit) FontUnderline() bool {
	if ptr.Pointer() != nil {
		return C.QTextEdit_FontUnderline(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextEdit) FontWeight() int {
	if ptr.Pointer() != nil {
		return int(C.QTextEdit_FontWeight(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextEdit) InputMethodQuery(property core.Qt__InputMethodQuery) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextEdit_InputMethodQuery(C.QtObjectPtr(ptr.Pointer()), C.int(property)))
	}
	return ""
}

func (ptr *QTextEdit) InsertHtml(text string) {
	if ptr.Pointer() != nil {
		C.QTextEdit_InsertHtml(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QTextEdit) InsertPlainText(text string) {
	if ptr.Pointer() != nil {
		C.QTextEdit_InsertPlainText(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QTextEdit) IsUndoRedoEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QTextEdit_IsUndoRedoEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextEdit) LoadResource(ty int, name string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextEdit_LoadResource(C.QtObjectPtr(ptr.Pointer()), C.int(ty), C.CString(name)))
	}
	return ""
}

func (ptr *QTextEdit) MergeCurrentCharFormat(modifier gui.QTextCharFormatITF) {
	if ptr.Pointer() != nil {
		C.QTextEdit_MergeCurrentCharFormat(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQTextCharFormat(modifier)))
	}
}

func (ptr *QTextEdit) MoveCursor(operation gui.QTextCursor__MoveOperation, mode gui.QTextCursor__MoveMode) {
	if ptr.Pointer() != nil {
		C.QTextEdit_MoveCursor(C.QtObjectPtr(ptr.Pointer()), C.int(operation), C.int(mode))
	}
}

func (ptr *QTextEdit) Paste() {
	if ptr.Pointer() != nil {
		C.QTextEdit_Paste(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextEdit) Print(printer gui.QPagedPaintDeviceITF) {
	if ptr.Pointer() != nil {
		C.QTextEdit_Print(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPagedPaintDevice(printer)))
	}
}

func (ptr *QTextEdit) ConnectRedoAvailable(f func(available bool)) {
	if ptr.Pointer() != nil {
		C.QTextEdit_ConnectRedoAvailable(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "redoAvailable", f)
	}
}

func (ptr *QTextEdit) DisconnectRedoAvailable() {
	if ptr.Pointer() != nil {
		C.QTextEdit_DisconnectRedoAvailable(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "redoAvailable")
	}
}

//export callbackQTextEditRedoAvailable
func callbackQTextEditRedoAvailable(ptrName *C.char, available C.int) {
	qt.GetSignal(C.GoString(ptrName), "redoAvailable").(func(bool))(int(available) != 0)
}

func (ptr *QTextEdit) ScrollToAnchor(name string) {
	if ptr.Pointer() != nil {
		C.QTextEdit_ScrollToAnchor(C.QtObjectPtr(ptr.Pointer()), C.CString(name))
	}
}

func (ptr *QTextEdit) SelectAll() {
	if ptr.Pointer() != nil {
		C.QTextEdit_SelectAll(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextEdit) ConnectSelectionChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QTextEdit_ConnectSelectionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "selectionChanged", f)
	}
}

func (ptr *QTextEdit) DisconnectSelectionChanged() {
	if ptr.Pointer() != nil {
		C.QTextEdit_DisconnectSelectionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "selectionChanged")
	}
}

//export callbackQTextEditSelectionChanged
func callbackQTextEditSelectionChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "selectionChanged").(func())()
}

func (ptr *QTextEdit) SetAlignment(a core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetAlignment(C.QtObjectPtr(ptr.Pointer()), C.int(a))
	}
}

func (ptr *QTextEdit) SetCurrentCharFormat(format gui.QTextCharFormatITF) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetCurrentCharFormat(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQTextCharFormat(format)))
	}
}

func (ptr *QTextEdit) SetCurrentFont(f gui.QFontITF) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetCurrentFont(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQFont(f)))
	}
}

func (ptr *QTextEdit) SetDocumentTitle(title string) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetDocumentTitle(C.QtObjectPtr(ptr.Pointer()), C.CString(title))
	}
}

func (ptr *QTextEdit) SetFontFamily(fontFamily string) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetFontFamily(C.QtObjectPtr(ptr.Pointer()), C.CString(fontFamily))
	}
}

func (ptr *QTextEdit) SetFontItalic(italic bool) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetFontItalic(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(italic)))
	}
}

func (ptr *QTextEdit) SetFontUnderline(underline bool) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetFontUnderline(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(underline)))
	}
}

func (ptr *QTextEdit) SetPlainText(text string) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetPlainText(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QTextEdit) SetText(text string) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetText(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QTextEdit) SetTextBackgroundColor(c gui.QColorITF) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetTextBackgroundColor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQColor(c)))
	}
}

func (ptr *QTextEdit) SetTextColor(c gui.QColorITF) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetTextColor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQColor(c)))
	}
}

func (ptr *QTextEdit) SetTextCursor(cursor gui.QTextCursorITF) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetTextCursor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQTextCursor(cursor)))
	}
}

func (ptr *QTextEdit) SetUndoRedoEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetUndoRedoEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTextEdit) ConnectTextChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QTextEdit_ConnectTextChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "textChanged", f)
	}
}

func (ptr *QTextEdit) DisconnectTextChanged() {
	if ptr.Pointer() != nil {
		C.QTextEdit_DisconnectTextChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "textChanged")
	}
}

//export callbackQTextEditTextChanged
func callbackQTextEditTextChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "textChanged").(func())()
}

func (ptr *QTextEdit) ToPlainText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextEdit_ToPlainText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTextEdit) Undo() {
	if ptr.Pointer() != nil {
		C.QTextEdit_Undo(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextEdit) ConnectUndoAvailable(f func(available bool)) {
	if ptr.Pointer() != nil {
		C.QTextEdit_ConnectUndoAvailable(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "undoAvailable", f)
	}
}

func (ptr *QTextEdit) DisconnectUndoAvailable() {
	if ptr.Pointer() != nil {
		C.QTextEdit_DisconnectUndoAvailable(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "undoAvailable")
	}
}

//export callbackQTextEditUndoAvailable
func callbackQTextEditUndoAvailable(ptrName *C.char, available C.int) {
	qt.GetSignal(C.GoString(ptrName), "undoAvailable").(func(bool))(int(available) != 0)
}

func (ptr *QTextEdit) DestroyQTextEdit() {
	if ptr.Pointer() != nil {
		C.QTextEdit_DestroyQTextEdit(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
