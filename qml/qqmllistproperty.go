package qml

//#include "qml.h"
import "C"
import (
	"unsafe"
)

type QQmlListProperty struct {
	ptr unsafe.Pointer
}

type QQmlListProperty_ITF interface {
	QQmlListProperty_PTR() *QQmlListProperty
}

func (p *QQmlListProperty) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlListProperty) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlListProperty(ptr QQmlListProperty_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlListProperty_PTR().Pointer()
	}
	return nil
}

func NewQQmlListPropertyFromPointer(ptr unsafe.Pointer) *QQmlListProperty {
	var n = new(QQmlListProperty)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlListProperty) QQmlListProperty_PTR() *QQmlListProperty {
	return ptr
}
