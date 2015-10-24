package gui

//#include "qopenglshader.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QOpenGLShader struct {
	core.QObject
}

type QOpenGLShaderITF interface {
	core.QObjectITF
	QOpenGLShaderPTR() *QOpenGLShader
}

func PointerFromQOpenGLShader(ptr QOpenGLShaderITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLShaderPTR().Pointer()
	}
	return nil
}

func QOpenGLShaderFromPointer(ptr unsafe.Pointer) *QOpenGLShader {
	var n = new(QOpenGLShader)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QOpenGLShader_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QOpenGLShader) QOpenGLShaderPTR() *QOpenGLShader {
	return ptr
}

//QOpenGLShader::ShaderTypeBit
type QOpenGLShader__ShaderTypeBit int

var (
	QOpenGLShader__Vertex                 = QOpenGLShader__ShaderTypeBit(0x0001)
	QOpenGLShader__Fragment               = QOpenGLShader__ShaderTypeBit(0x0002)
	QOpenGLShader__Geometry               = QOpenGLShader__ShaderTypeBit(0x0004)
	QOpenGLShader__TessellationControl    = QOpenGLShader__ShaderTypeBit(0x0008)
	QOpenGLShader__TessellationEvaluation = QOpenGLShader__ShaderTypeBit(0x0010)
	QOpenGLShader__Compute                = QOpenGLShader__ShaderTypeBit(0x0020)
)

func NewQOpenGLShader(ty QOpenGLShader__ShaderTypeBit, parent core.QObjectITF) *QOpenGLShader {
	return QOpenGLShaderFromPointer(unsafe.Pointer(C.QOpenGLShader_NewQOpenGLShader(C.int(ty), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QOpenGLShader) CompileSourceCode2(source core.QByteArrayITF) bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLShader_CompileSourceCode2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(source))) != 0
	}
	return false
}

func (ptr *QOpenGLShader) CompileSourceCode3(source string) bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLShader_CompileSourceCode3(C.QtObjectPtr(ptr.Pointer()), C.CString(source)) != 0
	}
	return false
}

func (ptr *QOpenGLShader) CompileSourceCode(source string) bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLShader_CompileSourceCode(C.QtObjectPtr(ptr.Pointer()), C.CString(source)) != 0
	}
	return false
}

func (ptr *QOpenGLShader) CompileSourceFile(fileName string) bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLShader_CompileSourceFile(C.QtObjectPtr(ptr.Pointer()), C.CString(fileName)) != 0
	}
	return false
}

func QOpenGLShader_HasOpenGLShaders(ty QOpenGLShader__ShaderTypeBit, context QOpenGLContextITF) bool {
	return C.QOpenGLShader_QOpenGLShader_HasOpenGLShaders(C.int(ty), C.QtObjectPtr(PointerFromQOpenGLContext(context))) != 0
}

func (ptr *QOpenGLShader) IsCompiled() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLShader_IsCompiled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLShader) Log() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QOpenGLShader_Log(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QOpenGLShader) ShaderType() QOpenGLShader__ShaderTypeBit {
	if ptr.Pointer() != nil {
		return QOpenGLShader__ShaderTypeBit(C.QOpenGLShader_ShaderType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLShader) DestroyQOpenGLShader() {
	if ptr.Pointer() != nil {
		C.QOpenGLShader_DestroyQOpenGLShader(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
