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

type QNdefFilterITF interface {
	QNdefFilterPTR() *QNdefFilter
}

func (p *QNdefFilter) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNdefFilter) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQNdefFilter(ptr QNdefFilterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefFilterPTR().Pointer()
	}
	return nil
}

func QNdefFilterFromPointer(ptr unsafe.Pointer) *QNdefFilter {
	var n = new(QNdefFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNdefFilter) QNdefFilterPTR() *QNdefFilter {
	return ptr
}

func NewQNdefFilter() *QNdefFilter {
	return QNdefFilterFromPointer(unsafe.Pointer(C.QNdefFilter_NewQNdefFilter()))
}

func NewQNdefFilter2(other QNdefFilterITF) *QNdefFilter {
	return QNdefFilterFromPointer(unsafe.Pointer(C.QNdefFilter_NewQNdefFilter2(C.QtObjectPtr(PointerFromQNdefFilter(other)))))
}

func (ptr *QNdefFilter) Clear() {
	if ptr.Pointer() != nil {
		C.QNdefFilter_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QNdefFilter) OrderMatch() bool {
	if ptr.Pointer() != nil {
		return C.QNdefFilter_OrderMatch(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNdefFilter) RecordCount() int {
	if ptr.Pointer() != nil {
		return int(C.QNdefFilter_RecordCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNdefFilter) SetOrderMatch(on bool) {
	if ptr.Pointer() != nil {
		C.QNdefFilter_SetOrderMatch(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QNdefFilter) DestroyQNdefFilter() {
	if ptr.Pointer() != nil {
		C.QNdefFilter_DestroyQNdefFilter(C.QtObjectPtr(ptr.Pointer()))
	}
}
