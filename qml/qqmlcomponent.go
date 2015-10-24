package qml

//#include "qqmlcomponent.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QQmlComponent struct {
	core.QObject
}

type QQmlComponentITF interface {
	core.QObjectITF
	QQmlComponentPTR() *QQmlComponent
}

func PointerFromQQmlComponent(ptr QQmlComponentITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlComponentPTR().Pointer()
	}
	return nil
}

func QQmlComponentFromPointer(ptr unsafe.Pointer) *QQmlComponent {
	var n = new(QQmlComponent)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QQmlComponent_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQmlComponent) QQmlComponentPTR() *QQmlComponent {
	return ptr
}

//QQmlComponent::CompilationMode
type QQmlComponent__CompilationMode int

var (
	QQmlComponent__PreferSynchronous = QQmlComponent__CompilationMode(0)
	QQmlComponent__Asynchronous      = QQmlComponent__CompilationMode(1)
)

//QQmlComponent::Status
type QQmlComponent__Status int

var (
	QQmlComponent__Null    = QQmlComponent__Status(0)
	QQmlComponent__Ready   = QQmlComponent__Status(1)
	QQmlComponent__Loading = QQmlComponent__Status(2)
	QQmlComponent__Error   = QQmlComponent__Status(3)
)

func (ptr *QQmlComponent) Status() QQmlComponent__Status {
	if ptr.Pointer() != nil {
		return QQmlComponent__Status(C.QQmlComponent_Status(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlComponent) Url() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlComponent_Url(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func NewQQmlComponent(engine QQmlEngineITF, parent core.QObjectITF) *QQmlComponent {
	return QQmlComponentFromPointer(unsafe.Pointer(C.QQmlComponent_NewQQmlComponent(C.QtObjectPtr(PointerFromQQmlEngine(engine)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQQmlComponent4(engine QQmlEngineITF, fileName string, mode QQmlComponent__CompilationMode, parent core.QObjectITF) *QQmlComponent {
	return QQmlComponentFromPointer(unsafe.Pointer(C.QQmlComponent_NewQQmlComponent4(C.QtObjectPtr(PointerFromQQmlEngine(engine)), C.CString(fileName), C.int(mode), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQQmlComponent3(engine QQmlEngineITF, fileName string, parent core.QObjectITF) *QQmlComponent {
	return QQmlComponentFromPointer(unsafe.Pointer(C.QQmlComponent_NewQQmlComponent3(C.QtObjectPtr(PointerFromQQmlEngine(engine)), C.CString(fileName), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQQmlComponent6(engine QQmlEngineITF, url string, mode QQmlComponent__CompilationMode, parent core.QObjectITF) *QQmlComponent {
	return QQmlComponentFromPointer(unsafe.Pointer(C.QQmlComponent_NewQQmlComponent6(C.QtObjectPtr(PointerFromQQmlEngine(engine)), C.CString(url), C.int(mode), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQQmlComponent5(engine QQmlEngineITF, url string, parent core.QObjectITF) *QQmlComponent {
	return QQmlComponentFromPointer(unsafe.Pointer(C.QQmlComponent_NewQQmlComponent5(C.QtObjectPtr(PointerFromQQmlEngine(engine)), C.CString(url), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QQmlComponent) BeginCreate(publicContext QQmlContextITF) *core.QObject {
	if ptr.Pointer() != nil {
		return core.QObjectFromPointer(unsafe.Pointer(C.QQmlComponent_BeginCreate(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQQmlContext(publicContext)))))
	}
	return nil
}

func (ptr *QQmlComponent) CompleteCreate() {
	if ptr.Pointer() != nil {
		C.QQmlComponent_CompleteCreate(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QQmlComponent) Create(context QQmlContextITF) *core.QObject {
	if ptr.Pointer() != nil {
		return core.QObjectFromPointer(unsafe.Pointer(C.QQmlComponent_Create(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQQmlContext(context)))))
	}
	return nil
}

func (ptr *QQmlComponent) Create2(incubator QQmlIncubatorITF, context QQmlContextITF, forContext QQmlContextITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_Create2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQQmlIncubator(incubator)), C.QtObjectPtr(PointerFromQQmlContext(context)), C.QtObjectPtr(PointerFromQQmlContext(forContext)))
	}
}

func (ptr *QQmlComponent) CreationContext() *QQmlContext {
	if ptr.Pointer() != nil {
		return QQmlContextFromPointer(unsafe.Pointer(C.QQmlComponent_CreationContext(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QQmlComponent) IsError() bool {
	if ptr.Pointer() != nil {
		return C.QQmlComponent_IsError(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlComponent) IsLoading() bool {
	if ptr.Pointer() != nil {
		return C.QQmlComponent_IsLoading(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlComponent) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QQmlComponent_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlComponent) IsReady() bool {
	if ptr.Pointer() != nil {
		return C.QQmlComponent_IsReady(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlComponent) LoadUrl(url string) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_LoadUrl(C.QtObjectPtr(ptr.Pointer()), C.CString(url))
	}
}

func (ptr *QQmlComponent) LoadUrl2(url string, mode QQmlComponent__CompilationMode) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_LoadUrl2(C.QtObjectPtr(ptr.Pointer()), C.CString(url), C.int(mode))
	}
}

func (ptr *QQmlComponent) SetData(data core.QByteArrayITF, url string) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_SetData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(data)), C.CString(url))
	}
}

func (ptr *QQmlComponent) ConnectStatusChanged(f func(status QQmlComponent__Status)) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_ConnectStatusChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "statusChanged", f)
	}
}

func (ptr *QQmlComponent) DisconnectStatusChanged() {
	if ptr.Pointer() != nil {
		C.QQmlComponent_DisconnectStatusChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "statusChanged")
	}
}

//export callbackQQmlComponentStatusChanged
func callbackQQmlComponentStatusChanged(ptrName *C.char, status C.int) {
	qt.GetSignal(C.GoString(ptrName), "statusChanged").(func(QQmlComponent__Status))(QQmlComponent__Status(status))
}

func (ptr *QQmlComponent) DestroyQQmlComponent() {
	if ptr.Pointer() != nil {
		C.QQmlComponent_DestroyQQmlComponent(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
