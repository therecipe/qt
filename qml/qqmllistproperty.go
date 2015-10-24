package qml

//#include "qqmllistproperty.h"
import "C"
import (
	"unsafe"
)

type QQmlListProperty struct {
	ptr unsafe.Pointer
}

type QQmlListPropertyITF interface {
	QQmlListPropertyPTR() *QQmlListProperty
}

func (p *QQmlListProperty) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlListProperty) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlListProperty(ptr QQmlListPropertyITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlListPropertyPTR().Pointer()
	}
	return nil
}

func QQmlListPropertyFromPointer(ptr unsafe.Pointer) *QQmlListProperty {
	var n = new(QQmlListProperty)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlListProperty) QQmlListPropertyPTR() *QQmlListProperty {
	return ptr
}
