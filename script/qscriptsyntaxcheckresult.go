package script

//#include "qscriptsyntaxcheckresult.h"
import "C"
import (
	"unsafe"
)

type QScriptSyntaxCheckResult struct {
	ptr unsafe.Pointer
}

type QScriptSyntaxCheckResultITF interface {
	QScriptSyntaxCheckResultPTR() *QScriptSyntaxCheckResult
}

func (p *QScriptSyntaxCheckResult) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptSyntaxCheckResult) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQScriptSyntaxCheckResult(ptr QScriptSyntaxCheckResultITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptSyntaxCheckResultPTR().Pointer()
	}
	return nil
}

func QScriptSyntaxCheckResultFromPointer(ptr unsafe.Pointer) *QScriptSyntaxCheckResult {
	var n = new(QScriptSyntaxCheckResult)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScriptSyntaxCheckResult) QScriptSyntaxCheckResultPTR() *QScriptSyntaxCheckResult {
	return ptr
}

//QScriptSyntaxCheckResult::State
type QScriptSyntaxCheckResult__State int

var (
	QScriptSyntaxCheckResult__Error        = QScriptSyntaxCheckResult__State(0)
	QScriptSyntaxCheckResult__Intermediate = QScriptSyntaxCheckResult__State(1)
	QScriptSyntaxCheckResult__Valid        = QScriptSyntaxCheckResult__State(2)
)

func NewQScriptSyntaxCheckResult(other QScriptSyntaxCheckResultITF) *QScriptSyntaxCheckResult {
	return QScriptSyntaxCheckResultFromPointer(unsafe.Pointer(C.QScriptSyntaxCheckResult_NewQScriptSyntaxCheckResult(C.QtObjectPtr(PointerFromQScriptSyntaxCheckResult(other)))))
}

func (ptr *QScriptSyntaxCheckResult) ErrorColumnNumber() int {
	if ptr.Pointer() != nil {
		return int(C.QScriptSyntaxCheckResult_ErrorColumnNumber(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScriptSyntaxCheckResult) ErrorLineNumber() int {
	if ptr.Pointer() != nil {
		return int(C.QScriptSyntaxCheckResult_ErrorLineNumber(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScriptSyntaxCheckResult) ErrorMessage() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptSyntaxCheckResult_ErrorMessage(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QScriptSyntaxCheckResult) DestroyQScriptSyntaxCheckResult() {
	if ptr.Pointer() != nil {
		C.QScriptSyntaxCheckResult_DestroyQScriptSyntaxCheckResult(C.QtObjectPtr(ptr.Pointer()))
	}
}
