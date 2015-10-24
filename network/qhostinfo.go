package network

//#include "qhostinfo.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QHostInfo struct {
	ptr unsafe.Pointer
}

type QHostInfoITF interface {
	QHostInfoPTR() *QHostInfo
}

func (p *QHostInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QHostInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQHostInfo(ptr QHostInfoITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHostInfoPTR().Pointer()
	}
	return nil
}

func QHostInfoFromPointer(ptr unsafe.Pointer) *QHostInfo {
	var n = new(QHostInfo)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHostInfo) QHostInfoPTR() *QHostInfo {
	return ptr
}

//QHostInfo::HostInfoError
type QHostInfo__HostInfoError int

var (
	QHostInfo__NoError      = QHostInfo__HostInfoError(0)
	QHostInfo__HostNotFound = QHostInfo__HostInfoError(1)
	QHostInfo__UnknownError = QHostInfo__HostInfoError(2)
)

func NewQHostInfo2(other QHostInfoITF) *QHostInfo {
	return QHostInfoFromPointer(unsafe.Pointer(C.QHostInfo_NewQHostInfo2(C.QtObjectPtr(PointerFromQHostInfo(other)))))
}

func NewQHostInfo(id int) *QHostInfo {
	return QHostInfoFromPointer(unsafe.Pointer(C.QHostInfo_NewQHostInfo(C.int(id))))
}

func QHostInfo_AbortHostLookup(id int) {
	C.QHostInfo_QHostInfo_AbortHostLookup(C.int(id))
}

func (ptr *QHostInfo) Error() QHostInfo__HostInfoError {
	if ptr.Pointer() != nil {
		return QHostInfo__HostInfoError(C.QHostInfo_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHostInfo) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QHostInfo_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QHostInfo) HostName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QHostInfo_HostName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func QHostInfo_LookupHost(name string, receiver core.QObjectITF, member string) int {
	return int(C.QHostInfo_QHostInfo_LookupHost(C.CString(name), C.QtObjectPtr(core.PointerFromQObject(receiver)), C.CString(member)))
}

func (ptr *QHostInfo) LookupId() int {
	if ptr.Pointer() != nil {
		return int(C.QHostInfo_LookupId(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHostInfo) SetError(error QHostInfo__HostInfoError) {
	if ptr.Pointer() != nil {
		C.QHostInfo_SetError(C.QtObjectPtr(ptr.Pointer()), C.int(error))
	}
}

func (ptr *QHostInfo) SetErrorString(str string) {
	if ptr.Pointer() != nil {
		C.QHostInfo_SetErrorString(C.QtObjectPtr(ptr.Pointer()), C.CString(str))
	}
}

func (ptr *QHostInfo) SetHostName(hostName string) {
	if ptr.Pointer() != nil {
		C.QHostInfo_SetHostName(C.QtObjectPtr(ptr.Pointer()), C.CString(hostName))
	}
}

func (ptr *QHostInfo) SetLookupId(id int) {
	if ptr.Pointer() != nil {
		C.QHostInfo_SetLookupId(C.QtObjectPtr(ptr.Pointer()), C.int(id))
	}
}

func (ptr *QHostInfo) DestroyQHostInfo() {
	if ptr.Pointer() != nil {
		C.QHostInfo_DestroyQHostInfo(C.QtObjectPtr(ptr.Pointer()))
	}
}

func QHostInfo_LocalHostName() string {
	return C.GoString(C.QHostInfo_QHostInfo_LocalHostName())
}

func QHostInfo_LocalDomainName() string {
	return C.GoString(C.QHostInfo_QHostInfo_LocalDomainName())
}
