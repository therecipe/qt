package nfc

//#include "nfc.h"
import "C"
import (
	"log"
	"unsafe"
)

type QNdefNfcTextRecord struct {
	QNdefRecord
}

type QNdefNfcTextRecord_ITF interface {
	QNdefRecord_ITF
	QNdefNfcTextRecord_PTR() *QNdefNfcTextRecord
}

func PointerFromQNdefNfcTextRecord(ptr QNdefNfcTextRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefNfcTextRecord_PTR().Pointer()
	}
	return nil
}

func NewQNdefNfcTextRecordFromPointer(ptr unsafe.Pointer) *QNdefNfcTextRecord {
	var n = new(QNdefNfcTextRecord)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNdefNfcTextRecord) QNdefNfcTextRecord_PTR() *QNdefNfcTextRecord {
	return ptr
}

//QNdefNfcTextRecord::Encoding
type QNdefNfcTextRecord__Encoding int64

const (
	QNdefNfcTextRecord__Utf8  = QNdefNfcTextRecord__Encoding(0)
	QNdefNfcTextRecord__Utf16 = QNdefNfcTextRecord__Encoding(1)
)

func NewQNdefNfcTextRecord() *QNdefNfcTextRecord {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcTextRecord::QNdefNfcTextRecord")
		}
	}()

	return NewQNdefNfcTextRecordFromPointer(C.QNdefNfcTextRecord_NewQNdefNfcTextRecord())
}

func NewQNdefNfcTextRecord2(other QNdefRecord_ITF) *QNdefNfcTextRecord {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcTextRecord::QNdefNfcTextRecord")
		}
	}()

	return NewQNdefNfcTextRecordFromPointer(C.QNdefNfcTextRecord_NewQNdefNfcTextRecord2(PointerFromQNdefRecord(other)))
}

func (ptr *QNdefNfcTextRecord) Encoding() QNdefNfcTextRecord__Encoding {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcTextRecord::encoding")
		}
	}()

	if ptr.Pointer() != nil {
		return QNdefNfcTextRecord__Encoding(C.QNdefNfcTextRecord_Encoding(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNdefNfcTextRecord) Locale() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcTextRecord::locale")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QNdefNfcTextRecord_Locale(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNdefNfcTextRecord) SetEncoding(encoding QNdefNfcTextRecord__Encoding) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcTextRecord::setEncoding")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNdefNfcTextRecord_SetEncoding(ptr.Pointer(), C.int(encoding))
	}
}

func (ptr *QNdefNfcTextRecord) SetLocale(locale string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcTextRecord::setLocale")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNdefNfcTextRecord_SetLocale(ptr.Pointer(), C.CString(locale))
	}
}

func (ptr *QNdefNfcTextRecord) SetText(text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcTextRecord::setText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNdefNfcTextRecord_SetText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QNdefNfcTextRecord) Text() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcTextRecord::text")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QNdefNfcTextRecord_Text(ptr.Pointer()))
	}
	return ""
}
