package sensors

//#include "sensors.h"
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

type QSensorGestureRecognizer_ITF interface {
	core.QObject_ITF
	QSensorGestureRecognizer_PTR() *QSensorGestureRecognizer
}

func PointerFromQSensorGestureRecognizer(ptr QSensorGestureRecognizer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorGestureRecognizer_PTR().Pointer()
	}
	return nil
}

func NewQSensorGestureRecognizerFromPointer(ptr unsafe.Pointer) *QSensorGestureRecognizer {
	var n = new(QSensorGestureRecognizer)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSensorGestureRecognizer_") {
		n.SetObjectName("QSensorGestureRecognizer_" + qt.Identifier())
	}
	return n
}

func (ptr *QSensorGestureRecognizer) QSensorGestureRecognizer_PTR() *QSensorGestureRecognizer {
	return ptr
}

func (ptr *QSensorGestureRecognizer) CreateBackend() {
	defer qt.Recovering("QSensorGestureRecognizer::createBackend")

	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_CreateBackend(ptr.Pointer())
	}
}

func (ptr *QSensorGestureRecognizer) GestureSignals() []string {
	defer qt.Recovering("QSensorGestureRecognizer::gestureSignals")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSensorGestureRecognizer_GestureSignals(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QSensorGestureRecognizer) Id() string {
	defer qt.Recovering("QSensorGestureRecognizer::id")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSensorGestureRecognizer_Id(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSensorGestureRecognizer) IsActive() bool {
	defer qt.Recovering("QSensorGestureRecognizer::isActive")

	if ptr.Pointer() != nil {
		return C.QSensorGestureRecognizer_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensorGestureRecognizer) StartBackend() {
	defer qt.Recovering("QSensorGestureRecognizer::startBackend")

	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_StartBackend(ptr.Pointer())
	}
}

func (ptr *QSensorGestureRecognizer) StopBackend() {
	defer qt.Recovering("QSensorGestureRecognizer::stopBackend")

	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_StopBackend(ptr.Pointer())
	}
}

func (ptr *QSensorGestureRecognizer) DestroyQSensorGestureRecognizer() {
	defer qt.Recovering("QSensorGestureRecognizer::~QSensorGestureRecognizer")

	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_DestroyQSensorGestureRecognizer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
