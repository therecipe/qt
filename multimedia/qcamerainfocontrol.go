package multimedia

//#include "qcamerainfocontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QCameraInfoControl struct {
	QMediaControl
}

type QCameraInfoControl_ITF interface {
	QMediaControl_ITF
	QCameraInfoControl_PTR() *QCameraInfoControl
}

func PointerFromQCameraInfoControl(ptr QCameraInfoControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraInfoControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraInfoControlFromPointer(ptr unsafe.Pointer) *QCameraInfoControl {
	var n = new(QCameraInfoControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QCameraInfoControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCameraInfoControl) QCameraInfoControl_PTR() *QCameraInfoControl {
	return ptr
}

func (ptr *QCameraInfoControl) CameraOrientation(deviceName string) int {
	if ptr.Pointer() != nil {
		return int(C.QCameraInfoControl_CameraOrientation(ptr.Pointer(), C.CString(deviceName)))
	}
	return 0
}

func (ptr *QCameraInfoControl) CameraPosition(deviceName string) QCamera__Position {
	if ptr.Pointer() != nil {
		return QCamera__Position(C.QCameraInfoControl_CameraPosition(ptr.Pointer(), C.CString(deviceName)))
	}
	return 0
}

func (ptr *QCameraInfoControl) DestroyQCameraInfoControl() {
	if ptr.Pointer() != nil {
		C.QCameraInfoControl_DestroyQCameraInfoControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
