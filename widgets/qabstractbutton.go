package widgets

//#include "qabstractbutton.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QAbstractButton struct {
	QWidget
}

type QAbstractButton_ITF interface {
	QWidget_ITF
	QAbstractButton_PTR() *QAbstractButton
}

func PointerFromQAbstractButton(ptr QAbstractButton_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractButton_PTR().Pointer()
	}
	return nil
}

func NewQAbstractButtonFromPointer(ptr unsafe.Pointer) *QAbstractButton {
	var n = new(QAbstractButton)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAbstractButton_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractButton) QAbstractButton_PTR() *QAbstractButton {
	return ptr
}

func (ptr *QAbstractButton) AutoExclusive() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractButton_AutoExclusive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractButton) AutoRepeat() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractButton_AutoRepeat(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractButton) AutoRepeatDelay() int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractButton_AutoRepeatDelay(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractButton) AutoRepeatInterval() int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractButton_AutoRepeatInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractButton) IsCheckable() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractButton_IsCheckable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractButton) IsChecked() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractButton_IsChecked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractButton) IsDown() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractButton_IsDown(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractButton) SetAutoExclusive(v bool) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_SetAutoExclusive(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractButton) SetAutoRepeat(v bool) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_SetAutoRepeat(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractButton) SetAutoRepeatDelay(v int) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_SetAutoRepeatDelay(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractButton) SetAutoRepeatInterval(v int) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_SetAutoRepeatInterval(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractButton) SetCheckable(v bool) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_SetCheckable(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractButton) SetChecked(v bool) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_SetChecked(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractButton) SetDown(v bool) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_SetDown(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractButton) SetIcon(icon gui.QIcon_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_SetIcon(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QAbstractButton) SetIconSize(size core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_SetIconSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QAbstractButton) SetShortcut(key gui.QKeySequence_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_SetShortcut(ptr.Pointer(), gui.PointerFromQKeySequence(key))
	}
}

func (ptr *QAbstractButton) SetText(text string) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_SetText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QAbstractButton) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractButton_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstractButton) Toggle() {
	if ptr.Pointer() != nil {
		C.QAbstractButton_Toggle(ptr.Pointer())
	}
}

func (ptr *QAbstractButton) AnimateClick(msec int) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_AnimateClick(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QAbstractButton) Click() {
	if ptr.Pointer() != nil {
		C.QAbstractButton_Click(ptr.Pointer())
	}
}

func (ptr *QAbstractButton) ConnectClicked(f func(checked bool)) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_ConnectClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "clicked", f)
	}
}

func (ptr *QAbstractButton) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QAbstractButton_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "clicked")
	}
}

//export callbackQAbstractButtonClicked
func callbackQAbstractButtonClicked(ptrName *C.char, checked C.int) {
	qt.GetSignal(C.GoString(ptrName), "clicked").(func(bool))(int(checked) != 0)
}

func (ptr *QAbstractButton) Group() *QButtonGroup {
	if ptr.Pointer() != nil {
		return NewQButtonGroupFromPointer(C.QAbstractButton_Group(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractButton) ConnectPressed(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_ConnectPressed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "pressed", f)
	}
}

func (ptr *QAbstractButton) DisconnectPressed() {
	if ptr.Pointer() != nil {
		C.QAbstractButton_DisconnectPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "pressed")
	}
}

//export callbackQAbstractButtonPressed
func callbackQAbstractButtonPressed(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "pressed").(func())()
}

func (ptr *QAbstractButton) ConnectReleased(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_ConnectReleased(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "released", f)
	}
}

func (ptr *QAbstractButton) DisconnectReleased() {
	if ptr.Pointer() != nil {
		C.QAbstractButton_DisconnectReleased(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "released")
	}
}

//export callbackQAbstractButtonReleased
func callbackQAbstractButtonReleased(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "released").(func())()
}

func (ptr *QAbstractButton) ConnectToggled(f func(checked bool)) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_ConnectToggled(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "toggled", f)
	}
}

func (ptr *QAbstractButton) DisconnectToggled() {
	if ptr.Pointer() != nil {
		C.QAbstractButton_DisconnectToggled(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "toggled")
	}
}

//export callbackQAbstractButtonToggled
func callbackQAbstractButtonToggled(ptrName *C.char, checked C.int) {
	qt.GetSignal(C.GoString(ptrName), "toggled").(func(bool))(int(checked) != 0)
}

func (ptr *QAbstractButton) DestroyQAbstractButton() {
	if ptr.Pointer() != nil {
		C.QAbstractButton_DestroyQAbstractButton(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
