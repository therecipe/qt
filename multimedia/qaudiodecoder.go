package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

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

	return NewQAudioDecoderFromPointer(C.QAudioDecoder_NewQAudioDecoder(core.PointerFromQObject(parent)))
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

func QAudioDecoder_HasSupport(mimeType string, codecs []string) QMultimedia__SupportEstimate {
	defer qt.Recovering("QAudioDecoder::hasSupport")

	return QMultimedia__SupportEstimate(C.QAudioDecoder_QAudioDecoder_HasSupport(C.CString(mimeType), C.CString(strings.Join(codecs, ",,,"))))
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
