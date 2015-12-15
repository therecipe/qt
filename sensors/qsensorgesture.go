package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QSensorGesture struct {
	core.QObject
}

type QSensorGesture_ITF interface {
	core.QObject_ITF
	QSensorGesture_PTR() *QSensorGesture
}

func PointerFromQSensorGesture(ptr QSensorGesture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorGesture_PTR().Pointer()
	}
	return nil
}

func NewQSensorGestureFromPointer(ptr unsafe.Pointer) *QSensorGesture {
	var n = new(QSensorGesture)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSensorGesture_") {
		n.SetObjectName("QSensorGesture_" + qt.Identifier())
	}
	return n
}

func (ptr *QSensorGesture) QSensorGesture_PTR() *QSensorGesture {
	return ptr
}

func NewQSensorGesture(ids []string, parent core.QObject_ITF) *QSensorGesture {
	defer qt.Recovering("QSensorGesture::QSensorGesture")

	return NewQSensorGestureFromPointer(C.QSensorGesture_NewQSensorGesture(C.CString(strings.Join(ids, ",,,")), core.PointerFromQObject(parent)))
}

func (ptr *QSensorGesture) GestureSignals() []string {
	defer qt.Recovering("QSensorGesture::gestureSignals")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSensorGesture_GestureSignals(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QSensorGesture) InvalidIds() []string {
	defer qt.Recovering("QSensorGesture::invalidIds")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSensorGesture_InvalidIds(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QSensorGesture) IsActive() bool {
	defer qt.Recovering("QSensorGesture::isActive")

	if ptr.Pointer() != nil {
		return C.QSensorGesture_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensorGesture) StartDetection() {
	defer qt.Recovering("QSensorGesture::startDetection")

	if ptr.Pointer() != nil {
		C.QSensorGesture_StartDetection(ptr.Pointer())
	}
}

func (ptr *QSensorGesture) StopDetection() {
	defer qt.Recovering("QSensorGesture::stopDetection")

	if ptr.Pointer() != nil {
		C.QSensorGesture_StopDetection(ptr.Pointer())
	}
}

func (ptr *QSensorGesture) ValidIds() []string {
	defer qt.Recovering("QSensorGesture::validIds")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSensorGesture_ValidIds(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QSensorGesture) DestroyQSensorGesture() {
	defer qt.Recovering("QSensorGesture::~QSensorGesture")

	if ptr.Pointer() != nil {
		C.QSensorGesture_DestroyQSensorGesture(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
