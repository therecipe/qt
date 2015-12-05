package location

//#include "location.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QGeoRouteReply struct {
	core.QObject
}

type QGeoRouteReply_ITF interface {
	core.QObject_ITF
	QGeoRouteReply_PTR() *QGeoRouteReply
}

func PointerFromQGeoRouteReply(ptr QGeoRouteReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoRouteReply_PTR().Pointer()
	}
	return nil
}

func NewQGeoRouteReplyFromPointer(ptr unsafe.Pointer) *QGeoRouteReply {
	var n = new(QGeoRouteReply)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGeoRouteReply_") {
		n.SetObjectName("QGeoRouteReply_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGeoRouteReply) QGeoRouteReply_PTR() *QGeoRouteReply {
	return ptr
}

//QGeoRouteReply::Error
type QGeoRouteReply__Error int64

const (
	QGeoRouteReply__NoError                = QGeoRouteReply__Error(0)
	QGeoRouteReply__EngineNotSetError      = QGeoRouteReply__Error(1)
	QGeoRouteReply__CommunicationError     = QGeoRouteReply__Error(2)
	QGeoRouteReply__ParseError             = QGeoRouteReply__Error(3)
	QGeoRouteReply__UnsupportedOptionError = QGeoRouteReply__Error(4)
	QGeoRouteReply__UnknownError           = QGeoRouteReply__Error(5)
)

func NewQGeoRouteReply(error QGeoRouteReply__Error, errorString string, parent core.QObject_ITF) *QGeoRouteReply {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRouteReply::QGeoRouteReply")
		}
	}()

	return NewQGeoRouteReplyFromPointer(C.QGeoRouteReply_NewQGeoRouteReply(C.int(error), C.CString(errorString), core.PointerFromQObject(parent)))
}

func (ptr *QGeoRouteReply) Abort() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRouteReply::abort")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoRouteReply_Abort(ptr.Pointer())
	}
}

func (ptr *QGeoRouteReply) Error() QGeoRouteReply__Error {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRouteReply::error")
		}
	}()

	if ptr.Pointer() != nil {
		return QGeoRouteReply__Error(C.QGeoRouteReply_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRouteReply) ErrorString() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRouteReply::errorString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoRouteReply_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoRouteReply) ConnectFinished(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRouteReply::finished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoRouteReply_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QGeoRouteReply) DisconnectFinished() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRouteReply::finished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoRouteReply_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQGeoRouteReplyFinished
func callbackQGeoRouteReplyFinished(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRouteReply::finished")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "finished").(func())()
}

func (ptr *QGeoRouteReply) IsFinished() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRouteReply::isFinished")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QGeoRouteReply_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGeoRouteReply) DestroyQGeoRouteReply() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRouteReply::~QGeoRouteReply")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoRouteReply_DestroyQGeoRouteReply(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
