//go:build !minimal
// +build !minimal

package printsupport

import (
	"unsafe"

	"github.com/akiyosi/qt/core"
	"github.com/akiyosi/qt/gui"
	"github.com/akiyosi/qt/internal"
	"github.com/akiyosi/qt/widgets"
)

type QAbstractPrintDialog struct {
	widgets.QDialog
}

type QAbstractPrintDialog_ITF interface {
	widgets.QDialog_ITF
	QAbstractPrintDialog_PTR() *QAbstractPrintDialog
}

func (ptr *QAbstractPrintDialog) QAbstractPrintDialog_PTR() *QAbstractPrintDialog {
	return ptr
}

func (ptr *QAbstractPrintDialog) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDialog_PTR().Pointer()
	}
	return nil
}

func (ptr *QAbstractPrintDialog) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDialog_PTR().SetPointer(p)
	}
}

func PointerFromQAbstractPrintDialog(ptr QAbstractPrintDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractPrintDialog_PTR().Pointer()
	}
	return nil
}

func (n *QAbstractPrintDialog) InitFromInternal(ptr uintptr, name string) {
	n.QDialog_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QAbstractPrintDialog) ClassNameInternalF() string {
	return n.QDialog_PTR().ClassNameInternalF()
}

func NewQAbstractPrintDialogFromPointer(ptr unsafe.Pointer) (n *QAbstractPrintDialog) {
	n = new(QAbstractPrintDialog)
	n.InitFromInternal(uintptr(ptr), "printsupport.QAbstractPrintDialog")
	return
}

func (ptr *QAbstractPrintDialog) DestroyQAbstractPrintDialog() {
}

//go:generate stringer -type=QAbstractPrintDialog__PrintDialogOption
//QAbstractPrintDialog::PrintDialogOption
type QAbstractPrintDialog__PrintDialogOption int64

const (
	QAbstractPrintDialog__None               QAbstractPrintDialog__PrintDialogOption = QAbstractPrintDialog__PrintDialogOption(0x0000)
	QAbstractPrintDialog__PrintToFile        QAbstractPrintDialog__PrintDialogOption = QAbstractPrintDialog__PrintDialogOption(0x0001)
	QAbstractPrintDialog__PrintSelection     QAbstractPrintDialog__PrintDialogOption = QAbstractPrintDialog__PrintDialogOption(0x0002)
	QAbstractPrintDialog__PrintPageRange     QAbstractPrintDialog__PrintDialogOption = QAbstractPrintDialog__PrintDialogOption(0x0004)
	QAbstractPrintDialog__PrintShowPageSize  QAbstractPrintDialog__PrintDialogOption = QAbstractPrintDialog__PrintDialogOption(0x0008)
	QAbstractPrintDialog__PrintCollateCopies QAbstractPrintDialog__PrintDialogOption = QAbstractPrintDialog__PrintDialogOption(0x0010)
	QAbstractPrintDialog__DontUseSheet       QAbstractPrintDialog__PrintDialogOption = QAbstractPrintDialog__PrintDialogOption(0x0020)
	QAbstractPrintDialog__PrintCurrentPage   QAbstractPrintDialog__PrintDialogOption = QAbstractPrintDialog__PrintDialogOption(0x0040)
)

//go:generate stringer -type=QAbstractPrintDialog__PrintRange
//QAbstractPrintDialog::PrintRange
type QAbstractPrintDialog__PrintRange int64

const (
	QAbstractPrintDialog__AllPages    QAbstractPrintDialog__PrintRange = QAbstractPrintDialog__PrintRange(0)
	QAbstractPrintDialog__Selection   QAbstractPrintDialog__PrintRange = QAbstractPrintDialog__PrintRange(1)
	QAbstractPrintDialog__PageRange   QAbstractPrintDialog__PrintRange = QAbstractPrintDialog__PrintRange(2)
	QAbstractPrintDialog__CurrentPage QAbstractPrintDialog__PrintRange = QAbstractPrintDialog__PrintRange(3)
)

func NewQAbstractPrintDialog(printer QPrinter_ITF, parent widgets.QWidget_ITF) *QAbstractPrintDialog {

	return internal.CallLocalFunction([]interface{}{"", "", "printsupport.NewQAbstractPrintDialog", "", printer, parent}).(*QAbstractPrintDialog)
}

func (ptr *QAbstractPrintDialog) ConnectExec(f func() int) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectExec", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractPrintDialog) DisconnectExec() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectExec"})
}

func (ptr *QAbstractPrintDialog) Exec() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Exec"}).(float64))
}

func (ptr *QAbstractPrintDialog) FromPage() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FromPage"}).(float64))
}

func (ptr *QAbstractPrintDialog) MaxPage() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaxPage"}).(float64))
}

func (ptr *QAbstractPrintDialog) MinPage() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinPage"}).(float64))
}

func (ptr *QAbstractPrintDialog) PrintRange() QAbstractPrintDialog__PrintRange {

	return QAbstractPrintDialog__PrintRange(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrintRange"}).(float64))
}

func (ptr *QAbstractPrintDialog) Printer() *QPrinter {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Printer"}).(*QPrinter)
}

func (ptr *QAbstractPrintDialog) SetFromTo(from int, to int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFromTo", from, to})
}

func (ptr *QAbstractPrintDialog) SetMinMax(min int, max int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMinMax", min, max})
}

func (ptr *QAbstractPrintDialog) SetOptionTabs(tabs []*widgets.QWidget) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOptionTabs", tabs})
}

func (ptr *QAbstractPrintDialog) SetPrintRange(ran QAbstractPrintDialog__PrintRange) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPrintRange", ran})
}

func (ptr *QAbstractPrintDialog) ToPage() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToPage"}).(float64))
}

func (ptr *QAbstractPrintDialog) __setOptionTabs_tabs_atList(i int) *widgets.QWidget {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setOptionTabs_tabs_atList", i}).(*widgets.QWidget)
}

func (ptr *QAbstractPrintDialog) __setOptionTabs_tabs_setList(i widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setOptionTabs_tabs_setList", i})
}

func (ptr *QAbstractPrintDialog) __setOptionTabs_tabs_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setOptionTabs_tabs_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractPrintDialog) __actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_atList", i}).(*widgets.QAction)
}

func (ptr *QAbstractPrintDialog) __actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_setList", i})
}

func (ptr *QAbstractPrintDialog) __actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractPrintDialog) __addActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QAbstractPrintDialog) __addActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_setList", i})
}

func (ptr *QAbstractPrintDialog) __addActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractPrintDialog) __insertActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QAbstractPrintDialog) __insertActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_setList", i})
}

func (ptr *QAbstractPrintDialog) __insertActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractPrintDialog) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QAbstractPrintDialog) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QAbstractPrintDialog) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractPrintDialog) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QAbstractPrintDialog) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QAbstractPrintDialog) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractPrintDialog) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QAbstractPrintDialog) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QAbstractPrintDialog) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractPrintDialog) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QAbstractPrintDialog) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QAbstractPrintDialog) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QAbstractPrintDialog) AcceptDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AcceptDefault"})
}

func (ptr *QAbstractPrintDialog) CloseEventDefault(e gui.QCloseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseEventDefault", e})
}

func (ptr *QAbstractPrintDialog) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContextMenuEventDefault", e})
}

func (ptr *QAbstractPrintDialog) DoneDefault(r int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DoneDefault", r})
}

func (ptr *QAbstractPrintDialog) EventFilterDefault(o core.QObject_ITF, e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", o, e}).(bool)
}

func (ptr *QAbstractPrintDialog) KeyPressEventDefault(e gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyPressEventDefault", e})
}

func (ptr *QAbstractPrintDialog) MinimumSizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinimumSizeHintDefault"}).(*core.QSize)
}

func (ptr *QAbstractPrintDialog) OpenDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OpenDefault"})
}

func (ptr *QAbstractPrintDialog) RejectDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RejectDefault"})
}

func (ptr *QAbstractPrintDialog) ResizeEventDefault(vqr gui.QResizeEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResizeEventDefault", vqr})
}

func (ptr *QAbstractPrintDialog) SetVisibleDefault(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVisibleDefault", visible})
}

func (ptr *QAbstractPrintDialog) ShowEventDefault(event gui.QShowEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowEventDefault", event})
}

func (ptr *QAbstractPrintDialog) SizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeHintDefault"}).(*core.QSize)
}

func (ptr *QAbstractPrintDialog) ActionEventDefault(event gui.QActionEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActionEventDefault", event})
}

func (ptr *QAbstractPrintDialog) ChangeEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChangeEventDefault", event})
}

func (ptr *QAbstractPrintDialog) CloseDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"}).(bool)
}

func (ptr *QAbstractPrintDialog) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragEnterEventDefault", event})
}

func (ptr *QAbstractPrintDialog) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragLeaveEventDefault", event})
}

func (ptr *QAbstractPrintDialog) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragMoveEventDefault", event})
}

func (ptr *QAbstractPrintDialog) DropEventDefault(event gui.QDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DropEventDefault", event})
}

func (ptr *QAbstractPrintDialog) EnterEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EnterEventDefault", event})
}

func (ptr *QAbstractPrintDialog) EventDefault(event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", event}).(bool)
}

func (ptr *QAbstractPrintDialog) FocusInEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusInEventDefault", event})
}

func (ptr *QAbstractPrintDialog) FocusNextPrevChildDefault(next bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusNextPrevChildDefault", next}).(bool)
}

func (ptr *QAbstractPrintDialog) FocusOutEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusOutEventDefault", event})
}

func (ptr *QAbstractPrintDialog) HasHeightForWidthDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasHeightForWidthDefault"}).(bool)
}

func (ptr *QAbstractPrintDialog) HeightForWidthDefault(w int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HeightForWidthDefault", w}).(float64))
}

func (ptr *QAbstractPrintDialog) HideDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideDefault"})
}

func (ptr *QAbstractPrintDialog) HideEventDefault(event gui.QHideEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideEventDefault", event})
}

func (ptr *QAbstractPrintDialog) InitPainterDefault(painter gui.QPainter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InitPainterDefault", painter})
}

func (ptr *QAbstractPrintDialog) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodEventDefault", event})
}

func (ptr *QAbstractPrintDialog) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodQueryDefault", query}).(*core.QVariant)
}

func (ptr *QAbstractPrintDialog) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyReleaseEventDefault", event})
}

func (ptr *QAbstractPrintDialog) LeaveEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LeaveEventDefault", event})
}

func (ptr *QAbstractPrintDialog) LowerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LowerDefault"})
}

func (ptr *QAbstractPrintDialog) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetricDefault", m}).(float64))
}

func (ptr *QAbstractPrintDialog) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseDoubleClickEventDefault", event})
}

func (ptr *QAbstractPrintDialog) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseMoveEventDefault", event})
}

func (ptr *QAbstractPrintDialog) MousePressEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MousePressEventDefault", event})
}

