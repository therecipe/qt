package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/gui"
	"strings"
	"unsafe"
)

type QAccessibleWidget struct {
	gui.QAccessibleObject
	gui.QAccessibleActionInterface
}

type QAccessibleWidget_ITF interface {
	gui.QAccessibleObject_ITF
	gui.QAccessibleActionInterface_ITF
	QAccessibleWidget_PTR() *QAccessibleWidget
}

func (p *QAccessibleWidget) Pointer() unsafe.Pointer {
	return p.QAccessibleObject_PTR().Pointer()
}

func (p *QAccessibleWidget) SetPointer(ptr unsafe.Pointer) {
	p.QAccessibleObject_PTR().SetPointer(ptr)
	p.QAccessibleActionInterface_PTR().SetPointer(ptr)
}

func PointerFromQAccessibleWidget(ptr QAccessibleWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleWidget_PTR().Pointer()
	}
	return nil
}

func NewQAccessibleWidgetFromPointer(ptr unsafe.Pointer) *QAccessibleWidget {
	var n = new(QAccessibleWidget)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QAccessibleWidget_") {
		n.SetObjectNameAbs("QAccessibleWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QAccessibleWidget) QAccessibleWidget_PTR() *QAccessibleWidget {
	return ptr
}

func NewQAccessibleWidget(w QWidget_ITF, role gui.QAccessible__Role, name string) *QAccessibleWidget {
	defer qt.Recovering("QAccessibleWidget::QAccessibleWidget")

	return NewQAccessibleWidgetFromPointer(C.QAccessibleWidget_NewQAccessibleWidget(PointerFromQWidget(w), C.int(role), C.CString(name)))
}

func (ptr *QAccessibleWidget) ActionNames() []string {
	defer qt.Recovering("QAccessibleWidget::actionNames")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QAccessibleWidget_ActionNames(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QAccessibleWidget) BackgroundColor() *gui.QColor {
	defer qt.Recovering("QAccessibleWidget::backgroundColor")

	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QAccessibleWidget_BackgroundColor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleWidget) Child(index int) *gui.QAccessibleInterface {
	defer qt.Recovering("QAccessibleWidget::child")

	if ptr.Pointer() != nil {
		return gui.NewQAccessibleInterfaceFromPointer(C.QAccessibleWidget_Child(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QAccessibleWidget) ChildCount() int {
	defer qt.Recovering("QAccessibleWidget::childCount")

	if ptr.Pointer() != nil {
		return int(C.QAccessibleWidget_ChildCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleWidget) ConnectDoAction(f func(actionName string)) {
	defer qt.Recovering("connect QAccessibleWidget::doAction")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "doAction", f)
	}
}

func (ptr *QAccessibleWidget) DisconnectDoAction() {
	defer qt.Recovering("disconnect QAccessibleWidget::doAction")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "doAction")
	}
}

//export callbackQAccessibleWidgetDoAction
func callbackQAccessibleWidgetDoAction(ptrName *C.char, actionName *C.char) bool {
	defer qt.Recovering("callback QAccessibleWidget::doAction")

	var signal = qt.GetSignal(C.GoString(ptrName), "doAction")
	if signal != nil {
		defer signal.(func(string))(C.GoString(actionName))
		return true
	}
	return false

}

func (ptr *QAccessibleWidget) FocusChild() *gui.QAccessibleInterface {
	defer qt.Recovering("QAccessibleWidget::focusChild")

	if ptr.Pointer() != nil {
		return gui.NewQAccessibleInterfaceFromPointer(C.QAccessibleWidget_FocusChild(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleWidget) ForegroundColor() *gui.QColor {
	defer qt.Recovering("QAccessibleWidget::foregroundColor")

	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QAccessibleWidget_ForegroundColor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleWidget) IndexOfChild(child gui.QAccessibleInterface_ITF) int {
	defer qt.Recovering("QAccessibleWidget::indexOfChild")

	if ptr.Pointer() != nil {
		return int(C.QAccessibleWidget_IndexOfChild(ptr.Pointer(), gui.PointerFromQAccessibleInterface(child)))
	}
	return 0
}

func (ptr *QAccessibleWidget) Interface_cast(t gui.QAccessible__InterfaceType) unsafe.Pointer {
	defer qt.Recovering("QAccessibleWidget::interface_cast")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QAccessibleWidget_Interface_cast(ptr.Pointer(), C.int(t)))
	}
	return nil
}

func (ptr *QAccessibleWidget) IsValid() bool {
	defer qt.Recovering("QAccessibleWidget::isValid")

	if ptr.Pointer() != nil {
		return C.QAccessibleWidget_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAccessibleWidget) KeyBindingsForAction(actionName string) []string {
	defer qt.Recovering("QAccessibleWidget::keyBindingsForAction")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QAccessibleWidget_KeyBindingsForAction(ptr.Pointer(), C.CString(actionName))), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QAccessibleWidget) Parent() *gui.QAccessibleInterface {
	defer qt.Recovering("QAccessibleWidget::parent")

	if ptr.Pointer() != nil {
		return gui.NewQAccessibleInterfaceFromPointer(C.QAccessibleWidget_Parent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleWidget) Role() gui.QAccessible__Role {
	defer qt.Recovering("QAccessibleWidget::role")

	if ptr.Pointer() != nil {
		return gui.QAccessible__Role(C.QAccessibleWidget_Role(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleWidget) Text(t gui.QAccessible__Text) string {
	defer qt.Recovering("QAccessibleWidget::text")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleWidget_Text(ptr.Pointer(), C.int(t)))
	}
	return ""
}

func (ptr *QAccessibleWidget) Window() *gui.QWindow {
	defer qt.Recovering("QAccessibleWidget::window")

	if ptr.Pointer() != nil {
		return gui.NewQWindowFromPointer(C.QAccessibleWidget_Window(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleWidget) ObjectNameAbs() string {
	defer qt.Recovering("QAccessibleWidget::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleWidget_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAccessibleWidget) SetObjectNameAbs(name string) {
	defer qt.Recovering("QAccessibleWidget::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QAccessibleWidget_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
