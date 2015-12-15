package location

//#include "location.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGeoCodeReply struct {
	core.QObject
}

type QGeoCodeReply_ITF interface {
	core.QObject_ITF
	QGeoCodeReply_PTR() *QGeoCodeReply
}

func PointerFromQGeoCodeReply(ptr QGeoCodeReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoCodeReply_PTR().Pointer()
	}
	return nil
}

func NewQGeoCodeReplyFromPointer(ptr unsafe.Pointer) *QGeoCodeReply {
	var n = new(QGeoCodeReply)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGeoCodeReply_") {
		n.SetObjectName("QGeoCodeReply_" + qt.Identifier())
	}
	return n
}

func (ptr *QGeoCodeReply) QGeoCodeReply_PTR() *QGeoCodeReply {
	return ptr
}

//QGeoCodeReply::Error
type QGeoCodeReply__Error int64

const (
	QGeoCodeReply__NoError                = QGeoCodeReply__Error(0)
	QGeoCodeReply__EngineNotSetError      = QGeoCodeReply__Error(1)
	QGeoCodeReply__CommunicationError     = QGeoCodeReply__Error(2)
	QGeoCodeReply__ParseError             = QGeoCodeReply__Error(3)
	QGeoCodeReply__UnsupportedOptionError = QGeoCodeReply__Error(4)
	QGeoCodeReply__CombinationError       = QGeoCodeReply__Error(5)
	QGeoCodeReply__UnknownError           = QGeoCodeReply__Error(6)
)
