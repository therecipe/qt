package core

//#include "core.h"
import "C"
import (
	"log"
	"unsafe"
)

type QEasingCurve struct {
	ptr unsafe.Pointer
}

type QEasingCurve_ITF interface {
	QEasingCurve_PTR() *QEasingCurve
}

func (p *QEasingCurve) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QEasingCurve) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQEasingCurve(ptr QEasingCurve_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QEasingCurve_PTR().Pointer()
	}
	return nil
}

func NewQEasingCurveFromPointer(ptr unsafe.Pointer) *QEasingCurve {
	var n = new(QEasingCurve)
	n.SetPointer(ptr)
	return n
}

func (ptr *QEasingCurve) QEasingCurve_PTR() *QEasingCurve {
	return ptr
}

//QEasingCurve::Type
type QEasingCurve__Type int64

const (
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

func NewQEasingCurve3(other QEasingCurve_ITF) *QEasingCurve {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QEasingCurve::QEasingCurve")
		}
	}()

	return NewQEasingCurveFromPointer(C.QEasingCurve_NewQEasingCurve3(PointerFromQEasingCurve(other)))
}

func NewQEasingCurve(ty QEasingCurve__Type) *QEasingCurve {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QEasingCurve::QEasingCurve")
		}
	}()

	return NewQEasingCurveFromPointer(C.QEasingCurve_NewQEasingCurve(C.int(ty)))
}

func NewQEasingCurve2(other QEasingCurve_ITF) *QEasingCurve {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QEasingCurve::QEasingCurve")
		}
	}()

	return NewQEasingCurveFromPointer(C.QEasingCurve_NewQEasingCurve2(PointerFromQEasingCurve(other)))
}

func (ptr *QEasingCurve) AddCubicBezierSegment(c1 QPointF_ITF, c2 QPointF_ITF, endPoint QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QEasingCurve::addCubicBezierSegment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QEasingCurve_AddCubicBezierSegment(ptr.Pointer(), PointerFromQPointF(c1), PointerFromQPointF(c2), PointerFromQPointF(endPoint))
	}
}

func (ptr *QEasingCurve) AddTCBSegment(nextPoint QPointF_ITF, t float64, c float64, b float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QEasingCurve::addTCBSegment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QEasingCurve_AddTCBSegment(ptr.Pointer(), PointerFromQPointF(nextPoint), C.double(t), C.double(c), C.double(b))
	}
}

func (ptr *QEasingCurve) Amplitude() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QEasingCurve::amplitude")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QEasingCurve_Amplitude(ptr.Pointer()))
	}
	return 0
}

func (ptr *QEasingCurve) Overshoot() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QEasingCurve::overshoot")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QEasingCurve_Overshoot(ptr.Pointer()))
	}
	return 0
}

func (ptr *QEasingCurve) Period() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QEasingCurve::period")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QEasingCurve_Period(ptr.Pointer()))
	}
	return 0
}

func (ptr *QEasingCurve) SetAmplitude(amplitude float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QEasingCurve::setAmplitude")
		}
	}()

	if ptr.Pointer() != nil {
		C.QEasingCurve_SetAmplitude(ptr.Pointer(), C.double(amplitude))
	}
}

func (ptr *QEasingCurve) SetOvershoot(overshoot float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QEasingCurve::setOvershoot")
		}
	}()

	if ptr.Pointer() != nil {
		C.QEasingCurve_SetOvershoot(ptr.Pointer(), C.double(overshoot))
	}
}

func (ptr *QEasingCurve) SetPeriod(period float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QEasingCurve::setPeriod")
		}
	}()

	if ptr.Pointer() != nil {
		C.QEasingCurve_SetPeriod(ptr.Pointer(), C.double(period))
	}
}

func (ptr *QEasingCurve) SetType(ty QEasingCurve__Type) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QEasingCurve::setType")
		}
	}()

	if ptr.Pointer() != nil {
		C.QEasingCurve_SetType(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QEasingCurve) Swap(other QEasingCurve_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QEasingCurve::swap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QEasingCurve_Swap(ptr.Pointer(), PointerFromQEasingCurve(other))
	}
}

func (ptr *QEasingCurve) Type() QEasingCurve__Type {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QEasingCurve::type")
		}
	}()

	if ptr.Pointer() != nil {
		return QEasingCurve__Type(C.QEasingCurve_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QEasingCurve) ValueForProgress(progress float64) float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QEasingCurve::valueForProgress")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QEasingCurve_ValueForProgress(ptr.Pointer(), C.double(progress)))
	}
	return 0
}

func (ptr *QEasingCurve) DestroyQEasingCurve() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QEasingCurve::~QEasingCurve")
		}
	}()

	if ptr.Pointer() != nil {
		C.QEasingCurve_DestroyQEasingCurve(ptr.Pointer())
	}
}
