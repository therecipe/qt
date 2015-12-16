package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QPushButton struct {
	QAbstractButton
}

type QPushButton_ITF interface {
	QAbstractButton_ITF
	QPushButton_PTR() *QPushButton
}

func PointerFromQPushButton(ptr QPushButton_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPushButton_PTR().Pointer()
	}
	return nil
}

func NewQPushButtonFromPointer(ptr unsafe.Pointer) *QPushButton {
	var n = new(QPushButton)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QPushButton_") {
		n.SetObjectName("QPushButton_" + qt.Identifier())
	}
	return n
}

func (ptr *QPushButton) QPushButton_PTR() *QPushButton {
	return ptr
}

func (ptr *QPushButton) AutoDefault() bool {
	defer qt.Recovering("QPushButton::autoDefault")

	if ptr.Pointer() != nil {
		return C.QPushButton_AutoDefault(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPushButton) IsDefault() bool {
	defer qt.Recovering("QPushButton::isDefault")

	if ptr.Pointer() != nil {
		return C.QPushButton_IsDefault(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPushButton) IsFlat() bool {
	defer qt.Recovering("QPushButton::isFlat")

	if ptr.Pointer() != nil {
		return C.QPushButton_IsFlat(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPushButton) SetAutoDefault(v bool) {
	defer qt.Recovering("QPushButton::setAutoDefault")

	if ptr.Pointer() != nil {
		C.QPushButton_SetAutoDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QPushButton) SetDefault(v bool) {
	defer qt.Recovering("QPushButton::setDefault")

	if ptr.Pointer() != nil {
		C.QPushButton_SetDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QPushButton) SetFlat(v bool) {
	defer qt.Recovering("QPushButton::setFlat")

	if ptr.Pointer() != nil {
		C.QPushButton_SetFlat(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func NewQPushButton(parent QWidget_ITF) *QPushButton {
	defer qt.Recovering("QPushButton::QPushButton")

	return NewQPushButtonFromPointer(C.QPushButton_NewQPushButton(PointerFromQWidget(parent)))
}

func NewQPushButton3(icon gui.QIcon_ITF, text string, parent QWidget_ITF) *QPushButton {
	defer qt.Recovering("QPushButton::QPushButton")

	return NewQPushButtonFromPointer(C.QPushButton_NewQPushButton3(gui.PointerFromQIcon(icon), C.CString(text), PointerFromQWidget(parent)))
}

func NewQPushButton2(text string, parent QWidget_ITF) *QPushButton {
	defer qt.Recovering("QPushButton::QPushButton")

	return NewQPushButtonFromPointer(C.QPushButton_NewQPushButton2(C.CString(text), PointerFromQWidget(parent)))
}

func (ptr *QPushButton) ConnectFocusInEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QPushButton::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QPushButton) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QPushButton::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQPushButtonFocusInEvent
func callbackQPushButtonFocusInEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::focusInEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "focusInEvent")
	if signal != nil {
		defer signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectFocusOutEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QPushButton::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QPushButton) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QPushButton::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQPushButtonFocusOutEvent
func callbackQPushButtonFocusOutEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::focusOutEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "focusOutEvent")
	if signal != nil {
		defer signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QPushButton::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QPushButton) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QPushButton::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQPushButtonKeyPressEvent
func callbackQPushButtonKeyPressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::keyPressEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "keyPressEvent")
	if signal != nil {
		defer signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QPushButton) Menu() *QMenu {
	defer qt.Recovering("QPushButton::menu")

	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QPushButton_Menu(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPushButton) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QPushButton::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QPushButton_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPushButton) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QPushButton::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QPushButton) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QPushButton::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQPushButtonPaintEvent
func callbackQPushButtonPaintEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::paintEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "paintEvent")
	if signal != nil {
		defer signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QPushButton) SetMenu(menu QMenu_ITF) {
	defer qt.Recovering("QPushButton::setMenu")

	if ptr.Pointer() != nil {
		C.QPushButton_SetMenu(ptr.Pointer(), PointerFromQMenu(menu))
	}
}

func (ptr *QPushButton) ShowMenu() {
	defer qt.Recovering("QPushButton::showMenu")

	if ptr.Pointer() != nil {
		C.QPushButton_ShowMenu(ptr.Pointer())
	}
}

func (ptr *QPushButton) SizeHint() *core.QSize {
	defer qt.Recovering("QPushButton::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QPushButton_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPushButton) DestroyQPushButton() {
	defer qt.Recovering("QPushButton::~QPushButton")

	if ptr.Pointer() != nil {
		C.QPushButton_DestroyQPushButton(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
