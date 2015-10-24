package script

//#include "qscriptclass.h"
import "C"
import (
	"unsafe"
)

type QScriptClass struct {
	ptr unsafe.Pointer
}

type QScriptClassITF interface {
	QScriptClassPTR() *QScriptClass
}

func (p *QScriptClass) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptClass) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQScriptClass(ptr QScriptClassITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptClassPTR().Pointer()
	}
	return nil
}

func QScriptClassFromPointer(ptr unsafe.Pointer) *QScriptClass {
	var n = new(QScriptClass)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScriptClass) QScriptClassPTR() *QScriptClass {
	return ptr
}

//QScriptClass::Extension
type QScriptClass__Extension int

var (
	QScriptClass__Callable    = QScriptClass__Extension(0)
	QScriptClass__HasInstance = QScriptClass__Extension(1)
)

//QScriptClass::QueryFlag
type QScriptClass__QueryFlag int

var (
	QScriptClass__HandlesReadAccess  = QScriptClass__QueryFlag(0x01)
	QScriptClass__HandlesWriteAccess = QScriptClass__QueryFlag(0x02)
)

func NewQScriptClass(engine QScriptEngineITF) *QScriptClass {
	return QScriptClassFromPointer(unsafe.Pointer(C.QScriptClass_NewQScriptClass(C.QtObjectPtr(PointerFromQScriptEngine(engine)))))
}

func (ptr *QScriptClass) Engine() *QScriptEngine {
	if ptr.Pointer() != nil {
		return QScriptEngineFromPointer(unsafe.Pointer(C.QScriptClass_Engine(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QScriptClass) Extension(extension QScriptClass__Extension, argument string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptClass_Extension(C.QtObjectPtr(ptr.Pointer()), C.int(extension), C.CString(argument)))
	}
	return ""
}

func (ptr *QScriptClass) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptClass_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QScriptClass) NewIterator(object QScriptValueITF) *QScriptClassPropertyIterator {
	if ptr.Pointer() != nil {
		return QScriptClassPropertyIteratorFromPointer(unsafe.Pointer(C.QScriptClass_NewIterator(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQScriptValue(object)))))
	}
	return nil
}

func (ptr *QScriptClass) SupportsExtension(extension QScriptClass__Extension) bool {
	if ptr.Pointer() != nil {
		return C.QScriptClass_SupportsExtension(C.QtObjectPtr(ptr.Pointer()), C.int(extension)) != 0
	}
	return false
}

func (ptr *QScriptClass) DestroyQScriptClass() {
	if ptr.Pointer() != nil {
		C.QScriptClass_DestroyQScriptClass(C.QtObjectPtr(ptr.Pointer()))
	}
}
