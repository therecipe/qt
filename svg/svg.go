package svg

//#include "svg.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"unsafe"
)

type QGraphicsSvgItem struct {
	widgets.QGraphicsObject
}

type QGraphicsSvgItem_ITF interface {
	widgets.QGraphicsObject_ITF
	QGraphicsSvgItem_PTR() *QGraphicsSvgItem
}

func PointerFromQGraphicsSvgItem(ptr QGraphicsSvgItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSvgItem_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsSvgItemFromPointer(ptr unsafe.Pointer) *QGraphicsSvgItem {
	var n = new(QGraphicsSvgItem)
	n.SetPointer(ptr)
	return n
}

func newQGraphicsSvgItemFromPointer(ptr unsafe.Pointer) *QGraphicsSvgItem {
	var n = NewQGraphicsSvgItemFromPointer(ptr)
	for len(n.ObjectName()) < len("QGraphicsSvgItem_") {
		n.SetObjectName("QGraphicsSvgItem_" + qt.Identifier())
	}
	return n
}

func (ptr *QGraphicsSvgItem) QGraphicsSvgItem_PTR() *QGraphicsSvgItem {
	return ptr
}

func (ptr *QGraphicsSvgItem) ElementId() string {
	defer qt.Recovering("QGraphicsSvgItem::elementId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsSvgItem_ElementId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGraphicsSvgItem) MaximumCacheSize() *core.QSize {
	defer qt.Recovering("QGraphicsSvgItem::maximumCacheSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QGraphicsSvgItem_MaximumCacheSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsSvgItem) ConnectPaint(f func(painter *gui.QPainter, option *widgets.QStyleOptionGraphicsItem, widget *widgets.QWidget)) {
	defer qt.Recovering("connect QGraphicsSvgItem::paint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paint", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectPaint() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::paint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paint")
	}
}

//export callbackQGraphicsSvgItemPaint
func callbackQGraphicsSvgItemPaint(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSvgItem::paint")

	if signal := qt.GetSignal(C.GoString(ptrName), "paint"); signal != nil {
		signal.(func(*gui.QPainter, *widgets.QStyleOptionGraphicsItem, *widgets.QWidget))(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionGraphicsItemFromPointer(option), widgets.NewQWidgetFromPointer(widget))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).PaintDefault(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionGraphicsItemFromPointer(option), widgets.NewQWidgetFromPointer(widget))
	}
}

func (ptr *QGraphicsSvgItem) Paint(painter gui.QPainter_ITF, option widgets.QStyleOptionGraphicsItem_ITF, widget widgets.QWidget_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::paint")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionGraphicsItem(option), widgets.PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsSvgItem) PaintDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionGraphicsItem_ITF, widget widgets.QWidget_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::paint")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_PaintDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionGraphicsItem(option), widgets.PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsSvgItem) Renderer() *QSvgRenderer {
	defer qt.Recovering("QGraphicsSvgItem::renderer")

	if ptr.Pointer() != nil {
		return NewQSvgRendererFromPointer(C.QGraphicsSvgItem_Renderer(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsSvgItem) SetElementId(id string) {
	defer qt.Recovering("QGraphicsSvgItem::setElementId")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_SetElementId(ptr.Pointer(), C.CString(id))
	}
}

func (ptr *QGraphicsSvgItem) SetMaximumCacheSize(size core.QSize_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::setMaximumCacheSize")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_SetMaximumCacheSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QGraphicsSvgItem) SetSharedRenderer(renderer QSvgRenderer_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::setSharedRenderer")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_SetSharedRenderer(ptr.Pointer(), PointerFromQSvgRenderer(renderer))
	}
}

func (ptr *QGraphicsSvgItem) Type() int {
	defer qt.Recovering("QGraphicsSvgItem::type")

	if ptr.Pointer() != nil {
		return int(C.QGraphicsSvgItem_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsSvgItem) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGraphicsSvgItem::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGraphicsSvgItemTimerEvent
func callbackQGraphicsSvgItemTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSvgItem::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::timerEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGraphicsSvgItem) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::timerEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGraphicsSvgItem) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGraphicsSvgItem::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGraphicsSvgItemChildEvent
func callbackQGraphicsSvgItemChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSvgItem::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::childEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGraphicsSvgItem) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::childEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGraphicsSvgItem) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsSvgItem::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGraphicsSvgItemCustomEvent
func callbackQGraphicsSvgItemCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSvgItem::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::customEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsSvgItem) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::customEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QSvgGenerator struct {
	gui.QPaintDevice
}

type QSvgGenerator_ITF interface {
	gui.QPaintDevice_ITF
	QSvgGenerator_PTR() *QSvgGenerator
}

