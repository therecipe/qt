package widgets

//#include "qactiongroup.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QActionGroup struct {
	core.QObject
}

type QActionGroup_ITF interface {
	core.QObject_ITF
	QActionGroup_PTR() *QActionGroup
}

func PointerFromQActionGroup(ptr QActionGroup_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QActionGroup_PTR().Pointer()
	}
	return nil
}

func NewQActionGroupFromPointer(ptr unsafe.Pointer) *QActionGroup {
	var n = new(QActionGroup)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QActionGroup_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QActionGroup) QActionGroup_PTR() *QActionGroup {
	return ptr
}

func (ptr *QActionGroup) AddAction(action QAction_ITF) *QAction {
	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QActionGroup_AddAction(ptr.Pointer(), PointerFromQAction(action)))
	}
	return nil
}

func (ptr *QActionGroup) IsEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QActionGroup_IsEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QActionGroup) IsExclusive() bool {
	if ptr.Pointer() != nil {
		return C.QActionGroup_IsExclusive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QActionGroup) IsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QActionGroup_IsVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QActionGroup) SetEnabled(v bool) {
	if ptr.Pointer() != nil {
		C.QActionGroup_SetEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QActionGroup) SetExclusive(v bool) {
	if ptr.Pointer() != nil {
		C.QActionGroup_SetExclusive(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QActionGroup) SetVisible(v bool) {
	if ptr.Pointer() != nil {
		C.QActionGroup_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func NewQActionGroup(parent core.QObject_ITF) *QActionGroup {
	return NewQActionGroupFromPointer(C.QActionGroup_NewQActionGroup(core.PointerFromQObject(parent)))
}

func (ptr *QActionGroup) AddAction3(icon gui.QIcon_ITF, text string) *QAction {
	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QActionGroup_AddAction3(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(text)))
	}
	return nil
}

func (ptr *QActionGroup) AddAction2(text string) *QAction {
	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QActionGroup_AddAction2(ptr.Pointer(), C.CString(text)))
	}
	return nil
}

func (ptr *QActionGroup) CheckedAction() *QAction {
	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QActionGroup_CheckedAction(ptr.Pointer()))
	}
	return nil
}

func (ptr *QActionGroup) ConnectHovered(f func(action *QAction)) {
	if ptr.Pointer() != nil {
		C.QActionGroup_ConnectHovered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "hovered", f)
	}
}

func (ptr *QActionGroup) DisconnectHovered() {
	if ptr.Pointer() != nil {
		C.QActionGroup_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "hovered")
	}
}

//export callbackQActionGroupHovered
func callbackQActionGroupHovered(ptrName *C.char, action unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "hovered").(func(*QAction))(NewQActionFromPointer(action))
}

func (ptr *QActionGroup) RemoveAction(action QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QActionGroup_RemoveAction(ptr.Pointer(), PointerFromQAction(action))
	}
}

func (ptr *QActionGroup) SetDisabled(b bool) {
	if ptr.Pointer() != nil {
		C.QActionGroup_SetDisabled(ptr.Pointer(), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QActionGroup) ConnectTriggered(f func(action *QAction)) {
	if ptr.Pointer() != nil {
		C.QActionGroup_ConnectTriggered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "triggered", f)
	}
}

func (ptr *QActionGroup) DisconnectTriggered() {
	if ptr.Pointer() != nil {
		C.QActionGroup_DisconnectTriggered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "triggered")
	}
}

//export callbackQActionGroupTriggered
func callbackQActionGroupTriggered(ptrName *C.char, action unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "triggered").(func(*QAction))(NewQActionFromPointer(action))
}

func (ptr *QActionGroup) DestroyQActionGroup() {
	if ptr.Pointer() != nil {
		C.QActionGroup_DestroyQActionGroup(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
