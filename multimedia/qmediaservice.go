package multimedia

//#include "qmediaservice.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMediaService struct {
	core.QObject
}

type QMediaServiceITF interface {
	core.QObjectITF
	QMediaServicePTR() *QMediaService
}

func PointerFromQMediaService(ptr QMediaServiceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaServicePTR().Pointer()
	}
	return nil
}

func QMediaServiceFromPointer(ptr unsafe.Pointer) *QMediaService {
	var n = new(QMediaService)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMediaService_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaService) QMediaServicePTR() *QMediaService {
	return ptr
}

func (ptr *QMediaService) ReleaseControl(control QMediaControlITF) {
	if ptr.Pointer() != nil {
		C.QMediaService_ReleaseControl(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMediaControl(control)))
	}
}

func (ptr *QMediaService) RequestControl(interfa string) *QMediaControl {
	if ptr.Pointer() != nil {
		return QMediaControlFromPointer(unsafe.Pointer(C.QMediaService_RequestControl(C.QtObjectPtr(ptr.Pointer()), C.CString(interfa))))
	}
	return nil
}

func (ptr *QMediaService) DestroyQMediaService() {
	if ptr.Pointer() != nil {
		C.QMediaService_DestroyQMediaService(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
