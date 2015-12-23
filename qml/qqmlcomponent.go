package qml

//#include "qml.h"
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
	for len(n.ObjectName()) < len("QQmlComponent_") {
		n.SetObjectName("QQmlComponent_" + qt.Identifier())
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
	defer qt.Recovering("QQmlComponent::progress")

	if ptr.Pointer() != nil {
		return float64(C.QQmlComponent_Progress(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlComponent) Status() QQmlComponent__Status {
	defer qt.Recovering("QQmlComponent::status")

	if ptr.Pointer() != nil {
		return QQmlComponent__Status(C.QQmlComponent_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlComponent) Url() *core.QUrl {
	defer qt.Recovering("QQmlComponent::url")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QQmlComponent_Url(ptr.Pointer()))
	}
	return nil
}

func NewQQmlComponent(engine QQmlEngine_ITF, parent core.QObject_ITF) *QQmlComponent {
	defer qt.Recovering("QQmlComponent::QQmlComponent")

	return NewQQmlComponentFromPointer(C.QQmlComponent_NewQQmlComponent(PointerFromQQmlEngine(engine), core.PointerFromQObject(parent)))
}

func NewQQmlComponent4(engine QQmlEngine_ITF, fileName string, mode QQmlComponent__CompilationMode, parent core.QObject_ITF) *QQmlComponent {
	defer qt.Recovering("QQmlComponent::QQmlComponent")

	return NewQQmlComponentFromPointer(C.QQmlComponent_NewQQmlComponent4(PointerFromQQmlEngine(engine), C.CString(fileName), C.int(mode), core.PointerFromQObject(parent)))
}

func NewQQmlComponent3(engine QQmlEngine_ITF, fileName string, parent core.QObject_ITF) *QQmlComponent {
	defer qt.Recovering("QQmlComponent::QQmlComponent")

	return NewQQmlComponentFromPointer(C.QQmlComponent_NewQQmlComponent3(PointerFromQQmlEngine(engine), C.CString(fileName), core.PointerFromQObject(parent)))
}

func NewQQmlComponent6(engine QQmlEngine_ITF, url core.QUrl_ITF, mode QQmlComponent__CompilationMode, parent core.QObject_ITF) *QQmlComponent {
	defer qt.Recovering("QQmlComponent::QQmlComponent")

	return NewQQmlComponentFromPointer(C.QQmlComponent_NewQQmlComponent6(PointerFromQQmlEngine(engine), core.PointerFromQUrl(url), C.int(mode), core.PointerFromQObject(parent)))
}

func NewQQmlComponent5(engine QQmlEngine_ITF, url core.QUrl_ITF, parent core.QObject_ITF) *QQmlComponent {
	defer qt.Recovering("QQmlComponent::QQmlComponent")

	return NewQQmlComponentFromPointer(C.QQmlComponent_NewQQmlComponent5(PointerFromQQmlEngine(engine), core.PointerFromQUrl(url), core.PointerFromQObject(parent)))
}

func (ptr *QQmlComponent) BeginCreate(publicContext QQmlContext_ITF) *core.QObject {
	defer qt.Recovering("QQmlComponent::beginCreate")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QQmlComponent_BeginCreate(ptr.Pointer(), PointerFromQQmlContext(publicContext)))
	}
	return nil
}

func (ptr *QQmlComponent) ConnectCompleteCreate(f func()) {
	defer qt.Recovering("connect QQmlComponent::completeCreate")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "completeCreate", f)
	}
}

func (ptr *QQmlComponent) DisconnectCompleteCreate() {
	defer qt.Recovering("disconnect QQmlComponent::completeCreate")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "completeCreate")
	}
}

//export callbackQQmlComponentCompleteCreate
func callbackQQmlComponentCompleteCreate(ptrName *C.char) bool {
	defer qt.Recovering("callback QQmlComponent::completeCreate")

	if signal := qt.GetSignal(C.GoString(ptrName), "completeCreate"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QQmlComponent) Create(context QQmlContext_ITF) *core.QObject {
	defer qt.Recovering("QQmlComponent::create")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QQmlComponent_Create(ptr.Pointer(), PointerFromQQmlContext(context)))
	}
	return nil
}

func (ptr *QQmlComponent) Create2(incubator QQmlIncubator_ITF, context QQmlContext_ITF, forContext QQmlContext_ITF) {
	defer qt.Recovering("QQmlComponent::create")

	if ptr.Pointer() != nil {
		C.QQmlComponent_Create2(ptr.Pointer(), PointerFromQQmlIncubator(incubator), PointerFromQQmlContext(context), PointerFromQQmlContext(forContext))
	}
}

