package gui

//#include "qplatformsurfaceevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPlatformSurfaceEvent struct {
	core.QEvent
}

type QPlatformSurfaceEventITF interface {
	core.QEventITF
	QPlatformSurfaceEventPTR() *QPlatformSurfaceEvent
}

func PointerFromQPlatformSurfaceEvent(ptr QPlatformSurfaceEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlatformSurfaceEventPTR().Pointer()
	}
	return nil
}

func QPlatformSurfaceEventFromPointer(ptr unsafe.Pointer) *QPlatformSurfaceEvent {
	var n = new(QPlatformSurfaceEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlatformSurfaceEvent) QPlatformSurfaceEventPTR() *QPlatformSurfaceEvent {
	return ptr
}

//QPlatformSurfaceEvent::SurfaceEventType
type QPlatformSurfaceEvent__SurfaceEventType int

var (
	QPlatformSurfaceEvent__SurfaceCreated            = QPlatformSurfaceEvent__SurfaceEventType(0)
	QPlatformSurfaceEvent__SurfaceAboutToBeDestroyed = QPlatformSurfaceEvent__SurfaceEventType(1)
)

func NewQPlatformSurfaceEvent(surfaceEventType QPlatformSurfaceEvent__SurfaceEventType) *QPlatformSurfaceEvent {
	return QPlatformSurfaceEventFromPointer(unsafe.Pointer(C.QPlatformSurfaceEvent_NewQPlatformSurfaceEvent(C.int(surfaceEventType))))
}

func (ptr *QPlatformSurfaceEvent) SurfaceEventType() QPlatformSurfaceEvent__SurfaceEventType {
	if ptr.Pointer() != nil {
		return QPlatformSurfaceEvent__SurfaceEventType(C.QPlatformSurfaceEvent_SurfaceEventType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
