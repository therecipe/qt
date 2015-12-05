package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"log"
	"unsafe"
)

type QMouseEventTransition struct {
	core.QEventTransition
}

type QMouseEventTransition_ITF interface {
	core.QEventTransition_ITF
	QMouseEventTransition_PTR() *QMouseEventTransition
}

func PointerFromQMouseEventTransition(ptr QMouseEventTransition_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMouseEventTransition_PTR().Pointer()
	}
	return nil
}

func NewQMouseEventTransitionFromPointer(ptr unsafe.Pointer) *QMouseEventTransition {
	var n = new(QMouseEventTransition)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMouseEventTransition_") {
		n.SetObjectName("QMouseEventTransition_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMouseEventTransition) QMouseEventTransition_PTR() *QMouseEventTransition {
	return ptr
}

func NewQMouseEventTransition2(object core.QObject_ITF, ty core.QEvent__Type, button core.Qt__MouseButton, sourceState core.QState_ITF) *QMouseEventTransition {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMouseEventTransition::QMouseEventTransition")
		}
	}()

	return NewQMouseEventTransitionFromPointer(C.QMouseEventTransition_NewQMouseEventTransition2(core.PointerFromQObject(object), C.int(ty), C.int(button), core.PointerFromQState(sourceState)))
}

func NewQMouseEventTransition(sourceState core.QState_ITF) *QMouseEventTransition {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMouseEventTransition::QMouseEventTransition")
		}
	}()

	return NewQMouseEventTransitionFromPointer(C.QMouseEventTransition_NewQMouseEventTransition(core.PointerFromQState(sourceState)))
}

func (ptr *QMouseEventTransition) Button() core.Qt__MouseButton {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMouseEventTransition::button")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QMouseEventTransition_Button(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMouseEventTransition) ModifierMask() core.Qt__KeyboardModifier {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMouseEventTransition::modifierMask")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__KeyboardModifier(C.QMouseEventTransition_ModifierMask(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMouseEventTransition) SetButton(button core.Qt__MouseButton) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMouseEventTransition::setButton")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMouseEventTransition_SetButton(ptr.Pointer(), C.int(button))
	}
}

func (ptr *QMouseEventTransition) SetHitTestPath(path gui.QPainterPath_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMouseEventTransition::setHitTestPath")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMouseEventTransition_SetHitTestPath(ptr.Pointer(), gui.PointerFromQPainterPath(path))
	}
}

func (ptr *QMouseEventTransition) SetModifierMask(modifierMask core.Qt__KeyboardModifier) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMouseEventTransition::setModifierMask")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMouseEventTransition_SetModifierMask(ptr.Pointer(), C.int(modifierMask))
	}
}

func (ptr *QMouseEventTransition) DestroyQMouseEventTransition() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMouseEventTransition::~QMouseEventTransition")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMouseEventTransition_DestroyQMouseEventTransition(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
