package location

//#include "location.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QPlaceDetailsReply struct {
	QPlaceReply
}

type QPlaceDetailsReply_ITF interface {
	QPlaceReply_ITF
	QPlaceDetailsReply_PTR() *QPlaceDetailsReply
}

func PointerFromQPlaceDetailsReply(ptr QPlaceDetailsReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceDetailsReply_PTR().Pointer()
	}
	return nil
}

func NewQPlaceDetailsReplyFromPointer(ptr unsafe.Pointer) *QPlaceDetailsReply {
	var n = new(QPlaceDetailsReply)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QPlaceDetailsReply_") {
		n.SetObjectName("QPlaceDetailsReply_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPlaceDetailsReply) QPlaceDetailsReply_PTR() *QPlaceDetailsReply {
	return ptr
}
