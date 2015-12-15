package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QGradient struct {
	ptr unsafe.Pointer
}

type QGradient_ITF interface {
	QGradient_PTR() *QGradient
}

func (p *QGradient) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGradient) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGradient(ptr QGradient_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGradient_PTR().Pointer()
	}
	return nil
}

func NewQGradientFromPointer(ptr unsafe.Pointer) *QGradient {
	var n = new(QGradient)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGradient) QGradient_PTR() *QGradient {
	return ptr
}

//QGradient::CoordinateMode
type QGradient__CoordinateMode int64

const (
	QGradient__LogicalMode         = QGradient__CoordinateMode(0)
	QGradient__StretchToDeviceMode = QGradient__CoordinateMode(1)
	QGradient__ObjectBoundingMode  = QGradient__CoordinateMode(2)
)

//QGradient::Spread
type QGradient__Spread int64

const (
	QGradient__PadSpread     = QGradient__Spread(0)
	QGradient__ReflectSpread = QGradient__Spread(1)
	QGradient__RepeatSpread  = QGradient__Spread(2)
)

//QGradient::Type
type QGradient__Type int64

const (
	QGradient__LinearGradient  = QGradient__Type(0)
	QGradient__RadialGradient  = QGradient__Type(1)
	QGradient__ConicalGradient = QGradient__Type(2)
	QGradient__NoGradient      = QGradient__Type(3)
)

func (ptr *QGradient) SetColorAt(position float64, color QColor_ITF) {
	defer qt.Recovering("QGradient::setColorAt")

	if ptr.Pointer() != nil {
		C.QGradient_SetColorAt(ptr.Pointer(), C.double(position), PointerFromQColor(color))
	}
}

func (ptr *QGradient) CoordinateMode() QGradient__CoordinateMode {
	defer qt.Recovering("QGradient::coordinateMode")

	if ptr.Pointer() != nil {
		return QGradient__CoordinateMode(C.QGradient_CoordinateMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGradient) SetCoordinateMode(mode QGradient__CoordinateMode) {
	defer qt.Recovering("QGradient::setCoordinateMode")

	if ptr.Pointer() != nil {
		C.QGradient_SetCoordinateMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QGradient) SetSpread(method QGradient__Spread) {
	defer qt.Recovering("QGradient::setSpread")

	if ptr.Pointer() != nil {
		C.QGradient_SetSpread(ptr.Pointer(), C.int(method))
	}
}

func (ptr *QGradient) Spread() QGradient__Spread {
	defer qt.Recovering("QGradient::spread")

	if ptr.Pointer() != nil {
		return QGradient__Spread(C.QGradient_Spread(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGradient) Type() QGradient__Type {
	defer qt.Recovering("QGradient::type")

	if ptr.Pointer() != nil {
		return QGradient__Type(C.QGradient_Type(ptr.Pointer()))
	}
	return 0
}
