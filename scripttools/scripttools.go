// +build !minimal

package scripttools

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "scripttools.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/script"
	"github.com/therecipe/qt/widgets"
	"strings"
	"unsafe"
)

func cGoFreePacked(ptr unsafe.Pointer) { core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_QtScriptTools_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtScriptTools_PackedString) []byte {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return *(*[]byte)(unsafe.Pointer(&gs))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type QScriptEngineDebugger struct {
	core.QObject
}

type QScriptEngineDebugger_ITF interface {
	core.QObject_ITF
	QScriptEngineDebugger_PTR() *QScriptEngineDebugger
}

func (ptr *QScriptEngineDebugger) QScriptEngineDebugger_PTR() *QScriptEngineDebugger {
	return ptr
}

func (ptr *QScriptEngineDebugger) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QScriptEngineDebugger) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQScriptEngineDebugger(ptr QScriptEngineDebugger_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptEngineDebugger_PTR().Pointer()
	}
	return nil
}

func NewQScriptEngineDebuggerFromPointer(ptr unsafe.Pointer) (n *QScriptEngineDebugger) {
	n = new(QScriptEngineDebugger)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QScriptEngineDebugger__DebuggerAction
//QScriptEngineDebugger::DebuggerAction
type QScriptEngineDebugger__DebuggerAction int64

const (
	QScriptEngineDebugger__InterruptAction            QScriptEngineDebugger__DebuggerAction = QScriptEngineDebugger__DebuggerAction(0)
	QScriptEngineDebugger__ContinueAction             QScriptEngineDebugger__DebuggerAction = QScriptEngineDebugger__DebuggerAction(1)
	QScriptEngineDebugger__StepIntoAction             QScriptEngineDebugger__DebuggerAction = QScriptEngineDebugger__DebuggerAction(2)
	QScriptEngineDebugger__StepOverAction             QScriptEngineDebugger__DebuggerAction = QScriptEngineDebugger__DebuggerAction(3)
	QScriptEngineDebugger__StepOutAction              QScriptEngineDebugger__DebuggerAction = QScriptEngineDebugger__DebuggerAction(4)
	QScriptEngineDebugger__RunToCursorAction          QScriptEngineDebugger__DebuggerAction = QScriptEngineDebugger__DebuggerAction(5)
	QScriptEngineDebugger__RunToNewScriptAction       QScriptEngineDebugger__DebuggerAction = QScriptEngineDebugger__DebuggerAction(6)
	QScriptEngineDebugger__ToggleBreakpointAction     QScriptEngineDebugger__DebuggerAction = QScriptEngineDebugger__DebuggerAction(7)
	QScriptEngineDebugger__ClearDebugOutputAction     QScriptEngineDebugger__DebuggerAction = QScriptEngineDebugger__DebuggerAction(8)
	QScriptEngineDebugger__ClearErrorLogAction        QScriptEngineDebugger__DebuggerAction = QScriptEngineDebugger__DebuggerAction(9)
	QScriptEngineDebugger__ClearConsoleAction         QScriptEngineDebugger__DebuggerAction = QScriptEngineDebugger__DebuggerAction(10)
	QScriptEngineDebugger__FindInScriptAction         QScriptEngineDebugger__DebuggerAction = QScriptEngineDebugger__DebuggerAction(11)
	QScriptEngineDebugger__FindNextInScriptAction     QScriptEngineDebugger__DebuggerAction = QScriptEngineDebugger__DebuggerAction(12)
	QScriptEngineDebugger__FindPreviousInScriptAction QScriptEngineDebugger__DebuggerAction = QScriptEngineDebugger__DebuggerAction(13)
	QScriptEngineDebugger__GoToLineAction             QScriptEngineDebugger__DebuggerAction = QScriptEngineDebugger__DebuggerAction(14)
)

//go:generate stringer -type=QScriptEngineDebugger__DebuggerWidget
//QScriptEngineDebugger::DebuggerWidget
type QScriptEngineDebugger__DebuggerWidget int64

const (
	QScriptEngineDebugger__ConsoleWidget     QScriptEngineDebugger__DebuggerWidget = QScriptEngineDebugger__DebuggerWidget(0)
	QScriptEngineDebugger__StackWidget       QScriptEngineDebugger__DebuggerWidget = QScriptEngineDebugger__DebuggerWidget(1)
	QScriptEngineDebugger__ScriptsWidget     QScriptEngineDebugger__DebuggerWidget = QScriptEngineDebugger__DebuggerWidget(2)
	QScriptEngineDebugger__LocalsWidget      QScriptEngineDebugger__DebuggerWidget = QScriptEngineDebugger__DebuggerWidget(3)
	QScriptEngineDebugger__CodeWidget        QScriptEngineDebugger__DebuggerWidget = QScriptEngineDebugger__DebuggerWidget(4)
	QScriptEngineDebugger__CodeFinderWidget  QScriptEngineDebugger__DebuggerWidget = QScriptEngineDebugger__DebuggerWidget(5)
	QScriptEngineDebugger__BreakpointsWidget QScriptEngineDebugger__DebuggerWidget = QScriptEngineDebugger__DebuggerWidget(6)
	QScriptEngineDebugger__DebugOutputWidget QScriptEngineDebugger__DebuggerWidget = QScriptEngineDebugger__DebuggerWidget(7)
	QScriptEngineDebugger__ErrorLogWidget    QScriptEngineDebugger__DebuggerWidget = QScriptEngineDebugger__DebuggerWidget(8)
)

//go:generate stringer -type=QScriptEngineDebugger__DebuggerState
//QScriptEngineDebugger::DebuggerState
type QScriptEngineDebugger__DebuggerState int64

const (
	QScriptEngineDebugger__RunningState   QScriptEngineDebugger__DebuggerState = QScriptEngineDebugger__DebuggerState(0)
	QScriptEngineDebugger__SuspendedState QScriptEngineDebugger__DebuggerState = QScriptEngineDebugger__DebuggerState(1)
)

func NewQScriptEngineDebugger(parent core.QObject_ITF) *QScriptEngineDebugger {
	tmpValue := NewQScriptEngineDebuggerFromPointer(C.QScriptEngineDebugger_NewQScriptEngineDebugger(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QScriptEngineDebugger) Action(action QScriptEngineDebugger__DebuggerAction) *widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQActionFromPointer(C.QScriptEngineDebugger_Action(ptr.Pointer(), C.longlong(action)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngineDebugger) AttachTo(engine script.QScriptEngine_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_AttachTo(ptr.Pointer(), script.PointerFromQScriptEngine(engine))
	}
}

func (ptr *QScriptEngineDebugger) AutoShowStandardWindow() bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptEngineDebugger_AutoShowStandardWindow(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptEngineDebugger) CreateStandardMenu(parent widgets.QWidget_ITF) *widgets.QMenu {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQMenuFromPointer(C.QScriptEngineDebugger_CreateStandardMenu(ptr.Pointer(), widgets.PointerFromQWidget(parent)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngineDebugger) CreateStandardToolBar(parent widgets.QWidget_ITF) *widgets.QToolBar {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQToolBarFromPointer(C.QScriptEngineDebugger_CreateStandardToolBar(ptr.Pointer(), widgets.PointerFromQWidget(parent)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngineDebugger) Detach() {
	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_Detach(ptr.Pointer())
	}
}

//export callbackQScriptEngineDebugger_EvaluationResumed
func callbackQScriptEngineDebugger_EvaluationResumed(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "evaluationResumed"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QScriptEngineDebugger) ConnectEvaluationResumed(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "evaluationResumed") {
			C.QScriptEngineDebugger_ConnectEvaluationResumed(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "evaluationResumed")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "evaluationResumed"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "evaluationResumed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "evaluationResumed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScriptEngineDebugger) DisconnectEvaluationResumed() {
	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_DisconnectEvaluationResumed(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "evaluationResumed")
	}
}

func (ptr *QScriptEngineDebugger) EvaluationResumed() {
	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_EvaluationResumed(ptr.Pointer())
	}
}

//export callbackQScriptEngineDebugger_EvaluationSuspended
func callbackQScriptEngineDebugger_EvaluationSuspended(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "evaluationSuspended"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QScriptEngineDebugger) ConnectEvaluationSuspended(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "evaluationSuspended") {
			C.QScriptEngineDebugger_ConnectEvaluationSuspended(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "evaluationSuspended")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "evaluationSuspended"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "evaluationSuspended", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "evaluationSuspended", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScriptEngineDebugger) DisconnectEvaluationSuspended() {
	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_DisconnectEvaluationSuspended(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "evaluationSuspended")
	}
}

func (ptr *QScriptEngineDebugger) EvaluationSuspended() {
	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_EvaluationSuspended(ptr.Pointer())
	}
}

func (ptr *QScriptEngineDebugger) SetAutoShowStandardWindow(autoShow bool) {
	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_SetAutoShowStandardWindow(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(autoShow))))
	}
}

