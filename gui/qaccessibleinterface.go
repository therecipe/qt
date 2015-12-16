package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAccessibleInterface struct {
	ptr unsafe.Pointer
}

type QAccessibleInterface_ITF interface {
	QAccessibleInterface_PTR() *QAccessibleInterface
}

func (p *QAccessibleInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAccessibleInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAccessibleInterface(ptr QAccessibleInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleInterface_PTR().Pointer()
	}
	return nil
}

func NewQAccessibleInterfaceFromPointer(ptr unsafe.Pointer) *QAccessibleInterface {
	var n = new(QAccessibleInterface)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QAccessibleInterface_") {
		n.SetObjectNameAbs("QAccessibleInterface_" + qt.Identifier())
	}
	return n
}

func (ptr *QAccessibleInterface) QAccessibleInterface_PTR() *QAccessibleInterface {
	return ptr
}

func (ptr *QAccessibleInterface) ActionInterface() *QAccessibleActionInterface {
	defer qt.Recovering("QAccessibleInterface::actionInterface")

	if ptr.Pointer() != nil {
		return NewQAccessibleActionInterfaceFromPointer(C.QAccessibleInterface_ActionInterface(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleInterface) BackgroundColor() *QColor {
	defer qt.Recovering("QAccessibleInterface::backgroundColor")

	if ptr.Pointer() != nil {
		return NewQColorFromPointer(C.QAccessibleInterface_BackgroundColor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleInterface) Child(index int) *QAccessibleInterface {
	defer qt.Recovering("QAccessibleInterface::child")

	if ptr.Pointer() != nil {
		return NewQAccessibleInterfaceFromPointer(C.QAccessibleInterface_Child(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QAccessibleInterface) ChildAt(x int, y int) *QAccessibleInterface {
	defer qt.Recovering("QAccessibleInterface::childAt")

	if ptr.Pointer() != nil {
		return NewQAccessibleInterfaceFromPointer(C.QAccessibleInterface_ChildAt(ptr.Pointer(), C.int(x), C.int(y)))
	}
	return nil
}

func (ptr *QAccessibleInterface) ChildCount() int {
	defer qt.Recovering("QAccessibleInterface::childCount")

	if ptr.Pointer() != nil {
		return int(C.QAccessibleInterface_ChildCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleInterface) FocusChild() *QAccessibleInterface {
	defer qt.Recovering("QAccessibleInterface::focusChild")

	if ptr.Pointer() != nil {
		return NewQAccessibleInterfaceFromPointer(C.QAccessibleInterface_FocusChild(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleInterface) ForegroundColor() *QColor {
	defer qt.Recovering("QAccessibleInterface::foregroundColor")

	if ptr.Pointer() != nil {
		return NewQColorFromPointer(C.QAccessibleInterface_ForegroundColor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleInterface) IndexOfChild(child QAccessibleInterface_ITF) int {
	defer qt.Recovering("QAccessibleInterface::indexOfChild")

	if ptr.Pointer() != nil {
		return int(C.QAccessibleInterface_IndexOfChild(ptr.Pointer(), PointerFromQAccessibleInterface(child)))
	}
	return 0
}

func (ptr *QAccessibleInterface) Interface_cast(ty QAccessible__InterfaceType) unsafe.Pointer {
	defer qt.Recovering("QAccessibleInterface::interface_cast")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QAccessibleInterface_Interface_cast(ptr.Pointer(), C.int(ty)))
	}
	return nil
}

func (ptr *QAccessibleInterface) IsValid() bool {
	defer qt.Recovering("QAccessibleInterface::isValid")

	if ptr.Pointer() != nil {
		return C.QAccessibleInterface_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAccessibleInterface) Object() *core.QObject {
	defer qt.Recovering("QAccessibleInterface::object")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QAccessibleInterface_Object(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleInterface) Parent() *QAccessibleInterface {
	defer qt.Recovering("QAccessibleInterface::parent")

	if ptr.Pointer() != nil {
		return NewQAccessibleInterfaceFromPointer(C.QAccessibleInterface_Parent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleInterface) Rect() *core.QRect {
	defer qt.Recovering("QAccessibleInterface::rect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QAccessibleInterface_Rect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleInterface) Role() QAccessible__Role {
	defer qt.Recovering("QAccessibleInterface::role")

	if ptr.Pointer() != nil {
		return QAccessible__Role(C.QAccessibleInterface_Role(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleInterface) SetText(t QAccessible__Text, text string) {
	defer qt.Recovering("QAccessibleInterface::setText")

	if ptr.Pointer() != nil {
		C.QAccessibleInterface_SetText(ptr.Pointer(), C.int(t), C.CString(text))
	}
}

func (ptr *QAccessibleInterface) TableCellInterface() *QAccessibleTableCellInterface {
	defer qt.Recovering("QAccessibleInterface::tableCellInterface")

	if ptr.Pointer() != nil {
		return NewQAccessibleTableCellInterfaceFromPointer(C.QAccessibleInterface_TableCellInterface(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleInterface) TableInterface() *QAccessibleTableInterface {
	defer qt.Recovering("QAccessibleInterface::tableInterface")

	if ptr.Pointer() != nil {
		return NewQAccessibleTableInterfaceFromPointer(C.QAccessibleInterface_TableInterface(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleInterface) Text(t QAccessible__Text) string {
	defer qt.Recovering("QAccessibleInterface::text")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleInterface_Text(ptr.Pointer(), C.int(t)))
	}
	return ""
}

func (ptr *QAccessibleInterface) TextInterface() *QAccessibleTextInterface {
	defer qt.Recovering("QAccessibleInterface::textInterface")

	if ptr.Pointer() != nil {
		return NewQAccessibleTextInterfaceFromPointer(C.QAccessibleInterface_TextInterface(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleInterface) ValueInterface() *QAccessibleValueInterface {
	defer qt.Recovering("QAccessibleInterface::valueInterface")

	if ptr.Pointer() != nil {
		return NewQAccessibleValueInterfaceFromPointer(C.QAccessibleInterface_ValueInterface(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleInterface) Window() *QWindow {
	defer qt.Recovering("QAccessibleInterface::window")

	if ptr.Pointer() != nil {
		return NewQWindowFromPointer(C.QAccessibleInterface_Window(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleInterface) ObjectNameAbs() string {
	defer qt.Recovering("QAccessibleInterface::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleInterface_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAccessibleInterface) SetObjectNameAbs(name string) {
	defer qt.Recovering("QAccessibleInterface::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QAccessibleInterface_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
