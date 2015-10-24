package core

//#include "qpointf.h"
import "C"
import (
	"unsafe"
)

type QPointF struct {
	ptr unsafe.Pointer
}

type QPointFITF interface {
	QPointFPTR() *QPointF
}

func (p *QPointF) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPointF) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPointF(ptr QPointFITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPointFPTR().Pointer()
	}
	return nil
}

func QPointFFromPointer(ptr unsafe.Pointer) *QPointF {
	var n = new(QPointF)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPointF) QPointFPTR() *QPointF {
	return ptr
}

func NewQPointF() *QPointF {
	return QPointFFromPointer(unsafe.Pointer(C.QPointF_NewQPointF()))
}

func NewQPointF2(point QPointITF) *QPointF {
	return QPointFFromPointer(unsafe.Pointer(C.QPointF_NewQPointF2(C.QtObjectPtr(PointerFromQPoint(point)))))
}

func (ptr *QPointF) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QPointF_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}
