package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSGTextureProvider struct {
	core.QObject
}

type QSGTextureProvider_ITF interface {
	core.QObject_ITF
	QSGTextureProvider_PTR() *QSGTextureProvider
}

func PointerFromQSGTextureProvider(ptr QSGTextureProvider_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGTextureProvider_PTR().Pointer()
	}
	return nil
}

func NewQSGTextureProviderFromPointer(ptr unsafe.Pointer) *QSGTextureProvider {
	var n = new(QSGTextureProvider)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSGTextureProvider_") {
		n.SetObjectName("QSGTextureProvider_" + qt.Identifier())
	}
	return n
}

func (ptr *QSGTextureProvider) QSGTextureProvider_PTR() *QSGTextureProvider {
	return ptr
}

func (ptr *QSGTextureProvider) Texture() *QSGTexture {
	defer qt.Recovering("QSGTextureProvider::texture")

	if ptr.Pointer() != nil {
		return NewQSGTextureFromPointer(C.QSGTextureProvider_Texture(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGTextureProvider) ConnectTextureChanged(f func()) {
	defer qt.Recovering("connect QSGTextureProvider::textureChanged")

	if ptr.Pointer() != nil {
		C.QSGTextureProvider_ConnectTextureChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "textureChanged", f)
	}
}

func (ptr *QSGTextureProvider) DisconnectTextureChanged() {
	defer qt.Recovering("disconnect QSGTextureProvider::textureChanged")

	if ptr.Pointer() != nil {
		C.QSGTextureProvider_DisconnectTextureChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "textureChanged")
	}
}

//export callbackQSGTextureProviderTextureChanged
func callbackQSGTextureProviderTextureChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGTextureProvider::textureChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "textureChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSGTextureProvider) TextureChanged() {
	defer qt.Recovering("QSGTextureProvider::textureChanged")

	if ptr.Pointer() != nil {
		C.QSGTextureProvider_TextureChanged(ptr.Pointer())
	}
}

func (ptr *QSGTextureProvider) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSGTextureProvider::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSGTextureProvider) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSGTextureProvider::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSGTextureProviderTimerEvent
func callbackQSGTextureProviderTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGTextureProvider::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSGTextureProviderFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSGTextureProvider) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSGTextureProvider::timerEvent")

	if ptr.Pointer() != nil {
		C.QSGTextureProvider_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSGTextureProvider) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSGTextureProvider::timerEvent")

	if ptr.Pointer() != nil {
		C.QSGTextureProvider_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSGTextureProvider) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSGTextureProvider::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSGTextureProvider) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSGTextureProvider::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSGTextureProviderChildEvent
func callbackQSGTextureProviderChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGTextureProvider::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSGTextureProviderFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSGTextureProvider) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSGTextureProvider::childEvent")

	if ptr.Pointer() != nil {
		C.QSGTextureProvider_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSGTextureProvider) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSGTextureProvider::childEvent")

	if ptr.Pointer() != nil {
		C.QSGTextureProvider_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSGTextureProvider) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSGTextureProvider::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSGTextureProvider) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSGTextureProvider::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSGTextureProviderCustomEvent
func callbackQSGTextureProviderCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGTextureProvider::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSGTextureProviderFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSGTextureProvider) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSGTextureProvider::customEvent")

	if ptr.Pointer() != nil {
		C.QSGTextureProvider_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSGTextureProvider) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSGTextureProvider::customEvent")

	if ptr.Pointer() != nil {
		C.QSGTextureProvider_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
