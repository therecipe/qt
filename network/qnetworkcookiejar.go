package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QNetworkCookieJar struct {
	core.QObject
}

type QNetworkCookieJar_ITF interface {
	core.QObject_ITF
	QNetworkCookieJar_PTR() *QNetworkCookieJar
}

func PointerFromQNetworkCookieJar(ptr QNetworkCookieJar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkCookieJar_PTR().Pointer()
	}
	return nil
}

func NewQNetworkCookieJarFromPointer(ptr unsafe.Pointer) *QNetworkCookieJar {
	var n = new(QNetworkCookieJar)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QNetworkCookieJar_") {
		n.SetObjectName("QNetworkCookieJar_" + qt.Identifier())
	}
	return n
}

func (ptr *QNetworkCookieJar) QNetworkCookieJar_PTR() *QNetworkCookieJar {
	return ptr
}

func NewQNetworkCookieJar(parent core.QObject_ITF) *QNetworkCookieJar {
	defer qt.Recovering("QNetworkCookieJar::QNetworkCookieJar")

	return NewQNetworkCookieJarFromPointer(C.QNetworkCookieJar_NewQNetworkCookieJar(core.PointerFromQObject(parent)))
}

func (ptr *QNetworkCookieJar) DeleteCookie(cookie QNetworkCookie_ITF) bool {
	defer qt.Recovering("QNetworkCookieJar::deleteCookie")

	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_DeleteCookie(ptr.Pointer(), PointerFromQNetworkCookie(cookie)) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) InsertCookie(cookie QNetworkCookie_ITF) bool {
	defer qt.Recovering("QNetworkCookieJar::insertCookie")

	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_InsertCookie(ptr.Pointer(), PointerFromQNetworkCookie(cookie)) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) UpdateCookie(cookie QNetworkCookie_ITF) bool {
	defer qt.Recovering("QNetworkCookieJar::updateCookie")

	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_UpdateCookie(ptr.Pointer(), PointerFromQNetworkCookie(cookie)) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) ValidateCookie(cookie QNetworkCookie_ITF, url core.QUrl_ITF) bool {
	defer qt.Recovering("QNetworkCookieJar::validateCookie")

	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_ValidateCookie(ptr.Pointer(), PointerFromQNetworkCookie(cookie), core.PointerFromQUrl(url)) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) DestroyQNetworkCookieJar() {
	defer qt.Recovering("QNetworkCookieJar::~QNetworkCookieJar")

	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_DestroyQNetworkCookieJar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkCookieJar) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QNetworkCookieJar::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QNetworkCookieJar::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQNetworkCookieJarTimerEvent
func callbackQNetworkCookieJarTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkCookieJar::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNetworkCookieJarFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNetworkCookieJar) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkCookieJar::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkCookieJar) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkCookieJar::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkCookieJar) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QNetworkCookieJar::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QNetworkCookieJar::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQNetworkCookieJarChildEvent
func callbackQNetworkCookieJarChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkCookieJar::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNetworkCookieJarFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNetworkCookieJar) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkCookieJar::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkCookieJar) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkCookieJar::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkCookieJar) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QNetworkCookieJar::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QNetworkCookieJar::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQNetworkCookieJarCustomEvent
func callbackQNetworkCookieJarCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkCookieJar::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNetworkCookieJarFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNetworkCookieJar) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkCookieJar::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNetworkCookieJar) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkCookieJar::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
