package nfc

//#include "nfc.h"
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
	for len(n.ObjectName()) < len("QQmlNdefRecord_") {
		n.SetObjectName("QQmlNdefRecord_" + qt.Identifier())
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
	defer qt.Recovering("QQmlNdefRecord::typeNameFormat")

	if ptr.Pointer() != nil {
		return QQmlNdefRecord__TypeNameFormat(C.QQmlNdefRecord_TypeNameFormat(ptr.Pointer()))
	}
	return 0
}

func NewQQmlNdefRecord(parent core.QObject_ITF) *QQmlNdefRecord {
	defer qt.Recovering("QQmlNdefRecord::QQmlNdefRecord")

	return NewQQmlNdefRecordFromPointer(C.QQmlNdefRecord_NewQQmlNdefRecord(core.PointerFromQObject(parent)))
}

func NewQQmlNdefRecord2(record QNdefRecord_ITF, parent core.QObject_ITF) *QQmlNdefRecord {
	defer qt.Recovering("QQmlNdefRecord::QQmlNdefRecord")

	return NewQQmlNdefRecordFromPointer(C.QQmlNdefRecord_NewQQmlNdefRecord2(PointerFromQNdefRecord(record), core.PointerFromQObject(parent)))
}

func (ptr *QQmlNdefRecord) ConnectRecordChanged(f func()) {
	defer qt.Recovering("connect QQmlNdefRecord::recordChanged")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_ConnectRecordChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "recordChanged", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectRecordChanged() {
	defer qt.Recovering("disconnect QQmlNdefRecord::recordChanged")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_DisconnectRecordChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "recordChanged")
	}
}

//export callbackQQmlNdefRecordRecordChanged
func callbackQQmlNdefRecordRecordChanged(ptrName *C.char) {
	defer qt.Recovering("callback QQmlNdefRecord::recordChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "recordChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQmlNdefRecord) SetRecord(record QNdefRecord_ITF) {
	defer qt.Recovering("QQmlNdefRecord::setRecord")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_SetRecord(ptr.Pointer(), PointerFromQNdefRecord(record))
	}
}

func (ptr *QQmlNdefRecord) SetType(newtype string) {
	defer qt.Recovering("QQmlNdefRecord::setType")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_SetType(ptr.Pointer(), C.CString(newtype))
	}
}

func (ptr *QQmlNdefRecord) SetTypeNameFormat(newTypeNameFormat QQmlNdefRecord__TypeNameFormat) {
	defer qt.Recovering("QQmlNdefRecord::setTypeNameFormat")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_SetTypeNameFormat(ptr.Pointer(), C.int(newTypeNameFormat))
	}
}

func (ptr *QQmlNdefRecord) Type() string {
	defer qt.Recovering("QQmlNdefRecord::type")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlNdefRecord_Type(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlNdefRecord) ConnectTypeChanged(f func()) {
	defer qt.Recovering("connect QQmlNdefRecord::typeChanged")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_ConnectTypeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "typeChanged", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectTypeChanged() {
	defer qt.Recovering("disconnect QQmlNdefRecord::typeChanged")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_DisconnectTypeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "typeChanged")
	}
}

//export callbackQQmlNdefRecordTypeChanged
func callbackQQmlNdefRecordTypeChanged(ptrName *C.char) {
	defer qt.Recovering("callback QQmlNdefRecord::typeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "typeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQmlNdefRecord) ConnectTypeNameFormatChanged(f func()) {
	defer qt.Recovering("connect QQmlNdefRecord::typeNameFormatChanged")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_ConnectTypeNameFormatChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "typeNameFormatChanged", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectTypeNameFormatChanged() {
	defer qt.Recovering("disconnect QQmlNdefRecord::typeNameFormatChanged")

	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_DisconnectTypeNameFormatChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "typeNameFormatChanged")
	}
}

//export callbackQQmlNdefRecordTypeNameFormatChanged
func callbackQQmlNdefRecordTypeNameFormatChanged(ptrName *C.char) {
	defer qt.Recovering("callback QQmlNdefRecord::typeNameFormatChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "typeNameFormatChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQmlNdefRecord) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQmlNdefRecord::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQmlNdefRecord::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQmlNdefRecordTimerEvent
func callbackQQmlNdefRecordTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQmlNdefRecord::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQmlNdefRecord) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQmlNdefRecord::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQmlNdefRecord::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQmlNdefRecordChildEvent
func callbackQQmlNdefRecordChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQmlNdefRecord::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQmlNdefRecord) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQmlNdefRecord::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQmlNdefRecord) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQmlNdefRecord::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQmlNdefRecordCustomEvent
func callbackQQmlNdefRecordCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQmlNdefRecord::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
