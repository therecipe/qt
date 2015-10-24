package gui

//#include "qrasterpaintengine.h"
import "C"
import (
	"unsafe"
)

type QRasterPaintEngine struct {
	QPaintEngine
}

type QRasterPaintEngineITF interface {
	QPaintEngineITF
	QRasterPaintEnginePTR() *QRasterPaintEngine
}

func PointerFromQRasterPaintEngine(ptr QRasterPaintEngineITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRasterPaintEnginePTR().Pointer()
	}
	return nil
}

func QRasterPaintEngineFromPointer(ptr unsafe.Pointer) *QRasterPaintEngine {
	var n = new(QRasterPaintEngine)
	n.SetPointer(ptr)
	return n
}

func (ptr *QRasterPaintEngine) QRasterPaintEnginePTR() *QRasterPaintEngine {
	return ptr
}
