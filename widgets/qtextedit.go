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

type QTextEdit_ITF interface {
	QAbstractScrollArea_ITF
	QTextEdit_PTR() *QTextEdit
}

func PointerFromQTextEdit(ptr QTextEdit_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextEdit_PTR().Pointer()
	}
	return nil
}

func NewQTextEditFromPointer(ptr unsafe.Pointer) *QTextEdit {
	var n = new(QTextEdit)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTextEdit_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTextEdit) QTextEdit_PTR() *QTextEdit {
	return ptr
}

//QTextEdit::AutoFormattingFlag
type QTextEdit__AutoFormattingFlag int64

const (
	QTextEdit__AutoNone       = QTextEdit__AutoFormattingFlag(0)
	QTextEdit__AutoBulletList = QTextEdit__AutoFormattingFlag(0x00000001)
	QTextEdit__AutoAll        = QTextEdit__AutoFormattingFlag(0xffffffff)
)

//QTextEdit::LineWrapMode
type QTextEdit__LineWrapMode int64

const (
	QTextEdit__NoWrap           = QTextEdit__LineWrapMode(0)
	QTextEdit__WidgetWidth      = QTextEdit__LineWrapMode(1)
	QTextEdit__FixedPixelWidth  = QTextEdit__LineWrapMode(2)
	QTextEdit__FixedColumnWidth = QTextEdit__LineWrapMode(3)
)

func (ptr *QTextEdit) AcceptRichText() bool {
	if ptr.Pointer() != nil {
		return C.QTextEdit_AcceptRichText(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextEdit) AutoFormatting() QTextEdit__AutoFormattingFlag {
	if ptr.Pointer() != nil {
		return QTextEdit__AutoFormattingFlag(C.QTextEdit_AutoFormatting(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) CursorWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QTextEdit_CursorWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) Document() *gui.QTextDocument {
	if ptr.Pointer() != nil {
		return gui.NewQTextDocumentFromPointer(C.QTextEdit_Document(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextEdit) IsReadOnly() bool {
	if ptr.Pointer() != nil {
		return C.QTextEdit_IsReadOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextEdit) LineWrapColumnOrWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QTextEdit_LineWrapColumnOrWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) LineWrapMode() QTextEdit__LineWrapMode {
	if ptr.Pointer() != nil {
		return QTextEdit__LineWrapMode(C.QTextEdit_LineWrapMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) OverwriteMode() bool {
	if ptr.Pointer() != nil {
		return C.QTextEdit_OverwriteMode(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextEdit) PlaceholderText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextEdit_PlaceholderText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextEdit) Redo() {
	if ptr.Pointer() != nil {
		C.QTextEdit_Redo(ptr.Pointer())
	}
}

func (ptr *QTextEdit) SetAcceptRichText(accept bool) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetAcceptRichText(ptr.Pointer(), C.int(qt.GoBoolToInt(accept)))
	}
}

func (ptr *QTextEdit) SetAutoFormatting(features QTextEdit__AutoFormattingFlag) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetAutoFormatting(ptr.Pointer(), C.int(features))
	}
}

func (ptr *QTextEdit) SetCursorWidth(width int) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetCursorWidth(ptr.Pointer(), C.int(width))
	}
}

func (ptr *QTextEdit) SetDocument(document gui.QTextDocument_ITF) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetDocument(ptr.Pointer(), gui.PointerFromQTextDocument(document))
	}
}

func (ptr *QTextEdit) SetFontWeight(weight int) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetFontWeight(ptr.Pointer(), C.int(weight))
	}
}

func (ptr *QTextEdit) SetHtml(text string) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetHtml(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QTextEdit) SetLineWrapColumnOrWidth(w int) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetLineWrapColumnOrWidth(ptr.Pointer(), C.int(w))
	}
}

func (ptr *QTextEdit) SetLineWrapMode(mode QTextEdit__LineWrapMode) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetLineWrapMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QTextEdit) SetOverwriteMode(overwrite bool) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetOverwriteMode(ptr.Pointer(), C.int(qt.GoBoolToInt(overwrite)))
	}
}

