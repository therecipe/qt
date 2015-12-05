package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QMediaAvailabilityControl struct {
	QMediaControl
}

type QMediaAvailabilityControl_ITF interface {
	QMediaControl_ITF
	QMediaAvailabilityControl_PTR() *QMediaAvailabilityControl
}

func PointerFromQMediaAvailabilityControl(ptr QMediaAvailabilityControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaAvailabilityControl_PTR().Pointer()
	}
	return nil
}

func NewQMediaAvailabilityControlFromPointer(ptr unsafe.Pointer) *QMediaAvailabilityControl {
	var n = new(QMediaAvailabilityControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMediaAvailabilityControl_") {
		n.SetObjectName("QMediaAvailabilityControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaAvailabilityControl) QMediaAvailabilityControl_PTR() *QMediaAvailabilityControl {
	return ptr
}

func (ptr *QMediaAvailabilityControl) DestroyQMediaAvailabilityControl() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaAvailabilityControl::~QMediaAvailabilityControl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaAvailabilityControl_DestroyQMediaAvailabilityControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
