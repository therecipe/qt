package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSGDynamicTexture struct {
	QSGTexture
}

type QSGDynamicTexture_ITF interface {
	QSGTexture_ITF
	QSGDynamicTexture_PTR() *QSGDynamicTexture
}

func PointerFromQSGDynamicTexture(ptr QSGDynamicTexture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGDynamicTexture_PTR().Pointer()
	}
	return nil
}

func NewQSGDynamicTextureFromPointer(ptr unsafe.Pointer) *QSGDynamicTexture {
	var n = new(QSGDynamicTexture)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSGDynamicTexture_") {
		n.SetObjectName("QSGDynamicTexture_" + qt.Identifier())
	}
	return n
}

func (ptr *QSGDynamicTexture) QSGDynamicTexture_PTR() *QSGDynamicTexture {
	return ptr
}

func (ptr *QSGDynamicTexture) UpdateTexture() bool {
	defer qt.Recovering("QSGDynamicTexture::updateTexture")

	if ptr.Pointer() != nil {
		return C.QSGDynamicTexture_UpdateTexture(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGDynamicTexture) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSGDynamicTexture::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSGDynamicTexture::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSGDynamicTextureTimerEvent
func callbackQSGDynamicTextureTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGDynamicTexture::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSGDynamicTextureFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSGDynamicTexture) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSGDynamicTexture::timerEvent")

	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSGDynamicTexture) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSGDynamicTexture::timerEvent")

	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSGDynamicTexture) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSGDynamicTexture::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSGDynamicTexture::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSGDynamicTextureChildEvent
func callbackQSGDynamicTextureChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGDynamicTexture::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSGDynamicTextureFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSGDynamicTexture) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSGDynamicTexture::childEvent")

	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSGDynamicTexture) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSGDynamicTexture::childEvent")

	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSGDynamicTexture) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSGDynamicTexture::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSGDynamicTexture::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSGDynamicTextureCustomEvent
func callbackQSGDynamicTextureCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGDynamicTexture::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSGDynamicTextureFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSGDynamicTexture) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSGDynamicTexture::customEvent")

	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSGDynamicTexture) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSGDynamicTexture::customEvent")

	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
