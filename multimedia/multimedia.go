package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/network"
	"strings"
	"unsafe"
)

type QAbstractPlanarVideoBuffer struct {
	QAbstractVideoBuffer
}

type QAbstractPlanarVideoBuffer_ITF interface {
	QAbstractVideoBuffer_ITF
	QAbstractPlanarVideoBuffer_PTR() *QAbstractPlanarVideoBuffer
}

func PointerFromQAbstractPlanarVideoBuffer(ptr QAbstractPlanarVideoBuffer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractPlanarVideoBuffer_PTR().Pointer()
	}
	return nil
}

func NewQAbstractPlanarVideoBufferFromPointer(ptr unsafe.Pointer) *QAbstractPlanarVideoBuffer {
	var n = new(QAbstractPlanarVideoBuffer)
	n.SetPointer(ptr)
	return n
}

func newQAbstractPlanarVideoBufferFromPointer(ptr unsafe.Pointer) *QAbstractPlanarVideoBuffer {
	var n = NewQAbstractPlanarVideoBufferFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QAbstractPlanarVideoBuffer_") {
		n.SetObjectNameAbs("QAbstractPlanarVideoBuffer_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractPlanarVideoBuffer) QAbstractPlanarVideoBuffer_PTR() *QAbstractPlanarVideoBuffer {
	return ptr
}

func (ptr *QAbstractPlanarVideoBuffer) DestroyQAbstractPlanarVideoBuffer() {
	defer qt.Recovering("QAbstractPlanarVideoBuffer::~QAbstractPlanarVideoBuffer")

	if ptr.Pointer() != nil {
		C.QAbstractPlanarVideoBuffer_DestroyQAbstractPlanarVideoBuffer(ptr.Pointer())
	}
}

func (ptr *QAbstractPlanarVideoBuffer) ObjectNameAbs() string {
	defer qt.Recovering("QAbstractPlanarVideoBuffer::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractPlanarVideoBuffer_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstractPlanarVideoBuffer) SetObjectNameAbs(name string) {
	defer qt.Recovering("QAbstractPlanarVideoBuffer::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QAbstractPlanarVideoBuffer_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QAbstractPlanarVideoBuffer) ConnectRelease(f func()) {
	defer qt.Recovering("connect QAbstractPlanarVideoBuffer::release")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "release", f)
	}
}

func (ptr *QAbstractPlanarVideoBuffer) DisconnectRelease() {
	defer qt.Recovering("disconnect QAbstractPlanarVideoBuffer::release")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "release")
	}
}

//export callbackQAbstractPlanarVideoBufferRelease
func callbackQAbstractPlanarVideoBufferRelease(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractPlanarVideoBuffer::release")

	if signal := qt.GetSignal(C.GoString(ptrName), "release"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPlanarVideoBufferFromPointer(ptr).ReleaseDefault()
	}
}

func (ptr *QAbstractPlanarVideoBuffer) Release() {
	defer qt.Recovering("QAbstractPlanarVideoBuffer::release")

	if ptr.Pointer() != nil {
		C.QAbstractPlanarVideoBuffer_Release(ptr.Pointer())
	}
}

func (ptr *QAbstractPlanarVideoBuffer) ReleaseDefault() {
	defer qt.Recovering("QAbstractPlanarVideoBuffer::release")

	if ptr.Pointer() != nil {
		C.QAbstractPlanarVideoBuffer_ReleaseDefault(ptr.Pointer())
	}
}

type QAbstractVideoBuffer struct {
	ptr unsafe.Pointer
}

type QAbstractVideoBuffer_ITF interface {
	QAbstractVideoBuffer_PTR() *QAbstractVideoBuffer
}

func (p *QAbstractVideoBuffer) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAbstractVideoBuffer) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAbstractVideoBuffer(ptr QAbstractVideoBuffer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractVideoBuffer_PTR().Pointer()
	}
	return nil
}

func NewQAbstractVideoBufferFromPointer(ptr unsafe.Pointer) *QAbstractVideoBuffer {
	var n = new(QAbstractVideoBuffer)
	n.SetPointer(ptr)
	return n
}

func newQAbstractVideoBufferFromPointer(ptr unsafe.Pointer) *QAbstractVideoBuffer {
	var n = NewQAbstractVideoBufferFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QAbstractVideoBuffer_") {
		n.SetObjectNameAbs("QAbstractVideoBuffer_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractVideoBuffer) QAbstractVideoBuffer_PTR() *QAbstractVideoBuffer {
	return ptr
}

//QAbstractVideoBuffer::HandleType
type QAbstractVideoBuffer__HandleType int64

const (
	QAbstractVideoBuffer__NoHandle         = QAbstractVideoBuffer__HandleType(0)
	QAbstractVideoBuffer__GLTextureHandle  = QAbstractVideoBuffer__HandleType(1)
	QAbstractVideoBuffer__XvShmImageHandle = QAbstractVideoBuffer__HandleType(2)
	QAbstractVideoBuffer__CoreImageHandle  = QAbstractVideoBuffer__HandleType(3)
	QAbstractVideoBuffer__QPixmapHandle    = QAbstractVideoBuffer__HandleType(4)
	QAbstractVideoBuffer__EGLImageHandle   = QAbstractVideoBuffer__HandleType(5)
	QAbstractVideoBuffer__UserHandle       = QAbstractVideoBuffer__HandleType(1000)
)

//QAbstractVideoBuffer::MapMode
type QAbstractVideoBuffer__MapMode int64

const (
	QAbstractVideoBuffer__NotMapped = QAbstractVideoBuffer__MapMode(0x00)
	QAbstractVideoBuffer__ReadOnly  = QAbstractVideoBuffer__MapMode(0x01)
	QAbstractVideoBuffer__WriteOnly = QAbstractVideoBuffer__MapMode(0x02)
	QAbstractVideoBuffer__ReadWrite = QAbstractVideoBuffer__MapMode(QAbstractVideoBuffer__ReadOnly | QAbstractVideoBuffer__WriteOnly)
)

func (ptr *QAbstractVideoBuffer) Handle() *core.QVariant {
	defer qt.Recovering("QAbstractVideoBuffer::handle")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QAbstractVideoBuffer_Handle(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractVideoBuffer) HandleType() QAbstractVideoBuffer__HandleType {
	defer qt.Recovering("QAbstractVideoBuffer::handleType")

	if ptr.Pointer() != nil {
		return QAbstractVideoBuffer__HandleType(C.QAbstractVideoBuffer_HandleType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractVideoBuffer) MapMode() QAbstractVideoBuffer__MapMode {
	defer qt.Recovering("QAbstractVideoBuffer::mapMode")

	if ptr.Pointer() != nil {
		return QAbstractVideoBuffer__MapMode(C.QAbstractVideoBuffer_MapMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractVideoBuffer) ConnectRelease(f func()) {
	defer qt.Recovering("connect QAbstractVideoBuffer::release")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "release", f)
	}
}

func (ptr *QAbstractVideoBuffer) DisconnectRelease() {
	defer qt.Recovering("disconnect QAbstractVideoBuffer::release")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "release")
	}
}

//export callbackQAbstractVideoBufferRelease
func callbackQAbstractVideoBufferRelease(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractVideoBuffer::release")

	if signal := qt.GetSignal(C.GoString(ptrName), "release"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractVideoBufferFromPointer(ptr).ReleaseDefault()
	}
}

func (ptr *QAbstractVideoBuffer) Release() {
	defer qt.Recovering("QAbstractVideoBuffer::release")

	if ptr.Pointer() != nil {
		C.QAbstractVideoBuffer_Release(ptr.Pointer())
	}
}

func (ptr *QAbstractVideoBuffer) ReleaseDefault() {
	defer qt.Recovering("QAbstractVideoBuffer::release")

	if ptr.Pointer() != nil {
		C.QAbstractVideoBuffer_ReleaseDefault(ptr.Pointer())
	}
}

func (ptr *QAbstractVideoBuffer) Unmap() {
	defer qt.Recovering("QAbstractVideoBuffer::unmap")

	if ptr.Pointer() != nil {
		C.QAbstractVideoBuffer_Unmap(ptr.Pointer())
	}
}

func (ptr *QAbstractVideoBuffer) DestroyQAbstractVideoBuffer() {
	defer qt.Recovering("QAbstractVideoBuffer::~QAbstractVideoBuffer")

	if ptr.Pointer() != nil {
		C.QAbstractVideoBuffer_DestroyQAbstractVideoBuffer(ptr.Pointer())
	}
}

func (ptr *QAbstractVideoBuffer) ObjectNameAbs() string {
	defer qt.Recovering("QAbstractVideoBuffer::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractVideoBuffer_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstractVideoBuffer) SetObjectNameAbs(name string) {
	defer qt.Recovering("QAbstractVideoBuffer::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QAbstractVideoBuffer_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QAbstractVideoFilter struct {
	core.QObject
}

type QAbstractVideoFilter_ITF interface {
	core.QObject_ITF
	QAbstractVideoFilter_PTR() *QAbstractVideoFilter
}

func PointerFromQAbstractVideoFilter(ptr QAbstractVideoFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractVideoFilter_PTR().Pointer()
	}
	return nil
}

func NewQAbstractVideoFilterFromPointer(ptr unsafe.Pointer) *QAbstractVideoFilter {
	var n = new(QAbstractVideoFilter)
	n.SetPointer(ptr)
	return n
}

func newQAbstractVideoFilterFromPointer(ptr unsafe.Pointer) *QAbstractVideoFilter {
	var n = NewQAbstractVideoFilterFromPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractVideoFilter_") {
		n.SetObjectName("QAbstractVideoFilter_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractVideoFilter) QAbstractVideoFilter_PTR() *QAbstractVideoFilter {
	return ptr
}

func (ptr *QAbstractVideoFilter) IsActive() bool {
	defer qt.Recovering("QAbstractVideoFilter::isActive")

	if ptr.Pointer() != nil {
		return C.QAbstractVideoFilter_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractVideoFilter) SetActive(v bool) {
	defer qt.Recovering("QAbstractVideoFilter::setActive")

	if ptr.Pointer() != nil {
		C.QAbstractVideoFilter_SetActive(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractVideoFilter) ConnectActiveChanged(f func()) {
	defer qt.Recovering("connect QAbstractVideoFilter::activeChanged")

	if ptr.Pointer() != nil {
		C.QAbstractVideoFilter_ConnectActiveChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeChanged", f)
	}
}

func (ptr *QAbstractVideoFilter) DisconnectActiveChanged() {
	defer qt.Recovering("disconnect QAbstractVideoFilter::activeChanged")

	if ptr.Pointer() != nil {
		C.QAbstractVideoFilter_DisconnectActiveChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeChanged")
	}
}

//export callbackQAbstractVideoFilterActiveChanged
func callbackQAbstractVideoFilterActiveChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractVideoFilter::activeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "activeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractVideoFilter) ActiveChanged() {
	defer qt.Recovering("QAbstractVideoFilter::activeChanged")

	if ptr.Pointer() != nil {
		C.QAbstractVideoFilter_ActiveChanged(ptr.Pointer())
	}
}

func (ptr *QAbstractVideoFilter) CreateFilterRunnable() *QVideoFilterRunnable {
	defer qt.Recovering("QAbstractVideoFilter::createFilterRunnable")

	if ptr.Pointer() != nil {
		return NewQVideoFilterRunnableFromPointer(C.QAbstractVideoFilter_CreateFilterRunnable(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractVideoFilter) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAbstractVideoFilter::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAbstractVideoFilter) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAbstractVideoFilter::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAbstractVideoFilterTimerEvent
func callbackQAbstractVideoFilterTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractVideoFilter::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractVideoFilterFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractVideoFilter) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractVideoFilter::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractVideoFilter_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractVideoFilter) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractVideoFilter::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractVideoFilter_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractVideoFilter) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAbstractVideoFilter::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAbstractVideoFilter) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAbstractVideoFilter::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAbstractVideoFilterChildEvent
func callbackQAbstractVideoFilterChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractVideoFilter::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractVideoFilterFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractVideoFilter) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractVideoFilter::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractVideoFilter_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractVideoFilter) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractVideoFilter::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractVideoFilter_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractVideoFilter) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractVideoFilter::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAbstractVideoFilter) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAbstractVideoFilter::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAbstractVideoFilterCustomEvent
func callbackQAbstractVideoFilterCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractVideoFilter::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractVideoFilterFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractVideoFilter) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractVideoFilter::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractVideoFilter_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractVideoFilter) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractVideoFilter::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractVideoFilter_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QAbstractVideoSurface struct {
	core.QObject
}

type QAbstractVideoSurface_ITF interface {
	core.QObject_ITF
	QAbstractVideoSurface_PTR() *QAbstractVideoSurface
}

func PointerFromQAbstractVideoSurface(ptr QAbstractVideoSurface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractVideoSurface_PTR().Pointer()
	}
	return nil
}

func NewQAbstractVideoSurfaceFromPointer(ptr unsafe.Pointer) *QAbstractVideoSurface {
	var n = new(QAbstractVideoSurface)
	n.SetPointer(ptr)
	return n
}

func newQAbstractVideoSurfaceFromPointer(ptr unsafe.Pointer) *QAbstractVideoSurface {
	var n = NewQAbstractVideoSurfaceFromPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractVideoSurface_") {
		n.SetObjectName("QAbstractVideoSurface_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractVideoSurface) QAbstractVideoSurface_PTR() *QAbstractVideoSurface {
	return ptr
}

//QAbstractVideoSurface::Error
type QAbstractVideoSurface__Error int64

const (
	QAbstractVideoSurface__NoError                = QAbstractVideoSurface__Error(0)
	QAbstractVideoSurface__UnsupportedFormatError = QAbstractVideoSurface__Error(1)
	QAbstractVideoSurface__IncorrectFormatError   = QAbstractVideoSurface__Error(2)
	QAbstractVideoSurface__StoppedError           = QAbstractVideoSurface__Error(3)
	QAbstractVideoSurface__ResourceError          = QAbstractVideoSurface__Error(4)
)

func (ptr *QAbstractVideoSurface) NativeResolution() *core.QSize {
	defer qt.Recovering("QAbstractVideoSurface::nativeResolution")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QAbstractVideoSurface_NativeResolution(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractVideoSurface) ConnectActiveChanged(f func(active bool)) {
	defer qt.Recovering("connect QAbstractVideoSurface::activeChanged")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_ConnectActiveChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeChanged", f)
	}
}

func (ptr *QAbstractVideoSurface) DisconnectActiveChanged() {
	defer qt.Recovering("disconnect QAbstractVideoSurface::activeChanged")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_DisconnectActiveChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeChanged")
	}
}

//export callbackQAbstractVideoSurfaceActiveChanged
func callbackQAbstractVideoSurfaceActiveChanged(ptr unsafe.Pointer, ptrName *C.char, active C.int) {
	defer qt.Recovering("callback QAbstractVideoSurface::activeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "activeChanged"); signal != nil {
		signal.(func(bool))(int(active) != 0)
	}

}

func (ptr *QAbstractVideoSurface) ActiveChanged(active bool) {
	defer qt.Recovering("QAbstractVideoSurface::activeChanged")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_ActiveChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(active)))
	}
}

func (ptr *QAbstractVideoSurface) Error() QAbstractVideoSurface__Error {
	defer qt.Recovering("QAbstractVideoSurface::error")

	if ptr.Pointer() != nil {
		return QAbstractVideoSurface__Error(C.QAbstractVideoSurface_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractVideoSurface) IsActive() bool {
	defer qt.Recovering("QAbstractVideoSurface::isActive")

	if ptr.Pointer() != nil {
		return C.QAbstractVideoSurface_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractVideoSurface) IsFormatSupported(format QVideoSurfaceFormat_ITF) bool {
	defer qt.Recovering("QAbstractVideoSurface::isFormatSupported")

	if ptr.Pointer() != nil {
		return C.QAbstractVideoSurface_IsFormatSupported(ptr.Pointer(), PointerFromQVideoSurfaceFormat(format)) != 0
	}
	return false
}

func (ptr *QAbstractVideoSurface) ConnectNativeResolutionChanged(f func(resolution *core.QSize)) {
	defer qt.Recovering("connect QAbstractVideoSurface::nativeResolutionChanged")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_ConnectNativeResolutionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "nativeResolutionChanged", f)
	}
}

func (ptr *QAbstractVideoSurface) DisconnectNativeResolutionChanged() {
	defer qt.Recovering("disconnect QAbstractVideoSurface::nativeResolutionChanged")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_DisconnectNativeResolutionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "nativeResolutionChanged")
	}
}

//export callbackQAbstractVideoSurfaceNativeResolutionChanged
func callbackQAbstractVideoSurfaceNativeResolutionChanged(ptr unsafe.Pointer, ptrName *C.char, resolution unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractVideoSurface::nativeResolutionChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "nativeResolutionChanged"); signal != nil {
		signal.(func(*core.QSize))(core.NewQSizeFromPointer(resolution))
	}

}

func (ptr *QAbstractVideoSurface) NativeResolutionChanged(resolution core.QSize_ITF) {
	defer qt.Recovering("QAbstractVideoSurface::nativeResolutionChanged")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_NativeResolutionChanged(ptr.Pointer(), core.PointerFromQSize(resolution))
	}
}

func (ptr *QAbstractVideoSurface) NearestFormat(format QVideoSurfaceFormat_ITF) *QVideoSurfaceFormat {
	defer qt.Recovering("QAbstractVideoSurface::nearestFormat")

	if ptr.Pointer() != nil {
		return NewQVideoSurfaceFormatFromPointer(C.QAbstractVideoSurface_NearestFormat(ptr.Pointer(), PointerFromQVideoSurfaceFormat(format)))
	}
	return nil
}

func (ptr *QAbstractVideoSurface) Present(frame QVideoFrame_ITF) bool {
	defer qt.Recovering("QAbstractVideoSurface::present")

	if ptr.Pointer() != nil {
		return C.QAbstractVideoSurface_Present(ptr.Pointer(), PointerFromQVideoFrame(frame)) != 0
	}
	return false
}

func (ptr *QAbstractVideoSurface) Start(format QVideoSurfaceFormat_ITF) bool {
	defer qt.Recovering("QAbstractVideoSurface::start")

	if ptr.Pointer() != nil {
		return C.QAbstractVideoSurface_Start(ptr.Pointer(), PointerFromQVideoSurfaceFormat(format)) != 0
	}
	return false
}

func (ptr *QAbstractVideoSurface) ConnectStop(f func()) {
	defer qt.Recovering("connect QAbstractVideoSurface::stop")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "stop", f)
	}
}

func (ptr *QAbstractVideoSurface) DisconnectStop() {
	defer qt.Recovering("disconnect QAbstractVideoSurface::stop")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "stop")
	}
}

//export callbackQAbstractVideoSurfaceStop
func callbackQAbstractVideoSurfaceStop(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractVideoSurface::stop")

	if signal := qt.GetSignal(C.GoString(ptrName), "stop"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractVideoSurfaceFromPointer(ptr).StopDefault()
	}
}

func (ptr *QAbstractVideoSurface) Stop() {
	defer qt.Recovering("QAbstractVideoSurface::stop")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_Stop(ptr.Pointer())
	}
}

func (ptr *QAbstractVideoSurface) StopDefault() {
	defer qt.Recovering("QAbstractVideoSurface::stop")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_StopDefault(ptr.Pointer())
	}
}

func (ptr *QAbstractVideoSurface) ConnectSupportedFormatsChanged(f func()) {
	defer qt.Recovering("connect QAbstractVideoSurface::supportedFormatsChanged")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_ConnectSupportedFormatsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "supportedFormatsChanged", f)
	}
}

func (ptr *QAbstractVideoSurface) DisconnectSupportedFormatsChanged() {
	defer qt.Recovering("disconnect QAbstractVideoSurface::supportedFormatsChanged")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_DisconnectSupportedFormatsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "supportedFormatsChanged")
	}
}

//export callbackQAbstractVideoSurfaceSupportedFormatsChanged
func callbackQAbstractVideoSurfaceSupportedFormatsChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractVideoSurface::supportedFormatsChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "supportedFormatsChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractVideoSurface) SupportedFormatsChanged() {
	defer qt.Recovering("QAbstractVideoSurface::supportedFormatsChanged")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_SupportedFormatsChanged(ptr.Pointer())
	}
}

func (ptr *QAbstractVideoSurface) SurfaceFormat() *QVideoSurfaceFormat {
	defer qt.Recovering("QAbstractVideoSurface::surfaceFormat")

	if ptr.Pointer() != nil {
		return NewQVideoSurfaceFormatFromPointer(C.QAbstractVideoSurface_SurfaceFormat(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractVideoSurface) ConnectSurfaceFormatChanged(f func(format *QVideoSurfaceFormat)) {
	defer qt.Recovering("connect QAbstractVideoSurface::surfaceFormatChanged")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_ConnectSurfaceFormatChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "surfaceFormatChanged", f)
	}
}

func (ptr *QAbstractVideoSurface) DisconnectSurfaceFormatChanged() {
	defer qt.Recovering("disconnect QAbstractVideoSurface::surfaceFormatChanged")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_DisconnectSurfaceFormatChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "surfaceFormatChanged")
	}
}

//export callbackQAbstractVideoSurfaceSurfaceFormatChanged
func callbackQAbstractVideoSurfaceSurfaceFormatChanged(ptr unsafe.Pointer, ptrName *C.char, format unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractVideoSurface::surfaceFormatChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "surfaceFormatChanged"); signal != nil {
		signal.(func(*QVideoSurfaceFormat))(NewQVideoSurfaceFormatFromPointer(format))
	}

}

func (ptr *QAbstractVideoSurface) SurfaceFormatChanged(format QVideoSurfaceFormat_ITF) {
	defer qt.Recovering("QAbstractVideoSurface::surfaceFormatChanged")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_SurfaceFormatChanged(ptr.Pointer(), PointerFromQVideoSurfaceFormat(format))
	}
}

func (ptr *QAbstractVideoSurface) DestroyQAbstractVideoSurface() {
	defer qt.Recovering("QAbstractVideoSurface::~QAbstractVideoSurface")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_DestroyQAbstractVideoSurface(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractVideoSurface) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAbstractVideoSurface::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAbstractVideoSurface) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAbstractVideoSurface::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAbstractVideoSurfaceTimerEvent
func callbackQAbstractVideoSurfaceTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractVideoSurface::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractVideoSurfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractVideoSurface) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractVideoSurface::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractVideoSurface) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractVideoSurface::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractVideoSurface) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAbstractVideoSurface::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAbstractVideoSurface) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAbstractVideoSurface::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAbstractVideoSurfaceChildEvent
func callbackQAbstractVideoSurfaceChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractVideoSurface::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractVideoSurfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractVideoSurface) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractVideoSurface::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractVideoSurface) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractVideoSurface::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractVideoSurface) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractVideoSurface::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAbstractVideoSurface) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAbstractVideoSurface::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAbstractVideoSurfaceCustomEvent
func callbackQAbstractVideoSurfaceCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractVideoSurface::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractVideoSurfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractVideoSurface) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractVideoSurface::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractVideoSurface) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractVideoSurface::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//QAudio::Error
type QAudio__Error int64

const (
	QAudio__NoError       = QAudio__Error(0)
	QAudio__OpenError     = QAudio__Error(1)
	QAudio__IOError       = QAudio__Error(2)
	QAudio__UnderrunError = QAudio__Error(3)
	QAudio__FatalError    = QAudio__Error(4)
)

//QAudio::Mode
type QAudio__Mode int64

const (
	QAudio__AudioInput  = QAudio__Mode(0)
	QAudio__AudioOutput = QAudio__Mode(1)
)

//QAudio::State
type QAudio__State int64

const (
	QAudio__ActiveState    = QAudio__State(0)
	QAudio__SuspendedState = QAudio__State(1)
	QAudio__StoppedState   = QAudio__State(2)
	QAudio__IdleState      = QAudio__State(3)
)

type QAudioBuffer struct {
	ptr unsafe.Pointer
}

type QAudioBuffer_ITF interface {
	QAudioBuffer_PTR() *QAudioBuffer
}

func (p *QAudioBuffer) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAudioBuffer) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAudioBuffer(ptr QAudioBuffer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioBuffer_PTR().Pointer()
	}
	return nil
}

func NewQAudioBufferFromPointer(ptr unsafe.Pointer) *QAudioBuffer {
	var n = new(QAudioBuffer)
	n.SetPointer(ptr)
	return n
}

func newQAudioBufferFromPointer(ptr unsafe.Pointer) *QAudioBuffer {
	var n = NewQAudioBufferFromPointer(ptr)
	return n
}

func (ptr *QAudioBuffer) QAudioBuffer_PTR() *QAudioBuffer {
	return ptr
}

func NewQAudioBuffer() *QAudioBuffer {
	defer qt.Recovering("QAudioBuffer::QAudioBuffer")

	return newQAudioBufferFromPointer(C.QAudioBuffer_NewQAudioBuffer())
}

func NewQAudioBuffer3(other QAudioBuffer_ITF) *QAudioBuffer {
	defer qt.Recovering("QAudioBuffer::QAudioBuffer")

	return newQAudioBufferFromPointer(C.QAudioBuffer_NewQAudioBuffer3(PointerFromQAudioBuffer(other)))
}

func NewQAudioBuffer4(data string, format QAudioFormat_ITF, startTime int64) *QAudioBuffer {
	defer qt.Recovering("QAudioBuffer::QAudioBuffer")

	return newQAudioBufferFromPointer(C.QAudioBuffer_NewQAudioBuffer4(C.CString(data), PointerFromQAudioFormat(format), C.longlong(startTime)))
}

func NewQAudioBuffer5(numFrames int, format QAudioFormat_ITF, startTime int64) *QAudioBuffer {
	defer qt.Recovering("QAudioBuffer::QAudioBuffer")

	return newQAudioBufferFromPointer(C.QAudioBuffer_NewQAudioBuffer5(C.int(numFrames), PointerFromQAudioFormat(format), C.longlong(startTime)))
}

func (ptr *QAudioBuffer) ByteCount() int {
	defer qt.Recovering("QAudioBuffer::byteCount")

	if ptr.Pointer() != nil {
		return int(C.QAudioBuffer_ByteCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioBuffer) ConstData() unsafe.Pointer {
	defer qt.Recovering("QAudioBuffer::constData")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QAudioBuffer_ConstData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAudioBuffer) Data2() unsafe.Pointer {
	defer qt.Recovering("QAudioBuffer::data")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QAudioBuffer_Data2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAudioBuffer) Data() unsafe.Pointer {
	defer qt.Recovering("QAudioBuffer::data")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QAudioBuffer_Data(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAudioBuffer) Duration() int64 {
	defer qt.Recovering("QAudioBuffer::duration")

	if ptr.Pointer() != nil {
		return int64(C.QAudioBuffer_Duration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioBuffer) Format() *QAudioFormat {
	defer qt.Recovering("QAudioBuffer::format")

	if ptr.Pointer() != nil {
		return NewQAudioFormatFromPointer(C.QAudioBuffer_Format(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAudioBuffer) FrameCount() int {
	defer qt.Recovering("QAudioBuffer::frameCount")

	if ptr.Pointer() != nil {
		return int(C.QAudioBuffer_FrameCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioBuffer) IsValid() bool {
	defer qt.Recovering("QAudioBuffer::isValid")

	if ptr.Pointer() != nil {
		return C.QAudioBuffer_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAudioBuffer) SampleCount() int {
	defer qt.Recovering("QAudioBuffer::sampleCount")

	if ptr.Pointer() != nil {
		return int(C.QAudioBuffer_SampleCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioBuffer) StartTime() int64 {
	defer qt.Recovering("QAudioBuffer::startTime")

	if ptr.Pointer() != nil {
		return int64(C.QAudioBuffer_StartTime(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioBuffer) DestroyQAudioBuffer() {
	defer qt.Recovering("QAudioBuffer::~QAudioBuffer")

	if ptr.Pointer() != nil {
		C.QAudioBuffer_DestroyQAudioBuffer(ptr.Pointer())
	}
}

type QAudioDecoder struct {
	QMediaObject
}

type QAudioDecoder_ITF interface {
	QMediaObject_ITF
	QAudioDecoder_PTR() *QAudioDecoder
}

func PointerFromQAudioDecoder(ptr QAudioDecoder_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioDecoder_PTR().Pointer()
	}
	return nil
}

func NewQAudioDecoderFromPointer(ptr unsafe.Pointer) *QAudioDecoder {
	var n = new(QAudioDecoder)
	n.SetPointer(ptr)
	return n
}

func newQAudioDecoderFromPointer(ptr unsafe.Pointer) *QAudioDecoder {
	var n = NewQAudioDecoderFromPointer(ptr)
	for len(n.ObjectName()) < len("QAudioDecoder_") {
		n.SetObjectName("QAudioDecoder_" + qt.Identifier())
	}
	return n
}

func (ptr *QAudioDecoder) QAudioDecoder_PTR() *QAudioDecoder {
	return ptr
}

//QAudioDecoder::Error
type QAudioDecoder__Error int64

const (
	QAudioDecoder__NoError             = QAudioDecoder__Error(0)
	QAudioDecoder__ResourceError       = QAudioDecoder__Error(1)
	QAudioDecoder__FormatError         = QAudioDecoder__Error(2)
	QAudioDecoder__AccessDeniedError   = QAudioDecoder__Error(3)
	QAudioDecoder__ServiceMissingError = QAudioDecoder__Error(4)
)

//QAudioDecoder::State
type QAudioDecoder__State int64

const (
	QAudioDecoder__StoppedState  = QAudioDecoder__State(0)
	QAudioDecoder__DecodingState = QAudioDecoder__State(1)
)

func (ptr *QAudioDecoder) ErrorString() string {
	defer qt.Recovering("QAudioDecoder::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioDecoder_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioDecoder) State() QAudioDecoder__State {
	defer qt.Recovering("QAudioDecoder::state")

	if ptr.Pointer() != nil {
		return QAudioDecoder__State(C.QAudioDecoder_State(ptr.Pointer()))
	}
	return 0
}

func NewQAudioDecoder(parent core.QObject_ITF) *QAudioDecoder {
	defer qt.Recovering("QAudioDecoder::QAudioDecoder")

	return newQAudioDecoderFromPointer(C.QAudioDecoder_NewQAudioDecoder(core.PointerFromQObject(parent)))
}

func (ptr *QAudioDecoder) AudioFormat() *QAudioFormat {
	defer qt.Recovering("QAudioDecoder::audioFormat")

	if ptr.Pointer() != nil {
		return NewQAudioFormatFromPointer(C.QAudioDecoder_AudioFormat(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAudioDecoder) BufferAvailable() bool {
	defer qt.Recovering("QAudioDecoder::bufferAvailable")

	if ptr.Pointer() != nil {
		return C.QAudioDecoder_BufferAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAudioDecoder) ConnectBufferAvailableChanged(f func(available bool)) {
	defer qt.Recovering("connect QAudioDecoder::bufferAvailableChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_ConnectBufferAvailableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bufferAvailableChanged", f)
	}
}

func (ptr *QAudioDecoder) DisconnectBufferAvailableChanged() {
	defer qt.Recovering("disconnect QAudioDecoder::bufferAvailableChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_DisconnectBufferAvailableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bufferAvailableChanged")
	}
}

//export callbackQAudioDecoderBufferAvailableChanged
func callbackQAudioDecoderBufferAvailableChanged(ptr unsafe.Pointer, ptrName *C.char, available C.int) {
	defer qt.Recovering("callback QAudioDecoder::bufferAvailableChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "bufferAvailableChanged"); signal != nil {
		signal.(func(bool))(int(available) != 0)
	}

}

func (ptr *QAudioDecoder) BufferAvailableChanged(available bool) {
	defer qt.Recovering("QAudioDecoder::bufferAvailableChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_BufferAvailableChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(available)))
	}
}

func (ptr *QAudioDecoder) ConnectBufferReady(f func()) {
	defer qt.Recovering("connect QAudioDecoder::bufferReady")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_ConnectBufferReady(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bufferReady", f)
	}
}

func (ptr *QAudioDecoder) DisconnectBufferReady() {
	defer qt.Recovering("disconnect QAudioDecoder::bufferReady")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_DisconnectBufferReady(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bufferReady")
	}
}

//export callbackQAudioDecoderBufferReady
func callbackQAudioDecoderBufferReady(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAudioDecoder::bufferReady")

	if signal := qt.GetSignal(C.GoString(ptrName), "bufferReady"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAudioDecoder) BufferReady() {
	defer qt.Recovering("QAudioDecoder::bufferReady")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_BufferReady(ptr.Pointer())
	}
}

func (ptr *QAudioDecoder) Duration() int64 {
	defer qt.Recovering("QAudioDecoder::duration")

	if ptr.Pointer() != nil {
		return int64(C.QAudioDecoder_Duration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioDecoder) ConnectDurationChanged(f func(duration int64)) {
	defer qt.Recovering("connect QAudioDecoder::durationChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_ConnectDurationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "durationChanged", f)
	}
}

func (ptr *QAudioDecoder) DisconnectDurationChanged() {
	defer qt.Recovering("disconnect QAudioDecoder::durationChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_DisconnectDurationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "durationChanged")
	}
}

//export callbackQAudioDecoderDurationChanged
func callbackQAudioDecoderDurationChanged(ptr unsafe.Pointer, ptrName *C.char, duration C.longlong) {
	defer qt.Recovering("callback QAudioDecoder::durationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "durationChanged"); signal != nil {
		signal.(func(int64))(int64(duration))
	}

}

func (ptr *QAudioDecoder) DurationChanged(duration int64) {
	defer qt.Recovering("QAudioDecoder::durationChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_DurationChanged(ptr.Pointer(), C.longlong(duration))
	}
}

func (ptr *QAudioDecoder) ConnectError2(f func(error QAudioDecoder__Error)) {
	defer qt.Recovering("connect QAudioDecoder::error")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QAudioDecoder) DisconnectError2() {
	defer qt.Recovering("disconnect QAudioDecoder::error")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQAudioDecoderError2
func callbackQAudioDecoderError2(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QAudioDecoder::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QAudioDecoder__Error))(QAudioDecoder__Error(error))
	}

}

func (ptr *QAudioDecoder) Error2(error QAudioDecoder__Error) {
	defer qt.Recovering("QAudioDecoder::error")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_Error2(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QAudioDecoder) Error() QAudioDecoder__Error {
	defer qt.Recovering("QAudioDecoder::error")

	if ptr.Pointer() != nil {
		return QAudioDecoder__Error(C.QAudioDecoder_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioDecoder) ConnectFinished(f func()) {
	defer qt.Recovering("connect QAudioDecoder::finished")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QAudioDecoder) DisconnectFinished() {
	defer qt.Recovering("disconnect QAudioDecoder::finished")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQAudioDecoderFinished
func callbackQAudioDecoderFinished(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAudioDecoder::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAudioDecoder) Finished() {
	defer qt.Recovering("QAudioDecoder::finished")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_Finished(ptr.Pointer())
	}
}

func (ptr *QAudioDecoder) ConnectFormatChanged(f func(format *QAudioFormat)) {
	defer qt.Recovering("connect QAudioDecoder::formatChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_ConnectFormatChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "formatChanged", f)
	}
}

func (ptr *QAudioDecoder) DisconnectFormatChanged() {
	defer qt.Recovering("disconnect QAudioDecoder::formatChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_DisconnectFormatChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "formatChanged")
	}
}

//export callbackQAudioDecoderFormatChanged
func callbackQAudioDecoderFormatChanged(ptr unsafe.Pointer, ptrName *C.char, format unsafe.Pointer) {
	defer qt.Recovering("callback QAudioDecoder::formatChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "formatChanged"); signal != nil {
		signal.(func(*QAudioFormat))(NewQAudioFormatFromPointer(format))
	}

}

func (ptr *QAudioDecoder) FormatChanged(format QAudioFormat_ITF) {
	defer qt.Recovering("QAudioDecoder::formatChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_FormatChanged(ptr.Pointer(), PointerFromQAudioFormat(format))
	}
}

func QAudioDecoder_HasSupport(mimeType string, codecs []string) QMultimedia__SupportEstimate {
	defer qt.Recovering("QAudioDecoder::hasSupport")

	return QMultimedia__SupportEstimate(C.QAudioDecoder_QAudioDecoder_HasSupport(C.CString(mimeType), C.CString(strings.Join(codecs, "|"))))
}

func (ptr *QAudioDecoder) Position() int64 {
	defer qt.Recovering("QAudioDecoder::position")

	if ptr.Pointer() != nil {
		return int64(C.QAudioDecoder_Position(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioDecoder) ConnectPositionChanged(f func(position int64)) {
	defer qt.Recovering("connect QAudioDecoder::positionChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_ConnectPositionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "positionChanged", f)
	}
}

func (ptr *QAudioDecoder) DisconnectPositionChanged() {
	defer qt.Recovering("disconnect QAudioDecoder::positionChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_DisconnectPositionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "positionChanged")
	}
}

//export callbackQAudioDecoderPositionChanged
func callbackQAudioDecoderPositionChanged(ptr unsafe.Pointer, ptrName *C.char, position C.longlong) {
	defer qt.Recovering("callback QAudioDecoder::positionChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "positionChanged"); signal != nil {
		signal.(func(int64))(int64(position))
	}

}

func (ptr *QAudioDecoder) PositionChanged(position int64) {
	defer qt.Recovering("QAudioDecoder::positionChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_PositionChanged(ptr.Pointer(), C.longlong(position))
	}
}

func (ptr *QAudioDecoder) Read() *QAudioBuffer {
	defer qt.Recovering("QAudioDecoder::read")

	if ptr.Pointer() != nil {
		return NewQAudioBufferFromPointer(C.QAudioDecoder_Read(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAudioDecoder) SetAudioFormat(format QAudioFormat_ITF) {
	defer qt.Recovering("QAudioDecoder::setAudioFormat")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_SetAudioFormat(ptr.Pointer(), PointerFromQAudioFormat(format))
	}
}

func (ptr *QAudioDecoder) SetSourceDevice(device core.QIODevice_ITF) {
	defer qt.Recovering("QAudioDecoder::setSourceDevice")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_SetSourceDevice(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

func (ptr *QAudioDecoder) SetSourceFilename(fileName string) {
	defer qt.Recovering("QAudioDecoder::setSourceFilename")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_SetSourceFilename(ptr.Pointer(), C.CString(fileName))
	}
}

func (ptr *QAudioDecoder) ConnectSourceChanged(f func()) {
	defer qt.Recovering("connect QAudioDecoder::sourceChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_ConnectSourceChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sourceChanged", f)
	}
}

func (ptr *QAudioDecoder) DisconnectSourceChanged() {
	defer qt.Recovering("disconnect QAudioDecoder::sourceChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_DisconnectSourceChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sourceChanged")
	}
}

//export callbackQAudioDecoderSourceChanged
func callbackQAudioDecoderSourceChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAudioDecoder::sourceChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "sourceChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAudioDecoder) SourceChanged() {
	defer qt.Recovering("QAudioDecoder::sourceChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_SourceChanged(ptr.Pointer())
	}
}

func (ptr *QAudioDecoder) SourceDevice() *core.QIODevice {
	defer qt.Recovering("QAudioDecoder::sourceDevice")

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QAudioDecoder_SourceDevice(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAudioDecoder) SourceFilename() string {
	defer qt.Recovering("QAudioDecoder::sourceFilename")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioDecoder_SourceFilename(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioDecoder) Start() {
	defer qt.Recovering("QAudioDecoder::start")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_Start(ptr.Pointer())
	}
}

func (ptr *QAudioDecoder) ConnectStateChanged(f func(state QAudioDecoder__State)) {
	defer qt.Recovering("connect QAudioDecoder::stateChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QAudioDecoder) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QAudioDecoder::stateChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQAudioDecoderStateChanged
func callbackQAudioDecoderStateChanged(ptr unsafe.Pointer, ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QAudioDecoder::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QAudioDecoder__State))(QAudioDecoder__State(state))
	}

}

func (ptr *QAudioDecoder) StateChanged(state QAudioDecoder__State) {
	defer qt.Recovering("QAudioDecoder::stateChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_StateChanged(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QAudioDecoder) Stop() {
	defer qt.Recovering("QAudioDecoder::stop")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_Stop(ptr.Pointer())
	}
}

func (ptr *QAudioDecoder) DestroyQAudioDecoder() {
	defer qt.Recovering("QAudioDecoder::~QAudioDecoder")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_DestroyQAudioDecoder(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAudioDecoder) ConnectUnbind(f func(object *core.QObject)) {
	defer qt.Recovering("connect QAudioDecoder::unbind")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "unbind", f)
	}
}

func (ptr *QAudioDecoder) DisconnectUnbind() {
	defer qt.Recovering("disconnect QAudioDecoder::unbind")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "unbind")
	}
}

//export callbackQAudioDecoderUnbind
func callbackQAudioDecoderUnbind(ptr unsafe.Pointer, ptrName *C.char, object unsafe.Pointer) {
	defer qt.Recovering("callback QAudioDecoder::unbind")

	if signal := qt.GetSignal(C.GoString(ptrName), "unbind"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(object))
	} else {
		NewQAudioDecoderFromPointer(ptr).UnbindDefault(core.NewQObjectFromPointer(object))
	}
}

func (ptr *QAudioDecoder) Unbind(object core.QObject_ITF) {
	defer qt.Recovering("QAudioDecoder::unbind")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_Unbind(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QAudioDecoder) UnbindDefault(object core.QObject_ITF) {
	defer qt.Recovering("QAudioDecoder::unbind")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_UnbindDefault(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QAudioDecoder) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAudioDecoder::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAudioDecoder) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAudioDecoder::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAudioDecoderTimerEvent
func callbackQAudioDecoderTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioDecoder::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAudioDecoderFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAudioDecoder) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAudioDecoder::timerEvent")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAudioDecoder) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAudioDecoder::timerEvent")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAudioDecoder) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAudioDecoder::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAudioDecoder) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAudioDecoder::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAudioDecoderChildEvent
func callbackQAudioDecoderChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioDecoder::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAudioDecoderFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAudioDecoder) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAudioDecoder::childEvent")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAudioDecoder) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAudioDecoder::childEvent")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAudioDecoder) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAudioDecoder::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAudioDecoder) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAudioDecoder::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAudioDecoderCustomEvent
func callbackQAudioDecoderCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioDecoder::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAudioDecoderFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAudioDecoder) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAudioDecoder::customEvent")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAudioDecoder) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAudioDecoder::customEvent")

	if ptr.Pointer() != nil {
		C.QAudioDecoder_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QAudioDecoderControl struct {
	QMediaControl
}

type QAudioDecoderControl_ITF interface {
	QMediaControl_ITF
	QAudioDecoderControl_PTR() *QAudioDecoderControl
}

func PointerFromQAudioDecoderControl(ptr QAudioDecoderControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioDecoderControl_PTR().Pointer()
	}
	return nil
}

func NewQAudioDecoderControlFromPointer(ptr unsafe.Pointer) *QAudioDecoderControl {
	var n = new(QAudioDecoderControl)
	n.SetPointer(ptr)
	return n
}

func newQAudioDecoderControlFromPointer(ptr unsafe.Pointer) *QAudioDecoderControl {
	var n = NewQAudioDecoderControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QAudioDecoderControl_") {
		n.SetObjectName("QAudioDecoderControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QAudioDecoderControl) QAudioDecoderControl_PTR() *QAudioDecoderControl {
	return ptr
}

func (ptr *QAudioDecoderControl) AudioFormat() *QAudioFormat {
	defer qt.Recovering("QAudioDecoderControl::audioFormat")

	if ptr.Pointer() != nil {
		return NewQAudioFormatFromPointer(C.QAudioDecoderControl_AudioFormat(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAudioDecoderControl) BufferAvailable() bool {
	defer qt.Recovering("QAudioDecoderControl::bufferAvailable")

	if ptr.Pointer() != nil {
		return C.QAudioDecoderControl_BufferAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAudioDecoderControl) ConnectBufferAvailableChanged(f func(available bool)) {
	defer qt.Recovering("connect QAudioDecoderControl::bufferAvailableChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectBufferAvailableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bufferAvailableChanged", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectBufferAvailableChanged() {
	defer qt.Recovering("disconnect QAudioDecoderControl::bufferAvailableChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectBufferAvailableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bufferAvailableChanged")
	}
}

//export callbackQAudioDecoderControlBufferAvailableChanged
func callbackQAudioDecoderControlBufferAvailableChanged(ptr unsafe.Pointer, ptrName *C.char, available C.int) {
	defer qt.Recovering("callback QAudioDecoderControl::bufferAvailableChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "bufferAvailableChanged"); signal != nil {
		signal.(func(bool))(int(available) != 0)
	}

}

func (ptr *QAudioDecoderControl) BufferAvailableChanged(available bool) {
	defer qt.Recovering("QAudioDecoderControl::bufferAvailableChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_BufferAvailableChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(available)))
	}
}

func (ptr *QAudioDecoderControl) ConnectBufferReady(f func()) {
	defer qt.Recovering("connect QAudioDecoderControl::bufferReady")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectBufferReady(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bufferReady", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectBufferReady() {
	defer qt.Recovering("disconnect QAudioDecoderControl::bufferReady")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectBufferReady(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bufferReady")
	}
}

//export callbackQAudioDecoderControlBufferReady
func callbackQAudioDecoderControlBufferReady(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAudioDecoderControl::bufferReady")

	if signal := qt.GetSignal(C.GoString(ptrName), "bufferReady"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAudioDecoderControl) BufferReady() {
	defer qt.Recovering("QAudioDecoderControl::bufferReady")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_BufferReady(ptr.Pointer())
	}
}

func (ptr *QAudioDecoderControl) Duration() int64 {
	defer qt.Recovering("QAudioDecoderControl::duration")

	if ptr.Pointer() != nil {
		return int64(C.QAudioDecoderControl_Duration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioDecoderControl) ConnectDurationChanged(f func(duration int64)) {
	defer qt.Recovering("connect QAudioDecoderControl::durationChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectDurationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "durationChanged", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectDurationChanged() {
	defer qt.Recovering("disconnect QAudioDecoderControl::durationChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectDurationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "durationChanged")
	}
}

//export callbackQAudioDecoderControlDurationChanged
func callbackQAudioDecoderControlDurationChanged(ptr unsafe.Pointer, ptrName *C.char, duration C.longlong) {
	defer qt.Recovering("callback QAudioDecoderControl::durationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "durationChanged"); signal != nil {
		signal.(func(int64))(int64(duration))
	}

}

func (ptr *QAudioDecoderControl) DurationChanged(duration int64) {
	defer qt.Recovering("QAudioDecoderControl::durationChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DurationChanged(ptr.Pointer(), C.longlong(duration))
	}
}

func (ptr *QAudioDecoderControl) ConnectError(f func(error int, errorString string)) {
	defer qt.Recovering("connect QAudioDecoderControl::error")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectError() {
	defer qt.Recovering("disconnect QAudioDecoderControl::error")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQAudioDecoderControlError
func callbackQAudioDecoderControlError(ptr unsafe.Pointer, ptrName *C.char, error C.int, errorString *C.char) {
	defer qt.Recovering("callback QAudioDecoderControl::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error"); signal != nil {
		signal.(func(int, string))(int(error), C.GoString(errorString))
	}

}

func (ptr *QAudioDecoderControl) Error(error int, errorString string) {
	defer qt.Recovering("QAudioDecoderControl::error")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_Error(ptr.Pointer(), C.int(error), C.CString(errorString))
	}
}

func (ptr *QAudioDecoderControl) ConnectFinished(f func()) {
	defer qt.Recovering("connect QAudioDecoderControl::finished")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectFinished() {
	defer qt.Recovering("disconnect QAudioDecoderControl::finished")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQAudioDecoderControlFinished
func callbackQAudioDecoderControlFinished(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAudioDecoderControl::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAudioDecoderControl) Finished() {
	defer qt.Recovering("QAudioDecoderControl::finished")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_Finished(ptr.Pointer())
	}
}

func (ptr *QAudioDecoderControl) ConnectFormatChanged(f func(format *QAudioFormat)) {
	defer qt.Recovering("connect QAudioDecoderControl::formatChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectFormatChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "formatChanged", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectFormatChanged() {
	defer qt.Recovering("disconnect QAudioDecoderControl::formatChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectFormatChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "formatChanged")
	}
}

//export callbackQAudioDecoderControlFormatChanged
func callbackQAudioDecoderControlFormatChanged(ptr unsafe.Pointer, ptrName *C.char, format unsafe.Pointer) {
	defer qt.Recovering("callback QAudioDecoderControl::formatChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "formatChanged"); signal != nil {
		signal.(func(*QAudioFormat))(NewQAudioFormatFromPointer(format))
	}

}

func (ptr *QAudioDecoderControl) FormatChanged(format QAudioFormat_ITF) {
	defer qt.Recovering("QAudioDecoderControl::formatChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_FormatChanged(ptr.Pointer(), PointerFromQAudioFormat(format))
	}
}

func (ptr *QAudioDecoderControl) Position() int64 {
	defer qt.Recovering("QAudioDecoderControl::position")

	if ptr.Pointer() != nil {
		return int64(C.QAudioDecoderControl_Position(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioDecoderControl) ConnectPositionChanged(f func(position int64)) {
	defer qt.Recovering("connect QAudioDecoderControl::positionChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectPositionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "positionChanged", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectPositionChanged() {
	defer qt.Recovering("disconnect QAudioDecoderControl::positionChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectPositionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "positionChanged")
	}
}

//export callbackQAudioDecoderControlPositionChanged
func callbackQAudioDecoderControlPositionChanged(ptr unsafe.Pointer, ptrName *C.char, position C.longlong) {
	defer qt.Recovering("callback QAudioDecoderControl::positionChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "positionChanged"); signal != nil {
		signal.(func(int64))(int64(position))
	}

}

func (ptr *QAudioDecoderControl) PositionChanged(position int64) {
	defer qt.Recovering("QAudioDecoderControl::positionChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_PositionChanged(ptr.Pointer(), C.longlong(position))
	}
}

func (ptr *QAudioDecoderControl) Read() *QAudioBuffer {
	defer qt.Recovering("QAudioDecoderControl::read")

	if ptr.Pointer() != nil {
		return NewQAudioBufferFromPointer(C.QAudioDecoderControl_Read(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAudioDecoderControl) SetAudioFormat(format QAudioFormat_ITF) {
	defer qt.Recovering("QAudioDecoderControl::setAudioFormat")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_SetAudioFormat(ptr.Pointer(), PointerFromQAudioFormat(format))
	}
}

func (ptr *QAudioDecoderControl) SetSourceDevice(device core.QIODevice_ITF) {
	defer qt.Recovering("QAudioDecoderControl::setSourceDevice")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_SetSourceDevice(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

func (ptr *QAudioDecoderControl) SetSourceFilename(fileName string) {
	defer qt.Recovering("QAudioDecoderControl::setSourceFilename")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_SetSourceFilename(ptr.Pointer(), C.CString(fileName))
	}
}

func (ptr *QAudioDecoderControl) ConnectSourceChanged(f func()) {
	defer qt.Recovering("connect QAudioDecoderControl::sourceChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectSourceChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sourceChanged", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectSourceChanged() {
	defer qt.Recovering("disconnect QAudioDecoderControl::sourceChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectSourceChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sourceChanged")
	}
}

//export callbackQAudioDecoderControlSourceChanged
func callbackQAudioDecoderControlSourceChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAudioDecoderControl::sourceChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "sourceChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAudioDecoderControl) SourceChanged() {
	defer qt.Recovering("QAudioDecoderControl::sourceChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_SourceChanged(ptr.Pointer())
	}
}

func (ptr *QAudioDecoderControl) SourceDevice() *core.QIODevice {
	defer qt.Recovering("QAudioDecoderControl::sourceDevice")

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QAudioDecoderControl_SourceDevice(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAudioDecoderControl) SourceFilename() string {
	defer qt.Recovering("QAudioDecoderControl::sourceFilename")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioDecoderControl_SourceFilename(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioDecoderControl) Start() {
	defer qt.Recovering("QAudioDecoderControl::start")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_Start(ptr.Pointer())
	}
}

func (ptr *QAudioDecoderControl) State() QAudioDecoder__State {
	defer qt.Recovering("QAudioDecoderControl::state")

	if ptr.Pointer() != nil {
		return QAudioDecoder__State(C.QAudioDecoderControl_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioDecoderControl) ConnectStateChanged(f func(state QAudioDecoder__State)) {
	defer qt.Recovering("connect QAudioDecoderControl::stateChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QAudioDecoderControl::stateChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQAudioDecoderControlStateChanged
func callbackQAudioDecoderControlStateChanged(ptr unsafe.Pointer, ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QAudioDecoderControl::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QAudioDecoder__State))(QAudioDecoder__State(state))
	}

}

func (ptr *QAudioDecoderControl) StateChanged(state QAudioDecoder__State) {
	defer qt.Recovering("QAudioDecoderControl::stateChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_StateChanged(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QAudioDecoderControl) Stop() {
	defer qt.Recovering("QAudioDecoderControl::stop")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_Stop(ptr.Pointer())
	}
}

func (ptr *QAudioDecoderControl) DestroyQAudioDecoderControl() {
	defer qt.Recovering("QAudioDecoderControl::~QAudioDecoderControl")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DestroyQAudioDecoderControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAudioDecoderControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAudioDecoderControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAudioDecoderControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAudioDecoderControlTimerEvent
func callbackQAudioDecoderControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioDecoderControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAudioDecoderControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAudioDecoderControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAudioDecoderControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAudioDecoderControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAudioDecoderControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAudioDecoderControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAudioDecoderControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAudioDecoderControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAudioDecoderControlChildEvent
func callbackQAudioDecoderControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioDecoderControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAudioDecoderControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAudioDecoderControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAudioDecoderControl::childEvent")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAudioDecoderControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAudioDecoderControl::childEvent")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAudioDecoderControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAudioDecoderControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAudioDecoderControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAudioDecoderControlCustomEvent
func callbackQAudioDecoderControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioDecoderControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAudioDecoderControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAudioDecoderControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAudioDecoderControl::customEvent")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAudioDecoderControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAudioDecoderControl::customEvent")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QAudioDeviceInfo struct {
	ptr unsafe.Pointer
}

type QAudioDeviceInfo_ITF interface {
	QAudioDeviceInfo_PTR() *QAudioDeviceInfo
}

func (p *QAudioDeviceInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAudioDeviceInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAudioDeviceInfo(ptr QAudioDeviceInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioDeviceInfo_PTR().Pointer()
	}
	return nil
}

func NewQAudioDeviceInfoFromPointer(ptr unsafe.Pointer) *QAudioDeviceInfo {
	var n = new(QAudioDeviceInfo)
	n.SetPointer(ptr)
	return n
}

func newQAudioDeviceInfoFromPointer(ptr unsafe.Pointer) *QAudioDeviceInfo {
	var n = NewQAudioDeviceInfoFromPointer(ptr)
	return n
}

func (ptr *QAudioDeviceInfo) QAudioDeviceInfo_PTR() *QAudioDeviceInfo {
	return ptr
}

func NewQAudioDeviceInfo() *QAudioDeviceInfo {
	defer qt.Recovering("QAudioDeviceInfo::QAudioDeviceInfo")

	return newQAudioDeviceInfoFromPointer(C.QAudioDeviceInfo_NewQAudioDeviceInfo())
}

func NewQAudioDeviceInfo2(other QAudioDeviceInfo_ITF) *QAudioDeviceInfo {
	defer qt.Recovering("QAudioDeviceInfo::QAudioDeviceInfo")

	return newQAudioDeviceInfoFromPointer(C.QAudioDeviceInfo_NewQAudioDeviceInfo2(PointerFromQAudioDeviceInfo(other)))
}

func QAudioDeviceInfo_DefaultInputDevice() *QAudioDeviceInfo {
	defer qt.Recovering("QAudioDeviceInfo::defaultInputDevice")

	return NewQAudioDeviceInfoFromPointer(C.QAudioDeviceInfo_QAudioDeviceInfo_DefaultInputDevice())
}

func QAudioDeviceInfo_DefaultOutputDevice() *QAudioDeviceInfo {
	defer qt.Recovering("QAudioDeviceInfo::defaultOutputDevice")

	return NewQAudioDeviceInfoFromPointer(C.QAudioDeviceInfo_QAudioDeviceInfo_DefaultOutputDevice())
}

func (ptr *QAudioDeviceInfo) DeviceName() string {
	defer qt.Recovering("QAudioDeviceInfo::deviceName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioDeviceInfo_DeviceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioDeviceInfo) IsFormatSupported(settings QAudioFormat_ITF) bool {
	defer qt.Recovering("QAudioDeviceInfo::isFormatSupported")

	if ptr.Pointer() != nil {
		return C.QAudioDeviceInfo_IsFormatSupported(ptr.Pointer(), PointerFromQAudioFormat(settings)) != 0
	}
	return false
}

func (ptr *QAudioDeviceInfo) IsNull() bool {
	defer qt.Recovering("QAudioDeviceInfo::isNull")

	if ptr.Pointer() != nil {
		return C.QAudioDeviceInfo_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAudioDeviceInfo) NearestFormat(settings QAudioFormat_ITF) *QAudioFormat {
	defer qt.Recovering("QAudioDeviceInfo::nearestFormat")

	if ptr.Pointer() != nil {
		return NewQAudioFormatFromPointer(C.QAudioDeviceInfo_NearestFormat(ptr.Pointer(), PointerFromQAudioFormat(settings)))
	}
	return nil
}

func (ptr *QAudioDeviceInfo) PreferredFormat() *QAudioFormat {
	defer qt.Recovering("QAudioDeviceInfo::preferredFormat")

	if ptr.Pointer() != nil {
		return NewQAudioFormatFromPointer(C.QAudioDeviceInfo_PreferredFormat(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAudioDeviceInfo) SupportedCodecs() []string {
	defer qt.Recovering("QAudioDeviceInfo::supportedCodecs")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QAudioDeviceInfo_SupportedCodecs(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QAudioDeviceInfo) DestroyQAudioDeviceInfo() {
	defer qt.Recovering("QAudioDeviceInfo::~QAudioDeviceInfo")

	if ptr.Pointer() != nil {
		C.QAudioDeviceInfo_DestroyQAudioDeviceInfo(ptr.Pointer())
	}
}

type QAudioEncoderSettings struct {
	ptr unsafe.Pointer
}

type QAudioEncoderSettings_ITF interface {
	QAudioEncoderSettings_PTR() *QAudioEncoderSettings
}

func (p *QAudioEncoderSettings) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAudioEncoderSettings) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAudioEncoderSettings(ptr QAudioEncoderSettings_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioEncoderSettings_PTR().Pointer()
	}
	return nil
}

func NewQAudioEncoderSettingsFromPointer(ptr unsafe.Pointer) *QAudioEncoderSettings {
	var n = new(QAudioEncoderSettings)
	n.SetPointer(ptr)
	return n
}

func newQAudioEncoderSettingsFromPointer(ptr unsafe.Pointer) *QAudioEncoderSettings {
	var n = NewQAudioEncoderSettingsFromPointer(ptr)
	return n
}

func (ptr *QAudioEncoderSettings) QAudioEncoderSettings_PTR() *QAudioEncoderSettings {
	return ptr
}

func NewQAudioEncoderSettings() *QAudioEncoderSettings {
	defer qt.Recovering("QAudioEncoderSettings::QAudioEncoderSettings")

	return newQAudioEncoderSettingsFromPointer(C.QAudioEncoderSettings_NewQAudioEncoderSettings())
}

func NewQAudioEncoderSettings2(other QAudioEncoderSettings_ITF) *QAudioEncoderSettings {
	defer qt.Recovering("QAudioEncoderSettings::QAudioEncoderSettings")

	return newQAudioEncoderSettingsFromPointer(C.QAudioEncoderSettings_NewQAudioEncoderSettings2(PointerFromQAudioEncoderSettings(other)))
}

func (ptr *QAudioEncoderSettings) BitRate() int {
	defer qt.Recovering("QAudioEncoderSettings::bitRate")

	if ptr.Pointer() != nil {
		return int(C.QAudioEncoderSettings_BitRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioEncoderSettings) ChannelCount() int {
	defer qt.Recovering("QAudioEncoderSettings::channelCount")

	if ptr.Pointer() != nil {
		return int(C.QAudioEncoderSettings_ChannelCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioEncoderSettings) Codec() string {
	defer qt.Recovering("QAudioEncoderSettings::codec")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioEncoderSettings_Codec(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioEncoderSettings) EncodingMode() QMultimedia__EncodingMode {
	defer qt.Recovering("QAudioEncoderSettings::encodingMode")

	if ptr.Pointer() != nil {
		return QMultimedia__EncodingMode(C.QAudioEncoderSettings_EncodingMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioEncoderSettings) EncodingOption(option string) *core.QVariant {
	defer qt.Recovering("QAudioEncoderSettings::encodingOption")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QAudioEncoderSettings_EncodingOption(ptr.Pointer(), C.CString(option)))
	}
	return nil
}

func (ptr *QAudioEncoderSettings) IsNull() bool {
	defer qt.Recovering("QAudioEncoderSettings::isNull")

	if ptr.Pointer() != nil {
		return C.QAudioEncoderSettings_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAudioEncoderSettings) Quality() QMultimedia__EncodingQuality {
	defer qt.Recovering("QAudioEncoderSettings::quality")

	if ptr.Pointer() != nil {
		return QMultimedia__EncodingQuality(C.QAudioEncoderSettings_Quality(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioEncoderSettings) SampleRate() int {
	defer qt.Recovering("QAudioEncoderSettings::sampleRate")

	if ptr.Pointer() != nil {
		return int(C.QAudioEncoderSettings_SampleRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioEncoderSettings) SetBitRate(rate int) {
	defer qt.Recovering("QAudioEncoderSettings::setBitRate")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettings_SetBitRate(ptr.Pointer(), C.int(rate))
	}
}

func (ptr *QAudioEncoderSettings) SetChannelCount(channels int) {
	defer qt.Recovering("QAudioEncoderSettings::setChannelCount")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettings_SetChannelCount(ptr.Pointer(), C.int(channels))
	}
}

func (ptr *QAudioEncoderSettings) SetCodec(codec string) {
	defer qt.Recovering("QAudioEncoderSettings::setCodec")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettings_SetCodec(ptr.Pointer(), C.CString(codec))
	}
}

func (ptr *QAudioEncoderSettings) SetEncodingMode(mode QMultimedia__EncodingMode) {
	defer qt.Recovering("QAudioEncoderSettings::setEncodingMode")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettings_SetEncodingMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QAudioEncoderSettings) SetEncodingOption(option string, value core.QVariant_ITF) {
	defer qt.Recovering("QAudioEncoderSettings::setEncodingOption")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettings_SetEncodingOption(ptr.Pointer(), C.CString(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QAudioEncoderSettings) SetQuality(quality QMultimedia__EncodingQuality) {
	defer qt.Recovering("QAudioEncoderSettings::setQuality")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettings_SetQuality(ptr.Pointer(), C.int(quality))
	}
}

func (ptr *QAudioEncoderSettings) SetSampleRate(rate int) {
	defer qt.Recovering("QAudioEncoderSettings::setSampleRate")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettings_SetSampleRate(ptr.Pointer(), C.int(rate))
	}
}

func (ptr *QAudioEncoderSettings) DestroyQAudioEncoderSettings() {
	defer qt.Recovering("QAudioEncoderSettings::~QAudioEncoderSettings")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettings_DestroyQAudioEncoderSettings(ptr.Pointer())
	}
}

type QAudioEncoderSettingsControl struct {
	QMediaControl
}

type QAudioEncoderSettingsControl_ITF interface {
	QMediaControl_ITF
	QAudioEncoderSettingsControl_PTR() *QAudioEncoderSettingsControl
}

func PointerFromQAudioEncoderSettingsControl(ptr QAudioEncoderSettingsControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioEncoderSettingsControl_PTR().Pointer()
	}
	return nil
}

func NewQAudioEncoderSettingsControlFromPointer(ptr unsafe.Pointer) *QAudioEncoderSettingsControl {
	var n = new(QAudioEncoderSettingsControl)
	n.SetPointer(ptr)
	return n
}

func newQAudioEncoderSettingsControlFromPointer(ptr unsafe.Pointer) *QAudioEncoderSettingsControl {
	var n = NewQAudioEncoderSettingsControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QAudioEncoderSettingsControl_") {
		n.SetObjectName("QAudioEncoderSettingsControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QAudioEncoderSettingsControl) QAudioEncoderSettingsControl_PTR() *QAudioEncoderSettingsControl {
	return ptr
}

func (ptr *QAudioEncoderSettingsControl) AudioSettings() *QAudioEncoderSettings {
	defer qt.Recovering("QAudioEncoderSettingsControl::audioSettings")

	if ptr.Pointer() != nil {
		return NewQAudioEncoderSettingsFromPointer(C.QAudioEncoderSettingsControl_AudioSettings(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAudioEncoderSettingsControl) CodecDescription(codec string) string {
	defer qt.Recovering("QAudioEncoderSettingsControl::codecDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioEncoderSettingsControl_CodecDescription(ptr.Pointer(), C.CString(codec)))
	}
	return ""
}

func (ptr *QAudioEncoderSettingsControl) SetAudioSettings(settings QAudioEncoderSettings_ITF) {
	defer qt.Recovering("QAudioEncoderSettingsControl::setAudioSettings")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettingsControl_SetAudioSettings(ptr.Pointer(), PointerFromQAudioEncoderSettings(settings))
	}
}

func (ptr *QAudioEncoderSettingsControl) SupportedAudioCodecs() []string {
	defer qt.Recovering("QAudioEncoderSettingsControl::supportedAudioCodecs")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QAudioEncoderSettingsControl_SupportedAudioCodecs(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QAudioEncoderSettingsControl) DestroyQAudioEncoderSettingsControl() {
	defer qt.Recovering("QAudioEncoderSettingsControl::~QAudioEncoderSettingsControl")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettingsControl_DestroyQAudioEncoderSettingsControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAudioEncoderSettingsControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAudioEncoderSettingsControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAudioEncoderSettingsControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAudioEncoderSettingsControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAudioEncoderSettingsControlTimerEvent
func callbackQAudioEncoderSettingsControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioEncoderSettingsControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAudioEncoderSettingsControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAudioEncoderSettingsControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAudioEncoderSettingsControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettingsControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAudioEncoderSettingsControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAudioEncoderSettingsControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettingsControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAudioEncoderSettingsControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAudioEncoderSettingsControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAudioEncoderSettingsControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAudioEncoderSettingsControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAudioEncoderSettingsControlChildEvent
func callbackQAudioEncoderSettingsControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioEncoderSettingsControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAudioEncoderSettingsControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAudioEncoderSettingsControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAudioEncoderSettingsControl::childEvent")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettingsControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAudioEncoderSettingsControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAudioEncoderSettingsControl::childEvent")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettingsControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAudioEncoderSettingsControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAudioEncoderSettingsControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAudioEncoderSettingsControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAudioEncoderSettingsControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAudioEncoderSettingsControlCustomEvent
func callbackQAudioEncoderSettingsControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioEncoderSettingsControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAudioEncoderSettingsControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAudioEncoderSettingsControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAudioEncoderSettingsControl::customEvent")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettingsControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAudioEncoderSettingsControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAudioEncoderSettingsControl::customEvent")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettingsControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QAudioFormat struct {
	ptr unsafe.Pointer
}

type QAudioFormat_ITF interface {
	QAudioFormat_PTR() *QAudioFormat
}

func (p *QAudioFormat) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAudioFormat) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAudioFormat(ptr QAudioFormat_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioFormat_PTR().Pointer()
	}
	return nil
}

func NewQAudioFormatFromPointer(ptr unsafe.Pointer) *QAudioFormat {
	var n = new(QAudioFormat)
	n.SetPointer(ptr)
	return n
}

func newQAudioFormatFromPointer(ptr unsafe.Pointer) *QAudioFormat {
	var n = NewQAudioFormatFromPointer(ptr)
	return n
}

func (ptr *QAudioFormat) QAudioFormat_PTR() *QAudioFormat {
	return ptr
}

//QAudioFormat::Endian
type QAudioFormat__Endian int64

const (
	QAudioFormat__BigEndian    = QAudioFormat__Endian(core.QSysInfo__BigEndian)
	QAudioFormat__LittleEndian = QAudioFormat__Endian(core.QSysInfo__LittleEndian)
)

//QAudioFormat::SampleType
type QAudioFormat__SampleType int64

const (
	QAudioFormat__Unknown     = QAudioFormat__SampleType(0)
	QAudioFormat__SignedInt   = QAudioFormat__SampleType(1)
	QAudioFormat__UnSignedInt = QAudioFormat__SampleType(2)
	QAudioFormat__Float       = QAudioFormat__SampleType(3)
)

func NewQAudioFormat() *QAudioFormat {
	defer qt.Recovering("QAudioFormat::QAudioFormat")

	return newQAudioFormatFromPointer(C.QAudioFormat_NewQAudioFormat())
}

func NewQAudioFormat2(other QAudioFormat_ITF) *QAudioFormat {
	defer qt.Recovering("QAudioFormat::QAudioFormat")

	return newQAudioFormatFromPointer(C.QAudioFormat_NewQAudioFormat2(PointerFromQAudioFormat(other)))
}

func (ptr *QAudioFormat) ByteOrder() QAudioFormat__Endian {
	defer qt.Recovering("QAudioFormat::byteOrder")

	if ptr.Pointer() != nil {
		return QAudioFormat__Endian(C.QAudioFormat_ByteOrder(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioFormat) BytesPerFrame() int {
	defer qt.Recovering("QAudioFormat::bytesPerFrame")

	if ptr.Pointer() != nil {
		return int(C.QAudioFormat_BytesPerFrame(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioFormat) ChannelCount() int {
	defer qt.Recovering("QAudioFormat::channelCount")

	if ptr.Pointer() != nil {
		return int(C.QAudioFormat_ChannelCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioFormat) Codec() string {
	defer qt.Recovering("QAudioFormat::codec")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioFormat_Codec(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioFormat) IsValid() bool {
	defer qt.Recovering("QAudioFormat::isValid")

	if ptr.Pointer() != nil {
		return C.QAudioFormat_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAudioFormat) SampleRate() int {
	defer qt.Recovering("QAudioFormat::sampleRate")

	if ptr.Pointer() != nil {
		return int(C.QAudioFormat_SampleRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioFormat) SampleSize() int {
	defer qt.Recovering("QAudioFormat::sampleSize")

	if ptr.Pointer() != nil {
		return int(C.QAudioFormat_SampleSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioFormat) SampleType() QAudioFormat__SampleType {
	defer qt.Recovering("QAudioFormat::sampleType")

	if ptr.Pointer() != nil {
		return QAudioFormat__SampleType(C.QAudioFormat_SampleType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioFormat) SetByteOrder(byteOrder QAudioFormat__Endian) {
	defer qt.Recovering("QAudioFormat::setByteOrder")

	if ptr.Pointer() != nil {
		C.QAudioFormat_SetByteOrder(ptr.Pointer(), C.int(byteOrder))
	}
}

func (ptr *QAudioFormat) SetChannelCount(channels int) {
	defer qt.Recovering("QAudioFormat::setChannelCount")

	if ptr.Pointer() != nil {
		C.QAudioFormat_SetChannelCount(ptr.Pointer(), C.int(channels))
	}
}

func (ptr *QAudioFormat) SetCodec(codec string) {
	defer qt.Recovering("QAudioFormat::setCodec")

	if ptr.Pointer() != nil {
		C.QAudioFormat_SetCodec(ptr.Pointer(), C.CString(codec))
	}
}

func (ptr *QAudioFormat) SetSampleRate(samplerate int) {
	defer qt.Recovering("QAudioFormat::setSampleRate")

	if ptr.Pointer() != nil {
		C.QAudioFormat_SetSampleRate(ptr.Pointer(), C.int(samplerate))
	}
}

func (ptr *QAudioFormat) SetSampleSize(sampleSize int) {
	defer qt.Recovering("QAudioFormat::setSampleSize")

	if ptr.Pointer() != nil {
		C.QAudioFormat_SetSampleSize(ptr.Pointer(), C.int(sampleSize))
	}
}

func (ptr *QAudioFormat) SetSampleType(sampleType QAudioFormat__SampleType) {
	defer qt.Recovering("QAudioFormat::setSampleType")

	if ptr.Pointer() != nil {
		C.QAudioFormat_SetSampleType(ptr.Pointer(), C.int(sampleType))
	}
}

func (ptr *QAudioFormat) DestroyQAudioFormat() {
	defer qt.Recovering("QAudioFormat::~QAudioFormat")

	if ptr.Pointer() != nil {
		C.QAudioFormat_DestroyQAudioFormat(ptr.Pointer())
	}
}

type QAudioInput struct {
	core.QObject
}

type QAudioInput_ITF interface {
	core.QObject_ITF
	QAudioInput_PTR() *QAudioInput
}

func PointerFromQAudioInput(ptr QAudioInput_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioInput_PTR().Pointer()
	}
	return nil
}

func NewQAudioInputFromPointer(ptr unsafe.Pointer) *QAudioInput {
	var n = new(QAudioInput)
	n.SetPointer(ptr)
	return n
}

func newQAudioInputFromPointer(ptr unsafe.Pointer) *QAudioInput {
	var n = NewQAudioInputFromPointer(ptr)
	for len(n.ObjectName()) < len("QAudioInput_") {
		n.SetObjectName("QAudioInput_" + qt.Identifier())
	}
	return n
}

func (ptr *QAudioInput) QAudioInput_PTR() *QAudioInput {
	return ptr
}

func NewQAudioInput2(audioDevice QAudioDeviceInfo_ITF, format QAudioFormat_ITF, parent core.QObject_ITF) *QAudioInput {
	defer qt.Recovering("QAudioInput::QAudioInput")

	return newQAudioInputFromPointer(C.QAudioInput_NewQAudioInput2(PointerFromQAudioDeviceInfo(audioDevice), PointerFromQAudioFormat(format), core.PointerFromQObject(parent)))
}

func NewQAudioInput(format QAudioFormat_ITF, parent core.QObject_ITF) *QAudioInput {
	defer qt.Recovering("QAudioInput::QAudioInput")

	return newQAudioInputFromPointer(C.QAudioInput_NewQAudioInput(PointerFromQAudioFormat(format), core.PointerFromQObject(parent)))
}

func (ptr *QAudioInput) BufferSize() int {
	defer qt.Recovering("QAudioInput::bufferSize")

	if ptr.Pointer() != nil {
		return int(C.QAudioInput_BufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioInput) BytesReady() int {
	defer qt.Recovering("QAudioInput::bytesReady")

	if ptr.Pointer() != nil {
		return int(C.QAudioInput_BytesReady(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioInput) ElapsedUSecs() int64 {
	defer qt.Recovering("QAudioInput::elapsedUSecs")

	if ptr.Pointer() != nil {
		return int64(C.QAudioInput_ElapsedUSecs(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioInput) Error() QAudio__Error {
	defer qt.Recovering("QAudioInput::error")

	if ptr.Pointer() != nil {
		return QAudio__Error(C.QAudioInput_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioInput) Format() *QAudioFormat {
	defer qt.Recovering("QAudioInput::format")

	if ptr.Pointer() != nil {
		return NewQAudioFormatFromPointer(C.QAudioInput_Format(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAudioInput) ConnectNotify(f func()) {
	defer qt.Recovering("connect QAudioInput::notify")

	if ptr.Pointer() != nil {
		C.QAudioInput_ConnectNotify(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "notify", f)
	}
}

func (ptr *QAudioInput) DisconnectNotify() {
	defer qt.Recovering("disconnect QAudioInput::notify")

	if ptr.Pointer() != nil {
		C.QAudioInput_DisconnectNotify(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "notify")
	}
}

//export callbackQAudioInputNotify
func callbackQAudioInputNotify(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAudioInput::notify")

	if signal := qt.GetSignal(C.GoString(ptrName), "notify"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAudioInput) Notify() {
	defer qt.Recovering("QAudioInput::notify")

	if ptr.Pointer() != nil {
		C.QAudioInput_Notify(ptr.Pointer())
	}
}

func (ptr *QAudioInput) NotifyInterval() int {
	defer qt.Recovering("QAudioInput::notifyInterval")

	if ptr.Pointer() != nil {
		return int(C.QAudioInput_NotifyInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioInput) PeriodSize() int {
	defer qt.Recovering("QAudioInput::periodSize")

	if ptr.Pointer() != nil {
		return int(C.QAudioInput_PeriodSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioInput) ProcessedUSecs() int64 {
	defer qt.Recovering("QAudioInput::processedUSecs")

	if ptr.Pointer() != nil {
		return int64(C.QAudioInput_ProcessedUSecs(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioInput) Reset() {
	defer qt.Recovering("QAudioInput::reset")

	if ptr.Pointer() != nil {
		C.QAudioInput_Reset(ptr.Pointer())
	}
}

func (ptr *QAudioInput) Resume() {
	defer qt.Recovering("QAudioInput::resume")

	if ptr.Pointer() != nil {
		C.QAudioInput_Resume(ptr.Pointer())
	}
}

func (ptr *QAudioInput) SetBufferSize(value int) {
	defer qt.Recovering("QAudioInput::setBufferSize")

	if ptr.Pointer() != nil {
		C.QAudioInput_SetBufferSize(ptr.Pointer(), C.int(value))
	}
}

func (ptr *QAudioInput) SetNotifyInterval(ms int) {
	defer qt.Recovering("QAudioInput::setNotifyInterval")

	if ptr.Pointer() != nil {
		C.QAudioInput_SetNotifyInterval(ptr.Pointer(), C.int(ms))
	}
}

func (ptr *QAudioInput) SetVolume(volume float64) {
	defer qt.Recovering("QAudioInput::setVolume")

	if ptr.Pointer() != nil {
		C.QAudioInput_SetVolume(ptr.Pointer(), C.double(volume))
	}
}

func (ptr *QAudioInput) Start2() *core.QIODevice {
	defer qt.Recovering("QAudioInput::start")

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QAudioInput_Start2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAudioInput) Start(device core.QIODevice_ITF) {
	defer qt.Recovering("QAudioInput::start")

	if ptr.Pointer() != nil {
		C.QAudioInput_Start(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

func (ptr *QAudioInput) State() QAudio__State {
	defer qt.Recovering("QAudioInput::state")

	if ptr.Pointer() != nil {
		return QAudio__State(C.QAudioInput_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioInput) ConnectStateChanged(f func(state QAudio__State)) {
	defer qt.Recovering("connect QAudioInput::stateChanged")

	if ptr.Pointer() != nil {
		C.QAudioInput_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QAudioInput) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QAudioInput::stateChanged")

	if ptr.Pointer() != nil {
		C.QAudioInput_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQAudioInputStateChanged
func callbackQAudioInputStateChanged(ptr unsafe.Pointer, ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QAudioInput::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QAudio__State))(QAudio__State(state))
	}

}

func (ptr *QAudioInput) StateChanged(state QAudio__State) {
	defer qt.Recovering("QAudioInput::stateChanged")

	if ptr.Pointer() != nil {
		C.QAudioInput_StateChanged(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QAudioInput) Stop() {
	defer qt.Recovering("QAudioInput::stop")

	if ptr.Pointer() != nil {
		C.QAudioInput_Stop(ptr.Pointer())
	}
}

func (ptr *QAudioInput) Suspend() {
	defer qt.Recovering("QAudioInput::suspend")

	if ptr.Pointer() != nil {
		C.QAudioInput_Suspend(ptr.Pointer())
	}
}

func (ptr *QAudioInput) Volume() float64 {
	defer qt.Recovering("QAudioInput::volume")

	if ptr.Pointer() != nil {
		return float64(C.QAudioInput_Volume(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioInput) DestroyQAudioInput() {
	defer qt.Recovering("QAudioInput::~QAudioInput")

	if ptr.Pointer() != nil {
		C.QAudioInput_DestroyQAudioInput(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAudioInput) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAudioInput::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAudioInput) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAudioInput::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAudioInputTimerEvent
func callbackQAudioInputTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioInput::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAudioInputFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAudioInput) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAudioInput::timerEvent")

	if ptr.Pointer() != nil {
		C.QAudioInput_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAudioInput) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAudioInput::timerEvent")

	if ptr.Pointer() != nil {
		C.QAudioInput_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAudioInput) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAudioInput::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAudioInput) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAudioInput::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAudioInputChildEvent
func callbackQAudioInputChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioInput::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAudioInputFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAudioInput) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAudioInput::childEvent")

	if ptr.Pointer() != nil {
		C.QAudioInput_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAudioInput) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAudioInput::childEvent")

	if ptr.Pointer() != nil {
		C.QAudioInput_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAudioInput) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAudioInput::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAudioInput) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAudioInput::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAudioInputCustomEvent
func callbackQAudioInputCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioInput::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAudioInputFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAudioInput) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAudioInput::customEvent")

	if ptr.Pointer() != nil {
		C.QAudioInput_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAudioInput) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAudioInput::customEvent")

	if ptr.Pointer() != nil {
		C.QAudioInput_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QAudioInputSelectorControl struct {
	QMediaControl
}

type QAudioInputSelectorControl_ITF interface {
	QMediaControl_ITF
	QAudioInputSelectorControl_PTR() *QAudioInputSelectorControl
}

func PointerFromQAudioInputSelectorControl(ptr QAudioInputSelectorControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioInputSelectorControl_PTR().Pointer()
	}
	return nil
}

func NewQAudioInputSelectorControlFromPointer(ptr unsafe.Pointer) *QAudioInputSelectorControl {
	var n = new(QAudioInputSelectorControl)
	n.SetPointer(ptr)
	return n
}

func newQAudioInputSelectorControlFromPointer(ptr unsafe.Pointer) *QAudioInputSelectorControl {
	var n = NewQAudioInputSelectorControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QAudioInputSelectorControl_") {
		n.SetObjectName("QAudioInputSelectorControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QAudioInputSelectorControl) QAudioInputSelectorControl_PTR() *QAudioInputSelectorControl {
	return ptr
}

func (ptr *QAudioInputSelectorControl) ActiveInput() string {
	defer qt.Recovering("QAudioInputSelectorControl::activeInput")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioInputSelectorControl_ActiveInput(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioInputSelectorControl) ConnectActiveInputChanged(f func(name string)) {
	defer qt.Recovering("connect QAudioInputSelectorControl::activeInputChanged")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_ConnectActiveInputChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeInputChanged", f)
	}
}

func (ptr *QAudioInputSelectorControl) DisconnectActiveInputChanged() {
	defer qt.Recovering("disconnect QAudioInputSelectorControl::activeInputChanged")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_DisconnectActiveInputChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeInputChanged")
	}
}

//export callbackQAudioInputSelectorControlActiveInputChanged
func callbackQAudioInputSelectorControlActiveInputChanged(ptr unsafe.Pointer, ptrName *C.char, name *C.char) {
	defer qt.Recovering("callback QAudioInputSelectorControl::activeInputChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "activeInputChanged"); signal != nil {
		signal.(func(string))(C.GoString(name))
	}

}

func (ptr *QAudioInputSelectorControl) ActiveInputChanged(name string) {
	defer qt.Recovering("QAudioInputSelectorControl::activeInputChanged")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_ActiveInputChanged(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QAudioInputSelectorControl) ConnectAvailableInputsChanged(f func()) {
	defer qt.Recovering("connect QAudioInputSelectorControl::availableInputsChanged")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_ConnectAvailableInputsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "availableInputsChanged", f)
	}
}

func (ptr *QAudioInputSelectorControl) DisconnectAvailableInputsChanged() {
	defer qt.Recovering("disconnect QAudioInputSelectorControl::availableInputsChanged")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_DisconnectAvailableInputsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "availableInputsChanged")
	}
}

//export callbackQAudioInputSelectorControlAvailableInputsChanged
func callbackQAudioInputSelectorControlAvailableInputsChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAudioInputSelectorControl::availableInputsChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "availableInputsChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAudioInputSelectorControl) AvailableInputsChanged() {
	defer qt.Recovering("QAudioInputSelectorControl::availableInputsChanged")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_AvailableInputsChanged(ptr.Pointer())
	}
}

func (ptr *QAudioInputSelectorControl) DefaultInput() string {
	defer qt.Recovering("QAudioInputSelectorControl::defaultInput")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioInputSelectorControl_DefaultInput(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioInputSelectorControl) InputDescription(name string) string {
	defer qt.Recovering("QAudioInputSelectorControl::inputDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioInputSelectorControl_InputDescription(ptr.Pointer(), C.CString(name)))
	}
	return ""
}

func (ptr *QAudioInputSelectorControl) SetActiveInput(name string) {
	defer qt.Recovering("QAudioInputSelectorControl::setActiveInput")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_SetActiveInput(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QAudioInputSelectorControl) DestroyQAudioInputSelectorControl() {
	defer qt.Recovering("QAudioInputSelectorControl::~QAudioInputSelectorControl")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_DestroyQAudioInputSelectorControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAudioInputSelectorControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAudioInputSelectorControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAudioInputSelectorControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAudioInputSelectorControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAudioInputSelectorControlTimerEvent
func callbackQAudioInputSelectorControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioInputSelectorControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAudioInputSelectorControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAudioInputSelectorControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAudioInputSelectorControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAudioInputSelectorControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAudioInputSelectorControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAudioInputSelectorControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAudioInputSelectorControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAudioInputSelectorControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAudioInputSelectorControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAudioInputSelectorControlChildEvent
func callbackQAudioInputSelectorControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioInputSelectorControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAudioInputSelectorControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAudioInputSelectorControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAudioInputSelectorControl::childEvent")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAudioInputSelectorControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAudioInputSelectorControl::childEvent")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAudioInputSelectorControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAudioInputSelectorControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAudioInputSelectorControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAudioInputSelectorControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAudioInputSelectorControlCustomEvent
func callbackQAudioInputSelectorControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioInputSelectorControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAudioInputSelectorControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAudioInputSelectorControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAudioInputSelectorControl::customEvent")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAudioInputSelectorControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAudioInputSelectorControl::customEvent")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QAudioOutput struct {
	core.QObject
}

type QAudioOutput_ITF interface {
	core.QObject_ITF
	QAudioOutput_PTR() *QAudioOutput
}

func PointerFromQAudioOutput(ptr QAudioOutput_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioOutput_PTR().Pointer()
	}
	return nil
}

func NewQAudioOutputFromPointer(ptr unsafe.Pointer) *QAudioOutput {
	var n = new(QAudioOutput)
	n.SetPointer(ptr)
	return n
}

func newQAudioOutputFromPointer(ptr unsafe.Pointer) *QAudioOutput {
	var n = NewQAudioOutputFromPointer(ptr)
	for len(n.ObjectName()) < len("QAudioOutput_") {
		n.SetObjectName("QAudioOutput_" + qt.Identifier())
	}
	return n
}

func (ptr *QAudioOutput) QAudioOutput_PTR() *QAudioOutput {
	return ptr
}

func NewQAudioOutput2(audioDevice QAudioDeviceInfo_ITF, format QAudioFormat_ITF, parent core.QObject_ITF) *QAudioOutput {
	defer qt.Recovering("QAudioOutput::QAudioOutput")

	return newQAudioOutputFromPointer(C.QAudioOutput_NewQAudioOutput2(PointerFromQAudioDeviceInfo(audioDevice), PointerFromQAudioFormat(format), core.PointerFromQObject(parent)))
}

func NewQAudioOutput(format QAudioFormat_ITF, parent core.QObject_ITF) *QAudioOutput {
	defer qt.Recovering("QAudioOutput::QAudioOutput")

	return newQAudioOutputFromPointer(C.QAudioOutput_NewQAudioOutput(PointerFromQAudioFormat(format), core.PointerFromQObject(parent)))
}

func (ptr *QAudioOutput) BufferSize() int {
	defer qt.Recovering("QAudioOutput::bufferSize")

	if ptr.Pointer() != nil {
		return int(C.QAudioOutput_BufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioOutput) BytesFree() int {
	defer qt.Recovering("QAudioOutput::bytesFree")

	if ptr.Pointer() != nil {
		return int(C.QAudioOutput_BytesFree(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioOutput) Category() string {
	defer qt.Recovering("QAudioOutput::category")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioOutput_Category(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioOutput) ElapsedUSecs() int64 {
	defer qt.Recovering("QAudioOutput::elapsedUSecs")

	if ptr.Pointer() != nil {
		return int64(C.QAudioOutput_ElapsedUSecs(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioOutput) Error() QAudio__Error {
	defer qt.Recovering("QAudioOutput::error")

	if ptr.Pointer() != nil {
		return QAudio__Error(C.QAudioOutput_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioOutput) Format() *QAudioFormat {
	defer qt.Recovering("QAudioOutput::format")

	if ptr.Pointer() != nil {
		return NewQAudioFormatFromPointer(C.QAudioOutput_Format(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAudioOutput) ConnectNotify(f func()) {
	defer qt.Recovering("connect QAudioOutput::notify")

	if ptr.Pointer() != nil {
		C.QAudioOutput_ConnectNotify(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "notify", f)
	}
}

func (ptr *QAudioOutput) DisconnectNotify() {
	defer qt.Recovering("disconnect QAudioOutput::notify")

	if ptr.Pointer() != nil {
		C.QAudioOutput_DisconnectNotify(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "notify")
	}
}

//export callbackQAudioOutputNotify
func callbackQAudioOutputNotify(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAudioOutput::notify")

	if signal := qt.GetSignal(C.GoString(ptrName), "notify"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAudioOutput) Notify() {
	defer qt.Recovering("QAudioOutput::notify")

	if ptr.Pointer() != nil {
		C.QAudioOutput_Notify(ptr.Pointer())
	}
}

func (ptr *QAudioOutput) NotifyInterval() int {
	defer qt.Recovering("QAudioOutput::notifyInterval")

	if ptr.Pointer() != nil {
		return int(C.QAudioOutput_NotifyInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioOutput) PeriodSize() int {
	defer qt.Recovering("QAudioOutput::periodSize")

	if ptr.Pointer() != nil {
		return int(C.QAudioOutput_PeriodSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioOutput) ProcessedUSecs() int64 {
	defer qt.Recovering("QAudioOutput::processedUSecs")

	if ptr.Pointer() != nil {
		return int64(C.QAudioOutput_ProcessedUSecs(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioOutput) Reset() {
	defer qt.Recovering("QAudioOutput::reset")

	if ptr.Pointer() != nil {
		C.QAudioOutput_Reset(ptr.Pointer())
	}
}

func (ptr *QAudioOutput) Resume() {
	defer qt.Recovering("QAudioOutput::resume")

	if ptr.Pointer() != nil {
		C.QAudioOutput_Resume(ptr.Pointer())
	}
}

func (ptr *QAudioOutput) SetBufferSize(value int) {
	defer qt.Recovering("QAudioOutput::setBufferSize")

	if ptr.Pointer() != nil {
		C.QAudioOutput_SetBufferSize(ptr.Pointer(), C.int(value))
	}
}

func (ptr *QAudioOutput) SetCategory(category string) {
	defer qt.Recovering("QAudioOutput::setCategory")

	if ptr.Pointer() != nil {
		C.QAudioOutput_SetCategory(ptr.Pointer(), C.CString(category))
	}
}

func (ptr *QAudioOutput) SetNotifyInterval(ms int) {
	defer qt.Recovering("QAudioOutput::setNotifyInterval")

	if ptr.Pointer() != nil {
		C.QAudioOutput_SetNotifyInterval(ptr.Pointer(), C.int(ms))
	}
}

func (ptr *QAudioOutput) SetVolume(volume float64) {
	defer qt.Recovering("QAudioOutput::setVolume")

	if ptr.Pointer() != nil {
		C.QAudioOutput_SetVolume(ptr.Pointer(), C.double(volume))
	}
}

func (ptr *QAudioOutput) Start2() *core.QIODevice {
	defer qt.Recovering("QAudioOutput::start")

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QAudioOutput_Start2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAudioOutput) Start(device core.QIODevice_ITF) {
	defer qt.Recovering("QAudioOutput::start")

	if ptr.Pointer() != nil {
		C.QAudioOutput_Start(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

func (ptr *QAudioOutput) State() QAudio__State {
	defer qt.Recovering("QAudioOutput::state")

	if ptr.Pointer() != nil {
		return QAudio__State(C.QAudioOutput_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioOutput) ConnectStateChanged(f func(state QAudio__State)) {
	defer qt.Recovering("connect QAudioOutput::stateChanged")

	if ptr.Pointer() != nil {
		C.QAudioOutput_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QAudioOutput) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QAudioOutput::stateChanged")

	if ptr.Pointer() != nil {
		C.QAudioOutput_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQAudioOutputStateChanged
func callbackQAudioOutputStateChanged(ptr unsafe.Pointer, ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QAudioOutput::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QAudio__State))(QAudio__State(state))
	}

}

func (ptr *QAudioOutput) StateChanged(state QAudio__State) {
	defer qt.Recovering("QAudioOutput::stateChanged")

	if ptr.Pointer() != nil {
		C.QAudioOutput_StateChanged(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QAudioOutput) Stop() {
	defer qt.Recovering("QAudioOutput::stop")

	if ptr.Pointer() != nil {
		C.QAudioOutput_Stop(ptr.Pointer())
	}
}

func (ptr *QAudioOutput) Suspend() {
	defer qt.Recovering("QAudioOutput::suspend")

	if ptr.Pointer() != nil {
		C.QAudioOutput_Suspend(ptr.Pointer())
	}
}

func (ptr *QAudioOutput) Volume() float64 {
	defer qt.Recovering("QAudioOutput::volume")

	if ptr.Pointer() != nil {
		return float64(C.QAudioOutput_Volume(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioOutput) DestroyQAudioOutput() {
	defer qt.Recovering("QAudioOutput::~QAudioOutput")

	if ptr.Pointer() != nil {
		C.QAudioOutput_DestroyQAudioOutput(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAudioOutput) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAudioOutput::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAudioOutput) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAudioOutput::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAudioOutputTimerEvent
func callbackQAudioOutputTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioOutput::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAudioOutputFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAudioOutput) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAudioOutput::timerEvent")

	if ptr.Pointer() != nil {
		C.QAudioOutput_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAudioOutput) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAudioOutput::timerEvent")

	if ptr.Pointer() != nil {
		C.QAudioOutput_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAudioOutput) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAudioOutput::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAudioOutput) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAudioOutput::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAudioOutputChildEvent
func callbackQAudioOutputChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioOutput::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAudioOutputFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAudioOutput) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAudioOutput::childEvent")

	if ptr.Pointer() != nil {
		C.QAudioOutput_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAudioOutput) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAudioOutput::childEvent")

	if ptr.Pointer() != nil {
		C.QAudioOutput_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAudioOutput) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAudioOutput::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAudioOutput) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAudioOutput::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAudioOutputCustomEvent
func callbackQAudioOutputCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioOutput::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAudioOutputFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAudioOutput) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAudioOutput::customEvent")

	if ptr.Pointer() != nil {
		C.QAudioOutput_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAudioOutput) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAudioOutput::customEvent")

	if ptr.Pointer() != nil {
		C.QAudioOutput_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QAudioOutputSelectorControl struct {
	QMediaControl
}

type QAudioOutputSelectorControl_ITF interface {
	QMediaControl_ITF
	QAudioOutputSelectorControl_PTR() *QAudioOutputSelectorControl
}

func PointerFromQAudioOutputSelectorControl(ptr QAudioOutputSelectorControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioOutputSelectorControl_PTR().Pointer()
	}
	return nil
}

func NewQAudioOutputSelectorControlFromPointer(ptr unsafe.Pointer) *QAudioOutputSelectorControl {
	var n = new(QAudioOutputSelectorControl)
	n.SetPointer(ptr)
	return n
}

func newQAudioOutputSelectorControlFromPointer(ptr unsafe.Pointer) *QAudioOutputSelectorControl {
	var n = NewQAudioOutputSelectorControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QAudioOutputSelectorControl_") {
		n.SetObjectName("QAudioOutputSelectorControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QAudioOutputSelectorControl) QAudioOutputSelectorControl_PTR() *QAudioOutputSelectorControl {
	return ptr
}

func (ptr *QAudioOutputSelectorControl) ActiveOutput() string {
	defer qt.Recovering("QAudioOutputSelectorControl::activeOutput")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioOutputSelectorControl_ActiveOutput(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioOutputSelectorControl) ConnectActiveOutputChanged(f func(name string)) {
	defer qt.Recovering("connect QAudioOutputSelectorControl::activeOutputChanged")

	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_ConnectActiveOutputChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeOutputChanged", f)
	}
}

func (ptr *QAudioOutputSelectorControl) DisconnectActiveOutputChanged() {
	defer qt.Recovering("disconnect QAudioOutputSelectorControl::activeOutputChanged")

	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_DisconnectActiveOutputChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeOutputChanged")
	}
}

//export callbackQAudioOutputSelectorControlActiveOutputChanged
func callbackQAudioOutputSelectorControlActiveOutputChanged(ptr unsafe.Pointer, ptrName *C.char, name *C.char) {
	defer qt.Recovering("callback QAudioOutputSelectorControl::activeOutputChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "activeOutputChanged"); signal != nil {
		signal.(func(string))(C.GoString(name))
	}

}

func (ptr *QAudioOutputSelectorControl) ActiveOutputChanged(name string) {
	defer qt.Recovering("QAudioOutputSelectorControl::activeOutputChanged")

	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_ActiveOutputChanged(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QAudioOutputSelectorControl) ConnectAvailableOutputsChanged(f func()) {
	defer qt.Recovering("connect QAudioOutputSelectorControl::availableOutputsChanged")

	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_ConnectAvailableOutputsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "availableOutputsChanged", f)
	}
}

func (ptr *QAudioOutputSelectorControl) DisconnectAvailableOutputsChanged() {
	defer qt.Recovering("disconnect QAudioOutputSelectorControl::availableOutputsChanged")

	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_DisconnectAvailableOutputsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "availableOutputsChanged")
	}
}

//export callbackQAudioOutputSelectorControlAvailableOutputsChanged
func callbackQAudioOutputSelectorControlAvailableOutputsChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAudioOutputSelectorControl::availableOutputsChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "availableOutputsChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAudioOutputSelectorControl) AvailableOutputsChanged() {
	defer qt.Recovering("QAudioOutputSelectorControl::availableOutputsChanged")

	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_AvailableOutputsChanged(ptr.Pointer())
	}
}

func (ptr *QAudioOutputSelectorControl) DefaultOutput() string {
	defer qt.Recovering("QAudioOutputSelectorControl::defaultOutput")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioOutputSelectorControl_DefaultOutput(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioOutputSelectorControl) OutputDescription(name string) string {
	defer qt.Recovering("QAudioOutputSelectorControl::outputDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioOutputSelectorControl_OutputDescription(ptr.Pointer(), C.CString(name)))
	}
	return ""
}

func (ptr *QAudioOutputSelectorControl) SetActiveOutput(name string) {
	defer qt.Recovering("QAudioOutputSelectorControl::setActiveOutput")

	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_SetActiveOutput(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QAudioOutputSelectorControl) DestroyQAudioOutputSelectorControl() {
	defer qt.Recovering("QAudioOutputSelectorControl::~QAudioOutputSelectorControl")

	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_DestroyQAudioOutputSelectorControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAudioOutputSelectorControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAudioOutputSelectorControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAudioOutputSelectorControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAudioOutputSelectorControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAudioOutputSelectorControlTimerEvent
func callbackQAudioOutputSelectorControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioOutputSelectorControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAudioOutputSelectorControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAudioOutputSelectorControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAudioOutputSelectorControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAudioOutputSelectorControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAudioOutputSelectorControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAudioOutputSelectorControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAudioOutputSelectorControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAudioOutputSelectorControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAudioOutputSelectorControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAudioOutputSelectorControlChildEvent
func callbackQAudioOutputSelectorControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioOutputSelectorControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAudioOutputSelectorControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAudioOutputSelectorControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAudioOutputSelectorControl::childEvent")

	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAudioOutputSelectorControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAudioOutputSelectorControl::childEvent")

	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAudioOutputSelectorControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAudioOutputSelectorControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAudioOutputSelectorControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAudioOutputSelectorControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAudioOutputSelectorControlCustomEvent
func callbackQAudioOutputSelectorControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioOutputSelectorControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAudioOutputSelectorControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAudioOutputSelectorControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAudioOutputSelectorControl::customEvent")

	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAudioOutputSelectorControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAudioOutputSelectorControl::customEvent")

	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QAudioProbe struct {
	core.QObject
}

type QAudioProbe_ITF interface {
	core.QObject_ITF
	QAudioProbe_PTR() *QAudioProbe
}

func PointerFromQAudioProbe(ptr QAudioProbe_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioProbe_PTR().Pointer()
	}
	return nil
}

func NewQAudioProbeFromPointer(ptr unsafe.Pointer) *QAudioProbe {
	var n = new(QAudioProbe)
	n.SetPointer(ptr)
	return n
}

func newQAudioProbeFromPointer(ptr unsafe.Pointer) *QAudioProbe {
	var n = NewQAudioProbeFromPointer(ptr)
	for len(n.ObjectName()) < len("QAudioProbe_") {
		n.SetObjectName("QAudioProbe_" + qt.Identifier())
	}
	return n
}

func (ptr *QAudioProbe) QAudioProbe_PTR() *QAudioProbe {
	return ptr
}

func NewQAudioProbe(parent core.QObject_ITF) *QAudioProbe {
	defer qt.Recovering("QAudioProbe::QAudioProbe")

	return newQAudioProbeFromPointer(C.QAudioProbe_NewQAudioProbe(core.PointerFromQObject(parent)))
}

func (ptr *QAudioProbe) ConnectAudioBufferProbed(f func(buffer *QAudioBuffer)) {
	defer qt.Recovering("connect QAudioProbe::audioBufferProbed")

	if ptr.Pointer() != nil {
		C.QAudioProbe_ConnectAudioBufferProbed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "audioBufferProbed", f)
	}
}

func (ptr *QAudioProbe) DisconnectAudioBufferProbed() {
	defer qt.Recovering("disconnect QAudioProbe::audioBufferProbed")

	if ptr.Pointer() != nil {
		C.QAudioProbe_DisconnectAudioBufferProbed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "audioBufferProbed")
	}
}

//export callbackQAudioProbeAudioBufferProbed
func callbackQAudioProbeAudioBufferProbed(ptr unsafe.Pointer, ptrName *C.char, buffer unsafe.Pointer) {
	defer qt.Recovering("callback QAudioProbe::audioBufferProbed")

	if signal := qt.GetSignal(C.GoString(ptrName), "audioBufferProbed"); signal != nil {
		signal.(func(*QAudioBuffer))(NewQAudioBufferFromPointer(buffer))
	}

}

func (ptr *QAudioProbe) AudioBufferProbed(buffer QAudioBuffer_ITF) {
	defer qt.Recovering("QAudioProbe::audioBufferProbed")

	if ptr.Pointer() != nil {
		C.QAudioProbe_AudioBufferProbed(ptr.Pointer(), PointerFromQAudioBuffer(buffer))
	}
}

func (ptr *QAudioProbe) ConnectFlush(f func()) {
	defer qt.Recovering("connect QAudioProbe::flush")

	if ptr.Pointer() != nil {
		C.QAudioProbe_ConnectFlush(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "flush", f)
	}
}

func (ptr *QAudioProbe) DisconnectFlush() {
	defer qt.Recovering("disconnect QAudioProbe::flush")

	if ptr.Pointer() != nil {
		C.QAudioProbe_DisconnectFlush(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "flush")
	}
}

//export callbackQAudioProbeFlush
func callbackQAudioProbeFlush(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAudioProbe::flush")

	if signal := qt.GetSignal(C.GoString(ptrName), "flush"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAudioProbe) Flush() {
	defer qt.Recovering("QAudioProbe::flush")

	if ptr.Pointer() != nil {
		C.QAudioProbe_Flush(ptr.Pointer())
	}
}

func (ptr *QAudioProbe) IsActive() bool {
	defer qt.Recovering("QAudioProbe::isActive")

	if ptr.Pointer() != nil {
		return C.QAudioProbe_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAudioProbe) SetSource(source QMediaObject_ITF) bool {
	defer qt.Recovering("QAudioProbe::setSource")

	if ptr.Pointer() != nil {
		return C.QAudioProbe_SetSource(ptr.Pointer(), PointerFromQMediaObject(source)) != 0
	}
	return false
}

func (ptr *QAudioProbe) SetSource2(mediaRecorder QMediaRecorder_ITF) bool {
	defer qt.Recovering("QAudioProbe::setSource")

	if ptr.Pointer() != nil {
		return C.QAudioProbe_SetSource2(ptr.Pointer(), PointerFromQMediaRecorder(mediaRecorder)) != 0
	}
	return false
}

func (ptr *QAudioProbe) DestroyQAudioProbe() {
	defer qt.Recovering("QAudioProbe::~QAudioProbe")

	if ptr.Pointer() != nil {
		C.QAudioProbe_DestroyQAudioProbe(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAudioProbe) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAudioProbe::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAudioProbe) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAudioProbe::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAudioProbeTimerEvent
func callbackQAudioProbeTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioProbe::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAudioProbeFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAudioProbe) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAudioProbe::timerEvent")

	if ptr.Pointer() != nil {
		C.QAudioProbe_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAudioProbe) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAudioProbe::timerEvent")

	if ptr.Pointer() != nil {
		C.QAudioProbe_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAudioProbe) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAudioProbe::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAudioProbe) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAudioProbe::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAudioProbeChildEvent
func callbackQAudioProbeChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioProbe::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAudioProbeFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAudioProbe) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAudioProbe::childEvent")

	if ptr.Pointer() != nil {
		C.QAudioProbe_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAudioProbe) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAudioProbe::childEvent")

	if ptr.Pointer() != nil {
		C.QAudioProbe_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAudioProbe) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAudioProbe::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAudioProbe) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAudioProbe::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAudioProbeCustomEvent
func callbackQAudioProbeCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioProbe::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAudioProbeFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAudioProbe) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAudioProbe::customEvent")

	if ptr.Pointer() != nil {
		C.QAudioProbe_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAudioProbe) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAudioProbe::customEvent")

	if ptr.Pointer() != nil {
		C.QAudioProbe_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QAudioRecorder struct {
	QMediaRecorder
}

type QAudioRecorder_ITF interface {
	QMediaRecorder_ITF
	QAudioRecorder_PTR() *QAudioRecorder
}

func PointerFromQAudioRecorder(ptr QAudioRecorder_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioRecorder_PTR().Pointer()
	}
	return nil
}

func NewQAudioRecorderFromPointer(ptr unsafe.Pointer) *QAudioRecorder {
	var n = new(QAudioRecorder)
	n.SetPointer(ptr)
	return n
}

func newQAudioRecorderFromPointer(ptr unsafe.Pointer) *QAudioRecorder {
	var n = NewQAudioRecorderFromPointer(ptr)
	for len(n.ObjectName()) < len("QAudioRecorder_") {
		n.SetObjectName("QAudioRecorder_" + qt.Identifier())
	}
	return n
}

func (ptr *QAudioRecorder) QAudioRecorder_PTR() *QAudioRecorder {
	return ptr
}

func NewQAudioRecorder(parent core.QObject_ITF) *QAudioRecorder {
	defer qt.Recovering("QAudioRecorder::QAudioRecorder")

	return newQAudioRecorderFromPointer(C.QAudioRecorder_NewQAudioRecorder(core.PointerFromQObject(parent)))
}

func (ptr *QAudioRecorder) AudioInput() string {
	defer qt.Recovering("QAudioRecorder::audioInput")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioRecorder_AudioInput(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioRecorder) ConnectAudioInputChanged(f func(name string)) {
	defer qt.Recovering("connect QAudioRecorder::audioInputChanged")

	if ptr.Pointer() != nil {
		C.QAudioRecorder_ConnectAudioInputChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "audioInputChanged", f)
	}
}

func (ptr *QAudioRecorder) DisconnectAudioInputChanged() {
	defer qt.Recovering("disconnect QAudioRecorder::audioInputChanged")

	if ptr.Pointer() != nil {
		C.QAudioRecorder_DisconnectAudioInputChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "audioInputChanged")
	}
}

//export callbackQAudioRecorderAudioInputChanged
func callbackQAudioRecorderAudioInputChanged(ptr unsafe.Pointer, ptrName *C.char, name *C.char) {
	defer qt.Recovering("callback QAudioRecorder::audioInputChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "audioInputChanged"); signal != nil {
		signal.(func(string))(C.GoString(name))
	}

}

func (ptr *QAudioRecorder) AudioInputChanged(name string) {
	defer qt.Recovering("QAudioRecorder::audioInputChanged")

	if ptr.Pointer() != nil {
		C.QAudioRecorder_AudioInputChanged(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QAudioRecorder) AudioInputDescription(name string) string {
	defer qt.Recovering("QAudioRecorder::audioInputDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioRecorder_AudioInputDescription(ptr.Pointer(), C.CString(name)))
	}
	return ""
}

func (ptr *QAudioRecorder) AudioInputs() []string {
	defer qt.Recovering("QAudioRecorder::audioInputs")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QAudioRecorder_AudioInputs(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QAudioRecorder) ConnectAvailableAudioInputsChanged(f func()) {
	defer qt.Recovering("connect QAudioRecorder::availableAudioInputsChanged")

	if ptr.Pointer() != nil {
		C.QAudioRecorder_ConnectAvailableAudioInputsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "availableAudioInputsChanged", f)
	}
}

func (ptr *QAudioRecorder) DisconnectAvailableAudioInputsChanged() {
	defer qt.Recovering("disconnect QAudioRecorder::availableAudioInputsChanged")

	if ptr.Pointer() != nil {
		C.QAudioRecorder_DisconnectAvailableAudioInputsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "availableAudioInputsChanged")
	}
}

//export callbackQAudioRecorderAvailableAudioInputsChanged
func callbackQAudioRecorderAvailableAudioInputsChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAudioRecorder::availableAudioInputsChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "availableAudioInputsChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAudioRecorder) AvailableAudioInputsChanged() {
	defer qt.Recovering("QAudioRecorder::availableAudioInputsChanged")

	if ptr.Pointer() != nil {
		C.QAudioRecorder_AvailableAudioInputsChanged(ptr.Pointer())
	}
}

func (ptr *QAudioRecorder) DefaultAudioInput() string {
	defer qt.Recovering("QAudioRecorder::defaultAudioInput")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioRecorder_DefaultAudioInput(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioRecorder) SetAudioInput(name string) {
	defer qt.Recovering("QAudioRecorder::setAudioInput")

	if ptr.Pointer() != nil {
		C.QAudioRecorder_SetAudioInput(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QAudioRecorder) DestroyQAudioRecorder() {
	defer qt.Recovering("QAudioRecorder::~QAudioRecorder")

	if ptr.Pointer() != nil {
		C.QAudioRecorder_DestroyQAudioRecorder(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAudioRecorder) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAudioRecorder::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAudioRecorder) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAudioRecorder::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAudioRecorderTimerEvent
func callbackQAudioRecorderTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioRecorder::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAudioRecorderFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAudioRecorder) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAudioRecorder::timerEvent")

	if ptr.Pointer() != nil {
		C.QAudioRecorder_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAudioRecorder) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAudioRecorder::timerEvent")

	if ptr.Pointer() != nil {
		C.QAudioRecorder_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAudioRecorder) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAudioRecorder::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAudioRecorder) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAudioRecorder::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAudioRecorderChildEvent
func callbackQAudioRecorderChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioRecorder::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAudioRecorderFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAudioRecorder) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAudioRecorder::childEvent")

	if ptr.Pointer() != nil {
		C.QAudioRecorder_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAudioRecorder) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAudioRecorder::childEvent")

	if ptr.Pointer() != nil {
		C.QAudioRecorder_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAudioRecorder) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAudioRecorder::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAudioRecorder) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAudioRecorder::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAudioRecorderCustomEvent
func callbackQAudioRecorderCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioRecorder::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAudioRecorderFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAudioRecorder) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAudioRecorder::customEvent")

	if ptr.Pointer() != nil {
		C.QAudioRecorder_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAudioRecorder) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAudioRecorder::customEvent")

	if ptr.Pointer() != nil {
		C.QAudioRecorder_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QCamera struct {
	QMediaObject
}

type QCamera_ITF interface {
	QMediaObject_ITF
	QCamera_PTR() *QCamera
}

func PointerFromQCamera(ptr QCamera_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCamera_PTR().Pointer()
	}
	return nil
}

func NewQCameraFromPointer(ptr unsafe.Pointer) *QCamera {
	var n = new(QCamera)
	n.SetPointer(ptr)
	return n
}

func newQCameraFromPointer(ptr unsafe.Pointer) *QCamera {
	var n = NewQCameraFromPointer(ptr)
	for len(n.ObjectName()) < len("QCamera_") {
		n.SetObjectName("QCamera_" + qt.Identifier())
	}
	return n
}

func (ptr *QCamera) QCamera_PTR() *QCamera {
	return ptr
}

//QCamera::CaptureMode
type QCamera__CaptureMode int64

const (
	QCamera__CaptureViewfinder = QCamera__CaptureMode(0)
	QCamera__CaptureStillImage = QCamera__CaptureMode(0x01)
	QCamera__CaptureVideo      = QCamera__CaptureMode(0x02)
)

//QCamera::Error
type QCamera__Error int64

const (
	QCamera__NoError                  = QCamera__Error(0)
	QCamera__CameraError              = QCamera__Error(1)
	QCamera__InvalidRequestError      = QCamera__Error(2)
	QCamera__ServiceMissingError      = QCamera__Error(3)
	QCamera__NotSupportedFeatureError = QCamera__Error(4)
)

//QCamera::LockChangeReason
type QCamera__LockChangeReason int64

const (
	QCamera__UserRequest       = QCamera__LockChangeReason(0)
	QCamera__LockAcquired      = QCamera__LockChangeReason(1)
	QCamera__LockFailed        = QCamera__LockChangeReason(2)
	QCamera__LockLost          = QCamera__LockChangeReason(3)
	QCamera__LockTemporaryLost = QCamera__LockChangeReason(4)
)

//QCamera::LockStatus
type QCamera__LockStatus int64

const (
	QCamera__Unlocked  = QCamera__LockStatus(0)
	QCamera__Searching = QCamera__LockStatus(1)
	QCamera__Locked    = QCamera__LockStatus(2)
)

//QCamera::LockType
type QCamera__LockType int64

const (
	QCamera__NoLock           = QCamera__LockType(0)
	QCamera__LockExposure     = QCamera__LockType(0x01)
	QCamera__LockWhiteBalance = QCamera__LockType(0x02)
	QCamera__LockFocus        = QCamera__LockType(0x04)
)

//QCamera::Position
type QCamera__Position int64

const (
	QCamera__UnspecifiedPosition = QCamera__Position(0)
	QCamera__BackFace            = QCamera__Position(1)
	QCamera__FrontFace           = QCamera__Position(2)
)

//QCamera::State
type QCamera__State int64

const (
	QCamera__UnloadedState = QCamera__State(0)
	QCamera__LoadedState   = QCamera__State(1)
	QCamera__ActiveState   = QCamera__State(2)
)

//QCamera::Status
type QCamera__Status int64

const (
	QCamera__UnavailableStatus = QCamera__Status(0)
	QCamera__UnloadedStatus    = QCamera__Status(1)
	QCamera__LoadingStatus     = QCamera__Status(2)
	QCamera__UnloadingStatus   = QCamera__Status(3)
	QCamera__LoadedStatus      = QCamera__Status(4)
	QCamera__StandbyStatus     = QCamera__Status(5)
	QCamera__StartingStatus    = QCamera__Status(6)
	QCamera__StoppingStatus    = QCamera__Status(7)
	QCamera__ActiveStatus      = QCamera__Status(8)
)

func (ptr *QCamera) CaptureMode() QCamera__CaptureMode {
	defer qt.Recovering("QCamera::captureMode")

	if ptr.Pointer() != nil {
		return QCamera__CaptureMode(C.QCamera_CaptureMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCamera) SearchAndLock2(locks QCamera__LockType) {
	defer qt.Recovering("QCamera::searchAndLock")

	if ptr.Pointer() != nil {
		C.QCamera_SearchAndLock2(ptr.Pointer(), C.int(locks))
	}
}

func (ptr *QCamera) SetCaptureMode(mode QCamera__CaptureMode) {
	defer qt.Recovering("QCamera::setCaptureMode")

	if ptr.Pointer() != nil {
		C.QCamera_SetCaptureMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCamera) State() QCamera__State {
	defer qt.Recovering("QCamera::state")

	if ptr.Pointer() != nil {
		return QCamera__State(C.QCamera_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCamera) Status() QCamera__Status {
	defer qt.Recovering("QCamera::status")

	if ptr.Pointer() != nil {
		return QCamera__Status(C.QCamera_Status(ptr.Pointer()))
	}
	return 0
}

func NewQCamera4(position QCamera__Position, parent core.QObject_ITF) *QCamera {
	defer qt.Recovering("QCamera::QCamera")

	return newQCameraFromPointer(C.QCamera_NewQCamera4(C.int(position), core.PointerFromQObject(parent)))
}

func NewQCamera(parent core.QObject_ITF) *QCamera {
	defer qt.Recovering("QCamera::QCamera")

	return newQCameraFromPointer(C.QCamera_NewQCamera(core.PointerFromQObject(parent)))
}

func NewQCamera2(deviceName string, parent core.QObject_ITF) *QCamera {
	defer qt.Recovering("QCamera::QCamera")

	return newQCameraFromPointer(C.QCamera_NewQCamera2(C.CString(deviceName), core.PointerFromQObject(parent)))
}

func NewQCamera3(cameraInfo QCameraInfo_ITF, parent core.QObject_ITF) *QCamera {
	defer qt.Recovering("QCamera::QCamera")

	return newQCameraFromPointer(C.QCamera_NewQCamera3(PointerFromQCameraInfo(cameraInfo), core.PointerFromQObject(parent)))
}

func (ptr *QCamera) Availability() QMultimedia__AvailabilityStatus {
	defer qt.Recovering("QCamera::availability")

	if ptr.Pointer() != nil {
		return QMultimedia__AvailabilityStatus(C.QCamera_Availability(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCamera) ConnectCaptureModeChanged(f func(mode QCamera__CaptureMode)) {
	defer qt.Recovering("connect QCamera::captureModeChanged")

	if ptr.Pointer() != nil {
		C.QCamera_ConnectCaptureModeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "captureModeChanged", f)
	}
}

func (ptr *QCamera) DisconnectCaptureModeChanged() {
	defer qt.Recovering("disconnect QCamera::captureModeChanged")

	if ptr.Pointer() != nil {
		C.QCamera_DisconnectCaptureModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "captureModeChanged")
	}
}

//export callbackQCameraCaptureModeChanged
func callbackQCameraCaptureModeChanged(ptr unsafe.Pointer, ptrName *C.char, mode C.int) {
	defer qt.Recovering("callback QCamera::captureModeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "captureModeChanged"); signal != nil {
		signal.(func(QCamera__CaptureMode))(QCamera__CaptureMode(mode))
	}

}

func (ptr *QCamera) CaptureModeChanged(mode QCamera__CaptureMode) {
	defer qt.Recovering("QCamera::captureModeChanged")

	if ptr.Pointer() != nil {
		C.QCamera_CaptureModeChanged(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCamera) ConnectError2(f func(value QCamera__Error)) {
	defer qt.Recovering("connect QCamera::error")

	if ptr.Pointer() != nil {
		C.QCamera_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QCamera) DisconnectError2() {
	defer qt.Recovering("disconnect QCamera::error")

	if ptr.Pointer() != nil {
		C.QCamera_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQCameraError2
func callbackQCameraError2(ptr unsafe.Pointer, ptrName *C.char, value C.int) {
	defer qt.Recovering("callback QCamera::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QCamera__Error))(QCamera__Error(value))
	}

}

func (ptr *QCamera) Error2(value QCamera__Error) {
	defer qt.Recovering("QCamera::error")

	if ptr.Pointer() != nil {
		C.QCamera_Error2(ptr.Pointer(), C.int(value))
	}
}

func (ptr *QCamera) Error() QCamera__Error {
	defer qt.Recovering("QCamera::error")

	if ptr.Pointer() != nil {
		return QCamera__Error(C.QCamera_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCamera) ErrorString() string {
	defer qt.Recovering("QCamera::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QCamera_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QCamera) Exposure() *QCameraExposure {
	defer qt.Recovering("QCamera::exposure")

	if ptr.Pointer() != nil {
		return NewQCameraExposureFromPointer(C.QCamera_Exposure(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCamera) Focus() *QCameraFocus {
	defer qt.Recovering("QCamera::focus")

	if ptr.Pointer() != nil {
		return NewQCameraFocusFromPointer(C.QCamera_Focus(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCamera) ImageProcessing() *QCameraImageProcessing {
	defer qt.Recovering("QCamera::imageProcessing")

	if ptr.Pointer() != nil {
		return NewQCameraImageProcessingFromPointer(C.QCamera_ImageProcessing(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCamera) IsCaptureModeSupported(mode QCamera__CaptureMode) bool {
	defer qt.Recovering("QCamera::isCaptureModeSupported")

	if ptr.Pointer() != nil {
		return C.QCamera_IsCaptureModeSupported(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCamera) Load() {
	defer qt.Recovering("QCamera::load")

	if ptr.Pointer() != nil {
		C.QCamera_Load(ptr.Pointer())
	}
}

func (ptr *QCamera) ConnectLockFailed(f func()) {
	defer qt.Recovering("connect QCamera::lockFailed")

	if ptr.Pointer() != nil {
		C.QCamera_ConnectLockFailed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "lockFailed", f)
	}
}

func (ptr *QCamera) DisconnectLockFailed() {
	defer qt.Recovering("disconnect QCamera::lockFailed")

	if ptr.Pointer() != nil {
		C.QCamera_DisconnectLockFailed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "lockFailed")
	}
}

//export callbackQCameraLockFailed
func callbackQCameraLockFailed(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QCamera::lockFailed")

	if signal := qt.GetSignal(C.GoString(ptrName), "lockFailed"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCamera) LockFailed() {
	defer qt.Recovering("QCamera::lockFailed")

	if ptr.Pointer() != nil {
		C.QCamera_LockFailed(ptr.Pointer())
	}
}

func (ptr *QCamera) LockStatus() QCamera__LockStatus {
	defer qt.Recovering("QCamera::lockStatus")

	if ptr.Pointer() != nil {
		return QCamera__LockStatus(C.QCamera_LockStatus(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCamera) LockStatus2(lockType QCamera__LockType) QCamera__LockStatus {
	defer qt.Recovering("QCamera::lockStatus")

	if ptr.Pointer() != nil {
		return QCamera__LockStatus(C.QCamera_LockStatus2(ptr.Pointer(), C.int(lockType)))
	}
	return 0
}

func (ptr *QCamera) ConnectLockStatusChanged(f func(status QCamera__LockStatus, reason QCamera__LockChangeReason)) {
	defer qt.Recovering("connect QCamera::lockStatusChanged")

	if ptr.Pointer() != nil {
		C.QCamera_ConnectLockStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "lockStatusChanged", f)
	}
}

func (ptr *QCamera) DisconnectLockStatusChanged() {
	defer qt.Recovering("disconnect QCamera::lockStatusChanged")

	if ptr.Pointer() != nil {
		C.QCamera_DisconnectLockStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "lockStatusChanged")
	}
}

//export callbackQCameraLockStatusChanged
func callbackQCameraLockStatusChanged(ptr unsafe.Pointer, ptrName *C.char, status C.int, reason C.int) {
	defer qt.Recovering("callback QCamera::lockStatusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "lockStatusChanged"); signal != nil {
		signal.(func(QCamera__LockStatus, QCamera__LockChangeReason))(QCamera__LockStatus(status), QCamera__LockChangeReason(reason))
	}

}

func (ptr *QCamera) LockStatusChanged(status QCamera__LockStatus, reason QCamera__LockChangeReason) {
	defer qt.Recovering("QCamera::lockStatusChanged")

	if ptr.Pointer() != nil {
		C.QCamera_LockStatusChanged(ptr.Pointer(), C.int(status), C.int(reason))
	}
}

func (ptr *QCamera) ConnectLockStatusChanged2(f func(lock QCamera__LockType, status QCamera__LockStatus, reason QCamera__LockChangeReason)) {
	defer qt.Recovering("connect QCamera::lockStatusChanged")

	if ptr.Pointer() != nil {
		C.QCamera_ConnectLockStatusChanged2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "lockStatusChanged2", f)
	}
}

func (ptr *QCamera) DisconnectLockStatusChanged2() {
	defer qt.Recovering("disconnect QCamera::lockStatusChanged")

	if ptr.Pointer() != nil {
		C.QCamera_DisconnectLockStatusChanged2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "lockStatusChanged2")
	}
}

//export callbackQCameraLockStatusChanged2
func callbackQCameraLockStatusChanged2(ptr unsafe.Pointer, ptrName *C.char, lock C.int, status C.int, reason C.int) {
	defer qt.Recovering("callback QCamera::lockStatusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "lockStatusChanged2"); signal != nil {
		signal.(func(QCamera__LockType, QCamera__LockStatus, QCamera__LockChangeReason))(QCamera__LockType(lock), QCamera__LockStatus(status), QCamera__LockChangeReason(reason))
	}

}

func (ptr *QCamera) LockStatusChanged2(lock QCamera__LockType, status QCamera__LockStatus, reason QCamera__LockChangeReason) {
	defer qt.Recovering("QCamera::lockStatusChanged")

	if ptr.Pointer() != nil {
		C.QCamera_LockStatusChanged2(ptr.Pointer(), C.int(lock), C.int(status), C.int(reason))
	}
}

func (ptr *QCamera) ConnectLocked(f func()) {
	defer qt.Recovering("connect QCamera::locked")

	if ptr.Pointer() != nil {
		C.QCamera_ConnectLocked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "locked", f)
	}
}

func (ptr *QCamera) DisconnectLocked() {
	defer qt.Recovering("disconnect QCamera::locked")

	if ptr.Pointer() != nil {
		C.QCamera_DisconnectLocked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "locked")
	}
}

//export callbackQCameraLocked
func callbackQCameraLocked(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QCamera::locked")

	if signal := qt.GetSignal(C.GoString(ptrName), "locked"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCamera) Locked() {
	defer qt.Recovering("QCamera::locked")

	if ptr.Pointer() != nil {
		C.QCamera_Locked(ptr.Pointer())
	}
}

func (ptr *QCamera) RequestedLocks() QCamera__LockType {
	defer qt.Recovering("QCamera::requestedLocks")

	if ptr.Pointer() != nil {
		return QCamera__LockType(C.QCamera_RequestedLocks(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCamera) SearchAndLock() {
	defer qt.Recovering("QCamera::searchAndLock")

	if ptr.Pointer() != nil {
		C.QCamera_SearchAndLock(ptr.Pointer())
	}
}

func (ptr *QCamera) SetViewfinder3(surface QAbstractVideoSurface_ITF) {
	defer qt.Recovering("QCamera::setViewfinder")

	if ptr.Pointer() != nil {
		C.QCamera_SetViewfinder3(ptr.Pointer(), PointerFromQAbstractVideoSurface(surface))
	}
}

func (ptr *QCamera) SetViewfinder2(viewfinder unsafe.Pointer) {
	defer qt.Recovering("QCamera::setViewfinder")

	if ptr.Pointer() != nil {
		C.QCamera_SetViewfinder2(ptr.Pointer(), viewfinder)
	}
}

func (ptr *QCamera) SetViewfinder(viewfinder unsafe.Pointer) {
	defer qt.Recovering("QCamera::setViewfinder")

	if ptr.Pointer() != nil {
		C.QCamera_SetViewfinder(ptr.Pointer(), viewfinder)
	}
}

func (ptr *QCamera) SetViewfinderSettings(settings QCameraViewfinderSettings_ITF) {
	defer qt.Recovering("QCamera::setViewfinderSettings")

	if ptr.Pointer() != nil {
		C.QCamera_SetViewfinderSettings(ptr.Pointer(), PointerFromQCameraViewfinderSettings(settings))
	}
}

func (ptr *QCamera) Start() {
	defer qt.Recovering("QCamera::start")

	if ptr.Pointer() != nil {
		C.QCamera_Start(ptr.Pointer())
	}
}

func (ptr *QCamera) ConnectStateChanged(f func(state QCamera__State)) {
	defer qt.Recovering("connect QCamera::stateChanged")

	if ptr.Pointer() != nil {
		C.QCamera_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QCamera) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QCamera::stateChanged")

	if ptr.Pointer() != nil {
		C.QCamera_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQCameraStateChanged
func callbackQCameraStateChanged(ptr unsafe.Pointer, ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QCamera::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QCamera__State))(QCamera__State(state))
	}

}

func (ptr *QCamera) StateChanged(state QCamera__State) {
	defer qt.Recovering("QCamera::stateChanged")

	if ptr.Pointer() != nil {
		C.QCamera_StateChanged(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QCamera) ConnectStatusChanged(f func(status QCamera__Status)) {
	defer qt.Recovering("connect QCamera::statusChanged")

	if ptr.Pointer() != nil {
		C.QCamera_ConnectStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "statusChanged", f)
	}
}

func (ptr *QCamera) DisconnectStatusChanged() {
	defer qt.Recovering("disconnect QCamera::statusChanged")

	if ptr.Pointer() != nil {
		C.QCamera_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "statusChanged")
	}
}

//export callbackQCameraStatusChanged
func callbackQCameraStatusChanged(ptr unsafe.Pointer, ptrName *C.char, status C.int) {
	defer qt.Recovering("callback QCamera::statusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "statusChanged"); signal != nil {
		signal.(func(QCamera__Status))(QCamera__Status(status))
	}

}

func (ptr *QCamera) StatusChanged(status QCamera__Status) {
	defer qt.Recovering("QCamera::statusChanged")

	if ptr.Pointer() != nil {
		C.QCamera_StatusChanged(ptr.Pointer(), C.int(status))
	}
}

func (ptr *QCamera) Stop() {
	defer qt.Recovering("QCamera::stop")

	if ptr.Pointer() != nil {
		C.QCamera_Stop(ptr.Pointer())
	}
}

func (ptr *QCamera) SupportedLocks() QCamera__LockType {
	defer qt.Recovering("QCamera::supportedLocks")

	if ptr.Pointer() != nil {
		return QCamera__LockType(C.QCamera_SupportedLocks(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCamera) Unload() {
	defer qt.Recovering("QCamera::unload")

	if ptr.Pointer() != nil {
		C.QCamera_Unload(ptr.Pointer())
	}
}

func (ptr *QCamera) Unlock() {
	defer qt.Recovering("QCamera::unlock")

	if ptr.Pointer() != nil {
		C.QCamera_Unlock(ptr.Pointer())
	}
}

func (ptr *QCamera) Unlock2(locks QCamera__LockType) {
	defer qt.Recovering("QCamera::unlock")

	if ptr.Pointer() != nil {
		C.QCamera_Unlock2(ptr.Pointer(), C.int(locks))
	}
}

func (ptr *QCamera) ViewfinderSettings() *QCameraViewfinderSettings {
	defer qt.Recovering("QCamera::viewfinderSettings")

	if ptr.Pointer() != nil {
		return NewQCameraViewfinderSettingsFromPointer(C.QCamera_ViewfinderSettings(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCamera) DestroyQCamera() {
	defer qt.Recovering("QCamera::~QCamera")

	if ptr.Pointer() != nil {
		C.QCamera_DestroyQCamera(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCamera) ConnectUnbind(f func(object *core.QObject)) {
	defer qt.Recovering("connect QCamera::unbind")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "unbind", f)
	}
}

func (ptr *QCamera) DisconnectUnbind() {
	defer qt.Recovering("disconnect QCamera::unbind")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "unbind")
	}
}

//export callbackQCameraUnbind
func callbackQCameraUnbind(ptr unsafe.Pointer, ptrName *C.char, object unsafe.Pointer) {
	defer qt.Recovering("callback QCamera::unbind")

	if signal := qt.GetSignal(C.GoString(ptrName), "unbind"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(object))
	} else {
		NewQCameraFromPointer(ptr).UnbindDefault(core.NewQObjectFromPointer(object))
	}
}

func (ptr *QCamera) Unbind(object core.QObject_ITF) {
	defer qt.Recovering("QCamera::unbind")

	if ptr.Pointer() != nil {
		C.QCamera_Unbind(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QCamera) UnbindDefault(object core.QObject_ITF) {
	defer qt.Recovering("QCamera::unbind")

	if ptr.Pointer() != nil {
		C.QCamera_UnbindDefault(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QCamera) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCamera::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCamera) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCamera::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraTimerEvent
func callbackQCameraTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCamera::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCamera) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCamera::timerEvent")

	if ptr.Pointer() != nil {
		C.QCamera_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCamera) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCamera::timerEvent")

	if ptr.Pointer() != nil {
		C.QCamera_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCamera) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCamera::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCamera) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCamera::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraChildEvent
func callbackQCameraChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCamera::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCamera) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCamera::childEvent")

	if ptr.Pointer() != nil {
		C.QCamera_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCamera) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCamera::childEvent")

	if ptr.Pointer() != nil {
		C.QCamera_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCamera) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCamera::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCamera) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCamera::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraCustomEvent
func callbackQCameraCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCamera::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCamera) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCamera::customEvent")

	if ptr.Pointer() != nil {
		C.QCamera_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCamera) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCamera::customEvent")

	if ptr.Pointer() != nil {
		C.QCamera_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QCameraCaptureBufferFormatControl struct {
	QMediaControl
}

type QCameraCaptureBufferFormatControl_ITF interface {
	QMediaControl_ITF
	QCameraCaptureBufferFormatControl_PTR() *QCameraCaptureBufferFormatControl
}

func PointerFromQCameraCaptureBufferFormatControl(ptr QCameraCaptureBufferFormatControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraCaptureBufferFormatControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraCaptureBufferFormatControlFromPointer(ptr unsafe.Pointer) *QCameraCaptureBufferFormatControl {
	var n = new(QCameraCaptureBufferFormatControl)
	n.SetPointer(ptr)
	return n
}

func newQCameraCaptureBufferFormatControlFromPointer(ptr unsafe.Pointer) *QCameraCaptureBufferFormatControl {
	var n = NewQCameraCaptureBufferFormatControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QCameraCaptureBufferFormatControl_") {
		n.SetObjectName("QCameraCaptureBufferFormatControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraCaptureBufferFormatControl) QCameraCaptureBufferFormatControl_PTR() *QCameraCaptureBufferFormatControl {
	return ptr
}

func (ptr *QCameraCaptureBufferFormatControl) BufferFormat() QVideoFrame__PixelFormat {
	defer qt.Recovering("QCameraCaptureBufferFormatControl::bufferFormat")

	if ptr.Pointer() != nil {
		return QVideoFrame__PixelFormat(C.QCameraCaptureBufferFormatControl_BufferFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraCaptureBufferFormatControl) ConnectBufferFormatChanged(f func(format QVideoFrame__PixelFormat)) {
	defer qt.Recovering("connect QCameraCaptureBufferFormatControl::bufferFormatChanged")

	if ptr.Pointer() != nil {
		C.QCameraCaptureBufferFormatControl_ConnectBufferFormatChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bufferFormatChanged", f)
	}
}

func (ptr *QCameraCaptureBufferFormatControl) DisconnectBufferFormatChanged() {
	defer qt.Recovering("disconnect QCameraCaptureBufferFormatControl::bufferFormatChanged")

	if ptr.Pointer() != nil {
		C.QCameraCaptureBufferFormatControl_DisconnectBufferFormatChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bufferFormatChanged")
	}
}

//export callbackQCameraCaptureBufferFormatControlBufferFormatChanged
func callbackQCameraCaptureBufferFormatControlBufferFormatChanged(ptr unsafe.Pointer, ptrName *C.char, format C.int) {
	defer qt.Recovering("callback QCameraCaptureBufferFormatControl::bufferFormatChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "bufferFormatChanged"); signal != nil {
		signal.(func(QVideoFrame__PixelFormat))(QVideoFrame__PixelFormat(format))
	}

}

func (ptr *QCameraCaptureBufferFormatControl) BufferFormatChanged(format QVideoFrame__PixelFormat) {
	defer qt.Recovering("QCameraCaptureBufferFormatControl::bufferFormatChanged")

	if ptr.Pointer() != nil {
		C.QCameraCaptureBufferFormatControl_BufferFormatChanged(ptr.Pointer(), C.int(format))
	}
}

func (ptr *QCameraCaptureBufferFormatControl) SetBufferFormat(format QVideoFrame__PixelFormat) {
	defer qt.Recovering("QCameraCaptureBufferFormatControl::setBufferFormat")

	if ptr.Pointer() != nil {
		C.QCameraCaptureBufferFormatControl_SetBufferFormat(ptr.Pointer(), C.int(format))
	}
}

func (ptr *QCameraCaptureBufferFormatControl) DestroyQCameraCaptureBufferFormatControl() {
	defer qt.Recovering("QCameraCaptureBufferFormatControl::~QCameraCaptureBufferFormatControl")

	if ptr.Pointer() != nil {
		C.QCameraCaptureBufferFormatControl_DestroyQCameraCaptureBufferFormatControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraCaptureBufferFormatControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraCaptureBufferFormatControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraCaptureBufferFormatControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraCaptureBufferFormatControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraCaptureBufferFormatControlTimerEvent
func callbackQCameraCaptureBufferFormatControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraCaptureBufferFormatControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraCaptureBufferFormatControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraCaptureBufferFormatControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraCaptureBufferFormatControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraCaptureBufferFormatControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraCaptureBufferFormatControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraCaptureBufferFormatControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraCaptureBufferFormatControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraCaptureBufferFormatControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraCaptureBufferFormatControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraCaptureBufferFormatControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraCaptureBufferFormatControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraCaptureBufferFormatControlChildEvent
func callbackQCameraCaptureBufferFormatControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraCaptureBufferFormatControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraCaptureBufferFormatControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraCaptureBufferFormatControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraCaptureBufferFormatControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraCaptureBufferFormatControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraCaptureBufferFormatControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraCaptureBufferFormatControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraCaptureBufferFormatControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraCaptureBufferFormatControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraCaptureBufferFormatControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraCaptureBufferFormatControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraCaptureBufferFormatControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraCaptureBufferFormatControlCustomEvent
func callbackQCameraCaptureBufferFormatControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraCaptureBufferFormatControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraCaptureBufferFormatControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraCaptureBufferFormatControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraCaptureBufferFormatControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraCaptureBufferFormatControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraCaptureBufferFormatControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraCaptureBufferFormatControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraCaptureBufferFormatControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QCameraCaptureDestinationControl struct {
	QMediaControl
}

type QCameraCaptureDestinationControl_ITF interface {
	QMediaControl_ITF
	QCameraCaptureDestinationControl_PTR() *QCameraCaptureDestinationControl
}

func PointerFromQCameraCaptureDestinationControl(ptr QCameraCaptureDestinationControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraCaptureDestinationControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraCaptureDestinationControlFromPointer(ptr unsafe.Pointer) *QCameraCaptureDestinationControl {
	var n = new(QCameraCaptureDestinationControl)
	n.SetPointer(ptr)
	return n
}

func newQCameraCaptureDestinationControlFromPointer(ptr unsafe.Pointer) *QCameraCaptureDestinationControl {
	var n = NewQCameraCaptureDestinationControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QCameraCaptureDestinationControl_") {
		n.SetObjectName("QCameraCaptureDestinationControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraCaptureDestinationControl) QCameraCaptureDestinationControl_PTR() *QCameraCaptureDestinationControl {
	return ptr
}

func (ptr *QCameraCaptureDestinationControl) CaptureDestination() QCameraImageCapture__CaptureDestination {
	defer qt.Recovering("QCameraCaptureDestinationControl::captureDestination")

	if ptr.Pointer() != nil {
		return QCameraImageCapture__CaptureDestination(C.QCameraCaptureDestinationControl_CaptureDestination(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraCaptureDestinationControl) ConnectCaptureDestinationChanged(f func(destination QCameraImageCapture__CaptureDestination)) {
	defer qt.Recovering("connect QCameraCaptureDestinationControl::captureDestinationChanged")

	if ptr.Pointer() != nil {
		C.QCameraCaptureDestinationControl_ConnectCaptureDestinationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "captureDestinationChanged", f)
	}
}

func (ptr *QCameraCaptureDestinationControl) DisconnectCaptureDestinationChanged() {
	defer qt.Recovering("disconnect QCameraCaptureDestinationControl::captureDestinationChanged")

	if ptr.Pointer() != nil {
		C.QCameraCaptureDestinationControl_DisconnectCaptureDestinationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "captureDestinationChanged")
	}
}

//export callbackQCameraCaptureDestinationControlCaptureDestinationChanged
func callbackQCameraCaptureDestinationControlCaptureDestinationChanged(ptr unsafe.Pointer, ptrName *C.char, destination C.int) {
	defer qt.Recovering("callback QCameraCaptureDestinationControl::captureDestinationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "captureDestinationChanged"); signal != nil {
		signal.(func(QCameraImageCapture__CaptureDestination))(QCameraImageCapture__CaptureDestination(destination))
	}

}

func (ptr *QCameraCaptureDestinationControl) CaptureDestinationChanged(destination QCameraImageCapture__CaptureDestination) {
	defer qt.Recovering("QCameraCaptureDestinationControl::captureDestinationChanged")

	if ptr.Pointer() != nil {
		C.QCameraCaptureDestinationControl_CaptureDestinationChanged(ptr.Pointer(), C.int(destination))
	}
}

func (ptr *QCameraCaptureDestinationControl) IsCaptureDestinationSupported(destination QCameraImageCapture__CaptureDestination) bool {
	defer qt.Recovering("QCameraCaptureDestinationControl::isCaptureDestinationSupported")

	if ptr.Pointer() != nil {
		return C.QCameraCaptureDestinationControl_IsCaptureDestinationSupported(ptr.Pointer(), C.int(destination)) != 0
	}
	return false
}

func (ptr *QCameraCaptureDestinationControl) SetCaptureDestination(destination QCameraImageCapture__CaptureDestination) {
	defer qt.Recovering("QCameraCaptureDestinationControl::setCaptureDestination")

	if ptr.Pointer() != nil {
		C.QCameraCaptureDestinationControl_SetCaptureDestination(ptr.Pointer(), C.int(destination))
	}
}

func (ptr *QCameraCaptureDestinationControl) DestroyQCameraCaptureDestinationControl() {
	defer qt.Recovering("QCameraCaptureDestinationControl::~QCameraCaptureDestinationControl")

	if ptr.Pointer() != nil {
		C.QCameraCaptureDestinationControl_DestroyQCameraCaptureDestinationControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraCaptureDestinationControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraCaptureDestinationControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraCaptureDestinationControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraCaptureDestinationControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraCaptureDestinationControlTimerEvent
func callbackQCameraCaptureDestinationControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraCaptureDestinationControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraCaptureDestinationControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraCaptureDestinationControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraCaptureDestinationControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraCaptureDestinationControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraCaptureDestinationControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraCaptureDestinationControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraCaptureDestinationControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraCaptureDestinationControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraCaptureDestinationControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraCaptureDestinationControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraCaptureDestinationControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraCaptureDestinationControlChildEvent
func callbackQCameraCaptureDestinationControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraCaptureDestinationControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraCaptureDestinationControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraCaptureDestinationControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraCaptureDestinationControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraCaptureDestinationControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraCaptureDestinationControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraCaptureDestinationControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraCaptureDestinationControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraCaptureDestinationControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraCaptureDestinationControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraCaptureDestinationControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraCaptureDestinationControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraCaptureDestinationControlCustomEvent
func callbackQCameraCaptureDestinationControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraCaptureDestinationControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraCaptureDestinationControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraCaptureDestinationControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraCaptureDestinationControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraCaptureDestinationControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraCaptureDestinationControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraCaptureDestinationControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraCaptureDestinationControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QCameraControl struct {
	QMediaControl
}

type QCameraControl_ITF interface {
	QMediaControl_ITF
	QCameraControl_PTR() *QCameraControl
}

func PointerFromQCameraControl(ptr QCameraControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraControlFromPointer(ptr unsafe.Pointer) *QCameraControl {
	var n = new(QCameraControl)
	n.SetPointer(ptr)
	return n
}

func newQCameraControlFromPointer(ptr unsafe.Pointer) *QCameraControl {
	var n = NewQCameraControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QCameraControl_") {
		n.SetObjectName("QCameraControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraControl) QCameraControl_PTR() *QCameraControl {
	return ptr
}

//QCameraControl::PropertyChangeType
type QCameraControl__PropertyChangeType int64

const (
	QCameraControl__CaptureMode           = QCameraControl__PropertyChangeType(1)
	QCameraControl__ImageEncodingSettings = QCameraControl__PropertyChangeType(2)
	QCameraControl__VideoEncodingSettings = QCameraControl__PropertyChangeType(3)
	QCameraControl__Viewfinder            = QCameraControl__PropertyChangeType(4)
	QCameraControl__ViewfinderSettings    = QCameraControl__PropertyChangeType(5)
)

func (ptr *QCameraControl) CanChangeProperty(changeType QCameraControl__PropertyChangeType, status QCamera__Status) bool {
	defer qt.Recovering("QCameraControl::canChangeProperty")

	if ptr.Pointer() != nil {
		return C.QCameraControl_CanChangeProperty(ptr.Pointer(), C.int(changeType), C.int(status)) != 0
	}
	return false
}

func (ptr *QCameraControl) CaptureMode() QCamera__CaptureMode {
	defer qt.Recovering("QCameraControl::captureMode")

	if ptr.Pointer() != nil {
		return QCamera__CaptureMode(C.QCameraControl_CaptureMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraControl) ConnectCaptureModeChanged(f func(mode QCamera__CaptureMode)) {
	defer qt.Recovering("connect QCameraControl::captureModeChanged")

	if ptr.Pointer() != nil {
		C.QCameraControl_ConnectCaptureModeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "captureModeChanged", f)
	}
}

func (ptr *QCameraControl) DisconnectCaptureModeChanged() {
	defer qt.Recovering("disconnect QCameraControl::captureModeChanged")

	if ptr.Pointer() != nil {
		C.QCameraControl_DisconnectCaptureModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "captureModeChanged")
	}
}

//export callbackQCameraControlCaptureModeChanged
func callbackQCameraControlCaptureModeChanged(ptr unsafe.Pointer, ptrName *C.char, mode C.int) {
	defer qt.Recovering("callback QCameraControl::captureModeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "captureModeChanged"); signal != nil {
		signal.(func(QCamera__CaptureMode))(QCamera__CaptureMode(mode))
	}

}

func (ptr *QCameraControl) CaptureModeChanged(mode QCamera__CaptureMode) {
	defer qt.Recovering("QCameraControl::captureModeChanged")

	if ptr.Pointer() != nil {
		C.QCameraControl_CaptureModeChanged(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraControl) ConnectError(f func(error int, errorString string)) {
	defer qt.Recovering("connect QCameraControl::error")

	if ptr.Pointer() != nil {
		C.QCameraControl_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QCameraControl) DisconnectError() {
	defer qt.Recovering("disconnect QCameraControl::error")

	if ptr.Pointer() != nil {
		C.QCameraControl_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQCameraControlError
func callbackQCameraControlError(ptr unsafe.Pointer, ptrName *C.char, error C.int, errorString *C.char) {
	defer qt.Recovering("callback QCameraControl::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error"); signal != nil {
		signal.(func(int, string))(int(error), C.GoString(errorString))
	}

}

func (ptr *QCameraControl) Error(error int, errorString string) {
	defer qt.Recovering("QCameraControl::error")

	if ptr.Pointer() != nil {
		C.QCameraControl_Error(ptr.Pointer(), C.int(error), C.CString(errorString))
	}
}

func (ptr *QCameraControl) IsCaptureModeSupported(mode QCamera__CaptureMode) bool {
	defer qt.Recovering("QCameraControl::isCaptureModeSupported")

	if ptr.Pointer() != nil {
		return C.QCameraControl_IsCaptureModeSupported(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraControl) SetCaptureMode(mode QCamera__CaptureMode) {
	defer qt.Recovering("QCameraControl::setCaptureMode")

	if ptr.Pointer() != nil {
		C.QCameraControl_SetCaptureMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraControl) SetState(state QCamera__State) {
	defer qt.Recovering("QCameraControl::setState")

	if ptr.Pointer() != nil {
		C.QCameraControl_SetState(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QCameraControl) State() QCamera__State {
	defer qt.Recovering("QCameraControl::state")

	if ptr.Pointer() != nil {
		return QCamera__State(C.QCameraControl_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraControl) ConnectStateChanged(f func(state QCamera__State)) {
	defer qt.Recovering("connect QCameraControl::stateChanged")

	if ptr.Pointer() != nil {
		C.QCameraControl_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QCameraControl) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QCameraControl::stateChanged")

	if ptr.Pointer() != nil {
		C.QCameraControl_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQCameraControlStateChanged
func callbackQCameraControlStateChanged(ptr unsafe.Pointer, ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QCameraControl::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QCamera__State))(QCamera__State(state))
	}

}

func (ptr *QCameraControl) StateChanged(state QCamera__State) {
	defer qt.Recovering("QCameraControl::stateChanged")

	if ptr.Pointer() != nil {
		C.QCameraControl_StateChanged(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QCameraControl) Status() QCamera__Status {
	defer qt.Recovering("QCameraControl::status")

	if ptr.Pointer() != nil {
		return QCamera__Status(C.QCameraControl_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraControl) ConnectStatusChanged(f func(status QCamera__Status)) {
	defer qt.Recovering("connect QCameraControl::statusChanged")

	if ptr.Pointer() != nil {
		C.QCameraControl_ConnectStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "statusChanged", f)
	}
}

func (ptr *QCameraControl) DisconnectStatusChanged() {
	defer qt.Recovering("disconnect QCameraControl::statusChanged")

	if ptr.Pointer() != nil {
		C.QCameraControl_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "statusChanged")
	}
}

//export callbackQCameraControlStatusChanged
func callbackQCameraControlStatusChanged(ptr unsafe.Pointer, ptrName *C.char, status C.int) {
	defer qt.Recovering("callback QCameraControl::statusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "statusChanged"); signal != nil {
		signal.(func(QCamera__Status))(QCamera__Status(status))
	}

}

func (ptr *QCameraControl) StatusChanged(status QCamera__Status) {
	defer qt.Recovering("QCameraControl::statusChanged")

	if ptr.Pointer() != nil {
		C.QCameraControl_StatusChanged(ptr.Pointer(), C.int(status))
	}
}

func (ptr *QCameraControl) DestroyQCameraControl() {
	defer qt.Recovering("QCameraControl::~QCameraControl")

	if ptr.Pointer() != nil {
		C.QCameraControl_DestroyQCameraControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraControlTimerEvent
func callbackQCameraControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraControlChildEvent
func callbackQCameraControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraControlCustomEvent
func callbackQCameraControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QCameraExposure struct {
	core.QObject
}

type QCameraExposure_ITF interface {
	core.QObject_ITF
	QCameraExposure_PTR() *QCameraExposure
}

func PointerFromQCameraExposure(ptr QCameraExposure_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraExposure_PTR().Pointer()
	}
	return nil
}

func NewQCameraExposureFromPointer(ptr unsafe.Pointer) *QCameraExposure {
	var n = new(QCameraExposure)
	n.SetPointer(ptr)
	return n
}

func newQCameraExposureFromPointer(ptr unsafe.Pointer) *QCameraExposure {
	var n = NewQCameraExposureFromPointer(ptr)
	for len(n.ObjectName()) < len("QCameraExposure_") {
		n.SetObjectName("QCameraExposure_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraExposure) QCameraExposure_PTR() *QCameraExposure {
	return ptr
}

//QCameraExposure::ExposureMode
type QCameraExposure__ExposureMode int64

const (
	QCameraExposure__ExposureAuto          = QCameraExposure__ExposureMode(0)
	QCameraExposure__ExposureManual        = QCameraExposure__ExposureMode(1)
	QCameraExposure__ExposurePortrait      = QCameraExposure__ExposureMode(2)
	QCameraExposure__ExposureNight         = QCameraExposure__ExposureMode(3)
	QCameraExposure__ExposureBacklight     = QCameraExposure__ExposureMode(4)
	QCameraExposure__ExposureSpotlight     = QCameraExposure__ExposureMode(5)
	QCameraExposure__ExposureSports        = QCameraExposure__ExposureMode(6)
	QCameraExposure__ExposureSnow          = QCameraExposure__ExposureMode(7)
	QCameraExposure__ExposureBeach         = QCameraExposure__ExposureMode(8)
	QCameraExposure__ExposureLargeAperture = QCameraExposure__ExposureMode(9)
	QCameraExposure__ExposureSmallAperture = QCameraExposure__ExposureMode(10)
	QCameraExposure__ExposureAction        = QCameraExposure__ExposureMode(11)
	QCameraExposure__ExposureLandscape     = QCameraExposure__ExposureMode(12)
	QCameraExposure__ExposureNightPortrait = QCameraExposure__ExposureMode(13)
	QCameraExposure__ExposureTheatre       = QCameraExposure__ExposureMode(14)
	QCameraExposure__ExposureSunset        = QCameraExposure__ExposureMode(15)
	QCameraExposure__ExposureSteadyPhoto   = QCameraExposure__ExposureMode(16)
	QCameraExposure__ExposureFireworks     = QCameraExposure__ExposureMode(17)
	QCameraExposure__ExposureParty         = QCameraExposure__ExposureMode(18)
	QCameraExposure__ExposureCandlelight   = QCameraExposure__ExposureMode(19)
	QCameraExposure__ExposureBarcode       = QCameraExposure__ExposureMode(20)
	QCameraExposure__ExposureModeVendor    = QCameraExposure__ExposureMode(1000)
)

//QCameraExposure::FlashMode
type QCameraExposure__FlashMode int64

const (
	QCameraExposure__FlashAuto                 = QCameraExposure__FlashMode(0x1)
	QCameraExposure__FlashOff                  = QCameraExposure__FlashMode(0x2)
	QCameraExposure__FlashOn                   = QCameraExposure__FlashMode(0x4)
	QCameraExposure__FlashRedEyeReduction      = QCameraExposure__FlashMode(0x8)
	QCameraExposure__FlashFill                 = QCameraExposure__FlashMode(0x10)
	QCameraExposure__FlashTorch                = QCameraExposure__FlashMode(0x20)
	QCameraExposure__FlashVideoLight           = QCameraExposure__FlashMode(0x40)
	QCameraExposure__FlashSlowSyncFrontCurtain = QCameraExposure__FlashMode(0x80)
	QCameraExposure__FlashSlowSyncRearCurtain  = QCameraExposure__FlashMode(0x100)
	QCameraExposure__FlashManual               = QCameraExposure__FlashMode(0x200)
)

//QCameraExposure::MeteringMode
type QCameraExposure__MeteringMode int64

const (
	QCameraExposure__MeteringMatrix  = QCameraExposure__MeteringMode(1)
	QCameraExposure__MeteringAverage = QCameraExposure__MeteringMode(2)
	QCameraExposure__MeteringSpot    = QCameraExposure__MeteringMode(3)
)

func (ptr *QCameraExposure) Aperture() float64 {
	defer qt.Recovering("QCameraExposure::aperture")

	if ptr.Pointer() != nil {
		return float64(C.QCameraExposure_Aperture(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraExposure) ExposureCompensation() float64 {
	defer qt.Recovering("QCameraExposure::exposureCompensation")

	if ptr.Pointer() != nil {
		return float64(C.QCameraExposure_ExposureCompensation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraExposure) ExposureMode() QCameraExposure__ExposureMode {
	defer qt.Recovering("QCameraExposure::exposureMode")

	if ptr.Pointer() != nil {
		return QCameraExposure__ExposureMode(C.QCameraExposure_ExposureMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraExposure) FlashMode() QCameraExposure__FlashMode {
	defer qt.Recovering("QCameraExposure::flashMode")

	if ptr.Pointer() != nil {
		return QCameraExposure__FlashMode(C.QCameraExposure_FlashMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraExposure) IsoSensitivity() int {
	defer qt.Recovering("QCameraExposure::isoSensitivity")

	if ptr.Pointer() != nil {
		return int(C.QCameraExposure_IsoSensitivity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraExposure) MeteringMode() QCameraExposure__MeteringMode {
	defer qt.Recovering("QCameraExposure::meteringMode")

	if ptr.Pointer() != nil {
		return QCameraExposure__MeteringMode(C.QCameraExposure_MeteringMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraExposure) SetAutoAperture() {
	defer qt.Recovering("QCameraExposure::setAutoAperture")

	if ptr.Pointer() != nil {
		C.QCameraExposure_SetAutoAperture(ptr.Pointer())
	}
}

func (ptr *QCameraExposure) SetAutoIsoSensitivity() {
	defer qt.Recovering("QCameraExposure::setAutoIsoSensitivity")

	if ptr.Pointer() != nil {
		C.QCameraExposure_SetAutoIsoSensitivity(ptr.Pointer())
	}
}

func (ptr *QCameraExposure) SetExposureCompensation(ev float64) {
	defer qt.Recovering("QCameraExposure::setExposureCompensation")

	if ptr.Pointer() != nil {
		C.QCameraExposure_SetExposureCompensation(ptr.Pointer(), C.double(ev))
	}
}

func (ptr *QCameraExposure) SetExposureMode(mode QCameraExposure__ExposureMode) {
	defer qt.Recovering("QCameraExposure::setExposureMode")

	if ptr.Pointer() != nil {
		C.QCameraExposure_SetExposureMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraExposure) SetFlashMode(mode QCameraExposure__FlashMode) {
	defer qt.Recovering("QCameraExposure::setFlashMode")

	if ptr.Pointer() != nil {
		C.QCameraExposure_SetFlashMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraExposure) SetManualAperture(aperture float64) {
	defer qt.Recovering("QCameraExposure::setManualAperture")

	if ptr.Pointer() != nil {
		C.QCameraExposure_SetManualAperture(ptr.Pointer(), C.double(aperture))
	}
}

func (ptr *QCameraExposure) SetManualIsoSensitivity(iso int) {
	defer qt.Recovering("QCameraExposure::setManualIsoSensitivity")

	if ptr.Pointer() != nil {
		C.QCameraExposure_SetManualIsoSensitivity(ptr.Pointer(), C.int(iso))
	}
}

func (ptr *QCameraExposure) SetMeteringMode(mode QCameraExposure__MeteringMode) {
	defer qt.Recovering("QCameraExposure::setMeteringMode")

	if ptr.Pointer() != nil {
		C.QCameraExposure_SetMeteringMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraExposure) SetSpotMeteringPoint(point core.QPointF_ITF) {
	defer qt.Recovering("QCameraExposure::setSpotMeteringPoint")

	if ptr.Pointer() != nil {
		C.QCameraExposure_SetSpotMeteringPoint(ptr.Pointer(), core.PointerFromQPointF(point))
	}
}

func (ptr *QCameraExposure) ConnectApertureChanged(f func(value float64)) {
	defer qt.Recovering("connect QCameraExposure::apertureChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposure_ConnectApertureChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "apertureChanged", f)
	}
}

func (ptr *QCameraExposure) DisconnectApertureChanged() {
	defer qt.Recovering("disconnect QCameraExposure::apertureChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposure_DisconnectApertureChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "apertureChanged")
	}
}

//export callbackQCameraExposureApertureChanged
func callbackQCameraExposureApertureChanged(ptr unsafe.Pointer, ptrName *C.char, value C.double) {
	defer qt.Recovering("callback QCameraExposure::apertureChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "apertureChanged"); signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QCameraExposure) ApertureChanged(value float64) {
	defer qt.Recovering("QCameraExposure::apertureChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposure_ApertureChanged(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QCameraExposure) ConnectApertureRangeChanged(f func()) {
	defer qt.Recovering("connect QCameraExposure::apertureRangeChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposure_ConnectApertureRangeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "apertureRangeChanged", f)
	}
}

func (ptr *QCameraExposure) DisconnectApertureRangeChanged() {
	defer qt.Recovering("disconnect QCameraExposure::apertureRangeChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposure_DisconnectApertureRangeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "apertureRangeChanged")
	}
}

//export callbackQCameraExposureApertureRangeChanged
func callbackQCameraExposureApertureRangeChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QCameraExposure::apertureRangeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "apertureRangeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCameraExposure) ApertureRangeChanged() {
	defer qt.Recovering("QCameraExposure::apertureRangeChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposure_ApertureRangeChanged(ptr.Pointer())
	}
}

func (ptr *QCameraExposure) ConnectExposureCompensationChanged(f func(value float64)) {
	defer qt.Recovering("connect QCameraExposure::exposureCompensationChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposure_ConnectExposureCompensationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "exposureCompensationChanged", f)
	}
}

func (ptr *QCameraExposure) DisconnectExposureCompensationChanged() {
	defer qt.Recovering("disconnect QCameraExposure::exposureCompensationChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposure_DisconnectExposureCompensationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "exposureCompensationChanged")
	}
}

//export callbackQCameraExposureExposureCompensationChanged
func callbackQCameraExposureExposureCompensationChanged(ptr unsafe.Pointer, ptrName *C.char, value C.double) {
	defer qt.Recovering("callback QCameraExposure::exposureCompensationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "exposureCompensationChanged"); signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QCameraExposure) ExposureCompensationChanged(value float64) {
	defer qt.Recovering("QCameraExposure::exposureCompensationChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposure_ExposureCompensationChanged(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QCameraExposure) ConnectFlashReady(f func(ready bool)) {
	defer qt.Recovering("connect QCameraExposure::flashReady")

	if ptr.Pointer() != nil {
		C.QCameraExposure_ConnectFlashReady(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "flashReady", f)
	}
}

func (ptr *QCameraExposure) DisconnectFlashReady() {
	defer qt.Recovering("disconnect QCameraExposure::flashReady")

	if ptr.Pointer() != nil {
		C.QCameraExposure_DisconnectFlashReady(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "flashReady")
	}
}

//export callbackQCameraExposureFlashReady
func callbackQCameraExposureFlashReady(ptr unsafe.Pointer, ptrName *C.char, ready C.int) {
	defer qt.Recovering("callback QCameraExposure::flashReady")

	if signal := qt.GetSignal(C.GoString(ptrName), "flashReady"); signal != nil {
		signal.(func(bool))(int(ready) != 0)
	}

}

func (ptr *QCameraExposure) FlashReady(ready bool) {
	defer qt.Recovering("QCameraExposure::flashReady")

	if ptr.Pointer() != nil {
		C.QCameraExposure_FlashReady(ptr.Pointer(), C.int(qt.GoBoolToInt(ready)))
	}
}

func (ptr *QCameraExposure) IsAvailable() bool {
	defer qt.Recovering("QCameraExposure::isAvailable")

	if ptr.Pointer() != nil {
		return C.QCameraExposure_IsAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCameraExposure) IsExposureModeSupported(mode QCameraExposure__ExposureMode) bool {
	defer qt.Recovering("QCameraExposure::isExposureModeSupported")

	if ptr.Pointer() != nil {
		return C.QCameraExposure_IsExposureModeSupported(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraExposure) IsFlashModeSupported(mode QCameraExposure__FlashMode) bool {
	defer qt.Recovering("QCameraExposure::isFlashModeSupported")

	if ptr.Pointer() != nil {
		return C.QCameraExposure_IsFlashModeSupported(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraExposure) IsFlashReady() bool {
	defer qt.Recovering("QCameraExposure::isFlashReady")

	if ptr.Pointer() != nil {
		return C.QCameraExposure_IsFlashReady(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCameraExposure) IsMeteringModeSupported(mode QCameraExposure__MeteringMode) bool {
	defer qt.Recovering("QCameraExposure::isMeteringModeSupported")

	if ptr.Pointer() != nil {
		return C.QCameraExposure_IsMeteringModeSupported(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraExposure) ConnectIsoSensitivityChanged(f func(value int)) {
	defer qt.Recovering("connect QCameraExposure::isoSensitivityChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposure_ConnectIsoSensitivityChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "isoSensitivityChanged", f)
	}
}

func (ptr *QCameraExposure) DisconnectIsoSensitivityChanged() {
	defer qt.Recovering("disconnect QCameraExposure::isoSensitivityChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposure_DisconnectIsoSensitivityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "isoSensitivityChanged")
	}
}

//export callbackQCameraExposureIsoSensitivityChanged
func callbackQCameraExposureIsoSensitivityChanged(ptr unsafe.Pointer, ptrName *C.char, value C.int) {
	defer qt.Recovering("callback QCameraExposure::isoSensitivityChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "isoSensitivityChanged"); signal != nil {
		signal.(func(int))(int(value))
	}

}

func (ptr *QCameraExposure) IsoSensitivityChanged(value int) {
	defer qt.Recovering("QCameraExposure::isoSensitivityChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposure_IsoSensitivityChanged(ptr.Pointer(), C.int(value))
	}
}

func (ptr *QCameraExposure) RequestedAperture() float64 {
	defer qt.Recovering("QCameraExposure::requestedAperture")

	if ptr.Pointer() != nil {
		return float64(C.QCameraExposure_RequestedAperture(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraExposure) RequestedIsoSensitivity() int {
	defer qt.Recovering("QCameraExposure::requestedIsoSensitivity")

	if ptr.Pointer() != nil {
		return int(C.QCameraExposure_RequestedIsoSensitivity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraExposure) RequestedShutterSpeed() float64 {
	defer qt.Recovering("QCameraExposure::requestedShutterSpeed")

	if ptr.Pointer() != nil {
		return float64(C.QCameraExposure_RequestedShutterSpeed(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraExposure) SetAutoShutterSpeed() {
	defer qt.Recovering("QCameraExposure::setAutoShutterSpeed")

	if ptr.Pointer() != nil {
		C.QCameraExposure_SetAutoShutterSpeed(ptr.Pointer())
	}
}

func (ptr *QCameraExposure) SetManualShutterSpeed(seconds float64) {
	defer qt.Recovering("QCameraExposure::setManualShutterSpeed")

	if ptr.Pointer() != nil {
		C.QCameraExposure_SetManualShutterSpeed(ptr.Pointer(), C.double(seconds))
	}
}

func (ptr *QCameraExposure) ShutterSpeed() float64 {
	defer qt.Recovering("QCameraExposure::shutterSpeed")

	if ptr.Pointer() != nil {
		return float64(C.QCameraExposure_ShutterSpeed(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraExposure) ConnectShutterSpeedChanged(f func(speed float64)) {
	defer qt.Recovering("connect QCameraExposure::shutterSpeedChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposure_ConnectShutterSpeedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "shutterSpeedChanged", f)
	}
}

func (ptr *QCameraExposure) DisconnectShutterSpeedChanged() {
	defer qt.Recovering("disconnect QCameraExposure::shutterSpeedChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposure_DisconnectShutterSpeedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "shutterSpeedChanged")
	}
}

//export callbackQCameraExposureShutterSpeedChanged
func callbackQCameraExposureShutterSpeedChanged(ptr unsafe.Pointer, ptrName *C.char, speed C.double) {
	defer qt.Recovering("callback QCameraExposure::shutterSpeedChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "shutterSpeedChanged"); signal != nil {
		signal.(func(float64))(float64(speed))
	}

}

func (ptr *QCameraExposure) ShutterSpeedChanged(speed float64) {
	defer qt.Recovering("QCameraExposure::shutterSpeedChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposure_ShutterSpeedChanged(ptr.Pointer(), C.double(speed))
	}
}

func (ptr *QCameraExposure) ConnectShutterSpeedRangeChanged(f func()) {
	defer qt.Recovering("connect QCameraExposure::shutterSpeedRangeChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposure_ConnectShutterSpeedRangeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "shutterSpeedRangeChanged", f)
	}
}

func (ptr *QCameraExposure) DisconnectShutterSpeedRangeChanged() {
	defer qt.Recovering("disconnect QCameraExposure::shutterSpeedRangeChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposure_DisconnectShutterSpeedRangeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "shutterSpeedRangeChanged")
	}
}

//export callbackQCameraExposureShutterSpeedRangeChanged
func callbackQCameraExposureShutterSpeedRangeChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QCameraExposure::shutterSpeedRangeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "shutterSpeedRangeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCameraExposure) ShutterSpeedRangeChanged() {
	defer qt.Recovering("QCameraExposure::shutterSpeedRangeChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposure_ShutterSpeedRangeChanged(ptr.Pointer())
	}
}

func (ptr *QCameraExposure) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraExposure::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraExposure) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraExposure::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraExposureTimerEvent
func callbackQCameraExposureTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraExposure::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraExposureFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraExposure) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraExposure::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraExposure_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraExposure) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraExposure::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraExposure_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraExposure) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraExposure::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraExposure) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraExposure::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraExposureChildEvent
func callbackQCameraExposureChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraExposure::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraExposureFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraExposure) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraExposure::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraExposure_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraExposure) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraExposure::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraExposure_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraExposure) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraExposure::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraExposure) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraExposure::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraExposureCustomEvent
func callbackQCameraExposureCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraExposure::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraExposureFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraExposure) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraExposure::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraExposure_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraExposure) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraExposure::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraExposure_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QCameraExposureControl struct {
	QMediaControl
}

type QCameraExposureControl_ITF interface {
	QMediaControl_ITF
	QCameraExposureControl_PTR() *QCameraExposureControl
}

func PointerFromQCameraExposureControl(ptr QCameraExposureControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraExposureControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraExposureControlFromPointer(ptr unsafe.Pointer) *QCameraExposureControl {
	var n = new(QCameraExposureControl)
	n.SetPointer(ptr)
	return n
}

func newQCameraExposureControlFromPointer(ptr unsafe.Pointer) *QCameraExposureControl {
	var n = NewQCameraExposureControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QCameraExposureControl_") {
		n.SetObjectName("QCameraExposureControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraExposureControl) QCameraExposureControl_PTR() *QCameraExposureControl {
	return ptr
}

//QCameraExposureControl::ExposureParameter
type QCameraExposureControl__ExposureParameter int64

const (
	QCameraExposureControl__ISO                       = QCameraExposureControl__ExposureParameter(0)
	QCameraExposureControl__Aperture                  = QCameraExposureControl__ExposureParameter(1)
	QCameraExposureControl__ShutterSpeed              = QCameraExposureControl__ExposureParameter(2)
	QCameraExposureControl__ExposureCompensation      = QCameraExposureControl__ExposureParameter(3)
	QCameraExposureControl__FlashPower                = QCameraExposureControl__ExposureParameter(4)
	QCameraExposureControl__FlashCompensation         = QCameraExposureControl__ExposureParameter(5)
	QCameraExposureControl__TorchPower                = QCameraExposureControl__ExposureParameter(6)
	QCameraExposureControl__SpotMeteringPoint         = QCameraExposureControl__ExposureParameter(7)
	QCameraExposureControl__ExposureMode              = QCameraExposureControl__ExposureParameter(8)
	QCameraExposureControl__MeteringMode              = QCameraExposureControl__ExposureParameter(9)
	QCameraExposureControl__ExtendedExposureParameter = QCameraExposureControl__ExposureParameter(1000)
)

func (ptr *QCameraExposureControl) ActualValue(parameter QCameraExposureControl__ExposureParameter) *core.QVariant {
	defer qt.Recovering("QCameraExposureControl::actualValue")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QCameraExposureControl_ActualValue(ptr.Pointer(), C.int(parameter)))
	}
	return nil
}

func (ptr *QCameraExposureControl) ConnectActualValueChanged(f func(parameter int)) {
	defer qt.Recovering("connect QCameraExposureControl::actualValueChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_ConnectActualValueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "actualValueChanged", f)
	}
}

func (ptr *QCameraExposureControl) DisconnectActualValueChanged() {
	defer qt.Recovering("disconnect QCameraExposureControl::actualValueChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_DisconnectActualValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "actualValueChanged")
	}
}

//export callbackQCameraExposureControlActualValueChanged
func callbackQCameraExposureControlActualValueChanged(ptr unsafe.Pointer, ptrName *C.char, parameter C.int) {
	defer qt.Recovering("callback QCameraExposureControl::actualValueChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "actualValueChanged"); signal != nil {
		signal.(func(int))(int(parameter))
	}

}

func (ptr *QCameraExposureControl) ActualValueChanged(parameter int) {
	defer qt.Recovering("QCameraExposureControl::actualValueChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_ActualValueChanged(ptr.Pointer(), C.int(parameter))
	}
}

func (ptr *QCameraExposureControl) IsParameterSupported(parameter QCameraExposureControl__ExposureParameter) bool {
	defer qt.Recovering("QCameraExposureControl::isParameterSupported")

	if ptr.Pointer() != nil {
		return C.QCameraExposureControl_IsParameterSupported(ptr.Pointer(), C.int(parameter)) != 0
	}
	return false
}

func (ptr *QCameraExposureControl) ConnectParameterRangeChanged(f func(parameter int)) {
	defer qt.Recovering("connect QCameraExposureControl::parameterRangeChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_ConnectParameterRangeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "parameterRangeChanged", f)
	}
}

func (ptr *QCameraExposureControl) DisconnectParameterRangeChanged() {
	defer qt.Recovering("disconnect QCameraExposureControl::parameterRangeChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_DisconnectParameterRangeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "parameterRangeChanged")
	}
}

//export callbackQCameraExposureControlParameterRangeChanged
func callbackQCameraExposureControlParameterRangeChanged(ptr unsafe.Pointer, ptrName *C.char, parameter C.int) {
	defer qt.Recovering("callback QCameraExposureControl::parameterRangeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "parameterRangeChanged"); signal != nil {
		signal.(func(int))(int(parameter))
	}

}

func (ptr *QCameraExposureControl) ParameterRangeChanged(parameter int) {
	defer qt.Recovering("QCameraExposureControl::parameterRangeChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_ParameterRangeChanged(ptr.Pointer(), C.int(parameter))
	}
}

func (ptr *QCameraExposureControl) RequestedValue(parameter QCameraExposureControl__ExposureParameter) *core.QVariant {
	defer qt.Recovering("QCameraExposureControl::requestedValue")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QCameraExposureControl_RequestedValue(ptr.Pointer(), C.int(parameter)))
	}
	return nil
}

func (ptr *QCameraExposureControl) ConnectRequestedValueChanged(f func(parameter int)) {
	defer qt.Recovering("connect QCameraExposureControl::requestedValueChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_ConnectRequestedValueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "requestedValueChanged", f)
	}
}

func (ptr *QCameraExposureControl) DisconnectRequestedValueChanged() {
	defer qt.Recovering("disconnect QCameraExposureControl::requestedValueChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_DisconnectRequestedValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "requestedValueChanged")
	}
}

//export callbackQCameraExposureControlRequestedValueChanged
func callbackQCameraExposureControlRequestedValueChanged(ptr unsafe.Pointer, ptrName *C.char, parameter C.int) {
	defer qt.Recovering("callback QCameraExposureControl::requestedValueChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "requestedValueChanged"); signal != nil {
		signal.(func(int))(int(parameter))
	}

}

func (ptr *QCameraExposureControl) RequestedValueChanged(parameter int) {
	defer qt.Recovering("QCameraExposureControl::requestedValueChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_RequestedValueChanged(ptr.Pointer(), C.int(parameter))
	}
}

func (ptr *QCameraExposureControl) SetValue(parameter QCameraExposureControl__ExposureParameter, value core.QVariant_ITF) bool {
	defer qt.Recovering("QCameraExposureControl::setValue")

	if ptr.Pointer() != nil {
		return C.QCameraExposureControl_SetValue(ptr.Pointer(), C.int(parameter), core.PointerFromQVariant(value)) != 0
	}
	return false
}

func (ptr *QCameraExposureControl) DestroyQCameraExposureControl() {
	defer qt.Recovering("QCameraExposureControl::~QCameraExposureControl")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_DestroyQCameraExposureControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraExposureControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraExposureControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraExposureControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraExposureControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraExposureControlTimerEvent
func callbackQCameraExposureControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraExposureControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraExposureControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraExposureControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraExposureControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraExposureControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraExposureControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraExposureControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraExposureControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraExposureControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraExposureControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraExposureControlChildEvent
func callbackQCameraExposureControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraExposureControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraExposureControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraExposureControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraExposureControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraExposureControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraExposureControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraExposureControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraExposureControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraExposureControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraExposureControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraExposureControlCustomEvent
func callbackQCameraExposureControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraExposureControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraExposureControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraExposureControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraExposureControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraExposureControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraExposureControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QCameraFeedbackControl struct {
	QMediaControl
}

type QCameraFeedbackControl_ITF interface {
	QMediaControl_ITF
	QCameraFeedbackControl_PTR() *QCameraFeedbackControl
}

func PointerFromQCameraFeedbackControl(ptr QCameraFeedbackControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraFeedbackControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraFeedbackControlFromPointer(ptr unsafe.Pointer) *QCameraFeedbackControl {
	var n = new(QCameraFeedbackControl)
	n.SetPointer(ptr)
	return n
}

func newQCameraFeedbackControlFromPointer(ptr unsafe.Pointer) *QCameraFeedbackControl {
	var n = NewQCameraFeedbackControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QCameraFeedbackControl_") {
		n.SetObjectName("QCameraFeedbackControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraFeedbackControl) QCameraFeedbackControl_PTR() *QCameraFeedbackControl {
	return ptr
}

//QCameraFeedbackControl::EventType
type QCameraFeedbackControl__EventType int64

const (
	QCameraFeedbackControl__ViewfinderStarted   = QCameraFeedbackControl__EventType(1)
	QCameraFeedbackControl__ViewfinderStopped   = QCameraFeedbackControl__EventType(2)
	QCameraFeedbackControl__ImageCaptured       = QCameraFeedbackControl__EventType(3)
	QCameraFeedbackControl__ImageSaved          = QCameraFeedbackControl__EventType(4)
	QCameraFeedbackControl__ImageError          = QCameraFeedbackControl__EventType(5)
	QCameraFeedbackControl__RecordingStarted    = QCameraFeedbackControl__EventType(6)
	QCameraFeedbackControl__RecordingInProgress = QCameraFeedbackControl__EventType(7)
	QCameraFeedbackControl__RecordingStopped    = QCameraFeedbackControl__EventType(8)
	QCameraFeedbackControl__AutoFocusInProgress = QCameraFeedbackControl__EventType(9)
	QCameraFeedbackControl__AutoFocusLocked     = QCameraFeedbackControl__EventType(10)
	QCameraFeedbackControl__AutoFocusFailed     = QCameraFeedbackControl__EventType(11)
)

func (ptr *QCameraFeedbackControl) IsEventFeedbackEnabled(event QCameraFeedbackControl__EventType) bool {
	defer qt.Recovering("QCameraFeedbackControl::isEventFeedbackEnabled")

	if ptr.Pointer() != nil {
		return C.QCameraFeedbackControl_IsEventFeedbackEnabled(ptr.Pointer(), C.int(event)) != 0
	}
	return false
}

func (ptr *QCameraFeedbackControl) IsEventFeedbackLocked(event QCameraFeedbackControl__EventType) bool {
	defer qt.Recovering("QCameraFeedbackControl::isEventFeedbackLocked")

	if ptr.Pointer() != nil {
		return C.QCameraFeedbackControl_IsEventFeedbackLocked(ptr.Pointer(), C.int(event)) != 0
	}
	return false
}

func (ptr *QCameraFeedbackControl) ResetEventFeedback(event QCameraFeedbackControl__EventType) {
	defer qt.Recovering("QCameraFeedbackControl::resetEventFeedback")

	if ptr.Pointer() != nil {
		C.QCameraFeedbackControl_ResetEventFeedback(ptr.Pointer(), C.int(event))
	}
}

func (ptr *QCameraFeedbackControl) SetEventFeedbackEnabled(event QCameraFeedbackControl__EventType, enabled bool) bool {
	defer qt.Recovering("QCameraFeedbackControl::setEventFeedbackEnabled")

	if ptr.Pointer() != nil {
		return C.QCameraFeedbackControl_SetEventFeedbackEnabled(ptr.Pointer(), C.int(event), C.int(qt.GoBoolToInt(enabled))) != 0
	}
	return false
}

func (ptr *QCameraFeedbackControl) SetEventFeedbackSound(event QCameraFeedbackControl__EventType, filePath string) bool {
	defer qt.Recovering("QCameraFeedbackControl::setEventFeedbackSound")

	if ptr.Pointer() != nil {
		return C.QCameraFeedbackControl_SetEventFeedbackSound(ptr.Pointer(), C.int(event), C.CString(filePath)) != 0
	}
	return false
}

func (ptr *QCameraFeedbackControl) DestroyQCameraFeedbackControl() {
	defer qt.Recovering("QCameraFeedbackControl::~QCameraFeedbackControl")

	if ptr.Pointer() != nil {
		C.QCameraFeedbackControl_DestroyQCameraFeedbackControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraFeedbackControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraFeedbackControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraFeedbackControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraFeedbackControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraFeedbackControlTimerEvent
func callbackQCameraFeedbackControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraFeedbackControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraFeedbackControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraFeedbackControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraFeedbackControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraFeedbackControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraFeedbackControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraFeedbackControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraFeedbackControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraFeedbackControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraFeedbackControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraFeedbackControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraFeedbackControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraFeedbackControlChildEvent
func callbackQCameraFeedbackControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraFeedbackControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraFeedbackControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraFeedbackControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraFeedbackControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraFeedbackControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraFeedbackControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraFeedbackControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraFeedbackControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraFeedbackControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraFeedbackControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraFeedbackControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraFeedbackControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraFeedbackControlCustomEvent
func callbackQCameraFeedbackControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraFeedbackControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraFeedbackControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraFeedbackControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraFeedbackControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraFeedbackControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraFeedbackControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraFeedbackControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraFeedbackControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QCameraFlashControl struct {
	QMediaControl
}

type QCameraFlashControl_ITF interface {
	QMediaControl_ITF
	QCameraFlashControl_PTR() *QCameraFlashControl
}

func PointerFromQCameraFlashControl(ptr QCameraFlashControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraFlashControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraFlashControlFromPointer(ptr unsafe.Pointer) *QCameraFlashControl {
	var n = new(QCameraFlashControl)
	n.SetPointer(ptr)
	return n
}

func newQCameraFlashControlFromPointer(ptr unsafe.Pointer) *QCameraFlashControl {
	var n = NewQCameraFlashControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QCameraFlashControl_") {
		n.SetObjectName("QCameraFlashControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraFlashControl) QCameraFlashControl_PTR() *QCameraFlashControl {
	return ptr
}

func (ptr *QCameraFlashControl) FlashMode() QCameraExposure__FlashMode {
	defer qt.Recovering("QCameraFlashControl::flashMode")

	if ptr.Pointer() != nil {
		return QCameraExposure__FlashMode(C.QCameraFlashControl_FlashMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFlashControl) ConnectFlashReady(f func(ready bool)) {
	defer qt.Recovering("connect QCameraFlashControl::flashReady")

	if ptr.Pointer() != nil {
		C.QCameraFlashControl_ConnectFlashReady(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "flashReady", f)
	}
}

func (ptr *QCameraFlashControl) DisconnectFlashReady() {
	defer qt.Recovering("disconnect QCameraFlashControl::flashReady")

	if ptr.Pointer() != nil {
		C.QCameraFlashControl_DisconnectFlashReady(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "flashReady")
	}
}

//export callbackQCameraFlashControlFlashReady
func callbackQCameraFlashControlFlashReady(ptr unsafe.Pointer, ptrName *C.char, ready C.int) {
	defer qt.Recovering("callback QCameraFlashControl::flashReady")

	if signal := qt.GetSignal(C.GoString(ptrName), "flashReady"); signal != nil {
		signal.(func(bool))(int(ready) != 0)
	}

}

func (ptr *QCameraFlashControl) FlashReady(ready bool) {
	defer qt.Recovering("QCameraFlashControl::flashReady")

	if ptr.Pointer() != nil {
		C.QCameraFlashControl_FlashReady(ptr.Pointer(), C.int(qt.GoBoolToInt(ready)))
	}
}

func (ptr *QCameraFlashControl) IsFlashModeSupported(mode QCameraExposure__FlashMode) bool {
	defer qt.Recovering("QCameraFlashControl::isFlashModeSupported")

	if ptr.Pointer() != nil {
		return C.QCameraFlashControl_IsFlashModeSupported(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraFlashControl) IsFlashReady() bool {
	defer qt.Recovering("QCameraFlashControl::isFlashReady")

	if ptr.Pointer() != nil {
		return C.QCameraFlashControl_IsFlashReady(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCameraFlashControl) SetFlashMode(mode QCameraExposure__FlashMode) {
	defer qt.Recovering("QCameraFlashControl::setFlashMode")

	if ptr.Pointer() != nil {
		C.QCameraFlashControl_SetFlashMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraFlashControl) DestroyQCameraFlashControl() {
	defer qt.Recovering("QCameraFlashControl::~QCameraFlashControl")

	if ptr.Pointer() != nil {
		C.QCameraFlashControl_DestroyQCameraFlashControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraFlashControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraFlashControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraFlashControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraFlashControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraFlashControlTimerEvent
func callbackQCameraFlashControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraFlashControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraFlashControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraFlashControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraFlashControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraFlashControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraFlashControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraFlashControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraFlashControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraFlashControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraFlashControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraFlashControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraFlashControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraFlashControlChildEvent
func callbackQCameraFlashControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraFlashControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraFlashControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraFlashControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraFlashControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraFlashControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraFlashControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraFlashControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraFlashControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraFlashControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraFlashControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraFlashControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraFlashControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraFlashControlCustomEvent
func callbackQCameraFlashControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraFlashControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraFlashControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraFlashControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraFlashControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraFlashControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraFlashControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraFlashControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraFlashControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QCameraFocus struct {
	core.QObject
}

type QCameraFocus_ITF interface {
	core.QObject_ITF
	QCameraFocus_PTR() *QCameraFocus
}

func PointerFromQCameraFocus(ptr QCameraFocus_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraFocus_PTR().Pointer()
	}
	return nil
}

func NewQCameraFocusFromPointer(ptr unsafe.Pointer) *QCameraFocus {
	var n = new(QCameraFocus)
	n.SetPointer(ptr)
	return n
}

func newQCameraFocusFromPointer(ptr unsafe.Pointer) *QCameraFocus {
	var n = NewQCameraFocusFromPointer(ptr)
	for len(n.ObjectName()) < len("QCameraFocus_") {
		n.SetObjectName("QCameraFocus_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraFocus) QCameraFocus_PTR() *QCameraFocus {
	return ptr
}

//QCameraFocus::FocusMode
type QCameraFocus__FocusMode int64

const (
	QCameraFocus__ManualFocus     = QCameraFocus__FocusMode(0x1)
	QCameraFocus__HyperfocalFocus = QCameraFocus__FocusMode(0x02)
	QCameraFocus__InfinityFocus   = QCameraFocus__FocusMode(0x04)
	QCameraFocus__AutoFocus       = QCameraFocus__FocusMode(0x8)
	QCameraFocus__ContinuousFocus = QCameraFocus__FocusMode(0x10)
	QCameraFocus__MacroFocus      = QCameraFocus__FocusMode(0x20)
)

//QCameraFocus::FocusPointMode
type QCameraFocus__FocusPointMode int64

const (
	QCameraFocus__FocusPointAuto          = QCameraFocus__FocusPointMode(0)
	QCameraFocus__FocusPointCenter        = QCameraFocus__FocusPointMode(1)
	QCameraFocus__FocusPointFaceDetection = QCameraFocus__FocusPointMode(2)
	QCameraFocus__FocusPointCustom        = QCameraFocus__FocusPointMode(3)
)

func (ptr *QCameraFocus) DigitalZoom() float64 {
	defer qt.Recovering("QCameraFocus::digitalZoom")

	if ptr.Pointer() != nil {
		return float64(C.QCameraFocus_DigitalZoom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFocus) FocusMode() QCameraFocus__FocusMode {
	defer qt.Recovering("QCameraFocus::focusMode")

	if ptr.Pointer() != nil {
		return QCameraFocus__FocusMode(C.QCameraFocus_FocusMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFocus) FocusPointMode() QCameraFocus__FocusPointMode {
	defer qt.Recovering("QCameraFocus::focusPointMode")

	if ptr.Pointer() != nil {
		return QCameraFocus__FocusPointMode(C.QCameraFocus_FocusPointMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFocus) OpticalZoom() float64 {
	defer qt.Recovering("QCameraFocus::opticalZoom")

	if ptr.Pointer() != nil {
		return float64(C.QCameraFocus_OpticalZoom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFocus) SetCustomFocusPoint(point core.QPointF_ITF) {
	defer qt.Recovering("QCameraFocus::setCustomFocusPoint")

	if ptr.Pointer() != nil {
		C.QCameraFocus_SetCustomFocusPoint(ptr.Pointer(), core.PointerFromQPointF(point))
	}
}

func (ptr *QCameraFocus) SetFocusMode(mode QCameraFocus__FocusMode) {
	defer qt.Recovering("QCameraFocus::setFocusMode")

	if ptr.Pointer() != nil {
		C.QCameraFocus_SetFocusMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraFocus) SetFocusPointMode(mode QCameraFocus__FocusPointMode) {
	defer qt.Recovering("QCameraFocus::setFocusPointMode")

	if ptr.Pointer() != nil {
		C.QCameraFocus_SetFocusPointMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraFocus) ConnectDigitalZoomChanged(f func(value float64)) {
	defer qt.Recovering("connect QCameraFocus::digitalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocus_ConnectDigitalZoomChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "digitalZoomChanged", f)
	}
}

func (ptr *QCameraFocus) DisconnectDigitalZoomChanged() {
	defer qt.Recovering("disconnect QCameraFocus::digitalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocus_DisconnectDigitalZoomChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "digitalZoomChanged")
	}
}

//export callbackQCameraFocusDigitalZoomChanged
func callbackQCameraFocusDigitalZoomChanged(ptr unsafe.Pointer, ptrName *C.char, value C.double) {
	defer qt.Recovering("callback QCameraFocus::digitalZoomChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "digitalZoomChanged"); signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QCameraFocus) DigitalZoomChanged(value float64) {
	defer qt.Recovering("QCameraFocus::digitalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocus_DigitalZoomChanged(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QCameraFocus) ConnectFocusZonesChanged(f func()) {
	defer qt.Recovering("connect QCameraFocus::focusZonesChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocus_ConnectFocusZonesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "focusZonesChanged", f)
	}
}

func (ptr *QCameraFocus) DisconnectFocusZonesChanged() {
	defer qt.Recovering("disconnect QCameraFocus::focusZonesChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocus_DisconnectFocusZonesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "focusZonesChanged")
	}
}

//export callbackQCameraFocusFocusZonesChanged
func callbackQCameraFocusFocusZonesChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QCameraFocus::focusZonesChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusZonesChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCameraFocus) FocusZonesChanged() {
	defer qt.Recovering("QCameraFocus::focusZonesChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocus_FocusZonesChanged(ptr.Pointer())
	}
}

func (ptr *QCameraFocus) IsAvailable() bool {
	defer qt.Recovering("QCameraFocus::isAvailable")

	if ptr.Pointer() != nil {
		return C.QCameraFocus_IsAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCameraFocus) IsFocusModeSupported(mode QCameraFocus__FocusMode) bool {
	defer qt.Recovering("QCameraFocus::isFocusModeSupported")

	if ptr.Pointer() != nil {
		return C.QCameraFocus_IsFocusModeSupported(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraFocus) IsFocusPointModeSupported(mode QCameraFocus__FocusPointMode) bool {
	defer qt.Recovering("QCameraFocus::isFocusPointModeSupported")

	if ptr.Pointer() != nil {
		return C.QCameraFocus_IsFocusPointModeSupported(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraFocus) MaximumDigitalZoom() float64 {
	defer qt.Recovering("QCameraFocus::maximumDigitalZoom")

	if ptr.Pointer() != nil {
		return float64(C.QCameraFocus_MaximumDigitalZoom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFocus) ConnectMaximumDigitalZoomChanged(f func(zoom float64)) {
	defer qt.Recovering("connect QCameraFocus::maximumDigitalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocus_ConnectMaximumDigitalZoomChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "maximumDigitalZoomChanged", f)
	}
}

func (ptr *QCameraFocus) DisconnectMaximumDigitalZoomChanged() {
	defer qt.Recovering("disconnect QCameraFocus::maximumDigitalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocus_DisconnectMaximumDigitalZoomChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "maximumDigitalZoomChanged")
	}
}

//export callbackQCameraFocusMaximumDigitalZoomChanged
func callbackQCameraFocusMaximumDigitalZoomChanged(ptr unsafe.Pointer, ptrName *C.char, zoom C.double) {
	defer qt.Recovering("callback QCameraFocus::maximumDigitalZoomChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "maximumDigitalZoomChanged"); signal != nil {
		signal.(func(float64))(float64(zoom))
	}

}

func (ptr *QCameraFocus) MaximumDigitalZoomChanged(zoom float64) {
	defer qt.Recovering("QCameraFocus::maximumDigitalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocus_MaximumDigitalZoomChanged(ptr.Pointer(), C.double(zoom))
	}
}

func (ptr *QCameraFocus) MaximumOpticalZoom() float64 {
	defer qt.Recovering("QCameraFocus::maximumOpticalZoom")

	if ptr.Pointer() != nil {
		return float64(C.QCameraFocus_MaximumOpticalZoom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFocus) ConnectMaximumOpticalZoomChanged(f func(zoom float64)) {
	defer qt.Recovering("connect QCameraFocus::maximumOpticalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocus_ConnectMaximumOpticalZoomChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "maximumOpticalZoomChanged", f)
	}
}

func (ptr *QCameraFocus) DisconnectMaximumOpticalZoomChanged() {
	defer qt.Recovering("disconnect QCameraFocus::maximumOpticalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocus_DisconnectMaximumOpticalZoomChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "maximumOpticalZoomChanged")
	}
}

//export callbackQCameraFocusMaximumOpticalZoomChanged
func callbackQCameraFocusMaximumOpticalZoomChanged(ptr unsafe.Pointer, ptrName *C.char, zoom C.double) {
	defer qt.Recovering("callback QCameraFocus::maximumOpticalZoomChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "maximumOpticalZoomChanged"); signal != nil {
		signal.(func(float64))(float64(zoom))
	}

}

func (ptr *QCameraFocus) MaximumOpticalZoomChanged(zoom float64) {
	defer qt.Recovering("QCameraFocus::maximumOpticalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocus_MaximumOpticalZoomChanged(ptr.Pointer(), C.double(zoom))
	}
}

func (ptr *QCameraFocus) ConnectOpticalZoomChanged(f func(value float64)) {
	defer qt.Recovering("connect QCameraFocus::opticalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocus_ConnectOpticalZoomChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "opticalZoomChanged", f)
	}
}

func (ptr *QCameraFocus) DisconnectOpticalZoomChanged() {
	defer qt.Recovering("disconnect QCameraFocus::opticalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocus_DisconnectOpticalZoomChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "opticalZoomChanged")
	}
}

//export callbackQCameraFocusOpticalZoomChanged
func callbackQCameraFocusOpticalZoomChanged(ptr unsafe.Pointer, ptrName *C.char, value C.double) {
	defer qt.Recovering("callback QCameraFocus::opticalZoomChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "opticalZoomChanged"); signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QCameraFocus) OpticalZoomChanged(value float64) {
	defer qt.Recovering("QCameraFocus::opticalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocus_OpticalZoomChanged(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QCameraFocus) ZoomTo(optical float64, digital float64) {
	defer qt.Recovering("QCameraFocus::zoomTo")

	if ptr.Pointer() != nil {
		C.QCameraFocus_ZoomTo(ptr.Pointer(), C.double(optical), C.double(digital))
	}
}

func (ptr *QCameraFocus) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraFocus::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraFocus) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraFocus::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraFocusTimerEvent
func callbackQCameraFocusTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraFocus::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraFocusFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraFocus) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraFocus::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraFocus_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraFocus) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraFocus::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraFocus_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraFocus) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraFocus::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraFocus) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraFocus::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraFocusChildEvent
func callbackQCameraFocusChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraFocus::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraFocusFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraFocus) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraFocus::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraFocus_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraFocus) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraFocus::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraFocus_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraFocus) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraFocus::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraFocus) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraFocus::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraFocusCustomEvent
func callbackQCameraFocusCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraFocus::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraFocusFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraFocus) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraFocus::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraFocus_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraFocus) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraFocus::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraFocus_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QCameraFocusControl struct {
	QMediaControl
}

type QCameraFocusControl_ITF interface {
	QMediaControl_ITF
	QCameraFocusControl_PTR() *QCameraFocusControl
}

func PointerFromQCameraFocusControl(ptr QCameraFocusControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraFocusControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraFocusControlFromPointer(ptr unsafe.Pointer) *QCameraFocusControl {
	var n = new(QCameraFocusControl)
	n.SetPointer(ptr)
	return n
}

func newQCameraFocusControlFromPointer(ptr unsafe.Pointer) *QCameraFocusControl {
	var n = NewQCameraFocusControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QCameraFocusControl_") {
		n.SetObjectName("QCameraFocusControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraFocusControl) QCameraFocusControl_PTR() *QCameraFocusControl {
	return ptr
}

func (ptr *QCameraFocusControl) FocusMode() QCameraFocus__FocusMode {
	defer qt.Recovering("QCameraFocusControl::focusMode")

	if ptr.Pointer() != nil {
		return QCameraFocus__FocusMode(C.QCameraFocusControl_FocusMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFocusControl) ConnectFocusModeChanged(f func(mode QCameraFocus__FocusMode)) {
	defer qt.Recovering("connect QCameraFocusControl::focusModeChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_ConnectFocusModeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "focusModeChanged", f)
	}
}

func (ptr *QCameraFocusControl) DisconnectFocusModeChanged() {
	defer qt.Recovering("disconnect QCameraFocusControl::focusModeChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_DisconnectFocusModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "focusModeChanged")
	}
}

//export callbackQCameraFocusControlFocusModeChanged
func callbackQCameraFocusControlFocusModeChanged(ptr unsafe.Pointer, ptrName *C.char, mode C.int) {
	defer qt.Recovering("callback QCameraFocusControl::focusModeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusModeChanged"); signal != nil {
		signal.(func(QCameraFocus__FocusMode))(QCameraFocus__FocusMode(mode))
	}

}

func (ptr *QCameraFocusControl) FocusModeChanged(mode QCameraFocus__FocusMode) {
	defer qt.Recovering("QCameraFocusControl::focusModeChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_FocusModeChanged(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraFocusControl) FocusPointMode() QCameraFocus__FocusPointMode {
	defer qt.Recovering("QCameraFocusControl::focusPointMode")

	if ptr.Pointer() != nil {
		return QCameraFocus__FocusPointMode(C.QCameraFocusControl_FocusPointMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFocusControl) ConnectFocusPointModeChanged(f func(mode QCameraFocus__FocusPointMode)) {
	defer qt.Recovering("connect QCameraFocusControl::focusPointModeChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_ConnectFocusPointModeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "focusPointModeChanged", f)
	}
}

func (ptr *QCameraFocusControl) DisconnectFocusPointModeChanged() {
	defer qt.Recovering("disconnect QCameraFocusControl::focusPointModeChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_DisconnectFocusPointModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "focusPointModeChanged")
	}
}

//export callbackQCameraFocusControlFocusPointModeChanged
func callbackQCameraFocusControlFocusPointModeChanged(ptr unsafe.Pointer, ptrName *C.char, mode C.int) {
	defer qt.Recovering("callback QCameraFocusControl::focusPointModeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusPointModeChanged"); signal != nil {
		signal.(func(QCameraFocus__FocusPointMode))(QCameraFocus__FocusPointMode(mode))
	}

}

func (ptr *QCameraFocusControl) FocusPointModeChanged(mode QCameraFocus__FocusPointMode) {
	defer qt.Recovering("QCameraFocusControl::focusPointModeChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_FocusPointModeChanged(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraFocusControl) ConnectFocusZonesChanged(f func()) {
	defer qt.Recovering("connect QCameraFocusControl::focusZonesChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_ConnectFocusZonesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "focusZonesChanged", f)
	}
}

func (ptr *QCameraFocusControl) DisconnectFocusZonesChanged() {
	defer qt.Recovering("disconnect QCameraFocusControl::focusZonesChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_DisconnectFocusZonesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "focusZonesChanged")
	}
}

//export callbackQCameraFocusControlFocusZonesChanged
func callbackQCameraFocusControlFocusZonesChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QCameraFocusControl::focusZonesChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusZonesChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCameraFocusControl) FocusZonesChanged() {
	defer qt.Recovering("QCameraFocusControl::focusZonesChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_FocusZonesChanged(ptr.Pointer())
	}
}

func (ptr *QCameraFocusControl) IsFocusModeSupported(mode QCameraFocus__FocusMode) bool {
	defer qt.Recovering("QCameraFocusControl::isFocusModeSupported")

	if ptr.Pointer() != nil {
		return C.QCameraFocusControl_IsFocusModeSupported(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraFocusControl) IsFocusPointModeSupported(mode QCameraFocus__FocusPointMode) bool {
	defer qt.Recovering("QCameraFocusControl::isFocusPointModeSupported")

	if ptr.Pointer() != nil {
		return C.QCameraFocusControl_IsFocusPointModeSupported(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraFocusControl) SetCustomFocusPoint(point core.QPointF_ITF) {
	defer qt.Recovering("QCameraFocusControl::setCustomFocusPoint")

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_SetCustomFocusPoint(ptr.Pointer(), core.PointerFromQPointF(point))
	}
}

func (ptr *QCameraFocusControl) SetFocusMode(mode QCameraFocus__FocusMode) {
	defer qt.Recovering("QCameraFocusControl::setFocusMode")

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_SetFocusMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraFocusControl) SetFocusPointMode(mode QCameraFocus__FocusPointMode) {
	defer qt.Recovering("QCameraFocusControl::setFocusPointMode")

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_SetFocusPointMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraFocusControl) DestroyQCameraFocusControl() {
	defer qt.Recovering("QCameraFocusControl::~QCameraFocusControl")

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_DestroyQCameraFocusControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraFocusControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraFocusControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraFocusControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraFocusControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraFocusControlTimerEvent
func callbackQCameraFocusControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraFocusControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraFocusControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraFocusControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraFocusControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraFocusControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraFocusControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraFocusControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraFocusControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraFocusControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraFocusControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraFocusControlChildEvent
func callbackQCameraFocusControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraFocusControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraFocusControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraFocusControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraFocusControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraFocusControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraFocusControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraFocusControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraFocusControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraFocusControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraFocusControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraFocusControlCustomEvent
func callbackQCameraFocusControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraFocusControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraFocusControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraFocusControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraFocusControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraFocusControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraFocusControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QCameraFocusZone struct {
	ptr unsafe.Pointer
}

type QCameraFocusZone_ITF interface {
	QCameraFocusZone_PTR() *QCameraFocusZone
}

func (p *QCameraFocusZone) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QCameraFocusZone) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQCameraFocusZone(ptr QCameraFocusZone_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraFocusZone_PTR().Pointer()
	}
	return nil
}

func NewQCameraFocusZoneFromPointer(ptr unsafe.Pointer) *QCameraFocusZone {
	var n = new(QCameraFocusZone)
	n.SetPointer(ptr)
	return n
}

func newQCameraFocusZoneFromPointer(ptr unsafe.Pointer) *QCameraFocusZone {
	var n = NewQCameraFocusZoneFromPointer(ptr)
	return n
}

func (ptr *QCameraFocusZone) QCameraFocusZone_PTR() *QCameraFocusZone {
	return ptr
}

//QCameraFocusZone::FocusZoneStatus
type QCameraFocusZone__FocusZoneStatus int64

const (
	QCameraFocusZone__Invalid  = QCameraFocusZone__FocusZoneStatus(0)
	QCameraFocusZone__Unused   = QCameraFocusZone__FocusZoneStatus(1)
	QCameraFocusZone__Selected = QCameraFocusZone__FocusZoneStatus(2)
	QCameraFocusZone__Focused  = QCameraFocusZone__FocusZoneStatus(3)
)

func NewQCameraFocusZone(other QCameraFocusZone_ITF) *QCameraFocusZone {
	defer qt.Recovering("QCameraFocusZone::QCameraFocusZone")

	return newQCameraFocusZoneFromPointer(C.QCameraFocusZone_NewQCameraFocusZone(PointerFromQCameraFocusZone(other)))
}

func (ptr *QCameraFocusZone) IsValid() bool {
	defer qt.Recovering("QCameraFocusZone::isValid")

	if ptr.Pointer() != nil {
		return C.QCameraFocusZone_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCameraFocusZone) Status() QCameraFocusZone__FocusZoneStatus {
	defer qt.Recovering("QCameraFocusZone::status")

	if ptr.Pointer() != nil {
		return QCameraFocusZone__FocusZoneStatus(C.QCameraFocusZone_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFocusZone) DestroyQCameraFocusZone() {
	defer qt.Recovering("QCameraFocusZone::~QCameraFocusZone")

	if ptr.Pointer() != nil {
		C.QCameraFocusZone_DestroyQCameraFocusZone(ptr.Pointer())
	}
}

type QCameraImageCapture struct {
	core.QObject
	QMediaBindableInterface
}

type QCameraImageCapture_ITF interface {
	core.QObject_ITF
	QMediaBindableInterface_ITF
	QCameraImageCapture_PTR() *QCameraImageCapture
}

func (p *QCameraImageCapture) Pointer() unsafe.Pointer {
	return p.QObject_PTR().Pointer()
}

func (p *QCameraImageCapture) SetPointer(ptr unsafe.Pointer) {
	p.QObject_PTR().SetPointer(ptr)
	p.QMediaBindableInterface_PTR().SetPointer(ptr)
}

func PointerFromQCameraImageCapture(ptr QCameraImageCapture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraImageCapture_PTR().Pointer()
	}
	return nil
}

func NewQCameraImageCaptureFromPointer(ptr unsafe.Pointer) *QCameraImageCapture {
	var n = new(QCameraImageCapture)
	n.SetPointer(ptr)
	return n
}

func newQCameraImageCaptureFromPointer(ptr unsafe.Pointer) *QCameraImageCapture {
	var n = NewQCameraImageCaptureFromPointer(ptr)
	for len(n.ObjectName()) < len("QCameraImageCapture_") {
		n.SetObjectName("QCameraImageCapture_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraImageCapture) QCameraImageCapture_PTR() *QCameraImageCapture {
	return ptr
}

//QCameraImageCapture::CaptureDestination
type QCameraImageCapture__CaptureDestination int64

const (
	QCameraImageCapture__CaptureToFile   = QCameraImageCapture__CaptureDestination(0x01)
	QCameraImageCapture__CaptureToBuffer = QCameraImageCapture__CaptureDestination(0x02)
)

//QCameraImageCapture::DriveMode
type QCameraImageCapture__DriveMode int64

const (
	QCameraImageCapture__SingleImageCapture = QCameraImageCapture__DriveMode(0)
)

//QCameraImageCapture::Error
type QCameraImageCapture__Error int64

const (
	QCameraImageCapture__NoError                  = QCameraImageCapture__Error(0)
	QCameraImageCapture__NotReadyError            = QCameraImageCapture__Error(1)
	QCameraImageCapture__ResourceError            = QCameraImageCapture__Error(2)
	QCameraImageCapture__OutOfSpaceError          = QCameraImageCapture__Error(3)
	QCameraImageCapture__NotSupportedFeatureError = QCameraImageCapture__Error(4)
	QCameraImageCapture__FormatError              = QCameraImageCapture__Error(5)
)

func (ptr *QCameraImageCapture) IsReadyForCapture() bool {
	defer qt.Recovering("QCameraImageCapture::isReadyForCapture")

	if ptr.Pointer() != nil {
		return C.QCameraImageCapture_IsReadyForCapture(ptr.Pointer()) != 0
	}
	return false
}

func NewQCameraImageCapture(mediaObject QMediaObject_ITF, parent core.QObject_ITF) *QCameraImageCapture {
	defer qt.Recovering("QCameraImageCapture::QCameraImageCapture")

	return newQCameraImageCaptureFromPointer(C.QCameraImageCapture_NewQCameraImageCapture(PointerFromQMediaObject(mediaObject), core.PointerFromQObject(parent)))
}

func (ptr *QCameraImageCapture) Availability() QMultimedia__AvailabilityStatus {
	defer qt.Recovering("QCameraImageCapture::availability")

	if ptr.Pointer() != nil {
		return QMultimedia__AvailabilityStatus(C.QCameraImageCapture_Availability(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraImageCapture) BufferFormat() QVideoFrame__PixelFormat {
	defer qt.Recovering("QCameraImageCapture::bufferFormat")

	if ptr.Pointer() != nil {
		return QVideoFrame__PixelFormat(C.QCameraImageCapture_BufferFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraImageCapture) ConnectBufferFormatChanged(f func(format QVideoFrame__PixelFormat)) {
	defer qt.Recovering("connect QCameraImageCapture::bufferFormatChanged")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ConnectBufferFormatChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bufferFormatChanged", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectBufferFormatChanged() {
	defer qt.Recovering("disconnect QCameraImageCapture::bufferFormatChanged")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DisconnectBufferFormatChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bufferFormatChanged")
	}
}

//export callbackQCameraImageCaptureBufferFormatChanged
func callbackQCameraImageCaptureBufferFormatChanged(ptr unsafe.Pointer, ptrName *C.char, format C.int) {
	defer qt.Recovering("callback QCameraImageCapture::bufferFormatChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "bufferFormatChanged"); signal != nil {
		signal.(func(QVideoFrame__PixelFormat))(QVideoFrame__PixelFormat(format))
	}

}

func (ptr *QCameraImageCapture) BufferFormatChanged(format QVideoFrame__PixelFormat) {
	defer qt.Recovering("QCameraImageCapture::bufferFormatChanged")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_BufferFormatChanged(ptr.Pointer(), C.int(format))
	}
}

func (ptr *QCameraImageCapture) CancelCapture() {
	defer qt.Recovering("QCameraImageCapture::cancelCapture")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_CancelCapture(ptr.Pointer())
	}
}

func (ptr *QCameraImageCapture) Capture(file string) int {
	defer qt.Recovering("QCameraImageCapture::capture")

	if ptr.Pointer() != nil {
		return int(C.QCameraImageCapture_Capture(ptr.Pointer(), C.CString(file)))
	}
	return 0
}

func (ptr *QCameraImageCapture) CaptureDestination() QCameraImageCapture__CaptureDestination {
	defer qt.Recovering("QCameraImageCapture::captureDestination")

	if ptr.Pointer() != nil {
		return QCameraImageCapture__CaptureDestination(C.QCameraImageCapture_CaptureDestination(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraImageCapture) ConnectCaptureDestinationChanged(f func(destination QCameraImageCapture__CaptureDestination)) {
	defer qt.Recovering("connect QCameraImageCapture::captureDestinationChanged")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ConnectCaptureDestinationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "captureDestinationChanged", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectCaptureDestinationChanged() {
	defer qt.Recovering("disconnect QCameraImageCapture::captureDestinationChanged")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DisconnectCaptureDestinationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "captureDestinationChanged")
	}
}

//export callbackQCameraImageCaptureCaptureDestinationChanged
func callbackQCameraImageCaptureCaptureDestinationChanged(ptr unsafe.Pointer, ptrName *C.char, destination C.int) {
	defer qt.Recovering("callback QCameraImageCapture::captureDestinationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "captureDestinationChanged"); signal != nil {
		signal.(func(QCameraImageCapture__CaptureDestination))(QCameraImageCapture__CaptureDestination(destination))
	}

}

func (ptr *QCameraImageCapture) CaptureDestinationChanged(destination QCameraImageCapture__CaptureDestination) {
	defer qt.Recovering("QCameraImageCapture::captureDestinationChanged")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_CaptureDestinationChanged(ptr.Pointer(), C.int(destination))
	}
}

func (ptr *QCameraImageCapture) EncodingSettings() *QImageEncoderSettings {
	defer qt.Recovering("QCameraImageCapture::encodingSettings")

	if ptr.Pointer() != nil {
		return NewQImageEncoderSettingsFromPointer(C.QCameraImageCapture_EncodingSettings(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCameraImageCapture) ConnectError2(f func(id int, error QCameraImageCapture__Error, errorString string)) {
	defer qt.Recovering("connect QCameraImageCapture::error")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectError2() {
	defer qt.Recovering("disconnect QCameraImageCapture::error")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQCameraImageCaptureError2
func callbackQCameraImageCaptureError2(ptr unsafe.Pointer, ptrName *C.char, id C.int, error C.int, errorString *C.char) {
	defer qt.Recovering("callback QCameraImageCapture::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(int, QCameraImageCapture__Error, string))(int(id), QCameraImageCapture__Error(error), C.GoString(errorString))
	}

}

func (ptr *QCameraImageCapture) Error2(id int, error QCameraImageCapture__Error, errorString string) {
	defer qt.Recovering("QCameraImageCapture::error")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_Error2(ptr.Pointer(), C.int(id), C.int(error), C.CString(errorString))
	}
}

func (ptr *QCameraImageCapture) Error() QCameraImageCapture__Error {
	defer qt.Recovering("QCameraImageCapture::error")

	if ptr.Pointer() != nil {
		return QCameraImageCapture__Error(C.QCameraImageCapture_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraImageCapture) ErrorString() string {
	defer qt.Recovering("QCameraImageCapture::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QCameraImageCapture_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QCameraImageCapture) ConnectImageAvailable(f func(id int, buffer *QVideoFrame)) {
	defer qt.Recovering("connect QCameraImageCapture::imageAvailable")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ConnectImageAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "imageAvailable", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectImageAvailable() {
	defer qt.Recovering("disconnect QCameraImageCapture::imageAvailable")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DisconnectImageAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "imageAvailable")
	}
}

//export callbackQCameraImageCaptureImageAvailable
func callbackQCameraImageCaptureImageAvailable(ptr unsafe.Pointer, ptrName *C.char, id C.int, buffer unsafe.Pointer) {
	defer qt.Recovering("callback QCameraImageCapture::imageAvailable")

	if signal := qt.GetSignal(C.GoString(ptrName), "imageAvailable"); signal != nil {
		signal.(func(int, *QVideoFrame))(int(id), NewQVideoFrameFromPointer(buffer))
	}

}

func (ptr *QCameraImageCapture) ImageAvailable(id int, buffer QVideoFrame_ITF) {
	defer qt.Recovering("QCameraImageCapture::imageAvailable")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ImageAvailable(ptr.Pointer(), C.int(id), PointerFromQVideoFrame(buffer))
	}
}

func (ptr *QCameraImageCapture) ConnectImageCaptured(f func(id int, preview *gui.QImage)) {
	defer qt.Recovering("connect QCameraImageCapture::imageCaptured")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ConnectImageCaptured(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "imageCaptured", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectImageCaptured() {
	defer qt.Recovering("disconnect QCameraImageCapture::imageCaptured")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DisconnectImageCaptured(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "imageCaptured")
	}
}

//export callbackQCameraImageCaptureImageCaptured
func callbackQCameraImageCaptureImageCaptured(ptr unsafe.Pointer, ptrName *C.char, id C.int, preview unsafe.Pointer) {
	defer qt.Recovering("callback QCameraImageCapture::imageCaptured")

	if signal := qt.GetSignal(C.GoString(ptrName), "imageCaptured"); signal != nil {
		signal.(func(int, *gui.QImage))(int(id), gui.NewQImageFromPointer(preview))
	}

}

func (ptr *QCameraImageCapture) ImageCaptured(id int, preview gui.QImage_ITF) {
	defer qt.Recovering("QCameraImageCapture::imageCaptured")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ImageCaptured(ptr.Pointer(), C.int(id), gui.PointerFromQImage(preview))
	}
}

func (ptr *QCameraImageCapture) ImageCodecDescription(codec string) string {
	defer qt.Recovering("QCameraImageCapture::imageCodecDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QCameraImageCapture_ImageCodecDescription(ptr.Pointer(), C.CString(codec)))
	}
	return ""
}

func (ptr *QCameraImageCapture) ConnectImageExposed(f func(id int)) {
	defer qt.Recovering("connect QCameraImageCapture::imageExposed")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ConnectImageExposed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "imageExposed", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectImageExposed() {
	defer qt.Recovering("disconnect QCameraImageCapture::imageExposed")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DisconnectImageExposed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "imageExposed")
	}
}

//export callbackQCameraImageCaptureImageExposed
func callbackQCameraImageCaptureImageExposed(ptr unsafe.Pointer, ptrName *C.char, id C.int) {
	defer qt.Recovering("callback QCameraImageCapture::imageExposed")

	if signal := qt.GetSignal(C.GoString(ptrName), "imageExposed"); signal != nil {
		signal.(func(int))(int(id))
	}

}

func (ptr *QCameraImageCapture) ImageExposed(id int) {
	defer qt.Recovering("QCameraImageCapture::imageExposed")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ImageExposed(ptr.Pointer(), C.int(id))
	}
}

func (ptr *QCameraImageCapture) ConnectImageMetadataAvailable(f func(id int, key string, value *core.QVariant)) {
	defer qt.Recovering("connect QCameraImageCapture::imageMetadataAvailable")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ConnectImageMetadataAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "imageMetadataAvailable", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectImageMetadataAvailable() {
	defer qt.Recovering("disconnect QCameraImageCapture::imageMetadataAvailable")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DisconnectImageMetadataAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "imageMetadataAvailable")
	}
}

//export callbackQCameraImageCaptureImageMetadataAvailable
func callbackQCameraImageCaptureImageMetadataAvailable(ptr unsafe.Pointer, ptrName *C.char, id C.int, key *C.char, value unsafe.Pointer) {
	defer qt.Recovering("callback QCameraImageCapture::imageMetadataAvailable")

	if signal := qt.GetSignal(C.GoString(ptrName), "imageMetadataAvailable"); signal != nil {
		signal.(func(int, string, *core.QVariant))(int(id), C.GoString(key), core.NewQVariantFromPointer(value))
	}

}

func (ptr *QCameraImageCapture) ImageMetadataAvailable(id int, key string, value core.QVariant_ITF) {
	defer qt.Recovering("QCameraImageCapture::imageMetadataAvailable")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ImageMetadataAvailable(ptr.Pointer(), C.int(id), C.CString(key), core.PointerFromQVariant(value))
	}
}

func (ptr *QCameraImageCapture) ConnectImageSaved(f func(id int, fileName string)) {
	defer qt.Recovering("connect QCameraImageCapture::imageSaved")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ConnectImageSaved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "imageSaved", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectImageSaved() {
	defer qt.Recovering("disconnect QCameraImageCapture::imageSaved")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DisconnectImageSaved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "imageSaved")
	}
}

//export callbackQCameraImageCaptureImageSaved
func callbackQCameraImageCaptureImageSaved(ptr unsafe.Pointer, ptrName *C.char, id C.int, fileName *C.char) {
	defer qt.Recovering("callback QCameraImageCapture::imageSaved")

	if signal := qt.GetSignal(C.GoString(ptrName), "imageSaved"); signal != nil {
		signal.(func(int, string))(int(id), C.GoString(fileName))
	}

}

func (ptr *QCameraImageCapture) ImageSaved(id int, fileName string) {
	defer qt.Recovering("QCameraImageCapture::imageSaved")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ImageSaved(ptr.Pointer(), C.int(id), C.CString(fileName))
	}
}

func (ptr *QCameraImageCapture) IsAvailable() bool {
	defer qt.Recovering("QCameraImageCapture::isAvailable")

	if ptr.Pointer() != nil {
		return C.QCameraImageCapture_IsAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCameraImageCapture) IsCaptureDestinationSupported(destination QCameraImageCapture__CaptureDestination) bool {
	defer qt.Recovering("QCameraImageCapture::isCaptureDestinationSupported")

	if ptr.Pointer() != nil {
		return C.QCameraImageCapture_IsCaptureDestinationSupported(ptr.Pointer(), C.int(destination)) != 0
	}
	return false
}

func (ptr *QCameraImageCapture) MediaObject() *QMediaObject {
	defer qt.Recovering("QCameraImageCapture::mediaObject")

	if ptr.Pointer() != nil {
		return NewQMediaObjectFromPointer(C.QCameraImageCapture_MediaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCameraImageCapture) ConnectReadyForCaptureChanged(f func(ready bool)) {
	defer qt.Recovering("connect QCameraImageCapture::readyForCaptureChanged")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ConnectReadyForCaptureChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "readyForCaptureChanged", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectReadyForCaptureChanged() {
	defer qt.Recovering("disconnect QCameraImageCapture::readyForCaptureChanged")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DisconnectReadyForCaptureChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "readyForCaptureChanged")
	}
}

//export callbackQCameraImageCaptureReadyForCaptureChanged
func callbackQCameraImageCaptureReadyForCaptureChanged(ptr unsafe.Pointer, ptrName *C.char, ready C.int) {
	defer qt.Recovering("callback QCameraImageCapture::readyForCaptureChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "readyForCaptureChanged"); signal != nil {
		signal.(func(bool))(int(ready) != 0)
	}

}

func (ptr *QCameraImageCapture) ReadyForCaptureChanged(ready bool) {
	defer qt.Recovering("QCameraImageCapture::readyForCaptureChanged")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ReadyForCaptureChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(ready)))
	}
}

func (ptr *QCameraImageCapture) SetBufferFormat(format QVideoFrame__PixelFormat) {
	defer qt.Recovering("QCameraImageCapture::setBufferFormat")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_SetBufferFormat(ptr.Pointer(), C.int(format))
	}
}

func (ptr *QCameraImageCapture) SetCaptureDestination(destination QCameraImageCapture__CaptureDestination) {
	defer qt.Recovering("QCameraImageCapture::setCaptureDestination")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_SetCaptureDestination(ptr.Pointer(), C.int(destination))
	}
}

func (ptr *QCameraImageCapture) SetEncodingSettings(settings QImageEncoderSettings_ITF) {
	defer qt.Recovering("QCameraImageCapture::setEncodingSettings")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_SetEncodingSettings(ptr.Pointer(), PointerFromQImageEncoderSettings(settings))
	}
}

func (ptr *QCameraImageCapture) SetMediaObject(mediaObject QMediaObject_ITF) bool {
	defer qt.Recovering("QCameraImageCapture::setMediaObject")

	if ptr.Pointer() != nil {
		return C.QCameraImageCapture_SetMediaObject(ptr.Pointer(), PointerFromQMediaObject(mediaObject)) != 0
	}
	return false
}

func (ptr *QCameraImageCapture) SupportedImageCodecs() []string {
	defer qt.Recovering("QCameraImageCapture::supportedImageCodecs")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QCameraImageCapture_SupportedImageCodecs(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QCameraImageCapture) DestroyQCameraImageCapture() {
	defer qt.Recovering("QCameraImageCapture::~QCameraImageCapture")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DestroyQCameraImageCapture(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraImageCapture) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraImageCapture::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraImageCapture::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraImageCaptureTimerEvent
func callbackQCameraImageCaptureTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraImageCapture::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraImageCaptureFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraImageCapture) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraImageCapture::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraImageCapture) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraImageCapture::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraImageCapture) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraImageCapture::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraImageCapture::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraImageCaptureChildEvent
func callbackQCameraImageCaptureChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraImageCapture::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraImageCaptureFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraImageCapture) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraImageCapture::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraImageCapture) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraImageCapture::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraImageCapture) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraImageCapture::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraImageCapture::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraImageCaptureCustomEvent
func callbackQCameraImageCaptureCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraImageCapture::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraImageCaptureFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraImageCapture) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraImageCapture::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraImageCapture) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraImageCapture::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QCameraImageCaptureControl struct {
	QMediaControl
}

type QCameraImageCaptureControl_ITF interface {
	QMediaControl_ITF
	QCameraImageCaptureControl_PTR() *QCameraImageCaptureControl
}

func PointerFromQCameraImageCaptureControl(ptr QCameraImageCaptureControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraImageCaptureControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraImageCaptureControlFromPointer(ptr unsafe.Pointer) *QCameraImageCaptureControl {
	var n = new(QCameraImageCaptureControl)
	n.SetPointer(ptr)
	return n
}

func newQCameraImageCaptureControlFromPointer(ptr unsafe.Pointer) *QCameraImageCaptureControl {
	var n = NewQCameraImageCaptureControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QCameraImageCaptureControl_") {
		n.SetObjectName("QCameraImageCaptureControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraImageCaptureControl) QCameraImageCaptureControl_PTR() *QCameraImageCaptureControl {
	return ptr
}

func (ptr *QCameraImageCaptureControl) CancelCapture() {
	defer qt.Recovering("QCameraImageCaptureControl::cancelCapture")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_CancelCapture(ptr.Pointer())
	}
}

func (ptr *QCameraImageCaptureControl) Capture(fileName string) int {
	defer qt.Recovering("QCameraImageCaptureControl::capture")

	if ptr.Pointer() != nil {
		return int(C.QCameraImageCaptureControl_Capture(ptr.Pointer(), C.CString(fileName)))
	}
	return 0
}

func (ptr *QCameraImageCaptureControl) DriveMode() QCameraImageCapture__DriveMode {
	defer qt.Recovering("QCameraImageCaptureControl::driveMode")

	if ptr.Pointer() != nil {
		return QCameraImageCapture__DriveMode(C.QCameraImageCaptureControl_DriveMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraImageCaptureControl) ConnectError(f func(id int, error int, errorString string)) {
	defer qt.Recovering("connect QCameraImageCaptureControl::error")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QCameraImageCaptureControl) DisconnectError() {
	defer qt.Recovering("disconnect QCameraImageCaptureControl::error")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQCameraImageCaptureControlError
func callbackQCameraImageCaptureControlError(ptr unsafe.Pointer, ptrName *C.char, id C.int, error C.int, errorString *C.char) {
	defer qt.Recovering("callback QCameraImageCaptureControl::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error"); signal != nil {
		signal.(func(int, int, string))(int(id), int(error), C.GoString(errorString))
	}

}

func (ptr *QCameraImageCaptureControl) Error(id int, error int, errorString string) {
	defer qt.Recovering("QCameraImageCaptureControl::error")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_Error(ptr.Pointer(), C.int(id), C.int(error), C.CString(errorString))
	}
}

func (ptr *QCameraImageCaptureControl) ConnectImageAvailable(f func(requestId int, buffer *QVideoFrame)) {
	defer qt.Recovering("connect QCameraImageCaptureControl::imageAvailable")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ConnectImageAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "imageAvailable", f)
	}
}

func (ptr *QCameraImageCaptureControl) DisconnectImageAvailable() {
	defer qt.Recovering("disconnect QCameraImageCaptureControl::imageAvailable")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_DisconnectImageAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "imageAvailable")
	}
}

//export callbackQCameraImageCaptureControlImageAvailable
func callbackQCameraImageCaptureControlImageAvailable(ptr unsafe.Pointer, ptrName *C.char, requestId C.int, buffer unsafe.Pointer) {
	defer qt.Recovering("callback QCameraImageCaptureControl::imageAvailable")

	if signal := qt.GetSignal(C.GoString(ptrName), "imageAvailable"); signal != nil {
		signal.(func(int, *QVideoFrame))(int(requestId), NewQVideoFrameFromPointer(buffer))
	}

}

func (ptr *QCameraImageCaptureControl) ImageAvailable(requestId int, buffer QVideoFrame_ITF) {
	defer qt.Recovering("QCameraImageCaptureControl::imageAvailable")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ImageAvailable(ptr.Pointer(), C.int(requestId), PointerFromQVideoFrame(buffer))
	}
}

func (ptr *QCameraImageCaptureControl) ConnectImageCaptured(f func(requestId int, preview *gui.QImage)) {
	defer qt.Recovering("connect QCameraImageCaptureControl::imageCaptured")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ConnectImageCaptured(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "imageCaptured", f)
	}
}

func (ptr *QCameraImageCaptureControl) DisconnectImageCaptured() {
	defer qt.Recovering("disconnect QCameraImageCaptureControl::imageCaptured")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_DisconnectImageCaptured(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "imageCaptured")
	}
}

//export callbackQCameraImageCaptureControlImageCaptured
func callbackQCameraImageCaptureControlImageCaptured(ptr unsafe.Pointer, ptrName *C.char, requestId C.int, preview unsafe.Pointer) {
	defer qt.Recovering("callback QCameraImageCaptureControl::imageCaptured")

	if signal := qt.GetSignal(C.GoString(ptrName), "imageCaptured"); signal != nil {
		signal.(func(int, *gui.QImage))(int(requestId), gui.NewQImageFromPointer(preview))
	}

}

func (ptr *QCameraImageCaptureControl) ImageCaptured(requestId int, preview gui.QImage_ITF) {
	defer qt.Recovering("QCameraImageCaptureControl::imageCaptured")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ImageCaptured(ptr.Pointer(), C.int(requestId), gui.PointerFromQImage(preview))
	}
}

func (ptr *QCameraImageCaptureControl) ConnectImageExposed(f func(requestId int)) {
	defer qt.Recovering("connect QCameraImageCaptureControl::imageExposed")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ConnectImageExposed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "imageExposed", f)
	}
}

func (ptr *QCameraImageCaptureControl) DisconnectImageExposed() {
	defer qt.Recovering("disconnect QCameraImageCaptureControl::imageExposed")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_DisconnectImageExposed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "imageExposed")
	}
}

//export callbackQCameraImageCaptureControlImageExposed
func callbackQCameraImageCaptureControlImageExposed(ptr unsafe.Pointer, ptrName *C.char, requestId C.int) {
	defer qt.Recovering("callback QCameraImageCaptureControl::imageExposed")

	if signal := qt.GetSignal(C.GoString(ptrName), "imageExposed"); signal != nil {
		signal.(func(int))(int(requestId))
	}

}

func (ptr *QCameraImageCaptureControl) ImageExposed(requestId int) {
	defer qt.Recovering("QCameraImageCaptureControl::imageExposed")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ImageExposed(ptr.Pointer(), C.int(requestId))
	}
}

func (ptr *QCameraImageCaptureControl) ConnectImageMetadataAvailable(f func(id int, key string, value *core.QVariant)) {
	defer qt.Recovering("connect QCameraImageCaptureControl::imageMetadataAvailable")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ConnectImageMetadataAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "imageMetadataAvailable", f)
	}
}

func (ptr *QCameraImageCaptureControl) DisconnectImageMetadataAvailable() {
	defer qt.Recovering("disconnect QCameraImageCaptureControl::imageMetadataAvailable")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_DisconnectImageMetadataAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "imageMetadataAvailable")
	}
}

//export callbackQCameraImageCaptureControlImageMetadataAvailable
func callbackQCameraImageCaptureControlImageMetadataAvailable(ptr unsafe.Pointer, ptrName *C.char, id C.int, key *C.char, value unsafe.Pointer) {
	defer qt.Recovering("callback QCameraImageCaptureControl::imageMetadataAvailable")

	if signal := qt.GetSignal(C.GoString(ptrName), "imageMetadataAvailable"); signal != nil {
		signal.(func(int, string, *core.QVariant))(int(id), C.GoString(key), core.NewQVariantFromPointer(value))
	}

}

func (ptr *QCameraImageCaptureControl) ImageMetadataAvailable(id int, key string, value core.QVariant_ITF) {
	defer qt.Recovering("QCameraImageCaptureControl::imageMetadataAvailable")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ImageMetadataAvailable(ptr.Pointer(), C.int(id), C.CString(key), core.PointerFromQVariant(value))
	}
}

func (ptr *QCameraImageCaptureControl) ConnectImageSaved(f func(requestId int, fileName string)) {
	defer qt.Recovering("connect QCameraImageCaptureControl::imageSaved")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ConnectImageSaved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "imageSaved", f)
	}
}

func (ptr *QCameraImageCaptureControl) DisconnectImageSaved() {
	defer qt.Recovering("disconnect QCameraImageCaptureControl::imageSaved")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_DisconnectImageSaved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "imageSaved")
	}
}

//export callbackQCameraImageCaptureControlImageSaved
func callbackQCameraImageCaptureControlImageSaved(ptr unsafe.Pointer, ptrName *C.char, requestId C.int, fileName *C.char) {
	defer qt.Recovering("callback QCameraImageCaptureControl::imageSaved")

	if signal := qt.GetSignal(C.GoString(ptrName), "imageSaved"); signal != nil {
		signal.(func(int, string))(int(requestId), C.GoString(fileName))
	}

}

func (ptr *QCameraImageCaptureControl) ImageSaved(requestId int, fileName string) {
	defer qt.Recovering("QCameraImageCaptureControl::imageSaved")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ImageSaved(ptr.Pointer(), C.int(requestId), C.CString(fileName))
	}
}

func (ptr *QCameraImageCaptureControl) IsReadyForCapture() bool {
	defer qt.Recovering("QCameraImageCaptureControl::isReadyForCapture")

	if ptr.Pointer() != nil {
		return C.QCameraImageCaptureControl_IsReadyForCapture(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCameraImageCaptureControl) ConnectReadyForCaptureChanged(f func(ready bool)) {
	defer qt.Recovering("connect QCameraImageCaptureControl::readyForCaptureChanged")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ConnectReadyForCaptureChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "readyForCaptureChanged", f)
	}
}

func (ptr *QCameraImageCaptureControl) DisconnectReadyForCaptureChanged() {
	defer qt.Recovering("disconnect QCameraImageCaptureControl::readyForCaptureChanged")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_DisconnectReadyForCaptureChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "readyForCaptureChanged")
	}
}

//export callbackQCameraImageCaptureControlReadyForCaptureChanged
func callbackQCameraImageCaptureControlReadyForCaptureChanged(ptr unsafe.Pointer, ptrName *C.char, ready C.int) {
	defer qt.Recovering("callback QCameraImageCaptureControl::readyForCaptureChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "readyForCaptureChanged"); signal != nil {
		signal.(func(bool))(int(ready) != 0)
	}

}

func (ptr *QCameraImageCaptureControl) ReadyForCaptureChanged(ready bool) {
	defer qt.Recovering("QCameraImageCaptureControl::readyForCaptureChanged")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ReadyForCaptureChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(ready)))
	}
}

func (ptr *QCameraImageCaptureControl) SetDriveMode(mode QCameraImageCapture__DriveMode) {
	defer qt.Recovering("QCameraImageCaptureControl::setDriveMode")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_SetDriveMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraImageCaptureControl) DestroyQCameraImageCaptureControl() {
	defer qt.Recovering("QCameraImageCaptureControl::~QCameraImageCaptureControl")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_DestroyQCameraImageCaptureControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraImageCaptureControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraImageCaptureControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraImageCaptureControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraImageCaptureControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraImageCaptureControlTimerEvent
func callbackQCameraImageCaptureControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraImageCaptureControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraImageCaptureControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraImageCaptureControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraImageCaptureControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraImageCaptureControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraImageCaptureControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraImageCaptureControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraImageCaptureControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraImageCaptureControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraImageCaptureControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraImageCaptureControlChildEvent
func callbackQCameraImageCaptureControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraImageCaptureControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraImageCaptureControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraImageCaptureControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraImageCaptureControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraImageCaptureControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraImageCaptureControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraImageCaptureControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraImageCaptureControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraImageCaptureControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraImageCaptureControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraImageCaptureControlCustomEvent
func callbackQCameraImageCaptureControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraImageCaptureControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraImageCaptureControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraImageCaptureControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraImageCaptureControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraImageCaptureControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraImageCaptureControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QCameraImageProcessing struct {
	core.QObject
}

type QCameraImageProcessing_ITF interface {
	core.QObject_ITF
	QCameraImageProcessing_PTR() *QCameraImageProcessing
}

func PointerFromQCameraImageProcessing(ptr QCameraImageProcessing_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraImageProcessing_PTR().Pointer()
	}
	return nil
}

func NewQCameraImageProcessingFromPointer(ptr unsafe.Pointer) *QCameraImageProcessing {
	var n = new(QCameraImageProcessing)
	n.SetPointer(ptr)
	return n
}

func newQCameraImageProcessingFromPointer(ptr unsafe.Pointer) *QCameraImageProcessing {
	var n = NewQCameraImageProcessingFromPointer(ptr)
	for len(n.ObjectName()) < len("QCameraImageProcessing_") {
		n.SetObjectName("QCameraImageProcessing_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraImageProcessing) QCameraImageProcessing_PTR() *QCameraImageProcessing {
	return ptr
}

//QCameraImageProcessing::ColorFilter
type QCameraImageProcessing__ColorFilter int64

const (
	QCameraImageProcessing__ColorFilterNone       = QCameraImageProcessing__ColorFilter(0)
	QCameraImageProcessing__ColorFilterGrayscale  = QCameraImageProcessing__ColorFilter(1)
	QCameraImageProcessing__ColorFilterNegative   = QCameraImageProcessing__ColorFilter(2)
	QCameraImageProcessing__ColorFilterSolarize   = QCameraImageProcessing__ColorFilter(3)
	QCameraImageProcessing__ColorFilterSepia      = QCameraImageProcessing__ColorFilter(4)
	QCameraImageProcessing__ColorFilterPosterize  = QCameraImageProcessing__ColorFilter(5)
	QCameraImageProcessing__ColorFilterWhiteboard = QCameraImageProcessing__ColorFilter(6)
	QCameraImageProcessing__ColorFilterBlackboard = QCameraImageProcessing__ColorFilter(7)
	QCameraImageProcessing__ColorFilterAqua       = QCameraImageProcessing__ColorFilter(8)
	QCameraImageProcessing__ColorFilterVendor     = QCameraImageProcessing__ColorFilter(1000)
)

//QCameraImageProcessing::WhiteBalanceMode
type QCameraImageProcessing__WhiteBalanceMode int64

const (
	QCameraImageProcessing__WhiteBalanceAuto        = QCameraImageProcessing__WhiteBalanceMode(0)
	QCameraImageProcessing__WhiteBalanceManual      = QCameraImageProcessing__WhiteBalanceMode(1)
	QCameraImageProcessing__WhiteBalanceSunlight    = QCameraImageProcessing__WhiteBalanceMode(2)
	QCameraImageProcessing__WhiteBalanceCloudy      = QCameraImageProcessing__WhiteBalanceMode(3)
	QCameraImageProcessing__WhiteBalanceShade       = QCameraImageProcessing__WhiteBalanceMode(4)
	QCameraImageProcessing__WhiteBalanceTungsten    = QCameraImageProcessing__WhiteBalanceMode(5)
	QCameraImageProcessing__WhiteBalanceFluorescent = QCameraImageProcessing__WhiteBalanceMode(6)
	QCameraImageProcessing__WhiteBalanceFlash       = QCameraImageProcessing__WhiteBalanceMode(7)
	QCameraImageProcessing__WhiteBalanceSunset      = QCameraImageProcessing__WhiteBalanceMode(8)
	QCameraImageProcessing__WhiteBalanceVendor      = QCameraImageProcessing__WhiteBalanceMode(1000)
)

func (ptr *QCameraImageProcessing) ColorFilter() QCameraImageProcessing__ColorFilter {
	defer qt.Recovering("QCameraImageProcessing::colorFilter")

	if ptr.Pointer() != nil {
		return QCameraImageProcessing__ColorFilter(C.QCameraImageProcessing_ColorFilter(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraImageProcessing) Contrast() float64 {
	defer qt.Recovering("QCameraImageProcessing::contrast")

	if ptr.Pointer() != nil {
		return float64(C.QCameraImageProcessing_Contrast(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraImageProcessing) DenoisingLevel() float64 {
	defer qt.Recovering("QCameraImageProcessing::denoisingLevel")

	if ptr.Pointer() != nil {
		return float64(C.QCameraImageProcessing_DenoisingLevel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraImageProcessing) IsAvailable() bool {
	defer qt.Recovering("QCameraImageProcessing::isAvailable")

	if ptr.Pointer() != nil {
		return C.QCameraImageProcessing_IsAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCameraImageProcessing) IsColorFilterSupported(filter QCameraImageProcessing__ColorFilter) bool {
	defer qt.Recovering("QCameraImageProcessing::isColorFilterSupported")

	if ptr.Pointer() != nil {
		return C.QCameraImageProcessing_IsColorFilterSupported(ptr.Pointer(), C.int(filter)) != 0
	}
	return false
}

func (ptr *QCameraImageProcessing) IsWhiteBalanceModeSupported(mode QCameraImageProcessing__WhiteBalanceMode) bool {
	defer qt.Recovering("QCameraImageProcessing::isWhiteBalanceModeSupported")

	if ptr.Pointer() != nil {
		return C.QCameraImageProcessing_IsWhiteBalanceModeSupported(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraImageProcessing) ManualWhiteBalance() float64 {
	defer qt.Recovering("QCameraImageProcessing::manualWhiteBalance")

	if ptr.Pointer() != nil {
		return float64(C.QCameraImageProcessing_ManualWhiteBalance(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraImageProcessing) Saturation() float64 {
	defer qt.Recovering("QCameraImageProcessing::saturation")

	if ptr.Pointer() != nil {
		return float64(C.QCameraImageProcessing_Saturation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraImageProcessing) SetColorFilter(filter QCameraImageProcessing__ColorFilter) {
	defer qt.Recovering("QCameraImageProcessing::setColorFilter")

	if ptr.Pointer() != nil {
		C.QCameraImageProcessing_SetColorFilter(ptr.Pointer(), C.int(filter))
	}
}

func (ptr *QCameraImageProcessing) SetContrast(value float64) {
	defer qt.Recovering("QCameraImageProcessing::setContrast")

	if ptr.Pointer() != nil {
		C.QCameraImageProcessing_SetContrast(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QCameraImageProcessing) SetDenoisingLevel(level float64) {
	defer qt.Recovering("QCameraImageProcessing::setDenoisingLevel")

	if ptr.Pointer() != nil {
		C.QCameraImageProcessing_SetDenoisingLevel(ptr.Pointer(), C.double(level))
	}
}

func (ptr *QCameraImageProcessing) SetManualWhiteBalance(colorTemperature float64) {
	defer qt.Recovering("QCameraImageProcessing::setManualWhiteBalance")

	if ptr.Pointer() != nil {
		C.QCameraImageProcessing_SetManualWhiteBalance(ptr.Pointer(), C.double(colorTemperature))
	}
}

func (ptr *QCameraImageProcessing) SetSaturation(value float64) {
	defer qt.Recovering("QCameraImageProcessing::setSaturation")

	if ptr.Pointer() != nil {
		C.QCameraImageProcessing_SetSaturation(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QCameraImageProcessing) SetSharpeningLevel(level float64) {
	defer qt.Recovering("QCameraImageProcessing::setSharpeningLevel")

	if ptr.Pointer() != nil {
		C.QCameraImageProcessing_SetSharpeningLevel(ptr.Pointer(), C.double(level))
	}
}

func (ptr *QCameraImageProcessing) SetWhiteBalanceMode(mode QCameraImageProcessing__WhiteBalanceMode) {
	defer qt.Recovering("QCameraImageProcessing::setWhiteBalanceMode")

	if ptr.Pointer() != nil {
		C.QCameraImageProcessing_SetWhiteBalanceMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraImageProcessing) SharpeningLevel() float64 {
	defer qt.Recovering("QCameraImageProcessing::sharpeningLevel")

	if ptr.Pointer() != nil {
		return float64(C.QCameraImageProcessing_SharpeningLevel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraImageProcessing) WhiteBalanceMode() QCameraImageProcessing__WhiteBalanceMode {
	defer qt.Recovering("QCameraImageProcessing::whiteBalanceMode")

	if ptr.Pointer() != nil {
		return QCameraImageProcessing__WhiteBalanceMode(C.QCameraImageProcessing_WhiteBalanceMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraImageProcessing) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraImageProcessing::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraImageProcessing) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraImageProcessing::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraImageProcessingTimerEvent
func callbackQCameraImageProcessingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraImageProcessing::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraImageProcessingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraImageProcessing) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraImageProcessing::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageProcessing_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraImageProcessing) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraImageProcessing::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageProcessing_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraImageProcessing) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraImageProcessing::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraImageProcessing) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraImageProcessing::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraImageProcessingChildEvent
func callbackQCameraImageProcessingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraImageProcessing::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraImageProcessingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraImageProcessing) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraImageProcessing::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageProcessing_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraImageProcessing) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraImageProcessing::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageProcessing_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraImageProcessing) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraImageProcessing::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraImageProcessing) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraImageProcessing::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraImageProcessingCustomEvent
func callbackQCameraImageProcessingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraImageProcessing::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraImageProcessingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraImageProcessing) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraImageProcessing::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageProcessing_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraImageProcessing) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraImageProcessing::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageProcessing_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QCameraImageProcessingControl struct {
	QMediaControl
}

type QCameraImageProcessingControl_ITF interface {
	QMediaControl_ITF
	QCameraImageProcessingControl_PTR() *QCameraImageProcessingControl
}

func PointerFromQCameraImageProcessingControl(ptr QCameraImageProcessingControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraImageProcessingControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraImageProcessingControlFromPointer(ptr unsafe.Pointer) *QCameraImageProcessingControl {
	var n = new(QCameraImageProcessingControl)
	n.SetPointer(ptr)
	return n
}

func newQCameraImageProcessingControlFromPointer(ptr unsafe.Pointer) *QCameraImageProcessingControl {
	var n = NewQCameraImageProcessingControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QCameraImageProcessingControl_") {
		n.SetObjectName("QCameraImageProcessingControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraImageProcessingControl) QCameraImageProcessingControl_PTR() *QCameraImageProcessingControl {
	return ptr
}

//QCameraImageProcessingControl::ProcessingParameter
type QCameraImageProcessingControl__ProcessingParameter int64

const (
	QCameraImageProcessingControl__WhiteBalancePreset   = QCameraImageProcessingControl__ProcessingParameter(0)
	QCameraImageProcessingControl__ColorTemperature     = QCameraImageProcessingControl__ProcessingParameter(1)
	QCameraImageProcessingControl__Contrast             = QCameraImageProcessingControl__ProcessingParameter(2)
	QCameraImageProcessingControl__Saturation           = QCameraImageProcessingControl__ProcessingParameter(3)
	QCameraImageProcessingControl__Brightness           = QCameraImageProcessingControl__ProcessingParameter(4)
	QCameraImageProcessingControl__Sharpening           = QCameraImageProcessingControl__ProcessingParameter(5)
	QCameraImageProcessingControl__Denoising            = QCameraImageProcessingControl__ProcessingParameter(6)
	QCameraImageProcessingControl__ContrastAdjustment   = QCameraImageProcessingControl__ProcessingParameter(7)
	QCameraImageProcessingControl__SaturationAdjustment = QCameraImageProcessingControl__ProcessingParameter(8)
	QCameraImageProcessingControl__BrightnessAdjustment = QCameraImageProcessingControl__ProcessingParameter(9)
	QCameraImageProcessingControl__SharpeningAdjustment = QCameraImageProcessingControl__ProcessingParameter(10)
	QCameraImageProcessingControl__DenoisingAdjustment  = QCameraImageProcessingControl__ProcessingParameter(11)
	QCameraImageProcessingControl__ColorFilter          = QCameraImageProcessingControl__ProcessingParameter(12)
	QCameraImageProcessingControl__ExtendedParameter    = QCameraImageProcessingControl__ProcessingParameter(1000)
)

func (ptr *QCameraImageProcessingControl) IsParameterSupported(parameter QCameraImageProcessingControl__ProcessingParameter) bool {
	defer qt.Recovering("QCameraImageProcessingControl::isParameterSupported")

	if ptr.Pointer() != nil {
		return C.QCameraImageProcessingControl_IsParameterSupported(ptr.Pointer(), C.int(parameter)) != 0
	}
	return false
}

func (ptr *QCameraImageProcessingControl) IsParameterValueSupported(parameter QCameraImageProcessingControl__ProcessingParameter, value core.QVariant_ITF) bool {
	defer qt.Recovering("QCameraImageProcessingControl::isParameterValueSupported")

	if ptr.Pointer() != nil {
		return C.QCameraImageProcessingControl_IsParameterValueSupported(ptr.Pointer(), C.int(parameter), core.PointerFromQVariant(value)) != 0
	}
	return false
}

func (ptr *QCameraImageProcessingControl) Parameter(parameter QCameraImageProcessingControl__ProcessingParameter) *core.QVariant {
	defer qt.Recovering("QCameraImageProcessingControl::parameter")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QCameraImageProcessingControl_Parameter(ptr.Pointer(), C.int(parameter)))
	}
	return nil
}

func (ptr *QCameraImageProcessingControl) SetParameter(parameter QCameraImageProcessingControl__ProcessingParameter, value core.QVariant_ITF) {
	defer qt.Recovering("QCameraImageProcessingControl::setParameter")

	if ptr.Pointer() != nil {
		C.QCameraImageProcessingControl_SetParameter(ptr.Pointer(), C.int(parameter), core.PointerFromQVariant(value))
	}
}

func (ptr *QCameraImageProcessingControl) DestroyQCameraImageProcessingControl() {
	defer qt.Recovering("QCameraImageProcessingControl::~QCameraImageProcessingControl")

	if ptr.Pointer() != nil {
		C.QCameraImageProcessingControl_DestroyQCameraImageProcessingControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraImageProcessingControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraImageProcessingControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraImageProcessingControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraImageProcessingControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraImageProcessingControlTimerEvent
func callbackQCameraImageProcessingControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraImageProcessingControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraImageProcessingControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraImageProcessingControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraImageProcessingControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageProcessingControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraImageProcessingControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraImageProcessingControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageProcessingControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraImageProcessingControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraImageProcessingControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraImageProcessingControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraImageProcessingControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraImageProcessingControlChildEvent
func callbackQCameraImageProcessingControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraImageProcessingControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraImageProcessingControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraImageProcessingControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraImageProcessingControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageProcessingControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraImageProcessingControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraImageProcessingControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageProcessingControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraImageProcessingControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraImageProcessingControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraImageProcessingControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraImageProcessingControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraImageProcessingControlCustomEvent
func callbackQCameraImageProcessingControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraImageProcessingControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraImageProcessingControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraImageProcessingControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraImageProcessingControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageProcessingControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraImageProcessingControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraImageProcessingControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageProcessingControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QCameraInfo struct {
	ptr unsafe.Pointer
}

type QCameraInfo_ITF interface {
	QCameraInfo_PTR() *QCameraInfo
}

func (p *QCameraInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QCameraInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQCameraInfo(ptr QCameraInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraInfo_PTR().Pointer()
	}
	return nil
}

func NewQCameraInfoFromPointer(ptr unsafe.Pointer) *QCameraInfo {
	var n = new(QCameraInfo)
	n.SetPointer(ptr)
	return n
}

func newQCameraInfoFromPointer(ptr unsafe.Pointer) *QCameraInfo {
	var n = NewQCameraInfoFromPointer(ptr)
	return n
}

func (ptr *QCameraInfo) QCameraInfo_PTR() *QCameraInfo {
	return ptr
}

func NewQCameraInfo(name string) *QCameraInfo {
	defer qt.Recovering("QCameraInfo::QCameraInfo")

	return newQCameraInfoFromPointer(C.QCameraInfo_NewQCameraInfo(C.CString(name)))
}

func NewQCameraInfo2(camera QCamera_ITF) *QCameraInfo {
	defer qt.Recovering("QCameraInfo::QCameraInfo")

	return newQCameraInfoFromPointer(C.QCameraInfo_NewQCameraInfo2(PointerFromQCamera(camera)))
}

func NewQCameraInfo3(other QCameraInfo_ITF) *QCameraInfo {
	defer qt.Recovering("QCameraInfo::QCameraInfo")

	return newQCameraInfoFromPointer(C.QCameraInfo_NewQCameraInfo3(PointerFromQCameraInfo(other)))
}

func QCameraInfo_DefaultCamera() *QCameraInfo {
	defer qt.Recovering("QCameraInfo::defaultCamera")

	return NewQCameraInfoFromPointer(C.QCameraInfo_QCameraInfo_DefaultCamera())
}

func (ptr *QCameraInfo) Description() string {
	defer qt.Recovering("QCameraInfo::description")

	if ptr.Pointer() != nil {
		return C.GoString(C.QCameraInfo_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QCameraInfo) DeviceName() string {
	defer qt.Recovering("QCameraInfo::deviceName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QCameraInfo_DeviceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QCameraInfo) IsNull() bool {
	defer qt.Recovering("QCameraInfo::isNull")

	if ptr.Pointer() != nil {
		return C.QCameraInfo_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCameraInfo) Orientation() int {
	defer qt.Recovering("QCameraInfo::orientation")

	if ptr.Pointer() != nil {
		return int(C.QCameraInfo_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraInfo) Position() QCamera__Position {
	defer qt.Recovering("QCameraInfo::position")

	if ptr.Pointer() != nil {
		return QCamera__Position(C.QCameraInfo_Position(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraInfo) DestroyQCameraInfo() {
	defer qt.Recovering("QCameraInfo::~QCameraInfo")

	if ptr.Pointer() != nil {
		C.QCameraInfo_DestroyQCameraInfo(ptr.Pointer())
	}
}

type QCameraInfoControl struct {
	QMediaControl
}

type QCameraInfoControl_ITF interface {
	QMediaControl_ITF
	QCameraInfoControl_PTR() *QCameraInfoControl
}

func PointerFromQCameraInfoControl(ptr QCameraInfoControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraInfoControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraInfoControlFromPointer(ptr unsafe.Pointer) *QCameraInfoControl {
	var n = new(QCameraInfoControl)
	n.SetPointer(ptr)
	return n
}

func newQCameraInfoControlFromPointer(ptr unsafe.Pointer) *QCameraInfoControl {
	var n = NewQCameraInfoControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QCameraInfoControl_") {
		n.SetObjectName("QCameraInfoControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraInfoControl) QCameraInfoControl_PTR() *QCameraInfoControl {
	return ptr
}

func (ptr *QCameraInfoControl) CameraOrientation(deviceName string) int {
	defer qt.Recovering("QCameraInfoControl::cameraOrientation")

	if ptr.Pointer() != nil {
		return int(C.QCameraInfoControl_CameraOrientation(ptr.Pointer(), C.CString(deviceName)))
	}
	return 0
}

func (ptr *QCameraInfoControl) CameraPosition(deviceName string) QCamera__Position {
	defer qt.Recovering("QCameraInfoControl::cameraPosition")

	if ptr.Pointer() != nil {
		return QCamera__Position(C.QCameraInfoControl_CameraPosition(ptr.Pointer(), C.CString(deviceName)))
	}
	return 0
}

func (ptr *QCameraInfoControl) DestroyQCameraInfoControl() {
	defer qt.Recovering("QCameraInfoControl::~QCameraInfoControl")

	if ptr.Pointer() != nil {
		C.QCameraInfoControl_DestroyQCameraInfoControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraInfoControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraInfoControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraInfoControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraInfoControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraInfoControlTimerEvent
func callbackQCameraInfoControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraInfoControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraInfoControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraInfoControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraInfoControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraInfoControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraInfoControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraInfoControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraInfoControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraInfoControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraInfoControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraInfoControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraInfoControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraInfoControlChildEvent
func callbackQCameraInfoControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraInfoControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraInfoControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraInfoControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraInfoControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraInfoControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraInfoControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraInfoControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraInfoControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraInfoControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraInfoControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraInfoControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraInfoControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraInfoControlCustomEvent
func callbackQCameraInfoControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraInfoControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraInfoControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraInfoControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraInfoControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraInfoControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraInfoControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraInfoControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraInfoControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QCameraLocksControl struct {
	QMediaControl
}

type QCameraLocksControl_ITF interface {
	QMediaControl_ITF
	QCameraLocksControl_PTR() *QCameraLocksControl
}

func PointerFromQCameraLocksControl(ptr QCameraLocksControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraLocksControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraLocksControlFromPointer(ptr unsafe.Pointer) *QCameraLocksControl {
	var n = new(QCameraLocksControl)
	n.SetPointer(ptr)
	return n
}

func newQCameraLocksControlFromPointer(ptr unsafe.Pointer) *QCameraLocksControl {
	var n = NewQCameraLocksControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QCameraLocksControl_") {
		n.SetObjectName("QCameraLocksControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraLocksControl) QCameraLocksControl_PTR() *QCameraLocksControl {
	return ptr
}

func (ptr *QCameraLocksControl) LockStatus(lock QCamera__LockType) QCamera__LockStatus {
	defer qt.Recovering("QCameraLocksControl::lockStatus")

	if ptr.Pointer() != nil {
		return QCamera__LockStatus(C.QCameraLocksControl_LockStatus(ptr.Pointer(), C.int(lock)))
	}
	return 0
}

func (ptr *QCameraLocksControl) ConnectLockStatusChanged(f func(lock QCamera__LockType, status QCamera__LockStatus, reason QCamera__LockChangeReason)) {
	defer qt.Recovering("connect QCameraLocksControl::lockStatusChanged")

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_ConnectLockStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "lockStatusChanged", f)
	}
}

func (ptr *QCameraLocksControl) DisconnectLockStatusChanged() {
	defer qt.Recovering("disconnect QCameraLocksControl::lockStatusChanged")

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_DisconnectLockStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "lockStatusChanged")
	}
}

//export callbackQCameraLocksControlLockStatusChanged
func callbackQCameraLocksControlLockStatusChanged(ptr unsafe.Pointer, ptrName *C.char, lock C.int, status C.int, reason C.int) {
	defer qt.Recovering("callback QCameraLocksControl::lockStatusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "lockStatusChanged"); signal != nil {
		signal.(func(QCamera__LockType, QCamera__LockStatus, QCamera__LockChangeReason))(QCamera__LockType(lock), QCamera__LockStatus(status), QCamera__LockChangeReason(reason))
	}

}

func (ptr *QCameraLocksControl) LockStatusChanged(lock QCamera__LockType, status QCamera__LockStatus, reason QCamera__LockChangeReason) {
	defer qt.Recovering("QCameraLocksControl::lockStatusChanged")

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_LockStatusChanged(ptr.Pointer(), C.int(lock), C.int(status), C.int(reason))
	}
}

func (ptr *QCameraLocksControl) SearchAndLock(locks QCamera__LockType) {
	defer qt.Recovering("QCameraLocksControl::searchAndLock")

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_SearchAndLock(ptr.Pointer(), C.int(locks))
	}
}

func (ptr *QCameraLocksControl) SupportedLocks() QCamera__LockType {
	defer qt.Recovering("QCameraLocksControl::supportedLocks")

	if ptr.Pointer() != nil {
		return QCamera__LockType(C.QCameraLocksControl_SupportedLocks(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraLocksControl) Unlock(locks QCamera__LockType) {
	defer qt.Recovering("QCameraLocksControl::unlock")

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_Unlock(ptr.Pointer(), C.int(locks))
	}
}

func (ptr *QCameraLocksControl) DestroyQCameraLocksControl() {
	defer qt.Recovering("QCameraLocksControl::~QCameraLocksControl")

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_DestroyQCameraLocksControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraLocksControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraLocksControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraLocksControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraLocksControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraLocksControlTimerEvent
func callbackQCameraLocksControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraLocksControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraLocksControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraLocksControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraLocksControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraLocksControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraLocksControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraLocksControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraLocksControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraLocksControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraLocksControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraLocksControlChildEvent
func callbackQCameraLocksControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraLocksControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraLocksControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraLocksControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraLocksControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraLocksControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraLocksControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraLocksControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraLocksControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraLocksControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraLocksControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraLocksControlCustomEvent
func callbackQCameraLocksControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraLocksControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraLocksControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraLocksControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraLocksControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraLocksControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraLocksControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QCameraViewfinderSettings struct {
	ptr unsafe.Pointer
}

type QCameraViewfinderSettings_ITF interface {
	QCameraViewfinderSettings_PTR() *QCameraViewfinderSettings
}

func (p *QCameraViewfinderSettings) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QCameraViewfinderSettings) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQCameraViewfinderSettings(ptr QCameraViewfinderSettings_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraViewfinderSettings_PTR().Pointer()
	}
	return nil
}

func NewQCameraViewfinderSettingsFromPointer(ptr unsafe.Pointer) *QCameraViewfinderSettings {
	var n = new(QCameraViewfinderSettings)
	n.SetPointer(ptr)
	return n
}

func newQCameraViewfinderSettingsFromPointer(ptr unsafe.Pointer) *QCameraViewfinderSettings {
	var n = NewQCameraViewfinderSettingsFromPointer(ptr)
	return n
}

func (ptr *QCameraViewfinderSettings) QCameraViewfinderSettings_PTR() *QCameraViewfinderSettings {
	return ptr
}

func NewQCameraViewfinderSettings() *QCameraViewfinderSettings {
	defer qt.Recovering("QCameraViewfinderSettings::QCameraViewfinderSettings")

	return newQCameraViewfinderSettingsFromPointer(C.QCameraViewfinderSettings_NewQCameraViewfinderSettings())
}

func NewQCameraViewfinderSettings2(other QCameraViewfinderSettings_ITF) *QCameraViewfinderSettings {
	defer qt.Recovering("QCameraViewfinderSettings::QCameraViewfinderSettings")

	return newQCameraViewfinderSettingsFromPointer(C.QCameraViewfinderSettings_NewQCameraViewfinderSettings2(PointerFromQCameraViewfinderSettings(other)))
}

func (ptr *QCameraViewfinderSettings) IsNull() bool {
	defer qt.Recovering("QCameraViewfinderSettings::isNull")

	if ptr.Pointer() != nil {
		return C.QCameraViewfinderSettings_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCameraViewfinderSettings) MaximumFrameRate() float64 {
	defer qt.Recovering("QCameraViewfinderSettings::maximumFrameRate")

	if ptr.Pointer() != nil {
		return float64(C.QCameraViewfinderSettings_MaximumFrameRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraViewfinderSettings) MinimumFrameRate() float64 {
	defer qt.Recovering("QCameraViewfinderSettings::minimumFrameRate")

	if ptr.Pointer() != nil {
		return float64(C.QCameraViewfinderSettings_MinimumFrameRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraViewfinderSettings) PixelAspectRatio() *core.QSize {
	defer qt.Recovering("QCameraViewfinderSettings::pixelAspectRatio")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QCameraViewfinderSettings_PixelAspectRatio(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCameraViewfinderSettings) PixelFormat() QVideoFrame__PixelFormat {
	defer qt.Recovering("QCameraViewfinderSettings::pixelFormat")

	if ptr.Pointer() != nil {
		return QVideoFrame__PixelFormat(C.QCameraViewfinderSettings_PixelFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraViewfinderSettings) Resolution() *core.QSize {
	defer qt.Recovering("QCameraViewfinderSettings::resolution")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QCameraViewfinderSettings_Resolution(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCameraViewfinderSettings) SetMaximumFrameRate(rate float64) {
	defer qt.Recovering("QCameraViewfinderSettings::setMaximumFrameRate")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettings_SetMaximumFrameRate(ptr.Pointer(), C.double(rate))
	}
}

func (ptr *QCameraViewfinderSettings) SetMinimumFrameRate(rate float64) {
	defer qt.Recovering("QCameraViewfinderSettings::setMinimumFrameRate")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettings_SetMinimumFrameRate(ptr.Pointer(), C.double(rate))
	}
}

func (ptr *QCameraViewfinderSettings) SetPixelAspectRatio(ratio core.QSize_ITF) {
	defer qt.Recovering("QCameraViewfinderSettings::setPixelAspectRatio")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettings_SetPixelAspectRatio(ptr.Pointer(), core.PointerFromQSize(ratio))
	}
}

func (ptr *QCameraViewfinderSettings) SetPixelAspectRatio2(horizontal int, vertical int) {
	defer qt.Recovering("QCameraViewfinderSettings::setPixelAspectRatio")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettings_SetPixelAspectRatio2(ptr.Pointer(), C.int(horizontal), C.int(vertical))
	}
}

func (ptr *QCameraViewfinderSettings) SetPixelFormat(format QVideoFrame__PixelFormat) {
	defer qt.Recovering("QCameraViewfinderSettings::setPixelFormat")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettings_SetPixelFormat(ptr.Pointer(), C.int(format))
	}
}

func (ptr *QCameraViewfinderSettings) SetResolution(resolution core.QSize_ITF) {
	defer qt.Recovering("QCameraViewfinderSettings::setResolution")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettings_SetResolution(ptr.Pointer(), core.PointerFromQSize(resolution))
	}
}

func (ptr *QCameraViewfinderSettings) SetResolution2(width int, height int) {
	defer qt.Recovering("QCameraViewfinderSettings::setResolution")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettings_SetResolution2(ptr.Pointer(), C.int(width), C.int(height))
	}
}

func (ptr *QCameraViewfinderSettings) Swap(other QCameraViewfinderSettings_ITF) {
	defer qt.Recovering("QCameraViewfinderSettings::swap")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettings_Swap(ptr.Pointer(), PointerFromQCameraViewfinderSettings(other))
	}
}

func (ptr *QCameraViewfinderSettings) DestroyQCameraViewfinderSettings() {
	defer qt.Recovering("QCameraViewfinderSettings::~QCameraViewfinderSettings")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettings_DestroyQCameraViewfinderSettings(ptr.Pointer())
	}
}

type QCameraViewfinderSettingsControl struct {
	QMediaControl
}

type QCameraViewfinderSettingsControl_ITF interface {
	QMediaControl_ITF
	QCameraViewfinderSettingsControl_PTR() *QCameraViewfinderSettingsControl
}

func PointerFromQCameraViewfinderSettingsControl(ptr QCameraViewfinderSettingsControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraViewfinderSettingsControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraViewfinderSettingsControlFromPointer(ptr unsafe.Pointer) *QCameraViewfinderSettingsControl {
	var n = new(QCameraViewfinderSettingsControl)
	n.SetPointer(ptr)
	return n
}

func newQCameraViewfinderSettingsControlFromPointer(ptr unsafe.Pointer) *QCameraViewfinderSettingsControl {
	var n = NewQCameraViewfinderSettingsControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QCameraViewfinderSettingsControl_") {
		n.SetObjectName("QCameraViewfinderSettingsControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraViewfinderSettingsControl) QCameraViewfinderSettingsControl_PTR() *QCameraViewfinderSettingsControl {
	return ptr
}

//QCameraViewfinderSettingsControl::ViewfinderParameter
type QCameraViewfinderSettingsControl__ViewfinderParameter int64

const (
	QCameraViewfinderSettingsControl__Resolution       = QCameraViewfinderSettingsControl__ViewfinderParameter(0)
	QCameraViewfinderSettingsControl__PixelAspectRatio = QCameraViewfinderSettingsControl__ViewfinderParameter(1)
	QCameraViewfinderSettingsControl__MinimumFrameRate = QCameraViewfinderSettingsControl__ViewfinderParameter(2)
	QCameraViewfinderSettingsControl__MaximumFrameRate = QCameraViewfinderSettingsControl__ViewfinderParameter(3)
	QCameraViewfinderSettingsControl__PixelFormat      = QCameraViewfinderSettingsControl__ViewfinderParameter(4)
	QCameraViewfinderSettingsControl__UserParameter    = QCameraViewfinderSettingsControl__ViewfinderParameter(1000)
)

func (ptr *QCameraViewfinderSettingsControl) IsViewfinderParameterSupported(parameter QCameraViewfinderSettingsControl__ViewfinderParameter) bool {
	defer qt.Recovering("QCameraViewfinderSettingsControl::isViewfinderParameterSupported")

	if ptr.Pointer() != nil {
		return C.QCameraViewfinderSettingsControl_IsViewfinderParameterSupported(ptr.Pointer(), C.int(parameter)) != 0
	}
	return false
}

func (ptr *QCameraViewfinderSettingsControl) SetViewfinderParameter(parameter QCameraViewfinderSettingsControl__ViewfinderParameter, value core.QVariant_ITF) {
	defer qt.Recovering("QCameraViewfinderSettingsControl::setViewfinderParameter")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl_SetViewfinderParameter(ptr.Pointer(), C.int(parameter), core.PointerFromQVariant(value))
	}
}

func (ptr *QCameraViewfinderSettingsControl) ViewfinderParameter(parameter QCameraViewfinderSettingsControl__ViewfinderParameter) *core.QVariant {
	defer qt.Recovering("QCameraViewfinderSettingsControl::viewfinderParameter")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QCameraViewfinderSettingsControl_ViewfinderParameter(ptr.Pointer(), C.int(parameter)))
	}
	return nil
}

func (ptr *QCameraViewfinderSettingsControl) DestroyQCameraViewfinderSettingsControl() {
	defer qt.Recovering("QCameraViewfinderSettingsControl::~QCameraViewfinderSettingsControl")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl_DestroyQCameraViewfinderSettingsControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraViewfinderSettingsControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraViewfinderSettingsControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraViewfinderSettingsControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraViewfinderSettingsControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraViewfinderSettingsControlTimerEvent
func callbackQCameraViewfinderSettingsControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinderSettingsControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraViewfinderSettingsControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinderSettingsControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraViewfinderSettingsControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraViewfinderSettingsControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraViewfinderSettingsControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraViewfinderSettingsControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraViewfinderSettingsControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraViewfinderSettingsControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraViewfinderSettingsControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraViewfinderSettingsControlChildEvent
func callbackQCameraViewfinderSettingsControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinderSettingsControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraViewfinderSettingsControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinderSettingsControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraViewfinderSettingsControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraViewfinderSettingsControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraViewfinderSettingsControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraViewfinderSettingsControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraViewfinderSettingsControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraViewfinderSettingsControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraViewfinderSettingsControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraViewfinderSettingsControlCustomEvent
func callbackQCameraViewfinderSettingsControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinderSettingsControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraViewfinderSettingsControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinderSettingsControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraViewfinderSettingsControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraViewfinderSettingsControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraViewfinderSettingsControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QCameraViewfinderSettingsControl2 struct {
	QMediaControl
}

type QCameraViewfinderSettingsControl2_ITF interface {
	QMediaControl_ITF
	QCameraViewfinderSettingsControl2_PTR() *QCameraViewfinderSettingsControl2
}

func PointerFromQCameraViewfinderSettingsControl2(ptr QCameraViewfinderSettingsControl2_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraViewfinderSettingsControl2_PTR().Pointer()
	}
	return nil
}

func NewQCameraViewfinderSettingsControl2FromPointer(ptr unsafe.Pointer) *QCameraViewfinderSettingsControl2 {
	var n = new(QCameraViewfinderSettingsControl2)
	n.SetPointer(ptr)
	return n
}

func newQCameraViewfinderSettingsControl2FromPointer(ptr unsafe.Pointer) *QCameraViewfinderSettingsControl2 {
	var n = NewQCameraViewfinderSettingsControl2FromPointer(ptr)
	for len(n.ObjectName()) < len("QCameraViewfinderSettingsControl2_") {
		n.SetObjectName("QCameraViewfinderSettingsControl2_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraViewfinderSettingsControl2) QCameraViewfinderSettingsControl2_PTR() *QCameraViewfinderSettingsControl2 {
	return ptr
}

func (ptr *QCameraViewfinderSettingsControl2) SetViewfinderSettings(settings QCameraViewfinderSettings_ITF) {
	defer qt.Recovering("QCameraViewfinderSettingsControl2::setViewfinderSettings")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl2_SetViewfinderSettings(ptr.Pointer(), PointerFromQCameraViewfinderSettings(settings))
	}
}

func (ptr *QCameraViewfinderSettingsControl2) ViewfinderSettings() *QCameraViewfinderSettings {
	defer qt.Recovering("QCameraViewfinderSettingsControl2::viewfinderSettings")

	if ptr.Pointer() != nil {
		return NewQCameraViewfinderSettingsFromPointer(C.QCameraViewfinderSettingsControl2_ViewfinderSettings(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCameraViewfinderSettingsControl2) DestroyQCameraViewfinderSettingsControl2() {
	defer qt.Recovering("QCameraViewfinderSettingsControl2::~QCameraViewfinderSettingsControl2")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl2_DestroyQCameraViewfinderSettingsControl2(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraViewfinderSettingsControl2) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraViewfinderSettingsControl2::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraViewfinderSettingsControl2) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraViewfinderSettingsControl2::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraViewfinderSettingsControl2TimerEvent
func callbackQCameraViewfinderSettingsControl2TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinderSettingsControl2::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraViewfinderSettingsControl2FromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinderSettingsControl2) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraViewfinderSettingsControl2::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl2_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraViewfinderSettingsControl2) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraViewfinderSettingsControl2::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl2_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraViewfinderSettingsControl2) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraViewfinderSettingsControl2::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraViewfinderSettingsControl2) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraViewfinderSettingsControl2::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraViewfinderSettingsControl2ChildEvent
func callbackQCameraViewfinderSettingsControl2ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinderSettingsControl2::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraViewfinderSettingsControl2FromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinderSettingsControl2) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraViewfinderSettingsControl2::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl2_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraViewfinderSettingsControl2) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraViewfinderSettingsControl2::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl2_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraViewfinderSettingsControl2) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraViewfinderSettingsControl2::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraViewfinderSettingsControl2) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraViewfinderSettingsControl2::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraViewfinderSettingsControl2CustomEvent
func callbackQCameraViewfinderSettingsControl2CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinderSettingsControl2::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraViewfinderSettingsControl2FromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinderSettingsControl2) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraViewfinderSettingsControl2::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl2_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraViewfinderSettingsControl2) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraViewfinderSettingsControl2::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl2_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QCameraZoomControl struct {
	QMediaControl
}

type QCameraZoomControl_ITF interface {
	QMediaControl_ITF
	QCameraZoomControl_PTR() *QCameraZoomControl
}

func PointerFromQCameraZoomControl(ptr QCameraZoomControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraZoomControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraZoomControlFromPointer(ptr unsafe.Pointer) *QCameraZoomControl {
	var n = new(QCameraZoomControl)
	n.SetPointer(ptr)
	return n
}

func newQCameraZoomControlFromPointer(ptr unsafe.Pointer) *QCameraZoomControl {
	var n = NewQCameraZoomControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QCameraZoomControl_") {
		n.SetObjectName("QCameraZoomControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraZoomControl) QCameraZoomControl_PTR() *QCameraZoomControl {
	return ptr
}

func (ptr *QCameraZoomControl) CurrentDigitalZoom() float64 {
	defer qt.Recovering("QCameraZoomControl::currentDigitalZoom")

	if ptr.Pointer() != nil {
		return float64(C.QCameraZoomControl_CurrentDigitalZoom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraZoomControl) ConnectCurrentDigitalZoomChanged(f func(zoom float64)) {
	defer qt.Recovering("connect QCameraZoomControl::currentDigitalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_ConnectCurrentDigitalZoomChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentDigitalZoomChanged", f)
	}
}

func (ptr *QCameraZoomControl) DisconnectCurrentDigitalZoomChanged() {
	defer qt.Recovering("disconnect QCameraZoomControl::currentDigitalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_DisconnectCurrentDigitalZoomChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentDigitalZoomChanged")
	}
}

//export callbackQCameraZoomControlCurrentDigitalZoomChanged
func callbackQCameraZoomControlCurrentDigitalZoomChanged(ptr unsafe.Pointer, ptrName *C.char, zoom C.double) {
	defer qt.Recovering("callback QCameraZoomControl::currentDigitalZoomChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentDigitalZoomChanged"); signal != nil {
		signal.(func(float64))(float64(zoom))
	}

}

func (ptr *QCameraZoomControl) CurrentDigitalZoomChanged(zoom float64) {
	defer qt.Recovering("QCameraZoomControl::currentDigitalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_CurrentDigitalZoomChanged(ptr.Pointer(), C.double(zoom))
	}
}

func (ptr *QCameraZoomControl) CurrentOpticalZoom() float64 {
	defer qt.Recovering("QCameraZoomControl::currentOpticalZoom")

	if ptr.Pointer() != nil {
		return float64(C.QCameraZoomControl_CurrentOpticalZoom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraZoomControl) ConnectCurrentOpticalZoomChanged(f func(zoom float64)) {
	defer qt.Recovering("connect QCameraZoomControl::currentOpticalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_ConnectCurrentOpticalZoomChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentOpticalZoomChanged", f)
	}
}

func (ptr *QCameraZoomControl) DisconnectCurrentOpticalZoomChanged() {
	defer qt.Recovering("disconnect QCameraZoomControl::currentOpticalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_DisconnectCurrentOpticalZoomChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentOpticalZoomChanged")
	}
}

//export callbackQCameraZoomControlCurrentOpticalZoomChanged
func callbackQCameraZoomControlCurrentOpticalZoomChanged(ptr unsafe.Pointer, ptrName *C.char, zoom C.double) {
	defer qt.Recovering("callback QCameraZoomControl::currentOpticalZoomChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentOpticalZoomChanged"); signal != nil {
		signal.(func(float64))(float64(zoom))
	}

}

func (ptr *QCameraZoomControl) CurrentOpticalZoomChanged(zoom float64) {
	defer qt.Recovering("QCameraZoomControl::currentOpticalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_CurrentOpticalZoomChanged(ptr.Pointer(), C.double(zoom))
	}
}

func (ptr *QCameraZoomControl) MaximumDigitalZoom() float64 {
	defer qt.Recovering("QCameraZoomControl::maximumDigitalZoom")

	if ptr.Pointer() != nil {
		return float64(C.QCameraZoomControl_MaximumDigitalZoom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraZoomControl) ConnectMaximumDigitalZoomChanged(f func(zoom float64)) {
	defer qt.Recovering("connect QCameraZoomControl::maximumDigitalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_ConnectMaximumDigitalZoomChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "maximumDigitalZoomChanged", f)
	}
}

func (ptr *QCameraZoomControl) DisconnectMaximumDigitalZoomChanged() {
	defer qt.Recovering("disconnect QCameraZoomControl::maximumDigitalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_DisconnectMaximumDigitalZoomChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "maximumDigitalZoomChanged")
	}
}

//export callbackQCameraZoomControlMaximumDigitalZoomChanged
func callbackQCameraZoomControlMaximumDigitalZoomChanged(ptr unsafe.Pointer, ptrName *C.char, zoom C.double) {
	defer qt.Recovering("callback QCameraZoomControl::maximumDigitalZoomChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "maximumDigitalZoomChanged"); signal != nil {
		signal.(func(float64))(float64(zoom))
	}

}

func (ptr *QCameraZoomControl) MaximumDigitalZoomChanged(zoom float64) {
	defer qt.Recovering("QCameraZoomControl::maximumDigitalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_MaximumDigitalZoomChanged(ptr.Pointer(), C.double(zoom))
	}
}

func (ptr *QCameraZoomControl) MaximumOpticalZoom() float64 {
	defer qt.Recovering("QCameraZoomControl::maximumOpticalZoom")

	if ptr.Pointer() != nil {
		return float64(C.QCameraZoomControl_MaximumOpticalZoom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraZoomControl) ConnectMaximumOpticalZoomChanged(f func(zoom float64)) {
	defer qt.Recovering("connect QCameraZoomControl::maximumOpticalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_ConnectMaximumOpticalZoomChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "maximumOpticalZoomChanged", f)
	}
}

func (ptr *QCameraZoomControl) DisconnectMaximumOpticalZoomChanged() {
	defer qt.Recovering("disconnect QCameraZoomControl::maximumOpticalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_DisconnectMaximumOpticalZoomChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "maximumOpticalZoomChanged")
	}
}

//export callbackQCameraZoomControlMaximumOpticalZoomChanged
func callbackQCameraZoomControlMaximumOpticalZoomChanged(ptr unsafe.Pointer, ptrName *C.char, zoom C.double) {
	defer qt.Recovering("callback QCameraZoomControl::maximumOpticalZoomChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "maximumOpticalZoomChanged"); signal != nil {
		signal.(func(float64))(float64(zoom))
	}

}

func (ptr *QCameraZoomControl) MaximumOpticalZoomChanged(zoom float64) {
	defer qt.Recovering("QCameraZoomControl::maximumOpticalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_MaximumOpticalZoomChanged(ptr.Pointer(), C.double(zoom))
	}
}

func (ptr *QCameraZoomControl) RequestedDigitalZoom() float64 {
	defer qt.Recovering("QCameraZoomControl::requestedDigitalZoom")

	if ptr.Pointer() != nil {
		return float64(C.QCameraZoomControl_RequestedDigitalZoom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraZoomControl) ConnectRequestedDigitalZoomChanged(f func(zoom float64)) {
	defer qt.Recovering("connect QCameraZoomControl::requestedDigitalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_ConnectRequestedDigitalZoomChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "requestedDigitalZoomChanged", f)
	}
}

func (ptr *QCameraZoomControl) DisconnectRequestedDigitalZoomChanged() {
	defer qt.Recovering("disconnect QCameraZoomControl::requestedDigitalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_DisconnectRequestedDigitalZoomChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "requestedDigitalZoomChanged")
	}
}

//export callbackQCameraZoomControlRequestedDigitalZoomChanged
func callbackQCameraZoomControlRequestedDigitalZoomChanged(ptr unsafe.Pointer, ptrName *C.char, zoom C.double) {
	defer qt.Recovering("callback QCameraZoomControl::requestedDigitalZoomChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "requestedDigitalZoomChanged"); signal != nil {
		signal.(func(float64))(float64(zoom))
	}

}

func (ptr *QCameraZoomControl) RequestedDigitalZoomChanged(zoom float64) {
	defer qt.Recovering("QCameraZoomControl::requestedDigitalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_RequestedDigitalZoomChanged(ptr.Pointer(), C.double(zoom))
	}
}

func (ptr *QCameraZoomControl) RequestedOpticalZoom() float64 {
	defer qt.Recovering("QCameraZoomControl::requestedOpticalZoom")

	if ptr.Pointer() != nil {
		return float64(C.QCameraZoomControl_RequestedOpticalZoom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraZoomControl) ConnectRequestedOpticalZoomChanged(f func(zoom float64)) {
	defer qt.Recovering("connect QCameraZoomControl::requestedOpticalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_ConnectRequestedOpticalZoomChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "requestedOpticalZoomChanged", f)
	}
}

func (ptr *QCameraZoomControl) DisconnectRequestedOpticalZoomChanged() {
	defer qt.Recovering("disconnect QCameraZoomControl::requestedOpticalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_DisconnectRequestedOpticalZoomChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "requestedOpticalZoomChanged")
	}
}

//export callbackQCameraZoomControlRequestedOpticalZoomChanged
func callbackQCameraZoomControlRequestedOpticalZoomChanged(ptr unsafe.Pointer, ptrName *C.char, zoom C.double) {
	defer qt.Recovering("callback QCameraZoomControl::requestedOpticalZoomChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "requestedOpticalZoomChanged"); signal != nil {
		signal.(func(float64))(float64(zoom))
	}

}

func (ptr *QCameraZoomControl) RequestedOpticalZoomChanged(zoom float64) {
	defer qt.Recovering("QCameraZoomControl::requestedOpticalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_RequestedOpticalZoomChanged(ptr.Pointer(), C.double(zoom))
	}
}

func (ptr *QCameraZoomControl) ZoomTo(optical float64, digital float64) {
	defer qt.Recovering("QCameraZoomControl::zoomTo")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_ZoomTo(ptr.Pointer(), C.double(optical), C.double(digital))
	}
}

func (ptr *QCameraZoomControl) DestroyQCameraZoomControl() {
	defer qt.Recovering("QCameraZoomControl::~QCameraZoomControl")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_DestroyQCameraZoomControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraZoomControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraZoomControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraZoomControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraZoomControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraZoomControlTimerEvent
func callbackQCameraZoomControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraZoomControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraZoomControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraZoomControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraZoomControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraZoomControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraZoomControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraZoomControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraZoomControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraZoomControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraZoomControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraZoomControlChildEvent
func callbackQCameraZoomControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraZoomControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraZoomControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraZoomControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraZoomControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraZoomControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraZoomControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraZoomControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraZoomControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraZoomControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraZoomControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraZoomControlCustomEvent
func callbackQCameraZoomControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraZoomControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraZoomControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraZoomControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraZoomControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraZoomControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraZoomControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QImageEncoderControl struct {
	QMediaControl
}

type QImageEncoderControl_ITF interface {
	QMediaControl_ITF
	QImageEncoderControl_PTR() *QImageEncoderControl
}

func PointerFromQImageEncoderControl(ptr QImageEncoderControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QImageEncoderControl_PTR().Pointer()
	}
	return nil
}

func NewQImageEncoderControlFromPointer(ptr unsafe.Pointer) *QImageEncoderControl {
	var n = new(QImageEncoderControl)
	n.SetPointer(ptr)
	return n
}

func newQImageEncoderControlFromPointer(ptr unsafe.Pointer) *QImageEncoderControl {
	var n = NewQImageEncoderControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QImageEncoderControl_") {
		n.SetObjectName("QImageEncoderControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QImageEncoderControl) QImageEncoderControl_PTR() *QImageEncoderControl {
	return ptr
}

func (ptr *QImageEncoderControl) ImageCodecDescription(codec string) string {
	defer qt.Recovering("QImageEncoderControl::imageCodecDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QImageEncoderControl_ImageCodecDescription(ptr.Pointer(), C.CString(codec)))
	}
	return ""
}

func (ptr *QImageEncoderControl) ImageSettings() *QImageEncoderSettings {
	defer qt.Recovering("QImageEncoderControl::imageSettings")

	if ptr.Pointer() != nil {
		return NewQImageEncoderSettingsFromPointer(C.QImageEncoderControl_ImageSettings(ptr.Pointer()))
	}
	return nil
}

func (ptr *QImageEncoderControl) SetImageSettings(settings QImageEncoderSettings_ITF) {
	defer qt.Recovering("QImageEncoderControl::setImageSettings")

	if ptr.Pointer() != nil {
		C.QImageEncoderControl_SetImageSettings(ptr.Pointer(), PointerFromQImageEncoderSettings(settings))
	}
}

func (ptr *QImageEncoderControl) SupportedImageCodecs() []string {
	defer qt.Recovering("QImageEncoderControl::supportedImageCodecs")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QImageEncoderControl_SupportedImageCodecs(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QImageEncoderControl) DestroyQImageEncoderControl() {
	defer qt.Recovering("QImageEncoderControl::~QImageEncoderControl")

	if ptr.Pointer() != nil {
		C.QImageEncoderControl_DestroyQImageEncoderControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QImageEncoderControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QImageEncoderControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QImageEncoderControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QImageEncoderControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQImageEncoderControlTimerEvent
func callbackQImageEncoderControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QImageEncoderControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQImageEncoderControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QImageEncoderControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QImageEncoderControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QImageEncoderControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QImageEncoderControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QImageEncoderControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QImageEncoderControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QImageEncoderControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QImageEncoderControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QImageEncoderControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QImageEncoderControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQImageEncoderControlChildEvent
func callbackQImageEncoderControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QImageEncoderControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQImageEncoderControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QImageEncoderControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QImageEncoderControl::childEvent")

	if ptr.Pointer() != nil {
		C.QImageEncoderControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QImageEncoderControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QImageEncoderControl::childEvent")

	if ptr.Pointer() != nil {
		C.QImageEncoderControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QImageEncoderControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QImageEncoderControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QImageEncoderControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QImageEncoderControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQImageEncoderControlCustomEvent
func callbackQImageEncoderControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QImageEncoderControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQImageEncoderControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QImageEncoderControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QImageEncoderControl::customEvent")

	if ptr.Pointer() != nil {
		C.QImageEncoderControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QImageEncoderControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QImageEncoderControl::customEvent")

	if ptr.Pointer() != nil {
		C.QImageEncoderControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QImageEncoderSettings struct {
	ptr unsafe.Pointer
}

type QImageEncoderSettings_ITF interface {
	QImageEncoderSettings_PTR() *QImageEncoderSettings
}

func (p *QImageEncoderSettings) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QImageEncoderSettings) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQImageEncoderSettings(ptr QImageEncoderSettings_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QImageEncoderSettings_PTR().Pointer()
	}
	return nil
}

func NewQImageEncoderSettingsFromPointer(ptr unsafe.Pointer) *QImageEncoderSettings {
	var n = new(QImageEncoderSettings)
	n.SetPointer(ptr)
	return n
}

func newQImageEncoderSettingsFromPointer(ptr unsafe.Pointer) *QImageEncoderSettings {
	var n = NewQImageEncoderSettingsFromPointer(ptr)
	return n
}

func (ptr *QImageEncoderSettings) QImageEncoderSettings_PTR() *QImageEncoderSettings {
	return ptr
}

func NewQImageEncoderSettings() *QImageEncoderSettings {
	defer qt.Recovering("QImageEncoderSettings::QImageEncoderSettings")

	return newQImageEncoderSettingsFromPointer(C.QImageEncoderSettings_NewQImageEncoderSettings())
}

func NewQImageEncoderSettings2(other QImageEncoderSettings_ITF) *QImageEncoderSettings {
	defer qt.Recovering("QImageEncoderSettings::QImageEncoderSettings")

	return newQImageEncoderSettingsFromPointer(C.QImageEncoderSettings_NewQImageEncoderSettings2(PointerFromQImageEncoderSettings(other)))
}

func (ptr *QImageEncoderSettings) Codec() string {
	defer qt.Recovering("QImageEncoderSettings::codec")

	if ptr.Pointer() != nil {
		return C.GoString(C.QImageEncoderSettings_Codec(ptr.Pointer()))
	}
	return ""
}

func (ptr *QImageEncoderSettings) EncodingOption(option string) *core.QVariant {
	defer qt.Recovering("QImageEncoderSettings::encodingOption")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QImageEncoderSettings_EncodingOption(ptr.Pointer(), C.CString(option)))
	}
	return nil
}

func (ptr *QImageEncoderSettings) IsNull() bool {
	defer qt.Recovering("QImageEncoderSettings::isNull")

	if ptr.Pointer() != nil {
		return C.QImageEncoderSettings_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QImageEncoderSettings) Quality() QMultimedia__EncodingQuality {
	defer qt.Recovering("QImageEncoderSettings::quality")

	if ptr.Pointer() != nil {
		return QMultimedia__EncodingQuality(C.QImageEncoderSettings_Quality(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImageEncoderSettings) Resolution() *core.QSize {
	defer qt.Recovering("QImageEncoderSettings::resolution")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QImageEncoderSettings_Resolution(ptr.Pointer()))
	}
	return nil
}

func (ptr *QImageEncoderSettings) SetCodec(codec string) {
	defer qt.Recovering("QImageEncoderSettings::setCodec")

	if ptr.Pointer() != nil {
		C.QImageEncoderSettings_SetCodec(ptr.Pointer(), C.CString(codec))
	}
}

func (ptr *QImageEncoderSettings) SetEncodingOption(option string, value core.QVariant_ITF) {
	defer qt.Recovering("QImageEncoderSettings::setEncodingOption")

	if ptr.Pointer() != nil {
		C.QImageEncoderSettings_SetEncodingOption(ptr.Pointer(), C.CString(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QImageEncoderSettings) SetQuality(quality QMultimedia__EncodingQuality) {
	defer qt.Recovering("QImageEncoderSettings::setQuality")

	if ptr.Pointer() != nil {
		C.QImageEncoderSettings_SetQuality(ptr.Pointer(), C.int(quality))
	}
}

func (ptr *QImageEncoderSettings) SetResolution(resolution core.QSize_ITF) {
	defer qt.Recovering("QImageEncoderSettings::setResolution")

	if ptr.Pointer() != nil {
		C.QImageEncoderSettings_SetResolution(ptr.Pointer(), core.PointerFromQSize(resolution))
	}
}

func (ptr *QImageEncoderSettings) SetResolution2(width int, height int) {
	defer qt.Recovering("QImageEncoderSettings::setResolution")

	if ptr.Pointer() != nil {
		C.QImageEncoderSettings_SetResolution2(ptr.Pointer(), C.int(width), C.int(height))
	}
}

func (ptr *QImageEncoderSettings) DestroyQImageEncoderSettings() {
	defer qt.Recovering("QImageEncoderSettings::~QImageEncoderSettings")

	if ptr.Pointer() != nil {
		C.QImageEncoderSettings_DestroyQImageEncoderSettings(ptr.Pointer())
	}
}

type QMediaAudioProbeControl struct {
	QMediaControl
}

type QMediaAudioProbeControl_ITF interface {
	QMediaControl_ITF
	QMediaAudioProbeControl_PTR() *QMediaAudioProbeControl
}

func PointerFromQMediaAudioProbeControl(ptr QMediaAudioProbeControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaAudioProbeControl_PTR().Pointer()
	}
	return nil
}

func NewQMediaAudioProbeControlFromPointer(ptr unsafe.Pointer) *QMediaAudioProbeControl {
	var n = new(QMediaAudioProbeControl)
	n.SetPointer(ptr)
	return n
}

func newQMediaAudioProbeControlFromPointer(ptr unsafe.Pointer) *QMediaAudioProbeControl {
	var n = NewQMediaAudioProbeControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QMediaAudioProbeControl_") {
		n.SetObjectName("QMediaAudioProbeControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaAudioProbeControl) QMediaAudioProbeControl_PTR() *QMediaAudioProbeControl {
	return ptr
}

func (ptr *QMediaAudioProbeControl) ConnectAudioBufferProbed(f func(buffer *QAudioBuffer)) {
	defer qt.Recovering("connect QMediaAudioProbeControl::audioBufferProbed")

	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_ConnectAudioBufferProbed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "audioBufferProbed", f)
	}
}

func (ptr *QMediaAudioProbeControl) DisconnectAudioBufferProbed() {
	defer qt.Recovering("disconnect QMediaAudioProbeControl::audioBufferProbed")

	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_DisconnectAudioBufferProbed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "audioBufferProbed")
	}
}

//export callbackQMediaAudioProbeControlAudioBufferProbed
func callbackQMediaAudioProbeControlAudioBufferProbed(ptr unsafe.Pointer, ptrName *C.char, buffer unsafe.Pointer) {
	defer qt.Recovering("callback QMediaAudioProbeControl::audioBufferProbed")

	if signal := qt.GetSignal(C.GoString(ptrName), "audioBufferProbed"); signal != nil {
		signal.(func(*QAudioBuffer))(NewQAudioBufferFromPointer(buffer))
	}

}

func (ptr *QMediaAudioProbeControl) AudioBufferProbed(buffer QAudioBuffer_ITF) {
	defer qt.Recovering("QMediaAudioProbeControl::audioBufferProbed")

	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_AudioBufferProbed(ptr.Pointer(), PointerFromQAudioBuffer(buffer))
	}
}

func (ptr *QMediaAudioProbeControl) ConnectFlush(f func()) {
	defer qt.Recovering("connect QMediaAudioProbeControl::flush")

	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_ConnectFlush(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "flush", f)
	}
}

func (ptr *QMediaAudioProbeControl) DisconnectFlush() {
	defer qt.Recovering("disconnect QMediaAudioProbeControl::flush")

	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_DisconnectFlush(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "flush")
	}
}

//export callbackQMediaAudioProbeControlFlush
func callbackQMediaAudioProbeControlFlush(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QMediaAudioProbeControl::flush")

	if signal := qt.GetSignal(C.GoString(ptrName), "flush"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QMediaAudioProbeControl) Flush() {
	defer qt.Recovering("QMediaAudioProbeControl::flush")

	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_Flush(ptr.Pointer())
	}
}

func (ptr *QMediaAudioProbeControl) DestroyQMediaAudioProbeControl() {
	defer qt.Recovering("QMediaAudioProbeControl::~QMediaAudioProbeControl")

	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_DestroyQMediaAudioProbeControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMediaAudioProbeControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMediaAudioProbeControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMediaAudioProbeControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMediaAudioProbeControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMediaAudioProbeControlTimerEvent
func callbackQMediaAudioProbeControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaAudioProbeControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMediaAudioProbeControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMediaAudioProbeControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaAudioProbeControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaAudioProbeControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaAudioProbeControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaAudioProbeControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMediaAudioProbeControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMediaAudioProbeControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMediaAudioProbeControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMediaAudioProbeControlChildEvent
func callbackQMediaAudioProbeControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaAudioProbeControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMediaAudioProbeControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMediaAudioProbeControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaAudioProbeControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaAudioProbeControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaAudioProbeControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaAudioProbeControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMediaAudioProbeControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMediaAudioProbeControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMediaAudioProbeControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMediaAudioProbeControlCustomEvent
func callbackQMediaAudioProbeControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaAudioProbeControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMediaAudioProbeControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMediaAudioProbeControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaAudioProbeControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMediaAudioProbeControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaAudioProbeControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QMediaAvailabilityControl struct {
	QMediaControl
}

type QMediaAvailabilityControl_ITF interface {
	QMediaControl_ITF
	QMediaAvailabilityControl_PTR() *QMediaAvailabilityControl
}

func PointerFromQMediaAvailabilityControl(ptr QMediaAvailabilityControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaAvailabilityControl_PTR().Pointer()
	}
	return nil
}

func NewQMediaAvailabilityControlFromPointer(ptr unsafe.Pointer) *QMediaAvailabilityControl {
	var n = new(QMediaAvailabilityControl)
	n.SetPointer(ptr)
	return n
}

func newQMediaAvailabilityControlFromPointer(ptr unsafe.Pointer) *QMediaAvailabilityControl {
	var n = NewQMediaAvailabilityControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QMediaAvailabilityControl_") {
		n.SetObjectName("QMediaAvailabilityControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaAvailabilityControl) QMediaAvailabilityControl_PTR() *QMediaAvailabilityControl {
	return ptr
}

func (ptr *QMediaAvailabilityControl) Availability() QMultimedia__AvailabilityStatus {
	defer qt.Recovering("QMediaAvailabilityControl::availability")

	if ptr.Pointer() != nil {
		return QMultimedia__AvailabilityStatus(C.QMediaAvailabilityControl_Availability(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaAvailabilityControl) ConnectAvailabilityChanged(f func(availability QMultimedia__AvailabilityStatus)) {
	defer qt.Recovering("connect QMediaAvailabilityControl::availabilityChanged")

	if ptr.Pointer() != nil {
		C.QMediaAvailabilityControl_ConnectAvailabilityChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "availabilityChanged", f)
	}
}

func (ptr *QMediaAvailabilityControl) DisconnectAvailabilityChanged() {
	defer qt.Recovering("disconnect QMediaAvailabilityControl::availabilityChanged")

	if ptr.Pointer() != nil {
		C.QMediaAvailabilityControl_DisconnectAvailabilityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "availabilityChanged")
	}
}

//export callbackQMediaAvailabilityControlAvailabilityChanged
func callbackQMediaAvailabilityControlAvailabilityChanged(ptr unsafe.Pointer, ptrName *C.char, availability C.int) {
	defer qt.Recovering("callback QMediaAvailabilityControl::availabilityChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "availabilityChanged"); signal != nil {
		signal.(func(QMultimedia__AvailabilityStatus))(QMultimedia__AvailabilityStatus(availability))
	}

}

func (ptr *QMediaAvailabilityControl) AvailabilityChanged(availability QMultimedia__AvailabilityStatus) {
	defer qt.Recovering("QMediaAvailabilityControl::availabilityChanged")

	if ptr.Pointer() != nil {
		C.QMediaAvailabilityControl_AvailabilityChanged(ptr.Pointer(), C.int(availability))
	}
}

func (ptr *QMediaAvailabilityControl) DestroyQMediaAvailabilityControl() {
	defer qt.Recovering("QMediaAvailabilityControl::~QMediaAvailabilityControl")

	if ptr.Pointer() != nil {
		C.QMediaAvailabilityControl_DestroyQMediaAvailabilityControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMediaAvailabilityControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMediaAvailabilityControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMediaAvailabilityControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMediaAvailabilityControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMediaAvailabilityControlTimerEvent
func callbackQMediaAvailabilityControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaAvailabilityControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMediaAvailabilityControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMediaAvailabilityControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaAvailabilityControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaAvailabilityControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaAvailabilityControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaAvailabilityControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaAvailabilityControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaAvailabilityControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMediaAvailabilityControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMediaAvailabilityControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMediaAvailabilityControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMediaAvailabilityControlChildEvent
func callbackQMediaAvailabilityControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaAvailabilityControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMediaAvailabilityControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMediaAvailabilityControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaAvailabilityControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaAvailabilityControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaAvailabilityControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaAvailabilityControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaAvailabilityControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaAvailabilityControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMediaAvailabilityControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMediaAvailabilityControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMediaAvailabilityControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMediaAvailabilityControlCustomEvent
func callbackQMediaAvailabilityControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaAvailabilityControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMediaAvailabilityControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMediaAvailabilityControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaAvailabilityControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaAvailabilityControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMediaAvailabilityControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaAvailabilityControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaAvailabilityControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QMediaBindableInterface struct {
	ptr unsafe.Pointer
}

type QMediaBindableInterface_ITF interface {
	QMediaBindableInterface_PTR() *QMediaBindableInterface
}

func (p *QMediaBindableInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMediaBindableInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMediaBindableInterface(ptr QMediaBindableInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaBindableInterface_PTR().Pointer()
	}
	return nil
}

func NewQMediaBindableInterfaceFromPointer(ptr unsafe.Pointer) *QMediaBindableInterface {
	var n = new(QMediaBindableInterface)
	n.SetPointer(ptr)
	return n
}

func newQMediaBindableInterfaceFromPointer(ptr unsafe.Pointer) *QMediaBindableInterface {
	var n = NewQMediaBindableInterfaceFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QMediaBindableInterface_") {
		n.SetObjectNameAbs("QMediaBindableInterface_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaBindableInterface) QMediaBindableInterface_PTR() *QMediaBindableInterface {
	return ptr
}

func (ptr *QMediaBindableInterface) MediaObject() *QMediaObject {
	defer qt.Recovering("QMediaBindableInterface::mediaObject")

	if ptr.Pointer() != nil {
		return NewQMediaObjectFromPointer(C.QMediaBindableInterface_MediaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaBindableInterface) DestroyQMediaBindableInterface() {
	defer qt.Recovering("QMediaBindableInterface::~QMediaBindableInterface")

	if ptr.Pointer() != nil {
		C.QMediaBindableInterface_DestroyQMediaBindableInterface(ptr.Pointer())
	}
}

func (ptr *QMediaBindableInterface) ObjectNameAbs() string {
	defer qt.Recovering("QMediaBindableInterface::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaBindableInterface_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaBindableInterface) SetObjectNameAbs(name string) {
	defer qt.Recovering("QMediaBindableInterface::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QMediaBindableInterface_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QMediaContainerControl struct {
	QMediaControl
}

type QMediaContainerControl_ITF interface {
	QMediaControl_ITF
	QMediaContainerControl_PTR() *QMediaContainerControl
}

func PointerFromQMediaContainerControl(ptr QMediaContainerControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaContainerControl_PTR().Pointer()
	}
	return nil
}

func NewQMediaContainerControlFromPointer(ptr unsafe.Pointer) *QMediaContainerControl {
	var n = new(QMediaContainerControl)
	n.SetPointer(ptr)
	return n
}

func newQMediaContainerControlFromPointer(ptr unsafe.Pointer) *QMediaContainerControl {
	var n = NewQMediaContainerControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QMediaContainerControl_") {
		n.SetObjectName("QMediaContainerControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaContainerControl) QMediaContainerControl_PTR() *QMediaContainerControl {
	return ptr
}

func (ptr *QMediaContainerControl) ContainerDescription(format string) string {
	defer qt.Recovering("QMediaContainerControl::containerDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaContainerControl_ContainerDescription(ptr.Pointer(), C.CString(format)))
	}
	return ""
}

func (ptr *QMediaContainerControl) ContainerFormat() string {
	defer qt.Recovering("QMediaContainerControl::containerFormat")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaContainerControl_ContainerFormat(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaContainerControl) SetContainerFormat(format string) {
	defer qt.Recovering("QMediaContainerControl::setContainerFormat")

	if ptr.Pointer() != nil {
		C.QMediaContainerControl_SetContainerFormat(ptr.Pointer(), C.CString(format))
	}
}

func (ptr *QMediaContainerControl) SupportedContainers() []string {
	defer qt.Recovering("QMediaContainerControl::supportedContainers")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMediaContainerControl_SupportedContainers(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QMediaContainerControl) DestroyQMediaContainerControl() {
	defer qt.Recovering("QMediaContainerControl::~QMediaContainerControl")

	if ptr.Pointer() != nil {
		C.QMediaContainerControl_DestroyQMediaContainerControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMediaContainerControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMediaContainerControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMediaContainerControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMediaContainerControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMediaContainerControlTimerEvent
func callbackQMediaContainerControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaContainerControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMediaContainerControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMediaContainerControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaContainerControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaContainerControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaContainerControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaContainerControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaContainerControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaContainerControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMediaContainerControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMediaContainerControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMediaContainerControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMediaContainerControlChildEvent
func callbackQMediaContainerControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaContainerControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMediaContainerControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMediaContainerControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaContainerControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaContainerControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaContainerControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaContainerControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaContainerControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaContainerControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMediaContainerControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMediaContainerControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMediaContainerControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMediaContainerControlCustomEvent
func callbackQMediaContainerControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaContainerControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMediaContainerControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMediaContainerControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaContainerControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaContainerControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMediaContainerControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaContainerControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaContainerControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QMediaContent struct {
	ptr unsafe.Pointer
}

type QMediaContent_ITF interface {
	QMediaContent_PTR() *QMediaContent
}

func (p *QMediaContent) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMediaContent) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMediaContent(ptr QMediaContent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaContent_PTR().Pointer()
	}
	return nil
}

func NewQMediaContentFromPointer(ptr unsafe.Pointer) *QMediaContent {
	var n = new(QMediaContent)
	n.SetPointer(ptr)
	return n
}

func newQMediaContentFromPointer(ptr unsafe.Pointer) *QMediaContent {
	var n = NewQMediaContentFromPointer(ptr)
	return n
}

func (ptr *QMediaContent) QMediaContent_PTR() *QMediaContent {
	return ptr
}

func NewQMediaContent() *QMediaContent {
	defer qt.Recovering("QMediaContent::QMediaContent")

	return newQMediaContentFromPointer(C.QMediaContent_NewQMediaContent())
}

func NewQMediaContent7(playlist QMediaPlaylist_ITF, contentUrl core.QUrl_ITF, takeOwnership bool) *QMediaContent {
	defer qt.Recovering("QMediaContent::QMediaContent")

	return newQMediaContentFromPointer(C.QMediaContent_NewQMediaContent7(PointerFromQMediaPlaylist(playlist), core.PointerFromQUrl(contentUrl), C.int(qt.GoBoolToInt(takeOwnership))))
}

func NewQMediaContent6(other QMediaContent_ITF) *QMediaContent {
	defer qt.Recovering("QMediaContent::QMediaContent")

	return newQMediaContentFromPointer(C.QMediaContent_NewQMediaContent6(PointerFromQMediaContent(other)))
}

func NewQMediaContent4(resource QMediaResource_ITF) *QMediaContent {
	defer qt.Recovering("QMediaContent::QMediaContent")

	return newQMediaContentFromPointer(C.QMediaContent_NewQMediaContent4(PointerFromQMediaResource(resource)))
}

func NewQMediaContent3(request network.QNetworkRequest_ITF) *QMediaContent {
	defer qt.Recovering("QMediaContent::QMediaContent")

	return newQMediaContentFromPointer(C.QMediaContent_NewQMediaContent3(network.PointerFromQNetworkRequest(request)))
}

func NewQMediaContent2(url core.QUrl_ITF) *QMediaContent {
	defer qt.Recovering("QMediaContent::QMediaContent")

	return newQMediaContentFromPointer(C.QMediaContent_NewQMediaContent2(core.PointerFromQUrl(url)))
}

func (ptr *QMediaContent) CanonicalRequest() *network.QNetworkRequest {
	defer qt.Recovering("QMediaContent::canonicalRequest")

	if ptr.Pointer() != nil {
		return network.NewQNetworkRequestFromPointer(C.QMediaContent_CanonicalRequest(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaContent) CanonicalResource() *QMediaResource {
	defer qt.Recovering("QMediaContent::canonicalResource")

	if ptr.Pointer() != nil {
		return NewQMediaResourceFromPointer(C.QMediaContent_CanonicalResource(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaContent) CanonicalUrl() *core.QUrl {
	defer qt.Recovering("QMediaContent::canonicalUrl")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QMediaContent_CanonicalUrl(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaContent) IsNull() bool {
	defer qt.Recovering("QMediaContent::isNull")

	if ptr.Pointer() != nil {
		return C.QMediaContent_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaContent) Playlist() *QMediaPlaylist {
	defer qt.Recovering("QMediaContent::playlist")

	if ptr.Pointer() != nil {
		return NewQMediaPlaylistFromPointer(C.QMediaContent_Playlist(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaContent) DestroyQMediaContent() {
	defer qt.Recovering("QMediaContent::~QMediaContent")

	if ptr.Pointer() != nil {
		C.QMediaContent_DestroyQMediaContent(ptr.Pointer())
	}
}

type QMediaControl struct {
	core.QObject
}

type QMediaControl_ITF interface {
	core.QObject_ITF
	QMediaControl_PTR() *QMediaControl
}

func PointerFromQMediaControl(ptr QMediaControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaControl_PTR().Pointer()
	}
	return nil
}

func NewQMediaControlFromPointer(ptr unsafe.Pointer) *QMediaControl {
	var n = new(QMediaControl)
	n.SetPointer(ptr)
	return n
}

func newQMediaControlFromPointer(ptr unsafe.Pointer) *QMediaControl {
	var n = NewQMediaControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QMediaControl_") {
		n.SetObjectName("QMediaControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaControl) QMediaControl_PTR() *QMediaControl {
	return ptr
}

func (ptr *QMediaControl) DestroyQMediaControl() {
	defer qt.Recovering("QMediaControl::~QMediaControl")

	if ptr.Pointer() != nil {
		C.QMediaControl_DestroyQMediaControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMediaControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMediaControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMediaControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMediaControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMediaControlTimerEvent
func callbackQMediaControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMediaControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMediaControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMediaControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMediaControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMediaControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMediaControlChildEvent
func callbackQMediaControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMediaControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMediaControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMediaControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMediaControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMediaControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMediaControlCustomEvent
func callbackQMediaControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMediaControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMediaControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMediaControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QMediaGaplessPlaybackControl struct {
	QMediaControl
}

type QMediaGaplessPlaybackControl_ITF interface {
	QMediaControl_ITF
	QMediaGaplessPlaybackControl_PTR() *QMediaGaplessPlaybackControl
}

func PointerFromQMediaGaplessPlaybackControl(ptr QMediaGaplessPlaybackControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaGaplessPlaybackControl_PTR().Pointer()
	}
	return nil
}

func NewQMediaGaplessPlaybackControlFromPointer(ptr unsafe.Pointer) *QMediaGaplessPlaybackControl {
	var n = new(QMediaGaplessPlaybackControl)
	n.SetPointer(ptr)
	return n
}

func newQMediaGaplessPlaybackControlFromPointer(ptr unsafe.Pointer) *QMediaGaplessPlaybackControl {
	var n = NewQMediaGaplessPlaybackControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QMediaGaplessPlaybackControl_") {
		n.SetObjectName("QMediaGaplessPlaybackControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaGaplessPlaybackControl) QMediaGaplessPlaybackControl_PTR() *QMediaGaplessPlaybackControl {
	return ptr
}

func (ptr *QMediaGaplessPlaybackControl) ConnectAdvancedToNextMedia(f func()) {
	defer qt.Recovering("connect QMediaGaplessPlaybackControl::advancedToNextMedia")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_ConnectAdvancedToNextMedia(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "advancedToNextMedia", f)
	}
}

func (ptr *QMediaGaplessPlaybackControl) DisconnectAdvancedToNextMedia() {
	defer qt.Recovering("disconnect QMediaGaplessPlaybackControl::advancedToNextMedia")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_DisconnectAdvancedToNextMedia(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "advancedToNextMedia")
	}
}

//export callbackQMediaGaplessPlaybackControlAdvancedToNextMedia
func callbackQMediaGaplessPlaybackControlAdvancedToNextMedia(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QMediaGaplessPlaybackControl::advancedToNextMedia")

	if signal := qt.GetSignal(C.GoString(ptrName), "advancedToNextMedia"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QMediaGaplessPlaybackControl) AdvancedToNextMedia() {
	defer qt.Recovering("QMediaGaplessPlaybackControl::advancedToNextMedia")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_AdvancedToNextMedia(ptr.Pointer())
	}
}

func (ptr *QMediaGaplessPlaybackControl) CrossfadeTime() float64 {
	defer qt.Recovering("QMediaGaplessPlaybackControl::crossfadeTime")

	if ptr.Pointer() != nil {
		return float64(C.QMediaGaplessPlaybackControl_CrossfadeTime(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaGaplessPlaybackControl) ConnectCrossfadeTimeChanged(f func(crossfadeTime float64)) {
	defer qt.Recovering("connect QMediaGaplessPlaybackControl::crossfadeTimeChanged")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_ConnectCrossfadeTimeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "crossfadeTimeChanged", f)
	}
}

func (ptr *QMediaGaplessPlaybackControl) DisconnectCrossfadeTimeChanged() {
	defer qt.Recovering("disconnect QMediaGaplessPlaybackControl::crossfadeTimeChanged")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_DisconnectCrossfadeTimeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "crossfadeTimeChanged")
	}
}

//export callbackQMediaGaplessPlaybackControlCrossfadeTimeChanged
func callbackQMediaGaplessPlaybackControlCrossfadeTimeChanged(ptr unsafe.Pointer, ptrName *C.char, crossfadeTime C.double) {
	defer qt.Recovering("callback QMediaGaplessPlaybackControl::crossfadeTimeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "crossfadeTimeChanged"); signal != nil {
		signal.(func(float64))(float64(crossfadeTime))
	}

}

func (ptr *QMediaGaplessPlaybackControl) CrossfadeTimeChanged(crossfadeTime float64) {
	defer qt.Recovering("QMediaGaplessPlaybackControl::crossfadeTimeChanged")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_CrossfadeTimeChanged(ptr.Pointer(), C.double(crossfadeTime))
	}
}

func (ptr *QMediaGaplessPlaybackControl) IsCrossfadeSupported() bool {
	defer qt.Recovering("QMediaGaplessPlaybackControl::isCrossfadeSupported")

	if ptr.Pointer() != nil {
		return C.QMediaGaplessPlaybackControl_IsCrossfadeSupported(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaGaplessPlaybackControl) NextMedia() *QMediaContent {
	defer qt.Recovering("QMediaGaplessPlaybackControl::nextMedia")

	if ptr.Pointer() != nil {
		return NewQMediaContentFromPointer(C.QMediaGaplessPlaybackControl_NextMedia(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaGaplessPlaybackControl) ConnectNextMediaChanged(f func(media *QMediaContent)) {
	defer qt.Recovering("connect QMediaGaplessPlaybackControl::nextMediaChanged")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_ConnectNextMediaChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "nextMediaChanged", f)
	}
}

func (ptr *QMediaGaplessPlaybackControl) DisconnectNextMediaChanged() {
	defer qt.Recovering("disconnect QMediaGaplessPlaybackControl::nextMediaChanged")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_DisconnectNextMediaChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "nextMediaChanged")
	}
}

//export callbackQMediaGaplessPlaybackControlNextMediaChanged
func callbackQMediaGaplessPlaybackControlNextMediaChanged(ptr unsafe.Pointer, ptrName *C.char, media unsafe.Pointer) {
	defer qt.Recovering("callback QMediaGaplessPlaybackControl::nextMediaChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "nextMediaChanged"); signal != nil {
		signal.(func(*QMediaContent))(NewQMediaContentFromPointer(media))
	}

}

func (ptr *QMediaGaplessPlaybackControl) NextMediaChanged(media QMediaContent_ITF) {
	defer qt.Recovering("QMediaGaplessPlaybackControl::nextMediaChanged")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_NextMediaChanged(ptr.Pointer(), PointerFromQMediaContent(media))
	}
}

func (ptr *QMediaGaplessPlaybackControl) SetCrossfadeTime(crossfadeTime float64) {
	defer qt.Recovering("QMediaGaplessPlaybackControl::setCrossfadeTime")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_SetCrossfadeTime(ptr.Pointer(), C.double(crossfadeTime))
	}
}

func (ptr *QMediaGaplessPlaybackControl) SetNextMedia(media QMediaContent_ITF) {
	defer qt.Recovering("QMediaGaplessPlaybackControl::setNextMedia")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_SetNextMedia(ptr.Pointer(), PointerFromQMediaContent(media))
	}
}

func (ptr *QMediaGaplessPlaybackControl) DestroyQMediaGaplessPlaybackControl() {
	defer qt.Recovering("QMediaGaplessPlaybackControl::~QMediaGaplessPlaybackControl")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_DestroyQMediaGaplessPlaybackControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMediaGaplessPlaybackControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMediaGaplessPlaybackControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMediaGaplessPlaybackControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMediaGaplessPlaybackControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMediaGaplessPlaybackControlTimerEvent
func callbackQMediaGaplessPlaybackControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaGaplessPlaybackControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMediaGaplessPlaybackControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMediaGaplessPlaybackControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaGaplessPlaybackControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaGaplessPlaybackControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaGaplessPlaybackControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaGaplessPlaybackControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMediaGaplessPlaybackControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMediaGaplessPlaybackControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMediaGaplessPlaybackControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMediaGaplessPlaybackControlChildEvent
func callbackQMediaGaplessPlaybackControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaGaplessPlaybackControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMediaGaplessPlaybackControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMediaGaplessPlaybackControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaGaplessPlaybackControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaGaplessPlaybackControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaGaplessPlaybackControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaGaplessPlaybackControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMediaGaplessPlaybackControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMediaGaplessPlaybackControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMediaGaplessPlaybackControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMediaGaplessPlaybackControlCustomEvent
func callbackQMediaGaplessPlaybackControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaGaplessPlaybackControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMediaGaplessPlaybackControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMediaGaplessPlaybackControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaGaplessPlaybackControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMediaGaplessPlaybackControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaGaplessPlaybackControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QMediaNetworkAccessControl struct {
	QMediaControl
}

type QMediaNetworkAccessControl_ITF interface {
	QMediaControl_ITF
	QMediaNetworkAccessControl_PTR() *QMediaNetworkAccessControl
}

func PointerFromQMediaNetworkAccessControl(ptr QMediaNetworkAccessControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaNetworkAccessControl_PTR().Pointer()
	}
	return nil
}

func NewQMediaNetworkAccessControlFromPointer(ptr unsafe.Pointer) *QMediaNetworkAccessControl {
	var n = new(QMediaNetworkAccessControl)
	n.SetPointer(ptr)
	return n
}

func newQMediaNetworkAccessControlFromPointer(ptr unsafe.Pointer) *QMediaNetworkAccessControl {
	var n = NewQMediaNetworkAccessControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QMediaNetworkAccessControl_") {
		n.SetObjectName("QMediaNetworkAccessControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaNetworkAccessControl) QMediaNetworkAccessControl_PTR() *QMediaNetworkAccessControl {
	return ptr
}

func (ptr *QMediaNetworkAccessControl) ConnectConfigurationChanged(f func(configuration *network.QNetworkConfiguration)) {
	defer qt.Recovering("connect QMediaNetworkAccessControl::configurationChanged")

	if ptr.Pointer() != nil {
		C.QMediaNetworkAccessControl_ConnectConfigurationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "configurationChanged", f)
	}
}

func (ptr *QMediaNetworkAccessControl) DisconnectConfigurationChanged() {
	defer qt.Recovering("disconnect QMediaNetworkAccessControl::configurationChanged")

	if ptr.Pointer() != nil {
		C.QMediaNetworkAccessControl_DisconnectConfigurationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "configurationChanged")
	}
}

//export callbackQMediaNetworkAccessControlConfigurationChanged
func callbackQMediaNetworkAccessControlConfigurationChanged(ptr unsafe.Pointer, ptrName *C.char, configuration unsafe.Pointer) {
	defer qt.Recovering("callback QMediaNetworkAccessControl::configurationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "configurationChanged"); signal != nil {
		signal.(func(*network.QNetworkConfiguration))(network.NewQNetworkConfigurationFromPointer(configuration))
	}

}

func (ptr *QMediaNetworkAccessControl) ConfigurationChanged(configuration network.QNetworkConfiguration_ITF) {
	defer qt.Recovering("QMediaNetworkAccessControl::configurationChanged")

	if ptr.Pointer() != nil {
		C.QMediaNetworkAccessControl_ConfigurationChanged(ptr.Pointer(), network.PointerFromQNetworkConfiguration(configuration))
	}
}

func (ptr *QMediaNetworkAccessControl) CurrentConfiguration() *network.QNetworkConfiguration {
	defer qt.Recovering("QMediaNetworkAccessControl::currentConfiguration")

	if ptr.Pointer() != nil {
		return network.NewQNetworkConfigurationFromPointer(C.QMediaNetworkAccessControl_CurrentConfiguration(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaNetworkAccessControl) DestroyQMediaNetworkAccessControl() {
	defer qt.Recovering("QMediaNetworkAccessControl::~QMediaNetworkAccessControl")

	if ptr.Pointer() != nil {
		C.QMediaNetworkAccessControl_DestroyQMediaNetworkAccessControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMediaNetworkAccessControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMediaNetworkAccessControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMediaNetworkAccessControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMediaNetworkAccessControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMediaNetworkAccessControlTimerEvent
func callbackQMediaNetworkAccessControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaNetworkAccessControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMediaNetworkAccessControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMediaNetworkAccessControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaNetworkAccessControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaNetworkAccessControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaNetworkAccessControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaNetworkAccessControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaNetworkAccessControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaNetworkAccessControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMediaNetworkAccessControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMediaNetworkAccessControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMediaNetworkAccessControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMediaNetworkAccessControlChildEvent
func callbackQMediaNetworkAccessControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaNetworkAccessControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMediaNetworkAccessControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMediaNetworkAccessControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaNetworkAccessControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaNetworkAccessControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaNetworkAccessControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaNetworkAccessControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaNetworkAccessControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaNetworkAccessControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMediaNetworkAccessControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMediaNetworkAccessControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMediaNetworkAccessControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMediaNetworkAccessControlCustomEvent
func callbackQMediaNetworkAccessControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaNetworkAccessControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMediaNetworkAccessControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMediaNetworkAccessControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaNetworkAccessControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaNetworkAccessControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMediaNetworkAccessControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaNetworkAccessControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaNetworkAccessControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QMediaObject struct {
	core.QObject
}

type QMediaObject_ITF interface {
	core.QObject_ITF
	QMediaObject_PTR() *QMediaObject
}

func PointerFromQMediaObject(ptr QMediaObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaObject_PTR().Pointer()
	}
	return nil
}

func NewQMediaObjectFromPointer(ptr unsafe.Pointer) *QMediaObject {
	var n = new(QMediaObject)
	n.SetPointer(ptr)
	return n
}

func newQMediaObjectFromPointer(ptr unsafe.Pointer) *QMediaObject {
	var n = NewQMediaObjectFromPointer(ptr)
	for len(n.ObjectName()) < len("QMediaObject_") {
		n.SetObjectName("QMediaObject_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaObject) QMediaObject_PTR() *QMediaObject {
	return ptr
}

func (ptr *QMediaObject) NotifyInterval() int {
	defer qt.Recovering("QMediaObject::notifyInterval")

	if ptr.Pointer() != nil {
		return int(C.QMediaObject_NotifyInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaObject) SetNotifyInterval(milliSeconds int) {
	defer qt.Recovering("QMediaObject::setNotifyInterval")

	if ptr.Pointer() != nil {
		C.QMediaObject_SetNotifyInterval(ptr.Pointer(), C.int(milliSeconds))
	}
}

func (ptr *QMediaObject) Availability() QMultimedia__AvailabilityStatus {
	defer qt.Recovering("QMediaObject::availability")

	if ptr.Pointer() != nil {
		return QMultimedia__AvailabilityStatus(C.QMediaObject_Availability(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaObject) ConnectAvailabilityChanged2(f func(availability QMultimedia__AvailabilityStatus)) {
	defer qt.Recovering("connect QMediaObject::availabilityChanged")

	if ptr.Pointer() != nil {
		C.QMediaObject_ConnectAvailabilityChanged2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "availabilityChanged2", f)
	}
}

func (ptr *QMediaObject) DisconnectAvailabilityChanged2() {
	defer qt.Recovering("disconnect QMediaObject::availabilityChanged")

	if ptr.Pointer() != nil {
		C.QMediaObject_DisconnectAvailabilityChanged2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "availabilityChanged2")
	}
}

//export callbackQMediaObjectAvailabilityChanged2
func callbackQMediaObjectAvailabilityChanged2(ptr unsafe.Pointer, ptrName *C.char, availability C.int) {
	defer qt.Recovering("callback QMediaObject::availabilityChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "availabilityChanged2"); signal != nil {
		signal.(func(QMultimedia__AvailabilityStatus))(QMultimedia__AvailabilityStatus(availability))
	}

}

func (ptr *QMediaObject) AvailabilityChanged2(availability QMultimedia__AvailabilityStatus) {
	defer qt.Recovering("QMediaObject::availabilityChanged")

	if ptr.Pointer() != nil {
		C.QMediaObject_AvailabilityChanged2(ptr.Pointer(), C.int(availability))
	}
}

func (ptr *QMediaObject) ConnectAvailabilityChanged(f func(available bool)) {
	defer qt.Recovering("connect QMediaObject::availabilityChanged")

	if ptr.Pointer() != nil {
		C.QMediaObject_ConnectAvailabilityChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "availabilityChanged", f)
	}
}

func (ptr *QMediaObject) DisconnectAvailabilityChanged() {
	defer qt.Recovering("disconnect QMediaObject::availabilityChanged")

	if ptr.Pointer() != nil {
		C.QMediaObject_DisconnectAvailabilityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "availabilityChanged")
	}
}

//export callbackQMediaObjectAvailabilityChanged
func callbackQMediaObjectAvailabilityChanged(ptr unsafe.Pointer, ptrName *C.char, available C.int) {
	defer qt.Recovering("callback QMediaObject::availabilityChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "availabilityChanged"); signal != nil {
		signal.(func(bool))(int(available) != 0)
	}

}

func (ptr *QMediaObject) AvailabilityChanged(available bool) {
	defer qt.Recovering("QMediaObject::availabilityChanged")

	if ptr.Pointer() != nil {
		C.QMediaObject_AvailabilityChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(available)))
	}
}

func (ptr *QMediaObject) AvailableMetaData() []string {
	defer qt.Recovering("QMediaObject::availableMetaData")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMediaObject_AvailableMetaData(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QMediaObject) Bind(object core.QObject_ITF) bool {
	defer qt.Recovering("QMediaObject::bind")

	if ptr.Pointer() != nil {
		return C.QMediaObject_Bind(ptr.Pointer(), core.PointerFromQObject(object)) != 0
	}
	return false
}

func (ptr *QMediaObject) IsAvailable() bool {
	defer qt.Recovering("QMediaObject::isAvailable")

	if ptr.Pointer() != nil {
		return C.QMediaObject_IsAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaObject) IsMetaDataAvailable() bool {
	defer qt.Recovering("QMediaObject::isMetaDataAvailable")

	if ptr.Pointer() != nil {
		return C.QMediaObject_IsMetaDataAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaObject) MetaData(key string) *core.QVariant {
	defer qt.Recovering("QMediaObject::metaData")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QMediaObject_MetaData(ptr.Pointer(), C.CString(key)))
	}
	return nil
}

func (ptr *QMediaObject) ConnectMetaDataAvailableChanged(f func(available bool)) {
	defer qt.Recovering("connect QMediaObject::metaDataAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMediaObject_ConnectMetaDataAvailableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataAvailableChanged", f)
	}
}

func (ptr *QMediaObject) DisconnectMetaDataAvailableChanged() {
	defer qt.Recovering("disconnect QMediaObject::metaDataAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMediaObject_DisconnectMetaDataAvailableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataAvailableChanged")
	}
}

//export callbackQMediaObjectMetaDataAvailableChanged
func callbackQMediaObjectMetaDataAvailableChanged(ptr unsafe.Pointer, ptrName *C.char, available C.int) {
	defer qt.Recovering("callback QMediaObject::metaDataAvailableChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaDataAvailableChanged"); signal != nil {
		signal.(func(bool))(int(available) != 0)
	}

}

func (ptr *QMediaObject) MetaDataAvailableChanged(available bool) {
	defer qt.Recovering("QMediaObject::metaDataAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMediaObject_MetaDataAvailableChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(available)))
	}
}

func (ptr *QMediaObject) ConnectMetaDataChanged(f func()) {
	defer qt.Recovering("connect QMediaObject::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMediaObject_ConnectMetaDataChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataChanged", f)
	}
}

func (ptr *QMediaObject) DisconnectMetaDataChanged() {
	defer qt.Recovering("disconnect QMediaObject::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMediaObject_DisconnectMetaDataChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataChanged")
	}
}

//export callbackQMediaObjectMetaDataChanged
func callbackQMediaObjectMetaDataChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QMediaObject::metaDataChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaDataChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QMediaObject) MetaDataChanged() {
	defer qt.Recovering("QMediaObject::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMediaObject_MetaDataChanged(ptr.Pointer())
	}
}

func (ptr *QMediaObject) ConnectMetaDataChanged2(f func(key string, value *core.QVariant)) {
	defer qt.Recovering("connect QMediaObject::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMediaObject_ConnectMetaDataChanged2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataChanged2", f)
	}
}

func (ptr *QMediaObject) DisconnectMetaDataChanged2() {
	defer qt.Recovering("disconnect QMediaObject::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMediaObject_DisconnectMetaDataChanged2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataChanged2")
	}
}

//export callbackQMediaObjectMetaDataChanged2
func callbackQMediaObjectMetaDataChanged2(ptr unsafe.Pointer, ptrName *C.char, key *C.char, value unsafe.Pointer) {
	defer qt.Recovering("callback QMediaObject::metaDataChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaDataChanged2"); signal != nil {
		signal.(func(string, *core.QVariant))(C.GoString(key), core.NewQVariantFromPointer(value))
	}

}

func (ptr *QMediaObject) MetaDataChanged2(key string, value core.QVariant_ITF) {
	defer qt.Recovering("QMediaObject::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMediaObject_MetaDataChanged2(ptr.Pointer(), C.CString(key), core.PointerFromQVariant(value))
	}
}

func (ptr *QMediaObject) ConnectNotifyIntervalChanged(f func(milliseconds int)) {
	defer qt.Recovering("connect QMediaObject::notifyIntervalChanged")

	if ptr.Pointer() != nil {
		C.QMediaObject_ConnectNotifyIntervalChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "notifyIntervalChanged", f)
	}
}

func (ptr *QMediaObject) DisconnectNotifyIntervalChanged() {
	defer qt.Recovering("disconnect QMediaObject::notifyIntervalChanged")

	if ptr.Pointer() != nil {
		C.QMediaObject_DisconnectNotifyIntervalChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "notifyIntervalChanged")
	}
}

//export callbackQMediaObjectNotifyIntervalChanged
func callbackQMediaObjectNotifyIntervalChanged(ptr unsafe.Pointer, ptrName *C.char, milliseconds C.int) {
	defer qt.Recovering("callback QMediaObject::notifyIntervalChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "notifyIntervalChanged"); signal != nil {
		signal.(func(int))(int(milliseconds))
	}

}

func (ptr *QMediaObject) NotifyIntervalChanged(milliseconds int) {
	defer qt.Recovering("QMediaObject::notifyIntervalChanged")

	if ptr.Pointer() != nil {
		C.QMediaObject_NotifyIntervalChanged(ptr.Pointer(), C.int(milliseconds))
	}
}

func (ptr *QMediaObject) Service() *QMediaService {
	defer qt.Recovering("QMediaObject::service")

	if ptr.Pointer() != nil {
		return NewQMediaServiceFromPointer(C.QMediaObject_Service(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaObject) ConnectUnbind(f func(object *core.QObject)) {
	defer qt.Recovering("connect QMediaObject::unbind")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "unbind", f)
	}
}

func (ptr *QMediaObject) DisconnectUnbind() {
	defer qt.Recovering("disconnect QMediaObject::unbind")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "unbind")
	}
}

//export callbackQMediaObjectUnbind
func callbackQMediaObjectUnbind(ptr unsafe.Pointer, ptrName *C.char, object unsafe.Pointer) {
	defer qt.Recovering("callback QMediaObject::unbind")

	if signal := qt.GetSignal(C.GoString(ptrName), "unbind"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(object))
	} else {
		NewQMediaObjectFromPointer(ptr).UnbindDefault(core.NewQObjectFromPointer(object))
	}
}

func (ptr *QMediaObject) Unbind(object core.QObject_ITF) {
	defer qt.Recovering("QMediaObject::unbind")

	if ptr.Pointer() != nil {
		C.QMediaObject_Unbind(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QMediaObject) UnbindDefault(object core.QObject_ITF) {
	defer qt.Recovering("QMediaObject::unbind")

	if ptr.Pointer() != nil {
		C.QMediaObject_UnbindDefault(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QMediaObject) DestroyQMediaObject() {
	defer qt.Recovering("QMediaObject::~QMediaObject")

	if ptr.Pointer() != nil {
		C.QMediaObject_DestroyQMediaObject(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMediaObject) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMediaObject::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMediaObject) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMediaObject::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMediaObjectTimerEvent
func callbackQMediaObjectTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaObject::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMediaObjectFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMediaObject) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaObject::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaObject_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaObject) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaObject::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaObject_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaObject) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMediaObject::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMediaObject) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMediaObject::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMediaObjectChildEvent
func callbackQMediaObjectChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaObject::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMediaObjectFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMediaObject) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaObject::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaObject_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaObject) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaObject::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaObject_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaObject) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMediaObject::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMediaObject) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMediaObject::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMediaObjectCustomEvent
func callbackQMediaObjectCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaObject::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMediaObjectFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMediaObject) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaObject::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaObject_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMediaObject) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaObject::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaObject_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QMediaPlayer struct {
	QMediaObject
}

type QMediaPlayer_ITF interface {
	QMediaObject_ITF
	QMediaPlayer_PTR() *QMediaPlayer
}

func PointerFromQMediaPlayer(ptr QMediaPlayer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaPlayer_PTR().Pointer()
	}
	return nil
}

func NewQMediaPlayerFromPointer(ptr unsafe.Pointer) *QMediaPlayer {
	var n = new(QMediaPlayer)
	n.SetPointer(ptr)
	return n
}

func newQMediaPlayerFromPointer(ptr unsafe.Pointer) *QMediaPlayer {
	var n = NewQMediaPlayerFromPointer(ptr)
	for len(n.ObjectName()) < len("QMediaPlayer_") {
		n.SetObjectName("QMediaPlayer_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaPlayer) QMediaPlayer_PTR() *QMediaPlayer {
	return ptr
}

//QMediaPlayer::Error
type QMediaPlayer__Error int64

const (
	QMediaPlayer__NoError             = QMediaPlayer__Error(0)
	QMediaPlayer__ResourceError       = QMediaPlayer__Error(1)
	QMediaPlayer__FormatError         = QMediaPlayer__Error(2)
	QMediaPlayer__NetworkError        = QMediaPlayer__Error(3)
	QMediaPlayer__AccessDeniedError   = QMediaPlayer__Error(4)
	QMediaPlayer__ServiceMissingError = QMediaPlayer__Error(5)
	QMediaPlayer__MediaIsPlaylist     = QMediaPlayer__Error(6)
)

//QMediaPlayer::Flag
type QMediaPlayer__Flag int64

const (
	QMediaPlayer__LowLatency     = QMediaPlayer__Flag(0x01)
	QMediaPlayer__StreamPlayback = QMediaPlayer__Flag(0x02)
	QMediaPlayer__VideoSurface   = QMediaPlayer__Flag(0x04)
)

//QMediaPlayer::MediaStatus
type QMediaPlayer__MediaStatus int64

const (
	QMediaPlayer__UnknownMediaStatus = QMediaPlayer__MediaStatus(0)
	QMediaPlayer__NoMedia            = QMediaPlayer__MediaStatus(1)
	QMediaPlayer__LoadingMedia       = QMediaPlayer__MediaStatus(2)
	QMediaPlayer__LoadedMedia        = QMediaPlayer__MediaStatus(3)
	QMediaPlayer__StalledMedia       = QMediaPlayer__MediaStatus(4)
	QMediaPlayer__BufferingMedia     = QMediaPlayer__MediaStatus(5)
	QMediaPlayer__BufferedMedia      = QMediaPlayer__MediaStatus(6)
	QMediaPlayer__EndOfMedia         = QMediaPlayer__MediaStatus(7)
	QMediaPlayer__InvalidMedia       = QMediaPlayer__MediaStatus(8)
)

//QMediaPlayer::State
type QMediaPlayer__State int64

const (
	QMediaPlayer__StoppedState = QMediaPlayer__State(0)
	QMediaPlayer__PlayingState = QMediaPlayer__State(1)
	QMediaPlayer__PausedState  = QMediaPlayer__State(2)
)

func (ptr *QMediaPlayer) BufferStatus() int {
	defer qt.Recovering("QMediaPlayer::bufferStatus")

	if ptr.Pointer() != nil {
		return int(C.QMediaPlayer_BufferStatus(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayer) CurrentMedia() *QMediaContent {
	defer qt.Recovering("QMediaPlayer::currentMedia")

	if ptr.Pointer() != nil {
		return NewQMediaContentFromPointer(C.QMediaPlayer_CurrentMedia(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaPlayer) Duration() int64 {
	defer qt.Recovering("QMediaPlayer::duration")

	if ptr.Pointer() != nil {
		return int64(C.QMediaPlayer_Duration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayer) ErrorString() string {
	defer qt.Recovering("QMediaPlayer::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaPlayer_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaPlayer) IsAudioAvailable() bool {
	defer qt.Recovering("QMediaPlayer::isAudioAvailable")

	if ptr.Pointer() != nil {
		return C.QMediaPlayer_IsAudioAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlayer) IsMuted() bool {
	defer qt.Recovering("QMediaPlayer::isMuted")

	if ptr.Pointer() != nil {
		return C.QMediaPlayer_IsMuted(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlayer) IsSeekable() bool {
	defer qt.Recovering("QMediaPlayer::isSeekable")

	if ptr.Pointer() != nil {
		return C.QMediaPlayer_IsSeekable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlayer) IsVideoAvailable() bool {
	defer qt.Recovering("QMediaPlayer::isVideoAvailable")

	if ptr.Pointer() != nil {
		return C.QMediaPlayer_IsVideoAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlayer) Media() *QMediaContent {
	defer qt.Recovering("QMediaPlayer::media")

	if ptr.Pointer() != nil {
		return NewQMediaContentFromPointer(C.QMediaPlayer_Media(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaPlayer) MediaStatus() QMediaPlayer__MediaStatus {
	defer qt.Recovering("QMediaPlayer::mediaStatus")

	if ptr.Pointer() != nil {
		return QMediaPlayer__MediaStatus(C.QMediaPlayer_MediaStatus(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayer) PlaybackRate() float64 {
	defer qt.Recovering("QMediaPlayer::playbackRate")

	if ptr.Pointer() != nil {
		return float64(C.QMediaPlayer_PlaybackRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayer) Playlist() *QMediaPlaylist {
	defer qt.Recovering("QMediaPlayer::playlist")

	if ptr.Pointer() != nil {
		return NewQMediaPlaylistFromPointer(C.QMediaPlayer_Playlist(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaPlayer) Position() int64 {
	defer qt.Recovering("QMediaPlayer::position")

	if ptr.Pointer() != nil {
		return int64(C.QMediaPlayer_Position(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayer) SetMuted(muted bool) {
	defer qt.Recovering("QMediaPlayer::setMuted")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_SetMuted(ptr.Pointer(), C.int(qt.GoBoolToInt(muted)))
	}
}

func (ptr *QMediaPlayer) SetPlaybackRate(rate float64) {
	defer qt.Recovering("QMediaPlayer::setPlaybackRate")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_SetPlaybackRate(ptr.Pointer(), C.double(rate))
	}
}

func (ptr *QMediaPlayer) SetPlaylist(playlist QMediaPlaylist_ITF) {
	defer qt.Recovering("QMediaPlayer::setPlaylist")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_SetPlaylist(ptr.Pointer(), PointerFromQMediaPlaylist(playlist))
	}
}

func (ptr *QMediaPlayer) SetPosition(position int64) {
	defer qt.Recovering("QMediaPlayer::setPosition")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_SetPosition(ptr.Pointer(), C.longlong(position))
	}
}

func (ptr *QMediaPlayer) SetVideoOutput2(output unsafe.Pointer) {
	defer qt.Recovering("QMediaPlayer::setVideoOutput")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_SetVideoOutput2(ptr.Pointer(), output)
	}
}

func (ptr *QMediaPlayer) SetVideoOutput(output unsafe.Pointer) {
	defer qt.Recovering("QMediaPlayer::setVideoOutput")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_SetVideoOutput(ptr.Pointer(), output)
	}
}

func (ptr *QMediaPlayer) SetVolume(volume int) {
	defer qt.Recovering("QMediaPlayer::setVolume")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_SetVolume(ptr.Pointer(), C.int(volume))
	}
}

func (ptr *QMediaPlayer) State() QMediaPlayer__State {
	defer qt.Recovering("QMediaPlayer::state")

	if ptr.Pointer() != nil {
		return QMediaPlayer__State(C.QMediaPlayer_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayer) Volume() int {
	defer qt.Recovering("QMediaPlayer::volume")

	if ptr.Pointer() != nil {
		return int(C.QMediaPlayer_Volume(ptr.Pointer()))
	}
	return 0
}

func NewQMediaPlayer(parent core.QObject_ITF, flags QMediaPlayer__Flag) *QMediaPlayer {
	defer qt.Recovering("QMediaPlayer::QMediaPlayer")

	return newQMediaPlayerFromPointer(C.QMediaPlayer_NewQMediaPlayer(core.PointerFromQObject(parent), C.int(flags)))
}

func (ptr *QMediaPlayer) ConnectAudioAvailableChanged(f func(available bool)) {
	defer qt.Recovering("connect QMediaPlayer::audioAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_ConnectAudioAvailableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "audioAvailableChanged", f)
	}
}

func (ptr *QMediaPlayer) DisconnectAudioAvailableChanged() {
	defer qt.Recovering("disconnect QMediaPlayer::audioAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_DisconnectAudioAvailableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "audioAvailableChanged")
	}
}

//export callbackQMediaPlayerAudioAvailableChanged
func callbackQMediaPlayerAudioAvailableChanged(ptr unsafe.Pointer, ptrName *C.char, available C.int) {
	defer qt.Recovering("callback QMediaPlayer::audioAvailableChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "audioAvailableChanged"); signal != nil {
		signal.(func(bool))(int(available) != 0)
	}

}

func (ptr *QMediaPlayer) AudioAvailableChanged(available bool) {
	defer qt.Recovering("QMediaPlayer::audioAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_AudioAvailableChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(available)))
	}
}

func (ptr *QMediaPlayer) Availability() QMultimedia__AvailabilityStatus {
	defer qt.Recovering("QMediaPlayer::availability")

	if ptr.Pointer() != nil {
		return QMultimedia__AvailabilityStatus(C.QMediaPlayer_Availability(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayer) ConnectBufferStatusChanged(f func(percentFilled int)) {
	defer qt.Recovering("connect QMediaPlayer::bufferStatusChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_ConnectBufferStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bufferStatusChanged", f)
	}
}

func (ptr *QMediaPlayer) DisconnectBufferStatusChanged() {
	defer qt.Recovering("disconnect QMediaPlayer::bufferStatusChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_DisconnectBufferStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bufferStatusChanged")
	}
}

//export callbackQMediaPlayerBufferStatusChanged
func callbackQMediaPlayerBufferStatusChanged(ptr unsafe.Pointer, ptrName *C.char, percentFilled C.int) {
	defer qt.Recovering("callback QMediaPlayer::bufferStatusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "bufferStatusChanged"); signal != nil {
		signal.(func(int))(int(percentFilled))
	}

}

func (ptr *QMediaPlayer) BufferStatusChanged(percentFilled int) {
	defer qt.Recovering("QMediaPlayer::bufferStatusChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_BufferStatusChanged(ptr.Pointer(), C.int(percentFilled))
	}
}

func (ptr *QMediaPlayer) ConnectCurrentMediaChanged(f func(media *QMediaContent)) {
	defer qt.Recovering("connect QMediaPlayer::currentMediaChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_ConnectCurrentMediaChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentMediaChanged", f)
	}
}

func (ptr *QMediaPlayer) DisconnectCurrentMediaChanged() {
	defer qt.Recovering("disconnect QMediaPlayer::currentMediaChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_DisconnectCurrentMediaChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentMediaChanged")
	}
}

//export callbackQMediaPlayerCurrentMediaChanged
func callbackQMediaPlayerCurrentMediaChanged(ptr unsafe.Pointer, ptrName *C.char, media unsafe.Pointer) {
	defer qt.Recovering("callback QMediaPlayer::currentMediaChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentMediaChanged"); signal != nil {
		signal.(func(*QMediaContent))(NewQMediaContentFromPointer(media))
	}

}

func (ptr *QMediaPlayer) CurrentMediaChanged(media QMediaContent_ITF) {
	defer qt.Recovering("QMediaPlayer::currentMediaChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_CurrentMediaChanged(ptr.Pointer(), PointerFromQMediaContent(media))
	}
}

func (ptr *QMediaPlayer) CurrentNetworkConfiguration() *network.QNetworkConfiguration {
	defer qt.Recovering("QMediaPlayer::currentNetworkConfiguration")

	if ptr.Pointer() != nil {
		return network.NewQNetworkConfigurationFromPointer(C.QMediaPlayer_CurrentNetworkConfiguration(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaPlayer) ConnectDurationChanged(f func(duration int64)) {
	defer qt.Recovering("connect QMediaPlayer::durationChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_ConnectDurationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "durationChanged", f)
	}
}

func (ptr *QMediaPlayer) DisconnectDurationChanged() {
	defer qt.Recovering("disconnect QMediaPlayer::durationChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_DisconnectDurationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "durationChanged")
	}
}

//export callbackQMediaPlayerDurationChanged
func callbackQMediaPlayerDurationChanged(ptr unsafe.Pointer, ptrName *C.char, duration C.longlong) {
	defer qt.Recovering("callback QMediaPlayer::durationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "durationChanged"); signal != nil {
		signal.(func(int64))(int64(duration))
	}

}

func (ptr *QMediaPlayer) DurationChanged(duration int64) {
	defer qt.Recovering("QMediaPlayer::durationChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_DurationChanged(ptr.Pointer(), C.longlong(duration))
	}
}

func (ptr *QMediaPlayer) ConnectError2(f func(error QMediaPlayer__Error)) {
	defer qt.Recovering("connect QMediaPlayer::error")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QMediaPlayer) DisconnectError2() {
	defer qt.Recovering("disconnect QMediaPlayer::error")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQMediaPlayerError2
func callbackQMediaPlayerError2(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QMediaPlayer::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QMediaPlayer__Error))(QMediaPlayer__Error(error))
	}

}

func (ptr *QMediaPlayer) Error2(error QMediaPlayer__Error) {
	defer qt.Recovering("QMediaPlayer::error")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_Error2(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QMediaPlayer) Error() QMediaPlayer__Error {
	defer qt.Recovering("QMediaPlayer::error")

	if ptr.Pointer() != nil {
		return QMediaPlayer__Error(C.QMediaPlayer_Error(ptr.Pointer()))
	}
	return 0
}

func QMediaPlayer_HasSupport(mimeType string, codecs []string, flags QMediaPlayer__Flag) QMultimedia__SupportEstimate {
	defer qt.Recovering("QMediaPlayer::hasSupport")

	return QMultimedia__SupportEstimate(C.QMediaPlayer_QMediaPlayer_HasSupport(C.CString(mimeType), C.CString(strings.Join(codecs, "|")), C.int(flags)))
}

func (ptr *QMediaPlayer) ConnectMediaChanged(f func(media *QMediaContent)) {
	defer qt.Recovering("connect QMediaPlayer::mediaChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_ConnectMediaChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mediaChanged", f)
	}
}

func (ptr *QMediaPlayer) DisconnectMediaChanged() {
	defer qt.Recovering("disconnect QMediaPlayer::mediaChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_DisconnectMediaChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mediaChanged")
	}
}

//export callbackQMediaPlayerMediaChanged
func callbackQMediaPlayerMediaChanged(ptr unsafe.Pointer, ptrName *C.char, media unsafe.Pointer) {
	defer qt.Recovering("callback QMediaPlayer::mediaChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "mediaChanged"); signal != nil {
		signal.(func(*QMediaContent))(NewQMediaContentFromPointer(media))
	}

}

func (ptr *QMediaPlayer) MediaChanged(media QMediaContent_ITF) {
	defer qt.Recovering("QMediaPlayer::mediaChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_MediaChanged(ptr.Pointer(), PointerFromQMediaContent(media))
	}
}

func (ptr *QMediaPlayer) ConnectMediaStatusChanged(f func(status QMediaPlayer__MediaStatus)) {
	defer qt.Recovering("connect QMediaPlayer::mediaStatusChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_ConnectMediaStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mediaStatusChanged", f)
	}
}

func (ptr *QMediaPlayer) DisconnectMediaStatusChanged() {
	defer qt.Recovering("disconnect QMediaPlayer::mediaStatusChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_DisconnectMediaStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mediaStatusChanged")
	}
}

//export callbackQMediaPlayerMediaStatusChanged
func callbackQMediaPlayerMediaStatusChanged(ptr unsafe.Pointer, ptrName *C.char, status C.int) {
	defer qt.Recovering("callback QMediaPlayer::mediaStatusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "mediaStatusChanged"); signal != nil {
		signal.(func(QMediaPlayer__MediaStatus))(QMediaPlayer__MediaStatus(status))
	}

}

func (ptr *QMediaPlayer) MediaStatusChanged(status QMediaPlayer__MediaStatus) {
	defer qt.Recovering("QMediaPlayer::mediaStatusChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_MediaStatusChanged(ptr.Pointer(), C.int(status))
	}
}

func (ptr *QMediaPlayer) MediaStream() *core.QIODevice {
	defer qt.Recovering("QMediaPlayer::mediaStream")

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QMediaPlayer_MediaStream(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaPlayer) ConnectMutedChanged(f func(muted bool)) {
	defer qt.Recovering("connect QMediaPlayer::mutedChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_ConnectMutedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mutedChanged", f)
	}
}

func (ptr *QMediaPlayer) DisconnectMutedChanged() {
	defer qt.Recovering("disconnect QMediaPlayer::mutedChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_DisconnectMutedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mutedChanged")
	}
}

//export callbackQMediaPlayerMutedChanged
func callbackQMediaPlayerMutedChanged(ptr unsafe.Pointer, ptrName *C.char, muted C.int) {
	defer qt.Recovering("callback QMediaPlayer::mutedChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "mutedChanged"); signal != nil {
		signal.(func(bool))(int(muted) != 0)
	}

}

func (ptr *QMediaPlayer) MutedChanged(muted bool) {
	defer qt.Recovering("QMediaPlayer::mutedChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_MutedChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(muted)))
	}
}

func (ptr *QMediaPlayer) ConnectNetworkConfigurationChanged(f func(configuration *network.QNetworkConfiguration)) {
	defer qt.Recovering("connect QMediaPlayer::networkConfigurationChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_ConnectNetworkConfigurationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "networkConfigurationChanged", f)
	}
}

func (ptr *QMediaPlayer) DisconnectNetworkConfigurationChanged() {
	defer qt.Recovering("disconnect QMediaPlayer::networkConfigurationChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_DisconnectNetworkConfigurationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "networkConfigurationChanged")
	}
}

//export callbackQMediaPlayerNetworkConfigurationChanged
func callbackQMediaPlayerNetworkConfigurationChanged(ptr unsafe.Pointer, ptrName *C.char, configuration unsafe.Pointer) {
	defer qt.Recovering("callback QMediaPlayer::networkConfigurationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "networkConfigurationChanged"); signal != nil {
		signal.(func(*network.QNetworkConfiguration))(network.NewQNetworkConfigurationFromPointer(configuration))
	}

}

func (ptr *QMediaPlayer) NetworkConfigurationChanged(configuration network.QNetworkConfiguration_ITF) {
	defer qt.Recovering("QMediaPlayer::networkConfigurationChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_NetworkConfigurationChanged(ptr.Pointer(), network.PointerFromQNetworkConfiguration(configuration))
	}
}

func (ptr *QMediaPlayer) Pause() {
	defer qt.Recovering("QMediaPlayer::pause")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_Pause(ptr.Pointer())
	}
}

func (ptr *QMediaPlayer) Play() {
	defer qt.Recovering("QMediaPlayer::play")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_Play(ptr.Pointer())
	}
}

func (ptr *QMediaPlayer) ConnectPlaybackRateChanged(f func(rate float64)) {
	defer qt.Recovering("connect QMediaPlayer::playbackRateChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_ConnectPlaybackRateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "playbackRateChanged", f)
	}
}

func (ptr *QMediaPlayer) DisconnectPlaybackRateChanged() {
	defer qt.Recovering("disconnect QMediaPlayer::playbackRateChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_DisconnectPlaybackRateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "playbackRateChanged")
	}
}

//export callbackQMediaPlayerPlaybackRateChanged
func callbackQMediaPlayerPlaybackRateChanged(ptr unsafe.Pointer, ptrName *C.char, rate C.double) {
	defer qt.Recovering("callback QMediaPlayer::playbackRateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "playbackRateChanged"); signal != nil {
		signal.(func(float64))(float64(rate))
	}

}

func (ptr *QMediaPlayer) PlaybackRateChanged(rate float64) {
	defer qt.Recovering("QMediaPlayer::playbackRateChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_PlaybackRateChanged(ptr.Pointer(), C.double(rate))
	}
}

func (ptr *QMediaPlayer) ConnectPositionChanged(f func(position int64)) {
	defer qt.Recovering("connect QMediaPlayer::positionChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_ConnectPositionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "positionChanged", f)
	}
}

func (ptr *QMediaPlayer) DisconnectPositionChanged() {
	defer qt.Recovering("disconnect QMediaPlayer::positionChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_DisconnectPositionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "positionChanged")
	}
}

//export callbackQMediaPlayerPositionChanged
func callbackQMediaPlayerPositionChanged(ptr unsafe.Pointer, ptrName *C.char, position C.longlong) {
	defer qt.Recovering("callback QMediaPlayer::positionChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "positionChanged"); signal != nil {
		signal.(func(int64))(int64(position))
	}

}

func (ptr *QMediaPlayer) PositionChanged(position int64) {
	defer qt.Recovering("QMediaPlayer::positionChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_PositionChanged(ptr.Pointer(), C.longlong(position))
	}
}

func (ptr *QMediaPlayer) ConnectSeekableChanged(f func(seekable bool)) {
	defer qt.Recovering("connect QMediaPlayer::seekableChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_ConnectSeekableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "seekableChanged", f)
	}
}

func (ptr *QMediaPlayer) DisconnectSeekableChanged() {
	defer qt.Recovering("disconnect QMediaPlayer::seekableChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_DisconnectSeekableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "seekableChanged")
	}
}

//export callbackQMediaPlayerSeekableChanged
func callbackQMediaPlayerSeekableChanged(ptr unsafe.Pointer, ptrName *C.char, seekable C.int) {
	defer qt.Recovering("callback QMediaPlayer::seekableChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "seekableChanged"); signal != nil {
		signal.(func(bool))(int(seekable) != 0)
	}

}

func (ptr *QMediaPlayer) SeekableChanged(seekable bool) {
	defer qt.Recovering("QMediaPlayer::seekableChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_SeekableChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(seekable)))
	}
}

func (ptr *QMediaPlayer) SetMedia(media QMediaContent_ITF, stream core.QIODevice_ITF) {
	defer qt.Recovering("QMediaPlayer::setMedia")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_SetMedia(ptr.Pointer(), PointerFromQMediaContent(media), core.PointerFromQIODevice(stream))
	}
}

func (ptr *QMediaPlayer) SetVideoOutput3(surface QAbstractVideoSurface_ITF) {
	defer qt.Recovering("QMediaPlayer::setVideoOutput")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_SetVideoOutput3(ptr.Pointer(), PointerFromQAbstractVideoSurface(surface))
	}
}

func (ptr *QMediaPlayer) ConnectStateChanged(f func(state QMediaPlayer__State)) {
	defer qt.Recovering("connect QMediaPlayer::stateChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QMediaPlayer) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QMediaPlayer::stateChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQMediaPlayerStateChanged
func callbackQMediaPlayerStateChanged(ptr unsafe.Pointer, ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QMediaPlayer::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QMediaPlayer__State))(QMediaPlayer__State(state))
	}

}

func (ptr *QMediaPlayer) StateChanged(state QMediaPlayer__State) {
	defer qt.Recovering("QMediaPlayer::stateChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_StateChanged(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QMediaPlayer) Stop() {
	defer qt.Recovering("QMediaPlayer::stop")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_Stop(ptr.Pointer())
	}
}

func (ptr *QMediaPlayer) ConnectVideoAvailableChanged(f func(videoAvailable bool)) {
	defer qt.Recovering("connect QMediaPlayer::videoAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_ConnectVideoAvailableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "videoAvailableChanged", f)
	}
}

func (ptr *QMediaPlayer) DisconnectVideoAvailableChanged() {
	defer qt.Recovering("disconnect QMediaPlayer::videoAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_DisconnectVideoAvailableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "videoAvailableChanged")
	}
}

//export callbackQMediaPlayerVideoAvailableChanged
func callbackQMediaPlayerVideoAvailableChanged(ptr unsafe.Pointer, ptrName *C.char, videoAvailable C.int) {
	defer qt.Recovering("callback QMediaPlayer::videoAvailableChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "videoAvailableChanged"); signal != nil {
		signal.(func(bool))(int(videoAvailable) != 0)
	}

}

func (ptr *QMediaPlayer) VideoAvailableChanged(videoAvailable bool) {
	defer qt.Recovering("QMediaPlayer::videoAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_VideoAvailableChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(videoAvailable)))
	}
}

func (ptr *QMediaPlayer) ConnectVolumeChanged(f func(volume int)) {
	defer qt.Recovering("connect QMediaPlayer::volumeChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_ConnectVolumeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "volumeChanged", f)
	}
}

func (ptr *QMediaPlayer) DisconnectVolumeChanged() {
	defer qt.Recovering("disconnect QMediaPlayer::volumeChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_DisconnectVolumeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "volumeChanged")
	}
}

//export callbackQMediaPlayerVolumeChanged
func callbackQMediaPlayerVolumeChanged(ptr unsafe.Pointer, ptrName *C.char, volume C.int) {
	defer qt.Recovering("callback QMediaPlayer::volumeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "volumeChanged"); signal != nil {
		signal.(func(int))(int(volume))
	}

}

func (ptr *QMediaPlayer) VolumeChanged(volume int) {
	defer qt.Recovering("QMediaPlayer::volumeChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_VolumeChanged(ptr.Pointer(), C.int(volume))
	}
}

func (ptr *QMediaPlayer) DestroyQMediaPlayer() {
	defer qt.Recovering("QMediaPlayer::~QMediaPlayer")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_DestroyQMediaPlayer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMediaPlayer) ConnectUnbind(f func(object *core.QObject)) {
	defer qt.Recovering("connect QMediaPlayer::unbind")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "unbind", f)
	}
}

func (ptr *QMediaPlayer) DisconnectUnbind() {
	defer qt.Recovering("disconnect QMediaPlayer::unbind")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "unbind")
	}
}

//export callbackQMediaPlayerUnbind
func callbackQMediaPlayerUnbind(ptr unsafe.Pointer, ptrName *C.char, object unsafe.Pointer) {
	defer qt.Recovering("callback QMediaPlayer::unbind")

	if signal := qt.GetSignal(C.GoString(ptrName), "unbind"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(object))
	} else {
		NewQMediaPlayerFromPointer(ptr).UnbindDefault(core.NewQObjectFromPointer(object))
	}
}

func (ptr *QMediaPlayer) Unbind(object core.QObject_ITF) {
	defer qt.Recovering("QMediaPlayer::unbind")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_Unbind(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QMediaPlayer) UnbindDefault(object core.QObject_ITF) {
	defer qt.Recovering("QMediaPlayer::unbind")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_UnbindDefault(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QMediaPlayer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMediaPlayer::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMediaPlayer) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMediaPlayer::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMediaPlayerTimerEvent
func callbackQMediaPlayerTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaPlayer::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMediaPlayerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMediaPlayer) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaPlayer::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaPlayer) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaPlayer::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaPlayer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMediaPlayer::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMediaPlayer) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMediaPlayer::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMediaPlayerChildEvent
func callbackQMediaPlayerChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaPlayer::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMediaPlayerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMediaPlayer) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaPlayer::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaPlayer) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaPlayer::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaPlayer) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMediaPlayer::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMediaPlayer) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMediaPlayer::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMediaPlayerCustomEvent
func callbackQMediaPlayerCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaPlayer::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMediaPlayerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMediaPlayer) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaPlayer::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMediaPlayer) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaPlayer::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaPlayer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QMediaPlayerControl struct {
	QMediaControl
}

type QMediaPlayerControl_ITF interface {
	QMediaControl_ITF
	QMediaPlayerControl_PTR() *QMediaPlayerControl
}

func PointerFromQMediaPlayerControl(ptr QMediaPlayerControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaPlayerControl_PTR().Pointer()
	}
	return nil
}

func NewQMediaPlayerControlFromPointer(ptr unsafe.Pointer) *QMediaPlayerControl {
	var n = new(QMediaPlayerControl)
	n.SetPointer(ptr)
	return n
}

func newQMediaPlayerControlFromPointer(ptr unsafe.Pointer) *QMediaPlayerControl {
	var n = NewQMediaPlayerControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QMediaPlayerControl_") {
		n.SetObjectName("QMediaPlayerControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaPlayerControl) QMediaPlayerControl_PTR() *QMediaPlayerControl {
	return ptr
}

func (ptr *QMediaPlayerControl) ConnectAudioAvailableChanged(f func(audio bool)) {
	defer qt.Recovering("connect QMediaPlayerControl::audioAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectAudioAvailableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "audioAvailableChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectAudioAvailableChanged() {
	defer qt.Recovering("disconnect QMediaPlayerControl::audioAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectAudioAvailableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "audioAvailableChanged")
	}
}

//export callbackQMediaPlayerControlAudioAvailableChanged
func callbackQMediaPlayerControlAudioAvailableChanged(ptr unsafe.Pointer, ptrName *C.char, audio C.int) {
	defer qt.Recovering("callback QMediaPlayerControl::audioAvailableChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "audioAvailableChanged"); signal != nil {
		signal.(func(bool))(int(audio) != 0)
	}

}

func (ptr *QMediaPlayerControl) AudioAvailableChanged(audio bool) {
	defer qt.Recovering("QMediaPlayerControl::audioAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_AudioAvailableChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(audio)))
	}
}

func (ptr *QMediaPlayerControl) AvailablePlaybackRanges() *QMediaTimeRange {
	defer qt.Recovering("QMediaPlayerControl::availablePlaybackRanges")

	if ptr.Pointer() != nil {
		return NewQMediaTimeRangeFromPointer(C.QMediaPlayerControl_AvailablePlaybackRanges(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaPlayerControl) ConnectAvailablePlaybackRangesChanged(f func(ranges *QMediaTimeRange)) {
	defer qt.Recovering("connect QMediaPlayerControl::availablePlaybackRangesChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectAvailablePlaybackRangesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "availablePlaybackRangesChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectAvailablePlaybackRangesChanged() {
	defer qt.Recovering("disconnect QMediaPlayerControl::availablePlaybackRangesChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectAvailablePlaybackRangesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "availablePlaybackRangesChanged")
	}
}

//export callbackQMediaPlayerControlAvailablePlaybackRangesChanged
func callbackQMediaPlayerControlAvailablePlaybackRangesChanged(ptr unsafe.Pointer, ptrName *C.char, ranges unsafe.Pointer) {
	defer qt.Recovering("callback QMediaPlayerControl::availablePlaybackRangesChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "availablePlaybackRangesChanged"); signal != nil {
		signal.(func(*QMediaTimeRange))(NewQMediaTimeRangeFromPointer(ranges))
	}

}

func (ptr *QMediaPlayerControl) AvailablePlaybackRangesChanged(ranges QMediaTimeRange_ITF) {
	defer qt.Recovering("QMediaPlayerControl::availablePlaybackRangesChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_AvailablePlaybackRangesChanged(ptr.Pointer(), PointerFromQMediaTimeRange(ranges))
	}
}

func (ptr *QMediaPlayerControl) BufferStatus() int {
	defer qt.Recovering("QMediaPlayerControl::bufferStatus")

	if ptr.Pointer() != nil {
		return int(C.QMediaPlayerControl_BufferStatus(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayerControl) ConnectBufferStatusChanged(f func(progress int)) {
	defer qt.Recovering("connect QMediaPlayerControl::bufferStatusChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectBufferStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bufferStatusChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectBufferStatusChanged() {
	defer qt.Recovering("disconnect QMediaPlayerControl::bufferStatusChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectBufferStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bufferStatusChanged")
	}
}

//export callbackQMediaPlayerControlBufferStatusChanged
func callbackQMediaPlayerControlBufferStatusChanged(ptr unsafe.Pointer, ptrName *C.char, progress C.int) {
	defer qt.Recovering("callback QMediaPlayerControl::bufferStatusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "bufferStatusChanged"); signal != nil {
		signal.(func(int))(int(progress))
	}

}

func (ptr *QMediaPlayerControl) BufferStatusChanged(progress int) {
	defer qt.Recovering("QMediaPlayerControl::bufferStatusChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_BufferStatusChanged(ptr.Pointer(), C.int(progress))
	}
}

func (ptr *QMediaPlayerControl) Duration() int64 {
	defer qt.Recovering("QMediaPlayerControl::duration")

	if ptr.Pointer() != nil {
		return int64(C.QMediaPlayerControl_Duration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayerControl) ConnectDurationChanged(f func(duration int64)) {
	defer qt.Recovering("connect QMediaPlayerControl::durationChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectDurationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "durationChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectDurationChanged() {
	defer qt.Recovering("disconnect QMediaPlayerControl::durationChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectDurationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "durationChanged")
	}
}

//export callbackQMediaPlayerControlDurationChanged
func callbackQMediaPlayerControlDurationChanged(ptr unsafe.Pointer, ptrName *C.char, duration C.longlong) {
	defer qt.Recovering("callback QMediaPlayerControl::durationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "durationChanged"); signal != nil {
		signal.(func(int64))(int64(duration))
	}

}

func (ptr *QMediaPlayerControl) DurationChanged(duration int64) {
	defer qt.Recovering("QMediaPlayerControl::durationChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DurationChanged(ptr.Pointer(), C.longlong(duration))
	}
}

func (ptr *QMediaPlayerControl) ConnectError(f func(error int, errorString string)) {
	defer qt.Recovering("connect QMediaPlayerControl::error")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectError() {
	defer qt.Recovering("disconnect QMediaPlayerControl::error")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQMediaPlayerControlError
func callbackQMediaPlayerControlError(ptr unsafe.Pointer, ptrName *C.char, error C.int, errorString *C.char) {
	defer qt.Recovering("callback QMediaPlayerControl::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error"); signal != nil {
		signal.(func(int, string))(int(error), C.GoString(errorString))
	}

}

func (ptr *QMediaPlayerControl) Error(error int, errorString string) {
	defer qt.Recovering("QMediaPlayerControl::error")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_Error(ptr.Pointer(), C.int(error), C.CString(errorString))
	}
}

func (ptr *QMediaPlayerControl) IsAudioAvailable() bool {
	defer qt.Recovering("QMediaPlayerControl::isAudioAvailable")

	if ptr.Pointer() != nil {
		return C.QMediaPlayerControl_IsAudioAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlayerControl) IsMuted() bool {
	defer qt.Recovering("QMediaPlayerControl::isMuted")

	if ptr.Pointer() != nil {
		return C.QMediaPlayerControl_IsMuted(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlayerControl) IsSeekable() bool {
	defer qt.Recovering("QMediaPlayerControl::isSeekable")

	if ptr.Pointer() != nil {
		return C.QMediaPlayerControl_IsSeekable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlayerControl) IsVideoAvailable() bool {
	defer qt.Recovering("QMediaPlayerControl::isVideoAvailable")

	if ptr.Pointer() != nil {
		return C.QMediaPlayerControl_IsVideoAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlayerControl) Media() *QMediaContent {
	defer qt.Recovering("QMediaPlayerControl::media")

	if ptr.Pointer() != nil {
		return NewQMediaContentFromPointer(C.QMediaPlayerControl_Media(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaPlayerControl) ConnectMediaChanged(f func(content *QMediaContent)) {
	defer qt.Recovering("connect QMediaPlayerControl::mediaChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectMediaChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mediaChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectMediaChanged() {
	defer qt.Recovering("disconnect QMediaPlayerControl::mediaChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectMediaChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mediaChanged")
	}
}

//export callbackQMediaPlayerControlMediaChanged
func callbackQMediaPlayerControlMediaChanged(ptr unsafe.Pointer, ptrName *C.char, content unsafe.Pointer) {
	defer qt.Recovering("callback QMediaPlayerControl::mediaChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "mediaChanged"); signal != nil {
		signal.(func(*QMediaContent))(NewQMediaContentFromPointer(content))
	}

}

func (ptr *QMediaPlayerControl) MediaChanged(content QMediaContent_ITF) {
	defer qt.Recovering("QMediaPlayerControl::mediaChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_MediaChanged(ptr.Pointer(), PointerFromQMediaContent(content))
	}
}

func (ptr *QMediaPlayerControl) MediaStatus() QMediaPlayer__MediaStatus {
	defer qt.Recovering("QMediaPlayerControl::mediaStatus")

	if ptr.Pointer() != nil {
		return QMediaPlayer__MediaStatus(C.QMediaPlayerControl_MediaStatus(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayerControl) ConnectMediaStatusChanged(f func(status QMediaPlayer__MediaStatus)) {
	defer qt.Recovering("connect QMediaPlayerControl::mediaStatusChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectMediaStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mediaStatusChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectMediaStatusChanged() {
	defer qt.Recovering("disconnect QMediaPlayerControl::mediaStatusChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectMediaStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mediaStatusChanged")
	}
}

//export callbackQMediaPlayerControlMediaStatusChanged
func callbackQMediaPlayerControlMediaStatusChanged(ptr unsafe.Pointer, ptrName *C.char, status C.int) {
	defer qt.Recovering("callback QMediaPlayerControl::mediaStatusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "mediaStatusChanged"); signal != nil {
		signal.(func(QMediaPlayer__MediaStatus))(QMediaPlayer__MediaStatus(status))
	}

}

func (ptr *QMediaPlayerControl) MediaStatusChanged(status QMediaPlayer__MediaStatus) {
	defer qt.Recovering("QMediaPlayerControl::mediaStatusChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_MediaStatusChanged(ptr.Pointer(), C.int(status))
	}
}

func (ptr *QMediaPlayerControl) MediaStream() *core.QIODevice {
	defer qt.Recovering("QMediaPlayerControl::mediaStream")

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QMediaPlayerControl_MediaStream(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaPlayerControl) ConnectMutedChanged(f func(mute bool)) {
	defer qt.Recovering("connect QMediaPlayerControl::mutedChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectMutedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mutedChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectMutedChanged() {
	defer qt.Recovering("disconnect QMediaPlayerControl::mutedChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectMutedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mutedChanged")
	}
}

//export callbackQMediaPlayerControlMutedChanged
func callbackQMediaPlayerControlMutedChanged(ptr unsafe.Pointer, ptrName *C.char, mute C.int) {
	defer qt.Recovering("callback QMediaPlayerControl::mutedChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "mutedChanged"); signal != nil {
		signal.(func(bool))(int(mute) != 0)
	}

}

func (ptr *QMediaPlayerControl) MutedChanged(mute bool) {
	defer qt.Recovering("QMediaPlayerControl::mutedChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_MutedChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(mute)))
	}
}

func (ptr *QMediaPlayerControl) Pause() {
	defer qt.Recovering("QMediaPlayerControl::pause")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_Pause(ptr.Pointer())
	}
}

func (ptr *QMediaPlayerControl) Play() {
	defer qt.Recovering("QMediaPlayerControl::play")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_Play(ptr.Pointer())
	}
}

func (ptr *QMediaPlayerControl) PlaybackRate() float64 {
	defer qt.Recovering("QMediaPlayerControl::playbackRate")

	if ptr.Pointer() != nil {
		return float64(C.QMediaPlayerControl_PlaybackRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayerControl) ConnectPlaybackRateChanged(f func(rate float64)) {
	defer qt.Recovering("connect QMediaPlayerControl::playbackRateChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectPlaybackRateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "playbackRateChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectPlaybackRateChanged() {
	defer qt.Recovering("disconnect QMediaPlayerControl::playbackRateChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectPlaybackRateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "playbackRateChanged")
	}
}

//export callbackQMediaPlayerControlPlaybackRateChanged
func callbackQMediaPlayerControlPlaybackRateChanged(ptr unsafe.Pointer, ptrName *C.char, rate C.double) {
	defer qt.Recovering("callback QMediaPlayerControl::playbackRateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "playbackRateChanged"); signal != nil {
		signal.(func(float64))(float64(rate))
	}

}

func (ptr *QMediaPlayerControl) PlaybackRateChanged(rate float64) {
	defer qt.Recovering("QMediaPlayerControl::playbackRateChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_PlaybackRateChanged(ptr.Pointer(), C.double(rate))
	}
}

func (ptr *QMediaPlayerControl) Position() int64 {
	defer qt.Recovering("QMediaPlayerControl::position")

	if ptr.Pointer() != nil {
		return int64(C.QMediaPlayerControl_Position(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayerControl) ConnectPositionChanged(f func(position int64)) {
	defer qt.Recovering("connect QMediaPlayerControl::positionChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectPositionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "positionChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectPositionChanged() {
	defer qt.Recovering("disconnect QMediaPlayerControl::positionChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectPositionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "positionChanged")
	}
}

//export callbackQMediaPlayerControlPositionChanged
func callbackQMediaPlayerControlPositionChanged(ptr unsafe.Pointer, ptrName *C.char, position C.longlong) {
	defer qt.Recovering("callback QMediaPlayerControl::positionChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "positionChanged"); signal != nil {
		signal.(func(int64))(int64(position))
	}

}

func (ptr *QMediaPlayerControl) PositionChanged(position int64) {
	defer qt.Recovering("QMediaPlayerControl::positionChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_PositionChanged(ptr.Pointer(), C.longlong(position))
	}
}

func (ptr *QMediaPlayerControl) ConnectSeekableChanged(f func(seekable bool)) {
	defer qt.Recovering("connect QMediaPlayerControl::seekableChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectSeekableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "seekableChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectSeekableChanged() {
	defer qt.Recovering("disconnect QMediaPlayerControl::seekableChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectSeekableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "seekableChanged")
	}
}

//export callbackQMediaPlayerControlSeekableChanged
func callbackQMediaPlayerControlSeekableChanged(ptr unsafe.Pointer, ptrName *C.char, seekable C.int) {
	defer qt.Recovering("callback QMediaPlayerControl::seekableChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "seekableChanged"); signal != nil {
		signal.(func(bool))(int(seekable) != 0)
	}

}

func (ptr *QMediaPlayerControl) SeekableChanged(seekable bool) {
	defer qt.Recovering("QMediaPlayerControl::seekableChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_SeekableChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(seekable)))
	}
}

func (ptr *QMediaPlayerControl) SetMedia(media QMediaContent_ITF, stream core.QIODevice_ITF) {
	defer qt.Recovering("QMediaPlayerControl::setMedia")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_SetMedia(ptr.Pointer(), PointerFromQMediaContent(media), core.PointerFromQIODevice(stream))
	}
}

func (ptr *QMediaPlayerControl) SetMuted(mute bool) {
	defer qt.Recovering("QMediaPlayerControl::setMuted")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_SetMuted(ptr.Pointer(), C.int(qt.GoBoolToInt(mute)))
	}
}

func (ptr *QMediaPlayerControl) SetPlaybackRate(rate float64) {
	defer qt.Recovering("QMediaPlayerControl::setPlaybackRate")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_SetPlaybackRate(ptr.Pointer(), C.double(rate))
	}
}

func (ptr *QMediaPlayerControl) SetPosition(position int64) {
	defer qt.Recovering("QMediaPlayerControl::setPosition")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_SetPosition(ptr.Pointer(), C.longlong(position))
	}
}

func (ptr *QMediaPlayerControl) SetVolume(volume int) {
	defer qt.Recovering("QMediaPlayerControl::setVolume")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_SetVolume(ptr.Pointer(), C.int(volume))
	}
}

func (ptr *QMediaPlayerControl) State() QMediaPlayer__State {
	defer qt.Recovering("QMediaPlayerControl::state")

	if ptr.Pointer() != nil {
		return QMediaPlayer__State(C.QMediaPlayerControl_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayerControl) ConnectStateChanged(f func(state QMediaPlayer__State)) {
	defer qt.Recovering("connect QMediaPlayerControl::stateChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QMediaPlayerControl::stateChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQMediaPlayerControlStateChanged
func callbackQMediaPlayerControlStateChanged(ptr unsafe.Pointer, ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QMediaPlayerControl::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QMediaPlayer__State))(QMediaPlayer__State(state))
	}

}

func (ptr *QMediaPlayerControl) StateChanged(state QMediaPlayer__State) {
	defer qt.Recovering("QMediaPlayerControl::stateChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_StateChanged(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QMediaPlayerControl) Stop() {
	defer qt.Recovering("QMediaPlayerControl::stop")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_Stop(ptr.Pointer())
	}
}

func (ptr *QMediaPlayerControl) ConnectVideoAvailableChanged(f func(video bool)) {
	defer qt.Recovering("connect QMediaPlayerControl::videoAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectVideoAvailableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "videoAvailableChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectVideoAvailableChanged() {
	defer qt.Recovering("disconnect QMediaPlayerControl::videoAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectVideoAvailableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "videoAvailableChanged")
	}
}

//export callbackQMediaPlayerControlVideoAvailableChanged
func callbackQMediaPlayerControlVideoAvailableChanged(ptr unsafe.Pointer, ptrName *C.char, video C.int) {
	defer qt.Recovering("callback QMediaPlayerControl::videoAvailableChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "videoAvailableChanged"); signal != nil {
		signal.(func(bool))(int(video) != 0)
	}

}

func (ptr *QMediaPlayerControl) VideoAvailableChanged(video bool) {
	defer qt.Recovering("QMediaPlayerControl::videoAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_VideoAvailableChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(video)))
	}
}

func (ptr *QMediaPlayerControl) Volume() int {
	defer qt.Recovering("QMediaPlayerControl::volume")

	if ptr.Pointer() != nil {
		return int(C.QMediaPlayerControl_Volume(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayerControl) ConnectVolumeChanged(f func(volume int)) {
	defer qt.Recovering("connect QMediaPlayerControl::volumeChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectVolumeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "volumeChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectVolumeChanged() {
	defer qt.Recovering("disconnect QMediaPlayerControl::volumeChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectVolumeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "volumeChanged")
	}
}

//export callbackQMediaPlayerControlVolumeChanged
func callbackQMediaPlayerControlVolumeChanged(ptr unsafe.Pointer, ptrName *C.char, volume C.int) {
	defer qt.Recovering("callback QMediaPlayerControl::volumeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "volumeChanged"); signal != nil {
		signal.(func(int))(int(volume))
	}

}

func (ptr *QMediaPlayerControl) VolumeChanged(volume int) {
	defer qt.Recovering("QMediaPlayerControl::volumeChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_VolumeChanged(ptr.Pointer(), C.int(volume))
	}
}

func (ptr *QMediaPlayerControl) DestroyQMediaPlayerControl() {
	defer qt.Recovering("QMediaPlayerControl::~QMediaPlayerControl")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DestroyQMediaPlayerControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMediaPlayerControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMediaPlayerControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMediaPlayerControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMediaPlayerControlTimerEvent
func callbackQMediaPlayerControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaPlayerControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMediaPlayerControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMediaPlayerControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaPlayerControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaPlayerControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaPlayerControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaPlayerControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMediaPlayerControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMediaPlayerControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMediaPlayerControlChildEvent
func callbackQMediaPlayerControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaPlayerControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMediaPlayerControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMediaPlayerControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaPlayerControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaPlayerControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaPlayerControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaPlayerControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMediaPlayerControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMediaPlayerControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMediaPlayerControlCustomEvent
func callbackQMediaPlayerControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaPlayerControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMediaPlayerControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMediaPlayerControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaPlayerControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMediaPlayerControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaPlayerControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QMediaPlaylist struct {
	core.QObject
	QMediaBindableInterface
}

type QMediaPlaylist_ITF interface {
	core.QObject_ITF
	QMediaBindableInterface_ITF
	QMediaPlaylist_PTR() *QMediaPlaylist
}

func (p *QMediaPlaylist) Pointer() unsafe.Pointer {
	return p.QObject_PTR().Pointer()
}

func (p *QMediaPlaylist) SetPointer(ptr unsafe.Pointer) {
	p.QObject_PTR().SetPointer(ptr)
	p.QMediaBindableInterface_PTR().SetPointer(ptr)
}

func PointerFromQMediaPlaylist(ptr QMediaPlaylist_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaPlaylist_PTR().Pointer()
	}
	return nil
}

func NewQMediaPlaylistFromPointer(ptr unsafe.Pointer) *QMediaPlaylist {
	var n = new(QMediaPlaylist)
	n.SetPointer(ptr)
	return n
}

func newQMediaPlaylistFromPointer(ptr unsafe.Pointer) *QMediaPlaylist {
	var n = NewQMediaPlaylistFromPointer(ptr)
	for len(n.ObjectName()) < len("QMediaPlaylist_") {
		n.SetObjectName("QMediaPlaylist_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaPlaylist) QMediaPlaylist_PTR() *QMediaPlaylist {
	return ptr
}

//QMediaPlaylist::Error
type QMediaPlaylist__Error int64

const (
	QMediaPlaylist__NoError                 = QMediaPlaylist__Error(0)
	QMediaPlaylist__FormatError             = QMediaPlaylist__Error(1)
	QMediaPlaylist__FormatNotSupportedError = QMediaPlaylist__Error(2)
	QMediaPlaylist__NetworkError            = QMediaPlaylist__Error(3)
	QMediaPlaylist__AccessDeniedError       = QMediaPlaylist__Error(4)
)

//QMediaPlaylist::PlaybackMode
type QMediaPlaylist__PlaybackMode int64

const (
	QMediaPlaylist__CurrentItemOnce   = QMediaPlaylist__PlaybackMode(0)
	QMediaPlaylist__CurrentItemInLoop = QMediaPlaylist__PlaybackMode(1)
	QMediaPlaylist__Sequential        = QMediaPlaylist__PlaybackMode(2)
	QMediaPlaylist__Loop              = QMediaPlaylist__PlaybackMode(3)
	QMediaPlaylist__Random            = QMediaPlaylist__PlaybackMode(4)
)

func (ptr *QMediaPlaylist) PlaybackMode() QMediaPlaylist__PlaybackMode {
	defer qt.Recovering("QMediaPlaylist::playbackMode")

	if ptr.Pointer() != nil {
		return QMediaPlaylist__PlaybackMode(C.QMediaPlaylist_PlaybackMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlaylist) SetPlaybackMode(mode QMediaPlaylist__PlaybackMode) {
	defer qt.Recovering("QMediaPlaylist::setPlaybackMode")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_SetPlaybackMode(ptr.Pointer(), C.int(mode))
	}
}

func NewQMediaPlaylist(parent core.QObject_ITF) *QMediaPlaylist {
	defer qt.Recovering("QMediaPlaylist::QMediaPlaylist")

	return newQMediaPlaylistFromPointer(C.QMediaPlaylist_NewQMediaPlaylist(core.PointerFromQObject(parent)))
}

func (ptr *QMediaPlaylist) AddMedia(content QMediaContent_ITF) bool {
	defer qt.Recovering("QMediaPlaylist::addMedia")

	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_AddMedia(ptr.Pointer(), PointerFromQMediaContent(content)) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) Clear() bool {
	defer qt.Recovering("QMediaPlaylist::clear")

	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_Clear(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) CurrentIndex() int {
	defer qt.Recovering("QMediaPlaylist::currentIndex")

	if ptr.Pointer() != nil {
		return int(C.QMediaPlaylist_CurrentIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlaylist) ConnectCurrentIndexChanged(f func(position int)) {
	defer qt.Recovering("connect QMediaPlaylist::currentIndexChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectCurrentIndexChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentIndexChanged", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectCurrentIndexChanged() {
	defer qt.Recovering("disconnect QMediaPlaylist::currentIndexChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectCurrentIndexChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentIndexChanged")
	}
}

//export callbackQMediaPlaylistCurrentIndexChanged
func callbackQMediaPlaylistCurrentIndexChanged(ptr unsafe.Pointer, ptrName *C.char, position C.int) {
	defer qt.Recovering("callback QMediaPlaylist::currentIndexChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentIndexChanged"); signal != nil {
		signal.(func(int))(int(position))
	}

}

func (ptr *QMediaPlaylist) CurrentIndexChanged(position int) {
	defer qt.Recovering("QMediaPlaylist::currentIndexChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_CurrentIndexChanged(ptr.Pointer(), C.int(position))
	}
}

func (ptr *QMediaPlaylist) CurrentMedia() *QMediaContent {
	defer qt.Recovering("QMediaPlaylist::currentMedia")

	if ptr.Pointer() != nil {
		return NewQMediaContentFromPointer(C.QMediaPlaylist_CurrentMedia(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaPlaylist) ConnectCurrentMediaChanged(f func(content *QMediaContent)) {
	defer qt.Recovering("connect QMediaPlaylist::currentMediaChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectCurrentMediaChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentMediaChanged", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectCurrentMediaChanged() {
	defer qt.Recovering("disconnect QMediaPlaylist::currentMediaChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectCurrentMediaChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentMediaChanged")
	}
}

//export callbackQMediaPlaylistCurrentMediaChanged
func callbackQMediaPlaylistCurrentMediaChanged(ptr unsafe.Pointer, ptrName *C.char, content unsafe.Pointer) {
	defer qt.Recovering("callback QMediaPlaylist::currentMediaChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentMediaChanged"); signal != nil {
		signal.(func(*QMediaContent))(NewQMediaContentFromPointer(content))
	}

}

func (ptr *QMediaPlaylist) CurrentMediaChanged(content QMediaContent_ITF) {
	defer qt.Recovering("QMediaPlaylist::currentMediaChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_CurrentMediaChanged(ptr.Pointer(), PointerFromQMediaContent(content))
	}
}

func (ptr *QMediaPlaylist) Error() QMediaPlaylist__Error {
	defer qt.Recovering("QMediaPlaylist::error")

	if ptr.Pointer() != nil {
		return QMediaPlaylist__Error(C.QMediaPlaylist_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlaylist) ErrorString() string {
	defer qt.Recovering("QMediaPlaylist::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaPlaylist_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaPlaylist) InsertMedia(pos int, content QMediaContent_ITF) bool {
	defer qt.Recovering("QMediaPlaylist::insertMedia")

	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_InsertMedia(ptr.Pointer(), C.int(pos), PointerFromQMediaContent(content)) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) IsEmpty() bool {
	defer qt.Recovering("QMediaPlaylist::isEmpty")

	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) IsReadOnly() bool {
	defer qt.Recovering("QMediaPlaylist::isReadOnly")

	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_IsReadOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) Load3(device core.QIODevice_ITF, format string) {
	defer qt.Recovering("QMediaPlaylist::load")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_Load3(ptr.Pointer(), core.PointerFromQIODevice(device), C.CString(format))
	}
}

func (ptr *QMediaPlaylist) Load(request network.QNetworkRequest_ITF, format string) {
	defer qt.Recovering("QMediaPlaylist::load")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_Load(ptr.Pointer(), network.PointerFromQNetworkRequest(request), C.CString(format))
	}
}

func (ptr *QMediaPlaylist) Load2(location core.QUrl_ITF, format string) {
	defer qt.Recovering("QMediaPlaylist::load")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_Load2(ptr.Pointer(), core.PointerFromQUrl(location), C.CString(format))
	}
}

func (ptr *QMediaPlaylist) ConnectLoadFailed(f func()) {
	defer qt.Recovering("connect QMediaPlaylist::loadFailed")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectLoadFailed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loadFailed", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectLoadFailed() {
	defer qt.Recovering("disconnect QMediaPlaylist::loadFailed")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectLoadFailed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loadFailed")
	}
}

//export callbackQMediaPlaylistLoadFailed
func callbackQMediaPlaylistLoadFailed(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QMediaPlaylist::loadFailed")

	if signal := qt.GetSignal(C.GoString(ptrName), "loadFailed"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QMediaPlaylist) LoadFailed() {
	defer qt.Recovering("QMediaPlaylist::loadFailed")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_LoadFailed(ptr.Pointer())
	}
}

func (ptr *QMediaPlaylist) ConnectLoaded(f func()) {
	defer qt.Recovering("connect QMediaPlaylist::loaded")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectLoaded(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loaded", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectLoaded() {
	defer qt.Recovering("disconnect QMediaPlaylist::loaded")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectLoaded(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loaded")
	}
}

//export callbackQMediaPlaylistLoaded
func callbackQMediaPlaylistLoaded(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QMediaPlaylist::loaded")

	if signal := qt.GetSignal(C.GoString(ptrName), "loaded"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QMediaPlaylist) Loaded() {
	defer qt.Recovering("QMediaPlaylist::loaded")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_Loaded(ptr.Pointer())
	}
}

func (ptr *QMediaPlaylist) Media(index int) *QMediaContent {
	defer qt.Recovering("QMediaPlaylist::media")

	if ptr.Pointer() != nil {
		return NewQMediaContentFromPointer(C.QMediaPlaylist_Media(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QMediaPlaylist) ConnectMediaAboutToBeInserted(f func(start int, end int)) {
	defer qt.Recovering("connect QMediaPlaylist::mediaAboutToBeInserted")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectMediaAboutToBeInserted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mediaAboutToBeInserted", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectMediaAboutToBeInserted() {
	defer qt.Recovering("disconnect QMediaPlaylist::mediaAboutToBeInserted")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectMediaAboutToBeInserted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mediaAboutToBeInserted")
	}
}

//export callbackQMediaPlaylistMediaAboutToBeInserted
func callbackQMediaPlaylistMediaAboutToBeInserted(ptr unsafe.Pointer, ptrName *C.char, start C.int, end C.int) {
	defer qt.Recovering("callback QMediaPlaylist::mediaAboutToBeInserted")

	if signal := qt.GetSignal(C.GoString(ptrName), "mediaAboutToBeInserted"); signal != nil {
		signal.(func(int, int))(int(start), int(end))
	}

}

func (ptr *QMediaPlaylist) MediaAboutToBeInserted(start int, end int) {
	defer qt.Recovering("QMediaPlaylist::mediaAboutToBeInserted")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_MediaAboutToBeInserted(ptr.Pointer(), C.int(start), C.int(end))
	}
}

func (ptr *QMediaPlaylist) ConnectMediaAboutToBeRemoved(f func(start int, end int)) {
	defer qt.Recovering("connect QMediaPlaylist::mediaAboutToBeRemoved")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectMediaAboutToBeRemoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mediaAboutToBeRemoved", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectMediaAboutToBeRemoved() {
	defer qt.Recovering("disconnect QMediaPlaylist::mediaAboutToBeRemoved")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectMediaAboutToBeRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mediaAboutToBeRemoved")
	}
}

//export callbackQMediaPlaylistMediaAboutToBeRemoved
func callbackQMediaPlaylistMediaAboutToBeRemoved(ptr unsafe.Pointer, ptrName *C.char, start C.int, end C.int) {
	defer qt.Recovering("callback QMediaPlaylist::mediaAboutToBeRemoved")

	if signal := qt.GetSignal(C.GoString(ptrName), "mediaAboutToBeRemoved"); signal != nil {
		signal.(func(int, int))(int(start), int(end))
	}

}

func (ptr *QMediaPlaylist) MediaAboutToBeRemoved(start int, end int) {
	defer qt.Recovering("QMediaPlaylist::mediaAboutToBeRemoved")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_MediaAboutToBeRemoved(ptr.Pointer(), C.int(start), C.int(end))
	}
}

func (ptr *QMediaPlaylist) ConnectMediaChanged(f func(start int, end int)) {
	defer qt.Recovering("connect QMediaPlaylist::mediaChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectMediaChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mediaChanged", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectMediaChanged() {
	defer qt.Recovering("disconnect QMediaPlaylist::mediaChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectMediaChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mediaChanged")
	}
}

//export callbackQMediaPlaylistMediaChanged
func callbackQMediaPlaylistMediaChanged(ptr unsafe.Pointer, ptrName *C.char, start C.int, end C.int) {
	defer qt.Recovering("callback QMediaPlaylist::mediaChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "mediaChanged"); signal != nil {
		signal.(func(int, int))(int(start), int(end))
	}

}

func (ptr *QMediaPlaylist) MediaChanged(start int, end int) {
	defer qt.Recovering("QMediaPlaylist::mediaChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_MediaChanged(ptr.Pointer(), C.int(start), C.int(end))
	}
}

func (ptr *QMediaPlaylist) MediaCount() int {
	defer qt.Recovering("QMediaPlaylist::mediaCount")

	if ptr.Pointer() != nil {
		return int(C.QMediaPlaylist_MediaCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlaylist) ConnectMediaInserted(f func(start int, end int)) {
	defer qt.Recovering("connect QMediaPlaylist::mediaInserted")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectMediaInserted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mediaInserted", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectMediaInserted() {
	defer qt.Recovering("disconnect QMediaPlaylist::mediaInserted")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectMediaInserted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mediaInserted")
	}
}

//export callbackQMediaPlaylistMediaInserted
func callbackQMediaPlaylistMediaInserted(ptr unsafe.Pointer, ptrName *C.char, start C.int, end C.int) {
	defer qt.Recovering("callback QMediaPlaylist::mediaInserted")

	if signal := qt.GetSignal(C.GoString(ptrName), "mediaInserted"); signal != nil {
		signal.(func(int, int))(int(start), int(end))
	}

}

func (ptr *QMediaPlaylist) MediaInserted(start int, end int) {
	defer qt.Recovering("QMediaPlaylist::mediaInserted")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_MediaInserted(ptr.Pointer(), C.int(start), C.int(end))
	}
}

func (ptr *QMediaPlaylist) MediaObject() *QMediaObject {
	defer qt.Recovering("QMediaPlaylist::mediaObject")

	if ptr.Pointer() != nil {
		return NewQMediaObjectFromPointer(C.QMediaPlaylist_MediaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaPlaylist) ConnectMediaRemoved(f func(start int, end int)) {
	defer qt.Recovering("connect QMediaPlaylist::mediaRemoved")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectMediaRemoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mediaRemoved", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectMediaRemoved() {
	defer qt.Recovering("disconnect QMediaPlaylist::mediaRemoved")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectMediaRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mediaRemoved")
	}
}

//export callbackQMediaPlaylistMediaRemoved
func callbackQMediaPlaylistMediaRemoved(ptr unsafe.Pointer, ptrName *C.char, start C.int, end C.int) {
	defer qt.Recovering("callback QMediaPlaylist::mediaRemoved")

	if signal := qt.GetSignal(C.GoString(ptrName), "mediaRemoved"); signal != nil {
		signal.(func(int, int))(int(start), int(end))
	}

}

func (ptr *QMediaPlaylist) MediaRemoved(start int, end int) {
	defer qt.Recovering("QMediaPlaylist::mediaRemoved")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_MediaRemoved(ptr.Pointer(), C.int(start), C.int(end))
	}
}

func (ptr *QMediaPlaylist) Next() {
	defer qt.Recovering("QMediaPlaylist::next")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_Next(ptr.Pointer())
	}
}

func (ptr *QMediaPlaylist) NextIndex(steps int) int {
	defer qt.Recovering("QMediaPlaylist::nextIndex")

	if ptr.Pointer() != nil {
		return int(C.QMediaPlaylist_NextIndex(ptr.Pointer(), C.int(steps)))
	}
	return 0
}

func (ptr *QMediaPlaylist) ConnectPlaybackModeChanged(f func(mode QMediaPlaylist__PlaybackMode)) {
	defer qt.Recovering("connect QMediaPlaylist::playbackModeChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectPlaybackModeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "playbackModeChanged", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectPlaybackModeChanged() {
	defer qt.Recovering("disconnect QMediaPlaylist::playbackModeChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectPlaybackModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "playbackModeChanged")
	}
}

//export callbackQMediaPlaylistPlaybackModeChanged
func callbackQMediaPlaylistPlaybackModeChanged(ptr unsafe.Pointer, ptrName *C.char, mode C.int) {
	defer qt.Recovering("callback QMediaPlaylist::playbackModeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "playbackModeChanged"); signal != nil {
		signal.(func(QMediaPlaylist__PlaybackMode))(QMediaPlaylist__PlaybackMode(mode))
	}

}

func (ptr *QMediaPlaylist) PlaybackModeChanged(mode QMediaPlaylist__PlaybackMode) {
	defer qt.Recovering("QMediaPlaylist::playbackModeChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_PlaybackModeChanged(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QMediaPlaylist) Previous() {
	defer qt.Recovering("QMediaPlaylist::previous")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_Previous(ptr.Pointer())
	}
}

func (ptr *QMediaPlaylist) PreviousIndex(steps int) int {
	defer qt.Recovering("QMediaPlaylist::previousIndex")

	if ptr.Pointer() != nil {
		return int(C.QMediaPlaylist_PreviousIndex(ptr.Pointer(), C.int(steps)))
	}
	return 0
}

func (ptr *QMediaPlaylist) RemoveMedia(pos int) bool {
	defer qt.Recovering("QMediaPlaylist::removeMedia")

	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_RemoveMedia(ptr.Pointer(), C.int(pos)) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) RemoveMedia2(start int, end int) bool {
	defer qt.Recovering("QMediaPlaylist::removeMedia")

	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_RemoveMedia2(ptr.Pointer(), C.int(start), C.int(end)) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) Save2(device core.QIODevice_ITF, format string) bool {
	defer qt.Recovering("QMediaPlaylist::save")

	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_Save2(ptr.Pointer(), core.PointerFromQIODevice(device), C.CString(format)) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) Save(location core.QUrl_ITF, format string) bool {
	defer qt.Recovering("QMediaPlaylist::save")

	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_Save(ptr.Pointer(), core.PointerFromQUrl(location), C.CString(format)) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) SetCurrentIndex(playlistPosition int) {
	defer qt.Recovering("QMediaPlaylist::setCurrentIndex")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_SetCurrentIndex(ptr.Pointer(), C.int(playlistPosition))
	}
}

func (ptr *QMediaPlaylist) Shuffle() {
	defer qt.Recovering("QMediaPlaylist::shuffle")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_Shuffle(ptr.Pointer())
	}
}

func (ptr *QMediaPlaylist) DestroyQMediaPlaylist() {
	defer qt.Recovering("QMediaPlaylist::~QMediaPlaylist")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DestroyQMediaPlaylist(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMediaPlaylist) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMediaPlaylist::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMediaPlaylist::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMediaPlaylistTimerEvent
func callbackQMediaPlaylistTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaPlaylist::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMediaPlaylistFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMediaPlaylist) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaPlaylist::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaPlaylist) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaPlaylist::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaPlaylist) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMediaPlaylist::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMediaPlaylist::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMediaPlaylistChildEvent
func callbackQMediaPlaylistChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaPlaylist::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMediaPlaylistFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMediaPlaylist) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaPlaylist::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaPlaylist) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaPlaylist::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaPlaylist) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMediaPlaylist::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMediaPlaylist::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMediaPlaylistCustomEvent
func callbackQMediaPlaylistCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaPlaylist::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMediaPlaylistFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMediaPlaylist) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaPlaylist::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMediaPlaylist) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaPlaylist::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QMediaRecorder struct {
	core.QObject
	QMediaBindableInterface
}

type QMediaRecorder_ITF interface {
	core.QObject_ITF
	QMediaBindableInterface_ITF
	QMediaRecorder_PTR() *QMediaRecorder
}

func (p *QMediaRecorder) Pointer() unsafe.Pointer {
	return p.QObject_PTR().Pointer()
}

func (p *QMediaRecorder) SetPointer(ptr unsafe.Pointer) {
	p.QObject_PTR().SetPointer(ptr)
	p.QMediaBindableInterface_PTR().SetPointer(ptr)
}

func PointerFromQMediaRecorder(ptr QMediaRecorder_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaRecorder_PTR().Pointer()
	}
	return nil
}

func NewQMediaRecorderFromPointer(ptr unsafe.Pointer) *QMediaRecorder {
	var n = new(QMediaRecorder)
	n.SetPointer(ptr)
	return n
}

func newQMediaRecorderFromPointer(ptr unsafe.Pointer) *QMediaRecorder {
	var n = NewQMediaRecorderFromPointer(ptr)
	for len(n.ObjectName()) < len("QMediaRecorder_") {
		n.SetObjectName("QMediaRecorder_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaRecorder) QMediaRecorder_PTR() *QMediaRecorder {
	return ptr
}

//QMediaRecorder::Error
type QMediaRecorder__Error int64

const (
	QMediaRecorder__NoError         = QMediaRecorder__Error(0)
	QMediaRecorder__ResourceError   = QMediaRecorder__Error(1)
	QMediaRecorder__FormatError     = QMediaRecorder__Error(2)
	QMediaRecorder__OutOfSpaceError = QMediaRecorder__Error(3)
)

//QMediaRecorder::State
type QMediaRecorder__State int64

const (
	QMediaRecorder__StoppedState   = QMediaRecorder__State(0)
	QMediaRecorder__RecordingState = QMediaRecorder__State(1)
	QMediaRecorder__PausedState    = QMediaRecorder__State(2)
)

//QMediaRecorder::Status
type QMediaRecorder__Status int64

const (
	QMediaRecorder__UnavailableStatus = QMediaRecorder__Status(0)
	QMediaRecorder__UnloadedStatus    = QMediaRecorder__Status(1)
	QMediaRecorder__LoadingStatus     = QMediaRecorder__Status(2)
	QMediaRecorder__LoadedStatus      = QMediaRecorder__Status(3)
	QMediaRecorder__StartingStatus    = QMediaRecorder__Status(4)
	QMediaRecorder__RecordingStatus   = QMediaRecorder__Status(5)
	QMediaRecorder__PausedStatus      = QMediaRecorder__Status(6)
	QMediaRecorder__FinalizingStatus  = QMediaRecorder__Status(7)
)

func (ptr *QMediaRecorder) ActualLocation() *core.QUrl {
	defer qt.Recovering("QMediaRecorder::actualLocation")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QMediaRecorder_ActualLocation(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaRecorder) Duration() int64 {
	defer qt.Recovering("QMediaRecorder::duration")

	if ptr.Pointer() != nil {
		return int64(C.QMediaRecorder_Duration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaRecorder) IsMetaDataAvailable() bool {
	defer qt.Recovering("QMediaRecorder::isMetaDataAvailable")

	if ptr.Pointer() != nil {
		return C.QMediaRecorder_IsMetaDataAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaRecorder) IsMetaDataWritable() bool {
	defer qt.Recovering("QMediaRecorder::isMetaDataWritable")

	if ptr.Pointer() != nil {
		return C.QMediaRecorder_IsMetaDataWritable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaRecorder) IsMuted() bool {
	defer qt.Recovering("QMediaRecorder::isMuted")

	if ptr.Pointer() != nil {
		return C.QMediaRecorder_IsMuted(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaRecorder) OutputLocation() *core.QUrl {
	defer qt.Recovering("QMediaRecorder::outputLocation")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QMediaRecorder_OutputLocation(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaRecorder) SetMuted(muted bool) {
	defer qt.Recovering("QMediaRecorder::setMuted")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_SetMuted(ptr.Pointer(), C.int(qt.GoBoolToInt(muted)))
	}
}

func (ptr *QMediaRecorder) SetOutputLocation(location core.QUrl_ITF) bool {
	defer qt.Recovering("QMediaRecorder::setOutputLocation")

	if ptr.Pointer() != nil {
		return C.QMediaRecorder_SetOutputLocation(ptr.Pointer(), core.PointerFromQUrl(location)) != 0
	}
	return false
}

func (ptr *QMediaRecorder) SetVolume(volume float64) {
	defer qt.Recovering("QMediaRecorder::setVolume")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_SetVolume(ptr.Pointer(), C.double(volume))
	}
}

func (ptr *QMediaRecorder) Volume() float64 {
	defer qt.Recovering("QMediaRecorder::volume")

	if ptr.Pointer() != nil {
		return float64(C.QMediaRecorder_Volume(ptr.Pointer()))
	}
	return 0
}

func NewQMediaRecorder(mediaObject QMediaObject_ITF, parent core.QObject_ITF) *QMediaRecorder {
	defer qt.Recovering("QMediaRecorder::QMediaRecorder")

	return newQMediaRecorderFromPointer(C.QMediaRecorder_NewQMediaRecorder(PointerFromQMediaObject(mediaObject), core.PointerFromQObject(parent)))
}

func (ptr *QMediaRecorder) ConnectActualLocationChanged(f func(location *core.QUrl)) {
	defer qt.Recovering("connect QMediaRecorder::actualLocationChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_ConnectActualLocationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "actualLocationChanged", f)
	}
}

func (ptr *QMediaRecorder) DisconnectActualLocationChanged() {
	defer qt.Recovering("disconnect QMediaRecorder::actualLocationChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_DisconnectActualLocationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "actualLocationChanged")
	}
}

//export callbackQMediaRecorderActualLocationChanged
func callbackQMediaRecorderActualLocationChanged(ptr unsafe.Pointer, ptrName *C.char, location unsafe.Pointer) {
	defer qt.Recovering("callback QMediaRecorder::actualLocationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "actualLocationChanged"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(location))
	}

}

func (ptr *QMediaRecorder) ActualLocationChanged(location core.QUrl_ITF) {
	defer qt.Recovering("QMediaRecorder::actualLocationChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_ActualLocationChanged(ptr.Pointer(), core.PointerFromQUrl(location))
	}
}

func (ptr *QMediaRecorder) AudioCodecDescription(codec string) string {
	defer qt.Recovering("QMediaRecorder::audioCodecDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaRecorder_AudioCodecDescription(ptr.Pointer(), C.CString(codec)))
	}
	return ""
}

func (ptr *QMediaRecorder) AudioSettings() *QAudioEncoderSettings {
	defer qt.Recovering("QMediaRecorder::audioSettings")

	if ptr.Pointer() != nil {
		return NewQAudioEncoderSettingsFromPointer(C.QMediaRecorder_AudioSettings(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaRecorder) Availability() QMultimedia__AvailabilityStatus {
	defer qt.Recovering("QMediaRecorder::availability")

	if ptr.Pointer() != nil {
		return QMultimedia__AvailabilityStatus(C.QMediaRecorder_Availability(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaRecorder) ConnectAvailabilityChanged2(f func(availability QMultimedia__AvailabilityStatus)) {
	defer qt.Recovering("connect QMediaRecorder::availabilityChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_ConnectAvailabilityChanged2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "availabilityChanged2", f)
	}
}

func (ptr *QMediaRecorder) DisconnectAvailabilityChanged2() {
	defer qt.Recovering("disconnect QMediaRecorder::availabilityChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_DisconnectAvailabilityChanged2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "availabilityChanged2")
	}
}

//export callbackQMediaRecorderAvailabilityChanged2
func callbackQMediaRecorderAvailabilityChanged2(ptr unsafe.Pointer, ptrName *C.char, availability C.int) {
	defer qt.Recovering("callback QMediaRecorder::availabilityChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "availabilityChanged2"); signal != nil {
		signal.(func(QMultimedia__AvailabilityStatus))(QMultimedia__AvailabilityStatus(availability))
	}

}

func (ptr *QMediaRecorder) AvailabilityChanged2(availability QMultimedia__AvailabilityStatus) {
	defer qt.Recovering("QMediaRecorder::availabilityChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_AvailabilityChanged2(ptr.Pointer(), C.int(availability))
	}
}

func (ptr *QMediaRecorder) ConnectAvailabilityChanged(f func(available bool)) {
	defer qt.Recovering("connect QMediaRecorder::availabilityChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_ConnectAvailabilityChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "availabilityChanged", f)
	}
}

func (ptr *QMediaRecorder) DisconnectAvailabilityChanged() {
	defer qt.Recovering("disconnect QMediaRecorder::availabilityChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_DisconnectAvailabilityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "availabilityChanged")
	}
}

//export callbackQMediaRecorderAvailabilityChanged
func callbackQMediaRecorderAvailabilityChanged(ptr unsafe.Pointer, ptrName *C.char, available C.int) {
	defer qt.Recovering("callback QMediaRecorder::availabilityChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "availabilityChanged"); signal != nil {
		signal.(func(bool))(int(available) != 0)
	}

}

func (ptr *QMediaRecorder) AvailabilityChanged(available bool) {
	defer qt.Recovering("QMediaRecorder::availabilityChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_AvailabilityChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(available)))
	}
}

func (ptr *QMediaRecorder) AvailableMetaData() []string {
	defer qt.Recovering("QMediaRecorder::availableMetaData")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMediaRecorder_AvailableMetaData(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QMediaRecorder) ContainerDescription(format string) string {
	defer qt.Recovering("QMediaRecorder::containerDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaRecorder_ContainerDescription(ptr.Pointer(), C.CString(format)))
	}
	return ""
}

func (ptr *QMediaRecorder) ContainerFormat() string {
	defer qt.Recovering("QMediaRecorder::containerFormat")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaRecorder_ContainerFormat(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaRecorder) ConnectDurationChanged(f func(duration int64)) {
	defer qt.Recovering("connect QMediaRecorder::durationChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_ConnectDurationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "durationChanged", f)
	}
}

func (ptr *QMediaRecorder) DisconnectDurationChanged() {
	defer qt.Recovering("disconnect QMediaRecorder::durationChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_DisconnectDurationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "durationChanged")
	}
}

//export callbackQMediaRecorderDurationChanged
func callbackQMediaRecorderDurationChanged(ptr unsafe.Pointer, ptrName *C.char, duration C.longlong) {
	defer qt.Recovering("callback QMediaRecorder::durationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "durationChanged"); signal != nil {
		signal.(func(int64))(int64(duration))
	}

}

func (ptr *QMediaRecorder) DurationChanged(duration int64) {
	defer qt.Recovering("QMediaRecorder::durationChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_DurationChanged(ptr.Pointer(), C.longlong(duration))
	}
}

func (ptr *QMediaRecorder) ConnectError2(f func(error QMediaRecorder__Error)) {
	defer qt.Recovering("connect QMediaRecorder::error")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QMediaRecorder) DisconnectError2() {
	defer qt.Recovering("disconnect QMediaRecorder::error")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQMediaRecorderError2
func callbackQMediaRecorderError2(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QMediaRecorder::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QMediaRecorder__Error))(QMediaRecorder__Error(error))
	}

}

func (ptr *QMediaRecorder) Error2(error QMediaRecorder__Error) {
	defer qt.Recovering("QMediaRecorder::error")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_Error2(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QMediaRecorder) Error() QMediaRecorder__Error {
	defer qt.Recovering("QMediaRecorder::error")

	if ptr.Pointer() != nil {
		return QMediaRecorder__Error(C.QMediaRecorder_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaRecorder) ErrorString() string {
	defer qt.Recovering("QMediaRecorder::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaRecorder_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaRecorder) IsAvailable() bool {
	defer qt.Recovering("QMediaRecorder::isAvailable")

	if ptr.Pointer() != nil {
		return C.QMediaRecorder_IsAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaRecorder) MediaObject() *QMediaObject {
	defer qt.Recovering("QMediaRecorder::mediaObject")

	if ptr.Pointer() != nil {
		return NewQMediaObjectFromPointer(C.QMediaRecorder_MediaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaRecorder) MetaData(key string) *core.QVariant {
	defer qt.Recovering("QMediaRecorder::metaData")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QMediaRecorder_MetaData(ptr.Pointer(), C.CString(key)))
	}
	return nil
}

func (ptr *QMediaRecorder) ConnectMetaDataAvailableChanged(f func(available bool)) {
	defer qt.Recovering("connect QMediaRecorder::metaDataAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_ConnectMetaDataAvailableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataAvailableChanged", f)
	}
}

func (ptr *QMediaRecorder) DisconnectMetaDataAvailableChanged() {
	defer qt.Recovering("disconnect QMediaRecorder::metaDataAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_DisconnectMetaDataAvailableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataAvailableChanged")
	}
}

//export callbackQMediaRecorderMetaDataAvailableChanged
func callbackQMediaRecorderMetaDataAvailableChanged(ptr unsafe.Pointer, ptrName *C.char, available C.int) {
	defer qt.Recovering("callback QMediaRecorder::metaDataAvailableChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaDataAvailableChanged"); signal != nil {
		signal.(func(bool))(int(available) != 0)
	}

}

func (ptr *QMediaRecorder) MetaDataAvailableChanged(available bool) {
	defer qt.Recovering("QMediaRecorder::metaDataAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_MetaDataAvailableChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(available)))
	}
}

func (ptr *QMediaRecorder) ConnectMetaDataChanged(f func()) {
	defer qt.Recovering("connect QMediaRecorder::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_ConnectMetaDataChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataChanged", f)
	}
}

func (ptr *QMediaRecorder) DisconnectMetaDataChanged() {
	defer qt.Recovering("disconnect QMediaRecorder::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_DisconnectMetaDataChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataChanged")
	}
}

//export callbackQMediaRecorderMetaDataChanged
func callbackQMediaRecorderMetaDataChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QMediaRecorder::metaDataChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaDataChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QMediaRecorder) MetaDataChanged() {
	defer qt.Recovering("QMediaRecorder::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_MetaDataChanged(ptr.Pointer())
	}
}

func (ptr *QMediaRecorder) ConnectMetaDataChanged2(f func(key string, value *core.QVariant)) {
	defer qt.Recovering("connect QMediaRecorder::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_ConnectMetaDataChanged2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataChanged2", f)
	}
}

func (ptr *QMediaRecorder) DisconnectMetaDataChanged2() {
	defer qt.Recovering("disconnect QMediaRecorder::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_DisconnectMetaDataChanged2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataChanged2")
	}
}

//export callbackQMediaRecorderMetaDataChanged2
func callbackQMediaRecorderMetaDataChanged2(ptr unsafe.Pointer, ptrName *C.char, key *C.char, value unsafe.Pointer) {
	defer qt.Recovering("callback QMediaRecorder::metaDataChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaDataChanged2"); signal != nil {
		signal.(func(string, *core.QVariant))(C.GoString(key), core.NewQVariantFromPointer(value))
	}

}

func (ptr *QMediaRecorder) MetaDataChanged2(key string, value core.QVariant_ITF) {
	defer qt.Recovering("QMediaRecorder::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_MetaDataChanged2(ptr.Pointer(), C.CString(key), core.PointerFromQVariant(value))
	}
}

func (ptr *QMediaRecorder) ConnectMetaDataWritableChanged(f func(writable bool)) {
	defer qt.Recovering("connect QMediaRecorder::metaDataWritableChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_ConnectMetaDataWritableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataWritableChanged", f)
	}
}

func (ptr *QMediaRecorder) DisconnectMetaDataWritableChanged() {
	defer qt.Recovering("disconnect QMediaRecorder::metaDataWritableChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_DisconnectMetaDataWritableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataWritableChanged")
	}
}

//export callbackQMediaRecorderMetaDataWritableChanged
func callbackQMediaRecorderMetaDataWritableChanged(ptr unsafe.Pointer, ptrName *C.char, writable C.int) {
	defer qt.Recovering("callback QMediaRecorder::metaDataWritableChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaDataWritableChanged"); signal != nil {
		signal.(func(bool))(int(writable) != 0)
	}

}

func (ptr *QMediaRecorder) MetaDataWritableChanged(writable bool) {
	defer qt.Recovering("QMediaRecorder::metaDataWritableChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_MetaDataWritableChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(writable)))
	}
}

func (ptr *QMediaRecorder) ConnectMutedChanged(f func(muted bool)) {
	defer qt.Recovering("connect QMediaRecorder::mutedChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_ConnectMutedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mutedChanged", f)
	}
}

func (ptr *QMediaRecorder) DisconnectMutedChanged() {
	defer qt.Recovering("disconnect QMediaRecorder::mutedChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_DisconnectMutedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mutedChanged")
	}
}

//export callbackQMediaRecorderMutedChanged
func callbackQMediaRecorderMutedChanged(ptr unsafe.Pointer, ptrName *C.char, muted C.int) {
	defer qt.Recovering("callback QMediaRecorder::mutedChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "mutedChanged"); signal != nil {
		signal.(func(bool))(int(muted) != 0)
	}

}

func (ptr *QMediaRecorder) MutedChanged(muted bool) {
	defer qt.Recovering("QMediaRecorder::mutedChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_MutedChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(muted)))
	}
}

func (ptr *QMediaRecorder) Pause() {
	defer qt.Recovering("QMediaRecorder::pause")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_Pause(ptr.Pointer())
	}
}

func (ptr *QMediaRecorder) Record() {
	defer qt.Recovering("QMediaRecorder::record")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_Record(ptr.Pointer())
	}
}

func (ptr *QMediaRecorder) SetAudioSettings(settings QAudioEncoderSettings_ITF) {
	defer qt.Recovering("QMediaRecorder::setAudioSettings")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_SetAudioSettings(ptr.Pointer(), PointerFromQAudioEncoderSettings(settings))
	}
}

func (ptr *QMediaRecorder) SetContainerFormat(container string) {
	defer qt.Recovering("QMediaRecorder::setContainerFormat")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_SetContainerFormat(ptr.Pointer(), C.CString(container))
	}
}

func (ptr *QMediaRecorder) SetEncodingSettings(audio QAudioEncoderSettings_ITF, video QVideoEncoderSettings_ITF, container string) {
	defer qt.Recovering("QMediaRecorder::setEncodingSettings")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_SetEncodingSettings(ptr.Pointer(), PointerFromQAudioEncoderSettings(audio), PointerFromQVideoEncoderSettings(video), C.CString(container))
	}
}

func (ptr *QMediaRecorder) SetMetaData(key string, value core.QVariant_ITF) {
	defer qt.Recovering("QMediaRecorder::setMetaData")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_SetMetaData(ptr.Pointer(), C.CString(key), core.PointerFromQVariant(value))
	}
}

func (ptr *QMediaRecorder) SetVideoSettings(settings QVideoEncoderSettings_ITF) {
	defer qt.Recovering("QMediaRecorder::setVideoSettings")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_SetVideoSettings(ptr.Pointer(), PointerFromQVideoEncoderSettings(settings))
	}
}

func (ptr *QMediaRecorder) State() QMediaRecorder__State {
	defer qt.Recovering("QMediaRecorder::state")

	if ptr.Pointer() != nil {
		return QMediaRecorder__State(C.QMediaRecorder_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaRecorder) ConnectStateChanged(f func(state QMediaRecorder__State)) {
	defer qt.Recovering("connect QMediaRecorder::stateChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QMediaRecorder) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QMediaRecorder::stateChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQMediaRecorderStateChanged
func callbackQMediaRecorderStateChanged(ptr unsafe.Pointer, ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QMediaRecorder::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QMediaRecorder__State))(QMediaRecorder__State(state))
	}

}

func (ptr *QMediaRecorder) StateChanged(state QMediaRecorder__State) {
	defer qt.Recovering("QMediaRecorder::stateChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_StateChanged(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QMediaRecorder) Status() QMediaRecorder__Status {
	defer qt.Recovering("QMediaRecorder::status")

	if ptr.Pointer() != nil {
		return QMediaRecorder__Status(C.QMediaRecorder_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaRecorder) ConnectStatusChanged(f func(status QMediaRecorder__Status)) {
	defer qt.Recovering("connect QMediaRecorder::statusChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_ConnectStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "statusChanged", f)
	}
}

func (ptr *QMediaRecorder) DisconnectStatusChanged() {
	defer qt.Recovering("disconnect QMediaRecorder::statusChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "statusChanged")
	}
}

//export callbackQMediaRecorderStatusChanged
func callbackQMediaRecorderStatusChanged(ptr unsafe.Pointer, ptrName *C.char, status C.int) {
	defer qt.Recovering("callback QMediaRecorder::statusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "statusChanged"); signal != nil {
		signal.(func(QMediaRecorder__Status))(QMediaRecorder__Status(status))
	}

}

func (ptr *QMediaRecorder) StatusChanged(status QMediaRecorder__Status) {
	defer qt.Recovering("QMediaRecorder::statusChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_StatusChanged(ptr.Pointer(), C.int(status))
	}
}

func (ptr *QMediaRecorder) Stop() {
	defer qt.Recovering("QMediaRecorder::stop")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_Stop(ptr.Pointer())
	}
}

func (ptr *QMediaRecorder) SupportedAudioCodecs() []string {
	defer qt.Recovering("QMediaRecorder::supportedAudioCodecs")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMediaRecorder_SupportedAudioCodecs(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QMediaRecorder) SupportedContainers() []string {
	defer qt.Recovering("QMediaRecorder::supportedContainers")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMediaRecorder_SupportedContainers(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QMediaRecorder) SupportedVideoCodecs() []string {
	defer qt.Recovering("QMediaRecorder::supportedVideoCodecs")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMediaRecorder_SupportedVideoCodecs(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QMediaRecorder) VideoCodecDescription(codec string) string {
	defer qt.Recovering("QMediaRecorder::videoCodecDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaRecorder_VideoCodecDescription(ptr.Pointer(), C.CString(codec)))
	}
	return ""
}

func (ptr *QMediaRecorder) VideoSettings() *QVideoEncoderSettings {
	defer qt.Recovering("QMediaRecorder::videoSettings")

	if ptr.Pointer() != nil {
		return NewQVideoEncoderSettingsFromPointer(C.QMediaRecorder_VideoSettings(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaRecorder) ConnectVolumeChanged(f func(volume float64)) {
	defer qt.Recovering("connect QMediaRecorder::volumeChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_ConnectVolumeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "volumeChanged", f)
	}
}

func (ptr *QMediaRecorder) DisconnectVolumeChanged() {
	defer qt.Recovering("disconnect QMediaRecorder::volumeChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_DisconnectVolumeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "volumeChanged")
	}
}

//export callbackQMediaRecorderVolumeChanged
func callbackQMediaRecorderVolumeChanged(ptr unsafe.Pointer, ptrName *C.char, volume C.double) {
	defer qt.Recovering("callback QMediaRecorder::volumeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "volumeChanged"); signal != nil {
		signal.(func(float64))(float64(volume))
	}

}

func (ptr *QMediaRecorder) VolumeChanged(volume float64) {
	defer qt.Recovering("QMediaRecorder::volumeChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_VolumeChanged(ptr.Pointer(), C.double(volume))
	}
}

func (ptr *QMediaRecorder) DestroyQMediaRecorder() {
	defer qt.Recovering("QMediaRecorder::~QMediaRecorder")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_DestroyQMediaRecorder(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMediaRecorder) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMediaRecorder::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMediaRecorder) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMediaRecorder::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMediaRecorderTimerEvent
func callbackQMediaRecorderTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaRecorder::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMediaRecorderFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMediaRecorder) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaRecorder::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaRecorder) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaRecorder::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaRecorder) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMediaRecorder::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMediaRecorder) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMediaRecorder::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMediaRecorderChildEvent
func callbackQMediaRecorderChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaRecorder::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMediaRecorderFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMediaRecorder) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaRecorder::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaRecorder) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaRecorder::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaRecorder) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMediaRecorder::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMediaRecorder) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMediaRecorder::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMediaRecorderCustomEvent
func callbackQMediaRecorderCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaRecorder::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMediaRecorderFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMediaRecorder) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaRecorder::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMediaRecorder) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaRecorder::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaRecorder_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QMediaRecorderControl struct {
	QMediaControl
}

type QMediaRecorderControl_ITF interface {
	QMediaControl_ITF
	QMediaRecorderControl_PTR() *QMediaRecorderControl
}

func PointerFromQMediaRecorderControl(ptr QMediaRecorderControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaRecorderControl_PTR().Pointer()
	}
	return nil
}

func NewQMediaRecorderControlFromPointer(ptr unsafe.Pointer) *QMediaRecorderControl {
	var n = new(QMediaRecorderControl)
	n.SetPointer(ptr)
	return n
}

func newQMediaRecorderControlFromPointer(ptr unsafe.Pointer) *QMediaRecorderControl {
	var n = NewQMediaRecorderControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QMediaRecorderControl_") {
		n.SetObjectName("QMediaRecorderControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaRecorderControl) QMediaRecorderControl_PTR() *QMediaRecorderControl {
	return ptr
}

func (ptr *QMediaRecorderControl) ConnectActualLocationChanged(f func(location *core.QUrl)) {
	defer qt.Recovering("connect QMediaRecorderControl::actualLocationChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ConnectActualLocationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "actualLocationChanged", f)
	}
}

func (ptr *QMediaRecorderControl) DisconnectActualLocationChanged() {
	defer qt.Recovering("disconnect QMediaRecorderControl::actualLocationChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DisconnectActualLocationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "actualLocationChanged")
	}
}

//export callbackQMediaRecorderControlActualLocationChanged
func callbackQMediaRecorderControlActualLocationChanged(ptr unsafe.Pointer, ptrName *C.char, location unsafe.Pointer) {
	defer qt.Recovering("callback QMediaRecorderControl::actualLocationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "actualLocationChanged"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(location))
	}

}

func (ptr *QMediaRecorderControl) ActualLocationChanged(location core.QUrl_ITF) {
	defer qt.Recovering("QMediaRecorderControl::actualLocationChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ActualLocationChanged(ptr.Pointer(), core.PointerFromQUrl(location))
	}
}

func (ptr *QMediaRecorderControl) ApplySettings() {
	defer qt.Recovering("QMediaRecorderControl::applySettings")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ApplySettings(ptr.Pointer())
	}
}

func (ptr *QMediaRecorderControl) Duration() int64 {
	defer qt.Recovering("QMediaRecorderControl::duration")

	if ptr.Pointer() != nil {
		return int64(C.QMediaRecorderControl_Duration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaRecorderControl) ConnectDurationChanged(f func(duration int64)) {
	defer qt.Recovering("connect QMediaRecorderControl::durationChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ConnectDurationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "durationChanged", f)
	}
}

func (ptr *QMediaRecorderControl) DisconnectDurationChanged() {
	defer qt.Recovering("disconnect QMediaRecorderControl::durationChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DisconnectDurationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "durationChanged")
	}
}

//export callbackQMediaRecorderControlDurationChanged
func callbackQMediaRecorderControlDurationChanged(ptr unsafe.Pointer, ptrName *C.char, duration C.longlong) {
	defer qt.Recovering("callback QMediaRecorderControl::durationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "durationChanged"); signal != nil {
		signal.(func(int64))(int64(duration))
	}

}

func (ptr *QMediaRecorderControl) DurationChanged(duration int64) {
	defer qt.Recovering("QMediaRecorderControl::durationChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DurationChanged(ptr.Pointer(), C.longlong(duration))
	}
}

func (ptr *QMediaRecorderControl) ConnectError(f func(error int, errorString string)) {
	defer qt.Recovering("connect QMediaRecorderControl::error")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QMediaRecorderControl) DisconnectError() {
	defer qt.Recovering("disconnect QMediaRecorderControl::error")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQMediaRecorderControlError
func callbackQMediaRecorderControlError(ptr unsafe.Pointer, ptrName *C.char, error C.int, errorString *C.char) {
	defer qt.Recovering("callback QMediaRecorderControl::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error"); signal != nil {
		signal.(func(int, string))(int(error), C.GoString(errorString))
	}

}

func (ptr *QMediaRecorderControl) Error(error int, errorString string) {
	defer qt.Recovering("QMediaRecorderControl::error")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_Error(ptr.Pointer(), C.int(error), C.CString(errorString))
	}
}

func (ptr *QMediaRecorderControl) IsMuted() bool {
	defer qt.Recovering("QMediaRecorderControl::isMuted")

	if ptr.Pointer() != nil {
		return C.QMediaRecorderControl_IsMuted(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaRecorderControl) ConnectMutedChanged(f func(muted bool)) {
	defer qt.Recovering("connect QMediaRecorderControl::mutedChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ConnectMutedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mutedChanged", f)
	}
}

func (ptr *QMediaRecorderControl) DisconnectMutedChanged() {
	defer qt.Recovering("disconnect QMediaRecorderControl::mutedChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DisconnectMutedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mutedChanged")
	}
}

//export callbackQMediaRecorderControlMutedChanged
func callbackQMediaRecorderControlMutedChanged(ptr unsafe.Pointer, ptrName *C.char, muted C.int) {
	defer qt.Recovering("callback QMediaRecorderControl::mutedChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "mutedChanged"); signal != nil {
		signal.(func(bool))(int(muted) != 0)
	}

}

func (ptr *QMediaRecorderControl) MutedChanged(muted bool) {
	defer qt.Recovering("QMediaRecorderControl::mutedChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_MutedChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(muted)))
	}
}

func (ptr *QMediaRecorderControl) OutputLocation() *core.QUrl {
	defer qt.Recovering("QMediaRecorderControl::outputLocation")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QMediaRecorderControl_OutputLocation(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaRecorderControl) SetMuted(muted bool) {
	defer qt.Recovering("QMediaRecorderControl::setMuted")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_SetMuted(ptr.Pointer(), C.int(qt.GoBoolToInt(muted)))
	}
}

func (ptr *QMediaRecorderControl) SetOutputLocation(location core.QUrl_ITF) bool {
	defer qt.Recovering("QMediaRecorderControl::setOutputLocation")

	if ptr.Pointer() != nil {
		return C.QMediaRecorderControl_SetOutputLocation(ptr.Pointer(), core.PointerFromQUrl(location)) != 0
	}
	return false
}

func (ptr *QMediaRecorderControl) SetState(state QMediaRecorder__State) {
	defer qt.Recovering("QMediaRecorderControl::setState")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_SetState(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QMediaRecorderControl) SetVolume(gain float64) {
	defer qt.Recovering("QMediaRecorderControl::setVolume")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_SetVolume(ptr.Pointer(), C.double(gain))
	}
}

func (ptr *QMediaRecorderControl) State() QMediaRecorder__State {
	defer qt.Recovering("QMediaRecorderControl::state")

	if ptr.Pointer() != nil {
		return QMediaRecorder__State(C.QMediaRecorderControl_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaRecorderControl) ConnectStateChanged(f func(state QMediaRecorder__State)) {
	defer qt.Recovering("connect QMediaRecorderControl::stateChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QMediaRecorderControl) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QMediaRecorderControl::stateChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQMediaRecorderControlStateChanged
func callbackQMediaRecorderControlStateChanged(ptr unsafe.Pointer, ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QMediaRecorderControl::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QMediaRecorder__State))(QMediaRecorder__State(state))
	}

}

func (ptr *QMediaRecorderControl) StateChanged(state QMediaRecorder__State) {
	defer qt.Recovering("QMediaRecorderControl::stateChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_StateChanged(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QMediaRecorderControl) Status() QMediaRecorder__Status {
	defer qt.Recovering("QMediaRecorderControl::status")

	if ptr.Pointer() != nil {
		return QMediaRecorder__Status(C.QMediaRecorderControl_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaRecorderControl) ConnectStatusChanged(f func(status QMediaRecorder__Status)) {
	defer qt.Recovering("connect QMediaRecorderControl::statusChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ConnectStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "statusChanged", f)
	}
}

func (ptr *QMediaRecorderControl) DisconnectStatusChanged() {
	defer qt.Recovering("disconnect QMediaRecorderControl::statusChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "statusChanged")
	}
}

//export callbackQMediaRecorderControlStatusChanged
func callbackQMediaRecorderControlStatusChanged(ptr unsafe.Pointer, ptrName *C.char, status C.int) {
	defer qt.Recovering("callback QMediaRecorderControl::statusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "statusChanged"); signal != nil {
		signal.(func(QMediaRecorder__Status))(QMediaRecorder__Status(status))
	}

}

func (ptr *QMediaRecorderControl) StatusChanged(status QMediaRecorder__Status) {
	defer qt.Recovering("QMediaRecorderControl::statusChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_StatusChanged(ptr.Pointer(), C.int(status))
	}
}

func (ptr *QMediaRecorderControl) Volume() float64 {
	defer qt.Recovering("QMediaRecorderControl::volume")

	if ptr.Pointer() != nil {
		return float64(C.QMediaRecorderControl_Volume(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaRecorderControl) ConnectVolumeChanged(f func(gain float64)) {
	defer qt.Recovering("connect QMediaRecorderControl::volumeChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ConnectVolumeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "volumeChanged", f)
	}
}

func (ptr *QMediaRecorderControl) DisconnectVolumeChanged() {
	defer qt.Recovering("disconnect QMediaRecorderControl::volumeChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DisconnectVolumeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "volumeChanged")
	}
}

//export callbackQMediaRecorderControlVolumeChanged
func callbackQMediaRecorderControlVolumeChanged(ptr unsafe.Pointer, ptrName *C.char, gain C.double) {
	defer qt.Recovering("callback QMediaRecorderControl::volumeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "volumeChanged"); signal != nil {
		signal.(func(float64))(float64(gain))
	}

}

func (ptr *QMediaRecorderControl) VolumeChanged(gain float64) {
	defer qt.Recovering("QMediaRecorderControl::volumeChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_VolumeChanged(ptr.Pointer(), C.double(gain))
	}
}

func (ptr *QMediaRecorderControl) DestroyQMediaRecorderControl() {
	defer qt.Recovering("QMediaRecorderControl::~QMediaRecorderControl")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DestroyQMediaRecorderControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMediaRecorderControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMediaRecorderControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMediaRecorderControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMediaRecorderControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMediaRecorderControlTimerEvent
func callbackQMediaRecorderControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaRecorderControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMediaRecorderControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMediaRecorderControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaRecorderControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaRecorderControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaRecorderControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaRecorderControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMediaRecorderControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMediaRecorderControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMediaRecorderControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMediaRecorderControlChildEvent
func callbackQMediaRecorderControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaRecorderControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMediaRecorderControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMediaRecorderControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaRecorderControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaRecorderControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaRecorderControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaRecorderControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMediaRecorderControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMediaRecorderControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMediaRecorderControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMediaRecorderControlCustomEvent
func callbackQMediaRecorderControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaRecorderControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMediaRecorderControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMediaRecorderControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaRecorderControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMediaRecorderControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaRecorderControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QMediaResource struct {
	ptr unsafe.Pointer
}

type QMediaResource_ITF interface {
	QMediaResource_PTR() *QMediaResource
}

func (p *QMediaResource) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMediaResource) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMediaResource(ptr QMediaResource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaResource_PTR().Pointer()
	}
	return nil
}

func NewQMediaResourceFromPointer(ptr unsafe.Pointer) *QMediaResource {
	var n = new(QMediaResource)
	n.SetPointer(ptr)
	return n
}

func newQMediaResourceFromPointer(ptr unsafe.Pointer) *QMediaResource {
	var n = NewQMediaResourceFromPointer(ptr)
	return n
}

func (ptr *QMediaResource) QMediaResource_PTR() *QMediaResource {
	return ptr
}

func NewQMediaResource() *QMediaResource {
	defer qt.Recovering("QMediaResource::QMediaResource")

	return newQMediaResourceFromPointer(C.QMediaResource_NewQMediaResource())
}

func NewQMediaResource4(other QMediaResource_ITF) *QMediaResource {
	defer qt.Recovering("QMediaResource::QMediaResource")

	return newQMediaResourceFromPointer(C.QMediaResource_NewQMediaResource4(PointerFromQMediaResource(other)))
}

func NewQMediaResource3(request network.QNetworkRequest_ITF, mimeType string) *QMediaResource {
	defer qt.Recovering("QMediaResource::QMediaResource")

	return newQMediaResourceFromPointer(C.QMediaResource_NewQMediaResource3(network.PointerFromQNetworkRequest(request), C.CString(mimeType)))
}

func NewQMediaResource2(url core.QUrl_ITF, mimeType string) *QMediaResource {
	defer qt.Recovering("QMediaResource::QMediaResource")

	return newQMediaResourceFromPointer(C.QMediaResource_NewQMediaResource2(core.PointerFromQUrl(url), C.CString(mimeType)))
}

func (ptr *QMediaResource) AudioBitRate() int {
	defer qt.Recovering("QMediaResource::audioBitRate")

	if ptr.Pointer() != nil {
		return int(C.QMediaResource_AudioBitRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaResource) AudioCodec() string {
	defer qt.Recovering("QMediaResource::audioCodec")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaResource_AudioCodec(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaResource) ChannelCount() int {
	defer qt.Recovering("QMediaResource::channelCount")

	if ptr.Pointer() != nil {
		return int(C.QMediaResource_ChannelCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaResource) DataSize() int64 {
	defer qt.Recovering("QMediaResource::dataSize")

	if ptr.Pointer() != nil {
		return int64(C.QMediaResource_DataSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaResource) IsNull() bool {
	defer qt.Recovering("QMediaResource::isNull")

	if ptr.Pointer() != nil {
		return C.QMediaResource_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaResource) Language() string {
	defer qt.Recovering("QMediaResource::language")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaResource_Language(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaResource) MimeType() string {
	defer qt.Recovering("QMediaResource::mimeType")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaResource_MimeType(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaResource) Request() *network.QNetworkRequest {
	defer qt.Recovering("QMediaResource::request")

	if ptr.Pointer() != nil {
		return network.NewQNetworkRequestFromPointer(C.QMediaResource_Request(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaResource) Resolution() *core.QSize {
	defer qt.Recovering("QMediaResource::resolution")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QMediaResource_Resolution(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaResource) SampleRate() int {
	defer qt.Recovering("QMediaResource::sampleRate")

	if ptr.Pointer() != nil {
		return int(C.QMediaResource_SampleRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaResource) SetAudioBitRate(rate int) {
	defer qt.Recovering("QMediaResource::setAudioBitRate")

	if ptr.Pointer() != nil {
		C.QMediaResource_SetAudioBitRate(ptr.Pointer(), C.int(rate))
	}
}

func (ptr *QMediaResource) SetAudioCodec(codec string) {
	defer qt.Recovering("QMediaResource::setAudioCodec")

	if ptr.Pointer() != nil {
		C.QMediaResource_SetAudioCodec(ptr.Pointer(), C.CString(codec))
	}
}

func (ptr *QMediaResource) SetChannelCount(channels int) {
	defer qt.Recovering("QMediaResource::setChannelCount")

	if ptr.Pointer() != nil {
		C.QMediaResource_SetChannelCount(ptr.Pointer(), C.int(channels))
	}
}

func (ptr *QMediaResource) SetDataSize(size int64) {
	defer qt.Recovering("QMediaResource::setDataSize")

	if ptr.Pointer() != nil {
		C.QMediaResource_SetDataSize(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QMediaResource) SetLanguage(language string) {
	defer qt.Recovering("QMediaResource::setLanguage")

	if ptr.Pointer() != nil {
		C.QMediaResource_SetLanguage(ptr.Pointer(), C.CString(language))
	}
}

func (ptr *QMediaResource) SetResolution(resolution core.QSize_ITF) {
	defer qt.Recovering("QMediaResource::setResolution")

	if ptr.Pointer() != nil {
		C.QMediaResource_SetResolution(ptr.Pointer(), core.PointerFromQSize(resolution))
	}
}

func (ptr *QMediaResource) SetResolution2(width int, height int) {
	defer qt.Recovering("QMediaResource::setResolution")

	if ptr.Pointer() != nil {
		C.QMediaResource_SetResolution2(ptr.Pointer(), C.int(width), C.int(height))
	}
}

func (ptr *QMediaResource) SetSampleRate(sampleRate int) {
	defer qt.Recovering("QMediaResource::setSampleRate")

	if ptr.Pointer() != nil {
		C.QMediaResource_SetSampleRate(ptr.Pointer(), C.int(sampleRate))
	}
}

func (ptr *QMediaResource) SetVideoBitRate(rate int) {
	defer qt.Recovering("QMediaResource::setVideoBitRate")

	if ptr.Pointer() != nil {
		C.QMediaResource_SetVideoBitRate(ptr.Pointer(), C.int(rate))
	}
}

func (ptr *QMediaResource) SetVideoCodec(codec string) {
	defer qt.Recovering("QMediaResource::setVideoCodec")

	if ptr.Pointer() != nil {
		C.QMediaResource_SetVideoCodec(ptr.Pointer(), C.CString(codec))
	}
}

func (ptr *QMediaResource) Url() *core.QUrl {
	defer qt.Recovering("QMediaResource::url")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QMediaResource_Url(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaResource) VideoBitRate() int {
	defer qt.Recovering("QMediaResource::videoBitRate")

	if ptr.Pointer() != nil {
		return int(C.QMediaResource_VideoBitRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaResource) VideoCodec() string {
	defer qt.Recovering("QMediaResource::videoCodec")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaResource_VideoCodec(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaResource) DestroyQMediaResource() {
	defer qt.Recovering("QMediaResource::~QMediaResource")

	if ptr.Pointer() != nil {
		C.QMediaResource_DestroyQMediaResource(ptr.Pointer())
	}
}

type QMediaService struct {
	core.QObject
}

type QMediaService_ITF interface {
	core.QObject_ITF
	QMediaService_PTR() *QMediaService
}

func PointerFromQMediaService(ptr QMediaService_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaService_PTR().Pointer()
	}
	return nil
}

func NewQMediaServiceFromPointer(ptr unsafe.Pointer) *QMediaService {
	var n = new(QMediaService)
	n.SetPointer(ptr)
	return n
}

func newQMediaServiceFromPointer(ptr unsafe.Pointer) *QMediaService {
	var n = NewQMediaServiceFromPointer(ptr)
	for len(n.ObjectName()) < len("QMediaService_") {
		n.SetObjectName("QMediaService_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaService) QMediaService_PTR() *QMediaService {
	return ptr
}

func (ptr *QMediaService) ReleaseControl(control QMediaControl_ITF) {
	defer qt.Recovering("QMediaService::releaseControl")

	if ptr.Pointer() != nil {
		C.QMediaService_ReleaseControl(ptr.Pointer(), PointerFromQMediaControl(control))
	}
}

func (ptr *QMediaService) RequestControl(interfa string) *QMediaControl {
	defer qt.Recovering("QMediaService::requestControl")

	if ptr.Pointer() != nil {
		return NewQMediaControlFromPointer(C.QMediaService_RequestControl(ptr.Pointer(), C.CString(interfa)))
	}
	return nil
}

func (ptr *QMediaService) RequestControl2() unsafe.Pointer {
	defer qt.Recovering("QMediaService::requestControl")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QMediaService_RequestControl2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaService) DestroyQMediaService() {
	defer qt.Recovering("QMediaService::~QMediaService")

	if ptr.Pointer() != nil {
		C.QMediaService_DestroyQMediaService(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMediaService) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMediaService::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMediaService) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMediaService::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMediaServiceTimerEvent
func callbackQMediaServiceTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaService::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMediaServiceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMediaService) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaService::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaService_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaService) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaService::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaService_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaService) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMediaService::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMediaService) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMediaService::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMediaServiceChildEvent
func callbackQMediaServiceChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaService::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMediaServiceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMediaService) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaService::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaService_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaService) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaService::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaService_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaService) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMediaService::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMediaService) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMediaService::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMediaServiceCustomEvent
func callbackQMediaServiceCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaService::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMediaServiceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMediaService) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaService::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaService_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMediaService) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaService::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaService_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QMediaServiceCameraInfoInterface struct {
	ptr unsafe.Pointer
}

type QMediaServiceCameraInfoInterface_ITF interface {
	QMediaServiceCameraInfoInterface_PTR() *QMediaServiceCameraInfoInterface
}

func (p *QMediaServiceCameraInfoInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMediaServiceCameraInfoInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMediaServiceCameraInfoInterface(ptr QMediaServiceCameraInfoInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaServiceCameraInfoInterface_PTR().Pointer()
	}
	return nil
}

func NewQMediaServiceCameraInfoInterfaceFromPointer(ptr unsafe.Pointer) *QMediaServiceCameraInfoInterface {
	var n = new(QMediaServiceCameraInfoInterface)
	n.SetPointer(ptr)
	return n
}

func newQMediaServiceCameraInfoInterfaceFromPointer(ptr unsafe.Pointer) *QMediaServiceCameraInfoInterface {
	var n = NewQMediaServiceCameraInfoInterfaceFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QMediaServiceCameraInfoInterface_") {
		n.SetObjectNameAbs("QMediaServiceCameraInfoInterface_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaServiceCameraInfoInterface) QMediaServiceCameraInfoInterface_PTR() *QMediaServiceCameraInfoInterface {
	return ptr
}

func (ptr *QMediaServiceCameraInfoInterface) CameraOrientation(device string) int {
	defer qt.Recovering("QMediaServiceCameraInfoInterface::cameraOrientation")

	if ptr.Pointer() != nil {
		return int(C.QMediaServiceCameraInfoInterface_CameraOrientation(ptr.Pointer(), C.CString(device)))
	}
	return 0
}

func (ptr *QMediaServiceCameraInfoInterface) CameraPosition(device string) QCamera__Position {
	defer qt.Recovering("QMediaServiceCameraInfoInterface::cameraPosition")

	if ptr.Pointer() != nil {
		return QCamera__Position(C.QMediaServiceCameraInfoInterface_CameraPosition(ptr.Pointer(), C.CString(device)))
	}
	return 0
}

func (ptr *QMediaServiceCameraInfoInterface) DestroyQMediaServiceCameraInfoInterface() {
	defer qt.Recovering("QMediaServiceCameraInfoInterface::~QMediaServiceCameraInfoInterface")

	if ptr.Pointer() != nil {
		C.QMediaServiceCameraInfoInterface_DestroyQMediaServiceCameraInfoInterface(ptr.Pointer())
	}
}

func (ptr *QMediaServiceCameraInfoInterface) ObjectNameAbs() string {
	defer qt.Recovering("QMediaServiceCameraInfoInterface::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaServiceCameraInfoInterface_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaServiceCameraInfoInterface) SetObjectNameAbs(name string) {
	defer qt.Recovering("QMediaServiceCameraInfoInterface::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QMediaServiceCameraInfoInterface_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QMediaServiceDefaultDeviceInterface struct {
	ptr unsafe.Pointer
}

type QMediaServiceDefaultDeviceInterface_ITF interface {
	QMediaServiceDefaultDeviceInterface_PTR() *QMediaServiceDefaultDeviceInterface
}

func (p *QMediaServiceDefaultDeviceInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMediaServiceDefaultDeviceInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMediaServiceDefaultDeviceInterface(ptr QMediaServiceDefaultDeviceInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaServiceDefaultDeviceInterface_PTR().Pointer()
	}
	return nil
}

func NewQMediaServiceDefaultDeviceInterfaceFromPointer(ptr unsafe.Pointer) *QMediaServiceDefaultDeviceInterface {
	var n = new(QMediaServiceDefaultDeviceInterface)
	n.SetPointer(ptr)
	return n
}

func newQMediaServiceDefaultDeviceInterfaceFromPointer(ptr unsafe.Pointer) *QMediaServiceDefaultDeviceInterface {
	var n = NewQMediaServiceDefaultDeviceInterfaceFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QMediaServiceDefaultDeviceInterface_") {
		n.SetObjectNameAbs("QMediaServiceDefaultDeviceInterface_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaServiceDefaultDeviceInterface) QMediaServiceDefaultDeviceInterface_PTR() *QMediaServiceDefaultDeviceInterface {
	return ptr
}

func (ptr *QMediaServiceDefaultDeviceInterface) DefaultDevice(service string) string {
	defer qt.Recovering("QMediaServiceDefaultDeviceInterface::defaultDevice")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaServiceDefaultDeviceInterface_DefaultDevice(ptr.Pointer(), C.CString(service)))
	}
	return ""
}

func (ptr *QMediaServiceDefaultDeviceInterface) DestroyQMediaServiceDefaultDeviceInterface() {
	defer qt.Recovering("QMediaServiceDefaultDeviceInterface::~QMediaServiceDefaultDeviceInterface")

	if ptr.Pointer() != nil {
		C.QMediaServiceDefaultDeviceInterface_DestroyQMediaServiceDefaultDeviceInterface(ptr.Pointer())
	}
}

func (ptr *QMediaServiceDefaultDeviceInterface) ObjectNameAbs() string {
	defer qt.Recovering("QMediaServiceDefaultDeviceInterface::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaServiceDefaultDeviceInterface_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaServiceDefaultDeviceInterface) SetObjectNameAbs(name string) {
	defer qt.Recovering("QMediaServiceDefaultDeviceInterface::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QMediaServiceDefaultDeviceInterface_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QMediaServiceFeaturesInterface struct {
	ptr unsafe.Pointer
}

type QMediaServiceFeaturesInterface_ITF interface {
	QMediaServiceFeaturesInterface_PTR() *QMediaServiceFeaturesInterface
}

func (p *QMediaServiceFeaturesInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMediaServiceFeaturesInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMediaServiceFeaturesInterface(ptr QMediaServiceFeaturesInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaServiceFeaturesInterface_PTR().Pointer()
	}
	return nil
}

func NewQMediaServiceFeaturesInterfaceFromPointer(ptr unsafe.Pointer) *QMediaServiceFeaturesInterface {
	var n = new(QMediaServiceFeaturesInterface)
	n.SetPointer(ptr)
	return n
}

func newQMediaServiceFeaturesInterfaceFromPointer(ptr unsafe.Pointer) *QMediaServiceFeaturesInterface {
	var n = NewQMediaServiceFeaturesInterfaceFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QMediaServiceFeaturesInterface_") {
		n.SetObjectNameAbs("QMediaServiceFeaturesInterface_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaServiceFeaturesInterface) QMediaServiceFeaturesInterface_PTR() *QMediaServiceFeaturesInterface {
	return ptr
}

func (ptr *QMediaServiceFeaturesInterface) DestroyQMediaServiceFeaturesInterface() {
	defer qt.Recovering("QMediaServiceFeaturesInterface::~QMediaServiceFeaturesInterface")

	if ptr.Pointer() != nil {
		C.QMediaServiceFeaturesInterface_DestroyQMediaServiceFeaturesInterface(ptr.Pointer())
	}
}

func (ptr *QMediaServiceFeaturesInterface) ObjectNameAbs() string {
	defer qt.Recovering("QMediaServiceFeaturesInterface::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaServiceFeaturesInterface_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaServiceFeaturesInterface) SetObjectNameAbs(name string) {
	defer qt.Recovering("QMediaServiceFeaturesInterface::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QMediaServiceFeaturesInterface_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QMediaServiceProviderPlugin struct {
	core.QObject
}

type QMediaServiceProviderPlugin_ITF interface {
	core.QObject_ITF
	QMediaServiceProviderPlugin_PTR() *QMediaServiceProviderPlugin
}

func PointerFromQMediaServiceProviderPlugin(ptr QMediaServiceProviderPlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaServiceProviderPlugin_PTR().Pointer()
	}
	return nil
}

func NewQMediaServiceProviderPluginFromPointer(ptr unsafe.Pointer) *QMediaServiceProviderPlugin {
	var n = new(QMediaServiceProviderPlugin)
	n.SetPointer(ptr)
	return n
}

func newQMediaServiceProviderPluginFromPointer(ptr unsafe.Pointer) *QMediaServiceProviderPlugin {
	var n = NewQMediaServiceProviderPluginFromPointer(ptr)
	for len(n.ObjectName()) < len("QMediaServiceProviderPlugin_") {
		n.SetObjectName("QMediaServiceProviderPlugin_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaServiceProviderPlugin) QMediaServiceProviderPlugin_PTR() *QMediaServiceProviderPlugin {
	return ptr
}

func (ptr *QMediaServiceProviderPlugin) Create(key string) *QMediaService {
	defer qt.Recovering("QMediaServiceProviderPlugin::create")

	if ptr.Pointer() != nil {
		return NewQMediaServiceFromPointer(C.QMediaServiceProviderPlugin_Create(ptr.Pointer(), C.CString(key)))
	}
	return nil
}

func (ptr *QMediaServiceProviderPlugin) Release(service QMediaService_ITF) {
	defer qt.Recovering("QMediaServiceProviderPlugin::release")

	if ptr.Pointer() != nil {
		C.QMediaServiceProviderPlugin_Release(ptr.Pointer(), PointerFromQMediaService(service))
	}
}

func (ptr *QMediaServiceProviderPlugin) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMediaServiceProviderPlugin::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMediaServiceProviderPlugin) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMediaServiceProviderPlugin::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMediaServiceProviderPluginTimerEvent
func callbackQMediaServiceProviderPluginTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaServiceProviderPlugin::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMediaServiceProviderPluginFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMediaServiceProviderPlugin) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaServiceProviderPlugin::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaServiceProviderPlugin_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaServiceProviderPlugin) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaServiceProviderPlugin::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaServiceProviderPlugin_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaServiceProviderPlugin) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMediaServiceProviderPlugin::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMediaServiceProviderPlugin) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMediaServiceProviderPlugin::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMediaServiceProviderPluginChildEvent
func callbackQMediaServiceProviderPluginChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaServiceProviderPlugin::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMediaServiceProviderPluginFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMediaServiceProviderPlugin) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaServiceProviderPlugin::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaServiceProviderPlugin_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaServiceProviderPlugin) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaServiceProviderPlugin::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaServiceProviderPlugin_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaServiceProviderPlugin) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMediaServiceProviderPlugin::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMediaServiceProviderPlugin) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMediaServiceProviderPlugin::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMediaServiceProviderPluginCustomEvent
func callbackQMediaServiceProviderPluginCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaServiceProviderPlugin::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMediaServiceProviderPluginFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMediaServiceProviderPlugin) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaServiceProviderPlugin::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaServiceProviderPlugin_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMediaServiceProviderPlugin) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaServiceProviderPlugin::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaServiceProviderPlugin_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QMediaServiceSupportedDevicesInterface struct {
	ptr unsafe.Pointer
}

type QMediaServiceSupportedDevicesInterface_ITF interface {
	QMediaServiceSupportedDevicesInterface_PTR() *QMediaServiceSupportedDevicesInterface
}

func (p *QMediaServiceSupportedDevicesInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMediaServiceSupportedDevicesInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMediaServiceSupportedDevicesInterface(ptr QMediaServiceSupportedDevicesInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaServiceSupportedDevicesInterface_PTR().Pointer()
	}
	return nil
}

func NewQMediaServiceSupportedDevicesInterfaceFromPointer(ptr unsafe.Pointer) *QMediaServiceSupportedDevicesInterface {
	var n = new(QMediaServiceSupportedDevicesInterface)
	n.SetPointer(ptr)
	return n
}

func newQMediaServiceSupportedDevicesInterfaceFromPointer(ptr unsafe.Pointer) *QMediaServiceSupportedDevicesInterface {
	var n = NewQMediaServiceSupportedDevicesInterfaceFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QMediaServiceSupportedDevicesInterface_") {
		n.SetObjectNameAbs("QMediaServiceSupportedDevicesInterface_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaServiceSupportedDevicesInterface) QMediaServiceSupportedDevicesInterface_PTR() *QMediaServiceSupportedDevicesInterface {
	return ptr
}

func (ptr *QMediaServiceSupportedDevicesInterface) DeviceDescription(service string, device string) string {
	defer qt.Recovering("QMediaServiceSupportedDevicesInterface::deviceDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaServiceSupportedDevicesInterface_DeviceDescription(ptr.Pointer(), C.CString(service), C.CString(device)))
	}
	return ""
}

func (ptr *QMediaServiceSupportedDevicesInterface) DestroyQMediaServiceSupportedDevicesInterface() {
	defer qt.Recovering("QMediaServiceSupportedDevicesInterface::~QMediaServiceSupportedDevicesInterface")

	if ptr.Pointer() != nil {
		C.QMediaServiceSupportedDevicesInterface_DestroyQMediaServiceSupportedDevicesInterface(ptr.Pointer())
	}
}

func (ptr *QMediaServiceSupportedDevicesInterface) ObjectNameAbs() string {
	defer qt.Recovering("QMediaServiceSupportedDevicesInterface::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaServiceSupportedDevicesInterface_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaServiceSupportedDevicesInterface) SetObjectNameAbs(name string) {
	defer qt.Recovering("QMediaServiceSupportedDevicesInterface::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QMediaServiceSupportedDevicesInterface_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QMediaServiceSupportedFormatsInterface struct {
	ptr unsafe.Pointer
}

type QMediaServiceSupportedFormatsInterface_ITF interface {
	QMediaServiceSupportedFormatsInterface_PTR() *QMediaServiceSupportedFormatsInterface
}

func (p *QMediaServiceSupportedFormatsInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMediaServiceSupportedFormatsInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMediaServiceSupportedFormatsInterface(ptr QMediaServiceSupportedFormatsInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaServiceSupportedFormatsInterface_PTR().Pointer()
	}
	return nil
}

func NewQMediaServiceSupportedFormatsInterfaceFromPointer(ptr unsafe.Pointer) *QMediaServiceSupportedFormatsInterface {
	var n = new(QMediaServiceSupportedFormatsInterface)
	n.SetPointer(ptr)
	return n
}

func newQMediaServiceSupportedFormatsInterfaceFromPointer(ptr unsafe.Pointer) *QMediaServiceSupportedFormatsInterface {
	var n = NewQMediaServiceSupportedFormatsInterfaceFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QMediaServiceSupportedFormatsInterface_") {
		n.SetObjectNameAbs("QMediaServiceSupportedFormatsInterface_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaServiceSupportedFormatsInterface) QMediaServiceSupportedFormatsInterface_PTR() *QMediaServiceSupportedFormatsInterface {
	return ptr
}

func (ptr *QMediaServiceSupportedFormatsInterface) HasSupport(mimeType string, codecs []string) QMultimedia__SupportEstimate {
	defer qt.Recovering("QMediaServiceSupportedFormatsInterface::hasSupport")

	if ptr.Pointer() != nil {
		return QMultimedia__SupportEstimate(C.QMediaServiceSupportedFormatsInterface_HasSupport(ptr.Pointer(), C.CString(mimeType), C.CString(strings.Join(codecs, "|"))))
	}
	return 0
}

func (ptr *QMediaServiceSupportedFormatsInterface) SupportedMimeTypes() []string {
	defer qt.Recovering("QMediaServiceSupportedFormatsInterface::supportedMimeTypes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMediaServiceSupportedFormatsInterface_SupportedMimeTypes(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QMediaServiceSupportedFormatsInterface) DestroyQMediaServiceSupportedFormatsInterface() {
	defer qt.Recovering("QMediaServiceSupportedFormatsInterface::~QMediaServiceSupportedFormatsInterface")

	if ptr.Pointer() != nil {
		C.QMediaServiceSupportedFormatsInterface_DestroyQMediaServiceSupportedFormatsInterface(ptr.Pointer())
	}
}

func (ptr *QMediaServiceSupportedFormatsInterface) ObjectNameAbs() string {
	defer qt.Recovering("QMediaServiceSupportedFormatsInterface::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaServiceSupportedFormatsInterface_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaServiceSupportedFormatsInterface) SetObjectNameAbs(name string) {
	defer qt.Recovering("QMediaServiceSupportedFormatsInterface::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QMediaServiceSupportedFormatsInterface_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QMediaStreamsControl struct {
	QMediaControl
}

type QMediaStreamsControl_ITF interface {
	QMediaControl_ITF
	QMediaStreamsControl_PTR() *QMediaStreamsControl
}

func PointerFromQMediaStreamsControl(ptr QMediaStreamsControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaStreamsControl_PTR().Pointer()
	}
	return nil
}

func NewQMediaStreamsControlFromPointer(ptr unsafe.Pointer) *QMediaStreamsControl {
	var n = new(QMediaStreamsControl)
	n.SetPointer(ptr)
	return n
}

func newQMediaStreamsControlFromPointer(ptr unsafe.Pointer) *QMediaStreamsControl {
	var n = NewQMediaStreamsControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QMediaStreamsControl_") {
		n.SetObjectName("QMediaStreamsControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaStreamsControl) QMediaStreamsControl_PTR() *QMediaStreamsControl {
	return ptr
}

//QMediaStreamsControl::StreamType
type QMediaStreamsControl__StreamType int64

const (
	QMediaStreamsControl__UnknownStream    = QMediaStreamsControl__StreamType(0)
	QMediaStreamsControl__VideoStream      = QMediaStreamsControl__StreamType(1)
	QMediaStreamsControl__AudioStream      = QMediaStreamsControl__StreamType(2)
	QMediaStreamsControl__SubPictureStream = QMediaStreamsControl__StreamType(3)
	QMediaStreamsControl__DataStream       = QMediaStreamsControl__StreamType(4)
)

func (ptr *QMediaStreamsControl) ConnectActiveStreamsChanged(f func()) {
	defer qt.Recovering("connect QMediaStreamsControl::activeStreamsChanged")

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_ConnectActiveStreamsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeStreamsChanged", f)
	}
}

func (ptr *QMediaStreamsControl) DisconnectActiveStreamsChanged() {
	defer qt.Recovering("disconnect QMediaStreamsControl::activeStreamsChanged")

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_DisconnectActiveStreamsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeStreamsChanged")
	}
}

//export callbackQMediaStreamsControlActiveStreamsChanged
func callbackQMediaStreamsControlActiveStreamsChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QMediaStreamsControl::activeStreamsChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "activeStreamsChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QMediaStreamsControl) ActiveStreamsChanged() {
	defer qt.Recovering("QMediaStreamsControl::activeStreamsChanged")

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_ActiveStreamsChanged(ptr.Pointer())
	}
}

func (ptr *QMediaStreamsControl) IsActive(stream int) bool {
	defer qt.Recovering("QMediaStreamsControl::isActive")

	if ptr.Pointer() != nil {
		return C.QMediaStreamsControl_IsActive(ptr.Pointer(), C.int(stream)) != 0
	}
	return false
}

func (ptr *QMediaStreamsControl) MetaData(stream int, key string) *core.QVariant {
	defer qt.Recovering("QMediaStreamsControl::metaData")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QMediaStreamsControl_MetaData(ptr.Pointer(), C.int(stream), C.CString(key)))
	}
	return nil
}

func (ptr *QMediaStreamsControl) SetActive(stream int, state bool) {
	defer qt.Recovering("QMediaStreamsControl::setActive")

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_SetActive(ptr.Pointer(), C.int(stream), C.int(qt.GoBoolToInt(state)))
	}
}

func (ptr *QMediaStreamsControl) StreamCount() int {
	defer qt.Recovering("QMediaStreamsControl::streamCount")

	if ptr.Pointer() != nil {
		return int(C.QMediaStreamsControl_StreamCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaStreamsControl) StreamType(stream int) QMediaStreamsControl__StreamType {
	defer qt.Recovering("QMediaStreamsControl::streamType")

	if ptr.Pointer() != nil {
		return QMediaStreamsControl__StreamType(C.QMediaStreamsControl_StreamType(ptr.Pointer(), C.int(stream)))
	}
	return 0
}

func (ptr *QMediaStreamsControl) ConnectStreamsChanged(f func()) {
	defer qt.Recovering("connect QMediaStreamsControl::streamsChanged")

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_ConnectStreamsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "streamsChanged", f)
	}
}

func (ptr *QMediaStreamsControl) DisconnectStreamsChanged() {
	defer qt.Recovering("disconnect QMediaStreamsControl::streamsChanged")

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_DisconnectStreamsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "streamsChanged")
	}
}

//export callbackQMediaStreamsControlStreamsChanged
func callbackQMediaStreamsControlStreamsChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QMediaStreamsControl::streamsChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "streamsChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QMediaStreamsControl) StreamsChanged() {
	defer qt.Recovering("QMediaStreamsControl::streamsChanged")

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_StreamsChanged(ptr.Pointer())
	}
}

func (ptr *QMediaStreamsControl) DestroyQMediaStreamsControl() {
	defer qt.Recovering("QMediaStreamsControl::~QMediaStreamsControl")

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_DestroyQMediaStreamsControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMediaStreamsControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMediaStreamsControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMediaStreamsControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMediaStreamsControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMediaStreamsControlTimerEvent
func callbackQMediaStreamsControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaStreamsControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMediaStreamsControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMediaStreamsControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaStreamsControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaStreamsControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaStreamsControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaStreamsControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMediaStreamsControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMediaStreamsControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMediaStreamsControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMediaStreamsControlChildEvent
func callbackQMediaStreamsControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaStreamsControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMediaStreamsControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMediaStreamsControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaStreamsControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaStreamsControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaStreamsControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaStreamsControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMediaStreamsControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMediaStreamsControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMediaStreamsControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMediaStreamsControlCustomEvent
func callbackQMediaStreamsControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaStreamsControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMediaStreamsControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMediaStreamsControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaStreamsControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMediaStreamsControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaStreamsControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QMediaTimeInterval struct {
	ptr unsafe.Pointer
}

type QMediaTimeInterval_ITF interface {
	QMediaTimeInterval_PTR() *QMediaTimeInterval
}

func (p *QMediaTimeInterval) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMediaTimeInterval) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMediaTimeInterval(ptr QMediaTimeInterval_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaTimeInterval_PTR().Pointer()
	}
	return nil
}

func NewQMediaTimeIntervalFromPointer(ptr unsafe.Pointer) *QMediaTimeInterval {
	var n = new(QMediaTimeInterval)
	n.SetPointer(ptr)
	return n
}

func newQMediaTimeIntervalFromPointer(ptr unsafe.Pointer) *QMediaTimeInterval {
	var n = NewQMediaTimeIntervalFromPointer(ptr)
	return n
}

func (ptr *QMediaTimeInterval) QMediaTimeInterval_PTR() *QMediaTimeInterval {
	return ptr
}

func NewQMediaTimeInterval() *QMediaTimeInterval {
	defer qt.Recovering("QMediaTimeInterval::QMediaTimeInterval")

	return newQMediaTimeIntervalFromPointer(C.QMediaTimeInterval_NewQMediaTimeInterval())
}

func NewQMediaTimeInterval3(other QMediaTimeInterval_ITF) *QMediaTimeInterval {
	defer qt.Recovering("QMediaTimeInterval::QMediaTimeInterval")

	return newQMediaTimeIntervalFromPointer(C.QMediaTimeInterval_NewQMediaTimeInterval3(PointerFromQMediaTimeInterval(other)))
}

func NewQMediaTimeInterval2(start int64, end int64) *QMediaTimeInterval {
	defer qt.Recovering("QMediaTimeInterval::QMediaTimeInterval")

	return newQMediaTimeIntervalFromPointer(C.QMediaTimeInterval_NewQMediaTimeInterval2(C.longlong(start), C.longlong(end)))
}

func (ptr *QMediaTimeInterval) Contains(time int64) bool {
	defer qt.Recovering("QMediaTimeInterval::contains")

	if ptr.Pointer() != nil {
		return C.QMediaTimeInterval_Contains(ptr.Pointer(), C.longlong(time)) != 0
	}
	return false
}

func (ptr *QMediaTimeInterval) End() int64 {
	defer qt.Recovering("QMediaTimeInterval::end")

	if ptr.Pointer() != nil {
		return int64(C.QMediaTimeInterval_End(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaTimeInterval) IsNormal() bool {
	defer qt.Recovering("QMediaTimeInterval::isNormal")

	if ptr.Pointer() != nil {
		return C.QMediaTimeInterval_IsNormal(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaTimeInterval) Normalized() *QMediaTimeInterval {
	defer qt.Recovering("QMediaTimeInterval::normalized")

	if ptr.Pointer() != nil {
		return NewQMediaTimeIntervalFromPointer(C.QMediaTimeInterval_Normalized(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaTimeInterval) Start() int64 {
	defer qt.Recovering("QMediaTimeInterval::start")

	if ptr.Pointer() != nil {
		return int64(C.QMediaTimeInterval_Start(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaTimeInterval) Translated(offset int64) *QMediaTimeInterval {
	defer qt.Recovering("QMediaTimeInterval::translated")

	if ptr.Pointer() != nil {
		return NewQMediaTimeIntervalFromPointer(C.QMediaTimeInterval_Translated(ptr.Pointer(), C.longlong(offset)))
	}
	return nil
}

type QMediaTimeRange struct {
	ptr unsafe.Pointer
}

type QMediaTimeRange_ITF interface {
	QMediaTimeRange_PTR() *QMediaTimeRange
}

func (p *QMediaTimeRange) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMediaTimeRange) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMediaTimeRange(ptr QMediaTimeRange_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaTimeRange_PTR().Pointer()
	}
	return nil
}

func NewQMediaTimeRangeFromPointer(ptr unsafe.Pointer) *QMediaTimeRange {
	var n = new(QMediaTimeRange)
	n.SetPointer(ptr)
	return n
}

func newQMediaTimeRangeFromPointer(ptr unsafe.Pointer) *QMediaTimeRange {
	var n = NewQMediaTimeRangeFromPointer(ptr)
	return n
}

func (ptr *QMediaTimeRange) QMediaTimeRange_PTR() *QMediaTimeRange {
	return ptr
}

func NewQMediaTimeRange() *QMediaTimeRange {
	defer qt.Recovering("QMediaTimeRange::QMediaTimeRange")

	return newQMediaTimeRangeFromPointer(C.QMediaTimeRange_NewQMediaTimeRange())
}

func NewQMediaTimeRange3(interval QMediaTimeInterval_ITF) *QMediaTimeRange {
	defer qt.Recovering("QMediaTimeRange::QMediaTimeRange")

	return newQMediaTimeRangeFromPointer(C.QMediaTimeRange_NewQMediaTimeRange3(PointerFromQMediaTimeInterval(interval)))
}

func NewQMediaTimeRange4(ran QMediaTimeRange_ITF) *QMediaTimeRange {
	defer qt.Recovering("QMediaTimeRange::QMediaTimeRange")

	return newQMediaTimeRangeFromPointer(C.QMediaTimeRange_NewQMediaTimeRange4(PointerFromQMediaTimeRange(ran)))
}

func NewQMediaTimeRange2(start int64, end int64) *QMediaTimeRange {
	defer qt.Recovering("QMediaTimeRange::QMediaTimeRange")

	return newQMediaTimeRangeFromPointer(C.QMediaTimeRange_NewQMediaTimeRange2(C.longlong(start), C.longlong(end)))
}

func (ptr *QMediaTimeRange) AddInterval(interval QMediaTimeInterval_ITF) {
	defer qt.Recovering("QMediaTimeRange::addInterval")

	if ptr.Pointer() != nil {
		C.QMediaTimeRange_AddInterval(ptr.Pointer(), PointerFromQMediaTimeInterval(interval))
	}
}

func (ptr *QMediaTimeRange) AddInterval2(start int64, end int64) {
	defer qt.Recovering("QMediaTimeRange::addInterval")

	if ptr.Pointer() != nil {
		C.QMediaTimeRange_AddInterval2(ptr.Pointer(), C.longlong(start), C.longlong(end))
	}
}

func (ptr *QMediaTimeRange) AddTimeRange(ran QMediaTimeRange_ITF) {
	defer qt.Recovering("QMediaTimeRange::addTimeRange")

	if ptr.Pointer() != nil {
		C.QMediaTimeRange_AddTimeRange(ptr.Pointer(), PointerFromQMediaTimeRange(ran))
	}
}

func (ptr *QMediaTimeRange) Clear() {
	defer qt.Recovering("QMediaTimeRange::clear")

	if ptr.Pointer() != nil {
		C.QMediaTimeRange_Clear(ptr.Pointer())
	}
}

func (ptr *QMediaTimeRange) Contains(time int64) bool {
	defer qt.Recovering("QMediaTimeRange::contains")

	if ptr.Pointer() != nil {
		return C.QMediaTimeRange_Contains(ptr.Pointer(), C.longlong(time)) != 0
	}
	return false
}

func (ptr *QMediaTimeRange) EarliestTime() int64 {
	defer qt.Recovering("QMediaTimeRange::earliestTime")

	if ptr.Pointer() != nil {
		return int64(C.QMediaTimeRange_EarliestTime(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaTimeRange) IsContinuous() bool {
	defer qt.Recovering("QMediaTimeRange::isContinuous")

	if ptr.Pointer() != nil {
		return C.QMediaTimeRange_IsContinuous(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaTimeRange) IsEmpty() bool {
	defer qt.Recovering("QMediaTimeRange::isEmpty")

	if ptr.Pointer() != nil {
		return C.QMediaTimeRange_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaTimeRange) LatestTime() int64 {
	defer qt.Recovering("QMediaTimeRange::latestTime")

	if ptr.Pointer() != nil {
		return int64(C.QMediaTimeRange_LatestTime(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaTimeRange) RemoveInterval(interval QMediaTimeInterval_ITF) {
	defer qt.Recovering("QMediaTimeRange::removeInterval")

	if ptr.Pointer() != nil {
		C.QMediaTimeRange_RemoveInterval(ptr.Pointer(), PointerFromQMediaTimeInterval(interval))
	}
}

func (ptr *QMediaTimeRange) RemoveInterval2(start int64, end int64) {
	defer qt.Recovering("QMediaTimeRange::removeInterval")

	if ptr.Pointer() != nil {
		C.QMediaTimeRange_RemoveInterval2(ptr.Pointer(), C.longlong(start), C.longlong(end))
	}
}

func (ptr *QMediaTimeRange) RemoveTimeRange(ran QMediaTimeRange_ITF) {
	defer qt.Recovering("QMediaTimeRange::removeTimeRange")

	if ptr.Pointer() != nil {
		C.QMediaTimeRange_RemoveTimeRange(ptr.Pointer(), PointerFromQMediaTimeRange(ran))
	}
}

func (ptr *QMediaTimeRange) DestroyQMediaTimeRange() {
	defer qt.Recovering("QMediaTimeRange::~QMediaTimeRange")

	if ptr.Pointer() != nil {
		C.QMediaTimeRange_DestroyQMediaTimeRange(ptr.Pointer())
	}
}

type QMediaVideoProbeControl struct {
	QMediaControl
}

type QMediaVideoProbeControl_ITF interface {
	QMediaControl_ITF
	QMediaVideoProbeControl_PTR() *QMediaVideoProbeControl
}

func PointerFromQMediaVideoProbeControl(ptr QMediaVideoProbeControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaVideoProbeControl_PTR().Pointer()
	}
	return nil
}

func NewQMediaVideoProbeControlFromPointer(ptr unsafe.Pointer) *QMediaVideoProbeControl {
	var n = new(QMediaVideoProbeControl)
	n.SetPointer(ptr)
	return n
}

func newQMediaVideoProbeControlFromPointer(ptr unsafe.Pointer) *QMediaVideoProbeControl {
	var n = NewQMediaVideoProbeControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QMediaVideoProbeControl_") {
		n.SetObjectName("QMediaVideoProbeControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaVideoProbeControl) QMediaVideoProbeControl_PTR() *QMediaVideoProbeControl {
	return ptr
}

func (ptr *QMediaVideoProbeControl) ConnectFlush(f func()) {
	defer qt.Recovering("connect QMediaVideoProbeControl::flush")

	if ptr.Pointer() != nil {
		C.QMediaVideoProbeControl_ConnectFlush(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "flush", f)
	}
}

func (ptr *QMediaVideoProbeControl) DisconnectFlush() {
	defer qt.Recovering("disconnect QMediaVideoProbeControl::flush")

	if ptr.Pointer() != nil {
		C.QMediaVideoProbeControl_DisconnectFlush(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "flush")
	}
}

//export callbackQMediaVideoProbeControlFlush
func callbackQMediaVideoProbeControlFlush(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QMediaVideoProbeControl::flush")

	if signal := qt.GetSignal(C.GoString(ptrName), "flush"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QMediaVideoProbeControl) Flush() {
	defer qt.Recovering("QMediaVideoProbeControl::flush")

	if ptr.Pointer() != nil {
		C.QMediaVideoProbeControl_Flush(ptr.Pointer())
	}
}

func (ptr *QMediaVideoProbeControl) ConnectVideoFrameProbed(f func(frame *QVideoFrame)) {
	defer qt.Recovering("connect QMediaVideoProbeControl::videoFrameProbed")

	if ptr.Pointer() != nil {
		C.QMediaVideoProbeControl_ConnectVideoFrameProbed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "videoFrameProbed", f)
	}
}

func (ptr *QMediaVideoProbeControl) DisconnectVideoFrameProbed() {
	defer qt.Recovering("disconnect QMediaVideoProbeControl::videoFrameProbed")

	if ptr.Pointer() != nil {
		C.QMediaVideoProbeControl_DisconnectVideoFrameProbed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "videoFrameProbed")
	}
}

//export callbackQMediaVideoProbeControlVideoFrameProbed
func callbackQMediaVideoProbeControlVideoFrameProbed(ptr unsafe.Pointer, ptrName *C.char, frame unsafe.Pointer) {
	defer qt.Recovering("callback QMediaVideoProbeControl::videoFrameProbed")

	if signal := qt.GetSignal(C.GoString(ptrName), "videoFrameProbed"); signal != nil {
		signal.(func(*QVideoFrame))(NewQVideoFrameFromPointer(frame))
	}

}

func (ptr *QMediaVideoProbeControl) VideoFrameProbed(frame QVideoFrame_ITF) {
	defer qt.Recovering("QMediaVideoProbeControl::videoFrameProbed")

	if ptr.Pointer() != nil {
		C.QMediaVideoProbeControl_VideoFrameProbed(ptr.Pointer(), PointerFromQVideoFrame(frame))
	}
}

func (ptr *QMediaVideoProbeControl) DestroyQMediaVideoProbeControl() {
	defer qt.Recovering("QMediaVideoProbeControl::~QMediaVideoProbeControl")

	if ptr.Pointer() != nil {
		C.QMediaVideoProbeControl_DestroyQMediaVideoProbeControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMediaVideoProbeControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMediaVideoProbeControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMediaVideoProbeControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMediaVideoProbeControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMediaVideoProbeControlTimerEvent
func callbackQMediaVideoProbeControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaVideoProbeControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMediaVideoProbeControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMediaVideoProbeControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaVideoProbeControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaVideoProbeControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaVideoProbeControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaVideoProbeControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaVideoProbeControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaVideoProbeControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMediaVideoProbeControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMediaVideoProbeControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMediaVideoProbeControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMediaVideoProbeControlChildEvent
func callbackQMediaVideoProbeControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaVideoProbeControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMediaVideoProbeControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMediaVideoProbeControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaVideoProbeControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaVideoProbeControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaVideoProbeControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaVideoProbeControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaVideoProbeControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaVideoProbeControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMediaVideoProbeControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMediaVideoProbeControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMediaVideoProbeControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMediaVideoProbeControlCustomEvent
func callbackQMediaVideoProbeControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaVideoProbeControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMediaVideoProbeControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMediaVideoProbeControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaVideoProbeControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaVideoProbeControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMediaVideoProbeControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaVideoProbeControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaVideoProbeControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QMetaDataReaderControl struct {
	QMediaControl
}

type QMetaDataReaderControl_ITF interface {
	QMediaControl_ITF
	QMetaDataReaderControl_PTR() *QMetaDataReaderControl
}

func PointerFromQMetaDataReaderControl(ptr QMetaDataReaderControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMetaDataReaderControl_PTR().Pointer()
	}
	return nil
}

func NewQMetaDataReaderControlFromPointer(ptr unsafe.Pointer) *QMetaDataReaderControl {
	var n = new(QMetaDataReaderControl)
	n.SetPointer(ptr)
	return n
}

func newQMetaDataReaderControlFromPointer(ptr unsafe.Pointer) *QMetaDataReaderControl {
	var n = NewQMetaDataReaderControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QMetaDataReaderControl_") {
		n.SetObjectName("QMetaDataReaderControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QMetaDataReaderControl) QMetaDataReaderControl_PTR() *QMetaDataReaderControl {
	return ptr
}

func (ptr *QMetaDataReaderControl) AvailableMetaData() []string {
	defer qt.Recovering("QMetaDataReaderControl::availableMetaData")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMetaDataReaderControl_AvailableMetaData(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QMetaDataReaderControl) IsMetaDataAvailable() bool {
	defer qt.Recovering("QMetaDataReaderControl::isMetaDataAvailable")

	if ptr.Pointer() != nil {
		return C.QMetaDataReaderControl_IsMetaDataAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMetaDataReaderControl) MetaData(key string) *core.QVariant {
	defer qt.Recovering("QMetaDataReaderControl::metaData")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QMetaDataReaderControl_MetaData(ptr.Pointer(), C.CString(key)))
	}
	return nil
}

func (ptr *QMetaDataReaderControl) ConnectMetaDataAvailableChanged(f func(available bool)) {
	defer qt.Recovering("connect QMetaDataReaderControl::metaDataAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_ConnectMetaDataAvailableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataAvailableChanged", f)
	}
}

func (ptr *QMetaDataReaderControl) DisconnectMetaDataAvailableChanged() {
	defer qt.Recovering("disconnect QMetaDataReaderControl::metaDataAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_DisconnectMetaDataAvailableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataAvailableChanged")
	}
}

//export callbackQMetaDataReaderControlMetaDataAvailableChanged
func callbackQMetaDataReaderControlMetaDataAvailableChanged(ptr unsafe.Pointer, ptrName *C.char, available C.int) {
	defer qt.Recovering("callback QMetaDataReaderControl::metaDataAvailableChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaDataAvailableChanged"); signal != nil {
		signal.(func(bool))(int(available) != 0)
	}

}

func (ptr *QMetaDataReaderControl) MetaDataAvailableChanged(available bool) {
	defer qt.Recovering("QMetaDataReaderControl::metaDataAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_MetaDataAvailableChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(available)))
	}
}

func (ptr *QMetaDataReaderControl) ConnectMetaDataChanged(f func()) {
	defer qt.Recovering("connect QMetaDataReaderControl::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_ConnectMetaDataChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataChanged", f)
	}
}

func (ptr *QMetaDataReaderControl) DisconnectMetaDataChanged() {
	defer qt.Recovering("disconnect QMetaDataReaderControl::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_DisconnectMetaDataChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataChanged")
	}
}

//export callbackQMetaDataReaderControlMetaDataChanged
func callbackQMetaDataReaderControlMetaDataChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QMetaDataReaderControl::metaDataChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaDataChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QMetaDataReaderControl) MetaDataChanged() {
	defer qt.Recovering("QMetaDataReaderControl::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_MetaDataChanged(ptr.Pointer())
	}
}

func (ptr *QMetaDataReaderControl) ConnectMetaDataChanged2(f func(key string, value *core.QVariant)) {
	defer qt.Recovering("connect QMetaDataReaderControl::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_ConnectMetaDataChanged2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataChanged2", f)
	}
}

func (ptr *QMetaDataReaderControl) DisconnectMetaDataChanged2() {
	defer qt.Recovering("disconnect QMetaDataReaderControl::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_DisconnectMetaDataChanged2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataChanged2")
	}
}

//export callbackQMetaDataReaderControlMetaDataChanged2
func callbackQMetaDataReaderControlMetaDataChanged2(ptr unsafe.Pointer, ptrName *C.char, key *C.char, value unsafe.Pointer) {
	defer qt.Recovering("callback QMetaDataReaderControl::metaDataChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaDataChanged2"); signal != nil {
		signal.(func(string, *core.QVariant))(C.GoString(key), core.NewQVariantFromPointer(value))
	}

}

func (ptr *QMetaDataReaderControl) MetaDataChanged2(key string, value core.QVariant_ITF) {
	defer qt.Recovering("QMetaDataReaderControl::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_MetaDataChanged2(ptr.Pointer(), C.CString(key), core.PointerFromQVariant(value))
	}
}

func (ptr *QMetaDataReaderControl) DestroyQMetaDataReaderControl() {
	defer qt.Recovering("QMetaDataReaderControl::~QMetaDataReaderControl")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_DestroyQMetaDataReaderControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMetaDataReaderControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMetaDataReaderControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMetaDataReaderControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMetaDataReaderControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMetaDataReaderControlTimerEvent
func callbackQMetaDataReaderControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMetaDataReaderControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMetaDataReaderControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMetaDataReaderControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMetaDataReaderControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMetaDataReaderControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMetaDataReaderControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMetaDataReaderControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMetaDataReaderControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMetaDataReaderControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMetaDataReaderControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMetaDataReaderControlChildEvent
func callbackQMetaDataReaderControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMetaDataReaderControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMetaDataReaderControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMetaDataReaderControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMetaDataReaderControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMetaDataReaderControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMetaDataReaderControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMetaDataReaderControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMetaDataReaderControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMetaDataReaderControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMetaDataReaderControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMetaDataReaderControlCustomEvent
func callbackQMetaDataReaderControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMetaDataReaderControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMetaDataReaderControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMetaDataReaderControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMetaDataReaderControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMetaDataReaderControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMetaDataReaderControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QMetaDataWriterControl struct {
	QMediaControl
}

type QMetaDataWriterControl_ITF interface {
	QMediaControl_ITF
	QMetaDataWriterControl_PTR() *QMetaDataWriterControl
}

func PointerFromQMetaDataWriterControl(ptr QMetaDataWriterControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMetaDataWriterControl_PTR().Pointer()
	}
	return nil
}

func NewQMetaDataWriterControlFromPointer(ptr unsafe.Pointer) *QMetaDataWriterControl {
	var n = new(QMetaDataWriterControl)
	n.SetPointer(ptr)
	return n
}

func newQMetaDataWriterControlFromPointer(ptr unsafe.Pointer) *QMetaDataWriterControl {
	var n = NewQMetaDataWriterControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QMetaDataWriterControl_") {
		n.SetObjectName("QMetaDataWriterControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QMetaDataWriterControl) QMetaDataWriterControl_PTR() *QMetaDataWriterControl {
	return ptr
}

func (ptr *QMetaDataWriterControl) AvailableMetaData() []string {
	defer qt.Recovering("QMetaDataWriterControl::availableMetaData")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMetaDataWriterControl_AvailableMetaData(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QMetaDataWriterControl) IsMetaDataAvailable() bool {
	defer qt.Recovering("QMetaDataWriterControl::isMetaDataAvailable")

	if ptr.Pointer() != nil {
		return C.QMetaDataWriterControl_IsMetaDataAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMetaDataWriterControl) IsWritable() bool {
	defer qt.Recovering("QMetaDataWriterControl::isWritable")

	if ptr.Pointer() != nil {
		return C.QMetaDataWriterControl_IsWritable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMetaDataWriterControl) MetaData(key string) *core.QVariant {
	defer qt.Recovering("QMetaDataWriterControl::metaData")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QMetaDataWriterControl_MetaData(ptr.Pointer(), C.CString(key)))
	}
	return nil
}

func (ptr *QMetaDataWriterControl) ConnectMetaDataAvailableChanged(f func(available bool)) {
	defer qt.Recovering("connect QMetaDataWriterControl::metaDataAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_ConnectMetaDataAvailableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataAvailableChanged", f)
	}
}

func (ptr *QMetaDataWriterControl) DisconnectMetaDataAvailableChanged() {
	defer qt.Recovering("disconnect QMetaDataWriterControl::metaDataAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_DisconnectMetaDataAvailableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataAvailableChanged")
	}
}

//export callbackQMetaDataWriterControlMetaDataAvailableChanged
func callbackQMetaDataWriterControlMetaDataAvailableChanged(ptr unsafe.Pointer, ptrName *C.char, available C.int) {
	defer qt.Recovering("callback QMetaDataWriterControl::metaDataAvailableChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaDataAvailableChanged"); signal != nil {
		signal.(func(bool))(int(available) != 0)
	}

}

func (ptr *QMetaDataWriterControl) MetaDataAvailableChanged(available bool) {
	defer qt.Recovering("QMetaDataWriterControl::metaDataAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_MetaDataAvailableChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(available)))
	}
}

func (ptr *QMetaDataWriterControl) ConnectMetaDataChanged(f func()) {
	defer qt.Recovering("connect QMetaDataWriterControl::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_ConnectMetaDataChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataChanged", f)
	}
}

func (ptr *QMetaDataWriterControl) DisconnectMetaDataChanged() {
	defer qt.Recovering("disconnect QMetaDataWriterControl::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_DisconnectMetaDataChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataChanged")
	}
}

//export callbackQMetaDataWriterControlMetaDataChanged
func callbackQMetaDataWriterControlMetaDataChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QMetaDataWriterControl::metaDataChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaDataChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QMetaDataWriterControl) MetaDataChanged() {
	defer qt.Recovering("QMetaDataWriterControl::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_MetaDataChanged(ptr.Pointer())
	}
}

func (ptr *QMetaDataWriterControl) ConnectMetaDataChanged2(f func(key string, value *core.QVariant)) {
	defer qt.Recovering("connect QMetaDataWriterControl::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_ConnectMetaDataChanged2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataChanged2", f)
	}
}

func (ptr *QMetaDataWriterControl) DisconnectMetaDataChanged2() {
	defer qt.Recovering("disconnect QMetaDataWriterControl::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_DisconnectMetaDataChanged2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataChanged2")
	}
}

//export callbackQMetaDataWriterControlMetaDataChanged2
func callbackQMetaDataWriterControlMetaDataChanged2(ptr unsafe.Pointer, ptrName *C.char, key *C.char, value unsafe.Pointer) {
	defer qt.Recovering("callback QMetaDataWriterControl::metaDataChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaDataChanged2"); signal != nil {
		signal.(func(string, *core.QVariant))(C.GoString(key), core.NewQVariantFromPointer(value))
	}

}

func (ptr *QMetaDataWriterControl) MetaDataChanged2(key string, value core.QVariant_ITF) {
	defer qt.Recovering("QMetaDataWriterControl::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_MetaDataChanged2(ptr.Pointer(), C.CString(key), core.PointerFromQVariant(value))
	}
}

func (ptr *QMetaDataWriterControl) SetMetaData(key string, value core.QVariant_ITF) {
	defer qt.Recovering("QMetaDataWriterControl::setMetaData")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_SetMetaData(ptr.Pointer(), C.CString(key), core.PointerFromQVariant(value))
	}
}

func (ptr *QMetaDataWriterControl) ConnectWritableChanged(f func(writable bool)) {
	defer qt.Recovering("connect QMetaDataWriterControl::writableChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_ConnectWritableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "writableChanged", f)
	}
}

func (ptr *QMetaDataWriterControl) DisconnectWritableChanged() {
	defer qt.Recovering("disconnect QMetaDataWriterControl::writableChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_DisconnectWritableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "writableChanged")
	}
}

//export callbackQMetaDataWriterControlWritableChanged
func callbackQMetaDataWriterControlWritableChanged(ptr unsafe.Pointer, ptrName *C.char, writable C.int) {
	defer qt.Recovering("callback QMetaDataWriterControl::writableChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "writableChanged"); signal != nil {
		signal.(func(bool))(int(writable) != 0)
	}

}

func (ptr *QMetaDataWriterControl) WritableChanged(writable bool) {
	defer qt.Recovering("QMetaDataWriterControl::writableChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_WritableChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(writable)))
	}
}

func (ptr *QMetaDataWriterControl) DestroyQMetaDataWriterControl() {
	defer qt.Recovering("QMetaDataWriterControl::~QMetaDataWriterControl")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_DestroyQMetaDataWriterControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMetaDataWriterControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMetaDataWriterControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMetaDataWriterControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMetaDataWriterControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMetaDataWriterControlTimerEvent
func callbackQMetaDataWriterControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMetaDataWriterControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMetaDataWriterControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMetaDataWriterControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMetaDataWriterControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMetaDataWriterControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMetaDataWriterControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMetaDataWriterControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMetaDataWriterControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMetaDataWriterControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMetaDataWriterControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMetaDataWriterControlChildEvent
func callbackQMetaDataWriterControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMetaDataWriterControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMetaDataWriterControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMetaDataWriterControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMetaDataWriterControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMetaDataWriterControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMetaDataWriterControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMetaDataWriterControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMetaDataWriterControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMetaDataWriterControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMetaDataWriterControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMetaDataWriterControlCustomEvent
func callbackQMetaDataWriterControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMetaDataWriterControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMetaDataWriterControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMetaDataWriterControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMetaDataWriterControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMetaDataWriterControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMetaDataWriterControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//QMultimedia::AvailabilityStatus
type QMultimedia__AvailabilityStatus int64

const (
	QMultimedia__Available      = QMultimedia__AvailabilityStatus(0)
	QMultimedia__ServiceMissing = QMultimedia__AvailabilityStatus(1)
	QMultimedia__Busy           = QMultimedia__AvailabilityStatus(2)
	QMultimedia__ResourceError  = QMultimedia__AvailabilityStatus(3)
)

//QMultimedia::EncodingMode
type QMultimedia__EncodingMode int64

const (
	QMultimedia__ConstantQualityEncoding = QMultimedia__EncodingMode(0)
	QMultimedia__ConstantBitRateEncoding = QMultimedia__EncodingMode(1)
	QMultimedia__AverageBitRateEncoding  = QMultimedia__EncodingMode(2)
	QMultimedia__TwoPassEncoding         = QMultimedia__EncodingMode(3)
)

//QMultimedia::EncodingQuality
type QMultimedia__EncodingQuality int64

const (
	QMultimedia__VeryLowQuality  = QMultimedia__EncodingQuality(0)
	QMultimedia__LowQuality      = QMultimedia__EncodingQuality(1)
	QMultimedia__NormalQuality   = QMultimedia__EncodingQuality(2)
	QMultimedia__HighQuality     = QMultimedia__EncodingQuality(3)
	QMultimedia__VeryHighQuality = QMultimedia__EncodingQuality(4)
)

//QMultimedia::SupportEstimate
type QMultimedia__SupportEstimate int64

const (
	QMultimedia__NotSupported      = QMultimedia__SupportEstimate(0)
	QMultimedia__MaybeSupported    = QMultimedia__SupportEstimate(1)
	QMultimedia__ProbablySupported = QMultimedia__SupportEstimate(2)
	QMultimedia__PreferredService  = QMultimedia__SupportEstimate(3)
)

type QRadioData struct {
	core.QObject
	QMediaBindableInterface
}

type QRadioData_ITF interface {
	core.QObject_ITF
	QMediaBindableInterface_ITF
	QRadioData_PTR() *QRadioData
}

func (p *QRadioData) Pointer() unsafe.Pointer {
	return p.QObject_PTR().Pointer()
}

func (p *QRadioData) SetPointer(ptr unsafe.Pointer) {
	p.QObject_PTR().SetPointer(ptr)
	p.QMediaBindableInterface_PTR().SetPointer(ptr)
}

func PointerFromQRadioData(ptr QRadioData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRadioData_PTR().Pointer()
	}
	return nil
}

func NewQRadioDataFromPointer(ptr unsafe.Pointer) *QRadioData {
	var n = new(QRadioData)
	n.SetPointer(ptr)
	return n
}

func newQRadioDataFromPointer(ptr unsafe.Pointer) *QRadioData {
	var n = NewQRadioDataFromPointer(ptr)
	for len(n.ObjectName()) < len("QRadioData_") {
		n.SetObjectName("QRadioData_" + qt.Identifier())
	}
	return n
}

func (ptr *QRadioData) QRadioData_PTR() *QRadioData {
	return ptr
}

//QRadioData::Error
type QRadioData__Error int64

const (
	QRadioData__NoError         = QRadioData__Error(0)
	QRadioData__ResourceError   = QRadioData__Error(1)
	QRadioData__OpenError       = QRadioData__Error(2)
	QRadioData__OutOfRangeError = QRadioData__Error(3)
)

//QRadioData::ProgramType
type QRadioData__ProgramType int64

const (
	QRadioData__Undefined           = QRadioData__ProgramType(0)
	QRadioData__News                = QRadioData__ProgramType(1)
	QRadioData__CurrentAffairs      = QRadioData__ProgramType(2)
	QRadioData__Information         = QRadioData__ProgramType(3)
	QRadioData__Sport               = QRadioData__ProgramType(4)
	QRadioData__Education           = QRadioData__ProgramType(5)
	QRadioData__Drama               = QRadioData__ProgramType(6)
	QRadioData__Culture             = QRadioData__ProgramType(7)
	QRadioData__Science             = QRadioData__ProgramType(8)
	QRadioData__Varied              = QRadioData__ProgramType(9)
	QRadioData__PopMusic            = QRadioData__ProgramType(10)
	QRadioData__RockMusic           = QRadioData__ProgramType(11)
	QRadioData__EasyListening       = QRadioData__ProgramType(12)
	QRadioData__LightClassical      = QRadioData__ProgramType(13)
	QRadioData__SeriousClassical    = QRadioData__ProgramType(14)
	QRadioData__OtherMusic          = QRadioData__ProgramType(15)
	QRadioData__Weather             = QRadioData__ProgramType(16)
	QRadioData__Finance             = QRadioData__ProgramType(17)
	QRadioData__ChildrensProgrammes = QRadioData__ProgramType(18)
	QRadioData__SocialAffairs       = QRadioData__ProgramType(19)
	QRadioData__Religion            = QRadioData__ProgramType(20)
	QRadioData__PhoneIn             = QRadioData__ProgramType(21)
	QRadioData__Travel              = QRadioData__ProgramType(22)
	QRadioData__Leisure             = QRadioData__ProgramType(23)
	QRadioData__JazzMusic           = QRadioData__ProgramType(24)
	QRadioData__CountryMusic        = QRadioData__ProgramType(25)
	QRadioData__NationalMusic       = QRadioData__ProgramType(26)
	QRadioData__OldiesMusic         = QRadioData__ProgramType(27)
	QRadioData__FolkMusic           = QRadioData__ProgramType(28)
	QRadioData__Documentary         = QRadioData__ProgramType(29)
	QRadioData__AlarmTest           = QRadioData__ProgramType(30)
	QRadioData__Alarm               = QRadioData__ProgramType(31)
	QRadioData__Talk                = QRadioData__ProgramType(32)
	QRadioData__ClassicRock         = QRadioData__ProgramType(33)
	QRadioData__AdultHits           = QRadioData__ProgramType(34)
	QRadioData__SoftRock            = QRadioData__ProgramType(35)
	QRadioData__Top40               = QRadioData__ProgramType(36)
	QRadioData__Soft                = QRadioData__ProgramType(37)
	QRadioData__Nostalgia           = QRadioData__ProgramType(38)
	QRadioData__Classical           = QRadioData__ProgramType(39)
	QRadioData__RhythmAndBlues      = QRadioData__ProgramType(40)
	QRadioData__SoftRhythmAndBlues  = QRadioData__ProgramType(41)
	QRadioData__Language            = QRadioData__ProgramType(42)
	QRadioData__ReligiousMusic      = QRadioData__ProgramType(43)
	QRadioData__ReligiousTalk       = QRadioData__ProgramType(44)
	QRadioData__Personality         = QRadioData__ProgramType(45)
	QRadioData__Public              = QRadioData__ProgramType(46)
	QRadioData__College             = QRadioData__ProgramType(47)
)

func (ptr *QRadioData) IsAlternativeFrequenciesEnabled() bool {
	defer qt.Recovering("QRadioData::isAlternativeFrequenciesEnabled")

	if ptr.Pointer() != nil {
		return C.QRadioData_IsAlternativeFrequenciesEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRadioData) ProgramType() QRadioData__ProgramType {
	defer qt.Recovering("QRadioData::programType")

	if ptr.Pointer() != nil {
		return QRadioData__ProgramType(C.QRadioData_ProgramType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioData) ProgramTypeName() string {
	defer qt.Recovering("QRadioData::programTypeName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioData_ProgramTypeName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QRadioData) RadioText() string {
	defer qt.Recovering("QRadioData::radioText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioData_RadioText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QRadioData) SetAlternativeFrequenciesEnabled(enabled bool) {
	defer qt.Recovering("QRadioData::setAlternativeFrequenciesEnabled")

	if ptr.Pointer() != nil {
		C.QRadioData_SetAlternativeFrequenciesEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QRadioData) StationId() string {
	defer qt.Recovering("QRadioData::stationId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioData_StationId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QRadioData) StationName() string {
	defer qt.Recovering("QRadioData::stationName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioData_StationName(ptr.Pointer()))
	}
	return ""
}

func NewQRadioData(mediaObject QMediaObject_ITF, parent core.QObject_ITF) *QRadioData {
	defer qt.Recovering("QRadioData::QRadioData")

	return newQRadioDataFromPointer(C.QRadioData_NewQRadioData(PointerFromQMediaObject(mediaObject), core.PointerFromQObject(parent)))
}

func (ptr *QRadioData) ConnectAlternativeFrequenciesEnabledChanged(f func(enabled bool)) {
	defer qt.Recovering("connect QRadioData::alternativeFrequenciesEnabledChanged")

	if ptr.Pointer() != nil {
		C.QRadioData_ConnectAlternativeFrequenciesEnabledChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "alternativeFrequenciesEnabledChanged", f)
	}
}

func (ptr *QRadioData) DisconnectAlternativeFrequenciesEnabledChanged() {
	defer qt.Recovering("disconnect QRadioData::alternativeFrequenciesEnabledChanged")

	if ptr.Pointer() != nil {
		C.QRadioData_DisconnectAlternativeFrequenciesEnabledChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "alternativeFrequenciesEnabledChanged")
	}
}

//export callbackQRadioDataAlternativeFrequenciesEnabledChanged
func callbackQRadioDataAlternativeFrequenciesEnabledChanged(ptr unsafe.Pointer, ptrName *C.char, enabled C.int) {
	defer qt.Recovering("callback QRadioData::alternativeFrequenciesEnabledChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "alternativeFrequenciesEnabledChanged"); signal != nil {
		signal.(func(bool))(int(enabled) != 0)
	}

}

func (ptr *QRadioData) AlternativeFrequenciesEnabledChanged(enabled bool) {
	defer qt.Recovering("QRadioData::alternativeFrequenciesEnabledChanged")

	if ptr.Pointer() != nil {
		C.QRadioData_AlternativeFrequenciesEnabledChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QRadioData) Availability() QMultimedia__AvailabilityStatus {
	defer qt.Recovering("QRadioData::availability")

	if ptr.Pointer() != nil {
		return QMultimedia__AvailabilityStatus(C.QRadioData_Availability(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioData) ConnectError2(f func(error QRadioData__Error)) {
	defer qt.Recovering("connect QRadioData::error")

	if ptr.Pointer() != nil {
		C.QRadioData_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QRadioData) DisconnectError2() {
	defer qt.Recovering("disconnect QRadioData::error")

	if ptr.Pointer() != nil {
		C.QRadioData_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQRadioDataError2
func callbackQRadioDataError2(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QRadioData::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QRadioData__Error))(QRadioData__Error(error))
	}

}

func (ptr *QRadioData) Error2(error QRadioData__Error) {
	defer qt.Recovering("QRadioData::error")

	if ptr.Pointer() != nil {
		C.QRadioData_Error2(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QRadioData) Error() QRadioData__Error {
	defer qt.Recovering("QRadioData::error")

	if ptr.Pointer() != nil {
		return QRadioData__Error(C.QRadioData_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioData) ErrorString() string {
	defer qt.Recovering("QRadioData::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioData_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QRadioData) MediaObject() *QMediaObject {
	defer qt.Recovering("QRadioData::mediaObject")

	if ptr.Pointer() != nil {
		return NewQMediaObjectFromPointer(C.QRadioData_MediaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QRadioData) ConnectProgramTypeChanged(f func(programType QRadioData__ProgramType)) {
	defer qt.Recovering("connect QRadioData::programTypeChanged")

	if ptr.Pointer() != nil {
		C.QRadioData_ConnectProgramTypeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "programTypeChanged", f)
	}
}

func (ptr *QRadioData) DisconnectProgramTypeChanged() {
	defer qt.Recovering("disconnect QRadioData::programTypeChanged")

	if ptr.Pointer() != nil {
		C.QRadioData_DisconnectProgramTypeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "programTypeChanged")
	}
}

//export callbackQRadioDataProgramTypeChanged
func callbackQRadioDataProgramTypeChanged(ptr unsafe.Pointer, ptrName *C.char, programType C.int) {
	defer qt.Recovering("callback QRadioData::programTypeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "programTypeChanged"); signal != nil {
		signal.(func(QRadioData__ProgramType))(QRadioData__ProgramType(programType))
	}

}

func (ptr *QRadioData) ProgramTypeChanged(programType QRadioData__ProgramType) {
	defer qt.Recovering("QRadioData::programTypeChanged")

	if ptr.Pointer() != nil {
		C.QRadioData_ProgramTypeChanged(ptr.Pointer(), C.int(programType))
	}
}

func (ptr *QRadioData) ConnectProgramTypeNameChanged(f func(programTypeName string)) {
	defer qt.Recovering("connect QRadioData::programTypeNameChanged")

	if ptr.Pointer() != nil {
		C.QRadioData_ConnectProgramTypeNameChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "programTypeNameChanged", f)
	}
}

func (ptr *QRadioData) DisconnectProgramTypeNameChanged() {
	defer qt.Recovering("disconnect QRadioData::programTypeNameChanged")

	if ptr.Pointer() != nil {
		C.QRadioData_DisconnectProgramTypeNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "programTypeNameChanged")
	}
}

//export callbackQRadioDataProgramTypeNameChanged
func callbackQRadioDataProgramTypeNameChanged(ptr unsafe.Pointer, ptrName *C.char, programTypeName *C.char) {
	defer qt.Recovering("callback QRadioData::programTypeNameChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "programTypeNameChanged"); signal != nil {
		signal.(func(string))(C.GoString(programTypeName))
	}

}

func (ptr *QRadioData) ProgramTypeNameChanged(programTypeName string) {
	defer qt.Recovering("QRadioData::programTypeNameChanged")

	if ptr.Pointer() != nil {
		C.QRadioData_ProgramTypeNameChanged(ptr.Pointer(), C.CString(programTypeName))
	}
}

func (ptr *QRadioData) ConnectRadioTextChanged(f func(radioText string)) {
	defer qt.Recovering("connect QRadioData::radioTextChanged")

	if ptr.Pointer() != nil {
		C.QRadioData_ConnectRadioTextChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "radioTextChanged", f)
	}
}

func (ptr *QRadioData) DisconnectRadioTextChanged() {
	defer qt.Recovering("disconnect QRadioData::radioTextChanged")

	if ptr.Pointer() != nil {
		C.QRadioData_DisconnectRadioTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "radioTextChanged")
	}
}

//export callbackQRadioDataRadioTextChanged
func callbackQRadioDataRadioTextChanged(ptr unsafe.Pointer, ptrName *C.char, radioText *C.char) {
	defer qt.Recovering("callback QRadioData::radioTextChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "radioTextChanged"); signal != nil {
		signal.(func(string))(C.GoString(radioText))
	}

}

func (ptr *QRadioData) RadioTextChanged(radioText string) {
	defer qt.Recovering("QRadioData::radioTextChanged")

	if ptr.Pointer() != nil {
		C.QRadioData_RadioTextChanged(ptr.Pointer(), C.CString(radioText))
	}
}

func (ptr *QRadioData) SetMediaObject(mediaObject QMediaObject_ITF) bool {
	defer qt.Recovering("QRadioData::setMediaObject")

	if ptr.Pointer() != nil {
		return C.QRadioData_SetMediaObject(ptr.Pointer(), PointerFromQMediaObject(mediaObject)) != 0
	}
	return false
}

func (ptr *QRadioData) ConnectStationIdChanged(f func(stationId string)) {
	defer qt.Recovering("connect QRadioData::stationIdChanged")

	if ptr.Pointer() != nil {
		C.QRadioData_ConnectStationIdChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stationIdChanged", f)
	}
}

func (ptr *QRadioData) DisconnectStationIdChanged() {
	defer qt.Recovering("disconnect QRadioData::stationIdChanged")

	if ptr.Pointer() != nil {
		C.QRadioData_DisconnectStationIdChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stationIdChanged")
	}
}

//export callbackQRadioDataStationIdChanged
func callbackQRadioDataStationIdChanged(ptr unsafe.Pointer, ptrName *C.char, stationId *C.char) {
	defer qt.Recovering("callback QRadioData::stationIdChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stationIdChanged"); signal != nil {
		signal.(func(string))(C.GoString(stationId))
	}

}

func (ptr *QRadioData) StationIdChanged(stationId string) {
	defer qt.Recovering("QRadioData::stationIdChanged")

	if ptr.Pointer() != nil {
		C.QRadioData_StationIdChanged(ptr.Pointer(), C.CString(stationId))
	}
}

func (ptr *QRadioData) ConnectStationNameChanged(f func(stationName string)) {
	defer qt.Recovering("connect QRadioData::stationNameChanged")

	if ptr.Pointer() != nil {
		C.QRadioData_ConnectStationNameChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stationNameChanged", f)
	}
}

func (ptr *QRadioData) DisconnectStationNameChanged() {
	defer qt.Recovering("disconnect QRadioData::stationNameChanged")

	if ptr.Pointer() != nil {
		C.QRadioData_DisconnectStationNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stationNameChanged")
	}
}

//export callbackQRadioDataStationNameChanged
func callbackQRadioDataStationNameChanged(ptr unsafe.Pointer, ptrName *C.char, stationName *C.char) {
	defer qt.Recovering("callback QRadioData::stationNameChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stationNameChanged"); signal != nil {
		signal.(func(string))(C.GoString(stationName))
	}

}

func (ptr *QRadioData) StationNameChanged(stationName string) {
	defer qt.Recovering("QRadioData::stationNameChanged")

	if ptr.Pointer() != nil {
		C.QRadioData_StationNameChanged(ptr.Pointer(), C.CString(stationName))
	}
}

func (ptr *QRadioData) DestroyQRadioData() {
	defer qt.Recovering("QRadioData::~QRadioData")

	if ptr.Pointer() != nil {
		C.QRadioData_DestroyQRadioData(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QRadioData) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QRadioData::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QRadioData) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QRadioData::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQRadioDataTimerEvent
func callbackQRadioDataTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioData::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQRadioDataFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QRadioData) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QRadioData::timerEvent")

	if ptr.Pointer() != nil {
		C.QRadioData_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QRadioData) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QRadioData::timerEvent")

	if ptr.Pointer() != nil {
		C.QRadioData_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QRadioData) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QRadioData::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QRadioData) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QRadioData::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQRadioDataChildEvent
func callbackQRadioDataChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioData::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQRadioDataFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QRadioData) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QRadioData::childEvent")

	if ptr.Pointer() != nil {
		C.QRadioData_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QRadioData) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QRadioData::childEvent")

	if ptr.Pointer() != nil {
		C.QRadioData_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QRadioData) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QRadioData::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QRadioData) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QRadioData::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQRadioDataCustomEvent
func callbackQRadioDataCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioData::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQRadioDataFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QRadioData) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QRadioData::customEvent")

	if ptr.Pointer() != nil {
		C.QRadioData_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QRadioData) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QRadioData::customEvent")

	if ptr.Pointer() != nil {
		C.QRadioData_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QRadioDataControl struct {
	QMediaControl
}

type QRadioDataControl_ITF interface {
	QMediaControl_ITF
	QRadioDataControl_PTR() *QRadioDataControl
}

func PointerFromQRadioDataControl(ptr QRadioDataControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRadioDataControl_PTR().Pointer()
	}
	return nil
}

func NewQRadioDataControlFromPointer(ptr unsafe.Pointer) *QRadioDataControl {
	var n = new(QRadioDataControl)
	n.SetPointer(ptr)
	return n
}

func newQRadioDataControlFromPointer(ptr unsafe.Pointer) *QRadioDataControl {
	var n = NewQRadioDataControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QRadioDataControl_") {
		n.SetObjectName("QRadioDataControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QRadioDataControl) QRadioDataControl_PTR() *QRadioDataControl {
	return ptr
}

func (ptr *QRadioDataControl) ConnectAlternativeFrequenciesEnabledChanged(f func(enabled bool)) {
	defer qt.Recovering("connect QRadioDataControl::alternativeFrequenciesEnabledChanged")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_ConnectAlternativeFrequenciesEnabledChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "alternativeFrequenciesEnabledChanged", f)
	}
}

func (ptr *QRadioDataControl) DisconnectAlternativeFrequenciesEnabledChanged() {
	defer qt.Recovering("disconnect QRadioDataControl::alternativeFrequenciesEnabledChanged")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_DisconnectAlternativeFrequenciesEnabledChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "alternativeFrequenciesEnabledChanged")
	}
}

//export callbackQRadioDataControlAlternativeFrequenciesEnabledChanged
func callbackQRadioDataControlAlternativeFrequenciesEnabledChanged(ptr unsafe.Pointer, ptrName *C.char, enabled C.int) {
	defer qt.Recovering("callback QRadioDataControl::alternativeFrequenciesEnabledChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "alternativeFrequenciesEnabledChanged"); signal != nil {
		signal.(func(bool))(int(enabled) != 0)
	}

}

func (ptr *QRadioDataControl) AlternativeFrequenciesEnabledChanged(enabled bool) {
	defer qt.Recovering("QRadioDataControl::alternativeFrequenciesEnabledChanged")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_AlternativeFrequenciesEnabledChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QRadioDataControl) ConnectError2(f func(error QRadioData__Error)) {
	defer qt.Recovering("connect QRadioDataControl::error")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QRadioDataControl) DisconnectError2() {
	defer qt.Recovering("disconnect QRadioDataControl::error")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQRadioDataControlError2
func callbackQRadioDataControlError2(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QRadioDataControl::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QRadioData__Error))(QRadioData__Error(error))
	}

}

func (ptr *QRadioDataControl) Error2(error QRadioData__Error) {
	defer qt.Recovering("QRadioDataControl::error")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_Error2(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QRadioDataControl) Error() QRadioData__Error {
	defer qt.Recovering("QRadioDataControl::error")

	if ptr.Pointer() != nil {
		return QRadioData__Error(C.QRadioDataControl_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioDataControl) ErrorString() string {
	defer qt.Recovering("QRadioDataControl::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioDataControl_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QRadioDataControl) IsAlternativeFrequenciesEnabled() bool {
	defer qt.Recovering("QRadioDataControl::isAlternativeFrequenciesEnabled")

	if ptr.Pointer() != nil {
		return C.QRadioDataControl_IsAlternativeFrequenciesEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRadioDataControl) ProgramType() QRadioData__ProgramType {
	defer qt.Recovering("QRadioDataControl::programType")

	if ptr.Pointer() != nil {
		return QRadioData__ProgramType(C.QRadioDataControl_ProgramType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioDataControl) ConnectProgramTypeChanged(f func(programType QRadioData__ProgramType)) {
	defer qt.Recovering("connect QRadioDataControl::programTypeChanged")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_ConnectProgramTypeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "programTypeChanged", f)
	}
}

func (ptr *QRadioDataControl) DisconnectProgramTypeChanged() {
	defer qt.Recovering("disconnect QRadioDataControl::programTypeChanged")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_DisconnectProgramTypeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "programTypeChanged")
	}
}

//export callbackQRadioDataControlProgramTypeChanged
func callbackQRadioDataControlProgramTypeChanged(ptr unsafe.Pointer, ptrName *C.char, programType C.int) {
	defer qt.Recovering("callback QRadioDataControl::programTypeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "programTypeChanged"); signal != nil {
		signal.(func(QRadioData__ProgramType))(QRadioData__ProgramType(programType))
	}

}

func (ptr *QRadioDataControl) ProgramTypeChanged(programType QRadioData__ProgramType) {
	defer qt.Recovering("QRadioDataControl::programTypeChanged")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_ProgramTypeChanged(ptr.Pointer(), C.int(programType))
	}
}

func (ptr *QRadioDataControl) ProgramTypeName() string {
	defer qt.Recovering("QRadioDataControl::programTypeName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioDataControl_ProgramTypeName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QRadioDataControl) ConnectProgramTypeNameChanged(f func(programTypeName string)) {
	defer qt.Recovering("connect QRadioDataControl::programTypeNameChanged")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_ConnectProgramTypeNameChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "programTypeNameChanged", f)
	}
}

func (ptr *QRadioDataControl) DisconnectProgramTypeNameChanged() {
	defer qt.Recovering("disconnect QRadioDataControl::programTypeNameChanged")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_DisconnectProgramTypeNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "programTypeNameChanged")
	}
}

//export callbackQRadioDataControlProgramTypeNameChanged
func callbackQRadioDataControlProgramTypeNameChanged(ptr unsafe.Pointer, ptrName *C.char, programTypeName *C.char) {
	defer qt.Recovering("callback QRadioDataControl::programTypeNameChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "programTypeNameChanged"); signal != nil {
		signal.(func(string))(C.GoString(programTypeName))
	}

}

func (ptr *QRadioDataControl) ProgramTypeNameChanged(programTypeName string) {
	defer qt.Recovering("QRadioDataControl::programTypeNameChanged")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_ProgramTypeNameChanged(ptr.Pointer(), C.CString(programTypeName))
	}
}

func (ptr *QRadioDataControl) RadioText() string {
	defer qt.Recovering("QRadioDataControl::radioText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioDataControl_RadioText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QRadioDataControl) ConnectRadioTextChanged(f func(radioText string)) {
	defer qt.Recovering("connect QRadioDataControl::radioTextChanged")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_ConnectRadioTextChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "radioTextChanged", f)
	}
}

func (ptr *QRadioDataControl) DisconnectRadioTextChanged() {
	defer qt.Recovering("disconnect QRadioDataControl::radioTextChanged")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_DisconnectRadioTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "radioTextChanged")
	}
}

//export callbackQRadioDataControlRadioTextChanged
func callbackQRadioDataControlRadioTextChanged(ptr unsafe.Pointer, ptrName *C.char, radioText *C.char) {
	defer qt.Recovering("callback QRadioDataControl::radioTextChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "radioTextChanged"); signal != nil {
		signal.(func(string))(C.GoString(radioText))
	}

}

func (ptr *QRadioDataControl) RadioTextChanged(radioText string) {
	defer qt.Recovering("QRadioDataControl::radioTextChanged")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_RadioTextChanged(ptr.Pointer(), C.CString(radioText))
	}
}

func (ptr *QRadioDataControl) SetAlternativeFrequenciesEnabled(enabled bool) {
	defer qt.Recovering("QRadioDataControl::setAlternativeFrequenciesEnabled")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_SetAlternativeFrequenciesEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QRadioDataControl) StationId() string {
	defer qt.Recovering("QRadioDataControl::stationId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioDataControl_StationId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QRadioDataControl) ConnectStationIdChanged(f func(stationId string)) {
	defer qt.Recovering("connect QRadioDataControl::stationIdChanged")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_ConnectStationIdChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stationIdChanged", f)
	}
}

func (ptr *QRadioDataControl) DisconnectStationIdChanged() {
	defer qt.Recovering("disconnect QRadioDataControl::stationIdChanged")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_DisconnectStationIdChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stationIdChanged")
	}
}

//export callbackQRadioDataControlStationIdChanged
func callbackQRadioDataControlStationIdChanged(ptr unsafe.Pointer, ptrName *C.char, stationId *C.char) {
	defer qt.Recovering("callback QRadioDataControl::stationIdChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stationIdChanged"); signal != nil {
		signal.(func(string))(C.GoString(stationId))
	}

}

func (ptr *QRadioDataControl) StationIdChanged(stationId string) {
	defer qt.Recovering("QRadioDataControl::stationIdChanged")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_StationIdChanged(ptr.Pointer(), C.CString(stationId))
	}
}

func (ptr *QRadioDataControl) StationName() string {
	defer qt.Recovering("QRadioDataControl::stationName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioDataControl_StationName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QRadioDataControl) ConnectStationNameChanged(f func(stationName string)) {
	defer qt.Recovering("connect QRadioDataControl::stationNameChanged")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_ConnectStationNameChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stationNameChanged", f)
	}
}

func (ptr *QRadioDataControl) DisconnectStationNameChanged() {
	defer qt.Recovering("disconnect QRadioDataControl::stationNameChanged")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_DisconnectStationNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stationNameChanged")
	}
}

//export callbackQRadioDataControlStationNameChanged
func callbackQRadioDataControlStationNameChanged(ptr unsafe.Pointer, ptrName *C.char, stationName *C.char) {
	defer qt.Recovering("callback QRadioDataControl::stationNameChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stationNameChanged"); signal != nil {
		signal.(func(string))(C.GoString(stationName))
	}

}

func (ptr *QRadioDataControl) StationNameChanged(stationName string) {
	defer qt.Recovering("QRadioDataControl::stationNameChanged")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_StationNameChanged(ptr.Pointer(), C.CString(stationName))
	}
}

func (ptr *QRadioDataControl) DestroyQRadioDataControl() {
	defer qt.Recovering("QRadioDataControl::~QRadioDataControl")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_DestroyQRadioDataControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QRadioDataControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QRadioDataControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QRadioDataControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QRadioDataControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQRadioDataControlTimerEvent
func callbackQRadioDataControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioDataControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQRadioDataControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QRadioDataControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QRadioDataControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QRadioDataControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QRadioDataControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QRadioDataControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QRadioDataControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QRadioDataControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QRadioDataControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQRadioDataControlChildEvent
func callbackQRadioDataControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioDataControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQRadioDataControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QRadioDataControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QRadioDataControl::childEvent")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QRadioDataControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QRadioDataControl::childEvent")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QRadioDataControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QRadioDataControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QRadioDataControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QRadioDataControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQRadioDataControlCustomEvent
func callbackQRadioDataControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioDataControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQRadioDataControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QRadioDataControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QRadioDataControl::customEvent")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QRadioDataControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QRadioDataControl::customEvent")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QRadioTuner struct {
	QMediaObject
}

type QRadioTuner_ITF interface {
	QMediaObject_ITF
	QRadioTuner_PTR() *QRadioTuner
}

func PointerFromQRadioTuner(ptr QRadioTuner_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRadioTuner_PTR().Pointer()
	}
	return nil
}

func NewQRadioTunerFromPointer(ptr unsafe.Pointer) *QRadioTuner {
	var n = new(QRadioTuner)
	n.SetPointer(ptr)
	return n
}

func newQRadioTunerFromPointer(ptr unsafe.Pointer) *QRadioTuner {
	var n = NewQRadioTunerFromPointer(ptr)
	for len(n.ObjectName()) < len("QRadioTuner_") {
		n.SetObjectName("QRadioTuner_" + qt.Identifier())
	}
	return n
}

func (ptr *QRadioTuner) QRadioTuner_PTR() *QRadioTuner {
	return ptr
}

//QRadioTuner::Band
type QRadioTuner__Band int64

const (
	QRadioTuner__AM  = QRadioTuner__Band(0)
	QRadioTuner__FM  = QRadioTuner__Band(1)
	QRadioTuner__SW  = QRadioTuner__Band(2)
	QRadioTuner__LW  = QRadioTuner__Band(3)
	QRadioTuner__FM2 = QRadioTuner__Band(4)
)

//QRadioTuner::Error
type QRadioTuner__Error int64

const (
	QRadioTuner__NoError         = QRadioTuner__Error(0)
	QRadioTuner__ResourceError   = QRadioTuner__Error(1)
	QRadioTuner__OpenError       = QRadioTuner__Error(2)
	QRadioTuner__OutOfRangeError = QRadioTuner__Error(3)
)

//QRadioTuner::SearchMode
type QRadioTuner__SearchMode int64

const (
	QRadioTuner__SearchFast         = QRadioTuner__SearchMode(0)
	QRadioTuner__SearchGetStationId = QRadioTuner__SearchMode(1)
)

//QRadioTuner::State
type QRadioTuner__State int64

const (
	QRadioTuner__ActiveState  = QRadioTuner__State(0)
	QRadioTuner__StoppedState = QRadioTuner__State(1)
)

//QRadioTuner::StereoMode
type QRadioTuner__StereoMode int64

const (
	QRadioTuner__ForceStereo = QRadioTuner__StereoMode(0)
	QRadioTuner__ForceMono   = QRadioTuner__StereoMode(1)
	QRadioTuner__Auto        = QRadioTuner__StereoMode(2)
)

func (ptr *QRadioTuner) Band() QRadioTuner__Band {
	defer qt.Recovering("QRadioTuner::band")

	if ptr.Pointer() != nil {
		return QRadioTuner__Band(C.QRadioTuner_Band(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioTuner) Frequency() int {
	defer qt.Recovering("QRadioTuner::frequency")

	if ptr.Pointer() != nil {
		return int(C.QRadioTuner_Frequency(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioTuner) IsAntennaConnected() bool {
	defer qt.Recovering("QRadioTuner::isAntennaConnected")

	if ptr.Pointer() != nil {
		return C.QRadioTuner_IsAntennaConnected(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRadioTuner) IsMuted() bool {
	defer qt.Recovering("QRadioTuner::isMuted")

	if ptr.Pointer() != nil {
		return C.QRadioTuner_IsMuted(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRadioTuner) IsSearching() bool {
	defer qt.Recovering("QRadioTuner::isSearching")

	if ptr.Pointer() != nil {
		return C.QRadioTuner_IsSearching(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRadioTuner) IsStereo() bool {
	defer qt.Recovering("QRadioTuner::isStereo")

	if ptr.Pointer() != nil {
		return C.QRadioTuner_IsStereo(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRadioTuner) RadioData() *QRadioData {
	defer qt.Recovering("QRadioTuner::radioData")

	if ptr.Pointer() != nil {
		return NewQRadioDataFromPointer(C.QRadioTuner_RadioData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QRadioTuner) SetMuted(muted bool) {
	defer qt.Recovering("QRadioTuner::setMuted")

	if ptr.Pointer() != nil {
		C.QRadioTuner_SetMuted(ptr.Pointer(), C.int(qt.GoBoolToInt(muted)))
	}
}

func (ptr *QRadioTuner) SetStereoMode(mode QRadioTuner__StereoMode) {
	defer qt.Recovering("QRadioTuner::setStereoMode")

	if ptr.Pointer() != nil {
		C.QRadioTuner_SetStereoMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QRadioTuner) SetVolume(volume int) {
	defer qt.Recovering("QRadioTuner::setVolume")

	if ptr.Pointer() != nil {
		C.QRadioTuner_SetVolume(ptr.Pointer(), C.int(volume))
	}
}

func (ptr *QRadioTuner) SignalStrength() int {
	defer qt.Recovering("QRadioTuner::signalStrength")

	if ptr.Pointer() != nil {
		return int(C.QRadioTuner_SignalStrength(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioTuner) State() QRadioTuner__State {
	defer qt.Recovering("QRadioTuner::state")

	if ptr.Pointer() != nil {
		return QRadioTuner__State(C.QRadioTuner_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioTuner) StereoMode() QRadioTuner__StereoMode {
	defer qt.Recovering("QRadioTuner::stereoMode")

	if ptr.Pointer() != nil {
		return QRadioTuner__StereoMode(C.QRadioTuner_StereoMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioTuner) Volume() int {
	defer qt.Recovering("QRadioTuner::volume")

	if ptr.Pointer() != nil {
		return int(C.QRadioTuner_Volume(ptr.Pointer()))
	}
	return 0
}

func NewQRadioTuner(parent core.QObject_ITF) *QRadioTuner {
	defer qt.Recovering("QRadioTuner::QRadioTuner")

	return newQRadioTunerFromPointer(C.QRadioTuner_NewQRadioTuner(core.PointerFromQObject(parent)))
}

func (ptr *QRadioTuner) ConnectAntennaConnectedChanged(f func(connectionStatus bool)) {
	defer qt.Recovering("connect QRadioTuner::antennaConnectedChanged")

	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectAntennaConnectedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "antennaConnectedChanged", f)
	}
}

func (ptr *QRadioTuner) DisconnectAntennaConnectedChanged() {
	defer qt.Recovering("disconnect QRadioTuner::antennaConnectedChanged")

	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectAntennaConnectedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "antennaConnectedChanged")
	}
}

//export callbackQRadioTunerAntennaConnectedChanged
func callbackQRadioTunerAntennaConnectedChanged(ptr unsafe.Pointer, ptrName *C.char, connectionStatus C.int) {
	defer qt.Recovering("callback QRadioTuner::antennaConnectedChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "antennaConnectedChanged"); signal != nil {
		signal.(func(bool))(int(connectionStatus) != 0)
	}

}

func (ptr *QRadioTuner) AntennaConnectedChanged(connectionStatus bool) {
	defer qt.Recovering("QRadioTuner::antennaConnectedChanged")

	if ptr.Pointer() != nil {
		C.QRadioTuner_AntennaConnectedChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(connectionStatus)))
	}
}

func (ptr *QRadioTuner) Availability() QMultimedia__AvailabilityStatus {
	defer qt.Recovering("QRadioTuner::availability")

	if ptr.Pointer() != nil {
		return QMultimedia__AvailabilityStatus(C.QRadioTuner_Availability(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioTuner) ConnectBandChanged(f func(band QRadioTuner__Band)) {
	defer qt.Recovering("connect QRadioTuner::bandChanged")

	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectBandChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bandChanged", f)
	}
}

func (ptr *QRadioTuner) DisconnectBandChanged() {
	defer qt.Recovering("disconnect QRadioTuner::bandChanged")

	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectBandChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bandChanged")
	}
}

//export callbackQRadioTunerBandChanged
func callbackQRadioTunerBandChanged(ptr unsafe.Pointer, ptrName *C.char, band C.int) {
	defer qt.Recovering("callback QRadioTuner::bandChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "bandChanged"); signal != nil {
		signal.(func(QRadioTuner__Band))(QRadioTuner__Band(band))
	}

}

func (ptr *QRadioTuner) BandChanged(band QRadioTuner__Band) {
	defer qt.Recovering("QRadioTuner::bandChanged")

	if ptr.Pointer() != nil {
		C.QRadioTuner_BandChanged(ptr.Pointer(), C.int(band))
	}
}

func (ptr *QRadioTuner) CancelSearch() {
	defer qt.Recovering("QRadioTuner::cancelSearch")

	if ptr.Pointer() != nil {
		C.QRadioTuner_CancelSearch(ptr.Pointer())
	}
}

func (ptr *QRadioTuner) ConnectError2(f func(error QRadioTuner__Error)) {
	defer qt.Recovering("connect QRadioTuner::error")

	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QRadioTuner) DisconnectError2() {
	defer qt.Recovering("disconnect QRadioTuner::error")

	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQRadioTunerError2
func callbackQRadioTunerError2(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QRadioTuner::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QRadioTuner__Error))(QRadioTuner__Error(error))
	}

}

func (ptr *QRadioTuner) Error2(error QRadioTuner__Error) {
	defer qt.Recovering("QRadioTuner::error")

	if ptr.Pointer() != nil {
		C.QRadioTuner_Error2(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QRadioTuner) Error() QRadioTuner__Error {
	defer qt.Recovering("QRadioTuner::error")

	if ptr.Pointer() != nil {
		return QRadioTuner__Error(C.QRadioTuner_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioTuner) ErrorString() string {
	defer qt.Recovering("QRadioTuner::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioTuner_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QRadioTuner) ConnectFrequencyChanged(f func(frequency int)) {
	defer qt.Recovering("connect QRadioTuner::frequencyChanged")

	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectFrequencyChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "frequencyChanged", f)
	}
}

func (ptr *QRadioTuner) DisconnectFrequencyChanged() {
	defer qt.Recovering("disconnect QRadioTuner::frequencyChanged")

	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectFrequencyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "frequencyChanged")
	}
}

//export callbackQRadioTunerFrequencyChanged
func callbackQRadioTunerFrequencyChanged(ptr unsafe.Pointer, ptrName *C.char, frequency C.int) {
	defer qt.Recovering("callback QRadioTuner::frequencyChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "frequencyChanged"); signal != nil {
		signal.(func(int))(int(frequency))
	}

}

func (ptr *QRadioTuner) FrequencyChanged(frequency int) {
	defer qt.Recovering("QRadioTuner::frequencyChanged")

	if ptr.Pointer() != nil {
		C.QRadioTuner_FrequencyChanged(ptr.Pointer(), C.int(frequency))
	}
}

func (ptr *QRadioTuner) FrequencyStep(band QRadioTuner__Band) int {
	defer qt.Recovering("QRadioTuner::frequencyStep")

	if ptr.Pointer() != nil {
		return int(C.QRadioTuner_FrequencyStep(ptr.Pointer(), C.int(band)))
	}
	return 0
}

func (ptr *QRadioTuner) IsBandSupported(band QRadioTuner__Band) bool {
	defer qt.Recovering("QRadioTuner::isBandSupported")

	if ptr.Pointer() != nil {
		return C.QRadioTuner_IsBandSupported(ptr.Pointer(), C.int(band)) != 0
	}
	return false
}

func (ptr *QRadioTuner) ConnectMutedChanged(f func(muted bool)) {
	defer qt.Recovering("connect QRadioTuner::mutedChanged")

	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectMutedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mutedChanged", f)
	}
}

func (ptr *QRadioTuner) DisconnectMutedChanged() {
	defer qt.Recovering("disconnect QRadioTuner::mutedChanged")

	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectMutedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mutedChanged")
	}
}

//export callbackQRadioTunerMutedChanged
func callbackQRadioTunerMutedChanged(ptr unsafe.Pointer, ptrName *C.char, muted C.int) {
	defer qt.Recovering("callback QRadioTuner::mutedChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "mutedChanged"); signal != nil {
		signal.(func(bool))(int(muted) != 0)
	}

}

func (ptr *QRadioTuner) MutedChanged(muted bool) {
	defer qt.Recovering("QRadioTuner::mutedChanged")

	if ptr.Pointer() != nil {
		C.QRadioTuner_MutedChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(muted)))
	}
}

func (ptr *QRadioTuner) SearchAllStations(searchMode QRadioTuner__SearchMode) {
	defer qt.Recovering("QRadioTuner::searchAllStations")

	if ptr.Pointer() != nil {
		C.QRadioTuner_SearchAllStations(ptr.Pointer(), C.int(searchMode))
	}
}

func (ptr *QRadioTuner) SearchBackward() {
	defer qt.Recovering("QRadioTuner::searchBackward")

	if ptr.Pointer() != nil {
		C.QRadioTuner_SearchBackward(ptr.Pointer())
	}
}

func (ptr *QRadioTuner) SearchForward() {
	defer qt.Recovering("QRadioTuner::searchForward")

	if ptr.Pointer() != nil {
		C.QRadioTuner_SearchForward(ptr.Pointer())
	}
}

func (ptr *QRadioTuner) ConnectSearchingChanged(f func(searching bool)) {
	defer qt.Recovering("connect QRadioTuner::searchingChanged")

	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectSearchingChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "searchingChanged", f)
	}
}

func (ptr *QRadioTuner) DisconnectSearchingChanged() {
	defer qt.Recovering("disconnect QRadioTuner::searchingChanged")

	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectSearchingChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "searchingChanged")
	}
}

//export callbackQRadioTunerSearchingChanged
func callbackQRadioTunerSearchingChanged(ptr unsafe.Pointer, ptrName *C.char, searching C.int) {
	defer qt.Recovering("callback QRadioTuner::searchingChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "searchingChanged"); signal != nil {
		signal.(func(bool))(int(searching) != 0)
	}

}

func (ptr *QRadioTuner) SearchingChanged(searching bool) {
	defer qt.Recovering("QRadioTuner::searchingChanged")

	if ptr.Pointer() != nil {
		C.QRadioTuner_SearchingChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(searching)))
	}
}

func (ptr *QRadioTuner) SetBand(band QRadioTuner__Band) {
	defer qt.Recovering("QRadioTuner::setBand")

	if ptr.Pointer() != nil {
		C.QRadioTuner_SetBand(ptr.Pointer(), C.int(band))
	}
}

func (ptr *QRadioTuner) SetFrequency(frequency int) {
	defer qt.Recovering("QRadioTuner::setFrequency")

	if ptr.Pointer() != nil {
		C.QRadioTuner_SetFrequency(ptr.Pointer(), C.int(frequency))
	}
}

func (ptr *QRadioTuner) ConnectSignalStrengthChanged(f func(strength int)) {
	defer qt.Recovering("connect QRadioTuner::signalStrengthChanged")

	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectSignalStrengthChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "signalStrengthChanged", f)
	}
}

func (ptr *QRadioTuner) DisconnectSignalStrengthChanged() {
	defer qt.Recovering("disconnect QRadioTuner::signalStrengthChanged")

	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectSignalStrengthChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "signalStrengthChanged")
	}
}

//export callbackQRadioTunerSignalStrengthChanged
func callbackQRadioTunerSignalStrengthChanged(ptr unsafe.Pointer, ptrName *C.char, strength C.int) {
	defer qt.Recovering("callback QRadioTuner::signalStrengthChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "signalStrengthChanged"); signal != nil {
		signal.(func(int))(int(strength))
	}

}

func (ptr *QRadioTuner) SignalStrengthChanged(strength int) {
	defer qt.Recovering("QRadioTuner::signalStrengthChanged")

	if ptr.Pointer() != nil {
		C.QRadioTuner_SignalStrengthChanged(ptr.Pointer(), C.int(strength))
	}
}

func (ptr *QRadioTuner) Start() {
	defer qt.Recovering("QRadioTuner::start")

	if ptr.Pointer() != nil {
		C.QRadioTuner_Start(ptr.Pointer())
	}
}

func (ptr *QRadioTuner) ConnectStateChanged(f func(state QRadioTuner__State)) {
	defer qt.Recovering("connect QRadioTuner::stateChanged")

	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QRadioTuner) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QRadioTuner::stateChanged")

	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQRadioTunerStateChanged
func callbackQRadioTunerStateChanged(ptr unsafe.Pointer, ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QRadioTuner::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QRadioTuner__State))(QRadioTuner__State(state))
	}

}

func (ptr *QRadioTuner) StateChanged(state QRadioTuner__State) {
	defer qt.Recovering("QRadioTuner::stateChanged")

	if ptr.Pointer() != nil {
		C.QRadioTuner_StateChanged(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QRadioTuner) ConnectStationFound(f func(frequency int, stationId string)) {
	defer qt.Recovering("connect QRadioTuner::stationFound")

	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectStationFound(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stationFound", f)
	}
}

func (ptr *QRadioTuner) DisconnectStationFound() {
	defer qt.Recovering("disconnect QRadioTuner::stationFound")

	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectStationFound(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stationFound")
	}
}

//export callbackQRadioTunerStationFound
func callbackQRadioTunerStationFound(ptr unsafe.Pointer, ptrName *C.char, frequency C.int, stationId *C.char) {
	defer qt.Recovering("callback QRadioTuner::stationFound")

	if signal := qt.GetSignal(C.GoString(ptrName), "stationFound"); signal != nil {
		signal.(func(int, string))(int(frequency), C.GoString(stationId))
	}

}

func (ptr *QRadioTuner) StationFound(frequency int, stationId string) {
	defer qt.Recovering("QRadioTuner::stationFound")

	if ptr.Pointer() != nil {
		C.QRadioTuner_StationFound(ptr.Pointer(), C.int(frequency), C.CString(stationId))
	}
}

func (ptr *QRadioTuner) ConnectStereoStatusChanged(f func(stereo bool)) {
	defer qt.Recovering("connect QRadioTuner::stereoStatusChanged")

	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectStereoStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stereoStatusChanged", f)
	}
}

func (ptr *QRadioTuner) DisconnectStereoStatusChanged() {
	defer qt.Recovering("disconnect QRadioTuner::stereoStatusChanged")

	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectStereoStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stereoStatusChanged")
	}
}

//export callbackQRadioTunerStereoStatusChanged
func callbackQRadioTunerStereoStatusChanged(ptr unsafe.Pointer, ptrName *C.char, stereo C.int) {
	defer qt.Recovering("callback QRadioTuner::stereoStatusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stereoStatusChanged"); signal != nil {
		signal.(func(bool))(int(stereo) != 0)
	}

}

func (ptr *QRadioTuner) StereoStatusChanged(stereo bool) {
	defer qt.Recovering("QRadioTuner::stereoStatusChanged")

	if ptr.Pointer() != nil {
		C.QRadioTuner_StereoStatusChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(stereo)))
	}
}

func (ptr *QRadioTuner) Stop() {
	defer qt.Recovering("QRadioTuner::stop")

	if ptr.Pointer() != nil {
		C.QRadioTuner_Stop(ptr.Pointer())
	}
}

func (ptr *QRadioTuner) ConnectVolumeChanged(f func(volume int)) {
	defer qt.Recovering("connect QRadioTuner::volumeChanged")

	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectVolumeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "volumeChanged", f)
	}
}

func (ptr *QRadioTuner) DisconnectVolumeChanged() {
	defer qt.Recovering("disconnect QRadioTuner::volumeChanged")

	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectVolumeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "volumeChanged")
	}
}

//export callbackQRadioTunerVolumeChanged
func callbackQRadioTunerVolumeChanged(ptr unsafe.Pointer, ptrName *C.char, volume C.int) {
	defer qt.Recovering("callback QRadioTuner::volumeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "volumeChanged"); signal != nil {
		signal.(func(int))(int(volume))
	}

}

func (ptr *QRadioTuner) VolumeChanged(volume int) {
	defer qt.Recovering("QRadioTuner::volumeChanged")

	if ptr.Pointer() != nil {
		C.QRadioTuner_VolumeChanged(ptr.Pointer(), C.int(volume))
	}
}

func (ptr *QRadioTuner) DestroyQRadioTuner() {
	defer qt.Recovering("QRadioTuner::~QRadioTuner")

	if ptr.Pointer() != nil {
		C.QRadioTuner_DestroyQRadioTuner(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QRadioTuner) ConnectUnbind(f func(object *core.QObject)) {
	defer qt.Recovering("connect QRadioTuner::unbind")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "unbind", f)
	}
}

func (ptr *QRadioTuner) DisconnectUnbind() {
	defer qt.Recovering("disconnect QRadioTuner::unbind")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "unbind")
	}
}

//export callbackQRadioTunerUnbind
func callbackQRadioTunerUnbind(ptr unsafe.Pointer, ptrName *C.char, object unsafe.Pointer) {
	defer qt.Recovering("callback QRadioTuner::unbind")

	if signal := qt.GetSignal(C.GoString(ptrName), "unbind"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(object))
	} else {
		NewQRadioTunerFromPointer(ptr).UnbindDefault(core.NewQObjectFromPointer(object))
	}
}

func (ptr *QRadioTuner) Unbind(object core.QObject_ITF) {
	defer qt.Recovering("QRadioTuner::unbind")

	if ptr.Pointer() != nil {
		C.QRadioTuner_Unbind(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QRadioTuner) UnbindDefault(object core.QObject_ITF) {
	defer qt.Recovering("QRadioTuner::unbind")

	if ptr.Pointer() != nil {
		C.QRadioTuner_UnbindDefault(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QRadioTuner) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QRadioTuner::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QRadioTuner) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QRadioTuner::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQRadioTunerTimerEvent
func callbackQRadioTunerTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioTuner::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQRadioTunerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QRadioTuner) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QRadioTuner::timerEvent")

	if ptr.Pointer() != nil {
		C.QRadioTuner_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QRadioTuner) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QRadioTuner::timerEvent")

	if ptr.Pointer() != nil {
		C.QRadioTuner_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QRadioTuner) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QRadioTuner::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QRadioTuner) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QRadioTuner::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQRadioTunerChildEvent
func callbackQRadioTunerChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioTuner::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQRadioTunerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QRadioTuner) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QRadioTuner::childEvent")

	if ptr.Pointer() != nil {
		C.QRadioTuner_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QRadioTuner) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QRadioTuner::childEvent")

	if ptr.Pointer() != nil {
		C.QRadioTuner_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QRadioTuner) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QRadioTuner::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QRadioTuner) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QRadioTuner::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQRadioTunerCustomEvent
func callbackQRadioTunerCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioTuner::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQRadioTunerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QRadioTuner) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QRadioTuner::customEvent")

	if ptr.Pointer() != nil {
		C.QRadioTuner_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QRadioTuner) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QRadioTuner::customEvent")

	if ptr.Pointer() != nil {
		C.QRadioTuner_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QRadioTunerControl struct {
	QMediaControl
}

type QRadioTunerControl_ITF interface {
	QMediaControl_ITF
	QRadioTunerControl_PTR() *QRadioTunerControl
}

func PointerFromQRadioTunerControl(ptr QRadioTunerControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRadioTunerControl_PTR().Pointer()
	}
	return nil
}

func NewQRadioTunerControlFromPointer(ptr unsafe.Pointer) *QRadioTunerControl {
	var n = new(QRadioTunerControl)
	n.SetPointer(ptr)
	return n
}

func newQRadioTunerControlFromPointer(ptr unsafe.Pointer) *QRadioTunerControl {
	var n = NewQRadioTunerControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QRadioTunerControl_") {
		n.SetObjectName("QRadioTunerControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QRadioTunerControl) QRadioTunerControl_PTR() *QRadioTunerControl {
	return ptr
}

func (ptr *QRadioTunerControl) ConnectAntennaConnectedChanged(f func(connectionStatus bool)) {
	defer qt.Recovering("connect QRadioTunerControl::antennaConnectedChanged")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectAntennaConnectedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "antennaConnectedChanged", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectAntennaConnectedChanged() {
	defer qt.Recovering("disconnect QRadioTunerControl::antennaConnectedChanged")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectAntennaConnectedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "antennaConnectedChanged")
	}
}

//export callbackQRadioTunerControlAntennaConnectedChanged
func callbackQRadioTunerControlAntennaConnectedChanged(ptr unsafe.Pointer, ptrName *C.char, connectionStatus C.int) {
	defer qt.Recovering("callback QRadioTunerControl::antennaConnectedChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "antennaConnectedChanged"); signal != nil {
		signal.(func(bool))(int(connectionStatus) != 0)
	}

}

func (ptr *QRadioTunerControl) AntennaConnectedChanged(connectionStatus bool) {
	defer qt.Recovering("QRadioTunerControl::antennaConnectedChanged")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_AntennaConnectedChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(connectionStatus)))
	}
}

func (ptr *QRadioTunerControl) Band() QRadioTuner__Band {
	defer qt.Recovering("QRadioTunerControl::band")

	if ptr.Pointer() != nil {
		return QRadioTuner__Band(C.QRadioTunerControl_Band(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioTunerControl) ConnectBandChanged(f func(band QRadioTuner__Band)) {
	defer qt.Recovering("connect QRadioTunerControl::bandChanged")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectBandChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bandChanged", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectBandChanged() {
	defer qt.Recovering("disconnect QRadioTunerControl::bandChanged")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectBandChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bandChanged")
	}
}

//export callbackQRadioTunerControlBandChanged
func callbackQRadioTunerControlBandChanged(ptr unsafe.Pointer, ptrName *C.char, band C.int) {
	defer qt.Recovering("callback QRadioTunerControl::bandChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "bandChanged"); signal != nil {
		signal.(func(QRadioTuner__Band))(QRadioTuner__Band(band))
	}

}

func (ptr *QRadioTunerControl) BandChanged(band QRadioTuner__Band) {
	defer qt.Recovering("QRadioTunerControl::bandChanged")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_BandChanged(ptr.Pointer(), C.int(band))
	}
}

func (ptr *QRadioTunerControl) CancelSearch() {
	defer qt.Recovering("QRadioTunerControl::cancelSearch")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_CancelSearch(ptr.Pointer())
	}
}

func (ptr *QRadioTunerControl) ConnectError2(f func(error QRadioTuner__Error)) {
	defer qt.Recovering("connect QRadioTunerControl::error")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectError2() {
	defer qt.Recovering("disconnect QRadioTunerControl::error")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQRadioTunerControlError2
func callbackQRadioTunerControlError2(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QRadioTunerControl::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QRadioTuner__Error))(QRadioTuner__Error(error))
	}

}

func (ptr *QRadioTunerControl) Error2(error QRadioTuner__Error) {
	defer qt.Recovering("QRadioTunerControl::error")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_Error2(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QRadioTunerControl) Error() QRadioTuner__Error {
	defer qt.Recovering("QRadioTunerControl::error")

	if ptr.Pointer() != nil {
		return QRadioTuner__Error(C.QRadioTunerControl_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioTunerControl) ErrorString() string {
	defer qt.Recovering("QRadioTunerControl::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioTunerControl_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QRadioTunerControl) Frequency() int {
	defer qt.Recovering("QRadioTunerControl::frequency")

	if ptr.Pointer() != nil {
		return int(C.QRadioTunerControl_Frequency(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioTunerControl) ConnectFrequencyChanged(f func(frequency int)) {
	defer qt.Recovering("connect QRadioTunerControl::frequencyChanged")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectFrequencyChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "frequencyChanged", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectFrequencyChanged() {
	defer qt.Recovering("disconnect QRadioTunerControl::frequencyChanged")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectFrequencyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "frequencyChanged")
	}
}

//export callbackQRadioTunerControlFrequencyChanged
func callbackQRadioTunerControlFrequencyChanged(ptr unsafe.Pointer, ptrName *C.char, frequency C.int) {
	defer qt.Recovering("callback QRadioTunerControl::frequencyChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "frequencyChanged"); signal != nil {
		signal.(func(int))(int(frequency))
	}

}

func (ptr *QRadioTunerControl) FrequencyChanged(frequency int) {
	defer qt.Recovering("QRadioTunerControl::frequencyChanged")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_FrequencyChanged(ptr.Pointer(), C.int(frequency))
	}
}

func (ptr *QRadioTunerControl) FrequencyStep(band QRadioTuner__Band) int {
	defer qt.Recovering("QRadioTunerControl::frequencyStep")

	if ptr.Pointer() != nil {
		return int(C.QRadioTunerControl_FrequencyStep(ptr.Pointer(), C.int(band)))
	}
	return 0
}

func (ptr *QRadioTunerControl) IsAntennaConnected() bool {
	defer qt.Recovering("QRadioTunerControl::isAntennaConnected")

	if ptr.Pointer() != nil {
		return C.QRadioTunerControl_IsAntennaConnected(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRadioTunerControl) IsBandSupported(band QRadioTuner__Band) bool {
	defer qt.Recovering("QRadioTunerControl::isBandSupported")

	if ptr.Pointer() != nil {
		return C.QRadioTunerControl_IsBandSupported(ptr.Pointer(), C.int(band)) != 0
	}
	return false
}

func (ptr *QRadioTunerControl) IsMuted() bool {
	defer qt.Recovering("QRadioTunerControl::isMuted")

	if ptr.Pointer() != nil {
		return C.QRadioTunerControl_IsMuted(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRadioTunerControl) IsSearching() bool {
	defer qt.Recovering("QRadioTunerControl::isSearching")

	if ptr.Pointer() != nil {
		return C.QRadioTunerControl_IsSearching(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRadioTunerControl) IsStereo() bool {
	defer qt.Recovering("QRadioTunerControl::isStereo")

	if ptr.Pointer() != nil {
		return C.QRadioTunerControl_IsStereo(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRadioTunerControl) ConnectMutedChanged(f func(muted bool)) {
	defer qt.Recovering("connect QRadioTunerControl::mutedChanged")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectMutedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mutedChanged", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectMutedChanged() {
	defer qt.Recovering("disconnect QRadioTunerControl::mutedChanged")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectMutedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mutedChanged")
	}
}

//export callbackQRadioTunerControlMutedChanged
func callbackQRadioTunerControlMutedChanged(ptr unsafe.Pointer, ptrName *C.char, muted C.int) {
	defer qt.Recovering("callback QRadioTunerControl::mutedChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "mutedChanged"); signal != nil {
		signal.(func(bool))(int(muted) != 0)
	}

}

func (ptr *QRadioTunerControl) MutedChanged(muted bool) {
	defer qt.Recovering("QRadioTunerControl::mutedChanged")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_MutedChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(muted)))
	}
}

func (ptr *QRadioTunerControl) SearchAllStations(searchMode QRadioTuner__SearchMode) {
	defer qt.Recovering("QRadioTunerControl::searchAllStations")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_SearchAllStations(ptr.Pointer(), C.int(searchMode))
	}
}

func (ptr *QRadioTunerControl) SearchBackward() {
	defer qt.Recovering("QRadioTunerControl::searchBackward")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_SearchBackward(ptr.Pointer())
	}
}

func (ptr *QRadioTunerControl) SearchForward() {
	defer qt.Recovering("QRadioTunerControl::searchForward")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_SearchForward(ptr.Pointer())
	}
}

func (ptr *QRadioTunerControl) ConnectSearchingChanged(f func(searching bool)) {
	defer qt.Recovering("connect QRadioTunerControl::searchingChanged")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectSearchingChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "searchingChanged", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectSearchingChanged() {
	defer qt.Recovering("disconnect QRadioTunerControl::searchingChanged")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectSearchingChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "searchingChanged")
	}
}

//export callbackQRadioTunerControlSearchingChanged
func callbackQRadioTunerControlSearchingChanged(ptr unsafe.Pointer, ptrName *C.char, searching C.int) {
	defer qt.Recovering("callback QRadioTunerControl::searchingChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "searchingChanged"); signal != nil {
		signal.(func(bool))(int(searching) != 0)
	}

}

func (ptr *QRadioTunerControl) SearchingChanged(searching bool) {
	defer qt.Recovering("QRadioTunerControl::searchingChanged")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_SearchingChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(searching)))
	}
}

func (ptr *QRadioTunerControl) SetBand(band QRadioTuner__Band) {
	defer qt.Recovering("QRadioTunerControl::setBand")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_SetBand(ptr.Pointer(), C.int(band))
	}
}

func (ptr *QRadioTunerControl) SetFrequency(frequency int) {
	defer qt.Recovering("QRadioTunerControl::setFrequency")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_SetFrequency(ptr.Pointer(), C.int(frequency))
	}
}

func (ptr *QRadioTunerControl) SetMuted(muted bool) {
	defer qt.Recovering("QRadioTunerControl::setMuted")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_SetMuted(ptr.Pointer(), C.int(qt.GoBoolToInt(muted)))
	}
}

func (ptr *QRadioTunerControl) SetStereoMode(mode QRadioTuner__StereoMode) {
	defer qt.Recovering("QRadioTunerControl::setStereoMode")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_SetStereoMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QRadioTunerControl) SetVolume(volume int) {
	defer qt.Recovering("QRadioTunerControl::setVolume")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_SetVolume(ptr.Pointer(), C.int(volume))
	}
}

func (ptr *QRadioTunerControl) SignalStrength() int {
	defer qt.Recovering("QRadioTunerControl::signalStrength")

	if ptr.Pointer() != nil {
		return int(C.QRadioTunerControl_SignalStrength(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioTunerControl) ConnectSignalStrengthChanged(f func(strength int)) {
	defer qt.Recovering("connect QRadioTunerControl::signalStrengthChanged")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectSignalStrengthChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "signalStrengthChanged", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectSignalStrengthChanged() {
	defer qt.Recovering("disconnect QRadioTunerControl::signalStrengthChanged")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectSignalStrengthChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "signalStrengthChanged")
	}
}

//export callbackQRadioTunerControlSignalStrengthChanged
func callbackQRadioTunerControlSignalStrengthChanged(ptr unsafe.Pointer, ptrName *C.char, strength C.int) {
	defer qt.Recovering("callback QRadioTunerControl::signalStrengthChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "signalStrengthChanged"); signal != nil {
		signal.(func(int))(int(strength))
	}

}

func (ptr *QRadioTunerControl) SignalStrengthChanged(strength int) {
	defer qt.Recovering("QRadioTunerControl::signalStrengthChanged")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_SignalStrengthChanged(ptr.Pointer(), C.int(strength))
	}
}

func (ptr *QRadioTunerControl) Start() {
	defer qt.Recovering("QRadioTunerControl::start")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_Start(ptr.Pointer())
	}
}

func (ptr *QRadioTunerControl) State() QRadioTuner__State {
	defer qt.Recovering("QRadioTunerControl::state")

	if ptr.Pointer() != nil {
		return QRadioTuner__State(C.QRadioTunerControl_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioTunerControl) ConnectStateChanged(f func(state QRadioTuner__State)) {
	defer qt.Recovering("connect QRadioTunerControl::stateChanged")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QRadioTunerControl::stateChanged")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQRadioTunerControlStateChanged
func callbackQRadioTunerControlStateChanged(ptr unsafe.Pointer, ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QRadioTunerControl::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QRadioTuner__State))(QRadioTuner__State(state))
	}

}

func (ptr *QRadioTunerControl) StateChanged(state QRadioTuner__State) {
	defer qt.Recovering("QRadioTunerControl::stateChanged")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_StateChanged(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QRadioTunerControl) ConnectStationFound(f func(frequency int, stationId string)) {
	defer qt.Recovering("connect QRadioTunerControl::stationFound")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectStationFound(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stationFound", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectStationFound() {
	defer qt.Recovering("disconnect QRadioTunerControl::stationFound")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectStationFound(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stationFound")
	}
}

//export callbackQRadioTunerControlStationFound
func callbackQRadioTunerControlStationFound(ptr unsafe.Pointer, ptrName *C.char, frequency C.int, stationId *C.char) {
	defer qt.Recovering("callback QRadioTunerControl::stationFound")

	if signal := qt.GetSignal(C.GoString(ptrName), "stationFound"); signal != nil {
		signal.(func(int, string))(int(frequency), C.GoString(stationId))
	}

}

func (ptr *QRadioTunerControl) StationFound(frequency int, stationId string) {
	defer qt.Recovering("QRadioTunerControl::stationFound")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_StationFound(ptr.Pointer(), C.int(frequency), C.CString(stationId))
	}
}

func (ptr *QRadioTunerControl) StereoMode() QRadioTuner__StereoMode {
	defer qt.Recovering("QRadioTunerControl::stereoMode")

	if ptr.Pointer() != nil {
		return QRadioTuner__StereoMode(C.QRadioTunerControl_StereoMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioTunerControl) ConnectStereoStatusChanged(f func(stereo bool)) {
	defer qt.Recovering("connect QRadioTunerControl::stereoStatusChanged")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectStereoStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stereoStatusChanged", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectStereoStatusChanged() {
	defer qt.Recovering("disconnect QRadioTunerControl::stereoStatusChanged")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectStereoStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stereoStatusChanged")
	}
}

//export callbackQRadioTunerControlStereoStatusChanged
func callbackQRadioTunerControlStereoStatusChanged(ptr unsafe.Pointer, ptrName *C.char, stereo C.int) {
	defer qt.Recovering("callback QRadioTunerControl::stereoStatusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stereoStatusChanged"); signal != nil {
		signal.(func(bool))(int(stereo) != 0)
	}

}

func (ptr *QRadioTunerControl) StereoStatusChanged(stereo bool) {
	defer qt.Recovering("QRadioTunerControl::stereoStatusChanged")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_StereoStatusChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(stereo)))
	}
}

func (ptr *QRadioTunerControl) Stop() {
	defer qt.Recovering("QRadioTunerControl::stop")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_Stop(ptr.Pointer())
	}
}

func (ptr *QRadioTunerControl) Volume() int {
	defer qt.Recovering("QRadioTunerControl::volume")

	if ptr.Pointer() != nil {
		return int(C.QRadioTunerControl_Volume(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioTunerControl) ConnectVolumeChanged(f func(volume int)) {
	defer qt.Recovering("connect QRadioTunerControl::volumeChanged")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectVolumeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "volumeChanged", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectVolumeChanged() {
	defer qt.Recovering("disconnect QRadioTunerControl::volumeChanged")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectVolumeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "volumeChanged")
	}
}

//export callbackQRadioTunerControlVolumeChanged
func callbackQRadioTunerControlVolumeChanged(ptr unsafe.Pointer, ptrName *C.char, volume C.int) {
	defer qt.Recovering("callback QRadioTunerControl::volumeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "volumeChanged"); signal != nil {
		signal.(func(int))(int(volume))
	}

}

func (ptr *QRadioTunerControl) VolumeChanged(volume int) {
	defer qt.Recovering("QRadioTunerControl::volumeChanged")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_VolumeChanged(ptr.Pointer(), C.int(volume))
	}
}

func (ptr *QRadioTunerControl) DestroyQRadioTunerControl() {
	defer qt.Recovering("QRadioTunerControl::~QRadioTunerControl")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DestroyQRadioTunerControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QRadioTunerControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QRadioTunerControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QRadioTunerControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQRadioTunerControlTimerEvent
func callbackQRadioTunerControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioTunerControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQRadioTunerControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QRadioTunerControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QRadioTunerControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QRadioTunerControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QRadioTunerControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QRadioTunerControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QRadioTunerControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QRadioTunerControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQRadioTunerControlChildEvent
func callbackQRadioTunerControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioTunerControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQRadioTunerControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QRadioTunerControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QRadioTunerControl::childEvent")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QRadioTunerControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QRadioTunerControl::childEvent")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QRadioTunerControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QRadioTunerControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QRadioTunerControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQRadioTunerControlCustomEvent
func callbackQRadioTunerControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRadioTunerControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQRadioTunerControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QRadioTunerControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QRadioTunerControl::customEvent")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QRadioTunerControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QRadioTunerControl::customEvent")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QSound struct {
	core.QObject
}

type QSound_ITF interface {
	core.QObject_ITF
	QSound_PTR() *QSound
}

func PointerFromQSound(ptr QSound_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSound_PTR().Pointer()
	}
	return nil
}

func NewQSoundFromPointer(ptr unsafe.Pointer) *QSound {
	var n = new(QSound)
	n.SetPointer(ptr)
	return n
}

func newQSoundFromPointer(ptr unsafe.Pointer) *QSound {
	var n = NewQSoundFromPointer(ptr)
	for len(n.ObjectName()) < len("QSound_") {
		n.SetObjectName("QSound_" + qt.Identifier())
	}
	return n
}

func (ptr *QSound) QSound_PTR() *QSound {
	return ptr
}

//QSound::Loop
type QSound__Loop int64

const (
	QSound__Infinite = QSound__Loop(-1)
)

func (ptr *QSound) SetLoops(number int) {
	defer qt.Recovering("QSound::setLoops")

	if ptr.Pointer() != nil {
		C.QSound_SetLoops(ptr.Pointer(), C.int(number))
	}
}

func NewQSound(filename string, parent core.QObject_ITF) *QSound {
	defer qt.Recovering("QSound::QSound")

	return newQSoundFromPointer(C.QSound_NewQSound(C.CString(filename), core.PointerFromQObject(parent)))
}

func (ptr *QSound) FileName() string {
	defer qt.Recovering("QSound::fileName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSound_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSound) IsFinished() bool {
	defer qt.Recovering("QSound::isFinished")

	if ptr.Pointer() != nil {
		return C.QSound_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSound) Loops() int {
	defer qt.Recovering("QSound::loops")

	if ptr.Pointer() != nil {
		return int(C.QSound_Loops(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSound) LoopsRemaining() int {
	defer qt.Recovering("QSound::loopsRemaining")

	if ptr.Pointer() != nil {
		return int(C.QSound_LoopsRemaining(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSound) Play2() {
	defer qt.Recovering("QSound::play")

	if ptr.Pointer() != nil {
		C.QSound_Play2(ptr.Pointer())
	}
}

func QSound_Play(filename string) {
	defer qt.Recovering("QSound::play")

	C.QSound_QSound_Play(C.CString(filename))
}

func (ptr *QSound) Stop() {
	defer qt.Recovering("QSound::stop")

	if ptr.Pointer() != nil {
		C.QSound_Stop(ptr.Pointer())
	}
}

func (ptr *QSound) DestroyQSound() {
	defer qt.Recovering("QSound::~QSound")

	if ptr.Pointer() != nil {
		C.QSound_DestroyQSound(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSound) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSound::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSound) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSound::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSoundTimerEvent
func callbackQSoundTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSound::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSoundFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSound) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSound::timerEvent")

	if ptr.Pointer() != nil {
		C.QSound_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSound) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSound::timerEvent")

	if ptr.Pointer() != nil {
		C.QSound_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSound) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSound::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSound) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSound::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSoundChildEvent
func callbackQSoundChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSound::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSoundFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSound) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSound::childEvent")

	if ptr.Pointer() != nil {
		C.QSound_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSound) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSound::childEvent")

	if ptr.Pointer() != nil {
		C.QSound_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSound) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSound::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSound) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSound::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSoundCustomEvent
func callbackQSoundCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSound::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSoundFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSound) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSound::customEvent")

	if ptr.Pointer() != nil {
		C.QSound_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSound) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSound::customEvent")

	if ptr.Pointer() != nil {
		C.QSound_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QSoundEffect struct {
	core.QObject
}

type QSoundEffect_ITF interface {
	core.QObject_ITF
	QSoundEffect_PTR() *QSoundEffect
}

func PointerFromQSoundEffect(ptr QSoundEffect_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSoundEffect_PTR().Pointer()
	}
	return nil
}

func NewQSoundEffectFromPointer(ptr unsafe.Pointer) *QSoundEffect {
	var n = new(QSoundEffect)
	n.SetPointer(ptr)
	return n
}

func newQSoundEffectFromPointer(ptr unsafe.Pointer) *QSoundEffect {
	var n = NewQSoundEffectFromPointer(ptr)
	for len(n.ObjectName()) < len("QSoundEffect_") {
		n.SetObjectName("QSoundEffect_" + qt.Identifier())
	}
	return n
}

func (ptr *QSoundEffect) QSoundEffect_PTR() *QSoundEffect {
	return ptr
}

//QSoundEffect::Loop
type QSoundEffect__Loop int64

const (
	QSoundEffect__Infinite = QSoundEffect__Loop(-2)
)

//QSoundEffect::Status
type QSoundEffect__Status int64

const (
	QSoundEffect__Null    = QSoundEffect__Status(0)
	QSoundEffect__Loading = QSoundEffect__Status(1)
	QSoundEffect__Ready   = QSoundEffect__Status(2)
	QSoundEffect__Error   = QSoundEffect__Status(3)
)

func (ptr *QSoundEffect) IsLoaded() bool {
	defer qt.Recovering("QSoundEffect::isLoaded")

	if ptr.Pointer() != nil {
		return C.QSoundEffect_IsLoaded(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSoundEffect) LoopsRemaining() int {
	defer qt.Recovering("QSoundEffect::loopsRemaining")

	if ptr.Pointer() != nil {
		return int(C.QSoundEffect_LoopsRemaining(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSoundEffect) Play() {
	defer qt.Recovering("QSoundEffect::play")

	if ptr.Pointer() != nil {
		C.QSoundEffect_Play(ptr.Pointer())
	}
}

func (ptr *QSoundEffect) Stop() {
	defer qt.Recovering("QSoundEffect::stop")

	if ptr.Pointer() != nil {
		C.QSoundEffect_Stop(ptr.Pointer())
	}
}

func QSoundEffect_SupportedMimeTypes() []string {
	defer qt.Recovering("QSoundEffect::supportedMimeTypes")

	return strings.Split(C.GoString(C.QSoundEffect_QSoundEffect_SupportedMimeTypes()), "|")
}

func NewQSoundEffect(parent core.QObject_ITF) *QSoundEffect {
	defer qt.Recovering("QSoundEffect::QSoundEffect")

	return newQSoundEffectFromPointer(C.QSoundEffect_NewQSoundEffect(core.PointerFromQObject(parent)))
}

func (ptr *QSoundEffect) Category() string {
	defer qt.Recovering("QSoundEffect::category")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSoundEffect_Category(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSoundEffect) ConnectCategoryChanged(f func()) {
	defer qt.Recovering("connect QSoundEffect::categoryChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectCategoryChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "categoryChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectCategoryChanged() {
	defer qt.Recovering("disconnect QSoundEffect::categoryChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectCategoryChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "categoryChanged")
	}
}

//export callbackQSoundEffectCategoryChanged
func callbackQSoundEffectCategoryChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSoundEffect::categoryChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "categoryChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSoundEffect) CategoryChanged() {
	defer qt.Recovering("QSoundEffect::categoryChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_CategoryChanged(ptr.Pointer())
	}
}

func (ptr *QSoundEffect) IsMuted() bool {
	defer qt.Recovering("QSoundEffect::isMuted")

	if ptr.Pointer() != nil {
		return C.QSoundEffect_IsMuted(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSoundEffect) IsPlaying() bool {
	defer qt.Recovering("QSoundEffect::isPlaying")

	if ptr.Pointer() != nil {
		return C.QSoundEffect_IsPlaying(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSoundEffect) ConnectLoadedChanged(f func()) {
	defer qt.Recovering("connect QSoundEffect::loadedChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectLoadedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loadedChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectLoadedChanged() {
	defer qt.Recovering("disconnect QSoundEffect::loadedChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectLoadedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loadedChanged")
	}
}

//export callbackQSoundEffectLoadedChanged
func callbackQSoundEffectLoadedChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSoundEffect::loadedChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "loadedChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSoundEffect) LoadedChanged() {
	defer qt.Recovering("QSoundEffect::loadedChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_LoadedChanged(ptr.Pointer())
	}
}

func (ptr *QSoundEffect) LoopCount() int {
	defer qt.Recovering("QSoundEffect::loopCount")

	if ptr.Pointer() != nil {
		return int(C.QSoundEffect_LoopCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSoundEffect) ConnectLoopCountChanged(f func()) {
	defer qt.Recovering("connect QSoundEffect::loopCountChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectLoopCountChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loopCountChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectLoopCountChanged() {
	defer qt.Recovering("disconnect QSoundEffect::loopCountChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectLoopCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loopCountChanged")
	}
}

//export callbackQSoundEffectLoopCountChanged
func callbackQSoundEffectLoopCountChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSoundEffect::loopCountChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "loopCountChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSoundEffect) LoopCountChanged() {
	defer qt.Recovering("QSoundEffect::loopCountChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_LoopCountChanged(ptr.Pointer())
	}
}

func (ptr *QSoundEffect) ConnectLoopsRemainingChanged(f func()) {
	defer qt.Recovering("connect QSoundEffect::loopsRemainingChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectLoopsRemainingChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loopsRemainingChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectLoopsRemainingChanged() {
	defer qt.Recovering("disconnect QSoundEffect::loopsRemainingChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectLoopsRemainingChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loopsRemainingChanged")
	}
}

//export callbackQSoundEffectLoopsRemainingChanged
func callbackQSoundEffectLoopsRemainingChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSoundEffect::loopsRemainingChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "loopsRemainingChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSoundEffect) LoopsRemainingChanged() {
	defer qt.Recovering("QSoundEffect::loopsRemainingChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_LoopsRemainingChanged(ptr.Pointer())
	}
}

func (ptr *QSoundEffect) ConnectMutedChanged(f func()) {
	defer qt.Recovering("connect QSoundEffect::mutedChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectMutedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mutedChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectMutedChanged() {
	defer qt.Recovering("disconnect QSoundEffect::mutedChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectMutedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mutedChanged")
	}
}

//export callbackQSoundEffectMutedChanged
func callbackQSoundEffectMutedChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSoundEffect::mutedChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "mutedChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSoundEffect) MutedChanged() {
	defer qt.Recovering("QSoundEffect::mutedChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_MutedChanged(ptr.Pointer())
	}
}

func (ptr *QSoundEffect) ConnectPlayingChanged(f func()) {
	defer qt.Recovering("connect QSoundEffect::playingChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectPlayingChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "playingChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectPlayingChanged() {
	defer qt.Recovering("disconnect QSoundEffect::playingChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectPlayingChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "playingChanged")
	}
}

//export callbackQSoundEffectPlayingChanged
func callbackQSoundEffectPlayingChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSoundEffect::playingChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "playingChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSoundEffect) PlayingChanged() {
	defer qt.Recovering("QSoundEffect::playingChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_PlayingChanged(ptr.Pointer())
	}
}

func (ptr *QSoundEffect) SetCategory(category string) {
	defer qt.Recovering("QSoundEffect::setCategory")

	if ptr.Pointer() != nil {
		C.QSoundEffect_SetCategory(ptr.Pointer(), C.CString(category))
	}
}

func (ptr *QSoundEffect) SetLoopCount(loopCount int) {
	defer qt.Recovering("QSoundEffect::setLoopCount")

	if ptr.Pointer() != nil {
		C.QSoundEffect_SetLoopCount(ptr.Pointer(), C.int(loopCount))
	}
}

func (ptr *QSoundEffect) SetMuted(muted bool) {
	defer qt.Recovering("QSoundEffect::setMuted")

	if ptr.Pointer() != nil {
		C.QSoundEffect_SetMuted(ptr.Pointer(), C.int(qt.GoBoolToInt(muted)))
	}
}

func (ptr *QSoundEffect) SetSource(url core.QUrl_ITF) {
	defer qt.Recovering("QSoundEffect::setSource")

	if ptr.Pointer() != nil {
		C.QSoundEffect_SetSource(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QSoundEffect) SetVolume(volume float64) {
	defer qt.Recovering("QSoundEffect::setVolume")

	if ptr.Pointer() != nil {
		C.QSoundEffect_SetVolume(ptr.Pointer(), C.double(volume))
	}
}

func (ptr *QSoundEffect) Source() *core.QUrl {
	defer qt.Recovering("QSoundEffect::source")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QSoundEffect_Source(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSoundEffect) ConnectSourceChanged(f func()) {
	defer qt.Recovering("connect QSoundEffect::sourceChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectSourceChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sourceChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectSourceChanged() {
	defer qt.Recovering("disconnect QSoundEffect::sourceChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectSourceChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sourceChanged")
	}
}

//export callbackQSoundEffectSourceChanged
func callbackQSoundEffectSourceChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSoundEffect::sourceChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "sourceChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSoundEffect) SourceChanged() {
	defer qt.Recovering("QSoundEffect::sourceChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_SourceChanged(ptr.Pointer())
	}
}

func (ptr *QSoundEffect) Status() QSoundEffect__Status {
	defer qt.Recovering("QSoundEffect::status")

	if ptr.Pointer() != nil {
		return QSoundEffect__Status(C.QSoundEffect_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSoundEffect) ConnectStatusChanged(f func()) {
	defer qt.Recovering("connect QSoundEffect::statusChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "statusChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectStatusChanged() {
	defer qt.Recovering("disconnect QSoundEffect::statusChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "statusChanged")
	}
}

//export callbackQSoundEffectStatusChanged
func callbackQSoundEffectStatusChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSoundEffect::statusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "statusChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSoundEffect) StatusChanged() {
	defer qt.Recovering("QSoundEffect::statusChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_StatusChanged(ptr.Pointer())
	}
}

func (ptr *QSoundEffect) Volume() float64 {
	defer qt.Recovering("QSoundEffect::volume")

	if ptr.Pointer() != nil {
		return float64(C.QSoundEffect_Volume(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSoundEffect) ConnectVolumeChanged(f func()) {
	defer qt.Recovering("connect QSoundEffect::volumeChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectVolumeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "volumeChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectVolumeChanged() {
	defer qt.Recovering("disconnect QSoundEffect::volumeChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectVolumeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "volumeChanged")
	}
}

//export callbackQSoundEffectVolumeChanged
func callbackQSoundEffectVolumeChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSoundEffect::volumeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "volumeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSoundEffect) VolumeChanged() {
	defer qt.Recovering("QSoundEffect::volumeChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_VolumeChanged(ptr.Pointer())
	}
}

func (ptr *QSoundEffect) DestroyQSoundEffect() {
	defer qt.Recovering("QSoundEffect::~QSoundEffect")

	if ptr.Pointer() != nil {
		C.QSoundEffect_DestroyQSoundEffect(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSoundEffect) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSoundEffect::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSoundEffect) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSoundEffect::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSoundEffectTimerEvent
func callbackQSoundEffectTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSoundEffect::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSoundEffectFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSoundEffect) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSoundEffect::timerEvent")

	if ptr.Pointer() != nil {
		C.QSoundEffect_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSoundEffect) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSoundEffect::timerEvent")

	if ptr.Pointer() != nil {
		C.QSoundEffect_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSoundEffect) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSoundEffect::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSoundEffect) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSoundEffect::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSoundEffectChildEvent
func callbackQSoundEffectChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSoundEffect::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSoundEffectFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSoundEffect) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSoundEffect::childEvent")

	if ptr.Pointer() != nil {
		C.QSoundEffect_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSoundEffect) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSoundEffect::childEvent")

	if ptr.Pointer() != nil {
		C.QSoundEffect_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSoundEffect) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSoundEffect::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSoundEffect) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSoundEffect::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSoundEffectCustomEvent
func callbackQSoundEffectCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSoundEffect::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSoundEffectFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSoundEffect) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSoundEffect::customEvent")

	if ptr.Pointer() != nil {
		C.QSoundEffect_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSoundEffect) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSoundEffect::customEvent")

	if ptr.Pointer() != nil {
		C.QSoundEffect_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QVideoDeviceSelectorControl struct {
	QMediaControl
}

type QVideoDeviceSelectorControl_ITF interface {
	QMediaControl_ITF
	QVideoDeviceSelectorControl_PTR() *QVideoDeviceSelectorControl
}

func PointerFromQVideoDeviceSelectorControl(ptr QVideoDeviceSelectorControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoDeviceSelectorControl_PTR().Pointer()
	}
	return nil
}

func NewQVideoDeviceSelectorControlFromPointer(ptr unsafe.Pointer) *QVideoDeviceSelectorControl {
	var n = new(QVideoDeviceSelectorControl)
	n.SetPointer(ptr)
	return n
}

func newQVideoDeviceSelectorControlFromPointer(ptr unsafe.Pointer) *QVideoDeviceSelectorControl {
	var n = NewQVideoDeviceSelectorControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QVideoDeviceSelectorControl_") {
		n.SetObjectName("QVideoDeviceSelectorControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QVideoDeviceSelectorControl) QVideoDeviceSelectorControl_PTR() *QVideoDeviceSelectorControl {
	return ptr
}

func (ptr *QVideoDeviceSelectorControl) DefaultDevice() int {
	defer qt.Recovering("QVideoDeviceSelectorControl::defaultDevice")

	if ptr.Pointer() != nil {
		return int(C.QVideoDeviceSelectorControl_DefaultDevice(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoDeviceSelectorControl) DeviceCount() int {
	defer qt.Recovering("QVideoDeviceSelectorControl::deviceCount")

	if ptr.Pointer() != nil {
		return int(C.QVideoDeviceSelectorControl_DeviceCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoDeviceSelectorControl) DeviceDescription(index int) string {
	defer qt.Recovering("QVideoDeviceSelectorControl::deviceDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QVideoDeviceSelectorControl_DeviceDescription(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QVideoDeviceSelectorControl) DeviceName(index int) string {
	defer qt.Recovering("QVideoDeviceSelectorControl::deviceName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QVideoDeviceSelectorControl_DeviceName(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QVideoDeviceSelectorControl) ConnectDevicesChanged(f func()) {
	defer qt.Recovering("connect QVideoDeviceSelectorControl::devicesChanged")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_ConnectDevicesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "devicesChanged", f)
	}
}

func (ptr *QVideoDeviceSelectorControl) DisconnectDevicesChanged() {
	defer qt.Recovering("disconnect QVideoDeviceSelectorControl::devicesChanged")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_DisconnectDevicesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "devicesChanged")
	}
}

//export callbackQVideoDeviceSelectorControlDevicesChanged
func callbackQVideoDeviceSelectorControlDevicesChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QVideoDeviceSelectorControl::devicesChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "devicesChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVideoDeviceSelectorControl) DevicesChanged() {
	defer qt.Recovering("QVideoDeviceSelectorControl::devicesChanged")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_DevicesChanged(ptr.Pointer())
	}
}

func (ptr *QVideoDeviceSelectorControl) SelectedDevice() int {
	defer qt.Recovering("QVideoDeviceSelectorControl::selectedDevice")

	if ptr.Pointer() != nil {
		return int(C.QVideoDeviceSelectorControl_SelectedDevice(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoDeviceSelectorControl) ConnectSelectedDeviceChanged2(f func(name string)) {
	defer qt.Recovering("connect QVideoDeviceSelectorControl::selectedDeviceChanged")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_ConnectSelectedDeviceChanged2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "selectedDeviceChanged2", f)
	}
}

func (ptr *QVideoDeviceSelectorControl) DisconnectSelectedDeviceChanged2() {
	defer qt.Recovering("disconnect QVideoDeviceSelectorControl::selectedDeviceChanged")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_DisconnectSelectedDeviceChanged2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "selectedDeviceChanged2")
	}
}

//export callbackQVideoDeviceSelectorControlSelectedDeviceChanged2
func callbackQVideoDeviceSelectorControlSelectedDeviceChanged2(ptr unsafe.Pointer, ptrName *C.char, name *C.char) {
	defer qt.Recovering("callback QVideoDeviceSelectorControl::selectedDeviceChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectedDeviceChanged2"); signal != nil {
		signal.(func(string))(C.GoString(name))
	}

}

func (ptr *QVideoDeviceSelectorControl) SelectedDeviceChanged2(name string) {
	defer qt.Recovering("QVideoDeviceSelectorControl::selectedDeviceChanged")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_SelectedDeviceChanged2(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QVideoDeviceSelectorControl) ConnectSelectedDeviceChanged(f func(index int)) {
	defer qt.Recovering("connect QVideoDeviceSelectorControl::selectedDeviceChanged")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_ConnectSelectedDeviceChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "selectedDeviceChanged", f)
	}
}

func (ptr *QVideoDeviceSelectorControl) DisconnectSelectedDeviceChanged() {
	defer qt.Recovering("disconnect QVideoDeviceSelectorControl::selectedDeviceChanged")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_DisconnectSelectedDeviceChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "selectedDeviceChanged")
	}
}

//export callbackQVideoDeviceSelectorControlSelectedDeviceChanged
func callbackQVideoDeviceSelectorControlSelectedDeviceChanged(ptr unsafe.Pointer, ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QVideoDeviceSelectorControl::selectedDeviceChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectedDeviceChanged"); signal != nil {
		signal.(func(int))(int(index))
	}

}

func (ptr *QVideoDeviceSelectorControl) SelectedDeviceChanged(index int) {
	defer qt.Recovering("QVideoDeviceSelectorControl::selectedDeviceChanged")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_SelectedDeviceChanged(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QVideoDeviceSelectorControl) SetSelectedDevice(index int) {
	defer qt.Recovering("QVideoDeviceSelectorControl::setSelectedDevice")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_SetSelectedDevice(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QVideoDeviceSelectorControl) DestroyQVideoDeviceSelectorControl() {
	defer qt.Recovering("QVideoDeviceSelectorControl::~QVideoDeviceSelectorControl")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_DestroyQVideoDeviceSelectorControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QVideoDeviceSelectorControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QVideoDeviceSelectorControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QVideoDeviceSelectorControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QVideoDeviceSelectorControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQVideoDeviceSelectorControlTimerEvent
func callbackQVideoDeviceSelectorControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoDeviceSelectorControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQVideoDeviceSelectorControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QVideoDeviceSelectorControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QVideoDeviceSelectorControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QVideoDeviceSelectorControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QVideoDeviceSelectorControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QVideoDeviceSelectorControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QVideoDeviceSelectorControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QVideoDeviceSelectorControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QVideoDeviceSelectorControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQVideoDeviceSelectorControlChildEvent
func callbackQVideoDeviceSelectorControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoDeviceSelectorControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQVideoDeviceSelectorControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QVideoDeviceSelectorControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QVideoDeviceSelectorControl::childEvent")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QVideoDeviceSelectorControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QVideoDeviceSelectorControl::childEvent")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QVideoDeviceSelectorControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QVideoDeviceSelectorControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QVideoDeviceSelectorControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QVideoDeviceSelectorControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQVideoDeviceSelectorControlCustomEvent
func callbackQVideoDeviceSelectorControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoDeviceSelectorControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQVideoDeviceSelectorControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QVideoDeviceSelectorControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoDeviceSelectorControl::customEvent")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QVideoDeviceSelectorControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoDeviceSelectorControl::customEvent")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QVideoEncoderSettings struct {
	ptr unsafe.Pointer
}

type QVideoEncoderSettings_ITF interface {
	QVideoEncoderSettings_PTR() *QVideoEncoderSettings
}

func (p *QVideoEncoderSettings) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QVideoEncoderSettings) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQVideoEncoderSettings(ptr QVideoEncoderSettings_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoEncoderSettings_PTR().Pointer()
	}
	return nil
}

func NewQVideoEncoderSettingsFromPointer(ptr unsafe.Pointer) *QVideoEncoderSettings {
	var n = new(QVideoEncoderSettings)
	n.SetPointer(ptr)
	return n
}

func newQVideoEncoderSettingsFromPointer(ptr unsafe.Pointer) *QVideoEncoderSettings {
	var n = NewQVideoEncoderSettingsFromPointer(ptr)
	return n
}

func (ptr *QVideoEncoderSettings) QVideoEncoderSettings_PTR() *QVideoEncoderSettings {
	return ptr
}

func (ptr *QVideoEncoderSettings) SetFrameRate(rate float64) {
	defer qt.Recovering("QVideoEncoderSettings::setFrameRate")

	if ptr.Pointer() != nil {
		C.QVideoEncoderSettings_SetFrameRate(ptr.Pointer(), C.double(rate))
	}
}

func NewQVideoEncoderSettings() *QVideoEncoderSettings {
	defer qt.Recovering("QVideoEncoderSettings::QVideoEncoderSettings")

	return newQVideoEncoderSettingsFromPointer(C.QVideoEncoderSettings_NewQVideoEncoderSettings())
}

func NewQVideoEncoderSettings2(other QVideoEncoderSettings_ITF) *QVideoEncoderSettings {
	defer qt.Recovering("QVideoEncoderSettings::QVideoEncoderSettings")

	return newQVideoEncoderSettingsFromPointer(C.QVideoEncoderSettings_NewQVideoEncoderSettings2(PointerFromQVideoEncoderSettings(other)))
}

func (ptr *QVideoEncoderSettings) BitRate() int {
	defer qt.Recovering("QVideoEncoderSettings::bitRate")

	if ptr.Pointer() != nil {
		return int(C.QVideoEncoderSettings_BitRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoEncoderSettings) Codec() string {
	defer qt.Recovering("QVideoEncoderSettings::codec")

	if ptr.Pointer() != nil {
		return C.GoString(C.QVideoEncoderSettings_Codec(ptr.Pointer()))
	}
	return ""
}

func (ptr *QVideoEncoderSettings) EncodingMode() QMultimedia__EncodingMode {
	defer qt.Recovering("QVideoEncoderSettings::encodingMode")

	if ptr.Pointer() != nil {
		return QMultimedia__EncodingMode(C.QVideoEncoderSettings_EncodingMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoEncoderSettings) EncodingOption(option string) *core.QVariant {
	defer qt.Recovering("QVideoEncoderSettings::encodingOption")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QVideoEncoderSettings_EncodingOption(ptr.Pointer(), C.CString(option)))
	}
	return nil
}

func (ptr *QVideoEncoderSettings) FrameRate() float64 {
	defer qt.Recovering("QVideoEncoderSettings::frameRate")

	if ptr.Pointer() != nil {
		return float64(C.QVideoEncoderSettings_FrameRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoEncoderSettings) IsNull() bool {
	defer qt.Recovering("QVideoEncoderSettings::isNull")

	if ptr.Pointer() != nil {
		return C.QVideoEncoderSettings_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVideoEncoderSettings) Quality() QMultimedia__EncodingQuality {
	defer qt.Recovering("QVideoEncoderSettings::quality")

	if ptr.Pointer() != nil {
		return QMultimedia__EncodingQuality(C.QVideoEncoderSettings_Quality(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoEncoderSettings) Resolution() *core.QSize {
	defer qt.Recovering("QVideoEncoderSettings::resolution")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QVideoEncoderSettings_Resolution(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVideoEncoderSettings) SetBitRate(value int) {
	defer qt.Recovering("QVideoEncoderSettings::setBitRate")

	if ptr.Pointer() != nil {
		C.QVideoEncoderSettings_SetBitRate(ptr.Pointer(), C.int(value))
	}
}

func (ptr *QVideoEncoderSettings) SetCodec(codec string) {
	defer qt.Recovering("QVideoEncoderSettings::setCodec")

	if ptr.Pointer() != nil {
		C.QVideoEncoderSettings_SetCodec(ptr.Pointer(), C.CString(codec))
	}
}

func (ptr *QVideoEncoderSettings) SetEncodingMode(mode QMultimedia__EncodingMode) {
	defer qt.Recovering("QVideoEncoderSettings::setEncodingMode")

	if ptr.Pointer() != nil {
		C.QVideoEncoderSettings_SetEncodingMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QVideoEncoderSettings) SetEncodingOption(option string, value core.QVariant_ITF) {
	defer qt.Recovering("QVideoEncoderSettings::setEncodingOption")

	if ptr.Pointer() != nil {
		C.QVideoEncoderSettings_SetEncodingOption(ptr.Pointer(), C.CString(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QVideoEncoderSettings) SetQuality(quality QMultimedia__EncodingQuality) {
	defer qt.Recovering("QVideoEncoderSettings::setQuality")

	if ptr.Pointer() != nil {
		C.QVideoEncoderSettings_SetQuality(ptr.Pointer(), C.int(quality))
	}
}

func (ptr *QVideoEncoderSettings) SetResolution(resolution core.QSize_ITF) {
	defer qt.Recovering("QVideoEncoderSettings::setResolution")

	if ptr.Pointer() != nil {
		C.QVideoEncoderSettings_SetResolution(ptr.Pointer(), core.PointerFromQSize(resolution))
	}
}

func (ptr *QVideoEncoderSettings) SetResolution2(width int, height int) {
	defer qt.Recovering("QVideoEncoderSettings::setResolution")

	if ptr.Pointer() != nil {
		C.QVideoEncoderSettings_SetResolution2(ptr.Pointer(), C.int(width), C.int(height))
	}
}

func (ptr *QVideoEncoderSettings) DestroyQVideoEncoderSettings() {
	defer qt.Recovering("QVideoEncoderSettings::~QVideoEncoderSettings")

	if ptr.Pointer() != nil {
		C.QVideoEncoderSettings_DestroyQVideoEncoderSettings(ptr.Pointer())
	}
}

type QVideoEncoderSettingsControl struct {
	QMediaControl
}

type QVideoEncoderSettingsControl_ITF interface {
	QMediaControl_ITF
	QVideoEncoderSettingsControl_PTR() *QVideoEncoderSettingsControl
}

func PointerFromQVideoEncoderSettingsControl(ptr QVideoEncoderSettingsControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoEncoderSettingsControl_PTR().Pointer()
	}
	return nil
}

func NewQVideoEncoderSettingsControlFromPointer(ptr unsafe.Pointer) *QVideoEncoderSettingsControl {
	var n = new(QVideoEncoderSettingsControl)
	n.SetPointer(ptr)
	return n
}

func newQVideoEncoderSettingsControlFromPointer(ptr unsafe.Pointer) *QVideoEncoderSettingsControl {
	var n = NewQVideoEncoderSettingsControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QVideoEncoderSettingsControl_") {
		n.SetObjectName("QVideoEncoderSettingsControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QVideoEncoderSettingsControl) QVideoEncoderSettingsControl_PTR() *QVideoEncoderSettingsControl {
	return ptr
}

func (ptr *QVideoEncoderSettingsControl) SetVideoSettings(settings QVideoEncoderSettings_ITF) {
	defer qt.Recovering("QVideoEncoderSettingsControl::setVideoSettings")

	if ptr.Pointer() != nil {
		C.QVideoEncoderSettingsControl_SetVideoSettings(ptr.Pointer(), PointerFromQVideoEncoderSettings(settings))
	}
}

func (ptr *QVideoEncoderSettingsControl) SupportedVideoCodecs() []string {
	defer qt.Recovering("QVideoEncoderSettingsControl::supportedVideoCodecs")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QVideoEncoderSettingsControl_SupportedVideoCodecs(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QVideoEncoderSettingsControl) VideoCodecDescription(codec string) string {
	defer qt.Recovering("QVideoEncoderSettingsControl::videoCodecDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QVideoEncoderSettingsControl_VideoCodecDescription(ptr.Pointer(), C.CString(codec)))
	}
	return ""
}

func (ptr *QVideoEncoderSettingsControl) VideoSettings() *QVideoEncoderSettings {
	defer qt.Recovering("QVideoEncoderSettingsControl::videoSettings")

	if ptr.Pointer() != nil {
		return NewQVideoEncoderSettingsFromPointer(C.QVideoEncoderSettingsControl_VideoSettings(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVideoEncoderSettingsControl) DestroyQVideoEncoderSettingsControl() {
	defer qt.Recovering("QVideoEncoderSettingsControl::~QVideoEncoderSettingsControl")

	if ptr.Pointer() != nil {
		C.QVideoEncoderSettingsControl_DestroyQVideoEncoderSettingsControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QVideoEncoderSettingsControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QVideoEncoderSettingsControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QVideoEncoderSettingsControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QVideoEncoderSettingsControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQVideoEncoderSettingsControlTimerEvent
func callbackQVideoEncoderSettingsControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoEncoderSettingsControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQVideoEncoderSettingsControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QVideoEncoderSettingsControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QVideoEncoderSettingsControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QVideoEncoderSettingsControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QVideoEncoderSettingsControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QVideoEncoderSettingsControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QVideoEncoderSettingsControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QVideoEncoderSettingsControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QVideoEncoderSettingsControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QVideoEncoderSettingsControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QVideoEncoderSettingsControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQVideoEncoderSettingsControlChildEvent
func callbackQVideoEncoderSettingsControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoEncoderSettingsControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQVideoEncoderSettingsControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QVideoEncoderSettingsControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QVideoEncoderSettingsControl::childEvent")

	if ptr.Pointer() != nil {
		C.QVideoEncoderSettingsControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QVideoEncoderSettingsControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QVideoEncoderSettingsControl::childEvent")

	if ptr.Pointer() != nil {
		C.QVideoEncoderSettingsControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QVideoEncoderSettingsControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QVideoEncoderSettingsControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QVideoEncoderSettingsControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QVideoEncoderSettingsControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQVideoEncoderSettingsControlCustomEvent
func callbackQVideoEncoderSettingsControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoEncoderSettingsControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQVideoEncoderSettingsControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QVideoEncoderSettingsControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoEncoderSettingsControl::customEvent")

	if ptr.Pointer() != nil {
		C.QVideoEncoderSettingsControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QVideoEncoderSettingsControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoEncoderSettingsControl::customEvent")

	if ptr.Pointer() != nil {
		C.QVideoEncoderSettingsControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QVideoFilterRunnable struct {
	ptr unsafe.Pointer
}

type QVideoFilterRunnable_ITF interface {
	QVideoFilterRunnable_PTR() *QVideoFilterRunnable
}

func (p *QVideoFilterRunnable) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QVideoFilterRunnable) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQVideoFilterRunnable(ptr QVideoFilterRunnable_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoFilterRunnable_PTR().Pointer()
	}
	return nil
}

func NewQVideoFilterRunnableFromPointer(ptr unsafe.Pointer) *QVideoFilterRunnable {
	var n = new(QVideoFilterRunnable)
	n.SetPointer(ptr)
	return n
}

func newQVideoFilterRunnableFromPointer(ptr unsafe.Pointer) *QVideoFilterRunnable {
	var n = NewQVideoFilterRunnableFromPointer(ptr)
	return n
}

func (ptr *QVideoFilterRunnable) QVideoFilterRunnable_PTR() *QVideoFilterRunnable {
	return ptr
}

//QVideoFilterRunnable::RunFlag
type QVideoFilterRunnable__RunFlag int64

const (
	QVideoFilterRunnable__LastInChain = QVideoFilterRunnable__RunFlag(0x01)
)

func (ptr *QVideoFilterRunnable) Run(input QVideoFrame_ITF, surfaceFormat QVideoSurfaceFormat_ITF, flags QVideoFilterRunnable__RunFlag) *QVideoFrame {
	defer qt.Recovering("QVideoFilterRunnable::run")

	if ptr.Pointer() != nil {
		return NewQVideoFrameFromPointer(C.QVideoFilterRunnable_Run(ptr.Pointer(), PointerFromQVideoFrame(input), PointerFromQVideoSurfaceFormat(surfaceFormat), C.int(flags)))
	}
	return nil
}

type QVideoFrame struct {
	ptr unsafe.Pointer
}

type QVideoFrame_ITF interface {
	QVideoFrame_PTR() *QVideoFrame
}

func (p *QVideoFrame) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QVideoFrame) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQVideoFrame(ptr QVideoFrame_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoFrame_PTR().Pointer()
	}
	return nil
}

func NewQVideoFrameFromPointer(ptr unsafe.Pointer) *QVideoFrame {
	var n = new(QVideoFrame)
	n.SetPointer(ptr)
	return n
}

func newQVideoFrameFromPointer(ptr unsafe.Pointer) *QVideoFrame {
	var n = NewQVideoFrameFromPointer(ptr)
	return n
}

func (ptr *QVideoFrame) QVideoFrame_PTR() *QVideoFrame {
	return ptr
}

//QVideoFrame::FieldType
type QVideoFrame__FieldType int64

const (
	QVideoFrame__ProgressiveFrame = QVideoFrame__FieldType(0)
	QVideoFrame__TopField         = QVideoFrame__FieldType(1)
	QVideoFrame__BottomField      = QVideoFrame__FieldType(2)
	QVideoFrame__InterlacedFrame  = QVideoFrame__FieldType(3)
)

//QVideoFrame::PixelFormat
type QVideoFrame__PixelFormat int64

const (
	QVideoFrame__Format_Invalid                = QVideoFrame__PixelFormat(0)
	QVideoFrame__Format_ARGB32                 = QVideoFrame__PixelFormat(1)
	QVideoFrame__Format_ARGB32_Premultiplied   = QVideoFrame__PixelFormat(2)
	QVideoFrame__Format_RGB32                  = QVideoFrame__PixelFormat(3)
	QVideoFrame__Format_RGB24                  = QVideoFrame__PixelFormat(4)
	QVideoFrame__Format_RGB565                 = QVideoFrame__PixelFormat(5)
	QVideoFrame__Format_RGB555                 = QVideoFrame__PixelFormat(6)
	QVideoFrame__Format_ARGB8565_Premultiplied = QVideoFrame__PixelFormat(7)
	QVideoFrame__Format_BGRA32                 = QVideoFrame__PixelFormat(8)
	QVideoFrame__Format_BGRA32_Premultiplied   = QVideoFrame__PixelFormat(9)
	QVideoFrame__Format_BGR32                  = QVideoFrame__PixelFormat(10)
	QVideoFrame__Format_BGR24                  = QVideoFrame__PixelFormat(11)
	QVideoFrame__Format_BGR565                 = QVideoFrame__PixelFormat(12)
	QVideoFrame__Format_BGR555                 = QVideoFrame__PixelFormat(13)
	QVideoFrame__Format_BGRA5658_Premultiplied = QVideoFrame__PixelFormat(14)
	QVideoFrame__Format_AYUV444                = QVideoFrame__PixelFormat(15)
	QVideoFrame__Format_AYUV444_Premultiplied  = QVideoFrame__PixelFormat(16)
	QVideoFrame__Format_YUV444                 = QVideoFrame__PixelFormat(17)
	QVideoFrame__Format_YUV420P                = QVideoFrame__PixelFormat(18)
	QVideoFrame__Format_YV12                   = QVideoFrame__PixelFormat(19)
	QVideoFrame__Format_UYVY                   = QVideoFrame__PixelFormat(20)
	QVideoFrame__Format_YUYV                   = QVideoFrame__PixelFormat(21)
	QVideoFrame__Format_NV12                   = QVideoFrame__PixelFormat(22)
	QVideoFrame__Format_NV21                   = QVideoFrame__PixelFormat(23)
	QVideoFrame__Format_IMC1                   = QVideoFrame__PixelFormat(24)
	QVideoFrame__Format_IMC2                   = QVideoFrame__PixelFormat(25)
	QVideoFrame__Format_IMC3                   = QVideoFrame__PixelFormat(26)
	QVideoFrame__Format_IMC4                   = QVideoFrame__PixelFormat(27)
	QVideoFrame__Format_Y8                     = QVideoFrame__PixelFormat(28)
	QVideoFrame__Format_Y16                    = QVideoFrame__PixelFormat(29)
	QVideoFrame__Format_Jpeg                   = QVideoFrame__PixelFormat(30)
	QVideoFrame__Format_CameraRaw              = QVideoFrame__PixelFormat(31)
	QVideoFrame__Format_AdobeDng               = QVideoFrame__PixelFormat(32)
	QVideoFrame__Format_User                   = QVideoFrame__PixelFormat(1000)
)

func NewQVideoFrame() *QVideoFrame {
	defer qt.Recovering("QVideoFrame::QVideoFrame")

	return newQVideoFrameFromPointer(C.QVideoFrame_NewQVideoFrame())
}

func NewQVideoFrame2(buffer QAbstractVideoBuffer_ITF, size core.QSize_ITF, format QVideoFrame__PixelFormat) *QVideoFrame {
	defer qt.Recovering("QVideoFrame::QVideoFrame")

	return newQVideoFrameFromPointer(C.QVideoFrame_NewQVideoFrame2(PointerFromQAbstractVideoBuffer(buffer), core.PointerFromQSize(size), C.int(format)))
}

func NewQVideoFrame4(image gui.QImage_ITF) *QVideoFrame {
	defer qt.Recovering("QVideoFrame::QVideoFrame")

	return newQVideoFrameFromPointer(C.QVideoFrame_NewQVideoFrame4(gui.PointerFromQImage(image)))
}

func NewQVideoFrame5(other QVideoFrame_ITF) *QVideoFrame {
	defer qt.Recovering("QVideoFrame::QVideoFrame")

	return newQVideoFrameFromPointer(C.QVideoFrame_NewQVideoFrame5(PointerFromQVideoFrame(other)))
}

func NewQVideoFrame3(bytes int, size core.QSize_ITF, bytesPerLine int, format QVideoFrame__PixelFormat) *QVideoFrame {
	defer qt.Recovering("QVideoFrame::QVideoFrame")

	return newQVideoFrameFromPointer(C.QVideoFrame_NewQVideoFrame3(C.int(bytes), core.PointerFromQSize(size), C.int(bytesPerLine), C.int(format)))
}

func (ptr *QVideoFrame) BytesPerLine() int {
	defer qt.Recovering("QVideoFrame::bytesPerLine")

	if ptr.Pointer() != nil {
		return int(C.QVideoFrame_BytesPerLine(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoFrame) BytesPerLine2(plane int) int {
	defer qt.Recovering("QVideoFrame::bytesPerLine")

	if ptr.Pointer() != nil {
		return int(C.QVideoFrame_BytesPerLine2(ptr.Pointer(), C.int(plane)))
	}
	return 0
}

func (ptr *QVideoFrame) EndTime() int64 {
	defer qt.Recovering("QVideoFrame::endTime")

	if ptr.Pointer() != nil {
		return int64(C.QVideoFrame_EndTime(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoFrame) FieldType() QVideoFrame__FieldType {
	defer qt.Recovering("QVideoFrame::fieldType")

	if ptr.Pointer() != nil {
		return QVideoFrame__FieldType(C.QVideoFrame_FieldType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoFrame) Handle() *core.QVariant {
	defer qt.Recovering("QVideoFrame::handle")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QVideoFrame_Handle(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVideoFrame) HandleType() QAbstractVideoBuffer__HandleType {
	defer qt.Recovering("QVideoFrame::handleType")

	if ptr.Pointer() != nil {
		return QAbstractVideoBuffer__HandleType(C.QVideoFrame_HandleType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoFrame) Height() int {
	defer qt.Recovering("QVideoFrame::height")

	if ptr.Pointer() != nil {
		return int(C.QVideoFrame_Height(ptr.Pointer()))
	}
	return 0
}

func QVideoFrame_ImageFormatFromPixelFormat(format QVideoFrame__PixelFormat) gui.QImage__Format {
	defer qt.Recovering("QVideoFrame::imageFormatFromPixelFormat")

	return gui.QImage__Format(C.QVideoFrame_QVideoFrame_ImageFormatFromPixelFormat(C.int(format)))
}

func (ptr *QVideoFrame) IsMapped() bool {
	defer qt.Recovering("QVideoFrame::isMapped")

	if ptr.Pointer() != nil {
		return C.QVideoFrame_IsMapped(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVideoFrame) IsReadable() bool {
	defer qt.Recovering("QVideoFrame::isReadable")

	if ptr.Pointer() != nil {
		return C.QVideoFrame_IsReadable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVideoFrame) IsValid() bool {
	defer qt.Recovering("QVideoFrame::isValid")

	if ptr.Pointer() != nil {
		return C.QVideoFrame_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVideoFrame) IsWritable() bool {
	defer qt.Recovering("QVideoFrame::isWritable")

	if ptr.Pointer() != nil {
		return C.QVideoFrame_IsWritable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVideoFrame) Map(mode QAbstractVideoBuffer__MapMode) bool {
	defer qt.Recovering("QVideoFrame::map")

	if ptr.Pointer() != nil {
		return C.QVideoFrame_Map(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QVideoFrame) MapMode() QAbstractVideoBuffer__MapMode {
	defer qt.Recovering("QVideoFrame::mapMode")

	if ptr.Pointer() != nil {
		return QAbstractVideoBuffer__MapMode(C.QVideoFrame_MapMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoFrame) MappedBytes() int {
	defer qt.Recovering("QVideoFrame::mappedBytes")

	if ptr.Pointer() != nil {
		return int(C.QVideoFrame_MappedBytes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoFrame) MetaData(key string) *core.QVariant {
	defer qt.Recovering("QVideoFrame::metaData")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QVideoFrame_MetaData(ptr.Pointer(), C.CString(key)))
	}
	return nil
}

func (ptr *QVideoFrame) PixelFormat() QVideoFrame__PixelFormat {
	defer qt.Recovering("QVideoFrame::pixelFormat")

	if ptr.Pointer() != nil {
		return QVideoFrame__PixelFormat(C.QVideoFrame_PixelFormat(ptr.Pointer()))
	}
	return 0
}

func QVideoFrame_PixelFormatFromImageFormat(format gui.QImage__Format) QVideoFrame__PixelFormat {
	defer qt.Recovering("QVideoFrame::pixelFormatFromImageFormat")

	return QVideoFrame__PixelFormat(C.QVideoFrame_QVideoFrame_PixelFormatFromImageFormat(C.int(format)))
}

func (ptr *QVideoFrame) PlaneCount() int {
	defer qt.Recovering("QVideoFrame::planeCount")

	if ptr.Pointer() != nil {
		return int(C.QVideoFrame_PlaneCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoFrame) SetEndTime(time int64) {
	defer qt.Recovering("QVideoFrame::setEndTime")

	if ptr.Pointer() != nil {
		C.QVideoFrame_SetEndTime(ptr.Pointer(), C.longlong(time))
	}
}

func (ptr *QVideoFrame) SetFieldType(field QVideoFrame__FieldType) {
	defer qt.Recovering("QVideoFrame::setFieldType")

	if ptr.Pointer() != nil {
		C.QVideoFrame_SetFieldType(ptr.Pointer(), C.int(field))
	}
}

func (ptr *QVideoFrame) SetMetaData(key string, value core.QVariant_ITF) {
	defer qt.Recovering("QVideoFrame::setMetaData")

	if ptr.Pointer() != nil {
		C.QVideoFrame_SetMetaData(ptr.Pointer(), C.CString(key), core.PointerFromQVariant(value))
	}
}

func (ptr *QVideoFrame) SetStartTime(time int64) {
	defer qt.Recovering("QVideoFrame::setStartTime")

	if ptr.Pointer() != nil {
		C.QVideoFrame_SetStartTime(ptr.Pointer(), C.longlong(time))
	}
}

func (ptr *QVideoFrame) Size() *core.QSize {
	defer qt.Recovering("QVideoFrame::size")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QVideoFrame_Size(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVideoFrame) StartTime() int64 {
	defer qt.Recovering("QVideoFrame::startTime")

	if ptr.Pointer() != nil {
		return int64(C.QVideoFrame_StartTime(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoFrame) Unmap() {
	defer qt.Recovering("QVideoFrame::unmap")

	if ptr.Pointer() != nil {
		C.QVideoFrame_Unmap(ptr.Pointer())
	}
}

func (ptr *QVideoFrame) Width() int {
	defer qt.Recovering("QVideoFrame::width")

	if ptr.Pointer() != nil {
		return int(C.QVideoFrame_Width(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoFrame) DestroyQVideoFrame() {
	defer qt.Recovering("QVideoFrame::~QVideoFrame")

	if ptr.Pointer() != nil {
		C.QVideoFrame_DestroyQVideoFrame(ptr.Pointer())
	}
}

type QVideoProbe struct {
	core.QObject
}

type QVideoProbe_ITF interface {
	core.QObject_ITF
	QVideoProbe_PTR() *QVideoProbe
}

func PointerFromQVideoProbe(ptr QVideoProbe_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoProbe_PTR().Pointer()
	}
	return nil
}

func NewQVideoProbeFromPointer(ptr unsafe.Pointer) *QVideoProbe {
	var n = new(QVideoProbe)
	n.SetPointer(ptr)
	return n
}

func newQVideoProbeFromPointer(ptr unsafe.Pointer) *QVideoProbe {
	var n = NewQVideoProbeFromPointer(ptr)
	for len(n.ObjectName()) < len("QVideoProbe_") {
		n.SetObjectName("QVideoProbe_" + qt.Identifier())
	}
	return n
}

func (ptr *QVideoProbe) QVideoProbe_PTR() *QVideoProbe {
	return ptr
}

func NewQVideoProbe(parent core.QObject_ITF) *QVideoProbe {
	defer qt.Recovering("QVideoProbe::QVideoProbe")

	return newQVideoProbeFromPointer(C.QVideoProbe_NewQVideoProbe(core.PointerFromQObject(parent)))
}

func (ptr *QVideoProbe) ConnectFlush(f func()) {
	defer qt.Recovering("connect QVideoProbe::flush")

	if ptr.Pointer() != nil {
		C.QVideoProbe_ConnectFlush(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "flush", f)
	}
}

func (ptr *QVideoProbe) DisconnectFlush() {
	defer qt.Recovering("disconnect QVideoProbe::flush")

	if ptr.Pointer() != nil {
		C.QVideoProbe_DisconnectFlush(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "flush")
	}
}

//export callbackQVideoProbeFlush
func callbackQVideoProbeFlush(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QVideoProbe::flush")

	if signal := qt.GetSignal(C.GoString(ptrName), "flush"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVideoProbe) Flush() {
	defer qt.Recovering("QVideoProbe::flush")

	if ptr.Pointer() != nil {
		C.QVideoProbe_Flush(ptr.Pointer())
	}
}

func (ptr *QVideoProbe) IsActive() bool {
	defer qt.Recovering("QVideoProbe::isActive")

	if ptr.Pointer() != nil {
		return C.QVideoProbe_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVideoProbe) SetSource(source QMediaObject_ITF) bool {
	defer qt.Recovering("QVideoProbe::setSource")

	if ptr.Pointer() != nil {
		return C.QVideoProbe_SetSource(ptr.Pointer(), PointerFromQMediaObject(source)) != 0
	}
	return false
}

func (ptr *QVideoProbe) SetSource2(mediaRecorder QMediaRecorder_ITF) bool {
	defer qt.Recovering("QVideoProbe::setSource")

	if ptr.Pointer() != nil {
		return C.QVideoProbe_SetSource2(ptr.Pointer(), PointerFromQMediaRecorder(mediaRecorder)) != 0
	}
	return false
}

func (ptr *QVideoProbe) ConnectVideoFrameProbed(f func(frame *QVideoFrame)) {
	defer qt.Recovering("connect QVideoProbe::videoFrameProbed")

	if ptr.Pointer() != nil {
		C.QVideoProbe_ConnectVideoFrameProbed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "videoFrameProbed", f)
	}
}

func (ptr *QVideoProbe) DisconnectVideoFrameProbed() {
	defer qt.Recovering("disconnect QVideoProbe::videoFrameProbed")

	if ptr.Pointer() != nil {
		C.QVideoProbe_DisconnectVideoFrameProbed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "videoFrameProbed")
	}
}

//export callbackQVideoProbeVideoFrameProbed
func callbackQVideoProbeVideoFrameProbed(ptr unsafe.Pointer, ptrName *C.char, frame unsafe.Pointer) {
	defer qt.Recovering("callback QVideoProbe::videoFrameProbed")

	if signal := qt.GetSignal(C.GoString(ptrName), "videoFrameProbed"); signal != nil {
		signal.(func(*QVideoFrame))(NewQVideoFrameFromPointer(frame))
	}

}

func (ptr *QVideoProbe) VideoFrameProbed(frame QVideoFrame_ITF) {
	defer qt.Recovering("QVideoProbe::videoFrameProbed")

	if ptr.Pointer() != nil {
		C.QVideoProbe_VideoFrameProbed(ptr.Pointer(), PointerFromQVideoFrame(frame))
	}
}

func (ptr *QVideoProbe) DestroyQVideoProbe() {
	defer qt.Recovering("QVideoProbe::~QVideoProbe")

	if ptr.Pointer() != nil {
		C.QVideoProbe_DestroyQVideoProbe(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QVideoProbe) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QVideoProbe::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QVideoProbe) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QVideoProbe::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQVideoProbeTimerEvent
func callbackQVideoProbeTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoProbe::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQVideoProbeFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QVideoProbe) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QVideoProbe::timerEvent")

	if ptr.Pointer() != nil {
		C.QVideoProbe_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QVideoProbe) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QVideoProbe::timerEvent")

	if ptr.Pointer() != nil {
		C.QVideoProbe_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QVideoProbe) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QVideoProbe::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QVideoProbe) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QVideoProbe::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQVideoProbeChildEvent
func callbackQVideoProbeChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoProbe::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQVideoProbeFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QVideoProbe) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QVideoProbe::childEvent")

	if ptr.Pointer() != nil {
		C.QVideoProbe_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QVideoProbe) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QVideoProbe::childEvent")

	if ptr.Pointer() != nil {
		C.QVideoProbe_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QVideoProbe) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QVideoProbe::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QVideoProbe) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QVideoProbe::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQVideoProbeCustomEvent
func callbackQVideoProbeCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoProbe::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQVideoProbeFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QVideoProbe) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoProbe::customEvent")

	if ptr.Pointer() != nil {
		C.QVideoProbe_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QVideoProbe) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoProbe::customEvent")

	if ptr.Pointer() != nil {
		C.QVideoProbe_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QVideoRendererControl struct {
	QMediaControl
}

type QVideoRendererControl_ITF interface {
	QMediaControl_ITF
	QVideoRendererControl_PTR() *QVideoRendererControl
}

func PointerFromQVideoRendererControl(ptr QVideoRendererControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoRendererControl_PTR().Pointer()
	}
	return nil
}

func NewQVideoRendererControlFromPointer(ptr unsafe.Pointer) *QVideoRendererControl {
	var n = new(QVideoRendererControl)
	n.SetPointer(ptr)
	return n
}

func newQVideoRendererControlFromPointer(ptr unsafe.Pointer) *QVideoRendererControl {
	var n = NewQVideoRendererControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QVideoRendererControl_") {
		n.SetObjectName("QVideoRendererControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QVideoRendererControl) QVideoRendererControl_PTR() *QVideoRendererControl {
	return ptr
}

func (ptr *QVideoRendererControl) SetSurface(surface QAbstractVideoSurface_ITF) {
	defer qt.Recovering("QVideoRendererControl::setSurface")

	if ptr.Pointer() != nil {
		C.QVideoRendererControl_SetSurface(ptr.Pointer(), PointerFromQAbstractVideoSurface(surface))
	}
}

func (ptr *QVideoRendererControl) Surface() *QAbstractVideoSurface {
	defer qt.Recovering("QVideoRendererControl::surface")

	if ptr.Pointer() != nil {
		return NewQAbstractVideoSurfaceFromPointer(C.QVideoRendererControl_Surface(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVideoRendererControl) DestroyQVideoRendererControl() {
	defer qt.Recovering("QVideoRendererControl::~QVideoRendererControl")

	if ptr.Pointer() != nil {
		C.QVideoRendererControl_DestroyQVideoRendererControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QVideoRendererControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QVideoRendererControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QVideoRendererControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QVideoRendererControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQVideoRendererControlTimerEvent
func callbackQVideoRendererControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoRendererControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQVideoRendererControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QVideoRendererControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QVideoRendererControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QVideoRendererControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QVideoRendererControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QVideoRendererControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QVideoRendererControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QVideoRendererControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QVideoRendererControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QVideoRendererControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QVideoRendererControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQVideoRendererControlChildEvent
func callbackQVideoRendererControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoRendererControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQVideoRendererControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QVideoRendererControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QVideoRendererControl::childEvent")

	if ptr.Pointer() != nil {
		C.QVideoRendererControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QVideoRendererControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QVideoRendererControl::childEvent")

	if ptr.Pointer() != nil {
		C.QVideoRendererControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QVideoRendererControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QVideoRendererControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QVideoRendererControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QVideoRendererControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQVideoRendererControlCustomEvent
func callbackQVideoRendererControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoRendererControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQVideoRendererControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QVideoRendererControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoRendererControl::customEvent")

	if ptr.Pointer() != nil {
		C.QVideoRendererControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QVideoRendererControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoRendererControl::customEvent")

	if ptr.Pointer() != nil {
		C.QVideoRendererControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QVideoSurfaceFormat struct {
	ptr unsafe.Pointer
}

type QVideoSurfaceFormat_ITF interface {
	QVideoSurfaceFormat_PTR() *QVideoSurfaceFormat
}

func (p *QVideoSurfaceFormat) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QVideoSurfaceFormat) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQVideoSurfaceFormat(ptr QVideoSurfaceFormat_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoSurfaceFormat_PTR().Pointer()
	}
	return nil
}

func NewQVideoSurfaceFormatFromPointer(ptr unsafe.Pointer) *QVideoSurfaceFormat {
	var n = new(QVideoSurfaceFormat)
	n.SetPointer(ptr)
	return n
}

func newQVideoSurfaceFormatFromPointer(ptr unsafe.Pointer) *QVideoSurfaceFormat {
	var n = NewQVideoSurfaceFormatFromPointer(ptr)
	return n
}

func (ptr *QVideoSurfaceFormat) QVideoSurfaceFormat_PTR() *QVideoSurfaceFormat {
	return ptr
}

//QVideoSurfaceFormat::Direction
type QVideoSurfaceFormat__Direction int64

const (
	QVideoSurfaceFormat__TopToBottom = QVideoSurfaceFormat__Direction(0)
	QVideoSurfaceFormat__BottomToTop = QVideoSurfaceFormat__Direction(1)
)

//QVideoSurfaceFormat::YCbCrColorSpace
type QVideoSurfaceFormat__YCbCrColorSpace int64

const (
	QVideoSurfaceFormat__YCbCr_Undefined = QVideoSurfaceFormat__YCbCrColorSpace(0)
	QVideoSurfaceFormat__YCbCr_BT601     = QVideoSurfaceFormat__YCbCrColorSpace(1)
	QVideoSurfaceFormat__YCbCr_BT709     = QVideoSurfaceFormat__YCbCrColorSpace(2)
	QVideoSurfaceFormat__YCbCr_xvYCC601  = QVideoSurfaceFormat__YCbCrColorSpace(3)
	QVideoSurfaceFormat__YCbCr_xvYCC709  = QVideoSurfaceFormat__YCbCrColorSpace(4)
	QVideoSurfaceFormat__YCbCr_JPEG      = QVideoSurfaceFormat__YCbCrColorSpace(5)
)

func NewQVideoSurfaceFormat() *QVideoSurfaceFormat {
	defer qt.Recovering("QVideoSurfaceFormat::QVideoSurfaceFormat")

	return newQVideoSurfaceFormatFromPointer(C.QVideoSurfaceFormat_NewQVideoSurfaceFormat())
}

func NewQVideoSurfaceFormat2(size core.QSize_ITF, format QVideoFrame__PixelFormat, ty QAbstractVideoBuffer__HandleType) *QVideoSurfaceFormat {
	defer qt.Recovering("QVideoSurfaceFormat::QVideoSurfaceFormat")

	return newQVideoSurfaceFormatFromPointer(C.QVideoSurfaceFormat_NewQVideoSurfaceFormat2(core.PointerFromQSize(size), C.int(format), C.int(ty)))
}

func NewQVideoSurfaceFormat3(other QVideoSurfaceFormat_ITF) *QVideoSurfaceFormat {
	defer qt.Recovering("QVideoSurfaceFormat::QVideoSurfaceFormat")

	return newQVideoSurfaceFormatFromPointer(C.QVideoSurfaceFormat_NewQVideoSurfaceFormat3(PointerFromQVideoSurfaceFormat(other)))
}

func (ptr *QVideoSurfaceFormat) FrameHeight() int {
	defer qt.Recovering("QVideoSurfaceFormat::frameHeight")

	if ptr.Pointer() != nil {
		return int(C.QVideoSurfaceFormat_FrameHeight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoSurfaceFormat) FrameRate() float64 {
	defer qt.Recovering("QVideoSurfaceFormat::frameRate")

	if ptr.Pointer() != nil {
		return float64(C.QVideoSurfaceFormat_FrameRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoSurfaceFormat) FrameSize() *core.QSize {
	defer qt.Recovering("QVideoSurfaceFormat::frameSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QVideoSurfaceFormat_FrameSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVideoSurfaceFormat) FrameWidth() int {
	defer qt.Recovering("QVideoSurfaceFormat::frameWidth")

	if ptr.Pointer() != nil {
		return int(C.QVideoSurfaceFormat_FrameWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoSurfaceFormat) HandleType() QAbstractVideoBuffer__HandleType {
	defer qt.Recovering("QVideoSurfaceFormat::handleType")

	if ptr.Pointer() != nil {
		return QAbstractVideoBuffer__HandleType(C.QVideoSurfaceFormat_HandleType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoSurfaceFormat) IsValid() bool {
	defer qt.Recovering("QVideoSurfaceFormat::isValid")

	if ptr.Pointer() != nil {
		return C.QVideoSurfaceFormat_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVideoSurfaceFormat) PixelAspectRatio() *core.QSize {
	defer qt.Recovering("QVideoSurfaceFormat::pixelAspectRatio")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QVideoSurfaceFormat_PixelAspectRatio(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVideoSurfaceFormat) PixelFormat() QVideoFrame__PixelFormat {
	defer qt.Recovering("QVideoSurfaceFormat::pixelFormat")

	if ptr.Pointer() != nil {
		return QVideoFrame__PixelFormat(C.QVideoSurfaceFormat_PixelFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoSurfaceFormat) Property(name string) *core.QVariant {
	defer qt.Recovering("QVideoSurfaceFormat::property")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QVideoSurfaceFormat_Property(ptr.Pointer(), C.CString(name)))
	}
	return nil
}

func (ptr *QVideoSurfaceFormat) ScanLineDirection() QVideoSurfaceFormat__Direction {
	defer qt.Recovering("QVideoSurfaceFormat::scanLineDirection")

	if ptr.Pointer() != nil {
		return QVideoSurfaceFormat__Direction(C.QVideoSurfaceFormat_ScanLineDirection(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoSurfaceFormat) SetFrameRate(rate float64) {
	defer qt.Recovering("QVideoSurfaceFormat::setFrameRate")

	if ptr.Pointer() != nil {
		C.QVideoSurfaceFormat_SetFrameRate(ptr.Pointer(), C.double(rate))
	}
}

func (ptr *QVideoSurfaceFormat) SetFrameSize(size core.QSize_ITF) {
	defer qt.Recovering("QVideoSurfaceFormat::setFrameSize")

	if ptr.Pointer() != nil {
		C.QVideoSurfaceFormat_SetFrameSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QVideoSurfaceFormat) SetFrameSize2(width int, height int) {
	defer qt.Recovering("QVideoSurfaceFormat::setFrameSize")

	if ptr.Pointer() != nil {
		C.QVideoSurfaceFormat_SetFrameSize2(ptr.Pointer(), C.int(width), C.int(height))
	}
}

func (ptr *QVideoSurfaceFormat) SetPixelAspectRatio(ratio core.QSize_ITF) {
	defer qt.Recovering("QVideoSurfaceFormat::setPixelAspectRatio")

	if ptr.Pointer() != nil {
		C.QVideoSurfaceFormat_SetPixelAspectRatio(ptr.Pointer(), core.PointerFromQSize(ratio))
	}
}

func (ptr *QVideoSurfaceFormat) SetPixelAspectRatio2(horizontal int, vertical int) {
	defer qt.Recovering("QVideoSurfaceFormat::setPixelAspectRatio")

	if ptr.Pointer() != nil {
		C.QVideoSurfaceFormat_SetPixelAspectRatio2(ptr.Pointer(), C.int(horizontal), C.int(vertical))
	}
}

func (ptr *QVideoSurfaceFormat) SetProperty(name string, value core.QVariant_ITF) {
	defer qt.Recovering("QVideoSurfaceFormat::setProperty")

	if ptr.Pointer() != nil {
		C.QVideoSurfaceFormat_SetProperty(ptr.Pointer(), C.CString(name), core.PointerFromQVariant(value))
	}
}

func (ptr *QVideoSurfaceFormat) SetScanLineDirection(direction QVideoSurfaceFormat__Direction) {
	defer qt.Recovering("QVideoSurfaceFormat::setScanLineDirection")

	if ptr.Pointer() != nil {
		C.QVideoSurfaceFormat_SetScanLineDirection(ptr.Pointer(), C.int(direction))
	}
}

func (ptr *QVideoSurfaceFormat) SetViewport(viewport core.QRect_ITF) {
	defer qt.Recovering("QVideoSurfaceFormat::setViewport")

	if ptr.Pointer() != nil {
		C.QVideoSurfaceFormat_SetViewport(ptr.Pointer(), core.PointerFromQRect(viewport))
	}
}

func (ptr *QVideoSurfaceFormat) SetYCbCrColorSpace(space QVideoSurfaceFormat__YCbCrColorSpace) {
	defer qt.Recovering("QVideoSurfaceFormat::setYCbCrColorSpace")

	if ptr.Pointer() != nil {
		C.QVideoSurfaceFormat_SetYCbCrColorSpace(ptr.Pointer(), C.int(space))
	}
}

func (ptr *QVideoSurfaceFormat) SizeHint() *core.QSize {
	defer qt.Recovering("QVideoSurfaceFormat::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QVideoSurfaceFormat_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVideoSurfaceFormat) Viewport() *core.QRect {
	defer qt.Recovering("QVideoSurfaceFormat::viewport")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QVideoSurfaceFormat_Viewport(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVideoSurfaceFormat) YCbCrColorSpace() QVideoSurfaceFormat__YCbCrColorSpace {
	defer qt.Recovering("QVideoSurfaceFormat::yCbCrColorSpace")

	if ptr.Pointer() != nil {
		return QVideoSurfaceFormat__YCbCrColorSpace(C.QVideoSurfaceFormat_YCbCrColorSpace(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoSurfaceFormat) DestroyQVideoSurfaceFormat() {
	defer qt.Recovering("QVideoSurfaceFormat::~QVideoSurfaceFormat")

	if ptr.Pointer() != nil {
		C.QVideoSurfaceFormat_DestroyQVideoSurfaceFormat(ptr.Pointer())
	}
}

type QVideoWindowControl struct {
	QMediaControl
}

type QVideoWindowControl_ITF interface {
	QMediaControl_ITF
	QVideoWindowControl_PTR() *QVideoWindowControl
}

func PointerFromQVideoWindowControl(ptr QVideoWindowControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoWindowControl_PTR().Pointer()
	}
	return nil
}

func NewQVideoWindowControlFromPointer(ptr unsafe.Pointer) *QVideoWindowControl {
	var n = new(QVideoWindowControl)
	n.SetPointer(ptr)
	return n
}

func newQVideoWindowControlFromPointer(ptr unsafe.Pointer) *QVideoWindowControl {
	var n = NewQVideoWindowControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QVideoWindowControl_") {
		n.SetObjectName("QVideoWindowControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QVideoWindowControl) QVideoWindowControl_PTR() *QVideoWindowControl {
	return ptr
}

func (ptr *QVideoWindowControl) AspectRatioMode() core.Qt__AspectRatioMode {
	defer qt.Recovering("QVideoWindowControl::aspectRatioMode")

	if ptr.Pointer() != nil {
		return core.Qt__AspectRatioMode(C.QVideoWindowControl_AspectRatioMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWindowControl) Brightness() int {
	defer qt.Recovering("QVideoWindowControl::brightness")

	if ptr.Pointer() != nil {
		return int(C.QVideoWindowControl_Brightness(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWindowControl) ConnectBrightnessChanged(f func(brightness int)) {
	defer qt.Recovering("connect QVideoWindowControl::brightnessChanged")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_ConnectBrightnessChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "brightnessChanged", f)
	}
}

func (ptr *QVideoWindowControl) DisconnectBrightnessChanged() {
	defer qt.Recovering("disconnect QVideoWindowControl::brightnessChanged")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_DisconnectBrightnessChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "brightnessChanged")
	}
}

//export callbackQVideoWindowControlBrightnessChanged
func callbackQVideoWindowControlBrightnessChanged(ptr unsafe.Pointer, ptrName *C.char, brightness C.int) {
	defer qt.Recovering("callback QVideoWindowControl::brightnessChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "brightnessChanged"); signal != nil {
		signal.(func(int))(int(brightness))
	}

}

func (ptr *QVideoWindowControl) BrightnessChanged(brightness int) {
	defer qt.Recovering("QVideoWindowControl::brightnessChanged")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_BrightnessChanged(ptr.Pointer(), C.int(brightness))
	}
}

func (ptr *QVideoWindowControl) Contrast() int {
	defer qt.Recovering("QVideoWindowControl::contrast")

	if ptr.Pointer() != nil {
		return int(C.QVideoWindowControl_Contrast(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWindowControl) ConnectContrastChanged(f func(contrast int)) {
	defer qt.Recovering("connect QVideoWindowControl::contrastChanged")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_ConnectContrastChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "contrastChanged", f)
	}
}

func (ptr *QVideoWindowControl) DisconnectContrastChanged() {
	defer qt.Recovering("disconnect QVideoWindowControl::contrastChanged")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_DisconnectContrastChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "contrastChanged")
	}
}

//export callbackQVideoWindowControlContrastChanged
func callbackQVideoWindowControlContrastChanged(ptr unsafe.Pointer, ptrName *C.char, contrast C.int) {
	defer qt.Recovering("callback QVideoWindowControl::contrastChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "contrastChanged"); signal != nil {
		signal.(func(int))(int(contrast))
	}

}

func (ptr *QVideoWindowControl) ContrastChanged(contrast int) {
	defer qt.Recovering("QVideoWindowControl::contrastChanged")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_ContrastChanged(ptr.Pointer(), C.int(contrast))
	}
}

func (ptr *QVideoWindowControl) DisplayRect() *core.QRect {
	defer qt.Recovering("QVideoWindowControl::displayRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QVideoWindowControl_DisplayRect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVideoWindowControl) ConnectFullScreenChanged(f func(fullScreen bool)) {
	defer qt.Recovering("connect QVideoWindowControl::fullScreenChanged")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_ConnectFullScreenChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "fullScreenChanged", f)
	}
}

func (ptr *QVideoWindowControl) DisconnectFullScreenChanged() {
	defer qt.Recovering("disconnect QVideoWindowControl::fullScreenChanged")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_DisconnectFullScreenChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "fullScreenChanged")
	}
}

//export callbackQVideoWindowControlFullScreenChanged
func callbackQVideoWindowControlFullScreenChanged(ptr unsafe.Pointer, ptrName *C.char, fullScreen C.int) {
	defer qt.Recovering("callback QVideoWindowControl::fullScreenChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "fullScreenChanged"); signal != nil {
		signal.(func(bool))(int(fullScreen) != 0)
	}

}

func (ptr *QVideoWindowControl) FullScreenChanged(fullScreen bool) {
	defer qt.Recovering("QVideoWindowControl::fullScreenChanged")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_FullScreenChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(fullScreen)))
	}
}

func (ptr *QVideoWindowControl) Hue() int {
	defer qt.Recovering("QVideoWindowControl::hue")

	if ptr.Pointer() != nil {
		return int(C.QVideoWindowControl_Hue(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWindowControl) ConnectHueChanged(f func(hue int)) {
	defer qt.Recovering("connect QVideoWindowControl::hueChanged")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_ConnectHueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "hueChanged", f)
	}
}

func (ptr *QVideoWindowControl) DisconnectHueChanged() {
	defer qt.Recovering("disconnect QVideoWindowControl::hueChanged")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_DisconnectHueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "hueChanged")
	}
}

//export callbackQVideoWindowControlHueChanged
func callbackQVideoWindowControlHueChanged(ptr unsafe.Pointer, ptrName *C.char, hue C.int) {
	defer qt.Recovering("callback QVideoWindowControl::hueChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "hueChanged"); signal != nil {
		signal.(func(int))(int(hue))
	}

}

func (ptr *QVideoWindowControl) HueChanged(hue int) {
	defer qt.Recovering("QVideoWindowControl::hueChanged")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_HueChanged(ptr.Pointer(), C.int(hue))
	}
}

func (ptr *QVideoWindowControl) IsFullScreen() bool {
	defer qt.Recovering("QVideoWindowControl::isFullScreen")

	if ptr.Pointer() != nil {
		return C.QVideoWindowControl_IsFullScreen(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVideoWindowControl) NativeSize() *core.QSize {
	defer qt.Recovering("QVideoWindowControl::nativeSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QVideoWindowControl_NativeSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVideoWindowControl) ConnectNativeSizeChanged(f func()) {
	defer qt.Recovering("connect QVideoWindowControl::nativeSizeChanged")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_ConnectNativeSizeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "nativeSizeChanged", f)
	}
}

func (ptr *QVideoWindowControl) DisconnectNativeSizeChanged() {
	defer qt.Recovering("disconnect QVideoWindowControl::nativeSizeChanged")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_DisconnectNativeSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "nativeSizeChanged")
	}
}

//export callbackQVideoWindowControlNativeSizeChanged
func callbackQVideoWindowControlNativeSizeChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QVideoWindowControl::nativeSizeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "nativeSizeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVideoWindowControl) NativeSizeChanged() {
	defer qt.Recovering("QVideoWindowControl::nativeSizeChanged")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_NativeSizeChanged(ptr.Pointer())
	}
}

func (ptr *QVideoWindowControl) Repaint() {
	defer qt.Recovering("QVideoWindowControl::repaint")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_Repaint(ptr.Pointer())
	}
}

func (ptr *QVideoWindowControl) Saturation() int {
	defer qt.Recovering("QVideoWindowControl::saturation")

	if ptr.Pointer() != nil {
		return int(C.QVideoWindowControl_Saturation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWindowControl) ConnectSaturationChanged(f func(saturation int)) {
	defer qt.Recovering("connect QVideoWindowControl::saturationChanged")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_ConnectSaturationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "saturationChanged", f)
	}
}

func (ptr *QVideoWindowControl) DisconnectSaturationChanged() {
	defer qt.Recovering("disconnect QVideoWindowControl::saturationChanged")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_DisconnectSaturationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "saturationChanged")
	}
}

//export callbackQVideoWindowControlSaturationChanged
func callbackQVideoWindowControlSaturationChanged(ptr unsafe.Pointer, ptrName *C.char, saturation C.int) {
	defer qt.Recovering("callback QVideoWindowControl::saturationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "saturationChanged"); signal != nil {
		signal.(func(int))(int(saturation))
	}

}

func (ptr *QVideoWindowControl) SaturationChanged(saturation int) {
	defer qt.Recovering("QVideoWindowControl::saturationChanged")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_SaturationChanged(ptr.Pointer(), C.int(saturation))
	}
}

func (ptr *QVideoWindowControl) SetAspectRatioMode(mode core.Qt__AspectRatioMode) {
	defer qt.Recovering("QVideoWindowControl::setAspectRatioMode")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_SetAspectRatioMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QVideoWindowControl) SetBrightness(brightness int) {
	defer qt.Recovering("QVideoWindowControl::setBrightness")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_SetBrightness(ptr.Pointer(), C.int(brightness))
	}
}

func (ptr *QVideoWindowControl) SetContrast(contrast int) {
	defer qt.Recovering("QVideoWindowControl::setContrast")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_SetContrast(ptr.Pointer(), C.int(contrast))
	}
}

func (ptr *QVideoWindowControl) SetDisplayRect(rect core.QRect_ITF) {
	defer qt.Recovering("QVideoWindowControl::setDisplayRect")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_SetDisplayRect(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QVideoWindowControl) SetFullScreen(fullScreen bool) {
	defer qt.Recovering("QVideoWindowControl::setFullScreen")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_SetFullScreen(ptr.Pointer(), C.int(qt.GoBoolToInt(fullScreen)))
	}
}

func (ptr *QVideoWindowControl) SetHue(hue int) {
	defer qt.Recovering("QVideoWindowControl::setHue")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_SetHue(ptr.Pointer(), C.int(hue))
	}
}

func (ptr *QVideoWindowControl) SetSaturation(saturation int) {
	defer qt.Recovering("QVideoWindowControl::setSaturation")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_SetSaturation(ptr.Pointer(), C.int(saturation))
	}
}

func (ptr *QVideoWindowControl) SetWinId(id uintptr) {
	defer qt.Recovering("QVideoWindowControl::setWinId")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_SetWinId(ptr.Pointer(), C.ulonglong(id))
	}
}

func (ptr *QVideoWindowControl) WinId() uintptr {
	defer qt.Recovering("QVideoWindowControl::winId")

	if ptr.Pointer() != nil {
		return uintptr(C.QVideoWindowControl_WinId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWindowControl) DestroyQVideoWindowControl() {
	defer qt.Recovering("QVideoWindowControl::~QVideoWindowControl")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_DestroyQVideoWindowControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QVideoWindowControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QVideoWindowControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QVideoWindowControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QVideoWindowControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQVideoWindowControlTimerEvent
func callbackQVideoWindowControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWindowControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQVideoWindowControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QVideoWindowControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QVideoWindowControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QVideoWindowControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QVideoWindowControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QVideoWindowControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QVideoWindowControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QVideoWindowControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QVideoWindowControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQVideoWindowControlChildEvent
func callbackQVideoWindowControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWindowControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQVideoWindowControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QVideoWindowControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QVideoWindowControl::childEvent")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QVideoWindowControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QVideoWindowControl::childEvent")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QVideoWindowControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QVideoWindowControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QVideoWindowControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QVideoWindowControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQVideoWindowControlCustomEvent
func callbackQVideoWindowControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWindowControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQVideoWindowControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QVideoWindowControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoWindowControl::customEvent")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QVideoWindowControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoWindowControl::customEvent")

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
