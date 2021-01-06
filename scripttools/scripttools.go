// +build !minimal

package scripttools

import (
	"github.com/dev-drprasad/qt/core"
	"github.com/dev-drprasad/qt/internal"
	"github.com/dev-drprasad/qt/script"
	"github.com/dev-drprasad/qt/widgets"
	"unsafe"
)

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

func (n *QScriptEngineDebugger) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QScriptEngineDebugger) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQScriptEngineDebuggerFromPointer(ptr unsafe.Pointer) (n *QScriptEngineDebugger) {
	n = new(QScriptEngineDebugger)
	n.InitFromInternal(uintptr(ptr), "scripttools.QScriptEngineDebugger")
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

	return internal.CallLocalFunction([]interface{}{"", "", "scripttools.NewQScriptEngineDebugger", "", parent}).(*QScriptEngineDebugger)
}

func (ptr *QScriptEngineDebugger) Action(action QScriptEngineDebugger__DebuggerAction) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Action", action}).(*widgets.QAction)
}

func (ptr *QScriptEngineDebugger) AttachTo(engine script.QScriptEngine_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AttachTo", engine})
}

func (ptr *QScriptEngineDebugger) AutoShowStandardWindow() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AutoShowStandardWindow"}).(bool)
}

func (ptr *QScriptEngineDebugger) CreateStandardMenu(parent widgets.QWidget_ITF) *widgets.QMenu {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateStandardMenu", parent}).(*widgets.QMenu)
}

func (ptr *QScriptEngineDebugger) CreateStandardToolBar(parent widgets.QWidget_ITF) *widgets.QToolBar {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateStandardToolBar", parent}).(*widgets.QToolBar)
}

func (ptr *QScriptEngineDebugger) Detach() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Detach"})
}

func (ptr *QScriptEngineDebugger) ConnectEvaluationResumed(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectEvaluationResumed", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScriptEngineDebugger) DisconnectEvaluationResumed() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectEvaluationResumed"})
}

func (ptr *QScriptEngineDebugger) EvaluationResumed() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EvaluationResumed"})
}

func (ptr *QScriptEngineDebugger) ConnectEvaluationSuspended(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectEvaluationSuspended", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScriptEngineDebugger) DisconnectEvaluationSuspended() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectEvaluationSuspended"})
}

func (ptr *QScriptEngineDebugger) EvaluationSuspended() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EvaluationSuspended"})
}

func (ptr *QScriptEngineDebugger) SetAutoShowStandardWindow(autoShow bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAutoShowStandardWindow", autoShow})
}

func (ptr *QScriptEngineDebugger) StandardWindow() *widgets.QMainWindow {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StandardWindow"}).(*widgets.QMainWindow)
}

func (ptr *QScriptEngineDebugger) State() QScriptEngineDebugger__DebuggerState {

	return QScriptEngineDebugger__DebuggerState(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "State"}).(float64))
}

func (ptr *QScriptEngineDebugger) Widget(widget QScriptEngineDebugger__DebuggerWidget) *widgets.QWidget {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Widget", widget}).(*widgets.QWidget)
}

func (ptr *QScriptEngineDebugger) ConnectDestroyQScriptEngineDebugger(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQScriptEngineDebugger", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QScriptEngineDebugger) DisconnectDestroyQScriptEngineDebugger() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQScriptEngineDebugger"})
}

func (ptr *QScriptEngineDebugger) DestroyQScriptEngineDebugger() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQScriptEngineDebugger"})
}

func (ptr *QScriptEngineDebugger) DestroyQScriptEngineDebuggerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQScriptEngineDebuggerDefault"})
}

func (ptr *QScriptEngineDebugger) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QScriptEngineDebugger) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QScriptEngineDebugger) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QScriptEngineDebugger) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QScriptEngineDebugger) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QScriptEngineDebugger) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QScriptEngineDebugger) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QScriptEngineDebugger) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QScriptEngineDebugger) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QScriptEngineDebugger) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QScriptEngineDebugger) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QScriptEngineDebugger) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QScriptEngineDebugger) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QScriptEngineDebugger) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QScriptEngineDebugger) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QScriptEngineDebugger) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QScriptEngineDebugger) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QScriptEngineDebugger) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QScriptEngineDebugger) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QScriptEngineDebugger) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QScriptEngineDebugger) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

func init() {
	internal.ConstructorTable["scripttools.QScriptEngineDebugger"] = NewQScriptEngineDebuggerFromPointer
}
