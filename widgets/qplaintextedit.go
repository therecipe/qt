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

type QPlainTextEditITF interface {
	QAbstractScrollAreaITF
	QPlainTextEditPTR() *QPlainTextEdit
}

func PointerFromQPlainTextEdit(ptr QPlainTextEditITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlainTextEditPTR().Pointer()
	}
	return nil
}

func QPlainTextEditFromPointer(ptr unsafe.Pointer) *QPlainTextEdit {
	var n = new(QPlainTextEdit)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QPlainTextEdit_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPlainTextEdit) QPlainTextEditPTR() *QPlainTextEdit {
	return ptr
}

//QPlainTextEdit::LineWrapMode
type QPlainTextEdit__LineWrapMode int

var (
	QPlainTextEdit__NoWrap      = QPlainTextEdit__LineWrapMode(0)
	QPlainTextEdit__WidgetWidth = QPlainTextEdit__LineWrapMode(1)
)

func (ptr *QPlainTextEdit) BackgroundVisible() bool {
	if ptr.Pointer() != nil {
		return C.QPlainTextEdit_BackgroundVisible(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPlainTextEdit) BlockCount() int {
	if ptr.Pointer() != nil {
		return int(C.QPlainTextEdit_BlockCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPlainTextEdit) CenterOnScroll() bool {
	if ptr.Pointer() != nil {
		return C.QPlainTextEdit_CenterOnScroll(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPlainTextEdit) CursorWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QPlainTextEdit_CursorWidth(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPlainTextEdit) IsReadOnly() bool {
	if ptr.Pointer() != nil {
		return C.QPlainTextEdit_IsReadOnly(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPlainTextEdit) LineWrapMode() QPlainTextEdit__LineWrapMode {
	if ptr.Pointer() != nil {
		return QPlainTextEdit__LineWrapMode(C.QPlainTextEdit_LineWrapMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPlainTextEdit) OverwriteMode() bool {
	if ptr.Pointer() != nil {
		return C.QPlainTextEdit_OverwriteMode(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPlainTextEdit) PlaceholderText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QPlainTextEdit_PlaceholderText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QPlainTextEdit) Redo() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_Redo(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QPlainTextEdit) SetBackgroundVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetBackgroundVisible(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QPlainTextEdit) SetCenterOnScroll(enabled bool) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetCenterOnScroll(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QPlainTextEdit) SetCursorWidth(width int) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetCursorWidth(C.QtObjectPtr(ptr.Pointer()), C.int(width))
	}
}

func (ptr *QPlainTextEdit) SetLineWrapMode(mode QPlainTextEdit__LineWrapMode) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetLineWrapMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QPlainTextEdit) SetOverwriteMode(overwrite bool) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetOverwriteMode(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(overwrite)))
	}
}

func (ptr *QPlainTextEdit) SetPlaceholderText(placeholderText string) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetPlaceholderText(C.QtObjectPtr(ptr.Pointer()), C.CString(placeholderText))
	}
}

func (ptr *QPlainTextEdit) SetReadOnly(ro bool) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetReadOnly(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(ro)))
	}
}

func (ptr *QPlainTextEdit) SetTabChangesFocus(b bool) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetTabChangesFocus(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QPlainTextEdit) SetTabStopWidth(width int) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetTabStopWidth(C.QtObjectPtr(ptr.Pointer()), C.int(width))
	}
}

func (ptr *QPlainTextEdit) SetTextInteractionFlags(flags core.Qt__TextInteractionFlag) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetTextInteractionFlags(C.QtObjectPtr(ptr.Pointer()), C.int(flags))
	}
}

func (ptr *QPlainTextEdit) SetWordWrapMode(policy gui.QTextOption__WrapMode) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetWordWrapMode(C.QtObjectPtr(ptr.Pointer()), C.int(policy))
	}
}

