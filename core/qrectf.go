package core

//#include "qrectf.h"
import "C"
import (
	"unsafe"
)

type QRectF struct {
	ptr unsafe.Pointer
}

type QRectFITF interface {
	QRectFPTR() *QRectF
}

func (p *QRectF) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QRectF) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQRectF(ptr QRectFITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRectFPTR().Pointer()
	}
	return nil
}

func QRectFFromPointer(ptr unsafe.Pointer) *QRectF {
	var n = new(QRectF)
	n.SetPointer(ptr)
	return n
}

func (ptr *QRectF) QRectFPTR() *QRectF {
	return ptr
}

func (ptr *QRectF) Contains(point QPointFITF) bool {
	if ptr.Pointer() != nil {
		return C.QRectF_Contains(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPointF(point))) != 0
	}
	return false
}

func (ptr *QRectF) Contains3(rectangle QRectFITF) bool {
	if ptr.Pointer() != nil {
		return C.QRectF_Contains3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQRectF(rectangle))) != 0
	}
	return false
}

func (ptr *QRectF) Intersects(rectangle QRectFITF) bool {
	if ptr.Pointer() != nil {
		return C.QRectF_Intersects(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQRectF(rectangle))) != 0
	}
	return false
}

func NewQRectF() *QRectF {
	return QRectFFromPointer(unsafe.Pointer(C.QRectF_NewQRectF()))
}

func NewQRectF3(topLeft QPointFITF, bottomRight QPointFITF) *QRectF {
	return QRectFFromPointer(unsafe.Pointer(C.QRectF_NewQRectF3(C.QtObjectPtr(PointerFromQPointF(topLeft)), C.QtObjectPtr(PointerFromQPointF(bottomRight)))))
}

func NewQRectF2(topLeft QPointFITF, size QSizeFITF) *QRectF {
	return QRectFFromPointer(unsafe.Pointer(C.QRectF_NewQRectF2(C.QtObjectPtr(PointerFromQPointF(topLeft)), C.QtObjectPtr(PointerFromQSizeF(size)))))
}

func NewQRectF5(rectangle QRectITF) *QRectF {
	return QRectFFromPointer(unsafe.Pointer(C.QRectF_NewQRectF5(C.QtObjectPtr(PointerFromQRect(rectangle)))))
}

func (ptr *QRectF) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QRectF_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRectF) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QRectF_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRectF) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QRectF_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRectF) MoveBottomLeft(position QPointFITF) {
	if ptr.Pointer() != nil {
		C.QRectF_MoveBottomLeft(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPointF(position)))
	}
}

func (ptr *QRectF) MoveBottomRight(position QPointFITF) {
	if ptr.Pointer() != nil {
		C.QRectF_MoveBottomRight(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPointF(position)))
	}
}

func (ptr *QRectF) MoveCenter(position QPointFITF) {
	if ptr.Pointer() != nil {
		C.QRectF_MoveCenter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPointF(position)))
	}
}

func (ptr *QRectF) MoveTo2(position QPointFITF) {
	if ptr.Pointer() != nil {
		C.QRectF_MoveTo2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPointF(position)))
	}
}

func (ptr *QRectF) MoveTopLeft(position QPointFITF) {
	if ptr.Pointer() != nil {
		C.QRectF_MoveTopLeft(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPointF(position)))
	}
}

func (ptr *QRectF) MoveTopRight(position QPointFITF) {
	if ptr.Pointer() != nil {
		C.QRectF_MoveTopRight(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPointF(position)))
	}
}

func (ptr *QRectF) SetBottomLeft(position QPointFITF) {
	if ptr.Pointer() != nil {
		C.QRectF_SetBottomLeft(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPointF(position)))
	}
}

func (ptr *QRectF) SetBottomRight(position QPointFITF) {
	if ptr.Pointer() != nil {
		C.QRectF_SetBottomRight(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPointF(position)))
	}
}

func (ptr *QRectF) SetSize(size QSizeFITF) {
	if ptr.Pointer() != nil {
		C.QRectF_SetSize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSizeF(size)))
	}
}

func (ptr *QRectF) SetTopLeft(position QPointFITF) {
	if ptr.Pointer() != nil {
		C.QRectF_SetTopLeft(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPointF(position)))
	}
}

func (ptr *QRectF) SetTopRight(position QPointFITF) {
	if ptr.Pointer() != nil {
		C.QRectF_SetTopRight(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPointF(position)))
	}
}

func (ptr *QRectF) Translate2(offset QPointFITF) {
	if ptr.Pointer() != nil {
		C.QRectF_Translate2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPointF(offset)))
	}
}
