package widgets

//#include "qgraphicsdropshadoweffect.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QGraphicsDropShadowEffect struct {
	QGraphicsEffect
}

type QGraphicsDropShadowEffectITF interface {
	QGraphicsEffectITF
	QGraphicsDropShadowEffectPTR() *QGraphicsDropShadowEffect
}

func PointerFromQGraphicsDropShadowEffect(ptr QGraphicsDropShadowEffectITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsDropShadowEffectPTR().Pointer()
	}
	return nil
}

func QGraphicsDropShadowEffectFromPointer(ptr unsafe.Pointer) *QGraphicsDropShadowEffect {
	var n = new(QGraphicsDropShadowEffect)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGraphicsDropShadowEffect_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsDropShadowEffect) QGraphicsDropShadowEffectPTR() *QGraphicsDropShadowEffect {
	return ptr
}

func (ptr *QGraphicsDropShadowEffect) SetColor(color gui.QColorITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsDropShadowEffect_SetColor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQColor(color)))
	}
}

func (ptr *QGraphicsDropShadowEffect) SetOffset(ofs core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsDropShadowEffect_SetOffset(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(ofs)))
	}
}

func NewQGraphicsDropShadowEffect(parent core.QObjectITF) *QGraphicsDropShadowEffect {
	return QGraphicsDropShadowEffectFromPointer(unsafe.Pointer(C.QGraphicsDropShadowEffect_NewQGraphicsDropShadowEffect(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QGraphicsDropShadowEffect) DestroyQGraphicsDropShadowEffect() {
	if ptr.Pointer() != nil {
		C.QGraphicsDropShadowEffect_DestroyQGraphicsDropShadowEffect(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