func (ptr *QAbstractPrintDialog) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseReleaseEventDefault", event})
}

func (ptr *QAbstractPrintDialog) MoveEventDefault(event gui.QMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveEventDefault", event})
}

func (ptr *QAbstractPrintDialog) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NativeEventDefault", eventType, message, result}).(bool)
}

func (ptr *QAbstractPrintDialog) PaintEngineDefault() *gui.QPaintEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEngineDefault"}).(*gui.QPaintEngine)
}

func (ptr *QAbstractPrintDialog) PaintEventDefault(event gui.QPaintEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEventDefault", event})
}

func (ptr *QAbstractPrintDialog) RaiseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RaiseDefault"})
}

func (ptr *QAbstractPrintDialog) RepaintDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RepaintDefault"})
}

func (ptr *QAbstractPrintDialog) SetDisabledDefault(disable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDisabledDefault", disable})
}

func (ptr *QAbstractPrintDialog) SetEnabledDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetEnabledDefault", vbo})
}

func (ptr *QAbstractPrintDialog) SetFocus2Default() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFocus2Default"})
}

func (ptr *QAbstractPrintDialog) SetHiddenDefault(hidden bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHiddenDefault", hidden})
}

func (ptr *QAbstractPrintDialog) SetStyleSheetDefault(styleSheet string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetStyleSheetDefault", styleSheet})
}

func (ptr *QAbstractPrintDialog) SetWindowModifiedDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowModifiedDefault", vbo})
}

func (ptr *QAbstractPrintDialog) SetWindowTitleDefault(vqs string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowTitleDefault", vqs})
}

func (ptr *QAbstractPrintDialog) ShowDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowDefault"})
}

func (ptr *QAbstractPrintDialog) ShowFullScreenDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowFullScreenDefault"})
}

func (ptr *QAbstractPrintDialog) ShowMaximizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMaximizedDefault"})
}

func (ptr *QAbstractPrintDialog) ShowMinimizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMinimizedDefault"})
}

func (ptr *QAbstractPrintDialog) ShowNormalDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowNormalDefault"})
}

func (ptr *QAbstractPrintDialog) TabletEventDefault(event gui.QTabletEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TabletEventDefault", event})
}

func (ptr *QAbstractPrintDialog) UpdateDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateDefault"})
}

func (ptr *QAbstractPrintDialog) UpdateMicroFocusDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateMicroFocusDefault"})
}

func (ptr *QAbstractPrintDialog) WheelEventDefault(event gui.QWheelEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WheelEventDefault", event})
}

func (ptr *QAbstractPrintDialog) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QAbstractPrintDialog) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QAbstractPrintDialog) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QAbstractPrintDialog) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QAbstractPrintDialog) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QAbstractPrintDialog) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QAbstractPrintDialog) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QPageSetupDialog struct {
	widgets.QDialog
}

type QPageSetupDialog_ITF interface {
	widgets.QDialog_ITF
	QPageSetupDialog_PTR() *QPageSetupDialog
}

func (ptr *QPageSetupDialog) QPageSetupDialog_PTR() *QPageSetupDialog {
	return ptr
}

func (ptr *QPageSetupDialog) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDialog_PTR().Pointer()
	}
	return nil
}

func (ptr *QPageSetupDialog) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDialog_PTR().SetPointer(p)
	}
}

func PointerFromQPageSetupDialog(ptr QPageSetupDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPageSetupDialog_PTR().Pointer()
	}
	return nil
}

func (n *QPageSetupDialog) InitFromInternal(ptr uintptr, name string) {
	n.QDialog_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QPageSetupDialog) ClassNameInternalF() string {
	return n.QDialog_PTR().ClassNameInternalF()
}

func NewQPageSetupDialogFromPointer(ptr unsafe.Pointer) (n *QPageSetupDialog) {
	n = new(QPageSetupDialog)
	n.InitFromInternal(uintptr(ptr), "printsupport.QPageSetupDialog")
	return
}
func NewQPageSetupDialog(printer QPrinter_ITF, parent widgets.QWidget_ITF) *QPageSetupDialog {

	return internal.CallLocalFunction([]interface{}{"", "", "printsupport.NewQPageSetupDialog", "", printer, parent}).(*QPageSetupDialog)
}

func NewQPageSetupDialog2(parent widgets.QWidget_ITF) *QPageSetupDialog {

	return internal.CallLocalFunction([]interface{}{"", "", "printsupport.NewQPageSetupDialog2", "", parent}).(*QPageSetupDialog)
}

func (ptr *QPageSetupDialog) ConnectDone(f func(result int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDone", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPageSetupDialog) DisconnectDone() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDone"})
}

func (ptr *QPageSetupDialog) Done(result int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Done", result})
}

func (ptr *QPageSetupDialog) DoneDefault(result int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DoneDefault", result})
}

func (ptr *QPageSetupDialog) ConnectExec(f func() int) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectExec", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPageSetupDialog) DisconnectExec() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectExec"})
}

func (ptr *QPageSetupDialog) Exec() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Exec"}).(float64))
}

func (ptr *QPageSetupDialog) ExecDefault() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExecDefault"}).(float64))
}

func (ptr *QPageSetupDialog) Printer() *QPrinter {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Printer"}).(*QPrinter)
}

func (ptr *QPageSetupDialog) ConnectDestroyQPageSetupDialog(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQPageSetupDialog", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPageSetupDialog) DisconnectDestroyQPageSetupDialog() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQPageSetupDialog"})
}

func (ptr *QPageSetupDialog) DestroyQPageSetupDialog() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQPageSetupDialog"})
}

func (ptr *QPageSetupDialog) DestroyQPageSetupDialogDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQPageSetupDialogDefault"})
}

func (ptr *QPageSetupDialog) __actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_atList", i}).(*widgets.QAction)
}

func (ptr *QPageSetupDialog) __actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_setList", i})
}

func (ptr *QPageSetupDialog) __actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_newList"}).(unsafe.Pointer)
}

func (ptr *QPageSetupDialog) __addActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QPageSetupDialog) __addActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_setList", i})
}

func (ptr *QPageSetupDialog) __addActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QPageSetupDialog) __insertActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QPageSetupDialog) __insertActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_setList", i})
}

func (ptr *QPageSetupDialog) __insertActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QPageSetupDialog) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QPageSetupDialog) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QPageSetupDialog) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QPageSetupDialog) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QPageSetupDialog) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QPageSetupDialog) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QPageSetupDialog) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QPageSetupDialog) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QPageSetupDialog) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QPageSetupDialog) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QPageSetupDialog) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QPageSetupDialog) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QPageSetupDialog) AcceptDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AcceptDefault"})
}

func (ptr *QPageSetupDialog) CloseEventDefault(e gui.QCloseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseEventDefault", e})
}

func (ptr *QPageSetupDialog) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContextMenuEventDefault", e})
}

func (ptr *QPageSetupDialog) EventFilterDefault(o core.QObject_ITF, e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", o, e}).(bool)
}

func (ptr *QPageSetupDialog) KeyPressEventDefault(e gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyPressEventDefault", e})
}

func (ptr *QPageSetupDialog) MinimumSizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinimumSizeHintDefault"}).(*core.QSize)
}

func (ptr *QPageSetupDialog) RejectDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RejectDefault"})
}

func (ptr *QPageSetupDialog) ResizeEventDefault(vqr gui.QResizeEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResizeEventDefault", vqr})
}

func (ptr *QPageSetupDialog) SetVisibleDefault(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVisibleDefault", visible})
}

func (ptr *QPageSetupDialog) ShowEventDefault(event gui.QShowEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowEventDefault", event})
}

func (ptr *QPageSetupDialog) SizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeHintDefault"}).(*core.QSize)
}

func (ptr *QPageSetupDialog) ActionEventDefault(event gui.QActionEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActionEventDefault", event})
}

func (ptr *QPageSetupDialog) ChangeEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChangeEventDefault", event})
}

func (ptr *QPageSetupDialog) CloseDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"}).(bool)
}

func (ptr *QPageSetupDialog) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragEnterEventDefault", event})
}

func (ptr *QPageSetupDialog) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragLeaveEventDefault", event})
}

func (ptr *QPageSetupDialog) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragMoveEventDefault", event})
}

func (ptr *QPageSetupDialog) DropEventDefault(event gui.QDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DropEventDefault", event})
}

func (ptr *QPageSetupDialog) EnterEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EnterEventDefault", event})
}

func (ptr *QPageSetupDialog) EventDefault(event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", event}).(bool)
}

func (ptr *QPageSetupDialog) FocusInEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusInEventDefault", event})
}

func (ptr *QPageSetupDialog) FocusNextPrevChildDefault(next bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusNextPrevChildDefault", next}).(bool)
}

func (ptr *QPageSetupDialog) FocusOutEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusOutEventDefault", event})
}

func (ptr *QPageSetupDialog) HasHeightForWidthDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasHeightForWidthDefault"}).(bool)
}

func (ptr *QPageSetupDialog) HeightForWidthDefault(w int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HeightForWidthDefault", w}).(float64))
}

func (ptr *QPageSetupDialog) HideDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideDefault"})
}

func (ptr *QPageSetupDialog) HideEventDefault(event gui.QHideEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideEventDefault", event})
}

func (ptr *QPageSetupDialog) InitPainterDefault(painter gui.QPainter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InitPainterDefault", painter})
}

func (ptr *QPageSetupDialog) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodEventDefault", event})
}

func (ptr *QPageSetupDialog) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodQueryDefault", query}).(*core.QVariant)
}

func (ptr *QPageSetupDialog) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyReleaseEventDefault", event})
}

func (ptr *QPageSetupDialog) LeaveEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LeaveEventDefault", event})
}

func (ptr *QPageSetupDialog) LowerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LowerDefault"})
}

func (ptr *QPageSetupDialog) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetricDefault", m}).(float64))
}

func (ptr *QPageSetupDialog) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseDoubleClickEventDefault", event})
}

func (ptr *QPageSetupDialog) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseMoveEventDefault", event})
}

func (ptr *QPageSetupDialog) MousePressEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MousePressEventDefault", event})
}

func (ptr *QPageSetupDialog) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseReleaseEventDefault", event})
}

func (ptr *QPageSetupDialog) MoveEventDefault(event gui.QMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveEventDefault", event})
}

func (ptr *QPageSetupDialog) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NativeEventDefault", eventType, message, result}).(bool)
}

func (ptr *QPageSetupDialog) PaintEngineDefault() *gui.QPaintEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEngineDefault"}).(*gui.QPaintEngine)
}

func (ptr *QPageSetupDialog) PaintEventDefault(event gui.QPaintEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEventDefault", event})
}

func (ptr *QPageSetupDialog) RaiseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RaiseDefault"})
}

func (ptr *QPageSetupDialog) RepaintDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RepaintDefault"})
}

