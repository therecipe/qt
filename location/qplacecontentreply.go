package location

//#include "qplacecontentreply.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QPlaceContentReply struct {
	QPlaceReply
}

type QPlaceContentReplyITF interface {
	QPlaceReplyITF
	QPlaceContentReplyPTR() *QPlaceContentReply
}

func PointerFromQPlaceContentReply(ptr QPlaceContentReplyITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceContentReplyPTR().Pointer()
	}
	return nil
}

func QPlaceContentReplyFromPointer(ptr unsafe.Pointer) *QPlaceContentReply {
	var n = new(QPlaceContentReply)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QPlaceContentReply_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPlaceContentReply) QPlaceContentReplyPTR() *QPlaceContentReply {
	return ptr
}
