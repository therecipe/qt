package gui

//#include "qtextitem.h"
import "C"
import (
	"unsafe"
)

type QTextItem struct {
	ptr unsafe.Pointer
}

type QTextItemITF interface {
	QTextItemPTR() *QTextItem
}

func (p *QTextItem) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextItem) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextItem(ptr QTextItemITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextItemPTR().Pointer()
	}
	return nil
}

func QTextItemFromPointer(ptr unsafe.Pointer) *QTextItem {
	var n = new(QTextItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextItem) QTextItemPTR() *QTextItem {
	return ptr
}

//QTextItem::RenderFlag
type QTextItem__RenderFlag int

var (
	QTextItem__RightToLeft = QTextItem__RenderFlag(0x1)
	QTextItem__Overline    = QTextItem__RenderFlag(0x10)
	QTextItem__Underline   = QTextItem__RenderFlag(0x20)
	QTextItem__StrikeOut   = QTextItem__RenderFlag(0x40)
	QTextItem__Dummy       = QTextItem__RenderFlag(0xffffffff)
)

func (ptr *QTextItem) RenderFlags() QTextItem__RenderFlag {
	if ptr.Pointer() != nil {
		return QTextItem__RenderFlag(C.QTextItem_RenderFlags(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextItem) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextItem_Text(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}
