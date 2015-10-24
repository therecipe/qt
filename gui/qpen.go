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

type QPenITF interface {
	QPenPTR() *QPen
}

func (p *QPen) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPen) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPen(ptr QPenITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPenPTR().Pointer()
	}
	return nil
}

func QPenFromPointer(ptr unsafe.Pointer) *QPen {
	var n = new(QPen)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPen) QPenPTR() *QPen {
	return ptr
}

func NewQPen5(pen QPenITF) *QPen {
	return QPenFromPointer(unsafe.Pointer(C.QPen_NewQPen5(C.QtObjectPtr(PointerFromQPen(pen)))))
}

func (ptr *QPen) SetCapStyle(style core.Qt__PenCapStyle) {
	if ptr.Pointer() != nil {
		C.QPen_SetCapStyle(C.QtObjectPtr(ptr.Pointer()), C.int(style))
	}
}

func (ptr *QPen) SetColor(color QColorITF) {
	if ptr.Pointer() != nil {
		C.QPen_SetColor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQColor(color)))
	}
}

func (ptr *QPen) SetJoinStyle(style core.Qt__PenJoinStyle) {
	if ptr.Pointer() != nil {
		C.QPen_SetJoinStyle(C.QtObjectPtr(ptr.Pointer()), C.int(style))
	}
}

func (ptr *QPen) SetStyle(style core.Qt__PenStyle) {
	if ptr.Pointer() != nil {
		C.QPen_SetStyle(C.QtObjectPtr(ptr.Pointer()), C.int(style))
	}
}

func (ptr *QPen) SetWidth(width int) {
	if ptr.Pointer() != nil {
		C.QPen_SetWidth(C.QtObjectPtr(ptr.Pointer()), C.int(width))
	}
}

func (ptr *QPen) Style() core.Qt__PenStyle {
	if ptr.Pointer() != nil {
		return core.Qt__PenStyle(C.QPen_Style(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPen) Width() int {
	if ptr.Pointer() != nil {
		return int(C.QPen_Width(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQPen() *QPen {
	return QPenFromPointer(unsafe.Pointer(C.QPen_NewQPen()))
}

func NewQPen6(pen QPenITF) *QPen {
	return QPenFromPointer(unsafe.Pointer(C.QPen_NewQPen6(C.QtObjectPtr(PointerFromQPen(pen)))))
}

func NewQPen2(style core.Qt__PenStyle) *QPen {
	return QPenFromPointer(unsafe.Pointer(C.QPen_NewQPen2(C.int(style))))
}

func NewQPen3(color QColorITF) *QPen {
	return QPenFromPointer(unsafe.Pointer(C.QPen_NewQPen3(C.QtObjectPtr(PointerFromQColor(color)))))
}

func (ptr *QPen) CapStyle() core.Qt__PenCapStyle {
	if ptr.Pointer() != nil {
		return core.Qt__PenCapStyle(C.QPen_CapStyle(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPen) IsCosmetic() bool {
	if ptr.Pointer() != nil {
		return C.QPen_IsCosmetic(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPen) IsSolid() bool {
	if ptr.Pointer() != nil {
		return C.QPen_IsSolid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPen) JoinStyle() core.Qt__PenJoinStyle {
	if ptr.Pointer() != nil {
		return core.Qt__PenJoinStyle(C.QPen_JoinStyle(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPen) SetBrush(brush QBrushITF) {
	if ptr.Pointer() != nil {
		C.QPen_SetBrush(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQBrush(brush)))
	}
}

func (ptr *QPen) SetCosmetic(cosmetic bool) {
	if ptr.Pointer() != nil {
		C.QPen_SetCosmetic(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(cosmetic)))
	}
}

func (ptr *QPen) Swap(other QPenITF) {
	if ptr.Pointer() != nil {
		C.QPen_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPen(other)))
	}
}

func (ptr *QPen) DestroyQPen() {
	if ptr.Pointer() != nil {
		C.QPen_DestroyQPen(C.QtObjectPtr(ptr.Pointer()))
	}
}
