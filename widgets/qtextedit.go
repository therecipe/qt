package widgets

//#include "widgets.h"
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
	for len(n.ObjectName()) < len("QTextEdit_") {
		n.SetObjectName("QTextEdit_" + qt.Identifier())
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
	defer qt.Recovering("QTextEdit::acceptRichText")

	if ptr.Pointer() != nil {
		return C.QTextEdit_AcceptRichText(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextEdit) AutoFormatting() QTextEdit__AutoFormattingFlag {
	defer qt.Recovering("QTextEdit::autoFormatting")

	if ptr.Pointer() != nil {
		return QTextEdit__AutoFormattingFlag(C.QTextEdit_AutoFormatting(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QTextEdit::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QTextEdit) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QTextEdit::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQTextEditContextMenuEvent
func callbackQTextEditContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQTextEditFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QTextEdit) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QTextEdit::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QTextEdit) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QTextEdit::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QTextEdit) CursorWidth() int {
	defer qt.Recovering("QTextEdit::cursorWidth")

	if ptr.Pointer() != nil {
		return int(C.QTextEdit_CursorWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) Document() *gui.QTextDocument {
	defer qt.Recovering("QTextEdit::document")

	if ptr.Pointer() != nil {
		return gui.NewQTextDocumentFromPointer(C.QTextEdit_Document(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextEdit) IsReadOnly() bool {
	defer qt.Recovering("QTextEdit::isReadOnly")

	if ptr.Pointer() != nil {
		return C.QTextEdit_IsReadOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextEdit) LineWrapColumnOrWidth() int {
	defer qt.Recovering("QTextEdit::lineWrapColumnOrWidth")

	if ptr.Pointer() != nil {
		return int(C.QTextEdit_LineWrapColumnOrWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) LineWrapMode() QTextEdit__LineWrapMode {
	defer qt.Recovering("QTextEdit::lineWrapMode")

	if ptr.Pointer() != nil {
		return QTextEdit__LineWrapMode(C.QTextEdit_LineWrapMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) OverwriteMode() bool {
	defer qt.Recovering("QTextEdit::overwriteMode")

	if ptr.Pointer() != nil {
		return C.QTextEdit_OverwriteMode(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextEdit) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QTextEdit::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QTextEdit) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QTextEdit::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQTextEditPaintEvent
func callbackQTextEditPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQTextEditFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QTextEdit) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QTextEdit::paintEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QTextEdit) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QTextEdit::paintEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QTextEdit) PlaceholderText() string {
	defer qt.Recovering("QTextEdit::placeholderText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextEdit_PlaceholderText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextEdit) Redo() {
	defer qt.Recovering("QTextEdit::redo")

	if ptr.Pointer() != nil {
		C.QTextEdit_Redo(ptr.Pointer())
	}
}

func (ptr *QTextEdit) SetAcceptRichText(accept bool) {
	defer qt.Recovering("QTextEdit::setAcceptRichText")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetAcceptRichText(ptr.Pointer(), C.int(qt.GoBoolToInt(accept)))
	}
}

func (ptr *QTextEdit) SetAutoFormatting(features QTextEdit__AutoFormattingFlag) {
	defer qt.Recovering("QTextEdit::setAutoFormatting")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetAutoFormatting(ptr.Pointer(), C.int(features))
	}
}

func (ptr *QTextEdit) SetCursorWidth(width int) {
	defer qt.Recovering("QTextEdit::setCursorWidth")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetCursorWidth(ptr.Pointer(), C.int(width))
	}
}

func (ptr *QTextEdit) SetDocument(document gui.QTextDocument_ITF) {
	defer qt.Recovering("QTextEdit::setDocument")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetDocument(ptr.Pointer(), gui.PointerFromQTextDocument(document))
	}
}

func (ptr *QTextEdit) SetFontWeight(weight int) {
	defer qt.Recovering("QTextEdit::setFontWeight")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetFontWeight(ptr.Pointer(), C.int(weight))
	}
}

func (ptr *QTextEdit) SetHtml(text string) {
	defer qt.Recovering("QTextEdit::setHtml")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetHtml(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QTextEdit) SetLineWrapColumnOrWidth(w int) {
	defer qt.Recovering("QTextEdit::setLineWrapColumnOrWidth")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetLineWrapColumnOrWidth(ptr.Pointer(), C.int(w))
	}
}

func (ptr *QTextEdit) SetLineWrapMode(mode QTextEdit__LineWrapMode) {
	defer qt.Recovering("QTextEdit::setLineWrapMode")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetLineWrapMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QTextEdit) SetOverwriteMode(overwrite bool) {
	defer qt.Recovering("QTextEdit::setOverwriteMode")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetOverwriteMode(ptr.Pointer(), C.int(qt.GoBoolToInt(overwrite)))
	}
}

func (ptr *QTextEdit) SetPlaceholderText(placeholderText string) {
	defer qt.Recovering("QTextEdit::setPlaceholderText")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetPlaceholderText(ptr.Pointer(), C.CString(placeholderText))
	}
}

func (ptr *QTextEdit) SetReadOnly(ro bool) {
	defer qt.Recovering("QTextEdit::setReadOnly")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetReadOnly(ptr.Pointer(), C.int(qt.GoBoolToInt(ro)))
	}
}

func (ptr *QTextEdit) SetTabChangesFocus(b bool) {
	defer qt.Recovering("QTextEdit::setTabChangesFocus")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetTabChangesFocus(ptr.Pointer(), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QTextEdit) SetTabStopWidth(width int) {
	defer qt.Recovering("QTextEdit::setTabStopWidth")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetTabStopWidth(ptr.Pointer(), C.int(width))
	}
}

func (ptr *QTextEdit) SetTextInteractionFlags(flags core.Qt__TextInteractionFlag) {
	defer qt.Recovering("QTextEdit::setTextInteractionFlags")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetTextInteractionFlags(ptr.Pointer(), C.int(flags))
	}
}

func (ptr *QTextEdit) SetWordWrapMode(policy gui.QTextOption__WrapMode) {
	defer qt.Recovering("QTextEdit::setWordWrapMode")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetWordWrapMode(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QTextEdit) TabChangesFocus() bool {
	defer qt.Recovering("QTextEdit::tabChangesFocus")

	if ptr.Pointer() != nil {
		return C.QTextEdit_TabChangesFocus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextEdit) TabStopWidth() int {
	defer qt.Recovering("QTextEdit::tabStopWidth")

	if ptr.Pointer() != nil {
		return int(C.QTextEdit_TabStopWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) TextInteractionFlags() core.Qt__TextInteractionFlag {
	defer qt.Recovering("QTextEdit::textInteractionFlags")

	if ptr.Pointer() != nil {
		return core.Qt__TextInteractionFlag(C.QTextEdit_TextInteractionFlags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) ToHtml() string {
	defer qt.Recovering("QTextEdit::toHtml")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextEdit_ToHtml(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextEdit) WordWrapMode() gui.QTextOption__WrapMode {
	defer qt.Recovering("QTextEdit::wordWrapMode")

	if ptr.Pointer() != nil {
		return gui.QTextOption__WrapMode(C.QTextEdit_WordWrapMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) ZoomIn(ran int) {
	defer qt.Recovering("QTextEdit::zoomIn")

	if ptr.Pointer() != nil {
		C.QTextEdit_ZoomIn(ptr.Pointer(), C.int(ran))
	}
}

func (ptr *QTextEdit) ZoomOut(ran int) {
	defer qt.Recovering("QTextEdit::zoomOut")

	if ptr.Pointer() != nil {
		C.QTextEdit_ZoomOut(ptr.Pointer(), C.int(ran))
	}
}

func NewQTextEdit(parent QWidget_ITF) *QTextEdit {
	defer qt.Recovering("QTextEdit::QTextEdit")

	return NewQTextEditFromPointer(C.QTextEdit_NewQTextEdit(PointerFromQWidget(parent)))
}

func NewQTextEdit2(text string, parent QWidget_ITF) *QTextEdit {
	defer qt.Recovering("QTextEdit::QTextEdit")

	return NewQTextEditFromPointer(C.QTextEdit_NewQTextEdit2(C.CString(text), PointerFromQWidget(parent)))
}

func (ptr *QTextEdit) Alignment() core.Qt__AlignmentFlag {
	defer qt.Recovering("QTextEdit::alignment")

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QTextEdit_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) AnchorAt(pos core.QPoint_ITF) string {
	defer qt.Recovering("QTextEdit::anchorAt")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextEdit_AnchorAt(ptr.Pointer(), core.PointerFromQPoint(pos)))
	}
	return ""
}

func (ptr *QTextEdit) Append(text string) {
	defer qt.Recovering("QTextEdit::append")

	if ptr.Pointer() != nil {
		C.QTextEdit_Append(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QTextEdit) CanInsertFromMimeData(source core.QMimeData_ITF) bool {
	defer qt.Recovering("QTextEdit::canInsertFromMimeData")

	if ptr.Pointer() != nil {
		return C.QTextEdit_CanInsertFromMimeData(ptr.Pointer(), core.PointerFromQMimeData(source)) != 0
	}
	return false
}

func (ptr *QTextEdit) CanPaste() bool {
	defer qt.Recovering("QTextEdit::canPaste")

	if ptr.Pointer() != nil {
		return C.QTextEdit_CanPaste(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextEdit) ConnectChangeEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QTextEdit::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QTextEdit) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QTextEdit::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQTextEditChangeEvent
func callbackQTextEditChangeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
	} else {
		NewQTextEditFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(e))
	}
}

func (ptr *QTextEdit) ChangeEvent(e core.QEvent_ITF) {
	defer qt.Recovering("QTextEdit::changeEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

func (ptr *QTextEdit) ChangeEventDefault(e core.QEvent_ITF) {
	defer qt.Recovering("QTextEdit::changeEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

func (ptr *QTextEdit) Clear() {
	defer qt.Recovering("QTextEdit::clear")

	if ptr.Pointer() != nil {
		C.QTextEdit_Clear(ptr.Pointer())
	}
}

func (ptr *QTextEdit) Copy() {
	defer qt.Recovering("QTextEdit::copy")

	if ptr.Pointer() != nil {
		C.QTextEdit_Copy(ptr.Pointer())
	}
}

func (ptr *QTextEdit) ConnectCopyAvailable(f func(yes bool)) {
	defer qt.Recovering("connect QTextEdit::copyAvailable")

	if ptr.Pointer() != nil {
		C.QTextEdit_ConnectCopyAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "copyAvailable", f)
	}
}

func (ptr *QTextEdit) DisconnectCopyAvailable() {
	defer qt.Recovering("disconnect QTextEdit::copyAvailable")

	if ptr.Pointer() != nil {
		C.QTextEdit_DisconnectCopyAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "copyAvailable")
	}
}

//export callbackQTextEditCopyAvailable
func callbackQTextEditCopyAvailable(ptr unsafe.Pointer, ptrName *C.char, yes C.int) {
	defer qt.Recovering("callback QTextEdit::copyAvailable")

	if signal := qt.GetSignal(C.GoString(ptrName), "copyAvailable"); signal != nil {
		signal.(func(bool))(int(yes) != 0)
	}

}

func (ptr *QTextEdit) CopyAvailable(yes bool) {
	defer qt.Recovering("QTextEdit::copyAvailable")

	if ptr.Pointer() != nil {
		C.QTextEdit_CopyAvailable(ptr.Pointer(), C.int(qt.GoBoolToInt(yes)))
	}
}

func (ptr *QTextEdit) CreateMimeDataFromSelection() *core.QMimeData {
	defer qt.Recovering("QTextEdit::createMimeDataFromSelection")

	if ptr.Pointer() != nil {
		return core.NewQMimeDataFromPointer(C.QTextEdit_CreateMimeDataFromSelection(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextEdit) CreateStandardContextMenu() *QMenu {
	defer qt.Recovering("QTextEdit::createStandardContextMenu")

	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QTextEdit_CreateStandardContextMenu(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextEdit) CreateStandardContextMenu2(position core.QPoint_ITF) *QMenu {
	defer qt.Recovering("QTextEdit::createStandardContextMenu")

	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QTextEdit_CreateStandardContextMenu2(ptr.Pointer(), core.PointerFromQPoint(position)))
	}
	return nil
}

func (ptr *QTextEdit) ConnectCursorPositionChanged(f func()) {
	defer qt.Recovering("connect QTextEdit::cursorPositionChanged")

	if ptr.Pointer() != nil {
		C.QTextEdit_ConnectCursorPositionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cursorPositionChanged", f)
	}
}

func (ptr *QTextEdit) DisconnectCursorPositionChanged() {
	defer qt.Recovering("disconnect QTextEdit::cursorPositionChanged")

	if ptr.Pointer() != nil {
		C.QTextEdit_DisconnectCursorPositionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cursorPositionChanged")
	}
}

//export callbackQTextEditCursorPositionChanged
func callbackQTextEditCursorPositionChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QTextEdit::cursorPositionChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "cursorPositionChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QTextEdit) CursorPositionChanged() {
	defer qt.Recovering("QTextEdit::cursorPositionChanged")

	if ptr.Pointer() != nil {
		C.QTextEdit_CursorPositionChanged(ptr.Pointer())
	}
}

func (ptr *QTextEdit) CursorRect2() *core.QRect {
	defer qt.Recovering("QTextEdit::cursorRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QTextEdit_CursorRect2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextEdit) CursorRect(cursor gui.QTextCursor_ITF) *core.QRect {
	defer qt.Recovering("QTextEdit::cursorRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QTextEdit_CursorRect(ptr.Pointer(), gui.PointerFromQTextCursor(cursor)))
	}
	return nil
}

func (ptr *QTextEdit) Cut() {
	defer qt.Recovering("QTextEdit::cut")

	if ptr.Pointer() != nil {
		C.QTextEdit_Cut(ptr.Pointer())
	}
}

func (ptr *QTextEdit) DocumentTitle() string {
	defer qt.Recovering("QTextEdit::documentTitle")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextEdit_DocumentTitle(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextEdit) ConnectDragEnterEvent(f func(e *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QTextEdit::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QTextEdit) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QTextEdit::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQTextEditDragEnterEvent
func callbackQTextEditDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(e))
	} else {
		NewQTextEditFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(e))
	}
}

func (ptr *QTextEdit) DragEnterEvent(e gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QTextEdit::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(e))
	}
}

func (ptr *QTextEdit) DragEnterEventDefault(e gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QTextEdit::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(e))
	}
}

func (ptr *QTextEdit) ConnectDragLeaveEvent(f func(e *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QTextEdit::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QTextEdit) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QTextEdit::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQTextEditDragLeaveEvent
func callbackQTextEditDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(e))
	} else {
		NewQTextEditFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(e))
	}
}

func (ptr *QTextEdit) DragLeaveEvent(e gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QTextEdit::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

func (ptr *QTextEdit) DragLeaveEventDefault(e gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QTextEdit::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

func (ptr *QTextEdit) ConnectDragMoveEvent(f func(e *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QTextEdit::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QTextEdit) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QTextEdit::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQTextEditDragMoveEvent
func callbackQTextEditDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(e))
	} else {
		NewQTextEditFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(e))
	}
}

func (ptr *QTextEdit) DragMoveEvent(e gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QTextEdit::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

func (ptr *QTextEdit) DragMoveEventDefault(e gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QTextEdit::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

func (ptr *QTextEdit) ConnectDropEvent(f func(e *gui.QDropEvent)) {
	defer qt.Recovering("connect QTextEdit::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QTextEdit) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QTextEdit::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQTextEditDropEvent
func callbackQTextEditDropEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(e))
	} else {
		NewQTextEditFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(e))
	}
}

func (ptr *QTextEdit) DropEvent(e gui.QDropEvent_ITF) {
	defer qt.Recovering("QTextEdit::dropEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
}

func (ptr *QTextEdit) DropEventDefault(e gui.QDropEvent_ITF) {
	defer qt.Recovering("QTextEdit::dropEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
}

func (ptr *QTextEdit) EnsureCursorVisible() {
	defer qt.Recovering("QTextEdit::ensureCursorVisible")

	if ptr.Pointer() != nil {
		C.QTextEdit_EnsureCursorVisible(ptr.Pointer())
	}
}

func (ptr *QTextEdit) ConnectFocusInEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QTextEdit::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QTextEdit) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QTextEdit::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQTextEditFocusInEvent
func callbackQTextEditFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
	} else {
		NewQTextEditFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(e))
	}
}

func (ptr *QTextEdit) FocusInEvent(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QTextEdit::focusInEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QTextEdit) FocusInEventDefault(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QTextEdit::focusInEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QTextEdit) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QTextEdit::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QTextEdit_FocusNextPrevChild(ptr.Pointer(), C.int(qt.GoBoolToInt(next))) != 0
	}
	return false
}

