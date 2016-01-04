package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDnsLookup struct {
	core.QObject
}

type QDnsLookup_ITF interface {
	core.QObject_ITF
	QDnsLookup_PTR() *QDnsLookup
}

func PointerFromQDnsLookup(ptr QDnsLookup_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsLookup_PTR().Pointer()
	}
	return nil
}

func NewQDnsLookupFromPointer(ptr unsafe.Pointer) *QDnsLookup {
	var n = new(QDnsLookup)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDnsLookup_") {
		n.SetObjectName("QDnsLookup_" + qt.Identifier())
	}
	return n
}

func (ptr *QDnsLookup) QDnsLookup_PTR() *QDnsLookup {
	return ptr
}

//QDnsLookup::Error
type QDnsLookup__Error int64

const (
	QDnsLookup__NoError                 = QDnsLookup__Error(0)
	QDnsLookup__ResolverError           = QDnsLookup__Error(1)
	QDnsLookup__OperationCancelledError = QDnsLookup__Error(2)
	QDnsLookup__InvalidRequestError     = QDnsLookup__Error(3)
	QDnsLookup__InvalidReplyError       = QDnsLookup__Error(4)
	QDnsLookup__ServerFailureError      = QDnsLookup__Error(5)
	QDnsLookup__ServerRefusedError      = QDnsLookup__Error(6)
	QDnsLookup__NotFoundError           = QDnsLookup__Error(7)
)

//QDnsLookup::Type
type QDnsLookup__Type int64

const (
	QDnsLookup__A     = QDnsLookup__Type(1)
	QDnsLookup__AAAA  = QDnsLookup__Type(28)
	QDnsLookup__ANY   = QDnsLookup__Type(255)
	QDnsLookup__CNAME = QDnsLookup__Type(5)
	QDnsLookup__MX    = QDnsLookup__Type(15)
	QDnsLookup__NS    = QDnsLookup__Type(2)
	QDnsLookup__PTR   = QDnsLookup__Type(12)
	QDnsLookup__SRV   = QDnsLookup__Type(33)
	QDnsLookup__TXT   = QDnsLookup__Type(16)
)

func NewQDnsLookup3(ty QDnsLookup__Type, name string, nameserver QHostAddress_ITF, parent core.QObject_ITF) *QDnsLookup {
	defer qt.Recovering("QDnsLookup::QDnsLookup")

	return NewQDnsLookupFromPointer(C.QDnsLookup_NewQDnsLookup3(C.int(ty), C.CString(name), PointerFromQHostAddress(nameserver), core.PointerFromQObject(parent)))
}

