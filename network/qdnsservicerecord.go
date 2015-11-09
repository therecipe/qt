package network

//#include "qdnsservicerecord.h"
import "C"
import (
	"unsafe"
)

type QDnsServiceRecord struct {
	ptr unsafe.Pointer
}

type QDnsServiceRecord_ITF interface {
	QDnsServiceRecord_PTR() *QDnsServiceRecord
}

func (p *QDnsServiceRecord) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDnsServiceRecord) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDnsServiceRecord(ptr QDnsServiceRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsServiceRecord_PTR().Pointer()
	}
	return nil
}

func NewQDnsServiceRecordFromPointer(ptr unsafe.Pointer) *QDnsServiceRecord {
	var n = new(QDnsServiceRecord)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDnsServiceRecord) QDnsServiceRecord_PTR() *QDnsServiceRecord {
	return ptr
}

func NewQDnsServiceRecord() *QDnsServiceRecord {
	return NewQDnsServiceRecordFromPointer(C.QDnsServiceRecord_NewQDnsServiceRecord())
}

func NewQDnsServiceRecord2(other QDnsServiceRecord_ITF) *QDnsServiceRecord {
	return NewQDnsServiceRecordFromPointer(C.QDnsServiceRecord_NewQDnsServiceRecord2(PointerFromQDnsServiceRecord(other)))
}

func (ptr *QDnsServiceRecord) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsServiceRecord_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsServiceRecord) Swap(other QDnsServiceRecord_ITF) {
	if ptr.Pointer() != nil {
		C.QDnsServiceRecord_Swap(ptr.Pointer(), PointerFromQDnsServiceRecord(other))
	}
}

func (ptr *QDnsServiceRecord) Target() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsServiceRecord_Target(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsServiceRecord) DestroyQDnsServiceRecord() {
	if ptr.Pointer() != nil {
		C.QDnsServiceRecord_DestroyQDnsServiceRecord(ptr.Pointer())
	}
}
