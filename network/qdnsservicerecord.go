package network

//#include "qdnsservicerecord.h"
import "C"
import (
	"unsafe"
)

type QDnsServiceRecord struct {
	ptr unsafe.Pointer
}

type QDnsServiceRecordITF interface {
	QDnsServiceRecordPTR() *QDnsServiceRecord
}

func (p *QDnsServiceRecord) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDnsServiceRecord) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDnsServiceRecord(ptr QDnsServiceRecordITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsServiceRecordPTR().Pointer()
	}
	return nil
}

func QDnsServiceRecordFromPointer(ptr unsafe.Pointer) *QDnsServiceRecord {
	var n = new(QDnsServiceRecord)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDnsServiceRecord) QDnsServiceRecordPTR() *QDnsServiceRecord {
	return ptr
}

func NewQDnsServiceRecord() *QDnsServiceRecord {
	return QDnsServiceRecordFromPointer(unsafe.Pointer(C.QDnsServiceRecord_NewQDnsServiceRecord()))
}

func NewQDnsServiceRecord2(other QDnsServiceRecordITF) *QDnsServiceRecord {
	return QDnsServiceRecordFromPointer(unsafe.Pointer(C.QDnsServiceRecord_NewQDnsServiceRecord2(C.QtObjectPtr(PointerFromQDnsServiceRecord(other)))))
}

func (ptr *QDnsServiceRecord) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsServiceRecord_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDnsServiceRecord) Swap(other QDnsServiceRecordITF) {
	if ptr.Pointer() != nil {
		C.QDnsServiceRecord_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQDnsServiceRecord(other)))
	}
}

func (ptr *QDnsServiceRecord) Target() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsServiceRecord_Target(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDnsServiceRecord) DestroyQDnsServiceRecord() {
	if ptr.Pointer() != nil {
		C.QDnsServiceRecord_DestroyQDnsServiceRecord(C.QtObjectPtr(ptr.Pointer()))
	}
}
