package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTextBlock struct {
	ptr unsafe.Pointer
}

type QTextBlock_ITF interface {
	QTextBlock_PTR() *QTextBlock
}

func (p *QTextBlock) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextBlock) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextBlock(ptr QTextBlock_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextBlock_PTR().Pointer()
	}
	return nil
}

func NewQTextBlockFromPointer(ptr unsafe.Pointer) *QTextBlock {
	var n = new(QTextBlock)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextBlock) QTextBlock_PTR() *QTextBlock {
	return ptr
}

func (ptr *QTextBlock) IsValid() bool {
	defer qt.Recovering("QTextBlock::isValid")

	if ptr.Pointer() != nil {
		return C.QTextBlock_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func NewQTextBlock(other QTextBlock_ITF) *QTextBlock {
	defer qt.Recovering("QTextBlock::QTextBlock")

	return NewQTextBlockFromPointer(C.QTextBlock_NewQTextBlock(PointerFromQTextBlock(other)))
}

func (ptr *QTextBlock) BlockFormatIndex() int {
	defer qt.Recovering("QTextBlock::blockFormatIndex")

	if ptr.Pointer() != nil {
		return int(C.QTextBlock_BlockFormatIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextBlock) CharFormatIndex() int {
	defer qt.Recovering("QTextBlock::charFormatIndex")

	if ptr.Pointer() != nil {
		return int(C.QTextBlock_CharFormatIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextBlock) ClearLayout() {
	defer qt.Recovering("QTextBlock::clearLayout")

	if ptr.Pointer() != nil {
		C.QTextBlock_ClearLayout(ptr.Pointer())
	}
}

func (ptr *QTextBlock) Contains(position int) bool {
	defer qt.Recovering("QTextBlock::contains")

	if ptr.Pointer() != nil {
		return C.QTextBlock_Contains(ptr.Pointer(), C.int(position)) != 0
	}
	return false
}

func (ptr *QTextBlock) BlockNumber() int {
	defer qt.Recovering("QTextBlock::blockNumber")

	if ptr.Pointer() != nil {
		return int(C.QTextBlock_BlockNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextBlock) Document() *QTextDocument {
	defer qt.Recovering("QTextBlock::document")

	if ptr.Pointer() != nil {
		return NewQTextDocumentFromPointer(C.QTextBlock_Document(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextBlock) FirstLineNumber() int {
	defer qt.Recovering("QTextBlock::firstLineNumber")

	if ptr.Pointer() != nil {
		return int(C.QTextBlock_FirstLineNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextBlock) IsVisible() bool {
	defer qt.Recovering("QTextBlock::isVisible")

	if ptr.Pointer() != nil {
		return C.QTextBlock_IsVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextBlock) Layout() *QTextLayout {
	defer qt.Recovering("QTextBlock::layout")

	if ptr.Pointer() != nil {
		return NewQTextLayoutFromPointer(C.QTextBlock_Layout(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextBlock) Length() int {
	defer qt.Recovering("QTextBlock::length")

	if ptr.Pointer() != nil {
		return int(C.QTextBlock_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextBlock) LineCount() int {
	defer qt.Recovering("QTextBlock::lineCount")

	if ptr.Pointer() != nil {
		return int(C.QTextBlock_LineCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextBlock) Position() int {
	defer qt.Recovering("QTextBlock::position")

	if ptr.Pointer() != nil {
		return int(C.QTextBlock_Position(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextBlock) Revision() int {
	defer qt.Recovering("QTextBlock::revision")

	if ptr.Pointer() != nil {
		return int(C.QTextBlock_Revision(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextBlock) SetLineCount(count int) {
	defer qt.Recovering("QTextBlock::setLineCount")

	if ptr.Pointer() != nil {
		C.QTextBlock_SetLineCount(ptr.Pointer(), C.int(count))
	}
}

func (ptr *QTextBlock) SetRevision(rev int) {
	defer qt.Recovering("QTextBlock::setRevision")

	if ptr.Pointer() != nil {
		C.QTextBlock_SetRevision(ptr.Pointer(), C.int(rev))
	}
}

func (ptr *QTextBlock) SetUserData(data QTextBlockUserData_ITF) {
	defer qt.Recovering("QTextBlock::setUserData")

	if ptr.Pointer() != nil {
		C.QTextBlock_SetUserData(ptr.Pointer(), PointerFromQTextBlockUserData(data))
	}
}

func (ptr *QTextBlock) SetUserState(state int) {
	defer qt.Recovering("QTextBlock::setUserState")

	if ptr.Pointer() != nil {
		C.QTextBlock_SetUserState(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QTextBlock) SetVisible(visible bool) {
	defer qt.Recovering("QTextBlock::setVisible")

	if ptr.Pointer() != nil {
		C.QTextBlock_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QTextBlock) Text() string {
	defer qt.Recovering("QTextBlock::text")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextBlock_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextBlock) TextDirection() core.Qt__LayoutDirection {
	defer qt.Recovering("QTextBlock::textDirection")

	if ptr.Pointer() != nil {
		return core.Qt__LayoutDirection(C.QTextBlock_TextDirection(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextBlock) TextList() *QTextList {
	defer qt.Recovering("QTextBlock::textList")

	if ptr.Pointer() != nil {
		return NewQTextListFromPointer(C.QTextBlock_TextList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextBlock) UserData() *QTextBlockUserData {
	defer qt.Recovering("QTextBlock::userData")

	if ptr.Pointer() != nil {
		return NewQTextBlockUserDataFromPointer(C.QTextBlock_UserData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextBlock) UserState() int {
	defer qt.Recovering("QTextBlock::userState")

	if ptr.Pointer() != nil {
		return int(C.QTextBlock_UserState(ptr.Pointer()))
	}
	return 0
}
