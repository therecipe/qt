package sensors

//#include "qsensorgesturerecognizer.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QSensorGestureRecognizer struct {
	core.QObject
}

type QSensorGestureRecognizerITF interface {
	core.QObjectITF
	QSensorGestureRecognizerPTR() *QSensorGestureRecognizer
}

func PointerFromQSensorGestureRecognizer(ptr QSensorGestureRecognizerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorGestureRecognizerPTR().Pointer()
	}
	return nil
}

func QSensorGestureRecognizerFromPointer(ptr unsafe.Pointer) *QSensorGestureRecognizer {
	var n = new(QSensorGestureRecognizer)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSensorGestureRecognizer_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSensorGestureRecognizer) QSensorGestureRecognizerPTR() *QSensorGestureRecognizer {
	return ptr
}

func (ptr *QSensorGestureRecognizer) CreateBackend() {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_CreateBackend(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSensorGestureRecognizer) GestureSignals() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSensorGestureRecognizer_GestureSignals(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QSensorGestureRecognizer) Id() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSensorGestureRecognizer_Id(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSensorGestureRecognizer) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QSensorGestureRecognizer_IsActive(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSensorGestureRecognizer) StartBackend() {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_StartBackend(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSensorGestureRecognizer) StopBackend() {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_StopBackend(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSensorGestureRecognizer) DestroyQSensorGestureRecognizer() {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_DestroyQSensorGestureRecognizer(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
