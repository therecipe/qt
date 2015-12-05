package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt/gui"
	"log"
	"unsafe"
)

type QSGMaterialShader struct {
	ptr unsafe.Pointer
}

type QSGMaterialShader_ITF interface {
	QSGMaterialShader_PTR() *QSGMaterialShader
}

func (p *QSGMaterialShader) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSGMaterialShader) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSGMaterialShader(ptr QSGMaterialShader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGMaterialShader_PTR().Pointer()
	}
	return nil
}

func NewQSGMaterialShaderFromPointer(ptr unsafe.Pointer) *QSGMaterialShader {
	var n = new(QSGMaterialShader)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGMaterialShader) QSGMaterialShader_PTR() *QSGMaterialShader {
	return ptr
}

func (ptr *QSGMaterialShader) Activate() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGMaterialShader::activate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSGMaterialShader_Activate(ptr.Pointer())
	}
}

func (ptr *QSGMaterialShader) Deactivate() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGMaterialShader::deactivate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSGMaterialShader_Deactivate(ptr.Pointer())
	}
}

func (ptr *QSGMaterialShader) Program() *gui.QOpenGLShaderProgram {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGMaterialShader::program")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQOpenGLShaderProgramFromPointer(C.QSGMaterialShader_Program(ptr.Pointer()))
	}
	return nil
}
