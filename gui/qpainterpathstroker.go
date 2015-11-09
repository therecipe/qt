package gui

//#include "qpainterpathstroker.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPainterPathStroker struct {
	ptr unsafe.Pointer
}

type QPainterPathStroker_ITF interface {
	QPainterPathStroker_PTR() *QPainterPathStroker
}

func (p *QPainterPathStroker) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPainterPathStroker) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPainterPathStroker(ptr QPainterPathStroker_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPainterPathStroker_PTR().Pointer()
	}
	return nil
}

func NewQPainterPathStrokerFromPointer(ptr unsafe.Pointer) *QPainterPathStroker {
	var n = new(QPainterPathStroker)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPainterPathStroker) QPainterPathStroker_PTR() *QPainterPathStroker {
	return ptr
}

func NewQPainterPathStroker() *QPainterPathStroker {
	return NewQPainterPathStrokerFromPointer(C.QPainterPathStroker_NewQPainterPathStroker())
}

func NewQPainterPathStroker2(pen QPen_ITF) *QPainterPathStroker {
	return NewQPainterPathStrokerFromPointer(C.QPainterPathStroker_NewQPainterPathStroker2(PointerFromQPen(pen)))
}

func (ptr *QPainterPathStroker) CapStyle() core.Qt__PenCapStyle {
	if ptr.Pointer() != nil {
		return core.Qt__PenCapStyle(C.QPainterPathStroker_CapStyle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPainterPathStroker) CurveThreshold() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPainterPathStroker_CurveThreshold(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPainterPathStroker) DashOffset() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPainterPathStroker_DashOffset(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPainterPathStroker) JoinStyle() core.Qt__PenJoinStyle {
	if ptr.Pointer() != nil {
		return core.Qt__PenJoinStyle(C.QPainterPathStroker_JoinStyle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPainterPathStroker) MiterLimit() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPainterPathStroker_MiterLimit(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPainterPathStroker) SetCapStyle(style core.Qt__PenCapStyle) {
	if ptr.Pointer() != nil {
		C.QPainterPathStroker_SetCapStyle(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QPainterPathStroker) SetCurveThreshold(threshold float64) {
	if ptr.Pointer() != nil {
		C.QPainterPathStroker_SetCurveThreshold(ptr.Pointer(), C.double(threshold))
	}
}

func (ptr *QPainterPathStroker) SetDashOffset(offset float64) {
	if ptr.Pointer() != nil {
		C.QPainterPathStroker_SetDashOffset(ptr.Pointer(), C.double(offset))
	}
}

func (ptr *QPainterPathStroker) SetDashPattern(style core.Qt__PenStyle) {
	if ptr.Pointer() != nil {
		C.QPainterPathStroker_SetDashPattern(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QPainterPathStroker) SetJoinStyle(style core.Qt__PenJoinStyle) {
	if ptr.Pointer() != nil {
		C.QPainterPathStroker_SetJoinStyle(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QPainterPathStroker) SetMiterLimit(limit float64) {
	if ptr.Pointer() != nil {
		C.QPainterPathStroker_SetMiterLimit(ptr.Pointer(), C.double(limit))
	}
}

func (ptr *QPainterPathStroker) SetWidth(width float64) {
	if ptr.Pointer() != nil {
		C.QPainterPathStroker_SetWidth(ptr.Pointer(), C.double(width))
	}
}

func (ptr *QPainterPathStroker) Width() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPainterPathStroker_Width(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPainterPathStroker) DestroyQPainterPathStroker() {
	if ptr.Pointer() != nil {
		C.QPainterPathStroker_DestroyQPainterPathStroker(ptr.Pointer())
	}
}
