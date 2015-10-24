package widgets

//#include "qgraphicsobject.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGraphicsObject struct {
	QGraphicsItem
	core.QObject
}

type QGraphicsObjectITF interface {
	QGraphicsItemITF
	core.QObjectITF
	QGraphicsObjectPTR() *QGraphicsObject
}

func (p *QGraphicsObject) Pointer() unsafe.Pointer {
	return p.QGraphicsItemPTR().Pointer()
}

func (p *QGraphicsObject) SetPointer(ptr unsafe.Pointer) {
	p.QGraphicsItemPTR().SetPointer(ptr)
	p.QObjectPTR().SetPointer(ptr)
}

func PointerFromQGraphicsObject(ptr QGraphicsObjectITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsObjectPTR().Pointer()
	}
	return nil
}

func QGraphicsObjectFromPointer(ptr unsafe.Pointer) *QGraphicsObject {
	var n = new(QGraphicsObject)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGraphicsObject_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsObject) QGraphicsObjectPTR() *QGraphicsObject {
	return ptr
}
