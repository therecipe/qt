package script

//#include "script.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QScriptExtensionPlugin struct {
	core.QObject
}

type QScriptExtensionPlugin_ITF interface {
	core.QObject_ITF
	QScriptExtensionPlugin_PTR() *QScriptExtensionPlugin
}

func PointerFromQScriptExtensionPlugin(ptr QScriptExtensionPlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptExtensionPlugin_PTR().Pointer()
	}
	return nil
}

func NewQScriptExtensionPluginFromPointer(ptr unsafe.Pointer) *QScriptExtensionPlugin {
	var n = new(QScriptExtensionPlugin)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QScriptExtensionPlugin_") {
		n.SetObjectName("QScriptExtensionPlugin_" + qt.Identifier())
	}
	return n
}

func (ptr *QScriptExtensionPlugin) QScriptExtensionPlugin_PTR() *QScriptExtensionPlugin {
	return ptr
}

func (ptr *QScriptExtensionPlugin) Initialize(key string, engine QScriptEngine_ITF) {
	defer qt.Recovering("QScriptExtensionPlugin::initialize")

	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_Initialize(ptr.Pointer(), C.CString(key), PointerFromQScriptEngine(engine))
	}
}

func (ptr *QScriptExtensionPlugin) Keys() []string {
	defer qt.Recovering("QScriptExtensionPlugin::keys")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QScriptExtensionPlugin_Keys(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QScriptExtensionPlugin) SetupPackage(key string, engine QScriptEngine_ITF) *QScriptValue {
	defer qt.Recovering("QScriptExtensionPlugin::setupPackage")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptExtensionPlugin_SetupPackage(ptr.Pointer(), C.CString(key), PointerFromQScriptEngine(engine)))
	}
	return nil
}

func (ptr *QScriptExtensionPlugin) DestroyQScriptExtensionPlugin() {
	defer qt.Recovering("QScriptExtensionPlugin::~QScriptExtensionPlugin")

	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_DestroyQScriptExtensionPlugin(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScriptExtensionPlugin) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QScriptExtensionPlugin::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QScriptExtensionPlugin) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QScriptExtensionPlugin::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQScriptExtensionPluginTimerEvent
func callbackQScriptExtensionPluginTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScriptExtensionPlugin::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QScriptExtensionPlugin) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QScriptExtensionPlugin::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QScriptExtensionPlugin) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QScriptExtensionPlugin::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQScriptExtensionPluginChildEvent
func callbackQScriptExtensionPluginChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScriptExtensionPlugin::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QScriptExtensionPlugin) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QScriptExtensionPlugin::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QScriptExtensionPlugin) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QScriptExtensionPlugin::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQScriptExtensionPluginCustomEvent
func callbackQScriptExtensionPluginCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScriptExtensionPlugin::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
