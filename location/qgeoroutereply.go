package location

//#include "qgeoroutereply.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
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
	if n.ObjectName() == "" {
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
	return NewQGeoRouteReplyFromPointer(C.QGeoRouteReply_NewQGeoRouteReply(C.int(error), C.CString(errorString), core.PointerFromQObject(parent)))
}

func (ptr *QGeoRouteReply) Abort() {
	if ptr.Pointer() != nil {
		C.QGeoRouteReply_Abort(ptr.Pointer())
	}
}

func (ptr *QGeoRouteReply) Error() QGeoRouteReply__Error {
	if ptr.Pointer() != nil {
		return QGeoRouteReply__Error(C.QGeoRouteReply_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRouteReply) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoRouteReply_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoRouteReply) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QGeoRouteReply_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QGeoRouteReply) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QGeoRouteReply_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQGeoRouteReplyFinished
func callbackQGeoRouteReplyFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "finished").(func())()
}

func (ptr *QGeoRouteReply) IsFinished() bool {
	if ptr.Pointer() != nil {
		return C.QGeoRouteReply_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGeoRouteReply) DestroyQGeoRouteReply() {
	if ptr.Pointer() != nil {
		C.QGeoRouteReply_DestroyQGeoRouteReply(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
