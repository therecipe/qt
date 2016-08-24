// +build !minimal

package scripttools

//#include <stdint.h>
//#include <stdlib.h>
//#include "scripttools.h"
import "C"
import (
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/script"
	"github.com/therecipe/qt/widgets"
	"unsafe"
)

//QScriptEngineDebugger::DebuggerAction
type QScriptEngineDebugger__DebuggerAction int64

const (
	QScriptEngineDebugger__InterruptAction            = QScriptEngineDebugger__DebuggerAction(0)
	QScriptEngineDebugger__ContinueAction             = QScriptEngineDebugger__DebuggerAction(1)
	QScriptEngineDebugger__StepIntoAction             = QScriptEngineDebugger__DebuggerAction(2)
	QScriptEngineDebugger__StepOverAction             = QScriptEngineDebugger__DebuggerAction(3)
	QScriptEngineDebugger__StepOutAction              = QScriptEngineDebugger__DebuggerAction(4)
	QScriptEngineDebugger__RunToCursorAction          = QScriptEngineDebugger__DebuggerAction(5)
	QScriptEngineDebugger__RunToNewScriptAction       = QScriptEngineDebugger__DebuggerAction(6)
	QScriptEngineDebugger__ToggleBreakpointAction     = QScriptEngineDebugger__DebuggerAction(7)
	QScriptEngineDebugger__ClearDebugOutputAction     = QScriptEngineDebugger__DebuggerAction(8)
	QScriptEngineDebugger__ClearErrorLogAction        = QScriptEngineDebugger__DebuggerAction(9)
	QScriptEngineDebugger__ClearConsoleAction         = QScriptEngineDebugger__DebuggerAction(10)
	QScriptEngineDebugger__FindInScriptAction         = QScriptEngineDebugger__DebuggerAction(11)
	QScriptEngineDebugger__FindNextInScriptAction     = QScriptEngineDebugger__DebuggerAction(12)
	QScriptEngineDebugger__FindPreviousInScriptAction = QScriptEngineDebugger__DebuggerAction(13)
	QScriptEngineDebugger__GoToLineAction             = QScriptEngineDebugger__DebuggerAction(14)
)

//QScriptEngineDebugger::DebuggerState
type QScriptEngineDebugger__DebuggerState int64

const (
	QScriptEngineDebugger__RunningState   = QScriptEngineDebugger__DebuggerState(0)
	QScriptEngineDebugger__SuspendedState = QScriptEngineDebugger__DebuggerState(1)
)

//QScriptEngineDebugger::DebuggerWidget
type QScriptEngineDebugger__DebuggerWidget int64

const (
	QScriptEngineDebugger__ConsoleWidget     = QScriptEngineDebugger__DebuggerWidget(0)
	QScriptEngineDebugger__StackWidget       = QScriptEngineDebugger__DebuggerWidget(1)
	QScriptEngineDebugger__ScriptsWidget     = QScriptEngineDebugger__DebuggerWidget(2)
	QScriptEngineDebugger__LocalsWidget      = QScriptEngineDebugger__DebuggerWidget(3)
	QScriptEngineDebugger__CodeWidget        = QScriptEngineDebugger__DebuggerWidget(4)
	QScriptEngineDebugger__CodeFinderWidget  = QScriptEngineDebugger__DebuggerWidget(5)
	QScriptEngineDebugger__BreakpointsWidget = QScriptEngineDebugger__DebuggerWidget(6)
	QScriptEngineDebugger__DebugOutputWidget = QScriptEngineDebugger__DebuggerWidget(7)
	QScriptEngineDebugger__ErrorLogWidget    = QScriptEngineDebugger__DebuggerWidget(8)
)

type QScriptEngineDebugger struct {
	core.QObject
}

type QScriptEngineDebugger_ITF interface {
	core.QObject_ITF
	QScriptEngineDebugger_PTR() *QScriptEngineDebugger
}

func (p *QScriptEngineDebugger) QScriptEngineDebugger_PTR() *QScriptEngineDebugger {
	return p
}

