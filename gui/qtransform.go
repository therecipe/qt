package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTransform struct {
	ptr unsafe.Pointer
}

type QTransform_ITF interface {
	QTransform_PTR() *QTransform
}

func (p *QTransform) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTransform) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTransform(ptr QTransform_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTransform_PTR().Pointer()
	}
	return nil
}

func NewQTransformFromPointer(ptr unsafe.Pointer) *QTransform {
	var n = new(QTransform)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTransform) QTransform_PTR() *QTransform {
	return ptr
}

//QTransform::TransformationType
type QTransform__TransformationType int64

const (
	QTransform__TxNone      = QTransform__TransformationType(0x00)
	QTransform__TxTranslate = QTransform__TransformationType(0x01)
	QTransform__TxScale     = QTransform__TransformationType(0x02)
	QTransform__TxRotate    = QTransform__TransformationType(0x04)
	QTransform__TxShear     = QTransform__TransformationType(0x08)
	QTransform__TxProject   = QTransform__TransformationType(0x10)
)

func NewQTransform3(m11 float64, m12 float64, m13 float64, m21 float64, m22 float64, m23 float64, m31 float64, m32 float64, m33 float64) *QTransform {
	defer qt.Recovering("QTransform::QTransform")

	return NewQTransformFromPointer(C.QTransform_NewQTransform3(C.double(m11), C.double(m12), C.double(m13), C.double(m21), C.double(m22), C.double(m23), C.double(m31), C.double(m32), C.double(m33)))
}

func NewQTransform4(m11 float64, m12 float64, m21 float64, m22 float64, dx float64, dy float64) *QTransform {
	defer qt.Recovering("QTransform::QTransform")

	return NewQTransformFromPointer(C.QTransform_NewQTransform4(C.double(m11), C.double(m12), C.double(m21), C.double(m22), C.double(dx), C.double(dy)))
}

func (ptr *QTransform) Map3(point core.QPoint_ITF) *core.QPoint {
	defer qt.Recovering("QTransform::map")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QTransform_Map3(ptr.Pointer(), core.PointerFromQPoint(point)))
	}
	return nil
}

func (ptr *QTransform) Map8(region QRegion_ITF) *QRegion {
	defer qt.Recovering("QTransform::map")

	if ptr.Pointer() != nil {
		return NewQRegionFromPointer(C.QTransform_Map8(ptr.Pointer(), PointerFromQRegion(region)))
	}
	return nil
}

func (ptr *QTransform) MapRect2(rectangle core.QRect_ITF) *core.QRect {
	defer qt.Recovering("QTransform::mapRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QTransform_MapRect2(ptr.Pointer(), core.PointerFromQRect(rectangle)))
	}
	return nil
}

func QTransform_QuadToSquare(quad QPolygonF_ITF, trans QTransform_ITF) bool {
	defer qt.Recovering("QTransform::quadToSquare")

	return C.QTransform_QTransform_QuadToSquare(PointerFromQPolygonF(quad), PointerFromQTransform(trans)) != 0
}

func NewQTransform() *QTransform {
	defer qt.Recovering("QTransform::QTransform")

	return NewQTransformFromPointer(C.QTransform_NewQTransform())
}

