package gui

//#include "gui.h"
import "C"
import (
	"log"
	"strings"
	"unsafe"
)

type QAccessibleActionInterface struct {
	ptr unsafe.Pointer
}

type QAccessibleActionInterface_ITF interface {
	QAccessibleActionInterface_PTR() *QAccessibleActionInterface
}

func (p *QAccessibleActionInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAccessibleActionInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAccessibleActionInterface(ptr QAccessibleActionInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleActionInterface_PTR().Pointer()
	}
	return nil
}

func NewQAccessibleActionInterfaceFromPointer(ptr unsafe.Pointer) *QAccessibleActionInterface {
	var n = new(QAccessibleActionInterface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessibleActionInterface) QAccessibleActionInterface_PTR() *QAccessibleActionInterface {
	return ptr
}

func (ptr *QAccessibleActionInterface) LocalizedActionDescription(actionName string) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleActionInterface::localizedActionDescription")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleActionInterface_LocalizedActionDescription(ptr.Pointer(), C.CString(actionName)))
	}
	return ""
}

func (ptr *QAccessibleActionInterface) LocalizedActionName(actionName string) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleActionInterface::localizedActionName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleActionInterface_LocalizedActionName(ptr.Pointer(), C.CString(actionName)))
	}
	return ""
}

func (ptr *QAccessibleActionInterface) ActionNames() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleActionInterface::actionNames")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QAccessibleActionInterface_ActionNames(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func QAccessibleActionInterface_DecreaseAction() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleActionInterface::decreaseAction")
		}
	}()

	return C.GoString(C.QAccessibleActionInterface_QAccessibleActionInterface_DecreaseAction())
}

func (ptr *QAccessibleActionInterface) DoAction(actionName string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleActionInterface::doAction")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAccessibleActionInterface_DoAction(ptr.Pointer(), C.CString(actionName))
	}
}

func QAccessibleActionInterface_IncreaseAction() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleActionInterface::increaseAction")
		}
	}()

	return C.GoString(C.QAccessibleActionInterface_QAccessibleActionInterface_IncreaseAction())
}

func (ptr *QAccessibleActionInterface) KeyBindingsForAction(actionName string) []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleActionInterface::keyBindingsForAction")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QAccessibleActionInterface_KeyBindingsForAction(ptr.Pointer(), C.CString(actionName))), ",,,")
	}
	return make([]string, 0)
}

func QAccessibleActionInterface_NextPageAction() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleActionInterface::nextPageAction")
		}
	}()

	return C.GoString(C.QAccessibleActionInterface_QAccessibleActionInterface_NextPageAction())
}

func QAccessibleActionInterface_PressAction() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleActionInterface::pressAction")
		}
	}()

	return C.GoString(C.QAccessibleActionInterface_QAccessibleActionInterface_PressAction())
}

func QAccessibleActionInterface_PreviousPageAction() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleActionInterface::previousPageAction")
		}
	}()

	return C.GoString(C.QAccessibleActionInterface_QAccessibleActionInterface_PreviousPageAction())
}

func QAccessibleActionInterface_ScrollDownAction() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleActionInterface::scrollDownAction")
		}
	}()

	return C.GoString(C.QAccessibleActionInterface_QAccessibleActionInterface_ScrollDownAction())
}

func QAccessibleActionInterface_ScrollLeftAction() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleActionInterface::scrollLeftAction")
		}
	}()

	return C.GoString(C.QAccessibleActionInterface_QAccessibleActionInterface_ScrollLeftAction())
}

func QAccessibleActionInterface_ScrollRightAction() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleActionInterface::scrollRightAction")
		}
	}()

	return C.GoString(C.QAccessibleActionInterface_QAccessibleActionInterface_ScrollRightAction())
}

func QAccessibleActionInterface_ScrollUpAction() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleActionInterface::scrollUpAction")
		}
	}()

	return C.GoString(C.QAccessibleActionInterface_QAccessibleActionInterface_ScrollUpAction())
}

func QAccessibleActionInterface_SetFocusAction() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleActionInterface::setFocusAction")
		}
	}()

	return C.GoString(C.QAccessibleActionInterface_QAccessibleActionInterface_SetFocusAction())
}

func QAccessibleActionInterface_ShowMenuAction() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleActionInterface::showMenuAction")
		}
	}()

	return C.GoString(C.QAccessibleActionInterface_QAccessibleActionInterface_ShowMenuAction())
}

func QAccessibleActionInterface_ToggleAction() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleActionInterface::toggleAction")
		}
	}()

	return C.GoString(C.QAccessibleActionInterface_QAccessibleActionInterface_ToggleAction())
}

func (ptr *QAccessibleActionInterface) DestroyQAccessibleActionInterface() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleActionInterface::~QAccessibleActionInterface")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAccessibleActionInterface_DestroyQAccessibleActionInterface(ptr.Pointer())
	}
}
