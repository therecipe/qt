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

type QHostInfo_ITF interface {
	QHostInfo_PTR() *QHostInfo
}

func (p *QHostInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QHostInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQHostInfo(ptr QHostInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHostInfo_PTR().Pointer()
	}
	return nil
}

func NewQHostInfoFromPointer(ptr unsafe.Pointer) *QHostInfo {
	var n = new(QHostInfo)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHostInfo) QHostInfo_PTR() *QHostInfo {
	return ptr
}

//QHostInfo::HostInfoError
type QHostInfo__HostInfoError int64

const (
	QHostInfo__NoError      = QHostInfo__HostInfoError(0)
	QHostInfo__HostNotFound = QHostInfo__HostInfoError(1)
	QHostInfo__UnknownError = QHostInfo__HostInfoError(2)
)

func NewQHostInfo2(other QHostInfo_ITF) *QHostInfo {
	return NewQHostInfoFromPointer(C.QHostInfo_NewQHostInfo2(PointerFromQHostInfo(other)))
}

func NewQHostInfo(id int) *QHostInfo {
	return NewQHostInfoFromPointer(C.QHostInfo_NewQHostInfo(C.int(id)))
}

func QHostInfo_AbortHostLookup(id int) {
	C.QHostInfo_QHostInfo_AbortHostLookup(C.int(id))
}

func (ptr *QHostInfo) Error() QHostInfo__HostInfoError {
	if ptr.Pointer() != nil {
		return QHostInfo__HostInfoError(C.QHostInfo_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHostInfo) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QHostInfo_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHostInfo) HostName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QHostInfo_HostName(ptr.Pointer()))
	}
	return ""
}

func QHostInfo_LookupHost(name string, receiver core.QObject_ITF, member string) int {
	return int(C.QHostInfo_QHostInfo_LookupHost(C.CString(name), core.PointerFromQObject(receiver), C.CString(member)))
}

func (ptr *QHostInfo) LookupId() int {
	if ptr.Pointer() != nil {
		return int(C.QHostInfo_LookupId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHostInfo) SetError(error QHostInfo__HostInfoError) {
	if ptr.Pointer() != nil {
		C.QHostInfo_SetError(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QHostInfo) SetErrorString(str string) {
	if ptr.Pointer() != nil {
		C.QHostInfo_SetErrorString(ptr.Pointer(), C.CString(str))
	}
}

func (ptr *QHostInfo) SetHostName(hostName string) {
	if ptr.Pointer() != nil {
		C.QHostInfo_SetHostName(ptr.Pointer(), C.CString(hostName))
	}
}

func (ptr *QHostInfo) SetLookupId(id int) {
	if ptr.Pointer() != nil {
		C.QHostInfo_SetLookupId(ptr.Pointer(), C.int(id))
	}
}

func (ptr *QHostInfo) DestroyQHostInfo() {
	if ptr.Pointer() != nil {
		C.QHostInfo_DestroyQHostInfo(ptr.Pointer())
	}
}

func QHostInfo_LocalHostName() string {
	return C.GoString(C.QHostInfo_QHostInfo_LocalHostName())
}

func QHostInfo_LocalDomainName() string {
	return C.GoString(C.QHostInfo_QHostInfo_LocalDomainName())
}