func (ptr *QPlainTextEdit) TabChangesFocus() bool {
	if ptr.Pointer() != nil {
		return C.QPlainTextEdit_TabChangesFocus(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPlainTextEdit) TabStopWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QPlainTextEdit_TabStopWidth(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPlainTextEdit) TextInteractionFlags() core.Qt__TextInteractionFlag {
	if ptr.Pointer() != nil {
		return core.Qt__TextInteractionFlag(C.QPlainTextEdit_TextInteractionFlags(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPlainTextEdit) WordWrapMode() gui.QTextOption__WrapMode {
	if ptr.Pointer() != nil {
		return gui.QTextOption__WrapMode(C.QPlainTextEdit_WordWrapMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPlainTextEdit) ZoomIn(ran int) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ZoomIn(C.QtObjectPtr(ptr.Pointer()), C.int(ran))
	}
}

func (ptr *QPlainTextEdit) ZoomOut(ran int) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ZoomOut(C.QtObjectPtr(ptr.Pointer()), C.int(ran))
	}
}

func NewQPlainTextEdit(parent QWidgetITF) *QPlainTextEdit {
	return QPlainTextEditFromPointer(unsafe.Pointer(C.QPlainTextEdit_NewQPlainTextEdit(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func NewQPlainTextEdit2(text string, parent QWidgetITF) *QPlainTextEdit {
	return QPlainTextEditFromPointer(unsafe.Pointer(C.QPlainTextEdit_NewQPlainTextEdit2(C.CString(text), C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QPlainTextEdit) AnchorAt(pos core.QPointITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QPlainTextEdit_AnchorAt(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(pos))))
	}
	return ""
}

func (ptr *QPlainTextEdit) AppendPlainText(text string) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_AppendPlainText(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QPlainTextEdit) CenterCursor() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_CenterCursor(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QPlainTextEdit) AppendHtml(html string) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_AppendHtml(C.QtObjectPtr(ptr.Pointer()), C.CString(html))
	}
}

func (ptr *QPlainTextEdit) ConnectBlockCountChanged(f func(newBlockCount int)) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ConnectBlockCountChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "blockCountChanged", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectBlockCountChanged() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DisconnectBlockCountChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "blockCountChanged")
	}
}

//export callbackQPlainTextEditBlockCountChanged
func callbackQPlainTextEditBlockCountChanged(ptrName *C.char, newBlockCount C.int) {
	qt.GetSignal(C.GoString(ptrName), "blockCountChanged").(func(int))(int(newBlockCount))
}

func (ptr *QPlainTextEdit) CanPaste() bool {
	if ptr.Pointer() != nil {
		return C.QPlainTextEdit_CanPaste(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPlainTextEdit) Clear() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QPlainTextEdit) Copy() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_Copy(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QPlainTextEdit) ConnectCopyAvailable(f func(yes bool)) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ConnectCopyAvailable(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "copyAvailable", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectCopyAvailable() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DisconnectCopyAvailable(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "copyAvailable")
	}
}

//export callbackQPlainTextEditCopyAvailable
func callbackQPlainTextEditCopyAvailable(ptrName *C.char, yes C.int) {
	qt.GetSignal(C.GoString(ptrName), "copyAvailable").(func(bool))(int(yes) != 0)
}

func (ptr *QPlainTextEdit) CreateStandardContextMenu() *QMenu {
	if ptr.Pointer() != nil {
		return QMenuFromPointer(unsafe.Pointer(C.QPlainTextEdit_CreateStandardContextMenu(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QPlainTextEdit) CreateStandardContextMenu2(position core.QPointITF) *QMenu {
	if ptr.Pointer() != nil {
		return QMenuFromPointer(unsafe.Pointer(C.QPlainTextEdit_CreateStandardContextMenu2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(position)))))
	}
	return nil
}

func (ptr *QPlainTextEdit) ConnectCursorPositionChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ConnectCursorPositionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "cursorPositionChanged", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectCursorPositionChanged() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DisconnectCursorPositionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "cursorPositionChanged")
	}
}

//export callbackQPlainTextEditCursorPositionChanged
func callbackQPlainTextEditCursorPositionChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "cursorPositionChanged").(func())()
}

func (ptr *QPlainTextEdit) Cut() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_Cut(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QPlainTextEdit) Document() *gui.QTextDocument {
	if ptr.Pointer() != nil {
		return gui.QTextDocumentFromPointer(unsafe.Pointer(C.QPlainTextEdit_Document(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QPlainTextEdit) DocumentTitle() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QPlainTextEdit_DocumentTitle(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QPlainTextEdit) EnsureCursorVisible() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_EnsureCursorVisible(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QPlainTextEdit) InputMethodQuery(property core.Qt__InputMethodQuery) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QPlainTextEdit_InputMethodQuery(C.QtObjectPtr(ptr.Pointer()), C.int(property)))
	}
	return ""
}

func (ptr *QPlainTextEdit) InsertPlainText(text string) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_InsertPlainText(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QPlainTextEdit) IsUndoRedoEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QPlainTextEdit_IsUndoRedoEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPlainTextEdit) LoadResource(ty int, name string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QPlainTextEdit_LoadResource(C.QtObjectPtr(ptr.Pointer()), C.int(ty), C.CString(name)))
	}
	return ""
}

