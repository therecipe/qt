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

type QPainterPathStrokerITF interface {
	QPainterPathStrokerPTR() *QPainterPathStroker
}

func (p *QPainterPathStroker) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPainterPathStroker) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPainterPathStroker(ptr QPainterPathStrokerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPainterPathStrokerPTR().Pointer()
	}
	return nil
}

func QPainterPathStrokerFromPointer(ptr unsafe.Pointer) *QPainterPathStroker {
	var n = new(QPainterPathStroker)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPainterPathStroker) QPainterPathStrokerPTR() *QPainterPathStroker {
	return ptr
}

func NewQPainterPathStroker() *QPainterPathStroker {
	return QPainterPathStrokerFromPointer(unsafe.Pointer(C.QPainterPathStroker_NewQPainterPathStroker()))
}

func NewQPainterPathStroker2(pen QPenITF) *QPainterPathStroker {
	return QPainterPathStrokerFromPointer(unsafe.Pointer(C.QPainterPathStroker_NewQPainterPathStroker2(C.QtObjectPtr(PointerFromQPen(pen)))))
}

func (ptr *QPainterPathStroker) CapStyle() core.Qt__PenCapStyle {
	if ptr.Pointer() != nil {
		return core.Qt__PenCapStyle(C.QPainterPathStroker_CapStyle(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPainterPathStroker) JoinStyle() core.Qt__PenJoinStyle {
	if ptr.Pointer() != nil {
		return core.Qt__PenJoinStyle(C.QPainterPathStroker_JoinStyle(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPainterPathStroker) SetCapStyle(style core.Qt__PenCapStyle) {
	if ptr.Pointer() != nil {
		C.QPainterPathStroker_SetCapStyle(C.QtObjectPtr(ptr.Pointer()), C.int(style))
	}
}

func (ptr *QPainterPathStroker) SetDashPattern(style core.Qt__PenStyle) {
	if ptr.Pointer() != nil {
		C.QPainterPathStroker_SetDashPattern(C.QtObjectPtr(ptr.Pointer()), C.int(style))
	}
}

func (ptr *QPainterPathStroker) SetJoinStyle(style core.Qt__PenJoinStyle) {
	if ptr.Pointer() != nil {
		C.QPainterPathStroker_SetJoinStyle(C.QtObjectPtr(ptr.Pointer()), C.int(style))
	}
}

func (ptr *QPainterPathStroker) DestroyQPainterPathStroker() {
	if ptr.Pointer() != nil {
		C.QPainterPathStroker_DestroyQPainterPathStroker(C.QtObjectPtr(ptr.Pointer()))
	}
}
