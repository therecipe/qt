package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QHttpMultiPart struct {
	core.QObject
}

type QHttpMultiPart_ITF interface {
	core.QObject_ITF
	QHttpMultiPart_PTR() *QHttpMultiPart
}

func PointerFromQHttpMultiPart(ptr QHttpMultiPart_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHttpMultiPart_PTR().Pointer()
	}
	return nil
}

func NewQHttpMultiPartFromPointer(ptr unsafe.Pointer) *QHttpMultiPart {
	var n = new(QHttpMultiPart)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QHttpMultiPart_") {
		n.SetObjectName("QHttpMultiPart_" + qt.Identifier())
	}
	return n
}

func (ptr *QHttpMultiPart) QHttpMultiPart_PTR() *QHttpMultiPart {
	return ptr
}

//QHttpMultiPart::ContentType
type QHttpMultiPart__ContentType int64

const (
	QHttpMultiPart__MixedType       = QHttpMultiPart__ContentType(0)
	QHttpMultiPart__RelatedType     = QHttpMultiPart__ContentType(1)
	QHttpMultiPart__FormDataType    = QHttpMultiPart__ContentType(2)
	QHttpMultiPart__AlternativeType = QHttpMultiPart__ContentType(3)
)

func NewQHttpMultiPart2(contentType QHttpMultiPart__ContentType, parent core.QObject_ITF) *QHttpMultiPart {
	defer qt.Recovering("QHttpMultiPart::QHttpMultiPart")

	return NewQHttpMultiPartFromPointer(C.QHttpMultiPart_NewQHttpMultiPart2(C.int(contentType), core.PointerFromQObject(parent)))
}

func NewQHttpMultiPart(parent core.QObject_ITF) *QHttpMultiPart {
	defer qt.Recovering("QHttpMultiPart::QHttpMultiPart")

	return NewQHttpMultiPartFromPointer(C.QHttpMultiPart_NewQHttpMultiPart(core.PointerFromQObject(parent)))
}

func (ptr *QHttpMultiPart) Append(httpPart QHttpPart_ITF) {
	defer qt.Recovering("QHttpMultiPart::append")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_Append(ptr.Pointer(), PointerFromQHttpPart(httpPart))
	}
}

func (ptr *QHttpMultiPart) Boundary() *core.QByteArray {
	defer qt.Recovering("QHttpMultiPart::boundary")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QHttpMultiPart_Boundary(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHttpMultiPart) SetBoundary(boundary core.QByteArray_ITF) {
	defer qt.Recovering("QHttpMultiPart::setBoundary")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_SetBoundary(ptr.Pointer(), core.PointerFromQByteArray(boundary))
	}
}

func (ptr *QHttpMultiPart) SetContentType(contentType QHttpMultiPart__ContentType) {
	defer qt.Recovering("QHttpMultiPart::setContentType")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_SetContentType(ptr.Pointer(), C.int(contentType))
	}
}

func (ptr *QHttpMultiPart) DestroyQHttpMultiPart() {
	defer qt.Recovering("QHttpMultiPart::~QHttpMultiPart")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_DestroyQHttpMultiPart(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHttpMultiPart) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QHttpMultiPart::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QHttpMultiPart) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHttpMultiPart::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQHttpMultiPartTimerEvent
func callbackQHttpMultiPartTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHttpMultiPart::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHttpMultiPartFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHttpMultiPart) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHttpMultiPart::timerEvent")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHttpMultiPart) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHttpMultiPart::timerEvent")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHttpMultiPart) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHttpMultiPart::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QHttpMultiPart) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHttpMultiPart::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQHttpMultiPartChildEvent
func callbackQHttpMultiPartChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHttpMultiPart::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHttpMultiPartFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHttpMultiPart) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHttpMultiPart::childEvent")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHttpMultiPart) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHttpMultiPart::childEvent")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHttpMultiPart) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHttpMultiPart::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QHttpMultiPart) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHttpMultiPart::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQHttpMultiPartCustomEvent
func callbackQHttpMultiPartCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHttpMultiPart::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHttpMultiPartFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHttpMultiPart) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHttpMultiPart::customEvent")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHttpMultiPart) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHttpMultiPart::customEvent")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
