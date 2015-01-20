package qt

//#include "qabstractspinbox.h"
import "C"

type qabstractspinbox struct {
	qwidget
}

type QAbstractSpinBox interface {
	QWidget
	Alignment() AlignmentFlag
	ButtonSymbols() ButtonSymbols
	HasAcceptableInput() bool
	HasFrame() bool
	InterpretText()
	IsAccelerated() bool
	IsGroupSeparatorShown() bool
	IsReadOnly() bool
	KeyboardTracking() bool
	SetAccelerated(on bool)
	SetAlignment(flag AlignmentFlag)
	SetButtonSymbols(bs ButtonSymbols)
	SetFrame(frame bool)
	SetGroupSeparatorShown(shown bool)
	SetKeyboardTracking(kt bool)
	SetReadOnly(r bool)
	SetSpecialValueText(txt string)
	SetWrapping(w bool)
	SpecialValueText() string
	StepBy(steps int)
	Text() string
	Wrapping() bool
	ConnectSlotClear()
	DisconnectSlotClear()
	SlotClear()
	ConnectSlotSelectAll()
	DisconnectSlotSelectAll()
	SlotSelectAll()
	ConnectSlotStepDown()
	DisconnectSlotStepDown()
	SlotStepDown()
	ConnectSlotStepUp()
	DisconnectSlotStepUp()
	SlotStepUp()
	ConnectSignalEditingFinished(f func())
	DisconnectSignalEditingFinished()
	SignalEditingFinished() func()
}

func (p *qabstractspinbox) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qabstractspinbox) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

//ButtonSymbols
type ButtonSymbols int

var (
	UPDOWNARROWS = ButtonSymbols(C.QAbstractSpinBox_UpDownArrows())
	PLUSMINUS    = ButtonSymbols(C.QAbstractSpinBox_PlusMinus())
	NOBUTTONS    = ButtonSymbols(C.QAbstractSpinBox_NoButtons())
)

//CorrectionMode
type CorrectionMode int

var (
	CORRECTTOPREVIOUSVALUE = CorrectionMode(C.QAbstractSpinBox_CorrectToPreviousValue())
	CORRECTTONEARESTVALUE  = CorrectionMode(C.QAbstractSpinBox_CorrectToNearestValue())
)

//StepEnabledFlag
type StepEnabledFlag int

var (
	STEPNONE        = StepEnabledFlag(C.QAbstractSpinBox_StepNone())
	STEPUPENABLED   = StepEnabledFlag(C.QAbstractSpinBox_StepUpEnabled())
	STEPDOWNENABLED = StepEnabledFlag(C.QAbstractSpinBox_StepDownEnabled())
)

func NewQAbstractSpinBox(parent QWidget) QAbstractSpinBox {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qabstractspinbox = new(qabstractspinbox)
	qabstractspinbox.SetPointer(C.QAbstractSpinBox_New_QWidget(parentPtr))
	qabstractspinbox.SetObjectName("QAbstractSpinBox_" + randomIdentifier())
	return qabstractspinbox
}

func (p *qabstractspinbox) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QAbstractSpinBox_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qabstractspinbox) Alignment() AlignmentFlag {
	if p.Pointer() == nil {
		return 0
	}
	return AlignmentFlag(C.QAbstractSpinBox_Alignment(p.Pointer()))
}

func (p *qabstractspinbox) ButtonSymbols() ButtonSymbols {
	if p.Pointer() == nil {
		return 0
	}
	return ButtonSymbols(C.QAbstractSpinBox_ButtonSymbols(p.Pointer()))
}

func (p *qabstractspinbox) HasAcceptableInput() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QAbstractSpinBox_HasAcceptableInput(p.Pointer()) != 0
}

func (p *qabstractspinbox) HasFrame() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QAbstractSpinBox_HasFrame(p.Pointer()) != 0
}

func (p *qabstractspinbox) InterpretText() {
	if p.Pointer() != nil {
		C.QAbstractSpinBox_InterpretText(p.Pointer())
	}
}

func (p *qabstractspinbox) IsAccelerated() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QAbstractSpinBox_IsAccelerated(p.Pointer()) != 0
}

func (p *qabstractspinbox) IsGroupSeparatorShown() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QAbstractSpinBox_IsGroupSeparatorShown(p.Pointer()) != 0
}

func (p *qabstractspinbox) IsReadOnly() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QAbstractSpinBox_IsReadOnly(p.Pointer()) != 0
}

func (p *qabstractspinbox) KeyboardTracking() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QAbstractSpinBox_KeyboardTracking(p.Pointer()) != 0
}

func (p *qabstractspinbox) SetAccelerated(on bool) {
	if p.Pointer() != nil {
		C.QAbstractSpinBox_SetAccelerated_Bool(p.Pointer(), goBoolToCInt(on))
	}
}

func (p *qabstractspinbox) SetAlignment(flag AlignmentFlag) {
	if p.Pointer() != nil {
		C.QAbstractSpinBox_SetAlignment_AlignmentFlag(p.Pointer(), C.int(flag))
	}
}

