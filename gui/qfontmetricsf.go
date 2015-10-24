package gui

//#include "qfontmetricsf.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QFontMetricsF struct {
	ptr unsafe.Pointer
}

type QFontMetricsFITF interface {
	QFontMetricsFPTR() *QFontMetricsF
}

func (p *QFontMetricsF) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QFontMetricsF) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQFontMetricsF(ptr QFontMetricsFITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFontMetricsFPTR().Pointer()
	}
	return nil
}

func QFontMetricsFFromPointer(ptr unsafe.Pointer) *QFontMetricsF {
	var n = new(QFontMetricsF)
	n.SetPointer(ptr)
	return n
}

func (ptr *QFontMetricsF) QFontMetricsFPTR() *QFontMetricsF {
	return ptr
}

func NewQFontMetricsF(font QFontITF) *QFontMetricsF {
	return QFontMetricsFFromPointer(unsafe.Pointer(C.QFontMetricsF_NewQFontMetricsF(C.QtObjectPtr(PointerFromQFont(font)))))
}

func NewQFontMetricsF2(font QFontITF, paintdevice QPaintDeviceITF) *QFontMetricsF {
	return QFontMetricsFFromPointer(unsafe.Pointer(C.QFontMetricsF_NewQFontMetricsF2(C.QtObjectPtr(PointerFromQFont(font)), C.QtObjectPtr(PointerFromQPaintDevice(paintdevice)))))
}

func NewQFontMetricsF3(fontMetrics QFontMetricsITF) *QFontMetricsF {
	return QFontMetricsFFromPointer(unsafe.Pointer(C.QFontMetricsF_NewQFontMetricsF3(C.QtObjectPtr(PointerFromQFontMetrics(fontMetrics)))))
}

func NewQFontMetricsF4(fm QFontMetricsFITF) *QFontMetricsF {
	return QFontMetricsFFromPointer(unsafe.Pointer(C.QFontMetricsF_NewQFontMetricsF4(C.QtObjectPtr(PointerFromQFontMetricsF(fm)))))
}

func (ptr *QFontMetricsF) InFont(ch core.QCharITF) bool {
	if ptr.Pointer() != nil {
		return C.QFontMetricsF_InFont(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQChar(ch))) != 0
	}
	return false
}

func (ptr *QFontMetricsF) Swap(other QFontMetricsFITF) {
	if ptr.Pointer() != nil {
		C.QFontMetricsF_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQFontMetricsF(other)))
	}
}

func (ptr *QFontMetricsF) DestroyQFontMetricsF() {
	if ptr.Pointer() != nil {
		C.QFontMetricsF_DestroyQFontMetricsF(C.QtObjectPtr(ptr.Pointer()))
	}
}
