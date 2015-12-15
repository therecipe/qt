package location

//#include "location.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPlaceReply struct {
	core.QObject
}

type QPlaceReply_ITF interface {
	core.QObject_ITF
	QPlaceReply_PTR() *QPlaceReply
}

func PointerFromQPlaceReply(ptr QPlaceReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceReply_PTR().Pointer()
	}
	return nil
}

func NewQPlaceReplyFromPointer(ptr unsafe.Pointer) *QPlaceReply {
	var n = new(QPlaceReply)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QPlaceReply_") {
		n.SetObjectName("QPlaceReply_" + qt.Identifier())
	}
	return n
}

func (ptr *QPlaceReply) QPlaceReply_PTR() *QPlaceReply {
	return ptr
}

//QPlaceReply::Error
type QPlaceReply__Error int64

const (
	QPlaceReply__NoError                   = QPlaceReply__Error(0)
	QPlaceReply__PlaceDoesNotExistError    = QPlaceReply__Error(1)
	QPlaceReply__CategoryDoesNotExistError = QPlaceReply__Error(2)
	QPlaceReply__CommunicationError        = QPlaceReply__Error(3)
	QPlaceReply__ParseError                = QPlaceReply__Error(4)
	QPlaceReply__PermissionsError          = QPlaceReply__Error(5)
	QPlaceReply__UnsupportedError          = QPlaceReply__Error(6)
	QPlaceReply__BadArgumentError          = QPlaceReply__Error(7)
	QPlaceReply__CancelError               = QPlaceReply__Error(8)
	QPlaceReply__UnknownError              = QPlaceReply__Error(9)
)

//QPlaceReply::Type
type QPlaceReply__Type int64

const (
	QPlaceReply__Reply                 = QPlaceReply__Type(0)
	QPlaceReply__DetailsReply          = QPlaceReply__Type(1)
	QPlaceReply__SearchReply           = QPlaceReply__Type(2)
	QPlaceReply__SearchSuggestionReply = QPlaceReply__Type(3)
	QPlaceReply__ContentReply          = QPlaceReply__Type(4)
	QPlaceReply__IdReply               = QPlaceReply__Type(5)
	QPlaceReply__MatchReply            = QPlaceReply__Type(6)
)
