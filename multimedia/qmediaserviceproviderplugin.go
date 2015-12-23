package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMediaServiceProviderPlugin struct {
	core.QObject
}

type QMediaServiceProviderPlugin_ITF interface {
	core.QObject_ITF
	QMediaServiceProviderPlugin_PTR() *QMediaServiceProviderPlugin
}

func PointerFromQMediaServiceProviderPlugin(ptr QMediaServiceProviderPlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaServiceProviderPlugin_PTR().Pointer()
	}
	return nil
}

func NewQMediaServiceProviderPluginFromPointer(ptr unsafe.Pointer) *QMediaServiceProviderPlugin {
	var n = new(QMediaServiceProviderPlugin)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMediaServiceProviderPlugin_") {
		n.SetObjectName("QMediaServiceProviderPlugin_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaServiceProviderPlugin) QMediaServiceProviderPlugin_PTR() *QMediaServiceProviderPlugin {
	return ptr
}

func (ptr *QMediaServiceProviderPlugin) Create(key string) *QMediaService {
	defer qt.Recovering("QMediaServiceProviderPlugin::create")

	if ptr.Pointer() != nil {
		return NewQMediaServiceFromPointer(C.QMediaServiceProviderPlugin_Create(ptr.Pointer(), C.CString(key)))
	}
	return nil
}

func (ptr *QMediaServiceProviderPlugin) Release(service QMediaService_ITF) {
	defer qt.Recovering("QMediaServiceProviderPlugin::release")

	if ptr.Pointer() != nil {
		C.QMediaServiceProviderPlugin_Release(ptr.Pointer(), PointerFromQMediaService(service))
	}
}

func (ptr *QMediaServiceProviderPlugin) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMediaServiceProviderPlugin::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMediaServiceProviderPlugin) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMediaServiceProviderPlugin::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMediaServiceProviderPluginTimerEvent
func callbackQMediaServiceProviderPluginTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMediaServiceProviderPlugin::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMediaServiceProviderPlugin) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMediaServiceProviderPlugin::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMediaServiceProviderPlugin) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMediaServiceProviderPlugin::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMediaServiceProviderPluginChildEvent
func callbackQMediaServiceProviderPluginChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMediaServiceProviderPlugin::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMediaServiceProviderPlugin) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMediaServiceProviderPlugin::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMediaServiceProviderPlugin) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMediaServiceProviderPlugin::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMediaServiceProviderPluginCustomEvent
func callbackQMediaServiceProviderPluginCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMediaServiceProviderPlugin::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
