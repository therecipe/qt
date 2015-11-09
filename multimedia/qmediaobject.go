package multimedia

//#include "qmediaobject.h"
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
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QMediaObject_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaObject) QMediaObject_PTR() *QMediaObject {
	return ptr
}

func (ptr *QMediaObject) NotifyInterval() int {
	if ptr.Pointer() != nil {
		return int(C.QMediaObject_NotifyInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaObject) SetNotifyInterval(milliSeconds int) {
	if ptr.Pointer() != nil {
		C.QMediaObject_SetNotifyInterval(ptr.Pointer(), C.int(milliSeconds))
	}
}

func (ptr *QMediaObject) ConnectAvailabilityChanged(f func(available bool)) {
	if ptr.Pointer() != nil {
		C.QMediaObject_ConnectAvailabilityChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "availabilityChanged", f)
	}
}

func (ptr *QMediaObject) DisconnectAvailabilityChanged() {
	if ptr.Pointer() != nil {
		C.QMediaObject_DisconnectAvailabilityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "availabilityChanged")
	}
}

//export callbackQMediaObjectAvailabilityChanged
func callbackQMediaObjectAvailabilityChanged(ptrName *C.char, available C.int) {
	qt.GetSignal(C.GoString(ptrName), "availabilityChanged").(func(bool))(int(available) != 0)
}

func (ptr *QMediaObject) AvailableMetaData() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMediaObject_AvailableMetaData(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QMediaObject) Bind(object core.QObject_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QMediaObject_Bind(ptr.Pointer(), core.PointerFromQObject(object)) != 0
	}
	return false
}

func (ptr *QMediaObject) IsAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QMediaObject_IsAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaObject) IsMetaDataAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QMediaObject_IsMetaDataAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaObject) MetaData(key string) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QMediaObject_MetaData(ptr.Pointer(), C.CString(key)))
	}
	return nil
}

func (ptr *QMediaObject) ConnectMetaDataAvailableChanged(f func(available bool)) {
	if ptr.Pointer() != nil {
		C.QMediaObject_ConnectMetaDataAvailableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataAvailableChanged", f)
	}
}

func (ptr *QMediaObject) DisconnectMetaDataAvailableChanged() {
	if ptr.Pointer() != nil {
		C.QMediaObject_DisconnectMetaDataAvailableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataAvailableChanged")
	}
}

//export callbackQMediaObjectMetaDataAvailableChanged
func callbackQMediaObjectMetaDataAvailableChanged(ptrName *C.char, available C.int) {
	qt.GetSignal(C.GoString(ptrName), "metaDataAvailableChanged").(func(bool))(int(available) != 0)
}

func (ptr *QMediaObject) ConnectMetaDataChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QMediaObject_ConnectMetaDataChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataChanged", f)
	}
}

func (ptr *QMediaObject) DisconnectMetaDataChanged() {
	if ptr.Pointer() != nil {
		C.QMediaObject_DisconnectMetaDataChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataChanged")
	}
}

//export callbackQMediaObjectMetaDataChanged
func callbackQMediaObjectMetaDataChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "metaDataChanged").(func())()
}

func (ptr *QMediaObject) ConnectNotifyIntervalChanged(f func(milliseconds int)) {
	if ptr.Pointer() != nil {
		C.QMediaObject_ConnectNotifyIntervalChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "notifyIntervalChanged", f)
	}
}

func (ptr *QMediaObject) DisconnectNotifyIntervalChanged() {
	if ptr.Pointer() != nil {
		C.QMediaObject_DisconnectNotifyIntervalChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "notifyIntervalChanged")
	}
}

//export callbackQMediaObjectNotifyIntervalChanged
func callbackQMediaObjectNotifyIntervalChanged(ptrName *C.char, milliseconds C.int) {
	qt.GetSignal(C.GoString(ptrName), "notifyIntervalChanged").(func(int))(int(milliseconds))
}

func (ptr *QMediaObject) Service() *QMediaService {
	if ptr.Pointer() != nil {
		return NewQMediaServiceFromPointer(C.QMediaObject_Service(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaObject) Unbind(object core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QMediaObject_Unbind(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QMediaObject) DestroyQMediaObject() {
	if ptr.Pointer() != nil {
		C.QMediaObject_DestroyQMediaObject(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
