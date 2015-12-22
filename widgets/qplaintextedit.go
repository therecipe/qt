package widgets

//#include "widgets.h"
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
	for len(n.ObjectName()) < len("QPlainTextEdit_") {
		n.SetObjectName("QPlainTextEdit_" + qt.Identifier())
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
	defer qt.Recovering("QPlainTextEdit::backgroundVisible")

	if ptr.Pointer() != nil {
		return C.QPlainTextEdit_BackgroundVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPlainTextEdit) BlockCount() int {
	defer qt.Recovering("QPlainTextEdit::blockCount")

	if ptr.Pointer() != nil {
		return int(C.QPlainTextEdit_BlockCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPlainTextEdit) CenterOnScroll() bool {
	defer qt.Recovering("QPlainTextEdit::centerOnScroll")

	if ptr.Pointer() != nil {
		return C.QPlainTextEdit_CenterOnScroll(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPlainTextEdit) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QPlainTextEdit::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QPlainTextEdit::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQPlainTextEditContextMenuEvent
func callbackQPlainTextEditContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QPlainTextEdit::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QPlainTextEdit) CursorWidth() int {
	defer qt.Recovering("QPlainTextEdit::cursorWidth")

	if ptr.Pointer() != nil {
		return int(C.QPlainTextEdit_CursorWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPlainTextEdit) IsReadOnly() bool {
	defer qt.Recovering("QPlainTextEdit::isReadOnly")

	if ptr.Pointer() != nil {
		return C.QPlainTextEdit_IsReadOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPlainTextEdit) LineWrapMode() QPlainTextEdit__LineWrapMode {
	defer qt.Recovering("QPlainTextEdit::lineWrapMode")

	if ptr.Pointer() != nil {
		return QPlainTextEdit__LineWrapMode(C.QPlainTextEdit_LineWrapMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPlainTextEdit) OverwriteMode() bool {
	defer qt.Recovering("QPlainTextEdit::overwriteMode")

	if ptr.Pointer() != nil {
		return C.QPlainTextEdit_OverwriteMode(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPlainTextEdit) PlaceholderText() string {
	defer qt.Recovering("QPlainTextEdit::placeholderText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QPlainTextEdit_PlaceholderText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPlainTextEdit) Redo() {
	defer qt.Recovering("QPlainTextEdit::redo")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_Redo(ptr.Pointer())
	}
}

func (ptr *QPlainTextEdit) SetBackgroundVisible(visible bool) {
	defer qt.Recovering("QPlainTextEdit::setBackgroundVisible")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetBackgroundVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QPlainTextEdit) SetCenterOnScroll(enabled bool) {
	defer qt.Recovering("QPlainTextEdit::setCenterOnScroll")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetCenterOnScroll(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QPlainTextEdit) SetCursorWidth(width int) {
	defer qt.Recovering("QPlainTextEdit::setCursorWidth")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetCursorWidth(ptr.Pointer(), C.int(width))
	}
}

func (ptr *QPlainTextEdit) SetLineWrapMode(mode QPlainTextEdit__LineWrapMode) {
	defer qt.Recovering("QPlainTextEdit::setLineWrapMode")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetLineWrapMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QPlainTextEdit) SetOverwriteMode(overwrite bool) {
	defer qt.Recovering("QPlainTextEdit::setOverwriteMode")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetOverwriteMode(ptr.Pointer(), C.int(qt.GoBoolToInt(overwrite)))
	}
}

func (ptr *QPlainTextEdit) SetPlaceholderText(placeholderText string) {
	defer qt.Recovering("QPlainTextEdit::setPlaceholderText")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetPlaceholderText(ptr.Pointer(), C.CString(placeholderText))
	}
}

func (ptr *QPlainTextEdit) SetReadOnly(ro bool) {
	defer qt.Recovering("QPlainTextEdit::setReadOnly")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetReadOnly(ptr.Pointer(), C.int(qt.GoBoolToInt(ro)))
	}
}

func (ptr *QPlainTextEdit) SetTabChangesFocus(b bool) {
	defer qt.Recovering("QPlainTextEdit::setTabChangesFocus")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetTabChangesFocus(ptr.Pointer(), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QPlainTextEdit) SetTabStopWidth(width int) {
	defer qt.Recovering("QPlainTextEdit::setTabStopWidth")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetTabStopWidth(ptr.Pointer(), C.int(width))
	}
}

func (ptr *QPlainTextEdit) SetTextInteractionFlags(flags core.Qt__TextInteractionFlag) {
	defer qt.Recovering("QPlainTextEdit::setTextInteractionFlags")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetTextInteractionFlags(ptr.Pointer(), C.int(flags))
	}
}

func (ptr *QPlainTextEdit) SetWordWrapMode(policy gui.QTextOption__WrapMode) {
	defer qt.Recovering("QPlainTextEdit::setWordWrapMode")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetWordWrapMode(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QPlainTextEdit) TabChangesFocus() bool {
	defer qt.Recovering("QPlainTextEdit::tabChangesFocus")

	if ptr.Pointer() != nil {
		return C.QPlainTextEdit_TabChangesFocus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPlainTextEdit) TabStopWidth() int {
	defer qt.Recovering("QPlainTextEdit::tabStopWidth")

	if ptr.Pointer() != nil {
		return int(C.QPlainTextEdit_TabStopWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPlainTextEdit) TextInteractionFlags() core.Qt__TextInteractionFlag {
	defer qt.Recovering("QPlainTextEdit::textInteractionFlags")

	if ptr.Pointer() != nil {
		return core.Qt__TextInteractionFlag(C.QPlainTextEdit_TextInteractionFlags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPlainTextEdit) WordWrapMode() gui.QTextOption__WrapMode {
	defer qt.Recovering("QPlainTextEdit::wordWrapMode")

	if ptr.Pointer() != nil {
		return gui.QTextOption__WrapMode(C.QPlainTextEdit_WordWrapMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPlainTextEdit) ZoomIn(ran int) {
	defer qt.Recovering("QPlainTextEdit::zoomIn")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ZoomIn(ptr.Pointer(), C.int(ran))
	}
}

func (ptr *QPlainTextEdit) ZoomOut(ran int) {
	defer qt.Recovering("QPlainTextEdit::zoomOut")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ZoomOut(ptr.Pointer(), C.int(ran))
	}
}

func NewQPlainTextEdit(parent QWidget_ITF) *QPlainTextEdit {
	defer qt.Recovering("QPlainTextEdit::QPlainTextEdit")

	return NewQPlainTextEditFromPointer(C.QPlainTextEdit_NewQPlainTextEdit(PointerFromQWidget(parent)))
}

func NewQPlainTextEdit2(text string, parent QWidget_ITF) *QPlainTextEdit {
	defer qt.Recovering("QPlainTextEdit::QPlainTextEdit")

	return NewQPlainTextEditFromPointer(C.QPlainTextEdit_NewQPlainTextEdit2(C.CString(text), PointerFromQWidget(parent)))
}

func (ptr *QPlainTextEdit) AnchorAt(pos core.QPoint_ITF) string {
	defer qt.Recovering("QPlainTextEdit::anchorAt")

	if ptr.Pointer() != nil {
		return C.GoString(C.QPlainTextEdit_AnchorAt(ptr.Pointer(), core.PointerFromQPoint(pos)))
	}
	return ""
}

func (ptr *QPlainTextEdit) AppendPlainText(text string) {
	defer qt.Recovering("QPlainTextEdit::appendPlainText")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_AppendPlainText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QPlainTextEdit) CenterCursor() {
	defer qt.Recovering("QPlainTextEdit::centerCursor")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_CenterCursor(ptr.Pointer())
	}
}

func (ptr *QPlainTextEdit) AppendHtml(html string) {
	defer qt.Recovering("QPlainTextEdit::appendHtml")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_AppendHtml(ptr.Pointer(), C.CString(html))
	}
}

func (ptr *QPlainTextEdit) ConnectBlockCountChanged(f func(newBlockCount int)) {
	defer qt.Recovering("connect QPlainTextEdit::blockCountChanged")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ConnectBlockCountChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "blockCountChanged", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectBlockCountChanged() {
	defer qt.Recovering("disconnect QPlainTextEdit::blockCountChanged")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DisconnectBlockCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "blockCountChanged")
	}
}

//export callbackQPlainTextEditBlockCountChanged
func callbackQPlainTextEditBlockCountChanged(ptrName *C.char, newBlockCount C.int) {
	defer qt.Recovering("callback QPlainTextEdit::blockCountChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "blockCountChanged"); signal != nil {
		signal.(func(int))(int(newBlockCount))
	}

}

func (ptr *QPlainTextEdit) CanPaste() bool {
	defer qt.Recovering("QPlainTextEdit::canPaste")

	if ptr.Pointer() != nil {
		return C.QPlainTextEdit_CanPaste(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPlainTextEdit) ConnectChangeEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QPlainTextEdit::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QPlainTextEdit::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQPlainTextEditChangeEvent
func callbackQPlainTextEditChangeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QPlainTextEdit::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QPlainTextEdit) Clear() {
	defer qt.Recovering("QPlainTextEdit::clear")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_Clear(ptr.Pointer())
	}
}

func (ptr *QPlainTextEdit) Copy() {
	defer qt.Recovering("QPlainTextEdit::copy")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_Copy(ptr.Pointer())
	}
}

func (ptr *QPlainTextEdit) ConnectCopyAvailable(f func(yes bool)) {
	defer qt.Recovering("connect QPlainTextEdit::copyAvailable")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ConnectCopyAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "copyAvailable", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectCopyAvailable() {
	defer qt.Recovering("disconnect QPlainTextEdit::copyAvailable")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DisconnectCopyAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "copyAvailable")
	}
}

//export callbackQPlainTextEditCopyAvailable
func callbackQPlainTextEditCopyAvailable(ptrName *C.char, yes C.int) {
	defer qt.Recovering("callback QPlainTextEdit::copyAvailable")

	if signal := qt.GetSignal(C.GoString(ptrName), "copyAvailable"); signal != nil {
		signal.(func(bool))(int(yes) != 0)
	}

}

func (ptr *QPlainTextEdit) CreateStandardContextMenu() *QMenu {
	defer qt.Recovering("QPlainTextEdit::createStandardContextMenu")

	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QPlainTextEdit_CreateStandardContextMenu(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPlainTextEdit) CreateStandardContextMenu2(position core.QPoint_ITF) *QMenu {
	defer qt.Recovering("QPlainTextEdit::createStandardContextMenu")

	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QPlainTextEdit_CreateStandardContextMenu2(ptr.Pointer(), core.PointerFromQPoint(position)))
	}
	return nil
}

func (ptr *QPlainTextEdit) ConnectCursorPositionChanged(f func()) {
	defer qt.Recovering("connect QPlainTextEdit::cursorPositionChanged")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ConnectCursorPositionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cursorPositionChanged", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectCursorPositionChanged() {
	defer qt.Recovering("disconnect QPlainTextEdit::cursorPositionChanged")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DisconnectCursorPositionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cursorPositionChanged")
	}
}

//export callbackQPlainTextEditCursorPositionChanged
func callbackQPlainTextEditCursorPositionChanged(ptrName *C.char) {
	defer qt.Recovering("callback QPlainTextEdit::cursorPositionChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "cursorPositionChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPlainTextEdit) CursorRect2() *core.QRect {
	defer qt.Recovering("QPlainTextEdit::cursorRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QPlainTextEdit_CursorRect2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPlainTextEdit) CursorRect(cursor gui.QTextCursor_ITF) *core.QRect {
	defer qt.Recovering("QPlainTextEdit::cursorRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QPlainTextEdit_CursorRect(ptr.Pointer(), gui.PointerFromQTextCursor(cursor)))
	}
	return nil
}

func (ptr *QPlainTextEdit) Cut() {
	defer qt.Recovering("QPlainTextEdit::cut")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_Cut(ptr.Pointer())
	}
}

func (ptr *QPlainTextEdit) Document() *gui.QTextDocument {
	defer qt.Recovering("QPlainTextEdit::document")

	if ptr.Pointer() != nil {
		return gui.NewQTextDocumentFromPointer(C.QPlainTextEdit_Document(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPlainTextEdit) DocumentTitle() string {
	defer qt.Recovering("QPlainTextEdit::documentTitle")

	if ptr.Pointer() != nil {
		return C.GoString(C.QPlainTextEdit_DocumentTitle(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPlainTextEdit) ConnectDragEnterEvent(f func(e *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QPlainTextEdit::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QPlainTextEdit::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQPlainTextEditDragEnterEvent
func callbackQPlainTextEditDragEnterEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QPlainTextEdit::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QPlainTextEdit) ConnectDragLeaveEvent(f func(e *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QPlainTextEdit::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QPlainTextEdit::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQPlainTextEditDragLeaveEvent
func callbackQPlainTextEditDragLeaveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QPlainTextEdit::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QPlainTextEdit) ConnectDragMoveEvent(f func(e *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QPlainTextEdit::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QPlainTextEdit::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQPlainTextEditDragMoveEvent
func callbackQPlainTextEditDragMoveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QPlainTextEdit::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QPlainTextEdit) ConnectDropEvent(f func(e *gui.QDropEvent)) {
	defer qt.Recovering("connect QPlainTextEdit::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QPlainTextEdit::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQPlainTextEditDropEvent
func callbackQPlainTextEditDropEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QPlainTextEdit::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QPlainTextEdit) EnsureCursorVisible() {
	defer qt.Recovering("QPlainTextEdit::ensureCursorVisible")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_EnsureCursorVisible(ptr.Pointer())
	}
}

func (ptr *QPlainTextEdit) ConnectFocusInEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QPlainTextEdit::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QPlainTextEdit::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQPlainTextEditFocusInEvent
func callbackQPlainTextEditFocusInEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QPlainTextEdit::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QPlainTextEdit) ConnectFocusOutEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QPlainTextEdit::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QPlainTextEdit::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQPlainTextEditFocusOutEvent
func callbackQPlainTextEditFocusOutEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QPlainTextEdit::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QPlainTextEdit) ConnectInputMethodEvent(f func(e *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QPlainTextEdit::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QPlainTextEdit::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQPlainTextEditInputMethodEvent
func callbackQPlainTextEditInputMethodEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QPlainTextEdit::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QPlainTextEdit) InputMethodQuery(property core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QPlainTextEdit::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QPlainTextEdit_InputMethodQuery(ptr.Pointer(), C.int(property)))
	}
	return nil
}

func (ptr *QPlainTextEdit) ConnectInsertFromMimeData(f func(source *core.QMimeData)) {
	defer qt.Recovering("connect QPlainTextEdit::insertFromMimeData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "insertFromMimeData", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectInsertFromMimeData() {
	defer qt.Recovering("disconnect QPlainTextEdit::insertFromMimeData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "insertFromMimeData")
	}
}

//export callbackQPlainTextEditInsertFromMimeData
func callbackQPlainTextEditInsertFromMimeData(ptrName *C.char, source unsafe.Pointer) bool {
	defer qt.Recovering("callback QPlainTextEdit::insertFromMimeData")

	if signal := qt.GetSignal(C.GoString(ptrName), "insertFromMimeData"); signal != nil {
		signal.(func(*core.QMimeData))(core.NewQMimeDataFromPointer(source))
		return true
	}
	return false

}

func (ptr *QPlainTextEdit) InsertPlainText(text string) {
	defer qt.Recovering("QPlainTextEdit::insertPlainText")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_InsertPlainText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QPlainTextEdit) IsUndoRedoEnabled() bool {
	defer qt.Recovering("QPlainTextEdit::isUndoRedoEnabled")

	if ptr.Pointer() != nil {
		return C.QPlainTextEdit_IsUndoRedoEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPlainTextEdit) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QPlainTextEdit::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QPlainTextEdit::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQPlainTextEditKeyPressEvent
func callbackQPlainTextEditKeyPressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QPlainTextEdit::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QPlainTextEdit) ConnectKeyReleaseEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QPlainTextEdit::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QPlainTextEdit::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQPlainTextEditKeyReleaseEvent
func callbackQPlainTextEditKeyReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QPlainTextEdit::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QPlainTextEdit) LoadResource(ty int, name core.QUrl_ITF) *core.QVariant {
	defer qt.Recovering("QPlainTextEdit::loadResource")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QPlainTextEdit_LoadResource(ptr.Pointer(), C.int(ty), core.PointerFromQUrl(name)))
	}
	return nil
}

func (ptr *QPlainTextEdit) MaximumBlockCount() int {
	defer qt.Recovering("QPlainTextEdit::maximumBlockCount")

	if ptr.Pointer() != nil {
		return int(C.QPlainTextEdit_MaximumBlockCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPlainTextEdit) MergeCurrentCharFormat(modifier gui.QTextCharFormat_ITF) {
	defer qt.Recovering("QPlainTextEdit::mergeCurrentCharFormat")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_MergeCurrentCharFormat(ptr.Pointer(), gui.PointerFromQTextCharFormat(modifier))
	}
}

func (ptr *QPlainTextEdit) ConnectModificationChanged(f func(changed bool)) {
	defer qt.Recovering("connect QPlainTextEdit::modificationChanged")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ConnectModificationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "modificationChanged", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectModificationChanged() {
	defer qt.Recovering("disconnect QPlainTextEdit::modificationChanged")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DisconnectModificationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "modificationChanged")
	}
}

//export callbackQPlainTextEditModificationChanged
func callbackQPlainTextEditModificationChanged(ptrName *C.char, changed C.int) {
	defer qt.Recovering("callback QPlainTextEdit::modificationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "modificationChanged"); signal != nil {
		signal.(func(bool))(int(changed) != 0)
	}

}

func (ptr *QPlainTextEdit) ConnectMouseDoubleClickEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QPlainTextEdit::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QPlainTextEdit::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQPlainTextEditMouseDoubleClickEvent
func callbackQPlainTextEditMouseDoubleClickEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QPlainTextEdit::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QPlainTextEdit) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QPlainTextEdit::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QPlainTextEdit::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQPlainTextEditMouseMoveEvent
func callbackQPlainTextEditMouseMoveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QPlainTextEdit::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QPlainTextEdit) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QPlainTextEdit::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QPlainTextEdit::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQPlainTextEditMousePressEvent
func callbackQPlainTextEditMousePressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QPlainTextEdit::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QPlainTextEdit) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QPlainTextEdit::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QPlainTextEdit::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQPlainTextEditMouseReleaseEvent
func callbackQPlainTextEditMouseReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QPlainTextEdit::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QPlainTextEdit) MoveCursor(operation gui.QTextCursor__MoveOperation, mode gui.QTextCursor__MoveMode) {
	defer qt.Recovering("QPlainTextEdit::moveCursor")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_MoveCursor(ptr.Pointer(), C.int(operation), C.int(mode))
	}
}

func (ptr *QPlainTextEdit) ConnectPaintEvent(f func(e *gui.QPaintEvent)) {
	defer qt.Recovering("connect QPlainTextEdit::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QPlainTextEdit::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQPlainTextEditPaintEvent
func callbackQPlainTextEditPaintEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QPlainTextEdit::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QPlainTextEdit) Paste() {
	defer qt.Recovering("QPlainTextEdit::paste")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_Paste(ptr.Pointer())
	}
}

func (ptr *QPlainTextEdit) Print(printer gui.QPagedPaintDevice_ITF) {
	defer qt.Recovering("QPlainTextEdit::print")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_Print(ptr.Pointer(), gui.PointerFromQPagedPaintDevice(printer))
	}
}

func (ptr *QPlainTextEdit) ConnectRedoAvailable(f func(available bool)) {
	defer qt.Recovering("connect QPlainTextEdit::redoAvailable")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ConnectRedoAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "redoAvailable", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectRedoAvailable() {
	defer qt.Recovering("disconnect QPlainTextEdit::redoAvailable")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DisconnectRedoAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "redoAvailable")
	}
}

//export callbackQPlainTextEditRedoAvailable
func callbackQPlainTextEditRedoAvailable(ptrName *C.char, available C.int) {
	defer qt.Recovering("callback QPlainTextEdit::redoAvailable")

	if signal := qt.GetSignal(C.GoString(ptrName), "redoAvailable"); signal != nil {
		signal.(func(bool))(int(available) != 0)
	}

}

func (ptr *QPlainTextEdit) ConnectResizeEvent(f func(e *gui.QResizeEvent)) {
	defer qt.Recovering("connect QPlainTextEdit::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QPlainTextEdit::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQPlainTextEditResizeEvent
func callbackQPlainTextEditResizeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QPlainTextEdit::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QPlainTextEdit) ConnectScrollContentsBy(f func(dx int, dy int)) {
	defer qt.Recovering("connect QPlainTextEdit::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollContentsBy", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectScrollContentsBy() {
	defer qt.Recovering("disconnect QPlainTextEdit::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollContentsBy")
	}
}

//export callbackQPlainTextEditScrollContentsBy
func callbackQPlainTextEditScrollContentsBy(ptrName *C.char, dx C.int, dy C.int) bool {
	defer qt.Recovering("callback QPlainTextEdit::scrollContentsBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(dx), int(dy))
		return true
	}
	return false

}

func (ptr *QPlainTextEdit) SelectAll() {
	defer qt.Recovering("QPlainTextEdit::selectAll")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SelectAll(ptr.Pointer())
	}
}

func (ptr *QPlainTextEdit) ConnectSelectionChanged(f func()) {
	defer qt.Recovering("connect QPlainTextEdit::selectionChanged")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ConnectSelectionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "selectionChanged", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectSelectionChanged() {
	defer qt.Recovering("disconnect QPlainTextEdit::selectionChanged")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DisconnectSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "selectionChanged")
	}
}

//export callbackQPlainTextEditSelectionChanged
func callbackQPlainTextEditSelectionChanged(ptrName *C.char) {
	defer qt.Recovering("callback QPlainTextEdit::selectionChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectionChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPlainTextEdit) SetCurrentCharFormat(format gui.QTextCharFormat_ITF) {
	defer qt.Recovering("QPlainTextEdit::setCurrentCharFormat")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetCurrentCharFormat(ptr.Pointer(), gui.PointerFromQTextCharFormat(format))
	}
}

func (ptr *QPlainTextEdit) SetDocument(document gui.QTextDocument_ITF) {
	defer qt.Recovering("QPlainTextEdit::setDocument")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetDocument(ptr.Pointer(), gui.PointerFromQTextDocument(document))
	}
}

func (ptr *QPlainTextEdit) SetDocumentTitle(title string) {
	defer qt.Recovering("QPlainTextEdit::setDocumentTitle")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetDocumentTitle(ptr.Pointer(), C.CString(title))
	}
}

func (ptr *QPlainTextEdit) SetMaximumBlockCount(maximum int) {
	defer qt.Recovering("QPlainTextEdit::setMaximumBlockCount")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetMaximumBlockCount(ptr.Pointer(), C.int(maximum))
	}
}

func (ptr *QPlainTextEdit) SetPlainText(text string) {
	defer qt.Recovering("QPlainTextEdit::setPlainText")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetPlainText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QPlainTextEdit) SetTextCursor(cursor gui.QTextCursor_ITF) {
	defer qt.Recovering("QPlainTextEdit::setTextCursor")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetTextCursor(ptr.Pointer(), gui.PointerFromQTextCursor(cursor))
	}
}

