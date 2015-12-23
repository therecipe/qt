package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QDialog struct {
	QWidget
}

type QDialog_ITF interface {
	QWidget_ITF
	QDialog_PTR() *QDialog
}

func PointerFromQDialog(ptr QDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDialog_PTR().Pointer()
	}
	return nil
}

func NewQDialogFromPointer(ptr unsafe.Pointer) *QDialog {
	var n = new(QDialog)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDialog_") {
		n.SetObjectName("QDialog_" + qt.Identifier())
	}
	return n
}

func (ptr *QDialog) QDialog_PTR() *QDialog {
	return ptr
}

//QDialog::DialogCode
type QDialog__DialogCode int64

const (
	QDialog__Rejected = QDialog__DialogCode(0)
	QDialog__Accepted = QDialog__DialogCode(1)
)

func (ptr *QDialog) IsSizeGripEnabled() bool {
	defer qt.Recovering("QDialog::isSizeGripEnabled")

	if ptr.Pointer() != nil {
		return C.QDialog_IsSizeGripEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDialog) SetModal(modal bool) {
	defer qt.Recovering("QDialog::setModal")

	if ptr.Pointer() != nil {
		C.QDialog_SetModal(ptr.Pointer(), C.int(qt.GoBoolToInt(modal)))
	}
}

func (ptr *QDialog) SetResult(i int) {
	defer qt.Recovering("QDialog::setResult")

	if ptr.Pointer() != nil {
		C.QDialog_SetResult(ptr.Pointer(), C.int(i))
	}
}

func (ptr *QDialog) SetSizeGripEnabled(v bool) {
	defer qt.Recovering("QDialog::setSizeGripEnabled")

	if ptr.Pointer() != nil {
		C.QDialog_SetSizeGripEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func NewQDialog(parent QWidget_ITF, f core.Qt__WindowType) *QDialog {
	defer qt.Recovering("QDialog::QDialog")

	return NewQDialogFromPointer(C.QDialog_NewQDialog(PointerFromQWidget(parent), C.int(f)))
}

func (ptr *QDialog) ConnectAccept(f func()) {
	defer qt.Recovering("connect QDialog::accept")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "accept", f)
	}
}

func (ptr *QDialog) DisconnectAccept() {
	defer qt.Recovering("disconnect QDialog::accept")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "accept")
	}
}

//export callbackQDialogAccept
func callbackQDialogAccept(ptrName *C.char) bool {
	defer qt.Recovering("callback QDialog::accept")

	if signal := qt.GetSignal(C.GoString(ptrName), "accept"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QDialog) ConnectAccepted(f func()) {
	defer qt.Recovering("connect QDialog::accepted")

	if ptr.Pointer() != nil {
		C.QDialog_ConnectAccepted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "accepted", f)
	}
}

func (ptr *QDialog) DisconnectAccepted() {
	defer qt.Recovering("disconnect QDialog::accepted")

	if ptr.Pointer() != nil {
		C.QDialog_DisconnectAccepted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "accepted")
	}
}

//export callbackQDialogAccepted
func callbackQDialogAccepted(ptrName *C.char) {
	defer qt.Recovering("callback QDialog::accepted")

	if signal := qt.GetSignal(C.GoString(ptrName), "accepted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QDialog) ConnectCloseEvent(f func(e *gui.QCloseEvent)) {
	defer qt.Recovering("connect QDialog::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QDialog) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QDialog::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQDialogCloseEvent
func callbackQDialogCloseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QDialog::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QDialog) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QDialog::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QDialog) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QDialog::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQDialogContextMenuEvent
func callbackQDialogContextMenuEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QDialog::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QDialog) ConnectDone(f func(r int)) {
	defer qt.Recovering("connect QDialog::done")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "done", f)
	}
}

func (ptr *QDialog) DisconnectDone() {
	defer qt.Recovering("disconnect QDialog::done")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "done")
	}
}

//export callbackQDialogDone
func callbackQDialogDone(ptrName *C.char, r C.int) bool {
	defer qt.Recovering("callback QDialog::done")

	if signal := qt.GetSignal(C.GoString(ptrName), "done"); signal != nil {
		signal.(func(int))(int(r))
		return true
	}
	return false

}

func (ptr *QDialog) Exec() int {
	defer qt.Recovering("QDialog::exec")

	if ptr.Pointer() != nil {
		return int(C.QDialog_Exec(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDialog) ConnectFinished(f func(result int)) {
	defer qt.Recovering("connect QDialog::finished")

	if ptr.Pointer() != nil {
		C.QDialog_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QDialog) DisconnectFinished() {
	defer qt.Recovering("disconnect QDialog::finished")

	if ptr.Pointer() != nil {
		C.QDialog_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQDialogFinished
func callbackQDialogFinished(ptrName *C.char, result C.int) {
	defer qt.Recovering("callback QDialog::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func(int))(int(result))
	}

}

func (ptr *QDialog) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QDialog::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QDialog) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QDialog::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQDialogKeyPressEvent
func callbackQDialogKeyPressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QDialog::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QDialog) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QDialog::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QDialog_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDialog) ConnectOpen(f func()) {
	defer qt.Recovering("connect QDialog::open")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "open", f)
	}
}

func (ptr *QDialog) DisconnectOpen() {
	defer qt.Recovering("disconnect QDialog::open")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "open")
	}
}

