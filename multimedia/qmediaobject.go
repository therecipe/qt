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

type QMediaObjectITF interface {
	core.QObjectITF
	QMediaObjectPTR() *QMediaObject
}

func PointerFromQMediaObject(ptr QMediaObjectITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaObjectPTR().Pointer()
	}
	return nil
}

func QMediaObjectFromPointer(ptr unsafe.Pointer) *QMediaObject {
	var n = new(QMediaObject)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMediaObject_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaObject) QMediaObjectPTR() *QMediaObject {
	return ptr
}

func (ptr *QMediaObject) NotifyInterval() int {
	if ptr.Pointer() != nil {
		return int(C.QMediaObject_NotifyInterval(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMediaObject) SetNotifyInterval(milliSeconds int) {
	if ptr.Pointer() != nil {
		C.QMediaObject_SetNotifyInterval(C.QtObjectPtr(ptr.Pointer()), C.int(milliSeconds))
	}
}

func (ptr *QMediaObject) ConnectAvailabilityChanged(f func(available bool)) {
	if ptr.Pointer() != nil {
		C.QMediaObject_ConnectAvailabilityChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "availabilityChanged", f)
	}
}

func (ptr *QMediaObject) DisconnectAvailabilityChanged() {
	if ptr.Pointer() != nil {
		C.QMediaObject_DisconnectAvailabilityChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "availabilityChanged")
	}
}

//export callbackQMediaObjectAvailabilityChanged
func callbackQMediaObjectAvailabilityChanged(ptrName *C.char, available C.int) {
	qt.GetSignal(C.GoString(ptrName), "availabilityChanged").(func(bool))(int(available) != 0)
}

func (ptr *QMediaObject) AvailableMetaData() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMediaObject_AvailableMetaData(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QMediaObject) Bind(object core.QObjectITF) bool {
	if ptr.Pointer() != nil {
		return C.QMediaObject_Bind(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQObject(object))) != 0
	}
	return false
}

func (ptr *QMediaObject) IsAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QMediaObject_IsAvailable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMediaObject) IsMetaDataAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QMediaObject_IsMetaDataAvailable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMediaObject) MetaData(key string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaObject_MetaData(C.QtObjectPtr(ptr.Pointer()), C.CString(key)))
	}
	return ""
}

func (ptr *QMediaObject) ConnectMetaDataAvailableChanged(f func(available bool)) {
	if ptr.Pointer() != nil {
		C.QMediaObject_ConnectMetaDataAvailableChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "metaDataAvailableChanged", f)
	}
}

func (ptr *QMediaObject) DisconnectMetaDataAvailableChanged() {
	if ptr.Pointer() != nil {
		C.QMediaObject_DisconnectMetaDataAvailableChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataAvailableChanged")
	}
}

//export callbackQMediaObjectMetaDataAvailableChanged
func callbackQMediaObjectMetaDataAvailableChanged(ptrName *C.char, available C.int) {
	qt.GetSignal(C.GoString(ptrName), "metaDataAvailableChanged").(func(bool))(int(available) != 0)
}

func (ptr *QMediaObject) ConnectMetaDataChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QMediaObject_ConnectMetaDataChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "metaDataChanged", f)
	}
}

func (ptr *QMediaObject) DisconnectMetaDataChanged() {
	if ptr.Pointer() != nil {
		C.QMediaObject_DisconnectMetaDataChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataChanged")
	}
}

//export callbackQMediaObjectMetaDataChanged
func callbackQMediaObjectMetaDataChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "metaDataChanged").(func())()
}

func (ptr *QMediaObject) ConnectNotifyIntervalChanged(f func(milliseconds int)) {
	if ptr.Pointer() != nil {
		C.QMediaObject_ConnectNotifyIntervalChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "notifyIntervalChanged", f)
	}
}

func (ptr *QMediaObject) DisconnectNotifyIntervalChanged() {
	if ptr.Pointer() != nil {
		C.QMediaObject_DisconnectNotifyIntervalChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "notifyIntervalChanged")
	}
}

//export callbackQMediaObjectNotifyIntervalChanged
func callbackQMediaObjectNotifyIntervalChanged(ptrName *C.char, milliseconds C.int) {
	qt.GetSignal(C.GoString(ptrName), "notifyIntervalChanged").(func(int))(int(milliseconds))
}

func (ptr *QMediaObject) Service() *QMediaService {
	if ptr.Pointer() != nil {
		return QMediaServiceFromPointer(unsafe.Pointer(C.QMediaObject_Service(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QMediaObject) Unbind(object core.QObjectITF) {
	if ptr.Pointer() != nil {
		C.QMediaObject_Unbind(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQObject(object)))
	}
}

func (ptr *QMediaObject) DestroyQMediaObject() {
	if ptr.Pointer() != nil {
		C.QMediaObject_DestroyQMediaObject(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
