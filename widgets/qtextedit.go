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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::acceptRichText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextEdit_AcceptRichText(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextEdit) AutoFormatting() QTextEdit__AutoFormattingFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::autoFormatting")
		}
	}()

	if ptr.Pointer() != nil {
		return QTextEdit__AutoFormattingFlag(C.QTextEdit_AutoFormatting(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) CursorWidth() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::cursorWidth")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextEdit_CursorWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) Document() *gui.QTextDocument {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::document")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQTextDocumentFromPointer(C.QTextEdit_Document(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextEdit) IsReadOnly() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::isReadOnly")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextEdit_IsReadOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextEdit) LineWrapColumnOrWidth() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::lineWrapColumnOrWidth")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextEdit_LineWrapColumnOrWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) LineWrapMode() QTextEdit__LineWrapMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::lineWrapMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QTextEdit__LineWrapMode(C.QTextEdit_LineWrapMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) OverwriteMode() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::overwriteMode")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextEdit_OverwriteMode(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextEdit) PlaceholderText() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::placeholderText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextEdit_PlaceholderText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextEdit) Redo() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::redo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_Redo(ptr.Pointer())
	}
}

func (ptr *QTextEdit) SetAcceptRichText(accept bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::setAcceptRichText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_SetAcceptRichText(ptr.Pointer(), C.int(qt.GoBoolToInt(accept)))
	}
}

func (ptr *QTextEdit) SetAutoFormatting(features QTextEdit__AutoFormattingFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::setAutoFormatting")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_SetAutoFormatting(ptr.Pointer(), C.int(features))
	}
}

func (ptr *QTextEdit) SetCursorWidth(width int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::setCursorWidth")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_SetCursorWidth(ptr.Pointer(), C.int(width))
	}
}

func (ptr *QTextEdit) SetDocument(document gui.QTextDocument_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::setDocument")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_SetDocument(ptr.Pointer(), gui.PointerFromQTextDocument(document))
	}
}

func (ptr *QTextEdit) SetFontWeight(weight int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::setFontWeight")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_SetFontWeight(ptr.Pointer(), C.int(weight))
	}
}

func (ptr *QTextEdit) SetHtml(text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::setHtml")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_SetHtml(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QTextEdit) SetLineWrapColumnOrWidth(w int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::setLineWrapColumnOrWidth")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_SetLineWrapColumnOrWidth(ptr.Pointer(), C.int(w))
	}
}

func (ptr *QTextEdit) SetLineWrapMode(mode QTextEdit__LineWrapMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::setLineWrapMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_SetLineWrapMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QTextEdit) SetOverwriteMode(overwrite bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::setOverwriteMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_SetOverwriteMode(ptr.Pointer(), C.int(qt.GoBoolToInt(overwrite)))
	}
}

func (ptr *QTextEdit) SetPlaceholderText(placeholderText string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::setPlaceholderText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_SetPlaceholderText(ptr.Pointer(), C.CString(placeholderText))
	}
}

func (ptr *QTextEdit) SetReadOnly(ro bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::setReadOnly")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_SetReadOnly(ptr.Pointer(), C.int(qt.GoBoolToInt(ro)))
	}
}

func (ptr *QTextEdit) SetTabChangesFocus(b bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::setTabChangesFocus")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_SetTabChangesFocus(ptr.Pointer(), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QTextEdit) SetTabStopWidth(width int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::setTabStopWidth")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_SetTabStopWidth(ptr.Pointer(), C.int(width))
	}
}

func (ptr *QTextEdit) SetTextInteractionFlags(flags core.Qt__TextInteractionFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::setTextInteractionFlags")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_SetTextInteractionFlags(ptr.Pointer(), C.int(flags))
	}
}

func (ptr *QTextEdit) SetWordWrapMode(policy gui.QTextOption__WrapMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::setWordWrapMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_SetWordWrapMode(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QTextEdit) TabChangesFocus() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::tabChangesFocus")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextEdit_TabChangesFocus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextEdit) TabStopWidth() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::tabStopWidth")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextEdit_TabStopWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) TextInteractionFlags() core.Qt__TextInteractionFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::textInteractionFlags")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__TextInteractionFlag(C.QTextEdit_TextInteractionFlags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) ToHtml() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::toHtml")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextEdit_ToHtml(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextEdit) WordWrapMode() gui.QTextOption__WrapMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::wordWrapMode")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.QTextOption__WrapMode(C.QTextEdit_WordWrapMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) ZoomIn(ran int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::zoomIn")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_ZoomIn(ptr.Pointer(), C.int(ran))
	}
}

func (ptr *QTextEdit) ZoomOut(ran int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::zoomOut")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_ZoomOut(ptr.Pointer(), C.int(ran))
	}
}

func NewQTextEdit(parent QWidget_ITF) *QTextEdit {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::QTextEdit")
		}
	}()

	return NewQTextEditFromPointer(C.QTextEdit_NewQTextEdit(PointerFromQWidget(parent)))
}

