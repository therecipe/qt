package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QAbstractNativeEventFilter struct {
	ptr unsafe.Pointer
}

type QAbstractNativeEventFilter_ITF interface {
	QAbstractNativeEventFilter_PTR() *QAbstractNativeEventFilter
}

func (p *QAbstractNativeEventFilter) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAbstractNativeEventFilter) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAbstractNativeEventFilter(ptr QAbstractNativeEventFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractNativeEventFilter_PTR().Pointer()
	}
	return nil
}

func NewQAbstractNativeEventFilterFromPointer(ptr unsafe.Pointer) *QAbstractNativeEventFilter {
	var n = new(QAbstractNativeEventFilter)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QAbstractNativeEventFilter_") {
		n.SetObjectNameAbs("QAbstractNativeEventFilter_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractNativeEventFilter) QAbstractNativeEventFilter_PTR() *QAbstractNativeEventFilter {
	return ptr
}

func (ptr *QAbstractNativeEventFilter) DestroyQAbstractNativeEventFilter() {
	defer qt.Recovering("QAbstractNativeEventFilter::~QAbstractNativeEventFilter")

	if ptr.Pointer() != nil {
		C.QAbstractNativeEventFilter_DestroyQAbstractNativeEventFilter(ptr.Pointer())
	}
}

func (ptr *QAbstractNativeEventFilter) ObjectNameAbs() string {
	defer qt.Recovering("QAbstractNativeEventFilter::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractNativeEventFilter_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstractNativeEventFilter) SetObjectNameAbs(name string) {
	defer qt.Recovering("QAbstractNativeEventFilter::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QAbstractNativeEventFilter_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
