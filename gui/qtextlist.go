package gui

//#include "qtextlist.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QTextList struct {
	QTextBlockGroup
}

type QTextList_ITF interface {
	QTextBlockGroup_ITF
	QTextList_PTR() *QTextList
}

func PointerFromQTextList(ptr QTextList_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextList_PTR().Pointer()
	}
	return nil
}

func NewQTextListFromPointer(ptr unsafe.Pointer) *QTextList {
	var n = new(QTextList)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QTextList_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTextList) QTextList_PTR() *QTextList {
	return ptr
}

func (ptr *QTextList) ItemNumber(block QTextBlock_ITF) int {
	if ptr.Pointer() != nil {
		return int(C.QTextList_ItemNumber(ptr.Pointer(), PointerFromQTextBlock(block)))
	}
	return 0
}

func (ptr *QTextList) ItemText(block QTextBlock_ITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextList_ItemText(ptr.Pointer(), PointerFromQTextBlock(block)))
	}
	return ""
}

func (ptr *QTextList) Add(block QTextBlock_ITF) {
	if ptr.Pointer() != nil {
		C.QTextList_Add(ptr.Pointer(), PointerFromQTextBlock(block))
	}
}

func (ptr *QTextList) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QTextList_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextList) RemoveItem(i int) {
	if ptr.Pointer() != nil {
		C.QTextList_RemoveItem(ptr.Pointer(), C.int(i))
	}
}

func (ptr *QTextList) SetFormat(format QTextListFormat_ITF) {
	if ptr.Pointer() != nil {
		C.QTextList_SetFormat(ptr.Pointer(), PointerFromQTextListFormat(format))
	}
}