func (ptr *QPageSetupDialog) SetDisabledDefault(disable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDisabledDefault", disable})
}

func (ptr *QPageSetupDialog) SetEnabledDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetEnabledDefault", vbo})
}

func (ptr *QPageSetupDialog) SetFocus2Default() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFocus2Default"})
}

func (ptr *QPageSetupDialog) SetHiddenDefault(hidden bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHiddenDefault", hidden})
}

func (ptr *QPageSetupDialog) SetStyleSheetDefault(styleSheet string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetStyleSheetDefault", styleSheet})
}

func (ptr *QPageSetupDialog) SetWindowModifiedDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowModifiedDefault", vbo})
}

func (ptr *QPageSetupDialog) SetWindowTitleDefault(vqs string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowTitleDefault", vqs})
}

func (ptr *QPageSetupDialog) ShowDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowDefault"})
}

func (ptr *QPageSetupDialog) ShowFullScreenDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowFullScreenDefault"})
}

func (ptr *QPageSetupDialog) ShowMaximizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMaximizedDefault"})
}

func (ptr *QPageSetupDialog) ShowMinimizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMinimizedDefault"})
}

func (ptr *QPageSetupDialog) ShowNormalDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowNormalDefault"})
}

func (ptr *QPageSetupDialog) TabletEventDefault(event gui.QTabletEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TabletEventDefault", event})
}

func (ptr *QPageSetupDialog) UpdateDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateDefault"})
}

func (ptr *QPageSetupDialog) UpdateMicroFocusDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateMicroFocusDefault"})
}

func (ptr *QPageSetupDialog) WheelEventDefault(event gui.QWheelEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WheelEventDefault", event})
}

func (ptr *QPageSetupDialog) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QPageSetupDialog) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QPageSetupDialog) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QPageSetupDialog) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QPageSetupDialog) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QPageSetupDialog) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QPageSetupDialog) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QPlatformPrintDevice struct {
	internal.Internal
}

type QPlatformPrintDevice_ITF interface {
	QPlatformPrintDevice_PTR() *QPlatformPrintDevice
}

func (ptr *QPlatformPrintDevice) QPlatformPrintDevice_PTR() *QPlatformPrintDevice {
	return ptr
}

func (ptr *QPlatformPrintDevice) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QPlatformPrintDevice) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQPlatformPrintDevice(ptr QPlatformPrintDevice_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlatformPrintDevice_PTR().Pointer()
	}
	return nil
}

func (n *QPlatformPrintDevice) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQPlatformPrintDeviceFromPointer(ptr unsafe.Pointer) (n *QPlatformPrintDevice) {
	n = new(QPlatformPrintDevice)
	n.InitFromInternal(uintptr(ptr), "printsupport.QPlatformPrintDevice")
	return
}

func (ptr *QPlatformPrintDevice) DestroyQPlatformPrintDevice() {
}

type QPlatformPrinterSupportPlugin struct {
	core.QObject
}

type QPlatformPrinterSupportPlugin_ITF interface {
	core.QObject_ITF
	QPlatformPrinterSupportPlugin_PTR() *QPlatformPrinterSupportPlugin
}

func (ptr *QPlatformPrinterSupportPlugin) QPlatformPrinterSupportPlugin_PTR() *QPlatformPrinterSupportPlugin {
	return ptr
}

func (ptr *QPlatformPrinterSupportPlugin) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QPlatformPrinterSupportPlugin) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQPlatformPrinterSupportPlugin(ptr QPlatformPrinterSupportPlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlatformPrinterSupportPlugin_PTR().Pointer()
	}
	return nil
}

func (n *QPlatformPrinterSupportPlugin) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QPlatformPrinterSupportPlugin) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQPlatformPrinterSupportPluginFromPointer(ptr unsafe.Pointer) (n *QPlatformPrinterSupportPlugin) {
	n = new(QPlatformPrinterSupportPlugin)
	n.InitFromInternal(uintptr(ptr), "printsupport.QPlatformPrinterSupportPlugin")
	return
}

func (ptr *QPlatformPrinterSupportPlugin) DestroyQPlatformPrinterSupportPlugin() {
}

type QPrintDialog struct {
	QAbstractPrintDialog
}

type QPrintDialog_ITF interface {
	QAbstractPrintDialog_ITF
	QPrintDialog_PTR() *QPrintDialog
}

func (ptr *QPrintDialog) QPrintDialog_PTR() *QPrintDialog {
	return ptr
}

func (ptr *QPrintDialog) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractPrintDialog_PTR().Pointer()
	}
	return nil
}

func (ptr *QPrintDialog) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractPrintDialog_PTR().SetPointer(p)
	}
}

func PointerFromQPrintDialog(ptr QPrintDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPrintDialog_PTR().Pointer()
	}
	return nil
}

func (n *QPrintDialog) InitFromInternal(ptr uintptr, name string) {
	n.QAbstractPrintDialog_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QPrintDialog) ClassNameInternalF() string {
	return n.QAbstractPrintDialog_PTR().ClassNameInternalF()
}

func NewQPrintDialogFromPointer(ptr unsafe.Pointer) (n *QPrintDialog) {
	n = new(QPrintDialog)
	n.InitFromInternal(uintptr(ptr), "printsupport.QPrintDialog")
	return
}
func NewQPrintDialog(printer QPrinter_ITF, parent widgets.QWidget_ITF) *QPrintDialog {

	return internal.CallLocalFunction([]interface{}{"", "", "printsupport.NewQPrintDialog", "", printer, parent}).(*QPrintDialog)
}

func (ptr *QPrintDialog) ConnectAccepted(f func(printer *QPrinter)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAccepted", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintDialog) DisconnectAccepted() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAccepted"})
}

func (ptr *QPrintDialog) Accepted(printer QPrinter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Accepted", printer})
}

func (ptr *QPrintDialog) ConnectDone(f func(result int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDone", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintDialog) DisconnectDone() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDone"})
}

func (ptr *QPrintDialog) Done(result int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Done", result})
}

func (ptr *QPrintDialog) DoneDefault(result int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DoneDefault", result})
}

func (ptr *QPrintDialog) ConnectExec(f func() int) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectExec", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintDialog) DisconnectExec() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectExec"})
}

func (ptr *QPrintDialog) Exec() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Exec"}).(float64))
}

func (ptr *QPrintDialog) ExecDefault() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExecDefault"}).(float64))
}

func (ptr *QPrintDialog) Open(receiver core.QObject_ITF, member string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Open", receiver, member})
}

func (ptr *QPrintDialog) Options() QAbstractPrintDialog__PrintDialogOption {

	return QAbstractPrintDialog__PrintDialogOption(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Options"}).(float64))
}

func (ptr *QPrintDialog) SetOption(option QAbstractPrintDialog__PrintDialogOption, on bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOption", option, on})
}

func (ptr *QPrintDialog) SetOptions(options QAbstractPrintDialog__PrintDialogOption) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOptions", options})
}

func (ptr *QPrintDialog) TestOption(option QAbstractPrintDialog__PrintDialogOption) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TestOption", option}).(bool)
}

func (ptr *QPrintDialog) ConnectDestroyQPrintDialog(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQPrintDialog", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintDialog) DisconnectDestroyQPrintDialog() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQPrintDialog"})
}

func (ptr *QPrintDialog) DestroyQPrintDialog() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQPrintDialog"})
}

func (ptr *QPrintDialog) DestroyQPrintDialogDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQPrintDialogDefault"})
}

type QPrintEngine struct {
	internal.Internal
}

type QPrintEngine_ITF interface {
	QPrintEngine_PTR() *QPrintEngine
}

func (ptr *QPrintEngine) QPrintEngine_PTR() *QPrintEngine {
	return ptr
}

func (ptr *QPrintEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QPrintEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQPrintEngine(ptr QPrintEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPrintEngine_PTR().Pointer()
	}
	return nil
}

func (n *QPrintEngine) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQPrintEngineFromPointer(ptr unsafe.Pointer) (n *QPrintEngine) {
	n = new(QPrintEngine)
	n.InitFromInternal(uintptr(ptr), "printsupport.QPrintEngine")
	return
}

//go:generate stringer -type=QPrintEngine__PrintEnginePropertyKey
//QPrintEngine::PrintEnginePropertyKey
type QPrintEngine__PrintEnginePropertyKey int64

const (
	QPrintEngine__PPK_CollateCopies          QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(0)
	QPrintEngine__PPK_ColorMode              QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(1)
	QPrintEngine__PPK_Creator                QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(2)
	QPrintEngine__PPK_DocumentName           QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(3)
	QPrintEngine__PPK_FullPage               QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(4)
	QPrintEngine__PPK_NumberOfCopies         QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(5)
	QPrintEngine__PPK_Orientation            QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(6)
	QPrintEngine__PPK_OutputFileName         QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(7)
	QPrintEngine__PPK_PageOrder              QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(8)
	QPrintEngine__PPK_PageRect               QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(9)
	QPrintEngine__PPK_PageSize               QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(10)
	QPrintEngine__PPK_PaperRect              QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(11)
	QPrintEngine__PPK_PaperSource            QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(12)
	QPrintEngine__PPK_PrinterName            QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(13)
	QPrintEngine__PPK_PrinterProgram         QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(14)
	QPrintEngine__PPK_Resolution             QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(15)
	QPrintEngine__PPK_SelectionOption        QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(16)
	QPrintEngine__PPK_SupportedResolutions   QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(17)
	QPrintEngine__PPK_WindowsPageSize        QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(18)
	QPrintEngine__PPK_FontEmbedding          QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(19)
	QPrintEngine__PPK_Duplex                 QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(20)
	QPrintEngine__PPK_PaperSources           QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(21)
	QPrintEngine__PPK_CustomPaperSize        QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(22)
	QPrintEngine__PPK_PageMargins            QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(23)
	QPrintEngine__PPK_CopyCount              QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(24)
	QPrintEngine__PPK_SupportsMultipleCopies QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(25)
	QPrintEngine__PPK_PaperName              QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(26)
	QPrintEngine__PPK_QPageSize              QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(27)
	QPrintEngine__PPK_QPageMargins           QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(28)
	QPrintEngine__PPK_QPageLayout            QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(29)
	QPrintEngine__PPK_PaperSize              QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(QPrintEngine__PPK_PageSize)
	QPrintEngine__PPK_CustomBase             QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(0xff00)
)

func (ptr *QPrintEngine) ConnectAbort(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAbort", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintEngine) DisconnectAbort() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAbort"})
}

func (ptr *QPrintEngine) Abort() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Abort"}).(bool)
}

func (ptr *QPrintEngine) ConnectMetric(f func(id gui.QPaintDevice__PaintDeviceMetric) int) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMetric", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintEngine) DisconnectMetric() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMetric"})
}

func (ptr *QPrintEngine) Metric(id gui.QPaintDevice__PaintDeviceMetric) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Metric", id}).(float64))
}

