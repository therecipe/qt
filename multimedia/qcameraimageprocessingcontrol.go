package multimedia

//#include "qcameraimageprocessingcontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QCameraImageProcessingControl struct {
	QMediaControl
}

type QCameraImageProcessingControl_ITF interface {
	QMediaControl_ITF
	QCameraImageProcessingControl_PTR() *QCameraImageProcessingControl
}

func PointerFromQCameraImageProcessingControl(ptr QCameraImageProcessingControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraImageProcessingControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraImageProcessingControlFromPointer(ptr unsafe.Pointer) *QCameraImageProcessingControl {
	var n = new(QCameraImageProcessingControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QCameraImageProcessingControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCameraImageProcessingControl) QCameraImageProcessingControl_PTR() *QCameraImageProcessingControl {
	return ptr
}

//QCameraImageProcessingControl::ProcessingParameter
type QCameraImageProcessingControl__ProcessingParameter int64

const (
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
		return C.QCameraImageProcessingControl_IsParameterSupported(ptr.Pointer(), C.int(parameter)) != 0
	}
	return false
}

func (ptr *QCameraImageProcessingControl) IsParameterValueSupported(parameter QCameraImageProcessingControl__ProcessingParameter, value core.QVariant_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QCameraImageProcessingControl_IsParameterValueSupported(ptr.Pointer(), C.int(parameter), core.PointerFromQVariant(value)) != 0
	}
	return false
}

func (ptr *QCameraImageProcessingControl) Parameter(parameter QCameraImageProcessingControl__ProcessingParameter) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QCameraImageProcessingControl_Parameter(ptr.Pointer(), C.int(parameter)))
	}
	return nil
}

func (ptr *QCameraImageProcessingControl) SetParameter(parameter QCameraImageProcessingControl__ProcessingParameter, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QCameraImageProcessingControl_SetParameter(ptr.Pointer(), C.int(parameter), core.PointerFromQVariant(value))
	}
}

func (ptr *QCameraImageProcessingControl) DestroyQCameraImageProcessingControl() {
	if ptr.Pointer() != nil {
		C.QCameraImageProcessingControl_DestroyQCameraImageProcessingControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
