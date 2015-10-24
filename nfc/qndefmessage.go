package nfc

//#include "qndefmessage.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QNdefMessage struct {
	core.QList
}

type QNdefMessageITF interface {
	core.QListITF
	QNdefMessagePTR() *QNdefMessage
}

func PointerFromQNdefMessage(ptr QNdefMessageITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefMessagePTR().Pointer()
	}
	return nil
}

func QNdefMessageFromPointer(ptr unsafe.Pointer) *QNdefMessage {
	var n = new(QNdefMessage)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNdefMessage) QNdefMessagePTR() *QNdefMessage {
	return ptr
}

func NewQNdefMessage() *QNdefMessage {
	return QNdefMessageFromPointer(unsafe.Pointer(C.QNdefMessage_NewQNdefMessage()))
}

func NewQNdefMessage3(message QNdefMessageITF) *QNdefMessage {
	return QNdefMessageFromPointer(unsafe.Pointer(C.QNdefMessage_NewQNdefMessage3(C.QtObjectPtr(PointerFromQNdefMessage(message)))))
}

func NewQNdefMessage2(record QNdefRecordITF) *QNdefMessage {
	return QNdefMessageFromPointer(unsafe.Pointer(C.QNdefMessage_NewQNdefMessage2(C.QtObjectPtr(PointerFromQNdefRecord(record)))))
}
