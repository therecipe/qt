package qt

//#include "qgroupbox.h"
import "C"

type qgroupbox struct {
	qwidget
}

type QGroupBox interface {
	QWidget
	Alignment() AlignmentFlag
	IsCheckable() bool
	IsChecked() bool
	IsFlat() bool
	SetAlignment(alignment int)
	SetCheckable(checkable bool)
	SetFlat(flat bool)
	SetTitle(title string)
	Title() string
	ConnectSlotSetChecked()
	DisconnectSlotSetChecked()
	SlotSetChecked(checked bool)
	ConnectSignalClicked(f func())
	DisconnectSignalClicked()
	SignalClicked() func()
	ConnectSignalToggled(f func())
	DisconnectSignalToggled()
	SignalToggled() func()
}

func (p *qgroupbox) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qgroupbox) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQGroupBox1(parent QWidget) QGroupBox {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qgroupbox = new(qgroupbox)
	qgroupbox.SetPointer(C.QGroupBox_New_QWidget(parentPtr))
	qgroupbox.SetObjectName("QGroupBox_" + randomIdentifier())
	return qgroupbox
}

func NewQGroupBox2(title string, parent QWidget) QGroupBox {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qgroupbox = new(qgroupbox)
	qgroupbox.SetPointer(C.QGroupBox_New_String_QWidget(C.CString(title), parentPtr))
	qgroupbox.SetObjectName("QGroupBox_" + randomIdentifier())
	return qgroupbox
}

func (p *qgroupbox) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QGroupBox_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qgroupbox) Alignment() AlignmentFlag {
	if p.Pointer() == nil {
		return 0
	}
	return AlignmentFlag(C.QGroupBox_Alignment(p.Pointer()))
}

func (p *qgroupbox) IsCheckable() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QGroupBox_IsCheckable(p.Pointer()) != 0
}

func (p *qgroupbox) IsChecked() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QGroupBox_IsChecked(p.Pointer()) != 0
}

func (p *qgroupbox) IsFlat() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QGroupBox_IsFlat(p.Pointer()) != 0
}

func (p *qgroupbox) SetAlignment(alignment int) {
	if p.Pointer() != nil {
		C.QGroupBox_SetAlignment_Int(p.Pointer(), C.int(alignment))
	}
}

func (p *qgroupbox) SetCheckable(checkable bool) {
	if p.Pointer() != nil {
		C.QGroupBox_SetCheckable_Bool(p.Pointer(), goBoolToCInt(checkable))
	}
}

func (p *qgroupbox) SetFlat(flat bool) {
	if p.Pointer() != nil {
		C.QGroupBox_SetFlat_Bool(p.Pointer(), goBoolToCInt(flat))
	}
}

func (p *qgroupbox) SetTitle(title string) {
	if p.Pointer() != nil {
		C.QGroupBox_SetTitle_String(p.Pointer(), C.CString(title))
	}
}

func (p *qgroupbox) Title() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QGroupBox_Title(p.Pointer()))
}

func (p *qgroupbox) ConnectSlotSetChecked() {
	C.QGroupBox_ConnectSlotSetChecked(p.Pointer())
}

func (p *qgroupbox) DisconnectSlotSetChecked() {
	C.QGroupBox_DisconnectSlotSetChecked(p.Pointer())
}

func (p *qgroupbox) SlotSetChecked(checked bool) {
	if p.Pointer() != nil {
		C.QGroupBox_SetChecked_Bool(p.Pointer(), goBoolToCInt(checked))
	}
}

func (p *qgroupbox) ConnectSignalClicked(f func()) {
	C.QGroupBox_ConnectSignalClicked(p.Pointer())
	connectSignal(p.ObjectName(), "clicked", f)
}

func (p *qgroupbox) DisconnectSignalClicked() {
	C.QGroupBox_DisconnectSignalClicked(p.Pointer())
	disconnectSignal(p.ObjectName(), "clicked")
}

func (p *qgroupbox) SignalClicked() func() {
	return getSignal(p.ObjectName(), "clicked")
}

func (p *qgroupbox) ConnectSignalToggled(f func()) {
	C.QGroupBox_ConnectSignalToggled(p.Pointer())
	connectSignal(p.ObjectName(), "toggled", f)
}

func (p *qgroupbox) DisconnectSignalToggled() {
	C.QGroupBox_DisconnectSignalToggled(p.Pointer())
	disconnectSignal(p.ObjectName(), "toggled")
}

func (p *qgroupbox) SignalToggled() func() {
	return getSignal(p.ObjectName(), "toggled")
}

//export callbackQGroupBox
func callbackQGroupBox(ptr C.QtObjectPtr, msg *C.char) {
	var qgroupbox = new(qgroupbox)
	qgroupbox.SetPointer(ptr)
	getSignal(qgroupbox.ObjectName(), C.GoString(msg))()
}