func (ptr *QTextEdit) ConnectFocusOutEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QTextEdit::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QTextEdit) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QTextEdit::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQTextEditFocusOutEvent
func callbackQTextEditFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
	} else {
		NewQTextEditFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(e))
	}
}

func (ptr *QTextEdit) FocusOutEvent(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QTextEdit::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QTextEdit) FocusOutEventDefault(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QTextEdit::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QTextEdit) FontFamily() string {
	defer qt.Recovering("QTextEdit::fontFamily")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextEdit_FontFamily(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextEdit) FontItalic() bool {
	defer qt.Recovering("QTextEdit::fontItalic")

	if ptr.Pointer() != nil {
		return C.QTextEdit_FontItalic(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextEdit) FontPointSize() float64 {
	defer qt.Recovering("QTextEdit::fontPointSize")

	if ptr.Pointer() != nil {
		return float64(C.QTextEdit_FontPointSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) FontUnderline() bool {
	defer qt.Recovering("QTextEdit::fontUnderline")

	if ptr.Pointer() != nil {
		return C.QTextEdit_FontUnderline(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextEdit) FontWeight() int {
	defer qt.Recovering("QTextEdit::fontWeight")

	if ptr.Pointer() != nil {
		return int(C.QTextEdit_FontWeight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) ConnectInputMethodEvent(f func(e *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QTextEdit::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QTextEdit) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QTextEdit::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQTextEditInputMethodEvent
func callbackQTextEditInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(e))
	} else {
		NewQTextEditFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(e))
	}
}

func (ptr *QTextEdit) InputMethodEvent(e gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QTextEdit::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(e))
	}
}

func (ptr *QTextEdit) InputMethodEventDefault(e gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QTextEdit::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(e))
	}
}

func (ptr *QTextEdit) InputMethodQuery(property core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QTextEdit::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QTextEdit_InputMethodQuery(ptr.Pointer(), C.int(property)))
	}
	return nil
}

