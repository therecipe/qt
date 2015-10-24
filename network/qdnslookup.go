package network

//#include "qdnslookup.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDnsLookup struct {
	core.QObject
}

type QDnsLookupITF interface {
	core.QObjectITF
	QDnsLookupPTR() *QDnsLookup
}

func PointerFromQDnsLookup(ptr QDnsLookupITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsLookupPTR().Pointer()
	}
	return nil
}

func QDnsLookupFromPointer(ptr unsafe.Pointer) *QDnsLookup {
	var n = new(QDnsLookup)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QDnsLookup_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDnsLookup) QDnsLookupPTR() *QDnsLookup {
	return ptr
}

//QDnsLookup::Error
type QDnsLookup__Error int

var (
	QDnsLookup__NoError                 = QDnsLookup__Error(0)
	QDnsLookup__ResolverError           = QDnsLookup__Error(1)
	QDnsLookup__OperationCancelledError = QDnsLookup__Error(2)
	QDnsLookup__InvalidRequestError     = QDnsLookup__Error(3)
	QDnsLookup__InvalidReplyError       = QDnsLookup__Error(4)
	QDnsLookup__ServerFailureError      = QDnsLookup__Error(5)
	QDnsLookup__ServerRefusedError      = QDnsLookup__Error(6)
	QDnsLookup__NotFoundError           = QDnsLookup__Error(7)
)

//QDnsLookup::Type
type QDnsLookup__Type int

var (
	QDnsLookup__A     = QDnsLookup__Type(1)
	QDnsLookup__AAAA  = QDnsLookup__Type(28)
	QDnsLookup__ANY   = QDnsLookup__Type(255)
	QDnsLookup__CNAME = QDnsLookup__Type(5)
	QDnsLookup__MX    = QDnsLookup__Type(15)
	QDnsLookup__NS    = QDnsLookup__Type(2)
	QDnsLookup__PTR   = QDnsLookup__Type(12)
	QDnsLookup__SRV   = QDnsLookup__Type(33)
	QDnsLookup__TXT   = QDnsLookup__Type(16)
)

func NewQDnsLookup3(ty QDnsLookup__Type, name string, nameserver QHostAddressITF, parent core.QObjectITF) *QDnsLookup {
	return QDnsLookupFromPointer(unsafe.Pointer(C.QDnsLookup_NewQDnsLookup3(C.int(ty), C.CString(name), C.QtObjectPtr(PointerFromQHostAddress(nameserver)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QDnsLookup) Error() QDnsLookup__Error {
	if ptr.Pointer() != nil {
		return QDnsLookup__Error(C.QDnsLookup_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDnsLookup) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsLookup_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDnsLookup) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsLookup_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDnsLookup) SetName(name string) {
	if ptr.Pointer() != nil {
		C.QDnsLookup_SetName(C.QtObjectPtr(ptr.Pointer()), C.CString(name))
	}
}

func (ptr *QDnsLookup) SetNameserver(nameserver QHostAddressITF) {
	if ptr.Pointer() != nil {
		C.QDnsLookup_SetNameserver(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQHostAddress(nameserver)))
	}
}

func (ptr *QDnsLookup) SetType(v QDnsLookup__Type) {
	if ptr.Pointer() != nil {
		C.QDnsLookup_SetType(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QDnsLookup) Type() QDnsLookup__Type {
	if ptr.Pointer() != nil {
		return QDnsLookup__Type(C.QDnsLookup_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQDnsLookup(parent core.QObjectITF) *QDnsLookup {
	return QDnsLookupFromPointer(unsafe.Pointer(C.QDnsLookup_NewQDnsLookup(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQDnsLookup2(ty QDnsLookup__Type, name string, parent core.QObjectITF) *QDnsLookup {
	return QDnsLookupFromPointer(unsafe.Pointer(C.QDnsLookup_NewQDnsLookup2(C.int(ty), C.CString(name), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QDnsLookup) Abort() {
	if ptr.Pointer() != nil {
		C.QDnsLookup_Abort(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDnsLookup) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QDnsLookup_ConnectFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QDnsLookup) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QDnsLookup_DisconnectFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQDnsLookupFinished
func callbackQDnsLookupFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "finished").(func())()
}

func (ptr *QDnsLookup) IsFinished() bool {
	if ptr.Pointer() != nil {
		return C.QDnsLookup_IsFinished(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDnsLookup) Lookup() {
	if ptr.Pointer() != nil {
		C.QDnsLookup_Lookup(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDnsLookup) ConnectNameChanged(f func(name string)) {
	if ptr.Pointer() != nil {
		C.QDnsLookup_ConnectNameChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "nameChanged", f)
	}
}

func (ptr *QDnsLookup) DisconnectNameChanged() {
	if ptr.Pointer() != nil {
		C.QDnsLookup_DisconnectNameChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "nameChanged")
	}
}

//export callbackQDnsLookupNameChanged
func callbackQDnsLookupNameChanged(ptrName *C.char, name *C.char) {
	qt.GetSignal(C.GoString(ptrName), "nameChanged").(func(string))(C.GoString(name))
}

func (ptr *QDnsLookup) ConnectTypeChanged(f func(ty QDnsLookup__Type)) {
	if ptr.Pointer() != nil {
		C.QDnsLookup_ConnectTypeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "typeChanged", f)
	}
}

func (ptr *QDnsLookup) DisconnectTypeChanged() {
	if ptr.Pointer() != nil {
		C.QDnsLookup_DisconnectTypeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "typeChanged")
	}
}

//export callbackQDnsLookupTypeChanged
func callbackQDnsLookupTypeChanged(ptrName *C.char, ty C.int) {
	qt.GetSignal(C.GoString(ptrName), "typeChanged").(func(QDnsLookup__Type))(QDnsLookup__Type(ty))
}

func (ptr *QDnsLookup) DestroyQDnsLookup() {
	if ptr.Pointer() != nil {
		C.QDnsLookup_DestroyQDnsLookup(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
