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

	var signal = qt.GetSignal(C.GoString(ptrName), "accept")
	if signal != nil {
		defer signal.(func())()
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

	var signal = qt.GetSignal(C.GoString(ptrName), "accepted")
	if signal != nil {
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

	var signal = qt.GetSignal(C.GoString(ptrName), "closeEvent")
	if signal != nil {
		defer signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(e))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "contextMenuEvent")
	if signal != nil {
		defer signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "done")
	if signal != nil {
		defer signal.(func(int))(int(r))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "finished")
	if signal != nil {
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

	var signal = qt.GetSignal(C.GoString(ptrName), "keyPressEvent")
	if signal != nil {
		defer signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "open")
	if signal != nil {
		defer signal.(func())()
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

	var signal = qt.GetSignal(C.GoString(ptrName), "reject")
	if signal != nil {
		defer signal.(func())()
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

	var signal = qt.GetSignal(C.GoString(ptrName), "rejected")
	if signal != nil {
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

	var signal = qt.GetSignal(C.GoString(ptrName), "resizeEvent")
	if signal != nil {
		defer signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(v))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "setVisible")
	if signal != nil {
		defer signal.(func(bool))(int(visible) != 0)
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

	var signal = qt.GetSignal(C.GoString(ptrName), "showEvent")
	if signal != nil {
		defer signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
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
