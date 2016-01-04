package location

//#include "location.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGeoRouteReply struct {
	core.QObject
}

type QGeoRouteReply_ITF interface {
	core.QObject_ITF
	QGeoRouteReply_PTR() *QGeoRouteReply
}

func PointerFromQGeoRouteReply(ptr QGeoRouteReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoRouteReply_PTR().Pointer()
	}
	return nil
}

func NewQGeoRouteReplyFromPointer(ptr unsafe.Pointer) *QGeoRouteReply {
	var n = new(QGeoRouteReply)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGeoRouteReply_") {
		n.SetObjectName("QGeoRouteReply_" + qt.Identifier())
	}
	return n
}

func (ptr *QGeoRouteReply) QGeoRouteReply_PTR() *QGeoRouteReply {
	return ptr
}

//QGeoRouteReply::Error
type QGeoRouteReply__Error int64

const (
	QGeoRouteReply__NoError                = QGeoRouteReply__Error(0)
	QGeoRouteReply__EngineNotSetError      = QGeoRouteReply__Error(1)
	QGeoRouteReply__CommunicationError     = QGeoRouteReply__Error(2)
	QGeoRouteReply__ParseError             = QGeoRouteReply__Error(3)
	QGeoRouteReply__UnsupportedOptionError = QGeoRouteReply__Error(4)
	QGeoRouteReply__UnknownError           = QGeoRouteReply__Error(5)
)

func NewQGeoRouteReply(error QGeoRouteReply__Error, errorString string, parent core.QObject_ITF) *QGeoRouteReply {
	defer qt.Recovering("QGeoRouteReply::QGeoRouteReply")

	return NewQGeoRouteReplyFromPointer(C.QGeoRouteReply_NewQGeoRouteReply(C.int(error), C.CString(errorString), core.PointerFromQObject(parent)))
}

func (ptr *QGeoRouteReply) ConnectAbort(f func()) {
	defer qt.Recovering("connect QGeoRouteReply::abort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "abort", f)
	}
}

func (ptr *QGeoRouteReply) DisconnectAbort() {
	defer qt.Recovering("disconnect QGeoRouteReply::abort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "abort")
	}
}

//export callbackQGeoRouteReplyAbort
func callbackQGeoRouteReplyAbort(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGeoRouteReply::abort")

	if signal := qt.GetSignal(C.GoString(ptrName), "abort"); signal != nil {
		signal.(func())()
	} else {
		NewQGeoRouteReplyFromPointer(ptr).AbortDefault()
	}
}

func (ptr *QGeoRouteReply) Abort() {
	defer qt.Recovering("QGeoRouteReply::abort")

	if ptr.Pointer() != nil {
		C.QGeoRouteReply_Abort(ptr.Pointer())
	}
}

func (ptr *QGeoRouteReply) AbortDefault() {
	defer qt.Recovering("QGeoRouteReply::abort")

	if ptr.Pointer() != nil {
		C.QGeoRouteReply_AbortDefault(ptr.Pointer())
	}
}

func (ptr *QGeoRouteReply) ConnectError2(f func(error QGeoRouteReply__Error, errorString string)) {
	defer qt.Recovering("connect QGeoRouteReply::error")

	if ptr.Pointer() != nil {
		C.QGeoRouteReply_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QGeoRouteReply) DisconnectError2() {
	defer qt.Recovering("disconnect QGeoRouteReply::error")

	if ptr.Pointer() != nil {
		C.QGeoRouteReply_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQGeoRouteReplyError2
func callbackQGeoRouteReplyError2(ptr unsafe.Pointer, ptrName *C.char, error C.int, errorString *C.char) {
	defer qt.Recovering("callback QGeoRouteReply::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QGeoRouteReply__Error, string))(QGeoRouteReply__Error(error), C.GoString(errorString))
	}

}

func (ptr *QGeoRouteReply) Error2(error QGeoRouteReply__Error, errorString string) {
	defer qt.Recovering("QGeoRouteReply::error")

	if ptr.Pointer() != nil {
		C.QGeoRouteReply_Error2(ptr.Pointer(), C.int(error), C.CString(errorString))
	}
}

func (ptr *QGeoRouteReply) Error() QGeoRouteReply__Error {
	defer qt.Recovering("QGeoRouteReply::error")

	if ptr.Pointer() != nil {
		return QGeoRouteReply__Error(C.QGeoRouteReply_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRouteReply) ErrorString() string {
	defer qt.Recovering("QGeoRouteReply::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoRouteReply_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoRouteReply) ConnectFinished(f func()) {
	defer qt.Recovering("connect QGeoRouteReply::finished")

	if ptr.Pointer() != nil {
		C.QGeoRouteReply_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QGeoRouteReply) DisconnectFinished() {
	defer qt.Recovering("disconnect QGeoRouteReply::finished")

	if ptr.Pointer() != nil {
		C.QGeoRouteReply_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQGeoRouteReplyFinished
func callbackQGeoRouteReplyFinished(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGeoRouteReply::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGeoRouteReply) Finished() {
	defer qt.Recovering("QGeoRouteReply::finished")

	if ptr.Pointer() != nil {
		C.QGeoRouteReply_Finished(ptr.Pointer())
	}
}

func (ptr *QGeoRouteReply) IsFinished() bool {
	defer qt.Recovering("QGeoRouteReply::isFinished")

	if ptr.Pointer() != nil {
		return C.QGeoRouteReply_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGeoRouteReply) DestroyQGeoRouteReply() {
	defer qt.Recovering("QGeoRouteReply::~QGeoRouteReply")

	if ptr.Pointer() != nil {
		C.QGeoRouteReply_DestroyQGeoRouteReply(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoRouteReply) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGeoRouteReply::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGeoRouteReply) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGeoRouteReply::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGeoRouteReplyTimerEvent
func callbackQGeoRouteReplyTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGeoRouteReply::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGeoRouteReplyFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGeoRouteReply) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGeoRouteReply::timerEvent")

	if ptr.Pointer() != nil {
		C.QGeoRouteReply_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGeoRouteReply) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGeoRouteReply::timerEvent")

	if ptr.Pointer() != nil {
		C.QGeoRouteReply_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGeoRouteReply) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGeoRouteReply::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGeoRouteReply) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGeoRouteReply::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGeoRouteReplyChildEvent
func callbackQGeoRouteReplyChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGeoRouteReply::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGeoRouteReplyFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGeoRouteReply) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGeoRouteReply::childEvent")

	if ptr.Pointer() != nil {
		C.QGeoRouteReply_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGeoRouteReply) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGeoRouteReply::childEvent")

	if ptr.Pointer() != nil {
		C.QGeoRouteReply_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGeoRouteReply) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGeoRouteReply::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGeoRouteReply) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGeoRouteReply::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGeoRouteReplyCustomEvent
func callbackQGeoRouteReplyCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGeoRouteReply::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGeoRouteReplyFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGeoRouteReply) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGeoRouteReply::customEvent")

	if ptr.Pointer() != nil {
		C.QGeoRouteReply_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGeoRouteReply) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGeoRouteReply::customEvent")

	if ptr.Pointer() != nil {
		C.QGeoRouteReply_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
