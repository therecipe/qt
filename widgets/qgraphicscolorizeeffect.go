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

type QGraphicsColorizeEffect_ITF interface {
	QGraphicsEffect_ITF
	QGraphicsColorizeEffect_PTR() *QGraphicsColorizeEffect
}

func PointerFromQGraphicsColorizeEffect(ptr QGraphicsColorizeEffect_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsColorizeEffect_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsColorizeEffectFromPointer(ptr unsafe.Pointer) *QGraphicsColorizeEffect {
	var n = new(QGraphicsColorizeEffect)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGraphicsColorizeEffect_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsColorizeEffect) QGraphicsColorizeEffect_PTR() *QGraphicsColorizeEffect {
	return ptr
}

func (ptr *QGraphicsColorizeEffect) Color() *gui.QColor {
	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QGraphicsColorizeEffect_Color(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsColorizeEffect) SetColor(c gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsColorizeEffect_SetColor(ptr.Pointer(), gui.PointerFromQColor(c))
	}
}

func (ptr *QGraphicsColorizeEffect) SetStrength(strength float64) {
	if ptr.Pointer() != nil {
		C.QGraphicsColorizeEffect_SetStrength(ptr.Pointer(), C.double(strength))
	}
}

func (ptr *QGraphicsColorizeEffect) Strength() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGraphicsColorizeEffect_Strength(ptr.Pointer()))
	}
	return 0
}

func NewQGraphicsColorizeEffect(parent core.QObject_ITF) *QGraphicsColorizeEffect {
	return NewQGraphicsColorizeEffectFromPointer(C.QGraphicsColorizeEffect_NewQGraphicsColorizeEffect(core.PointerFromQObject(parent)))
}

func (ptr *QGraphicsColorizeEffect) ConnectColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {
		C.QGraphicsColorizeEffect_ConnectColorChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "colorChanged", f)
	}
}

func (ptr *QGraphicsColorizeEffect) DisconnectColorChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsColorizeEffect_DisconnectColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "colorChanged")
	}
}

//export callbackQGraphicsColorizeEffectColorChanged
func callbackQGraphicsColorizeEffectColorChanged(ptrName *C.char, color unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "colorChanged").(func(*gui.QColor))(gui.NewQColorFromPointer(color))
}

func (ptr *QGraphicsColorizeEffect) DestroyQGraphicsColorizeEffect() {
	if ptr.Pointer() != nil {
		C.QGraphicsColorizeEffect_DestroyQGraphicsColorizeEffect(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
