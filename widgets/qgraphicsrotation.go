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
func callbackQGraphicsRotationAngleChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsRotation::angleChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "angleChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsRotation) AngleChanged() {
	defer qt.Recovering("QGraphicsRotation::angleChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsRotation_AngleChanged(ptr.Pointer())
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
func callbackQGraphicsRotationApplyTo(ptr unsafe.Pointer, ptrName *C.char, matrix unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsRotation::applyTo")

	if signal := qt.GetSignal(C.GoString(ptrName), "applyTo"); signal != nil {
		signal.(func(*gui.QMatrix4x4))(gui.NewQMatrix4x4FromPointer(matrix))
	} else {
		NewQGraphicsRotationFromPointer(ptr).ApplyToDefault(gui.NewQMatrix4x4FromPointer(matrix))
	}
}

func (ptr *QGraphicsRotation) ApplyTo(matrix gui.QMatrix4x4_ITF) {
	defer qt.Recovering("QGraphicsRotation::applyTo")

	if ptr.Pointer() != nil {
		C.QGraphicsRotation_ApplyTo(ptr.Pointer(), gui.PointerFromQMatrix4x4(matrix))
	}
}

func (ptr *QGraphicsRotation) ApplyToDefault(matrix gui.QMatrix4x4_ITF) {
	defer qt.Recovering("QGraphicsRotation::applyTo")

	if ptr.Pointer() != nil {
		C.QGraphicsRotation_ApplyToDefault(ptr.Pointer(), gui.PointerFromQMatrix4x4(matrix))
	}
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
func callbackQGraphicsRotationAxisChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsRotation::axisChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "axisChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsRotation) AxisChanged() {
	defer qt.Recovering("QGraphicsRotation::axisChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsRotation_AxisChanged(ptr.Pointer())
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
func callbackQGraphicsRotationOriginChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsRotation::originChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "originChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsRotation) OriginChanged() {
	defer qt.Recovering("QGraphicsRotation::originChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsRotation_OriginChanged(ptr.Pointer())
	}
}

func (ptr *QGraphicsRotation) DestroyQGraphicsRotation() {
	defer qt.Recovering("QGraphicsRotation::~QGraphicsRotation")

	if ptr.Pointer() != nil {
		C.QGraphicsRotation_DestroyQGraphicsRotation(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsRotation) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGraphicsRotation::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGraphicsRotation) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGraphicsRotation::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGraphicsRotationTimerEvent
func callbackQGraphicsRotationTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsRotation::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGraphicsRotationFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGraphicsRotation) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGraphicsRotation::timerEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsRotation_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGraphicsRotation) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGraphicsRotation::timerEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsRotation_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGraphicsRotation) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGraphicsRotation::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGraphicsRotation) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGraphicsRotation::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGraphicsRotationChildEvent
func callbackQGraphicsRotationChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsRotation::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGraphicsRotationFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGraphicsRotation) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGraphicsRotation::childEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsRotation_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGraphicsRotation) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGraphicsRotation::childEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsRotation_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGraphicsRotation) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsRotation::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGraphicsRotation) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGraphicsRotation::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGraphicsRotationCustomEvent
func callbackQGraphicsRotationCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsRotation::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsRotationFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsRotation) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsRotation::customEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsRotation_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsRotation) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsRotation::customEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsRotation_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
