package quick

//#include "qsgmaterialshader.h"
import "C"
import (
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QSGMaterialShader struct {
	ptr unsafe.Pointer
}

type QSGMaterialShaderITF interface {
	QSGMaterialShaderPTR() *QSGMaterialShader
}

func (p *QSGMaterialShader) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSGMaterialShader) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSGMaterialShader(ptr QSGMaterialShaderITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGMaterialShaderPTR().Pointer()
	}
	return nil
}

func QSGMaterialShaderFromPointer(ptr unsafe.Pointer) *QSGMaterialShader {
	var n = new(QSGMaterialShader)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGMaterialShader) QSGMaterialShaderPTR() *QSGMaterialShader {
	return ptr
}

func (ptr *QSGMaterialShader) Activate() {
	if ptr.Pointer() != nil {
		C.QSGMaterialShader_Activate(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSGMaterialShader) Deactivate() {
	if ptr.Pointer() != nil {
		C.QSGMaterialShader_Deactivate(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSGMaterialShader) Program() *gui.QOpenGLShaderProgram {
	if ptr.Pointer() != nil {
		return gui.QOpenGLShaderProgramFromPointer(unsafe.Pointer(C.QSGMaterialShader_Program(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}