func (ptr *QDnsLookup) Error() QDnsLookup__Error {
	defer qt.Recovering("QDnsLookup::error")

	if ptr.Pointer() != nil {
		return QDnsLookup__Error(C.QDnsLookup_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDnsLookup) ErrorString() string {
	defer qt.Recovering("QDnsLookup::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsLookup_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsLookup) Name() string {
	defer qt.Recovering("QDnsLookup::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsLookup_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsLookup) SetName(name string) {
	defer qt.Recovering("QDnsLookup::setName")

	if ptr.Pointer() != nil {
		C.QDnsLookup_SetName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QDnsLookup) SetNameserver(nameserver QHostAddress_ITF) {
	defer qt.Recovering("QDnsLookup::setNameserver")

	if ptr.Pointer() != nil {
		C.QDnsLookup_SetNameserver(ptr.Pointer(), PointerFromQHostAddress(nameserver))
	}
}

func (ptr *QDnsLookup) SetType(v QDnsLookup__Type) {
	defer qt.Recovering("QDnsLookup::setType")

	if ptr.Pointer() != nil {
		C.QDnsLookup_SetType(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QDnsLookup) Type() QDnsLookup__Type {
	defer qt.Recovering("QDnsLookup::type")

	if ptr.Pointer() != nil {
		return QDnsLookup__Type(C.QDnsLookup_Type(ptr.Pointer()))
	}
	return 0
}

func NewQDnsLookup(parent core.QObject_ITF) *QDnsLookup {
	defer qt.Recovering("QDnsLookup::QDnsLookup")

	return NewQDnsLookupFromPointer(C.QDnsLookup_NewQDnsLookup(core.PointerFromQObject(parent)))
}

func NewQDnsLookup2(ty QDnsLookup__Type, name string, parent core.QObject_ITF) *QDnsLookup {
	defer qt.Recovering("QDnsLookup::QDnsLookup")

	return NewQDnsLookupFromPointer(C.QDnsLookup_NewQDnsLookup2(C.int(ty), C.CString(name), core.PointerFromQObject(parent)))
}

func (ptr *QDnsLookup) Abort() {
	defer qt.Recovering("QDnsLookup::abort")

	if ptr.Pointer() != nil {
		C.QDnsLookup_Abort(ptr.Pointer())
	}
}

func (ptr *QDnsLookup) ConnectFinished(f func()) {
	defer qt.Recovering("connect QDnsLookup::finished")

	if ptr.Pointer() != nil {
		C.QDnsLookup_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QDnsLookup) DisconnectFinished() {
	defer qt.Recovering("disconnect QDnsLookup::finished")

	if ptr.Pointer() != nil {
		C.QDnsLookup_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQDnsLookupFinished
func callbackQDnsLookupFinished(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QDnsLookup::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QDnsLookup) Finished() {
	defer qt.Recovering("QDnsLookup::finished")

	if ptr.Pointer() != nil {
		C.QDnsLookup_Finished(ptr.Pointer())
	}
}

func (ptr *QDnsLookup) IsFinished() bool {
	defer qt.Recovering("QDnsLookup::isFinished")

	if ptr.Pointer() != nil {
		return C.QDnsLookup_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDnsLookup) Lookup() {
	defer qt.Recovering("QDnsLookup::lookup")

	if ptr.Pointer() != nil {
		C.QDnsLookup_Lookup(ptr.Pointer())
	}
}

func (ptr *QDnsLookup) ConnectNameChanged(f func(name string)) {
	defer qt.Recovering("connect QDnsLookup::nameChanged")

	if ptr.Pointer() != nil {
		C.QDnsLookup_ConnectNameChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "nameChanged", f)
	}
}

func (ptr *QDnsLookup) DisconnectNameChanged() {
	defer qt.Recovering("disconnect QDnsLookup::nameChanged")

	if ptr.Pointer() != nil {
		C.QDnsLookup_DisconnectNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "nameChanged")
	}
}

//export callbackQDnsLookupNameChanged
func callbackQDnsLookupNameChanged(ptr unsafe.Pointer, ptrName *C.char, name *C.char) {
	defer qt.Recovering("callback QDnsLookup::nameChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "nameChanged"); signal != nil {
		signal.(func(string))(C.GoString(name))
	}

}

func (ptr *QDnsLookup) NameChanged(name string) {
	defer qt.Recovering("QDnsLookup::nameChanged")

	if ptr.Pointer() != nil {
		C.QDnsLookup_NameChanged(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QDnsLookup) ConnectTypeChanged(f func(ty QDnsLookup__Type)) {
	defer qt.Recovering("connect QDnsLookup::typeChanged")

	if ptr.Pointer() != nil {
		C.QDnsLookup_ConnectTypeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "typeChanged", f)
	}
}

func (ptr *QDnsLookup) DisconnectTypeChanged() {
	defer qt.Recovering("disconnect QDnsLookup::typeChanged")

	if ptr.Pointer() != nil {
		C.QDnsLookup_DisconnectTypeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "typeChanged")
	}
}

//export callbackQDnsLookupTypeChanged
func callbackQDnsLookupTypeChanged(ptr unsafe.Pointer, ptrName *C.char, ty C.int) {
	defer qt.Recovering("callback QDnsLookup::typeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "typeChanged"); signal != nil {
		signal.(func(QDnsLookup__Type))(QDnsLookup__Type(ty))
	}

}

func (ptr *QDnsLookup) TypeChanged(ty QDnsLookup__Type) {
	defer qt.Recovering("QDnsLookup::typeChanged")

	if ptr.Pointer() != nil {
		C.QDnsLookup_TypeChanged(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QDnsLookup) DestroyQDnsLookup() {
	defer qt.Recovering("QDnsLookup::~QDnsLookup")

	if ptr.Pointer() != nil {
		C.QDnsLookup_DestroyQDnsLookup(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDnsLookup) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDnsLookup::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDnsLookup) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDnsLookup::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDnsLookupTimerEvent
func callbackQDnsLookupTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDnsLookup::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDnsLookupFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDnsLookup) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDnsLookup::timerEvent")

	if ptr.Pointer() != nil {
		C.QDnsLookup_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDnsLookup) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDnsLookup::timerEvent")

	if ptr.Pointer() != nil {
		C.QDnsLookup_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDnsLookup) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDnsLookup::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDnsLookup) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDnsLookup::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDnsLookupChildEvent
func callbackQDnsLookupChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDnsLookup::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDnsLookupFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDnsLookup) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDnsLookup::childEvent")

	if ptr.Pointer() != nil {
		C.QDnsLookup_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDnsLookup) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDnsLookup::childEvent")

	if ptr.Pointer() != nil {
		C.QDnsLookup_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDnsLookup) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDnsLookup::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDnsLookup) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDnsLookup::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDnsLookupCustomEvent
func callbackQDnsLookupCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDnsLookup::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDnsLookupFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDnsLookup) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDnsLookup::customEvent")

	if ptr.Pointer() != nil {
		C.QDnsLookup_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDnsLookup) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDnsLookup::customEvent")

	if ptr.Pointer() != nil {
		C.QDnsLookup_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
