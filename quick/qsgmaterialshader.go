package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/gui"
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
	for len(n.ObjectNameAbs()) < len("QSGMaterialShader_") {
		n.SetObjectNameAbs("QSGMaterialShader_" + qt.Identifier())
	}
	return n
}

func (ptr *QSGMaterialShader) QSGMaterialShader_PTR() *QSGMaterialShader {
	return ptr
}

func (ptr *QSGMaterialShader) ConnectActivate(f func()) {
	defer qt.Recovering("connect QSGMaterialShader::activate")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "activate", f)
	}
}

func (ptr *QSGMaterialShader) DisconnectActivate() {
	defer qt.Recovering("disconnect QSGMaterialShader::activate")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "activate")
	}
}

//export callbackQSGMaterialShaderActivate
func callbackQSGMaterialShaderActivate(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGMaterialShader::activate")

	if signal := qt.GetSignal(C.GoString(ptrName), "activate"); signal != nil {
		signal.(func())()
	} else {
		NewQSGMaterialShaderFromPointer(ptr).ActivateDefault()
	}
}

func (ptr *QSGMaterialShader) Activate() {
	defer qt.Recovering("QSGMaterialShader::activate")

	if ptr.Pointer() != nil {
		C.QSGMaterialShader_Activate(ptr.Pointer())
	}
}

func (ptr *QSGMaterialShader) ActivateDefault() {
	defer qt.Recovering("QSGMaterialShader::activate")

	if ptr.Pointer() != nil {
		C.QSGMaterialShader_ActivateDefault(ptr.Pointer())
	}
}

func (ptr *QSGMaterialShader) ConnectCompile(f func()) {
	defer qt.Recovering("connect QSGMaterialShader::compile")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "compile", f)
	}
}

func (ptr *QSGMaterialShader) DisconnectCompile() {
	defer qt.Recovering("disconnect QSGMaterialShader::compile")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "compile")
	}
}

//export callbackQSGMaterialShaderCompile
func callbackQSGMaterialShaderCompile(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGMaterialShader::compile")

	if signal := qt.GetSignal(C.GoString(ptrName), "compile"); signal != nil {
		signal.(func())()
	} else {
		NewQSGMaterialShaderFromPointer(ptr).CompileDefault()
	}
}

func (ptr *QSGMaterialShader) Compile() {
	defer qt.Recovering("QSGMaterialShader::compile")

	if ptr.Pointer() != nil {
		C.QSGMaterialShader_Compile(ptr.Pointer())
	}
}

func (ptr *QSGMaterialShader) CompileDefault() {
	defer qt.Recovering("QSGMaterialShader::compile")

	if ptr.Pointer() != nil {
		C.QSGMaterialShader_CompileDefault(ptr.Pointer())
	}
}

func (ptr *QSGMaterialShader) ConnectDeactivate(f func()) {
	defer qt.Recovering("connect QSGMaterialShader::deactivate")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "deactivate", f)
	}
}

func (ptr *QSGMaterialShader) DisconnectDeactivate() {
	defer qt.Recovering("disconnect QSGMaterialShader::deactivate")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "deactivate")
	}
}

//export callbackQSGMaterialShaderDeactivate
func callbackQSGMaterialShaderDeactivate(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGMaterialShader::deactivate")

	if signal := qt.GetSignal(C.GoString(ptrName), "deactivate"); signal != nil {
		signal.(func())()
	} else {
		NewQSGMaterialShaderFromPointer(ptr).DeactivateDefault()
	}
}

func (ptr *QSGMaterialShader) Deactivate() {
	defer qt.Recovering("QSGMaterialShader::deactivate")

	if ptr.Pointer() != nil {
		C.QSGMaterialShader_Deactivate(ptr.Pointer())
	}
}

func (ptr *QSGMaterialShader) DeactivateDefault() {
	defer qt.Recovering("QSGMaterialShader::deactivate")

	if ptr.Pointer() != nil {
		C.QSGMaterialShader_DeactivateDefault(ptr.Pointer())
	}
}

func (ptr *QSGMaterialShader) ConnectInitialize(f func()) {
	defer qt.Recovering("connect QSGMaterialShader::initialize")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "initialize", f)
	}
}

func (ptr *QSGMaterialShader) DisconnectInitialize() {
	defer qt.Recovering("disconnect QSGMaterialShader::initialize")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "initialize")
	}
}

//export callbackQSGMaterialShaderInitialize
func callbackQSGMaterialShaderInitialize(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGMaterialShader::initialize")

	if signal := qt.GetSignal(C.GoString(ptrName), "initialize"); signal != nil {
		signal.(func())()
	} else {
		NewQSGMaterialShaderFromPointer(ptr).InitializeDefault()
	}
}

func (ptr *QSGMaterialShader) Initialize() {
	defer qt.Recovering("QSGMaterialShader::initialize")

	if ptr.Pointer() != nil {
		C.QSGMaterialShader_Initialize(ptr.Pointer())
	}
}

func (ptr *QSGMaterialShader) InitializeDefault() {
	defer qt.Recovering("QSGMaterialShader::initialize")

	if ptr.Pointer() != nil {
		C.QSGMaterialShader_InitializeDefault(ptr.Pointer())
	}
}

func (ptr *QSGMaterialShader) Program() *gui.QOpenGLShaderProgram {
	defer qt.Recovering("QSGMaterialShader::program")

	if ptr.Pointer() != nil {
		return gui.NewQOpenGLShaderProgramFromPointer(C.QSGMaterialShader_Program(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGMaterialShader) ObjectNameAbs() string {
	defer qt.Recovering("QSGMaterialShader::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSGMaterialShader_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSGMaterialShader) SetObjectNameAbs(name string) {
	defer qt.Recovering("QSGMaterialShader::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QSGMaterialShader_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
