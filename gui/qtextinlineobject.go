package gui

//#include "qtextinlineobject.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTextInlineObject struct {
	ptr unsafe.Pointer
}

type QTextInlineObjectITF interface {
	QTextInlineObjectPTR() *QTextInlineObject
}

func (p *QTextInlineObject) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextInlineObject) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextInlineObject(ptr QTextInlineObjectITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextInlineObjectPTR().Pointer()
	}
	return nil
}

func QTextInlineObjectFromPointer(ptr unsafe.Pointer) *QTextInlineObject {
	var n = new(QTextInlineObject)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextInlineObject) QTextInlineObjectPTR() *QTextInlineObject {
	return ptr
}

func (ptr *QTextInlineObject) FormatIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QTextInlineObject_FormatIndex(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextInlineObject) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QTextInlineObject_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextInlineObject) TextDirection() core.Qt__LayoutDirection {
	if ptr.Pointer() != nil {
		return core.Qt__LayoutDirection(C.QTextInlineObject_TextDirection(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextInlineObject) TextPosition() int {
	if ptr.Pointer() != nil {
		return int(C.QTextInlineObject_TextPosition(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
