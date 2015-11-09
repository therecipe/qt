package gui

//#include "qrasterpaintengine.h"
import "C"
import (
	"unsafe"
)

type QRasterPaintEngine struct {
	QPaintEngine
}

type QRasterPaintEngine_ITF interface {
	QPaintEngine_ITF
	QRasterPaintEngine_PTR() *QRasterPaintEngine
}

func PointerFromQRasterPaintEngine(ptr QRasterPaintEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRasterPaintEngine_PTR().Pointer()
	}
	return nil
}

func NewQRasterPaintEngineFromPointer(ptr unsafe.Pointer) *QRasterPaintEngine {
	var n = new(QRasterPaintEngine)
	n.SetPointer(ptr)
	return n
}

func (ptr *QRasterPaintEngine) QRasterPaintEngine_PTR() *QRasterPaintEngine {
	return ptr
}
