package multimediawidgets

//#include "qcameraviewfinder.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/multimedia"
	"github.com/therecipe/qt/widgets"
	"unsafe"
)

type QCameraViewfinder struct {
	QVideoWidget
}

type QCameraViewfinder_ITF interface {
	QVideoWidget_ITF
	QCameraViewfinder_PTR() *QCameraViewfinder
}

func PointerFromQCameraViewfinder(ptr QCameraViewfinder_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraViewfinder_PTR().Pointer()
	}
	return nil
}

func NewQCameraViewfinderFromPointer(ptr unsafe.Pointer) *QCameraViewfinder {
	var n = new(QCameraViewfinder)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QCameraViewfinder_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCameraViewfinder) QCameraViewfinder_PTR() *QCameraViewfinder {
	return ptr
}

func NewQCameraViewfinder(parent widgets.QWidget_ITF) *QCameraViewfinder {
	return NewQCameraViewfinderFromPointer(C.QCameraViewfinder_NewQCameraViewfinder(widgets.PointerFromQWidget(parent)))
}

func (ptr *QCameraViewfinder) MediaObject() *multimedia.QMediaObject {
	if ptr.Pointer() != nil {
		return multimedia.NewQMediaObjectFromPointer(C.QCameraViewfinder_MediaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCameraViewfinder) DestroyQCameraViewfinder() {
	if ptr.Pointer() != nil {
		C.QCameraViewfinder_DestroyQCameraViewfinder(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