func (ptr *QTextEdit) ConnectInsertFromMimeData(f func(source *core.QMimeData)) {
	defer qt.Recovering("connect QTextEdit::insertFromMimeData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "insertFromMimeData", f)
	}
}

func (ptr *QTextEdit) DisconnectInsertFromMimeData() {
	defer qt.Recovering("disconnect QTextEdit::insertFromMimeData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "insertFromMimeData")
	}
}

//export callbackQTextEditInsertFromMimeData
func callbackQTextEditInsertFromMimeData(ptr unsafe.Pointer, ptrName *C.char, source unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::insertFromMimeData")

	if signal := qt.GetSignal(C.GoString(ptrName), "insertFromMimeData"); signal != nil {
		signal.(func(*core.QMimeData))(core.NewQMimeDataFromPointer(source))
	} else {
		NewQTextEditFromPointer(ptr).InsertFromMimeDataDefault(core.NewQMimeDataFromPointer(source))
	}
}

func (ptr *QTextEdit) InsertFromMimeData(source core.QMimeData_ITF) {
	defer qt.Recovering("QTextEdit::insertFromMimeData")

	if ptr.Pointer() != nil {
		C.QTextEdit_InsertFromMimeData(ptr.Pointer(), core.PointerFromQMimeData(source))
	}
}

