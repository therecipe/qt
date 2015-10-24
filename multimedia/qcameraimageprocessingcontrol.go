package multimedia

//#include "qcameraimageprocessingcontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QCameraImageProcessingControl struct {
	QMediaControl
}

type QCameraImageProcessingControlITF interface {
	QMediaControlITF
	QCameraImageProcessingControlPTR() *QCameraImageProcessingControl
}

func PointerFromQCameraImageProcessingControl(ptr QCameraImageProcessingControlITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraImageProcessingControlPTR().Pointer()
	}
	return nil
}

func QCameraImageProcessingControlFromPointer(ptr unsafe.Pointer) *QCameraImageProcessingControl {
	var n = new(QCameraImageProcessingControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QCameraImageProcessingControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCameraImageProcessingControl) QCameraImageProcessingControlPTR() *QCameraImageProcessingControl {
	return ptr
}

//QCameraImageProcessingControl::ProcessingParameter
type QCameraImageProcessingControl__ProcessingParameter int

var (
	QCameraImageProcessingControl__WhiteBalancePreset   = QCameraImageProcessingControl__ProcessingParameter(0)
	QCameraImageProcessingControl__ColorTemperature     = QCameraImageProcessingControl__ProcessingParameter(1)
	QCameraImageProcessingControl__Contrast             = QCameraImageProcessingControl__ProcessingParameter(2)
	QCameraImageProcessingControl__Saturation           = QCameraImageProcessingControl__ProcessingParameter(3)
	QCameraImageProcessingControl__Brightness           = QCameraImageProcessingControl__ProcessingParameter(4)
	QCameraImageProcessingControl__Sharpening           = QCameraImageProcessingControl__ProcessingParameter(5)
	QCameraImageProcessingControl__Denoising            = QCameraImageProcessingControl__ProcessingParameter(6)
	QCameraImageProcessingControl__ContrastAdjustment   = QCameraImageProcessingControl__ProcessingParameter(7)
	QCameraImageProcessingControl__SaturationAdjustment = QCameraImageProcessingControl__ProcessingParameter(8)
	QCameraImageProcessingControl__BrightnessAdjustment = QCameraImageProcessingControl__ProcessingParameter(9)
	QCameraImageProcessingControl__SharpeningAdjustment = QCameraImageProcessingControl__ProcessingParameter(10)
	QCameraImageProcessingControl__DenoisingAdjustment  = QCameraImageProcessingControl__ProcessingParameter(11)
	QCameraImageProcessingControl__ColorFilter          = QCameraImageProcessingControl__ProcessingParameter(12)
	QCameraImageProcessingControl__ExtendedParameter    = QCameraImageProcessingControl__ProcessingParameter(1000)
)

func (ptr *QCameraImageProcessingControl) IsParameterSupported(parameter QCameraImageProcessingControl__ProcessingParameter) bool {
	if ptr.Pointer() != nil {
		return C.QCameraImageProcessingControl_IsParameterSupported(C.QtObjectPtr(ptr.Pointer()), C.int(parameter)) != 0
	}
	return false
}

func (ptr *QCameraImageProcessingControl) IsParameterValueSupported(parameter QCameraImageProcessingControl__ProcessingParameter, value string) bool {
	if ptr.Pointer() != nil {
		return C.QCameraImageProcessingControl_IsParameterValueSupported(C.QtObjectPtr(ptr.Pointer()), C.int(parameter), C.CString(value)) != 0
	}
	return false
}

func (ptr *QCameraImageProcessingControl) Parameter(parameter QCameraImageProcessingControl__ProcessingParameter) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QCameraImageProcessingControl_Parameter(C.QtObjectPtr(ptr.Pointer()), C.int(parameter)))
	}
	return ""
}

func (ptr *QCameraImageProcessingControl) SetParameter(parameter QCameraImageProcessingControl__ProcessingParameter, value string) {
	if ptr.Pointer() != nil {
		C.QCameraImageProcessingControl_SetParameter(C.QtObjectPtr(ptr.Pointer()), C.int(parameter), C.CString(value))
	}
}

func (ptr *QCameraImageProcessingControl) DestroyQCameraImageProcessingControl() {
	if ptr.Pointer() != nil {
		C.QCameraImageProcessingControl_DestroyQCameraImageProcessingControl(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
