package widgets

//#include "qgraphicsscale.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QGraphicsScale struct {
	QGraphicsTransform
}

type QGraphicsScaleITF interface {
	QGraphicsTransformITF
	QGraphicsScalePTR() *QGraphicsScale
}

func PointerFromQGraphicsScale(ptr QGraphicsScaleITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsScalePTR().Pointer()
	}
	return nil
}

func QGraphicsScaleFromPointer(ptr unsafe.Pointer) *QGraphicsScale {
	var n = new(QGraphicsScale)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGraphicsScale_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsScale) QGraphicsScalePTR() *QGraphicsScale {
	return ptr
}

func (ptr *QGraphicsScale) SetOrigin(point gui.QVector3DITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScale_SetOrigin(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQVector3D(point)))
	}
}

func NewQGraphicsScale(parent core.QObjectITF) *QGraphicsScale {
	return QGraphicsScaleFromPointer(unsafe.Pointer(C.QGraphicsScale_NewQGraphicsScale(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QGraphicsScale) ApplyTo(matrix gui.QMatrix4x4ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScale_ApplyTo(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQMatrix4x4(matrix)))
	}
}

func (ptr *QGraphicsScale) ConnectOriginChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QGraphicsScale_ConnectOriginChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "originChanged", f)
	}
}

func (ptr *QGraphicsScale) DisconnectOriginChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsScale_DisconnectOriginChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "originChanged")
	}
}

//export callbackQGraphicsScaleOriginChanged
func callbackQGraphicsScaleOriginChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "originChanged").(func())()
}

func (ptr *QGraphicsScale) ConnectScaleChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QGraphicsScale_ConnectScaleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "scaleChanged", f)
	}
}

func (ptr *QGraphicsScale) DisconnectScaleChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsScale_DisconnectScaleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "scaleChanged")
	}
}

//export callbackQGraphicsScaleScaleChanged
func callbackQGraphicsScaleScaleChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "scaleChanged").(func())()
}

func (ptr *QGraphicsScale) ConnectXScaleChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QGraphicsScale_ConnectXScaleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "xScaleChanged", f)
	}
}

func (ptr *QGraphicsScale) DisconnectXScaleChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsScale_DisconnectXScaleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "xScaleChanged")
	}
}

//export callbackQGraphicsScaleXScaleChanged
func callbackQGraphicsScaleXScaleChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "xScaleChanged").(func())()
}

func (ptr *QGraphicsScale) ConnectYScaleChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QGraphicsScale_ConnectYScaleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "yScaleChanged", f)
	}
}

func (ptr *QGraphicsScale) DisconnectYScaleChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsScale_DisconnectYScaleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "yScaleChanged")
	}
}

//export callbackQGraphicsScaleYScaleChanged
func callbackQGraphicsScaleYScaleChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "yScaleChanged").(func())()
}

func (ptr *QGraphicsScale) ConnectZScaleChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QGraphicsScale_ConnectZScaleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "zScaleChanged", f)
	}
}

func (ptr *QGraphicsScale) DisconnectZScaleChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsScale_DisconnectZScaleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "zScaleChanged")
	}
}

//export callbackQGraphicsScaleZScaleChanged
func callbackQGraphicsScaleZScaleChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "zScaleChanged").(func())()
}

func (ptr *QGraphicsScale) DestroyQGraphicsScale() {
	if ptr.Pointer() != nil {
		C.QGraphicsScale_DestroyQGraphicsScale(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
