package qml

//#include "qml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QQmlPropertyMap struct {
	core.QObject
}

type QQmlPropertyMap_ITF interface {
	core.QObject_ITF
	QQmlPropertyMap_PTR() *QQmlPropertyMap
}

func PointerFromQQmlPropertyMap(ptr QQmlPropertyMap_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlPropertyMap_PTR().Pointer()
	}
	return nil
}

func NewQQmlPropertyMapFromPointer(ptr unsafe.Pointer) *QQmlPropertyMap {
	var n = new(QQmlPropertyMap)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QQmlPropertyMap_") {
		n.SetObjectName("QQmlPropertyMap_" + qt.Identifier())
	}
	return n
}

func (ptr *QQmlPropertyMap) QQmlPropertyMap_PTR() *QQmlPropertyMap {
	return ptr
}

func NewQQmlPropertyMap(parent core.QObject_ITF) *QQmlPropertyMap {
	defer qt.Recovering("QQmlPropertyMap::QQmlPropertyMap")

	return NewQQmlPropertyMapFromPointer(C.QQmlPropertyMap_NewQQmlPropertyMap(core.PointerFromQObject(parent)))
}

func (ptr *QQmlPropertyMap) Clear(key string) {
	defer qt.Recovering("QQmlPropertyMap::clear")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_Clear(ptr.Pointer(), C.CString(key))
	}
}

func (ptr *QQmlPropertyMap) Contains(key string) bool {
	defer qt.Recovering("QQmlPropertyMap::contains")

	if ptr.Pointer() != nil {
		return C.QQmlPropertyMap_Contains(ptr.Pointer(), C.CString(key)) != 0
	}
	return false
}

func (ptr *QQmlPropertyMap) Count() int {
	defer qt.Recovering("QQmlPropertyMap::count")

	if ptr.Pointer() != nil {
		return int(C.QQmlPropertyMap_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlPropertyMap) IsEmpty() bool {
	defer qt.Recovering("QQmlPropertyMap::isEmpty")

	if ptr.Pointer() != nil {
		return C.QQmlPropertyMap_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlPropertyMap) Keys() []string {
	defer qt.Recovering("QQmlPropertyMap::keys")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QQmlPropertyMap_Keys(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QQmlPropertyMap) Size() int {
	defer qt.Recovering("QQmlPropertyMap::size")

	if ptr.Pointer() != nil {
		return int(C.QQmlPropertyMap_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlPropertyMap) UpdateValue(key string, input core.QVariant_ITF) *core.QVariant {
	defer qt.Recovering("QQmlPropertyMap::updateValue")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QQmlPropertyMap_UpdateValue(ptr.Pointer(), C.CString(key), core.PointerFromQVariant(input)))
	}
	return nil
}

func (ptr *QQmlPropertyMap) Value(key string) *core.QVariant {
	defer qt.Recovering("QQmlPropertyMap::value")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QQmlPropertyMap_Value(ptr.Pointer(), C.CString(key)))
	}
	return nil
}

func (ptr *QQmlPropertyMap) ConnectValueChanged(f func(key string, value *core.QVariant)) {
	defer qt.Recovering("connect QQmlPropertyMap::valueChanged")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_ConnectValueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "valueChanged", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectValueChanged() {
	defer qt.Recovering("disconnect QQmlPropertyMap::valueChanged")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "valueChanged")
	}
}

//export callbackQQmlPropertyMapValueChanged
func callbackQQmlPropertyMapValueChanged(ptr unsafe.Pointer, ptrName *C.char, key *C.char, value unsafe.Pointer) {
	defer qt.Recovering("callback QQmlPropertyMap::valueChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "valueChanged"); signal != nil {
		signal.(func(string, *core.QVariant))(C.GoString(key), core.NewQVariantFromPointer(value))
	}

}

func (ptr *QQmlPropertyMap) ValueChanged(key string, value core.QVariant_ITF) {
	defer qt.Recovering("QQmlPropertyMap::valueChanged")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_ValueChanged(ptr.Pointer(), C.CString(key), core.PointerFromQVariant(value))
	}
}

func (ptr *QQmlPropertyMap) DestroyQQmlPropertyMap() {
	defer qt.Recovering("QQmlPropertyMap::~QQmlPropertyMap")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_DestroyQQmlPropertyMap(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlPropertyMap) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQmlPropertyMap::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQmlPropertyMap::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQmlPropertyMapTimerEvent
func callbackQQmlPropertyMapTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlPropertyMap::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlPropertyMapFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlPropertyMap) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlPropertyMap::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlPropertyMap) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlPropertyMap::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlPropertyMap) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQmlPropertyMap::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQmlPropertyMap::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQmlPropertyMapChildEvent
func callbackQQmlPropertyMapChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlPropertyMap::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlPropertyMapFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlPropertyMap) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlPropertyMap::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlPropertyMap) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlPropertyMap::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlPropertyMap) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQmlPropertyMap::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQmlPropertyMap::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQmlPropertyMapCustomEvent
func callbackQQmlPropertyMapCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlPropertyMap::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlPropertyMapFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlPropertyMap) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlPropertyMap::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQmlPropertyMap) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlPropertyMap::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
