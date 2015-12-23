package widgets

//#include "widgets.h"
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

type QGraphicsOpacityEffect_ITF interface {
	QGraphicsEffect_ITF
	QGraphicsOpacityEffect_PTR() *QGraphicsOpacityEffect
}

func PointerFromQGraphicsOpacityEffect(ptr QGraphicsOpacityEffect_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsOpacityEffect_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsOpacityEffectFromPointer(ptr unsafe.Pointer) *QGraphicsOpacityEffect {
	var n = new(QGraphicsOpacityEffect)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGraphicsOpacityEffect_") {
		n.SetObjectName("QGraphicsOpacityEffect_" + qt.Identifier())
	}
	return n
}

func (ptr *QGraphicsOpacityEffect) QGraphicsOpacityEffect_PTR() *QGraphicsOpacityEffect {
	return ptr
}

func (ptr *QGraphicsOpacityEffect) Opacity() float64 {
	defer qt.Recovering("QGraphicsOpacityEffect::opacity")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsOpacityEffect_Opacity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsOpacityEffect) OpacityMask() *gui.QBrush {
	defer qt.Recovering("QGraphicsOpacityEffect::opacityMask")

	if ptr.Pointer() != nil {
		return gui.NewQBrushFromPointer(C.QGraphicsOpacityEffect_OpacityMask(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsOpacityEffect) SetOpacity(opacity float64) {
	defer qt.Recovering("QGraphicsOpacityEffect::setOpacity")

	if ptr.Pointer() != nil {
		C.QGraphicsOpacityEffect_SetOpacity(ptr.Pointer(), C.double(opacity))
	}
}

func (ptr *QGraphicsOpacityEffect) SetOpacityMask(mask gui.QBrush_ITF) {
	defer qt.Recovering("QGraphicsOpacityEffect::setOpacityMask")

	if ptr.Pointer() != nil {
		C.QGraphicsOpacityEffect_SetOpacityMask(ptr.Pointer(), gui.PointerFromQBrush(mask))
	}
}

func NewQGraphicsOpacityEffect(parent core.QObject_ITF) *QGraphicsOpacityEffect {
	defer qt.Recovering("QGraphicsOpacityEffect::QGraphicsOpacityEffect")

	return NewQGraphicsOpacityEffectFromPointer(C.QGraphicsOpacityEffect_NewQGraphicsOpacityEffect(core.PointerFromQObject(parent)))
}

func (ptr *QGraphicsOpacityEffect) ConnectOpacityChanged(f func(opacity float64)) {
	defer qt.Recovering("connect QGraphicsOpacityEffect::opacityChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsOpacityEffect_ConnectOpacityChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "opacityChanged", f)
	}
}

func (ptr *QGraphicsOpacityEffect) DisconnectOpacityChanged() {
	defer qt.Recovering("disconnect QGraphicsOpacityEffect::opacityChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsOpacityEffect_DisconnectOpacityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "opacityChanged")
	}
}

//export callbackQGraphicsOpacityEffectOpacityChanged
func callbackQGraphicsOpacityEffectOpacityChanged(ptrName *C.char, opacity C.double) {
	defer qt.Recovering("callback QGraphicsOpacityEffect::opacityChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "opacityChanged"); signal != nil {
		signal.(func(float64))(float64(opacity))
	}

}

func (ptr *QGraphicsOpacityEffect) ConnectOpacityMaskChanged(f func(mask *gui.QBrush)) {
	defer qt.Recovering("connect QGraphicsOpacityEffect::opacityMaskChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsOpacityEffect_ConnectOpacityMaskChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "opacityMaskChanged", f)
	}
}

func (ptr *QGraphicsOpacityEffect) DisconnectOpacityMaskChanged() {
	defer qt.Recovering("disconnect QGraphicsOpacityEffect::opacityMaskChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsOpacityEffect_DisconnectOpacityMaskChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "opacityMaskChanged")
	}
}

//export callbackQGraphicsOpacityEffectOpacityMaskChanged
func callbackQGraphicsOpacityEffectOpacityMaskChanged(ptrName *C.char, mask unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsOpacityEffect::opacityMaskChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "opacityMaskChanged"); signal != nil {
		signal.(func(*gui.QBrush))(gui.NewQBrushFromPointer(mask))
	}

}

func (ptr *QGraphicsOpacityEffect) DestroyQGraphicsOpacityEffect() {
	defer qt.Recovering("QGraphicsOpacityEffect::~QGraphicsOpacityEffect")

	if ptr.Pointer() != nil {
		C.QGraphicsOpacityEffect_DestroyQGraphicsOpacityEffect(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsOpacityEffect) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGraphicsOpacityEffect::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGraphicsOpacityEffect) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGraphicsOpacityEffect::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGraphicsOpacityEffectTimerEvent
func callbackQGraphicsOpacityEffectTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsOpacityEffect::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsOpacityEffect) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGraphicsOpacityEffect::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGraphicsOpacityEffect) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGraphicsOpacityEffect::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGraphicsOpacityEffectChildEvent
func callbackQGraphicsOpacityEffectChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsOpacityEffect::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsOpacityEffect) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsOpacityEffect::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGraphicsOpacityEffect) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGraphicsOpacityEffect::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGraphicsOpacityEffectCustomEvent
func callbackQGraphicsOpacityEffectCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsOpacityEffect::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
