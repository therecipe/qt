package nfc

//#include "nfc.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QNdefNfcSmartPosterRecord struct {
	QNdefRecord
}

type QNdefNfcSmartPosterRecord_ITF interface {
	QNdefRecord_ITF
	QNdefNfcSmartPosterRecord_PTR() *QNdefNfcSmartPosterRecord
}

func PointerFromQNdefNfcSmartPosterRecord(ptr QNdefNfcSmartPosterRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefNfcSmartPosterRecord_PTR().Pointer()
	}
	return nil
}

func NewQNdefNfcSmartPosterRecordFromPointer(ptr unsafe.Pointer) *QNdefNfcSmartPosterRecord {
	var n = new(QNdefNfcSmartPosterRecord)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNdefNfcSmartPosterRecord) QNdefNfcSmartPosterRecord_PTR() *QNdefNfcSmartPosterRecord {
	return ptr
}

//QNdefNfcSmartPosterRecord::Action
type QNdefNfcSmartPosterRecord__Action int64

const (
	QNdefNfcSmartPosterRecord__UnspecifiedAction = QNdefNfcSmartPosterRecord__Action(-1)
	QNdefNfcSmartPosterRecord__DoAction          = QNdefNfcSmartPosterRecord__Action(0)
	QNdefNfcSmartPosterRecord__SaveAction        = QNdefNfcSmartPosterRecord__Action(1)
	QNdefNfcSmartPosterRecord__EditAction        = QNdefNfcSmartPosterRecord__Action(2)
)

func NewQNdefNfcSmartPosterRecord() *QNdefNfcSmartPosterRecord {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcSmartPosterRecord::QNdefNfcSmartPosterRecord")
		}
	}()

	return NewQNdefNfcSmartPosterRecordFromPointer(C.QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord())
}

func NewQNdefNfcSmartPosterRecord3(other QNdefNfcSmartPosterRecord_ITF) *QNdefNfcSmartPosterRecord {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcSmartPosterRecord::QNdefNfcSmartPosterRecord")
		}
	}()

	return NewQNdefNfcSmartPosterRecordFromPointer(C.QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord3(PointerFromQNdefNfcSmartPosterRecord(other)))
}

func NewQNdefNfcSmartPosterRecord2(other QNdefRecord_ITF) *QNdefNfcSmartPosterRecord {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcSmartPosterRecord::QNdefNfcSmartPosterRecord")
		}
	}()

	return NewQNdefNfcSmartPosterRecordFromPointer(C.QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord2(PointerFromQNdefRecord(other)))
}

func (ptr *QNdefNfcSmartPosterRecord) Action() QNdefNfcSmartPosterRecord__Action {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcSmartPosterRecord::action")
		}
	}()

	if ptr.Pointer() != nil {
		return QNdefNfcSmartPosterRecord__Action(C.QNdefNfcSmartPosterRecord_Action(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNdefNfcSmartPosterRecord) AddIcon2(ty core.QByteArray_ITF, data core.QByteArray_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcSmartPosterRecord::addIcon")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_AddIcon2(ptr.Pointer(), core.PointerFromQByteArray(ty), core.PointerFromQByteArray(data))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) AddTitle(text QNdefNfcTextRecord_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcSmartPosterRecord::addTitle")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_AddTitle(ptr.Pointer(), PointerFromQNdefNfcTextRecord(text)) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) AddTitle2(text string, locale string, encoding QNdefNfcTextRecord__Encoding) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcSmartPosterRecord::addTitle")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_AddTitle2(ptr.Pointer(), C.CString(text), C.CString(locale), C.int(encoding)) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) HasAction() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcSmartPosterRecord::hasAction")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_HasAction(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) HasIcon(mimetype core.QByteArray_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcSmartPosterRecord::hasIcon")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_HasIcon(ptr.Pointer(), core.PointerFromQByteArray(mimetype)) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) HasSize() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcSmartPosterRecord::hasSize")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_HasSize(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) HasTitle(locale string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcSmartPosterRecord::hasTitle")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_HasTitle(ptr.Pointer(), C.CString(locale)) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) HasTypeInfo() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcSmartPosterRecord::hasTypeInfo")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_HasTypeInfo(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) Icon(mimetype core.QByteArray_ITF) *core.QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcSmartPosterRecord::icon")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QNdefNfcSmartPosterRecord_Icon(ptr.Pointer(), core.PointerFromQByteArray(mimetype)))
	}
	return nil
}

func (ptr *QNdefNfcSmartPosterRecord) IconCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcSmartPosterRecord::iconCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QNdefNfcSmartPosterRecord_IconCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNdefNfcSmartPosterRecord) RemoveIcon2(ty core.QByteArray_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcSmartPosterRecord::removeIcon")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_RemoveIcon2(ptr.Pointer(), core.PointerFromQByteArray(ty)) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) RemoveTitle(text QNdefNfcTextRecord_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcSmartPosterRecord::removeTitle")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_RemoveTitle(ptr.Pointer(), PointerFromQNdefNfcTextRecord(text)) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) RemoveTitle2(locale string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcSmartPosterRecord::removeTitle")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QNdefNfcSmartPosterRecord_RemoveTitle2(ptr.Pointer(), C.CString(locale)) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) SetAction(act QNdefNfcSmartPosterRecord__Action) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcSmartPosterRecord::setAction")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_SetAction(ptr.Pointer(), C.int(act))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) SetTypeInfo(ty core.QByteArray_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcSmartPosterRecord::setTypeInfo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_SetTypeInfo(ptr.Pointer(), core.PointerFromQByteArray(ty))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) SetUri(url QNdefNfcUriRecord_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcSmartPosterRecord::setUri")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_SetUri(ptr.Pointer(), PointerFromQNdefNfcUriRecord(url))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) SetUri2(url core.QUrl_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcSmartPosterRecord::setUri")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_SetUri2(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) Title(locale string) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcSmartPosterRecord::title")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QNdefNfcSmartPosterRecord_Title(ptr.Pointer(), C.CString(locale)))
	}
	return ""
}

func (ptr *QNdefNfcSmartPosterRecord) TitleCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcSmartPosterRecord::titleCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QNdefNfcSmartPosterRecord_TitleCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNdefNfcSmartPosterRecord) TypeInfo() *core.QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcSmartPosterRecord::typeInfo")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QNdefNfcSmartPosterRecord_TypeInfo(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNdefNfcSmartPosterRecord) DestroyQNdefNfcSmartPosterRecord() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNdefNfcSmartPosterRecord::~QNdefNfcSmartPosterRecord")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_DestroyQNdefNfcSmartPosterRecord(ptr.Pointer())
	}
}
