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

type QQmlNdefRecordITF interface {
	core.QObjectITF
	QQmlNdefRecordPTR() *QQmlNdefRecord
}

func PointerFromQQmlNdefRecord(ptr QQmlNdefRecordITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlNdefRecordPTR().Pointer()
	}
	return nil
}

func QQmlNdefRecordFromPointer(ptr unsafe.Pointer) *QQmlNdefRecord {
	var n = new(QQmlNdefRecord)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QQmlNdefRecord_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQmlNdefRecord) QQmlNdefRecordPTR() *QQmlNdefRecord {
	return ptr
}

//QQmlNdefRecord::TypeNameFormat
type QQmlNdefRecord__TypeNameFormat int

var (
	QQmlNdefRecord__Empty       = QQmlNdefRecord__TypeNameFormat(QNdefRecord__Empty)
	QQmlNdefRecord__NfcRtd      = QQmlNdefRecord__TypeNameFormat(QNdefRecord__NfcRtd)
	QQmlNdefRecord__Mime        = QQmlNdefRecord__TypeNameFormat(QNdefRecord__Mime)
	QQmlNdefRecord__Uri         = QQmlNdefRecord__TypeNameFormat(QNdefRecord__Uri)
	QQmlNdefRecord__ExternalRtd = QQmlNdefRecord__TypeNameFormat(QNdefRecord__ExternalRtd)
	QQmlNdefRecord__Unknown     = QQmlNdefRecord__TypeNameFormat(QNdefRecord__Unknown)
)

func (ptr *QQmlNdefRecord) TypeNameFormat() QQmlNdefRecord__TypeNameFormat {
	if ptr.Pointer() != nil {
		return QQmlNdefRecord__TypeNameFormat(C.QQmlNdefRecord_TypeNameFormat(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQQmlNdefRecord(parent core.QObjectITF) *QQmlNdefRecord {
	return QQmlNdefRecordFromPointer(unsafe.Pointer(C.QQmlNdefRecord_NewQQmlNdefRecord(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQQmlNdefRecord2(record QNdefRecordITF, parent core.QObjectITF) *QQmlNdefRecord {
	return QQmlNdefRecordFromPointer(unsafe.Pointer(C.QQmlNdefRecord_NewQQmlNdefRecord2(C.QtObjectPtr(PointerFromQNdefRecord(record)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QQmlNdefRecord) ConnectRecordChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_ConnectRecordChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "recordChanged", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectRecordChanged() {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_DisconnectRecordChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "recordChanged")
	}
}

//export callbackQQmlNdefRecordRecordChanged
func callbackQQmlNdefRecordRecordChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "recordChanged").(func())()
}

func (ptr *QQmlNdefRecord) SetRecord(record QNdefRecordITF) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_SetRecord(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNdefRecord(record)))
	}
}

func (ptr *QQmlNdefRecord) SetType(newtype string) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_SetType(C.QtObjectPtr(ptr.Pointer()), C.CString(newtype))
	}
}

func (ptr *QQmlNdefRecord) SetTypeNameFormat(newTypeNameFormat QQmlNdefRecord__TypeNameFormat) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_SetTypeNameFormat(C.QtObjectPtr(ptr.Pointer()), C.int(newTypeNameFormat))
	}
}

func (ptr *QQmlNdefRecord) Type() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlNdefRecord_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QQmlNdefRecord) ConnectTypeChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_ConnectTypeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "typeChanged", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectTypeChanged() {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_DisconnectTypeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "typeChanged")
	}
}

//export callbackQQmlNdefRecordTypeChanged
func callbackQQmlNdefRecordTypeChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "typeChanged").(func())()
}

func (ptr *QQmlNdefRecord) ConnectTypeNameFormatChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_ConnectTypeNameFormatChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "typeNameFormatChanged", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectTypeNameFormatChanged() {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_DisconnectTypeNameFormatChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "typeNameFormatChanged")
	}
}

//export callbackQQmlNdefRecordTypeNameFormatChanged
func callbackQQmlNdefRecordTypeNameFormatChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "typeNameFormatChanged").(func())()
}
