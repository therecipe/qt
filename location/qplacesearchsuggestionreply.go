package location

//#include "qplacesearchsuggestionreply.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QPlaceSearchSuggestionReply struct {
	QPlaceReply
}

type QPlaceSearchSuggestionReplyITF interface {
	QPlaceReplyITF
	QPlaceSearchSuggestionReplyPTR() *QPlaceSearchSuggestionReply
}

func PointerFromQPlaceSearchSuggestionReply(ptr QPlaceSearchSuggestionReplyITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceSearchSuggestionReplyPTR().Pointer()
	}
	return nil
}

func QPlaceSearchSuggestionReplyFromPointer(ptr unsafe.Pointer) *QPlaceSearchSuggestionReply {
	var n = new(QPlaceSearchSuggestionReply)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QPlaceSearchSuggestionReply_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPlaceSearchSuggestionReply) QPlaceSearchSuggestionReplyPTR() *QPlaceSearchSuggestionReply {
	return ptr
}
