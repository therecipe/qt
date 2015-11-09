package multimedia

//#include "qcameraimagecapturecontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QCameraImageCaptureControl struct {
	QMediaControl
}

type QCameraImageCaptureControl_ITF interface {
	QMediaControl_ITF
	QCameraImageCaptureControl_PTR() *QCameraImageCaptureControl
}

func PointerFromQCameraImageCaptureControl(ptr QCameraImageCaptureControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraImageCaptureControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraImageCaptureControlFromPointer(ptr unsafe.Pointer) *QCameraImageCaptureControl {
	var n = new(QCameraImageCaptureControl)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QCameraImageCaptureControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCameraImageCaptureControl) QCameraImageCaptureControl_PTR() *QCameraImageCaptureControl {
	return ptr
}

func (ptr *QCameraImageCaptureControl) CancelCapture() {
	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_CancelCapture(ptr.Pointer())
	}
}

func (ptr *QCameraImageCaptureControl) Capture(fileName string) int {
	if ptr.Pointer() != nil {
		return int(C.QCameraImageCaptureControl_Capture(ptr.Pointer(), C.CString(fileName)))
	}
	return 0
}

func (ptr *QCameraImageCaptureControl) DriveMode() QCameraImageCapture__DriveMode {
	if ptr.Pointer() != nil {
		return QCameraImageCapture__DriveMode(C.QCameraImageCaptureControl_DriveMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraImageCaptureControl) ConnectError(f func(id int, error int, errorString string)) {
	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QCameraImageCaptureControl) DisconnectError() {
	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQCameraImageCaptureControlError
func callbackQCameraImageCaptureControlError(ptrName *C.char, id C.int, error C.int, errorString *C.char) {
	qt.GetSignal(C.GoString(ptrName), "error").(func(int, int, string))(int(id), int(error), C.GoString(errorString))
}

func (ptr *QCameraImageCaptureControl) ConnectImageExposed(f func(requestId int)) {
	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ConnectImageExposed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "imageExposed", f)
	}
}

func (ptr *QCameraImageCaptureControl) DisconnectImageExposed() {
	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_DisconnectImageExposed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "imageExposed")
	}
}

//export callbackQCameraImageCaptureControlImageExposed
func callbackQCameraImageCaptureControlImageExposed(ptrName *C.char, requestId C.int) {
	qt.GetSignal(C.GoString(ptrName), "imageExposed").(func(int))(int(requestId))
}

func (ptr *QCameraImageCaptureControl) ConnectImageMetadataAvailable(f func(id int, key string, value *core.QVariant)) {
	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ConnectImageMetadataAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "imageMetadataAvailable", f)
	}
}

func (ptr *QCameraImageCaptureControl) DisconnectImageMetadataAvailable() {
	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_DisconnectImageMetadataAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "imageMetadataAvailable")
	}
}

//export callbackQCameraImageCaptureControlImageMetadataAvailable
func callbackQCameraImageCaptureControlImageMetadataAvailable(ptrName *C.char, id C.int, key *C.char, value unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "imageMetadataAvailable").(func(int, string, *core.QVariant))(int(id), C.GoString(key), core.NewQVariantFromPointer(value))
}

func (ptr *QCameraImageCaptureControl) ConnectImageSaved(f func(requestId int, fileName string)) {
	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ConnectImageSaved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "imageSaved", f)
	}
}

func (ptr *QCameraImageCaptureControl) DisconnectImageSaved() {
	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_DisconnectImageSaved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "imageSaved")
	}
}

//export callbackQCameraImageCaptureControlImageSaved
func callbackQCameraImageCaptureControlImageSaved(ptrName *C.char, requestId C.int, fileName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "imageSaved").(func(int, string))(int(requestId), C.GoString(fileName))
}

func (ptr *QCameraImageCaptureControl) IsReadyForCapture() bool {
	if ptr.Pointer() != nil {
		return C.QCameraImageCaptureControl_IsReadyForCapture(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCameraImageCaptureControl) ConnectReadyForCaptureChanged(f func(ready bool)) {
	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ConnectReadyForCaptureChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "readyForCaptureChanged", f)
	}
}

func (ptr *QCameraImageCaptureControl) DisconnectReadyForCaptureChanged() {
	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_DisconnectReadyForCaptureChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "readyForCaptureChanged")
	}
}

//export callbackQCameraImageCaptureControlReadyForCaptureChanged
func callbackQCameraImageCaptureControlReadyForCaptureChanged(ptrName *C.char, ready C.int) {
	qt.GetSignal(C.GoString(ptrName), "readyForCaptureChanged").(func(bool))(int(ready) != 0)
}

func (ptr *QCameraImageCaptureControl) SetDriveMode(mode QCameraImageCapture__DriveMode) {
	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_SetDriveMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraImageCaptureControl) DestroyQCameraImageCaptureControl() {
	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_DestroyQCameraImageCaptureControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