func PointerFromQSvgGenerator(ptr QSvgGenerator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSvgGenerator_PTR().Pointer()
	}
	return nil
}

func NewQSvgGeneratorFromPointer(ptr unsafe.Pointer) *QSvgGenerator {
	var n = new(QSvgGenerator)
	n.SetPointer(ptr)
	return n
}

func newQSvgGeneratorFromPointer(ptr unsafe.Pointer) *QSvgGenerator {
	var n = NewQSvgGeneratorFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QSvgGenerator_") {
		n.SetObjectNameAbs("QSvgGenerator_" + qt.Identifier())
	}
	return n
}

func (ptr *QSvgGenerator) QSvgGenerator_PTR() *QSvgGenerator {
	return ptr
}

func (ptr *QSvgGenerator) Description() string {
	defer qt.Recovering("QSvgGenerator::description")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSvgGenerator_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSvgGenerator) FileName() string {
	defer qt.Recovering("QSvgGenerator::fileName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSvgGenerator_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSvgGenerator) OutputDevice() *core.QIODevice {
	defer qt.Recovering("QSvgGenerator::outputDevice")

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QSvgGenerator_OutputDevice(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSvgGenerator) Resolution() int {
	defer qt.Recovering("QSvgGenerator::resolution")

	if ptr.Pointer() != nil {
		return int(C.QSvgGenerator_Resolution(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSvgGenerator) SetDescription(description string) {
	defer qt.Recovering("QSvgGenerator::setDescription")

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetDescription(ptr.Pointer(), C.CString(description))
	}
}

func (ptr *QSvgGenerator) SetFileName(fileName string) {
	defer qt.Recovering("QSvgGenerator::setFileName")

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetFileName(ptr.Pointer(), C.CString(fileName))
	}
}

func (ptr *QSvgGenerator) SetOutputDevice(outputDevice core.QIODevice_ITF) {
	defer qt.Recovering("QSvgGenerator::setOutputDevice")

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetOutputDevice(ptr.Pointer(), core.PointerFromQIODevice(outputDevice))
	}
}

func (ptr *QSvgGenerator) SetResolution(dpi int) {
	defer qt.Recovering("QSvgGenerator::setResolution")

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetResolution(ptr.Pointer(), C.int(dpi))
	}
}

func (ptr *QSvgGenerator) SetSize(size core.QSize_ITF) {
	defer qt.Recovering("QSvgGenerator::setSize")

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QSvgGenerator) SetTitle(title string) {
	defer qt.Recovering("QSvgGenerator::setTitle")

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetTitle(ptr.Pointer(), C.CString(title))
	}
}

func (ptr *QSvgGenerator) SetViewBox(viewBox core.QRect_ITF) {
	defer qt.Recovering("QSvgGenerator::setViewBox")

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetViewBox(ptr.Pointer(), core.PointerFromQRect(viewBox))
	}
}

func (ptr *QSvgGenerator) SetViewBox2(viewBox core.QRectF_ITF) {
	defer qt.Recovering("QSvgGenerator::setViewBox")

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetViewBox2(ptr.Pointer(), core.PointerFromQRectF(viewBox))
	}
}

func (ptr *QSvgGenerator) Size() *core.QSize {
	defer qt.Recovering("QSvgGenerator::size")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QSvgGenerator_Size(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSvgGenerator) Title() string {
	defer qt.Recovering("QSvgGenerator::title")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSvgGenerator_Title(ptr.Pointer()))
	}
	return ""
}

func NewQSvgGenerator() *QSvgGenerator {
	defer qt.Recovering("QSvgGenerator::QSvgGenerator")

	return newQSvgGeneratorFromPointer(C.QSvgGenerator_NewQSvgGenerator())
}

func (ptr *QSvgGenerator) Metric(metric gui.QPaintDevice__PaintDeviceMetric) int {
	defer qt.Recovering("QSvgGenerator::metric")

	if ptr.Pointer() != nil {
		return int(C.QSvgGenerator_Metric(ptr.Pointer(), C.int(metric)))
	}
	return 0
}

