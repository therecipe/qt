package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTextDocument struct {
	core.QObject
}

type QTextDocument_ITF interface {
	core.QObject_ITF
	QTextDocument_PTR() *QTextDocument
}

func PointerFromQTextDocument(ptr QTextDocument_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextDocument_PTR().Pointer()
	}
	return nil
}

func NewQTextDocumentFromPointer(ptr unsafe.Pointer) *QTextDocument {
	var n = new(QTextDocument)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTextDocument_") {
		n.SetObjectName("QTextDocument_" + qt.Identifier())
	}
	return n
}

func (ptr *QTextDocument) QTextDocument_PTR() *QTextDocument {
	return ptr
}

//QTextDocument::FindFlag
type QTextDocument__FindFlag int64

const (
	QTextDocument__FindBackward        = QTextDocument__FindFlag(0x00001)
	QTextDocument__FindCaseSensitively = QTextDocument__FindFlag(0x00002)
	QTextDocument__FindWholeWords      = QTextDocument__FindFlag(0x00004)
)

//QTextDocument::MetaInformation
type QTextDocument__MetaInformation int64

const (
	QTextDocument__DocumentTitle = QTextDocument__MetaInformation(0)
	QTextDocument__DocumentUrl   = QTextDocument__MetaInformation(1)
)

//QTextDocument::ResourceType
type QTextDocument__ResourceType int64

const (
	QTextDocument__HtmlResource       = QTextDocument__ResourceType(1)
	QTextDocument__ImageResource      = QTextDocument__ResourceType(2)
	QTextDocument__StyleSheetResource = QTextDocument__ResourceType(3)
	QTextDocument__UserResource       = QTextDocument__ResourceType(100)
)

//QTextDocument::Stacks
type QTextDocument__Stacks int64

const (
	QTextDocument__UndoStack         = QTextDocument__Stacks(0x01)
	QTextDocument__RedoStack         = QTextDocument__Stacks(0x02)
	QTextDocument__UndoAndRedoStacks = QTextDocument__Stacks(QTextDocument__UndoStack | QTextDocument__RedoStack)
)

