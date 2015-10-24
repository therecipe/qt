package gui

//#include "qaccessibleinterface.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAccessibleInterface struct {
	ptr unsafe.Pointer
}

type QAccessibleInterfaceITF interface {
	QAccessibleInterfacePTR() *QAccessibleInterface
}

func (p *QAccessibleInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAccessibleInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAccessibleInterface(ptr QAccessibleInterfaceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleInterfacePTR().Pointer()
	}
	return nil
}

func QAccessibleInterfaceFromPointer(ptr unsafe.Pointer) *QAccessibleInterface {
	var n = new(QAccessibleInterface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessibleInterface) QAccessibleInterfacePTR() *QAccessibleInterface {
	return ptr
}

func (ptr *QAccessibleInterface) ActionInterface() *QAccessibleActionInterface {
	if ptr.Pointer() != nil {
		return QAccessibleActionInterfaceFromPointer(unsafe.Pointer(C.QAccessibleInterface_ActionInterface(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAccessibleInterface) Child(index int) *QAccessibleInterface {
	if ptr.Pointer() != nil {
		return QAccessibleInterfaceFromPointer(unsafe.Pointer(C.QAccessibleInterface_Child(C.QtObjectPtr(ptr.Pointer()), C.int(index))))
	}
	return nil
}

func (ptr *QAccessibleInterface) ChildAt(x int, y int) *QAccessibleInterface {
	if ptr.Pointer() != nil {
		return QAccessibleInterfaceFromPointer(unsafe.Pointer(C.QAccessibleInterface_ChildAt(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y))))
	}
	return nil
}

func (ptr *QAccessibleInterface) ChildCount() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleInterface_ChildCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAccessibleInterface) FocusChild() *QAccessibleInterface {
	if ptr.Pointer() != nil {
		return QAccessibleInterfaceFromPointer(unsafe.Pointer(C.QAccessibleInterface_FocusChild(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAccessibleInterface) IndexOfChild(child QAccessibleInterfaceITF) int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleInterface_IndexOfChild(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAccessibleInterface(child))))
	}
	return 0
}

func (ptr *QAccessibleInterface) Interface_cast(ty QAccessible__InterfaceType) {
	if ptr.Pointer() != nil {
		C.QAccessibleInterface_Interface_cast(C.QtObjectPtr(ptr.Pointer()), C.int(ty))
	}
}

func (ptr *QAccessibleInterface) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QAccessibleInterface_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAccessibleInterface) Object() *core.QObject {
	if ptr.Pointer() != nil {
		return core.QObjectFromPointer(unsafe.Pointer(C.QAccessibleInterface_Object(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAccessibleInterface) Parent() *QAccessibleInterface {
	if ptr.Pointer() != nil {
		return QAccessibleInterfaceFromPointer(unsafe.Pointer(C.QAccessibleInterface_Parent(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAccessibleInterface) Role() QAccessible__Role {
	if ptr.Pointer() != nil {
		return QAccessible__Role(C.QAccessibleInterface_Role(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAccessibleInterface) SetText(t QAccessible__Text, text string) {
	if ptr.Pointer() != nil {
		C.QAccessibleInterface_SetText(C.QtObjectPtr(ptr.Pointer()), C.int(t), C.CString(text))
	}
}

func (ptr *QAccessibleInterface) TableCellInterface() *QAccessibleTableCellInterface {
	if ptr.Pointer() != nil {
		return QAccessibleTableCellInterfaceFromPointer(unsafe.Pointer(C.QAccessibleInterface_TableCellInterface(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAccessibleInterface) TableInterface() *QAccessibleTableInterface {
	if ptr.Pointer() != nil {
		return QAccessibleTableInterfaceFromPointer(unsafe.Pointer(C.QAccessibleInterface_TableInterface(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAccessibleInterface) Text(t QAccessible__Text) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleInterface_Text(C.QtObjectPtr(ptr.Pointer()), C.int(t)))
	}
	return ""
}

func (ptr *QAccessibleInterface) TextInterface() *QAccessibleTextInterface {
	if ptr.Pointer() != nil {
		return QAccessibleTextInterfaceFromPointer(unsafe.Pointer(C.QAccessibleInterface_TextInterface(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAccessibleInterface) ValueInterface() *QAccessibleValueInterface {
	if ptr.Pointer() != nil {
		return QAccessibleValueInterfaceFromPointer(unsafe.Pointer(C.QAccessibleInterface_ValueInterface(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAccessibleInterface) Window() *QWindow {
	if ptr.Pointer() != nil {
		return QWindowFromPointer(unsafe.Pointer(C.QAccessibleInterface_Window(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}
