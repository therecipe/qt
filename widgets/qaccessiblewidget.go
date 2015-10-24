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

type QAccessibleWidgetITF interface {
	gui.QAccessibleObjectITF
	gui.QAccessibleActionInterfaceITF
	QAccessibleWidgetPTR() *QAccessibleWidget
}

func (p *QAccessibleWidget) Pointer() unsafe.Pointer {
	return p.QAccessibleObjectPTR().Pointer()
}

func (p *QAccessibleWidget) SetPointer(ptr unsafe.Pointer) {
	p.QAccessibleObjectPTR().SetPointer(ptr)
	p.QAccessibleActionInterfacePTR().SetPointer(ptr)
}

func PointerFromQAccessibleWidget(ptr QAccessibleWidgetITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleWidgetPTR().Pointer()
	}
	return nil
}

func QAccessibleWidgetFromPointer(ptr unsafe.Pointer) *QAccessibleWidget {
	var n = new(QAccessibleWidget)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessibleWidget) QAccessibleWidgetPTR() *QAccessibleWidget {
	return ptr
}

func NewQAccessibleWidget(w QWidgetITF, role gui.QAccessible__Role, name string) *QAccessibleWidget {
	return QAccessibleWidgetFromPointer(unsafe.Pointer(C.QAccessibleWidget_NewQAccessibleWidget(C.QtObjectPtr(PointerFromQWidget(w)), C.int(role), C.CString(name))))
}

func (ptr *QAccessibleWidget) ActionNames() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QAccessibleWidget_ActionNames(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QAccessibleWidget) Child(index int) *gui.QAccessibleInterface {
	if ptr.Pointer() != nil {
		return gui.QAccessibleInterfaceFromPointer(unsafe.Pointer(C.QAccessibleWidget_Child(C.QtObjectPtr(ptr.Pointer()), C.int(index))))
	}
	return nil
}

func (ptr *QAccessibleWidget) ChildCount() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleWidget_ChildCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAccessibleWidget) DoAction(actionName string) {
	if ptr.Pointer() != nil {
		C.QAccessibleWidget_DoAction(C.QtObjectPtr(ptr.Pointer()), C.CString(actionName))
	}
}

func (ptr *QAccessibleWidget) FocusChild() *gui.QAccessibleInterface {
	if ptr.Pointer() != nil {
		return gui.QAccessibleInterfaceFromPointer(unsafe.Pointer(C.QAccessibleWidget_FocusChild(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAccessibleWidget) IndexOfChild(child gui.QAccessibleInterfaceITF) int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleWidget_IndexOfChild(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQAccessibleInterface(child))))
	}
	return 0
}

func (ptr *QAccessibleWidget) Interface_cast(t gui.QAccessible__InterfaceType) {
	if ptr.Pointer() != nil {
		C.QAccessibleWidget_Interface_cast(C.QtObjectPtr(ptr.Pointer()), C.int(t))
	}
}

func (ptr *QAccessibleWidget) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QAccessibleWidget_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAccessibleWidget) KeyBindingsForAction(actionName string) []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QAccessibleWidget_KeyBindingsForAction(C.QtObjectPtr(ptr.Pointer()), C.CString(actionName))), "|")
	}
	return make([]string, 0)
}

func (ptr *QAccessibleWidget) Parent() *gui.QAccessibleInterface {
	if ptr.Pointer() != nil {
		return gui.QAccessibleInterfaceFromPointer(unsafe.Pointer(C.QAccessibleWidget_Parent(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAccessibleWidget) Role() gui.QAccessible__Role {
	if ptr.Pointer() != nil {
		return gui.QAccessible__Role(C.QAccessibleWidget_Role(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAccessibleWidget) Text(t gui.QAccessible__Text) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleWidget_Text(C.QtObjectPtr(ptr.Pointer()), C.int(t)))
	}
	return ""
}

func (ptr *QAccessibleWidget) Window() *gui.QWindow {
	if ptr.Pointer() != nil {
		return gui.QWindowFromPointer(unsafe.Pointer(C.QAccessibleWidget_Window(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}
