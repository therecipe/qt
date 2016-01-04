package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QMetaDataWriterControl struct {
	QMediaControl
}

type QMetaDataWriterControl_ITF interface {
	QMediaControl_ITF
	QMetaDataWriterControl_PTR() *QMetaDataWriterControl
}

func PointerFromQMetaDataWriterControl(ptr QMetaDataWriterControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMetaDataWriterControl_PTR().Pointer()
	}
	return nil
}

func NewQMetaDataWriterControlFromPointer(ptr unsafe.Pointer) *QMetaDataWriterControl {
	var n = new(QMetaDataWriterControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMetaDataWriterControl_") {
		n.SetObjectName("QMetaDataWriterControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QMetaDataWriterControl) QMetaDataWriterControl_PTR() *QMetaDataWriterControl {
	return ptr
}

func (ptr *QMetaDataWriterControl) AvailableMetaData() []string {
	defer qt.Recovering("QMetaDataWriterControl::availableMetaData")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMetaDataWriterControl_AvailableMetaData(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QMetaDataWriterControl) IsMetaDataAvailable() bool {
	defer qt.Recovering("QMetaDataWriterControl::isMetaDataAvailable")

	if ptr.Pointer() != nil {
		return C.QMetaDataWriterControl_IsMetaDataAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMetaDataWriterControl) IsWritable() bool {
	defer qt.Recovering("QMetaDataWriterControl::isWritable")

	if ptr.Pointer() != nil {
		return C.QMetaDataWriterControl_IsWritable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMetaDataWriterControl) MetaData(key string) *core.QVariant {
	defer qt.Recovering("QMetaDataWriterControl::metaData")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QMetaDataWriterControl_MetaData(ptr.Pointer(), C.CString(key)))
	}
	return nil
}

func (ptr *QMetaDataWriterControl) ConnectMetaDataAvailableChanged(f func(available bool)) {
	defer qt.Recovering("connect QMetaDataWriterControl::metaDataAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_ConnectMetaDataAvailableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataAvailableChanged", f)
	}
}

func (ptr *QMetaDataWriterControl) DisconnectMetaDataAvailableChanged() {
	defer qt.Recovering("disconnect QMetaDataWriterControl::metaDataAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_DisconnectMetaDataAvailableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataAvailableChanged")
	}
}

//export callbackQMetaDataWriterControlMetaDataAvailableChanged
func callbackQMetaDataWriterControlMetaDataAvailableChanged(ptr unsafe.Pointer, ptrName *C.char, available C.int) {
	defer qt.Recovering("callback QMetaDataWriterControl::metaDataAvailableChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaDataAvailableChanged"); signal != nil {
		signal.(func(bool))(int(available) != 0)
	}

}

func (ptr *QMetaDataWriterControl) MetaDataAvailableChanged(available bool) {
	defer qt.Recovering("QMetaDataWriterControl::metaDataAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_MetaDataAvailableChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(available)))
	}
}

func (ptr *QMetaDataWriterControl) ConnectMetaDataChanged(f func()) {
	defer qt.Recovering("connect QMetaDataWriterControl::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_ConnectMetaDataChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataChanged", f)
	}
}

func (ptr *QMetaDataWriterControl) DisconnectMetaDataChanged() {
	defer qt.Recovering("disconnect QMetaDataWriterControl::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_DisconnectMetaDataChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataChanged")
	}
}

//export callbackQMetaDataWriterControlMetaDataChanged
func callbackQMetaDataWriterControlMetaDataChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QMetaDataWriterControl::metaDataChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaDataChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QMetaDataWriterControl) MetaDataChanged() {
	defer qt.Recovering("QMetaDataWriterControl::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_MetaDataChanged(ptr.Pointer())
	}
}

func (ptr *QMetaDataWriterControl) ConnectMetaDataChanged2(f func(key string, value *core.QVariant)) {
	defer qt.Recovering("connect QMetaDataWriterControl::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_ConnectMetaDataChanged2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataChanged2", f)
	}
}

func (ptr *QMetaDataWriterControl) DisconnectMetaDataChanged2() {
	defer qt.Recovering("disconnect QMetaDataWriterControl::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_DisconnectMetaDataChanged2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataChanged2")
	}
}

//export callbackQMetaDataWriterControlMetaDataChanged2
func callbackQMetaDataWriterControlMetaDataChanged2(ptr unsafe.Pointer, ptrName *C.char, key *C.char, value unsafe.Pointer) {
	defer qt.Recovering("callback QMetaDataWriterControl::metaDataChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaDataChanged2"); signal != nil {
		signal.(func(string, *core.QVariant))(C.GoString(key), core.NewQVariantFromPointer(value))
	}

}

func (ptr *QMetaDataWriterControl) MetaDataChanged2(key string, value core.QVariant_ITF) {
	defer qt.Recovering("QMetaDataWriterControl::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_MetaDataChanged2(ptr.Pointer(), C.CString(key), core.PointerFromQVariant(value))
	}
}

func (ptr *QMetaDataWriterControl) SetMetaData(key string, value core.QVariant_ITF) {
	defer qt.Recovering("QMetaDataWriterControl::setMetaData")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_SetMetaData(ptr.Pointer(), C.CString(key), core.PointerFromQVariant(value))
	}
}

func (ptr *QMetaDataWriterControl) ConnectWritableChanged(f func(writable bool)) {
	defer qt.Recovering("connect QMetaDataWriterControl::writableChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_ConnectWritableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "writableChanged", f)
	}
}

func (ptr *QMetaDataWriterControl) DisconnectWritableChanged() {
	defer qt.Recovering("disconnect QMetaDataWriterControl::writableChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_DisconnectWritableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "writableChanged")
	}
}

//export callbackQMetaDataWriterControlWritableChanged
func callbackQMetaDataWriterControlWritableChanged(ptr unsafe.Pointer, ptrName *C.char, writable C.int) {
	defer qt.Recovering("callback QMetaDataWriterControl::writableChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "writableChanged"); signal != nil {
		signal.(func(bool))(int(writable) != 0)
	}

}

func (ptr *QMetaDataWriterControl) WritableChanged(writable bool) {
	defer qt.Recovering("QMetaDataWriterControl::writableChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_WritableChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(writable)))
	}
}

func (ptr *QMetaDataWriterControl) DestroyQMetaDataWriterControl() {
	defer qt.Recovering("QMetaDataWriterControl::~QMetaDataWriterControl")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_DestroyQMetaDataWriterControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMetaDataWriterControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMetaDataWriterControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMetaDataWriterControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMetaDataWriterControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMetaDataWriterControlTimerEvent
func callbackQMetaDataWriterControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMetaDataWriterControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMetaDataWriterControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMetaDataWriterControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMetaDataWriterControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMetaDataWriterControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMetaDataWriterControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMetaDataWriterControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMetaDataWriterControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMetaDataWriterControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMetaDataWriterControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMetaDataWriterControlChildEvent
func callbackQMetaDataWriterControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMetaDataWriterControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMetaDataWriterControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMetaDataWriterControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMetaDataWriterControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMetaDataWriterControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMetaDataWriterControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMetaDataWriterControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMetaDataWriterControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMetaDataWriterControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMetaDataWriterControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMetaDataWriterControlCustomEvent
func callbackQMetaDataWriterControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMetaDataWriterControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMetaDataWriterControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMetaDataWriterControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMetaDataWriterControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMetaDataWriterControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMetaDataWriterControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
