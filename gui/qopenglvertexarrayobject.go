package gui

//#include "qopenglvertexarrayobject.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QOpenGLVertexArrayObject struct {
	core.QObject
}

type QOpenGLVertexArrayObject_ITF interface {
	core.QObject_ITF
	QOpenGLVertexArrayObject_PTR() *QOpenGLVertexArrayObject
}

func PointerFromQOpenGLVertexArrayObject(ptr QOpenGLVertexArrayObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLVertexArrayObject_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLVertexArrayObjectFromPointer(ptr unsafe.Pointer) *QOpenGLVertexArrayObject {
	var n = new(QOpenGLVertexArrayObject)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QOpenGLVertexArrayObject_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QOpenGLVertexArrayObject) QOpenGLVertexArrayObject_PTR() *QOpenGLVertexArrayObject {
	return ptr
}
