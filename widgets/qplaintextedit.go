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
func callbackQPlainTextEditContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQPlainTextEditFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QPlainTextEdit) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QPlainTextEdit) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
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

func (ptr *QPlainTextEdit) CanInsertFromMimeData(source core.QMimeData_ITF) bool {
	defer qt.Recovering("QPlainTextEdit::canInsertFromMimeData")

	if ptr.Pointer() != nil {
		return C.QPlainTextEdit_CanInsertFromMimeData(ptr.Pointer(), core.PointerFromQMimeData(source)) != 0
	}
	return false
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
func callbackQPlainTextEditBlockCountChanged(ptr unsafe.Pointer, ptrName *C.char, newBlockCount C.int) {
	defer qt.Recovering("callback QPlainTextEdit::blockCountChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "blockCountChanged"); signal != nil {
		signal.(func(int))(int(newBlockCount))
	}

}

func (ptr *QPlainTextEdit) BlockCountChanged(newBlockCount int) {
	defer qt.Recovering("QPlainTextEdit::blockCountChanged")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_BlockCountChanged(ptr.Pointer(), C.int(newBlockCount))
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
func callbackQPlainTextEditChangeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
	} else {
		NewQPlainTextEditFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(e))
	}
}

func (ptr *QPlainTextEdit) ChangeEvent(e core.QEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::changeEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

func (ptr *QPlainTextEdit) ChangeEventDefault(e core.QEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::changeEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(e))
	}
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
func callbackQPlainTextEditCopyAvailable(ptr unsafe.Pointer, ptrName *C.char, yes C.int) {
	defer qt.Recovering("callback QPlainTextEdit::copyAvailable")

	if signal := qt.GetSignal(C.GoString(ptrName), "copyAvailable"); signal != nil {
		signal.(func(bool))(int(yes) != 0)
	}

}

func (ptr *QPlainTextEdit) CopyAvailable(yes bool) {
	defer qt.Recovering("QPlainTextEdit::copyAvailable")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_CopyAvailable(ptr.Pointer(), C.int(qt.GoBoolToInt(yes)))
	}
}

func (ptr *QPlainTextEdit) CreateMimeDataFromSelection() *core.QMimeData {
	defer qt.Recovering("QPlainTextEdit::createMimeDataFromSelection")

	if ptr.Pointer() != nil {
		return core.NewQMimeDataFromPointer(C.QPlainTextEdit_CreateMimeDataFromSelection(ptr.Pointer()))
	}
	return nil
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
func callbackQPlainTextEditCursorPositionChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QPlainTextEdit::cursorPositionChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "cursorPositionChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPlainTextEdit) CursorPositionChanged() {
	defer qt.Recovering("QPlainTextEdit::cursorPositionChanged")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_CursorPositionChanged(ptr.Pointer())
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
func callbackQPlainTextEditDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(e))
	} else {
		NewQPlainTextEditFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(e))
	}
}

func (ptr *QPlainTextEdit) DragEnterEvent(e gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(e))
	}
}

func (ptr *QPlainTextEdit) DragEnterEventDefault(e gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(e))
	}
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
func callbackQPlainTextEditDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(e))
	} else {
		NewQPlainTextEditFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(e))
	}
}

func (ptr *QPlainTextEdit) DragLeaveEvent(e gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

func (ptr *QPlainTextEdit) DragLeaveEventDefault(e gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
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
func callbackQPlainTextEditDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(e))
	} else {
		NewQPlainTextEditFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(e))
	}
}

func (ptr *QPlainTextEdit) DragMoveEvent(e gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

func (ptr *QPlainTextEdit) DragMoveEventDefault(e gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
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
func callbackQPlainTextEditDropEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(e))
	} else {
		NewQPlainTextEditFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(e))
	}
}

func (ptr *QPlainTextEdit) DropEvent(e gui.QDropEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::dropEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
}

func (ptr *QPlainTextEdit) DropEventDefault(e gui.QDropEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::dropEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
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
func callbackQPlainTextEditFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
	} else {
		NewQPlainTextEditFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(e))
	}
}