func (p *QScriptEngineDebugger) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QScriptEngineDebugger) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQScriptEngineDebugger(ptr QScriptEngineDebugger_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptEngineDebugger_PTR().Pointer()
	}
	return nil
}

func NewQScriptEngineDebuggerFromPointer(ptr unsafe.Pointer) *QScriptEngineDebugger {
	var n = new(QScriptEngineDebugger)
	n.SetPointer(ptr)
	return n
}
func NewQScriptEngineDebugger(parent core.QObject_ITF) *QScriptEngineDebugger {
	defer qt.Recovering("QScriptEngineDebugger::QScriptEngineDebugger")

	var tmpValue = NewQScriptEngineDebuggerFromPointer(C.QScriptEngineDebugger_NewQScriptEngineDebugger(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) {})
	}
	return tmpValue
}

func (ptr *QScriptEngineDebugger) Action(action QScriptEngineDebugger__DebuggerAction) *widgets.QAction {
	defer qt.Recovering("QScriptEngineDebugger::action")

	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QScriptEngineDebugger_Action(ptr.Pointer(), C.longlong(action)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngineDebugger) AttachTo(engine script.QScriptEngine_ITF) {
	defer qt.Recovering("QScriptEngineDebugger::attachTo")

	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_AttachTo(ptr.Pointer(), script.PointerFromQScriptEngine(engine))
	}
}

func (ptr *QScriptEngineDebugger) AutoShowStandardWindow() bool {
	defer qt.Recovering("QScriptEngineDebugger::autoShowStandardWindow")

	if ptr.Pointer() != nil {
		return C.QScriptEngineDebugger_AutoShowStandardWindow(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptEngineDebugger) CreateStandardMenu(parent widgets.QWidget_ITF) *widgets.QMenu {
	defer qt.Recovering("QScriptEngineDebugger::createStandardMenu")

	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQMenuFromPointer(C.QScriptEngineDebugger_CreateStandardMenu(ptr.Pointer(), widgets.PointerFromQWidget(parent)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngineDebugger) CreateStandardToolBar(parent widgets.QWidget_ITF) *widgets.QToolBar {
	defer qt.Recovering("QScriptEngineDebugger::createStandardToolBar")

	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQToolBarFromPointer(C.QScriptEngineDebugger_CreateStandardToolBar(ptr.Pointer(), widgets.PointerFromQWidget(parent)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngineDebugger) Detach() {
	defer qt.Recovering("QScriptEngineDebugger::detach")

	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_Detach(ptr.Pointer())
	}
}

//export callbackQScriptEngineDebugger_EvaluationResumed
func callbackQScriptEngineDebugger_EvaluationResumed(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QScriptEngineDebugger::evaluationResumed")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScriptEngineDebugger::evaluationResumed"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QScriptEngineDebugger) ConnectEvaluationResumed(f func()) {
	defer qt.Recovering("connect QScriptEngineDebugger::evaluationResumed")

	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_ConnectEvaluationResumed(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScriptEngineDebugger::evaluationResumed", f)
	}
}

func (ptr *QScriptEngineDebugger) DisconnectEvaluationResumed() {
	defer qt.Recovering("disconnect QScriptEngineDebugger::evaluationResumed")

	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_DisconnectEvaluationResumed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScriptEngineDebugger::evaluationResumed")
	}
}

func (ptr *QScriptEngineDebugger) EvaluationResumed() {
	defer qt.Recovering("QScriptEngineDebugger::evaluationResumed")

	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_EvaluationResumed(ptr.Pointer())
	}
}

//export callbackQScriptEngineDebugger_EvaluationSuspended
func callbackQScriptEngineDebugger_EvaluationSuspended(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QScriptEngineDebugger::evaluationSuspended")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScriptEngineDebugger::evaluationSuspended"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QScriptEngineDebugger) ConnectEvaluationSuspended(f func()) {
	defer qt.Recovering("connect QScriptEngineDebugger::evaluationSuspended")

	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_ConnectEvaluationSuspended(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScriptEngineDebugger::evaluationSuspended", f)
	}
}

func (ptr *QScriptEngineDebugger) DisconnectEvaluationSuspended() {
	defer qt.Recovering("disconnect QScriptEngineDebugger::evaluationSuspended")

	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_DisconnectEvaluationSuspended(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScriptEngineDebugger::evaluationSuspended")
	}
}

func (ptr *QScriptEngineDebugger) EvaluationSuspended() {
	defer qt.Recovering("QScriptEngineDebugger::evaluationSuspended")

	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_EvaluationSuspended(ptr.Pointer())
	}
}

func (ptr *QScriptEngineDebugger) SetAutoShowStandardWindow(autoShow bool) {
	defer qt.Recovering("QScriptEngineDebugger::setAutoShowStandardWindow")

	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_SetAutoShowStandardWindow(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(autoShow))))
	}
}

