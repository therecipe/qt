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

type QActionGroupITF interface {
	core.QObjectITF
	QActionGroupPTR() *QActionGroup
}

func PointerFromQActionGroup(ptr QActionGroupITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QActionGroupPTR().Pointer()
	}
	return nil
}

func QActionGroupFromPointer(ptr unsafe.Pointer) *QActionGroup {
	var n = new(QActionGroup)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QActionGroup_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QActionGroup) QActionGroupPTR() *QActionGroup {
	return ptr
}

func (ptr *QActionGroup) AddAction(action QActionITF) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QActionGroup_AddAction(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAction(action)))))
	}
	return nil
}

func (ptr *QActionGroup) IsEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QActionGroup_IsEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QActionGroup) IsExclusive() bool {
	if ptr.Pointer() != nil {
		return C.QActionGroup_IsExclusive(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QActionGroup) IsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QActionGroup_IsVisible(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QActionGroup) SetEnabled(v bool) {
	if ptr.Pointer() != nil {
		C.QActionGroup_SetEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QActionGroup) SetExclusive(v bool) {
	if ptr.Pointer() != nil {
		C.QActionGroup_SetExclusive(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QActionGroup) SetVisible(v bool) {
	if ptr.Pointer() != nil {
		C.QActionGroup_SetVisible(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func NewQActionGroup(parent core.QObjectITF) *QActionGroup {
	return QActionGroupFromPointer(unsafe.Pointer(C.QActionGroup_NewQActionGroup(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QActionGroup) AddAction3(icon gui.QIconITF, text string) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QActionGroup_AddAction3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQIcon(icon)), C.CString(text))))
	}
	return nil
}

func (ptr *QActionGroup) AddAction2(text string) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QActionGroup_AddAction2(C.QtObjectPtr(ptr.Pointer()), C.CString(text))))
	}
	return nil
}

func (ptr *QActionGroup) CheckedAction() *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QActionGroup_CheckedAction(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QActionGroup) ConnectHovered(f func(action QActionITF)) {
	if ptr.Pointer() != nil {
		C.QActionGroup_ConnectHovered(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "hovered", f)
	}
}

func (ptr *QActionGroup) DisconnectHovered() {
	if ptr.Pointer() != nil {
		C.QActionGroup_DisconnectHovered(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "hovered")
	}
}

//export callbackQActionGroupHovered
func callbackQActionGroupHovered(ptrName *C.char, action unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "hovered").(func(*QAction))(QActionFromPointer(action))
}

func (ptr *QActionGroup) RemoveAction(action QActionITF) {
	if ptr.Pointer() != nil {
		C.QActionGroup_RemoveAction(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAction(action)))
	}
}

func (ptr *QActionGroup) SetDisabled(b bool) {
	if ptr.Pointer() != nil {
		C.QActionGroup_SetDisabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QActionGroup) ConnectTriggered(f func(action QActionITF)) {
	if ptr.Pointer() != nil {
		C.QActionGroup_ConnectTriggered(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "triggered", f)
	}
}

func (ptr *QActionGroup) DisconnectTriggered() {
	if ptr.Pointer() != nil {
		C.QActionGroup_DisconnectTriggered(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "triggered")
	}
}

//export callbackQActionGroupTriggered
func callbackQActionGroupTriggered(ptrName *C.char, action unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "triggered").(func(*QAction))(QActionFromPointer(action))
}

func (ptr *QActionGroup) DestroyQActionGroup() {
	if ptr.Pointer() != nil {
		C.QActionGroup_DestroyQActionGroup(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
