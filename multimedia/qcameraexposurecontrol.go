package multimedia

//#include "qcameraexposurecontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QCameraExposureControl struct {
	QMediaControl
}

type QCameraExposureControlITF interface {
	QMediaControlITF
	QCameraExposureControlPTR() *QCameraExposureControl
}

func PointerFromQCameraExposureControl(ptr QCameraExposureControlITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraExposureControlPTR().Pointer()
	}
	return nil
}

func QCameraExposureControlFromPointer(ptr unsafe.Pointer) *QCameraExposureControl {
	var n = new(QCameraExposureControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QCameraExposureControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCameraExposureControl) QCameraExposureControlPTR() *QCameraExposureControl {
	return ptr
}

//QCameraExposureControl::ExposureParameter
type QCameraExposureControl__ExposureParameter int

var (
	QCameraExposureControl__ISO                       = QCameraExposureControl__ExposureParameter(0)
	QCameraExposureControl__Aperture                  = QCameraExposureControl__ExposureParameter(1)
	QCameraExposureControl__ShutterSpeed              = QCameraExposureControl__ExposureParameter(2)
	QCameraExposureControl__ExposureCompensation      = QCameraExposureControl__ExposureParameter(3)
	QCameraExposureControl__FlashPower                = QCameraExposureControl__ExposureParameter(4)
	QCameraExposureControl__FlashCompensation         = QCameraExposureControl__ExposureParameter(5)
	QCameraExposureControl__TorchPower                = QCameraExposureControl__ExposureParameter(6)
	QCameraExposureControl__SpotMeteringPoint         = QCameraExposureControl__ExposureParameter(7)
	QCameraExposureControl__ExposureMode              = QCameraExposureControl__ExposureParameter(8)
	QCameraExposureControl__MeteringMode              = QCameraExposureControl__ExposureParameter(9)
	QCameraExposureControl__ExtendedExposureParameter = QCameraExposureControl__ExposureParameter(1000)
)

func (ptr *QCameraExposureControl) ActualValue(parameter QCameraExposureControl__ExposureParameter) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QCameraExposureControl_ActualValue(C.QtObjectPtr(ptr.Pointer()), C.int(parameter)))
	}
	return ""
}

func (ptr *QCameraExposureControl) ConnectActualValueChanged(f func(parameter int)) {
	if ptr.Pointer() != nil {
		C.QCameraExposureControl_ConnectActualValueChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "actualValueChanged", f)
	}
}

func (ptr *QCameraExposureControl) DisconnectActualValueChanged() {
	if ptr.Pointer() != nil {
		C.QCameraExposureControl_DisconnectActualValueChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "actualValueChanged")
	}
}

//export callbackQCameraExposureControlActualValueChanged
func callbackQCameraExposureControlActualValueChanged(ptrName *C.char, parameter C.int) {
	qt.GetSignal(C.GoString(ptrName), "actualValueChanged").(func(int))(int(parameter))
}

func (ptr *QCameraExposureControl) IsParameterSupported(parameter QCameraExposureControl__ExposureParameter) bool {
	if ptr.Pointer() != nil {
		return C.QCameraExposureControl_IsParameterSupported(C.QtObjectPtr(ptr.Pointer()), C.int(parameter)) != 0
	}
	return false
}

func (ptr *QCameraExposureControl) ConnectParameterRangeChanged(f func(parameter int)) {
	if ptr.Pointer() != nil {
		C.QCameraExposureControl_ConnectParameterRangeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "parameterRangeChanged", f)
	}
}

func (ptr *QCameraExposureControl) DisconnectParameterRangeChanged() {
	if ptr.Pointer() != nil {
		C.QCameraExposureControl_DisconnectParameterRangeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "parameterRangeChanged")
	}
}

//export callbackQCameraExposureControlParameterRangeChanged
func callbackQCameraExposureControlParameterRangeChanged(ptrName *C.char, parameter C.int) {
	qt.GetSignal(C.GoString(ptrName), "parameterRangeChanged").(func(int))(int(parameter))
}

func (ptr *QCameraExposureControl) RequestedValue(parameter QCameraExposureControl__ExposureParameter) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QCameraExposureControl_RequestedValue(C.QtObjectPtr(ptr.Pointer()), C.int(parameter)))
	}
	return ""
}

func (ptr *QCameraExposureControl) ConnectRequestedValueChanged(f func(parameter int)) {
	if ptr.Pointer() != nil {
		C.QCameraExposureControl_ConnectRequestedValueChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "requestedValueChanged", f)
	}
}

func (ptr *QCameraExposureControl) DisconnectRequestedValueChanged() {
	if ptr.Pointer() != nil {
		C.QCameraExposureControl_DisconnectRequestedValueChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "requestedValueChanged")
	}
}

//export callbackQCameraExposureControlRequestedValueChanged
func callbackQCameraExposureControlRequestedValueChanged(ptrName *C.char, parameter C.int) {
	qt.GetSignal(C.GoString(ptrName), "requestedValueChanged").(func(int))(int(parameter))
}

func (ptr *QCameraExposureControl) SetValue(parameter QCameraExposureControl__ExposureParameter, value string) bool {
	if ptr.Pointer() != nil {
		return C.QCameraExposureControl_SetValue(C.QtObjectPtr(ptr.Pointer()), C.int(parameter), C.CString(value)) != 0
	}
	return false
}

func (ptr *QCameraExposureControl) DestroyQCameraExposureControl() {
	if ptr.Pointer() != nil {
		C.QCameraExposureControl_DestroyQCameraExposureControl(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