func (ptr *QScriptEngineDebugger) StandardWindow() *widgets.QMainWindow {
	defer qt.Recovering("QScriptEngineDebugger::standardWindow")

	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQMainWindowFromPointer(C.QScriptEngineDebugger_StandardWindow(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngineDebugger) State() QScriptEngineDebugger__DebuggerState {
	defer qt.Recovering("QScriptEngineDebugger::state")

	if ptr.Pointer() != nil {
		return QScriptEngineDebugger__DebuggerState(C.QScriptEngineDebugger_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptEngineDebugger) Widget(widget QScriptEngineDebugger__DebuggerWidget) *widgets.QWidget {
	defer qt.Recovering("QScriptEngineDebugger::widget")

	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQWidgetFromPointer(C.QScriptEngineDebugger_Widget(ptr.Pointer(), C.longlong(widget)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngineDebugger) DestroyQScriptEngineDebugger() {
	defer qt.Recovering("QScriptEngineDebugger::~QScriptEngineDebugger")

	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_DestroyQScriptEngineDebugger(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQScriptEngineDebugger_TimerEvent
func callbackQScriptEngineDebugger_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QScriptEngineDebugger::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScriptEngineDebugger::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScriptEngineDebuggerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QScriptEngineDebugger) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QScriptEngineDebugger::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScriptEngineDebugger::timerEvent", f)
	}
}

func (ptr *QScriptEngineDebugger) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QScriptEngineDebugger::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScriptEngineDebugger::timerEvent")
	}
}

func (ptr *QScriptEngineDebugger) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QScriptEngineDebugger::timerEvent")

	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QScriptEngineDebugger) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QScriptEngineDebugger::timerEvent")

	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQScriptEngineDebugger_ChildEvent
func callbackQScriptEngineDebugger_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QScriptEngineDebugger::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScriptEngineDebugger::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQScriptEngineDebuggerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QScriptEngineDebugger) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QScriptEngineDebugger::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScriptEngineDebugger::childEvent", f)
	}
}

func (ptr *QScriptEngineDebugger) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QScriptEngineDebugger::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScriptEngineDebugger::childEvent")
	}
}

func (ptr *QScriptEngineDebugger) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QScriptEngineDebugger::childEvent")

	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QScriptEngineDebugger) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QScriptEngineDebugger::childEvent")

	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQScriptEngineDebugger_ConnectNotify
func callbackQScriptEngineDebugger_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QScriptEngineDebugger::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScriptEngineDebugger::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScriptEngineDebuggerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScriptEngineDebugger) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QScriptEngineDebugger::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScriptEngineDebugger::connectNotify", f)
	}
}

func (ptr *QScriptEngineDebugger) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QScriptEngineDebugger::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScriptEngineDebugger::connectNotify")
	}
}

func (ptr *QScriptEngineDebugger) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScriptEngineDebugger::connectNotify")

	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScriptEngineDebugger) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScriptEngineDebugger::connectNotify")

	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScriptEngineDebugger_CustomEvent
func callbackQScriptEngineDebugger_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QScriptEngineDebugger::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScriptEngineDebugger::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQScriptEngineDebuggerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QScriptEngineDebugger) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QScriptEngineDebugger::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScriptEngineDebugger::customEvent", f)
	}
}

