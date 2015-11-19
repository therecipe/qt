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

type QQmlComponent_ITF interface {
	core.QObject_ITF
	QQmlComponent_PTR() *QQmlComponent
}

func PointerFromQQmlComponent(ptr QQmlComponent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlComponent_PTR().Pointer()
	}
	return nil
}

func NewQQmlComponentFromPointer(ptr unsafe.Pointer) *QQmlComponent {
	var n = new(QQmlComponent)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QQmlComponent_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQmlComponent) QQmlComponent_PTR() *QQmlComponent {
	return ptr
}

//QQmlComponent::CompilationMode
type QQmlComponent__CompilationMode int64

const (
	QQmlComponent__PreferSynchronous = QQmlComponent__CompilationMode(0)
	QQmlComponent__Asynchronous      = QQmlComponent__CompilationMode(1)
)

//QQmlComponent::Status
type QQmlComponent__Status int64

const (
	QQmlComponent__Null    = QQmlComponent__Status(0)
	QQmlComponent__Ready   = QQmlComponent__Status(1)
	QQmlComponent__Loading = QQmlComponent__Status(2)
	QQmlComponent__Error   = QQmlComponent__Status(3)
)

func (ptr *QQmlComponent) Progress() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQmlComponent_Progress(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlComponent) Status() QQmlComponent__Status {
	if ptr.Pointer() != nil {
		return QQmlComponent__Status(C.QQmlComponent_Status(ptr.Pointer()))
	}
	return 0
}

func NewQQmlComponent(engine QQmlEngine_ITF, parent core.QObject_ITF) *QQmlComponent {
	return NewQQmlComponentFromPointer(C.QQmlComponent_NewQQmlComponent(PointerFromQQmlEngine(engine), core.PointerFromQObject(parent)))
}

func NewQQmlComponent4(engine QQmlEngine_ITF, fileName string, mode QQmlComponent__CompilationMode, parent core.QObject_ITF) *QQmlComponent {
	return NewQQmlComponentFromPointer(C.QQmlComponent_NewQQmlComponent4(PointerFromQQmlEngine(engine), C.CString(fileName), C.int(mode), core.PointerFromQObject(parent)))
}

func NewQQmlComponent3(engine QQmlEngine_ITF, fileName string, parent core.QObject_ITF) *QQmlComponent {
	return NewQQmlComponentFromPointer(C.QQmlComponent_NewQQmlComponent3(PointerFromQQmlEngine(engine), C.CString(fileName), core.PointerFromQObject(parent)))
}

func NewQQmlComponent6(engine QQmlEngine_ITF, url core.QUrl_ITF, mode QQmlComponent__CompilationMode, parent core.QObject_ITF) *QQmlComponent {
	return NewQQmlComponentFromPointer(C.QQmlComponent_NewQQmlComponent6(PointerFromQQmlEngine(engine), core.PointerFromQUrl(url), C.int(mode), core.PointerFromQObject(parent)))
}

func NewQQmlComponent5(engine QQmlEngine_ITF, url core.QUrl_ITF, parent core.QObject_ITF) *QQmlComponent {
	return NewQQmlComponentFromPointer(C.QQmlComponent_NewQQmlComponent5(PointerFromQQmlEngine(engine), core.PointerFromQUrl(url), core.PointerFromQObject(parent)))
}

func (ptr *QQmlComponent) BeginCreate(publicContext QQmlContext_ITF) *core.QObject {
	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QQmlComponent_BeginCreate(ptr.Pointer(), PointerFromQQmlContext(publicContext)))
	}
	return nil
}

func (ptr *QQmlComponent) CompleteCreate() {
	if ptr.Pointer() != nil {
		C.QQmlComponent_CompleteCreate(ptr.Pointer())
	}
}

func (ptr *QQmlComponent) Create(context QQmlContext_ITF) *core.QObject {
	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QQmlComponent_Create(ptr.Pointer(), PointerFromQQmlContext(context)))
	}
	return nil
}

func (ptr *QQmlComponent) Create2(incubator QQmlIncubator_ITF, context QQmlContext_ITF, forContext QQmlContext_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_Create2(ptr.Pointer(), PointerFromQQmlIncubator(incubator), PointerFromQQmlContext(context), PointerFromQQmlContext(forContext))
	}
}

func (ptr *QQmlComponent) CreationContext() *QQmlContext {
	if ptr.Pointer() != nil {
		return NewQQmlContextFromPointer(C.QQmlComponent_CreationContext(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlComponent) IsError() bool {
	if ptr.Pointer() != nil {
		return C.QQmlComponent_IsError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlComponent) IsLoading() bool {
	if ptr.Pointer() != nil {
		return C.QQmlComponent_IsLoading(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlComponent) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QQmlComponent_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlComponent) IsReady() bool {
	if ptr.Pointer() != nil {
		return C.QQmlComponent_IsReady(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlComponent) LoadUrl(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_LoadUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QQmlComponent) LoadUrl2(url core.QUrl_ITF, mode QQmlComponent__CompilationMode) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_LoadUrl2(ptr.Pointer(), core.PointerFromQUrl(url), C.int(mode))
	}
}

func (ptr *QQmlComponent) SetData(data core.QByteArray_ITF, url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_SetData(ptr.Pointer(), core.PointerFromQByteArray(data), core.PointerFromQUrl(url))
	}
}

func (ptr *QQmlComponent) ConnectStatusChanged(f func(status QQmlComponent__Status)) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_ConnectStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "statusChanged", f)
	}
}

func (ptr *QQmlComponent) DisconnectStatusChanged() {
	if ptr.Pointer() != nil {
		C.QQmlComponent_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "statusChanged")
	}
}

//export callbackQQmlComponentStatusChanged
func callbackQQmlComponentStatusChanged(ptrName *C.char, status C.int) {
	qt.GetSignal(C.GoString(ptrName), "statusChanged").(func(QQmlComponent__Status))(QQmlComponent__Status(status))
}

func (ptr *QQmlComponent) DestroyQQmlComponent() {
	if ptr.Pointer() != nil {
		C.QQmlComponent_DestroyQQmlComponent(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
