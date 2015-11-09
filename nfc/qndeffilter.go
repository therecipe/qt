package nfc

//#include "qndeffilter.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QNdefFilter struct {
	ptr unsafe.Pointer
}

type QNdefFilter_ITF interface {
	QNdefFilter_PTR() *QNdefFilter
}

func (p *QNdefFilter) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNdefFilter) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQNdefFilter(ptr QNdefFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefFilter_PTR().Pointer()
	}
	return nil
}

func NewQNdefFilterFromPointer(ptr unsafe.Pointer) *QNdefFilter {
	var n = new(QNdefFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNdefFilter) QNdefFilter_PTR() *QNdefFilter {
	return ptr
}

func NewQNdefFilter() *QNdefFilter {
	return NewQNdefFilterFromPointer(C.QNdefFilter_NewQNdefFilter())
}

func NewQNdefFilter2(other QNdefFilter_ITF) *QNdefFilter {
	return NewQNdefFilterFromPointer(C.QNdefFilter_NewQNdefFilter2(PointerFromQNdefFilter(other)))
}

func (ptr *QNdefFilter) Clear() {
	if ptr.Pointer() != nil {
		C.QNdefFilter_Clear(ptr.Pointer())
	}
}

func (ptr *QNdefFilter) OrderMatch() bool {
	if ptr.Pointer() != nil {
		return C.QNdefFilter_OrderMatch(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNdefFilter) RecordCount() int {
	if ptr.Pointer() != nil {
		return int(C.QNdefFilter_RecordCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNdefFilter) SetOrderMatch(on bool) {
	if ptr.Pointer() != nil {
		C.QNdefFilter_SetOrderMatch(ptr.Pointer(), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QNdefFilter) DestroyQNdefFilter() {
	if ptr.Pointer() != nil {
		C.QNdefFilter_DestroyQNdefFilter(ptr.Pointer())
	}
}
