package nfc

//#include "qndefnfcurirecord.h"
import "C"
import (
	"unsafe"
)

type QNdefNfcUriRecord struct {
	QNdefRecord
}

type QNdefNfcUriRecordITF interface {
	QNdefRecordITF
	QNdefNfcUriRecordPTR() *QNdefNfcUriRecord
}

func PointerFromQNdefNfcUriRecord(ptr QNdefNfcUriRecordITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefNfcUriRecordPTR().Pointer()
	}
	return nil
}

func QNdefNfcUriRecordFromPointer(ptr unsafe.Pointer) *QNdefNfcUriRecord {
	var n = new(QNdefNfcUriRecord)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNdefNfcUriRecord) QNdefNfcUriRecordPTR() *QNdefNfcUriRecord {
	return ptr
}

func NewQNdefNfcUriRecord() *QNdefNfcUriRecord {
	return QNdefNfcUriRecordFromPointer(unsafe.Pointer(C.QNdefNfcUriRecord_NewQNdefNfcUriRecord()))
}

func NewQNdefNfcUriRecord2(other QNdefRecordITF) *QNdefNfcUriRecord {
	return QNdefNfcUriRecordFromPointer(unsafe.Pointer(C.QNdefNfcUriRecord_NewQNdefNfcUriRecord2(C.QtObjectPtr(PointerFromQNdefRecord(other)))))
}

func (ptr *QNdefNfcUriRecord) SetUri(uri string) {
	if ptr.Pointer() != nil {
		C.QNdefNfcUriRecord_SetUri(C.QtObjectPtr(ptr.Pointer()), C.CString(uri))
	}
}

func (ptr *QNdefNfcUriRecord) Uri() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNdefNfcUriRecord_Uri(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}
