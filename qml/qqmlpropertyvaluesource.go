package qml

//#include "qqmlpropertyvaluesource.h"
import "C"
import (
	"unsafe"
)

type QQmlPropertyValueSource struct {
	ptr unsafe.Pointer
}

type QQmlPropertyValueSource_ITF interface {
	QQmlPropertyValueSource_PTR() *QQmlPropertyValueSource
}

func (p *QQmlPropertyValueSource) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlPropertyValueSource) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlPropertyValueSource(ptr QQmlPropertyValueSource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlPropertyValueSource_PTR().Pointer()
	}
	return nil
}

func NewQQmlPropertyValueSourceFromPointer(ptr unsafe.Pointer) *QQmlPropertyValueSource {
	var n = new(QQmlPropertyValueSource)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlPropertyValueSource) QQmlPropertyValueSource_PTR() *QQmlPropertyValueSource {
	return ptr
}

func (ptr *QQmlPropertyValueSource) SetTarget(property QQmlProperty_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyValueSource_SetTarget(ptr.Pointer(), PointerFromQQmlProperty(property))
	}
}

func (ptr *QQmlPropertyValueSource) DestroyQQmlPropertyValueSource() {
	if ptr.Pointer() != nil {
		C.QQmlPropertyValueSource_DestroyQQmlPropertyValueSource(ptr.Pointer())
	}
}
