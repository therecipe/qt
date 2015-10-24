package location

//#include "qgeocodereply.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGeoCodeReply struct {
	core.QObject
}

type QGeoCodeReplyITF interface {
	core.QObjectITF
	QGeoCodeReplyPTR() *QGeoCodeReply
}

func PointerFromQGeoCodeReply(ptr QGeoCodeReplyITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoCodeReplyPTR().Pointer()
	}
	return nil
}

func QGeoCodeReplyFromPointer(ptr unsafe.Pointer) *QGeoCodeReply {
	var n = new(QGeoCodeReply)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGeoCodeReply_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGeoCodeReply) QGeoCodeReplyPTR() *QGeoCodeReply {
	return ptr
}

//QGeoCodeReply::Error
type QGeoCodeReply__Error int

var (
	QGeoCodeReply__NoError                = QGeoCodeReply__Error(0)
	QGeoCodeReply__EngineNotSetError      = QGeoCodeReply__Error(1)
	QGeoCodeReply__CommunicationError     = QGeoCodeReply__Error(2)
	QGeoCodeReply__ParseError             = QGeoCodeReply__Error(3)
	QGeoCodeReply__UnsupportedOptionError = QGeoCodeReply__Error(4)
	QGeoCodeReply__CombinationError       = QGeoCodeReply__Error(5)
	QGeoCodeReply__UnknownError           = QGeoCodeReply__Error(6)
)
