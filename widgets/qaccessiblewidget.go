package widgets

//#include "qaccessiblewidget.h"
import "C"
import (
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
	return n
}

func (ptr *QAccessibleWidget) QAccessibleWidget_PTR() *QAccessibleWidget {
	return ptr
}

func NewQAccessibleWidget(w QWidget_ITF, role gui.QAccessible__Role, name string) *QAccessibleWidget {
	return NewQAccessibleWidgetFromPointer(C.QAccessibleWidget_NewQAccessibleWidget(PointerFromQWidget(w), C.int(role), C.CString(name)))
}

func (ptr *QAccessibleWidget) ActionNames() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QAccessibleWidget_ActionNames(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QAccessibleWidget) BackgroundColor() *gui.QColor {
	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QAccessibleWidget_BackgroundColor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleWidget) Child(index int) *gui.QAccessibleInterface {
	if ptr.Pointer() != nil {
		return gui.NewQAccessibleInterfaceFromPointer(C.QAccessibleWidget_Child(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QAccessibleWidget) ChildCount() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleWidget_ChildCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleWidget) DoAction(actionName string) {
	if ptr.Pointer() != nil {
		C.QAccessibleWidget_DoAction(ptr.Pointer(), C.CString(actionName))
	}
}

func (ptr *QAccessibleWidget) FocusChild() *gui.QAccessibleInterface {
	if ptr.Pointer() != nil {
		return gui.NewQAccessibleInterfaceFromPointer(C.QAccessibleWidget_FocusChild(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleWidget) ForegroundColor() *gui.QColor {
	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QAccessibleWidget_ForegroundColor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleWidget) IndexOfChild(child gui.QAccessibleInterface_ITF) int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleWidget_IndexOfChild(ptr.Pointer(), gui.PointerFromQAccessibleInterface(child)))
	}
	return 0
}

func (ptr *QAccessibleWidget) Interface_cast(t gui.QAccessible__InterfaceType) unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QAccessibleWidget_Interface_cast(ptr.Pointer(), C.int(t)))
	}
	return nil
}

func (ptr *QAccessibleWidget) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QAccessibleWidget_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAccessibleWidget) KeyBindingsForAction(actionName string) []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QAccessibleWidget_KeyBindingsForAction(ptr.Pointer(), C.CString(actionName))), "|")
	}
	return make([]string, 0)
}

func (ptr *QAccessibleWidget) Parent() *gui.QAccessibleInterface {
	if ptr.Pointer() != nil {
		return gui.NewQAccessibleInterfaceFromPointer(C.QAccessibleWidget_Parent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleWidget) Role() gui.QAccessible__Role {
	if ptr.Pointer() != nil {
		return gui.QAccessible__Role(C.QAccessibleWidget_Role(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleWidget) Text(t gui.QAccessible__Text) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleWidget_Text(ptr.Pointer(), C.int(t)))
	}
	return ""
}

func (ptr *QAccessibleWidget) Window() *gui.QWindow {
	if ptr.Pointer() != nil {
		return gui.NewQWindowFromPointer(C.QAccessibleWidget_Window(ptr.Pointer()))
	}
	return nil
}
