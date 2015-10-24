package gui

//#include "qtextcursor.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QTextCursor struct {
	ptr unsafe.Pointer
}

type QTextCursorITF interface {
	QTextCursorPTR() *QTextCursor
}

func (p *QTextCursor) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextCursor) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextCursor(ptr QTextCursorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextCursorPTR().Pointer()
	}
	return nil
}

func QTextCursorFromPointer(ptr unsafe.Pointer) *QTextCursor {
	var n = new(QTextCursor)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextCursor) QTextCursorPTR() *QTextCursor {
	return ptr
}

//QTextCursor::MoveMode
type QTextCursor__MoveMode int

var (
	QTextCursor__MoveAnchor = QTextCursor__MoveMode(0)
	QTextCursor__KeepAnchor = QTextCursor__MoveMode(1)
)

//QTextCursor::MoveOperation
type QTextCursor__MoveOperation int

var (
	QTextCursor__NoMove            = QTextCursor__MoveOperation(0)
	QTextCursor__Start             = QTextCursor__MoveOperation(1)
	QTextCursor__Up                = QTextCursor__MoveOperation(2)
	QTextCursor__StartOfLine       = QTextCursor__MoveOperation(3)
	QTextCursor__StartOfBlock      = QTextCursor__MoveOperation(4)
	QTextCursor__StartOfWord       = QTextCursor__MoveOperation(5)
	QTextCursor__PreviousBlock     = QTextCursor__MoveOperation(6)
	QTextCursor__PreviousCharacter = QTextCursor__MoveOperation(7)
	QTextCursor__PreviousWord      = QTextCursor__MoveOperation(8)
	QTextCursor__Left              = QTextCursor__MoveOperation(9)
	QTextCursor__WordLeft          = QTextCursor__MoveOperation(10)
	QTextCursor__End               = QTextCursor__MoveOperation(11)
	QTextCursor__Down              = QTextCursor__MoveOperation(12)
	QTextCursor__EndOfLine         = QTextCursor__MoveOperation(13)
	QTextCursor__EndOfWord         = QTextCursor__MoveOperation(14)
	QTextCursor__EndOfBlock        = QTextCursor__MoveOperation(15)
	QTextCursor__NextBlock         = QTextCursor__MoveOperation(16)
	QTextCursor__NextCharacter     = QTextCursor__MoveOperation(17)
	QTextCursor__NextWord          = QTextCursor__MoveOperation(18)
	QTextCursor__Right             = QTextCursor__MoveOperation(19)
	QTextCursor__WordRight         = QTextCursor__MoveOperation(20)
	QTextCursor__NextCell          = QTextCursor__MoveOperation(21)
	QTextCursor__PreviousCell      = QTextCursor__MoveOperation(22)
	QTextCursor__NextRow           = QTextCursor__MoveOperation(23)
	QTextCursor__PreviousRow       = QTextCursor__MoveOperation(24)
)

//QTextCursor::SelectionType
type QTextCursor__SelectionType int

var (
	QTextCursor__WordUnderCursor  = QTextCursor__SelectionType(0)
	QTextCursor__LineUnderCursor  = QTextCursor__SelectionType(1)
	QTextCursor__BlockUnderCursor = QTextCursor__SelectionType(2)
	QTextCursor__Document         = QTextCursor__SelectionType(3)
)

func (ptr *QTextCursor) InsertBlock3(format QTextBlockFormatITF, charFormat QTextCharFormatITF) {
	if ptr.Pointer() != nil {
		C.QTextCursor_InsertBlock3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextBlockFormat(format)), C.QtObjectPtr(PointerFromQTextCharFormat(charFormat)))
	}
}

func (ptr *QTextCursor) InsertTable2(rows int, columns int) *QTextTable {
	if ptr.Pointer() != nil {
		return QTextTableFromPointer(unsafe.Pointer(C.QTextCursor_InsertTable2(C.QtObjectPtr(ptr.Pointer()), C.int(rows), C.int(columns))))
	}
	return nil
}

func (ptr *QTextCursor) InsertTable(rows int, columns int, format QTextTableFormatITF) *QTextTable {
	if ptr.Pointer() != nil {
		return QTextTableFromPointer(unsafe.Pointer(C.QTextCursor_InsertTable(C.QtObjectPtr(ptr.Pointer()), C.int(rows), C.int(columns), C.QtObjectPtr(PointerFromQTextTableFormat(format)))))
	}
	return nil
}

