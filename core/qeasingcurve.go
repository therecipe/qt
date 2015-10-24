package core

//#include "qeasingcurve.h"
import "C"
import (
	"unsafe"
)

type QEasingCurve struct {
	ptr unsafe.Pointer
}

type QEasingCurveITF interface {
	QEasingCurvePTR() *QEasingCurve
}

func (p *QEasingCurve) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QEasingCurve) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQEasingCurve(ptr QEasingCurveITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QEasingCurvePTR().Pointer()
	}
	return nil
}

func QEasingCurveFromPointer(ptr unsafe.Pointer) *QEasingCurve {
	var n = new(QEasingCurve)
	n.SetPointer(ptr)
	return n
}

func (ptr *QEasingCurve) QEasingCurvePTR() *QEasingCurve {
	return ptr
}

//QEasingCurve::Type
type QEasingCurve__Type int

var (
	QEasingCurve__Linear       = QEasingCurve__Type(0)
	QEasingCurve__InQuad       = QEasingCurve__Type(1)
	QEasingCurve__OutQuad      = QEasingCurve__Type(2)
	QEasingCurve__InOutQuad    = QEasingCurve__Type(3)
	QEasingCurve__OutInQuad    = QEasingCurve__Type(4)
	QEasingCurve__InCubic      = QEasingCurve__Type(5)
	QEasingCurve__OutCubic     = QEasingCurve__Type(6)
	QEasingCurve__InOutCubic   = QEasingCurve__Type(7)
	QEasingCurve__OutInCubic   = QEasingCurve__Type(8)
	QEasingCurve__InQuart      = QEasingCurve__Type(9)
	QEasingCurve__OutQuart     = QEasingCurve__Type(10)
	QEasingCurve__InOutQuart   = QEasingCurve__Type(11)
	QEasingCurve__OutInQuart   = QEasingCurve__Type(12)
	QEasingCurve__InQuint      = QEasingCurve__Type(13)
	QEasingCurve__OutQuint     = QEasingCurve__Type(14)
	QEasingCurve__InOutQuint   = QEasingCurve__Type(15)
	QEasingCurve__OutInQuint   = QEasingCurve__Type(16)
	QEasingCurve__InSine       = QEasingCurve__Type(17)
	QEasingCurve__OutSine      = QEasingCurve__Type(18)
	QEasingCurve__InOutSine    = QEasingCurve__Type(19)
	QEasingCurve__OutInSine    = QEasingCurve__Type(20)
	QEasingCurve__InExpo       = QEasingCurve__Type(21)
	QEasingCurve__OutExpo      = QEasingCurve__Type(22)
	QEasingCurve__InOutExpo    = QEasingCurve__Type(23)
	QEasingCurve__OutInExpo    = QEasingCurve__Type(24)
	QEasingCurve__InCirc       = QEasingCurve__Type(25)
	QEasingCurve__OutCirc      = QEasingCurve__Type(26)
	QEasingCurve__InOutCirc    = QEasingCurve__Type(27)
	QEasingCurve__OutInCirc    = QEasingCurve__Type(28)
	QEasingCurve__InElastic    = QEasingCurve__Type(29)
	QEasingCurve__OutElastic   = QEasingCurve__Type(30)
	QEasingCurve__InOutElastic = QEasingCurve__Type(31)
	QEasingCurve__OutInElastic = QEasingCurve__Type(32)
	QEasingCurve__InBack       = QEasingCurve__Type(33)
	QEasingCurve__OutBack      = QEasingCurve__Type(34)
	QEasingCurve__InOutBack    = QEasingCurve__Type(35)
	QEasingCurve__OutInBack    = QEasingCurve__Type(36)
	QEasingCurve__InBounce     = QEasingCurve__Type(37)
	QEasingCurve__OutBounce    = QEasingCurve__Type(38)
	QEasingCurve__InOutBounce  = QEasingCurve__Type(39)
	QEasingCurve__OutInBounce  = QEasingCurve__Type(40)
	QEasingCurve__InCurve      = QEasingCurve__Type(41)
	QEasingCurve__OutCurve     = QEasingCurve__Type(42)
	QEasingCurve__SineCurve    = QEasingCurve__Type(43)
	QEasingCurve__CosineCurve  = QEasingCurve__Type(44)
	QEasingCurve__BezierSpline = QEasingCurve__Type(45)
	QEasingCurve__TCBSpline    = QEasingCurve__Type(46)
	QEasingCurve__Custom       = QEasingCurve__Type(47)
	QEasingCurve__NCurveTypes  = QEasingCurve__Type(48)
)

func NewQEasingCurve3(other QEasingCurveITF) *QEasingCurve {
	return QEasingCurveFromPointer(unsafe.Pointer(C.QEasingCurve_NewQEasingCurve3(C.QtObjectPtr(PointerFromQEasingCurve(other)))))
}

func NewQEasingCurve(ty QEasingCurve__Type) *QEasingCurve {
	return QEasingCurveFromPointer(unsafe.Pointer(C.QEasingCurve_NewQEasingCurve(C.int(ty))))
}

func NewQEasingCurve2(other QEasingCurveITF) *QEasingCurve {
	return QEasingCurveFromPointer(unsafe.Pointer(C.QEasingCurve_NewQEasingCurve2(C.QtObjectPtr(PointerFromQEasingCurve(other)))))
}

func (ptr *QEasingCurve) AddCubicBezierSegment(c1 QPointFITF, c2 QPointFITF, endPoint QPointFITF) {
	if ptr.Pointer() != nil {
		C.QEasingCurve_AddCubicBezierSegment(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPointF(c1)), C.QtObjectPtr(PointerFromQPointF(c2)), C.QtObjectPtr(PointerFromQPointF(endPoint)))
	}
}

func (ptr *QEasingCurve) SetType(ty QEasingCurve__Type) {
	if ptr.Pointer() != nil {
		C.QEasingCurve_SetType(C.QtObjectPtr(ptr.Pointer()), C.int(ty))
	}
}

func (ptr *QEasingCurve) Swap(other QEasingCurveITF) {
	if ptr.Pointer() != nil {
		C.QEasingCurve_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQEasingCurve(other)))
	}
}

func (ptr *QEasingCurve) Type() QEasingCurve__Type {
	if ptr.Pointer() != nil {
		return QEasingCurve__Type(C.QEasingCurve_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QEasingCurve) DestroyQEasingCurve() {
	if ptr.Pointer() != nil {
		C.QEasingCurve_DestroyQEasingCurve(C.QtObjectPtr(ptr.Pointer()))
	}
}
