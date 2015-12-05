package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QCameraCaptureDestinationControl struct {
	QMediaControl
}

type QCameraCaptureDestinationControl_ITF interface {
	QMediaControl_ITF
	QCameraCaptureDestinationControl_PTR() *QCameraCaptureDestinationControl
}

func PointerFromQCameraCaptureDestinationControl(ptr QCameraCaptureDestinationControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraCaptureDestinationControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraCaptureDestinationControlFromPointer(ptr unsafe.Pointer) *QCameraCaptureDestinationControl {
	var n = new(QCameraCaptureDestinationControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCameraCaptureDestinationControl_") {
		n.SetObjectName("QCameraCaptureDestinationControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCameraCaptureDestinationControl) QCameraCaptureDestinationControl_PTR() *QCameraCaptureDestinationControl {
	return ptr
}

func (ptr *QCameraCaptureDestinationControl) CaptureDestination() QCameraImageCapture__CaptureDestination {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraCaptureDestinationControl::captureDestination")
		}
	}()

	if ptr.Pointer() != nil {
		return QCameraImageCapture__CaptureDestination(C.QCameraCaptureDestinationControl_CaptureDestination(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraCaptureDestinationControl) ConnectCaptureDestinationChanged(f func(destination QCameraImageCapture__CaptureDestination)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraCaptureDestinationControl::captureDestinationChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraCaptureDestinationControl_ConnectCaptureDestinationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "captureDestinationChanged", f)
	}
}

func (ptr *QCameraCaptureDestinationControl) DisconnectCaptureDestinationChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraCaptureDestinationControl::captureDestinationChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraCaptureDestinationControl_DisconnectCaptureDestinationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "captureDestinationChanged")
	}
}

//export callbackQCameraCaptureDestinationControlCaptureDestinationChanged
func callbackQCameraCaptureDestinationControlCaptureDestinationChanged(ptrName *C.char, destination C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraCaptureDestinationControl::captureDestinationChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "captureDestinationChanged").(func(QCameraImageCapture__CaptureDestination))(QCameraImageCapture__CaptureDestination(destination))
}

func (ptr *QCameraCaptureDestinationControl) IsCaptureDestinationSupported(destination QCameraImageCapture__CaptureDestination) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraCaptureDestinationControl::isCaptureDestinationSupported")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QCameraCaptureDestinationControl_IsCaptureDestinationSupported(ptr.Pointer(), C.int(destination)) != 0
	}
	return false
}

func (ptr *QCameraCaptureDestinationControl) SetCaptureDestination(destination QCameraImageCapture__CaptureDestination) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraCaptureDestinationControl::setCaptureDestination")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraCaptureDestinationControl_SetCaptureDestination(ptr.Pointer(), C.int(destination))
	}
}

func (ptr *QCameraCaptureDestinationControl) DestroyQCameraCaptureDestinationControl() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraCaptureDestinationControl::~QCameraCaptureDestinationControl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraCaptureDestinationControl_DestroyQCameraCaptureDestinationControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
