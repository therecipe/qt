package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QMediaService struct {
	core.QObject
}

type QMediaService_ITF interface {
	core.QObject_ITF
	QMediaService_PTR() *QMediaService
}

func PointerFromQMediaService(ptr QMediaService_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaService_PTR().Pointer()
	}
	return nil
}

func NewQMediaServiceFromPointer(ptr unsafe.Pointer) *QMediaService {
	var n = new(QMediaService)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMediaService_") {
		n.SetObjectName("QMediaService_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaService) QMediaService_PTR() *QMediaService {
	return ptr
}

func (ptr *QMediaService) ReleaseControl(control QMediaControl_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaService::releaseControl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaService_ReleaseControl(ptr.Pointer(), PointerFromQMediaControl(control))
	}
}

func (ptr *QMediaService) RequestControl(interfa string) *QMediaControl {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaService::requestControl")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQMediaControlFromPointer(C.QMediaService_RequestControl(ptr.Pointer(), C.CString(interfa)))
	}
	return nil
}

func (ptr *QMediaService) RequestControl2() unsafe.Pointer {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaService::requestControl")
		}
	}()

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QMediaService_RequestControl2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaService) DestroyQMediaService() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaService::~QMediaService")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaService_DestroyQMediaService(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
