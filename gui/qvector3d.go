package gui

//#include "qvector3d.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QVector3D struct {
	ptr unsafe.Pointer
}

type QVector3DITF interface {
	QVector3DPTR() *QVector3D
}

func (p *QVector3D) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QVector3D) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQVector3D(ptr QVector3DITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVector3DPTR().Pointer()
	}
	return nil
}

func QVector3DFromPointer(ptr unsafe.Pointer) *QVector3D {
	var n = new(QVector3D)
	n.SetPointer(ptr)
	return n
}

func (ptr *QVector3D) QVector3DPTR() *QVector3D {
	return ptr
}

func NewQVector3D() *QVector3D {
	return QVector3DFromPointer(unsafe.Pointer(C.QVector3D_NewQVector3D()))
}

func NewQVector3D4(point core.QPointITF) *QVector3D {
	return QVector3DFromPointer(unsafe.Pointer(C.QVector3D_NewQVector3D4(C.QtObjectPtr(core.PointerFromQPoint(point)))))
}

func NewQVector3D5(point core.QPointFITF) *QVector3D {
	return QVector3DFromPointer(unsafe.Pointer(C.QVector3D_NewQVector3D5(C.QtObjectPtr(core.PointerFromQPointF(point)))))
}

func NewQVector3D6(vector QVector2DITF) *QVector3D {
	return QVector3DFromPointer(unsafe.Pointer(C.QVector3D_NewQVector3D6(C.QtObjectPtr(PointerFromQVector2D(vector)))))
}

func NewQVector3D8(vector QVector4DITF) *QVector3D {
	return QVector3DFromPointer(unsafe.Pointer(C.QVector3D_NewQVector3D8(C.QtObjectPtr(PointerFromQVector4D(vector)))))
}

func (ptr *QVector3D) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QVector3D_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QVector3D) Normalize() {
	if ptr.Pointer() != nil {
		C.QVector3D_Normalize(C.QtObjectPtr(ptr.Pointer()))
	}
}
