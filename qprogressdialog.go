package qt

//#include "qprogressdialog.h"
import "C"

type qprogressdialog struct {
	qdialog
}

type QProgressDialog interface {
	QDialog
	AutoClose() bool
	AutoReset() bool
	LabelText() string
	Maximum() int
	Minimum() int
	MinimumDuration() int
	SetAutoClose_Bool(close bool)
	SetAutoReset_Bool(reset bool)
	SetBar_QProgressBar(bar QProgressBar)
	SetCancelButton_QPushButton(cancelButton QPushButton)
	SetLabel_QLabel(label QLabel)
	Value() int
	WasCanceled() bool
	ConnectSlotCancel()
	DisconnectSlotCancel()
	SlotCancel()
	ConnectSlotReset()
	DisconnectSlotReset()
	SlotReset()
	ConnectSlotSetMaximum()
	DisconnectSlotSetMaximum()
	SlotSetMaximum_Int(maximum int)
	ConnectSlotSetMinimum()
	DisconnectSlotSetMinimum()
	SlotSetMinimum_Int(minimum int)
	ConnectSlotSetMinimumDuration()
	DisconnectSlotSetMinimumDuration()
	SlotSetMinimumDuration_Int(ms int)
	ConnectSlotSetRange()
	DisconnectSlotSetRange()
	SlotSetRange_Int_Int(minimum int, maximum int)
	ConnectSlotSetValue()
	DisconnectSlotSetValue()
	SlotSetValue_Int(progress int)
	ConnectSignalCanceled(f func())
	DisconnectSignalCanceled()
	SignalCanceled() func()
}

func (p *qprogressdialog) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qprogressdialog) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQProgressDialog_QWidget_WindowType(parent QWidget, f WindowType) QProgressDialog {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qprogressdialog = new(qprogressdialog)
	qprogressdialog.SetPointer(C.QProgressDialog_New_QWidget_WindowType(parentPtr, C.int(f)))
	qprogressdialog.SetObjectName_String("QProgressDialog_" + randomIdentifier())
	return qprogressdialog
}

func NewQProgressDialog_String_String_Int_Int_QWidget_WindowType(labelText string, cancelButtonText string, minimum int, maximum int, parent QWidget, f WindowType) QProgressDialog {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qprogressdialog = new(qprogressdialog)
	qprogressdialog.SetPointer(C.QProgressDialog_New_String_String_Int_Int_QWidget_WindowType(C.CString(labelText), C.CString(cancelButtonText), C.int(minimum), C.int(maximum), parentPtr, C.int(f)))
	qprogressdialog.SetObjectName_String("QProgressDialog_" + randomIdentifier())
	return qprogressdialog
}

func (p *qprogressdialog) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QProgressDialog_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qprogressdialog) AutoClose() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QProgressDialog_AutoClose(p.Pointer()) != 0
	}
}

func (p *qprogressdialog) AutoReset() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QProgressDialog_AutoReset(p.Pointer()) != 0
	}
}

func (p *qprogressdialog) LabelText() string {
	if p.Pointer() == nil {
		return ""
	} else {
		return C.GoString(C.QProgressDialog_LabelText(p.Pointer()))
	}
}

func (p *qprogressdialog) Maximum() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QProgressDialog_Maximum(p.Pointer()))
	}
}

func (p *qprogressdialog) Minimum() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QProgressDialog_Minimum(p.Pointer()))
	}
}

func (p *qprogressdialog) MinimumDuration() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QProgressDialog_MinimumDuration(p.Pointer()))
	}
}

func (p *qprogressdialog) SetAutoClose_Bool(close bool) {
	if p.Pointer() != nil {
		C.QProgressDialog_SetAutoClose_Bool(p.Pointer(), goBoolToCInt(close))
	}
}

func (p *qprogressdialog) SetAutoReset_Bool(reset bool) {
	if p.Pointer() != nil {
		C.QProgressDialog_SetAutoReset_Bool(p.Pointer(), goBoolToCInt(reset))
	}
}

func (p *qprogressdialog) SetBar_QProgressBar(bar QProgressBar) {
	if p.Pointer() == nil {
	} else {
		var barPtr C.QtObjectPtr = nil
		if bar != nil {
			barPtr = bar.Pointer()
		}
		C.QProgressDialog_SetBar_QProgressBar(p.Pointer(), barPtr)
	}
}

func (p *qprogressdialog) SetCancelButton_QPushButton(cancelButton QPushButton) {
	if p.Pointer() == nil {
	} else {
		var cancelButtonPtr C.QtObjectPtr = nil
		if cancelButton != nil {
			cancelButtonPtr = cancelButton.Pointer()
		}
		C.QProgressDialog_SetCancelButton_QPushButton(p.Pointer(), cancelButtonPtr)
	}
}

