package widgets

//#include "qgraphicscolorizeeffect.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QGraphicsColorizeEffect struct {
	QGraphicsEffect
}

type QGraphicsColorizeEffectITF interface {
	QGraphicsEffectITF
	QGraphicsColorizeEffectPTR() *QGraphicsColorizeEffect
}

func PointerFromQGraphicsColorizeEffect(ptr QGraphicsColorizeEffectITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsColorizeEffectPTR().Pointer()
	}
	return nil
}

func QGraphicsColorizeEffectFromPointer(ptr unsafe.Pointer) *QGraphicsColorizeEffect {
	var n = new(QGraphicsColorizeEffect)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGraphicsColorizeEffect_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsColorizeEffect) QGraphicsColorizeEffectPTR() *QGraphicsColorizeEffect {
	return ptr
}

func (ptr *QGraphicsColorizeEffect) SetColor(c gui.QColorITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsColorizeEffect_SetColor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQColor(c)))
	}
}

func NewQGraphicsColorizeEffect(parent core.QObjectITF) *QGraphicsColorizeEffect {
	return QGraphicsColorizeEffectFromPointer(unsafe.Pointer(C.QGraphicsColorizeEffect_NewQGraphicsColorizeEffect(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QGraphicsColorizeEffect) DestroyQGraphicsColorizeEffect() {
	if ptr.Pointer() != nil {
		C.QGraphicsColorizeEffect_DestroyQGraphicsColorizeEffect(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
