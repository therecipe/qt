package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QPaintEngineState struct {
	ptr unsafe.Pointer
}

type QPaintEngineState_ITF interface {
	QPaintEngineState_PTR() *QPaintEngineState
}

func (p *QPaintEngineState) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPaintEngineState) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPaintEngineState(ptr QPaintEngineState_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPaintEngineState_PTR().Pointer()
	}
	return nil
}

func NewQPaintEngineStateFromPointer(ptr unsafe.Pointer) *QPaintEngineState {
	var n = new(QPaintEngineState)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPaintEngineState) QPaintEngineState_PTR() *QPaintEngineState {
	return ptr
}

func (ptr *QPaintEngineState) BackgroundBrush() *QBrush {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPaintEngineState::backgroundBrush")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QPaintEngineState_BackgroundBrush(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPaintEngineState) BackgroundMode() core.Qt__BGMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPaintEngineState::backgroundMode")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__BGMode(C.QPaintEngineState_BackgroundMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPaintEngineState) Brush() *QBrush {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPaintEngineState::brush")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QPaintEngineState_Brush(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPaintEngineState) BrushNeedsResolving() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPaintEngineState::brushNeedsResolving")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPaintEngineState_BrushNeedsResolving(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPaintEngineState) ClipOperation() core.Qt__ClipOperation {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPaintEngineState::clipOperation")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__ClipOperation(C.QPaintEngineState_ClipOperation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPaintEngineState) ClipRegion() *QRegion {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPaintEngineState::clipRegion")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQRegionFromPointer(C.QPaintEngineState_ClipRegion(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPaintEngineState) CompositionMode() QPainter__CompositionMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPaintEngineState::compositionMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QPainter__CompositionMode(C.QPaintEngineState_CompositionMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPaintEngineState) IsClipEnabled() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPaintEngineState::isClipEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPaintEngineState_IsClipEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPaintEngineState) Opacity() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPaintEngineState::opacity")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QPaintEngineState_Opacity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPaintEngineState) Painter() *QPainter {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPaintEngineState::painter")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQPainterFromPointer(C.QPaintEngineState_Painter(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPaintEngineState) PenNeedsResolving() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPaintEngineState::penNeedsResolving")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPaintEngineState_PenNeedsResolving(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPaintEngineState) RenderHints() QPainter__RenderHint {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPaintEngineState::renderHints")
		}
	}()

	if ptr.Pointer() != nil {
		return QPainter__RenderHint(C.QPaintEngineState_RenderHints(ptr.Pointer()))
	}
	return 0
}