func (ptr *QTextEdit) SetPlaceholderText(placeholderText string) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetPlaceholderText(ptr.Pointer(), C.CString(placeholderText))
	}
}

func (ptr *QTextEdit) SetReadOnly(ro bool) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetReadOnly(ptr.Pointer(), C.int(qt.GoBoolToInt(ro)))
	}
}

func (ptr *QTextEdit) SetTabChangesFocus(b bool) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetTabChangesFocus(ptr.Pointer(), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QTextEdit) SetTabStopWidth(width int) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetTabStopWidth(ptr.Pointer(), C.int(width))
	}
}

func (ptr *QTextEdit) SetTextInteractionFlags(flags core.Qt__TextInteractionFlag) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetTextInteractionFlags(ptr.Pointer(), C.int(flags))
	}
}

func (ptr *QTextEdit) SetWordWrapMode(policy gui.QTextOption__WrapMode) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetWordWrapMode(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QTextEdit) TabChangesFocus() bool {
	if ptr.Pointer() != nil {
		return C.QTextEdit_TabChangesFocus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextEdit) TabStopWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QTextEdit_TabStopWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) TextInteractionFlags() core.Qt__TextInteractionFlag {
	if ptr.Pointer() != nil {
		return core.Qt__TextInteractionFlag(C.QTextEdit_TextInteractionFlags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) ToHtml() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextEdit_ToHtml(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextEdit) WordWrapMode() gui.QTextOption__WrapMode {
	if ptr.Pointer() != nil {
		return gui.QTextOption__WrapMode(C.QTextEdit_WordWrapMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) ZoomIn(ran int) {
	if ptr.Pointer() != nil {
		C.QTextEdit_ZoomIn(ptr.Pointer(), C.int(ran))
	}
}

func (ptr *QTextEdit) ZoomOut(ran int) {
	if ptr.Pointer() != nil {
		C.QTextEdit_ZoomOut(ptr.Pointer(), C.int(ran))
	}
}

func NewQTextEdit(parent QWidget_ITF) *QTextEdit {
	return NewQTextEditFromPointer(C.QTextEdit_NewQTextEdit(PointerFromQWidget(parent)))
}

func NewQTextEdit2(text string, parent QWidget_ITF) *QTextEdit {
	return NewQTextEditFromPointer(C.QTextEdit_NewQTextEdit2(C.CString(text), PointerFromQWidget(parent)))
}

func (ptr *QTextEdit) Alignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QTextEdit_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) AnchorAt(pos core.QPoint_ITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextEdit_AnchorAt(ptr.Pointer(), core.PointerFromQPoint(pos)))
	}
	return ""
}

func (ptr *QTextEdit) Append(text string) {
	if ptr.Pointer() != nil {
		C.QTextEdit_Append(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QTextEdit) CanPaste() bool {
	if ptr.Pointer() != nil {
		return C.QTextEdit_CanPaste(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextEdit) Clear() {
	if ptr.Pointer() != nil {
		C.QTextEdit_Clear(ptr.Pointer())
	}
}

func (ptr *QTextEdit) Copy() {
	if ptr.Pointer() != nil {
		C.QTextEdit_Copy(ptr.Pointer())
	}
}

func (ptr *QTextEdit) ConnectCopyAvailable(f func(yes bool)) {
	if ptr.Pointer() != nil {
		C.QTextEdit_ConnectCopyAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "copyAvailable", f)
	}
}

func (ptr *QTextEdit) DisconnectCopyAvailable() {
	if ptr.Pointer() != nil {
		C.QTextEdit_DisconnectCopyAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "copyAvailable")
	}
}

//export callbackQTextEditCopyAvailable
func callbackQTextEditCopyAvailable(ptrName *C.char, yes C.int) {
	qt.GetSignal(C.GoString(ptrName), "copyAvailable").(func(bool))(int(yes) != 0)
}

func (ptr *QTextEdit) CreateStandardContextMenu() *QMenu {
	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QTextEdit_CreateStandardContextMenu(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextEdit) CreateStandardContextMenu2(position core.QPoint_ITF) *QMenu {
	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QTextEdit_CreateStandardContextMenu2(ptr.Pointer(), core.PointerFromQPoint(position)))
	}
	return nil
}

func (ptr *QTextEdit) ConnectCursorPositionChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QTextEdit_ConnectCursorPositionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cursorPositionChanged", f)
	}
}

