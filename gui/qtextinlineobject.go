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

type QTextInlineObject_ITF interface {
	QTextInlineObject_PTR() *QTextInlineObject
}

func (p *QTextInlineObject) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextInlineObject) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextInlineObject(ptr QTextInlineObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextInlineObject_PTR().Pointer()
	}
	return nil
}

func NewQTextInlineObjectFromPointer(ptr unsafe.Pointer) *QTextInlineObject {
	var n = new(QTextInlineObject)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextInlineObject) QTextInlineObject_PTR() *QTextInlineObject {
	return ptr
}

func (ptr *QTextInlineObject) Ascent() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextInlineObject_Ascent(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextInlineObject) Descent() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextInlineObject_Descent(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextInlineObject) FormatIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QTextInlineObject_FormatIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextInlineObject) Height() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextInlineObject_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextInlineObject) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QTextInlineObject_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextInlineObject) SetAscent(a float64) {
	if ptr.Pointer() != nil {
		C.QTextInlineObject_SetAscent(ptr.Pointer(), C.double(a))
	}
}

func (ptr *QTextInlineObject) SetDescent(d float64) {
	if ptr.Pointer() != nil {
		C.QTextInlineObject_SetDescent(ptr.Pointer(), C.double(d))
	}
}

func (ptr *QTextInlineObject) SetWidth(w float64) {
	if ptr.Pointer() != nil {
		C.QTextInlineObject_SetWidth(ptr.Pointer(), C.double(w))
	}
}

func (ptr *QTextInlineObject) TextDirection() core.Qt__LayoutDirection {
	if ptr.Pointer() != nil {
		return core.Qt__LayoutDirection(C.QTextInlineObject_TextDirection(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextInlineObject) TextPosition() int {
	if ptr.Pointer() != nil {
		return int(C.QTextInlineObject_TextPosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextInlineObject) Width() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextInlineObject_Width(ptr.Pointer()))
	}
	return 0
}
