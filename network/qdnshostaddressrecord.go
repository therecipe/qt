package network

//#include "qdnshostaddressrecord.h"
import "C"
import (
	"unsafe"
)

type QDnsHostAddressRecord struct {
	ptr unsafe.Pointer
}

type QDnsHostAddressRecordITF interface {
	QDnsHostAddressRecordPTR() *QDnsHostAddressRecord
}

func (p *QDnsHostAddressRecord) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDnsHostAddressRecord) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDnsHostAddressRecord(ptr QDnsHostAddressRecordITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsHostAddressRecordPTR().Pointer()
	}
	return nil
}

func QDnsHostAddressRecordFromPointer(ptr unsafe.Pointer) *QDnsHostAddressRecord {
	var n = new(QDnsHostAddressRecord)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDnsHostAddressRecord) QDnsHostAddressRecordPTR() *QDnsHostAddressRecord {
	return ptr
}

func NewQDnsHostAddressRecord() *QDnsHostAddressRecord {
	return QDnsHostAddressRecordFromPointer(unsafe.Pointer(C.QDnsHostAddressRecord_NewQDnsHostAddressRecord()))
}

func NewQDnsHostAddressRecord2(other QDnsHostAddressRecordITF) *QDnsHostAddressRecord {
	return QDnsHostAddressRecordFromPointer(unsafe.Pointer(C.QDnsHostAddressRecord_NewQDnsHostAddressRecord2(C.QtObjectPtr(PointerFromQDnsHostAddressRecord(other)))))
}

func (ptr *QDnsHostAddressRecord) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsHostAddressRecord_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDnsHostAddressRecord) Swap(other QDnsHostAddressRecordITF) {
	if ptr.Pointer() != nil {
		C.QDnsHostAddressRecord_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQDnsHostAddressRecord(other)))
	}
}

func (ptr *QDnsHostAddressRecord) DestroyQDnsHostAddressRecord() {
	if ptr.Pointer() != nil {
		C.QDnsHostAddressRecord_DestroyQDnsHostAddressRecord(C.QtObjectPtr(ptr.Pointer()))
	}
}
