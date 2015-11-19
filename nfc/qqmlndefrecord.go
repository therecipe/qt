package nfc

//#include "qqmlndefrecord.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QQmlNdefRecord struct {
	core.QObject
}

type QQmlNdefRecord_ITF interface {
	core.QObject_ITF
	QQmlNdefRecord_PTR() *QQmlNdefRecord
}

func PointerFromQQmlNdefRecord(ptr QQmlNdefRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlNdefRecord_PTR().Pointer()
	}
	return nil
}

func NewQQmlNdefRecordFromPointer(ptr unsafe.Pointer) *QQmlNdefRecord {
	var n = new(QQmlNdefRecord)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QQmlNdefRecord_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQmlNdefRecord) QQmlNdefRecord_PTR() *QQmlNdefRecord {
	return ptr
}

//QQmlNdefRecord::TypeNameFormat
type QQmlNdefRecord__TypeNameFormat int64

const (
	QQmlNdefRecord__Empty       = QQmlNdefRecord__TypeNameFormat(QNdefRecord__Empty)
	QQmlNdefRecord__NfcRtd      = QQmlNdefRecord__TypeNameFormat(QNdefRecord__NfcRtd)
	QQmlNdefRecord__Mime        = QQmlNdefRecord__TypeNameFormat(QNdefRecord__Mime)
	QQmlNdefRecord__Uri         = QQmlNdefRecord__TypeNameFormat(QNdefRecord__Uri)
	QQmlNdefRecord__ExternalRtd = QQmlNdefRecord__TypeNameFormat(QNdefRecord__ExternalRtd)
	QQmlNdefRecord__Unknown     = QQmlNdefRecord__TypeNameFormat(QNdefRecord__Unknown)
)

func (ptr *QQmlNdefRecord) TypeNameFormat() QQmlNdefRecord__TypeNameFormat {
	if ptr.Pointer() != nil {
		return QQmlNdefRecord__TypeNameFormat(C.QQmlNdefRecord_TypeNameFormat(ptr.Pointer()))
	}
	return 0
}

func NewQQmlNdefRecord(parent core.QObject_ITF) *QQmlNdefRecord {
	return NewQQmlNdefRecordFromPointer(C.QQmlNdefRecord_NewQQmlNdefRecord(core.PointerFromQObject(parent)))
}

func NewQQmlNdefRecord2(record QNdefRecord_ITF, parent core.QObject_ITF) *QQmlNdefRecord {
	return NewQQmlNdefRecordFromPointer(C.QQmlNdefRecord_NewQQmlNdefRecord2(PointerFromQNdefRecord(record), core.PointerFromQObject(parent)))
}

func (ptr *QQmlNdefRecord) ConnectRecordChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_ConnectRecordChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "recordChanged", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectRecordChanged() {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_DisconnectRecordChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "recordChanged")
	}
}

//export callbackQQmlNdefRecordRecordChanged
func callbackQQmlNdefRecordRecordChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "recordChanged").(func())()
}

func (ptr *QQmlNdefRecord) SetRecord(record QNdefRecord_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_SetRecord(ptr.Pointer(), PointerFromQNdefRecord(record))
	}
}

func (ptr *QQmlNdefRecord) SetType(newtype string) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_SetType(ptr.Pointer(), C.CString(newtype))
	}
}

func (ptr *QQmlNdefRecord) SetTypeNameFormat(newTypeNameFormat QQmlNdefRecord__TypeNameFormat) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_SetTypeNameFormat(ptr.Pointer(), C.int(newTypeNameFormat))
	}
}

func (ptr *QQmlNdefRecord) Type() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlNdefRecord_Type(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlNdefRecord) ConnectTypeChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_ConnectTypeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "typeChanged", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectTypeChanged() {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_DisconnectTypeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "typeChanged")
	}
}

//export callbackQQmlNdefRecordTypeChanged
func callbackQQmlNdefRecordTypeChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "typeChanged").(func())()
}

func (ptr *QQmlNdefRecord) ConnectTypeNameFormatChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_ConnectTypeNameFormatChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "typeNameFormatChanged", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectTypeNameFormatChanged() {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_DisconnectTypeNameFormatChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "typeNameFormatChanged")
	}
}

//export callbackQQmlNdefRecordTypeNameFormatChanged
func callbackQQmlNdefRecordTypeNameFormatChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "typeNameFormatChanged").(func())()
}
