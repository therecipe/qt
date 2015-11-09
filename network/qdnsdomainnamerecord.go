package network

//#include "qdnsdomainnamerecord.h"
import "C"
import (
	"unsafe"
)

type QDnsDomainNameRecord struct {
	ptr unsafe.Pointer
}

type QDnsDomainNameRecord_ITF interface {
	QDnsDomainNameRecord_PTR() *QDnsDomainNameRecord
}

func (p *QDnsDomainNameRecord) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDnsDomainNameRecord) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDnsDomainNameRecord(ptr QDnsDomainNameRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsDomainNameRecord_PTR().Pointer()
	}
	return nil
}

func NewQDnsDomainNameRecordFromPointer(ptr unsafe.Pointer) *QDnsDomainNameRecord {
	var n = new(QDnsDomainNameRecord)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDnsDomainNameRecord) QDnsDomainNameRecord_PTR() *QDnsDomainNameRecord {
	return ptr
}

func NewQDnsDomainNameRecord() *QDnsDomainNameRecord {
	return NewQDnsDomainNameRecordFromPointer(C.QDnsDomainNameRecord_NewQDnsDomainNameRecord())
}

func NewQDnsDomainNameRecord2(other QDnsDomainNameRecord_ITF) *QDnsDomainNameRecord {
	return NewQDnsDomainNameRecordFromPointer(C.QDnsDomainNameRecord_NewQDnsDomainNameRecord2(PointerFromQDnsDomainNameRecord(other)))
}

func (ptr *QDnsDomainNameRecord) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsDomainNameRecord_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsDomainNameRecord) Swap(other QDnsDomainNameRecord_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsDomainNameRecord_Swap(ptr.Pointer(), PointerFromQDnsDomainNameRecord(other))
	}
}

func (ptr *QDnsDomainNameRecord) Value() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsDomainNameRecord_Value(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsDomainNameRecord) DestroyQDnsDomainNameRecord() {
	if ptr.Pointer() != nil {
		C.QDnsDomainNameRecord_DestroyQDnsDomainNameRecord(ptr.Pointer())
	}
}
