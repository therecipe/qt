package gui

//#include "qgradient.h"
import "C"
import (
	"unsafe"
)

type QGradient struct {
	ptr unsafe.Pointer
}

type QGradientITF interface {
	QGradientPTR() *QGradient
}

func (p *QGradient) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGradient) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGradient(ptr QGradientITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGradientPTR().Pointer()
	}
	return nil
}

func QGradientFromPointer(ptr unsafe.Pointer) *QGradient {
	var n = new(QGradient)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGradient) QGradientPTR() *QGradient {
	return ptr
}

//QGradient::CoordinateMode
type QGradient__CoordinateMode int

var (
	QGradient__LogicalMode         = QGradient__CoordinateMode(0)
	QGradient__StretchToDeviceMode = QGradient__CoordinateMode(1)
	QGradient__ObjectBoundingMode  = QGradient__CoordinateMode(2)
)

//QGradient::Spread
type QGradient__Spread int

var (
	QGradient__PadSpread     = QGradient__Spread(0)
	QGradient__ReflectSpread = QGradient__Spread(1)
	QGradient__RepeatSpread  = QGradient__Spread(2)
)

//QGradient::Type
type QGradient__Type int

var (
	QGradient__LinearGradient  = QGradient__Type(0)
	QGradient__RadialGradient  = QGradient__Type(1)
	QGradient__ConicalGradient = QGradient__Type(2)
	QGradient__NoGradient      = QGradient__Type(3)
)

func (ptr *QGradient) CoordinateMode() QGradient__CoordinateMode {
	if ptr.Pointer() != nil {
		return QGradient__CoordinateMode(C.QGradient_CoordinateMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGradient) SetCoordinateMode(mode QGradient__CoordinateMode) {
	if ptr.Pointer() != nil {
		C.QGradient_SetCoordinateMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QGradient) SetSpread(method QGradient__Spread) {
	if ptr.Pointer() != nil {
		C.QGradient_SetSpread(C.QtObjectPtr(ptr.Pointer()), C.int(method))
	}
}

func (ptr *QGradient) Spread() QGradient__Spread {
	if ptr.Pointer() != nil {
		return QGradient__Spread(C.QGradient_Spread(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGradient) Type() QGradient__Type {
	if ptr.Pointer() != nil {
		return QGradient__Type(C.QGradient_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
