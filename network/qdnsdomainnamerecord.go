package network

//#include "qdnsdomainnamerecord.h"
import "C"
import (
	"unsafe"
)

type QDnsDomainNameRecord struct {
	ptr unsafe.Pointer
}

type QDnsDomainNameRecordITF interface {
	QDnsDomainNameRecordPTR() *QDnsDomainNameRecord
}

func (p *QDnsDomainNameRecord) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDnsDomainNameRecord) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDnsDomainNameRecord(ptr QDnsDomainNameRecordITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsDomainNameRecordPTR().Pointer()
	}
	return nil
}

func QDnsDomainNameRecordFromPointer(ptr unsafe.Pointer) *QDnsDomainNameRecord {
	var n = new(QDnsDomainNameRecord)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDnsDomainNameRecord) QDnsDomainNameRecordPTR() *QDnsDomainNameRecord {
	return ptr
}

func NewQDnsDomainNameRecord() *QDnsDomainNameRecord {
	return QDnsDomainNameRecordFromPointer(unsafe.Pointer(C.QDnsDomainNameRecord_NewQDnsDomainNameRecord()))
}

func NewQDnsDomainNameRecord2(other QDnsDomainNameRecordITF) *QDnsDomainNameRecord {
	return QDnsDomainNameRecordFromPointer(unsafe.Pointer(C.QDnsDomainNameRecord_NewQDnsDomainNameRecord2(C.QtObjectPtr(PointerFromQDnsDomainNameRecord(other)))))
}

func (ptr *QDnsDomainNameRecord) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsDomainNameRecord_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDnsDomainNameRecord) Swap(other QDnsDomainNameRecordITF) {
	if ptr.Pointer() != nil {
		C.QDnsDomainNameRecord_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQDnsDomainNameRecord(other)))
	}
}

func (ptr *QDnsDomainNameRecord) Value() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsDomainNameRecord_Value(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDnsDomainNameRecord) DestroyQDnsDomainNameRecord() {
	if ptr.Pointer() != nil {
		C.QDnsDomainNameRecord_DestroyQDnsDomainNameRecord(C.QtObjectPtr(ptr.Pointer()))
	}
}
