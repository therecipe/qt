package core

//#include "quuid.h"
import "C"
import (
	"unsafe"
)

type QUuid struct {
	ptr unsafe.Pointer
}

type QUuidITF interface {
	QUuidPTR() *QUuid
}

func (p *QUuid) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QUuid) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQUuid(ptr QUuidITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QUuidPTR().Pointer()
	}
	return nil
}

func QUuidFromPointer(ptr unsafe.Pointer) *QUuid {
	var n = new(QUuid)
	n.SetPointer(ptr)
	return n
}

func (ptr *QUuid) QUuidPTR() *QUuid {
	return ptr
}

//QUuid::Variant
type QUuid__Variant int

var (
	QUuid__VarUnknown = QUuid__Variant(-1)
	QUuid__NCS        = QUuid__Variant(0)
	QUuid__DCE        = QUuid__Variant(2)
	QUuid__Microsoft  = QUuid__Variant(6)
	QUuid__Reserved   = QUuid__Variant(7)
)

//QUuid::Version
type QUuid__Version int

var (
	QUuid__VerUnknown    = QUuid__Version(-1)
	QUuid__Time          = QUuid__Version(1)
	QUuid__EmbeddedPOSIX = QUuid__Version(2)
	QUuid__Md5           = QUuid__Version(3)
	QUuid__Name          = QUuid__Version(QUuid__Md5)
	QUuid__Random        = QUuid__Version(4)
	QUuid__Sha1          = QUuid__Version(5)
)

func (ptr *QUuid) Variant() QUuid__Variant {
	if ptr.Pointer() != nil {
		return QUuid__Variant(C.QUuid_Variant(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QUuid) Version() QUuid__Version {
	if ptr.Pointer() != nil {
		return QUuid__Version(C.QUuid_Version(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQUuid() *QUuid {
	return QUuidFromPointer(unsafe.Pointer(C.QUuid_NewQUuid()))
}

func NewQUuid5(text QByteArrayITF) *QUuid {
	return QUuidFromPointer(unsafe.Pointer(C.QUuid_NewQUuid5(C.QtObjectPtr(PointerFromQByteArray(text)))))
}

func NewQUuid3(text string) *QUuid {
	return QUuidFromPointer(unsafe.Pointer(C.QUuid_NewQUuid3(C.CString(text))))
}

func (ptr *QUuid) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QUuid_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QUuid) ToString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QUuid_ToString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}