func (ptr *QPlainTextEdit) FocusInEvent(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::focusInEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QPlainTextEdit) FocusInEventDefault(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::focusInEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QPlainTextEdit) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QPlainTextEdit::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QPlainTextEdit_FocusNextPrevChild(ptr.Pointer(), C.int(qt.GoBoolToInt(next))) != 0
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
func callbackQPlainTextEditFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
	} else {
		NewQPlainTextEditFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(e))
	}
}

func (ptr *QPlainTextEdit) FocusOutEvent(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QPlainTextEdit) FocusOutEventDefault(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
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
func callbackQPlainTextEditInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(e))
	} else {
		NewQPlainTextEditFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(e))
	}
}

func (ptr *QPlainTextEdit) InputMethodEvent(e gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(e))
	}
}

func (ptr *QPlainTextEdit) InputMethodEventDefault(e gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(e))
	}
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
func callbackQPlainTextEditInsertFromMimeData(ptr unsafe.Pointer, ptrName *C.char, source unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::insertFromMimeData")

	if signal := qt.GetSignal(C.GoString(ptrName), "insertFromMimeData"); signal != nil {
		signal.(func(*core.QMimeData))(core.NewQMimeDataFromPointer(source))
	} else {
		NewQPlainTextEditFromPointer(ptr).InsertFromMimeDataDefault(core.NewQMimeDataFromPointer(source))
	}
}

func (ptr *QPlainTextEdit) InsertFromMimeData(source core.QMimeData_ITF) {
	defer qt.Recovering("QPlainTextEdit::insertFromMimeData")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_InsertFromMimeData(ptr.Pointer(), core.PointerFromQMimeData(source))
	}
}

func (ptr *QPlainTextEdit) InsertFromMimeDataDefault(source core.QMimeData_ITF) {
	defer qt.Recovering("QPlainTextEdit::insertFromMimeData")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_InsertFromMimeDataDefault(ptr.Pointer(), core.PointerFromQMimeData(source))
	}
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
func callbackQPlainTextEditKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQPlainTextEditFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QPlainTextEdit) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QPlainTextEdit) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
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
func callbackQPlainTextEditKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQPlainTextEditFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QPlainTextEdit) KeyReleaseEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QPlainTextEdit) KeyReleaseEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
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
func callbackQPlainTextEditModificationChanged(ptr unsafe.Pointer, ptrName *C.char, changed C.int) {
	defer qt.Recovering("callback QPlainTextEdit::modificationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "modificationChanged"); signal != nil {
		signal.(func(bool))(int(changed) != 0)
	}

}

func (ptr *QPlainTextEdit) ModificationChanged(changed bool) {
	defer qt.Recovering("QPlainTextEdit::modificationChanged")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ModificationChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(changed)))
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
func callbackQPlainTextEditMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQPlainTextEditFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QPlainTextEdit) MouseDoubleClickEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QPlainTextEdit) MouseDoubleClickEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
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
func callbackQPlainTextEditMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQPlainTextEditFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QPlainTextEdit) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QPlainTextEdit) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
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
func callbackQPlainTextEditMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQPlainTextEditFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QPlainTextEdit) MousePressEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QPlainTextEdit) MousePressEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
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
func callbackQPlainTextEditMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQPlainTextEditFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QPlainTextEdit) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QPlainTextEdit) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
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
func callbackQPlainTextEditPaintEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(e))
	} else {
		NewQPlainTextEditFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(e))
	}
}

func (ptr *QPlainTextEdit) PaintEvent(e gui.QPaintEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::paintEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

func (ptr *QPlainTextEdit) PaintEventDefault(e gui.QPaintEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::paintEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
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
func callbackQPlainTextEditRedoAvailable(ptr unsafe.Pointer, ptrName *C.char, available C.int) {
	defer qt.Recovering("callback QPlainTextEdit::redoAvailable")

	if signal := qt.GetSignal(C.GoString(ptrName), "redoAvailable"); signal != nil {
		signal.(func(bool))(int(available) != 0)
	}

}

func (ptr *QPlainTextEdit) RedoAvailable(available bool) {
	defer qt.Recovering("QPlainTextEdit::redoAvailable")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_RedoAvailable(ptr.Pointer(), C.int(qt.GoBoolToInt(available)))
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
func callbackQPlainTextEditResizeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
	} else {
		NewQPlainTextEditFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(e))
	}
}

