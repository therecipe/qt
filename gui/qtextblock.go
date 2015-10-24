package gui

//#include "qtextblock.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTextBlock struct {
	ptr unsafe.Pointer
}

type QTextBlockITF interface {
	QTextBlockPTR() *QTextBlock
}

func (p *QTextBlock) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextBlock) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextBlock(ptr QTextBlockITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextBlockPTR().Pointer()
	}
	return nil
}

func QTextBlockFromPointer(ptr unsafe.Pointer) *QTextBlock {
	var n = new(QTextBlock)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextBlock) QTextBlockPTR() *QTextBlock {
	return ptr
}

func (ptr *QTextBlock) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QTextBlock_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func NewQTextBlock(other QTextBlockITF) *QTextBlock {
	return QTextBlockFromPointer(unsafe.Pointer(C.QTextBlock_NewQTextBlock(C.QtObjectPtr(PointerFromQTextBlock(other)))))
}

func (ptr *QTextBlock) BlockFormatIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QTextBlock_BlockFormatIndex(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextBlock) CharFormatIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QTextBlock_CharFormatIndex(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextBlock) ClearLayout() {
	if ptr.Pointer() != nil {
		C.QTextBlock_ClearLayout(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextBlock) Contains(position int) bool {
	if ptr.Pointer() != nil {
		return C.QTextBlock_Contains(C.QtObjectPtr(ptr.Pointer()), C.int(position)) != 0
	}
	return false
}

func (ptr *QTextBlock) BlockNumber() int {
	if ptr.Pointer() != nil {
		return int(C.QTextBlock_BlockNumber(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextBlock) Document() *QTextDocument {
	if ptr.Pointer() != nil {
		return QTextDocumentFromPointer(unsafe.Pointer(C.QTextBlock_Document(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTextBlock) FirstLineNumber() int {
	if ptr.Pointer() != nil {
		return int(C.QTextBlock_FirstLineNumber(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextBlock) IsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QTextBlock_IsVisible(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextBlock) Layout() *QTextLayout {
	if ptr.Pointer() != nil {
		return QTextLayoutFromPointer(unsafe.Pointer(C.QTextBlock_Layout(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTextBlock) Length() int {
	if ptr.Pointer() != nil {
		return int(C.QTextBlock_Length(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextBlock) LineCount() int {
	if ptr.Pointer() != nil {
		return int(C.QTextBlock_LineCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextBlock) Position() int {
	if ptr.Pointer() != nil {
		return int(C.QTextBlock_Position(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextBlock) Revision() int {
	if ptr.Pointer() != nil {
		return int(C.QTextBlock_Revision(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextBlock) SetLineCount(count int) {
	if ptr.Pointer() != nil {
		C.QTextBlock_SetLineCount(C.QtObjectPtr(ptr.Pointer()), C.int(count))
	}
}

func (ptr *QTextBlock) SetRevision(rev int) {
	if ptr.Pointer() != nil {
		C.QTextBlock_SetRevision(C.QtObjectPtr(ptr.Pointer()), C.int(rev))
	}
}

func (ptr *QTextBlock) SetUserData(data QTextBlockUserDataITF) {
	if ptr.Pointer() != nil {
		C.QTextBlock_SetUserData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextBlockUserData(data)))
	}
}

func (ptr *QTextBlock) SetUserState(state int) {
	if ptr.Pointer() != nil {
		C.QTextBlock_SetUserState(C.QtObjectPtr(ptr.Pointer()), C.int(state))
	}
}

func (ptr *QTextBlock) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QTextBlock_SetVisible(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QTextBlock) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextBlock_Text(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTextBlock) TextDirection() core.Qt__LayoutDirection {
	if ptr.Pointer() != nil {
		return core.Qt__LayoutDirection(C.QTextBlock_TextDirection(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextBlock) TextList() *QTextList {
	if ptr.Pointer() != nil {
		return QTextListFromPointer(unsafe.Pointer(C.QTextBlock_TextList(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTextBlock) UserData() *QTextBlockUserData {
	if ptr.Pointer() != nil {
		return QTextBlockUserDataFromPointer(unsafe.Pointer(C.QTextBlock_UserData(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTextBlock) UserState() int {
	if ptr.Pointer() != nil {
		return int(C.QTextBlock_UserState(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
