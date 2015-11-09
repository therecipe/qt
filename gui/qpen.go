package gui

//#include "qpen.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPen struct {
	ptr unsafe.Pointer
}

type QPen_ITF interface {
	QPen_PTR() *QPen
}

func (p *QPen) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPen) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPen(ptr QPen_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPen_PTR().Pointer()
	}
	return nil
}

func NewQPenFromPointer(ptr unsafe.Pointer) *QPen {
	var n = new(QPen)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPen) QPen_PTR() *QPen {
	return ptr
}

func NewQPen4(brush QBrush_ITF, width float64, style core.Qt__PenStyle, cap core.Qt__PenCapStyle, join core.Qt__PenJoinStyle) *QPen {
	return NewQPenFromPointer(C.QPen_NewQPen4(PointerFromQBrush(brush), C.double(width), C.int(style), C.int(cap), C.int(join)))
}

func NewQPen5(pen QPen_ITF) *QPen {
	return NewQPenFromPointer(C.QPen_NewQPen5(PointerFromQPen(pen)))
}

func (ptr *QPen) Color() *QColor {
	if ptr.Pointer() != nil {
		return NewQColorFromPointer(C.QPen_Color(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPen) SetCapStyle(style core.Qt__PenCapStyle) {
	if ptr.Pointer() != nil {
		C.QPen_SetCapStyle(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QPen) SetColor(color QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QPen_SetColor(ptr.Pointer(), PointerFromQColor(color))
	}
}

func (ptr *QPen) SetJoinStyle(style core.Qt__PenJoinStyle) {
	if ptr.Pointer() != nil {
		C.QPen_SetJoinStyle(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QPen) SetStyle(style core.Qt__PenStyle) {
	if ptr.Pointer() != nil {
		C.QPen_SetStyle(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QPen) SetWidth(width int) {
	if ptr.Pointer() != nil {
		C.QPen_SetWidth(ptr.Pointer(), C.int(width))
	}
}

func (ptr *QPen) Style() core.Qt__PenStyle {
	if ptr.Pointer() != nil {
		return core.Qt__PenStyle(C.QPen_Style(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPen) Width() int {
	if ptr.Pointer() != nil {
		return int(C.QPen_Width(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPen) WidthF() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPen_WidthF(ptr.Pointer()))
	}
	return 0
}

func NewQPen() *QPen {
	return NewQPenFromPointer(C.QPen_NewQPen())
}

func NewQPen6(pen QPen_ITF) *QPen {
	return NewQPenFromPointer(C.QPen_NewQPen6(PointerFromQPen(pen)))
}

func NewQPen2(style core.Qt__PenStyle) *QPen {
	return NewQPenFromPointer(C.QPen_NewQPen2(C.int(style)))
}

func NewQPen3(color QColor_ITF) *QPen {
	return NewQPenFromPointer(C.QPen_NewQPen3(PointerFromQColor(color)))
}

func (ptr *QPen) Brush() *QBrush {
	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QPen_Brush(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPen) CapStyle() core.Qt__PenCapStyle {
	if ptr.Pointer() != nil {
		return core.Qt__PenCapStyle(C.QPen_CapStyle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPen) DashOffset() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPen_DashOffset(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPen) IsCosmetic() bool {
	if ptr.Pointer() != nil {
		return C.QPen_IsCosmetic(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPen) IsSolid() bool {
	if ptr.Pointer() != nil {
		return C.QPen_IsSolid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPen) JoinStyle() core.Qt__PenJoinStyle {
	if ptr.Pointer() != nil {
		return core.Qt__PenJoinStyle(C.QPen_JoinStyle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPen) MiterLimit() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPen_MiterLimit(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPen) SetBrush(brush QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QPen_SetBrush(ptr.Pointer(), PointerFromQBrush(brush))
	}
}

func (ptr *QPen) SetCosmetic(cosmetic bool) {
	if ptr.Pointer() != nil {
		C.QPen_SetCosmetic(ptr.Pointer(), C.int(qt.GoBoolToInt(cosmetic)))
	}
}

func (ptr *QPen) SetDashOffset(offset float64) {
	if ptr.Pointer() != nil {
		C.QPen_SetDashOffset(ptr.Pointer(), C.double(offset))
	}
}

func (ptr *QPen) SetMiterLimit(limit float64) {
	if ptr.Pointer() != nil {
		C.QPen_SetMiterLimit(ptr.Pointer(), C.double(limit))
	}
}

func (ptr *QPen) SetWidthF(width float64) {
	if ptr.Pointer() != nil {
		C.QPen_SetWidthF(ptr.Pointer(), C.double(width))
	}
}

func (ptr *QPen) Swap(other QPen_ITF) {
	if ptr.Pointer() != nil {
		C.QPen_Swap(ptr.Pointer(), PointerFromQPen(other))
	}
}

func (ptr *QPen) DestroyQPen() {
	if ptr.Pointer() != nil {
		C.QPen_DestroyQPen(ptr.Pointer())
	}
}
