package network

//#include "qdnsmailexchangerecord.h"
import "C"
import (
	"unsafe"
)

type QDnsMailExchangeRecord struct {
	ptr unsafe.Pointer
}

type QDnsMailExchangeRecordITF interface {
	QDnsMailExchangeRecordPTR() *QDnsMailExchangeRecord
}

func (p *QDnsMailExchangeRecord) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDnsMailExchangeRecord) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDnsMailExchangeRecord(ptr QDnsMailExchangeRecordITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsMailExchangeRecordPTR().Pointer()
	}
	return nil
}

func QDnsMailExchangeRecordFromPointer(ptr unsafe.Pointer) *QDnsMailExchangeRecord {
	var n = new(QDnsMailExchangeRecord)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDnsMailExchangeRecord) QDnsMailExchangeRecordPTR() *QDnsMailExchangeRecord {
	return ptr
}

func NewQDnsMailExchangeRecord() *QDnsMailExchangeRecord {
	return QDnsMailExchangeRecordFromPointer(unsafe.Pointer(C.QDnsMailExchangeRecord_NewQDnsMailExchangeRecord()))
}

func NewQDnsMailExchangeRecord2(other QDnsMailExchangeRecordITF) *QDnsMailExchangeRecord {
	return QDnsMailExchangeRecordFromPointer(unsafe.Pointer(C.QDnsMailExchangeRecord_NewQDnsMailExchangeRecord2(C.QtObjectPtr(PointerFromQDnsMailExchangeRecord(other)))))
}

func (ptr *QDnsMailExchangeRecord) Exchange() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsMailExchangeRecord_Exchange(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDnsMailExchangeRecord) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsMailExchangeRecord_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDnsMailExchangeRecord) Swap(other QDnsMailExchangeRecordITF) {
	if ptr.Pointer() != nil {
		C.QDnsMailExchangeRecord_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQDnsMailExchangeRecord(other)))
	}
}

func (ptr *QDnsMailExchangeRecord) DestroyQDnsMailExchangeRecord() {
	if ptr.Pointer() != nil {
		C.QDnsMailExchangeRecord_DestroyQDnsMailExchangeRecord(C.QtObjectPtr(ptr.Pointer()))
	}
}
