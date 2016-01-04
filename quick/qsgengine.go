package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QSGEngine struct {
	core.QObject
}

type QSGEngine_ITF interface {
	core.QObject_ITF
	QSGEngine_PTR() *QSGEngine
}

func PointerFromQSGEngine(ptr QSGEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGEngine_PTR().Pointer()
	}
	return nil
}

func NewQSGEngineFromPointer(ptr unsafe.Pointer) *QSGEngine {
	var n = new(QSGEngine)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSGEngine_") {
		n.SetObjectName("QSGEngine_" + qt.Identifier())
	}
	return n
}

func (ptr *QSGEngine) QSGEngine_PTR() *QSGEngine {
	return ptr
}

//QSGEngine::CreateTextureOption
type QSGEngine__CreateTextureOption int64

const (
	QSGEngine__TextureHasAlphaChannel = QSGEngine__CreateTextureOption(0x0001)
	QSGEngine__TextureOwnsGLTexture   = QSGEngine__CreateTextureOption(0x0004)
	QSGEngine__TextureCanUseAtlas     = QSGEngine__CreateTextureOption(0x0008)
)

func NewQSGEngine(parent core.QObject_ITF) *QSGEngine {
	defer qt.Recovering("QSGEngine::QSGEngine")

	return NewQSGEngineFromPointer(C.QSGEngine_NewQSGEngine(core.PointerFromQObject(parent)))
}

func (ptr *QSGEngine) CreateRenderer() *QSGAbstractRenderer {
	defer qt.Recovering("QSGEngine::createRenderer")

	if ptr.Pointer() != nil {
		return NewQSGAbstractRendererFromPointer(C.QSGEngine_CreateRenderer(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGEngine) CreateTextureFromImage(image gui.QImage_ITF, options QSGEngine__CreateTextureOption) *QSGTexture {
	defer qt.Recovering("QSGEngine::createTextureFromImage")

	if ptr.Pointer() != nil {
		return NewQSGTextureFromPointer(C.QSGEngine_CreateTextureFromImage(ptr.Pointer(), gui.PointerFromQImage(image), C.int(options)))
	}
	return nil
}

func (ptr *QSGEngine) Initialize(context gui.QOpenGLContext_ITF) {
	defer qt.Recovering("QSGEngine::initialize")

	if ptr.Pointer() != nil {
		C.QSGEngine_Initialize(ptr.Pointer(), gui.PointerFromQOpenGLContext(context))
	}
}

func (ptr *QSGEngine) Invalidate() {
	defer qt.Recovering("QSGEngine::invalidate")

	if ptr.Pointer() != nil {
		C.QSGEngine_Invalidate(ptr.Pointer())
	}
}

func (ptr *QSGEngine) DestroyQSGEngine() {
	defer qt.Recovering("QSGEngine::~QSGEngine")

	if ptr.Pointer() != nil {
		C.QSGEngine_DestroyQSGEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSGEngine) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSGEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSGEngine) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSGEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSGEngineTimerEvent
func callbackQSGEngineTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGEngine::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSGEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSGEngine) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSGEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QSGEngine_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSGEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSGEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QSGEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSGEngine) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSGEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSGEngine) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSGEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSGEngineChildEvent
func callbackQSGEngineChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGEngine::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSGEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSGEngine) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSGEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QSGEngine_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSGEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSGEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QSGEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSGEngine) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSGEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSGEngine) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSGEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSGEngineCustomEvent
func callbackQSGEngineCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGEngine::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSGEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSGEngine) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSGEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QSGEngine_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSGEngine) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSGEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QSGEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