func (ptr *QScriptEngineDebugger) StandardWindow() *widgets.QMainWindow {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQMainWindowFromPointer(C.QScriptEngineDebugger_StandardWindow(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngineDebugger) State() QScriptEngineDebugger__DebuggerState {
	if ptr.Pointer() != nil {
		return QScriptEngineDebugger__DebuggerState(C.QScriptEngineDebugger_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptEngineDebugger) Widget(widget QScriptEngineDebugger__DebuggerWidget) *widgets.QWidget {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQWidgetFromPointer(C.QScriptEngineDebugger_Widget(ptr.Pointer(), C.longlong(widget)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQScriptEngineDebugger_DestroyQScriptEngineDebugger
func callbackQScriptEngineDebugger_DestroyQScriptEngineDebugger(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QScriptEngineDebugger"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQScriptEngineDebuggerFromPointer(ptr).DestroyQScriptEngineDebuggerDefault()
	}
}

func (ptr *QScriptEngineDebugger) ConnectDestroyQScriptEngineDebugger(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QScriptEngineDebugger"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QScriptEngineDebugger", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QScriptEngineDebugger", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScriptEngineDebugger) DisconnectDestroyQScriptEngineDebugger() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QScriptEngineDebugger")
	}
}

func (ptr *QScriptEngineDebugger) DestroyQScriptEngineDebugger() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QScriptEngineDebugger_DestroyQScriptEngineDebugger(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScriptEngineDebugger) DestroyQScriptEngineDebuggerDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QScriptEngineDebugger_DestroyQScriptEngineDebuggerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScriptEngineDebugger) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QScriptEngineDebugger___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngineDebugger) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScriptEngineDebugger) __children_newList() unsafe.Pointer {
	return C.QScriptEngineDebugger___children_newList(ptr.Pointer())
}

func (ptr *QScriptEngineDebugger) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QScriptEngineDebugger___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngineDebugger) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QScriptEngineDebugger) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QScriptEngineDebugger___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QScriptEngineDebugger) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QScriptEngineDebugger___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngineDebugger) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScriptEngineDebugger) __findChildren_newList() unsafe.Pointer {
	return C.QScriptEngineDebugger___findChildren_newList(ptr.Pointer())
}

