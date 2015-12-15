package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QTextCursor struct {
	ptr unsafe.Pointer
}

type QTextCursor_ITF interface {
	QTextCursor_PTR() *QTextCursor
}

func (p *QTextCursor) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextCursor) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextCursor(ptr QTextCursor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextCursor_PTR().Pointer()
	}
	return nil
}

func NewQTextCursorFromPointer(ptr unsafe.Pointer) *QTextCursor {
	var n = new(QTextCursor)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextCursor) QTextCursor_PTR() *QTextCursor {
	return ptr
}

//QTextCursor::MoveMode
type QTextCursor__MoveMode int64

const (
	QTextCursor__MoveAnchor = QTextCursor__MoveMode(0)
	QTextCursor__KeepAnchor = QTextCursor__MoveMode(1)
)

//QTextCursor::MoveOperation
type QTextCursor__MoveOperation int64

const (
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
type QTextCursor__SelectionType int64

const (
	QTextCursor__WordUnderCursor  = QTextCursor__SelectionType(0)
	QTextCursor__LineUnderCursor  = QTextCursor__SelectionType(1)
	QTextCursor__BlockUnderCursor = QTextCursor__SelectionType(2)
	QTextCursor__Document         = QTextCursor__SelectionType(3)
)

func (ptr *QTextCursor) InsertBlock3(format QTextBlockFormat_ITF, charFormat QTextCharFormat_ITF) {
	defer qt.Recovering("QTextCursor::insertBlock")

	if ptr.Pointer() != nil {
		C.QTextCursor_InsertBlock3(ptr.Pointer(), PointerFromQTextBlockFormat(format), PointerFromQTextCharFormat(charFormat))
	}
}

func (ptr *QTextCursor) InsertTable2(rows int, columns int) *QTextTable {
	defer qt.Recovering("QTextCursor::insertTable")

	if ptr.Pointer() != nil {
		return NewQTextTableFromPointer(C.QTextCursor_InsertTable2(ptr.Pointer(), C.int(rows), C.int(columns)))
	}
	return nil
}

func (ptr *QTextCursor) InsertTable(rows int, columns int, format QTextTableFormat_ITF) *QTextTable {
	defer qt.Recovering("QTextCursor::insertTable")

	if ptr.Pointer() != nil {
		return NewQTextTableFromPointer(C.QTextCursor_InsertTable(ptr.Pointer(), C.int(rows), C.int(columns), PointerFromQTextTableFormat(format)))
	}
	return nil
}

func (ptr *QTextCursor) InsertText2(text string, format QTextCharFormat_ITF) {
	defer qt.Recovering("QTextCursor::insertText")

	if ptr.Pointer() != nil {
		C.QTextCursor_InsertText2(ptr.Pointer(), C.CString(text), PointerFromQTextCharFormat(format))
	}
}

func (ptr *QTextCursor) MovePosition(operation QTextCursor__MoveOperation, mode QTextCursor__MoveMode, n int) bool {
	defer qt.Recovering("QTextCursor::movePosition")

	if ptr.Pointer() != nil {
		return C.QTextCursor_MovePosition(ptr.Pointer(), C.int(operation), C.int(mode), C.int(n)) != 0
	}
	return false
}

func NewQTextCursor() *QTextCursor {
	defer qt.Recovering("QTextCursor::QTextCursor")

	return NewQTextCursorFromPointer(C.QTextCursor_NewQTextCursor())
}

func NewQTextCursor2(document QTextDocument_ITF) *QTextCursor {
	defer qt.Recovering("QTextCursor::QTextCursor")

	return NewQTextCursorFromPointer(C.QTextCursor_NewQTextCursor2(PointerFromQTextDocument(document)))
}

func NewQTextCursor4(frame QTextFrame_ITF) *QTextCursor {
	defer qt.Recovering("QTextCursor::QTextCursor")

	return NewQTextCursorFromPointer(C.QTextCursor_NewQTextCursor4(PointerFromQTextFrame(frame)))
}

func NewQTextCursor5(block QTextBlock_ITF) *QTextCursor {
	defer qt.Recovering("QTextCursor::QTextCursor")

	return NewQTextCursorFromPointer(C.QTextCursor_NewQTextCursor5(PointerFromQTextBlock(block)))
}

func NewQTextCursor7(cursor QTextCursor_ITF) *QTextCursor {
	defer qt.Recovering("QTextCursor::QTextCursor")

	return NewQTextCursorFromPointer(C.QTextCursor_NewQTextCursor7(PointerFromQTextCursor(cursor)))
}

func (ptr *QTextCursor) Anchor() int {
	defer qt.Recovering("QTextCursor::anchor")

	if ptr.Pointer() != nil {
		return int(C.QTextCursor_Anchor(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextCursor) AtBlockEnd() bool {
	defer qt.Recovering("QTextCursor::atBlockEnd")

	if ptr.Pointer() != nil {
		return C.QTextCursor_AtBlockEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextCursor) AtBlockStart() bool {
	defer qt.Recovering("QTextCursor::atBlockStart")

	if ptr.Pointer() != nil {
		return C.QTextCursor_AtBlockStart(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextCursor) AtEnd() bool {
	defer qt.Recovering("QTextCursor::atEnd")

	if ptr.Pointer() != nil {
		return C.QTextCursor_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextCursor) AtStart() bool {
	defer qt.Recovering("QTextCursor::atStart")

	if ptr.Pointer() != nil {
		return C.QTextCursor_AtStart(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextCursor) BeginEditBlock() {
	defer qt.Recovering("QTextCursor::beginEditBlock")

	if ptr.Pointer() != nil {
		C.QTextCursor_BeginEditBlock(ptr.Pointer())
	}
}

func (ptr *QTextCursor) BlockNumber() int {
	defer qt.Recovering("QTextCursor::blockNumber")

	if ptr.Pointer() != nil {
		return int(C.QTextCursor_BlockNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextCursor) ClearSelection() {
	defer qt.Recovering("QTextCursor::clearSelection")

	if ptr.Pointer() != nil {
		C.QTextCursor_ClearSelection(ptr.Pointer())
	}
}

func (ptr *QTextCursor) ColumnNumber() int {
	defer qt.Recovering("QTextCursor::columnNumber")

	if ptr.Pointer() != nil {
		return int(C.QTextCursor_ColumnNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextCursor) CreateList2(style QTextListFormat__Style) *QTextList {
	defer qt.Recovering("QTextCursor::createList")

	if ptr.Pointer() != nil {
		return NewQTextListFromPointer(C.QTextCursor_CreateList2(ptr.Pointer(), C.int(style)))
	}
	return nil
}

func (ptr *QTextCursor) CreateList(format QTextListFormat_ITF) *QTextList {
	defer qt.Recovering("QTextCursor::createList")

	if ptr.Pointer() != nil {
		return NewQTextListFromPointer(C.QTextCursor_CreateList(ptr.Pointer(), PointerFromQTextListFormat(format)))
	}
	return nil
}

func (ptr *QTextCursor) CurrentFrame() *QTextFrame {
	defer qt.Recovering("QTextCursor::currentFrame")

	if ptr.Pointer() != nil {
		return NewQTextFrameFromPointer(C.QTextCursor_CurrentFrame(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextCursor) CurrentList() *QTextList {
	defer qt.Recovering("QTextCursor::currentList")

	if ptr.Pointer() != nil {
		return NewQTextListFromPointer(C.QTextCursor_CurrentList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextCursor) CurrentTable() *QTextTable {
	defer qt.Recovering("QTextCursor::currentTable")

	if ptr.Pointer() != nil {
		return NewQTextTableFromPointer(C.QTextCursor_CurrentTable(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextCursor) DeleteChar() {
	defer qt.Recovering("QTextCursor::deleteChar")

	if ptr.Pointer() != nil {
		C.QTextCursor_DeleteChar(ptr.Pointer())
	}
}

func (ptr *QTextCursor) DeletePreviousChar() {
	defer qt.Recovering("QTextCursor::deletePreviousChar")

	if ptr.Pointer() != nil {
		C.QTextCursor_DeletePreviousChar(ptr.Pointer())
	}
}

func (ptr *QTextCursor) Document() *QTextDocument {
	defer qt.Recovering("QTextCursor::document")

	if ptr.Pointer() != nil {
		return NewQTextDocumentFromPointer(C.QTextCursor_Document(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextCursor) EndEditBlock() {
	defer qt.Recovering("QTextCursor::endEditBlock")

	if ptr.Pointer() != nil {
		C.QTextCursor_EndEditBlock(ptr.Pointer())
	}
}

func (ptr *QTextCursor) HasComplexSelection() bool {
	defer qt.Recovering("QTextCursor::hasComplexSelection")

	if ptr.Pointer() != nil {
		return C.QTextCursor_HasComplexSelection(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextCursor) HasSelection() bool {
	defer qt.Recovering("QTextCursor::hasSelection")

	if ptr.Pointer() != nil {
		return C.QTextCursor_HasSelection(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextCursor) InsertBlock() {
	defer qt.Recovering("QTextCursor::insertBlock")

	if ptr.Pointer() != nil {
		C.QTextCursor_InsertBlock(ptr.Pointer())
	}
}

func (ptr *QTextCursor) InsertBlock2(format QTextBlockFormat_ITF) {
	defer qt.Recovering("QTextCursor::insertBlock")

	if ptr.Pointer() != nil {
		C.QTextCursor_InsertBlock2(ptr.Pointer(), PointerFromQTextBlockFormat(format))
	}
}

func (ptr *QTextCursor) InsertFragment(fragment QTextDocumentFragment_ITF) {
	defer qt.Recovering("QTextCursor::insertFragment")

	if ptr.Pointer() != nil {
		C.QTextCursor_InsertFragment(ptr.Pointer(), PointerFromQTextDocumentFragment(fragment))
	}
}

func (ptr *QTextCursor) InsertFrame(format QTextFrameFormat_ITF) *QTextFrame {
	defer qt.Recovering("QTextCursor::insertFrame")

	if ptr.Pointer() != nil {
		return NewQTextFrameFromPointer(C.QTextCursor_InsertFrame(ptr.Pointer(), PointerFromQTextFrameFormat(format)))
	}
	return nil
}

func (ptr *QTextCursor) InsertHtml(html string) {
	defer qt.Recovering("QTextCursor::insertHtml")

	if ptr.Pointer() != nil {
		C.QTextCursor_InsertHtml(ptr.Pointer(), C.CString(html))
	}
}

func (ptr *QTextCursor) InsertImage4(image QImage_ITF, name string) {
	defer qt.Recovering("QTextCursor::insertImage")

	if ptr.Pointer() != nil {
		C.QTextCursor_InsertImage4(ptr.Pointer(), PointerFromQImage(image), C.CString(name))
	}
}

func (ptr *QTextCursor) InsertImage3(name string) {
	defer qt.Recovering("QTextCursor::insertImage")

	if ptr.Pointer() != nil {
		C.QTextCursor_InsertImage3(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QTextCursor) InsertImage(format QTextImageFormat_ITF) {
	defer qt.Recovering("QTextCursor::insertImage")

	if ptr.Pointer() != nil {
		C.QTextCursor_InsertImage(ptr.Pointer(), PointerFromQTextImageFormat(format))
	}
}

func (ptr *QTextCursor) InsertImage2(format QTextImageFormat_ITF, alignment QTextFrameFormat__Position) {
	defer qt.Recovering("QTextCursor::insertImage")

	if ptr.Pointer() != nil {
		C.QTextCursor_InsertImage2(ptr.Pointer(), PointerFromQTextImageFormat(format), C.int(alignment))
	}
}

func (ptr *QTextCursor) InsertList2(style QTextListFormat__Style) *QTextList {
	defer qt.Recovering("QTextCursor::insertList")

	if ptr.Pointer() != nil {
		return NewQTextListFromPointer(C.QTextCursor_InsertList2(ptr.Pointer(), C.int(style)))
	}
	return nil
}

func (ptr *QTextCursor) InsertList(format QTextListFormat_ITF) *QTextList {
	defer qt.Recovering("QTextCursor::insertList")

	if ptr.Pointer() != nil {
		return NewQTextListFromPointer(C.QTextCursor_InsertList(ptr.Pointer(), PointerFromQTextListFormat(format)))
	}
	return nil
}

func (ptr *QTextCursor) InsertText(text string) {
	defer qt.Recovering("QTextCursor::insertText")

	if ptr.Pointer() != nil {
		C.QTextCursor_InsertText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QTextCursor) IsCopyOf(other QTextCursor_ITF) bool {
	defer qt.Recovering("QTextCursor::isCopyOf")

	if ptr.Pointer() != nil {
		return C.QTextCursor_IsCopyOf(ptr.Pointer(), PointerFromQTextCursor(other)) != 0
	}
	return false
}

func (ptr *QTextCursor) IsNull() bool {
	defer qt.Recovering("QTextCursor::isNull")

	if ptr.Pointer() != nil {
		return C.QTextCursor_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextCursor) JoinPreviousEditBlock() {
	defer qt.Recovering("QTextCursor::joinPreviousEditBlock")

	if ptr.Pointer() != nil {
		C.QTextCursor_JoinPreviousEditBlock(ptr.Pointer())
	}
}

func (ptr *QTextCursor) KeepPositionOnInsert() bool {
	defer qt.Recovering("QTextCursor::keepPositionOnInsert")

	if ptr.Pointer() != nil {
		return C.QTextCursor_KeepPositionOnInsert(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextCursor) MergeBlockCharFormat(modifier QTextCharFormat_ITF) {
	defer qt.Recovering("QTextCursor::mergeBlockCharFormat")

	if ptr.Pointer() != nil {
		C.QTextCursor_MergeBlockCharFormat(ptr.Pointer(), PointerFromQTextCharFormat(modifier))
	}
}

func (ptr *QTextCursor) MergeBlockFormat(modifier QTextBlockFormat_ITF) {
	defer qt.Recovering("QTextCursor::mergeBlockFormat")

	if ptr.Pointer() != nil {
		C.QTextCursor_MergeBlockFormat(ptr.Pointer(), PointerFromQTextBlockFormat(modifier))
	}
}

func (ptr *QTextCursor) MergeCharFormat(modifier QTextCharFormat_ITF) {
	defer qt.Recovering("QTextCursor::mergeCharFormat")

	if ptr.Pointer() != nil {
		C.QTextCursor_MergeCharFormat(ptr.Pointer(), PointerFromQTextCharFormat(modifier))
	}
}

func (ptr *QTextCursor) Position() int {
	defer qt.Recovering("QTextCursor::position")

	if ptr.Pointer() != nil {
		return int(C.QTextCursor_Position(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextCursor) PositionInBlock() int {
	defer qt.Recovering("QTextCursor::positionInBlock")

	if ptr.Pointer() != nil {
		return int(C.QTextCursor_PositionInBlock(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextCursor) RemoveSelectedText() {
	defer qt.Recovering("QTextCursor::removeSelectedText")

	if ptr.Pointer() != nil {
		C.QTextCursor_RemoveSelectedText(ptr.Pointer())
	}
}

func (ptr *QTextCursor) Select(selection QTextCursor__SelectionType) {
	defer qt.Recovering("QTextCursor::select")

	if ptr.Pointer() != nil {
		C.QTextCursor_Select(ptr.Pointer(), C.int(selection))
	}
}

func (ptr *QTextCursor) SelectedTableCells(firstRow int, numRows int, firstColumn int, numColumns int) {
	defer qt.Recovering("QTextCursor::selectedTableCells")

	if ptr.Pointer() != nil {
		C.QTextCursor_SelectedTableCells(ptr.Pointer(), C.int(firstRow), C.int(numRows), C.int(firstColumn), C.int(numColumns))
	}
}

func (ptr *QTextCursor) SelectedText() string {
	defer qt.Recovering("QTextCursor::selectedText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextCursor_SelectedText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextCursor) SelectionEnd() int {
	defer qt.Recovering("QTextCursor::selectionEnd")

	if ptr.Pointer() != nil {
		return int(C.QTextCursor_SelectionEnd(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextCursor) SelectionStart() int {
	defer qt.Recovering("QTextCursor::selectionStart")

	if ptr.Pointer() != nil {
		return int(C.QTextCursor_SelectionStart(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextCursor) SetBlockCharFormat(format QTextCharFormat_ITF) {
	defer qt.Recovering("QTextCursor::setBlockCharFormat")

	if ptr.Pointer() != nil {
		C.QTextCursor_SetBlockCharFormat(ptr.Pointer(), PointerFromQTextCharFormat(format))
	}
}

func (ptr *QTextCursor) SetBlockFormat(format QTextBlockFormat_ITF) {
	defer qt.Recovering("QTextCursor::setBlockFormat")

	if ptr.Pointer() != nil {
		C.QTextCursor_SetBlockFormat(ptr.Pointer(), PointerFromQTextBlockFormat(format))
	}
}

func (ptr *QTextCursor) SetCharFormat(format QTextCharFormat_ITF) {
	defer qt.Recovering("QTextCursor::setCharFormat")

	if ptr.Pointer() != nil {
		C.QTextCursor_SetCharFormat(ptr.Pointer(), PointerFromQTextCharFormat(format))
	}
}

func (ptr *QTextCursor) SetKeepPositionOnInsert(b bool) {
	defer qt.Recovering("QTextCursor::setKeepPositionOnInsert")

	if ptr.Pointer() != nil {
		C.QTextCursor_SetKeepPositionOnInsert(ptr.Pointer(), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QTextCursor) SetPosition(pos int, m QTextCursor__MoveMode) {
	defer qt.Recovering("QTextCursor::setPosition")

	if ptr.Pointer() != nil {
		C.QTextCursor_SetPosition(ptr.Pointer(), C.int(pos), C.int(m))
	}
}

func (ptr *QTextCursor) SetVerticalMovementX(x int) {
	defer qt.Recovering("QTextCursor::setVerticalMovementX")

	if ptr.Pointer() != nil {
		C.QTextCursor_SetVerticalMovementX(ptr.Pointer(), C.int(x))
	}
}

func (ptr *QTextCursor) SetVisualNavigation(b bool) {
	defer qt.Recovering("QTextCursor::setVisualNavigation")

	if ptr.Pointer() != nil {
		C.QTextCursor_SetVisualNavigation(ptr.Pointer(), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QTextCursor) Swap(other QTextCursor_ITF) {
	defer qt.Recovering("QTextCursor::swap")

	if ptr.Pointer() != nil {
		C.QTextCursor_Swap(ptr.Pointer(), PointerFromQTextCursor(other))
	}
}

func (ptr *QTextCursor) VerticalMovementX() int {
	defer qt.Recovering("QTextCursor::verticalMovementX")

	if ptr.Pointer() != nil {
		return int(C.QTextCursor_VerticalMovementX(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextCursor) VisualNavigation() bool {
	defer qt.Recovering("QTextCursor::visualNavigation")

	if ptr.Pointer() != nil {
		return C.QTextCursor_VisualNavigation(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextCursor) DestroyQTextCursor() {
	defer qt.Recovering("QTextCursor::~QTextCursor")

	if ptr.Pointer() != nil {
		C.QTextCursor_DestroyQTextCursor(ptr.Pointer())
	}
}
