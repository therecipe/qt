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

type QVector2DITF interface {
	QVector2DPTR() *QVector2D
}

func (p *QVector2D) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QVector2D) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQVector2D(ptr QVector2DITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVector2DPTR().Pointer()
	}
	return nil
}

func QVector2DFromPointer(ptr unsafe.Pointer) *QVector2D {
	var n = new(QVector2D)
	n.SetPointer(ptr)
	return n
}

func (ptr *QVector2D) QVector2DPTR() *QVector2D {
	return ptr
}

func NewQVector2D() *QVector2D {
	return QVector2DFromPointer(unsafe.Pointer(C.QVector2D_NewQVector2D()))
}

func NewQVector2D4(point core.QPointITF) *QVector2D {
	return QVector2DFromPointer(unsafe.Pointer(C.QVector2D_NewQVector2D4(C.QtObjectPtr(core.PointerFromQPoint(point)))))
}

func NewQVector2D5(point core.QPointFITF) *QVector2D {
	return QVector2DFromPointer(unsafe.Pointer(C.QVector2D_NewQVector2D5(C.QtObjectPtr(core.PointerFromQPointF(point)))))
}

func NewQVector2D6(vector QVector3DITF) *QVector2D {
	return QVector2DFromPointer(unsafe.Pointer(C.QVector2D_NewQVector2D6(C.QtObjectPtr(PointerFromQVector3D(vector)))))
}

func NewQVector2D7(vector QVector4DITF) *QVector2D {
	return QVector2DFromPointer(unsafe.Pointer(C.QVector2D_NewQVector2D7(C.QtObjectPtr(PointerFromQVector4D(vector)))))
}

func (ptr *QVector2D) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QVector2D_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QVector2D) Normalize() {
	if ptr.Pointer() != nil {
		C.QVector2D_Normalize(C.QtObjectPtr(ptr.Pointer()))
	}
}
