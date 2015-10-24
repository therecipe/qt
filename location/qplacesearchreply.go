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

type QPlaceSearchReplyITF interface {
	QPlaceReplyITF
	QPlaceSearchReplyPTR() *QPlaceSearchReply
}

func PointerFromQPlaceSearchReply(ptr QPlaceSearchReplyITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceSearchReplyPTR().Pointer()
	}
	return nil
}

func QPlaceSearchReplyFromPointer(ptr unsafe.Pointer) *QPlaceSearchReply {
	var n = new(QPlaceSearchReply)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QPlaceSearchReply_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPlaceSearchReply) QPlaceSearchReplyPTR() *QPlaceSearchReply {
	return ptr
}
