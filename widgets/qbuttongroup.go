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

type QButtonGroupITF interface {
	core.QObjectITF
	QButtonGroupPTR() *QButtonGroup
}

func PointerFromQButtonGroup(ptr QButtonGroupITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QButtonGroupPTR().Pointer()
	}
	return nil
}

func QButtonGroupFromPointer(ptr unsafe.Pointer) *QButtonGroup {
	var n = new(QButtonGroup)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QButtonGroup_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QButtonGroup) QButtonGroupPTR() *QButtonGroup {
	return ptr
}

func NewQButtonGroup(parent core.QObjectITF) *QButtonGroup {
	return QButtonGroupFromPointer(unsafe.Pointer(C.QButtonGroup_NewQButtonGroup(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QButtonGroup) AddButton(button QAbstractButtonITF, id int) {
	if ptr.Pointer() != nil {
		C.QButtonGroup_AddButton(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractButton(button)), C.int(id))
	}
}

func (ptr *QButtonGroup) Button(id int) *QAbstractButton {
	if ptr.Pointer() != nil {
		return QAbstractButtonFromPointer(unsafe.Pointer(C.QButtonGroup_Button(C.QtObjectPtr(ptr.Pointer()), C.int(id))))
	}
	return nil
}

func (ptr *QButtonGroup) CheckedButton() *QAbstractButton {
	if ptr.Pointer() != nil {
		return QAbstractButtonFromPointer(unsafe.Pointer(C.QButtonGroup_CheckedButton(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QButtonGroup) CheckedId() int {
	if ptr.Pointer() != nil {
		return int(C.QButtonGroup_CheckedId(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QButtonGroup) Exclusive() bool {
	if ptr.Pointer() != nil {
		return C.QButtonGroup_Exclusive(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QButtonGroup) Id(button QAbstractButtonITF) int {
	if ptr.Pointer() != nil {
		return int(C.QButtonGroup_Id(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractButton(button))))
	}
	return 0
}

func (ptr *QButtonGroup) RemoveButton(button QAbstractButtonITF) {
	if ptr.Pointer() != nil {
		C.QButtonGroup_RemoveButton(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractButton(button)))
	}
}

func (ptr *QButtonGroup) SetExclusive(v bool) {
	if ptr.Pointer() != nil {
		C.QButtonGroup_SetExclusive(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QButtonGroup) SetId(button QAbstractButtonITF, id int) {
	if ptr.Pointer() != nil {
		C.QButtonGroup_SetId(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractButton(button)), C.int(id))
	}
}

func (ptr *QButtonGroup) DestroyQButtonGroup() {
	if ptr.Pointer() != nil {
		C.QButtonGroup_DestroyQButtonGroup(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QButtonGroup) ConnectButtonClicked(f func(button QAbstractButtonITF)) {
	if ptr.Pointer() != nil {
		C.QButtonGroup_ConnectButtonClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "buttonClicked", f)
	}
}

func (ptr *QButtonGroup) DisconnectButtonClicked() {
	if ptr.Pointer() != nil {
		C.QButtonGroup_DisconnectButtonClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "buttonClicked")
	}
}

//export callbackQButtonGroupButtonClicked
func callbackQButtonGroupButtonClicked(ptrName *C.char, button unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "buttonClicked").(func(*QAbstractButton))(QAbstractButtonFromPointer(button))
}

func (ptr *QButtonGroup) ConnectButtonPressed(f func(button QAbstractButtonITF)) {
	if ptr.Pointer() != nil {
		C.QButtonGroup_ConnectButtonPressed(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "buttonPressed", f)
	}
}

func (ptr *QButtonGroup) DisconnectButtonPressed() {
	if ptr.Pointer() != nil {
		C.QButtonGroup_DisconnectButtonPressed(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "buttonPressed")
	}
}

//export callbackQButtonGroupButtonPressed
func callbackQButtonGroupButtonPressed(ptrName *C.char, button unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "buttonPressed").(func(*QAbstractButton))(QAbstractButtonFromPointer(button))
}

func (ptr *QButtonGroup) ConnectButtonReleased(f func(button QAbstractButtonITF)) {
	if ptr.Pointer() != nil {
		C.QButtonGroup_ConnectButtonReleased(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "buttonReleased", f)
	}
}

func (ptr *QButtonGroup) DisconnectButtonReleased() {
	if ptr.Pointer() != nil {
		C.QButtonGroup_DisconnectButtonReleased(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "buttonReleased")
	}
}

//export callbackQButtonGroupButtonReleased
func callbackQButtonGroupButtonReleased(ptrName *C.char, button unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "buttonReleased").(func(*QAbstractButton))(QAbstractButtonFromPointer(button))
}

func (ptr *QButtonGroup) ConnectButtonToggled(f func(button QAbstractButtonITF, checked bool)) {
	if ptr.Pointer() != nil {
		C.QButtonGroup_ConnectButtonToggled(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "buttonToggled", f)
	}
}

func (ptr *QButtonGroup) DisconnectButtonToggled() {
	if ptr.Pointer() != nil {
		C.QButtonGroup_DisconnectButtonToggled(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "buttonToggled")
	}
}

//export callbackQButtonGroupButtonToggled
func callbackQButtonGroupButtonToggled(ptrName *C.char, button unsafe.Pointer, checked C.int) {
	qt.GetSignal(C.GoString(ptrName), "buttonToggled").(func(*QAbstractButton, bool))(QAbstractButtonFromPointer(button), int(checked) != 0)
}
