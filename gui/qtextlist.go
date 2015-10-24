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

type QTextListITF interface {
	QTextBlockGroupITF
	QTextListPTR() *QTextList
}

func PointerFromQTextList(ptr QTextListITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextListPTR().Pointer()
	}
	return nil
}

func QTextListFromPointer(ptr unsafe.Pointer) *QTextList {
	var n = new(QTextList)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTextList_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTextList) QTextListPTR() *QTextList {
	return ptr
}

func (ptr *QTextList) ItemNumber(block QTextBlockITF) int {
	if ptr.Pointer() != nil {
		return int(C.QTextList_ItemNumber(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextBlock(block))))
	}
	return 0
}

func (ptr *QTextList) ItemText(block QTextBlockITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextList_ItemText(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextBlock(block))))
	}
	return ""
}

func (ptr *QTextList) Add(block QTextBlockITF) {
	if ptr.Pointer() != nil {
		C.QTextList_Add(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextBlock(block)))
	}
}

func (ptr *QTextList) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QTextList_Count(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextList) RemoveItem(i int) {
	if ptr.Pointer() != nil {
		C.QTextList_RemoveItem(C.QtObjectPtr(ptr.Pointer()), C.int(i))
	}
}

func (ptr *QTextList) SetFormat(format QTextListFormatITF) {
	if ptr.Pointer() != nil {
		C.QTextList_SetFormat(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextListFormat(format)))
	}
}
