package dbus

//#include "qdbuserror.h"
import "C"
import (
	"unsafe"
)

type QDBusError struct {
	ptr unsafe.Pointer
}

type QDBusErrorITF interface {
	QDBusErrorPTR() *QDBusError
}

func (p *QDBusError) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusError) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusError(ptr QDBusErrorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusErrorPTR().Pointer()
	}
	return nil
}

func QDBusErrorFromPointer(ptr unsafe.Pointer) *QDBusError {
	var n = new(QDBusError)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDBusError) QDBusErrorPTR() *QDBusError {
	return ptr
}

//QDBusError::ErrorType
type QDBusError__ErrorType int

var (
	QDBusError__NoError           = QDBusError__ErrorType(0)
	QDBusError__Other             = QDBusError__ErrorType(1)
	QDBusError__Failed            = QDBusError__ErrorType(2)
	QDBusError__NoMemory          = QDBusError__ErrorType(3)
	QDBusError__ServiceUnknown    = QDBusError__ErrorType(4)
	QDBusError__NoReply           = QDBusError__ErrorType(5)
	QDBusError__BadAddress        = QDBusError__ErrorType(6)
	QDBusError__NotSupported      = QDBusError__ErrorType(7)
	QDBusError__LimitsExceeded    = QDBusError__ErrorType(8)
	QDBusError__AccessDenied      = QDBusError__ErrorType(9)
	QDBusError__NoServer          = QDBusError__ErrorType(10)
	QDBusError__Timeout           = QDBusError__ErrorType(11)
	QDBusError__NoNetwork         = QDBusError__ErrorType(12)
	QDBusError__AddressInUse      = QDBusError__ErrorType(13)
	QDBusError__Disconnected      = QDBusError__ErrorType(14)
	QDBusError__InvalidArgs       = QDBusError__ErrorType(15)
	QDBusError__UnknownMethod     = QDBusError__ErrorType(16)
	QDBusError__TimedOut          = QDBusError__ErrorType(17)
	QDBusError__InvalidSignature  = QDBusError__ErrorType(18)
	QDBusError__UnknownInterface  = QDBusError__ErrorType(19)
	QDBusError__UnknownObject     = QDBusError__ErrorType(20)
	QDBusError__UnknownProperty   = QDBusError__ErrorType(21)
	QDBusError__PropertyReadOnly  = QDBusError__ErrorType(22)
	QDBusError__InternalError     = QDBusError__ErrorType(23)
	QDBusError__InvalidService    = QDBusError__ErrorType(24)
	QDBusError__InvalidObjectPath = QDBusError__ErrorType(25)
	QDBusError__InvalidInterface  = QDBusError__ErrorType(26)
	QDBusError__InvalidMember     = QDBusError__ErrorType(27)
)

func QDBusError_ErrorString(error QDBusError__ErrorType) string {
	return C.GoString(C.QDBusError_QDBusError_ErrorString(C.int(error)))
}

func (ptr *QDBusError) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QDBusError_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDBusError) Message() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusError_Message(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDBusError) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusError_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDBusError) Type() QDBusError__ErrorType {
	if ptr.Pointer() != nil {
		return QDBusError__ErrorType(C.QDBusError_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
