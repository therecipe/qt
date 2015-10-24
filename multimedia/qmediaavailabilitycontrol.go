package multimedia

//#include "qmediaavailabilitycontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QMediaAvailabilityControl struct {
	QMediaControl
}

type QMediaAvailabilityControlITF interface {
	QMediaControlITF
	QMediaAvailabilityControlPTR() *QMediaAvailabilityControl
}

func PointerFromQMediaAvailabilityControl(ptr QMediaAvailabilityControlITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaAvailabilityControlPTR().Pointer()
	}
	return nil
}

func QMediaAvailabilityControlFromPointer(ptr unsafe.Pointer) *QMediaAvailabilityControl {
	var n = new(QMediaAvailabilityControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMediaAvailabilityControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaAvailabilityControl) QMediaAvailabilityControlPTR() *QMediaAvailabilityControl {
	return ptr
}

func (ptr *QMediaAvailabilityControl) DestroyQMediaAvailabilityControl() {
	if ptr.Pointer() != nil {
		C.QMediaAvailabilityControl_DestroyQMediaAvailabilityControl(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
