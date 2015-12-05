package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
		n.SetObjectName("QTextDocument_" + qt.RandomIdentifier())
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

func (ptr *QTextDocument) BlockCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::blockCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextDocument_BlockCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextDocument) DefaultStyleSheet() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::defaultStyleSheet")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextDocument_DefaultStyleSheet(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextDocument) DocumentMargin() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::documentMargin")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextDocument_DocumentMargin(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextDocument) IndentWidth() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::indentWidth")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextDocument_IndentWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextDocument) IsModified() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::isModified")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextDocument_IsModified(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextDocument) IsUndoRedoEnabled() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::isUndoRedoEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextDocument_IsUndoRedoEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextDocument) MarkContentsDirty(position int, length int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::markContentsDirty")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_MarkContentsDirty(ptr.Pointer(), C.int(position), C.int(length))
	}
}

func (ptr *QTextDocument) MaximumBlockCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::maximumBlockCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextDocument_MaximumBlockCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextDocument) SetBaseUrl(url core.QUrl_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::setBaseUrl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_SetBaseUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QTextDocument) SetDefaultStyleSheet(sheet string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::setDefaultStyleSheet")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_SetDefaultStyleSheet(ptr.Pointer(), C.CString(sheet))
	}
}

func (ptr *QTextDocument) SetDocumentMargin(margin float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::setDocumentMargin")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_SetDocumentMargin(ptr.Pointer(), C.double(margin))
	}
}

func (ptr *QTextDocument) SetMaximumBlockCount(maximum int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::setMaximumBlockCount")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_SetMaximumBlockCount(ptr.Pointer(), C.int(maximum))
	}
}

func (ptr *QTextDocument) SetModified(m bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::setModified")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_SetModified(ptr.Pointer(), C.int(qt.GoBoolToInt(m)))
	}
}

func (ptr *QTextDocument) SetPageSize(size core.QSizeF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::setPageSize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_SetPageSize(ptr.Pointer(), core.PointerFromQSizeF(size))
	}
}

func (ptr *QTextDocument) SetTextWidth(width float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::setTextWidth")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_SetTextWidth(ptr.Pointer(), C.double(width))
	}
}

func (ptr *QTextDocument) SetUndoRedoEnabled(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::setUndoRedoEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_SetUndoRedoEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTextDocument) SetUseDesignMetrics(b bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::setUseDesignMetrics")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_SetUseDesignMetrics(ptr.Pointer(), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QTextDocument) TextWidth() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::textWidth")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextDocument_TextWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextDocument) UseDesignMetrics() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::useDesignMetrics")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextDocument_UseDesignMetrics(ptr.Pointer()) != 0
	}
	return false
}

func NewQTextDocument(parent core.QObject_ITF) *QTextDocument {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::QTextDocument")
		}
	}()

	return NewQTextDocumentFromPointer(C.QTextDocument_NewQTextDocument(core.PointerFromQObject(parent)))
}

func NewQTextDocument2(text string, parent core.QObject_ITF) *QTextDocument {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::QTextDocument")
		}
	}()

	return NewQTextDocumentFromPointer(C.QTextDocument_NewQTextDocument2(C.CString(text), core.PointerFromQObject(parent)))
}

func (ptr *QTextDocument) AddResource(ty int, name core.QUrl_ITF, resource core.QVariant_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::addResource")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_AddResource(ptr.Pointer(), C.int(ty), core.PointerFromQUrl(name), core.PointerFromQVariant(resource))
	}
}

func (ptr *QTextDocument) AdjustSize() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::adjustSize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_AdjustSize(ptr.Pointer())
	}
}

func (ptr *QTextDocument) AvailableRedoSteps() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::availableRedoSteps")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextDocument_AvailableRedoSteps(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextDocument) AvailableUndoSteps() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::availableUndoSteps")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextDocument_AvailableUndoSteps(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextDocument) ConnectBlockCountChanged(f func(newBlockCount int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::blockCountChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_ConnectBlockCountChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "blockCountChanged", f)
	}
}

func (ptr *QTextDocument) DisconnectBlockCountChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::blockCountChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_DisconnectBlockCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "blockCountChanged")
	}
}

//export callbackQTextDocumentBlockCountChanged
func callbackQTextDocumentBlockCountChanged(ptrName *C.char, newBlockCount C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::blockCountChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "blockCountChanged").(func(int))(int(newBlockCount))
}

func (ptr *QTextDocument) CharacterCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::characterCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextDocument_CharacterCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextDocument) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_Clear(ptr.Pointer())
	}
}

