package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QCameraViewfinderSettingsControl2 struct {
	QMediaControl
}

type QCameraViewfinderSettingsControl2_ITF interface {
	QMediaControl_ITF
	QCameraViewfinderSettingsControl2_PTR() *QCameraViewfinderSettingsControl2
}

func PointerFromQCameraViewfinderSettingsControl2(ptr QCameraViewfinderSettingsControl2_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraViewfinderSettingsControl2_PTR().Pointer()
	}
	return nil
}

func NewQCameraViewfinderSettingsControl2FromPointer(ptr unsafe.Pointer) *QCameraViewfinderSettingsControl2 {
	var n = new(QCameraViewfinderSettingsControl2)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCameraViewfinderSettingsControl2_") {
		n.SetObjectName("QCameraViewfinderSettingsControl2_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCameraViewfinderSettingsControl2) QCameraViewfinderSettingsControl2_PTR() *QCameraViewfinderSettingsControl2 {
	return ptr
}

func (ptr *QCameraViewfinderSettingsControl2) SetViewfinderSettings(settings QCameraViewfinderSettings_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraViewfinderSettingsControl2::setViewfinderSettings")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl2_SetViewfinderSettings(ptr.Pointer(), PointerFromQCameraViewfinderSettings(settings))
	}
}

func (ptr *QCameraViewfinderSettingsControl2) DestroyQCameraViewfinderSettingsControl2() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraViewfinderSettingsControl2::~QCameraViewfinderSettingsControl2")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl2_DestroyQCameraViewfinderSettingsControl2(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
