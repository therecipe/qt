package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMediaService struct {
	core.QObject
}

type QMediaService_ITF interface {
	core.QObject_ITF
	QMediaService_PTR() *QMediaService
}

func PointerFromQMediaService(ptr QMediaService_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaService_PTR().Pointer()
	}
	return nil
}

func NewQMediaServiceFromPointer(ptr unsafe.Pointer) *QMediaService {
	var n = new(QMediaService)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMediaService_") {
		n.SetObjectName("QMediaService_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaService) QMediaService_PTR() *QMediaService {
	return ptr
}

func (ptr *QMediaService) ReleaseControl(control QMediaControl_ITF) {
	defer qt.Recovering("QMediaService::releaseControl")

	if ptr.Pointer() != nil {
		C.QMediaService_ReleaseControl(ptr.Pointer(), PointerFromQMediaControl(control))
	}
}

func (ptr *QMediaService) RequestControl(interfa string) *QMediaControl {
	defer qt.Recovering("QMediaService::requestControl")

	if ptr.Pointer() != nil {
		return NewQMediaControlFromPointer(C.QMediaService_RequestControl(ptr.Pointer(), C.CString(interfa)))
	}
	return nil
}

func (ptr *QMediaService) RequestControl2() unsafe.Pointer {
	defer qt.Recovering("QMediaService::requestControl")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QMediaService_RequestControl2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaService) DestroyQMediaService() {
	defer qt.Recovering("QMediaService::~QMediaService")

	if ptr.Pointer() != nil {
		C.QMediaService_DestroyQMediaService(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMediaService) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMediaService::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMediaService) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMediaService::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMediaServiceTimerEvent
func callbackQMediaServiceTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMediaService::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMediaService) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMediaService::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMediaService) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMediaService::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMediaServiceChildEvent
func callbackQMediaServiceChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMediaService::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMediaService) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMediaService::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMediaService) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMediaService::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMediaServiceCustomEvent
func callbackQMediaServiceCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMediaService::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
