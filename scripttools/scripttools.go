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
	"runtime"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtScriptTools_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtScriptTools_PackedString) []byte {
	if int(s.len) == -1 {
		return []byte(C.GoString(s.data))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
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

//go:generate stringer -type=QScriptEngineDebugger__DebuggerState
//QScriptEngineDebugger::DebuggerState
type QScriptEngineDebugger__DebuggerState int64

const (
	QScriptEngineDebugger__RunningState   QScriptEngineDebugger__DebuggerState = QScriptEngineDebugger__DebuggerState(0)
	QScriptEngineDebugger__SuspendedState QScriptEngineDebugger__DebuggerState = QScriptEngineDebugger__DebuggerState(1)
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

func NewQScriptEngineDebugger(parent core.QObject_ITF) *QScriptEngineDebugger {
	tmpValue := NewQScriptEngineDebuggerFromPointer(C.QScriptEngineDebugger_NewQScriptEngineDebugger(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QScriptEngineDebugger_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QScriptEngineDebugger_QScriptEngineDebugger_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QScriptEngineDebugger) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QScriptEngineDebugger_QScriptEngineDebugger_Tr(sC, cC, C.int(int32(n))))
}

func QScriptEngineDebugger_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QScriptEngineDebugger_QScriptEngineDebugger_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QScriptEngineDebugger) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QScriptEngineDebugger_QScriptEngineDebugger_TrUtf8(sC, cC, C.int(int32(n))))
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

func (ptr *QScriptEngineDebugger) AttachTo(engine script.QScriptEngine_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_AttachTo(ptr.Pointer(), script.PointerFromQScriptEngine(engine))
	}
}

func (ptr *QScriptEngineDebugger) Detach() {
	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_Detach(ptr.Pointer())
	}
}

//export callbackQScriptEngineDebugger_EvaluationResumed
func callbackQScriptEngineDebugger_EvaluationResumed(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "evaluationResumed"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QScriptEngineDebugger) ConnectEvaluationResumed(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "evaluationResumed") {
			C.QScriptEngineDebugger_ConnectEvaluationResumed(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "evaluationResumed"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "evaluationResumed", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "evaluationResumed", f)
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
		signal.(func())()
	}

}

func (ptr *QScriptEngineDebugger) ConnectEvaluationSuspended(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "evaluationSuspended") {
			C.QScriptEngineDebugger_ConnectEvaluationSuspended(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "evaluationSuspended"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "evaluationSuspended", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "evaluationSuspended", f)
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

//export callbackQScriptEngineDebugger_DestroyQScriptEngineDebugger
func callbackQScriptEngineDebugger_DestroyQScriptEngineDebugger(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QScriptEngineDebugger"); signal != nil {
		signal.(func())()
	} else {
		NewQScriptEngineDebuggerFromPointer(ptr).DestroyQScriptEngineDebuggerDefault()
	}
}

func (ptr *QScriptEngineDebugger) ConnectDestroyQScriptEngineDebugger(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QScriptEngineDebugger"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QScriptEngineDebugger", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QScriptEngineDebugger", f)
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
		C.QScriptEngineDebugger_DestroyQScriptEngineDebugger(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QScriptEngineDebugger) DestroyQScriptEngineDebuggerDefault() {
	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_DestroyQScriptEngineDebuggerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
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

func (ptr *QScriptEngineDebugger) AutoShowStandardWindow() bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptEngineDebugger_AutoShowStandardWindow(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQScriptEngineDebugger_MetaObject
func callbackQScriptEngineDebugger_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQScriptEngineDebuggerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QScriptEngineDebugger) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScriptEngineDebugger_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptEngineDebugger) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QScriptEngineDebugger___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

func (ptr *QScriptEngineDebugger) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QScriptEngineDebugger___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngineDebugger) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScriptEngineDebugger) __findChildren_newList2() unsafe.Pointer {
	return C.QScriptEngineDebugger___findChildren_newList2(ptr.Pointer())
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

//export callbackQScriptEngineDebugger_Event
func callbackQScriptEngineDebugger_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
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
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScriptEngineDebuggerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QScriptEngineDebugger) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptEngineDebugger_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQScriptEngineDebugger_ChildEvent
func callbackQScriptEngineDebugger_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
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
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
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
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
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
		signal.(func())()
	} else {
		NewQScriptEngineDebuggerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QScriptEngineDebugger) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQScriptEngineDebugger_Destroyed
func callbackQScriptEngineDebugger_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQScriptEngineDebugger_DisconnectNotify
func callbackQScriptEngineDebugger_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScriptEngineDebuggerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScriptEngineDebugger) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScriptEngineDebugger_ObjectNameChanged
func callbackQScriptEngineDebugger_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtScriptTools_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQScriptEngineDebugger_TimerEvent
func callbackQScriptEngineDebugger_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScriptEngineDebuggerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QScriptEngineDebugger) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}