func (ptr *QTextDocument) ClearUndoRedoStacks(stacksToClear QTextDocument__Stacks) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::clearUndoRedoStacks")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_ClearUndoRedoStacks(ptr.Pointer(), C.int(stacksToClear))
	}
}

func (ptr *QTextDocument) Clone(parent core.QObject_ITF) *QTextDocument {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::clone")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTextDocumentFromPointer(C.QTextDocument_Clone(ptr.Pointer(), core.PointerFromQObject(parent)))
	}
	return nil
}

func (ptr *QTextDocument) ConnectContentsChange(f func(position int, charsRemoved int, charsAdded int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::contentsChange")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_ConnectContentsChange(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "contentsChange", f)
	}
}

func (ptr *QTextDocument) DisconnectContentsChange() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::contentsChange")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_DisconnectContentsChange(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "contentsChange")
	}
}

//export callbackQTextDocumentContentsChange
func callbackQTextDocumentContentsChange(ptrName *C.char, position C.int, charsRemoved C.int, charsAdded C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::contentsChange")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "contentsChange").(func(int, int, int))(int(position), int(charsRemoved), int(charsAdded))
}

func (ptr *QTextDocument) ConnectContentsChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::contentsChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_ConnectContentsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "contentsChanged", f)
	}
}

func (ptr *QTextDocument) DisconnectContentsChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::contentsChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_DisconnectContentsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "contentsChanged")
	}
}

//export callbackQTextDocumentContentsChanged
func callbackQTextDocumentContentsChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::contentsChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "contentsChanged").(func())()
}

func (ptr *QTextDocument) DefaultCursorMoveStyle() core.Qt__CursorMoveStyle {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::defaultCursorMoveStyle")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__CursorMoveStyle(C.QTextDocument_DefaultCursorMoveStyle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextDocument) DocumentLayout() *QAbstractTextDocumentLayout {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::documentLayout")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAbstractTextDocumentLayoutFromPointer(C.QTextDocument_DocumentLayout(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextDocument) ConnectDocumentLayoutChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::documentLayoutChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_ConnectDocumentLayoutChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "documentLayoutChanged", f)
	}
}

func (ptr *QTextDocument) DisconnectDocumentLayoutChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::documentLayoutChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_DisconnectDocumentLayoutChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "documentLayoutChanged")
	}
}

//export callbackQTextDocumentDocumentLayoutChanged
func callbackQTextDocumentDocumentLayoutChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::documentLayoutChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "documentLayoutChanged").(func())()
}

func (ptr *QTextDocument) DrawContents(p QPainter_ITF, rect core.QRectF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::drawContents")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_DrawContents(ptr.Pointer(), PointerFromQPainter(p), core.PointerFromQRectF(rect))
	}
}

func (ptr *QTextDocument) IdealWidth() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::idealWidth")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextDocument_IdealWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextDocument) IsEmpty() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::isEmpty")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextDocument_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextDocument) IsRedoAvailable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::isRedoAvailable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextDocument_IsRedoAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextDocument) IsUndoAvailable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::isUndoAvailable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextDocument_IsUndoAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextDocument) LineCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::lineCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextDocument_LineCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextDocument) MetaInformation(info QTextDocument__MetaInformation) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::metaInformation")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextDocument_MetaInformation(ptr.Pointer(), C.int(info)))
	}
	return ""
}

func (ptr *QTextDocument) ConnectModificationChanged(f func(changed bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::modificationChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_ConnectModificationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "modificationChanged", f)
	}
}

func (ptr *QTextDocument) DisconnectModificationChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::modificationChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_DisconnectModificationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "modificationChanged")
	}
}

//export callbackQTextDocumentModificationChanged
func callbackQTextDocumentModificationChanged(ptrName *C.char, changed C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::modificationChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "modificationChanged").(func(bool))(int(changed) != 0)
}

func (ptr *QTextDocument) Object(objectIndex int) *QTextObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::object")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTextObjectFromPointer(C.QTextDocument_Object(ptr.Pointer(), C.int(objectIndex)))
	}
	return nil
}

func (ptr *QTextDocument) ObjectForFormat(f QTextFormat_ITF) *QTextObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::objectForFormat")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTextObjectFromPointer(C.QTextDocument_ObjectForFormat(ptr.Pointer(), PointerFromQTextFormat(f)))
	}
	return nil
}

func (ptr *QTextDocument) PageCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::pageCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextDocument_PageCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextDocument) Print(printer QPagedPaintDevice_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::print")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_Print(ptr.Pointer(), PointerFromQPagedPaintDevice(printer))
	}
}

