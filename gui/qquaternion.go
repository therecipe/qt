package gui

//#include "qquaternion.h"
import "C"
import (
	"unsafe"
)

type QQuaternion struct {
	ptr unsafe.Pointer
}

type QQuaternionITF interface {
	QQuaternionPTR() *QQuaternion
}

func (p *QQuaternion) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQuaternion) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQuaternion(ptr QQuaternionITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuaternionPTR().Pointer()
	}
	return nil
}

func QQuaternionFromPointer(ptr unsafe.Pointer) *QQuaternion {
	var n = new(QQuaternion)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQuaternion) QQuaternionPTR() *QQuaternion {
	return ptr
}

func NewQQuaternion() *QQuaternion {
	return QQuaternionFromPointer(unsafe.Pointer(C.QQuaternion_NewQQuaternion()))
}

func NewQQuaternion5(vector QVector4DITF) *QQuaternion {
	return QQuaternionFromPointer(unsafe.Pointer(C.QQuaternion_NewQQuaternion5(C.QtObjectPtr(PointerFromQVector4D(vector)))))
}

func (ptr *QQuaternion) GetAxes(xAxis QVector3DITF, yAxis QVector3DITF, zAxis QVector3DITF) {
	if ptr.Pointer() != nil {
		C.QQuaternion_GetAxes(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQVector3D(xAxis)), C.QtObjectPtr(PointerFromQVector3D(yAxis)), C.QtObjectPtr(PointerFromQVector3D(zAxis)))
	}
}

func (ptr *QQuaternion) IsIdentity() bool {
	if ptr.Pointer() != nil {
		return C.QQuaternion_IsIdentity(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuaternion) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QQuaternion_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuaternion) Normalize() {
	if ptr.Pointer() != nil {
		C.QQuaternion_Normalize(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QQuaternion) SetVector(vector QVector3DITF) {
	if ptr.Pointer() != nil {
		C.QQuaternion_SetVector(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQVector3D(vector)))
	}
}