func (ptr *QPlainTextEdit) SetUndoRedoEnabled(enable bool) {
	defer qt.Recovering("QPlainTextEdit::setUndoRedoEnabled")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetUndoRedoEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QPlainTextEdit) ConnectShowEvent(f func(v *gui.QShowEvent)) {
	defer qt.Recovering("connect QPlainTextEdit::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QPlainTextEdit::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQPlainTextEditShowEvent
func callbackQPlainTextEditShowEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QPlainTextEdit::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QPlainTextEdit) ConnectTextChanged(f func()) {
	defer qt.Recovering("connect QPlainTextEdit::textChanged")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ConnectTextChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "textChanged", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectTextChanged() {
	defer qt.Recovering("disconnect QPlainTextEdit::textChanged")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DisconnectTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "textChanged")
	}
}

//export callbackQPlainTextEditTextChanged
func callbackQPlainTextEditTextChanged(ptrName *C.char) {
	defer qt.Recovering("callback QPlainTextEdit::textChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "textChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPlainTextEdit) ToPlainText() string {
	defer qt.Recovering("QPlainTextEdit::toPlainText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QPlainTextEdit_ToPlainText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPlainTextEdit) Undo() {
	defer qt.Recovering("QPlainTextEdit::undo")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_Undo(ptr.Pointer())
	}
}

func (ptr *QPlainTextEdit) ConnectUndoAvailable(f func(available bool)) {
	defer qt.Recovering("connect QPlainTextEdit::undoAvailable")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ConnectUndoAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "undoAvailable", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectUndoAvailable() {
	defer qt.Recovering("disconnect QPlainTextEdit::undoAvailable")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DisconnectUndoAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "undoAvailable")
	}
}

