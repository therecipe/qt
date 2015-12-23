package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSignalMapper struct {
	QObject
}

type QSignalMapper_ITF interface {
	QObject_ITF
	QSignalMapper_PTR() *QSignalMapper
}

func PointerFromQSignalMapper(ptr QSignalMapper_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSignalMapper_PTR().Pointer()
	}
	return nil
}

func NewQSignalMapperFromPointer(ptr unsafe.Pointer) *QSignalMapper {
	var n = new(QSignalMapper)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSignalMapper_") {
		n.SetObjectName("QSignalMapper_" + qt.Identifier())
	}
	return n
}

func (ptr *QSignalMapper) QSignalMapper_PTR() *QSignalMapper {
	return ptr
}

func NewQSignalMapper(parent QObject_ITF) *QSignalMapper {
	defer qt.Recovering("QSignalMapper::QSignalMapper")

	return NewQSignalMapperFromPointer(C.QSignalMapper_NewQSignalMapper(PointerFromQObject(parent)))
}

func (ptr *QSignalMapper) Map() {
	defer qt.Recovering("QSignalMapper::map")

	if ptr.Pointer() != nil {
		C.QSignalMapper_Map(ptr.Pointer())
	}
}

func (ptr *QSignalMapper) Map2(sender QObject_ITF) {
	defer qt.Recovering("QSignalMapper::map")

	if ptr.Pointer() != nil {
		C.QSignalMapper_Map2(ptr.Pointer(), PointerFromQObject(sender))
	}
}

