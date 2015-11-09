package widgets

//#include "qgraphicsobject.h"
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
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QGraphicsObject_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsObject) QGraphicsObject_PTR() *QGraphicsObject {
	return ptr
}

func (ptr *QGraphicsObject) ConnectEnabledChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_ConnectEnabledChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "enabledChanged", f)
	}
}

func (ptr *QGraphicsObject) DisconnectEnabledChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_DisconnectEnabledChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "enabledChanged")
	}
}

//export callbackQGraphicsObjectEnabledChanged
func callbackQGraphicsObjectEnabledChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "enabledChanged").(func())()
}

func (ptr *QGraphicsObject) GrabGesture(gesture core.Qt__GestureType, flags core.Qt__GestureFlag) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_GrabGesture(ptr.Pointer(), C.int(gesture), C.int(flags))
	}
}

func (ptr *QGraphicsObject) ConnectOpacityChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_ConnectOpacityChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "opacityChanged", f)
	}
}

func (ptr *QGraphicsObject) DisconnectOpacityChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_DisconnectOpacityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "opacityChanged")
	}
}

//export callbackQGraphicsObjectOpacityChanged
func callbackQGraphicsObjectOpacityChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "opacityChanged").(func())()
}

func (ptr *QGraphicsObject) ConnectParentChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_ConnectParentChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "parentChanged", f)
	}
}

func (ptr *QGraphicsObject) DisconnectParentChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_DisconnectParentChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "parentChanged")
	}
}

//export callbackQGraphicsObjectParentChanged
func callbackQGraphicsObjectParentChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "parentChanged").(func())()
}

func (ptr *QGraphicsObject) ConnectRotationChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_ConnectRotationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "rotationChanged", f)
	}
}

func (ptr *QGraphicsObject) DisconnectRotationChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_DisconnectRotationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "rotationChanged")
	}
}

//export callbackQGraphicsObjectRotationChanged
func callbackQGraphicsObjectRotationChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "rotationChanged").(func())()
}

func (ptr *QGraphicsObject) ConnectScaleChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_ConnectScaleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "scaleChanged", f)
	}
}

func (ptr *QGraphicsObject) DisconnectScaleChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_DisconnectScaleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "scaleChanged")
	}
}

//export callbackQGraphicsObjectScaleChanged
func callbackQGraphicsObjectScaleChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "scaleChanged").(func())()
}

func (ptr *QGraphicsObject) UngrabGesture(gesture core.Qt__GestureType) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_UngrabGesture(ptr.Pointer(), C.int(gesture))
	}
}

func (ptr *QGraphicsObject) ConnectVisibleChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_ConnectVisibleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "visibleChanged", f)
	}
}

func (ptr *QGraphicsObject) DisconnectVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_DisconnectVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "visibleChanged")
	}
}

//export callbackQGraphicsObjectVisibleChanged
func callbackQGraphicsObjectVisibleChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "visibleChanged").(func())()
}

func (ptr *QGraphicsObject) ConnectXChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_ConnectXChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "xChanged", f)
	}
}

func (ptr *QGraphicsObject) DisconnectXChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_DisconnectXChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "xChanged")
	}
}

//export callbackQGraphicsObjectXChanged
func callbackQGraphicsObjectXChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "xChanged").(func())()
}

func (ptr *QGraphicsObject) ConnectYChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_ConnectYChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "yChanged", f)
	}
}

func (ptr *QGraphicsObject) DisconnectYChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_DisconnectYChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "yChanged")
	}
}

//export callbackQGraphicsObjectYChanged
func callbackQGraphicsObjectYChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "yChanged").(func())()
}

func (ptr *QGraphicsObject) ConnectZChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_ConnectZChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "zChanged", f)
	}
}

func (ptr *QGraphicsObject) DisconnectZChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_DisconnectZChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "zChanged")
	}
}

//export callbackQGraphicsObjectZChanged
func callbackQGraphicsObjectZChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "zChanged").(func())()
}

func (ptr *QGraphicsObject) DestroyQGraphicsObject() {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_DestroyQGraphicsObject(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
