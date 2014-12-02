package qt

//#include "qcheckbox.h"
import "C"

type qcheckbox struct {
	qabstractbutton
}

type QCheckBox interface {
	QAbstractButton
	CheckState() CheckState
	IsTristate() bool
	SetCheckState_CheckState(state CheckState)
	SetTristate_Bool(y bool)
	ConnectSignalStateChanged(f func())
	DisconnectSignalStateChanged()
	SignalStateChanged() func()
}

func (p *qcheckbox) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qcheckbox) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQCheckBox_QWidget(parent QWidget) QCheckBox {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qcheckbox = new(qcheckbox)
	qcheckbox.SetPointer(C.QCheckBox_New_QWidget(parentPtr))
	qcheckbox.SetObjectName_String("QCheckBox_" + randomIdentifier())
	return qcheckbox
}

func NewQCheckBox_String_QWidget(text string, parent QWidget) QCheckBox {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qcheckbox = new(qcheckbox)
	qcheckbox.SetPointer(C.QCheckBox_New_String_QWidget(C.CString(text), parentPtr))
	qcheckbox.SetObjectName_String("QCheckBox_" + randomIdentifier())
	return qcheckbox
}

func (p *qcheckbox) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QCheckBox_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qcheckbox) CheckState() CheckState {
	if p.Pointer() == nil {
		return 0
	} else {
		return CheckState(C.QCheckBox_CheckState(p.Pointer()))
	}
}

func (p *qcheckbox) IsTristate() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QCheckBox_IsTristate(p.Pointer()) != 0
	}
}

func (p *qcheckbox) SetCheckState_CheckState(state CheckState) {
	if p.Pointer() != nil {
		C.QCheckBox_SetCheckState_CheckState(p.Pointer(), C.int(state))
	}
}

func (p *qcheckbox) SetTristate_Bool(y bool) {
	if p.Pointer() != nil {
		C.QCheckBox_SetTristate_Bool(p.Pointer(), goBoolToCInt(y))
	}
}

func (p *qcheckbox) ConnectSignalStateChanged(f func()) {
	C.QCheckBox_ConnectSignalStateChanged(p.Pointer())
	connectSignal(p.ObjectName(), "stateChanged", f)
}

func (p *qcheckbox) DisconnectSignalStateChanged() {
	C.QCheckBox_DisconnectSignalStateChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "stateChanged")
}

func (p *qcheckbox) SignalStateChanged() func() {
	return getSignal(p.ObjectName(), "stateChanged")
}

//export callbackQCheckBox
func callbackQCheckBox(ptr C.QtObjectPtr, msg *C.char) {
	var qcheckbox = new(qcheckbox)
	qcheckbox.SetPointer(ptr)
	getSignal(qcheckbox.ObjectName(), C.GoString(msg))()
}
