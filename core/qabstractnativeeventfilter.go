package core

//#include "qabstractnativeeventfilter.h"
import "C"
import (
	"unsafe"
)

type QAbstractNativeEventFilter struct {
	ptr unsafe.Pointer
}

type QAbstractNativeEventFilterITF interface {
	QAbstractNativeEventFilterPTR() *QAbstractNativeEventFilter
}

func (p *QAbstractNativeEventFilter) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAbstractNativeEventFilter) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAbstractNativeEventFilter(ptr QAbstractNativeEventFilterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractNativeEventFilterPTR().Pointer()
	}
	return nil
}

func QAbstractNativeEventFilterFromPointer(ptr unsafe.Pointer) *QAbstractNativeEventFilter {
	var n = new(QAbstractNativeEventFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAbstractNativeEventFilter) QAbstractNativeEventFilterPTR() *QAbstractNativeEventFilter {
	return ptr
}

func (ptr *QAbstractNativeEventFilter) DestroyQAbstractNativeEventFilter() {
	if ptr.Pointer() != nil {
		C.QAbstractNativeEventFilter_DestroyQAbstractNativeEventFilter(C.QtObjectPtr(ptr.Pointer()))
	}
}
