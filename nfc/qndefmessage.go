package nfc

//#include "nfc.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QNdefMessage struct {
	core.QList
}

type QNdefMessage_ITF interface {
	core.QList_ITF
	QNdefMessage_PTR() *QNdefMessage
}

func PointerFromQNdefMessage(ptr QNdefMessage_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefMessage_PTR().Pointer()
	}
	return nil
}

func NewQNdefMessageFromPointer(ptr unsafe.Pointer) *QNdefMessage {
	var n = new(QNdefMessage)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNdefMessage) QNdefMessage_PTR() *QNdefMessage {
	return ptr
}

func NewQNdefMessage() *QNdefMessage {
	defer qt.Recovering("QNdefMessage::QNdefMessage")

	return NewQNdefMessageFromPointer(C.QNdefMessage_NewQNdefMessage())
}

func NewQNdefMessage3(message QNdefMessage_ITF) *QNdefMessage {
	defer qt.Recovering("QNdefMessage::QNdefMessage")

	return NewQNdefMessageFromPointer(C.QNdefMessage_NewQNdefMessage3(PointerFromQNdefMessage(message)))
}

func NewQNdefMessage2(record QNdefRecord_ITF) *QNdefMessage {
	defer qt.Recovering("QNdefMessage::QNdefMessage")

	return NewQNdefMessageFromPointer(C.QNdefMessage_NewQNdefMessage2(PointerFromQNdefRecord(record)))
}

func (ptr *QNdefMessage) ToByteArray() *core.QByteArray {
	defer qt.Recovering("QNdefMessage::toByteArray")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QNdefMessage_ToByteArray(ptr.Pointer()))
	}
	return nil
}
