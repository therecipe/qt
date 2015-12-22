package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QGraphicsRotation struct {
	QGraphicsTransform
}

type QGraphicsRotation_ITF interface {
	QGraphicsTransform_ITF
	QGraphicsRotation_PTR() *QGraphicsRotation
}

func PointerFromQGraphicsRotation(ptr QGraphicsRotation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsRotation_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsRotationFromPointer(ptr unsafe.Pointer) *QGraphicsRotation {
	var n = new(QGraphicsRotation)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGraphicsRotation_") {
		n.SetObjectName("QGraphicsRotation_" + qt.Identifier())
	}
	return n
}

func (ptr *QGraphicsRotation) QGraphicsRotation_PTR() *QGraphicsRotation {
	return ptr
}

func (ptr *QGraphicsRotation) Angle() float64 {
	defer qt.Recovering("QGraphicsRotation::angle")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsRotation_Angle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsRotation) SetAngle(v float64) {
	defer qt.Recovering("QGraphicsRotation::setAngle")

	if ptr.Pointer() != nil {
		C.QGraphicsRotation_SetAngle(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QGraphicsRotation) SetAxis2(axis core.Qt__Axis) {
	defer qt.Recovering("QGraphicsRotation::setAxis")

	if ptr.Pointer() != nil {
		C.QGraphicsRotation_SetAxis2(ptr.Pointer(), C.int(axis))
	}
}

func (ptr *QGraphicsRotation) SetAxis(axis gui.QVector3D_ITF) {
	defer qt.Recovering("QGraphicsRotation::setAxis")

	if ptr.Pointer() != nil {
		C.QGraphicsRotation_SetAxis(ptr.Pointer(), gui.PointerFromQVector3D(axis))
	}
}

func (ptr *QGraphicsRotation) SetOrigin(point gui.QVector3D_ITF) {
	defer qt.Recovering("QGraphicsRotation::setOrigin")

	if ptr.Pointer() != nil {
		C.QGraphicsRotation_SetOrigin(ptr.Pointer(), gui.PointerFromQVector3D(point))
	}
}

func NewQGraphicsRotation(parent core.QObject_ITF) *QGraphicsRotation {
	defer qt.Recovering("QGraphicsRotation::QGraphicsRotation")

	return NewQGraphicsRotationFromPointer(C.QGraphicsRotation_NewQGraphicsRotation(core.PointerFromQObject(parent)))
}

func (ptr *QGraphicsRotation) ConnectAngleChanged(f func()) {
	defer qt.Recovering("connect QGraphicsRotation::angleChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsRotation_ConnectAngleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "angleChanged", f)
	}
}

func (ptr *QGraphicsRotation) DisconnectAngleChanged() {
	defer qt.Recovering("disconnect QGraphicsRotation::angleChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsRotation_DisconnectAngleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "angleChanged")
	}
}

//export callbackQGraphicsRotationAngleChanged
func callbackQGraphicsRotationAngleChanged(ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsRotation::angleChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "angleChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsRotation) ConnectApplyTo(f func(matrix *gui.QMatrix4x4)) {
	defer qt.Recovering("connect QGraphicsRotation::applyTo")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "applyTo", f)
	}
}

func (ptr *QGraphicsRotation) DisconnectApplyTo() {
	defer qt.Recovering("disconnect QGraphicsRotation::applyTo")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "applyTo")
	}
}

//export callbackQGraphicsRotationApplyTo
func callbackQGraphicsRotationApplyTo(ptrName *C.char, matrix unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsRotation::applyTo")

	if signal := qt.GetSignal(C.GoString(ptrName), "applyTo"); signal != nil {
		signal.(func(*gui.QMatrix4x4))(gui.NewQMatrix4x4FromPointer(matrix))
		return true
	}
	return false

}

func (ptr *QGraphicsRotation) ConnectAxisChanged(f func()) {
	defer qt.Recovering("connect QGraphicsRotation::axisChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsRotation_ConnectAxisChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "axisChanged", f)
	}
}

func (ptr *QGraphicsRotation) DisconnectAxisChanged() {
	defer qt.Recovering("disconnect QGraphicsRotation::axisChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsRotation_DisconnectAxisChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "axisChanged")
	}
}

//export callbackQGraphicsRotationAxisChanged
func callbackQGraphicsRotationAxisChanged(ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsRotation::axisChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "axisChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsRotation) ConnectOriginChanged(f func()) {
	defer qt.Recovering("connect QGraphicsRotation::originChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsRotation_ConnectOriginChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "originChanged", f)
	}
}

func (ptr *QGraphicsRotation) DisconnectOriginChanged() {
	defer qt.Recovering("disconnect QGraphicsRotation::originChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsRotation_DisconnectOriginChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "originChanged")
	}
}

//export callbackQGraphicsRotationOriginChanged
func callbackQGraphicsRotationOriginChanged(ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsRotation::originChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "originChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsRotation) DestroyQGraphicsRotation() {
	defer qt.Recovering("QGraphicsRotation::~QGraphicsRotation")

	if ptr.Pointer() != nil {
		C.QGraphicsRotation_DestroyQGraphicsRotation(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
