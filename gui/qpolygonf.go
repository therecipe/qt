package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QPolygonF struct {
	core.QVector
}

type QPolygonF_ITF interface {
	core.QVector_ITF
	QPolygonF_PTR() *QPolygonF
}

func PointerFromQPolygonF(ptr QPolygonF_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPolygonF_PTR().Pointer()
	}
	return nil
}

func NewQPolygonFFromPointer(ptr unsafe.Pointer) *QPolygonF {
	var n = new(QPolygonF)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPolygonF) QPolygonF_PTR() *QPolygonF {
	return ptr
}

func NewQPolygonF6(polygon QPolygon_ITF) *QPolygonF {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPolygonF::QPolygonF")
		}
	}()

	return NewQPolygonFFromPointer(C.QPolygonF_NewQPolygonF6(PointerFromQPolygon(polygon)))
}

func NewQPolygonF5(rectangle core.QRectF_ITF) *QPolygonF {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPolygonF::QPolygonF")
		}
	}()

	return NewQPolygonFFromPointer(C.QPolygonF_NewQPolygonF5(core.PointerFromQRectF(rectangle)))
}

func (ptr *QPolygonF) ContainsPoint(point core.QPointF_ITF, fillRule core.Qt__FillRule) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPolygonF::containsPoint")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPolygonF_ContainsPoint(ptr.Pointer(), core.PointerFromQPointF(point), C.int(fillRule)) != 0
	}
	return false
}

func NewQPolygonF() *QPolygonF {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPolygonF::QPolygonF")
		}
	}()

	return NewQPolygonFFromPointer(C.QPolygonF_NewQPolygonF())
}

func NewQPolygonF3(polygon QPolygonF_ITF) *QPolygonF {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPolygonF::QPolygonF")
		}
	}()

	return NewQPolygonFFromPointer(C.QPolygonF_NewQPolygonF3(PointerFromQPolygonF(polygon)))
}

func NewQPolygonF2(size int) *QPolygonF {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPolygonF::QPolygonF")
		}
	}()

	return NewQPolygonFFromPointer(C.QPolygonF_NewQPolygonF2(C.int(size)))
}

func (ptr *QPolygonF) IsClosed() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPolygonF::isClosed")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPolygonF_IsClosed(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPolygonF) Swap(other QPolygonF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPolygonF::swap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPolygonF_Swap(ptr.Pointer(), PointerFromQPolygonF(other))
	}
}

func (ptr *QPolygonF) Translate(offset core.QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPolygonF::translate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPolygonF_Translate(ptr.Pointer(), core.PointerFromQPointF(offset))
	}
}

func (ptr *QPolygonF) Translate2(dx float64, dy float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPolygonF::translate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPolygonF_Translate2(ptr.Pointer(), C.double(dx), C.double(dy))
	}
}

func (ptr *QPolygonF) DestroyQPolygonF() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPolygonF::~QPolygonF")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPolygonF_DestroyQPolygonF(ptr.Pointer())
	}
}
