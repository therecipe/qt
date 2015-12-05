package core

//#include "core.h"
import "C"
import (
	"log"
	"strings"
	"unsafe"
)

type QProcessEnvironment struct {
	ptr unsafe.Pointer
}

type QProcessEnvironment_ITF interface {
	QProcessEnvironment_PTR() *QProcessEnvironment
}

func (p *QProcessEnvironment) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QProcessEnvironment) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQProcessEnvironment(ptr QProcessEnvironment_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProcessEnvironment_PTR().Pointer()
	}
	return nil
}

func NewQProcessEnvironmentFromPointer(ptr unsafe.Pointer) *QProcessEnvironment {
	var n = new(QProcessEnvironment)
	n.SetPointer(ptr)
	return n
}

func (ptr *QProcessEnvironment) QProcessEnvironment_PTR() *QProcessEnvironment {
	return ptr
}

func NewQProcessEnvironment() *QProcessEnvironment {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcessEnvironment::QProcessEnvironment")
		}
	}()

	return NewQProcessEnvironmentFromPointer(C.QProcessEnvironment_NewQProcessEnvironment())
}

func NewQProcessEnvironment2(other QProcessEnvironment_ITF) *QProcessEnvironment {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcessEnvironment::QProcessEnvironment")
		}
	}()

	return NewQProcessEnvironmentFromPointer(C.QProcessEnvironment_NewQProcessEnvironment2(PointerFromQProcessEnvironment(other)))
}

func (ptr *QProcessEnvironment) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcessEnvironment::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcessEnvironment_Clear(ptr.Pointer())
	}
}

func (ptr *QProcessEnvironment) Contains(name string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcessEnvironment::contains")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QProcessEnvironment_Contains(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QProcessEnvironment) IsEmpty() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcessEnvironment::isEmpty")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QProcessEnvironment_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QProcessEnvironment) Keys() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcessEnvironment::keys")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QProcessEnvironment_Keys(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QProcessEnvironment) Swap(other QProcessEnvironment_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcessEnvironment::swap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcessEnvironment_Swap(ptr.Pointer(), PointerFromQProcessEnvironment(other))
	}
}

func (ptr *QProcessEnvironment) ToStringList() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcessEnvironment::toStringList")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QProcessEnvironment_ToStringList(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QProcessEnvironment) Value(name string, defaultValue string) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcessEnvironment::value")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QProcessEnvironment_Value(ptr.Pointer(), C.CString(name), C.CString(defaultValue)))
	}
	return ""
}

func (ptr *QProcessEnvironment) DestroyQProcessEnvironment() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcessEnvironment::~QProcessEnvironment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcessEnvironment_DestroyQProcessEnvironment(ptr.Pointer())
	}
}
