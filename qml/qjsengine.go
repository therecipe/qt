package qml

//#include "qml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QJSEngine struct {
	core.QObject
}

type QJSEngine_ITF interface {
	core.QObject_ITF
	QJSEngine_PTR() *QJSEngine
}

func PointerFromQJSEngine(ptr QJSEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJSEngine_PTR().Pointer()
	}
	return nil
}

func NewQJSEngineFromPointer(ptr unsafe.Pointer) *QJSEngine {
	var n = new(QJSEngine)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QJSEngine_") {
		n.SetObjectName("QJSEngine_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QJSEngine) QJSEngine_PTR() *QJSEngine {
	return ptr
}

func NewQJSEngine() *QJSEngine {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QJSEngine::QJSEngine")
		}
	}()

	return NewQJSEngineFromPointer(C.QJSEngine_NewQJSEngine())
}

func NewQJSEngine2(parent core.QObject_ITF) *QJSEngine {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QJSEngine::QJSEngine")
		}
	}()

	return NewQJSEngineFromPointer(C.QJSEngine_NewQJSEngine2(core.PointerFromQObject(parent)))
}

func (ptr *QJSEngine) CollectGarbage() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QJSEngine::collectGarbage")
		}
	}()

	if ptr.Pointer() != nil {
		C.QJSEngine_CollectGarbage(ptr.Pointer())
	}
}

func (ptr *QJSEngine) Evaluate(program string, fileName string, lineNumber int) *QJSValue {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QJSEngine::evaluate")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQJSValueFromPointer(C.QJSEngine_Evaluate(ptr.Pointer(), C.CString(program), C.CString(fileName), C.int(lineNumber)))
	}
	return nil
}

func (ptr *QJSEngine) GlobalObject() *QJSValue {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QJSEngine::globalObject")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQJSValueFromPointer(C.QJSEngine_GlobalObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QJSEngine) InstallTranslatorFunctions(object QJSValue_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QJSEngine::installTranslatorFunctions")
		}
	}()

	if ptr.Pointer() != nil {
		C.QJSEngine_InstallTranslatorFunctions(ptr.Pointer(), PointerFromQJSValue(object))
	}
}

func (ptr *QJSEngine) NewObject() *QJSValue {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QJSEngine::newObject")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQJSValueFromPointer(C.QJSEngine_NewObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QJSEngine) NewQObject(object core.QObject_ITF) *QJSValue {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QJSEngine::newQObject")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQJSValueFromPointer(C.QJSEngine_NewQObject(ptr.Pointer(), core.PointerFromQObject(object)))
	}
	return nil
}

func (ptr *QJSEngine) DestroyQJSEngine() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QJSEngine::~QJSEngine")
		}
	}()

	if ptr.Pointer() != nil {
		C.QJSEngine_DestroyQJSEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
