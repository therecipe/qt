package gui

//#include "qquaternion.h"
import "C"
import (
	"unsafe"
)

type QQuaternion struct {
	ptr unsafe.Pointer
}

type QQuaternion_ITF interface {
	QQuaternion_PTR() *QQuaternion
}

func (p *QQuaternion) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQuaternion) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQuaternion(ptr QQuaternion_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuaternion_PTR().Pointer()
	}
	return nil
}

func NewQQuaternionFromPointer(ptr unsafe.Pointer) *QQuaternion {
	var n = new(QQuaternion)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQuaternion) QQuaternion_PTR() *QQuaternion {
	return ptr
}

func NewQQuaternion() *QQuaternion {
	return NewQQuaternionFromPointer(C.QQuaternion_NewQQuaternion())
}

func NewQQuaternion5(vector QVector4D_ITF) *QQuaternion {
	return NewQQuaternionFromPointer(C.QQuaternion_NewQQuaternion5(PointerFromQVector4D(vector)))
}

func (ptr *QQuaternion) GetAxes(xAxis QVector3D_ITF, yAxis QVector3D_ITF, zAxis QVector3D_ITF) {
	if ptr.Pointer() != nil {
		C.QQuaternion_GetAxes(ptr.Pointer(), PointerFromQVector3D(xAxis), PointerFromQVector3D(yAxis), PointerFromQVector3D(zAxis))
	}
}

func (ptr *QQuaternion) IsIdentity() bool {
	if ptr.Pointer() != nil {
		return C.QQuaternion_IsIdentity(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuaternion) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QQuaternion_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuaternion) Normalize() {
	if ptr.Pointer() != nil {
		C.QQuaternion_Normalize(ptr.Pointer())
	}
}

func (ptr *QQuaternion) SetVector(vector QVector3D_ITF) {
	if ptr.Pointer() != nil {
		C.QQuaternion_SetVector(ptr.Pointer(), PointerFromQVector3D(vector))
	}
}
