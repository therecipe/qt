package nfc

//#include "qndefnfctextrecord.h"
import "C"
import (
	"unsafe"
)

type QNdefNfcTextRecord struct {
	QNdefRecord
}

type QNdefNfcTextRecordITF interface {
	QNdefRecordITF
	QNdefNfcTextRecordPTR() *QNdefNfcTextRecord
}

func PointerFromQNdefNfcTextRecord(ptr QNdefNfcTextRecordITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefNfcTextRecordPTR().Pointer()
	}
	return nil
}

func QNdefNfcTextRecordFromPointer(ptr unsafe.Pointer) *QNdefNfcTextRecord {
	var n = new(QNdefNfcTextRecord)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNdefNfcTextRecord) QNdefNfcTextRecordPTR() *QNdefNfcTextRecord {
	return ptr
}

//QNdefNfcTextRecord::Encoding
type QNdefNfcTextRecord__Encoding int

var (
	QNdefNfcTextRecord__Utf8  = QNdefNfcTextRecord__Encoding(0)
	QNdefNfcTextRecord__Utf16 = QNdefNfcTextRecord__Encoding(1)
)

func NewQNdefNfcTextRecord() *QNdefNfcTextRecord {
	return QNdefNfcTextRecordFromPointer(unsafe.Pointer(C.QNdefNfcTextRecord_NewQNdefNfcTextRecord()))
}

func NewQNdefNfcTextRecord2(other QNdefRecordITF) *QNdefNfcTextRecord {
	return QNdefNfcTextRecordFromPointer(unsafe.Pointer(C.QNdefNfcTextRecord_NewQNdefNfcTextRecord2(C.QtObjectPtr(PointerFromQNdefRecord(other)))))
}

func (ptr *QNdefNfcTextRecord) Encoding() QNdefNfcTextRecord__Encoding {
	if ptr.Pointer() != nil {
		return QNdefNfcTextRecord__Encoding(C.QNdefNfcTextRecord_Encoding(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNdefNfcTextRecord) Locale() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNdefNfcTextRecord_Locale(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QNdefNfcTextRecord) SetEncoding(encoding QNdefNfcTextRecord__Encoding) {
	if ptr.Pointer() != nil {
		C.QNdefNfcTextRecord_SetEncoding(C.QtObjectPtr(ptr.Pointer()), C.int(encoding))
	}
}

func (ptr *QNdefNfcTextRecord) SetLocale(locale string) {
	if ptr.Pointer() != nil {
		C.QNdefNfcTextRecord_SetLocale(C.QtObjectPtr(ptr.Pointer()), C.CString(locale))
	}
}

func (ptr *QNdefNfcTextRecord) SetText(text string) {
	if ptr.Pointer() != nil {
		C.QNdefNfcTextRecord_SetText(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QNdefNfcTextRecord) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNdefNfcTextRecord_Text(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}
