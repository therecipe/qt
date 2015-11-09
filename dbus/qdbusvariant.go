package dbus

//#include "qdbusvariant.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDBusVariant struct {
	ptr unsafe.Pointer
}

type QDBusVariant_ITF interface {
	QDBusVariant_PTR() *QDBusVariant
}

func (p *QDBusVariant) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusVariant) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusVariant(ptr QDBusVariant_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusVariant_PTR().Pointer()
	}
	return nil
}

func NewQDBusVariantFromPointer(ptr unsafe.Pointer) *QDBusVariant {
	var n = new(QDBusVariant)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDBusVariant) QDBusVariant_PTR() *QDBusVariant {
	return ptr
}

func NewQDBusVariant() *QDBusVariant {
	return NewQDBusVariantFromPointer(C.QDBusVariant_NewQDBusVariant())
}

func NewQDBusVariant2(variant core.QVariant_ITF) *QDBusVariant {
	return NewQDBusVariantFromPointer(C.QDBusVariant_NewQDBusVariant2(core.PointerFromQVariant(variant)))
}

func (ptr *QDBusVariant) SetVariant(variant core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusVariant_SetVariant(ptr.Pointer(), core.PointerFromQVariant(variant))
	}
}

func (ptr *QDBusVariant) Variant() *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QDBusVariant_Variant(ptr.Pointer()))
	}
	return nil
}
