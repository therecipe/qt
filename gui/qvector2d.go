package gui

//#include "qvector2d.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QVector2D struct {
	ptr unsafe.Pointer
}

type QVector2D_ITF interface {
	QVector2D_PTR() *QVector2D
}

func (p *QVector2D) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QVector2D) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQVector2D(ptr QVector2D_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVector2D_PTR().Pointer()
	}
	return nil
}

func NewQVector2DFromPointer(ptr unsafe.Pointer) *QVector2D {
	var n = new(QVector2D)
	n.SetPointer(ptr)
	return n
}

func (ptr *QVector2D) QVector2D_PTR() *QVector2D {
	return ptr
}

func NewQVector2D() *QVector2D {
	return NewQVector2DFromPointer(C.QVector2D_NewQVector2D())
}

func NewQVector2D4(point core.QPoint_ITF) *QVector2D {
	return NewQVector2DFromPointer(C.QVector2D_NewQVector2D4(core.PointerFromQPoint(point)))
}

func NewQVector2D5(point core.QPointF_ITF) *QVector2D {
	return NewQVector2DFromPointer(C.QVector2D_NewQVector2D5(core.PointerFromQPointF(point)))
}

func NewQVector2D6(vector QVector3D_ITF) *QVector2D {
	return NewQVector2DFromPointer(C.QVector2D_NewQVector2D6(PointerFromQVector3D(vector)))
}

func NewQVector2D7(vector QVector4D_ITF) *QVector2D {
	return NewQVector2DFromPointer(C.QVector2D_NewQVector2D7(PointerFromQVector4D(vector)))
}

func (ptr *QVector2D) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QVector2D_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVector2D) Normalize() {
	if ptr.Pointer() != nil {
		C.QVector2D_Normalize(ptr.Pointer())
	}
}
