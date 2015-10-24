package qml

//#include "qqmlcontext.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QQmlContext struct {
	core.QObject
}

type QQmlContextITF interface {
	core.QObjectITF
	QQmlContextPTR() *QQmlContext
}

func PointerFromQQmlContext(ptr QQmlContextITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlContextPTR().Pointer()
	}
	return nil
}

func QQmlContextFromPointer(ptr unsafe.Pointer) *QQmlContext {
	var n = new(QQmlContext)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QQmlContext_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQmlContext) QQmlContextPTR() *QQmlContext {
	return ptr
}

func NewQQmlContext2(parentContext QQmlContextITF, parent core.QObjectITF) *QQmlContext {
	return QQmlContextFromPointer(unsafe.Pointer(C.QQmlContext_NewQQmlContext2(C.QtObjectPtr(PointerFromQQmlContext(parentContext)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQQmlContext(engine QQmlEngineITF, parent core.QObjectITF) *QQmlContext {
	return QQmlContextFromPointer(unsafe.Pointer(C.QQmlContext_NewQQmlContext(C.QtObjectPtr(PointerFromQQmlEngine(engine)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QQmlContext) BaseUrl() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlContext_BaseUrl(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QQmlContext) ContextObject() *core.QObject {
	if ptr.Pointer() != nil {
		return core.QObjectFromPointer(unsafe.Pointer(C.QQmlContext_ContextObject(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QQmlContext) ContextProperty(name string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlContext_ContextProperty(C.QtObjectPtr(ptr.Pointer()), C.CString(name)))
	}
	return ""
}

func (ptr *QQmlContext) Engine() *QQmlEngine {
	if ptr.Pointer() != nil {
		return QQmlEngineFromPointer(unsafe.Pointer(C.QQmlContext_Engine(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QQmlContext) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QQmlContext_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlContext) NameForObject(object core.QObjectITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlContext_NameForObject(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQObject(object))))
	}
	return ""
}

func (ptr *QQmlContext) ParentContext() *QQmlContext {
	if ptr.Pointer() != nil {
		return QQmlContextFromPointer(unsafe.Pointer(C.QQmlContext_ParentContext(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QQmlContext) ResolvedUrl(src string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlContext_ResolvedUrl(C.QtObjectPtr(ptr.Pointer()), C.CString(src)))
	}
	return ""
}

func (ptr *QQmlContext) SetBaseUrl(baseUrl string) {
	if ptr.Pointer() != nil {
		C.QQmlContext_SetBaseUrl(C.QtObjectPtr(ptr.Pointer()), C.CString(baseUrl))
	}
}

func (ptr *QQmlContext) SetContextObject(object core.QObjectITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext_SetContextObject(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQObject(object)))
	}
}

func (ptr *QQmlContext) SetContextProperty(name string, value core.QObjectITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext_SetContextProperty(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.QtObjectPtr(core.PointerFromQObject(value)))
	}
}

func (ptr *QQmlContext) SetContextProperty2(name string, value string) {
	if ptr.Pointer() != nil {
		C.QQmlContext_SetContextProperty2(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.CString(value))
	}
}

func (ptr *QQmlContext) DestroyQQmlContext() {
	if ptr.Pointer() != nil {
		C.QQmlContext_DestroyQQmlContext(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
