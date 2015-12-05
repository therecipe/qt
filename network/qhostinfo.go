package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHostInfo::QHostInfo")
		}
	}()

	return NewQHostInfoFromPointer(C.QHostInfo_NewQHostInfo2(PointerFromQHostInfo(other)))
}

func NewQHostInfo(id int) *QHostInfo {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHostInfo::QHostInfo")
		}
	}()

	return NewQHostInfoFromPointer(C.QHostInfo_NewQHostInfo(C.int(id)))
}

func QHostInfo_AbortHostLookup(id int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHostInfo::abortHostLookup")
		}
	}()

	C.QHostInfo_QHostInfo_AbortHostLookup(C.int(id))
}

func (ptr *QHostInfo) Error() QHostInfo__HostInfoError {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHostInfo::error")
		}
	}()

	if ptr.Pointer() != nil {
		return QHostInfo__HostInfoError(C.QHostInfo_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHostInfo) ErrorString() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHostInfo::errorString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QHostInfo_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHostInfo) HostName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHostInfo::hostName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QHostInfo_HostName(ptr.Pointer()))
	}
	return ""
}

func QHostInfo_LookupHost(name string, receiver core.QObject_ITF, member string) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHostInfo::lookupHost")
		}
	}()

	return int(C.QHostInfo_QHostInfo_LookupHost(C.CString(name), core.PointerFromQObject(receiver), C.CString(member)))
}

func (ptr *QHostInfo) LookupId() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHostInfo::lookupId")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QHostInfo_LookupId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHostInfo) SetError(error QHostInfo__HostInfoError) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHostInfo::setError")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHostInfo_SetError(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QHostInfo) SetErrorString(str string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHostInfo::setErrorString")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHostInfo_SetErrorString(ptr.Pointer(), C.CString(str))
	}
}

func (ptr *QHostInfo) SetHostName(hostName string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHostInfo::setHostName")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHostInfo_SetHostName(ptr.Pointer(), C.CString(hostName))
	}
}

func (ptr *QHostInfo) SetLookupId(id int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHostInfo::setLookupId")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHostInfo_SetLookupId(ptr.Pointer(), C.int(id))
	}
}

func (ptr *QHostInfo) DestroyQHostInfo() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHostInfo::~QHostInfo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHostInfo_DestroyQHostInfo(ptr.Pointer())
	}
}

func QHostInfo_LocalHostName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHostInfo::localHostName")
		}
	}()

	return C.GoString(C.QHostInfo_QHostInfo_LocalHostName())
}

func QHostInfo_LocalDomainName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHostInfo::localDomainName")
		}
	}()

	return C.GoString(C.QHostInfo_QHostInfo_LocalDomainName())
}
