package script

//#include "script.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QScriptProgram struct {
	ptr unsafe.Pointer
}

type QScriptProgram_ITF interface {
	QScriptProgram_PTR() *QScriptProgram
}

func (p *QScriptProgram) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptProgram) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQScriptProgram(ptr QScriptProgram_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptProgram_PTR().Pointer()
	}
	return nil
}

func NewQScriptProgramFromPointer(ptr unsafe.Pointer) *QScriptProgram {
	var n = new(QScriptProgram)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScriptProgram) QScriptProgram_PTR() *QScriptProgram {
	return ptr
}

func NewQScriptProgram() *QScriptProgram {
	defer qt.Recovering("QScriptProgram::QScriptProgram")

	return NewQScriptProgramFromPointer(C.QScriptProgram_NewQScriptProgram())
}

func NewQScriptProgram3(other QScriptProgram_ITF) *QScriptProgram {
	defer qt.Recovering("QScriptProgram::QScriptProgram")

	return NewQScriptProgramFromPointer(C.QScriptProgram_NewQScriptProgram3(PointerFromQScriptProgram(other)))
}

func NewQScriptProgram2(sourceCode string, fileName string, firstLineNumber int) *QScriptProgram {
	defer qt.Recovering("QScriptProgram::QScriptProgram")

	return NewQScriptProgramFromPointer(C.QScriptProgram_NewQScriptProgram2(C.CString(sourceCode), C.CString(fileName), C.int(firstLineNumber)))
}

func (ptr *QScriptProgram) FileName() string {
	defer qt.Recovering("QScriptProgram::fileName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptProgram_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptProgram) FirstLineNumber() int {
	defer qt.Recovering("QScriptProgram::firstLineNumber")

	if ptr.Pointer() != nil {
		return int(C.QScriptProgram_FirstLineNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptProgram) IsNull() bool {
	defer qt.Recovering("QScriptProgram::isNull")

	if ptr.Pointer() != nil {
		return C.QScriptProgram_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptProgram) SourceCode() string {
	defer qt.Recovering("QScriptProgram::sourceCode")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptProgram_SourceCode(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptProgram) DestroyQScriptProgram() {
	defer qt.Recovering("QScriptProgram::~QScriptProgram")

	if ptr.Pointer() != nil {
		C.QScriptProgram_DestroyQScriptProgram(ptr.Pointer())
	}
}
