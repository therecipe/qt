package widgets

//#include "widgets.h"
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
	for len(n.ObjectName()) < len("QGraphicsScale_") {
		n.SetObjectName("QGraphicsScale_" + qt.Identifier())
	}
	return n
}

func (ptr *QGraphicsScale) QGraphicsScale_PTR() *QGraphicsScale {
	return ptr
}

func (ptr *QGraphicsScale) SetOrigin(point gui.QVector3D_ITF) {
	defer qt.Recovering("QGraphicsScale::setOrigin")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_SetOrigin(ptr.Pointer(), gui.PointerFromQVector3D(point))
	}
}

func (ptr *QGraphicsScale) SetXScale(v float64) {
	defer qt.Recovering("QGraphicsScale::setXScale")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_SetXScale(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QGraphicsScale) SetYScale(v float64) {
	defer qt.Recovering("QGraphicsScale::setYScale")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_SetYScale(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QGraphicsScale) SetZScale(v float64) {
	defer qt.Recovering("QGraphicsScale::setZScale")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_SetZScale(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QGraphicsScale) XScale() float64 {
	defer qt.Recovering("QGraphicsScale::xScale")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsScale_XScale(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsScale) YScale() float64 {
	defer qt.Recovering("QGraphicsScale::yScale")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsScale_YScale(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsScale) ZScale() float64 {
	defer qt.Recovering("QGraphicsScale::zScale")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsScale_ZScale(ptr.Pointer()))
	}
	return 0
}

func NewQGraphicsScale(parent core.QObject_ITF) *QGraphicsScale {
	defer qt.Recovering("QGraphicsScale::QGraphicsScale")

	return NewQGraphicsScaleFromPointer(C.QGraphicsScale_NewQGraphicsScale(core.PointerFromQObject(parent)))
}

func (ptr *QGraphicsScale) ConnectApplyTo(f func(matrix *gui.QMatrix4x4)) {
	defer qt.Recovering("connect QGraphicsScale::applyTo")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "applyTo", f)
	}
}

func (ptr *QGraphicsScale) DisconnectApplyTo() {
	defer qt.Recovering("disconnect QGraphicsScale::applyTo")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "applyTo")
	}
}

//export callbackQGraphicsScaleApplyTo
func callbackQGraphicsScaleApplyTo(ptrName *C.char, matrix unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsScale::applyTo")

	var signal = qt.GetSignal(C.GoString(ptrName), "applyTo")
	if signal != nil {
		defer signal.(func(*gui.QMatrix4x4))(gui.NewQMatrix4x4FromPointer(matrix))
		return true
	}
	return false

}

func (ptr *QGraphicsScale) ConnectOriginChanged(f func()) {
	defer qt.Recovering("connect QGraphicsScale::originChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_ConnectOriginChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "originChanged", f)
	}
}

func (ptr *QGraphicsScale) DisconnectOriginChanged() {
	defer qt.Recovering("disconnect QGraphicsScale::originChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_DisconnectOriginChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "originChanged")
	}
}

//export callbackQGraphicsScaleOriginChanged
func callbackQGraphicsScaleOriginChanged(ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsScale::originChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "originChanged")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsScale) ConnectScaleChanged(f func()) {
	defer qt.Recovering("connect QGraphicsScale::scaleChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_ConnectScaleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "scaleChanged", f)
	}
}

func (ptr *QGraphicsScale) DisconnectScaleChanged() {
	defer qt.Recovering("disconnect QGraphicsScale::scaleChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_DisconnectScaleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "scaleChanged")
	}
}

//export callbackQGraphicsScaleScaleChanged
func callbackQGraphicsScaleScaleChanged(ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsScale::scaleChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "scaleChanged")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsScale) ConnectXScaleChanged(f func()) {
	defer qt.Recovering("connect QGraphicsScale::xScaleChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_ConnectXScaleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "xScaleChanged", f)
	}
}

func (ptr *QGraphicsScale) DisconnectXScaleChanged() {
	defer qt.Recovering("disconnect QGraphicsScale::xScaleChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_DisconnectXScaleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "xScaleChanged")
	}
}

//export callbackQGraphicsScaleXScaleChanged
func callbackQGraphicsScaleXScaleChanged(ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsScale::xScaleChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "xScaleChanged")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsScale) ConnectYScaleChanged(f func()) {
	defer qt.Recovering("connect QGraphicsScale::yScaleChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_ConnectYScaleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "yScaleChanged", f)
	}
}

func (ptr *QGraphicsScale) DisconnectYScaleChanged() {
	defer qt.Recovering("disconnect QGraphicsScale::yScaleChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_DisconnectYScaleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "yScaleChanged")
	}
}

//export callbackQGraphicsScaleYScaleChanged
func callbackQGraphicsScaleYScaleChanged(ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsScale::yScaleChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "yScaleChanged")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsScale) ConnectZScaleChanged(f func()) {
	defer qt.Recovering("connect QGraphicsScale::zScaleChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_ConnectZScaleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "zScaleChanged", f)
	}
}

func (ptr *QGraphicsScale) DisconnectZScaleChanged() {
	defer qt.Recovering("disconnect QGraphicsScale::zScaleChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_DisconnectZScaleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "zScaleChanged")
	}
}

//export callbackQGraphicsScaleZScaleChanged
func callbackQGraphicsScaleZScaleChanged(ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsScale::zScaleChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "zScaleChanged")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsScale) DestroyQGraphicsScale() {
	defer qt.Recovering("QGraphicsScale::~QGraphicsScale")

	if ptr.Pointer() != nil {
		C.QGraphicsScale_DestroyQGraphicsScale(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