func (ptr *QPrintEngine) ConnectNewPage(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNewPage", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintEngine) DisconnectNewPage() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNewPage"})
}

func (ptr *QPrintEngine) NewPage() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NewPage"}).(bool)
}

func (ptr *QPrintEngine) ConnectPrinterState(f func() QPrinter__PrinterState) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPrinterState", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintEngine) DisconnectPrinterState() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPrinterState"})
}

func (ptr *QPrintEngine) PrinterState() QPrinter__PrinterState {

	return QPrinter__PrinterState(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrinterState"}).(float64))
}

func (ptr *QPrintEngine) ConnectProperty(f func(key QPrintEngine__PrintEnginePropertyKey) *core.QVariant) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintEngine) DisconnectProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectProperty"})
}

func (ptr *QPrintEngine) Property(key QPrintEngine__PrintEnginePropertyKey) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Property", key}).(*core.QVariant)
}

func (ptr *QPrintEngine) ConnectSetProperty(f func(key QPrintEngine__PrintEnginePropertyKey, value *core.QVariant)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintEngine) DisconnectSetProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetProperty"})
}

func (ptr *QPrintEngine) SetProperty(key QPrintEngine__PrintEnginePropertyKey, value core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetProperty", key, value})
}

func (ptr *QPrintEngine) ConnectDestroyQPrintEngine(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQPrintEngine", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintEngine) DisconnectDestroyQPrintEngine() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQPrintEngine"})
}

func (ptr *QPrintEngine) DestroyQPrintEngine() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQPrintEngine"})
}

func (ptr *QPrintEngine) DestroyQPrintEngineDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQPrintEngineDefault"})
}

type QPrintPreviewDialog struct {
	widgets.QDialog
}

type QPrintPreviewDialog_ITF interface {
	widgets.QDialog_ITF
	QPrintPreviewDialog_PTR() *QPrintPreviewDialog
}

func (ptr *QPrintPreviewDialog) QPrintPreviewDialog_PTR() *QPrintPreviewDialog {
	return ptr
}

func (ptr *QPrintPreviewDialog) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDialog_PTR().Pointer()
	}
	return nil
}

func (ptr *QPrintPreviewDialog) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDialog_PTR().SetPointer(p)
	}
}

func PointerFromQPrintPreviewDialog(ptr QPrintPreviewDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPrintPreviewDialog_PTR().Pointer()
	}
	return nil
}

func (n *QPrintPreviewDialog) InitFromInternal(ptr uintptr, name string) {
	n.QDialog_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QPrintPreviewDialog) ClassNameInternalF() string {
	return n.QDialog_PTR().ClassNameInternalF()
}

func NewQPrintPreviewDialogFromPointer(ptr unsafe.Pointer) (n *QPrintPreviewDialog) {
	n = new(QPrintPreviewDialog)
	n.InitFromInternal(uintptr(ptr), "printsupport.QPrintPreviewDialog")
	return
}
func NewQPrintPreviewDialog(printer QPrinter_ITF, parent widgets.QWidget_ITF, flags core.Qt__WindowType) *QPrintPreviewDialog {

	return internal.CallLocalFunction([]interface{}{"", "", "printsupport.NewQPrintPreviewDialog", "", printer, parent, flags}).(*QPrintPreviewDialog)
}

func NewQPrintPreviewDialog2(parent widgets.QWidget_ITF, flags core.Qt__WindowType) *QPrintPreviewDialog {

	return internal.CallLocalFunction([]interface{}{"", "", "printsupport.NewQPrintPreviewDialog2", "", parent, flags}).(*QPrintPreviewDialog)
}

func (ptr *QPrintPreviewDialog) ConnectDone(f func(result int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDone", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintPreviewDialog) DisconnectDone() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDone"})
}

func (ptr *QPrintPreviewDialog) Done(result int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Done", result})
}

func (ptr *QPrintPreviewDialog) DoneDefault(result int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DoneDefault", result})
}

func (ptr *QPrintPreviewDialog) ConnectPaintRequested(f func(printer *QPrinter)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPaintRequested", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintPreviewDialog) DisconnectPaintRequested() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPaintRequested"})
}

func (ptr *QPrintPreviewDialog) PaintRequested(printer QPrinter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintRequested", printer})
}

func (ptr *QPrintPreviewDialog) Printer() *QPrinter {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Printer"}).(*QPrinter)
}

func (ptr *QPrintPreviewDialog) SetVisibleDefault(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVisibleDefault", visible})
}

func (ptr *QPrintPreviewDialog) ConnectDestroyQPrintPreviewDialog(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQPrintPreviewDialog", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintPreviewDialog) DisconnectDestroyQPrintPreviewDialog() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQPrintPreviewDialog"})
}

func (ptr *QPrintPreviewDialog) DestroyQPrintPreviewDialog() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQPrintPreviewDialog"})
}

func (ptr *QPrintPreviewDialog) DestroyQPrintPreviewDialogDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQPrintPreviewDialogDefault"})
}

func (ptr *QPrintPreviewDialog) __actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_atList", i}).(*widgets.QAction)
}

func (ptr *QPrintPreviewDialog) __actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_setList", i})
}

func (ptr *QPrintPreviewDialog) __actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_newList"}).(unsafe.Pointer)
}

func (ptr *QPrintPreviewDialog) __addActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QPrintPreviewDialog) __addActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_setList", i})
}

func (ptr *QPrintPreviewDialog) __addActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QPrintPreviewDialog) __insertActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QPrintPreviewDialog) __insertActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_setList", i})
}

func (ptr *QPrintPreviewDialog) __insertActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QPrintPreviewDialog) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QPrintPreviewDialog) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QPrintPreviewDialog) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QPrintPreviewDialog) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QPrintPreviewDialog) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QPrintPreviewDialog) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QPrintPreviewDialog) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QPrintPreviewDialog) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QPrintPreviewDialog) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QPrintPreviewDialog) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QPrintPreviewDialog) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QPrintPreviewDialog) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QPrintPreviewDialog) AcceptDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AcceptDefault"})
}

func (ptr *QPrintPreviewDialog) CloseEventDefault(e gui.QCloseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseEventDefault", e})
}

func (ptr *QPrintPreviewDialog) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContextMenuEventDefault", e})
}

func (ptr *QPrintPreviewDialog) EventFilterDefault(o core.QObject_ITF, e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", o, e}).(bool)
}

func (ptr *QPrintPreviewDialog) ExecDefault() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExecDefault"}).(float64))
}

func (ptr *QPrintPreviewDialog) KeyPressEventDefault(e gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyPressEventDefault", e})
}

func (ptr *QPrintPreviewDialog) MinimumSizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinimumSizeHintDefault"}).(*core.QSize)
}

func (ptr *QPrintPreviewDialog) RejectDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RejectDefault"})
}

func (ptr *QPrintPreviewDialog) ResizeEventDefault(vqr gui.QResizeEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResizeEventDefault", vqr})
}

func (ptr *QPrintPreviewDialog) ShowEventDefault(event gui.QShowEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowEventDefault", event})
}

func (ptr *QPrintPreviewDialog) SizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeHintDefault"}).(*core.QSize)
}

func (ptr *QPrintPreviewDialog) ActionEventDefault(event gui.QActionEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActionEventDefault", event})
}

func (ptr *QPrintPreviewDialog) ChangeEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChangeEventDefault", event})
}

func (ptr *QPrintPreviewDialog) CloseDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"}).(bool)
}

func (ptr *QPrintPreviewDialog) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragEnterEventDefault", event})
}

func (ptr *QPrintPreviewDialog) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragLeaveEventDefault", event})
}

func (ptr *QPrintPreviewDialog) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragMoveEventDefault", event})
}

func (ptr *QPrintPreviewDialog) DropEventDefault(event gui.QDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DropEventDefault", event})
}

func (ptr *QPrintPreviewDialog) EnterEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EnterEventDefault", event})
}

func (ptr *QPrintPreviewDialog) EventDefault(event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", event}).(bool)
}

func (ptr *QPrintPreviewDialog) FocusInEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusInEventDefault", event})
}

func (ptr *QPrintPreviewDialog) FocusNextPrevChildDefault(next bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusNextPrevChildDefault", next}).(bool)
}

func (ptr *QPrintPreviewDialog) FocusOutEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusOutEventDefault", event})
}

func (ptr *QPrintPreviewDialog) HasHeightForWidthDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasHeightForWidthDefault"}).(bool)
}

func (ptr *QPrintPreviewDialog) HeightForWidthDefault(w int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HeightForWidthDefault", w}).(float64))
}

func (ptr *QPrintPreviewDialog) HideDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideDefault"})
}

func (ptr *QPrintPreviewDialog) HideEventDefault(event gui.QHideEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideEventDefault", event})
}

func (ptr *QPrintPreviewDialog) InitPainterDefault(painter gui.QPainter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InitPainterDefault", painter})
}

func (ptr *QPrintPreviewDialog) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodEventDefault", event})
}

func (ptr *QPrintPreviewDialog) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodQueryDefault", query}).(*core.QVariant)
}

func (ptr *QPrintPreviewDialog) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyReleaseEventDefault", event})
}

func (ptr *QPrintPreviewDialog) LeaveEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LeaveEventDefault", event})
}

func (ptr *QPrintPreviewDialog) LowerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LowerDefault"})
}

func (ptr *QPrintPreviewDialog) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetricDefault", m}).(float64))
}

func (ptr *QPrintPreviewDialog) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseDoubleClickEventDefault", event})
}

func (ptr *QPrintPreviewDialog) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseMoveEventDefault", event})
}

func (ptr *QPrintPreviewDialog) MousePressEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MousePressEventDefault", event})
}

func (ptr *QPrintPreviewDialog) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseReleaseEventDefault", event})
}

func (ptr *QPrintPreviewDialog) MoveEventDefault(event gui.QMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveEventDefault", event})
}

func (ptr *QPrintPreviewDialog) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NativeEventDefault", eventType, message, result}).(bool)
}

func (ptr *QPrintPreviewDialog) PaintEngineDefault() *gui.QPaintEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEngineDefault"}).(*gui.QPaintEngine)
}

func (ptr *QPrintPreviewDialog) PaintEventDefault(event gui.QPaintEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEventDefault", event})
}

func (ptr *QPrintPreviewDialog) RaiseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RaiseDefault"})
}

func (ptr *QPrintPreviewDialog) RepaintDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RepaintDefault"})
}

func (ptr *QPrintPreviewDialog) SetDisabledDefault(disable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDisabledDefault", disable})
}

func (ptr *QPrintPreviewDialog) SetEnabledDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetEnabledDefault", vbo})
}

func (ptr *QPrintPreviewDialog) SetFocus2Default() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFocus2Default"})
}

func (ptr *QPrintPreviewDialog) SetHiddenDefault(hidden bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHiddenDefault", hidden})
}

