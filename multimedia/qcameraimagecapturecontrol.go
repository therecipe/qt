package multimedia

//#include "qcameraimagecapturecontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QCameraImageCaptureControl struct {
	QMediaControl
}

type QCameraImageCaptureControlITF interface {
	QMediaControlITF
	QCameraImageCaptureControlPTR() *QCameraImageCaptureControl
}

func PointerFromQCameraImageCaptureControl(ptr QCameraImageCaptureControlITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraImageCaptureControlPTR().Pointer()
	}
	return nil
}

func QCameraImageCaptureControlFromPointer(ptr unsafe.Pointer) *QCameraImageCaptureControl {
	var n = new(QCameraImageCaptureControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QCameraImageCaptureControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCameraImageCaptureControl) QCameraImageCaptureControlPTR() *QCameraImageCaptureControl {
	return ptr
}

func (ptr *QCameraImageCaptureControl) CancelCapture() {
	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_CancelCapture(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QCameraImageCaptureControl) Capture(fileName string) int {
	if ptr.Pointer() != nil {
		return int(C.QCameraImageCaptureControl_Capture(C.QtObjectPtr(ptr.Pointer()), C.CString(fileName)))
	}
	return 0
}

func (ptr *QCameraImageCaptureControl) DriveMode() QCameraImageCapture__DriveMode {
	if ptr.Pointer() != nil {
		return QCameraImageCapture__DriveMode(C.QCameraImageCaptureControl_DriveMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCameraImageCaptureControl) ConnectError(f func(id int, error int, errorString string)) {
	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ConnectError(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QCameraImageCaptureControl) DisconnectError() {
	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_DisconnectError(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQCameraImageCaptureControlError
func callbackQCameraImageCaptureControlError(ptrName *C.char, id C.int, error C.int, errorString *C.char) {
	qt.GetSignal(C.GoString(ptrName), "error").(func(int, int, string))(int(id), int(error), C.GoString(errorString))
}

func (ptr *QCameraImageCaptureControl) ConnectImageExposed(f func(requestId int)) {
	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ConnectImageExposed(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "imageExposed", f)
	}
}

func (ptr *QCameraImageCaptureControl) DisconnectImageExposed() {
	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_DisconnectImageExposed(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "imageExposed")
	}
}

//export callbackQCameraImageCaptureControlImageExposed
func callbackQCameraImageCaptureControlImageExposed(ptrName *C.char, requestId C.int) {
	qt.GetSignal(C.GoString(ptrName), "imageExposed").(func(int))(int(requestId))
}

func (ptr *QCameraImageCaptureControl) ConnectImageMetadataAvailable(f func(id int, key string, value string)) {
	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ConnectImageMetadataAvailable(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "imageMetadataAvailable", f)
	}
}

func (ptr *QCameraImageCaptureControl) DisconnectImageMetadataAvailable() {
	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_DisconnectImageMetadataAvailable(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "imageMetadataAvailable")
	}
}

//export callbackQCameraImageCaptureControlImageMetadataAvailable
func callbackQCameraImageCaptureControlImageMetadataAvailable(ptrName *C.char, id C.int, key *C.char, value *C.char) {
	qt.GetSignal(C.GoString(ptrName), "imageMetadataAvailable").(func(int, string, string))(int(id), C.GoString(key), C.GoString(value))
}

func (ptr *QCameraImageCaptureControl) ConnectImageSaved(f func(requestId int, fileName string)) {
	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ConnectImageSaved(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "imageSaved", f)
	}
}

func (ptr *QCameraImageCaptureControl) DisconnectImageSaved() {
	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_DisconnectImageSaved(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "imageSaved")
	}
}

//export callbackQCameraImageCaptureControlImageSaved
func callbackQCameraImageCaptureControlImageSaved(ptrName *C.char, requestId C.int, fileName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "imageSaved").(func(int, string))(int(requestId), C.GoString(fileName))
}

func (ptr *QCameraImageCaptureControl) IsReadyForCapture() bool {
	if ptr.Pointer() != nil {
		return C.QCameraImageCaptureControl_IsReadyForCapture(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QCameraImageCaptureControl) ConnectReadyForCaptureChanged(f func(ready bool)) {
	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ConnectReadyForCaptureChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "readyForCaptureChanged", f)
	}
}

func (ptr *QCameraImageCaptureControl) DisconnectReadyForCaptureChanged() {
	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_DisconnectReadyForCaptureChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "readyForCaptureChanged")
	}
}

//export callbackQCameraImageCaptureControlReadyForCaptureChanged
func callbackQCameraImageCaptureControlReadyForCaptureChanged(ptrName *C.char, ready C.int) {
	qt.GetSignal(C.GoString(ptrName), "readyForCaptureChanged").(func(bool))(int(ready) != 0)
}

func (ptr *QCameraImageCaptureControl) SetDriveMode(mode QCameraImageCapture__DriveMode) {
	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_SetDriveMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QCameraImageCaptureControl) DestroyQCameraImageCaptureControl() {
	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_DestroyQCameraImageCaptureControl(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
