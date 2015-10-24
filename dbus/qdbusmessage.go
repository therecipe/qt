package dbus

//#include "qdbusmessage.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QDBusMessage struct {
	ptr unsafe.Pointer
}

type QDBusMessageITF interface {
	QDBusMessagePTR() *QDBusMessage
}

func (p *QDBusMessage) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusMessage) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusMessage(ptr QDBusMessageITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusMessagePTR().Pointer()
	}
	return nil
}

func QDBusMessageFromPointer(ptr unsafe.Pointer) *QDBusMessage {
	var n = new(QDBusMessage)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDBusMessage) QDBusMessagePTR() *QDBusMessage {
	return ptr
}

//QDBusMessage::MessageType
type QDBusMessage__MessageType int

var (
	QDBusMessage__InvalidMessage    = QDBusMessage__MessageType(0)
	QDBusMessage__MethodCallMessage = QDBusMessage__MessageType(1)
	QDBusMessage__ReplyMessage      = QDBusMessage__MessageType(2)
	QDBusMessage__ErrorMessage      = QDBusMessage__MessageType(3)
	QDBusMessage__SignalMessage     = QDBusMessage__MessageType(4)
)

func NewQDBusMessage() *QDBusMessage {
	return QDBusMessageFromPointer(unsafe.Pointer(C.QDBusMessage_NewQDBusMessage()))
}

func NewQDBusMessage2(other QDBusMessageITF) *QDBusMessage {
	return QDBusMessageFromPointer(unsafe.Pointer(C.QDBusMessage_NewQDBusMessage2(C.QtObjectPtr(PointerFromQDBusMessage(other)))))
}

func (ptr *QDBusMessage) AutoStartService() bool {
	if ptr.Pointer() != nil {
		return C.QDBusMessage_AutoStartService(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDBusMessage) ErrorMessage() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusMessage_ErrorMessage(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDBusMessage) ErrorName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusMessage_ErrorName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDBusMessage) Interface() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusMessage_Interface(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDBusMessage) IsDelayedReply() bool {
	if ptr.Pointer() != nil {
		return C.QDBusMessage_IsDelayedReply(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDBusMessage) IsReplyRequired() bool {
	if ptr.Pointer() != nil {
		return C.QDBusMessage_IsReplyRequired(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDBusMessage) Member() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusMessage_Member(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDBusMessage) Path() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusMessage_Path(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDBusMessage) Service() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusMessage_Service(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDBusMessage) SetAutoStartService(enable bool) {
	if ptr.Pointer() != nil {
		C.QDBusMessage_SetAutoStartService(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QDBusMessage) SetDelayedReply(enable bool) {
	if ptr.Pointer() != nil {
		C.QDBusMessage_SetDelayedReply(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QDBusMessage) Signature() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusMessage_Signature(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDBusMessage) Type() QDBusMessage__MessageType {
	if ptr.Pointer() != nil {
		return QDBusMessage__MessageType(C.QDBusMessage_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDBusMessage) DestroyQDBusMessage() {
	if ptr.Pointer() != nil {
		C.QDBusMessage_DestroyQDBusMessage(C.QtObjectPtr(ptr.Pointer()))
	}
}
