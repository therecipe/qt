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

type QOpenGLVertexArrayObjectITF interface {
	core.QObjectITF
	QOpenGLVertexArrayObjectPTR() *QOpenGLVertexArrayObject
}

func PointerFromQOpenGLVertexArrayObject(ptr QOpenGLVertexArrayObjectITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLVertexArrayObjectPTR().Pointer()
	}
	return nil
}

func QOpenGLVertexArrayObjectFromPointer(ptr unsafe.Pointer) *QOpenGLVertexArrayObject {
	var n = new(QOpenGLVertexArrayObject)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QOpenGLVertexArrayObject_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QOpenGLVertexArrayObject) QOpenGLVertexArrayObjectPTR() *QOpenGLVertexArrayObject {
	return ptr
}

func NewQOpenGLVertexArrayObject(parent core.QObjectITF) *QOpenGLVertexArrayObject {
	return QOpenGLVertexArrayObjectFromPointer(unsafe.Pointer(C.QOpenGLVertexArrayObject_NewQOpenGLVertexArrayObject(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QOpenGLVertexArrayObject) Bind() {
	if ptr.Pointer() != nil {
		C.QOpenGLVertexArrayObject_Bind(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLVertexArrayObject) Create() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLVertexArrayObject_Create(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLVertexArrayObject) Destroy() {
	if ptr.Pointer() != nil {
		C.QOpenGLVertexArrayObject_Destroy(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLVertexArrayObject) IsCreated() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLVertexArrayObject_IsCreated(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLVertexArrayObject) Release() {
	if ptr.Pointer() != nil {
		C.QOpenGLVertexArrayObject_Release(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLVertexArrayObject) DestroyQOpenGLVertexArrayObject() {
	if ptr.Pointer() != nil {
		C.QOpenGLVertexArrayObject_DestroyQOpenGLVertexArrayObject(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
