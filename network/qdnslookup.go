package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QDnsLookup struct {
	core.QObject
}

type QDnsLookup_ITF interface {
	core.QObject_ITF
	QDnsLookup_PTR() *QDnsLookup
}

func PointerFromQDnsLookup(ptr QDnsLookup_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsLookup_PTR().Pointer()
	}
	return nil
}

func NewQDnsLookupFromPointer(ptr unsafe.Pointer) *QDnsLookup {
	var n = new(QDnsLookup)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDnsLookup_") {
		n.SetObjectName("QDnsLookup_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDnsLookup) QDnsLookup_PTR() *QDnsLookup {
	return ptr
}

//QDnsLookup::Error
type QDnsLookup__Error int64

const (
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
type QDnsLookup__Type int64

const (
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

func NewQDnsLookup3(ty QDnsLookup__Type, name string, nameserver QHostAddress_ITF, parent core.QObject_ITF) *QDnsLookup {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDnsLookup::QDnsLookup")
		}
	}()

	return NewQDnsLookupFromPointer(C.QDnsLookup_NewQDnsLookup3(C.int(ty), C.CString(name), PointerFromQHostAddress(nameserver), core.PointerFromQObject(parent)))
}

func (ptr *QDnsLookup) Error() QDnsLookup__Error {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDnsLookup::error")
		}
	}()

	if ptr.Pointer() != nil {
		return QDnsLookup__Error(C.QDnsLookup_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDnsLookup) ErrorString() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDnsLookup::errorString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsLookup_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsLookup) Name() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDnsLookup::name")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsLookup_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsLookup) SetName(name string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDnsLookup::setName")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDnsLookup_SetName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QDnsLookup) SetNameserver(nameserver QHostAddress_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDnsLookup::setNameserver")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDnsLookup_SetNameserver(ptr.Pointer(), PointerFromQHostAddress(nameserver))
	}
}

func (ptr *QDnsLookup) SetType(v QDnsLookup__Type) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDnsLookup::setType")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDnsLookup_SetType(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QDnsLookup) Type() QDnsLookup__Type {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDnsLookup::type")
		}
	}()

	if ptr.Pointer() != nil {
		return QDnsLookup__Type(C.QDnsLookup_Type(ptr.Pointer()))
	}
	return 0
}

func NewQDnsLookup(parent core.QObject_ITF) *QDnsLookup {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDnsLookup::QDnsLookup")
		}
	}()

	return NewQDnsLookupFromPointer(C.QDnsLookup_NewQDnsLookup(core.PointerFromQObject(parent)))
}

func NewQDnsLookup2(ty QDnsLookup__Type, name string, parent core.QObject_ITF) *QDnsLookup {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDnsLookup::QDnsLookup")
		}
	}()

	return NewQDnsLookupFromPointer(C.QDnsLookup_NewQDnsLookup2(C.int(ty), C.CString(name), core.PointerFromQObject(parent)))
}

func (ptr *QDnsLookup) Abort() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDnsLookup::abort")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDnsLookup_Abort(ptr.Pointer())
	}
}

func (ptr *QDnsLookup) ConnectFinished(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDnsLookup::finished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDnsLookup_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QDnsLookup) DisconnectFinished() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDnsLookup::finished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDnsLookup_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQDnsLookupFinished
func callbackQDnsLookupFinished(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDnsLookup::finished")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "finished").(func())()
}

func (ptr *QDnsLookup) IsFinished() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDnsLookup::isFinished")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QDnsLookup_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDnsLookup) Lookup() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDnsLookup::lookup")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDnsLookup_Lookup(ptr.Pointer())
	}
}

func (ptr *QDnsLookup) ConnectNameChanged(f func(name string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDnsLookup::nameChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDnsLookup_ConnectNameChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "nameChanged", f)
	}
}

func (ptr *QDnsLookup) DisconnectNameChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDnsLookup::nameChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDnsLookup_DisconnectNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "nameChanged")
	}
}

//export callbackQDnsLookupNameChanged
func callbackQDnsLookupNameChanged(ptrName *C.char, name *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDnsLookup::nameChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "nameChanged").(func(string))(C.GoString(name))
}

func (ptr *QDnsLookup) ConnectTypeChanged(f func(ty QDnsLookup__Type)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDnsLookup::typeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDnsLookup_ConnectTypeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "typeChanged", f)
	}
}

func (ptr *QDnsLookup) DisconnectTypeChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDnsLookup::typeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDnsLookup_DisconnectTypeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "typeChanged")
	}
}

//export callbackQDnsLookupTypeChanged
func callbackQDnsLookupTypeChanged(ptrName *C.char, ty C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDnsLookup::typeChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "typeChanged").(func(QDnsLookup__Type))(QDnsLookup__Type(ty))
}

func (ptr *QDnsLookup) DestroyQDnsLookup() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDnsLookup::~QDnsLookup")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDnsLookup_DestroyQDnsLookup(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
