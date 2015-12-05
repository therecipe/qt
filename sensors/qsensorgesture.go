package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
		n.SetObjectName("QSensorGesture_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSensorGesture) QSensorGesture_PTR() *QSensorGesture {
	return ptr
}

func NewQSensorGesture(ids []string, parent core.QObject_ITF) *QSensorGesture {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensorGesture::QSensorGesture")
		}
	}()

	return NewQSensorGestureFromPointer(C.QSensorGesture_NewQSensorGesture(C.CString(strings.Join(ids, ",,,")), core.PointerFromQObject(parent)))
}

func (ptr *QSensorGesture) GestureSignals() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensorGesture::gestureSignals")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSensorGesture_GestureSignals(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QSensorGesture) InvalidIds() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensorGesture::invalidIds")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSensorGesture_InvalidIds(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QSensorGesture) IsActive() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensorGesture::isActive")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSensorGesture_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensorGesture) StartDetection() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensorGesture::startDetection")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensorGesture_StartDetection(ptr.Pointer())
	}
}

func (ptr *QSensorGesture) StopDetection() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensorGesture::stopDetection")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensorGesture_StopDetection(ptr.Pointer())
	}
}

func (ptr *QSensorGesture) ValidIds() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensorGesture::validIds")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSensorGesture_ValidIds(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QSensorGesture) DestroyQSensorGesture() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensorGesture::~QSensorGesture")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensorGesture_DestroyQSensorGesture(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
