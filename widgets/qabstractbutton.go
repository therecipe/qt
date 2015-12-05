package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"log"
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
	for len(n.ObjectName()) < len("QAbstractButton_") {
		n.SetObjectName("QAbstractButton_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractButton) QAbstractButton_PTR() *QAbstractButton {
	return ptr
}

func (ptr *QAbstractButton) AutoExclusive() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::autoExclusive")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractButton_AutoExclusive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractButton) AutoRepeat() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::autoRepeat")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractButton_AutoRepeat(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractButton) AutoRepeatDelay() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::autoRepeatDelay")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAbstractButton_AutoRepeatDelay(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractButton) AutoRepeatInterval() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::autoRepeatInterval")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAbstractButton_AutoRepeatInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractButton) IsCheckable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::isCheckable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractButton_IsCheckable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractButton) IsChecked() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::isChecked")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractButton_IsChecked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractButton) IsDown() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::isDown")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractButton_IsDown(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractButton) SetAutoExclusive(v bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::setAutoExclusive")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractButton_SetAutoExclusive(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractButton) SetAutoRepeat(v bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::setAutoRepeat")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractButton_SetAutoRepeat(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractButton) SetAutoRepeatDelay(v int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::setAutoRepeatDelay")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractButton_SetAutoRepeatDelay(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractButton) SetAutoRepeatInterval(v int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::setAutoRepeatInterval")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractButton_SetAutoRepeatInterval(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractButton) SetCheckable(v bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::setCheckable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractButton_SetCheckable(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractButton) SetChecked(v bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::setChecked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractButton_SetChecked(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractButton) SetDown(v bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::setDown")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractButton_SetDown(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractButton) SetIcon(icon gui.QIcon_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::setIcon")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractButton_SetIcon(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QAbstractButton) SetIconSize(size core.QSize_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::setIconSize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractButton_SetIconSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QAbstractButton) SetShortcut(key gui.QKeySequence_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::setShortcut")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractButton_SetShortcut(ptr.Pointer(), gui.PointerFromQKeySequence(key))
	}
}

func (ptr *QAbstractButton) SetText(text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::setText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractButton_SetText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QAbstractButton) Text() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::text")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractButton_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstractButton) Toggle() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::toggle")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractButton_Toggle(ptr.Pointer())
	}
}

func (ptr *QAbstractButton) AnimateClick(msec int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::animateClick")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractButton_AnimateClick(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QAbstractButton) Click() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::click")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractButton_Click(ptr.Pointer())
	}
}

func (ptr *QAbstractButton) ConnectClicked(f func(checked bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::clicked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractButton_ConnectClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "clicked", f)
	}
}

func (ptr *QAbstractButton) DisconnectClicked() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::clicked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractButton_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "clicked")
	}
}

//export callbackQAbstractButtonClicked
func callbackQAbstractButtonClicked(ptrName *C.char, checked C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::clicked")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "clicked").(func(bool))(int(checked) != 0)
}

func (ptr *QAbstractButton) Group() *QButtonGroup {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::group")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQButtonGroupFromPointer(C.QAbstractButton_Group(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractButton) ConnectPressed(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::pressed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractButton_ConnectPressed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "pressed", f)
	}
}

func (ptr *QAbstractButton) DisconnectPressed() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::pressed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractButton_DisconnectPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "pressed")
	}
}

//export callbackQAbstractButtonPressed
func callbackQAbstractButtonPressed(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::pressed")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "pressed").(func())()
}

func (ptr *QAbstractButton) ConnectReleased(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::released")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractButton_ConnectReleased(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "released", f)
	}
}

func (ptr *QAbstractButton) DisconnectReleased() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::released")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractButton_DisconnectReleased(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "released")
	}
}

//export callbackQAbstractButtonReleased
func callbackQAbstractButtonReleased(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::released")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "released").(func())()
}

func (ptr *QAbstractButton) ConnectToggled(f func(checked bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::toggled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractButton_ConnectToggled(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "toggled", f)
	}
}

func (ptr *QAbstractButton) DisconnectToggled() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::toggled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractButton_DisconnectToggled(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "toggled")
	}
}

//export callbackQAbstractButtonToggled
func callbackQAbstractButtonToggled(ptrName *C.char, checked C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::toggled")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "toggled").(func(bool))(int(checked) != 0)
}

func (ptr *QAbstractButton) DestroyQAbstractButton() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractButton::~QAbstractButton")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractButton_DestroyQAbstractButton(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
