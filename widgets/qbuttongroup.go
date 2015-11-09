package widgets

//#include "qbuttongroup.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QButtonGroup struct {
	core.QObject
}

type QButtonGroup_ITF interface {
	core.QObject_ITF
	QButtonGroup_PTR() *QButtonGroup
}

func PointerFromQButtonGroup(ptr QButtonGroup_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QButtonGroup_PTR().Pointer()
	}
	return nil
}

func NewQButtonGroupFromPointer(ptr unsafe.Pointer) *QButtonGroup {
	var n = new(QButtonGroup)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QButtonGroup_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QButtonGroup) QButtonGroup_PTR() *QButtonGroup {
	return ptr
}

func NewQButtonGroup(parent core.QObject_ITF) *QButtonGroup {
	return NewQButtonGroupFromPointer(C.QButtonGroup_NewQButtonGroup(core.PointerFromQObject(parent)))
}

func (ptr *QButtonGroup) AddButton(button QAbstractButton_ITF, id int) {
	if ptr.Pointer() != nil {
		C.QButtonGroup_AddButton(ptr.Pointer(), PointerFromQAbstractButton(button), C.int(id))
	}
}

func (ptr *QButtonGroup) Button(id int) *QAbstractButton {
	if ptr.Pointer() != nil {
		return NewQAbstractButtonFromPointer(C.QButtonGroup_Button(ptr.Pointer(), C.int(id)))
	}
	return nil
}

func (ptr *QButtonGroup) CheckedButton() *QAbstractButton {
	if ptr.Pointer() != nil {
		return NewQAbstractButtonFromPointer(C.QButtonGroup_CheckedButton(ptr.Pointer()))
	}
	return nil
}

func (ptr *QButtonGroup) CheckedId() int {
	if ptr.Pointer() != nil {
		return int(C.QButtonGroup_CheckedId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QButtonGroup) Exclusive() bool {
	if ptr.Pointer() != nil {
		return C.QButtonGroup_Exclusive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QButtonGroup) Id(button QAbstractButton_ITF) int {
	if ptr.Pointer() != nil {
		return int(C.QButtonGroup_Id(ptr.Pointer(), PointerFromQAbstractButton(button)))
	}
	return 0
}

func (ptr *QButtonGroup) RemoveButton(button QAbstractButton_ITF) {
	if ptr.Pointer() != nil {
		C.QButtonGroup_RemoveButton(ptr.Pointer(), PointerFromQAbstractButton(button))
	}
}

func (ptr *QButtonGroup) SetExclusive(v bool) {
	if ptr.Pointer() != nil {
		C.QButtonGroup_SetExclusive(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QButtonGroup) SetId(button QAbstractButton_ITF, id int) {
	if ptr.Pointer() != nil {
		C.QButtonGroup_SetId(ptr.Pointer(), PointerFromQAbstractButton(button), C.int(id))
	}
}

func (ptr *QButtonGroup) DestroyQButtonGroup() {
	if ptr.Pointer() != nil {
		C.QButtonGroup_DestroyQButtonGroup(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QButtonGroup) ConnectButtonClicked(f func(button *QAbstractButton)) {
	if ptr.Pointer() != nil {
		C.QButtonGroup_ConnectButtonClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "buttonClicked", f)
	}
}

func (ptr *QButtonGroup) DisconnectButtonClicked() {
	if ptr.Pointer() != nil {
		C.QButtonGroup_DisconnectButtonClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "buttonClicked")
	}
}

//export callbackQButtonGroupButtonClicked
func callbackQButtonGroupButtonClicked(ptrName *C.char, button unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "buttonClicked").(func(*QAbstractButton))(NewQAbstractButtonFromPointer(button))
}

func (ptr *QButtonGroup) ConnectButtonPressed(f func(button *QAbstractButton)) {
	if ptr.Pointer() != nil {
		C.QButtonGroup_ConnectButtonPressed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "buttonPressed", f)
	}
}

func (ptr *QButtonGroup) DisconnectButtonPressed() {
	if ptr.Pointer() != nil {
		C.QButtonGroup_DisconnectButtonPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "buttonPressed")
	}
}

//export callbackQButtonGroupButtonPressed
func callbackQButtonGroupButtonPressed(ptrName *C.char, button unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "buttonPressed").(func(*QAbstractButton))(NewQAbstractButtonFromPointer(button))
}

func (ptr *QButtonGroup) ConnectButtonReleased(f func(button *QAbstractButton)) {
	if ptr.Pointer() != nil {
		C.QButtonGroup_ConnectButtonReleased(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "buttonReleased", f)
	}
}

func (ptr *QButtonGroup) DisconnectButtonReleased() {
	if ptr.Pointer() != nil {
		C.QButtonGroup_DisconnectButtonReleased(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "buttonReleased")
	}
}

//export callbackQButtonGroupButtonReleased
func callbackQButtonGroupButtonReleased(ptrName *C.char, button unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "buttonReleased").(func(*QAbstractButton))(NewQAbstractButtonFromPointer(button))
}

func (ptr *QButtonGroup) ConnectButtonToggled(f func(button *QAbstractButton, checked bool)) {
	if ptr.Pointer() != nil {
		C.QButtonGroup_ConnectButtonToggled(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "buttonToggled", f)
	}
}

func (ptr *QButtonGroup) DisconnectButtonToggled() {
	if ptr.Pointer() != nil {
		C.QButtonGroup_DisconnectButtonToggled(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "buttonToggled")
	}
}

//export callbackQButtonGroupButtonToggled
func callbackQButtonGroupButtonToggled(ptrName *C.char, button unsafe.Pointer, checked C.int) {
	qt.GetSignal(C.GoString(ptrName), "buttonToggled").(func(*QAbstractButton, bool))(NewQAbstractButtonFromPointer(button), int(checked) != 0)
}
