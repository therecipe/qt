package location

//#include "qplacedetailsreply.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QPlaceDetailsReply struct {
	QPlaceReply
}

type QPlaceDetailsReplyITF interface {
	QPlaceReplyITF
	QPlaceDetailsReplyPTR() *QPlaceDetailsReply
}

func PointerFromQPlaceDetailsReply(ptr QPlaceDetailsReplyITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceDetailsReplyPTR().Pointer()
	}
	return nil
}

func QPlaceDetailsReplyFromPointer(ptr unsafe.Pointer) *QPlaceDetailsReply {
	var n = new(QPlaceDetailsReply)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QPlaceDetailsReply_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPlaceDetailsReply) QPlaceDetailsReplyPTR() *QPlaceDetailsReply {
	return ptr
}
