package gui

//#include "qtransform.h"
import "C"
import (
	"unsafe"
)

type QTransform struct {
	ptr unsafe.Pointer
}

type QTransformITF interface {
	QTransformPTR() *QTransform
}

func (p *QTransform) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTransform) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTransform(ptr QTransformITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTransformPTR().Pointer()
	}
	return nil
}

func QTransformFromPointer(ptr unsafe.Pointer) *QTransform {
	var n = new(QTransform)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTransform) QTransformPTR() *QTransform {
	return ptr
}

//QTransform::TransformationType
type QTransform__TransformationType int

var (
	QTransform__TxNone      = QTransform__TransformationType(0x00)
	QTransform__TxTranslate = QTransform__TransformationType(0x01)
	QTransform__TxScale     = QTransform__TransformationType(0x02)
	QTransform__TxRotate    = QTransform__TransformationType(0x04)
	QTransform__TxShear     = QTransform__TransformationType(0x08)
	QTransform__TxProject   = QTransform__TransformationType(0x10)
)

func QTransform_QuadToSquare(quad QPolygonFITF, trans QTransformITF) bool {
	return C.QTransform_QTransform_QuadToSquare(C.QtObjectPtr(PointerFromQPolygonF(quad)), C.QtObjectPtr(PointerFromQTransform(trans))) != 0
}

func NewQTransform() *QTransform {
	return QTransformFromPointer(unsafe.Pointer(C.QTransform_NewQTransform()))
}

func (ptr *QTransform) IsAffine() bool {
	if ptr.Pointer() != nil {
		return C.QTransform_IsAffine(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTransform) IsIdentity() bool {
	if ptr.Pointer() != nil {
		return C.QTransform_IsIdentity(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTransform) IsInvertible() bool {
	if ptr.Pointer() != nil {
		return C.QTransform_IsInvertible(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTransform) IsRotating() bool {
	if ptr.Pointer() != nil {
		return C.QTransform_IsRotating(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTransform) IsScaling() bool {
	if ptr.Pointer() != nil {
		return C.QTransform_IsScaling(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTransform) IsTranslating() bool {
	if ptr.Pointer() != nil {
		return C.QTransform_IsTranslating(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTransform) Map10(x int, y int, tx int, ty int) {
	if ptr.Pointer() != nil {
		C.QTransform_Map10(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.int(tx), C.int(ty))
	}
}

func QTransform_QuadToQuad(one QPolygonFITF, two QPolygonFITF, trans QTransformITF) bool {
	return C.QTransform_QTransform_QuadToQuad(C.QtObjectPtr(PointerFromQPolygonF(one)), C.QtObjectPtr(PointerFromQPolygonF(two)), C.QtObjectPtr(PointerFromQTransform(trans))) != 0
}

func (ptr *QTransform) Reset() {
	if ptr.Pointer() != nil {
		C.QTransform_Reset(C.QtObjectPtr(ptr.Pointer()))
	}
}

func QTransform_SquareToQuad(quad QPolygonFITF, trans QTransformITF) bool {
	return C.QTransform_QTransform_SquareToQuad(C.QtObjectPtr(PointerFromQPolygonF(quad)), C.QtObjectPtr(PointerFromQTransform(trans))) != 0
}

func (ptr *QTransform) Type() QTransform__TransformationType {
	if ptr.Pointer() != nil {
		return QTransform__TransformationType(C.QTransform_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