//export callbackQPlainTextEditUndoAvailable
func callbackQPlainTextEditUndoAvailable(ptrName *C.char, available C.int) {
	defer qt.Recovering("callback QPlainTextEdit::undoAvailable")

	if signal := qt.GetSignal(C.GoString(ptrName), "undoAvailable"); signal != nil {
		signal.(func(bool))(int(available) != 0)
	}

}

func (ptr *QPlainTextEdit) ConnectUpdateRequest(f func(rect *core.QRect, dy int)) {
	defer qt.Recovering("connect QPlainTextEdit::updateRequest")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ConnectUpdateRequest(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "updateRequest", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectUpdateRequest() {
	defer qt.Recovering("disconnect QPlainTextEdit::updateRequest")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DisconnectUpdateRequest(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "updateRequest")
	}
}

//export callbackQPlainTextEditUpdateRequest
func callbackQPlainTextEditUpdateRequest(ptrName *C.char, rect unsafe.Pointer, dy C.int) {
	defer qt.Recovering("callback QPlainTextEdit::updateRequest")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateRequest"); signal != nil {
		signal.(func(*core.QRect, int))(core.NewQRectFromPointer(rect), int(dy))
	}

}

func (ptr *QPlainTextEdit) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QPlainTextEdit::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QPlainTextEdit::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQPlainTextEditWheelEvent
func callbackQPlainTextEditWheelEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QPlainTextEdit::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QPlainTextEdit) DestroyQPlainTextEdit() {
	defer qt.Recovering("QPlainTextEdit::~QPlainTextEdit")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DestroyQPlainTextEdit(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
