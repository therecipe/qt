package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QVector3D struct {
	ptr unsafe.Pointer
}

type QVector3D_ITF interface {
	QVector3D_PTR() *QVector3D
}

func (p *QVector3D) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QVector3D) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQVector3D(ptr QVector3D_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVector3D_PTR().Pointer()
	}
	return nil
}

func NewQVector3DFromPointer(ptr unsafe.Pointer) *QVector3D {
	var n = new(QVector3D)
	n.SetPointer(ptr)
	return n
}

func (ptr *QVector3D) QVector3D_PTR() *QVector3D {
	return ptr
}

func NewQVector3D() *QVector3D {
	defer qt.Recovering("QVector3D::QVector3D")

	return NewQVector3DFromPointer(C.QVector3D_NewQVector3D())
}

func NewQVector3D4(point core.QPoint_ITF) *QVector3D {
	defer qt.Recovering("QVector3D::QVector3D")

	return NewQVector3DFromPointer(C.QVector3D_NewQVector3D4(core.PointerFromQPoint(point)))
}

func NewQVector3D5(point core.QPointF_ITF) *QVector3D {
	defer qt.Recovering("QVector3D::QVector3D")

	return NewQVector3DFromPointer(C.QVector3D_NewQVector3D5(core.PointerFromQPointF(point)))
}

func NewQVector3D6(vector QVector2D_ITF) *QVector3D {
	defer qt.Recovering("QVector3D::QVector3D")

	return NewQVector3DFromPointer(C.QVector3D_NewQVector3D6(PointerFromQVector2D(vector)))
}

func NewQVector3D8(vector QVector4D_ITF) *QVector3D {
	defer qt.Recovering("QVector3D::QVector3D")

	return NewQVector3DFromPointer(C.QVector3D_NewQVector3D8(PointerFromQVector4D(vector)))
}

func (ptr *QVector3D) IsNull() bool {
	defer qt.Recovering("QVector3D::isNull")

	if ptr.Pointer() != nil {
		return C.QVector3D_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVector3D) Normalize() {
	defer qt.Recovering("QVector3D::normalize")

	if ptr.Pointer() != nil {
		C.QVector3D_Normalize(ptr.Pointer())
	}
}

func (ptr *QVector3D) ToPoint() *core.QPoint {
	defer qt.Recovering("QVector3D::toPoint")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QVector3D_ToPoint(ptr.Pointer()))
	}
	return nil
}
