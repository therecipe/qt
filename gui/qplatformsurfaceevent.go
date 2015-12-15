package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPlatformSurfaceEvent struct {
	core.QEvent
}

type QPlatformSurfaceEvent_ITF interface {
	core.QEvent_ITF
	QPlatformSurfaceEvent_PTR() *QPlatformSurfaceEvent
}

func PointerFromQPlatformSurfaceEvent(ptr QPlatformSurfaceEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlatformSurfaceEvent_PTR().Pointer()
	}
	return nil
}

func NewQPlatformSurfaceEventFromPointer(ptr unsafe.Pointer) *QPlatformSurfaceEvent {
	var n = new(QPlatformSurfaceEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlatformSurfaceEvent) QPlatformSurfaceEvent_PTR() *QPlatformSurfaceEvent {
	return ptr
}

//QPlatformSurfaceEvent::SurfaceEventType
type QPlatformSurfaceEvent__SurfaceEventType int64

const (
	QPlatformSurfaceEvent__SurfaceCreated            = QPlatformSurfaceEvent__SurfaceEventType(0)
	QPlatformSurfaceEvent__SurfaceAboutToBeDestroyed = QPlatformSurfaceEvent__SurfaceEventType(1)
)

func NewQPlatformSurfaceEvent(surfaceEventType QPlatformSurfaceEvent__SurfaceEventType) *QPlatformSurfaceEvent {
	defer qt.Recovering("QPlatformSurfaceEvent::QPlatformSurfaceEvent")

	return NewQPlatformSurfaceEventFromPointer(C.QPlatformSurfaceEvent_NewQPlatformSurfaceEvent(C.int(surfaceEventType)))
}

func (ptr *QPlatformSurfaceEvent) SurfaceEventType() QPlatformSurfaceEvent__SurfaceEventType {
	defer qt.Recovering("QPlatformSurfaceEvent::surfaceEventType")

	if ptr.Pointer() != nil {
		return QPlatformSurfaceEvent__SurfaceEventType(C.QPlatformSurfaceEvent_SurfaceEventType(ptr.Pointer()))
	}
	return 0
}
