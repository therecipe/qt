package scripttools

//#include "scripttools.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/script"
	"github.com/therecipe/qt/widgets"
	"unsafe"
)

type QScriptEngineDebugger struct {
	core.QObject
}

type QScriptEngineDebugger_ITF interface {
	core.QObject_ITF
	QScriptEngineDebugger_PTR() *QScriptEngineDebugger
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
	for len(n.ObjectName()) < len("QScriptEngineDebugger_") {
		n.SetObjectName("QScriptEngineDebugger_" + qt.Identifier())
	}
	return n
}

func (ptr *QScriptEngineDebugger) QScriptEngineDebugger_PTR() *QScriptEngineDebugger {
	return ptr
}

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

func NewQScriptEngineDebugger(parent core.QObject_ITF) *QScriptEngineDebugger {
	defer qt.Recovering("QScriptEngineDebugger::QScriptEngineDebugger")

	return NewQScriptEngineDebuggerFromPointer(C.QScriptEngineDebugger_NewQScriptEngineDebugger(core.PointerFromQObject(parent)))
}

func (ptr *QScriptEngineDebugger) Action(action QScriptEngineDebugger__DebuggerAction) *widgets.QAction {
	defer qt.Recovering("QScriptEngineDebugger::action")

	if ptr.Pointer() != nil {
		return widgets.NewQActionFromPointer(C.QScriptEngineDebugger_Action(ptr.Pointer(), C.int(action)))
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
		return widgets.NewQMenuFromPointer(C.QScriptEngineDebugger_CreateStandardMenu(ptr.Pointer(), widgets.PointerFromQWidget(parent)))
	}
	return nil
}

func (ptr *QScriptEngineDebugger) CreateStandardToolBar(parent widgets.QWidget_ITF) *widgets.QToolBar {
	defer qt.Recovering("QScriptEngineDebugger::createStandardToolBar")

	if ptr.Pointer() != nil {
		return widgets.NewQToolBarFromPointer(C.QScriptEngineDebugger_CreateStandardToolBar(ptr.Pointer(), widgets.PointerFromQWidget(parent)))
	}
	return nil
}

func (ptr *QScriptEngineDebugger) Detach() {
	defer qt.Recovering("QScriptEngineDebugger::detach")

	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_Detach(ptr.Pointer())
	}
}

func (ptr *QScriptEngineDebugger) ConnectEvaluationResumed(f func()) {
	defer qt.Recovering("connect QScriptEngineDebugger::evaluationResumed")

	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_ConnectEvaluationResumed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "evaluationResumed", f)
	}
}

func (ptr *QScriptEngineDebugger) DisconnectEvaluationResumed() {
	defer qt.Recovering("disconnect QScriptEngineDebugger::evaluationResumed")

	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_DisconnectEvaluationResumed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "evaluationResumed")
	}
}

//export callbackQScriptEngineDebuggerEvaluationResumed
func callbackQScriptEngineDebuggerEvaluationResumed(ptrName *C.char) {
	defer qt.Recovering("callback QScriptEngineDebugger::evaluationResumed")

	if signal := qt.GetSignal(C.GoString(ptrName), "evaluationResumed"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QScriptEngineDebugger) ConnectEvaluationSuspended(f func()) {
	defer qt.Recovering("connect QScriptEngineDebugger::evaluationSuspended")

	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_ConnectEvaluationSuspended(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "evaluationSuspended", f)
	}
}

func (ptr *QScriptEngineDebugger) DisconnectEvaluationSuspended() {
	defer qt.Recovering("disconnect QScriptEngineDebugger::evaluationSuspended")

	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_DisconnectEvaluationSuspended(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "evaluationSuspended")
	}
}

//export callbackQScriptEngineDebuggerEvaluationSuspended
func callbackQScriptEngineDebuggerEvaluationSuspended(ptrName *C.char) {
	defer qt.Recovering("callback QScriptEngineDebugger::evaluationSuspended")

	if signal := qt.GetSignal(C.GoString(ptrName), "evaluationSuspended"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QScriptEngineDebugger) SetAutoShowStandardWindow(autoShow bool) {
	defer qt.Recovering("QScriptEngineDebugger::setAutoShowStandardWindow")

	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_SetAutoShowStandardWindow(ptr.Pointer(), C.int(qt.GoBoolToInt(autoShow)))
	}
}

func (ptr *QScriptEngineDebugger) StandardWindow() *widgets.QMainWindow {
	defer qt.Recovering("QScriptEngineDebugger::standardWindow")

	if ptr.Pointer() != nil {
		return widgets.NewQMainWindowFromPointer(C.QScriptEngineDebugger_StandardWindow(ptr.Pointer()))
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
		return widgets.NewQWidgetFromPointer(C.QScriptEngineDebugger_Widget(ptr.Pointer(), C.int(widget)))
	}
	return nil
}

func (ptr *QScriptEngineDebugger) DestroyQScriptEngineDebugger() {
	defer qt.Recovering("QScriptEngineDebugger::~QScriptEngineDebugger")

	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_DestroyQScriptEngineDebugger(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScriptEngineDebugger) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QScriptEngineDebugger::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QScriptEngineDebugger) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QScriptEngineDebugger::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQScriptEngineDebuggerTimerEvent
func callbackQScriptEngineDebuggerTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScriptEngineDebugger::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QScriptEngineDebugger) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QScriptEngineDebugger::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QScriptEngineDebugger) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QScriptEngineDebugger::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQScriptEngineDebuggerChildEvent
func callbackQScriptEngineDebuggerChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScriptEngineDebugger::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QScriptEngineDebugger) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QScriptEngineDebugger::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QScriptEngineDebugger) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QScriptEngineDebugger::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQScriptEngineDebuggerCustomEvent
func callbackQScriptEngineDebuggerCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScriptEngineDebugger::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
