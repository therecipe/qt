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

type QOpenGLContextGroupITF interface {
	core.QObjectITF
	QOpenGLContextGroupPTR() *QOpenGLContextGroup
}

func PointerFromQOpenGLContextGroup(ptr QOpenGLContextGroupITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLContextGroupPTR().Pointer()
	}
	return nil
}

func QOpenGLContextGroupFromPointer(ptr unsafe.Pointer) *QOpenGLContextGroup {
	var n = new(QOpenGLContextGroup)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QOpenGLContextGroup_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QOpenGLContextGroup) QOpenGLContextGroupPTR() *QOpenGLContextGroup {
	return ptr
}

func QOpenGLContextGroup_CurrentContextGroup() *QOpenGLContextGroup {
	return QOpenGLContextGroupFromPointer(unsafe.Pointer(C.QOpenGLContextGroup_QOpenGLContextGroup_CurrentContextGroup()))
}
