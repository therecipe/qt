package sensors

//#include "qsensorgesture.h"
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

type QSensorGestureITF interface {
	core.QObjectITF
	QSensorGesturePTR() *QSensorGesture
}

func PointerFromQSensorGesture(ptr QSensorGestureITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorGesturePTR().Pointer()
	}
	return nil
}

func QSensorGestureFromPointer(ptr unsafe.Pointer) *QSensorGesture {
	var n = new(QSensorGesture)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSensorGesture_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSensorGesture) QSensorGesturePTR() *QSensorGesture {
	return ptr
}

func NewQSensorGesture(ids []string, parent core.QObjectITF) *QSensorGesture {
	return QSensorGestureFromPointer(unsafe.Pointer(C.QSensorGesture_NewQSensorGesture(C.CString(strings.Join(ids, "|")), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QSensorGesture) GestureSignals() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSensorGesture_GestureSignals(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QSensorGesture) InvalidIds() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSensorGesture_InvalidIds(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QSensorGesture) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QSensorGesture_IsActive(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSensorGesture) StartDetection() {
	if ptr.Pointer() != nil {
		C.QSensorGesture_StartDetection(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSensorGesture) StopDetection() {
	if ptr.Pointer() != nil {
		C.QSensorGesture_StopDetection(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSensorGesture) ValidIds() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSensorGesture_ValidIds(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QSensorGesture) DestroyQSensorGesture() {
	if ptr.Pointer() != nil {
		C.QSensorGesture_DestroyQSensorGesture(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
