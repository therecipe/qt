package svg

//#include "svg.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QSvgRenderer struct {
	core.QObject
}

type QSvgRenderer_ITF interface {
	core.QObject_ITF
	QSvgRenderer_PTR() *QSvgRenderer
}

func PointerFromQSvgRenderer(ptr QSvgRenderer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSvgRenderer_PTR().Pointer()
	}
	return nil
}

func NewQSvgRendererFromPointer(ptr unsafe.Pointer) *QSvgRenderer {
	var n = new(QSvgRenderer)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSvgRenderer_") {
		n.SetObjectName("QSvgRenderer_" + qt.Identifier())
	}
	return n
}

func (ptr *QSvgRenderer) QSvgRenderer_PTR() *QSvgRenderer {
	return ptr
}

func (ptr *QSvgRenderer) FramesPerSecond() int {
	defer qt.Recovering("QSvgRenderer::framesPerSecond")

	if ptr.Pointer() != nil {
		return int(C.QSvgRenderer_FramesPerSecond(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSvgRenderer) SetFramesPerSecond(num int) {
	defer qt.Recovering("QSvgRenderer::setFramesPerSecond")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_SetFramesPerSecond(ptr.Pointer(), C.int(num))
	}
}

func (ptr *QSvgRenderer) SetViewBox(viewbox core.QRect_ITF) {
	defer qt.Recovering("QSvgRenderer::setViewBox")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_SetViewBox(ptr.Pointer(), core.PointerFromQRect(viewbox))
	}
}

func (ptr *QSvgRenderer) SetViewBox2(viewbox core.QRectF_ITF) {
	defer qt.Recovering("QSvgRenderer::setViewBox")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_SetViewBox2(ptr.Pointer(), core.PointerFromQRectF(viewbox))
	}
}

func NewQSvgRenderer(parent core.QObject_ITF) *QSvgRenderer {
	defer qt.Recovering("QSvgRenderer::QSvgRenderer")

	return NewQSvgRendererFromPointer(C.QSvgRenderer_NewQSvgRenderer(core.PointerFromQObject(parent)))
}

func NewQSvgRenderer4(contents core.QXmlStreamReader_ITF, parent core.QObject_ITF) *QSvgRenderer {
	defer qt.Recovering("QSvgRenderer::QSvgRenderer")

	return NewQSvgRendererFromPointer(C.QSvgRenderer_NewQSvgRenderer4(core.PointerFromQXmlStreamReader(contents), core.PointerFromQObject(parent)))
}

func NewQSvgRenderer3(contents core.QByteArray_ITF, parent core.QObject_ITF) *QSvgRenderer {
	defer qt.Recovering("QSvgRenderer::QSvgRenderer")

	return NewQSvgRendererFromPointer(C.QSvgRenderer_NewQSvgRenderer3(core.PointerFromQByteArray(contents), core.PointerFromQObject(parent)))
}

func NewQSvgRenderer2(filename string, parent core.QObject_ITF) *QSvgRenderer {
	defer qt.Recovering("QSvgRenderer::QSvgRenderer")

	return NewQSvgRendererFromPointer(C.QSvgRenderer_NewQSvgRenderer2(C.CString(filename), core.PointerFromQObject(parent)))
}

