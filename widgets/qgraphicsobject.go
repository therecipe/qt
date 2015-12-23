package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGraphicsObject struct {
	QGraphicsItem
	core.QObject
}

type QGraphicsObject_ITF interface {
	QGraphicsItem_ITF
	core.QObject_ITF
	QGraphicsObject_PTR() *QGraphicsObject
}

func (p *QGraphicsObject) Pointer() unsafe.Pointer {
	return p.QGraphicsItem_PTR().Pointer()
}

func (p *QGraphicsObject) SetPointer(ptr unsafe.Pointer) {
	p.QGraphicsItem_PTR().SetPointer(ptr)
	p.QObject_PTR().SetPointer(ptr)
}

func PointerFromQGraphicsObject(ptr QGraphicsObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsObject_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsObjectFromPointer(ptr unsafe.Pointer) *QGraphicsObject {
	var n = new(QGraphicsObject)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGraphicsObject_") {
		n.SetObjectName("QGraphicsObject_" + qt.Identifier())
	}
	return n
}

func (ptr *QGraphicsObject) QGraphicsObject_PTR() *QGraphicsObject {
	return ptr
}

func (ptr *QGraphicsObject) ConnectEnabledChanged(f func()) {
	defer qt.Recovering("connect QGraphicsObject::enabledChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsObject_ConnectEnabledChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "enabledChanged", f)
	}
}

func (ptr *QGraphicsObject) DisconnectEnabledChanged() {
	defer qt.Recovering("disconnect QGraphicsObject::enabledChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsObject_DisconnectEnabledChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "enabledChanged")
	}
}

//export callbackQGraphicsObjectEnabledChanged
func callbackQGraphicsObjectEnabledChanged(ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsObject::enabledChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "enabledChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsObject) GrabGesture(gesture core.Qt__GestureType, flags core.Qt__GestureFlag) {
	defer qt.Recovering("QGraphicsObject::grabGesture")

	if ptr.Pointer() != nil {
		C.QGraphicsObject_GrabGesture(ptr.Pointer(), C.int(gesture), C.int(flags))
	}
}

func (ptr *QGraphicsObject) ConnectOpacityChanged(f func()) {
	defer qt.Recovering("connect QGraphicsObject::opacityChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsObject_ConnectOpacityChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "opacityChanged", f)
	}
}

func (ptr *QGraphicsObject) DisconnectOpacityChanged() {
	defer qt.Recovering("disconnect QGraphicsObject::opacityChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsObject_DisconnectOpacityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "opacityChanged")
	}
}

//export callbackQGraphicsObjectOpacityChanged
func callbackQGraphicsObjectOpacityChanged(ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsObject::opacityChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "opacityChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsObject) ConnectParentChanged(f func()) {
	defer qt.Recovering("connect QGraphicsObject::parentChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsObject_ConnectParentChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "parentChanged", f)
	}
}

func (ptr *QGraphicsObject) DisconnectParentChanged() {
	defer qt.Recovering("disconnect QGraphicsObject::parentChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsObject_DisconnectParentChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "parentChanged")
	}
}

//export callbackQGraphicsObjectParentChanged
func callbackQGraphicsObjectParentChanged(ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsObject::parentChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "parentChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsObject) ConnectRotationChanged(f func()) {
	defer qt.Recovering("connect QGraphicsObject::rotationChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsObject_ConnectRotationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "rotationChanged", f)
	}
}

func (ptr *QGraphicsObject) DisconnectRotationChanged() {
	defer qt.Recovering("disconnect QGraphicsObject::rotationChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsObject_DisconnectRotationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "rotationChanged")
	}
}

//export callbackQGraphicsObjectRotationChanged
func callbackQGraphicsObjectRotationChanged(ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsObject::rotationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "rotationChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsObject) ConnectScaleChanged(f func()) {
	defer qt.Recovering("connect QGraphicsObject::scaleChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsObject_ConnectScaleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "scaleChanged", f)
	}
}

func (ptr *QGraphicsObject) DisconnectScaleChanged() {
	defer qt.Recovering("disconnect QGraphicsObject::scaleChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsObject_DisconnectScaleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "scaleChanged")
	}
}

//export callbackQGraphicsObjectScaleChanged
func callbackQGraphicsObjectScaleChanged(ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsObject::scaleChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "scaleChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsObject) UngrabGesture(gesture core.Qt__GestureType) {
	defer qt.Recovering("QGraphicsObject::ungrabGesture")

	if ptr.Pointer() != nil {
		C.QGraphicsObject_UngrabGesture(ptr.Pointer(), C.int(gesture))
	}
}

func (ptr *QGraphicsObject) ConnectVisibleChanged(f func()) {
	defer qt.Recovering("connect QGraphicsObject::visibleChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsObject_ConnectVisibleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "visibleChanged", f)
	}
}

func (ptr *QGraphicsObject) DisconnectVisibleChanged() {
	defer qt.Recovering("disconnect QGraphicsObject::visibleChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsObject_DisconnectVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "visibleChanged")
	}
}

//export callbackQGraphicsObjectVisibleChanged
func callbackQGraphicsObjectVisibleChanged(ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsObject::visibleChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "visibleChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsObject) ConnectXChanged(f func()) {
	defer qt.Recovering("connect QGraphicsObject::xChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsObject_ConnectXChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "xChanged", f)
	}
}

func (ptr *QGraphicsObject) DisconnectXChanged() {
	defer qt.Recovering("disconnect QGraphicsObject::xChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsObject_DisconnectXChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "xChanged")
	}
}

//export callbackQGraphicsObjectXChanged
func callbackQGraphicsObjectXChanged(ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsObject::xChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "xChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsObject) ConnectYChanged(f func()) {
	defer qt.Recovering("connect QGraphicsObject::yChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsObject_ConnectYChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "yChanged", f)
	}
}

func (ptr *QGraphicsObject) DisconnectYChanged() {
	defer qt.Recovering("disconnect QGraphicsObject::yChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsObject_DisconnectYChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "yChanged")
	}
}

//export callbackQGraphicsObjectYChanged
func callbackQGraphicsObjectYChanged(ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsObject::yChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "yChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsObject) ConnectZChanged(f func()) {
	defer qt.Recovering("connect QGraphicsObject::zChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsObject_ConnectZChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "zChanged", f)
	}
}

func (ptr *QGraphicsObject) DisconnectZChanged() {
	defer qt.Recovering("disconnect QGraphicsObject::zChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsObject_DisconnectZChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "zChanged")
	}
}

//export callbackQGraphicsObjectZChanged
func callbackQGraphicsObjectZChanged(ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsObject::zChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "zChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsObject) DestroyQGraphicsObject() {
	defer qt.Recovering("QGraphicsObject::~QGraphicsObject")

	if ptr.Pointer() != nil {
		C.QGraphicsObject_DestroyQGraphicsObject(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsObject) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGraphicsObject::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGraphicsObject) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGraphicsObject::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGraphicsObjectTimerEvent
func callbackQGraphicsObjectTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsObject::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsObject) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGraphicsObject::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGraphicsObject) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGraphicsObject::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGraphicsObjectChildEvent
func callbackQGraphicsObjectChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsObject::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsObject) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsObject::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGraphicsObject) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGraphicsObject::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGraphicsObjectCustomEvent
func callbackQGraphicsObjectCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsObject::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
