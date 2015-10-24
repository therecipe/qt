package widgets

//#include "qgesture.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGesture struct {
	core.QObject
}

type QGestureITF interface {
	core.QObjectITF
	QGesturePTR() *QGesture
}

func PointerFromQGesture(ptr QGestureITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGesturePTR().Pointer()
	}
	return nil
}

func QGestureFromPointer(ptr unsafe.Pointer) *QGesture {
	var n = new(QGesture)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGesture_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGesture) QGesturePTR() *QGesture {
	return ptr
}

//QGesture::GestureCancelPolicy
type QGesture__GestureCancelPolicy int

var (
	QGesture__CancelNone         = QGesture__GestureCancelPolicy(0)
	QGesture__CancelAllInContext = QGesture__GestureCancelPolicy(1)
)

func (ptr *QGesture) GestureCancelPolicy() QGesture__GestureCancelPolicy {
	if ptr.Pointer() != nil {
		return QGesture__GestureCancelPolicy(C.QGesture_GestureCancelPolicy(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGesture) GestureType() core.Qt__GestureType {
	if ptr.Pointer() != nil {
		return core.Qt__GestureType(C.QGesture_GestureType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGesture) HasHotSpot() bool {
	if ptr.Pointer() != nil {
		return C.QGesture_HasHotSpot(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGesture) SetGestureCancelPolicy(policy QGesture__GestureCancelPolicy) {
	if ptr.Pointer() != nil {
		C.QGesture_SetGestureCancelPolicy(C.QtObjectPtr(ptr.Pointer()), C.int(policy))
	}
}

func (ptr *QGesture) SetHotSpot(value core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QGesture_SetHotSpot(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(value)))
	}
}

func (ptr *QGesture) UnsetHotSpot() {
	if ptr.Pointer() != nil {
		C.QGesture_UnsetHotSpot(C.QtObjectPtr(ptr.Pointer()))
	}
}

func NewQGesture(parent core.QObjectITF) *QGesture {
	return QGestureFromPointer(unsafe.Pointer(C.QGesture_NewQGesture(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QGesture) DestroyQGesture() {
	if ptr.Pointer() != nil {
		C.QGesture_DestroyQGesture(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
