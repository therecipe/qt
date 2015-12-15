package qml

//#include "qml.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	for len(n.ObjectNameAbs()) < len("QQmlPropertyValueSource_") {
		n.SetObjectNameAbs("QQmlPropertyValueSource_" + qt.Identifier())
	}
	return n
}

func (ptr *QQmlPropertyValueSource) QQmlPropertyValueSource_PTR() *QQmlPropertyValueSource {
	return ptr
}

func (ptr *QQmlPropertyValueSource) SetTarget(property QQmlProperty_ITF) {
	defer qt.Recovering("QQmlPropertyValueSource::setTarget")

	if ptr.Pointer() != nil {
		C.QQmlPropertyValueSource_SetTarget(ptr.Pointer(), PointerFromQQmlProperty(property))
	}
}

func (ptr *QQmlPropertyValueSource) DestroyQQmlPropertyValueSource() {
	defer qt.Recovering("QQmlPropertyValueSource::~QQmlPropertyValueSource")

	if ptr.Pointer() != nil {
		C.QQmlPropertyValueSource_DestroyQQmlPropertyValueSource(ptr.Pointer())
	}
}

func (ptr *QQmlPropertyValueSource) ObjectNameAbs() string {
	defer qt.Recovering("QQmlPropertyValueSource::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlPropertyValueSource_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlPropertyValueSource) SetObjectNameAbs(name string) {
	defer qt.Recovering("QQmlPropertyValueSource::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QQmlPropertyValueSource_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
