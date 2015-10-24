package script

//#include "qscriptprogram.h"
import "C"
import (
	"unsafe"
)

type QScriptProgram struct {
	ptr unsafe.Pointer
}

type QScriptProgramITF interface {
	QScriptProgramPTR() *QScriptProgram
}

func (p *QScriptProgram) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptProgram) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQScriptProgram(ptr QScriptProgramITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptProgramPTR().Pointer()
	}
	return nil
}

func QScriptProgramFromPointer(ptr unsafe.Pointer) *QScriptProgram {
	var n = new(QScriptProgram)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScriptProgram) QScriptProgramPTR() *QScriptProgram {
	return ptr
}

func NewQScriptProgram() *QScriptProgram {
	return QScriptProgramFromPointer(unsafe.Pointer(C.QScriptProgram_NewQScriptProgram()))
}

func NewQScriptProgram3(other QScriptProgramITF) *QScriptProgram {
	return QScriptProgramFromPointer(unsafe.Pointer(C.QScriptProgram_NewQScriptProgram3(C.QtObjectPtr(PointerFromQScriptProgram(other)))))
}

func NewQScriptProgram2(sourceCode string, fileName string, firstLineNumber int) *QScriptProgram {
	return QScriptProgramFromPointer(unsafe.Pointer(C.QScriptProgram_NewQScriptProgram2(C.CString(sourceCode), C.CString(fileName), C.int(firstLineNumber))))
}

func (ptr *QScriptProgram) FileName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptProgram_FileName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QScriptProgram) FirstLineNumber() int {
	if ptr.Pointer() != nil {
		return int(C.QScriptProgram_FirstLineNumber(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScriptProgram) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QScriptProgram_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptProgram) SourceCode() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptProgram_SourceCode(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QScriptProgram) DestroyQScriptProgram() {
	if ptr.Pointer() != nil {
		C.QScriptProgram_DestroyQScriptProgram(C.QtObjectPtr(ptr.Pointer()))
	}
}