func (ptr *QTextEdit) DisconnectCursorPositionChanged() {
	if ptr.Pointer() != nil {
		C.QTextEdit_DisconnectCursorPositionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cursorPositionChanged")
	}
}

//export callbackQTextEditCursorPositionChanged
func callbackQTextEditCursorPositionChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "cursorPositionChanged").(func())()
}

func (ptr *QTextEdit) Cut() {
	if ptr.Pointer() != nil {
		C.QTextEdit_Cut(ptr.Pointer())
	}
}

func (ptr *QTextEdit) DocumentTitle() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextEdit_DocumentTitle(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextEdit) EnsureCursorVisible() {
	if ptr.Pointer() != nil {
		C.QTextEdit_EnsureCursorVisible(ptr.Pointer())
	}
}

func (ptr *QTextEdit) FontFamily() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextEdit_FontFamily(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextEdit) FontItalic() bool {
	if ptr.Pointer() != nil {
		return C.QTextEdit_FontItalic(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextEdit) FontPointSize() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextEdit_FontPointSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) FontUnderline() bool {
	if ptr.Pointer() != nil {
		return C.QTextEdit_FontUnderline(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextEdit) FontWeight() int {
	if ptr.Pointer() != nil {
		return int(C.QTextEdit_FontWeight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) InputMethodQuery(property core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QTextEdit_InputMethodQuery(ptr.Pointer(), C.int(property)))
	}
	return nil
}

func (ptr *QTextEdit) InsertHtml(text string) {
	if ptr.Pointer() != nil {
		C.QTextEdit_InsertHtml(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QTextEdit) InsertPlainText(text string) {
	if ptr.Pointer() != nil {
		C.QTextEdit_InsertPlainText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QTextEdit) IsUndoRedoEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QTextEdit_IsUndoRedoEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextEdit) LoadResource(ty int, name core.QUrl_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QTextEdit_LoadResource(ptr.Pointer(), C.int(ty), core.PointerFromQUrl(name)))
	}
	return nil
}

func (ptr *QTextEdit) MergeCurrentCharFormat(modifier gui.QTextCharFormat_ITF) {
	if ptr.Pointer() != nil {
		C.QTextEdit_MergeCurrentCharFormat(ptr.Pointer(), gui.PointerFromQTextCharFormat(modifier))
	}
}

func (ptr *QTextEdit) MoveCursor(operation gui.QTextCursor__MoveOperation, mode gui.QTextCursor__MoveMode) {
	if ptr.Pointer() != nil {
		C.QTextEdit_MoveCursor(ptr.Pointer(), C.int(operation), C.int(mode))
	}
}

func (ptr *QTextEdit) Paste() {
	if ptr.Pointer() != nil {
		C.QTextEdit_Paste(ptr.Pointer())
	}
}

func (ptr *QTextEdit) Print(printer gui.QPagedPaintDevice_ITF) {
	if ptr.Pointer() != nil {
		C.QTextEdit_Print(ptr.Pointer(), gui.PointerFromQPagedPaintDevice(printer))
	}
}

func (ptr *QTextEdit) ConnectRedoAvailable(f func(available bool)) {
	if ptr.Pointer() != nil {
		C.QTextEdit_ConnectRedoAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "redoAvailable", f)
	}
}

func (ptr *QTextEdit) DisconnectRedoAvailable() {
	if ptr.Pointer() != nil {
		C.QTextEdit_DisconnectRedoAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "redoAvailable")
	}
}

//export callbackQTextEditRedoAvailable
func callbackQTextEditRedoAvailable(ptrName *C.char, available C.int) {
	qt.GetSignal(C.GoString(ptrName), "redoAvailable").(func(bool))(int(available) != 0)
}

func (ptr *QTextEdit) ScrollToAnchor(name string) {
	if ptr.Pointer() != nil {
		C.QTextEdit_ScrollToAnchor(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QTextEdit) SelectAll() {
	if ptr.Pointer() != nil {
		C.QTextEdit_SelectAll(ptr.Pointer())
	}
}

func (ptr *QTextEdit) ConnectSelectionChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QTextEdit_ConnectSelectionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "selectionChanged", f)
	}
}

