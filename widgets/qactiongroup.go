package widgets

//#include "widgets.h"
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
	for len(n.ObjectName()) < len("QActionGroup_") {
		n.SetObjectName("QActionGroup_" + qt.Identifier())
	}
	return n
}

func (ptr *QActionGroup) QActionGroup_PTR() *QActionGroup {
	return ptr
}

func (ptr *QActionGroup) AddAction(action QAction_ITF) *QAction {
	defer qt.Recovering("QActionGroup::addAction")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QActionGroup_AddAction(ptr.Pointer(), PointerFromQAction(action)))
	}
	return nil
}

func (ptr *QActionGroup) IsEnabled() bool {
	defer qt.Recovering("QActionGroup::isEnabled")

	if ptr.Pointer() != nil {
		return C.QActionGroup_IsEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QActionGroup) IsExclusive() bool {
	defer qt.Recovering("QActionGroup::isExclusive")

	if ptr.Pointer() != nil {
		return C.QActionGroup_IsExclusive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QActionGroup) IsVisible() bool {
	defer qt.Recovering("QActionGroup::isVisible")

	if ptr.Pointer() != nil {
		return C.QActionGroup_IsVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QActionGroup) SetEnabled(v bool) {
	defer qt.Recovering("QActionGroup::setEnabled")

	if ptr.Pointer() != nil {
		C.QActionGroup_SetEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QActionGroup) SetExclusive(v bool) {
	defer qt.Recovering("QActionGroup::setExclusive")

	if ptr.Pointer() != nil {
		C.QActionGroup_SetExclusive(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QActionGroup) SetVisible(v bool) {
	defer qt.Recovering("QActionGroup::setVisible")

	if ptr.Pointer() != nil {
		C.QActionGroup_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func NewQActionGroup(parent core.QObject_ITF) *QActionGroup {
	defer qt.Recovering("QActionGroup::QActionGroup")

	return NewQActionGroupFromPointer(C.QActionGroup_NewQActionGroup(core.PointerFromQObject(parent)))
}

func (ptr *QActionGroup) AddAction3(icon gui.QIcon_ITF, text string) *QAction {
	defer qt.Recovering("QActionGroup::addAction")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QActionGroup_AddAction3(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(text)))
	}
	return nil
}

func (ptr *QActionGroup) AddAction2(text string) *QAction {
	defer qt.Recovering("QActionGroup::addAction")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QActionGroup_AddAction2(ptr.Pointer(), C.CString(text)))
	}
	return nil
}

func (ptr *QActionGroup) CheckedAction() *QAction {
	defer qt.Recovering("QActionGroup::checkedAction")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QActionGroup_CheckedAction(ptr.Pointer()))
	}
	return nil
}

func (ptr *QActionGroup) ConnectHovered(f func(action *QAction)) {
	defer qt.Recovering("connect QActionGroup::hovered")

	if ptr.Pointer() != nil {
		C.QActionGroup_ConnectHovered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "hovered", f)
	}
}

func (ptr *QActionGroup) DisconnectHovered() {
	defer qt.Recovering("disconnect QActionGroup::hovered")

	if ptr.Pointer() != nil {
		C.QActionGroup_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "hovered")
	}
}

//export callbackQActionGroupHovered
func callbackQActionGroupHovered(ptrName *C.char, action unsafe.Pointer) {
	defer qt.Recovering("callback QActionGroup::hovered")

	var signal = qt.GetSignal(C.GoString(ptrName), "hovered")
	if signal != nil {
		signal.(func(*QAction))(NewQActionFromPointer(action))
	}

}

func (ptr *QActionGroup) RemoveAction(action QAction_ITF) {
	defer qt.Recovering("QActionGroup::removeAction")

	if ptr.Pointer() != nil {
		C.QActionGroup_RemoveAction(ptr.Pointer(), PointerFromQAction(action))
	}
}

func (ptr *QActionGroup) SetDisabled(b bool) {
	defer qt.Recovering("QActionGroup::setDisabled")

	if ptr.Pointer() != nil {
		C.QActionGroup_SetDisabled(ptr.Pointer(), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QActionGroup) ConnectTriggered(f func(action *QAction)) {
	defer qt.Recovering("connect QActionGroup::triggered")

	if ptr.Pointer() != nil {
		C.QActionGroup_ConnectTriggered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "triggered", f)
	}
}

func (ptr *QActionGroup) DisconnectTriggered() {
	defer qt.Recovering("disconnect QActionGroup::triggered")

	if ptr.Pointer() != nil {
		C.QActionGroup_DisconnectTriggered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "triggered")
	}
}

//export callbackQActionGroupTriggered
func callbackQActionGroupTriggered(ptrName *C.char, action unsafe.Pointer) {
	defer qt.Recovering("callback QActionGroup::triggered")

	var signal = qt.GetSignal(C.GoString(ptrName), "triggered")
	if signal != nil {
		signal.(func(*QAction))(NewQActionFromPointer(action))
	}

}

func (ptr *QActionGroup) DestroyQActionGroup() {
	defer qt.Recovering("QActionGroup::~QActionGroup")

	if ptr.Pointer() != nil {
		C.QActionGroup_DestroyQActionGroup(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
