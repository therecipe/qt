package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QAuthenticator struct {
	ptr unsafe.Pointer
}

type QAuthenticator_ITF interface {
	QAuthenticator_PTR() *QAuthenticator
}

func (p *QAuthenticator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAuthenticator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAuthenticator(ptr QAuthenticator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAuthenticator_PTR().Pointer()
	}
	return nil
}

func NewQAuthenticatorFromPointer(ptr unsafe.Pointer) *QAuthenticator {
	var n = new(QAuthenticator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAuthenticator) QAuthenticator_PTR() *QAuthenticator {
	return ptr
}

func NewQAuthenticator() *QAuthenticator {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAuthenticator::QAuthenticator")
		}
	}()

	return NewQAuthenticatorFromPointer(C.QAuthenticator_NewQAuthenticator())
}

func NewQAuthenticator2(other QAuthenticator_ITF) *QAuthenticator {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAuthenticator::QAuthenticator")
		}
	}()

	return NewQAuthenticatorFromPointer(C.QAuthenticator_NewQAuthenticator2(PointerFromQAuthenticator(other)))
}

func (ptr *QAuthenticator) IsNull() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAuthenticator::isNull")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAuthenticator_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAuthenticator) Option(opt string) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAuthenticator::option")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QAuthenticator_Option(ptr.Pointer(), C.CString(opt)))
	}
	return nil
}

func (ptr *QAuthenticator) Password() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAuthenticator::password")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QAuthenticator_Password(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAuthenticator) Realm() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAuthenticator::realm")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QAuthenticator_Realm(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAuthenticator) SetOption(opt string, value core.QVariant_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAuthenticator::setOption")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAuthenticator_SetOption(ptr.Pointer(), C.CString(opt), core.PointerFromQVariant(value))
	}
}

func (ptr *QAuthenticator) SetPassword(password string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAuthenticator::setPassword")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAuthenticator_SetPassword(ptr.Pointer(), C.CString(password))
	}
}

func (ptr *QAuthenticator) SetUser(user string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAuthenticator::setUser")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAuthenticator_SetUser(ptr.Pointer(), C.CString(user))
	}
}

func (ptr *QAuthenticator) User() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAuthenticator::user")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QAuthenticator_User(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAuthenticator) DestroyQAuthenticator() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAuthenticator::~QAuthenticator")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAuthenticator_DestroyQAuthenticator(ptr.Pointer())
	}
}
