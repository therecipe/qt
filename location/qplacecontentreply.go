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

type QPlaceContentReply_ITF interface {
	QPlaceReply_ITF
	QPlaceContentReply_PTR() *QPlaceContentReply
}

func PointerFromQPlaceContentReply(ptr QPlaceContentReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceContentReply_PTR().Pointer()
	}
	return nil
}

func NewQPlaceContentReplyFromPointer(ptr unsafe.Pointer) *QPlaceContentReply {
	var n = new(QPlaceContentReply)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QPlaceContentReply_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPlaceContentReply) QPlaceContentReply_PTR() *QPlaceContentReply {
	return ptr
}
