package nfc

//#include "nfc.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QNdefNfcUriRecord struct {
	QNdefRecord
}

type QNdefNfcUriRecord_ITF interface {
	QNdefRecord_ITF
	QNdefNfcUriRecord_PTR() *QNdefNfcUriRecord
}

func PointerFromQNdefNfcUriRecord(ptr QNdefNfcUriRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefNfcUriRecord_PTR().Pointer()
	}
	return nil
}

func NewQNdefNfcUriRecordFromPointer(ptr unsafe.Pointer) *QNdefNfcUriRecord {
	var n = new(QNdefNfcUriRecord)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNdefNfcUriRecord) QNdefNfcUriRecord_PTR() *QNdefNfcUriRecord {
	return ptr
}

func NewQNdefNfcUriRecord() *QNdefNfcUriRecord {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcUriRecord::QNdefNfcUriRecord")
		}
	}()

	return NewQNdefNfcUriRecordFromPointer(C.QNdefNfcUriRecord_NewQNdefNfcUriRecord())
}

func NewQNdefNfcUriRecord2(other QNdefRecord_ITF) *QNdefNfcUriRecord {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcUriRecord::QNdefNfcUriRecord")
		}
	}()

	return NewQNdefNfcUriRecordFromPointer(C.QNdefNfcUriRecord_NewQNdefNfcUriRecord2(PointerFromQNdefRecord(other)))
}

func (ptr *QNdefNfcUriRecord) SetUri(uri core.QUrl_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcUriRecord::setUri")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNdefNfcUriRecord_SetUri(ptr.Pointer(), core.PointerFromQUrl(uri))
	}
}
