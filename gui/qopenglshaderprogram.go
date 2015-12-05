package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QOpenGLShaderProgram struct {
	core.QObject
}

type QOpenGLShaderProgram_ITF interface {
	core.QObject_ITF
	QOpenGLShaderProgram_PTR() *QOpenGLShaderProgram
}

func PointerFromQOpenGLShaderProgram(ptr QOpenGLShaderProgram_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLShaderProgram_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLShaderProgramFromPointer(ptr unsafe.Pointer) *QOpenGLShaderProgram {
	var n = new(QOpenGLShaderProgram)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QOpenGLShaderProgram_") {
		n.SetObjectName("QOpenGLShaderProgram_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QOpenGLShaderProgram) QOpenGLShaderProgram_PTR() *QOpenGLShaderProgram {
	return ptr
}