func (ptr *QScriptEngineDebugger) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QScriptEngineDebugger::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScriptEngineDebugger::customEvent")
	}
}

func (ptr *QScriptEngineDebugger) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QScriptEngineDebugger::customEvent")

	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QScriptEngineDebugger) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QScriptEngineDebugger::customEvent")

	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQScriptEngineDebugger_DeleteLater
func callbackQScriptEngineDebugger_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QScriptEngineDebugger::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScriptEngineDebugger::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQScriptEngineDebuggerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QScriptEngineDebugger) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QScriptEngineDebugger::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScriptEngineDebugger::deleteLater", f)
	}
}

func (ptr *QScriptEngineDebugger) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QScriptEngineDebugger::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScriptEngineDebugger::deleteLater")
	}
}

func (ptr *QScriptEngineDebugger) DeleteLater() {
	defer qt.Recovering("QScriptEngineDebugger::deleteLater")

	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QScriptEngineDebugger) DeleteLaterDefault() {
	defer qt.Recovering("QScriptEngineDebugger::deleteLater")

	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQScriptEngineDebugger_DisconnectNotify
func callbackQScriptEngineDebugger_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QScriptEngineDebugger::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScriptEngineDebugger::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScriptEngineDebuggerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScriptEngineDebugger) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QScriptEngineDebugger::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScriptEngineDebugger::disconnectNotify", f)
	}
}

func (ptr *QScriptEngineDebugger) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QScriptEngineDebugger::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScriptEngineDebugger::disconnectNotify")
	}
}

func (ptr *QScriptEngineDebugger) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScriptEngineDebugger::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScriptEngineDebugger) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScriptEngineDebugger::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScriptEngineDebugger_Event
func callbackQScriptEngineDebugger_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QScriptEngineDebugger::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScriptEngineDebugger::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScriptEngineDebuggerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QScriptEngineDebugger) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QScriptEngineDebugger::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScriptEngineDebugger::event", f)
	}
}

func (ptr *QScriptEngineDebugger) DisconnectEvent() {
	defer qt.Recovering("disconnect QScriptEngineDebugger::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScriptEngineDebugger::event")
	}
}

func (ptr *QScriptEngineDebugger) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QScriptEngineDebugger::event")

	if ptr.Pointer() != nil {
		return C.QScriptEngineDebugger_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QScriptEngineDebugger) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QScriptEngineDebugger::event")

	if ptr.Pointer() != nil {
		return C.QScriptEngineDebugger_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQScriptEngineDebugger_EventFilter
func callbackQScriptEngineDebugger_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QScriptEngineDebugger::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScriptEngineDebugger::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScriptEngineDebuggerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QScriptEngineDebugger) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QScriptEngineDebugger::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScriptEngineDebugger::eventFilter", f)
	}
}

func (ptr *QScriptEngineDebugger) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QScriptEngineDebugger::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScriptEngineDebugger::eventFilter")
	}
}

func (ptr *QScriptEngineDebugger) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QScriptEngineDebugger::eventFilter")

	if ptr.Pointer() != nil {
		return C.QScriptEngineDebugger_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QScriptEngineDebugger) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QScriptEngineDebugger::eventFilter")

	if ptr.Pointer() != nil {
		return C.QScriptEngineDebugger_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQScriptEngineDebugger_MetaObject
func callbackQScriptEngineDebugger_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QScriptEngineDebugger::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScriptEngineDebugger::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQScriptEngineDebuggerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QScriptEngineDebugger) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QScriptEngineDebugger::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScriptEngineDebugger::metaObject", f)
	}
}

func (ptr *QScriptEngineDebugger) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QScriptEngineDebugger::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScriptEngineDebugger::metaObject")
	}
}

func (ptr *QScriptEngineDebugger) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QScriptEngineDebugger::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScriptEngineDebugger_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptEngineDebugger) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QScriptEngineDebugger::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScriptEngineDebugger_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
