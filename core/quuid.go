package core

//#include "core.h"
import "C"
import (
	"log"
	"unsafe"
)

type QUuid struct {
	ptr unsafe.Pointer
}

type QUuid_ITF interface {
	QUuid_PTR() *QUuid
}

func (p *QUuid) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QUuid) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQUuid(ptr QUuid_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QUuid_PTR().Pointer()
	}
	return nil
}

func NewQUuidFromPointer(ptr unsafe.Pointer) *QUuid {
	var n = new(QUuid)
	n.SetPointer(ptr)
	return n
}

func (ptr *QUuid) QUuid_PTR() *QUuid {
	return ptr
}

//QUuid::Variant
type QUuid__Variant int64

const (
	QUuid__VarUnknown = QUuid__Variant(-1)
	QUuid__NCS        = QUuid__Variant(0)
	QUuid__DCE        = QUuid__Variant(2)
	QUuid__Microsoft  = QUuid__Variant(6)
	QUuid__Reserved   = QUuid__Variant(7)
)

//QUuid::Version
type QUuid__Version int64

const (
	QUuid__VerUnknown    = QUuid__Version(-1)
	QUuid__Time          = QUuid__Version(1)
	QUuid__EmbeddedPOSIX = QUuid__Version(2)
	QUuid__Md5           = QUuid__Version(3)
	QUuid__Name          = QUuid__Version(QUuid__Md5)
	QUuid__Random        = QUuid__Version(4)
	QUuid__Sha1          = QUuid__Version(5)
)

func (ptr *QUuid) Variant() QUuid__Variant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUuid::variant")
		}
	}()

	if ptr.Pointer() != nil {
		return QUuid__Variant(C.QUuid_Variant(ptr.Pointer()))
	}
	return 0
}

func (ptr *QUuid) Version() QUuid__Version {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUuid::version")
		}
	}()

	if ptr.Pointer() != nil {
		return QUuid__Version(C.QUuid_Version(ptr.Pointer()))
	}
	return 0
}

func NewQUuid() *QUuid {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUuid::QUuid")
		}
	}()

	return NewQUuidFromPointer(C.QUuid_NewQUuid())
}

func NewQUuid5(text QByteArray_ITF) *QUuid {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUuid::QUuid")
		}
	}()

	return NewQUuidFromPointer(C.QUuid_NewQUuid5(PointerFromQByteArray(text)))
}

func NewQUuid3(text string) *QUuid {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUuid::QUuid")
		}
	}()

	return NewQUuidFromPointer(C.QUuid_NewQUuid3(C.CString(text)))
}

func (ptr *QUuid) IsNull() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUuid::isNull")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QUuid_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUuid) ToByteArray() *QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUuid::toByteArray")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QUuid_ToByteArray(ptr.Pointer()))
	}
	return nil
}

func (ptr *QUuid) ToRfc4122() *QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUuid::toRfc4122")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QUuid_ToRfc4122(ptr.Pointer()))
	}
	return nil
}

func (ptr *QUuid) ToString() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUuid::toString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QUuid_ToString(ptr.Pointer()))
	}
	return ""
}