func (p *qprogressdialog) SetLabel_QLabel(label QLabel) {
	if p.Pointer() == nil {
	} else {
		var labelPtr C.QtObjectPtr = nil
		if label != nil {
			labelPtr = label.Pointer()
		}
		C.QProgressDialog_SetLabel_QLabel(p.Pointer(), labelPtr)
	}
}

func (p *qprogressdialog) Value() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QProgressDialog_Value(p.Pointer()))
	}
}

func (p *qprogressdialog) WasCanceled() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QProgressDialog_WasCanceled(p.Pointer()) != 0
	}
}

func (p *qprogressdialog) ConnectSlotCancel() {
	C.QProgressDialog_ConnectSlotCancel(p.Pointer())
}

func (p *qprogressdialog) DisconnectSlotCancel() {
	C.QProgressDialog_DisconnectSlotCancel(p.Pointer())
}

func (p *qprogressdialog) SlotCancel() {
	if p.Pointer() != nil {
		C.QProgressDialog_Cancel(p.Pointer())
	}
}

func (p *qprogressdialog) ConnectSlotReset() {
	C.QProgressDialog_ConnectSlotReset(p.Pointer())
}

func (p *qprogressdialog) DisconnectSlotReset() {
	C.QProgressDialog_DisconnectSlotReset(p.Pointer())
}

func (p *qprogressdialog) SlotReset() {
	if p.Pointer() != nil {
		C.QProgressDialog_Reset(p.Pointer())
	}
}

func (p *qprogressdialog) ConnectSlotSetMaximum() {
	C.QProgressDialog_ConnectSlotSetMaximum(p.Pointer())
}

func (p *qprogressdialog) DisconnectSlotSetMaximum() {
	C.QProgressDialog_DisconnectSlotSetMaximum(p.Pointer())
}

func (p *qprogressdialog) SlotSetMaximum_Int(maximum int) {
	if p.Pointer() != nil {
		C.QProgressDialog_SetMaximum_Int(p.Pointer(), C.int(maximum))
	}
}

func (p *qprogressdialog) ConnectSlotSetMinimum() {
	C.QProgressDialog_ConnectSlotSetMinimum(p.Pointer())
}

func (p *qprogressdialog) DisconnectSlotSetMinimum() {
	C.QProgressDialog_DisconnectSlotSetMinimum(p.Pointer())
}

func (p *qprogressdialog) SlotSetMinimum_Int(minimum int) {
	if p.Pointer() != nil {
		C.QProgressDialog_SetMinimum_Int(p.Pointer(), C.int(minimum))
	}
}

func (p *qprogressdialog) ConnectSlotSetMinimumDuration() {
	C.QProgressDialog_ConnectSlotSetMinimumDuration(p.Pointer())
}

func (p *qprogressdialog) DisconnectSlotSetMinimumDuration() {
	C.QProgressDialog_DisconnectSlotSetMinimumDuration(p.Pointer())
}

func (p *qprogressdialog) SlotSetMinimumDuration_Int(ms int) {
	if p.Pointer() != nil {
		C.QProgressDialog_SetMinimumDuration_Int(p.Pointer(), C.int(ms))
	}
}

func (p *qprogressdialog) ConnectSlotSetRange() {
	C.QProgressDialog_ConnectSlotSetRange(p.Pointer())
}

func (p *qprogressdialog) DisconnectSlotSetRange() {
	C.QProgressDialog_DisconnectSlotSetRange(p.Pointer())
}

func (p *qprogressdialog) SlotSetRange_Int_Int(minimum int, maximum int) {
	if p.Pointer() != nil {
		C.QProgressDialog_SetRange_Int_Int(p.Pointer(), C.int(minimum), C.int(maximum))
	}
}

func (p *qprogressdialog) ConnectSlotSetValue() {
	C.QProgressDialog_ConnectSlotSetValue(p.Pointer())
}

func (p *qprogressdialog) DisconnectSlotSetValue() {
	C.QProgressDialog_DisconnectSlotSetValue(p.Pointer())
}

func (p *qprogressdialog) SlotSetValue_Int(progress int) {
	if p.Pointer() != nil {
		C.QProgressDialog_SetValue_Int(p.Pointer(), C.int(progress))
	}
}

func (p *qprogressdialog) ConnectSignalCanceled(f func()) {
	C.QProgressDialog_ConnectSignalCanceled(p.Pointer())
	connectSignal(p.ObjectName(), "canceled", f)
}

func (p *qprogressdialog) DisconnectSignalCanceled() {
	C.QProgressDialog_DisconnectSignalCanceled(p.Pointer())
	disconnectSignal(p.ObjectName(), "canceled")
}

func (p *qprogressdialog) SignalCanceled() func() {
	return getSignal(p.ObjectName(), "canceled")
}

//export callbackQProgressDialog
func callbackQProgressDialog(ptr C.QtObjectPtr, msg *C.char) {
	var qprogressdialog = new(qprogressdialog)
	qprogressdialog.SetPointer(ptr)
	getSignal(qprogressdialog.ObjectName(), C.GoString(msg))()
}
