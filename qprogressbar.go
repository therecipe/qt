package qt

//#include "qprogressbar.h"
import "C"

type qprogressbar struct {
	qwidget
}

type QProgressBar interface {
	QWidget
	Alignment() AlignmentFlag
	Format() string
	InvertedAppearance() bool
	IsTextVisible() bool
	Maximum() int
	Minimum() int
	Orientation() Orientation
	ResetFormat()
	SetAlignment(alignment AlignmentFlag)
	SetFormat(format string)
	SetInvertedAppearance(invert bool)
	SetTextVisible(visible bool)
	Text() string
	Value() int
	ConnectSlotReset()
	DisconnectSlotReset()
	SlotReset()
	ConnectSlotSetMaximum()
	DisconnectSlotSetMaximum()
	SlotSetMaximum(maximum int)
	ConnectSlotSetMinimum()
	DisconnectSlotSetMinimum()
	SlotSetMinimum(minimum int)
	ConnectSlotSetOrientation()
	DisconnectSlotSetOrientation()
	SlotSetOrientation(orientation Orientation)
	ConnectSlotSetRange()
	DisconnectSlotSetRange()
	SlotSetRange(minimum int, maximum int)
	ConnectSlotSetValue()
	DisconnectSlotSetValue()
	SlotSetValue(value int)
	ConnectSignalValueChanged(f func())
	DisconnectSignalValueChanged()
	SignalValueChanged() func()
}

func (p *qprogressbar) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qprogressbar) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQProgressBar(parent QWidget) QProgressBar {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qprogressbar = new(qprogressbar)
	qprogressbar.SetPointer(C.QProgressBar_New_QWidget(parentPtr))
	qprogressbar.SetObjectName("QProgressBar_" + randomIdentifier())
	return qprogressbar
}

func (p *qprogressbar) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QProgressBar_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qprogressbar) Alignment() AlignmentFlag {
	if p.Pointer() == nil {
		return 0
	}
	return AlignmentFlag(C.QProgressBar_Alignment(p.Pointer()))
}

func (p *qprogressbar) Format() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QProgressBar_Format(p.Pointer()))
}

func (p *qprogressbar) InvertedAppearance() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QProgressBar_InvertedAppearance(p.Pointer()) != 0
}

func (p *qprogressbar) IsTextVisible() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QProgressBar_IsTextVisible(p.Pointer()) != 0
}

func (p *qprogressbar) Maximum() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QProgressBar_Maximum(p.Pointer()))
}

func (p *qprogressbar) Minimum() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QProgressBar_Minimum(p.Pointer()))
}

func (p *qprogressbar) Orientation() Orientation {
	if p.Pointer() == nil {
		return 0
	}
	return Orientation(C.QProgressBar_Orientation(p.Pointer()))
}

func (p *qprogressbar) ResetFormat() {
	if p.Pointer() != nil {
		C.QProgressBar_ResetFormat(p.Pointer())
	}
}

func (p *qprogressbar) SetAlignment(alignment AlignmentFlag) {
	if p.Pointer() != nil {
		C.QProgressBar_SetAlignment_AlignmentFlag(p.Pointer(), C.int(alignment))
	}
}

func (p *qprogressbar) SetFormat(format string) {
	if p.Pointer() != nil {
		C.QProgressBar_SetFormat_String(p.Pointer(), C.CString(format))
	}
}

func (p *qprogressbar) SetInvertedAppearance(invert bool) {
	if p.Pointer() != nil {
		C.QProgressBar_SetInvertedAppearance_Bool(p.Pointer(), goBoolToCInt(invert))
	}
}

func (p *qprogressbar) SetTextVisible(visible bool) {
	if p.Pointer() != nil {
		C.QProgressBar_SetTextVisible_Bool(p.Pointer(), goBoolToCInt(visible))
	}
}

func (p *qprogressbar) Text() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QProgressBar_Text(p.Pointer()))
}

