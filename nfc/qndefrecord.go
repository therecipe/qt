package nfc

//#include "nfc.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QNdefRecord struct {
	ptr unsafe.Pointer
}

type QNdefRecord_ITF interface {
	QNdefRecord_PTR() *QNdefRecord
}

func (p *QNdefRecord) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNdefRecord) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQNdefRecord(ptr QNdefRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefRecord_PTR().Pointer()
	}
	return nil
}

func NewQNdefRecordFromPointer(ptr unsafe.Pointer) *QNdefRecord {
	var n = new(QNdefRecord)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNdefRecord) QNdefRecord_PTR() *QNdefRecord {
	return ptr
}

//QNdefRecord::TypeNameFormat
type QNdefRecord__TypeNameFormat int64

const (
	QNdefRecord__Empty       = QNdefRecord__TypeNameFormat(0x00)
	QNdefRecord__NfcRtd      = QNdefRecord__TypeNameFormat(0x01)
	QNdefRecord__Mime        = QNdefRecord__TypeNameFormat(0x02)
	QNdefRecord__Uri         = QNdefRecord__TypeNameFormat(0x03)
	QNdefRecord__ExternalRtd = QNdefRecord__TypeNameFormat(0x04)
	QNdefRecord__Unknown     = QNdefRecord__TypeNameFormat(0x05)
)

func NewQNdefRecord() *QNdefRecord {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefRecord::QNdefRecord")
		}
	}()

	return NewQNdefRecordFromPointer(C.QNdefRecord_NewQNdefRecord())
}

func NewQNdefRecord2(other QNdefRecord_ITF) *QNdefRecord {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefRecord::QNdefRecord")
		}
	}()

	return NewQNdefRecordFromPointer(C.QNdefRecord_NewQNdefRecord2(PointerFromQNdefRecord(other)))
}

func (ptr *QNdefRecord) Id() *core.QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefRecord::id")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QNdefRecord_Id(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNdefRecord) IsEmpty() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefRecord::isEmpty")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QNdefRecord_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNdefRecord) Payload() *core.QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefRecord::payload")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QNdefRecord_Payload(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNdefRecord) SetId(id core.QByteArray_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefRecord::setId")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNdefRecord_SetId(ptr.Pointer(), core.PointerFromQByteArray(id))
	}
}

func (ptr *QNdefRecord) SetPayload(payload core.QByteArray_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefRecord::setPayload")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNdefRecord_SetPayload(ptr.Pointer(), core.PointerFromQByteArray(payload))
	}
}

func (ptr *QNdefRecord) SetType(ty core.QByteArray_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefRecord::setType")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNdefRecord_SetType(ptr.Pointer(), core.PointerFromQByteArray(ty))
	}
}

func (ptr *QNdefRecord) SetTypeNameFormat(typeNameFormat QNdefRecord__TypeNameFormat) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefRecord::setTypeNameFormat")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNdefRecord_SetTypeNameFormat(ptr.Pointer(), C.int(typeNameFormat))
	}
}

func (ptr *QNdefRecord) Type() *core.QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefRecord::type")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QNdefRecord_Type(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNdefRecord) TypeNameFormat() QNdefRecord__TypeNameFormat {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefRecord::typeNameFormat")
		}
	}()

	if ptr.Pointer() != nil {
		return QNdefRecord__TypeNameFormat(C.QNdefRecord_TypeNameFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNdefRecord) DestroyQNdefRecord() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefRecord::~QNdefRecord")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNdefRecord_DestroyQNdefRecord(ptr.Pointer())
	}
}
