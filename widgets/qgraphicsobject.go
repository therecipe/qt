package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
		n.SetObjectName("QGraphicsObject_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsObject) QGraphicsObject_PTR() *QGraphicsObject {
	return ptr
}

func (ptr *QGraphicsObject) ConnectEnabledChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsObject::enabledChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsObject_ConnectEnabledChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "enabledChanged", f)
	}
}

func (ptr *QGraphicsObject) DisconnectEnabledChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsObject::enabledChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsObject_DisconnectEnabledChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "enabledChanged")
	}
}

//export callbackQGraphicsObjectEnabledChanged
func callbackQGraphicsObjectEnabledChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsObject::enabledChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "enabledChanged").(func())()
}

func (ptr *QGraphicsObject) GrabGesture(gesture core.Qt__GestureType, flags core.Qt__GestureFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsObject::grabGesture")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsObject_GrabGesture(ptr.Pointer(), C.int(gesture), C.int(flags))
	}
}

func (ptr *QGraphicsObject) ConnectOpacityChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsObject::opacityChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsObject_ConnectOpacityChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "opacityChanged", f)
	}
}

func (ptr *QGraphicsObject) DisconnectOpacityChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsObject::opacityChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsObject_DisconnectOpacityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "opacityChanged")
	}
}

//export callbackQGraphicsObjectOpacityChanged
func callbackQGraphicsObjectOpacityChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsObject::opacityChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "opacityChanged").(func())()
}

func (ptr *QGraphicsObject) ConnectParentChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsObject::parentChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsObject_ConnectParentChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "parentChanged", f)
	}
}

func (ptr *QGraphicsObject) DisconnectParentChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsObject::parentChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsObject_DisconnectParentChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "parentChanged")
	}
}

//export callbackQGraphicsObjectParentChanged
func callbackQGraphicsObjectParentChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsObject::parentChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "parentChanged").(func())()
}

func (ptr *QGraphicsObject) ConnectRotationChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsObject::rotationChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsObject_ConnectRotationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "rotationChanged", f)
	}
}

func (ptr *QGraphicsObject) DisconnectRotationChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsObject::rotationChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsObject_DisconnectRotationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "rotationChanged")
	}
}

//export callbackQGraphicsObjectRotationChanged
func callbackQGraphicsObjectRotationChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsObject::rotationChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "rotationChanged").(func())()
}

func (ptr *QGraphicsObject) ConnectScaleChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsObject::scaleChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsObject_ConnectScaleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "scaleChanged", f)
	}
}

func (ptr *QGraphicsObject) DisconnectScaleChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsObject::scaleChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsObject_DisconnectScaleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "scaleChanged")
	}
}

//export callbackQGraphicsObjectScaleChanged
func callbackQGraphicsObjectScaleChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsObject::scaleChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "scaleChanged").(func())()
}

func (ptr *QGraphicsObject) UngrabGesture(gesture core.Qt__GestureType) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsObject::ungrabGesture")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsObject_UngrabGesture(ptr.Pointer(), C.int(gesture))
	}
}

func (ptr *QGraphicsObject) ConnectVisibleChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsObject::visibleChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsObject_ConnectVisibleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "visibleChanged", f)
	}
}

func (ptr *QGraphicsObject) DisconnectVisibleChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsObject::visibleChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsObject_DisconnectVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "visibleChanged")
	}
}

//export callbackQGraphicsObjectVisibleChanged
func callbackQGraphicsObjectVisibleChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsObject::visibleChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "visibleChanged").(func())()
}

func (ptr *QGraphicsObject) ConnectXChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsObject::xChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsObject_ConnectXChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "xChanged", f)
	}
}

func (ptr *QGraphicsObject) DisconnectXChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsObject::xChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsObject_DisconnectXChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "xChanged")
	}
}

//export callbackQGraphicsObjectXChanged
func callbackQGraphicsObjectXChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsObject::xChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "xChanged").(func())()
}

func (ptr *QGraphicsObject) ConnectYChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsObject::yChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsObject_ConnectYChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "yChanged", f)
	}
}

func (ptr *QGraphicsObject) DisconnectYChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsObject::yChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsObject_DisconnectYChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "yChanged")
	}
}

//export callbackQGraphicsObjectYChanged
func callbackQGraphicsObjectYChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsObject::yChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "yChanged").(func())()
}

func (ptr *QGraphicsObject) ConnectZChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsObject::zChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsObject_ConnectZChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "zChanged", f)
	}
}

func (ptr *QGraphicsObject) DisconnectZChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsObject::zChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsObject_DisconnectZChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "zChanged")
	}
}

//export callbackQGraphicsObjectZChanged
func callbackQGraphicsObjectZChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsObject::zChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "zChanged").(func())()
}

func (ptr *QGraphicsObject) DestroyQGraphicsObject() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsObject::~QGraphicsObject")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsObject_DestroyQGraphicsObject(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
