package gui

//#include "qoffscreensurface.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QOffscreenSurface struct {
	core.QObject
	QSurface
}

type QOffscreenSurfaceITF interface {
	core.QObjectITF
	QSurfaceITF
	QOffscreenSurfacePTR() *QOffscreenSurface
}

func (p *QOffscreenSurface) Pointer() unsafe.Pointer {
	return p.QObjectPTR().Pointer()
}

func (p *QOffscreenSurface) SetPointer(ptr unsafe.Pointer) {
	p.QObjectPTR().SetPointer(ptr)
	p.QSurfacePTR().SetPointer(ptr)
}

func PointerFromQOffscreenSurface(ptr QOffscreenSurfaceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOffscreenSurfacePTR().Pointer()
	}
	return nil
}

func QOffscreenSurfaceFromPointer(ptr unsafe.Pointer) *QOffscreenSurface {
	var n = new(QOffscreenSurface)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QOffscreenSurface_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QOffscreenSurface) QOffscreenSurfacePTR() *QOffscreenSurface {
	return ptr
}

func NewQOffscreenSurface(targetScreen QScreenITF) *QOffscreenSurface {
	return QOffscreenSurfaceFromPointer(unsafe.Pointer(C.QOffscreenSurface_NewQOffscreenSurface(C.QtObjectPtr(PointerFromQScreen(targetScreen)))))
}

func (ptr *QOffscreenSurface) Create() {
	if ptr.Pointer() != nil {
		C.QOffscreenSurface_Create(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOffscreenSurface) Destroy() {
	if ptr.Pointer() != nil {
		C.QOffscreenSurface_Destroy(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOffscreenSurface) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QOffscreenSurface_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOffscreenSurface) Screen() *QScreen {
	if ptr.Pointer() != nil {
		return QScreenFromPointer(unsafe.Pointer(C.QOffscreenSurface_Screen(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QOffscreenSurface) ConnectScreenChanged(f func(screen QScreenITF)) {
	if ptr.Pointer() != nil {
		C.QOffscreenSurface_ConnectScreenChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "screenChanged", f)
	}
}

func (ptr *QOffscreenSurface) DisconnectScreenChanged() {
	if ptr.Pointer() != nil {
		C.QOffscreenSurface_DisconnectScreenChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "screenChanged")
	}
}

//export callbackQOffscreenSurfaceScreenChanged
func callbackQOffscreenSurfaceScreenChanged(ptrName *C.char, screen unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "screenChanged").(func(*QScreen))(QScreenFromPointer(screen))
}

func (ptr *QOffscreenSurface) SetFormat(format QSurfaceFormatITF) {
	if ptr.Pointer() != nil {
		C.QOffscreenSurface_SetFormat(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSurfaceFormat(format)))
	}
}

func (ptr *QOffscreenSurface) SetScreen(newScreen QScreenITF) {
	if ptr.Pointer() != nil {
		C.QOffscreenSurface_SetScreen(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQScreen(newScreen)))
	}
}

func (ptr *QOffscreenSurface) SurfaceType() QSurface__SurfaceType {
	if ptr.Pointer() != nil {
		return QSurface__SurfaceType(C.QOffscreenSurface_SurfaceType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOffscreenSurface) DestroyQOffscreenSurface() {
	if ptr.Pointer() != nil {
		C.QOffscreenSurface_DestroyQOffscreenSurface(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