func (ptr *QTextEdit) DisconnectSelectionChanged() {
	if ptr.Pointer() != nil {
		C.QTextEdit_DisconnectSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "selectionChanged")
	}
}

//export callbackQTextEditSelectionChanged
func callbackQTextEditSelectionChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "selectionChanged").(func())()
}

func (ptr *QTextEdit) SetAlignment(a core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetAlignment(ptr.Pointer(), C.int(a))
	}
}

func (ptr *QTextEdit) SetCurrentCharFormat(format gui.QTextCharFormat_ITF) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetCurrentCharFormat(ptr.Pointer(), gui.PointerFromQTextCharFormat(format))
	}
}

func (ptr *QTextEdit) SetCurrentFont(f gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetCurrentFont(ptr.Pointer(), gui.PointerFromQFont(f))
	}
}

func (ptr *QTextEdit) SetDocumentTitle(title string) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetDocumentTitle(ptr.Pointer(), C.CString(title))
	}
}

func (ptr *QTextEdit) SetFontFamily(fontFamily string) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetFontFamily(ptr.Pointer(), C.CString(fontFamily))
	}
}

func (ptr *QTextEdit) SetFontItalic(italic bool) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetFontItalic(ptr.Pointer(), C.int(qt.GoBoolToInt(italic)))
	}
}

func (ptr *QTextEdit) SetFontPointSize(s float64) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetFontPointSize(ptr.Pointer(), C.double(s))
	}
}

func (ptr *QTextEdit) SetFontUnderline(underline bool) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetFontUnderline(ptr.Pointer(), C.int(qt.GoBoolToInt(underline)))
	}
}

func (ptr *QTextEdit) SetPlainText(text string) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetPlainText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QTextEdit) SetText(text string) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QTextEdit) SetTextBackgroundColor(c gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetTextBackgroundColor(ptr.Pointer(), gui.PointerFromQColor(c))
	}
}

func (ptr *QTextEdit) SetTextColor(c gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetTextColor(ptr.Pointer(), gui.PointerFromQColor(c))
	}
}

func (ptr *QTextEdit) SetTextCursor(cursor gui.QTextCursor_ITF) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetTextCursor(ptr.Pointer(), gui.PointerFromQTextCursor(cursor))
	}
}

func (ptr *QTextEdit) SetUndoRedoEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetUndoRedoEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTextEdit) TextBackgroundColor() *gui.QColor {
	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QTextEdit_TextBackgroundColor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextEdit) ConnectTextChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QTextEdit_ConnectTextChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "textChanged", f)
	}
}

func (ptr *QTextEdit) DisconnectTextChanged() {
	if ptr.Pointer() != nil {
		C.QTextEdit_DisconnectTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "textChanged")
	}
}

//export callbackQTextEditTextChanged
func callbackQTextEditTextChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "textChanged").(func())()
}

func (ptr *QTextEdit) TextColor() *gui.QColor {
	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QTextEdit_TextColor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextEdit) ToPlainText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextEdit_ToPlainText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextEdit) Undo() {
	if ptr.Pointer() != nil {
		C.QTextEdit_Undo(ptr.Pointer())
	}
}

func (ptr *QTextEdit) ConnectUndoAvailable(f func(available bool)) {
	if ptr.Pointer() != nil {
		C.QTextEdit_ConnectUndoAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "undoAvailable", f)
	}
}

func (ptr *QTextEdit) DisconnectUndoAvailable() {
	if ptr.Pointer() != nil {
		C.QTextEdit_DisconnectUndoAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "undoAvailable")
	}
}

//export callbackQTextEditUndoAvailable
func callbackQTextEditUndoAvailable(ptrName *C.char, available C.int) {
	qt.GetSignal(C.GoString(ptrName), "undoAvailable").(func(bool))(int(available) != 0)
}

func (ptr *QTextEdit) DestroyQTextEdit() {
	if ptr.Pointer() != nil {
		C.QTextEdit_DestroyQTextEdit(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