func (ptr *QPrintPreviewDialog) SetStyleSheetDefault(styleSheet string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetStyleSheetDefault", styleSheet})
}

func (ptr *QPrintPreviewDialog) SetWindowModifiedDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowModifiedDefault", vbo})
}

func (ptr *QPrintPreviewDialog) SetWindowTitleDefault(vqs string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowTitleDefault", vqs})
}

func (ptr *QPrintPreviewDialog) ShowDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowDefault"})
}

func (ptr *QPrintPreviewDialog) ShowFullScreenDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowFullScreenDefault"})
}

func (ptr *QPrintPreviewDialog) ShowMaximizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMaximizedDefault"})
}

func (ptr *QPrintPreviewDialog) ShowMinimizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMinimizedDefault"})
}

func (ptr *QPrintPreviewDialog) ShowNormalDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowNormalDefault"})
}

func (ptr *QPrintPreviewDialog) TabletEventDefault(event gui.QTabletEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TabletEventDefault", event})
}

func (ptr *QPrintPreviewDialog) UpdateDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateDefault"})
}

func (ptr *QPrintPreviewDialog) UpdateMicroFocusDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateMicroFocusDefault"})
}

func (ptr *QPrintPreviewDialog) WheelEventDefault(event gui.QWheelEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WheelEventDefault", event})
}

func (ptr *QPrintPreviewDialog) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QPrintPreviewDialog) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QPrintPreviewDialog) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QPrintPreviewDialog) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QPrintPreviewDialog) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QPrintPreviewDialog) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QPrintPreviewDialog) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QPrintPreviewWidget struct {
	widgets.QWidget
}

type QPrintPreviewWidget_ITF interface {
	widgets.QWidget_ITF
	QPrintPreviewWidget_PTR() *QPrintPreviewWidget
}

func (ptr *QPrintPreviewWidget) QPrintPreviewWidget_PTR() *QPrintPreviewWidget {
	return ptr
}

func (ptr *QPrintPreviewWidget) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QPrintPreviewWidget) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQPrintPreviewWidget(ptr QPrintPreviewWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPrintPreviewWidget_PTR().Pointer()
	}
	return nil
}

func (n *QPrintPreviewWidget) InitFromInternal(ptr uintptr, name string) {
	n.QWidget_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QPrintPreviewWidget) ClassNameInternalF() string {
	return n.QWidget_PTR().ClassNameInternalF()
}

func NewQPrintPreviewWidgetFromPointer(ptr unsafe.Pointer) (n *QPrintPreviewWidget) {
	n = new(QPrintPreviewWidget)
	n.InitFromInternal(uintptr(ptr), "printsupport.QPrintPreviewWidget")
	return
}

//go:generate stringer -type=QPrintPreviewWidget__ViewMode
//QPrintPreviewWidget::ViewMode
type QPrintPreviewWidget__ViewMode int64

const (
	QPrintPreviewWidget__SinglePageView  QPrintPreviewWidget__ViewMode = QPrintPreviewWidget__ViewMode(0)
	QPrintPreviewWidget__FacingPagesView QPrintPreviewWidget__ViewMode = QPrintPreviewWidget__ViewMode(1)
	QPrintPreviewWidget__AllPagesView    QPrintPreviewWidget__ViewMode = QPrintPreviewWidget__ViewMode(2)
)

//go:generate stringer -type=QPrintPreviewWidget__ZoomMode
//QPrintPreviewWidget::ZoomMode
type QPrintPreviewWidget__ZoomMode int64

const (
	QPrintPreviewWidget__CustomZoom QPrintPreviewWidget__ZoomMode = QPrintPreviewWidget__ZoomMode(0)
	QPrintPreviewWidget__FitToWidth QPrintPreviewWidget__ZoomMode = QPrintPreviewWidget__ZoomMode(1)
	QPrintPreviewWidget__FitInView  QPrintPreviewWidget__ZoomMode = QPrintPreviewWidget__ZoomMode(2)
)

func NewQPrintPreviewWidget(printer QPrinter_ITF, parent widgets.QWidget_ITF, flags core.Qt__WindowType) *QPrintPreviewWidget {

	return internal.CallLocalFunction([]interface{}{"", "", "printsupport.NewQPrintPreviewWidget", "", printer, parent, flags}).(*QPrintPreviewWidget)
}

func NewQPrintPreviewWidget2(parent widgets.QWidget_ITF, flags core.Qt__WindowType) *QPrintPreviewWidget {

	return internal.CallLocalFunction([]interface{}{"", "", "printsupport.NewQPrintPreviewWidget2", "", parent, flags}).(*QPrintPreviewWidget)
}

func (ptr *QPrintPreviewWidget) CurrentPage() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CurrentPage"}).(float64))
}

func (ptr *QPrintPreviewWidget) ConnectFitInView(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFitInView", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintPreviewWidget) DisconnectFitInView() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFitInView"})
}

func (ptr *QPrintPreviewWidget) FitInView() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FitInView"})
}

func (ptr *QPrintPreviewWidget) FitInViewDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FitInViewDefault"})
}

func (ptr *QPrintPreviewWidget) ConnectFitToWidth(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFitToWidth", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintPreviewWidget) DisconnectFitToWidth() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFitToWidth"})
}

func (ptr *QPrintPreviewWidget) FitToWidth() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FitToWidth"})
}

func (ptr *QPrintPreviewWidget) FitToWidthDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FitToWidthDefault"})
}

func (ptr *QPrintPreviewWidget) Orientation() QPrinter__Orientation {

	return QPrinter__Orientation(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Orientation"}).(float64))
}

func (ptr *QPrintPreviewWidget) PageCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PageCount"}).(float64))
}

func (ptr *QPrintPreviewWidget) ConnectPaintRequested(f func(printer *QPrinter)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPaintRequested", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintPreviewWidget) DisconnectPaintRequested() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPaintRequested"})
}

func (ptr *QPrintPreviewWidget) PaintRequested(printer QPrinter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintRequested", printer})
}

func (ptr *QPrintPreviewWidget) ConnectPreviewChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPreviewChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintPreviewWidget) DisconnectPreviewChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPreviewChanged"})
}

func (ptr *QPrintPreviewWidget) PreviewChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PreviewChanged"})
}

func (ptr *QPrintPreviewWidget) ConnectPrint(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPrint", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintPreviewWidget) DisconnectPrint() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPrint"})
}

func (ptr *QPrintPreviewWidget) Print() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Print"})
}

func (ptr *QPrintPreviewWidget) PrintDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrintDefault"})
}

func (ptr *QPrintPreviewWidget) ConnectSetAllPagesViewMode(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetAllPagesViewMode", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintPreviewWidget) DisconnectSetAllPagesViewMode() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetAllPagesViewMode"})
}

func (ptr *QPrintPreviewWidget) SetAllPagesViewMode() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAllPagesViewMode"})
}

func (ptr *QPrintPreviewWidget) SetAllPagesViewModeDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAllPagesViewModeDefault"})
}

func (ptr *QPrintPreviewWidget) ConnectSetCurrentPage(f func(page int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetCurrentPage", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintPreviewWidget) DisconnectSetCurrentPage() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetCurrentPage"})
}

func (ptr *QPrintPreviewWidget) SetCurrentPage(page int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCurrentPage", page})
}

func (ptr *QPrintPreviewWidget) SetCurrentPageDefault(page int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCurrentPageDefault", page})
}

func (ptr *QPrintPreviewWidget) ConnectSetFacingPagesViewMode(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetFacingPagesViewMode", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintPreviewWidget) DisconnectSetFacingPagesViewMode() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetFacingPagesViewMode"})
}

func (ptr *QPrintPreviewWidget) SetFacingPagesViewMode() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFacingPagesViewMode"})
}

func (ptr *QPrintPreviewWidget) SetFacingPagesViewModeDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFacingPagesViewModeDefault"})
}

func (ptr *QPrintPreviewWidget) ConnectSetLandscapeOrientation(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetLandscapeOrientation", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintPreviewWidget) DisconnectSetLandscapeOrientation() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetLandscapeOrientation"})
}

func (ptr *QPrintPreviewWidget) SetLandscapeOrientation() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLandscapeOrientation"})
}

func (ptr *QPrintPreviewWidget) SetLandscapeOrientationDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLandscapeOrientationDefault"})
}

func (ptr *QPrintPreviewWidget) ConnectSetOrientation(f func(orientation QPrinter__Orientation)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetOrientation", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintPreviewWidget) DisconnectSetOrientation() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetOrientation"})
}

func (ptr *QPrintPreviewWidget) SetOrientation(orientation QPrinter__Orientation) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOrientation", orientation})
}

func (ptr *QPrintPreviewWidget) SetOrientationDefault(orientation QPrinter__Orientation) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOrientationDefault", orientation})
}

func (ptr *QPrintPreviewWidget) ConnectSetPortraitOrientation(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetPortraitOrientation", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintPreviewWidget) DisconnectSetPortraitOrientation() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetPortraitOrientation"})
}

func (ptr *QPrintPreviewWidget) SetPortraitOrientation() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPortraitOrientation"})
}

func (ptr *QPrintPreviewWidget) SetPortraitOrientationDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPortraitOrientationDefault"})
}

func (ptr *QPrintPreviewWidget) ConnectSetSinglePageViewMode(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetSinglePageViewMode", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintPreviewWidget) DisconnectSetSinglePageViewMode() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetSinglePageViewMode"})
}

func (ptr *QPrintPreviewWidget) SetSinglePageViewMode() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSinglePageViewMode"})
}

func (ptr *QPrintPreviewWidget) SetSinglePageViewModeDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSinglePageViewModeDefault"})
}

func (ptr *QPrintPreviewWidget) ConnectSetViewMode(f func(mode QPrintPreviewWidget__ViewMode)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetViewMode", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintPreviewWidget) DisconnectSetViewMode() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetViewMode"})
}

func (ptr *QPrintPreviewWidget) SetViewMode(mode QPrintPreviewWidget__ViewMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetViewMode", mode})
}

func (ptr *QPrintPreviewWidget) SetViewModeDefault(mode QPrintPreviewWidget__ViewMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetViewModeDefault", mode})
}

func (ptr *QPrintPreviewWidget) ConnectSetVisible(f func(visible bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetVisible", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintPreviewWidget) DisconnectSetVisible() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetVisible"})
}

func (ptr *QPrintPreviewWidget) SetVisible(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVisible", visible})
}

func (ptr *QPrintPreviewWidget) SetVisibleDefault(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVisibleDefault", visible})
}

func (ptr *QPrintPreviewWidget) ConnectSetZoomFactor(f func(factor float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetZoomFactor", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintPreviewWidget) DisconnectSetZoomFactor() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetZoomFactor"})
}

func (ptr *QPrintPreviewWidget) SetZoomFactor(factor float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetZoomFactor", factor})
}

func (ptr *QPrintPreviewWidget) SetZoomFactorDefault(factor float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetZoomFactorDefault", factor})
}

