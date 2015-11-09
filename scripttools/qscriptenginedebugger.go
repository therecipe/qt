package scripttools

//#include "qscriptenginedebugger.h"
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
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QScriptEngineDebugger_" + qt.RandomIdentifier())
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
	return NewQScriptEngineDebuggerFromPointer(C.QScriptEngineDebugger_NewQScriptEngineDebugger(core.PointerFromQObject(parent)))
}

func (ptr *QScriptEngineDebugger) Action(action QScriptEngineDebugger__DebuggerAction) *widgets.QAction {
	if ptr.Pointer() != nil {
		return widgets.NewQActionFromPointer(C.QScriptEngineDebugger_Action(ptr.Pointer(), C.int(action)))
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
		return C.QScriptEngineDebugger_AutoShowStandardWindow(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptEngineDebugger) CreateStandardMenu(parent widgets.QWidget_ITF) *widgets.QMenu {
	if ptr.Pointer() != nil {
		return widgets.NewQMenuFromPointer(C.QScriptEngineDebugger_CreateStandardMenu(ptr.Pointer(), widgets.PointerFromQWidget(parent)))
	}
	return nil
}

func (ptr *QScriptEngineDebugger) CreateStandardToolBar(parent widgets.QWidget_ITF) *widgets.QToolBar {
	if ptr.Pointer() != nil {
		return widgets.NewQToolBarFromPointer(C.QScriptEngineDebugger_CreateStandardToolBar(ptr.Pointer(), widgets.PointerFromQWidget(parent)))
	}
	return nil
}

func (ptr *QScriptEngineDebugger) Detach() {
	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_Detach(ptr.Pointer())
	}
}

func (ptr *QScriptEngineDebugger) ConnectEvaluationResumed(f func()) {
	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_ConnectEvaluationResumed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "evaluationResumed", f)
	}
}

func (ptr *QScriptEngineDebugger) DisconnectEvaluationResumed() {
	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_DisconnectEvaluationResumed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "evaluationResumed")
	}
}

//export callbackQScriptEngineDebuggerEvaluationResumed
func callbackQScriptEngineDebuggerEvaluationResumed(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "evaluationResumed").(func())()
}

func (ptr *QScriptEngineDebugger) ConnectEvaluationSuspended(f func()) {
	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_ConnectEvaluationSuspended(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "evaluationSuspended", f)
	}
}

func (ptr *QScriptEngineDebugger) DisconnectEvaluationSuspended() {
	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_DisconnectEvaluationSuspended(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "evaluationSuspended")
	}
}

//export callbackQScriptEngineDebuggerEvaluationSuspended
func callbackQScriptEngineDebuggerEvaluationSuspended(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "evaluationSuspended").(func())()
}

func (ptr *QScriptEngineDebugger) SetAutoShowStandardWindow(autoShow bool) {
	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_SetAutoShowStandardWindow(ptr.Pointer(), C.int(qt.GoBoolToInt(autoShow)))
	}
}

func (ptr *QScriptEngineDebugger) StandardWindow() *widgets.QMainWindow {
	if ptr.Pointer() != nil {
		return widgets.NewQMainWindowFromPointer(C.QScriptEngineDebugger_StandardWindow(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptEngineDebugger) Widget(widget QScriptEngineDebugger__DebuggerWidget) *widgets.QWidget {
	if ptr.Pointer() != nil {
		return widgets.NewQWidgetFromPointer(C.QScriptEngineDebugger_Widget(ptr.Pointer(), C.int(widget)))
	}
	return nil
}

func (ptr *QScriptEngineDebugger) DestroyQScriptEngineDebugger() {
	if ptr.Pointer() != nil {
		C.QScriptEngineDebugger_DestroyQScriptEngineDebugger(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
