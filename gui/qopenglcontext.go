package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QOpenGLContext struct {
	core.QObject
}

type QOpenGLContext_ITF interface {
	core.QObject_ITF
	QOpenGLContext_PTR() *QOpenGLContext
}

func PointerFromQOpenGLContext(ptr QOpenGLContext_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLContext_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLContextFromPointer(ptr unsafe.Pointer) *QOpenGLContext {
	var n = new(QOpenGLContext)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QOpenGLContext_") {
		n.SetObjectName("QOpenGLContext_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QOpenGLContext) QOpenGLContext_PTR() *QOpenGLContext {
	return ptr
}

//QOpenGLContext::OpenGLModuleType
type QOpenGLContext__OpenGLModuleType int64

const (
	QOpenGLContext__LibGL   = QOpenGLContext__OpenGLModuleType(0)
	QOpenGLContext__LibGLES = QOpenGLContext__OpenGLModuleType(1)
)
