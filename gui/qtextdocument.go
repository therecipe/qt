package gui

//#include "qtextdocument.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTextDocument struct {
	core.QObject
}

type QTextDocumentITF interface {
	core.QObjectITF
	QTextDocumentPTR() *QTextDocument
}

func PointerFromQTextDocument(ptr QTextDocumentITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextDocumentPTR().Pointer()
	}
	return nil
}

func QTextDocumentFromPointer(ptr unsafe.Pointer) *QTextDocument {
	var n = new(QTextDocument)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTextDocument_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTextDocument) QTextDocumentPTR() *QTextDocument {
	return ptr
}

//QTextDocument::FindFlag
type QTextDocument__FindFlag int

var (
	QTextDocument__FindBackward        = QTextDocument__FindFlag(0x00001)
	QTextDocument__FindCaseSensitively = QTextDocument__FindFlag(0x00002)
	QTextDocument__FindWholeWords      = QTextDocument__FindFlag(0x00004)
)

//QTextDocument::MetaInformation
type QTextDocument__MetaInformation int

var (
	QTextDocument__DocumentTitle = QTextDocument__MetaInformation(0)
	QTextDocument__DocumentUrl   = QTextDocument__MetaInformation(1)
)

//QTextDocument::ResourceType
type QTextDocument__ResourceType int

var (
	QTextDocument__HtmlResource       = QTextDocument__ResourceType(1)
	QTextDocument__ImageResource      = QTextDocument__ResourceType(2)
	QTextDocument__StyleSheetResource = QTextDocument__ResourceType(3)
	QTextDocument__UserResource       = QTextDocument__ResourceType(100)
)

//QTextDocument::Stacks
type QTextDocument__Stacks int

var (
	QTextDocument__UndoStack         = QTextDocument__Stacks(0x01)
	QTextDocument__RedoStack         = QTextDocument__Stacks(0x02)
	QTextDocument__UndoAndRedoStacks = QTextDocument__Stacks(QTextDocument__UndoStack | QTextDocument__RedoStack)
)

