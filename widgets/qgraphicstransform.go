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

type QGraphicsTransformITF interface {
	core.QObjectITF
	QGraphicsTransformPTR() *QGraphicsTransform
}

func PointerFromQGraphicsTransform(ptr QGraphicsTransformITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsTransformPTR().Pointer()
	}
	return nil
}

func QGraphicsTransformFromPointer(ptr unsafe.Pointer) *QGraphicsTransform {
	var n = new(QGraphicsTransform)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGraphicsTransform_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsTransform) QGraphicsTransformPTR() *QGraphicsTransform {
	return ptr
}