func NewQTextEdit2(text string, parent QWidget_ITF) *QTextEdit {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::QTextEdit")
		}
	}()

	return NewQTextEditFromPointer(C.QTextEdit_NewQTextEdit2(C.CString(text), PointerFromQWidget(parent)))
}

func (ptr *QTextEdit) Alignment() core.Qt__AlignmentFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::alignment")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QTextEdit_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) AnchorAt(pos core.QPoint_ITF) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::anchorAt")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextEdit_AnchorAt(ptr.Pointer(), core.PointerFromQPoint(pos)))
	}
	return ""
}

func (ptr *QTextEdit) Append(text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::append")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_Append(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QTextEdit) CanPaste() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::canPaste")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextEdit_CanPaste(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextEdit) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_Clear(ptr.Pointer())
	}
}

func (ptr *QTextEdit) Copy() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::copy")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_Copy(ptr.Pointer())
	}
}

func (ptr *QTextEdit) ConnectCopyAvailable(f func(yes bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::copyAvailable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_ConnectCopyAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "copyAvailable", f)
	}
}

func (ptr *QTextEdit) DisconnectCopyAvailable() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::copyAvailable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_DisconnectCopyAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "copyAvailable")
	}
}

//export callbackQTextEditCopyAvailable
func callbackQTextEditCopyAvailable(ptrName *C.char, yes C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::copyAvailable")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "copyAvailable").(func(bool))(int(yes) != 0)
}

func (ptr *QTextEdit) CreateStandardContextMenu() *QMenu {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::createStandardContextMenu")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QTextEdit_CreateStandardContextMenu(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextEdit) CreateStandardContextMenu2(position core.QPoint_ITF) *QMenu {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::createStandardContextMenu")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QTextEdit_CreateStandardContextMenu2(ptr.Pointer(), core.PointerFromQPoint(position)))
	}
	return nil
}

func (ptr *QTextEdit) ConnectCursorPositionChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::cursorPositionChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_ConnectCursorPositionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cursorPositionChanged", f)
	}
}

func (ptr *QTextEdit) DisconnectCursorPositionChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::cursorPositionChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_DisconnectCursorPositionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cursorPositionChanged")
	}
}

//export callbackQTextEditCursorPositionChanged
func callbackQTextEditCursorPositionChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::cursorPositionChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "cursorPositionChanged").(func())()
}

func (ptr *QTextEdit) Cut() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::cut")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_Cut(ptr.Pointer())
	}
}

func (ptr *QTextEdit) DocumentTitle() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::documentTitle")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextEdit_DocumentTitle(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextEdit) EnsureCursorVisible() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::ensureCursorVisible")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_EnsureCursorVisible(ptr.Pointer())
	}
}

func (ptr *QTextEdit) FontFamily() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::fontFamily")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextEdit_FontFamily(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextEdit) FontItalic() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::fontItalic")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextEdit_FontItalic(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextEdit) FontPointSize() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::fontPointSize")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextEdit_FontPointSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) FontUnderline() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::fontUnderline")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextEdit_FontUnderline(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextEdit) FontWeight() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::fontWeight")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextEdit_FontWeight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) InputMethodQuery(property core.Qt__InputMethodQuery) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::inputMethodQuery")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QTextEdit_InputMethodQuery(ptr.Pointer(), C.int(property)))
	}
	return nil
}

func (ptr *QTextEdit) InsertHtml(text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::insertHtml")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_InsertHtml(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QTextEdit) InsertPlainText(text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::insertPlainText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_InsertPlainText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QTextEdit) IsUndoRedoEnabled() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::isUndoRedoEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextEdit_IsUndoRedoEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextEdit) LoadResource(ty int, name core.QUrl_ITF) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::loadResource")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QTextEdit_LoadResource(ptr.Pointer(), C.int(ty), core.PointerFromQUrl(name)))
	}
	return nil
}

func (ptr *QTextEdit) MergeCurrentCharFormat(modifier gui.QTextCharFormat_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::mergeCurrentCharFormat")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_MergeCurrentCharFormat(ptr.Pointer(), gui.PointerFromQTextCharFormat(modifier))
	}
}

func (ptr *QTextEdit) MoveCursor(operation gui.QTextCursor__MoveOperation, mode gui.QTextCursor__MoveMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::moveCursor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_MoveCursor(ptr.Pointer(), C.int(operation), C.int(mode))
	}
}

func (ptr *QTextEdit) Paste() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::paste")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_Paste(ptr.Pointer())
	}
}

func (ptr *QTextEdit) Print(printer gui.QPagedPaintDevice_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::print")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_Print(ptr.Pointer(), gui.PointerFromQPagedPaintDevice(printer))
	}
}

func (ptr *QTextEdit) ConnectRedoAvailable(f func(available bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::redoAvailable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_ConnectRedoAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "redoAvailable", f)
	}
}

