package location

//#include "location.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QPlaceIdReply struct {
	QPlaceReply
}

type QPlaceIdReply_ITF interface {
	QPlaceReply_ITF
	QPlaceIdReply_PTR() *QPlaceIdReply
}

func PointerFromQPlaceIdReply(ptr QPlaceIdReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceIdReply_PTR().Pointer()
	}
	return nil
}

func NewQPlaceIdReplyFromPointer(ptr unsafe.Pointer) *QPlaceIdReply {
	var n = new(QPlaceIdReply)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QPlaceIdReply_") {
		n.SetObjectName("QPlaceIdReply_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPlaceIdReply) QPlaceIdReply_PTR() *QPlaceIdReply {
	return ptr
}

//QPlaceIdReply::OperationType
type QPlaceIdReply__OperationType int64

const (
	QPlaceIdReply__SavePlace      = QPlaceIdReply__OperationType(0)
	QPlaceIdReply__SaveCategory   = QPlaceIdReply__OperationType(1)
	QPlaceIdReply__RemovePlace    = QPlaceIdReply__OperationType(2)
	QPlaceIdReply__RemoveCategory = QPlaceIdReply__OperationType(3)
)
