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

type QOffscreenSurface_ITF interface {
	core.QObject_ITF
	QSurface_ITF
	QOffscreenSurface_PTR() *QOffscreenSurface
}

func (p *QOffscreenSurface) Pointer() unsafe.Pointer {
	return p.QObject_PTR().Pointer()
}

func (p *QOffscreenSurface) SetPointer(ptr unsafe.Pointer) {
	p.QObject_PTR().SetPointer(ptr)
	p.QSurface_PTR().SetPointer(ptr)
}

func PointerFromQOffscreenSurface(ptr QOffscreenSurface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOffscreenSurface_PTR().Pointer()
	}
	return nil
}

func NewQOffscreenSurfaceFromPointer(ptr unsafe.Pointer) *QOffscreenSurface {
	var n = new(QOffscreenSurface)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QOffscreenSurface_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QOffscreenSurface) QOffscreenSurface_PTR() *QOffscreenSurface {
	return ptr
}

func NewQOffscreenSurface(targetScreen QScreen_ITF) *QOffscreenSurface {
	return NewQOffscreenSurfaceFromPointer(C.QOffscreenSurface_NewQOffscreenSurface(PointerFromQScreen(targetScreen)))
}

func (ptr *QOffscreenSurface) Create() {
	if ptr.Pointer() != nil {
		C.QOffscreenSurface_Create(ptr.Pointer())
	}
}

func (ptr *QOffscreenSurface) Destroy() {
	if ptr.Pointer() != nil {
		C.QOffscreenSurface_Destroy(ptr.Pointer())
	}
}

func (ptr *QOffscreenSurface) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QOffscreenSurface_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QOffscreenSurface) Screen() *QScreen {
	if ptr.Pointer() != nil {
		return NewQScreenFromPointer(C.QOffscreenSurface_Screen(ptr.Pointer()))
	}
	return nil
}

func (ptr *QOffscreenSurface) ConnectScreenChanged(f func(screen *QScreen)) {
	if ptr.Pointer() != nil {
		C.QOffscreenSurface_ConnectScreenChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "screenChanged", f)
	}
}

func (ptr *QOffscreenSurface) DisconnectScreenChanged() {
	if ptr.Pointer() != nil {
		C.QOffscreenSurface_DisconnectScreenChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "screenChanged")
	}
}

//export callbackQOffscreenSurfaceScreenChanged
func callbackQOffscreenSurfaceScreenChanged(ptrName *C.char, screen unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "screenChanged").(func(*QScreen))(NewQScreenFromPointer(screen))
}

func (ptr *QOffscreenSurface) SetFormat(format QSurfaceFormat_ITF) {
	if ptr.Pointer() != nil {
		C.QOffscreenSurface_SetFormat(ptr.Pointer(), PointerFromQSurfaceFormat(format))
	}
}

func (ptr *QOffscreenSurface) SetScreen(newScreen QScreen_ITF) {
	if ptr.Pointer() != nil {
		C.QOffscreenSurface_SetScreen(ptr.Pointer(), PointerFromQScreen(newScreen))
	}
}

func (ptr *QOffscreenSurface) SurfaceType() QSurface__SurfaceType {
	if ptr.Pointer() != nil {
		return QSurface__SurfaceType(C.QOffscreenSurface_SurfaceType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QOffscreenSurface) DestroyQOffscreenSurface() {
	if ptr.Pointer() != nil {
		C.QOffscreenSurface_DestroyQOffscreenSurface(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