func (ptr *QTextCursor) InsertText2(text string, format QTextCharFormatITF) {
	if ptr.Pointer() != nil {
		C.QTextCursor_InsertText2(C.QtObjectPtr(ptr.Pointer()), C.CString(text), C.QtObjectPtr(PointerFromQTextCharFormat(format)))
	}
}

func (ptr *QTextCursor) MovePosition(operation QTextCursor__MoveOperation, mode QTextCursor__MoveMode, n int) bool {
	if ptr.Pointer() != nil {
		return C.QTextCursor_MovePosition(C.QtObjectPtr(ptr.Pointer()), C.int(operation), C.int(mode), C.int(n)) != 0
	}
	return false
}

func NewQTextCursor() *QTextCursor {
	return QTextCursorFromPointer(unsafe.Pointer(C.QTextCursor_NewQTextCursor()))
}

func NewQTextCursor2(document QTextDocumentITF) *QTextCursor {
	return QTextCursorFromPointer(unsafe.Pointer(C.QTextCursor_NewQTextCursor2(C.QtObjectPtr(PointerFromQTextDocument(document)))))
}

func NewQTextCursor4(frame QTextFrameITF) *QTextCursor {
	return QTextCursorFromPointer(unsafe.Pointer(C.QTextCursor_NewQTextCursor4(C.QtObjectPtr(PointerFromQTextFrame(frame)))))
}

func NewQTextCursor5(block QTextBlockITF) *QTextCursor {
	return QTextCursorFromPointer(unsafe.Pointer(C.QTextCursor_NewQTextCursor5(C.QtObjectPtr(PointerFromQTextBlock(block)))))
}

func NewQTextCursor7(cursor QTextCursorITF) *QTextCursor {
	return QTextCursorFromPointer(unsafe.Pointer(C.QTextCursor_NewQTextCursor7(C.QtObjectPtr(PointerFromQTextCursor(cursor)))))
}