func (ptr *QTextDocument) BaseUrl() *core.QUrl {
	defer qt.Recovering("QTextDocument::baseUrl")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QTextDocument_BaseUrl(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextDocument) BlockCount() int {
	defer qt.Recovering("QTextDocument::blockCount")

	if ptr.Pointer() != nil {
		return int(C.QTextDocument_BlockCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextDocument) DefaultStyleSheet() string {
	defer qt.Recovering("QTextDocument::defaultStyleSheet")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextDocument_DefaultStyleSheet(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextDocument) DocumentMargin() float64 {
	defer qt.Recovering("QTextDocument::documentMargin")

	if ptr.Pointer() != nil {
		return float64(C.QTextDocument_DocumentMargin(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextDocument) IndentWidth() float64 {
	defer qt.Recovering("QTextDocument::indentWidth")

	if ptr.Pointer() != nil {
		return float64(C.QTextDocument_IndentWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextDocument) IsModified() bool {
	defer qt.Recovering("QTextDocument::isModified")

	if ptr.Pointer() != nil {
		return C.QTextDocument_IsModified(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextDocument) IsUndoRedoEnabled() bool {
	defer qt.Recovering("QTextDocument::isUndoRedoEnabled")

	if ptr.Pointer() != nil {
		return C.QTextDocument_IsUndoRedoEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextDocument) MarkContentsDirty(position int, length int) {
	defer qt.Recovering("QTextDocument::markContentsDirty")

	if ptr.Pointer() != nil {
		C.QTextDocument_MarkContentsDirty(ptr.Pointer(), C.int(position), C.int(length))
	}
}

func (ptr *QTextDocument) MaximumBlockCount() int {
	defer qt.Recovering("QTextDocument::maximumBlockCount")

	if ptr.Pointer() != nil {
		return int(C.QTextDocument_MaximumBlockCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextDocument) SetBaseUrl(url core.QUrl_ITF) {
	defer qt.Recovering("QTextDocument::setBaseUrl")

	if ptr.Pointer() != nil {
		C.QTextDocument_SetBaseUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QTextDocument) SetDefaultStyleSheet(sheet string) {
	defer qt.Recovering("QTextDocument::setDefaultStyleSheet")

	if ptr.Pointer() != nil {
		C.QTextDocument_SetDefaultStyleSheet(ptr.Pointer(), C.CString(sheet))
	}
}

func (ptr *QTextDocument) SetDocumentMargin(margin float64) {
	defer qt.Recovering("QTextDocument::setDocumentMargin")

	if ptr.Pointer() != nil {
		C.QTextDocument_SetDocumentMargin(ptr.Pointer(), C.double(margin))
	}
}

func (ptr *QTextDocument) SetMaximumBlockCount(maximum int) {
	defer qt.Recovering("QTextDocument::setMaximumBlockCount")

	if ptr.Pointer() != nil {
		C.QTextDocument_SetMaximumBlockCount(ptr.Pointer(), C.int(maximum))
	}
}

func (ptr *QTextDocument) SetModified(m bool) {
	defer qt.Recovering("QTextDocument::setModified")

	if ptr.Pointer() != nil {
		C.QTextDocument_SetModified(ptr.Pointer(), C.int(qt.GoBoolToInt(m)))
	}
}

func (ptr *QTextDocument) SetPageSize(size core.QSizeF_ITF) {
	defer qt.Recovering("QTextDocument::setPageSize")

	if ptr.Pointer() != nil {
		C.QTextDocument_SetPageSize(ptr.Pointer(), core.PointerFromQSizeF(size))
	}
}

func (ptr *QTextDocument) SetTextWidth(width float64) {
	defer qt.Recovering("QTextDocument::setTextWidth")

	if ptr.Pointer() != nil {
		C.QTextDocument_SetTextWidth(ptr.Pointer(), C.double(width))
	}
}

func (ptr *QTextDocument) SetUndoRedoEnabled(enable bool) {
	defer qt.Recovering("QTextDocument::setUndoRedoEnabled")

	if ptr.Pointer() != nil {
		C.QTextDocument_SetUndoRedoEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTextDocument) SetUseDesignMetrics(b bool) {
	defer qt.Recovering("QTextDocument::setUseDesignMetrics")

	if ptr.Pointer() != nil {
		C.QTextDocument_SetUseDesignMetrics(ptr.Pointer(), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QTextDocument) TextWidth() float64 {
	defer qt.Recovering("QTextDocument::textWidth")

	if ptr.Pointer() != nil {
		return float64(C.QTextDocument_TextWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextDocument) UseDesignMetrics() bool {
	defer qt.Recovering("QTextDocument::useDesignMetrics")

	if ptr.Pointer() != nil {
		return C.QTextDocument_UseDesignMetrics(ptr.Pointer()) != 0
	}
	return false
}

func NewQTextDocument(parent core.QObject_ITF) *QTextDocument {
	defer qt.Recovering("QTextDocument::QTextDocument")

	return NewQTextDocumentFromPointer(C.QTextDocument_NewQTextDocument(core.PointerFromQObject(parent)))
}

func NewQTextDocument2(text string, parent core.QObject_ITF) *QTextDocument {
	defer qt.Recovering("QTextDocument::QTextDocument")

	return NewQTextDocumentFromPointer(C.QTextDocument_NewQTextDocument2(C.CString(text), core.PointerFromQObject(parent)))
}

func (ptr *QTextDocument) AddResource(ty int, name core.QUrl_ITF, resource core.QVariant_ITF) {
	defer qt.Recovering("QTextDocument::addResource")

	if ptr.Pointer() != nil {
		C.QTextDocument_AddResource(ptr.Pointer(), C.int(ty), core.PointerFromQUrl(name), core.PointerFromQVariant(resource))
	}
}

func (ptr *QTextDocument) AdjustSize() {
	defer qt.Recovering("QTextDocument::adjustSize")

	if ptr.Pointer() != nil {
		C.QTextDocument_AdjustSize(ptr.Pointer())
	}
}

func (ptr *QTextDocument) AvailableRedoSteps() int {
	defer qt.Recovering("QTextDocument::availableRedoSteps")

	if ptr.Pointer() != nil {
		return int(C.QTextDocument_AvailableRedoSteps(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextDocument) AvailableUndoSteps() int {
	defer qt.Recovering("QTextDocument::availableUndoSteps")

	if ptr.Pointer() != nil {
		return int(C.QTextDocument_AvailableUndoSteps(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextDocument) ConnectBaseUrlChanged(f func(url *core.QUrl)) {
	defer qt.Recovering("connect QTextDocument::baseUrlChanged")

	if ptr.Pointer() != nil {
		C.QTextDocument_ConnectBaseUrlChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "baseUrlChanged", f)
	}
}

func (ptr *QTextDocument) DisconnectBaseUrlChanged() {
	defer qt.Recovering("disconnect QTextDocument::baseUrlChanged")

	if ptr.Pointer() != nil {
		C.QTextDocument_DisconnectBaseUrlChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "baseUrlChanged")
	}
}

//export callbackQTextDocumentBaseUrlChanged
func callbackQTextDocumentBaseUrlChanged(ptrName *C.char, url unsafe.Pointer) {
	defer qt.Recovering("callback QTextDocument::baseUrlChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "baseUrlChanged"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QTextDocument) ConnectBlockCountChanged(f func(newBlockCount int)) {
	defer qt.Recovering("connect QTextDocument::blockCountChanged")

	if ptr.Pointer() != nil {
		C.QTextDocument_ConnectBlockCountChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "blockCountChanged", f)
	}
}

func (ptr *QTextDocument) DisconnectBlockCountChanged() {
	defer qt.Recovering("disconnect QTextDocument::blockCountChanged")

	if ptr.Pointer() != nil {
		C.QTextDocument_DisconnectBlockCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "blockCountChanged")
	}
}

//export callbackQTextDocumentBlockCountChanged
func callbackQTextDocumentBlockCountChanged(ptrName *C.char, newBlockCount C.int) {
	defer qt.Recovering("callback QTextDocument::blockCountChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "blockCountChanged"); signal != nil {
		signal.(func(int))(int(newBlockCount))
	}

}

func (ptr *QTextDocument) CharacterCount() int {
	defer qt.Recovering("QTextDocument::characterCount")

	if ptr.Pointer() != nil {
		return int(C.QTextDocument_CharacterCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextDocument) ConnectClear(f func()) {
	defer qt.Recovering("connect QTextDocument::clear")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "clear", f)
	}
}

func (ptr *QTextDocument) DisconnectClear() {
	defer qt.Recovering("disconnect QTextDocument::clear")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "clear")
	}
}

//export callbackQTextDocumentClear
func callbackQTextDocumentClear(ptrName *C.char) bool {
	defer qt.Recovering("callback QTextDocument::clear")

	if signal := qt.GetSignal(C.GoString(ptrName), "clear"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QTextDocument) ClearUndoRedoStacks(stacksToClear QTextDocument__Stacks) {
	defer qt.Recovering("QTextDocument::clearUndoRedoStacks")

	if ptr.Pointer() != nil {
		C.QTextDocument_ClearUndoRedoStacks(ptr.Pointer(), C.int(stacksToClear))
	}
}

func (ptr *QTextDocument) Clone(parent core.QObject_ITF) *QTextDocument {
	defer qt.Recovering("QTextDocument::clone")

	if ptr.Pointer() != nil {
		return NewQTextDocumentFromPointer(C.QTextDocument_Clone(ptr.Pointer(), core.PointerFromQObject(parent)))
	}
	return nil
}

func (ptr *QTextDocument) ConnectContentsChange(f func(position int, charsRemoved int, charsAdded int)) {
	defer qt.Recovering("connect QTextDocument::contentsChange")

	if ptr.Pointer() != nil {
		C.QTextDocument_ConnectContentsChange(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "contentsChange", f)
	}
}

func (ptr *QTextDocument) DisconnectContentsChange() {
	defer qt.Recovering("disconnect QTextDocument::contentsChange")

	if ptr.Pointer() != nil {
		C.QTextDocument_DisconnectContentsChange(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "contentsChange")
	}
}

//export callbackQTextDocumentContentsChange
func callbackQTextDocumentContentsChange(ptrName *C.char, position C.int, charsRemoved C.int, charsAdded C.int) {
	defer qt.Recovering("callback QTextDocument::contentsChange")

	if signal := qt.GetSignal(C.GoString(ptrName), "contentsChange"); signal != nil {
		signal.(func(int, int, int))(int(position), int(charsRemoved), int(charsAdded))
	}

}

func (ptr *QTextDocument) ConnectContentsChanged(f func()) {
	defer qt.Recovering("connect QTextDocument::contentsChanged")

	if ptr.Pointer() != nil {
		C.QTextDocument_ConnectContentsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "contentsChanged", f)
	}
}

func (ptr *QTextDocument) DisconnectContentsChanged() {
	defer qt.Recovering("disconnect QTextDocument::contentsChanged")

	if ptr.Pointer() != nil {
		C.QTextDocument_DisconnectContentsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "contentsChanged")
	}
}

//export callbackQTextDocumentContentsChanged
func callbackQTextDocumentContentsChanged(ptrName *C.char) {
	defer qt.Recovering("callback QTextDocument::contentsChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "contentsChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QTextDocument) DefaultCursorMoveStyle() core.Qt__CursorMoveStyle {
	defer qt.Recovering("QTextDocument::defaultCursorMoveStyle")

	if ptr.Pointer() != nil {
		return core.Qt__CursorMoveStyle(C.QTextDocument_DefaultCursorMoveStyle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextDocument) DocumentLayout() *QAbstractTextDocumentLayout {
	defer qt.Recovering("QTextDocument::documentLayout")

	if ptr.Pointer() != nil {
		return NewQAbstractTextDocumentLayoutFromPointer(C.QTextDocument_DocumentLayout(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextDocument) ConnectDocumentLayoutChanged(f func()) {
	defer qt.Recovering("connect QTextDocument::documentLayoutChanged")

	if ptr.Pointer() != nil {
		C.QTextDocument_ConnectDocumentLayoutChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "documentLayoutChanged", f)
	}
}

func (ptr *QTextDocument) DisconnectDocumentLayoutChanged() {
	defer qt.Recovering("disconnect QTextDocument::documentLayoutChanged")

	if ptr.Pointer() != nil {
		C.QTextDocument_DisconnectDocumentLayoutChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "documentLayoutChanged")
	}
}

//export callbackQTextDocumentDocumentLayoutChanged
func callbackQTextDocumentDocumentLayoutChanged(ptrName *C.char) {
	defer qt.Recovering("callback QTextDocument::documentLayoutChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "documentLayoutChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QTextDocument) DrawContents(p QPainter_ITF, rect core.QRectF_ITF) {
	defer qt.Recovering("QTextDocument::drawContents")

	if ptr.Pointer() != nil {
		C.QTextDocument_DrawContents(ptr.Pointer(), PointerFromQPainter(p), core.PointerFromQRectF(rect))
	}
}

func (ptr *QTextDocument) IdealWidth() float64 {
	defer qt.Recovering("QTextDocument::idealWidth")

	if ptr.Pointer() != nil {
		return float64(C.QTextDocument_IdealWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextDocument) IsEmpty() bool {
	defer qt.Recovering("QTextDocument::isEmpty")

	if ptr.Pointer() != nil {
		return C.QTextDocument_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextDocument) IsRedoAvailable() bool {
	defer qt.Recovering("QTextDocument::isRedoAvailable")

	if ptr.Pointer() != nil {
		return C.QTextDocument_IsRedoAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextDocument) IsUndoAvailable() bool {
	defer qt.Recovering("QTextDocument::isUndoAvailable")

	if ptr.Pointer() != nil {
		return C.QTextDocument_IsUndoAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextDocument) LineCount() int {
	defer qt.Recovering("QTextDocument::lineCount")

	if ptr.Pointer() != nil {
		return int(C.QTextDocument_LineCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextDocument) MetaInformation(info QTextDocument__MetaInformation) string {
	defer qt.Recovering("QTextDocument::metaInformation")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextDocument_MetaInformation(ptr.Pointer(), C.int(info)))
	}
	return ""
}

func (ptr *QTextDocument) ConnectModificationChanged(f func(changed bool)) {
	defer qt.Recovering("connect QTextDocument::modificationChanged")

	if ptr.Pointer() != nil {
		C.QTextDocument_ConnectModificationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "modificationChanged", f)
	}
}

func (ptr *QTextDocument) DisconnectModificationChanged() {
	defer qt.Recovering("disconnect QTextDocument::modificationChanged")

	if ptr.Pointer() != nil {
		C.QTextDocument_DisconnectModificationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "modificationChanged")
	}
}

//export callbackQTextDocumentModificationChanged
func callbackQTextDocumentModificationChanged(ptrName *C.char, changed C.int) {
	defer qt.Recovering("callback QTextDocument::modificationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "modificationChanged"); signal != nil {
		signal.(func(bool))(int(changed) != 0)
	}

}

func (ptr *QTextDocument) Object(objectIndex int) *QTextObject {
	defer qt.Recovering("QTextDocument::object")

	if ptr.Pointer() != nil {
		return NewQTextObjectFromPointer(C.QTextDocument_Object(ptr.Pointer(), C.int(objectIndex)))
	}
	return nil
}

func (ptr *QTextDocument) ObjectForFormat(f QTextFormat_ITF) *QTextObject {
	defer qt.Recovering("QTextDocument::objectForFormat")

	if ptr.Pointer() != nil {
		return NewQTextObjectFromPointer(C.QTextDocument_ObjectForFormat(ptr.Pointer(), PointerFromQTextFormat(f)))
	}
	return nil
}

func (ptr *QTextDocument) PageCount() int {
	defer qt.Recovering("QTextDocument::pageCount")

	if ptr.Pointer() != nil {
		return int(C.QTextDocument_PageCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextDocument) Print(printer QPagedPaintDevice_ITF) {
	defer qt.Recovering("QTextDocument::print")

	if ptr.Pointer() != nil {
		C.QTextDocument_Print(ptr.Pointer(), PointerFromQPagedPaintDevice(printer))
	}
}

func (ptr *QTextDocument) Redo2() {
	defer qt.Recovering("QTextDocument::redo")

	if ptr.Pointer() != nil {
		C.QTextDocument_Redo2(ptr.Pointer())
	}
}

func (ptr *QTextDocument) Redo(cursor QTextCursor_ITF) {
	defer qt.Recovering("QTextDocument::redo")

	if ptr.Pointer() != nil {
		C.QTextDocument_Redo(ptr.Pointer(), PointerFromQTextCursor(cursor))
	}
}

func (ptr *QTextDocument) ConnectRedoAvailable(f func(available bool)) {
	defer qt.Recovering("connect QTextDocument::redoAvailable")

	if ptr.Pointer() != nil {
		C.QTextDocument_ConnectRedoAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "redoAvailable", f)
	}
}

func (ptr *QTextDocument) DisconnectRedoAvailable() {
	defer qt.Recovering("disconnect QTextDocument::redoAvailable")

	if ptr.Pointer() != nil {
		C.QTextDocument_DisconnectRedoAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "redoAvailable")
	}
}

//export callbackQTextDocumentRedoAvailable
func callbackQTextDocumentRedoAvailable(ptrName *C.char, available C.int) {
	defer qt.Recovering("callback QTextDocument::redoAvailable")

	if signal := qt.GetSignal(C.GoString(ptrName), "redoAvailable"); signal != nil {
		signal.(func(bool))(int(available) != 0)
	}

}

func (ptr *QTextDocument) Resource(ty int, name core.QUrl_ITF) *core.QVariant {
	defer qt.Recovering("QTextDocument::resource")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QTextDocument_Resource(ptr.Pointer(), C.int(ty), core.PointerFromQUrl(name)))
	}
	return nil
}

func (ptr *QTextDocument) Revision() int {
	defer qt.Recovering("QTextDocument::revision")

	if ptr.Pointer() != nil {
		return int(C.QTextDocument_Revision(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextDocument) RootFrame() *QTextFrame {
	defer qt.Recovering("QTextDocument::rootFrame")

	if ptr.Pointer() != nil {
		return NewQTextFrameFromPointer(C.QTextDocument_RootFrame(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextDocument) SetDefaultCursorMoveStyle(style core.Qt__CursorMoveStyle) {
	defer qt.Recovering("QTextDocument::setDefaultCursorMoveStyle")

	if ptr.Pointer() != nil {
		C.QTextDocument_SetDefaultCursorMoveStyle(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QTextDocument) SetDefaultFont(font QFont_ITF) {
	defer qt.Recovering("QTextDocument::setDefaultFont")

	if ptr.Pointer() != nil {
		C.QTextDocument_SetDefaultFont(ptr.Pointer(), PointerFromQFont(font))
	}
}

func (ptr *QTextDocument) SetDefaultTextOption(option QTextOption_ITF) {
	defer qt.Recovering("QTextDocument::setDefaultTextOption")

	if ptr.Pointer() != nil {
		C.QTextDocument_SetDefaultTextOption(ptr.Pointer(), PointerFromQTextOption(option))
	}
}

func (ptr *QTextDocument) SetDocumentLayout(layout QAbstractTextDocumentLayout_ITF) {
	defer qt.Recovering("QTextDocument::setDocumentLayout")

	if ptr.Pointer() != nil {
		C.QTextDocument_SetDocumentLayout(ptr.Pointer(), PointerFromQAbstractTextDocumentLayout(layout))
	}
}

func (ptr *QTextDocument) SetHtml(html string) {
	defer qt.Recovering("QTextDocument::setHtml")

	if ptr.Pointer() != nil {
		C.QTextDocument_SetHtml(ptr.Pointer(), C.CString(html))
	}
}

func (ptr *QTextDocument) SetIndentWidth(width float64) {
	defer qt.Recovering("QTextDocument::setIndentWidth")

	if ptr.Pointer() != nil {
		C.QTextDocument_SetIndentWidth(ptr.Pointer(), C.double(width))
	}
}

func (ptr *QTextDocument) SetMetaInformation(info QTextDocument__MetaInformation, stri string) {
	defer qt.Recovering("QTextDocument::setMetaInformation")

	if ptr.Pointer() != nil {
		C.QTextDocument_SetMetaInformation(ptr.Pointer(), C.int(info), C.CString(stri))
	}
}

func (ptr *QTextDocument) SetPlainText(text string) {
	defer qt.Recovering("QTextDocument::setPlainText")

	if ptr.Pointer() != nil {
		C.QTextDocument_SetPlainText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QTextDocument) ToHtml(encoding core.QByteArray_ITF) string {
	defer qt.Recovering("QTextDocument::toHtml")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextDocument_ToHtml(ptr.Pointer(), core.PointerFromQByteArray(encoding)))
	}
	return ""
}

func (ptr *QTextDocument) ToPlainText() string {
	defer qt.Recovering("QTextDocument::toPlainText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextDocument_ToPlainText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextDocument) Undo2() {
	defer qt.Recovering("QTextDocument::undo")

	if ptr.Pointer() != nil {
		C.QTextDocument_Undo2(ptr.Pointer())
	}
}

func (ptr *QTextDocument) Undo(cursor QTextCursor_ITF) {
	defer qt.Recovering("QTextDocument::undo")

	if ptr.Pointer() != nil {
		C.QTextDocument_Undo(ptr.Pointer(), PointerFromQTextCursor(cursor))
	}
}

func (ptr *QTextDocument) ConnectUndoAvailable(f func(available bool)) {
	defer qt.Recovering("connect QTextDocument::undoAvailable")

	if ptr.Pointer() != nil {
		C.QTextDocument_ConnectUndoAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "undoAvailable", f)
	}
}

func (ptr *QTextDocument) DisconnectUndoAvailable() {
	defer qt.Recovering("disconnect QTextDocument::undoAvailable")

	if ptr.Pointer() != nil {
		C.QTextDocument_DisconnectUndoAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "undoAvailable")
	}
}

//export callbackQTextDocumentUndoAvailable
func callbackQTextDocumentUndoAvailable(ptrName *C.char, available C.int) {
	defer qt.Recovering("callback QTextDocument::undoAvailable")

	if signal := qt.GetSignal(C.GoString(ptrName), "undoAvailable"); signal != nil {
		signal.(func(bool))(int(available) != 0)
	}

}

func (ptr *QTextDocument) ConnectUndoCommandAdded(f func()) {
	defer qt.Recovering("connect QTextDocument::undoCommandAdded")

	if ptr.Pointer() != nil {
		C.QTextDocument_ConnectUndoCommandAdded(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "undoCommandAdded", f)
	}
}

func (ptr *QTextDocument) DisconnectUndoCommandAdded() {
	defer qt.Recovering("disconnect QTextDocument::undoCommandAdded")

	if ptr.Pointer() != nil {
		C.QTextDocument_DisconnectUndoCommandAdded(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "undoCommandAdded")
	}
}

//export callbackQTextDocumentUndoCommandAdded
func callbackQTextDocumentUndoCommandAdded(ptrName *C.char) {
	defer qt.Recovering("callback QTextDocument::undoCommandAdded")

	if signal := qt.GetSignal(C.GoString(ptrName), "undoCommandAdded"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QTextDocument) DestroyQTextDocument() {
	defer qt.Recovering("QTextDocument::~QTextDocument")

	if ptr.Pointer() != nil {
		C.QTextDocument_DestroyQTextDocument(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTextDocument) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QTextDocument::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QTextDocument) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTextDocument::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQTextDocumentTimerEvent
func callbackQTextDocumentTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextDocument::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTextDocument) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QTextDocument::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QTextDocument) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QTextDocument::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQTextDocumentChildEvent
func callbackQTextDocumentChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextDocument::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTextDocument) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTextDocument::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QTextDocument) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QTextDocument::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQTextDocumentCustomEvent
func callbackQTextDocumentCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextDocument::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
