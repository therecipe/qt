package qml

//#include "qml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QQmlContext struct {
	core.QObject
}

type QQmlContext_ITF interface {
	core.QObject_ITF
	QQmlContext_PTR() *QQmlContext
}

func PointerFromQQmlContext(ptr QQmlContext_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlContext_PTR().Pointer()
	}
	return nil
}

func NewQQmlContextFromPointer(ptr unsafe.Pointer) *QQmlContext {
	var n = new(QQmlContext)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QQmlContext_") {
		n.SetObjectName("QQmlContext_" + qt.Identifier())
	}
	return n
}

func (ptr *QQmlContext) QQmlContext_PTR() *QQmlContext {
	return ptr
}

func NewQQmlContext2(parentContext QQmlContext_ITF, parent core.QObject_ITF) *QQmlContext {
	defer qt.Recovering("QQmlContext::QQmlContext")

	return NewQQmlContextFromPointer(C.QQmlContext_NewQQmlContext2(PointerFromQQmlContext(parentContext), core.PointerFromQObject(parent)))
}

func NewQQmlContext(engine QQmlEngine_ITF, parent core.QObject_ITF) *QQmlContext {
	defer qt.Recovering("QQmlContext::QQmlContext")

	return NewQQmlContextFromPointer(C.QQmlContext_NewQQmlContext(PointerFromQQmlEngine(engine), core.PointerFromQObject(parent)))
}

func (ptr *QQmlContext) ContextObject() *core.QObject {
	defer qt.Recovering("QQmlContext::contextObject")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QQmlContext_ContextObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlContext) ContextProperty(name string) *core.QVariant {
	defer qt.Recovering("QQmlContext::contextProperty")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QQmlContext_ContextProperty(ptr.Pointer(), C.CString(name)))
	}
	return nil
}

func (ptr *QQmlContext) Engine() *QQmlEngine {
	defer qt.Recovering("QQmlContext::engine")

	if ptr.Pointer() != nil {
		return NewQQmlEngineFromPointer(C.QQmlContext_Engine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlContext) IsValid() bool {
	defer qt.Recovering("QQmlContext::isValid")

	if ptr.Pointer() != nil {
		return C.QQmlContext_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlContext) NameForObject(object core.QObject_ITF) string {
	defer qt.Recovering("QQmlContext::nameForObject")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlContext_NameForObject(ptr.Pointer(), core.PointerFromQObject(object)))
	}
	return ""
}

func (ptr *QQmlContext) ParentContext() *QQmlContext {
	defer qt.Recovering("QQmlContext::parentContext")

	if ptr.Pointer() != nil {
		return NewQQmlContextFromPointer(C.QQmlContext_ParentContext(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlContext) SetBaseUrl(baseUrl core.QUrl_ITF) {
	defer qt.Recovering("QQmlContext::setBaseUrl")

	if ptr.Pointer() != nil {
		C.QQmlContext_SetBaseUrl(ptr.Pointer(), core.PointerFromQUrl(baseUrl))
	}
}

func (ptr *QQmlContext) SetContextObject(object core.QObject_ITF) {
	defer qt.Recovering("QQmlContext::setContextObject")

	if ptr.Pointer() != nil {
		C.QQmlContext_SetContextObject(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QQmlContext) SetContextProperty(name string, value core.QObject_ITF) {
	defer qt.Recovering("QQmlContext::setContextProperty")

	if ptr.Pointer() != nil {
		C.QQmlContext_SetContextProperty(ptr.Pointer(), C.CString(name), core.PointerFromQObject(value))
	}
}

func (ptr *QQmlContext) SetContextProperty2(name string, value core.QVariant_ITF) {
	defer qt.Recovering("QQmlContext::setContextProperty")

	if ptr.Pointer() != nil {
		C.QQmlContext_SetContextProperty2(ptr.Pointer(), C.CString(name), core.PointerFromQVariant(value))
	}
}

func (ptr *QQmlContext) DestroyQQmlContext() {
	defer qt.Recovering("QQmlContext::~QQmlContext")

	if ptr.Pointer() != nil {
		C.QQmlContext_DestroyQQmlContext(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
