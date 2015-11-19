package widgets

//#include "qrubberband.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QRubberBand struct {
	QWidget
}

type QRubberBand_ITF interface {
	QWidget_ITF
	QRubberBand_PTR() *QRubberBand
}

func PointerFromQRubberBand(ptr QRubberBand_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRubberBand_PTR().Pointer()
	}
	return nil
}

func NewQRubberBandFromPointer(ptr unsafe.Pointer) *QRubberBand {
	var n = new(QRubberBand)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QRubberBand_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QRubberBand) QRubberBand_PTR() *QRubberBand {
	return ptr
}

//QRubberBand::Shape
type QRubberBand__Shape int64

const (
	QRubberBand__Line      = QRubberBand__Shape(0)
	QRubberBand__Rectangle = QRubberBand__Shape(1)
)

func (ptr *QRubberBand) SetGeometry(rect core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QRubberBand_SetGeometry(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func NewQRubberBand(s QRubberBand__Shape, p QWidget_ITF) *QRubberBand {
	return NewQRubberBandFromPointer(C.QRubberBand_NewQRubberBand(C.int(s), PointerFromQWidget(p)))
}

func (ptr *QRubberBand) Move2(p core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QRubberBand_Move2(ptr.Pointer(), core.PointerFromQPoint(p))
	}
}

func (ptr *QRubberBand) Move(x int, y int) {
	if ptr.Pointer() != nil {
		C.QRubberBand_Move(ptr.Pointer(), C.int(x), C.int(y))
	}
}

func (ptr *QRubberBand) Resize2(size core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QRubberBand_Resize2(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QRubberBand) Resize(width int, height int) {
	if ptr.Pointer() != nil {
		C.QRubberBand_Resize(ptr.Pointer(), C.int(width), C.int(height))
	}
}

func (ptr *QRubberBand) SetGeometry2(x int, y int, width int, height int) {
	if ptr.Pointer() != nil {
		C.QRubberBand_SetGeometry2(ptr.Pointer(), C.int(x), C.int(y), C.int(width), C.int(height))
	}
}

func (ptr *QRubberBand) Shape() QRubberBand__Shape {
	if ptr.Pointer() != nil {
		return QRubberBand__Shape(C.QRubberBand_Shape(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRubberBand) DestroyQRubberBand() {
	if ptr.Pointer() != nil {
		C.QRubberBand_DestroyQRubberBand(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
