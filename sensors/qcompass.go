package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QCompass struct {
	QSensor
}

type QCompass_ITF interface {
	QSensor_ITF
	QCompass_PTR() *QCompass
}

func PointerFromQCompass(ptr QCompass_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCompass_PTR().Pointer()
	}
	return nil
}

func NewQCompassFromPointer(ptr unsafe.Pointer) *QCompass {
	var n = new(QCompass)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCompass_") {
		n.SetObjectName("QCompass_" + qt.Identifier())
	}
	return n
}

func (ptr *QCompass) QCompass_PTR() *QCompass {
	return ptr
}

func (ptr *QCompass) Reading() *QCompassReading {
	defer qt.Recovering("QCompass::reading")

	if ptr.Pointer() != nil {
		return NewQCompassReadingFromPointer(C.QCompass_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQCompass(parent core.QObject_ITF) *QCompass {
	defer qt.Recovering("QCompass::QCompass")

	return NewQCompassFromPointer(C.QCompass_NewQCompass(core.PointerFromQObject(parent)))
}

func (ptr *QCompass) DestroyQCompass() {
	defer qt.Recovering("QCompass::~QCompass")

	if ptr.Pointer() != nil {
		C.QCompass_DestroyQCompass(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCompass) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCompass::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCompass) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCompass::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCompassTimerEvent
func callbackQCompassTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCompass::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCompass) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCompass::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCompass) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCompass::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCompassChildEvent
func callbackQCompassChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCompass::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCompass) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCompass::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCompass) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCompass::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCompassCustomEvent
func callbackQCompassCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCompass::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
