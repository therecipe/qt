package sql

//#include "sql.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSqlDriverPlugin struct {
	core.QObject
}

type QSqlDriverPlugin_ITF interface {
	core.QObject_ITF
	QSqlDriverPlugin_PTR() *QSqlDriverPlugin
}

func PointerFromQSqlDriverPlugin(ptr QSqlDriverPlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlDriverPlugin_PTR().Pointer()
	}
	return nil
}

func NewQSqlDriverPluginFromPointer(ptr unsafe.Pointer) *QSqlDriverPlugin {
	var n = new(QSqlDriverPlugin)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSqlDriverPlugin_") {
		n.SetObjectName("QSqlDriverPlugin_" + qt.Identifier())
	}
	return n
}

func (ptr *QSqlDriverPlugin) QSqlDriverPlugin_PTR() *QSqlDriverPlugin {
	return ptr
}

func (ptr *QSqlDriverPlugin) Create(key string) *QSqlDriver {
	defer qt.Recovering("QSqlDriverPlugin::create")

	if ptr.Pointer() != nil {
		return NewQSqlDriverFromPointer(C.QSqlDriverPlugin_Create(ptr.Pointer(), C.CString(key)))
	}
	return nil
}

func (ptr *QSqlDriverPlugin) DestroyQSqlDriverPlugin() {
	defer qt.Recovering("QSqlDriverPlugin::~QSqlDriverPlugin")

	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_DestroyQSqlDriverPlugin(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSqlDriverPlugin) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSqlDriverPlugin::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSqlDriverPlugin) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSqlDriverPlugin::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSqlDriverPluginTimerEvent
func callbackQSqlDriverPluginTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlDriverPlugin::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSqlDriverPluginFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSqlDriverPlugin) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSqlDriverPlugin::timerEvent")

	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSqlDriverPlugin) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSqlDriverPlugin::timerEvent")

	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSqlDriverPlugin) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSqlDriverPlugin::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSqlDriverPlugin) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSqlDriverPlugin::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSqlDriverPluginChildEvent
func callbackQSqlDriverPluginChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlDriverPlugin::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSqlDriverPluginFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSqlDriverPlugin) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSqlDriverPlugin::childEvent")

	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSqlDriverPlugin) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSqlDriverPlugin::childEvent")

	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSqlDriverPlugin) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSqlDriverPlugin::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSqlDriverPlugin) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSqlDriverPlugin::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSqlDriverPluginCustomEvent
func callbackQSqlDriverPluginCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlDriverPlugin::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSqlDriverPluginFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSqlDriverPlugin) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSqlDriverPlugin::customEvent")

	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSqlDriverPlugin) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSqlDriverPlugin::customEvent")

	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
