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

type QGraphicsScale_ITF interface {
	QGraphicsTransform_ITF
	QGraphicsScale_PTR() *QGraphicsScale
}

func PointerFromQGraphicsScale(ptr QGraphicsScale_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsScale_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsScaleFromPointer(ptr unsafe.Pointer) *QGraphicsScale {
	var n = new(QGraphicsScale)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGraphicsScale_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsScale) QGraphicsScale_PTR() *QGraphicsScale {
	return ptr
}

func (ptr *QGraphicsScale) SetOrigin(point gui.QVector3D_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScale_SetOrigin(ptr.Pointer(), gui.PointerFromQVector3D(point))
	}
}

func (ptr *QGraphicsScale) SetXScale(v float64) {
	if ptr.Pointer() != nil {
		C.QGraphicsScale_SetXScale(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QGraphicsScale) SetYScale(v float64) {
	if ptr.Pointer() != nil {
		C.QGraphicsScale_SetYScale(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QGraphicsScale) SetZScale(v float64) {
	if ptr.Pointer() != nil {
		C.QGraphicsScale_SetZScale(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QGraphicsScale) XScale() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGraphicsScale_XScale(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsScale) YScale() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGraphicsScale_YScale(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsScale) ZScale() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGraphicsScale_ZScale(ptr.Pointer()))
	}
	return 0
}

func NewQGraphicsScale(parent core.QObject_ITF) *QGraphicsScale {
	return NewQGraphicsScaleFromPointer(C.QGraphicsScale_NewQGraphicsScale(core.PointerFromQObject(parent)))
}

func (ptr *QGraphicsScale) ApplyTo(matrix gui.QMatrix4x4_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScale_ApplyTo(ptr.Pointer(), gui.PointerFromQMatrix4x4(matrix))
	}
}

func (ptr *QGraphicsScale) ConnectOriginChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QGraphicsScale_ConnectOriginChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "originChanged", f)
	}
}

func (ptr *QGraphicsScale) DisconnectOriginChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsScale_DisconnectOriginChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "originChanged")
	}
}

//export callbackQGraphicsScaleOriginChanged
func callbackQGraphicsScaleOriginChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "originChanged").(func())()
}

func (ptr *QGraphicsScale) ConnectScaleChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QGraphicsScale_ConnectScaleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "scaleChanged", f)
	}
}

func (ptr *QGraphicsScale) DisconnectScaleChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsScale_DisconnectScaleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "scaleChanged")
	}
}

//export callbackQGraphicsScaleScaleChanged
func callbackQGraphicsScaleScaleChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "scaleChanged").(func())()
}

func (ptr *QGraphicsScale) ConnectXScaleChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QGraphicsScale_ConnectXScaleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "xScaleChanged", f)
	}
}

func (ptr *QGraphicsScale) DisconnectXScaleChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsScale_DisconnectXScaleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "xScaleChanged")
	}
}

//export callbackQGraphicsScaleXScaleChanged
func callbackQGraphicsScaleXScaleChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "xScaleChanged").(func())()
}

func (ptr *QGraphicsScale) ConnectYScaleChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QGraphicsScale_ConnectYScaleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "yScaleChanged", f)
	}
}

func (ptr *QGraphicsScale) DisconnectYScaleChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsScale_DisconnectYScaleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "yScaleChanged")
	}
}

//export callbackQGraphicsScaleYScaleChanged
func callbackQGraphicsScaleYScaleChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "yScaleChanged").(func())()
}

func (ptr *QGraphicsScale) ConnectZScaleChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QGraphicsScale_ConnectZScaleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "zScaleChanged", f)
	}
}

func (ptr *QGraphicsScale) DisconnectZScaleChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsScale_DisconnectZScaleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "zScaleChanged")
	}
}

//export callbackQGraphicsScaleZScaleChanged
func callbackQGraphicsScaleZScaleChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "zScaleChanged").(func())()
}

func (ptr *QGraphicsScale) DestroyQGraphicsScale() {
	if ptr.Pointer() != nil {
		C.QGraphicsScale_DestroyQGraphicsScale(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
