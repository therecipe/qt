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
	SetDisplayIntegerBase_Int(base int)
	SetMaximum_Int(max int)
	SetMinimum_Int(min int)
	SetPrefix_String(prefix string)
	SetRange_Int_Int(minimum int, maximum int)
	SetSingleStep_Int(val int)
	SetSuffix_String(suffix string)
	SingleStep() int
	Suffix() string
	Value() int
	ConnectSlotSetValue()
	DisconnectSlotSetValue()
	SlotSetValue_Int(val int)
}

func (p *qspinbox) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qspinbox) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQSpinBox_QWidget(parent QWidget) QSpinBox {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qspinbox = new(qspinbox)
	qspinbox.SetPointer(C.QSpinBox_New_QWidget(parentPtr))
	qspinbox.SetObjectName_String("QSpinBox_" + randomIdentifier())
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
	} else {
		return C.GoString(C.QSpinBox_CleanText(p.Pointer()))
	}
}

func (p *qspinbox) DisplayIntegerBase() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QSpinBox_DisplayIntegerBase(p.Pointer()))
	}
}

func (p *qspinbox) Maximum() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QSpinBox_Maximum(p.Pointer()))
	}
}

func (p *qspinbox) Minimum() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QSpinBox_Minimum(p.Pointer()))
	}
}

func (p *qspinbox) Prefix() string {
	if p.Pointer() == nil {
		return ""
	} else {
		return C.GoString(C.QSpinBox_Prefix(p.Pointer()))
	}
}

func (p *qspinbox) SetDisplayIntegerBase_Int(base int) {
	if p.Pointer() != nil {
		C.QSpinBox_SetDisplayIntegerBase_Int(p.Pointer(), C.int(base))
	}
}

func (p *qspinbox) SetMaximum_Int(max int) {
	if p.Pointer() != nil {
		C.QSpinBox_SetMaximum_Int(p.Pointer(), C.int(max))
	}
}

func (p *qspinbox) SetMinimum_Int(min int) {
	if p.Pointer() != nil {
		C.QSpinBox_SetMinimum_Int(p.Pointer(), C.int(min))
	}
}

func (p *qspinbox) SetPrefix_String(prefix string) {
	if p.Pointer() != nil {
		C.QSpinBox_SetPrefix_String(p.Pointer(), C.CString(prefix))
	}
}

func (p *qspinbox) SetRange_Int_Int(minimum int, maximum int) {
	if p.Pointer() != nil {
		C.QSpinBox_SetRange_Int_Int(p.Pointer(), C.int(minimum), C.int(maximum))
	}
}

func (p *qspinbox) SetSingleStep_Int(val int) {
	if p.Pointer() != nil {
		C.QSpinBox_SetSingleStep_Int(p.Pointer(), C.int(val))
	}
}

func (p *qspinbox) SetSuffix_String(suffix string) {
	if p.Pointer() != nil {
		C.QSpinBox_SetSuffix_String(p.Pointer(), C.CString(suffix))
	}
}

func (p *qspinbox) SingleStep() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QSpinBox_SingleStep(p.Pointer()))
	}
}

func (p *qspinbox) Suffix() string {
	if p.Pointer() == nil {
		return ""
	} else {
		return C.GoString(C.QSpinBox_Suffix(p.Pointer()))
	}
}

func (p *qspinbox) Value() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QSpinBox_Value(p.Pointer()))
	}
}

func (p *qspinbox) ConnectSlotSetValue() {
	C.QSpinBox_ConnectSlotSetValue(p.Pointer())
}

func (p *qspinbox) DisconnectSlotSetValue() {
	C.QSpinBox_DisconnectSlotSetValue(p.Pointer())
}

func (p *qspinbox) SlotSetValue_Int(val int) {
	if p.Pointer() != nil {
		C.QSpinBox_SetValue_Int(p.Pointer(), C.int(val))
	}
}
