package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QDnsMailExchangeRecord struct {
	ptr unsafe.Pointer
}

type QDnsMailExchangeRecord_ITF interface {
	QDnsMailExchangeRecord_PTR() *QDnsMailExchangeRecord
}

func (p *QDnsMailExchangeRecord) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDnsMailExchangeRecord) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDnsMailExchangeRecord(ptr QDnsMailExchangeRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsMailExchangeRecord_PTR().Pointer()
	}
	return nil
}

func NewQDnsMailExchangeRecordFromPointer(ptr unsafe.Pointer) *QDnsMailExchangeRecord {
	var n = new(QDnsMailExchangeRecord)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDnsMailExchangeRecord) QDnsMailExchangeRecord_PTR() *QDnsMailExchangeRecord {
	return ptr
}

func NewQDnsMailExchangeRecord() *QDnsMailExchangeRecord {
	defer qt.Recovering("QDnsMailExchangeRecord::QDnsMailExchangeRecord")

	return NewQDnsMailExchangeRecordFromPointer(C.QDnsMailExchangeRecord_NewQDnsMailExchangeRecord())
}

func NewQDnsMailExchangeRecord2(other QDnsMailExchangeRecord_ITF) *QDnsMailExchangeRecord {
	defer qt.Recovering("QDnsMailExchangeRecord::QDnsMailExchangeRecord")

	return NewQDnsMailExchangeRecordFromPointer(C.QDnsMailExchangeRecord_NewQDnsMailExchangeRecord2(PointerFromQDnsMailExchangeRecord(other)))
}

func (ptr *QDnsMailExchangeRecord) Exchange() string {
	defer qt.Recovering("QDnsMailExchangeRecord::exchange")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsMailExchangeRecord_Exchange(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsMailExchangeRecord) Name() string {
	defer qt.Recovering("QDnsMailExchangeRecord::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsMailExchangeRecord_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsMailExchangeRecord) Swap(other QDnsMailExchangeRecord_ITF) {
	defer qt.Recovering("QDnsMailExchangeRecord::swap")

	if ptr.Pointer() != nil {
		C.QDnsMailExchangeRecord_Swap(ptr.Pointer(), PointerFromQDnsMailExchangeRecord(other))
	}
}

func (ptr *QDnsMailExchangeRecord) DestroyQDnsMailExchangeRecord() {
	defer qt.Recovering("QDnsMailExchangeRecord::~QDnsMailExchangeRecord")

	if ptr.Pointer() != nil {
		C.QDnsMailExchangeRecord_DestroyQDnsMailExchangeRecord(ptr.Pointer())
	}
}