func (ptr *QSignalMapper) ConnectMapped4(f func(object *QObject)) {
	defer qt.Recovering("connect QSignalMapper::mapped")

	if ptr.Pointer() != nil {
		C.QSignalMapper_ConnectMapped4(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mapped4", f)
	}
}

func (ptr *QSignalMapper) DisconnectMapped4() {
	defer qt.Recovering("disconnect QSignalMapper::mapped")

	if ptr.Pointer() != nil {
		C.QSignalMapper_DisconnectMapped4(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mapped4")
	}
}

//export callbackQSignalMapperMapped4
func callbackQSignalMapperMapped4(ptrName *C.char, object unsafe.Pointer) {
	defer qt.Recovering("callback QSignalMapper::mapped")

	if signal := qt.GetSignal(C.GoString(ptrName), "mapped4"); signal != nil {
		signal.(func(*QObject))(NewQObjectFromPointer(object))
	}

}

func (ptr *QSignalMapper) ConnectMapped3(f func(widget unsafe.Pointer)) {
	defer qt.Recovering("connect QSignalMapper::mapped")

	if ptr.Pointer() != nil {
		C.QSignalMapper_ConnectMapped3(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mapped3", f)
	}
}

func (ptr *QSignalMapper) DisconnectMapped3() {
	defer qt.Recovering("disconnect QSignalMapper::mapped")

	if ptr.Pointer() != nil {
		C.QSignalMapper_DisconnectMapped3(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mapped3")
	}
}

//export callbackQSignalMapperMapped3
func callbackQSignalMapperMapped3(ptrName *C.char, widget unsafe.Pointer) {
	defer qt.Recovering("callback QSignalMapper::mapped")

	if signal := qt.GetSignal(C.GoString(ptrName), "mapped3"); signal != nil {
		signal.(func(unsafe.Pointer))(unsafe.Pointer(widget))
	}

}

func (ptr *QSignalMapper) ConnectMapped2(f func(text string)) {
	defer qt.Recovering("connect QSignalMapper::mapped")

	if ptr.Pointer() != nil {
		C.QSignalMapper_ConnectMapped2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mapped2", f)
	}
}

func (ptr *QSignalMapper) DisconnectMapped2() {
	defer qt.Recovering("disconnect QSignalMapper::mapped")

	if ptr.Pointer() != nil {
		C.QSignalMapper_DisconnectMapped2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mapped2")
	}
}

//export callbackQSignalMapperMapped2
func callbackQSignalMapperMapped2(ptrName *C.char, text *C.char) {
	defer qt.Recovering("callback QSignalMapper::mapped")

	if signal := qt.GetSignal(C.GoString(ptrName), "mapped2"); signal != nil {
		signal.(func(string))(C.GoString(text))
	}

}

func (ptr *QSignalMapper) ConnectMapped(f func(i int)) {
	defer qt.Recovering("connect QSignalMapper::mapped")

	if ptr.Pointer() != nil {
		C.QSignalMapper_ConnectMapped(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mapped", f)
	}
}

func (ptr *QSignalMapper) DisconnectMapped() {
	defer qt.Recovering("disconnect QSignalMapper::mapped")

	if ptr.Pointer() != nil {
		C.QSignalMapper_DisconnectMapped(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mapped")
	}
}

//export callbackQSignalMapperMapped
func callbackQSignalMapperMapped(ptrName *C.char, i C.int) {
	defer qt.Recovering("callback QSignalMapper::mapped")

	if signal := qt.GetSignal(C.GoString(ptrName), "mapped"); signal != nil {
		signal.(func(int))(int(i))
	}

}

func (ptr *QSignalMapper) Mapping4(object QObject_ITF) *QObject {
	defer qt.Recovering("QSignalMapper::mapping")

	if ptr.Pointer() != nil {
		return NewQObjectFromPointer(C.QSignalMapper_Mapping4(ptr.Pointer(), PointerFromQObject(object)))
	}
	return nil
}

func (ptr *QSignalMapper) Mapping3(widget unsafe.Pointer) *QObject {
	defer qt.Recovering("QSignalMapper::mapping")

	if ptr.Pointer() != nil {
		return NewQObjectFromPointer(C.QSignalMapper_Mapping3(ptr.Pointer(), widget))
	}
	return nil
}

func (ptr *QSignalMapper) Mapping2(id string) *QObject {
	defer qt.Recovering("QSignalMapper::mapping")

	if ptr.Pointer() != nil {
		return NewQObjectFromPointer(C.QSignalMapper_Mapping2(ptr.Pointer(), C.CString(id)))
	}
	return nil
}

func (ptr *QSignalMapper) Mapping(id int) *QObject {
	defer qt.Recovering("QSignalMapper::mapping")

	if ptr.Pointer() != nil {
		return NewQObjectFromPointer(C.QSignalMapper_Mapping(ptr.Pointer(), C.int(id)))
	}
	return nil
}

func (ptr *QSignalMapper) RemoveMappings(sender QObject_ITF) {
	defer qt.Recovering("QSignalMapper::removeMappings")

	if ptr.Pointer() != nil {
		C.QSignalMapper_RemoveMappings(ptr.Pointer(), PointerFromQObject(sender))
	}
}

func (ptr *QSignalMapper) SetMapping4(sender QObject_ITF, object QObject_ITF) {
	defer qt.Recovering("QSignalMapper::setMapping")

	if ptr.Pointer() != nil {
		C.QSignalMapper_SetMapping4(ptr.Pointer(), PointerFromQObject(sender), PointerFromQObject(object))
	}
}

func (ptr *QSignalMapper) SetMapping3(sender QObject_ITF, widget unsafe.Pointer) {
	defer qt.Recovering("QSignalMapper::setMapping")

	if ptr.Pointer() != nil {
		C.QSignalMapper_SetMapping3(ptr.Pointer(), PointerFromQObject(sender), widget)
	}
}

func (ptr *QSignalMapper) SetMapping2(sender QObject_ITF, text string) {
	defer qt.Recovering("QSignalMapper::setMapping")

	if ptr.Pointer() != nil {
		C.QSignalMapper_SetMapping2(ptr.Pointer(), PointerFromQObject(sender), C.CString(text))
	}
}

func (ptr *QSignalMapper) SetMapping(sender QObject_ITF, id int) {
	defer qt.Recovering("QSignalMapper::setMapping")

	if ptr.Pointer() != nil {
		C.QSignalMapper_SetMapping(ptr.Pointer(), PointerFromQObject(sender), C.int(id))
	}
}

func (ptr *QSignalMapper) DestroyQSignalMapper() {
	defer qt.Recovering("QSignalMapper::~QSignalMapper")

	if ptr.Pointer() != nil {
		C.QSignalMapper_DestroyQSignalMapper(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSignalMapper) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QSignalMapper::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSignalMapper) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSignalMapper::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSignalMapperTimerEvent
func callbackQSignalMapperTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSignalMapper::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSignalMapper) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QSignalMapper::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSignalMapper) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSignalMapper::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSignalMapperChildEvent
func callbackQSignalMapperChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSignalMapper::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSignalMapper) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QSignalMapper::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSignalMapper) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSignalMapper::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSignalMapperCustomEvent
func callbackQSignalMapperCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSignalMapper::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
		return true
	}
	return false

}
