package gui

//#include "qopenglcontextgroup.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QOpenGLContextGroup struct {
	core.QObject
}

type QOpenGLContextGroup_ITF interface {
	core.QObject_ITF
	QOpenGLContextGroup_PTR() *QOpenGLContextGroup
}

func PointerFromQOpenGLContextGroup(ptr QOpenGLContextGroup_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLContextGroup_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLContextGroupFromPointer(ptr unsafe.Pointer) *QOpenGLContextGroup {
	var n = new(QOpenGLContextGroup)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QOpenGLContextGroup_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QOpenGLContextGroup) QOpenGLContextGroup_PTR() *QOpenGLContextGroup {
	return ptr
}