func (ptr *QPrintPreviewWidget) ConnectSetZoomMode(f func(zoomMode QPrintPreviewWidget__ZoomMode)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetZoomMode", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintPreviewWidget) DisconnectSetZoomMode() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetZoomMode"})
}

func (ptr *QPrintPreviewWidget) SetZoomMode(zoomMode QPrintPreviewWidget__ZoomMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetZoomMode", zoomMode})
}

func (ptr *QPrintPreviewWidget) SetZoomModeDefault(zoomMode QPrintPreviewWidget__ZoomMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetZoomModeDefault", zoomMode})
}

func (ptr *QPrintPreviewWidget) ConnectUpdatePreview(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUpdatePreview", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintPreviewWidget) DisconnectUpdatePreview() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUpdatePreview"})
}

func (ptr *QPrintPreviewWidget) UpdatePreview() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdatePreview"})
}

func (ptr *QPrintPreviewWidget) UpdatePreviewDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdatePreviewDefault"})
}

func (ptr *QPrintPreviewWidget) ViewMode() QPrintPreviewWidget__ViewMode {

	return QPrintPreviewWidget__ViewMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ViewMode"}).(float64))
}

func (ptr *QPrintPreviewWidget) ZoomFactor() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ZoomFactor"}).(float64)
}

func (ptr *QPrintPreviewWidget) ConnectZoomIn(f func(factor float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectZoomIn", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintPreviewWidget) DisconnectZoomIn() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectZoomIn"})
}

func (ptr *QPrintPreviewWidget) ZoomIn(factor float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ZoomIn", factor})
}

func (ptr *QPrintPreviewWidget) ZoomInDefault(factor float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ZoomInDefault", factor})
}

func (ptr *QPrintPreviewWidget) ZoomMode() QPrintPreviewWidget__ZoomMode {

	return QPrintPreviewWidget__ZoomMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ZoomMode"}).(float64))
}

func (ptr *QPrintPreviewWidget) ConnectZoomOut(f func(factor float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectZoomOut", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintPreviewWidget) DisconnectZoomOut() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectZoomOut"})
}

func (ptr *QPrintPreviewWidget) ZoomOut(factor float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ZoomOut", factor})
}

func (ptr *QPrintPreviewWidget) ZoomOutDefault(factor float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ZoomOutDefault", factor})
}

func (ptr *QPrintPreviewWidget) ConnectDestroyQPrintPreviewWidget(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQPrintPreviewWidget", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrintPreviewWidget) DisconnectDestroyQPrintPreviewWidget() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQPrintPreviewWidget"})
}

func (ptr *QPrintPreviewWidget) DestroyQPrintPreviewWidget() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQPrintPreviewWidget"})
}

func (ptr *QPrintPreviewWidget) DestroyQPrintPreviewWidgetDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQPrintPreviewWidgetDefault"})
}

func (ptr *QPrintPreviewWidget) __actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_atList", i}).(*widgets.QAction)
}

func (ptr *QPrintPreviewWidget) __actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_setList", i})
}

func (ptr *QPrintPreviewWidget) __actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_newList"}).(unsafe.Pointer)
}

func (ptr *QPrintPreviewWidget) __addActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QPrintPreviewWidget) __addActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_setList", i})
}

func (ptr *QPrintPreviewWidget) __addActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QPrintPreviewWidget) __insertActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QPrintPreviewWidget) __insertActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_setList", i})
}

func (ptr *QPrintPreviewWidget) __insertActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QPrintPreviewWidget) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QPrintPreviewWidget) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QPrintPreviewWidget) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QPrintPreviewWidget) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QPrintPreviewWidget) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QPrintPreviewWidget) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QPrintPreviewWidget) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QPrintPreviewWidget) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QPrintPreviewWidget) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QPrintPreviewWidget) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QPrintPreviewWidget) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QPrintPreviewWidget) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QPrintPreviewWidget) ActionEventDefault(event gui.QActionEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActionEventDefault", event})
}

func (ptr *QPrintPreviewWidget) ChangeEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChangeEventDefault", event})
}

func (ptr *QPrintPreviewWidget) CloseDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"}).(bool)
}

func (ptr *QPrintPreviewWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseEventDefault", event})
}

func (ptr *QPrintPreviewWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContextMenuEventDefault", event})
}

func (ptr *QPrintPreviewWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragEnterEventDefault", event})
}

func (ptr *QPrintPreviewWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragLeaveEventDefault", event})
}

func (ptr *QPrintPreviewWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragMoveEventDefault", event})
}

func (ptr *QPrintPreviewWidget) DropEventDefault(event gui.QDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DropEventDefault", event})
}

func (ptr *QPrintPreviewWidget) EnterEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EnterEventDefault", event})
}

func (ptr *QPrintPreviewWidget) EventDefault(event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", event}).(bool)
}

func (ptr *QPrintPreviewWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusInEventDefault", event})
}

func (ptr *QPrintPreviewWidget) FocusNextPrevChildDefault(next bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusNextPrevChildDefault", next}).(bool)
}

func (ptr *QPrintPreviewWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusOutEventDefault", event})
}

func (ptr *QPrintPreviewWidget) HasHeightForWidthDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasHeightForWidthDefault"}).(bool)
}

func (ptr *QPrintPreviewWidget) HeightForWidthDefault(w int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HeightForWidthDefault", w}).(float64))
}

func (ptr *QPrintPreviewWidget) HideDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideDefault"})
}

func (ptr *QPrintPreviewWidget) HideEventDefault(event gui.QHideEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideEventDefault", event})
}

func (ptr *QPrintPreviewWidget) InitPainterDefault(painter gui.QPainter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InitPainterDefault", painter})
}

func (ptr *QPrintPreviewWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodEventDefault", event})
}

func (ptr *QPrintPreviewWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodQueryDefault", query}).(*core.QVariant)
}

func (ptr *QPrintPreviewWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyPressEventDefault", event})
}

func (ptr *QPrintPreviewWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyReleaseEventDefault", event})
}

func (ptr *QPrintPreviewWidget) LeaveEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LeaveEventDefault", event})
}

func (ptr *QPrintPreviewWidget) LowerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LowerDefault"})
}

func (ptr *QPrintPreviewWidget) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetricDefault", m}).(float64))
}

func (ptr *QPrintPreviewWidget) MinimumSizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinimumSizeHintDefault"}).(*core.QSize)
}

func (ptr *QPrintPreviewWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseDoubleClickEventDefault", event})
}

func (ptr *QPrintPreviewWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseMoveEventDefault", event})
}

func (ptr *QPrintPreviewWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MousePressEventDefault", event})
}

func (ptr *QPrintPreviewWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseReleaseEventDefault", event})
}

func (ptr *QPrintPreviewWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveEventDefault", event})
}

func (ptr *QPrintPreviewWidget) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NativeEventDefault", eventType, message, result}).(bool)
}

func (ptr *QPrintPreviewWidget) PaintEngineDefault() *gui.QPaintEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEngineDefault"}).(*gui.QPaintEngine)
}

func (ptr *QPrintPreviewWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEventDefault", event})
}

func (ptr *QPrintPreviewWidget) RaiseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RaiseDefault"})
}

func (ptr *QPrintPreviewWidget) RepaintDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RepaintDefault"})
}

func (ptr *QPrintPreviewWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResizeEventDefault", event})
}

func (ptr *QPrintPreviewWidget) SetDisabledDefault(disable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDisabledDefault", disable})
}

func (ptr *QPrintPreviewWidget) SetEnabledDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetEnabledDefault", vbo})
}

func (ptr *QPrintPreviewWidget) SetFocus2Default() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFocus2Default"})
}

func (ptr *QPrintPreviewWidget) SetHiddenDefault(hidden bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHiddenDefault", hidden})
}

func (ptr *QPrintPreviewWidget) SetStyleSheetDefault(styleSheet string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetStyleSheetDefault", styleSheet})
}

func (ptr *QPrintPreviewWidget) SetWindowModifiedDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowModifiedDefault", vbo})
}

func (ptr *QPrintPreviewWidget) SetWindowTitleDefault(vqs string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowTitleDefault", vqs})
}

func (ptr *QPrintPreviewWidget) ShowDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowDefault"})
}

func (ptr *QPrintPreviewWidget) ShowEventDefault(event gui.QShowEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowEventDefault", event})
}

func (ptr *QPrintPreviewWidget) ShowFullScreenDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowFullScreenDefault"})
}

func (ptr *QPrintPreviewWidget) ShowMaximizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMaximizedDefault"})
}

func (ptr *QPrintPreviewWidget) ShowMinimizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMinimizedDefault"})
}

func (ptr *QPrintPreviewWidget) ShowNormalDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowNormalDefault"})
}

func (ptr *QPrintPreviewWidget) SizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeHintDefault"}).(*core.QSize)
}

func (ptr *QPrintPreviewWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TabletEventDefault", event})
}

func (ptr *QPrintPreviewWidget) UpdateDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateDefault"})
}

func (ptr *QPrintPreviewWidget) UpdateMicroFocusDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateMicroFocusDefault"})
}

func (ptr *QPrintPreviewWidget) WheelEventDefault(event gui.QWheelEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WheelEventDefault", event})
}

func (ptr *QPrintPreviewWidget) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QPrintPreviewWidget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QPrintPreviewWidget) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QPrintPreviewWidget) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QPrintPreviewWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QPrintPreviewWidget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QPrintPreviewWidget) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QPrintPreviewWidget) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QPrinter struct {
	gui.QPagedPaintDevice
}

type QPrinter_ITF interface {
	gui.QPagedPaintDevice_ITF
	QPrinter_PTR() *QPrinter
}

func (ptr *QPrinter) QPrinter_PTR() *QPrinter {
	return ptr
}

func (ptr *QPrinter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QPagedPaintDevice_PTR().Pointer()
	}
	return nil
}

func (ptr *QPrinter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QPagedPaintDevice_PTR().SetPointer(p)
	}
}

func PointerFromQPrinter(ptr QPrinter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPrinter_PTR().Pointer()
	}
	return nil
}

func (n *QPrinter) InitFromInternal(ptr uintptr, name string) {
	n.QPagedPaintDevice_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QPrinter) ClassNameInternalF() string {
	return n.QPagedPaintDevice_PTR().ClassNameInternalF()
}

func NewQPrinterFromPointer(ptr unsafe.Pointer) (n *QPrinter) {
	n = new(QPrinter)
	n.InitFromInternal(uintptr(ptr), "printsupport.QPrinter")
	return
}

//go:generate stringer -type=QPrinter__PrinterMode
//QPrinter::PrinterMode
type QPrinter__PrinterMode int64

const (
	QPrinter__ScreenResolution  QPrinter__PrinterMode = QPrinter__PrinterMode(0)
	QPrinter__PrinterResolution QPrinter__PrinterMode = QPrinter__PrinterMode(1)
	QPrinter__HighResolution    QPrinter__PrinterMode = QPrinter__PrinterMode(2)
)

