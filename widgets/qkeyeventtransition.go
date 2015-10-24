package widgets

//#include "qkeyeventtransition.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QKeyEventTransition struct {
	core.QEventTransition
}

type QKeyEventTransitionITF interface {
	core.QEventTransitionITF
	QKeyEventTransitionPTR() *QKeyEventTransition
}

func PointerFromQKeyEventTransition(ptr QKeyEventTransitionITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QKeyEventTransitionPTR().Pointer()
	}
	return nil
}

func QKeyEventTransitionFromPointer(ptr unsafe.Pointer) *QKeyEventTransition {
	var n = new(QKeyEventTransition)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QKeyEventTransition_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QKeyEventTransition) QKeyEventTransitionPTR() *QKeyEventTransition {
	return ptr
}

func NewQKeyEventTransition2(object core.QObjectITF, ty core.QEvent__Type, key int, sourceState core.QStateITF) *QKeyEventTransition {
	return QKeyEventTransitionFromPointer(unsafe.Pointer(C.QKeyEventTransition_NewQKeyEventTransition2(C.QtObjectPtr(core.PointerFromQObject(object)), C.int(ty), C.int(key), C.QtObjectPtr(core.PointerFromQState(sourceState)))))
}

func NewQKeyEventTransition(sourceState core.QStateITF) *QKeyEventTransition {
	return QKeyEventTransitionFromPointer(unsafe.Pointer(C.QKeyEventTransition_NewQKeyEventTransition(C.QtObjectPtr(core.PointerFromQState(sourceState)))))
}

func (ptr *QKeyEventTransition) Key() int {
	if ptr.Pointer() != nil {
		return int(C.QKeyEventTransition_Key(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QKeyEventTransition) ModifierMask() core.Qt__KeyboardModifier {
	if ptr.Pointer() != nil {
		return core.Qt__KeyboardModifier(C.QKeyEventTransition_ModifierMask(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QKeyEventTransition) SetKey(key int) {
	if ptr.Pointer() != nil {
		C.QKeyEventTransition_SetKey(C.QtObjectPtr(ptr.Pointer()), C.int(key))
	}
}

func (ptr *QKeyEventTransition) SetModifierMask(modifierMask core.Qt__KeyboardModifier) {
	if ptr.Pointer() != nil {
		C.QKeyEventTransition_SetModifierMask(C.QtObjectPtr(ptr.Pointer()), C.int(modifierMask))
	}
}

func (ptr *QKeyEventTransition) DestroyQKeyEventTransition() {
	if ptr.Pointer() != nil {
		C.QKeyEventTransition_DestroyQKeyEventTransition(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
