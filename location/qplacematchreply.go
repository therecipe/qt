package location

//#include "qplacematchreply.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QPlaceMatchReply struct {
	QPlaceReply
}

type QPlaceMatchReplyITF interface {
	QPlaceReplyITF
	QPlaceMatchReplyPTR() *QPlaceMatchReply
}

func PointerFromQPlaceMatchReply(ptr QPlaceMatchReplyITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceMatchReplyPTR().Pointer()
	}
	return nil
}

func QPlaceMatchReplyFromPointer(ptr unsafe.Pointer) *QPlaceMatchReply {
	var n = new(QPlaceMatchReply)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QPlaceMatchReply_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPlaceMatchReply) QPlaceMatchReplyPTR() *QPlaceMatchReply {
	return ptr
}
