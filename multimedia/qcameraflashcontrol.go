package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QCameraFlashControl struct {
	QMediaControl
}

type QCameraFlashControl_ITF interface {
	QMediaControl_ITF
	QCameraFlashControl_PTR() *QCameraFlashControl
}

func PointerFromQCameraFlashControl(ptr QCameraFlashControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraFlashControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraFlashControlFromPointer(ptr unsafe.Pointer) *QCameraFlashControl {
	var n = new(QCameraFlashControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCameraFlashControl_") {
		n.SetObjectName("QCameraFlashControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCameraFlashControl) QCameraFlashControl_PTR() *QCameraFlashControl {
	return ptr
}

func (ptr *QCameraFlashControl) FlashMode() QCameraExposure__FlashMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFlashControl::flashMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QCameraExposure__FlashMode(C.QCameraFlashControl_FlashMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFlashControl) ConnectFlashReady(f func(ready bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFlashControl::flashReady")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraFlashControl_ConnectFlashReady(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "flashReady", f)
	}
}

func (ptr *QCameraFlashControl) DisconnectFlashReady() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFlashControl::flashReady")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraFlashControl_DisconnectFlashReady(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "flashReady")
	}
}

//export callbackQCameraFlashControlFlashReady
func callbackQCameraFlashControlFlashReady(ptrName *C.char, ready C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFlashControl::flashReady")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "flashReady").(func(bool))(int(ready) != 0)
}

func (ptr *QCameraFlashControl) IsFlashModeSupported(mode QCameraExposure__FlashMode) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFlashControl::isFlashModeSupported")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QCameraFlashControl_IsFlashModeSupported(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraFlashControl) IsFlashReady() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFlashControl::isFlashReady")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QCameraFlashControl_IsFlashReady(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCameraFlashControl) SetFlashMode(mode QCameraExposure__FlashMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFlashControl::setFlashMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraFlashControl_SetFlashMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraFlashControl) DestroyQCameraFlashControl() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFlashControl::~QCameraFlashControl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraFlashControl_DestroyQCameraFlashControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
