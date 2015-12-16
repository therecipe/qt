package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMatrix4x4 struct {
	ptr unsafe.Pointer
}

type QMatrix4x4_ITF interface {
	QMatrix4x4_PTR() *QMatrix4x4
}

func (p *QMatrix4x4) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMatrix4x4) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMatrix4x4(ptr QMatrix4x4_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMatrix4x4_PTR().Pointer()
	}
	return nil
}

func NewQMatrix4x4FromPointer(ptr unsafe.Pointer) *QMatrix4x4 {
	var n = new(QMatrix4x4)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMatrix4x4) QMatrix4x4_PTR() *QMatrix4x4 {
	return ptr
}

func NewQMatrix4x4() *QMatrix4x4 {
	defer qt.Recovering("QMatrix4x4::QMatrix4x4")

	return NewQMatrix4x4FromPointer(C.QMatrix4x4_NewQMatrix4x4())
}

func NewQMatrix4x47(transform QTransform_ITF) *QMatrix4x4 {
	defer qt.Recovering("QMatrix4x4::QMatrix4x4")

	return NewQMatrix4x4FromPointer(C.QMatrix4x4_NewQMatrix4x47(PointerFromQTransform(transform)))
}

func (ptr *QMatrix4x4) IsAffine() bool {
	defer qt.Recovering("QMatrix4x4::isAffine")

	if ptr.Pointer() != nil {
		return C.QMatrix4x4_IsAffine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMatrix4x4) IsIdentity() bool {
	defer qt.Recovering("QMatrix4x4::isIdentity")

	if ptr.Pointer() != nil {
		return C.QMatrix4x4_IsIdentity(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMatrix4x4) LookAt(eye QVector3D_ITF, center QVector3D_ITF, up QVector3D_ITF) {
	defer qt.Recovering("QMatrix4x4::lookAt")

	if ptr.Pointer() != nil {
		C.QMatrix4x4_LookAt(ptr.Pointer(), PointerFromQVector3D(eye), PointerFromQVector3D(center), PointerFromQVector3D(up))
	}
}

func (ptr *QMatrix4x4) Map(point core.QPoint_ITF) *core.QPoint {
	defer qt.Recovering("QMatrix4x4::map")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QMatrix4x4_Map(ptr.Pointer(), core.PointerFromQPoint(point)))
	}
	return nil
}

func (ptr *QMatrix4x4) MapRect(rect core.QRect_ITF) *core.QRect {
	defer qt.Recovering("QMatrix4x4::mapRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QMatrix4x4_MapRect(ptr.Pointer(), core.PointerFromQRect(rect)))
	}
	return nil
}

func (ptr *QMatrix4x4) Optimize() {
	defer qt.Recovering("QMatrix4x4::optimize")

	if ptr.Pointer() != nil {
		C.QMatrix4x4_Optimize(ptr.Pointer())
	}
}

func (ptr *QMatrix4x4) Ortho2(rect core.QRect_ITF) {
	defer qt.Recovering("QMatrix4x4::ortho")

	if ptr.Pointer() != nil {
		C.QMatrix4x4_Ortho2(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QMatrix4x4) Ortho3(rect core.QRectF_ITF) {
	defer qt.Recovering("QMatrix4x4::ortho")

	if ptr.Pointer() != nil {
		C.QMatrix4x4_Ortho3(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QMatrix4x4) Rotate2(quaternion QQuaternion_ITF) {
	defer qt.Recovering("QMatrix4x4::rotate")

	if ptr.Pointer() != nil {
		C.QMatrix4x4_Rotate2(ptr.Pointer(), PointerFromQQuaternion(quaternion))
	}
}

func (ptr *QMatrix4x4) Scale(vector QVector3D_ITF) {
	defer qt.Recovering("QMatrix4x4::scale")

	if ptr.Pointer() != nil {
		C.QMatrix4x4_Scale(ptr.Pointer(), PointerFromQVector3D(vector))
	}
}

func (ptr *QMatrix4x4) SetColumn(index int, value QVector4D_ITF) {
	defer qt.Recovering("QMatrix4x4::setColumn")

	if ptr.Pointer() != nil {
		C.QMatrix4x4_SetColumn(ptr.Pointer(), C.int(index), PointerFromQVector4D(value))
	}
}

func (ptr *QMatrix4x4) SetRow(index int, value QVector4D_ITF) {
	defer qt.Recovering("QMatrix4x4::setRow")

	if ptr.Pointer() != nil {
		C.QMatrix4x4_SetRow(ptr.Pointer(), C.int(index), PointerFromQVector4D(value))
	}
}

func (ptr *QMatrix4x4) SetToIdentity() {
	defer qt.Recovering("QMatrix4x4::setToIdentity")

	if ptr.Pointer() != nil {
		C.QMatrix4x4_SetToIdentity(ptr.Pointer())
	}
}

func (ptr *QMatrix4x4) Translate(vector QVector3D_ITF) {
	defer qt.Recovering("QMatrix4x4::translate")

	if ptr.Pointer() != nil {
		C.QMatrix4x4_Translate(ptr.Pointer(), PointerFromQVector3D(vector))
	}
}

func (ptr *QMatrix4x4) Viewport2(rect core.QRectF_ITF) {
	defer qt.Recovering("QMatrix4x4::viewport")

	if ptr.Pointer() != nil {
		C.QMatrix4x4_Viewport2(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}
