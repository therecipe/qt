package qt

//#include "qdialog.h"
import "C"

type qdialog struct {
	qwidget
}

type QDialog interface {
	QWidget
	IsSizeGripEnabled() bool
	Result() int
	SetModal_Bool(modal bool)
	SetResult_Int(i int)
	SetSizeGripEnabled_Bool(sizeGripEnabled bool)
	ConnectSlotAccept()
	DisconnectSlotAccept()
	SlotAccept()
	ConnectSlotDone()
	DisconnectSlotDone()
	SlotDone_Int(r int)
	ConnectSlotExec()
	DisconnectSlotExec()
	SlotExec()
	ConnectSlotOpen()
	DisconnectSlotOpen()
	SlotOpen()
	ConnectSlotReject()
	DisconnectSlotReject()
	SlotReject()
	ConnectSignalAccepted(f func())
	DisconnectSignalAccepted()
	SignalAccepted() func()
	ConnectSignalFinished(f func())
	DisconnectSignalFinished()
	SignalFinished() func()
	ConnectSignalRejected(f func())
	DisconnectSignalRejected()
	SignalRejected() func()
}

func (p *qdialog) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qdialog) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQDialog_QWidget_WindowType(parent QWidget, f WindowType) QDialog {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qdialog = new(qdialog)
	qdialog.SetPointer(C.QDialog_New_QWidget_WindowType(parentPtr, C.int(f)))
	qdialog.SetObjectName_String("QDialog_" + randomIdentifier())
	return qdialog
}

func (p *qdialog) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QDialog_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qdialog) IsSizeGripEnabled() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QDialog_IsSizeGripEnabled(p.Pointer()) != 0
	}
}

func (p *qdialog) Result() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QDialog_Result(p.Pointer()))
	}
}

func (p *qdialog) SetModal_Bool(modal bool) {
	if p.Pointer() != nil {
		C.QDialog_SetModal_Bool(p.Pointer(), goBoolToCInt(modal))
	}
}

func (p *qdialog) SetResult_Int(i int) {
	if p.Pointer() != nil {
		C.QDialog_SetResult_Int(p.Pointer(), C.int(i))
	}
}

func (p *qdialog) SetSizeGripEnabled_Bool(sizeGripEnabled bool) {
	if p.Pointer() != nil {
		C.QDialog_SetSizeGripEnabled_Bool(p.Pointer(), goBoolToCInt(sizeGripEnabled))
	}
}

func (p *qdialog) ConnectSlotAccept() {
	C.QDialog_ConnectSlotAccept(p.Pointer())
}

func (p *qdialog) DisconnectSlotAccept() {
	C.QDialog_DisconnectSlotAccept(p.Pointer())
}

func (p *qdialog) SlotAccept() {
	if p.Pointer() != nil {
		C.QDialog_Accept(p.Pointer())
	}
}

func (p *qdialog) ConnectSlotDone() {
	C.QDialog_ConnectSlotDone(p.Pointer())
}

func (p *qdialog) DisconnectSlotDone() {
	C.QDialog_DisconnectSlotDone(p.Pointer())
}

func (p *qdialog) SlotDone_Int(r int) {
	if p.Pointer() != nil {
		C.QDialog_Done_Int(p.Pointer(), C.int(r))
	}
}

func (p *qdialog) ConnectSlotExec() {
	C.QDialog_ConnectSlotExec(p.Pointer())
}

func (p *qdialog) DisconnectSlotExec() {
	C.QDialog_DisconnectSlotExec(p.Pointer())
}

func (p *qdialog) SlotExec() {
	if p.Pointer() != nil {
		C.QDialog_Exec(p.Pointer())
	}
}

func (p *qdialog) ConnectSlotOpen() {
	C.QDialog_ConnectSlotOpen(p.Pointer())
}

func (p *qdialog) DisconnectSlotOpen() {
	C.QDialog_DisconnectSlotOpen(p.Pointer())
}

func (p *qdialog) SlotOpen() {
	if p.Pointer() != nil {
		C.QDialog_Open(p.Pointer())
	}
}

func (p *qdialog) ConnectSlotReject() {
	C.QDialog_ConnectSlotReject(p.Pointer())
}

func (p *qdialog) DisconnectSlotReject() {
	C.QDialog_DisconnectSlotReject(p.Pointer())
}

func (p *qdialog) SlotReject() {
	if p.Pointer() != nil {
		C.QDialog_Reject(p.Pointer())
	}
}

func (p *qdialog) ConnectSignalAccepted(f func()) {
	C.QDialog_ConnectSignalAccepted(p.Pointer())
	connectSignal(p.ObjectName(), "accepted", f)
}

func (p *qdialog) DisconnectSignalAccepted() {
	C.QDialog_DisconnectSignalAccepted(p.Pointer())
	disconnectSignal(p.ObjectName(), "accepted")
}

func (p *qdialog) SignalAccepted() func() {
	return getSignal(p.ObjectName(), "accepted")
}

func (p *qdialog) ConnectSignalFinished(f func()) {
	C.QDialog_ConnectSignalFinished(p.Pointer())
	connectSignal(p.ObjectName(), "finished", f)
}

func (p *qdialog) DisconnectSignalFinished() {
	C.QDialog_DisconnectSignalFinished(p.Pointer())
	disconnectSignal(p.ObjectName(), "finished")
}

func (p *qdialog) SignalFinished() func() {
	return getSignal(p.ObjectName(), "finished")
}

func (p *qdialog) ConnectSignalRejected(f func()) {
	C.QDialog_ConnectSignalRejected(p.Pointer())
	connectSignal(p.ObjectName(), "rejected", f)
}

func (p *qdialog) DisconnectSignalRejected() {
	C.QDialog_DisconnectSignalRejected(p.Pointer())
	disconnectSignal(p.ObjectName(), "rejected")
}

func (p *qdialog) SignalRejected() func() {
	return getSignal(p.ObjectName(), "rejected")
}

//export callbackQDialog
func callbackQDialog(ptr C.QtObjectPtr, msg *C.char) {
	var qdialog = new(qdialog)
	qdialog.SetPointer(ptr)
	getSignal(qdialog.ObjectName(), C.GoString(msg))()
}
