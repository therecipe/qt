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

type QGraphicsRotationITF interface {
	QGraphicsTransformITF
	QGraphicsRotationPTR() *QGraphicsRotation
}

func PointerFromQGraphicsRotation(ptr QGraphicsRotationITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsRotationPTR().Pointer()
	}
	return nil
}

func QGraphicsRotationFromPointer(ptr unsafe.Pointer) *QGraphicsRotation {
	var n = new(QGraphicsRotation)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGraphicsRotation_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsRotation) QGraphicsRotationPTR() *QGraphicsRotation {
	return ptr
}

func (ptr *QGraphicsRotation) SetAxis2(axis core.Qt__Axis) {
	if ptr.Pointer() != nil {
		C.QGraphicsRotation_SetAxis2(C.QtObjectPtr(ptr.Pointer()), C.int(axis))
	}
}

func (ptr *QGraphicsRotation) SetAxis(axis gui.QVector3DITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsRotation_SetAxis(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQVector3D(axis)))
	}
}

func (ptr *QGraphicsRotation) SetOrigin(point gui.QVector3DITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsRotation_SetOrigin(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQVector3D(point)))
	}
}

func NewQGraphicsRotation(parent core.QObjectITF) *QGraphicsRotation {
	return QGraphicsRotationFromPointer(unsafe.Pointer(C.QGraphicsRotation_NewQGraphicsRotation(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QGraphicsRotation) ConnectAngleChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QGraphicsRotation_ConnectAngleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "angleChanged", f)
	}
}

func (ptr *QGraphicsRotation) DisconnectAngleChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsRotation_DisconnectAngleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "angleChanged")
	}
}

//export callbackQGraphicsRotationAngleChanged
func callbackQGraphicsRotationAngleChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "angleChanged").(func())()
}

func (ptr *QGraphicsRotation) ApplyTo(matrix gui.QMatrix4x4ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsRotation_ApplyTo(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQMatrix4x4(matrix)))
	}
}

func (ptr *QGraphicsRotation) ConnectAxisChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QGraphicsRotation_ConnectAxisChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "axisChanged", f)
	}
}

func (ptr *QGraphicsRotation) DisconnectAxisChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsRotation_DisconnectAxisChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "axisChanged")
	}
}

//export callbackQGraphicsRotationAxisChanged
func callbackQGraphicsRotationAxisChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "axisChanged").(func())()
}

func (ptr *QGraphicsRotation) ConnectOriginChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QGraphicsRotation_ConnectOriginChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "originChanged", f)
	}
}

func (ptr *QGraphicsRotation) DisconnectOriginChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsRotation_DisconnectOriginChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "originChanged")
	}
}

//export callbackQGraphicsRotationOriginChanged
func callbackQGraphicsRotationOriginChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "originChanged").(func())()
}

func (ptr *QGraphicsRotation) DestroyQGraphicsRotation() {
	if ptr.Pointer() != nil {
		C.QGraphicsRotation_DestroyQGraphicsRotation(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
