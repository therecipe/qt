package script

//#include "script.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QScriptContextInfo struct {
	ptr unsafe.Pointer
}

type QScriptContextInfo_ITF interface {
	QScriptContextInfo_PTR() *QScriptContextInfo
}

func (p *QScriptContextInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptContextInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQScriptContextInfo(ptr QScriptContextInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptContextInfo_PTR().Pointer()
	}
	return nil
}

func NewQScriptContextInfoFromPointer(ptr unsafe.Pointer) *QScriptContextInfo {
	var n = new(QScriptContextInfo)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScriptContextInfo) QScriptContextInfo_PTR() *QScriptContextInfo {
	return ptr
}

//QScriptContextInfo::FunctionType
type QScriptContextInfo__FunctionType int64

const (
	QScriptContextInfo__ScriptFunction     = QScriptContextInfo__FunctionType(0)
	QScriptContextInfo__QtFunction         = QScriptContextInfo__FunctionType(1)
	QScriptContextInfo__QtPropertyFunction = QScriptContextInfo__FunctionType(2)
	QScriptContextInfo__NativeFunction     = QScriptContextInfo__FunctionType(3)
)

func NewQScriptContextInfo3() *QScriptContextInfo {
	defer qt.Recovering("QScriptContextInfo::QScriptContextInfo")

	return NewQScriptContextInfoFromPointer(C.QScriptContextInfo_NewQScriptContextInfo3())
}

func NewQScriptContextInfo(context QScriptContext_ITF) *QScriptContextInfo {
	defer qt.Recovering("QScriptContextInfo::QScriptContextInfo")

	return NewQScriptContextInfoFromPointer(C.QScriptContextInfo_NewQScriptContextInfo(PointerFromQScriptContext(context)))
}

func NewQScriptContextInfo2(other QScriptContextInfo_ITF) *QScriptContextInfo {
	defer qt.Recovering("QScriptContextInfo::QScriptContextInfo")

	return NewQScriptContextInfoFromPointer(C.QScriptContextInfo_NewQScriptContextInfo2(PointerFromQScriptContextInfo(other)))
}

func (ptr *QScriptContextInfo) FileName() string {
	defer qt.Recovering("QScriptContextInfo::fileName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptContextInfo_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptContextInfo) FunctionEndLineNumber() int {
	defer qt.Recovering("QScriptContextInfo::functionEndLineNumber")

	if ptr.Pointer() != nil {
		return int(C.QScriptContextInfo_FunctionEndLineNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptContextInfo) FunctionMetaIndex() int {
	defer qt.Recovering("QScriptContextInfo::functionMetaIndex")

	if ptr.Pointer() != nil {
		return int(C.QScriptContextInfo_FunctionMetaIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptContextInfo) FunctionName() string {
	defer qt.Recovering("QScriptContextInfo::functionName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptContextInfo_FunctionName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptContextInfo) FunctionParameterNames() []string {
	defer qt.Recovering("QScriptContextInfo::functionParameterNames")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QScriptContextInfo_FunctionParameterNames(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QScriptContextInfo) FunctionStartLineNumber() int {
	defer qt.Recovering("QScriptContextInfo::functionStartLineNumber")

	if ptr.Pointer() != nil {
		return int(C.QScriptContextInfo_FunctionStartLineNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptContextInfo) FunctionType() QScriptContextInfo__FunctionType {
	defer qt.Recovering("QScriptContextInfo::functionType")

	if ptr.Pointer() != nil {
		return QScriptContextInfo__FunctionType(C.QScriptContextInfo_FunctionType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptContextInfo) IsNull() bool {
	defer qt.Recovering("QScriptContextInfo::isNull")

	if ptr.Pointer() != nil {
		return C.QScriptContextInfo_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptContextInfo) LineNumber() int {
	defer qt.Recovering("QScriptContextInfo::lineNumber")

	if ptr.Pointer() != nil {
		return int(C.QScriptContextInfo_LineNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptContextInfo) ScriptId() int64 {
	defer qt.Recovering("QScriptContextInfo::scriptId")

	if ptr.Pointer() != nil {
		return int64(C.QScriptContextInfo_ScriptId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptContextInfo) DestroyQScriptContextInfo() {
	defer qt.Recovering("QScriptContextInfo::~QScriptContextInfo")

	if ptr.Pointer() != nil {
		C.QScriptContextInfo_DestroyQScriptContextInfo(ptr.Pointer())
	}
}
