package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QRadialGradient struct {
	QGradient
}

type QRadialGradient_ITF interface {
	QGradient_ITF
	QRadialGradient_PTR() *QRadialGradient
}

func PointerFromQRadialGradient(ptr QRadialGradient_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRadialGradient_PTR().Pointer()
	}
	return nil
}

func NewQRadialGradientFromPointer(ptr unsafe.Pointer) *QRadialGradient {
	var n = new(QRadialGradient)
	n.SetPointer(ptr)
	return n
}

func (ptr *QRadialGradient) QRadialGradient_PTR() *QRadialGradient {
	return ptr
}

func NewQRadialGradient() *QRadialGradient {
	defer qt.Recovering("QRadialGradient::QRadialGradient")

	return NewQRadialGradientFromPointer(C.QRadialGradient_NewQRadialGradient())
}

func NewQRadialGradient6(center core.QPointF_ITF, centerRadius float64, focalPoint core.QPointF_ITF, focalRadius float64) *QRadialGradient {
	defer qt.Recovering("QRadialGradient::QRadialGradient")

	return NewQRadialGradientFromPointer(C.QRadialGradient_NewQRadialGradient6(core.PointerFromQPointF(center), C.double(centerRadius), core.PointerFromQPointF(focalPoint), C.double(focalRadius)))
}

func NewQRadialGradient4(center core.QPointF_ITF, radius float64) *QRadialGradient {
	defer qt.Recovering("QRadialGradient::QRadialGradient")

	return NewQRadialGradientFromPointer(C.QRadialGradient_NewQRadialGradient4(core.PointerFromQPointF(center), C.double(radius)))
}

func NewQRadialGradient2(center core.QPointF_ITF, radius float64, focalPoint core.QPointF_ITF) *QRadialGradient {
	defer qt.Recovering("QRadialGradient::QRadialGradient")

	return NewQRadialGradientFromPointer(C.QRadialGradient_NewQRadialGradient2(core.PointerFromQPointF(center), C.double(radius), core.PointerFromQPointF(focalPoint)))
}

func NewQRadialGradient7(cx float64, cy float64, centerRadius float64, fx float64, fy float64, focalRadius float64) *QRadialGradient {
	defer qt.Recovering("QRadialGradient::QRadialGradient")

	return NewQRadialGradientFromPointer(C.QRadialGradient_NewQRadialGradient7(C.double(cx), C.double(cy), C.double(centerRadius), C.double(fx), C.double(fy), C.double(focalRadius)))
}

func NewQRadialGradient5(cx float64, cy float64, radius float64) *QRadialGradient {
	defer qt.Recovering("QRadialGradient::QRadialGradient")

	return NewQRadialGradientFromPointer(C.QRadialGradient_NewQRadialGradient5(C.double(cx), C.double(cy), C.double(radius)))
}

func NewQRadialGradient3(cx float64, cy float64, radius float64, fx float64, fy float64) *QRadialGradient {
	defer qt.Recovering("QRadialGradient::QRadialGradient")

	return NewQRadialGradientFromPointer(C.QRadialGradient_NewQRadialGradient3(C.double(cx), C.double(cy), C.double(radius), C.double(fx), C.double(fy)))
}

func (ptr *QRadialGradient) CenterRadius() float64 {
	defer qt.Recovering("QRadialGradient::centerRadius")

	if ptr.Pointer() != nil {
		return float64(C.QRadialGradient_CenterRadius(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadialGradient) FocalRadius() float64 {
	defer qt.Recovering("QRadialGradient::focalRadius")

	if ptr.Pointer() != nil {
		return float64(C.QRadialGradient_FocalRadius(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadialGradient) Radius() float64 {
	defer qt.Recovering("QRadialGradient::radius")

	if ptr.Pointer() != nil {
		return float64(C.QRadialGradient_Radius(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadialGradient) SetCenter(center core.QPointF_ITF) {
	defer qt.Recovering("QRadialGradient::setCenter")

	if ptr.Pointer() != nil {
		C.QRadialGradient_SetCenter(ptr.Pointer(), core.PointerFromQPointF(center))
	}
}

func (ptr *QRadialGradient) SetCenter2(x float64, y float64) {
	defer qt.Recovering("QRadialGradient::setCenter")

	if ptr.Pointer() != nil {
		C.QRadialGradient_SetCenter2(ptr.Pointer(), C.double(x), C.double(y))
	}
}

func (ptr *QRadialGradient) SetCenterRadius(radius float64) {
	defer qt.Recovering("QRadialGradient::setCenterRadius")

	if ptr.Pointer() != nil {
		C.QRadialGradient_SetCenterRadius(ptr.Pointer(), C.double(radius))
	}
}

func (ptr *QRadialGradient) SetFocalPoint(focalPoint core.QPointF_ITF) {
	defer qt.Recovering("QRadialGradient::setFocalPoint")

	if ptr.Pointer() != nil {
		C.QRadialGradient_SetFocalPoint(ptr.Pointer(), core.PointerFromQPointF(focalPoint))
	}
}

func (ptr *QRadialGradient) SetFocalPoint2(x float64, y float64) {
	defer qt.Recovering("QRadialGradient::setFocalPoint")

	if ptr.Pointer() != nil {
		C.QRadialGradient_SetFocalPoint2(ptr.Pointer(), C.double(x), C.double(y))
	}
}

func (ptr *QRadialGradient) SetFocalRadius(radius float64) {
	defer qt.Recovering("QRadialGradient::setFocalRadius")

	if ptr.Pointer() != nil {
		C.QRadialGradient_SetFocalRadius(ptr.Pointer(), C.double(radius))
	}
}

func (ptr *QRadialGradient) SetRadius(radius float64) {
	defer qt.Recovering("QRadialGradient::setRadius")

	if ptr.Pointer() != nil {
		C.QRadialGradient_SetRadius(ptr.Pointer(), C.double(radius))
	}
}
