package qt

//#include "qabstractbutton.h"
import "C"

type qabstractbutton struct {
	qwidget
}

type QAbstractButton interface {
	QWidget
	AutoExclusive() bool
	AutoRepeat() bool
	AutoRepeatDelay() int
	AutoRepeatInterval() int
	IsCheckable() bool
	IsChecked() bool
	IsDown() bool
	SetAutoExclusive(autoExclusive bool)
	SetAutoRepeat(autoRepeat bool)
	SetAutoRepeatDelay(autoRepeatDelay int)
	SetAutoRepeatInterval(autoRepeatInterval int)
	SetCheckable(checkable bool)
	SetDown(down bool)
	SetText(text string)
	Text() string
	ConnectSlotAnimateClick()
	DisconnectSlotAnimateClick()
	SlotAnimateClick(msec int)
	ConnectSlotClick()
	DisconnectSlotClick()
	SlotClick()
	ConnectSlotSetChecked()
	DisconnectSlotSetChecked()
	SlotSetChecked(checked bool)
	ConnectSlotToggle()
	DisconnectSlotToggle()
	SlotToggle()
	ConnectSignalClicked(f func())
	DisconnectSignalClicked()
	SignalClicked() func()
	ConnectSignalPressed(f func())
	DisconnectSignalPressed()
	SignalPressed() func()
	ConnectSignalReleased(f func())
	DisconnectSignalReleased()
	SignalReleased() func()
	ConnectSignalToggled(f func())
	DisconnectSignalToggled()
	SignalToggled() func()
}

func (p *qabstractbutton) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qabstractbutton) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func (p *qabstractbutton) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QAbstractButton_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qabstractbutton) AutoExclusive() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QAbstractButton_AutoExclusive(p.Pointer()) != 0
}

func (p *qabstractbutton) AutoRepeat() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QAbstractButton_AutoRepeat(p.Pointer()) != 0
}

func (p *qabstractbutton) AutoRepeatDelay() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QAbstractButton_AutoRepeatDelay(p.Pointer()))
}

func (p *qabstractbutton) AutoRepeatInterval() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QAbstractButton_AutoRepeatInterval(p.Pointer()))
}

func (p *qabstractbutton) IsCheckable() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QAbstractButton_IsCheckable(p.Pointer()) != 0
}

func (p *qabstractbutton) IsChecked() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QAbstractButton_IsChecked(p.Pointer()) != 0
}

func (p *qabstractbutton) IsDown() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QAbstractButton_IsDown(p.Pointer()) != 0
}

func (p *qabstractbutton) SetAutoExclusive(autoExclusive bool) {
	if p.Pointer() != nil {
		C.QAbstractButton_SetAutoExclusive_Bool(p.Pointer(), goBoolToCInt(autoExclusive))
	}
}

func (p *qabstractbutton) SetAutoRepeat(autoRepeat bool) {
	if p.Pointer() != nil {
		C.QAbstractButton_SetAutoRepeat_Bool(p.Pointer(), goBoolToCInt(autoRepeat))
	}
}

func (p *qabstractbutton) SetAutoRepeatDelay(autoRepeatDelay int) {
	if p.Pointer() != nil {
		C.QAbstractButton_SetAutoRepeatDelay_Int(p.Pointer(), C.int(autoRepeatDelay))
	}
}

func (p *qabstractbutton) SetAutoRepeatInterval(autoRepeatInterval int) {
	if p.Pointer() != nil {
		C.QAbstractButton_SetAutoRepeatInterval_Int(p.Pointer(), C.int(autoRepeatInterval))
	}
}

func (p *qabstractbutton) SetCheckable(checkable bool) {
	if p.Pointer() != nil {
		C.QAbstractButton_SetCheckable_Bool(p.Pointer(), goBoolToCInt(checkable))
	}
}

func (p *qabstractbutton) SetDown(down bool) {
	if p.Pointer() != nil {
		C.QAbstractButton_SetDown_Bool(p.Pointer(), goBoolToCInt(down))
	}
}

func (p *qabstractbutton) SetText(text string) {
	if p.Pointer() != nil {
		C.QAbstractButton_SetText_String(p.Pointer(), C.CString(text))
	}
}

func (p *qabstractbutton) Text() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QAbstractButton_Text(p.Pointer()))
}

