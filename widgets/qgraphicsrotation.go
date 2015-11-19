package widgets

//#include "qgraphicsrotation.h"
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
	if n.ObjectName() == "" {
		n.SetObjectName("QGraphicsRotation_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsRotation) QGraphicsRotation_PTR() *QGraphicsRotation {
	return ptr
}

func (ptr *QGraphicsRotation) Angle() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGraphicsRotation_Angle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsRotation) SetAngle(v float64) {
	if ptr.Pointer() != nil {
		C.QGraphicsRotation_SetAngle(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QGraphicsRotation) SetAxis2(axis core.Qt__Axis) {
	if ptr.Pointer() != nil {
		C.QGraphicsRotation_SetAxis2(ptr.Pointer(), C.int(axis))
	}
}

func (ptr *QGraphicsRotation) SetAxis(axis gui.QVector3D_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsRotation_SetAxis(ptr.Pointer(), gui.PointerFromQVector3D(axis))
	}
}

func (ptr *QGraphicsRotation) SetOrigin(point gui.QVector3D_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsRotation_SetOrigin(ptr.Pointer(), gui.PointerFromQVector3D(point))
	}
}

func NewQGraphicsRotation(parent core.QObject_ITF) *QGraphicsRotation {
	return NewQGraphicsRotationFromPointer(C.QGraphicsRotation_NewQGraphicsRotation(core.PointerFromQObject(parent)))
}

func (ptr *QGraphicsRotation) ConnectAngleChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QGraphicsRotation_ConnectAngleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "angleChanged", f)
	}
}

func (ptr *QGraphicsRotation) DisconnectAngleChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsRotation_DisconnectAngleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "angleChanged")
	}
}

//export callbackQGraphicsRotationAngleChanged
func callbackQGraphicsRotationAngleChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "angleChanged").(func())()
}

func (ptr *QGraphicsRotation) ApplyTo(matrix gui.QMatrix4x4_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsRotation_ApplyTo(ptr.Pointer(), gui.PointerFromQMatrix4x4(matrix))
	}
}

func (ptr *QGraphicsRotation) ConnectAxisChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QGraphicsRotation_ConnectAxisChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "axisChanged", f)
	}
}

func (ptr *QGraphicsRotation) DisconnectAxisChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsRotation_DisconnectAxisChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "axisChanged")
	}
}

//export callbackQGraphicsRotationAxisChanged
func callbackQGraphicsRotationAxisChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "axisChanged").(func())()
}

func (ptr *QGraphicsRotation) ConnectOriginChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QGraphicsRotation_ConnectOriginChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "originChanged", f)
	}
}

func (ptr *QGraphicsRotation) DisconnectOriginChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsRotation_DisconnectOriginChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "originChanged")
	}
}

//export callbackQGraphicsRotationOriginChanged
func callbackQGraphicsRotationOriginChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "originChanged").(func())()
}

func (ptr *QGraphicsRotation) DestroyQGraphicsRotation() {
	if ptr.Pointer() != nil {
		C.QGraphicsRotation_DestroyQGraphicsRotation(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
