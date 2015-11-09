package network

//#include "qdnshostaddressrecord.h"
import "C"
import (
	"unsafe"
)

type QDnsHostAddressRecord struct {
	ptr unsafe.Pointer
}

type QDnsHostAddressRecord_ITF interface {
	QDnsHostAddressRecord_PTR() *QDnsHostAddressRecord
}

func (p *QDnsHostAddressRecord) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDnsHostAddressRecord) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDnsHostAddressRecord(ptr QDnsHostAddressRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsHostAddressRecord_PTR().Pointer()
	}
	return nil
}

func NewQDnsHostAddressRecordFromPointer(ptr unsafe.Pointer) *QDnsHostAddressRecord {
	var n = new(QDnsHostAddressRecord)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDnsHostAddressRecord) QDnsHostAddressRecord_PTR() *QDnsHostAddressRecord {
	return ptr
}

func NewQDnsHostAddressRecord() *QDnsHostAddressRecord {
	return NewQDnsHostAddressRecordFromPointer(C.QDnsHostAddressRecord_NewQDnsHostAddressRecord())
}

func NewQDnsHostAddressRecord2(other QDnsHostAddressRecord_ITF) *QDnsHostAddressRecord {
	return NewQDnsHostAddressRecordFromPointer(C.QDnsHostAddressRecord_NewQDnsHostAddressRecord2(PointerFromQDnsHostAddressRecord(other)))
}

func (ptr *QDnsHostAddressRecord) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsHostAddressRecord_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsHostAddressRecord) Swap(other QDnsHostAddressRecord_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsHostAddressRecord_Swap(ptr.Pointer(), PointerFromQDnsHostAddressRecord(other))
	}
}

func (ptr *QDnsHostAddressRecord) DestroyQDnsHostAddressRecord() {
	if ptr.Pointer() != nil {
		C.QDnsHostAddressRecord_DestroyQDnsHostAddressRecord(ptr.Pointer())
	}
}