//export callbackQDialogOpen
func callbackQDialogOpen(ptrName *C.char) bool {
	defer qt.Recovering("callback QDialog::open")

	if signal := qt.GetSignal(C.GoString(ptrName), "open"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QDialog) ConnectReject(f func()) {
	defer qt.Recovering("connect QDialog::reject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reject", f)
	}
}

func (ptr *QDialog) DisconnectReject() {
	defer qt.Recovering("disconnect QDialog::reject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reject")
	}
}

//export callbackQDialogReject
func callbackQDialogReject(ptrName *C.char) bool {
	defer qt.Recovering("callback QDialog::reject")

	if signal := qt.GetSignal(C.GoString(ptrName), "reject"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QDialog) ConnectRejected(f func()) {
	defer qt.Recovering("connect QDialog::rejected")

	if ptr.Pointer() != nil {
		C.QDialog_ConnectRejected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "rejected", f)
	}
}

func (ptr *QDialog) DisconnectRejected() {
	defer qt.Recovering("disconnect QDialog::rejected")

	if ptr.Pointer() != nil {
		C.QDialog_DisconnectRejected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "rejected")
	}
}

//export callbackQDialogRejected
func callbackQDialogRejected(ptrName *C.char) {
	defer qt.Recovering("callback QDialog::rejected")

	if signal := qt.GetSignal(C.GoString(ptrName), "rejected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QDialog) ConnectResizeEvent(f func(v *gui.QResizeEvent)) {
	defer qt.Recovering("connect QDialog::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QDialog) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QDialog::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQDialogResizeEvent
func callbackQDialogResizeEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QDialog::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QDialog) Result() int {
	defer qt.Recovering("QDialog::result")

	if ptr.Pointer() != nil {
		return int(C.QDialog_Result(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDialog) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QDialog::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QDialog) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QDialog::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQDialogSetVisible
func callbackQDialogSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QDialog::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QDialog) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QDialog::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QDialog) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QDialog::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQDialogShowEvent
func callbackQDialogShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDialog::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDialog) SizeHint() *core.QSize {
	defer qt.Recovering("QDialog::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QDialog_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDialog) DestroyQDialog() {
	defer qt.Recovering("QDialog::~QDialog")

	if ptr.Pointer() != nil {
		C.QDialog_DestroyQDialog(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDialog) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QDialog::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QDialog) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QDialog::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQDialogActionEvent
func callbackQDialogActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDialog::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDialog) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QDialog::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QDialog) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QDialog::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQDialogDragEnterEvent
func callbackQDialogDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDialog::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDialog) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QDialog) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQDialogDragLeaveEvent
func callbackQDialogDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDialog::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDialog) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QDialog::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QDialog) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QDialog::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQDialogDragMoveEvent
func callbackQDialogDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDialog::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDialog) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QDialog::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QDialog) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QDialog::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQDialogDropEvent
func callbackQDialogDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDialog::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDialog) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDialog::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QDialog) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QDialog::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQDialogEnterEvent
func callbackQDialogEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDialog::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDialog) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QDialog::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QDialog) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QDialog::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQDialogFocusInEvent
func callbackQDialogFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDialog::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDialog) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QDialog::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QDialog) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QDialog::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQDialogFocusOutEvent
func callbackQDialogFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDialog::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDialog) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QDialog::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QDialog) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QDialog::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQDialogHideEvent
func callbackQDialogHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDialog::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDialog) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDialog::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QDialog) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QDialog::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQDialogLeaveEvent
func callbackQDialogLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDialog::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDialog) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QDialog::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QDialog) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QDialog::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQDialogMoveEvent
func callbackQDialogMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDialog::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDialog) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QDialog::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QDialog) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QDialog::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQDialogPaintEvent
func callbackQDialogPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDialog::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDialog) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDialog::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QDialog) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QDialog::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQDialogChangeEvent
func callbackQDialogChangeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDialog::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDialog) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QDialog::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QDialog) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QDialog::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQDialogInitPainter
func callbackQDialogInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QDialog::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QDialog) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QDialog::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QDialog) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QDialog::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQDialogInputMethodEvent
func callbackQDialogInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDialog::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDialog) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QDialog) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQDialogKeyReleaseEvent
func callbackQDialogKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDialog::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDialog) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QDialog) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQDialogMouseDoubleClickEvent
func callbackQDialogMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDialog::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDialog) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QDialog) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQDialogMouseMoveEvent
func callbackQDialogMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDialog::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDialog) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDialog::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QDialog) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QDialog::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQDialogMousePressEvent
func callbackQDialogMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDialog::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDialog) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QDialog) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQDialogMouseReleaseEvent
func callbackQDialogMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDialog::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDialog) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QDialog::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QDialog) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QDialog::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQDialogTabletEvent
func callbackQDialogTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDialog::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDialog) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QDialog::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QDialog) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QDialog::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQDialogWheelEvent
func callbackQDialogWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDialog::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDialog) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDialog::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDialog) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDialog::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDialogTimerEvent
func callbackQDialogTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDialog::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDialog) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDialog::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDialog) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDialog::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDialogChildEvent
func callbackQDialogChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDialog::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDialog) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDialog::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDialog) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDialog::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDialogCustomEvent
func callbackQDialogCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDialog::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
