package multimedia

//#include "qcamerazoomcontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QCameraZoomControl struct {
	QMediaControl
}

type QCameraZoomControlITF interface {
	QMediaControlITF
	QCameraZoomControlPTR() *QCameraZoomControl
}

func PointerFromQCameraZoomControl(ptr QCameraZoomControlITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraZoomControlPTR().Pointer()
	}
	return nil
}

func QCameraZoomControlFromPointer(ptr unsafe.Pointer) *QCameraZoomControl {
	var n = new(QCameraZoomControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QCameraZoomControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCameraZoomControl) QCameraZoomControlPTR() *QCameraZoomControl {
	return ptr
}

func (ptr *QCameraZoomControl) DestroyQCameraZoomControl() {
	if ptr.Pointer() != nil {
		C.QCameraZoomControl_DestroyQCameraZoomControl(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