func (ptr *QTextDocument) Redo2() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::redo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_Redo2(ptr.Pointer())
	}
}

func (ptr *QTextDocument) Redo(cursor QTextCursor_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::redo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_Redo(ptr.Pointer(), PointerFromQTextCursor(cursor))
	}
}

func (ptr *QTextDocument) ConnectRedoAvailable(f func(available bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::redoAvailable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_ConnectRedoAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "redoAvailable", f)
	}
}

func (ptr *QTextDocument) DisconnectRedoAvailable() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::redoAvailable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_DisconnectRedoAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "redoAvailable")
	}
}

//export callbackQTextDocumentRedoAvailable
func callbackQTextDocumentRedoAvailable(ptrName *C.char, available C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::redoAvailable")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "redoAvailable").(func(bool))(int(available) != 0)
}

func (ptr *QTextDocument) Resource(ty int, name core.QUrl_ITF) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::resource")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QTextDocument_Resource(ptr.Pointer(), C.int(ty), core.PointerFromQUrl(name)))
	}
	return nil
}

func (ptr *QTextDocument) Revision() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::revision")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextDocument_Revision(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextDocument) RootFrame() *QTextFrame {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::rootFrame")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTextFrameFromPointer(C.QTextDocument_RootFrame(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextDocument) SetDefaultCursorMoveStyle(style core.Qt__CursorMoveStyle) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::setDefaultCursorMoveStyle")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_SetDefaultCursorMoveStyle(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QTextDocument) SetDefaultFont(font QFont_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::setDefaultFont")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_SetDefaultFont(ptr.Pointer(), PointerFromQFont(font))
	}
}

func (ptr *QTextDocument) SetDefaultTextOption(option QTextOption_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::setDefaultTextOption")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_SetDefaultTextOption(ptr.Pointer(), PointerFromQTextOption(option))
	}
}

func (ptr *QTextDocument) SetDocumentLayout(layout QAbstractTextDocumentLayout_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::setDocumentLayout")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_SetDocumentLayout(ptr.Pointer(), PointerFromQAbstractTextDocumentLayout(layout))
	}
}

func (ptr *QTextDocument) SetHtml(html string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::setHtml")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_SetHtml(ptr.Pointer(), C.CString(html))
	}
}

func (ptr *QTextDocument) SetIndentWidth(width float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::setIndentWidth")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_SetIndentWidth(ptr.Pointer(), C.double(width))
	}
}

func (ptr *QTextDocument) SetMetaInformation(info QTextDocument__MetaInformation, stri string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::setMetaInformation")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_SetMetaInformation(ptr.Pointer(), C.int(info), C.CString(stri))
	}
}

func (ptr *QTextDocument) SetPlainText(text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::setPlainText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_SetPlainText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QTextDocument) ToHtml(encoding core.QByteArray_ITF) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::toHtml")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextDocument_ToHtml(ptr.Pointer(), core.PointerFromQByteArray(encoding)))
	}
	return ""
}

func (ptr *QTextDocument) ToPlainText() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::toPlainText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextDocument_ToPlainText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextDocument) Undo2() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::undo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_Undo2(ptr.Pointer())
	}
}

func (ptr *QTextDocument) Undo(cursor QTextCursor_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::undo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_Undo(ptr.Pointer(), PointerFromQTextCursor(cursor))
	}
}

func (ptr *QTextDocument) ConnectUndoAvailable(f func(available bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::undoAvailable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_ConnectUndoAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "undoAvailable", f)
	}
}

func (ptr *QTextDocument) DisconnectUndoAvailable() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::undoAvailable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_DisconnectUndoAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "undoAvailable")
	}
}

//export callbackQTextDocumentUndoAvailable
func callbackQTextDocumentUndoAvailable(ptrName *C.char, available C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::undoAvailable")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "undoAvailable").(func(bool))(int(available) != 0)
}

func (ptr *QTextDocument) ConnectUndoCommandAdded(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::undoCommandAdded")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_ConnectUndoCommandAdded(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "undoCommandAdded", f)
	}
}

func (ptr *QTextDocument) DisconnectUndoCommandAdded() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::undoCommandAdded")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_DisconnectUndoCommandAdded(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "undoCommandAdded")
	}
}

//export callbackQTextDocumentUndoCommandAdded
func callbackQTextDocumentUndoCommandAdded(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::undoCommandAdded")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "undoCommandAdded").(func())()
}

func (ptr *QTextDocument) DestroyQTextDocument() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocument::~QTextDocument")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocument_DestroyQTextDocument(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
