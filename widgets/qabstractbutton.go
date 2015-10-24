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

type QAbstractButtonITF interface {
	QWidgetITF
	QAbstractButtonPTR() *QAbstractButton
}

func PointerFromQAbstractButton(ptr QAbstractButtonITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractButtonPTR().Pointer()
	}
	return nil
}

func QAbstractButtonFromPointer(ptr unsafe.Pointer) *QAbstractButton {
	var n = new(QAbstractButton)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAbstractButton_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractButton) QAbstractButtonPTR() *QAbstractButton {
	return ptr
}

func (ptr *QAbstractButton) AutoExclusive() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractButton_AutoExclusive(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractButton) AutoRepeat() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractButton_AutoRepeat(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractButton) AutoRepeatDelay() int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractButton_AutoRepeatDelay(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractButton) AutoRepeatInterval() int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractButton_AutoRepeatInterval(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractButton) IsCheckable() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractButton_IsCheckable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractButton) IsChecked() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractButton_IsChecked(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractButton) IsDown() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractButton_IsDown(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractButton) SetAutoExclusive(v bool) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_SetAutoExclusive(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractButton) SetAutoRepeat(v bool) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_SetAutoRepeat(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractButton) SetAutoRepeatDelay(v int) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_SetAutoRepeatDelay(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QAbstractButton) SetAutoRepeatInterval(v int) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_SetAutoRepeatInterval(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QAbstractButton) SetCheckable(v bool) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_SetCheckable(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractButton) SetChecked(v bool) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_SetChecked(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractButton) SetDown(v bool) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_SetDown(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractButton) SetIcon(icon gui.QIconITF) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_SetIcon(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQIcon(icon)))
	}
}

func (ptr *QAbstractButton) SetIconSize(size core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_SetIconSize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(size)))
	}
}

func (ptr *QAbstractButton) SetShortcut(key gui.QKeySequenceITF) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_SetShortcut(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQKeySequence(key)))
	}
}

func (ptr *QAbstractButton) SetText(text string) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_SetText(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QAbstractButton) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractButton_Text(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QAbstractButton) Toggle() {
	if ptr.Pointer() != nil {
		C.QAbstractButton_Toggle(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAbstractButton) AnimateClick(msec int) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_AnimateClick(C.QtObjectPtr(ptr.Pointer()), C.int(msec))
	}
}

func (ptr *QAbstractButton) Click() {
	if ptr.Pointer() != nil {
		C.QAbstractButton_Click(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAbstractButton) ConnectClicked(f func(checked bool)) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_ConnectClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "clicked", f)
	}
}

func (ptr *QAbstractButton) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QAbstractButton_DisconnectClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "clicked")
	}
}

//export callbackQAbstractButtonClicked
func callbackQAbstractButtonClicked(ptrName *C.char, checked C.int) {
	qt.GetSignal(C.GoString(ptrName), "clicked").(func(bool))(int(checked) != 0)
}

func (ptr *QAbstractButton) Group() *QButtonGroup {
	if ptr.Pointer() != nil {
		return QButtonGroupFromPointer(unsafe.Pointer(C.QAbstractButton_Group(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAbstractButton) ConnectPressed(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_ConnectPressed(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "pressed", f)
	}
}

func (ptr *QAbstractButton) DisconnectPressed() {
	if ptr.Pointer() != nil {
		C.QAbstractButton_DisconnectPressed(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "pressed")
	}
}

//export callbackQAbstractButtonPressed
func callbackQAbstractButtonPressed(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "pressed").(func())()
}

func (ptr *QAbstractButton) ConnectReleased(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_ConnectReleased(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "released", f)
	}
}

func (ptr *QAbstractButton) DisconnectReleased() {
	if ptr.Pointer() != nil {
		C.QAbstractButton_DisconnectReleased(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "released")
	}
}

//export callbackQAbstractButtonReleased
func callbackQAbstractButtonReleased(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "released").(func())()
}

func (ptr *QAbstractButton) ConnectToggled(f func(checked bool)) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_ConnectToggled(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "toggled", f)
	}
}

func (ptr *QAbstractButton) DisconnectToggled() {
	if ptr.Pointer() != nil {
		C.QAbstractButton_DisconnectToggled(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "toggled")
	}
}

//export callbackQAbstractButtonToggled
func callbackQAbstractButtonToggled(ptrName *C.char, checked C.int) {
	qt.GetSignal(C.GoString(ptrName), "toggled").(func(bool))(int(checked) != 0)
}

func (ptr *QAbstractButton) DestroyQAbstractButton() {
	if ptr.Pointer() != nil {
		C.QAbstractButton_DestroyQAbstractButton(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
