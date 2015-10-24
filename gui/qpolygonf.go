package gui

//#include "qpolygonf.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPolygonF struct {
	core.QVector
}

type QPolygonFITF interface {
	core.QVectorITF
	QPolygonFPTR() *QPolygonF
}

func PointerFromQPolygonF(ptr QPolygonFITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPolygonFPTR().Pointer()
	}
	return nil
}

func QPolygonFFromPointer(ptr unsafe.Pointer) *QPolygonF {
	var n = new(QPolygonF)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPolygonF) QPolygonFPTR() *QPolygonF {
	return ptr
}

func NewQPolygonF6(polygon QPolygonITF) *QPolygonF {
	return QPolygonFFromPointer(unsafe.Pointer(C.QPolygonF_NewQPolygonF6(C.QtObjectPtr(PointerFromQPolygon(polygon)))))
}

func NewQPolygonF5(rectangle core.QRectFITF) *QPolygonF {
	return QPolygonFFromPointer(unsafe.Pointer(C.QPolygonF_NewQPolygonF5(C.QtObjectPtr(core.PointerFromQRectF(rectangle)))))
}

func (ptr *QPolygonF) ContainsPoint(point core.QPointFITF, fillRule core.Qt__FillRule) bool {
	if ptr.Pointer() != nil {
		return C.QPolygonF_ContainsPoint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(point)), C.int(fillRule)) != 0
	}
	return false
}

func NewQPolygonF() *QPolygonF {
	return QPolygonFFromPointer(unsafe.Pointer(C.QPolygonF_NewQPolygonF()))
}

func NewQPolygonF3(polygon QPolygonFITF) *QPolygonF {
	return QPolygonFFromPointer(unsafe.Pointer(C.QPolygonF_NewQPolygonF3(C.QtObjectPtr(PointerFromQPolygonF(polygon)))))
}

func NewQPolygonF2(size int) *QPolygonF {
	return QPolygonFFromPointer(unsafe.Pointer(C.QPolygonF_NewQPolygonF2(C.int(size))))
}

func (ptr *QPolygonF) IsClosed() bool {
	if ptr.Pointer() != nil {
		return C.QPolygonF_IsClosed(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPolygonF) Swap(other QPolygonFITF) {
	if ptr.Pointer() != nil {
		C.QPolygonF_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPolygonF(other)))
	}
}

func (ptr *QPolygonF) Translate(offset core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QPolygonF_Translate(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(offset)))
	}
}

func (ptr *QPolygonF) DestroyQPolygonF() {
	if ptr.Pointer() != nil {
		C.QPolygonF_DestroyQPolygonF(C.QtObjectPtr(ptr.Pointer()))
	}
}
