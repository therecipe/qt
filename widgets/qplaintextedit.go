package widgets

//#include "qplaintextedit.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QPlainTextEdit struct {
	QAbstractScrollArea
}

type QPlainTextEdit_ITF interface {
	QAbstractScrollArea_ITF
	QPlainTextEdit_PTR() *QPlainTextEdit
}

func PointerFromQPlainTextEdit(ptr QPlainTextEdit_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlainTextEdit_PTR().Pointer()
	}
	return nil
}

func NewQPlainTextEditFromPointer(ptr unsafe.Pointer) *QPlainTextEdit {
	var n = new(QPlainTextEdit)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QPlainTextEdit_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPlainTextEdit) QPlainTextEdit_PTR() *QPlainTextEdit {
	return ptr
}

//QPlainTextEdit::LineWrapMode
type QPlainTextEdit__LineWrapMode int64

const (
	QPlainTextEdit__NoWrap      = QPlainTextEdit__LineWrapMode(0)
	QPlainTextEdit__WidgetWidth = QPlainTextEdit__LineWrapMode(1)
)

func (ptr *QPlainTextEdit) BackgroundVisible() bool {
	if ptr.Pointer() != nil {
		return C.QPlainTextEdit_BackgroundVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPlainTextEdit) BlockCount() int {
	if ptr.Pointer() != nil {
		return int(C.QPlainTextEdit_BlockCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPlainTextEdit) CenterOnScroll() bool {
	if ptr.Pointer() != nil {
		return C.QPlainTextEdit_CenterOnScroll(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPlainTextEdit) CursorWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QPlainTextEdit_CursorWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPlainTextEdit) IsReadOnly() bool {
	if ptr.Pointer() != nil {
		return C.QPlainTextEdit_IsReadOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPlainTextEdit) LineWrapMode() QPlainTextEdit__LineWrapMode {
	if ptr.Pointer() != nil {
		return QPlainTextEdit__LineWrapMode(C.QPlainTextEdit_LineWrapMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPlainTextEdit) OverwriteMode() bool {
	if ptr.Pointer() != nil {
		return C.QPlainTextEdit_OverwriteMode(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPlainTextEdit) PlaceholderText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QPlainTextEdit_PlaceholderText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPlainTextEdit) Redo() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_Redo(ptr.Pointer())
	}
}

func (ptr *QPlainTextEdit) SetBackgroundVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetBackgroundVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QPlainTextEdit) SetCenterOnScroll(enabled bool) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetCenterOnScroll(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QPlainTextEdit) SetCursorWidth(width int) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetCursorWidth(ptr.Pointer(), C.int(width))
	}
}

func (ptr *QPlainTextEdit) SetLineWrapMode(mode QPlainTextEdit__LineWrapMode) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetLineWrapMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QPlainTextEdit) SetOverwriteMode(overwrite bool) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetOverwriteMode(ptr.Pointer(), C.int(qt.GoBoolToInt(overwrite)))
	}
}

func (ptr *QPlainTextEdit) SetPlaceholderText(placeholderText string) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetPlaceholderText(ptr.Pointer(), C.CString(placeholderText))
	}
}

func (ptr *QPlainTextEdit) SetReadOnly(ro bool) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetReadOnly(ptr.Pointer(), C.int(qt.GoBoolToInt(ro)))
	}
}

func (ptr *QPlainTextEdit) SetTabChangesFocus(b bool) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetTabChangesFocus(ptr.Pointer(), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QPlainTextEdit) SetTabStopWidth(width int) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetTabStopWidth(ptr.Pointer(), C.int(width))
	}
}

func (ptr *QPlainTextEdit) SetTextInteractionFlags(flags core.Qt__TextInteractionFlag) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetTextInteractionFlags(ptr.Pointer(), C.int(flags))
	}
}

func (ptr *QPlainTextEdit) SetWordWrapMode(policy gui.QTextOption__WrapMode) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetWordWrapMode(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QPlainTextEdit) TabChangesFocus() bool {
	if ptr.Pointer() != nil {
		return C.QPlainTextEdit_TabChangesFocus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPlainTextEdit) TabStopWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QPlainTextEdit_TabStopWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPlainTextEdit) TextInteractionFlags() core.Qt__TextInteractionFlag {
	if ptr.Pointer() != nil {
		return core.Qt__TextInteractionFlag(C.QPlainTextEdit_TextInteractionFlags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPlainTextEdit) WordWrapMode() gui.QTextOption__WrapMode {
	if ptr.Pointer() != nil {
		return gui.QTextOption__WrapMode(C.QPlainTextEdit_WordWrapMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPlainTextEdit) ZoomIn(ran int) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ZoomIn(ptr.Pointer(), C.int(ran))
	}
}

func (ptr *QPlainTextEdit) ZoomOut(ran int) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ZoomOut(ptr.Pointer(), C.int(ran))
	}
}

