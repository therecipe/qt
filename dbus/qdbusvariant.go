package dbus

//#include "qdbusvariant.h"
import "C"
import (
	"unsafe"
)

type QDBusVariant struct {
	ptr unsafe.Pointer
}

type QDBusVariantITF interface {
	QDBusVariantPTR() *QDBusVariant
}

func (p *QDBusVariant) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusVariant) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusVariant(ptr QDBusVariantITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusVariantPTR().Pointer()
	}
	return nil
}

func QDBusVariantFromPointer(ptr unsafe.Pointer) *QDBusVariant {
	var n = new(QDBusVariant)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDBusVariant) QDBusVariantPTR() *QDBusVariant {
	return ptr
}

func NewQDBusVariant() *QDBusVariant {
	return QDBusVariantFromPointer(unsafe.Pointer(C.QDBusVariant_NewQDBusVariant()))
}

func NewQDBusVariant2(variant string) *QDBusVariant {
	return QDBusVariantFromPointer(unsafe.Pointer(C.QDBusVariant_NewQDBusVariant2(C.CString(variant))))
}

func (ptr *QDBusVariant) SetVariant(variant string) {
	if ptr.Pointer() != nil {
		C.QDBusVariant_SetVariant(C.QtObjectPtr(ptr.Pointer()), C.CString(variant))
	}
}

func (ptr *QDBusVariant) Variant() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusVariant_Variant(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}