//go:generate stringer -type=QPrinter__Orientation
//QPrinter::Orientation
type QPrinter__Orientation int64

const (
	QPrinter__Portrait  QPrinter__Orientation = QPrinter__Orientation(0)
	QPrinter__Landscape QPrinter__Orientation = QPrinter__Orientation(1)
)

//go:generate stringer -type=QPrinter__PageOrder
//QPrinter::PageOrder
type QPrinter__PageOrder int64

const (
	QPrinter__FirstPageFirst QPrinter__PageOrder = QPrinter__PageOrder(0)
	QPrinter__LastPageFirst  QPrinter__PageOrder = QPrinter__PageOrder(1)
)

//go:generate stringer -type=QPrinter__ColorMode
//QPrinter::ColorMode
type QPrinter__ColorMode int64

const (
	QPrinter__GrayScale QPrinter__ColorMode = QPrinter__ColorMode(0)
	QPrinter__Color     QPrinter__ColorMode = QPrinter__ColorMode(1)
)

//go:generate stringer -type=QPrinter__PaperSource
//QPrinter::PaperSource
type QPrinter__PaperSource int64

const (
	QPrinter__OnlyOne         QPrinter__PaperSource = QPrinter__PaperSource(0)
	QPrinter__Lower           QPrinter__PaperSource = QPrinter__PaperSource(1)
	QPrinter__Middle          QPrinter__PaperSource = QPrinter__PaperSource(2)
	QPrinter__Manual          QPrinter__PaperSource = QPrinter__PaperSource(3)
	QPrinter__Envelope        QPrinter__PaperSource = QPrinter__PaperSource(4)
	QPrinter__EnvelopeManual  QPrinter__PaperSource = QPrinter__PaperSource(5)
	QPrinter__Auto            QPrinter__PaperSource = QPrinter__PaperSource(6)
	QPrinter__Tractor         QPrinter__PaperSource = QPrinter__PaperSource(7)
	QPrinter__SmallFormat     QPrinter__PaperSource = QPrinter__PaperSource(8)
	QPrinter__LargeFormat     QPrinter__PaperSource = QPrinter__PaperSource(9)
	QPrinter__LargeCapacity   QPrinter__PaperSource = QPrinter__PaperSource(10)
	QPrinter__Cassette        QPrinter__PaperSource = QPrinter__PaperSource(11)
	QPrinter__FormSource      QPrinter__PaperSource = QPrinter__PaperSource(12)
	QPrinter__MaxPageSource   QPrinter__PaperSource = QPrinter__PaperSource(13)
	QPrinter__CustomSource    QPrinter__PaperSource = QPrinter__PaperSource(14)
	QPrinter__LastPaperSource QPrinter__PaperSource = QPrinter__PaperSource(QPrinter__CustomSource)
	QPrinter__Upper           QPrinter__PaperSource = QPrinter__PaperSource(QPrinter__OnlyOne)
)

//go:generate stringer -type=QPrinter__PrinterState
//QPrinter::PrinterState
type QPrinter__PrinterState int64

const (
	QPrinter__Idle    QPrinter__PrinterState = QPrinter__PrinterState(0)
	QPrinter__Active  QPrinter__PrinterState = QPrinter__PrinterState(1)
	QPrinter__Aborted QPrinter__PrinterState = QPrinter__PrinterState(2)
	QPrinter__Error   QPrinter__PrinterState = QPrinter__PrinterState(3)
)

//go:generate stringer -type=QPrinter__OutputFormat
//QPrinter::OutputFormat
type QPrinter__OutputFormat int64

const (
	QPrinter__NativeFormat QPrinter__OutputFormat = QPrinter__OutputFormat(0)
	QPrinter__PdfFormat    QPrinter__OutputFormat = QPrinter__OutputFormat(1)
)

//go:generate stringer -type=QPrinter__PrintRange
//QPrinter::PrintRange
type QPrinter__PrintRange int64

const (
	QPrinter__AllPages    QPrinter__PrintRange = QPrinter__PrintRange(0)
	QPrinter__Selection   QPrinter__PrintRange = QPrinter__PrintRange(1)
	QPrinter__PageRange   QPrinter__PrintRange = QPrinter__PrintRange(2)
	QPrinter__CurrentPage QPrinter__PrintRange = QPrinter__PrintRange(3)
)

//go:generate stringer -type=QPrinter__Unit
//QPrinter::Unit
type QPrinter__Unit int64

const (
	QPrinter__Millimeter  QPrinter__Unit = QPrinter__Unit(0)
	QPrinter__Point       QPrinter__Unit = QPrinter__Unit(1)
	QPrinter__Inch        QPrinter__Unit = QPrinter__Unit(2)
	QPrinter__Pica        QPrinter__Unit = QPrinter__Unit(3)
	QPrinter__Didot       QPrinter__Unit = QPrinter__Unit(4)
	QPrinter__Cicero      QPrinter__Unit = QPrinter__Unit(5)
	QPrinter__DevicePixel QPrinter__Unit = QPrinter__Unit(6)
)

//go:generate stringer -type=QPrinter__DuplexMode
//QPrinter::DuplexMode
type QPrinter__DuplexMode int64

const (
	QPrinter__DuplexNone      QPrinter__DuplexMode = QPrinter__DuplexMode(0)
	QPrinter__DuplexAuto      QPrinter__DuplexMode = QPrinter__DuplexMode(1)
	QPrinter__DuplexLongSide  QPrinter__DuplexMode = QPrinter__DuplexMode(2)
	QPrinter__DuplexShortSide QPrinter__DuplexMode = QPrinter__DuplexMode(3)
)

func NewQPrinter(mode QPrinter__PrinterMode) *QPrinter {

	return internal.CallLocalFunction([]interface{}{"", "", "printsupport.NewQPrinter", "", mode}).(*QPrinter)
}

func NewQPrinter2(printer QPrinterInfo_ITF, mode QPrinter__PrinterMode) *QPrinter {

	return internal.CallLocalFunction([]interface{}{"", "", "printsupport.NewQPrinter2", "", printer, mode}).(*QPrinter)
}

func (ptr *QPrinter) Abort() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Abort"}).(bool)
}

func (ptr *QPrinter) CollateCopies() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CollateCopies"}).(bool)
}

func (ptr *QPrinter) ColorMode() QPrinter__ColorMode {

	return QPrinter__ColorMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColorMode"}).(float64))
}

func (ptr *QPrinter) CopyCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CopyCount"}).(float64))
}

func (ptr *QPrinter) Creator() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Creator"}).(string)
}

func (ptr *QPrinter) DocName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DocName"}).(string)
}

func (ptr *QPrinter) Duplex() QPrinter__DuplexMode {

	return QPrinter__DuplexMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Duplex"}).(float64))
}

func (ptr *QPrinter) FontEmbeddingEnabled() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FontEmbeddingEnabled"}).(bool)
}

func (ptr *QPrinter) FromPage() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FromPage"}).(float64))
}

func (ptr *QPrinter) FullPage() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FullPage"}).(bool)
}

func (ptr *QPrinter) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QPrinter) ConnectNewPage(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNewPage", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrinter) DisconnectNewPage() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNewPage"})
}

func (ptr *QPrinter) NewPage() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NewPage"}).(bool)
}

func (ptr *QPrinter) NewPageDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NewPageDefault"}).(bool)
}

func (ptr *QPrinter) OutputFileName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OutputFileName"}).(string)
}

func (ptr *QPrinter) OutputFormat() QPrinter__OutputFormat {

	return QPrinter__OutputFormat(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OutputFormat"}).(float64))
}

func (ptr *QPrinter) PageOrder() QPrinter__PageOrder {

	return QPrinter__PageOrder(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PageOrder"}).(float64))
}

func (ptr *QPrinter) PageRect2(unit QPrinter__Unit) *core.QRectF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PageRect2", unit}).(*core.QRectF)
}

func (ptr *QPrinter) ConnectPaintEngine(f func() *gui.QPaintEngine) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPaintEngine", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrinter) DisconnectPaintEngine() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPaintEngine"})
}

func (ptr *QPrinter) PaintEngine() *gui.QPaintEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEngine"}).(*gui.QPaintEngine)
}

func (ptr *QPrinter) PaintEngineDefault() *gui.QPaintEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEngineDefault"}).(*gui.QPaintEngine)
}

func (ptr *QPrinter) PaperRect2(unit QPrinter__Unit) *core.QRectF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaperRect2", unit}).(*core.QRectF)
}

func (ptr *QPrinter) PaperSource() QPrinter__PaperSource {

	return QPrinter__PaperSource(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaperSource"}).(float64))
}

func (ptr *QPrinter) PdfVersion() gui.QPagedPaintDevice__PdfVersion {

	return gui.QPagedPaintDevice__PdfVersion(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PdfVersion"}).(float64))
}

func (ptr *QPrinter) PrintEngine() *QPrintEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrintEngine"}).(*QPrintEngine)
}

func (ptr *QPrinter) PrintProgram() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrintProgram"}).(string)
}

func (ptr *QPrinter) PrintRange() QPrinter__PrintRange {

	return QPrinter__PrintRange(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrintRange"}).(float64))
}

func (ptr *QPrinter) PrinterName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrinterName"}).(string)
}

func (ptr *QPrinter) PrinterSelectionOption() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrinterSelectionOption"}).(string)
}

func (ptr *QPrinter) PrinterState() QPrinter__PrinterState {

	return QPrinter__PrinterState(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrinterState"}).(float64))
}

func (ptr *QPrinter) Resolution() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Resolution"}).(float64))
}

func (ptr *QPrinter) SetCollateCopies(collate bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCollateCopies", collate})
}

func (ptr *QPrinter) SetColorMode(newColorMode QPrinter__ColorMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetColorMode", newColorMode})
}

func (ptr *QPrinter) SetCopyCount(count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCopyCount", count})
}

func (ptr *QPrinter) SetCreator(creator string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCreator", creator})
}

func (ptr *QPrinter) SetDocName(name string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDocName", name})
}

func (ptr *QPrinter) SetDuplex(duplex QPrinter__DuplexMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDuplex", duplex})
}

func (ptr *QPrinter) SetEngines(printEngine QPrintEngine_ITF, paintEngine gui.QPaintEngine_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetEngines", printEngine, paintEngine})
}

func (ptr *QPrinter) SetFontEmbeddingEnabled(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFontEmbeddingEnabled", enable})
}

func (ptr *QPrinter) SetFromTo(from int, to int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFromTo", from, to})
}

func (ptr *QPrinter) SetFullPage(fp bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFullPage", fp})
}

func (ptr *QPrinter) SetOutputFileName(fileName string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOutputFileName", fileName})
}

func (ptr *QPrinter) SetOutputFormat(format QPrinter__OutputFormat) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOutputFormat", format})
}

func (ptr *QPrinter) SetPageOrder(pageOrder QPrinter__PageOrder) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPageOrder", pageOrder})
}

