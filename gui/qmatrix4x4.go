package gui

//#include "qmatrix4x4.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMatrix4x4 struct {
	ptr unsafe.Pointer
}

type QMatrix4x4ITF interface {
	QMatrix4x4PTR() *QMatrix4x4
}

func (p *QMatrix4x4) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMatrix4x4) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMatrix4x4(ptr QMatrix4x4ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMatrix4x4PTR().Pointer()
	}
	return nil
}

func QMatrix4x4FromPointer(ptr unsafe.Pointer) *QMatrix4x4 {
	var n = new(QMatrix4x4)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMatrix4x4) QMatrix4x4PTR() *QMatrix4x4 {
	return ptr
}

func NewQMatrix4x4() *QMatrix4x4 {
	return QMatrix4x4FromPointer(unsafe.Pointer(C.QMatrix4x4_NewQMatrix4x4()))
}

func NewQMatrix4x47(transform QTransformITF) *QMatrix4x4 {
	return QMatrix4x4FromPointer(unsafe.Pointer(C.QMatrix4x4_NewQMatrix4x47(C.QtObjectPtr(PointerFromQTransform(transform)))))
}

func (ptr *QMatrix4x4) IsAffine() bool {
	if ptr.Pointer() != nil {
		return C.QMatrix4x4_IsAffine(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMatrix4x4) IsIdentity() bool {
	if ptr.Pointer() != nil {
		return C.QMatrix4x4_IsIdentity(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMatrix4x4) LookAt(eye QVector3DITF, center QVector3DITF, up QVector3DITF) {
	if ptr.Pointer() != nil {
		C.QMatrix4x4_LookAt(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQVector3D(eye)), C.QtObjectPtr(PointerFromQVector3D(center)), C.QtObjectPtr(PointerFromQVector3D(up)))
	}
}

func (ptr *QMatrix4x4) Optimize() {
	if ptr.Pointer() != nil {
		C.QMatrix4x4_Optimize(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMatrix4x4) Ortho2(rect core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QMatrix4x4_Ortho2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rect)))
	}
}

func (ptr *QMatrix4x4) Ortho3(rect core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QMatrix4x4_Ortho3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rect)))
	}
}

func (ptr *QMatrix4x4) Rotate2(quaternion QQuaternionITF) {
	if ptr.Pointer() != nil {
		C.QMatrix4x4_Rotate2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQQuaternion(quaternion)))
	}
}

func (ptr *QMatrix4x4) Scale(vector QVector3DITF) {
	if ptr.Pointer() != nil {
		C.QMatrix4x4_Scale(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQVector3D(vector)))
	}
}

func (ptr *QMatrix4x4) SetColumn(index int, value QVector4DITF) {
	if ptr.Pointer() != nil {
		C.QMatrix4x4_SetColumn(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.QtObjectPtr(PointerFromQVector4D(value)))
	}
}

func (ptr *QMatrix4x4) SetRow(index int, value QVector4DITF) {
	if ptr.Pointer() != nil {
		C.QMatrix4x4_SetRow(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.QtObjectPtr(PointerFromQVector4D(value)))
	}
}

func (ptr *QMatrix4x4) SetToIdentity() {
	if ptr.Pointer() != nil {
		C.QMatrix4x4_SetToIdentity(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMatrix4x4) Translate(vector QVector3DITF) {
	if ptr.Pointer() != nil {
		C.QMatrix4x4_Translate(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQVector3D(vector)))
	}
}

func (ptr *QMatrix4x4) Viewport2(rect core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QMatrix4x4_Viewport2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rect)))
	}
}
