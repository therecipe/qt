package location

//#include "location.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QPlaceMatchReply struct {
	QPlaceReply
}

type QPlaceMatchReply_ITF interface {
	QPlaceReply_ITF
	QPlaceMatchReply_PTR() *QPlaceMatchReply
}

func PointerFromQPlaceMatchReply(ptr QPlaceMatchReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceMatchReply_PTR().Pointer()
	}
	return nil
}

func NewQPlaceMatchReplyFromPointer(ptr unsafe.Pointer) *QPlaceMatchReply {
	var n = new(QPlaceMatchReply)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QPlaceMatchReply_") {
		n.SetObjectName("QPlaceMatchReply_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPlaceMatchReply) QPlaceMatchReply_PTR() *QPlaceMatchReply {
	return ptr
}