func (ptr *QTextEdit) InsertFromMimeDataDefault(source core.QMimeData_ITF) {
	defer qt.Recovering("QTextEdit::insertFromMimeData")

	if ptr.Pointer() != nil {
		C.QTextEdit_InsertFromMimeDataDefault(ptr.Pointer(), core.PointerFromQMimeData(source))
	}
}

func (ptr *QTextEdit) InsertHtml(text string) {
	defer qt.Recovering("QTextEdit::insertHtml")

	if ptr.Pointer() != nil {
		C.QTextEdit_InsertHtml(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QTextEdit) InsertPlainText(text string) {
	defer qt.Recovering("QTextEdit::insertPlainText")

	if ptr.Pointer() != nil {
		C.QTextEdit_InsertPlainText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QTextEdit) IsUndoRedoEnabled() bool {
	defer qt.Recovering("QTextEdit::isUndoRedoEnabled")

	if ptr.Pointer() != nil {
		return C.QTextEdit_IsUndoRedoEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextEdit) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QTextEdit::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QTextEdit) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QTextEdit::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQTextEditKeyPressEvent
func callbackQTextEditKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQTextEditFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QTextEdit) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QTextEdit::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QTextEdit) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QTextEdit::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QTextEdit) ConnectKeyReleaseEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QTextEdit::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QTextEdit) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QTextEdit::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQTextEditKeyReleaseEvent
func callbackQTextEditKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQTextEditFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QTextEdit) KeyReleaseEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QTextEdit::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QTextEdit) KeyReleaseEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QTextEdit::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QTextEdit) LoadResource(ty int, name core.QUrl_ITF) *core.QVariant {
	defer qt.Recovering("QTextEdit::loadResource")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QTextEdit_LoadResource(ptr.Pointer(), C.int(ty), core.PointerFromQUrl(name)))
	}
	return nil
}