func (ptr *QTextCursor) Anchor() int {
	if ptr.Pointer() != nil {
		return int(C.QTextCursor_Anchor(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextCursor) AtBlockEnd() bool {
	if ptr.Pointer() != nil {
		return C.QTextCursor_AtBlockEnd(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextCursor) AtBlockStart() bool {
	if ptr.Pointer() != nil {
		return C.QTextCursor_AtBlockStart(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextCursor) AtEnd() bool {
	if ptr.Pointer() != nil {
		return C.QTextCursor_AtEnd(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextCursor) AtStart() bool {
	if ptr.Pointer() != nil {
		return C.QTextCursor_AtStart(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextCursor) BeginEditBlock() {
	if ptr.Pointer() != nil {
		C.QTextCursor_BeginEditBlock(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextCursor) BlockNumber() int {
	if ptr.Pointer() != nil {
		return int(C.QTextCursor_BlockNumber(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextCursor) ClearSelection() {
	if ptr.Pointer() != nil {
		C.QTextCursor_ClearSelection(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextCursor) ColumnNumber() int {
	if ptr.Pointer() != nil {
		return int(C.QTextCursor_ColumnNumber(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextCursor) CreateList2(style QTextListFormat__Style) *QTextList {
	if ptr.Pointer() != nil {
		return QTextListFromPointer(unsafe.Pointer(C.QTextCursor_CreateList2(C.QtObjectPtr(ptr.Pointer()), C.int(style))))
	}
	return nil
}

func (ptr *QTextCursor) CreateList(format QTextListFormatITF) *QTextList {
	if ptr.Pointer() != nil {
		return QTextListFromPointer(unsafe.Pointer(C.QTextCursor_CreateList(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextListFormat(format)))))
	}
	return nil
}

func (ptr *QTextCursor) CurrentFrame() *QTextFrame {
	if ptr.Pointer() != nil {
		return QTextFrameFromPointer(unsafe.Pointer(C.QTextCursor_CurrentFrame(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTextCursor) CurrentList() *QTextList {
	if ptr.Pointer() != nil {
		return QTextListFromPointer(unsafe.Pointer(C.QTextCursor_CurrentList(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTextCursor) CurrentTable() *QTextTable {
	if ptr.Pointer() != nil {
		return QTextTableFromPointer(unsafe.Pointer(C.QTextCursor_CurrentTable(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTextCursor) DeleteChar() {
	if ptr.Pointer() != nil {
		C.QTextCursor_DeleteChar(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextCursor) DeletePreviousChar() {
	if ptr.Pointer() != nil {
		C.QTextCursor_DeletePreviousChar(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextCursor) Document() *QTextDocument {
	if ptr.Pointer() != nil {
		return QTextDocumentFromPointer(unsafe.Pointer(C.QTextCursor_Document(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTextCursor) EndEditBlock() {
	if ptr.Pointer() != nil {
		C.QTextCursor_EndEditBlock(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextCursor) HasComplexSelection() bool {
	if ptr.Pointer() != nil {
		return C.QTextCursor_HasComplexSelection(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextCursor) HasSelection() bool {
	if ptr.Pointer() != nil {
		return C.QTextCursor_HasSelection(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextCursor) InsertBlock() {
	if ptr.Pointer() != nil {
		C.QTextCursor_InsertBlock(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextCursor) InsertBlock2(format QTextBlockFormatITF) {
	if ptr.Pointer() != nil {
		C.QTextCursor_InsertBlock2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextBlockFormat(format)))
	}
}

func (ptr *QTextCursor) InsertFragment(fragment QTextDocumentFragmentITF) {
	if ptr.Pointer() != nil {
		C.QTextCursor_InsertFragment(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextDocumentFragment(fragment)))
	}
}

func (ptr *QTextCursor) InsertFrame(format QTextFrameFormatITF) *QTextFrame {
	if ptr.Pointer() != nil {
		return QTextFrameFromPointer(unsafe.Pointer(C.QTextCursor_InsertFrame(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextFrameFormat(format)))))
	}
	return nil
}

func (ptr *QTextCursor) InsertHtml(html string) {
	if ptr.Pointer() != nil {
		C.QTextCursor_InsertHtml(C.QtObjectPtr(ptr.Pointer()), C.CString(html))
	}
}

func (ptr *QTextCursor) InsertImage4(image QImageITF, name string) {
	if ptr.Pointer() != nil {
		C.QTextCursor_InsertImage4(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQImage(image)), C.CString(name))
	}
}

func (ptr *QTextCursor) InsertImage3(name string) {
	if ptr.Pointer() != nil {
		C.QTextCursor_InsertImage3(C.QtObjectPtr(ptr.Pointer()), C.CString(name))
	}
}

func (ptr *QTextCursor) InsertImage(format QTextImageFormatITF) {
	if ptr.Pointer() != nil {
		C.QTextCursor_InsertImage(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextImageFormat(format)))
	}
}

func (ptr *QTextCursor) InsertImage2(format QTextImageFormatITF, alignment QTextFrameFormat__Position) {
	if ptr.Pointer() != nil {
		C.QTextCursor_InsertImage2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextImageFormat(format)), C.int(alignment))
	}
}

func (ptr *QTextCursor) InsertList2(style QTextListFormat__Style) *QTextList {
	if ptr.Pointer() != nil {
		return QTextListFromPointer(unsafe.Pointer(C.QTextCursor_InsertList2(C.QtObjectPtr(ptr.Pointer()), C.int(style))))
	}
	return nil
}

func (ptr *QTextCursor) InsertList(format QTextListFormatITF) *QTextList {
	if ptr.Pointer() != nil {
		return QTextListFromPointer(unsafe.Pointer(C.QTextCursor_InsertList(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextListFormat(format)))))
	}
	return nil
}

func (ptr *QTextCursor) InsertText(text string) {
	if ptr.Pointer() != nil {
		C.QTextCursor_InsertText(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QTextCursor) IsCopyOf(other QTextCursorITF) bool {
	if ptr.Pointer() != nil {
		return C.QTextCursor_IsCopyOf(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextCursor(other))) != 0
	}
	return false
}

func (ptr *QTextCursor) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QTextCursor_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextCursor) JoinPreviousEditBlock() {
	if ptr.Pointer() != nil {
		C.QTextCursor_JoinPreviousEditBlock(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextCursor) KeepPositionOnInsert() bool {
	if ptr.Pointer() != nil {
		return C.QTextCursor_KeepPositionOnInsert(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextCursor) MergeBlockCharFormat(modifier QTextCharFormatITF) {
	if ptr.Pointer() != nil {
		C.QTextCursor_MergeBlockCharFormat(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextCharFormat(modifier)))
	}
}

func (ptr *QTextCursor) MergeBlockFormat(modifier QTextBlockFormatITF) {
	if ptr.Pointer() != nil {
		C.QTextCursor_MergeBlockFormat(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextBlockFormat(modifier)))
	}
}

func (ptr *QTextCursor) MergeCharFormat(modifier QTextCharFormatITF) {
	if ptr.Pointer() != nil {
		C.QTextCursor_MergeCharFormat(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextCharFormat(modifier)))
	}
}

func (ptr *QTextCursor) Position() int {
	if ptr.Pointer() != nil {
		return int(C.QTextCursor_Position(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextCursor) PositionInBlock() int {
	if ptr.Pointer() != nil {
		return int(C.QTextCursor_PositionInBlock(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextCursor) RemoveSelectedText() {
	if ptr.Pointer() != nil {
		C.QTextCursor_RemoveSelectedText(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextCursor) Select(selection QTextCursor__SelectionType) {
	if ptr.Pointer() != nil {
		C.QTextCursor_Select(C.QtObjectPtr(ptr.Pointer()), C.int(selection))
	}
}

func (ptr *QTextCursor) SelectedTableCells(firstRow int, numRows int, firstColumn int, numColumns int) {
	if ptr.Pointer() != nil {
		C.QTextCursor_SelectedTableCells(C.QtObjectPtr(ptr.Pointer()), C.int(firstRow), C.int(numRows), C.int(firstColumn), C.int(numColumns))
	}
}

func (ptr *QTextCursor) SelectedText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextCursor_SelectedText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTextCursor) SelectionEnd() int {
	if ptr.Pointer() != nil {
		return int(C.QTextCursor_SelectionEnd(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextCursor) SelectionStart() int {
	if ptr.Pointer() != nil {
		return int(C.QTextCursor_SelectionStart(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextCursor) SetBlockCharFormat(format QTextCharFormatITF) {
	if ptr.Pointer() != nil {
		C.QTextCursor_SetBlockCharFormat(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextCharFormat(format)))
	}
}

func (ptr *QTextCursor) SetBlockFormat(format QTextBlockFormatITF) {
	if ptr.Pointer() != nil {
		C.QTextCursor_SetBlockFormat(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextBlockFormat(format)))
	}
}

func (ptr *QTextCursor) SetCharFormat(format QTextCharFormatITF) {
	if ptr.Pointer() != nil {
		C.QTextCursor_SetCharFormat(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextCharFormat(format)))
	}
}

func (ptr *QTextCursor) SetKeepPositionOnInsert(b bool) {
	if ptr.Pointer() != nil {
		C.QTextCursor_SetKeepPositionOnInsert(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QTextCursor) SetPosition(pos int, m QTextCursor__MoveMode) {
	if ptr.Pointer() != nil {
		C.QTextCursor_SetPosition(C.QtObjectPtr(ptr.Pointer()), C.int(pos), C.int(m))
	}
}

func (ptr *QTextCursor) SetVerticalMovementX(x int) {
	if ptr.Pointer() != nil {
		C.QTextCursor_SetVerticalMovementX(C.QtObjectPtr(ptr.Pointer()), C.int(x))
	}
}

func (ptr *QTextCursor) SetVisualNavigation(b bool) {
	if ptr.Pointer() != nil {
		C.QTextCursor_SetVisualNavigation(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QTextCursor) Swap(other QTextCursorITF) {
	if ptr.Pointer() != nil {
		C.QTextCursor_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextCursor(other)))
	}
}

func (ptr *QTextCursor) VerticalMovementX() int {
	if ptr.Pointer() != nil {
		return int(C.QTextCursor_VerticalMovementX(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextCursor) VisualNavigation() bool {
	if ptr.Pointer() != nil {
		return C.QTextCursor_VisualNavigation(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextCursor) DestroyQTextCursor() {
	if ptr.Pointer() != nil {
		C.QTextCursor_DestroyQTextCursor(C.QtObjectPtr(ptr.Pointer()))
	}
}