func (ptr *QTextEdit) DisconnectRedoAvailable() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::redoAvailable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_DisconnectRedoAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "redoAvailable")
	}
}

//export callbackQTextEditRedoAvailable
func callbackQTextEditRedoAvailable(ptrName *C.char, available C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::redoAvailable")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "redoAvailable").(func(bool))(int(available) != 0)
}

func (ptr *QTextEdit) ScrollToAnchor(name string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::scrollToAnchor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_ScrollToAnchor(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QTextEdit) SelectAll() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::selectAll")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_SelectAll(ptr.Pointer())
	}
}

func (ptr *QTextEdit) ConnectSelectionChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::selectionChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_ConnectSelectionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "selectionChanged", f)
	}
}

func (ptr *QTextEdit) DisconnectSelectionChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::selectionChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_DisconnectSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "selectionChanged")
	}
}

//export callbackQTextEditSelectionChanged
func callbackQTextEditSelectionChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::selectionChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "selectionChanged").(func())()
}

func (ptr *QTextEdit) SetAlignment(a core.Qt__AlignmentFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::setAlignment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_SetAlignment(ptr.Pointer(), C.int(a))
	}
}

func (ptr *QTextEdit) SetCurrentCharFormat(format gui.QTextCharFormat_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::setCurrentCharFormat")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_SetCurrentCharFormat(ptr.Pointer(), gui.PointerFromQTextCharFormat(format))
	}
}

func (ptr *QTextEdit) SetCurrentFont(f gui.QFont_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::setCurrentFont")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_SetCurrentFont(ptr.Pointer(), gui.PointerFromQFont(f))
	}
}

func (ptr *QTextEdit) SetDocumentTitle(title string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::setDocumentTitle")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_SetDocumentTitle(ptr.Pointer(), C.CString(title))
	}
}

func (ptr *QTextEdit) SetFontFamily(fontFamily string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::setFontFamily")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_SetFontFamily(ptr.Pointer(), C.CString(fontFamily))
	}
}

func (ptr *QTextEdit) SetFontItalic(italic bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::setFontItalic")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_SetFontItalic(ptr.Pointer(), C.int(qt.GoBoolToInt(italic)))
	}
}

func (ptr *QTextEdit) SetFontPointSize(s float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::setFontPointSize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_SetFontPointSize(ptr.Pointer(), C.double(s))
	}
}

func (ptr *QTextEdit) SetFontUnderline(underline bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::setFontUnderline")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_SetFontUnderline(ptr.Pointer(), C.int(qt.GoBoolToInt(underline)))
	}
}

func (ptr *QTextEdit) SetPlainText(text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::setPlainText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_SetPlainText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QTextEdit) SetText(text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::setText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_SetText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QTextEdit) SetTextBackgroundColor(c gui.QColor_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::setTextBackgroundColor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_SetTextBackgroundColor(ptr.Pointer(), gui.PointerFromQColor(c))
	}
}

func (ptr *QTextEdit) SetTextColor(c gui.QColor_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::setTextColor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_SetTextColor(ptr.Pointer(), gui.PointerFromQColor(c))
	}
}

func (ptr *QTextEdit) SetTextCursor(cursor gui.QTextCursor_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::setTextCursor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_SetTextCursor(ptr.Pointer(), gui.PointerFromQTextCursor(cursor))
	}
}

func (ptr *QTextEdit) SetUndoRedoEnabled(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::setUndoRedoEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_SetUndoRedoEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTextEdit) TextBackgroundColor() *gui.QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::textBackgroundColor")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QTextEdit_TextBackgroundColor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextEdit) ConnectTextChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::textChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_ConnectTextChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "textChanged", f)
	}
}

func (ptr *QTextEdit) DisconnectTextChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::textChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_DisconnectTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "textChanged")
	}
}

//export callbackQTextEditTextChanged
func callbackQTextEditTextChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::textChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "textChanged").(func())()
}

func (ptr *QTextEdit) TextColor() *gui.QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::textColor")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QTextEdit_TextColor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextEdit) ToPlainText() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::toPlainText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextEdit_ToPlainText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextEdit) Undo() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::undo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_Undo(ptr.Pointer())
	}
}

func (ptr *QTextEdit) ConnectUndoAvailable(f func(available bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::undoAvailable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_ConnectUndoAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "undoAvailable", f)
	}
}

func (ptr *QTextEdit) DisconnectUndoAvailable() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::undoAvailable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_DisconnectUndoAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "undoAvailable")
	}
}

//export callbackQTextEditUndoAvailable
func callbackQTextEditUndoAvailable(ptrName *C.char, available C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::undoAvailable")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "undoAvailable").(func(bool))(int(available) != 0)
}

func (ptr *QTextEdit) DestroyQTextEdit() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEdit::~QTextEdit")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEdit_DestroyQTextEdit(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
