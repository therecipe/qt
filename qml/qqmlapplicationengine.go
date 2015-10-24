package qml

//#include "qqmlapplicationengine.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QQmlApplicationEngine struct {
	QQmlEngine
}

type QQmlApplicationEngineITF interface {
	QQmlEngineITF
	QQmlApplicationEnginePTR() *QQmlApplicationEngine
}

func PointerFromQQmlApplicationEngine(ptr QQmlApplicationEngineITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlApplicationEnginePTR().Pointer()
	}
	return nil
}

func QQmlApplicationEngineFromPointer(ptr unsafe.Pointer) *QQmlApplicationEngine {
	var n = new(QQmlApplicationEngine)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QQmlApplicationEngine_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQmlApplicationEngine) QQmlApplicationEnginePTR() *QQmlApplicationEngine {
	return ptr
}

func NewQQmlApplicationEngine(parent core.QObjectITF) *QQmlApplicationEngine {
	return QQmlApplicationEngineFromPointer(unsafe.Pointer(C.QQmlApplicationEngine_NewQQmlApplicationEngine(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQQmlApplicationEngine3(filePath string, parent core.QObjectITF) *QQmlApplicationEngine {
	return QQmlApplicationEngineFromPointer(unsafe.Pointer(C.QQmlApplicationEngine_NewQQmlApplicationEngine3(C.CString(filePath), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQQmlApplicationEngine2(url string, parent core.QObjectITF) *QQmlApplicationEngine {
	return QQmlApplicationEngineFromPointer(unsafe.Pointer(C.QQmlApplicationEngine_NewQQmlApplicationEngine2(C.CString(url), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QQmlApplicationEngine) Load2(filePath string) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_Load2(C.QtObjectPtr(ptr.Pointer()), C.CString(filePath))
	}
}

func (ptr *QQmlApplicationEngine) Load(url string) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_Load(C.QtObjectPtr(ptr.Pointer()), C.CString(url))
	}
}

func (ptr *QQmlApplicationEngine) LoadData(data core.QByteArrayITF, url string) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_LoadData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(data)), C.CString(url))
	}
}

func (ptr *QQmlApplicationEngine) ConnectObjectCreated(f func(object core.QObjectITF, url string)) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_ConnectObjectCreated(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "objectCreated", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectObjectCreated() {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_DisconnectObjectCreated(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "objectCreated")
	}
}

//export callbackQQmlApplicationEngineObjectCreated
func callbackQQmlApplicationEngineObjectCreated(ptrName *C.char, object unsafe.Pointer, url *C.char) {
	qt.GetSignal(C.GoString(ptrName), "objectCreated").(func(*core.QObject, string))(core.QObjectFromPointer(object), C.GoString(url))
}

func (ptr *QQmlApplicationEngine) DestroyQQmlApplicationEngine() {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_DestroyQQmlApplicationEngine(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