func NewQPlainTextEdit(parent QWidget_ITF) *QPlainTextEdit {
	return NewQPlainTextEditFromPointer(C.QPlainTextEdit_NewQPlainTextEdit(PointerFromQWidget(parent)))
}

func NewQPlainTextEdit2(text string, parent QWidget_ITF) *QPlainTextEdit {
	return NewQPlainTextEditFromPointer(C.QPlainTextEdit_NewQPlainTextEdit2(C.CString(text), PointerFromQWidget(parent)))
}

func (ptr *QPlainTextEdit) AnchorAt(pos core.QPoint_ITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QPlainTextEdit_AnchorAt(ptr.Pointer(), core.PointerFromQPoint(pos)))
	}
	return ""
}

func (ptr *QPlainTextEdit) AppendPlainText(text string) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_AppendPlainText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QPlainTextEdit) CenterCursor() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_CenterCursor(ptr.Pointer())
	}
}

func (ptr *QPlainTextEdit) AppendHtml(html string) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_AppendHtml(ptr.Pointer(), C.CString(html))
	}
}

func (ptr *QPlainTextEdit) ConnectBlockCountChanged(f func(newBlockCount int)) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ConnectBlockCountChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "blockCountChanged", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectBlockCountChanged() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DisconnectBlockCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "blockCountChanged")
	}
}

//export callbackQPlainTextEditBlockCountChanged
func callbackQPlainTextEditBlockCountChanged(ptrName *C.char, newBlockCount C.int) {
	qt.GetSignal(C.GoString(ptrName), "blockCountChanged").(func(int))(int(newBlockCount))
}

func (ptr *QPlainTextEdit) CanPaste() bool {
	if ptr.Pointer() != nil {
		return C.QPlainTextEdit_CanPaste(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPlainTextEdit) Clear() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_Clear(ptr.Pointer())
	}
}

func (ptr *QPlainTextEdit) Copy() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_Copy(ptr.Pointer())
	}
}

func (ptr *QPlainTextEdit) ConnectCopyAvailable(f func(yes bool)) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ConnectCopyAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "copyAvailable", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectCopyAvailable() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DisconnectCopyAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "copyAvailable")
	}
}

//export callbackQPlainTextEditCopyAvailable
func callbackQPlainTextEditCopyAvailable(ptrName *C.char, yes C.int) {
	qt.GetSignal(C.GoString(ptrName), "copyAvailable").(func(bool))(int(yes) != 0)
}

func (ptr *QPlainTextEdit) CreateStandardContextMenu() *QMenu {
	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QPlainTextEdit_CreateStandardContextMenu(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPlainTextEdit) CreateStandardContextMenu2(position core.QPoint_ITF) *QMenu {
	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QPlainTextEdit_CreateStandardContextMenu2(ptr.Pointer(), core.PointerFromQPoint(position)))
	}
	return nil
}

func (ptr *QPlainTextEdit) ConnectCursorPositionChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ConnectCursorPositionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cursorPositionChanged", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectCursorPositionChanged() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DisconnectCursorPositionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cursorPositionChanged")
	}
}

//export callbackQPlainTextEditCursorPositionChanged
func callbackQPlainTextEditCursorPositionChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "cursorPositionChanged").(func())()
}

func (ptr *QPlainTextEdit) Cut() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_Cut(ptr.Pointer())
	}
}

func (ptr *QPlainTextEdit) Document() *gui.QTextDocument {
	if ptr.Pointer() != nil {
		return gui.NewQTextDocumentFromPointer(C.QPlainTextEdit_Document(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPlainTextEdit) DocumentTitle() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QPlainTextEdit_DocumentTitle(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPlainTextEdit) EnsureCursorVisible() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_EnsureCursorVisible(ptr.Pointer())
	}
}

func (ptr *QPlainTextEdit) InputMethodQuery(property core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QPlainTextEdit_InputMethodQuery(ptr.Pointer(), C.int(property)))
	}
	return nil
}

func (ptr *QPlainTextEdit) InsertPlainText(text string) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_InsertPlainText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QPlainTextEdit) IsUndoRedoEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QPlainTextEdit_IsUndoRedoEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPlainTextEdit) LoadResource(ty int, name core.QUrl_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QPlainTextEdit_LoadResource(ptr.Pointer(), C.int(ty), core.PointerFromQUrl(name)))
	}
	return nil
}

