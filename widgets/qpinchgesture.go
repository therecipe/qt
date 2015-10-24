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

type QPinchGestureITF interface {
	QGestureITF
	QPinchGesturePTR() *QPinchGesture
}

func PointerFromQPinchGesture(ptr QPinchGestureITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPinchGesturePTR().Pointer()
	}
	return nil
}

func QPinchGestureFromPointer(ptr unsafe.Pointer) *QPinchGesture {
	var n = new(QPinchGesture)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QPinchGesture_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPinchGesture) QPinchGesturePTR() *QPinchGesture {
	return ptr
}

//QPinchGesture::ChangeFlag
type QPinchGesture__ChangeFlag int

var (
	QPinchGesture__ScaleFactorChanged   = QPinchGesture__ChangeFlag(0x1)
	QPinchGesture__RotationAngleChanged = QPinchGesture__ChangeFlag(0x2)
	QPinchGesture__CenterPointChanged   = QPinchGesture__ChangeFlag(0x4)
)

func (ptr *QPinchGesture) ChangeFlags() QPinchGesture__ChangeFlag {
	if ptr.Pointer() != nil {
		return QPinchGesture__ChangeFlag(C.QPinchGesture_ChangeFlags(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPinchGesture) SetCenterPoint(value core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QPinchGesture_SetCenterPoint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(value)))
	}
}

func (ptr *QPinchGesture) SetChangeFlags(value QPinchGesture__ChangeFlag) {
	if ptr.Pointer() != nil {
		C.QPinchGesture_SetChangeFlags(C.QtObjectPtr(ptr.Pointer()), C.int(value))
	}
}

func (ptr *QPinchGesture) SetLastCenterPoint(value core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QPinchGesture_SetLastCenterPoint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(value)))
	}
}

func (ptr *QPinchGesture) SetStartCenterPoint(value core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QPinchGesture_SetStartCenterPoint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(value)))
	}
}

func (ptr *QPinchGesture) SetTotalChangeFlags(value QPinchGesture__ChangeFlag) {
	if ptr.Pointer() != nil {
		C.QPinchGesture_SetTotalChangeFlags(C.QtObjectPtr(ptr.Pointer()), C.int(value))
	}
}

func (ptr *QPinchGesture) TotalChangeFlags() QPinchGesture__ChangeFlag {
	if ptr.Pointer() != nil {
		return QPinchGesture__ChangeFlag(C.QPinchGesture_TotalChangeFlags(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPinchGesture) DestroyQPinchGesture() {
	if ptr.Pointer() != nil {
		C.QPinchGesture_DestroyQPinchGesture(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
