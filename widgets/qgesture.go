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

type QGesture_ITF interface {
	core.QObject_ITF
	QGesture_PTR() *QGesture
}

func PointerFromQGesture(ptr QGesture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGesture_PTR().Pointer()
	}
	return nil
}

func NewQGestureFromPointer(ptr unsafe.Pointer) *QGesture {
	var n = new(QGesture)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGesture_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGesture) QGesture_PTR() *QGesture {
	return ptr
}

//QGesture::GestureCancelPolicy
type QGesture__GestureCancelPolicy int64

const (
	QGesture__CancelNone         = QGesture__GestureCancelPolicy(0)
	QGesture__CancelAllInContext = QGesture__GestureCancelPolicy(1)
)

func (ptr *QGesture) GestureCancelPolicy() QGesture__GestureCancelPolicy {
	if ptr.Pointer() != nil {
		return QGesture__GestureCancelPolicy(C.QGesture_GestureCancelPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGesture) GestureType() core.Qt__GestureType {
	if ptr.Pointer() != nil {
		return core.Qt__GestureType(C.QGesture_GestureType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGesture) HasHotSpot() bool {
	if ptr.Pointer() != nil {
		return C.QGesture_HasHotSpot(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGesture) SetGestureCancelPolicy(policy QGesture__GestureCancelPolicy) {
	if ptr.Pointer() != nil {
		C.QGesture_SetGestureCancelPolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QGesture) SetHotSpot(value core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QGesture_SetHotSpot(ptr.Pointer(), core.PointerFromQPointF(value))
	}
}

func (ptr *QGesture) UnsetHotSpot() {
	if ptr.Pointer() != nil {
		C.QGesture_UnsetHotSpot(ptr.Pointer())
	}
}

func NewQGesture(parent core.QObject_ITF) *QGesture {
	return NewQGestureFromPointer(C.QGesture_NewQGesture(core.PointerFromQObject(parent)))
}

func (ptr *QGesture) DestroyQGesture() {
	if ptr.Pointer() != nil {
		C.QGesture_DestroyQGesture(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
