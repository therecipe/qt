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

type QRubberBandITF interface {
	QWidgetITF
	QRubberBandPTR() *QRubberBand
}

func PointerFromQRubberBand(ptr QRubberBandITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRubberBandPTR().Pointer()
	}
	return nil
}

func QRubberBandFromPointer(ptr unsafe.Pointer) *QRubberBand {
	var n = new(QRubberBand)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QRubberBand_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QRubberBand) QRubberBandPTR() *QRubberBand {
	return ptr
}

//QRubberBand::Shape
type QRubberBand__Shape int

var (
	QRubberBand__Line      = QRubberBand__Shape(0)
	QRubberBand__Rectangle = QRubberBand__Shape(1)
)

func (ptr *QRubberBand) SetGeometry(rect core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QRubberBand_SetGeometry(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rect)))
	}
}

func NewQRubberBand(s QRubberBand__Shape, p QWidgetITF) *QRubberBand {
	return QRubberBandFromPointer(unsafe.Pointer(C.QRubberBand_NewQRubberBand(C.int(s), C.QtObjectPtr(PointerFromQWidget(p)))))
}

func (ptr *QRubberBand) Move2(p core.QPointITF) {
	if ptr.Pointer() != nil {
		C.QRubberBand_Move2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(p)))
	}
}

func (ptr *QRubberBand) Move(x int, y int) {
	if ptr.Pointer() != nil {
		C.QRubberBand_Move(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y))
	}
}

func (ptr *QRubberBand) Resize2(size core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QRubberBand_Resize2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(size)))
	}
}

func (ptr *QRubberBand) Resize(width int, height int) {
	if ptr.Pointer() != nil {
		C.QRubberBand_Resize(C.QtObjectPtr(ptr.Pointer()), C.int(width), C.int(height))
	}
}

func (ptr *QRubberBand) SetGeometry2(x int, y int, width int, height int) {
	if ptr.Pointer() != nil {
		C.QRubberBand_SetGeometry2(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.int(width), C.int(height))
	}
}

func (ptr *QRubberBand) Shape() QRubberBand__Shape {
	if ptr.Pointer() != nil {
		return QRubberBand__Shape(C.QRubberBand_Shape(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRubberBand) DestroyQRubberBand() {
	if ptr.Pointer() != nil {
		C.QRubberBand_DestroyQRubberBand(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