func (p *qabstractbutton) ConnectSlotAnimateClick() {
	C.QAbstractButton_ConnectSlotAnimateClick(p.Pointer())
}

func (p *qabstractbutton) DisconnectSlotAnimateClick() {
	C.QAbstractButton_DisconnectSlotAnimateClick(p.Pointer())
}

func (p *qabstractbutton) SlotAnimateClick(msec int) {
	if p.Pointer() != nil {
		C.QAbstractButton_AnimateClick_Int(p.Pointer(), C.int(msec))
	}
}

func (p *qabstractbutton) ConnectSlotClick() {
	C.QAbstractButton_ConnectSlotClick(p.Pointer())
}

func (p *qabstractbutton) DisconnectSlotClick() {
	C.QAbstractButton_DisconnectSlotClick(p.Pointer())
}

func (p *qabstractbutton) SlotClick() {
	if p.Pointer() != nil {
		C.QAbstractButton_Click(p.Pointer())
	}
}

func (p *qabstractbutton) ConnectSlotSetChecked() {
	C.QAbstractButton_ConnectSlotSetChecked(p.Pointer())
}

func (p *qabstractbutton) DisconnectSlotSetChecked() {
	C.QAbstractButton_DisconnectSlotSetChecked(p.Pointer())
}

func (p *qabstractbutton) SlotSetChecked(checked bool) {
	if p.Pointer() != nil {
		C.QAbstractButton_SetChecked_Bool(p.Pointer(), goBoolToCInt(checked))
	}
}

func (p *qabstractbutton) ConnectSlotToggle() {
	C.QAbstractButton_ConnectSlotToggle(p.Pointer())
}

func (p *qabstractbutton) DisconnectSlotToggle() {
	C.QAbstractButton_DisconnectSlotToggle(p.Pointer())
}

func (p *qabstractbutton) SlotToggle() {
	if p.Pointer() != nil {
		C.QAbstractButton_Toggle(p.Pointer())
	}
}

func (p *qabstractbutton) ConnectSignalClicked(f func()) {
	C.QAbstractButton_ConnectSignalClicked(p.Pointer())
	connectSignal(p.ObjectName(), "clicked", f)
}

func (p *qabstractbutton) DisconnectSignalClicked() {
	C.QAbstractButton_DisconnectSignalClicked(p.Pointer())
	disconnectSignal(p.ObjectName(), "clicked")
}

func (p *qabstractbutton) SignalClicked() func() {
	return getSignal(p.ObjectName(), "clicked")
}

func (p *qabstractbutton) ConnectSignalPressed(f func()) {
	C.QAbstractButton_ConnectSignalPressed(p.Pointer())
	connectSignal(p.ObjectName(), "pressed", f)
}

func (p *qabstractbutton) DisconnectSignalPressed() {
	C.QAbstractButton_DisconnectSignalPressed(p.Pointer())
	disconnectSignal(p.ObjectName(), "pressed")
}

func (p *qabstractbutton) SignalPressed() func() {
	return getSignal(p.ObjectName(), "pressed")
}

func (p *qabstractbutton) ConnectSignalReleased(f func()) {
	C.QAbstractButton_ConnectSignalReleased(p.Pointer())
	connectSignal(p.ObjectName(), "released", f)
}

func (p *qabstractbutton) DisconnectSignalReleased() {
	C.QAbstractButton_DisconnectSignalReleased(p.Pointer())
	disconnectSignal(p.ObjectName(), "released")
}

func (p *qabstractbutton) SignalReleased() func() {
	return getSignal(p.ObjectName(), "released")
}

func (p *qabstractbutton) ConnectSignalToggled(f func()) {
	C.QAbstractButton_ConnectSignalToggled(p.Pointer())
	connectSignal(p.ObjectName(), "toggled", f)
}

func (p *qabstractbutton) DisconnectSignalToggled() {
	C.QAbstractButton_DisconnectSignalToggled(p.Pointer())
	disconnectSignal(p.ObjectName(), "toggled")
}

func (p *qabstractbutton) SignalToggled() func() {
	return getSignal(p.ObjectName(), "toggled")
}

//export callbackQAbstractButton
func callbackQAbstractButton(ptr C.QtObjectPtr, msg *C.char) {
	var qabstractbutton = new(qabstractbutton)
	qabstractbutton.SetPointer(ptr)
	getSignal(qabstractbutton.ObjectName(), C.GoString(msg))()
}