func (ptr *QPlainTextEdit) ResizeEvent(e gui.QResizeEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::resizeEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
}

func (ptr *QPlainTextEdit) ResizeEventDefault(e gui.QResizeEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::resizeEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
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
func callbackQPlainTextEditScrollContentsBy(ptr unsafe.Pointer, ptrName *C.char, dx C.int, dy C.int) {
	defer qt.Recovering("callback QPlainTextEdit::scrollContentsBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(dx), int(dy))
	} else {
		NewQPlainTextEditFromPointer(ptr).ScrollContentsByDefault(int(dx), int(dy))
	}
}

func (ptr *QPlainTextEdit) ScrollContentsBy(dx int, dy int) {
	defer qt.Recovering("QPlainTextEdit::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ScrollContentsBy(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QPlainTextEdit) ScrollContentsByDefault(dx int, dy int) {
	defer qt.Recovering("QPlainTextEdit::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ScrollContentsByDefault(ptr.Pointer(), C.int(dx), C.int(dy))
	}
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
func callbackQPlainTextEditSelectionChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QPlainTextEdit::selectionChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectionChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPlainTextEdit) SelectionChanged() {
	defer qt.Recovering("QPlainTextEdit::selectionChanged")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SelectionChanged(ptr.Pointer())
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
func callbackQPlainTextEditShowEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(v))
	} else {
		NewQPlainTextEditFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(v))
	}
}

func (ptr *QPlainTextEdit) ShowEvent(v gui.QShowEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::showEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(v))
	}
}

func (ptr *QPlainTextEdit) ShowEventDefault(v gui.QShowEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::showEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(v))
	}
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
func callbackQPlainTextEditTextChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QPlainTextEdit::textChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "textChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QPlainTextEdit) TextChanged() {
	defer qt.Recovering("QPlainTextEdit::textChanged")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_TextChanged(ptr.Pointer())
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
func callbackQPlainTextEditUndoAvailable(ptr unsafe.Pointer, ptrName *C.char, available C.int) {
	defer qt.Recovering("callback QPlainTextEdit::undoAvailable")

	if signal := qt.GetSignal(C.GoString(ptrName), "undoAvailable"); signal != nil {
		signal.(func(bool))(int(available) != 0)
	}

}

func (ptr *QPlainTextEdit) UndoAvailable(available bool) {
	defer qt.Recovering("QPlainTextEdit::undoAvailable")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_UndoAvailable(ptr.Pointer(), C.int(qt.GoBoolToInt(available)))
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
func callbackQPlainTextEditUpdateRequest(ptr unsafe.Pointer, ptrName *C.char, rect unsafe.Pointer, dy C.int) {
	defer qt.Recovering("callback QPlainTextEdit::updateRequest")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateRequest"); signal != nil {
		signal.(func(*core.QRect, int))(core.NewQRectFromPointer(rect), int(dy))
	}

}

func (ptr *QPlainTextEdit) UpdateRequest(rect core.QRect_ITF, dy int) {
	defer qt.Recovering("QPlainTextEdit::updateRequest")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_UpdateRequest(ptr.Pointer(), core.PointerFromQRect(rect), C.int(dy))
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
func callbackQPlainTextEditWheelEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQPlainTextEditFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QPlainTextEdit) WheelEvent(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::wheelEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QPlainTextEdit) WheelEventDefault(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::wheelEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QPlainTextEdit) DestroyQPlainTextEdit() {
	defer qt.Recovering("QPlainTextEdit::~QPlainTextEdit")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_DestroyQPlainTextEdit(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPlainTextEdit) ConnectSetupViewport(f func(viewport *QWidget)) {
	defer qt.Recovering("connect QPlainTextEdit::setupViewport")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setupViewport", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectSetupViewport() {
	defer qt.Recovering("disconnect QPlainTextEdit::setupViewport")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setupViewport")
	}
}

