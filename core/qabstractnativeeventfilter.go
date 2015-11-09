package core

//#include "qabstractnativeeventfilter.h"
import "C"
import (
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
	return n
}

func (ptr *QAbstractNativeEventFilter) QAbstractNativeEventFilter_PTR() *QAbstractNativeEventFilter {
	return ptr
}

func (ptr *QAbstractNativeEventFilter) DestroyQAbstractNativeEventFilter() {
	if ptr.Pointer() != nil {
		C.QAbstractNativeEventFilter_DestroyQAbstractNativeEventFilter(ptr.Pointer())
	}
}