func (ptr *QTextEdit) MergeCurrentCharFormat(modifier gui.QTextCharFormat_ITF) {
	defer qt.Recovering("QTextEdit::mergeCurrentCharFormat")

	if ptr.Pointer() != nil {
		C.QTextEdit_MergeCurrentCharFormat(ptr.Pointer(), gui.PointerFromQTextCharFormat(modifier))
	}
}

func (ptr *QTextEdit) ConnectMouseDoubleClickEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTextEdit::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QTextEdit) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QTextEdit::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQTextEditMouseDoubleClickEvent
func callbackQTextEditMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQTextEditFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QTextEdit) MouseDoubleClickEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTextEdit::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QTextEdit) MouseDoubleClickEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTextEdit::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QTextEdit) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTextEdit::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QTextEdit) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QTextEdit::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQTextEditMouseMoveEvent
func callbackQTextEditMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQTextEditFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QTextEdit) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTextEdit::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QTextEdit) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTextEdit::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QTextEdit) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTextEdit::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QTextEdit) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QTextEdit::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQTextEditMousePressEvent
func callbackQTextEditMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQTextEditFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QTextEdit) MousePressEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTextEdit::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QTextEdit) MousePressEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTextEdit::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QTextEdit) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTextEdit::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QTextEdit) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QTextEdit::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQTextEditMouseReleaseEvent
func callbackQTextEditMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQTextEditFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QTextEdit) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTextEdit::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QTextEdit) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTextEdit::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QTextEdit) MoveCursor(operation gui.QTextCursor__MoveOperation, mode gui.QTextCursor__MoveMode) {
	defer qt.Recovering("QTextEdit::moveCursor")

	if ptr.Pointer() != nil {
		C.QTextEdit_MoveCursor(ptr.Pointer(), C.int(operation), C.int(mode))
	}
}

func (ptr *QTextEdit) Paste() {
	defer qt.Recovering("QTextEdit::paste")

	if ptr.Pointer() != nil {
		C.QTextEdit_Paste(ptr.Pointer())
	}
}

func (ptr *QTextEdit) Print(printer gui.QPagedPaintDevice_ITF) {
	defer qt.Recovering("QTextEdit::print")

	if ptr.Pointer() != nil {
		C.QTextEdit_Print(ptr.Pointer(), gui.PointerFromQPagedPaintDevice(printer))
	}
}

func (ptr *QTextEdit) ConnectRedoAvailable(f func(available bool)) {
	defer qt.Recovering("connect QTextEdit::redoAvailable")

	if ptr.Pointer() != nil {
		C.QTextEdit_ConnectRedoAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "redoAvailable", f)
	}
}

func (ptr *QTextEdit) DisconnectRedoAvailable() {
	defer qt.Recovering("disconnect QTextEdit::redoAvailable")

	if ptr.Pointer() != nil {
		C.QTextEdit_DisconnectRedoAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "redoAvailable")
	}
}

