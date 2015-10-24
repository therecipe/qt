package script

//#include "qscriptcontextinfo.h"
import "C"
import (
	"strings"
	"unsafe"
)

type QScriptContextInfo struct {
	ptr unsafe.Pointer
}

type QScriptContextInfoITF interface {
	QScriptContextInfoPTR() *QScriptContextInfo
}

func (p *QScriptContextInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptContextInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQScriptContextInfo(ptr QScriptContextInfoITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptContextInfoPTR().Pointer()
	}
	return nil
}

func QScriptContextInfoFromPointer(ptr unsafe.Pointer) *QScriptContextInfo {
	var n = new(QScriptContextInfo)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScriptContextInfo) QScriptContextInfoPTR() *QScriptContextInfo {
	return ptr
}

//QScriptContextInfo::FunctionType
type QScriptContextInfo__FunctionType int

var (
	QScriptContextInfo__ScriptFunction     = QScriptContextInfo__FunctionType(0)
	QScriptContextInfo__QtFunction         = QScriptContextInfo__FunctionType(1)
	QScriptContextInfo__QtPropertyFunction = QScriptContextInfo__FunctionType(2)
	QScriptContextInfo__NativeFunction     = QScriptContextInfo__FunctionType(3)
)

func NewQScriptContextInfo3() *QScriptContextInfo {
	return QScriptContextInfoFromPointer(unsafe.Pointer(C.QScriptContextInfo_NewQScriptContextInfo3()))
}

func NewQScriptContextInfo(context QScriptContextITF) *QScriptContextInfo {
	return QScriptContextInfoFromPointer(unsafe.Pointer(C.QScriptContextInfo_NewQScriptContextInfo(C.QtObjectPtr(PointerFromQScriptContext(context)))))
}

func NewQScriptContextInfo2(other QScriptContextInfoITF) *QScriptContextInfo {
	return QScriptContextInfoFromPointer(unsafe.Pointer(C.QScriptContextInfo_NewQScriptContextInfo2(C.QtObjectPtr(PointerFromQScriptContextInfo(other)))))
}

func (ptr *QScriptContextInfo) FileName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptContextInfo_FileName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QScriptContextInfo) FunctionEndLineNumber() int {
	if ptr.Pointer() != nil {
		return int(C.QScriptContextInfo_FunctionEndLineNumber(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScriptContextInfo) FunctionMetaIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QScriptContextInfo_FunctionMetaIndex(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScriptContextInfo) FunctionName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptContextInfo_FunctionName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QScriptContextInfo) FunctionParameterNames() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QScriptContextInfo_FunctionParameterNames(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QScriptContextInfo) FunctionStartLineNumber() int {
	if ptr.Pointer() != nil {
		return int(C.QScriptContextInfo_FunctionStartLineNumber(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScriptContextInfo) FunctionType() QScriptContextInfo__FunctionType {
	if ptr.Pointer() != nil {
		return QScriptContextInfo__FunctionType(C.QScriptContextInfo_FunctionType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScriptContextInfo) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QScriptContextInfo_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptContextInfo) LineNumber() int {
	if ptr.Pointer() != nil {
		return int(C.QScriptContextInfo_LineNumber(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScriptContextInfo) DestroyQScriptContextInfo() {
	if ptr.Pointer() != nil {
		C.QScriptContextInfo_DestroyQScriptContextInfo(C.QtObjectPtr(ptr.Pointer()))
	}
}