func (p *qprogressbar) Value() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QProgressBar_Value(p.Pointer()))
}

func (p *qprogressbar) ConnectSlotReset() {
	C.QProgressBar_ConnectSlotReset(p.Pointer())
}

func (p *qprogressbar) DisconnectSlotReset() {
	C.QProgressBar_DisconnectSlotReset(p.Pointer())
}

func (p *qprogressbar) SlotReset() {
	if p.Pointer() != nil {
		C.QProgressBar_Reset(p.Pointer())
	}
}

func (p *qprogressbar) ConnectSlotSetMaximum() {
	C.QProgressBar_ConnectSlotSetMaximum(p.Pointer())
}

func (p *qprogressbar) DisconnectSlotSetMaximum() {
	C.QProgressBar_DisconnectSlotSetMaximum(p.Pointer())
}

func (p *qprogressbar) SlotSetMaximum(maximum int) {
	if p.Pointer() != nil {
		C.QProgressBar_SetMaximum_Int(p.Pointer(), C.int(maximum))
	}
}

func (p *qprogressbar) ConnectSlotSetMinimum() {
	C.QProgressBar_ConnectSlotSetMinimum(p.Pointer())
}

func (p *qprogressbar) DisconnectSlotSetMinimum() {
	C.QProgressBar_DisconnectSlotSetMinimum(p.Pointer())
}

func (p *qprogressbar) SlotSetMinimum(minimum int) {
	if p.Pointer() != nil {
		C.QProgressBar_SetMinimum_Int(p.Pointer(), C.int(minimum))
	}
}

func (p *qprogressbar) ConnectSlotSetOrientation() {
	C.QProgressBar_ConnectSlotSetOrientation(p.Pointer())
}

func (p *qprogressbar) DisconnectSlotSetOrientation() {
	C.QProgressBar_DisconnectSlotSetOrientation(p.Pointer())
}

func (p *qprogressbar) SlotSetOrientation(orientation Orientation) {
	if p.Pointer() != nil {
		C.QProgressBar_SetOrientation_Orientation(p.Pointer(), C.int(orientation))
	}
}

func (p *qprogressbar) ConnectSlotSetRange() {
	C.QProgressBar_ConnectSlotSetRange(p.Pointer())
}

func (p *qprogressbar) DisconnectSlotSetRange() {
	C.QProgressBar_DisconnectSlotSetRange(p.Pointer())
}

func (p *qprogressbar) SlotSetRange(minimum int, maximum int) {
	if p.Pointer() != nil {
		C.QProgressBar_SetRange_Int_Int(p.Pointer(), C.int(minimum), C.int(maximum))
	}
}

func (p *qprogressbar) ConnectSlotSetValue() {
	C.QProgressBar_ConnectSlotSetValue(p.Pointer())
}

func (p *qprogressbar) DisconnectSlotSetValue() {
	C.QProgressBar_DisconnectSlotSetValue(p.Pointer())
}

func (p *qprogressbar) SlotSetValue(value int) {
	if p.Pointer() != nil {
		C.QProgressBar_SetValue_Int(p.Pointer(), C.int(value))
	}
}

func (p *qprogressbar) ConnectSignalValueChanged(f func()) {
	C.QProgressBar_ConnectSignalValueChanged(p.Pointer())
	connectSignal(p.ObjectName(), "valueChanged", f)
}

func (p *qprogressbar) DisconnectSignalValueChanged() {
	C.QProgressBar_DisconnectSignalValueChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "valueChanged")
}

func (p *qprogressbar) SignalValueChanged() func() {
	return getSignal(p.ObjectName(), "valueChanged")
}

//export callbackQProgressBar
func callbackQProgressBar(ptr C.QtObjectPtr, msg *C.char) {
	var qprogressbar = new(qprogressbar)
	qprogressbar.SetPointer(ptr)
	getSignal(qprogressbar.ObjectName(), C.GoString(msg))()
}
