package script

//#include "script.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QScriptClass struct {
	ptr unsafe.Pointer
}

type QScriptClass_ITF interface {
	QScriptClass_PTR() *QScriptClass
}

func (p *QScriptClass) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptClass) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQScriptClass(ptr QScriptClass_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptClass_PTR().Pointer()
	}
	return nil
}

func NewQScriptClassFromPointer(ptr unsafe.Pointer) *QScriptClass {
	var n = new(QScriptClass)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QScriptClass_") {
		n.SetObjectNameAbs("QScriptClass_" + qt.Identifier())
	}
	return n
}

func (ptr *QScriptClass) QScriptClass_PTR() *QScriptClass {
	return ptr
}

//QScriptClass::Extension
type QScriptClass__Extension int64

const (
	QScriptClass__Callable    = QScriptClass__Extension(0)
	QScriptClass__HasInstance = QScriptClass__Extension(1)
)

//QScriptClass::QueryFlag
type QScriptClass__QueryFlag int64

const (
	QScriptClass__HandlesReadAccess  = QScriptClass__QueryFlag(0x01)
	QScriptClass__HandlesWriteAccess = QScriptClass__QueryFlag(0x02)
)

func NewQScriptClass(engine QScriptEngine_ITF) *QScriptClass {
	defer qt.Recovering("QScriptClass::QScriptClass")

	return NewQScriptClassFromPointer(C.QScriptClass_NewQScriptClass(PointerFromQScriptEngine(engine)))
}

func (ptr *QScriptClass) Engine() *QScriptEngine {
	defer qt.Recovering("QScriptClass::engine")

	if ptr.Pointer() != nil {
		return NewQScriptEngineFromPointer(C.QScriptClass_Engine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptClass) Extension(extension QScriptClass__Extension, argument core.QVariant_ITF) *core.QVariant {
	defer qt.Recovering("QScriptClass::extension")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QScriptClass_Extension(ptr.Pointer(), C.int(extension), core.PointerFromQVariant(argument)))
	}
	return nil
}

func (ptr *QScriptClass) Name() string {
	defer qt.Recovering("QScriptClass::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptClass_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptClass) NewIterator(object QScriptValue_ITF) *QScriptClassPropertyIterator {
	defer qt.Recovering("QScriptClass::newIterator")

	if ptr.Pointer() != nil {
		return NewQScriptClassPropertyIteratorFromPointer(C.QScriptClass_NewIterator(ptr.Pointer(), PointerFromQScriptValue(object)))
	}
	return nil
}

func (ptr *QScriptClass) Prototype() *QScriptValue {
	defer qt.Recovering("QScriptClass::prototype")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptClass_Prototype(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptClass) SupportsExtension(extension QScriptClass__Extension) bool {
	defer qt.Recovering("QScriptClass::supportsExtension")

	if ptr.Pointer() != nil {
		return C.QScriptClass_SupportsExtension(ptr.Pointer(), C.int(extension)) != 0
	}
	return false
}

func (ptr *QScriptClass) DestroyQScriptClass() {
	defer qt.Recovering("QScriptClass::~QScriptClass")

	if ptr.Pointer() != nil {
		C.QScriptClass_DestroyQScriptClass(ptr.Pointer())
	}
}

func (ptr *QScriptClass) ObjectNameAbs() string {
	defer qt.Recovering("QScriptClass::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptClass_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptClass) SetObjectNameAbs(name string) {
	defer qt.Recovering("QScriptClass::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QScriptClass_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
