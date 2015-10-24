package location

//#include "qplaceidreply.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QPlaceIdReply struct {
	QPlaceReply
}

type QPlaceIdReplyITF interface {
	QPlaceReplyITF
	QPlaceIdReplyPTR() *QPlaceIdReply
}

func PointerFromQPlaceIdReply(ptr QPlaceIdReplyITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceIdReplyPTR().Pointer()
	}
	return nil
}

func QPlaceIdReplyFromPointer(ptr unsafe.Pointer) *QPlaceIdReply {
	var n = new(QPlaceIdReply)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QPlaceIdReply_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPlaceIdReply) QPlaceIdReplyPTR() *QPlaceIdReply {
	return ptr
}

//QPlaceIdReply::OperationType
type QPlaceIdReply__OperationType int

var (
	QPlaceIdReply__SavePlace      = QPlaceIdReply__OperationType(0)
	QPlaceIdReply__SaveCategory   = QPlaceIdReply__OperationType(1)
	QPlaceIdReply__RemovePlace    = QPlaceIdReply__OperationType(2)
	QPlaceIdReply__RemoveCategory = QPlaceIdReply__OperationType(3)
)
