package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt/gui"
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleWidget::QAccessibleWidget")
		}
	}()

	return NewQAccessibleWidgetFromPointer(C.QAccessibleWidget_NewQAccessibleWidget(PointerFromQWidget(w), C.int(role), C.CString(name)))
}

func (ptr *QAccessibleWidget) ActionNames() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleWidget::actionNames")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QAccessibleWidget_ActionNames(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QAccessibleWidget) BackgroundColor() *gui.QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleWidget::backgroundColor")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QAccessibleWidget_BackgroundColor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleWidget) Child(index int) *gui.QAccessibleInterface {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleWidget::child")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQAccessibleInterfaceFromPointer(C.QAccessibleWidget_Child(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QAccessibleWidget) ChildCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleWidget::childCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAccessibleWidget_ChildCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleWidget) DoAction(actionName string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleWidget::doAction")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAccessibleWidget_DoAction(ptr.Pointer(), C.CString(actionName))
	}
}

func (ptr *QAccessibleWidget) FocusChild() *gui.QAccessibleInterface {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleWidget::focusChild")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQAccessibleInterfaceFromPointer(C.QAccessibleWidget_FocusChild(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleWidget) ForegroundColor() *gui.QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleWidget::foregroundColor")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QAccessibleWidget_ForegroundColor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleWidget) IndexOfChild(child gui.QAccessibleInterface_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleWidget::indexOfChild")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAccessibleWidget_IndexOfChild(ptr.Pointer(), gui.PointerFromQAccessibleInterface(child)))
	}
	return 0
}

func (ptr *QAccessibleWidget) Interface_cast(t gui.QAccessible__InterfaceType) unsafe.Pointer {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleWidget::interface_cast")
		}
	}()

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QAccessibleWidget_Interface_cast(ptr.Pointer(), C.int(t)))
	}
	return nil
}

func (ptr *QAccessibleWidget) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleWidget::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAccessibleWidget_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAccessibleWidget) KeyBindingsForAction(actionName string) []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleWidget::keyBindingsForAction")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QAccessibleWidget_KeyBindingsForAction(ptr.Pointer(), C.CString(actionName))), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QAccessibleWidget) Parent() *gui.QAccessibleInterface {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleWidget::parent")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQAccessibleInterfaceFromPointer(C.QAccessibleWidget_Parent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleWidget) Role() gui.QAccessible__Role {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleWidget::role")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.QAccessible__Role(C.QAccessibleWidget_Role(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleWidget) Text(t gui.QAccessible__Text) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleWidget::text")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleWidget_Text(ptr.Pointer(), C.int(t)))
	}
	return ""
}

func (ptr *QAccessibleWidget) Window() *gui.QWindow {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleWidget::window")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQWindowFromPointer(C.QAccessibleWidget_Window(ptr.Pointer()))
	}
	return nil
}
