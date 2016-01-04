package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QQuickTextureFactory struct {
	core.QObject
}

type QQuickTextureFactory_ITF interface {
	core.QObject_ITF
	QQuickTextureFactory_PTR() *QQuickTextureFactory
}

func PointerFromQQuickTextureFactory(ptr QQuickTextureFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickTextureFactory_PTR().Pointer()
	}
	return nil
}

func NewQQuickTextureFactoryFromPointer(ptr unsafe.Pointer) *QQuickTextureFactory {
	var n = new(QQuickTextureFactory)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QQuickTextureFactory_") {
		n.SetObjectName("QQuickTextureFactory_" + qt.Identifier())
	}
	return n
}

func (ptr *QQuickTextureFactory) QQuickTextureFactory_PTR() *QQuickTextureFactory {
	return ptr
}

func (ptr *QQuickTextureFactory) CreateTexture(window QQuickWindow_ITF) *QSGTexture {
	defer qt.Recovering("QQuickTextureFactory::createTexture")

	if ptr.Pointer() != nil {
		return NewQSGTextureFromPointer(C.QQuickTextureFactory_CreateTexture(ptr.Pointer(), PointerFromQQuickWindow(window)))
	}
	return nil
}

func (ptr *QQuickTextureFactory) TextureByteCount() int {
	defer qt.Recovering("QQuickTextureFactory::textureByteCount")

	if ptr.Pointer() != nil {
		return int(C.QQuickTextureFactory_TextureByteCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickTextureFactory) TextureSize() *core.QSize {
	defer qt.Recovering("QQuickTextureFactory::textureSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QQuickTextureFactory_TextureSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickTextureFactory) DestroyQQuickTextureFactory() {
	defer qt.Recovering("QQuickTextureFactory::~QQuickTextureFactory")

	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_DestroyQQuickTextureFactory(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickTextureFactory) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQuickTextureFactory::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQuickTextureFactory) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQuickTextureFactory::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQuickTextureFactoryTimerEvent
func callbackQQuickTextureFactoryTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickTextureFactory::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickTextureFactoryFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickTextureFactory) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickTextureFactory::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickTextureFactory) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickTextureFactory::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickTextureFactory) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQuickTextureFactory::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQuickTextureFactory) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQuickTextureFactory::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQuickTextureFactoryChildEvent
func callbackQQuickTextureFactoryChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickTextureFactory::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickTextureFactoryFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickTextureFactory) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickTextureFactory::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickTextureFactory) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickTextureFactory::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickTextureFactory) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQuickTextureFactory::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQuickTextureFactory) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQuickTextureFactory::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQuickTextureFactoryCustomEvent
func callbackQQuickTextureFactoryCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickTextureFactory::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickTextureFactoryFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickTextureFactory) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickTextureFactory::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickTextureFactory) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickTextureFactory::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
