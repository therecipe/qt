package widgets

//#include "widgets.h"
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
	for len(n.ObjectName()) < len("QGraphicsColorizeEffect_") {
		n.SetObjectName("QGraphicsColorizeEffect_" + qt.Identifier())
	}
	return n
}

func (ptr *QGraphicsColorizeEffect) QGraphicsColorizeEffect_PTR() *QGraphicsColorizeEffect {
	return ptr
}

func (ptr *QGraphicsColorizeEffect) Color() *gui.QColor {
	defer qt.Recovering("QGraphicsColorizeEffect::color")

	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QGraphicsColorizeEffect_Color(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsColorizeEffect) SetColor(c gui.QColor_ITF) {
	defer qt.Recovering("QGraphicsColorizeEffect::setColor")

	if ptr.Pointer() != nil {
		C.QGraphicsColorizeEffect_SetColor(ptr.Pointer(), gui.PointerFromQColor(c))
	}
}

func (ptr *QGraphicsColorizeEffect) SetStrength(strength float64) {
	defer qt.Recovering("QGraphicsColorizeEffect::setStrength")

	if ptr.Pointer() != nil {
		C.QGraphicsColorizeEffect_SetStrength(ptr.Pointer(), C.double(strength))
	}
}

func (ptr *QGraphicsColorizeEffect) Strength() float64 {
	defer qt.Recovering("QGraphicsColorizeEffect::strength")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsColorizeEffect_Strength(ptr.Pointer()))
	}
	return 0
}

func NewQGraphicsColorizeEffect(parent core.QObject_ITF) *QGraphicsColorizeEffect {
	defer qt.Recovering("QGraphicsColorizeEffect::QGraphicsColorizeEffect")

	return NewQGraphicsColorizeEffectFromPointer(C.QGraphicsColorizeEffect_NewQGraphicsColorizeEffect(core.PointerFromQObject(parent)))
}

func (ptr *QGraphicsColorizeEffect) ConnectColorChanged(f func(color *gui.QColor)) {
	defer qt.Recovering("connect QGraphicsColorizeEffect::colorChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsColorizeEffect_ConnectColorChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "colorChanged", f)
	}
}

func (ptr *QGraphicsColorizeEffect) DisconnectColorChanged() {
	defer qt.Recovering("disconnect QGraphicsColorizeEffect::colorChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsColorizeEffect_DisconnectColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "colorChanged")
	}
}

//export callbackQGraphicsColorizeEffectColorChanged
func callbackQGraphicsColorizeEffectColorChanged(ptrName *C.char, color unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsColorizeEffect::colorChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "colorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QGraphicsColorizeEffect) ConnectStrengthChanged(f func(strength float64)) {
	defer qt.Recovering("connect QGraphicsColorizeEffect::strengthChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsColorizeEffect_ConnectStrengthChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "strengthChanged", f)
	}
}

func (ptr *QGraphicsColorizeEffect) DisconnectStrengthChanged() {
	defer qt.Recovering("disconnect QGraphicsColorizeEffect::strengthChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsColorizeEffect_DisconnectStrengthChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "strengthChanged")
	}
}

//export callbackQGraphicsColorizeEffectStrengthChanged
func callbackQGraphicsColorizeEffectStrengthChanged(ptrName *C.char, strength C.double) {
	defer qt.Recovering("callback QGraphicsColorizeEffect::strengthChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "strengthChanged"); signal != nil {
		signal.(func(float64))(float64(strength))
	}

}

func (ptr *QGraphicsColorizeEffect) DestroyQGraphicsColorizeEffect() {
	defer qt.Recovering("QGraphicsColorizeEffect::~QGraphicsColorizeEffect")

	if ptr.Pointer() != nil {
		C.QGraphicsColorizeEffect_DestroyQGraphicsColorizeEffect(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsColorizeEffect) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGraphicsColorizeEffect::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGraphicsColorizeEffect) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGraphicsColorizeEffect::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGraphicsColorizeEffectTimerEvent
func callbackQGraphicsColorizeEffectTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsColorizeEffect::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsColorizeEffect) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGraphicsColorizeEffect::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGraphicsColorizeEffect) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGraphicsColorizeEffect::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGraphicsColorizeEffectChildEvent
func callbackQGraphicsColorizeEffectChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsColorizeEffect::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsColorizeEffect) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsColorizeEffect::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGraphicsColorizeEffect) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGraphicsColorizeEffect::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGraphicsColorizeEffectCustomEvent
func callbackQGraphicsColorizeEffectCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsColorizeEffect::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