func (ptr *QScriptEngineDebugger) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QScriptEngineDebugger___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngineDebugger) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScriptEngineDebugger) __findChildren_newList3() unsafe.Pointer {
	return C.QScriptEngineDebugger___findChildren_newList3(ptr.Pointer())
}

//export callbackQScriptEngineDebugger_ChildEvent
func callbackQScriptEngineDebugger_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQScriptEngineDebuggerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QScriptEngineDebugger) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQScriptEngineDebugger_ConnectNotify
func callbackQScriptEngineDebugger_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScriptEngineDebuggerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScriptEngineDebugger) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScriptEngineDebugger_CustomEvent
func callbackQScriptEngineDebugger_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQScriptEngineDebuggerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QScriptEngineDebugger) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQScriptEngineDebugger_DeleteLater
func callbackQScriptEngineDebugger_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQScriptEngineDebuggerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QScriptEngineDebugger) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QScriptEngineDebugger_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQScriptEngineDebugger_Destroyed
func callbackQScriptEngineDebugger_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQScriptEngineDebugger_DisconnectNotify
func callbackQScriptEngineDebugger_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScriptEngineDebuggerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScriptEngineDebugger) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScriptEngineDebugger_Event
func callbackQScriptEngineDebugger_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScriptEngineDebuggerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QScriptEngineDebugger) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptEngineDebugger_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQScriptEngineDebugger_EventFilter
func callbackQScriptEngineDebugger_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScriptEngineDebuggerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QScriptEngineDebugger) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptEngineDebugger_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQScriptEngineDebugger_MetaObject
func callbackQScriptEngineDebugger_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQScriptEngineDebuggerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QScriptEngineDebugger) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScriptEngineDebugger_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQScriptEngineDebugger_ObjectNameChanged
func callbackQScriptEngineDebugger_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtScriptTools_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQScriptEngineDebugger_TimerEvent
func callbackQScriptEngineDebugger_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScriptEngineDebuggerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QScriptEngineDebugger) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func init() {
	qt.ItfMap["scripttools.QScriptEngineDebugger_ITF"] = QScriptEngineDebugger{}
	qt.FuncMap["scripttools.NewQScriptEngineDebugger"] = NewQScriptEngineDebugger
	qt.EnumMap["scripttools.QScriptEngineDebugger__InterruptAction"] = int64(QScriptEngineDebugger__InterruptAction)
	qt.EnumMap["scripttools.QScriptEngineDebugger__ContinueAction"] = int64(QScriptEngineDebugger__ContinueAction)
	qt.EnumMap["scripttools.QScriptEngineDebugger__StepIntoAction"] = int64(QScriptEngineDebugger__StepIntoAction)
	qt.EnumMap["scripttools.QScriptEngineDebugger__StepOverAction"] = int64(QScriptEngineDebugger__StepOverAction)
	qt.EnumMap["scripttools.QScriptEngineDebugger__StepOutAction"] = int64(QScriptEngineDebugger__StepOutAction)
	qt.EnumMap["scripttools.QScriptEngineDebugger__RunToCursorAction"] = int64(QScriptEngineDebugger__RunToCursorAction)
	qt.EnumMap["scripttools.QScriptEngineDebugger__RunToNewScriptAction"] = int64(QScriptEngineDebugger__RunToNewScriptAction)
	qt.EnumMap["scripttools.QScriptEngineDebugger__ToggleBreakpointAction"] = int64(QScriptEngineDebugger__ToggleBreakpointAction)
	qt.EnumMap["scripttools.QScriptEngineDebugger__ClearDebugOutputAction"] = int64(QScriptEngineDebugger__ClearDebugOutputAction)
	qt.EnumMap["scripttools.QScriptEngineDebugger__ClearErrorLogAction"] = int64(QScriptEngineDebugger__ClearErrorLogAction)
	qt.EnumMap["scripttools.QScriptEngineDebugger__ClearConsoleAction"] = int64(QScriptEngineDebugger__ClearConsoleAction)
	qt.EnumMap["scripttools.QScriptEngineDebugger__FindInScriptAction"] = int64(QScriptEngineDebugger__FindInScriptAction)
	qt.EnumMap["scripttools.QScriptEngineDebugger__FindNextInScriptAction"] = int64(QScriptEngineDebugger__FindNextInScriptAction)
	qt.EnumMap["scripttools.QScriptEngineDebugger__FindPreviousInScriptAction"] = int64(QScriptEngineDebugger__FindPreviousInScriptAction)
	qt.EnumMap["scripttools.QScriptEngineDebugger__GoToLineAction"] = int64(QScriptEngineDebugger__GoToLineAction)
	qt.EnumMap["scripttools.QScriptEngineDebugger__ConsoleWidget"] = int64(QScriptEngineDebugger__ConsoleWidget)
	qt.EnumMap["scripttools.QScriptEngineDebugger__StackWidget"] = int64(QScriptEngineDebugger__StackWidget)
	qt.EnumMap["scripttools.QScriptEngineDebugger__ScriptsWidget"] = int64(QScriptEngineDebugger__ScriptsWidget)
	qt.EnumMap["scripttools.QScriptEngineDebugger__LocalsWidget"] = int64(QScriptEngineDebugger__LocalsWidget)
	qt.EnumMap["scripttools.QScriptEngineDebugger__CodeWidget"] = int64(QScriptEngineDebugger__CodeWidget)
	qt.EnumMap["scripttools.QScriptEngineDebugger__CodeFinderWidget"] = int64(QScriptEngineDebugger__CodeFinderWidget)
	qt.EnumMap["scripttools.QScriptEngineDebugger__BreakpointsWidget"] = int64(QScriptEngineDebugger__BreakpointsWidget)
	qt.EnumMap["scripttools.QScriptEngineDebugger__DebugOutputWidget"] = int64(QScriptEngineDebugger__DebugOutputWidget)
	qt.EnumMap["scripttools.QScriptEngineDebugger__ErrorLogWidget"] = int64(QScriptEngineDebugger__ErrorLogWidget)
	qt.EnumMap["scripttools.QScriptEngineDebugger__RunningState"] = int64(QScriptEngineDebugger__RunningState)
	qt.EnumMap["scripttools.QScriptEngineDebugger__SuspendedState"] = int64(QScriptEngineDebugger__SuspendedState)
}