//export callbackQTextEditRedoAvailable
func callbackQTextEditRedoAvailable(ptr unsafe.Pointer, ptrName *C.char, available C.int) {
	defer qt.Recovering("callback QTextEdit::redoAvailable")

	if signal := qt.GetSignal(C.GoString(ptrName), "redoAvailable"); signal != nil {
		signal.(func(bool))(int(available) != 0)
	}

}

func (ptr *QTextEdit) RedoAvailable(available bool) {
	defer qt.Recovering("QTextEdit::redoAvailable")

	if ptr.Pointer() != nil {
		C.QTextEdit_RedoAvailable(ptr.Pointer(), C.int(qt.GoBoolToInt(available)))
	}
}

func (ptr *QTextEdit) ConnectResizeEvent(f func(e *gui.QResizeEvent)) {
	defer qt.Recovering("connect QTextEdit::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QTextEdit) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QTextEdit::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQTextEditResizeEvent
func callbackQTextEditResizeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
	} else {
		NewQTextEditFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(e))
	}
}

func (ptr *QTextEdit) ResizeEvent(e gui.QResizeEvent_ITF) {
	defer qt.Recovering("QTextEdit::resizeEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
}

func (ptr *QTextEdit) ResizeEventDefault(e gui.QResizeEvent_ITF) {
	defer qt.Recovering("QTextEdit::resizeEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
}

func (ptr *QTextEdit) ConnectScrollContentsBy(f func(dx int, dy int)) {
	defer qt.Recovering("connect QTextEdit::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollContentsBy", f)
	}
}

func (ptr *QTextEdit) DisconnectScrollContentsBy() {
	defer qt.Recovering("disconnect QTextEdit::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollContentsBy")
	}
}

//export callbackQTextEditScrollContentsBy
func callbackQTextEditScrollContentsBy(ptr unsafe.Pointer, ptrName *C.char, dx C.int, dy C.int) {
	defer qt.Recovering("callback QTextEdit::scrollContentsBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(dx), int(dy))
	} else {
		NewQTextEditFromPointer(ptr).ScrollContentsByDefault(int(dx), int(dy))
	}
}

func (ptr *QTextEdit) ScrollContentsBy(dx int, dy int) {
	defer qt.Recovering("QTextEdit::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QTextEdit_ScrollContentsBy(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QTextEdit) ScrollContentsByDefault(dx int, dy int) {
	defer qt.Recovering("QTextEdit::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QTextEdit_ScrollContentsByDefault(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QTextEdit) ScrollToAnchor(name string) {
	defer qt.Recovering("QTextEdit::scrollToAnchor")

	if ptr.Pointer() != nil {
		C.QTextEdit_ScrollToAnchor(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QTextEdit) SelectAll() {
	defer qt.Recovering("QTextEdit::selectAll")

	if ptr.Pointer() != nil {
		C.QTextEdit_SelectAll(ptr.Pointer())
	}
}

func (ptr *QTextEdit) ConnectSelectionChanged(f func()) {
	defer qt.Recovering("connect QTextEdit::selectionChanged")

	if ptr.Pointer() != nil {
		C.QTextEdit_ConnectSelectionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "selectionChanged", f)
	}
}

func (ptr *QTextEdit) DisconnectSelectionChanged() {
	defer qt.Recovering("disconnect QTextEdit::selectionChanged")

	if ptr.Pointer() != nil {
		C.QTextEdit_DisconnectSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "selectionChanged")
	}
}

//export callbackQTextEditSelectionChanged
func callbackQTextEditSelectionChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QTextEdit::selectionChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectionChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QTextEdit) SelectionChanged() {
	defer qt.Recovering("QTextEdit::selectionChanged")

	if ptr.Pointer() != nil {
		C.QTextEdit_SelectionChanged(ptr.Pointer())
	}
}

func (ptr *QTextEdit) SetAlignment(a core.Qt__AlignmentFlag) {
	defer qt.Recovering("QTextEdit::setAlignment")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetAlignment(ptr.Pointer(), C.int(a))
	}
}

func (ptr *QTextEdit) SetCurrentCharFormat(format gui.QTextCharFormat_ITF) {
	defer qt.Recovering("QTextEdit::setCurrentCharFormat")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetCurrentCharFormat(ptr.Pointer(), gui.PointerFromQTextCharFormat(format))
	}
}

func (ptr *QTextEdit) SetCurrentFont(f gui.QFont_ITF) {
	defer qt.Recovering("QTextEdit::setCurrentFont")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetCurrentFont(ptr.Pointer(), gui.PointerFromQFont(f))
	}
}

func (ptr *QTextEdit) SetDocumentTitle(title string) {
	defer qt.Recovering("QTextEdit::setDocumentTitle")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetDocumentTitle(ptr.Pointer(), C.CString(title))
	}
}

func (ptr *QTextEdit) SetFontFamily(fontFamily string) {
	defer qt.Recovering("QTextEdit::setFontFamily")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetFontFamily(ptr.Pointer(), C.CString(fontFamily))
	}
}

func (ptr *QTextEdit) SetFontItalic(italic bool) {
	defer qt.Recovering("QTextEdit::setFontItalic")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetFontItalic(ptr.Pointer(), C.int(qt.GoBoolToInt(italic)))
	}
}

