package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QStylePlugin struct {
	core.QObject
}

type QStylePlugin_ITF interface {
	core.QObject_ITF
	QStylePlugin_PTR() *QStylePlugin
}

func PointerFromQStylePlugin(ptr QStylePlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStylePlugin_PTR().Pointer()
	}
	return nil
}

func NewQStylePluginFromPointer(ptr unsafe.Pointer) *QStylePlugin {
	var n = new(QStylePlugin)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QStylePlugin_") {
		n.SetObjectName("QStylePlugin_" + qt.Identifier())
	}
	return n
}

func (ptr *QStylePlugin) QStylePlugin_PTR() *QStylePlugin {
	return ptr
}

func (ptr *QStylePlugin) Create(key string) *QStyle {
	defer qt.Recovering("QStylePlugin::create")

	if ptr.Pointer() != nil {
		return NewQStyleFromPointer(C.QStylePlugin_Create(ptr.Pointer(), C.CString(key)))
	}
	return nil
}

func (ptr *QStylePlugin) DestroyQStylePlugin() {
	defer qt.Recovering("QStylePlugin::~QStylePlugin")

	if ptr.Pointer() != nil {
		C.QStylePlugin_DestroyQStylePlugin(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QStylePlugin) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QStylePlugin::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QStylePlugin) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QStylePlugin::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQStylePluginTimerEvent
func callbackQStylePluginTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStylePlugin::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStylePlugin) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QStylePlugin::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QStylePlugin) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QStylePlugin::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQStylePluginChildEvent
func callbackQStylePluginChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStylePlugin::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStylePlugin) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QStylePlugin::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QStylePlugin) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QStylePlugin::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQStylePluginCustomEvent
func callbackQStylePluginCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStylePlugin::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
