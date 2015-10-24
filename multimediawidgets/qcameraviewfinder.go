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

type QCameraViewfinderITF interface {
	QVideoWidgetITF
	QCameraViewfinderPTR() *QCameraViewfinder
}

func PointerFromQCameraViewfinder(ptr QCameraViewfinderITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraViewfinderPTR().Pointer()
	}
	return nil
}

func QCameraViewfinderFromPointer(ptr unsafe.Pointer) *QCameraViewfinder {
	var n = new(QCameraViewfinder)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QCameraViewfinder_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCameraViewfinder) QCameraViewfinderPTR() *QCameraViewfinder {
	return ptr
}

func NewQCameraViewfinder(parent widgets.QWidgetITF) *QCameraViewfinder {
	return QCameraViewfinderFromPointer(unsafe.Pointer(C.QCameraViewfinder_NewQCameraViewfinder(C.QtObjectPtr(widgets.PointerFromQWidget(parent)))))
}

func (ptr *QCameraViewfinder) MediaObject() *multimedia.QMediaObject {
	if ptr.Pointer() != nil {
		return multimedia.QMediaObjectFromPointer(unsafe.Pointer(C.QCameraViewfinder_MediaObject(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QCameraViewfinder) DestroyQCameraViewfinder() {
	if ptr.Pointer() != nil {
		C.QCameraViewfinder_DestroyQCameraViewfinder(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
