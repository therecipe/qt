package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
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
	for len(n.ObjectName()) < len("QCameraInfoControl_") {
		n.SetObjectName("QCameraInfoControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCameraInfoControl) QCameraInfoControl_PTR() *QCameraInfoControl {
	return ptr
}

func (ptr *QCameraInfoControl) CameraOrientation(deviceName string) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraInfoControl::cameraOrientation")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QCameraInfoControl_CameraOrientation(ptr.Pointer(), C.CString(deviceName)))
	}
	return 0
}

func (ptr *QCameraInfoControl) CameraPosition(deviceName string) QCamera__Position {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraInfoControl::cameraPosition")
		}
	}()

	if ptr.Pointer() != nil {
		return QCamera__Position(C.QCameraInfoControl_CameraPosition(ptr.Pointer(), C.CString(deviceName)))
	}
	return 0
}

func (ptr *QCameraInfoControl) DestroyQCameraInfoControl() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraInfoControl::~QCameraInfoControl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraInfoControl_DestroyQCameraInfoControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
