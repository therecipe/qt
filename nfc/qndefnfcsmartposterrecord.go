package nfc

//#include "qndefnfcsmartposterrecord.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QNdefNfcSmartPosterRecord struct {
	QNdefRecord
}

type QNdefNfcSmartPosterRecordITF interface {
	QNdefRecordITF
	QNdefNfcSmartPosterRecordPTR() *QNdefNfcSmartPosterRecord
}

func PointerFromQNdefNfcSmartPosterRecord(ptr QNdefNfcSmartPosterRecordITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefNfcSmartPosterRecordPTR().Pointer()
	}
	return nil
}

func QNdefNfcSmartPosterRecordFromPointer(ptr unsafe.Pointer) *QNdefNfcSmartPosterRecord {
	var n = new(QNdefNfcSmartPosterRecord)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNdefNfcSmartPosterRecord) QNdefNfcSmartPosterRecordPTR() *QNdefNfcSmartPosterRecord {
	return ptr
}

//QNdefNfcSmartPosterRecord::Action
type QNdefNfcSmartPosterRecord__Action int

var (
	QNdefNfcSmartPosterRecord__UnspecifiedAction = QNdefNfcSmartPosterRecord__Action(-1)
	QNdefNfcSmartPosterRecord__DoAction          = QNdefNfcSmartPosterRecord__Action(0)
	QNdefNfcSmartPosterRecord__SaveAction        = QNdefNfcSmartPosterRecord__Action(1)
	QNdefNfcSmartPosterRecord__EditAction        = QNdefNfcSmartPosterRecord__Action(2)
)

func NewQNdefNfcSmartPosterRecord() *QNdefNfcSmartPosterRecord {
	return QNdefNfcSmartPosterRecordFromPointer(unsafe.Pointer(C.QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord()))
}

func NewQNdefNfcSmartPosterRecord3(other QNdefNfcSmartPosterRecordITF) *QNdefNfcSmartPosterRecord {
	return QNdefNfcSmartPosterRecordFromPointer(unsafe.Pointer(C.QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord3(C.QtObjectPtr(PointerFromQNdefNfcSmartPosterRecord(other)))))
}

func NewQNdefNfcSmartPosterRecord2(other QNdefRecordITF) *QNdefNfcSmartPosterRecord {
	return QNdefNfcSmartPosterRecordFromPointer(unsafe.Pointer(C.QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord2(C.QtObjectPtr(PointerFromQNdefRecord(other)))))
}

func (ptr *QNdefNfcSmartPosterRecord) Action() QNdefNfcSmartPosterRecord__Action {
	if ptr.Pointer() != nil {
		return QNdefNfcSmartPosterRecord__Action(C.QNdefNfcSmartPosterRecord_Action(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNdefNfcSmartPosterRecord) AddIcon2(ty core.QByteArrayITF, data core.QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_AddIcon2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(ty)), C.QtObjectPtr(core.PointerFromQByteArray(data)))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) AddTitle(text QNdefNfcTextRecordITF) bool {
	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_AddTitle(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNdefNfcTextRecord(text))) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) AddTitle2(text string, locale string, encoding QNdefNfcTextRecord__Encoding) bool {
	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_AddTitle2(C.QtObjectPtr(ptr.Pointer()), C.CString(text), C.CString(locale), C.int(encoding)) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) HasAction() bool {
	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_HasAction(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) HasIcon(mimetype core.QByteArrayITF) bool {
	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_HasIcon(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(mimetype))) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) HasSize() bool {
	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_HasSize(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) HasTitle(locale string) bool {
	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_HasTitle(C.QtObjectPtr(ptr.Pointer()), C.CString(locale)) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) HasTypeInfo() bool {
	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_HasTypeInfo(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) IconCount() int {
	if ptr.Pointer() != nil {
		return int(C.QNdefNfcSmartPosterRecord_IconCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNdefNfcSmartPosterRecord) RemoveIcon2(ty core.QByteArrayITF) bool {
	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_RemoveIcon2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(ty))) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) RemoveTitle(text QNdefNfcTextRecordITF) bool {
	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_RemoveTitle(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNdefNfcTextRecord(text))) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) RemoveTitle2(locale string) bool {
	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_RemoveTitle2(C.QtObjectPtr(ptr.Pointer()), C.CString(locale)) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) SetAction(act QNdefNfcSmartPosterRecord__Action) {
	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_SetAction(C.QtObjectPtr(ptr.Pointer()), C.int(act))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) SetTypeInfo(ty core.QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_SetTypeInfo(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(ty)))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) SetUri(url QNdefNfcUriRecordITF) {
	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_SetUri(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNdefNfcUriRecord(url)))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) SetUri2(url string) {
	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_SetUri2(C.QtObjectPtr(ptr.Pointer()), C.CString(url))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) Title(locale string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNdefNfcSmartPosterRecord_Title(C.QtObjectPtr(ptr.Pointer()), C.CString(locale)))
	}
	return ""
}

func (ptr *QNdefNfcSmartPosterRecord) TitleCount() int {
	if ptr.Pointer() != nil {
		return int(C.QNdefNfcSmartPosterRecord_TitleCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNdefNfcSmartPosterRecord) Uri() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNdefNfcSmartPosterRecord_Uri(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QNdefNfcSmartPosterRecord) DestroyQNdefNfcSmartPosterRecord() {
	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_DestroyQNdefNfcSmartPosterRecord(C.QtObjectPtr(ptr.Pointer()))
	}
}
