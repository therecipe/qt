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

type QPlaceSearchSuggestionReply_ITF interface {
	QPlaceReply_ITF
	QPlaceSearchSuggestionReply_PTR() *QPlaceSearchSuggestionReply
}

func PointerFromQPlaceSearchSuggestionReply(ptr QPlaceSearchSuggestionReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceSearchSuggestionReply_PTR().Pointer()
	}
	return nil
}

func NewQPlaceSearchSuggestionReplyFromPointer(ptr unsafe.Pointer) *QPlaceSearchSuggestionReply {
	var n = new(QPlaceSearchSuggestionReply)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QPlaceSearchSuggestionReply_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPlaceSearchSuggestionReply) QPlaceSearchSuggestionReply_PTR() *QPlaceSearchSuggestionReply {
	return ptr
}
