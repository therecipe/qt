package qt

//#include "qspinbox.h"
import "C"

type qspinbox struct {
	qabstractspinbox
}

type QSpinBox interface {
	QAbstractSpinBox
	CleanText() string
	DisplayIntegerBase() int
	Maximum() int
	Minimum() int
	Prefix() string
	SetDisplayIntegerBase(base int)
	SetMaximum(max int)
	SetMinimum(min int)
	SetPrefix(prefix string)
	SetRange(minimum int, maximum int)
	SetSingleStep(val int)
	SetSuffix(suffix string)
	SingleStep() int
	Suffix() string
	Value() int
	ConnectSlotSetValue()
	DisconnectSlotSetValue()
	SlotSetValue(val int)
}

func (p *qspinbox) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qspinbox) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQSpinBox(parent QWidget) QSpinBox {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qspinbox = new(qspinbox)
	qspinbox.SetPointer(C.QSpinBox_New_QWidget(parentPtr))
	qspinbox.SetObjectName("QSpinBox_" + randomIdentifier())
	return qspinbox
}

func (p *qspinbox) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QSpinBox_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qspinbox) CleanText() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QSpinBox_CleanText(p.Pointer()))
}

func (p *qspinbox) DisplayIntegerBase() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QSpinBox_DisplayIntegerBase(p.Pointer()))
}

func (p *qspinbox) Maximum() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QSpinBox_Maximum(p.Pointer()))
}

func (p *qspinbox) Minimum() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QSpinBox_Minimum(p.Pointer()))
}

func (p *qspinbox) Prefix() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QSpinBox_Prefix(p.Pointer()))
}

func (p *qspinbox) SetDisplayIntegerBase(base int) {
	if p.Pointer() != nil {
		C.QSpinBox_SetDisplayIntegerBase_Int(p.Pointer(), C.int(base))
	}
}

func (p *qspinbox) SetMaximum(max int) {
	if p.Pointer() != nil {
		C.QSpinBox_SetMaximum_Int(p.Pointer(), C.int(max))
	}
}

func (p *qspinbox) SetMinimum(min int) {
	if p.Pointer() != nil {
		C.QSpinBox_SetMinimum_Int(p.Pointer(), C.int(min))
	}
}

func (p *qspinbox) SetPrefix(prefix string) {
	if p.Pointer() != nil {
		C.QSpinBox_SetPrefix_String(p.Pointer(), C.CString(prefix))
	}
}

func (p *qspinbox) SetRange(minimum int, maximum int) {
	if p.Pointer() != nil {
		C.QSpinBox_SetRange_Int_Int(p.Pointer(), C.int(minimum), C.int(maximum))
	}
}

func (p *qspinbox) SetSingleStep(val int) {
	if p.Pointer() != nil {
		C.QSpinBox_SetSingleStep_Int(p.Pointer(), C.int(val))
	}
}

func (p *qspinbox) SetSuffix(suffix string) {
	if p.Pointer() != nil {
		C.QSpinBox_SetSuffix_String(p.Pointer(), C.CString(suffix))
	}
}

func (p *qspinbox) SingleStep() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QSpinBox_SingleStep(p.Pointer()))
}

func (p *qspinbox) Suffix() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QSpinBox_Suffix(p.Pointer()))
}

func (p *qspinbox) Value() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QSpinBox_Value(p.Pointer()))
}

func (p *qspinbox) ConnectSlotSetValue() {
	C.QSpinBox_ConnectSlotSetValue(p.Pointer())
}

func (p *qspinbox) DisconnectSlotSetValue() {
	C.QSpinBox_DisconnectSlotSetValue(p.Pointer())
}

func (p *qspinbox) SlotSetValue(val int) {
	if p.Pointer() != nil {
		C.QSpinBox_SetValue_Int(p.Pointer(), C.int(val))
	}
}