func (ptr *QTextDocument) BaseUrl() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextDocument_BaseUrl(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTextDocument) BlockCount() int {
	if ptr.Pointer() != nil {
		return int(C.QTextDocument_BlockCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextDocument) DefaultStyleSheet() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextDocument_DefaultStyleSheet(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTextDocument) IsModified() bool {
	if ptr.Pointer() != nil {
		return C.QTextDocument_IsModified(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextDocument) IsUndoRedoEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QTextDocument_IsUndoRedoEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextDocument) MarkContentsDirty(position int, length int) {
	if ptr.Pointer() != nil {
		C.QTextDocument_MarkContentsDirty(C.QtObjectPtr(ptr.Pointer()), C.int(position), C.int(length))
	}
}

func (ptr *QTextDocument) MaximumBlockCount() int {
	if ptr.Pointer() != nil {
		return int(C.QTextDocument_MaximumBlockCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextDocument) SetBaseUrl(url string) {
	if ptr.Pointer() != nil {
		C.QTextDocument_SetBaseUrl(C.QtObjectPtr(ptr.Pointer()), C.CString(url))
	}
}

func (ptr *QTextDocument) SetDefaultStyleSheet(sheet string) {
	if ptr.Pointer() != nil {
		C.QTextDocument_SetDefaultStyleSheet(C.QtObjectPtr(ptr.Pointer()), C.CString(sheet))
	}
}

func (ptr *QTextDocument) SetMaximumBlockCount(maximum int) {
	if ptr.Pointer() != nil {
		C.QTextDocument_SetMaximumBlockCount(C.QtObjectPtr(ptr.Pointer()), C.int(maximum))
	}
}

func (ptr *QTextDocument) SetModified(m bool) {
	if ptr.Pointer() != nil {
		C.QTextDocument_SetModified(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(m)))
	}
}

func (ptr *QTextDocument) SetPageSize(size core.QSizeFITF) {
	if ptr.Pointer() != nil {
		C.QTextDocument_SetPageSize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSizeF(size)))
	}
}

func (ptr *QTextDocument) SetUndoRedoEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QTextDocument_SetUndoRedoEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTextDocument) SetUseDesignMetrics(b bool) {
	if ptr.Pointer() != nil {
		C.QTextDocument_SetUseDesignMetrics(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QTextDocument) UseDesignMetrics() bool {
	if ptr.Pointer() != nil {
		return C.QTextDocument_UseDesignMetrics(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func NewQTextDocument(parent core.QObjectITF) *QTextDocument {
	return QTextDocumentFromPointer(unsafe.Pointer(C.QTextDocument_NewQTextDocument(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQTextDocument2(text string, parent core.QObjectITF) *QTextDocument {
	return QTextDocumentFromPointer(unsafe.Pointer(C.QTextDocument_NewQTextDocument2(C.CString(text), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QTextDocument) AddResource(ty int, name string, resource string) {
	if ptr.Pointer() != nil {
		C.QTextDocument_AddResource(C.QtObjectPtr(ptr.Pointer()), C.int(ty), C.CString(name), C.CString(resource))
	}
}

func (ptr *QTextDocument) AdjustSize() {
	if ptr.Pointer() != nil {
		C.QTextDocument_AdjustSize(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextDocument) AvailableRedoSteps() int {
	if ptr.Pointer() != nil {
		return int(C.QTextDocument_AvailableRedoSteps(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextDocument) AvailableUndoSteps() int {
	if ptr.Pointer() != nil {
		return int(C.QTextDocument_AvailableUndoSteps(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextDocument) ConnectBaseUrlChanged(f func(url string)) {
	if ptr.Pointer() != nil {
		C.QTextDocument_ConnectBaseUrlChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "baseUrlChanged", f)
	}
}

func (ptr *QTextDocument) DisconnectBaseUrlChanged() {
	if ptr.Pointer() != nil {
		C.QTextDocument_DisconnectBaseUrlChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "baseUrlChanged")
	}
}

//export callbackQTextDocumentBaseUrlChanged
func callbackQTextDocumentBaseUrlChanged(ptrName *C.char, url *C.char) {
	qt.GetSignal(C.GoString(ptrName), "baseUrlChanged").(func(string))(C.GoString(url))
}

func (ptr *QTextDocument) ConnectBlockCountChanged(f func(newBlockCount int)) {
	if ptr.Pointer() != nil {
		C.QTextDocument_ConnectBlockCountChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "blockCountChanged", f)
	}
}

func (ptr *QTextDocument) DisconnectBlockCountChanged() {
	if ptr.Pointer() != nil {
		C.QTextDocument_DisconnectBlockCountChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "blockCountChanged")
	}
}

//export callbackQTextDocumentBlockCountChanged
func callbackQTextDocumentBlockCountChanged(ptrName *C.char, newBlockCount C.int) {
	qt.GetSignal(C.GoString(ptrName), "blockCountChanged").(func(int))(int(newBlockCount))
}

func (ptr *QTextDocument) CharacterCount() int {
	if ptr.Pointer() != nil {
		return int(C.QTextDocument_CharacterCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextDocument) Clear() {
	if ptr.Pointer() != nil {
		C.QTextDocument_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextDocument) ClearUndoRedoStacks(stacksToClear QTextDocument__Stacks) {
	if ptr.Pointer() != nil {
		C.QTextDocument_ClearUndoRedoStacks(C.QtObjectPtr(ptr.Pointer()), C.int(stacksToClear))
	}
}

func (ptr *QTextDocument) Clone(parent core.QObjectITF) *QTextDocument {
	if ptr.Pointer() != nil {
		return QTextDocumentFromPointer(unsafe.Pointer(C.QTextDocument_Clone(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQObject(parent)))))
	}
	return nil
}

func (ptr *QTextDocument) ConnectContentsChange(f func(position int, charsRemoved int, charsAdded int)) {
	if ptr.Pointer() != nil {
		C.QTextDocument_ConnectContentsChange(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "contentsChange", f)
	}
}

func (ptr *QTextDocument) DisconnectContentsChange() {
	if ptr.Pointer() != nil {
		C.QTextDocument_DisconnectContentsChange(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "contentsChange")
	}
}

//export callbackQTextDocumentContentsChange
func callbackQTextDocumentContentsChange(ptrName *C.char, position C.int, charsRemoved C.int, charsAdded C.int) {
	qt.GetSignal(C.GoString(ptrName), "contentsChange").(func(int, int, int))(int(position), int(charsRemoved), int(charsAdded))
}

func (ptr *QTextDocument) ConnectContentsChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QTextDocument_ConnectContentsChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "contentsChanged", f)
	}
}

func (ptr *QTextDocument) DisconnectContentsChanged() {
	if ptr.Pointer() != nil {
		C.QTextDocument_DisconnectContentsChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "contentsChanged")
	}
}

//export callbackQTextDocumentContentsChanged
func callbackQTextDocumentContentsChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "contentsChanged").(func())()
}

func (ptr *QTextDocument) DefaultCursorMoveStyle() core.Qt__CursorMoveStyle {
	if ptr.Pointer() != nil {
		return core.Qt__CursorMoveStyle(C.QTextDocument_DefaultCursorMoveStyle(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextDocument) DocumentLayout() *QAbstractTextDocumentLayout {
	if ptr.Pointer() != nil {
		return QAbstractTextDocumentLayoutFromPointer(unsafe.Pointer(C.QTextDocument_DocumentLayout(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTextDocument) ConnectDocumentLayoutChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QTextDocument_ConnectDocumentLayoutChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "documentLayoutChanged", f)
	}
}

func (ptr *QTextDocument) DisconnectDocumentLayoutChanged() {
	if ptr.Pointer() != nil {
		C.QTextDocument_DisconnectDocumentLayoutChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "documentLayoutChanged")
	}
}

//export callbackQTextDocumentDocumentLayoutChanged
func callbackQTextDocumentDocumentLayoutChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "documentLayoutChanged").(func())()
}

func (ptr *QTextDocument) DrawContents(p QPainterITF, rect core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QTextDocument_DrawContents(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPainter(p)), C.QtObjectPtr(core.PointerFromQRectF(rect)))
	}
}

func (ptr *QTextDocument) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QTextDocument_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextDocument) IsRedoAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QTextDocument_IsRedoAvailable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextDocument) IsUndoAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QTextDocument_IsUndoAvailable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextDocument) LineCount() int {
	if ptr.Pointer() != nil {
		return int(C.QTextDocument_LineCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextDocument) MetaInformation(info QTextDocument__MetaInformation) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextDocument_MetaInformation(C.QtObjectPtr(ptr.Pointer()), C.int(info)))
	}
	return ""
}

func (ptr *QTextDocument) ConnectModificationChanged(f func(changed bool)) {
	if ptr.Pointer() != nil {
		C.QTextDocument_ConnectModificationChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "modificationChanged", f)
	}
}

func (ptr *QTextDocument) DisconnectModificationChanged() {
	if ptr.Pointer() != nil {
		C.QTextDocument_DisconnectModificationChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "modificationChanged")
	}
}

//export callbackQTextDocumentModificationChanged
func callbackQTextDocumentModificationChanged(ptrName *C.char, changed C.int) {
	qt.GetSignal(C.GoString(ptrName), "modificationChanged").(func(bool))(int(changed) != 0)
}

func (ptr *QTextDocument) Object(objectIndex int) *QTextObject {
	if ptr.Pointer() != nil {
		return QTextObjectFromPointer(unsafe.Pointer(C.QTextDocument_Object(C.QtObjectPtr(ptr.Pointer()), C.int(objectIndex))))
	}
	return nil
}

func (ptr *QTextDocument) ObjectForFormat(f QTextFormatITF) *QTextObject {
	if ptr.Pointer() != nil {
		return QTextObjectFromPointer(unsafe.Pointer(C.QTextDocument_ObjectForFormat(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextFormat(f)))))
	}
	return nil
}

func (ptr *QTextDocument) PageCount() int {
	if ptr.Pointer() != nil {
		return int(C.QTextDocument_PageCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextDocument) Print(printer QPagedPaintDeviceITF) {
	if ptr.Pointer() != nil {
		C.QTextDocument_Print(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPagedPaintDevice(printer)))
	}
}

func (ptr *QTextDocument) Redo2() {
	if ptr.Pointer() != nil {
		C.QTextDocument_Redo2(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextDocument) Redo(cursor QTextCursorITF) {
	if ptr.Pointer() != nil {
		C.QTextDocument_Redo(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextCursor(cursor)))
	}
}

func (ptr *QTextDocument) ConnectRedoAvailable(f func(available bool)) {
	if ptr.Pointer() != nil {
		C.QTextDocument_ConnectRedoAvailable(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "redoAvailable", f)
	}
}

func (ptr *QTextDocument) DisconnectRedoAvailable() {
	if ptr.Pointer() != nil {
		C.QTextDocument_DisconnectRedoAvailable(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "redoAvailable")
	}
}

//export callbackQTextDocumentRedoAvailable
func callbackQTextDocumentRedoAvailable(ptrName *C.char, available C.int) {
	qt.GetSignal(C.GoString(ptrName), "redoAvailable").(func(bool))(int(available) != 0)
}

func (ptr *QTextDocument) Resource(ty int, name string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextDocument_Resource(C.QtObjectPtr(ptr.Pointer()), C.int(ty), C.CString(name)))
	}
	return ""
}

func (ptr *QTextDocument) Revision() int {
	if ptr.Pointer() != nil {
		return int(C.QTextDocument_Revision(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextDocument) RootFrame() *QTextFrame {
	if ptr.Pointer() != nil {
		return QTextFrameFromPointer(unsafe.Pointer(C.QTextDocument_RootFrame(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTextDocument) SetDefaultCursorMoveStyle(style core.Qt__CursorMoveStyle) {
	if ptr.Pointer() != nil {
		C.QTextDocument_SetDefaultCursorMoveStyle(C.QtObjectPtr(ptr.Pointer()), C.int(style))
	}
}

func (ptr *QTextDocument) SetDefaultFont(font QFontITF) {
	if ptr.Pointer() != nil {
		C.QTextDocument_SetDefaultFont(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQFont(font)))
	}
}

func (ptr *QTextDocument) SetDefaultTextOption(option QTextOptionITF) {
	if ptr.Pointer() != nil {
		C.QTextDocument_SetDefaultTextOption(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextOption(option)))
	}
}

func (ptr *QTextDocument) SetDocumentLayout(layout QAbstractTextDocumentLayoutITF) {
	if ptr.Pointer() != nil {
		C.QTextDocument_SetDocumentLayout(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractTextDocumentLayout(layout)))
	}
}

func (ptr *QTextDocument) SetHtml(html string) {
	if ptr.Pointer() != nil {
		C.QTextDocument_SetHtml(C.QtObjectPtr(ptr.Pointer()), C.CString(html))
	}
}

func (ptr *QTextDocument) SetMetaInformation(info QTextDocument__MetaInformation, stri string) {
	if ptr.Pointer() != nil {
		C.QTextDocument_SetMetaInformation(C.QtObjectPtr(ptr.Pointer()), C.int(info), C.CString(stri))
	}
}

func (ptr *QTextDocument) SetPlainText(text string) {
	if ptr.Pointer() != nil {
		C.QTextDocument_SetPlainText(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QTextDocument) ToHtml(encoding core.QByteArrayITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextDocument_ToHtml(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(encoding))))
	}
	return ""
}

func (ptr *QTextDocument) ToPlainText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextDocument_ToPlainText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTextDocument) Undo2() {
	if ptr.Pointer() != nil {
		C.QTextDocument_Undo2(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextDocument) Undo(cursor QTextCursorITF) {
	if ptr.Pointer() != nil {
		C.QTextDocument_Undo(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextCursor(cursor)))
	}
}

func (ptr *QTextDocument) ConnectUndoAvailable(f func(available bool)) {
	if ptr.Pointer() != nil {
		C.QTextDocument_ConnectUndoAvailable(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "undoAvailable", f)
	}
}

func (ptr *QTextDocument) DisconnectUndoAvailable() {
	if ptr.Pointer() != nil {
		C.QTextDocument_DisconnectUndoAvailable(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "undoAvailable")
	}
}

//export callbackQTextDocumentUndoAvailable
func callbackQTextDocumentUndoAvailable(ptrName *C.char, available C.int) {
	qt.GetSignal(C.GoString(ptrName), "undoAvailable").(func(bool))(int(available) != 0)
}

func (ptr *QTextDocument) ConnectUndoCommandAdded(f func()) {
	if ptr.Pointer() != nil {
		C.QTextDocument_ConnectUndoCommandAdded(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "undoCommandAdded", f)
	}
}

func (ptr *QTextDocument) DisconnectUndoCommandAdded() {
	if ptr.Pointer() != nil {
		C.QTextDocument_DisconnectUndoCommandAdded(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "undoCommandAdded")
	}
}

//export callbackQTextDocumentUndoCommandAdded
func callbackQTextDocumentUndoCommandAdded(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "undoCommandAdded").(func())()
}

func (ptr *QTextDocument) DestroyQTextDocument() {
	if ptr.Pointer() != nil {
		C.QTextDocument_DestroyQTextDocument(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