func (ptr *QPlainTextEdit) MaximumBlockCount() int {
	if ptr.Pointer() != nil {
		return int(C.QPlainTextEdit_MaximumBlockCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPlainTextEdit) MergeCurrentCharFormat(modifier gui.QTextCharFormat_ITF) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_MergeCurrentCharFormat(ptr.Pointer(), gui.PointerFromQTextCharFormat(modifier))
	}
}

func (ptr *QPlainTextEdit) ConnectModificationChanged(f func(changed bool)) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ConnectModificationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "modificationChanged", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectModificationChanged() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DisconnectModificationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "modificationChanged")
	}
}

//export callbackQPlainTextEditModificationChanged
func callbackQPlainTextEditModificationChanged(ptrName *C.char, changed C.int) {
	qt.GetSignal(C.GoString(ptrName), "modificationChanged").(func(bool))(int(changed) != 0)
}

func (ptr *QPlainTextEdit) MoveCursor(operation gui.QTextCursor__MoveOperation, mode gui.QTextCursor__MoveMode) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_MoveCursor(ptr.Pointer(), C.int(operation), C.int(mode))
	}
}

func (ptr *QPlainTextEdit) Paste() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_Paste(ptr.Pointer())
	}
}

func (ptr *QPlainTextEdit) Print(printer gui.QPagedPaintDevice_ITF) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_Print(ptr.Pointer(), gui.PointerFromQPagedPaintDevice(printer))
	}
}

func (ptr *QPlainTextEdit) ConnectRedoAvailable(f func(available bool)) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ConnectRedoAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "redoAvailable", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectRedoAvailable() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DisconnectRedoAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "redoAvailable")
	}
}

//export callbackQPlainTextEditRedoAvailable
func callbackQPlainTextEditRedoAvailable(ptrName *C.char, available C.int) {
	qt.GetSignal(C.GoString(ptrName), "redoAvailable").(func(bool))(int(available) != 0)
}

func (ptr *QPlainTextEdit) SelectAll() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SelectAll(ptr.Pointer())
	}
}

func (ptr *QPlainTextEdit) ConnectSelectionChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ConnectSelectionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "selectionChanged", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectSelectionChanged() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DisconnectSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "selectionChanged")
	}
}

//export callbackQPlainTextEditSelectionChanged
func callbackQPlainTextEditSelectionChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "selectionChanged").(func())()
}

func (ptr *QPlainTextEdit) SetCurrentCharFormat(format gui.QTextCharFormat_ITF) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetCurrentCharFormat(ptr.Pointer(), gui.PointerFromQTextCharFormat(format))
	}
}

func (ptr *QPlainTextEdit) SetDocument(document gui.QTextDocument_ITF) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetDocument(ptr.Pointer(), gui.PointerFromQTextDocument(document))
	}
}

func (ptr *QPlainTextEdit) SetDocumentTitle(title string) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetDocumentTitle(ptr.Pointer(), C.CString(title))
	}
}

func (ptr *QPlainTextEdit) SetMaximumBlockCount(maximum int) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetMaximumBlockCount(ptr.Pointer(), C.int(maximum))
	}
}

func (ptr *QPlainTextEdit) SetPlainText(text string) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetPlainText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QPlainTextEdit) SetTextCursor(cursor gui.QTextCursor_ITF) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetTextCursor(ptr.Pointer(), gui.PointerFromQTextCursor(cursor))
	}
}

func (ptr *QPlainTextEdit) SetUndoRedoEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetUndoRedoEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QPlainTextEdit) ConnectTextChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ConnectTextChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "textChanged", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectTextChanged() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DisconnectTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "textChanged")
	}
}

//export callbackQPlainTextEditTextChanged
func callbackQPlainTextEditTextChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "textChanged").(func())()
}

func (ptr *QPlainTextEdit) ToPlainText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QPlainTextEdit_ToPlainText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPlainTextEdit) Undo() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_Undo(ptr.Pointer())
	}
}

func (ptr *QPlainTextEdit) ConnectUndoAvailable(f func(available bool)) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ConnectUndoAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "undoAvailable", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectUndoAvailable() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DisconnectUndoAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "undoAvailable")
	}
}

//export callbackQPlainTextEditUndoAvailable
func callbackQPlainTextEditUndoAvailable(ptrName *C.char, available C.int) {
	qt.GetSignal(C.GoString(ptrName), "undoAvailable").(func(bool))(int(available) != 0)
}

func (ptr *QPlainTextEdit) DestroyQPlainTextEdit() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DestroyQPlainTextEdit(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
