package widgets

//#include "qgraphicstransform.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGraphicsTransform struct {
	core.QObject
}

type QGraphicsTransform_ITF interface {
	core.QObject_ITF
	QGraphicsTransform_PTR() *QGraphicsTransform
}

func PointerFromQGraphicsTransform(ptr QGraphicsTransform_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsTransform_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsTransformFromPointer(ptr unsafe.Pointer) *QGraphicsTransform {
	var n = new(QGraphicsTransform)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QGraphicsTransform_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsTransform) QGraphicsTransform_PTR() *QGraphicsTransform {
	return ptr
}
