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

type QCameraInfoControlITF interface {
	QMediaControlITF
	QCameraInfoControlPTR() *QCameraInfoControl
}

func PointerFromQCameraInfoControl(ptr QCameraInfoControlITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraInfoControlPTR().Pointer()
	}
	return nil
}

func QCameraInfoControlFromPointer(ptr unsafe.Pointer) *QCameraInfoControl {
	var n = new(QCameraInfoControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QCameraInfoControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCameraInfoControl) QCameraInfoControlPTR() *QCameraInfoControl {
	return ptr
}

func (ptr *QCameraInfoControl) CameraOrientation(deviceName string) int {
	if ptr.Pointer() != nil {
		return int(C.QCameraInfoControl_CameraOrientation(C.QtObjectPtr(ptr.Pointer()), C.CString(deviceName)))
	}
	return 0
}

func (ptr *QCameraInfoControl) CameraPosition(deviceName string) QCamera__Position {
	if ptr.Pointer() != nil {
		return QCamera__Position(C.QCameraInfoControl_CameraPosition(C.QtObjectPtr(ptr.Pointer()), C.CString(deviceName)))
	}
	return 0
}

func (ptr *QCameraInfoControl) DestroyQCameraInfoControl() {
	if ptr.Pointer() != nil {
		C.QCameraInfoControl_DestroyQCameraInfoControl(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
