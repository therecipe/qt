package widgets

//#include "qmouseeventtransition.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QMouseEventTransition struct {
	core.QEventTransition
}

type QMouseEventTransitionITF interface {
	core.QEventTransitionITF
	QMouseEventTransitionPTR() *QMouseEventTransition
}

func PointerFromQMouseEventTransition(ptr QMouseEventTransitionITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMouseEventTransitionPTR().Pointer()
	}
	return nil
}

func QMouseEventTransitionFromPointer(ptr unsafe.Pointer) *QMouseEventTransition {
	var n = new(QMouseEventTransition)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMouseEventTransition_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMouseEventTransition) QMouseEventTransitionPTR() *QMouseEventTransition {
	return ptr
}

func NewQMouseEventTransition2(object core.QObjectITF, ty core.QEvent__Type, button core.Qt__MouseButton, sourceState core.QStateITF) *QMouseEventTransition {
	return QMouseEventTransitionFromPointer(unsafe.Pointer(C.QMouseEventTransition_NewQMouseEventTransition2(C.QtObjectPtr(core.PointerFromQObject(object)), C.int(ty), C.int(button), C.QtObjectPtr(core.PointerFromQState(sourceState)))))
}

func NewQMouseEventTransition(sourceState core.QStateITF) *QMouseEventTransition {
	return QMouseEventTransitionFromPointer(unsafe.Pointer(C.QMouseEventTransition_NewQMouseEventTransition(C.QtObjectPtr(core.PointerFromQState(sourceState)))))
}

func (ptr *QMouseEventTransition) Button() core.Qt__MouseButton {
	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QMouseEventTransition_Button(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMouseEventTransition) ModifierMask() core.Qt__KeyboardModifier {
	if ptr.Pointer() != nil {
		return core.Qt__KeyboardModifier(C.QMouseEventTransition_ModifierMask(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMouseEventTransition) SetButton(button core.Qt__MouseButton) {
	if ptr.Pointer() != nil {
		C.QMouseEventTransition_SetButton(C.QtObjectPtr(ptr.Pointer()), C.int(button))
	}
}

func (ptr *QMouseEventTransition) SetHitTestPath(path gui.QPainterPathITF) {
	if ptr.Pointer() != nil {
		C.QMouseEventTransition_SetHitTestPath(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainterPath(path)))
	}
}

func (ptr *QMouseEventTransition) SetModifierMask(modifierMask core.Qt__KeyboardModifier) {
	if ptr.Pointer() != nil {
		C.QMouseEventTransition_SetModifierMask(C.QtObjectPtr(ptr.Pointer()), C.int(modifierMask))
	}
}

func (ptr *QMouseEventTransition) DestroyQMouseEventTransition() {
	if ptr.Pointer() != nil {
		C.QMouseEventTransition_DestroyQMouseEventTransition(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
