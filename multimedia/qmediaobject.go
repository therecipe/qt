package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QMediaObject struct {
	core.QObject
}

type QMediaObject_ITF interface {
	core.QObject_ITF
	QMediaObject_PTR() *QMediaObject
}

func PointerFromQMediaObject(ptr QMediaObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaObject_PTR().Pointer()
	}
	return nil
}

func NewQMediaObjectFromPointer(ptr unsafe.Pointer) *QMediaObject {
	var n = new(QMediaObject)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMediaObject_") {
		n.SetObjectName("QMediaObject_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaObject) QMediaObject_PTR() *QMediaObject {
	return ptr
}

func (ptr *QMediaObject) NotifyInterval() int {
	defer qt.Recovering("QMediaObject::notifyInterval")

	if ptr.Pointer() != nil {
		return int(C.QMediaObject_NotifyInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaObject) SetNotifyInterval(milliSeconds int) {
	defer qt.Recovering("QMediaObject::setNotifyInterval")

	if ptr.Pointer() != nil {
		C.QMediaObject_SetNotifyInterval(ptr.Pointer(), C.int(milliSeconds))
	}
}

func (ptr *QMediaObject) ConnectAvailabilityChanged(f func(available bool)) {
	defer qt.Recovering("connect QMediaObject::availabilityChanged")

	if ptr.Pointer() != nil {
		C.QMediaObject_ConnectAvailabilityChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "availabilityChanged", f)
	}
}

func (ptr *QMediaObject) DisconnectAvailabilityChanged() {
	defer qt.Recovering("disconnect QMediaObject::availabilityChanged")

	if ptr.Pointer() != nil {
		C.QMediaObject_DisconnectAvailabilityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "availabilityChanged")
	}
}

//export callbackQMediaObjectAvailabilityChanged
func callbackQMediaObjectAvailabilityChanged(ptrName *C.char, available C.int) {
	defer qt.Recovering("callback QMediaObject::availabilityChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "availabilityChanged")
	if signal != nil {
		signal.(func(bool))(int(available) != 0)
	}

}

func (ptr *QMediaObject) AvailableMetaData() []string {
	defer qt.Recovering("QMediaObject::availableMetaData")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMediaObject_AvailableMetaData(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QMediaObject) Bind(object core.QObject_ITF) bool {
	defer qt.Recovering("QMediaObject::bind")

	if ptr.Pointer() != nil {
		return C.QMediaObject_Bind(ptr.Pointer(), core.PointerFromQObject(object)) != 0
	}
	return false
}

func (ptr *QMediaObject) IsAvailable() bool {
	defer qt.Recovering("QMediaObject::isAvailable")

	if ptr.Pointer() != nil {
		return C.QMediaObject_IsAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaObject) IsMetaDataAvailable() bool {
	defer qt.Recovering("QMediaObject::isMetaDataAvailable")

	if ptr.Pointer() != nil {
		return C.QMediaObject_IsMetaDataAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaObject) MetaData(key string) *core.QVariant {
	defer qt.Recovering("QMediaObject::metaData")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QMediaObject_MetaData(ptr.Pointer(), C.CString(key)))
	}
	return nil
}

func (ptr *QMediaObject) ConnectMetaDataAvailableChanged(f func(available bool)) {
	defer qt.Recovering("connect QMediaObject::metaDataAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMediaObject_ConnectMetaDataAvailableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataAvailableChanged", f)
	}
}

func (ptr *QMediaObject) DisconnectMetaDataAvailableChanged() {
	defer qt.Recovering("disconnect QMediaObject::metaDataAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMediaObject_DisconnectMetaDataAvailableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataAvailableChanged")
	}
}

//export callbackQMediaObjectMetaDataAvailableChanged
func callbackQMediaObjectMetaDataAvailableChanged(ptrName *C.char, available C.int) {
	defer qt.Recovering("callback QMediaObject::metaDataAvailableChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "metaDataAvailableChanged")
	if signal != nil {
		signal.(func(bool))(int(available) != 0)
	}

}

func (ptr *QMediaObject) ConnectMetaDataChanged(f func()) {
	defer qt.Recovering("connect QMediaObject::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMediaObject_ConnectMetaDataChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataChanged", f)
	}
}

func (ptr *QMediaObject) DisconnectMetaDataChanged() {
	defer qt.Recovering("disconnect QMediaObject::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMediaObject_DisconnectMetaDataChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataChanged")
	}
}

//export callbackQMediaObjectMetaDataChanged
func callbackQMediaObjectMetaDataChanged(ptrName *C.char) {
	defer qt.Recovering("callback QMediaObject::metaDataChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "metaDataChanged")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QMediaObject) ConnectNotifyIntervalChanged(f func(milliseconds int)) {
	defer qt.Recovering("connect QMediaObject::notifyIntervalChanged")

	if ptr.Pointer() != nil {
		C.QMediaObject_ConnectNotifyIntervalChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "notifyIntervalChanged", f)
	}
}

func (ptr *QMediaObject) DisconnectNotifyIntervalChanged() {
	defer qt.Recovering("disconnect QMediaObject::notifyIntervalChanged")

	if ptr.Pointer() != nil {
		C.QMediaObject_DisconnectNotifyIntervalChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "notifyIntervalChanged")
	}
}

//export callbackQMediaObjectNotifyIntervalChanged
func callbackQMediaObjectNotifyIntervalChanged(ptrName *C.char, milliseconds C.int) {
	defer qt.Recovering("callback QMediaObject::notifyIntervalChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "notifyIntervalChanged")
	if signal != nil {
		signal.(func(int))(int(milliseconds))
	}

}

func (ptr *QMediaObject) Service() *QMediaService {
	defer qt.Recovering("QMediaObject::service")

	if ptr.Pointer() != nil {
		return NewQMediaServiceFromPointer(C.QMediaObject_Service(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaObject) ConnectUnbind(f func(object *core.QObject)) {
	defer qt.Recovering("connect QMediaObject::unbind")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "unbind", f)
	}
}

func (ptr *QMediaObject) DisconnectUnbind() {
	defer qt.Recovering("disconnect QMediaObject::unbind")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "unbind")
	}
}

//export callbackQMediaObjectUnbind
func callbackQMediaObjectUnbind(ptrName *C.char, object unsafe.Pointer) bool {
	defer qt.Recovering("callback QMediaObject::unbind")

	var signal = qt.GetSignal(C.GoString(ptrName), "unbind")
	if signal != nil {
		defer signal.(func(*core.QObject))(core.NewQObjectFromPointer(object))
		return true
	}
	return false

}

func (ptr *QMediaObject) DestroyQMediaObject() {
	defer qt.Recovering("QMediaObject::~QMediaObject")

	if ptr.Pointer() != nil {
		C.QMediaObject_DestroyQMediaObject(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
