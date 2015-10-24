package gui

//#include "qpaintenginestate.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPaintEngineState struct {
	ptr unsafe.Pointer
}

type QPaintEngineStateITF interface {
	QPaintEngineStatePTR() *QPaintEngineState
}

func (p *QPaintEngineState) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPaintEngineState) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPaintEngineState(ptr QPaintEngineStateITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPaintEngineStatePTR().Pointer()
	}
	return nil
}

func QPaintEngineStateFromPointer(ptr unsafe.Pointer) *QPaintEngineState {
	var n = new(QPaintEngineState)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPaintEngineState) QPaintEngineStatePTR() *QPaintEngineState {
	return ptr
}

func (ptr *QPaintEngineState) BackgroundMode() core.Qt__BGMode {
	if ptr.Pointer() != nil {
		return core.Qt__BGMode(C.QPaintEngineState_BackgroundMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPaintEngineState) BrushNeedsResolving() bool {
	if ptr.Pointer() != nil {
		return C.QPaintEngineState_BrushNeedsResolving(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPaintEngineState) ClipOperation() core.Qt__ClipOperation {
	if ptr.Pointer() != nil {
		return core.Qt__ClipOperation(C.QPaintEngineState_ClipOperation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPaintEngineState) CompositionMode() QPainter__CompositionMode {
	if ptr.Pointer() != nil {
		return QPainter__CompositionMode(C.QPaintEngineState_CompositionMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPaintEngineState) IsClipEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QPaintEngineState_IsClipEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPaintEngineState) Painter() *QPainter {
	if ptr.Pointer() != nil {
		return QPainterFromPointer(unsafe.Pointer(C.QPaintEngineState_Painter(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QPaintEngineState) PenNeedsResolving() bool {
	if ptr.Pointer() != nil {
		return C.QPaintEngineState_PenNeedsResolving(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPaintEngineState) RenderHints() QPainter__RenderHint {
	if ptr.Pointer() != nil {
		return QPainter__RenderHint(C.QPaintEngineState_RenderHints(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