func (ptr *QTransform) Determinant() float64 {
	defer qt.Recovering("QTransform::determinant")

	if ptr.Pointer() != nil {
		return float64(C.QTransform_Determinant(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTransform) Dx() float64 {
	defer qt.Recovering("QTransform::dx")

	if ptr.Pointer() != nil {
		return float64(C.QTransform_Dx(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTransform) Dy() float64 {
	defer qt.Recovering("QTransform::dy")

	if ptr.Pointer() != nil {
		return float64(C.QTransform_Dy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTransform) IsAffine() bool {
	defer qt.Recovering("QTransform::isAffine")

	if ptr.Pointer() != nil {
		return C.QTransform_IsAffine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTransform) IsIdentity() bool {
	defer qt.Recovering("QTransform::isIdentity")

	if ptr.Pointer() != nil {
		return C.QTransform_IsIdentity(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTransform) IsInvertible() bool {
	defer qt.Recovering("QTransform::isInvertible")

	if ptr.Pointer() != nil {
		return C.QTransform_IsInvertible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTransform) IsRotating() bool {
	defer qt.Recovering("QTransform::isRotating")

	if ptr.Pointer() != nil {
		return C.QTransform_IsRotating(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTransform) IsScaling() bool {
	defer qt.Recovering("QTransform::isScaling")

	if ptr.Pointer() != nil {
		return C.QTransform_IsScaling(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTransform) IsTranslating() bool {
	defer qt.Recovering("QTransform::isTranslating")

	if ptr.Pointer() != nil {
		return C.QTransform_IsTranslating(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTransform) M11() float64 {
	defer qt.Recovering("QTransform::m11")

	if ptr.Pointer() != nil {
		return float64(C.QTransform_M11(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTransform) M12() float64 {
	defer qt.Recovering("QTransform::m12")

	if ptr.Pointer() != nil {
		return float64(C.QTransform_M12(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTransform) M13() float64 {
	defer qt.Recovering("QTransform::m13")

	if ptr.Pointer() != nil {
		return float64(C.QTransform_M13(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTransform) M21() float64 {
	defer qt.Recovering("QTransform::m21")

	if ptr.Pointer() != nil {
		return float64(C.QTransform_M21(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTransform) M22() float64 {
	defer qt.Recovering("QTransform::m22")

	if ptr.Pointer() != nil {
		return float64(C.QTransform_M22(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTransform) M23() float64 {
	defer qt.Recovering("QTransform::m23")

	if ptr.Pointer() != nil {
		return float64(C.QTransform_M23(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTransform) M31() float64 {
	defer qt.Recovering("QTransform::m31")

	if ptr.Pointer() != nil {
		return float64(C.QTransform_M31(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTransform) M32() float64 {
	defer qt.Recovering("QTransform::m32")

	if ptr.Pointer() != nil {
		return float64(C.QTransform_M32(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTransform) M33() float64 {
	defer qt.Recovering("QTransform::m33")

	if ptr.Pointer() != nil {
		return float64(C.QTransform_M33(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTransform) Map10(x int, y int, tx int, ty int) {
	defer qt.Recovering("QTransform::map")

	if ptr.Pointer() != nil {
		C.QTransform_Map10(ptr.Pointer(), C.int(x), C.int(y), C.int(tx), C.int(ty))
	}
}

func QTransform_QuadToQuad(one QPolygonF_ITF, two QPolygonF_ITF, trans QTransform_ITF) bool {
	defer qt.Recovering("QTransform::quadToQuad")

	return C.QTransform_QTransform_QuadToQuad(PointerFromQPolygonF(one), PointerFromQPolygonF(two), PointerFromQTransform(trans)) != 0
}

func (ptr *QTransform) Reset() {
	defer qt.Recovering("QTransform::reset")

	if ptr.Pointer() != nil {
		C.QTransform_Reset(ptr.Pointer())
	}
}

func (ptr *QTransform) SetMatrix(m11 float64, m12 float64, m13 float64, m21 float64, m22 float64, m23 float64, m31 float64, m32 float64, m33 float64) {
	defer qt.Recovering("QTransform::setMatrix")

	if ptr.Pointer() != nil {
		C.QTransform_SetMatrix(ptr.Pointer(), C.double(m11), C.double(m12), C.double(m13), C.double(m21), C.double(m22), C.double(m23), C.double(m31), C.double(m32), C.double(m33))
	}
}

func QTransform_SquareToQuad(quad QPolygonF_ITF, trans QTransform_ITF) bool {
	defer qt.Recovering("QTransform::squareToQuad")

	return C.QTransform_QTransform_SquareToQuad(PointerFromQPolygonF(quad), PointerFromQTransform(trans)) != 0
}

func (ptr *QTransform) Type() QTransform__TransformationType {
	defer qt.Recovering("QTransform::type")

	if ptr.Pointer() != nil {
		return QTransform__TransformationType(C.QTransform_Type(ptr.Pointer()))
	}
	return 0
}
