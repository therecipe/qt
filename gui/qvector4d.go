package gui

//#include "qvector4d.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QVector4D struct {
	ptr unsafe.Pointer
}

type QVector4DITF interface {
	QVector4DPTR() *QVector4D
}

func (p *QVector4D) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QVector4D) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQVector4D(ptr QVector4DITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVector4DPTR().Pointer()
	}
	return nil
}

func QVector4DFromPointer(ptr unsafe.Pointer) *QVector4D {
	var n = new(QVector4D)
	n.SetPointer(ptr)
	return n
}

func (ptr *QVector4D) QVector4DPTR() *QVector4D {
	return ptr
}

func NewQVector4D() *QVector4D {
	return QVector4DFromPointer(unsafe.Pointer(C.QVector4D_NewQVector4D()))
}

func NewQVector4D4(point core.QPointITF) *QVector4D {
	return QVector4DFromPointer(unsafe.Pointer(C.QVector4D_NewQVector4D4(C.QtObjectPtr(core.PointerFromQPoint(point)))))
}

func NewQVector4D5(point core.QPointFITF) *QVector4D {
	return QVector4DFromPointer(unsafe.Pointer(C.QVector4D_NewQVector4D5(C.QtObjectPtr(core.PointerFromQPointF(point)))))
}

func NewQVector4D6(vector QVector2DITF) *QVector4D {
	return QVector4DFromPointer(unsafe.Pointer(C.QVector4D_NewQVector4D6(C.QtObjectPtr(PointerFromQVector2D(vector)))))
}

func NewQVector4D8(vector QVector3DITF) *QVector4D {
	return QVector4DFromPointer(unsafe.Pointer(C.QVector4D_NewQVector4D8(C.QtObjectPtr(PointerFromQVector3D(vector)))))
}

func (ptr *QVector4D) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QVector4D_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QVector4D) Normalize() {
	if ptr.Pointer() != nil {
		C.QVector4D_Normalize(C.QtObjectPtr(ptr.Pointer()))
	}
}
