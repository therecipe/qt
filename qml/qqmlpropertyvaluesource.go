package qml

//#include "qqmlpropertyvaluesource.h"
import "C"
import (
	"unsafe"
)

type QQmlPropertyValueSource struct {
	ptr unsafe.Pointer
}

type QQmlPropertyValueSourceITF interface {
	QQmlPropertyValueSourcePTR() *QQmlPropertyValueSource
}

func (p *QQmlPropertyValueSource) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlPropertyValueSource) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlPropertyValueSource(ptr QQmlPropertyValueSourceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlPropertyValueSourcePTR().Pointer()
	}
	return nil
}

func QQmlPropertyValueSourceFromPointer(ptr unsafe.Pointer) *QQmlPropertyValueSource {
	var n = new(QQmlPropertyValueSource)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlPropertyValueSource) QQmlPropertyValueSourcePTR() *QQmlPropertyValueSource {
	return ptr
}

func (ptr *QQmlPropertyValueSource) SetTarget(property QQmlPropertyITF) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyValueSource_SetTarget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQQmlProperty(property)))
	}
}

func (ptr *QQmlPropertyValueSource) DestroyQQmlPropertyValueSource() {
	if ptr.Pointer() != nil {
		C.QQmlPropertyValueSource_DestroyQQmlPropertyValueSource(C.QtObjectPtr(ptr.Pointer()))
	}
}
