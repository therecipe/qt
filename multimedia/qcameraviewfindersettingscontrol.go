package multimedia

//#include "qcameraviewfindersettingscontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QCameraViewfinderSettingsControl struct {
	QMediaControl
}

type QCameraViewfinderSettingsControlITF interface {
	QMediaControlITF
	QCameraViewfinderSettingsControlPTR() *QCameraViewfinderSettingsControl
}

func PointerFromQCameraViewfinderSettingsControl(ptr QCameraViewfinderSettingsControlITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraViewfinderSettingsControlPTR().Pointer()
	}
	return nil
}

func QCameraViewfinderSettingsControlFromPointer(ptr unsafe.Pointer) *QCameraViewfinderSettingsControl {
	var n = new(QCameraViewfinderSettingsControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QCameraViewfinderSettingsControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCameraViewfinderSettingsControl) QCameraViewfinderSettingsControlPTR() *QCameraViewfinderSettingsControl {
	return ptr
}

//QCameraViewfinderSettingsControl::ViewfinderParameter
type QCameraViewfinderSettingsControl__ViewfinderParameter int

var (
	QCameraViewfinderSettingsControl__Resolution       = QCameraViewfinderSettingsControl__ViewfinderParameter(0)
	QCameraViewfinderSettingsControl__PixelAspectRatio = QCameraViewfinderSettingsControl__ViewfinderParameter(1)
	QCameraViewfinderSettingsControl__MinimumFrameRate = QCameraViewfinderSettingsControl__ViewfinderParameter(2)
	QCameraViewfinderSettingsControl__MaximumFrameRate = QCameraViewfinderSettingsControl__ViewfinderParameter(3)
	QCameraViewfinderSettingsControl__PixelFormat      = QCameraViewfinderSettingsControl__ViewfinderParameter(4)
	QCameraViewfinderSettingsControl__UserParameter    = QCameraViewfinderSettingsControl__ViewfinderParameter(1000)
)

func (ptr *QCameraViewfinderSettingsControl) IsViewfinderParameterSupported(parameter QCameraViewfinderSettingsControl__ViewfinderParameter) bool {
	if ptr.Pointer() != nil {
		return C.QCameraViewfinderSettingsControl_IsViewfinderParameterSupported(C.QtObjectPtr(ptr.Pointer()), C.int(parameter)) != 0
	}
	return false
}

func (ptr *QCameraViewfinderSettingsControl) SetViewfinderParameter(parameter QCameraViewfinderSettingsControl__ViewfinderParameter, value string) {
	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl_SetViewfinderParameter(C.QtObjectPtr(ptr.Pointer()), C.int(parameter), C.CString(value))
	}
}

func (ptr *QCameraViewfinderSettingsControl) ViewfinderParameter(parameter QCameraViewfinderSettingsControl__ViewfinderParameter) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QCameraViewfinderSettingsControl_ViewfinderParameter(C.QtObjectPtr(ptr.Pointer()), C.int(parameter)))
	}
	return ""
}

func (ptr *QCameraViewfinderSettingsControl) DestroyQCameraViewfinderSettingsControl() {
	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl_DestroyQCameraViewfinderSettingsControl(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
