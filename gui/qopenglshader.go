package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QOpenGLShader struct {
	core.QObject
}

type QOpenGLShader_ITF interface {
	core.QObject_ITF
	QOpenGLShader_PTR() *QOpenGLShader
}

func PointerFromQOpenGLShader(ptr QOpenGLShader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLShader_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLShaderFromPointer(ptr unsafe.Pointer) *QOpenGLShader {
	var n = new(QOpenGLShader)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QOpenGLShader_") {
		n.SetObjectName("QOpenGLShader_" + qt.Identifier())
	}
	return n
}

func (ptr *QOpenGLShader) QOpenGLShader_PTR() *QOpenGLShader {
	return ptr
}

//QOpenGLShader::ShaderTypeBit
type QOpenGLShader__ShaderTypeBit int64

const (
	QOpenGLShader__Vertex                 = QOpenGLShader__ShaderTypeBit(0x0001)
	QOpenGLShader__Fragment               = QOpenGLShader__ShaderTypeBit(0x0002)
	QOpenGLShader__Geometry               = QOpenGLShader__ShaderTypeBit(0x0004)
	QOpenGLShader__TessellationControl    = QOpenGLShader__ShaderTypeBit(0x0008)
	QOpenGLShader__TessellationEvaluation = QOpenGLShader__ShaderTypeBit(0x0010)
	QOpenGLShader__Compute                = QOpenGLShader__ShaderTypeBit(0x0020)
)