func (ptr *QTextEdit) SetFontPointSize(s float64) {
	defer qt.Recovering("QTextEdit::setFontPointSize")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetFontPointSize(ptr.Pointer(), C.double(s))
	}
}

func (ptr *QTextEdit) SetFontUnderline(underline bool) {
	defer qt.Recovering("QTextEdit::setFontUnderline")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetFontUnderline(ptr.Pointer(), C.int(qt.GoBoolToInt(underline)))
	}
}

func (ptr *QTextEdit) SetPlainText(text string) {
	defer qt.Recovering("QTextEdit::setPlainText")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetPlainText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QTextEdit) SetText(text string) {
	defer qt.Recovering("QTextEdit::setText")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QTextEdit) SetTextBackgroundColor(c gui.QColor_ITF) {
	defer qt.Recovering("QTextEdit::setTextBackgroundColor")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetTextBackgroundColor(ptr.Pointer(), gui.PointerFromQColor(c))
	}
}

func (ptr *QTextEdit) SetTextColor(c gui.QColor_ITF) {
	defer qt.Recovering("QTextEdit::setTextColor")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetTextColor(ptr.Pointer(), gui.PointerFromQColor(c))
	}
}

func (ptr *QTextEdit) SetTextCursor(cursor gui.QTextCursor_ITF) {
	defer qt.Recovering("QTextEdit::setTextCursor")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetTextCursor(ptr.Pointer(), gui.PointerFromQTextCursor(cursor))
	}
}

func (ptr *QTextEdit) SetUndoRedoEnabled(enable bool) {
	defer qt.Recovering("QTextEdit::setUndoRedoEnabled")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetUndoRedoEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTextEdit) ConnectShowEvent(f func(v *gui.QShowEvent)) {
	defer qt.Recovering("connect QTextEdit::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QTextEdit) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QTextEdit::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQTextEditShowEvent
func callbackQTextEditShowEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(v))
	} else {
		NewQTextEditFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(v))
	}
}

func (ptr *QTextEdit) ShowEvent(v gui.QShowEvent_ITF) {
	defer qt.Recovering("QTextEdit::showEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(v))
	}
}

func (ptr *QTextEdit) ShowEventDefault(v gui.QShowEvent_ITF) {
	defer qt.Recovering("QTextEdit::showEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(v))
	}
}

func (ptr *QTextEdit) TextBackgroundColor() *gui.QColor {
	defer qt.Recovering("QTextEdit::textBackgroundColor")

	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QTextEdit_TextBackgroundColor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextEdit) ConnectTextChanged(f func()) {
	defer qt.Recovering("connect QTextEdit::textChanged")

	if ptr.Pointer() != nil {
		C.QTextEdit_ConnectTextChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "textChanged", f)
	}
}

func (ptr *QTextEdit) DisconnectTextChanged() {
	defer qt.Recovering("disconnect QTextEdit::textChanged")

	if ptr.Pointer() != nil {
		C.QTextEdit_DisconnectTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "textChanged")
	}
}

//export callbackQTextEditTextChanged
func callbackQTextEditTextChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QTextEdit::textChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "textChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QTextEdit) TextChanged() {
	defer qt.Recovering("QTextEdit::textChanged")

	if ptr.Pointer() != nil {
		C.QTextEdit_TextChanged(ptr.Pointer())
	}
}

func (ptr *QTextEdit) TextColor() *gui.QColor {
	defer qt.Recovering("QTextEdit::textColor")

	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QTextEdit_TextColor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextEdit) ToPlainText() string {
	defer qt.Recovering("QTextEdit::toPlainText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextEdit_ToPlainText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextEdit) Undo() {
	defer qt.Recovering("QTextEdit::undo")

	if ptr.Pointer() != nil {
		C.QTextEdit_Undo(ptr.Pointer())
	}
}

func (ptr *QTextEdit) ConnectUndoAvailable(f func(available bool)) {
	defer qt.Recovering("connect QTextEdit::undoAvailable")

	if ptr.Pointer() != nil {
		C.QTextEdit_ConnectUndoAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "undoAvailable", f)
	}
}

func (ptr *QTextEdit) DisconnectUndoAvailable() {
	defer qt.Recovering("disconnect QTextEdit::undoAvailable")

	if ptr.Pointer() != nil {
		C.QTextEdit_DisconnectUndoAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "undoAvailable")
	}
}

//export callbackQTextEditUndoAvailable
func callbackQTextEditUndoAvailable(ptr unsafe.Pointer, ptrName *C.char, available C.int) {
	defer qt.Recovering("callback QTextEdit::undoAvailable")

	if signal := qt.GetSignal(C.GoString(ptrName), "undoAvailable"); signal != nil {
		signal.(func(bool))(int(available) != 0)
	}

}

func (ptr *QTextEdit) UndoAvailable(available bool) {
	defer qt.Recovering("QTextEdit::undoAvailable")

	if ptr.Pointer() != nil {
		C.QTextEdit_UndoAvailable(ptr.Pointer(), C.int(qt.GoBoolToInt(available)))
	}
}