func (ptr *QPrinter) SetPaperSource(source QPrinter__PaperSource) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPaperSource", source})
}

func (ptr *QPrinter) SetPdfVersion(version gui.QPagedPaintDevice__PdfVersion) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPdfVersion", version})
}

func (ptr *QPrinter) SetPrintProgram(printProg string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPrintProgram", printProg})
}

func (ptr *QPrinter) SetPrintRange(ran QPrinter__PrintRange) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPrintRange", ran})
}

func (ptr *QPrinter) SetPrinterName(name string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPrinterName", name})
}

func (ptr *QPrinter) SetPrinterSelectionOption(option string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPrinterSelectionOption", option})
}

func (ptr *QPrinter) SetResolution(dpi int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetResolution", dpi})
}

func (ptr *QPrinter) SupportedResolutions() []int {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedResolutions"}).([]int)
}

func (ptr *QPrinter) SupportsMultipleCopies() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportsMultipleCopies"}).(bool)
}

func (ptr *QPrinter) ToPage() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToPage"}).(float64))
}

func (ptr *QPrinter) ConnectDestroyQPrinter(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQPrinter", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPrinter) DisconnectDestroyQPrinter() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQPrinter"})
}

func (ptr *QPrinter) DestroyQPrinter() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQPrinter"})
}

func (ptr *QPrinter) DestroyQPrinterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQPrinterDefault"})
}

func (ptr *QPrinter) __supportedResolutions_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__supportedResolutions_atList", i}).(float64))
}

func (ptr *QPrinter) __supportedResolutions_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__supportedResolutions_setList", i})
}

func (ptr *QPrinter) __supportedResolutions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__supportedResolutions_newList"}).(unsafe.Pointer)
}

func (ptr *QPrinter) SetPageSize2Default(size gui.QPagedPaintDevice__PageSize) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPageSize2Default", size})
}

func (ptr *QPrinter) SetPageSizeMMDefault(size core.QSizeF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPageSizeMMDefault", size})
}

func (ptr *QPrinter) MetricDefault(metric gui.QPaintDevice__PaintDeviceMetric) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetricDefault", metric}).(float64))
}

type QPrinterInfo struct {
	internal.Internal
}

type QPrinterInfo_ITF interface {
	QPrinterInfo_PTR() *QPrinterInfo
}

func (ptr *QPrinterInfo) QPrinterInfo_PTR() *QPrinterInfo {
	return ptr
}

func (ptr *QPrinterInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QPrinterInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQPrinterInfo(ptr QPrinterInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPrinterInfo_PTR().Pointer()
	}
	return nil
}

func (n *QPrinterInfo) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQPrinterInfoFromPointer(ptr unsafe.Pointer) (n *QPrinterInfo) {
	n = new(QPrinterInfo)
	n.InitFromInternal(uintptr(ptr), "printsupport.QPrinterInfo")
	return
}
func NewQPrinterInfo() *QPrinterInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "printsupport.NewQPrinterInfo", ""}).(*QPrinterInfo)
}

func NewQPrinterInfo2(other QPrinterInfo_ITF) *QPrinterInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "printsupport.NewQPrinterInfo2", "", other}).(*QPrinterInfo)
}

func NewQPrinterInfo3(printer QPrinter_ITF) *QPrinterInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "printsupport.NewQPrinterInfo3", "", printer}).(*QPrinterInfo)
}

func QPrinterInfo_AvailablePrinterNames() []string {

	return internal.CallLocalFunction([]interface{}{"", "", "printsupport.QPrinterInfo_AvailablePrinterNames", ""}).([]string)
}

func (ptr *QPrinterInfo) AvailablePrinterNames() []string {

	return internal.CallLocalFunction([]interface{}{"", "", "printsupport.QPrinterInfo_AvailablePrinterNames", ""}).([]string)
}

func QPrinterInfo_AvailablePrinters() []*QPrinterInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "printsupport.QPrinterInfo_AvailablePrinters", ""}).([]*QPrinterInfo)
}

func (ptr *QPrinterInfo) AvailablePrinters() []*QPrinterInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "printsupport.QPrinterInfo_AvailablePrinters", ""}).([]*QPrinterInfo)
}

func (ptr *QPrinterInfo) DefaultColorMode() QPrinter__ColorMode {

	return QPrinter__ColorMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DefaultColorMode"}).(float64))
}

func (ptr *QPrinterInfo) DefaultDuplexMode() QPrinter__DuplexMode {

	return QPrinter__DuplexMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DefaultDuplexMode"}).(float64))
}

func (ptr *QPrinterInfo) DefaultPageSize() *gui.QPageSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DefaultPageSize"}).(*gui.QPageSize)
}

func QPrinterInfo_DefaultPrinter() *QPrinterInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "printsupport.QPrinterInfo_DefaultPrinter", ""}).(*QPrinterInfo)
}

func (ptr *QPrinterInfo) DefaultPrinter() *QPrinterInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "printsupport.QPrinterInfo_DefaultPrinter", ""}).(*QPrinterInfo)
}

func QPrinterInfo_DefaultPrinterName() string {

	return internal.CallLocalFunction([]interface{}{"", "", "printsupport.QPrinterInfo_DefaultPrinterName", ""}).(string)
}

func (ptr *QPrinterInfo) DefaultPrinterName() string {

	return internal.CallLocalFunction([]interface{}{"", "", "printsupport.QPrinterInfo_DefaultPrinterName", ""}).(string)
}

func (ptr *QPrinterInfo) Description() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Description"}).(string)
}

func (ptr *QPrinterInfo) IsDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsDefault"}).(bool)
}

func (ptr *QPrinterInfo) IsNull() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNull"}).(bool)
}

func (ptr *QPrinterInfo) IsRemote() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsRemote"}).(bool)
}

func (ptr *QPrinterInfo) Location() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Location"}).(string)
}

func (ptr *QPrinterInfo) MakeAndModel() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MakeAndModel"}).(string)
}

func (ptr *QPrinterInfo) MaximumPhysicalPageSize() *gui.QPageSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaximumPhysicalPageSize"}).(*gui.QPageSize)
}

func (ptr *QPrinterInfo) MinimumPhysicalPageSize() *gui.QPageSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinimumPhysicalPageSize"}).(*gui.QPageSize)
}

func QPrinterInfo_PrinterInfo(printerName string) *QPrinterInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "printsupport.QPrinterInfo_PrinterInfo", "", printerName}).(*QPrinterInfo)
}

func (ptr *QPrinterInfo) PrinterInfo(printerName string) *QPrinterInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "printsupport.QPrinterInfo_PrinterInfo", "", printerName}).(*QPrinterInfo)
}

func (ptr *QPrinterInfo) PrinterName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrinterName"}).(string)
}

func (ptr *QPrinterInfo) State() QPrinter__PrinterState {

	return QPrinter__PrinterState(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "State"}).(float64))
}

func (ptr *QPrinterInfo) SupportedColorModes() []QPrinter__ColorMode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedColorModes"}).([]QPrinter__ColorMode)
}

func (ptr *QPrinterInfo) SupportedDuplexModes() []QPrinter__DuplexMode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedDuplexModes"}).([]QPrinter__DuplexMode)
}

func (ptr *QPrinterInfo) SupportedPageSizes() []*gui.QPageSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedPageSizes"}).([]*gui.QPageSize)
}

func (ptr *QPrinterInfo) SupportedResolutions() []int {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedResolutions"}).([]int)
}

func (ptr *QPrinterInfo) SupportsCustomPageSizes() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportsCustomPageSizes"}).(bool)
}

func (ptr *QPrinterInfo) DestroyQPrinterInfo() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQPrinterInfo"})
}

func (ptr *QPrinterInfo) __availablePrinters_atList(i int) *QPrinterInfo {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__availablePrinters_atList", i}).(*QPrinterInfo)
}

func (ptr *QPrinterInfo) __availablePrinters_setList(i QPrinterInfo_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__availablePrinters_setList", i})
}

func (ptr *QPrinterInfo) __availablePrinters_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__availablePrinters_newList"}).(unsafe.Pointer)
}

func (ptr *QPrinterInfo) __supportedColorModes_atList(i int) QPrinter__ColorMode {

	return QPrinter__ColorMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__supportedColorModes_atList", i}).(float64))
}

func (ptr *QPrinterInfo) __supportedColorModes_setList(i QPrinter__ColorMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__supportedColorModes_setList", i})
}

func (ptr *QPrinterInfo) __supportedColorModes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__supportedColorModes_newList"}).(unsafe.Pointer)
}

func (ptr *QPrinterInfo) __supportedDuplexModes_atList(i int) QPrinter__DuplexMode {

	return QPrinter__DuplexMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__supportedDuplexModes_atList", i}).(float64))
}

func (ptr *QPrinterInfo) __supportedDuplexModes_setList(i QPrinter__DuplexMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__supportedDuplexModes_setList", i})
}

func (ptr *QPrinterInfo) __supportedDuplexModes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__supportedDuplexModes_newList"}).(unsafe.Pointer)
}

func (ptr *QPrinterInfo) __supportedPageSizes_atList(i int) *gui.QPageSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__supportedPageSizes_atList", i}).(*gui.QPageSize)
}

func (ptr *QPrinterInfo) __supportedPageSizes_setList(i gui.QPageSize_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__supportedPageSizes_setList", i})
}

func (ptr *QPrinterInfo) __supportedPageSizes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__supportedPageSizes_newList"}).(unsafe.Pointer)
}

func (ptr *QPrinterInfo) __supportedPaperSizes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__supportedPaperSizes_newList"}).(unsafe.Pointer)
}

func (ptr *QPrinterInfo) __supportedResolutions_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__supportedResolutions_atList", i}).(float64))
}

func (ptr *QPrinterInfo) __supportedResolutions_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__supportedResolutions_setList", i})
}

func (ptr *QPrinterInfo) __supportedResolutions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__supportedResolutions_newList"}).(unsafe.Pointer)
}

func init() {
	internal.ConstructorTable["printsupport.QAbstractPrintDialog"] = NewQAbstractPrintDialogFromPointer
	internal.ConstructorTable["printsupport.QPageSetupDialog"] = NewQPageSetupDialogFromPointer
	internal.ConstructorTable["printsupport.QPrintDialog"] = NewQPrintDialogFromPointer
	internal.ConstructorTable["printsupport.QPrintEngine"] = NewQPrintEngineFromPointer
	internal.ConstructorTable["printsupport.QPrintPreviewDialog"] = NewQPrintPreviewDialogFromPointer
	internal.ConstructorTable["printsupport.QPrintPreviewWidget"] = NewQPrintPreviewWidgetFromPointer
	internal.ConstructorTable["printsupport.QPrinter"] = NewQPrinterFromPointer
	internal.ConstructorTable["printsupport.QPrinterInfo"] = NewQPrinterInfoFromPointer
}
