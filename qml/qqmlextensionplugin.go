package qml

//#include "qml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QQmlExtensionPlugin struct {
	core.QObject
}

type QQmlExtensionPlugin_ITF interface {
	core.QObject_ITF
	QQmlExtensionPlugin_PTR() *QQmlExtensionPlugin
}

func PointerFromQQmlExtensionPlugin(ptr QQmlExtensionPlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlExtensionPlugin_PTR().Pointer()
	}
	return nil
}

func NewQQmlExtensionPluginFromPointer(ptr unsafe.Pointer) *QQmlExtensionPlugin {
	var n = new(QQmlExtensionPlugin)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QQmlExtensionPlugin_") {
		n.SetObjectName("QQmlExtensionPlugin_" + qt.Identifier())
	}
	return n
}

func (ptr *QQmlExtensionPlugin) QQmlExtensionPlugin_PTR() *QQmlExtensionPlugin {
	return ptr
}

func (ptr *QQmlExtensionPlugin) BaseUrl() *core.QUrl {
	defer qt.Recovering("QQmlExtensionPlugin::baseUrl")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QQmlExtensionPlugin_BaseUrl(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlExtensionPlugin) RegisterTypes(uri string) {
	defer qt.Recovering("QQmlExtensionPlugin::registerTypes")

	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_RegisterTypes(ptr.Pointer(), C.CString(uri))
	}
}

func (ptr *QQmlExtensionPlugin) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQmlExtensionPlugin::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQmlExtensionPlugin::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQmlExtensionPluginTimerEvent
func callbackQQmlExtensionPluginTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlExtensionPlugin::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlExtensionPluginFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlExtensionPlugin) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlExtensionPlugin::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlExtensionPlugin) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlExtensionPlugin::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlExtensionPlugin) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQmlExtensionPlugin::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQmlExtensionPlugin::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQmlExtensionPluginChildEvent
func callbackQQmlExtensionPluginChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlExtensionPlugin::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlExtensionPluginFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlExtensionPlugin) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlExtensionPlugin::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlExtensionPlugin) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlExtensionPlugin::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlExtensionPlugin) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQmlExtensionPlugin::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQmlExtensionPlugin::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQmlExtensionPluginCustomEvent
func callbackQQmlExtensionPluginCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlExtensionPlugin::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlExtensionPluginFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlExtensionPlugin) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlExtensionPlugin::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQmlExtensionPlugin) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlExtensionPlugin::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
