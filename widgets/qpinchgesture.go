package widgets

//#include "qpinchgesture.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPinchGesture struct {
	QGesture
}

type QPinchGesture_ITF interface {
	QGesture_ITF
	QPinchGesture_PTR() *QPinchGesture
}

func PointerFromQPinchGesture(ptr QPinchGesture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPinchGesture_PTR().Pointer()
	}
	return nil
}

func NewQPinchGestureFromPointer(ptr unsafe.Pointer) *QPinchGesture {
	var n = new(QPinchGesture)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QPinchGesture_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPinchGesture) QPinchGesture_PTR() *QPinchGesture {
	return ptr
}

//QPinchGesture::ChangeFlag
type QPinchGesture__ChangeFlag int64

const (
	QPinchGesture__ScaleFactorChanged   = QPinchGesture__ChangeFlag(0x1)
	QPinchGesture__RotationAngleChanged = QPinchGesture__ChangeFlag(0x2)
	QPinchGesture__CenterPointChanged   = QPinchGesture__ChangeFlag(0x4)
)

func (ptr *QPinchGesture) ChangeFlags() QPinchGesture__ChangeFlag {
	if ptr.Pointer() != nil {
		return QPinchGesture__ChangeFlag(C.QPinchGesture_ChangeFlags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPinchGesture) LastRotationAngle() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPinchGesture_LastRotationAngle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPinchGesture) LastScaleFactor() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPinchGesture_LastScaleFactor(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPinchGesture) RotationAngle() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPinchGesture_RotationAngle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPinchGesture) ScaleFactor() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPinchGesture_ScaleFactor(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPinchGesture) SetCenterPoint(value core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QPinchGesture_SetCenterPoint(ptr.Pointer(), core.PointerFromQPointF(value))
	}
}

func (ptr *QPinchGesture) SetChangeFlags(value QPinchGesture__ChangeFlag) {
	if ptr.Pointer() != nil {
		C.QPinchGesture_SetChangeFlags(ptr.Pointer(), C.int(value))
	}
}

func (ptr *QPinchGesture) SetLastCenterPoint(value core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QPinchGesture_SetLastCenterPoint(ptr.Pointer(), core.PointerFromQPointF(value))
	}
}

func (ptr *QPinchGesture) SetLastRotationAngle(value float64) {
	if ptr.Pointer() != nil {
		C.QPinchGesture_SetLastRotationAngle(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QPinchGesture) SetLastScaleFactor(value float64) {
	if ptr.Pointer() != nil {
		C.QPinchGesture_SetLastScaleFactor(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QPinchGesture) SetRotationAngle(value float64) {
	if ptr.Pointer() != nil {
		C.QPinchGesture_SetRotationAngle(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QPinchGesture) SetScaleFactor(value float64) {
	if ptr.Pointer() != nil {
		C.QPinchGesture_SetScaleFactor(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QPinchGesture) SetStartCenterPoint(value core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QPinchGesture_SetStartCenterPoint(ptr.Pointer(), core.PointerFromQPointF(value))
	}
}

func (ptr *QPinchGesture) SetTotalChangeFlags(value QPinchGesture__ChangeFlag) {
	if ptr.Pointer() != nil {
		C.QPinchGesture_SetTotalChangeFlags(ptr.Pointer(), C.int(value))
	}
}

func (ptr *QPinchGesture) SetTotalRotationAngle(value float64) {
	if ptr.Pointer() != nil {
		C.QPinchGesture_SetTotalRotationAngle(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QPinchGesture) SetTotalScaleFactor(value float64) {
	if ptr.Pointer() != nil {
		C.QPinchGesture_SetTotalScaleFactor(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QPinchGesture) TotalChangeFlags() QPinchGesture__ChangeFlag {
	if ptr.Pointer() != nil {
		return QPinchGesture__ChangeFlag(C.QPinchGesture_TotalChangeFlags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPinchGesture) TotalRotationAngle() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPinchGesture_TotalRotationAngle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPinchGesture) TotalScaleFactor() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPinchGesture_TotalScaleFactor(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPinchGesture) DestroyQPinchGesture() {
	if ptr.Pointer() != nil {
		C.QPinchGesture_DestroyQPinchGesture(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
