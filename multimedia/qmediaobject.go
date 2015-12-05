package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
		n.SetObjectName("QMediaObject_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaObject) QMediaObject_PTR() *QMediaObject {
	return ptr
}

func (ptr *QMediaObject) NotifyInterval() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaObject::notifyInterval")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QMediaObject_NotifyInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaObject) SetNotifyInterval(milliSeconds int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaObject::setNotifyInterval")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaObject_SetNotifyInterval(ptr.Pointer(), C.int(milliSeconds))
	}
}

func (ptr *QMediaObject) ConnectAvailabilityChanged(f func(available bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaObject::availabilityChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaObject_ConnectAvailabilityChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "availabilityChanged", f)
	}
}

func (ptr *QMediaObject) DisconnectAvailabilityChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaObject::availabilityChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaObject_DisconnectAvailabilityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "availabilityChanged")
	}
}

//export callbackQMediaObjectAvailabilityChanged
func callbackQMediaObjectAvailabilityChanged(ptrName *C.char, available C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaObject::availabilityChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "availabilityChanged").(func(bool))(int(available) != 0)
}

func (ptr *QMediaObject) AvailableMetaData() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaObject::availableMetaData")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMediaObject_AvailableMetaData(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QMediaObject) Bind(object core.QObject_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaObject::bind")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMediaObject_Bind(ptr.Pointer(), core.PointerFromQObject(object)) != 0
	}
	return false
}

func (ptr *QMediaObject) IsAvailable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaObject::isAvailable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMediaObject_IsAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaObject) IsMetaDataAvailable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaObject::isMetaDataAvailable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMediaObject_IsMetaDataAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaObject) MetaData(key string) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaObject::metaData")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QMediaObject_MetaData(ptr.Pointer(), C.CString(key)))
	}
	return nil
}

func (ptr *QMediaObject) ConnectMetaDataAvailableChanged(f func(available bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaObject::metaDataAvailableChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaObject_ConnectMetaDataAvailableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataAvailableChanged", f)
	}
}

func (ptr *QMediaObject) DisconnectMetaDataAvailableChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaObject::metaDataAvailableChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaObject_DisconnectMetaDataAvailableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataAvailableChanged")
	}
}

//export callbackQMediaObjectMetaDataAvailableChanged
func callbackQMediaObjectMetaDataAvailableChanged(ptrName *C.char, available C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaObject::metaDataAvailableChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "metaDataAvailableChanged").(func(bool))(int(available) != 0)
}

func (ptr *QMediaObject) ConnectMetaDataChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaObject::metaDataChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaObject_ConnectMetaDataChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataChanged", f)
	}
}

func (ptr *QMediaObject) DisconnectMetaDataChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaObject::metaDataChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaObject_DisconnectMetaDataChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataChanged")
	}
}

//export callbackQMediaObjectMetaDataChanged
func callbackQMediaObjectMetaDataChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaObject::metaDataChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "metaDataChanged").(func())()
}

func (ptr *QMediaObject) ConnectNotifyIntervalChanged(f func(milliseconds int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaObject::notifyIntervalChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaObject_ConnectNotifyIntervalChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "notifyIntervalChanged", f)
	}
}

func (ptr *QMediaObject) DisconnectNotifyIntervalChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaObject::notifyIntervalChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaObject_DisconnectNotifyIntervalChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "notifyIntervalChanged")
	}
}

//export callbackQMediaObjectNotifyIntervalChanged
func callbackQMediaObjectNotifyIntervalChanged(ptrName *C.char, milliseconds C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaObject::notifyIntervalChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "notifyIntervalChanged").(func(int))(int(milliseconds))
}

func (ptr *QMediaObject) Service() *QMediaService {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaObject::service")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQMediaServiceFromPointer(C.QMediaObject_Service(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaObject) Unbind(object core.QObject_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaObject::unbind")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaObject_Unbind(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QMediaObject) DestroyQMediaObject() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaObject::~QMediaObject")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaObject_DestroyQMediaObject(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
