package gui

//#include "gui.h"
import "C"
import (
	"log"
	"unsafe"
)

type QSurface struct {
	ptr unsafe.Pointer
}

type QSurface_ITF interface {
	QSurface_PTR() *QSurface
}

func (p *QSurface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSurface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSurface(ptr QSurface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSurface_PTR().Pointer()
	}
	return nil
}

func NewQSurfaceFromPointer(ptr unsafe.Pointer) *QSurface {
	var n = new(QSurface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSurface) QSurface_PTR() *QSurface {
	return ptr
}

//QSurface::SurfaceClass
type QSurface__SurfaceClass int64

const (
	QSurface__Window    = QSurface__SurfaceClass(0)
	QSurface__Offscreen = QSurface__SurfaceClass(1)
)

//QSurface::SurfaceType
type QSurface__SurfaceType int64

const (
	QSurface__RasterSurface   = QSurface__SurfaceType(0)
	QSurface__OpenGLSurface   = QSurface__SurfaceType(1)
	QSurface__RasterGLSurface = QSurface__SurfaceType(2)
)

func (ptr *QSurface) SupportsOpenGL() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSurface::supportsOpenGL")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSurface_SupportsOpenGL(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSurface) SurfaceClass() QSurface__SurfaceClass {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSurface::surfaceClass")
		}
	}()

	if ptr.Pointer() != nil {
		return QSurface__SurfaceClass(C.QSurface_SurfaceClass(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSurface) SurfaceType() QSurface__SurfaceType {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSurface::surfaceType")
		}
	}()

	if ptr.Pointer() != nil {
		return QSurface__SurfaceType(C.QSurface_SurfaceType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSurface) DestroyQSurface() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSurface::~QSurface")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSurface_DestroyQSurface(ptr.Pointer())
	}
}
