package widgets

//#include "qgraphicsopacityeffect.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QGraphicsOpacityEffect struct {
	QGraphicsEffect
}

type QGraphicsOpacityEffectITF interface {
	QGraphicsEffectITF
	QGraphicsOpacityEffectPTR() *QGraphicsOpacityEffect
}

func PointerFromQGraphicsOpacityEffect(ptr QGraphicsOpacityEffectITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsOpacityEffectPTR().Pointer()
	}
	return nil
}

func QGraphicsOpacityEffectFromPointer(ptr unsafe.Pointer) *QGraphicsOpacityEffect {
	var n = new(QGraphicsOpacityEffect)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGraphicsOpacityEffect_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsOpacityEffect) QGraphicsOpacityEffectPTR() *QGraphicsOpacityEffect {
	return ptr
}

func (ptr *QGraphicsOpacityEffect) SetOpacityMask(mask gui.QBrushITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsOpacityEffect_SetOpacityMask(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQBrush(mask)))
	}
}

func NewQGraphicsOpacityEffect(parent core.QObjectITF) *QGraphicsOpacityEffect {
	return QGraphicsOpacityEffectFromPointer(unsafe.Pointer(C.QGraphicsOpacityEffect_NewQGraphicsOpacityEffect(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QGraphicsOpacityEffect) DestroyQGraphicsOpacityEffect() {
	if ptr.Pointer() != nil {
		C.QGraphicsOpacityEffect_DestroyQGraphicsOpacityEffect(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
