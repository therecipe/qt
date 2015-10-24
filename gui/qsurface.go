package gui

//#include "qsurface.h"
import "C"
import (
	"unsafe"
)

type QSurface struct {
	ptr unsafe.Pointer
}

type QSurfaceITF interface {
	QSurfacePTR() *QSurface
}

func (p *QSurface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSurface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSurface(ptr QSurfaceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSurfacePTR().Pointer()
	}
	return nil
}

func QSurfaceFromPointer(ptr unsafe.Pointer) *QSurface {
	var n = new(QSurface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSurface) QSurfacePTR() *QSurface {
	return ptr
}

//QSurface::SurfaceClass
type QSurface__SurfaceClass int

var (
	QSurface__Window    = QSurface__SurfaceClass(0)
	QSurface__Offscreen = QSurface__SurfaceClass(1)
)

//QSurface::SurfaceType
type QSurface__SurfaceType int

var (
	QSurface__RasterSurface   = QSurface__SurfaceType(0)
	QSurface__OpenGLSurface   = QSurface__SurfaceType(1)
	QSurface__RasterGLSurface = QSurface__SurfaceType(2)
)

func (ptr *QSurface) SupportsOpenGL() bool {
	if ptr.Pointer() != nil {
		return C.QSurface_SupportsOpenGL(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSurface) SurfaceClass() QSurface__SurfaceClass {
	if ptr.Pointer() != nil {
		return QSurface__SurfaceClass(C.QSurface_SurfaceClass(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSurface) SurfaceType() QSurface__SurfaceType {
	if ptr.Pointer() != nil {
		return QSurface__SurfaceType(C.QSurface_SurfaceType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSurface) DestroyQSurface() {
	if ptr.Pointer() != nil {
		C.QSurface_DestroyQSurface(C.QtObjectPtr(ptr.Pointer()))
	}
}