//export callbackQPlainTextEditSetupViewport
func callbackQPlainTextEditSetupViewport(ptr unsafe.Pointer, ptrName *C.char, viewport unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::setupViewport")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupViewport"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(viewport))
	} else {
		NewQPlainTextEditFromPointer(ptr).SetupViewportDefault(NewQWidgetFromPointer(viewport))
	}
}

func (ptr *QPlainTextEdit) SetupViewport(viewport QWidget_ITF) {
	defer qt.Recovering("QPlainTextEdit::setupViewport")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetupViewport(ptr.Pointer(), PointerFromQWidget(viewport))
	}
}

func (ptr *QPlainTextEdit) SetupViewportDefault(viewport QWidget_ITF) {
	defer qt.Recovering("QPlainTextEdit::setupViewport")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetupViewportDefault(ptr.Pointer(), PointerFromQWidget(viewport))
	}
}

func (ptr *QPlainTextEdit) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QPlainTextEdit::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QPlainTextEdit::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQPlainTextEditActionEvent
func callbackQPlainTextEditActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQPlainTextEditFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QPlainTextEdit) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::actionEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QPlainTextEdit) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::actionEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QPlainTextEdit) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPlainTextEdit::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QPlainTextEdit::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQPlainTextEditEnterEvent
func callbackQPlainTextEditEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPlainTextEditFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPlainTextEdit) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::enterEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPlainTextEdit) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::enterEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPlainTextEdit) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QPlainTextEdit::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QPlainTextEdit::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQPlainTextEditHideEvent
func callbackQPlainTextEditHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQPlainTextEditFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QPlainTextEdit) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::hideEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QPlainTextEdit) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::hideEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QPlainTextEdit) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPlainTextEdit::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QPlainTextEdit::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQPlainTextEditLeaveEvent
func callbackQPlainTextEditLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPlainTextEditFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPlainTextEdit) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::leaveEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPlainTextEdit) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::leaveEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPlainTextEdit) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QPlainTextEdit::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QPlainTextEdit::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQPlainTextEditMoveEvent
func callbackQPlainTextEditMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQPlainTextEditFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QPlainTextEdit) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::moveEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QPlainTextEdit) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::moveEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QPlainTextEdit) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QPlainTextEdit::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QPlainTextEdit::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQPlainTextEditSetVisible
func callbackQPlainTextEditSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QPlainTextEdit::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QPlainTextEdit) SetVisible(visible bool) {
	defer qt.Recovering("QPlainTextEdit::setVisible")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QPlainTextEdit) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QPlainTextEdit::setVisible")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QPlainTextEdit) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QPlainTextEdit::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QPlainTextEdit::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQPlainTextEditCloseEvent
func callbackQPlainTextEditCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQPlainTextEditFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QPlainTextEdit) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::closeEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QPlainTextEdit) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::closeEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QPlainTextEdit) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QPlainTextEdit::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QPlainTextEdit::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQPlainTextEditInitPainter
func callbackQPlainTextEditInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQPlainTextEditFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QPlainTextEdit) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QPlainTextEdit::initPainter")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QPlainTextEdit) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QPlainTextEdit::initPainter")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QPlainTextEdit) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QPlainTextEdit::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QPlainTextEdit::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQPlainTextEditTabletEvent
func callbackQPlainTextEditTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQPlainTextEditFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QPlainTextEdit) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::tabletEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QPlainTextEdit) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::tabletEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QPlainTextEdit) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QPlainTextEdit::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QPlainTextEdit::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQPlainTextEditTimerEvent
func callbackQPlainTextEditTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQPlainTextEditFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QPlainTextEdit) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::timerEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QPlainTextEdit) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::timerEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QPlainTextEdit) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QPlainTextEdit::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QPlainTextEdit::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQPlainTextEditChildEvent
func callbackQPlainTextEditChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQPlainTextEditFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QPlainTextEdit) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::childEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QPlainTextEdit) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::childEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QPlainTextEdit) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPlainTextEdit::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QPlainTextEdit) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QPlainTextEdit::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQPlainTextEditCustomEvent
func callbackQPlainTextEditCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPlainTextEdit::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPlainTextEditFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPlainTextEdit) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::customEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPlainTextEdit) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QPlainTextEdit::customEvent")

	if ptr.Pointer() != nil {
		C.QPlainTextEdit_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