func (ptr *QQmlComponent) CreationContext() *QQmlContext {
	defer qt.Recovering("QQmlComponent::creationContext")

	if ptr.Pointer() != nil {
		return NewQQmlContextFromPointer(C.QQmlComponent_CreationContext(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlComponent) IsError() bool {
	defer qt.Recovering("QQmlComponent::isError")

	if ptr.Pointer() != nil {
		return C.QQmlComponent_IsError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlComponent) IsLoading() bool {
	defer qt.Recovering("QQmlComponent::isLoading")

	if ptr.Pointer() != nil {
		return C.QQmlComponent_IsLoading(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlComponent) IsNull() bool {
	defer qt.Recovering("QQmlComponent::isNull")

	if ptr.Pointer() != nil {
		return C.QQmlComponent_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlComponent) IsReady() bool {
	defer qt.Recovering("QQmlComponent::isReady")

	if ptr.Pointer() != nil {
		return C.QQmlComponent_IsReady(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlComponent) LoadUrl(url core.QUrl_ITF) {
	defer qt.Recovering("QQmlComponent::loadUrl")

	if ptr.Pointer() != nil {
		C.QQmlComponent_LoadUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QQmlComponent) LoadUrl2(url core.QUrl_ITF, mode QQmlComponent__CompilationMode) {
	defer qt.Recovering("QQmlComponent::loadUrl")

	if ptr.Pointer() != nil {
		C.QQmlComponent_LoadUrl2(ptr.Pointer(), core.PointerFromQUrl(url), C.int(mode))
	}
}

func (ptr *QQmlComponent) ConnectProgressChanged(f func(progress float64)) {
	defer qt.Recovering("connect QQmlComponent::progressChanged")

	if ptr.Pointer() != nil {
		C.QQmlComponent_ConnectProgressChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "progressChanged", f)
	}
}

func (ptr *QQmlComponent) DisconnectProgressChanged() {
	defer qt.Recovering("disconnect QQmlComponent::progressChanged")

	if ptr.Pointer() != nil {
		C.QQmlComponent_DisconnectProgressChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "progressChanged")
	}
}

//export callbackQQmlComponentProgressChanged
func callbackQQmlComponentProgressChanged(ptrName *C.char, progress C.double) {
	defer qt.Recovering("callback QQmlComponent::progressChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "progressChanged"); signal != nil {
		signal.(func(float64))(float64(progress))
	}

}

func (ptr *QQmlComponent) SetData(data core.QByteArray_ITF, url core.QUrl_ITF) {
	defer qt.Recovering("QQmlComponent::setData")

	if ptr.Pointer() != nil {
		C.QQmlComponent_SetData(ptr.Pointer(), core.PointerFromQByteArray(data), core.PointerFromQUrl(url))
	}
}

func (ptr *QQmlComponent) ConnectStatusChanged(f func(status QQmlComponent__Status)) {
	defer qt.Recovering("connect QQmlComponent::statusChanged")

	if ptr.Pointer() != nil {
		C.QQmlComponent_ConnectStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "statusChanged", f)
	}
}

func (ptr *QQmlComponent) DisconnectStatusChanged() {
	defer qt.Recovering("disconnect QQmlComponent::statusChanged")

	if ptr.Pointer() != nil {
		C.QQmlComponent_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "statusChanged")
	}
}

//export callbackQQmlComponentStatusChanged
func callbackQQmlComponentStatusChanged(ptrName *C.char, status C.int) {
	defer qt.Recovering("callback QQmlComponent::statusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "statusChanged"); signal != nil {
		signal.(func(QQmlComponent__Status))(QQmlComponent__Status(status))
	}

}

func (ptr *QQmlComponent) DestroyQQmlComponent() {
	defer qt.Recovering("QQmlComponent::~QQmlComponent")

	if ptr.Pointer() != nil {
		C.QQmlComponent_DestroyQQmlComponent(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlComponent) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQmlComponent::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQmlComponent) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQmlComponent::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQmlComponentTimerEvent
func callbackQQmlComponentTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQmlComponent::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQmlComponent) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQmlComponent::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQmlComponent) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQmlComponent::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQmlComponentChildEvent
func callbackQQmlComponentChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQmlComponent::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQmlComponent) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQmlComponent::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQmlComponent) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQmlComponent::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQmlComponentCustomEvent
func callbackQQmlComponentCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQmlComponent::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
