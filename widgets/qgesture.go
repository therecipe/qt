package widgets

//#include "widgets.h"
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
	for len(n.ObjectName()) < len("QGesture_") {
		n.SetObjectName("QGesture_" + qt.Identifier())
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
	defer qt.Recovering("QGesture::gestureCancelPolicy")

	if ptr.Pointer() != nil {
		return QGesture__GestureCancelPolicy(C.QGesture_GestureCancelPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGesture) GestureType() core.Qt__GestureType {
	defer qt.Recovering("QGesture::gestureType")

	if ptr.Pointer() != nil {
		return core.Qt__GestureType(C.QGesture_GestureType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGesture) HasHotSpot() bool {
	defer qt.Recovering("QGesture::hasHotSpot")

	if ptr.Pointer() != nil {
		return C.QGesture_HasHotSpot(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGesture) SetGestureCancelPolicy(policy QGesture__GestureCancelPolicy) {
	defer qt.Recovering("QGesture::setGestureCancelPolicy")

	if ptr.Pointer() != nil {
		C.QGesture_SetGestureCancelPolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QGesture) SetHotSpot(value core.QPointF_ITF) {
	defer qt.Recovering("QGesture::setHotSpot")

	if ptr.Pointer() != nil {
		C.QGesture_SetHotSpot(ptr.Pointer(), core.PointerFromQPointF(value))
	}
}

func (ptr *QGesture) State() core.Qt__GestureState {
	defer qt.Recovering("QGesture::state")

	if ptr.Pointer() != nil {
		return core.Qt__GestureState(C.QGesture_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGesture) UnsetHotSpot() {
	defer qt.Recovering("QGesture::unsetHotSpot")

	if ptr.Pointer() != nil {
		C.QGesture_UnsetHotSpot(ptr.Pointer())
	}
}

func NewQGesture(parent core.QObject_ITF) *QGesture {
	defer qt.Recovering("QGesture::QGesture")

	return NewQGestureFromPointer(C.QGesture_NewQGesture(core.PointerFromQObject(parent)))
}

func (ptr *QGesture) DestroyQGesture() {
	defer qt.Recovering("QGesture::~QGesture")

	if ptr.Pointer() != nil {
		C.QGesture_DestroyQGesture(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
