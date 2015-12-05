package dbus

//#include "dbus.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QDBusMessage struct {
	ptr unsafe.Pointer
}

type QDBusMessage_ITF interface {
	QDBusMessage_PTR() *QDBusMessage
}

func (p *QDBusMessage) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusMessage) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusMessage(ptr QDBusMessage_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusMessage_PTR().Pointer()
	}
	return nil
}

func NewQDBusMessageFromPointer(ptr unsafe.Pointer) *QDBusMessage {
	var n = new(QDBusMessage)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDBusMessage) QDBusMessage_PTR() *QDBusMessage {
	return ptr
}

//QDBusMessage::MessageType
type QDBusMessage__MessageType int64

const (
	QDBusMessage__InvalidMessage    = QDBusMessage__MessageType(0)
	QDBusMessage__MethodCallMessage = QDBusMessage__MessageType(1)
	QDBusMessage__ReplyMessage      = QDBusMessage__MessageType(2)
	QDBusMessage__ErrorMessage      = QDBusMessage__MessageType(3)
	QDBusMessage__SignalMessage     = QDBusMessage__MessageType(4)
)

func NewQDBusMessage() *QDBusMessage {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusMessage::QDBusMessage")
		}
	}()

	return NewQDBusMessageFromPointer(C.QDBusMessage_NewQDBusMessage())
}

func NewQDBusMessage2(other QDBusMessage_ITF) *QDBusMessage {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusMessage::QDBusMessage")
		}
	}()

	return NewQDBusMessageFromPointer(C.QDBusMessage_NewQDBusMessage2(PointerFromQDBusMessage(other)))
}

func (ptr *QDBusMessage) AutoStartService() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusMessage::autoStartService")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QDBusMessage_AutoStartService(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusMessage) ErrorMessage() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusMessage::errorMessage")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusMessage_ErrorMessage(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) ErrorName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusMessage::errorName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusMessage_ErrorName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) Interface() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusMessage::interface")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusMessage_Interface(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) IsDelayedReply() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusMessage::isDelayedReply")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QDBusMessage_IsDelayedReply(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusMessage) IsReplyRequired() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusMessage::isReplyRequired")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QDBusMessage_IsReplyRequired(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusMessage) Member() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusMessage::member")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusMessage_Member(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) Path() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusMessage::path")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusMessage_Path(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) Service() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusMessage::service")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusMessage_Service(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) SetAutoStartService(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusMessage::setAutoStartService")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDBusMessage_SetAutoStartService(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QDBusMessage) SetDelayedReply(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusMessage::setDelayedReply")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDBusMessage_SetDelayedReply(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QDBusMessage) Signature() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusMessage::signature")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusMessage_Signature(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) Type() QDBusMessage__MessageType {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusMessage::type")
		}
	}()

	if ptr.Pointer() != nil {
		return QDBusMessage__MessageType(C.QDBusMessage_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDBusMessage) DestroyQDBusMessage() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusMessage::~QDBusMessage")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDBusMessage_DestroyQDBusMessage(ptr.Pointer())
	}
}
