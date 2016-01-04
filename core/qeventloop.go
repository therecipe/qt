package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QEventLoop struct {
	QObject
}

type QEventLoop_ITF interface {
	QObject_ITF
	QEventLoop_PTR() *QEventLoop
}

func PointerFromQEventLoop(ptr QEventLoop_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QEventLoop_PTR().Pointer()
	}
	return nil
}

func NewQEventLoopFromPointer(ptr unsafe.Pointer) *QEventLoop {
	var n = new(QEventLoop)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QEventLoop_") {
		n.SetObjectName("QEventLoop_" + qt.Identifier())
	}
	return n
}

func (ptr *QEventLoop) QEventLoop_PTR() *QEventLoop {
	return ptr
}

//QEventLoop::ProcessEventsFlag
type QEventLoop__ProcessEventsFlag int64

const (
	QEventLoop__AllEvents              = QEventLoop__ProcessEventsFlag(0x00)
	QEventLoop__ExcludeUserInputEvents = QEventLoop__ProcessEventsFlag(0x01)
	QEventLoop__ExcludeSocketNotifiers = QEventLoop__ProcessEventsFlag(0x02)
	QEventLoop__WaitForMoreEvents      = QEventLoop__ProcessEventsFlag(0x04)
	QEventLoop__X11ExcludeTimers       = QEventLoop__ProcessEventsFlag(0x08)
	QEventLoop__EventLoopExec          = QEventLoop__ProcessEventsFlag(0x20)
	QEventLoop__DialogExec             = QEventLoop__ProcessEventsFlag(0x40)
)

func NewQEventLoop(parent QObject_ITF) *QEventLoop {
	defer qt.Recovering("QEventLoop::QEventLoop")

	return NewQEventLoopFromPointer(C.QEventLoop_NewQEventLoop(PointerFromQObject(parent)))
}

func (ptr *QEventLoop) Event(event QEvent_ITF) bool {
	defer qt.Recovering("QEventLoop::event")

	if ptr.Pointer() != nil {
		return C.QEventLoop_Event(ptr.Pointer(), PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QEventLoop) Exec(flags QEventLoop__ProcessEventsFlag) int {
	defer qt.Recovering("QEventLoop::exec")

	if ptr.Pointer() != nil {
		return int(C.QEventLoop_Exec(ptr.Pointer(), C.int(flags)))
	}
	return 0
}

func (ptr *QEventLoop) Exit(returnCode int) {
	defer qt.Recovering("QEventLoop::exit")

	if ptr.Pointer() != nil {
		C.QEventLoop_Exit(ptr.Pointer(), C.int(returnCode))
	}
}

func (ptr *QEventLoop) IsRunning() bool {
	defer qt.Recovering("QEventLoop::isRunning")

	if ptr.Pointer() != nil {
		return C.QEventLoop_IsRunning(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QEventLoop) ProcessEvents(flags QEventLoop__ProcessEventsFlag) bool {
	defer qt.Recovering("QEventLoop::processEvents")

	if ptr.Pointer() != nil {
		return C.QEventLoop_ProcessEvents(ptr.Pointer(), C.int(flags)) != 0
	}
	return false
}

func (ptr *QEventLoop) ProcessEvents2(flags QEventLoop__ProcessEventsFlag, maxTime int) {
	defer qt.Recovering("QEventLoop::processEvents")

	if ptr.Pointer() != nil {
		C.QEventLoop_ProcessEvents2(ptr.Pointer(), C.int(flags), C.int(maxTime))
	}
}

func (ptr *QEventLoop) Quit() {
	defer qt.Recovering("QEventLoop::quit")

	if ptr.Pointer() != nil {
		C.QEventLoop_Quit(ptr.Pointer())
	}
}

func (ptr *QEventLoop) WakeUp() {
	defer qt.Recovering("QEventLoop::wakeUp")

	if ptr.Pointer() != nil {
		C.QEventLoop_WakeUp(ptr.Pointer())
	}
}

func (ptr *QEventLoop) DestroyQEventLoop() {
	defer qt.Recovering("QEventLoop::~QEventLoop")

	if ptr.Pointer() != nil {
		C.QEventLoop_DestroyQEventLoop(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QEventLoop) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QEventLoop::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QEventLoop) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QEventLoop::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQEventLoopTimerEvent
func callbackQEventLoopTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QEventLoop::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
	} else {
		NewQEventLoopFromPointer(ptr).TimerEventDefault(NewQTimerEventFromPointer(event))
	}
}

func (ptr *QEventLoop) TimerEvent(event QTimerEvent_ITF) {
	defer qt.Recovering("QEventLoop::timerEvent")

	if ptr.Pointer() != nil {
		C.QEventLoop_TimerEvent(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QEventLoop) TimerEventDefault(event QTimerEvent_ITF) {
	defer qt.Recovering("QEventLoop::timerEvent")

	if ptr.Pointer() != nil {
		C.QEventLoop_TimerEventDefault(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QEventLoop) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QEventLoop::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QEventLoop) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QEventLoop::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQEventLoopChildEvent
func callbackQEventLoopChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QEventLoop::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
	} else {
		NewQEventLoopFromPointer(ptr).ChildEventDefault(NewQChildEventFromPointer(event))
	}
}

func (ptr *QEventLoop) ChildEvent(event QChildEvent_ITF) {
	defer qt.Recovering("QEventLoop::childEvent")

	if ptr.Pointer() != nil {
		C.QEventLoop_ChildEvent(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QEventLoop) ChildEventDefault(event QChildEvent_ITF) {
	defer qt.Recovering("QEventLoop::childEvent")

	if ptr.Pointer() != nil {
		C.QEventLoop_ChildEventDefault(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QEventLoop) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QEventLoop::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QEventLoop) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QEventLoop::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQEventLoopCustomEvent
func callbackQEventLoopCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QEventLoop::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
	} else {
		NewQEventLoopFromPointer(ptr).CustomEventDefault(NewQEventFromPointer(event))
	}
}

func (ptr *QEventLoop) CustomEvent(event QEvent_ITF) {
	defer qt.Recovering("QEventLoop::customEvent")

	if ptr.Pointer() != nil {
		C.QEventLoop_CustomEvent(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QEventLoop) CustomEventDefault(event QEvent_ITF) {
	defer qt.Recovering("QEventLoop::customEvent")

	if ptr.Pointer() != nil {
		C.QEventLoop_CustomEventDefault(ptr.Pointer(), PointerFromQEvent(event))
	}
}
