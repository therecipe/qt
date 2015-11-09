package location

//#include "qplacesearchreply.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QPlaceSearchReply struct {
	QPlaceReply
}

type QPlaceSearchReply_ITF interface {
	QPlaceReply_ITF
	QPlaceSearchReply_PTR() *QPlaceSearchReply
}

func PointerFromQPlaceSearchReply(ptr QPlaceSearchReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceSearchReply_PTR().Pointer()
	}
	return nil
}

func NewQPlaceSearchReplyFromPointer(ptr unsafe.Pointer) *QPlaceSearchReply {
	var n = new(QPlaceSearchReply)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QPlaceSearchReply_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPlaceSearchReply) QPlaceSearchReply_PTR() *QPlaceSearchReply {
	return ptr
}