func (ptr *QSvgRenderer) Animated() bool {
	defer qt.Recovering("QSvgRenderer::animated")

	if ptr.Pointer() != nil {
		return C.QSvgRenderer_Animated(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSvgRenderer) DefaultSize() *core.QSize {
	defer qt.Recovering("QSvgRenderer::defaultSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QSvgRenderer_DefaultSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSvgRenderer) ElementExists(id string) bool {
	defer qt.Recovering("QSvgRenderer::elementExists")

	if ptr.Pointer() != nil {
		return C.QSvgRenderer_ElementExists(ptr.Pointer(), C.CString(id)) != 0
	}
	return false
}

func (ptr *QSvgRenderer) IsValid() bool {
	defer qt.Recovering("QSvgRenderer::isValid")

	if ptr.Pointer() != nil {
		return C.QSvgRenderer_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSvgRenderer) Load3(contents core.QXmlStreamReader_ITF) bool {
	defer qt.Recovering("QSvgRenderer::load")

	if ptr.Pointer() != nil {
		return C.QSvgRenderer_Load3(ptr.Pointer(), core.PointerFromQXmlStreamReader(contents)) != 0
	}
	return false
}

func (ptr *QSvgRenderer) Load2(contents core.QByteArray_ITF) bool {
	defer qt.Recovering("QSvgRenderer::load")

	if ptr.Pointer() != nil {
		return C.QSvgRenderer_Load2(ptr.Pointer(), core.PointerFromQByteArray(contents)) != 0
	}
	return false
}

func (ptr *QSvgRenderer) Load(filename string) bool {
	defer qt.Recovering("QSvgRenderer::load")

	if ptr.Pointer() != nil {
		return C.QSvgRenderer_Load(ptr.Pointer(), C.CString(filename)) != 0
	}
	return false
}

func (ptr *QSvgRenderer) Render(painter gui.QPainter_ITF) {
	defer qt.Recovering("QSvgRenderer::render")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_Render(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QSvgRenderer) Render2(painter gui.QPainter_ITF, bounds core.QRectF_ITF) {
	defer qt.Recovering("QSvgRenderer::render")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_Render2(ptr.Pointer(), gui.PointerFromQPainter(painter), core.PointerFromQRectF(bounds))
	}
}

func (ptr *QSvgRenderer) Render3(painter gui.QPainter_ITF, elementId string, bounds core.QRectF_ITF) {
	defer qt.Recovering("QSvgRenderer::render")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_Render3(ptr.Pointer(), gui.PointerFromQPainter(painter), C.CString(elementId), core.PointerFromQRectF(bounds))
	}
}

func (ptr *QSvgRenderer) ConnectRepaintNeeded(f func()) {
	defer qt.Recovering("connect QSvgRenderer::repaintNeeded")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_ConnectRepaintNeeded(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "repaintNeeded", f)
	}
}

func (ptr *QSvgRenderer) DisconnectRepaintNeeded() {
	defer qt.Recovering("disconnect QSvgRenderer::repaintNeeded")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_DisconnectRepaintNeeded(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "repaintNeeded")
	}
}

//export callbackQSvgRendererRepaintNeeded
func callbackQSvgRendererRepaintNeeded(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSvgRenderer::repaintNeeded")

	if signal := qt.GetSignal(C.GoString(ptrName), "repaintNeeded"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSvgRenderer) RepaintNeeded() {
	defer qt.Recovering("QSvgRenderer::repaintNeeded")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_RepaintNeeded(ptr.Pointer())
	}
}

func (ptr *QSvgRenderer) ViewBox() *core.QRect {
	defer qt.Recovering("QSvgRenderer::viewBox")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QSvgRenderer_ViewBox(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSvgRenderer) DestroyQSvgRenderer() {
	defer qt.Recovering("QSvgRenderer::~QSvgRenderer")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_DestroyQSvgRenderer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSvgRenderer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSvgRenderer::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSvgRenderer) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSvgRenderer::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSvgRendererTimerEvent
func callbackQSvgRendererTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgRenderer::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSvgRendererFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSvgRenderer) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSvgRenderer::timerEvent")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSvgRenderer) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSvgRenderer::timerEvent")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSvgRenderer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSvgRenderer::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSvgRenderer) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSvgRenderer::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSvgRendererChildEvent
func callbackQSvgRendererChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgRenderer::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSvgRendererFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSvgRenderer) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSvgRenderer::childEvent")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSvgRenderer) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSvgRenderer::childEvent")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSvgRenderer) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSvgRenderer::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSvgRenderer) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSvgRenderer::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSvgRendererCustomEvent
func callbackQSvgRendererCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgRenderer::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSvgRendererFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSvgRenderer) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSvgRenderer::customEvent")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSvgRenderer) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSvgRenderer::customEvent")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
