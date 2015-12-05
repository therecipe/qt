package xml

//#include "xml.h"
import "C"
import (
	"log"
	"unsafe"
)

type QDomImplementation struct {
	ptr unsafe.Pointer
}

type QDomImplementation_ITF interface {
	QDomImplementation_PTR() *QDomImplementation
}

func (p *QDomImplementation) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDomImplementation) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDomImplementation(ptr QDomImplementation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomImplementation_PTR().Pointer()
	}
	return nil
}

func NewQDomImplementationFromPointer(ptr unsafe.Pointer) *QDomImplementation {
	var n = new(QDomImplementation)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomImplementation) QDomImplementation_PTR() *QDomImplementation {
	return ptr
}

//QDomImplementation::InvalidDataPolicy
type QDomImplementation__InvalidDataPolicy int64

const (
	QDomImplementation__AcceptInvalidChars = QDomImplementation__InvalidDataPolicy(0)
	QDomImplementation__DropInvalidChars   = QDomImplementation__InvalidDataPolicy(1)
	QDomImplementation__ReturnNullNode     = QDomImplementation__InvalidDataPolicy(2)
)

func NewQDomImplementation() *QDomImplementation {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomImplementation::QDomImplementation")
		}
	}()

	return NewQDomImplementationFromPointer(C.QDomImplementation_NewQDomImplementation())
}

func NewQDomImplementation2(x QDomImplementation_ITF) *QDomImplementation {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomImplementation::QDomImplementation")
		}
	}()

	return NewQDomImplementationFromPointer(C.QDomImplementation_NewQDomImplementation2(PointerFromQDomImplementation(x)))
}

func (ptr *QDomImplementation) HasFeature(feature string, version string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomImplementation::hasFeature")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QDomImplementation_HasFeature(ptr.Pointer(), C.CString(feature), C.CString(version)) != 0
	}
	return false
}

func QDomImplementation_InvalidDataPolicy() QDomImplementation__InvalidDataPolicy {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomImplementation::invalidDataPolicy")
		}
	}()

	return QDomImplementation__InvalidDataPolicy(C.QDomImplementation_QDomImplementation_InvalidDataPolicy())
}

func (ptr *QDomImplementation) IsNull() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomImplementation::isNull")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QDomImplementation_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func QDomImplementation_SetInvalidDataPolicy(policy QDomImplementation__InvalidDataPolicy) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomImplementation::setInvalidDataPolicy")
		}
	}()

	C.QDomImplementation_QDomImplementation_SetInvalidDataPolicy(C.int(policy))
}

func (ptr *QDomImplementation) DestroyQDomImplementation() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomImplementation::~QDomImplementation")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDomImplementation_DestroyQDomImplementation(ptr.Pointer())
	}
}