func (ptr *QPlainTextEdit) MaximumBlockCount() int {
	if ptr.Pointer() != nil {
		return int(C.QPlainTextEdit_MaximumBlockCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPlainTextEdit) MergeCurrentCharFormat(modifier gui.QTextCharFormatITF) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_MergeCurrentCharFormat(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQTextCharFormat(modifier)))
	}
}

func (ptr *QPlainTextEdit) ConnectModificationChanged(f func(changed bool)) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ConnectModificationChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "modificationChanged", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectModificationChanged() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DisconnectModificationChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "modificationChanged")
	}
}

//export callbackQPlainTextEditModificationChanged
func callbackQPlainTextEditModificationChanged(ptrName *C.char, changed C.int) {
	qt.GetSignal(C.GoString(ptrName), "modificationChanged").(func(bool))(int(changed) != 0)
}

func (ptr *QPlainTextEdit) MoveCursor(operation gui.QTextCursor__MoveOperation, mode gui.QTextCursor__MoveMode) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_MoveCursor(C.QtObjectPtr(ptr.Pointer()), C.int(operation), C.int(mode))
	}
}

func (ptr *QPlainTextEdit) Paste() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_Paste(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QPlainTextEdit) Print(printer gui.QPagedPaintDeviceITF) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_Print(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPagedPaintDevice(printer)))
	}
}

func (ptr *QPlainTextEdit) ConnectRedoAvailable(f func(available bool)) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ConnectRedoAvailable(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "redoAvailable", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectRedoAvailable() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DisconnectRedoAvailable(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "redoAvailable")
	}
}

//export callbackQPlainTextEditRedoAvailable
func callbackQPlainTextEditRedoAvailable(ptrName *C.char, available C.int) {
	qt.GetSignal(C.GoString(ptrName), "redoAvailable").(func(bool))(int(available) != 0)
}

func (ptr *QPlainTextEdit) SelectAll() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SelectAll(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QPlainTextEdit) ConnectSelectionChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ConnectSelectionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "selectionChanged", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectSelectionChanged() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DisconnectSelectionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "selectionChanged")
	}
}

//export callbackQPlainTextEditSelectionChanged
func callbackQPlainTextEditSelectionChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "selectionChanged").(func())()
}

func (ptr *QPlainTextEdit) SetCurrentCharFormat(format gui.QTextCharFormatITF) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetCurrentCharFormat(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQTextCharFormat(format)))
	}
}

func (ptr *QPlainTextEdit) SetDocument(document gui.QTextDocumentITF) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetDocument(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQTextDocument(document)))
	}
}

func (ptr *QPlainTextEdit) SetDocumentTitle(title string) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetDocumentTitle(C.QtObjectPtr(ptr.Pointer()), C.CString(title))
	}
}

func (ptr *QPlainTextEdit) SetMaximumBlockCount(maximum int) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetMaximumBlockCount(C.QtObjectPtr(ptr.Pointer()), C.int(maximum))
	}
}

func (ptr *QPlainTextEdit) SetPlainText(text string) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetPlainText(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QPlainTextEdit) SetTextCursor(cursor gui.QTextCursorITF) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetTextCursor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQTextCursor(cursor)))
	}
}

func (ptr *QPlainTextEdit) SetUndoRedoEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetUndoRedoEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QPlainTextEdit) ConnectTextChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ConnectTextChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "textChanged", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectTextChanged() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DisconnectTextChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "textChanged")
	}
}

//export callbackQPlainTextEditTextChanged
func callbackQPlainTextEditTextChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "textChanged").(func())()
}

func (ptr *QPlainTextEdit) ToPlainText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QPlainTextEdit_ToPlainText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QPlainTextEdit) Undo() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_Undo(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QPlainTextEdit) ConnectUndoAvailable(f func(available bool)) {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ConnectUndoAvailable(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "undoAvailable", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectUndoAvailable() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DisconnectUndoAvailable(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "undoAvailable")
	}
}

//export callbackQPlainTextEditUndoAvailable
func callbackQPlainTextEditUndoAvailable(ptrName *C.char, available C.int) {
	qt.GetSignal(C.GoString(ptrName), "undoAvailable").(func(bool))(int(available) != 0)
}

func (ptr *QPlainTextEdit) DestroyQPlainTextEdit() {
	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DestroyQPlainTextEdit(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
