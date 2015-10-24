package nfc

//#include "qndefrecord.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QNdefRecord struct {
	ptr unsafe.Pointer
}

type QNdefRecordITF interface {
	QNdefRecordPTR() *QNdefRecord
}

func (p *QNdefRecord) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNdefRecord) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQNdefRecord(ptr QNdefRecordITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefRecordPTR().Pointer()
	}
	return nil
}

func QNdefRecordFromPointer(ptr unsafe.Pointer) *QNdefRecord {
	var n = new(QNdefRecord)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNdefRecord) QNdefRecordPTR() *QNdefRecord {
	return ptr
}

//QNdefRecord::TypeNameFormat
type QNdefRecord__TypeNameFormat int

var (
	QNdefRecord__Empty       = QNdefRecord__TypeNameFormat(0x00)
	QNdefRecord__NfcRtd      = QNdefRecord__TypeNameFormat(0x01)
	QNdefRecord__Mime        = QNdefRecord__TypeNameFormat(0x02)
	QNdefRecord__Uri         = QNdefRecord__TypeNameFormat(0x03)
	QNdefRecord__ExternalRtd = QNdefRecord__TypeNameFormat(0x04)
	QNdefRecord__Unknown     = QNdefRecord__TypeNameFormat(0x05)
)

func NewQNdefRecord() *QNdefRecord {
	return QNdefRecordFromPointer(unsafe.Pointer(C.QNdefRecord_NewQNdefRecord()))
}

func NewQNdefRecord2(other QNdefRecordITF) *QNdefRecord {
	return QNdefRecordFromPointer(unsafe.Pointer(C.QNdefRecord_NewQNdefRecord2(C.QtObjectPtr(PointerFromQNdefRecord(other)))))
}

func (ptr *QNdefRecord) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QNdefRecord_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNdefRecord) SetId(id core.QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QNdefRecord_SetId(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(id)))
	}
}

func (ptr *QNdefRecord) SetPayload(payload core.QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QNdefRecord_SetPayload(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(payload)))
	}
}

func (ptr *QNdefRecord) SetType(ty core.QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QNdefRecord_SetType(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(ty)))
	}
}

func (ptr *QNdefRecord) SetTypeNameFormat(typeNameFormat QNdefRecord__TypeNameFormat) {
	if ptr.Pointer() != nil {
		C.QNdefRecord_SetTypeNameFormat(C.QtObjectPtr(ptr.Pointer()), C.int(typeNameFormat))
	}
}

func (ptr *QNdefRecord) TypeNameFormat() QNdefRecord__TypeNameFormat {
	if ptr.Pointer() != nil {
		return QNdefRecord__TypeNameFormat(C.QNdefRecord_TypeNameFormat(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNdefRecord) DestroyQNdefRecord() {
	if ptr.Pointer() != nil {
		C.QNdefRecord_DestroyQNdefRecord(C.QtObjectPtr(ptr.Pointer()))
	}
}
