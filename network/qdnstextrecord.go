package network

//#include "qdnstextrecord.h"
import "C"
import (
	"unsafe"
)

type QDnsTextRecord struct {
	ptr unsafe.Pointer
}

type QDnsTextRecordITF interface {
	QDnsTextRecordPTR() *QDnsTextRecord
}

func (p *QDnsTextRecord) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDnsTextRecord) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDnsTextRecord(ptr QDnsTextRecordITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsTextRecordPTR().Pointer()
	}
	return nil
}

func QDnsTextRecordFromPointer(ptr unsafe.Pointer) *QDnsTextRecord {
	var n = new(QDnsTextRecord)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDnsTextRecord) QDnsTextRecordPTR() *QDnsTextRecord {
	return ptr
}

func NewQDnsTextRecord() *QDnsTextRecord {
	return QDnsTextRecordFromPointer(unsafe.Pointer(C.QDnsTextRecord_NewQDnsTextRecord()))
}

func NewQDnsTextRecord2(other QDnsTextRecordITF) *QDnsTextRecord {
	return QDnsTextRecordFromPointer(unsafe.Pointer(C.QDnsTextRecord_NewQDnsTextRecord2(C.QtObjectPtr(PointerFromQDnsTextRecord(other)))))
}

func (ptr *QDnsTextRecord) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsTextRecord_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDnsTextRecord) Swap(other QDnsTextRecordITF) {
	if ptr.Pointer() != nil {
		C.QDnsTextRecord_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQDnsTextRecord(other)))
	}
}

func (ptr *QDnsTextRecord) DestroyQDnsTextRecord() {
	if ptr.Pointer() != nil {
		C.QDnsTextRecord_DestroyQDnsTextRecord(C.QtObjectPtr(ptr.Pointer()))
	}
}