func (ptr *QTextEdit) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QTextEdit::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QTextEdit) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QTextEdit::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQTextEditWheelEvent
func callbackQTextEditWheelEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQTextEditFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QTextEdit) WheelEvent(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QTextEdit::wheelEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QTextEdit) WheelEventDefault(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QTextEdit::wheelEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QTextEdit) DestroyQTextEdit() {
	defer qt.Recovering("QTextEdit::~QTextEdit")

	if ptr.Pointer() != nil {
		C.QTextEdit_DestroyQTextEdit(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTextEdit) ConnectSetupViewport(f func(viewport *QWidget)) {
	defer qt.Recovering("connect QTextEdit::setupViewport")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setupViewport", f)
	}
}

func (ptr *QTextEdit) DisconnectSetupViewport() {
	defer qt.Recovering("disconnect QTextEdit::setupViewport")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setupViewport")
	}
}

//export callbackQTextEditSetupViewport
func callbackQTextEditSetupViewport(ptr unsafe.Pointer, ptrName *C.char, viewport unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::setupViewport")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupViewport"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(viewport))
	} else {
		NewQTextEditFromPointer(ptr).SetupViewportDefault(NewQWidgetFromPointer(viewport))
	}
}

func (ptr *QTextEdit) SetupViewport(viewport QWidget_ITF) {
	defer qt.Recovering("QTextEdit::setupViewport")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetupViewport(ptr.Pointer(), PointerFromQWidget(viewport))
	}
}

func (ptr *QTextEdit) SetupViewportDefault(viewport QWidget_ITF) {
	defer qt.Recovering("QTextEdit::setupViewport")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetupViewportDefault(ptr.Pointer(), PointerFromQWidget(viewport))
	}
}

func (ptr *QTextEdit) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QTextEdit::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QTextEdit) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QTextEdit::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQTextEditActionEvent
func callbackQTextEditActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQTextEditFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QTextEdit) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QTextEdit::actionEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QTextEdit) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QTextEdit::actionEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QTextEdit) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTextEdit::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QTextEdit) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QTextEdit::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQTextEditEnterEvent
func callbackQTextEditEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTextEditFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTextEdit) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTextEdit::enterEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTextEdit) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTextEdit::enterEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTextEdit) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QTextEdit::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QTextEdit) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QTextEdit::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQTextEditHideEvent
func callbackQTextEditHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQTextEditFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QTextEdit) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QTextEdit::hideEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QTextEdit) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QTextEdit::hideEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QTextEdit) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTextEdit::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QTextEdit) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QTextEdit::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQTextEditLeaveEvent
func callbackQTextEditLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTextEditFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTextEdit) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTextEdit::leaveEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTextEdit) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTextEdit::leaveEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTextEdit) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QTextEdit::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QTextEdit) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QTextEdit::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQTextEditMoveEvent
func callbackQTextEditMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQTextEditFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QTextEdit) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QTextEdit::moveEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QTextEdit) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QTextEdit::moveEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QTextEdit) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QTextEdit::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QTextEdit) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QTextEdit::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQTextEditSetVisible
func callbackQTextEditSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QTextEdit::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QTextEdit) SetVisible(visible bool) {
	defer qt.Recovering("QTextEdit::setVisible")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QTextEdit) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QTextEdit::setVisible")

	if ptr.Pointer() != nil {
		C.QTextEdit_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QTextEdit) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QTextEdit::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QTextEdit) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QTextEdit::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQTextEditCloseEvent
func callbackQTextEditCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQTextEditFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QTextEdit) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QTextEdit::closeEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QTextEdit) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QTextEdit::closeEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QTextEdit) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QTextEdit::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QTextEdit) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QTextEdit::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQTextEditInitPainter
func callbackQTextEditInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQTextEditFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QTextEdit) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QTextEdit::initPainter")

	if ptr.Pointer() != nil {
		C.QTextEdit_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QTextEdit) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QTextEdit::initPainter")

	if ptr.Pointer() != nil {
		C.QTextEdit_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QTextEdit) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QTextEdit::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QTextEdit) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QTextEdit::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQTextEditTabletEvent
func callbackQTextEditTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQTextEditFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QTextEdit) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QTextEdit::tabletEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QTextEdit) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QTextEdit::tabletEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QTextEdit) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QTextEdit::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QTextEdit) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTextEdit::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQTextEditTimerEvent
func callbackQTextEditTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTextEditFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTextEdit) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTextEdit::timerEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTextEdit) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTextEdit::timerEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTextEdit) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QTextEdit::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QTextEdit) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QTextEdit::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQTextEditChildEvent
func callbackQTextEditChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTextEditFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTextEdit) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTextEdit::childEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTextEdit) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTextEdit::childEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTextEdit) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTextEdit::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QTextEdit) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QTextEdit::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQTextEditCustomEvent
func callbackQTextEditCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTextEdit::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTextEditFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTextEdit) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTextEdit::customEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTextEdit) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTextEdit::customEvent")

	if ptr.Pointer() != nil {
		C.QTextEdit_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
