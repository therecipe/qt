package qml

//#include "qjsengine.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QJSEngine struct {
	core.QObject
}

type QJSEngineITF interface {
	core.QObjectITF
	QJSEnginePTR() *QJSEngine
}

func PointerFromQJSEngine(ptr QJSEngineITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJSEnginePTR().Pointer()
	}
	return nil
}

func QJSEngineFromPointer(ptr unsafe.Pointer) *QJSEngine {
	var n = new(QJSEngine)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QJSEngine_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QJSEngine) QJSEnginePTR() *QJSEngine {
	return ptr
}

func NewQJSEngine() *QJSEngine {
	return QJSEngineFromPointer(unsafe.Pointer(C.QJSEngine_NewQJSEngine()))
}

func NewQJSEngine2(parent core.QObjectITF) *QJSEngine {
	return QJSEngineFromPointer(unsafe.Pointer(C.QJSEngine_NewQJSEngine2(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QJSEngine) CollectGarbage() {
	if ptr.Pointer() != nil {
		C.QJSEngine_CollectGarbage(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QJSEngine) InstallTranslatorFunctions(object QJSValueITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine_InstallTranslatorFunctions(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQJSValue(object)))
	}
}

func (ptr *QJSEngine) DestroyQJSEngine() {
	if ptr.Pointer() != nil {
		C.QJSEngine_DestroyQJSEngine(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