func (ptr *QSvgGenerator) PaintEngine() *gui.QPaintEngine {
	defer qt.Recovering("QSvgGenerator::paintEngine")

	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QSvgGenerator_PaintEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSvgGenerator) ViewBox() *core.QRect {
	defer qt.Recovering("QSvgGenerator::viewBox")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QSvgGenerator_ViewBox(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSvgGenerator) DestroyQSvgGenerator() {
	defer qt.Recovering("QSvgGenerator::~QSvgGenerator")

	if ptr.Pointer() != nil {
		C.QSvgGenerator_DestroyQSvgGenerator(ptr.Pointer())
	}
}

func (ptr *QSvgGenerator) ObjectNameAbs() string {
	defer qt.Recovering("QSvgGenerator::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSvgGenerator_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSvgGenerator) SetObjectNameAbs(name string) {
	defer qt.Recovering("QSvgGenerator::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

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
	return n
}

func newQSvgRendererFromPointer(ptr unsafe.Pointer) *QSvgRenderer {
	var n = NewQSvgRendererFromPointer(ptr)
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

	return newQSvgRendererFromPointer(C.QSvgRenderer_NewQSvgRenderer(core.PointerFromQObject(parent)))
}

func NewQSvgRenderer4(contents core.QXmlStreamReader_ITF, parent core.QObject_ITF) *QSvgRenderer {
	defer qt.Recovering("QSvgRenderer::QSvgRenderer")

	return newQSvgRendererFromPointer(C.QSvgRenderer_NewQSvgRenderer4(core.PointerFromQXmlStreamReader(contents), core.PointerFromQObject(parent)))
}

func NewQSvgRenderer3(contents core.QByteArray_ITF, parent core.QObject_ITF) *QSvgRenderer {
	defer qt.Recovering("QSvgRenderer::QSvgRenderer")

	return newQSvgRendererFromPointer(C.QSvgRenderer_NewQSvgRenderer3(core.PointerFromQByteArray(contents), core.PointerFromQObject(parent)))
}

func NewQSvgRenderer2(filename string, parent core.QObject_ITF) *QSvgRenderer {
	defer qt.Recovering("QSvgRenderer::QSvgRenderer")

	return newQSvgRendererFromPointer(C.QSvgRenderer_NewQSvgRenderer2(C.CString(filename), core.PointerFromQObject(parent)))
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

type QSvgWidget struct {
	widgets.QWidget
}

type QSvgWidget_ITF interface {
	widgets.QWidget_ITF
	QSvgWidget_PTR() *QSvgWidget
}

func PointerFromQSvgWidget(ptr QSvgWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSvgWidget_PTR().Pointer()
	}
	return nil
}

func NewQSvgWidgetFromPointer(ptr unsafe.Pointer) *QSvgWidget {
	var n = new(QSvgWidget)
	n.SetPointer(ptr)
	return n
}

func newQSvgWidgetFromPointer(ptr unsafe.Pointer) *QSvgWidget {
	var n = NewQSvgWidgetFromPointer(ptr)
	for len(n.ObjectName()) < len("QSvgWidget_") {
		n.SetObjectName("QSvgWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QSvgWidget) QSvgWidget_PTR() *QSvgWidget {
	return ptr
}

func NewQSvgWidget(parent widgets.QWidget_ITF) *QSvgWidget {
	defer qt.Recovering("QSvgWidget::QSvgWidget")

	return newQSvgWidgetFromPointer(C.QSvgWidget_NewQSvgWidget(widgets.PointerFromQWidget(parent)))
}

func NewQSvgWidget2(file string, parent widgets.QWidget_ITF) *QSvgWidget {
	defer qt.Recovering("QSvgWidget::QSvgWidget")

	return newQSvgWidgetFromPointer(C.QSvgWidget_NewQSvgWidget2(C.CString(file), widgets.PointerFromQWidget(parent)))
}

func (ptr *QSvgWidget) Load2(contents core.QByteArray_ITF) {
	defer qt.Recovering("QSvgWidget::load")

	if ptr.Pointer() != nil {
		C.QSvgWidget_Load2(ptr.Pointer(), core.PointerFromQByteArray(contents))
	}
}

func (ptr *QSvgWidget) Load(file string) {
	defer qt.Recovering("QSvgWidget::load")

	if ptr.Pointer() != nil {
		C.QSvgWidget_Load(ptr.Pointer(), C.CString(file))
	}
}

func (ptr *QSvgWidget) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QSvgWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QSvgWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQSvgWidgetPaintEvent
func callbackQSvgWidgetPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QSvgWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QSvgWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QSvgWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QSvgWidget) Renderer() *QSvgRenderer {
	defer qt.Recovering("QSvgWidget::renderer")

	if ptr.Pointer() != nil {
		return NewQSvgRendererFromPointer(C.QSvgWidget_Renderer(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSvgWidget) SizeHint() *core.QSize {
	defer qt.Recovering("QSvgWidget::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QSvgWidget_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSvgWidget) DestroyQSvgWidget() {
	defer qt.Recovering("QSvgWidget::~QSvgWidget")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DestroyQSvgWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSvgWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QSvgWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QSvgWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQSvgWidgetActionEvent
func callbackQSvgWidgetActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QSvgWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QSvgWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QSvgWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QSvgWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QSvgWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQSvgWidgetDragEnterEvent
func callbackQSvgWidgetDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QSvgWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QSvgWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QSvgWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QSvgWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QSvgWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQSvgWidgetDragLeaveEvent
func callbackQSvgWidgetDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QSvgWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QSvgWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QSvgWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QSvgWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QSvgWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQSvgWidgetDragMoveEvent
func callbackQSvgWidgetDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QSvgWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QSvgWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QSvgWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QSvgWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QSvgWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQSvgWidgetDropEvent
func callbackQSvgWidgetDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QSvgWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QSvgWidget) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QSvgWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSvgWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QSvgWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQSvgWidgetEnterEvent
func callbackQSvgWidgetEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSvgWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSvgWidget) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSvgWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QSvgWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QSvgWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQSvgWidgetFocusInEvent
func callbackQSvgWidgetFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSvgWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QSvgWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSvgWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QSvgWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QSvgWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQSvgWidgetFocusOutEvent
func callbackQSvgWidgetFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSvgWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QSvgWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSvgWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QSvgWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QSvgWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQSvgWidgetHideEvent
func callbackQSvgWidgetHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QSvgWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QSvgWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QSvgWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSvgWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QSvgWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQSvgWidgetLeaveEvent
func callbackQSvgWidgetLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSvgWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSvgWidget) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSvgWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QSvgWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QSvgWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQSvgWidgetMoveEvent
func callbackQSvgWidgetMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QSvgWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QSvgWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QSvgWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QSvgWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QSvgWidget) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QSvgWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQSvgWidgetSetVisible
func callbackQSvgWidgetSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QSvgWidget::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QSvgWidget) SetVisible(visible bool) {
	defer qt.Recovering("QSvgWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QSvgWidget_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QSvgWidget) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QSvgWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QSvgWidget_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QSvgWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QSvgWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QSvgWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQSvgWidgetShowEvent
func callbackQSvgWidgetShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QSvgWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QSvgWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QSvgWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSvgWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QSvgWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQSvgWidgetChangeEvent
func callbackQSvgWidgetChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSvgWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSvgWidget) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSvgWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QSvgWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QSvgWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQSvgWidgetCloseEvent
func callbackQSvgWidgetCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QSvgWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QSvgWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QSvgWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QSvgWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QSvgWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQSvgWidgetContextMenuEvent
func callbackQSvgWidgetContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QSvgWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QSvgWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QSvgWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QSvgWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QSvgWidget) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QSvgWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQSvgWidgetInitPainter
func callbackQSvgWidgetInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQSvgWidgetFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QSvgWidget) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QSvgWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QSvgWidget_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QSvgWidget) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QSvgWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QSvgWidget_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QSvgWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QSvgWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QSvgWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQSvgWidgetInputMethodEvent
func callbackQSvgWidgetInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QSvgWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QSvgWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QSvgWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QSvgWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QSvgWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQSvgWidgetKeyPressEvent
func callbackQSvgWidgetKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSvgWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QSvgWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSvgWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QSvgWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QSvgWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQSvgWidgetKeyReleaseEvent
func callbackQSvgWidgetKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSvgWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QSvgWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSvgWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSvgWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QSvgWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQSvgWidgetMouseDoubleClickEvent
func callbackQSvgWidgetMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSvgWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSvgWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSvgWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSvgWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QSvgWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQSvgWidgetMouseMoveEvent
func callbackQSvgWidgetMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSvgWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSvgWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSvgWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSvgWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QSvgWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQSvgWidgetMousePressEvent
func callbackQSvgWidgetMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSvgWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSvgWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSvgWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSvgWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QSvgWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQSvgWidgetMouseReleaseEvent
func callbackQSvgWidgetMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSvgWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSvgWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSvgWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QSvgWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QSvgWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQSvgWidgetResizeEvent
func callbackQSvgWidgetResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QSvgWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QSvgWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QSvgWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QSvgWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QSvgWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQSvgWidgetTabletEvent
func callbackQSvgWidgetTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QSvgWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QSvgWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QSvgWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QSvgWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QSvgWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQSvgWidgetWheelEvent
func callbackQSvgWidgetWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QSvgWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QSvgWidget) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QSvgWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSvgWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSvgWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSvgWidgetTimerEvent
func callbackQSvgWidgetTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSvgWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSvgWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSvgWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSvgWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSvgWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSvgWidgetChildEvent
func callbackQSvgWidgetChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSvgWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSvgWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSvgWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSvgWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSvgWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSvgWidgetCustomEvent
func callbackQSvgWidgetCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSvgWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSvgWidget) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSvgWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
