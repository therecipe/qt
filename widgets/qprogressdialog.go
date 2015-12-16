package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QProgressDialog struct {
	QDialog
}

type QProgressDialog_ITF interface {
	QDialog_ITF
	QProgressDialog_PTR() *QProgressDialog
}

func PointerFromQProgressDialog(ptr QProgressDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProgressDialog_PTR().Pointer()
	}
	return nil
}

func NewQProgressDialogFromPointer(ptr unsafe.Pointer) *QProgressDialog {
	var n = new(QProgressDialog)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QProgressDialog_") {
		n.SetObjectName("QProgressDialog_" + qt.Identifier())
	}
	return n
}

func (ptr *QProgressDialog) QProgressDialog_PTR() *QProgressDialog {
	return ptr
}

func (ptr *QProgressDialog) AutoClose() bool {
	defer qt.Recovering("QProgressDialog::autoClose")

	if ptr.Pointer() != nil {
		return C.QProgressDialog_AutoClose(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QProgressDialog) AutoReset() bool {
	defer qt.Recovering("QProgressDialog::autoReset")

	if ptr.Pointer() != nil {
		return C.QProgressDialog_AutoReset(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QProgressDialog) LabelText() string {
	defer qt.Recovering("QProgressDialog::labelText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QProgressDialog_LabelText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QProgressDialog) Maximum() int {
	defer qt.Recovering("QProgressDialog::maximum")

	if ptr.Pointer() != nil {
		return int(C.QProgressDialog_Maximum(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressDialog) Minimum() int {
	defer qt.Recovering("QProgressDialog::minimum")

	if ptr.Pointer() != nil {
		return int(C.QProgressDialog_Minimum(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressDialog) MinimumDuration() int {
	defer qt.Recovering("QProgressDialog::minimumDuration")

	if ptr.Pointer() != nil {
		return int(C.QProgressDialog_MinimumDuration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressDialog) SetAutoClose(close bool) {
	defer qt.Recovering("QProgressDialog::setAutoClose")

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetAutoClose(ptr.Pointer(), C.int(qt.GoBoolToInt(close)))
	}
}

func (ptr *QProgressDialog) SetAutoReset(reset bool) {
	defer qt.Recovering("QProgressDialog::setAutoReset")

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetAutoReset(ptr.Pointer(), C.int(qt.GoBoolToInt(reset)))
	}
}

func (ptr *QProgressDialog) SetLabelText(text string) {
	defer qt.Recovering("QProgressDialog::setLabelText")

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetLabelText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QProgressDialog) SetMaximum(maximum int) {
	defer qt.Recovering("QProgressDialog::setMaximum")

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetMaximum(ptr.Pointer(), C.int(maximum))
	}
}

func (ptr *QProgressDialog) SetMinimum(minimum int) {
	defer qt.Recovering("QProgressDialog::setMinimum")

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetMinimum(ptr.Pointer(), C.int(minimum))
	}
}

func (ptr *QProgressDialog) SetMinimumDuration(ms int) {
	defer qt.Recovering("QProgressDialog::setMinimumDuration")

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetMinimumDuration(ptr.Pointer(), C.int(ms))
	}
}

func (ptr *QProgressDialog) SetValue(progress int) {
	defer qt.Recovering("QProgressDialog::setValue")

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetValue(ptr.Pointer(), C.int(progress))
	}
}

func (ptr *QProgressDialog) Value() int {
	defer qt.Recovering("QProgressDialog::value")

	if ptr.Pointer() != nil {
		return int(C.QProgressDialog_Value(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressDialog) WasCanceled() bool {
	defer qt.Recovering("QProgressDialog::wasCanceled")

	if ptr.Pointer() != nil {
		return C.QProgressDialog_WasCanceled(ptr.Pointer()) != 0
	}
	return false
}

func NewQProgressDialog(parent QWidget_ITF, f core.Qt__WindowType) *QProgressDialog {
	defer qt.Recovering("QProgressDialog::QProgressDialog")

	return NewQProgressDialogFromPointer(C.QProgressDialog_NewQProgressDialog(PointerFromQWidget(parent), C.int(f)))
}

func NewQProgressDialog2(labelText string, cancelButtonText string, minimum int, maximum int, parent QWidget_ITF, f core.Qt__WindowType) *QProgressDialog {
	defer qt.Recovering("QProgressDialog::QProgressDialog")

	return NewQProgressDialogFromPointer(C.QProgressDialog_NewQProgressDialog2(C.CString(labelText), C.CString(cancelButtonText), C.int(minimum), C.int(maximum), PointerFromQWidget(parent), C.int(f)))
}

func (ptr *QProgressDialog) Cancel() {
	defer qt.Recovering("QProgressDialog::cancel")

	if ptr.Pointer() != nil {
		C.QProgressDialog_Cancel(ptr.Pointer())
	}
}

func (ptr *QProgressDialog) ConnectCanceled(f func()) {
	defer qt.Recovering("connect QProgressDialog::canceled")

	if ptr.Pointer() != nil {
		C.QProgressDialog_ConnectCanceled(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "canceled", f)
	}
}

func (ptr *QProgressDialog) DisconnectCanceled() {
	defer qt.Recovering("disconnect QProgressDialog::canceled")

	if ptr.Pointer() != nil {
		C.QProgressDialog_DisconnectCanceled(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "canceled")
	}
}

//export callbackQProgressDialogCanceled
func callbackQProgressDialogCanceled(ptrName *C.char) {
	defer qt.Recovering("callback QProgressDialog::canceled")

	var signal = qt.GetSignal(C.GoString(ptrName), "canceled")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QProgressDialog) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QProgressDialog::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QProgressDialog::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQProgressDialogChangeEvent
func callbackQProgressDialogChangeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::changeEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "changeEvent")
	if signal != nil {
		defer signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectCloseEvent(f func(e *gui.QCloseEvent)) {
	defer qt.Recovering("connect QProgressDialog::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QProgressDialog::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQProgressDialogCloseEvent
func callbackQProgressDialogCloseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::closeEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "closeEvent")
	if signal != nil {
		defer signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QProgressDialog) Open(receiver core.QObject_ITF, member string) {
	defer qt.Recovering("QProgressDialog::open")

	if ptr.Pointer() != nil {
		C.QProgressDialog_Open(ptr.Pointer(), core.PointerFromQObject(receiver), C.CString(member))
	}
}

func (ptr *QProgressDialog) Reset() {
	defer qt.Recovering("QProgressDialog::reset")

	if ptr.Pointer() != nil {
		C.QProgressDialog_Reset(ptr.Pointer())
	}
}

func (ptr *QProgressDialog) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QProgressDialog::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QProgressDialog::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQProgressDialogResizeEvent
func callbackQProgressDialogResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::resizeEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "resizeEvent")
	if signal != nil {
		defer signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QProgressDialog) SetBar(bar QProgressBar_ITF) {
	defer qt.Recovering("QProgressDialog::setBar")

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetBar(ptr.Pointer(), PointerFromQProgressBar(bar))
	}
}

func (ptr *QProgressDialog) SetCancelButton(cancelButton QPushButton_ITF) {
	defer qt.Recovering("QProgressDialog::setCancelButton")

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetCancelButton(ptr.Pointer(), PointerFromQPushButton(cancelButton))
	}
}

func (ptr *QProgressDialog) SetCancelButtonText(cancelButtonText string) {
	defer qt.Recovering("QProgressDialog::setCancelButtonText")

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetCancelButtonText(ptr.Pointer(), C.CString(cancelButtonText))
	}
}

func (ptr *QProgressDialog) SetLabel(label QLabel_ITF) {
	defer qt.Recovering("QProgressDialog::setLabel")

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetLabel(ptr.Pointer(), PointerFromQLabel(label))
	}
}

func (ptr *QProgressDialog) SetRange(minimum int, maximum int) {
	defer qt.Recovering("QProgressDialog::setRange")

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetRange(ptr.Pointer(), C.int(minimum), C.int(maximum))
	}
}

func (ptr *QProgressDialog) ConnectShowEvent(f func(e *gui.QShowEvent)) {
	defer qt.Recovering("connect QProgressDialog::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QProgressDialog::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQProgressDialogShowEvent
func callbackQProgressDialogShowEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::showEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "showEvent")
	if signal != nil {
		defer signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QProgressDialog) SizeHint() *core.QSize {
	defer qt.Recovering("QProgressDialog::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QProgressDialog_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QProgressDialog) DestroyQProgressDialog() {
	defer qt.Recovering("QProgressDialog::~QProgressDialog")

	if ptr.Pointer() != nil {
		C.QProgressDialog_DestroyQProgressDialog(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