func (p *qabstractspinbox) SetButtonSymbols(bs ButtonSymbols) {
	if p.Pointer() != nil {
		C.QAbstractSpinBox_SetButtonSymbols_ButtonSymbols(p.Pointer(), C.int(bs))
	}
}

func (p *qabstractspinbox) SetFrame(frame bool) {
	if p.Pointer() != nil {
		C.QAbstractSpinBox_SetFrame_Bool(p.Pointer(), goBoolToCInt(frame))
	}
}

func (p *qabstractspinbox) SetGroupSeparatorShown(shown bool) {
	if p.Pointer() != nil {
		C.QAbstractSpinBox_SetGroupSeparatorShown_Bool(p.Pointer(), goBoolToCInt(shown))
	}
}

func (p *qabstractspinbox) SetKeyboardTracking(kt bool) {
	if p.Pointer() != nil {
		C.QAbstractSpinBox_SetKeyboardTracking_Bool(p.Pointer(), goBoolToCInt(kt))
	}
}

func (p *qabstractspinbox) SetReadOnly(r bool) {
	if p.Pointer() != nil {
		C.QAbstractSpinBox_SetReadOnly_Bool(p.Pointer(), goBoolToCInt(r))
	}
}

func (p *qabstractspinbox) SetSpecialValueText(txt string) {
	if p.Pointer() != nil {
		C.QAbstractSpinBox_SetSpecialValueText_String(p.Pointer(), C.CString(txt))
	}
}

func (p *qabstractspinbox) SetWrapping(w bool) {
	if p.Pointer() != nil {
		C.QAbstractSpinBox_SetWrapping_Bool(p.Pointer(), goBoolToCInt(w))
	}
}

func (p *qabstractspinbox) SpecialValueText() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QAbstractSpinBox_SpecialValueText(p.Pointer()))
}

func (p *qabstractspinbox) StepBy(steps int) {
	if p.Pointer() != nil {
		C.QAbstractSpinBox_StepBy_Int(p.Pointer(), C.int(steps))
	}
}

func (p *qabstractspinbox) Text() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QAbstractSpinBox_Text(p.Pointer()))
}

func (p *qabstractspinbox) Wrapping() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QAbstractSpinBox_Wrapping(p.Pointer()) != 0
}

func (p *qabstractspinbox) ConnectSlotClear() {
	C.QAbstractSpinBox_ConnectSlotClear(p.Pointer())
}

func (p *qabstractspinbox) DisconnectSlotClear() {
	C.QAbstractSpinBox_DisconnectSlotClear(p.Pointer())
}

func (p *qabstractspinbox) SlotClear() {
	if p.Pointer() != nil {
		C.QAbstractSpinBox_Clear(p.Pointer())
	}
}

func (p *qabstractspinbox) ConnectSlotSelectAll() {
	C.QAbstractSpinBox_ConnectSlotSelectAll(p.Pointer())
}

func (p *qabstractspinbox) DisconnectSlotSelectAll() {
	C.QAbstractSpinBox_DisconnectSlotSelectAll(p.Pointer())
}

func (p *qabstractspinbox) SlotSelectAll() {
	if p.Pointer() != nil {
		C.QAbstractSpinBox_SelectAll(p.Pointer())
	}
}

func (p *qabstractspinbox) ConnectSlotStepDown() {
	C.QAbstractSpinBox_ConnectSlotStepDown(p.Pointer())
}

func (p *qabstractspinbox) DisconnectSlotStepDown() {
	C.QAbstractSpinBox_DisconnectSlotStepDown(p.Pointer())
}

func (p *qabstractspinbox) SlotStepDown() {
	if p.Pointer() != nil {
		C.QAbstractSpinBox_StepDown(p.Pointer())
	}
}

func (p *qabstractspinbox) ConnectSlotStepUp() {
	C.QAbstractSpinBox_ConnectSlotStepUp(p.Pointer())
}

func (p *qabstractspinbox) DisconnectSlotStepUp() {
	C.QAbstractSpinBox_DisconnectSlotStepUp(p.Pointer())
}

func (p *qabstractspinbox) SlotStepUp() {
	if p.Pointer() != nil {
		C.QAbstractSpinBox_StepUp(p.Pointer())
	}
}

func (p *qabstractspinbox) ConnectSignalEditingFinished(f func()) {
	C.QAbstractSpinBox_ConnectSignalEditingFinished(p.Pointer())
	connectSignal(p.ObjectName(), "editingFinished", f)
}

func (p *qabstractspinbox) DisconnectSignalEditingFinished() {
	C.QAbstractSpinBox_DisconnectSignalEditingFinished(p.Pointer())
	disconnectSignal(p.ObjectName(), "editingFinished")
}

func (p *qabstractspinbox) SignalEditingFinished() func() {
	return getSignal(p.ObjectName(), "editingFinished")
}

//export callbackQAbstractSpinBox
func callbackQAbstractSpinBox(ptr C.QtObjectPtr, msg *C.char) {
	var qabstractspinbox = new(qabstractspinbox)
	qabstractspinbox.SetPointer(ptr)
	getSignal(qabstractspinbox.ObjectName(), C.GoString(msg))()
}
