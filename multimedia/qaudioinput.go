package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

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

	return NewQAudioInputFromPointer(C.QAudioInput_NewQAudioInput2(PointerFromQAudioDeviceInfo(audioDevice), PointerFromQAudioFormat(format), core.PointerFromQObject(parent)))
}

func NewQAudioInput(format QAudioFormat_ITF, parent core.QObject_ITF) *QAudioInput {
	defer qt.Recovering("QAudioInput::QAudioInput")

	return NewQAudioInputFromPointer(C.QAudioInput_NewQAudioInput(PointerFromQAudioFormat(format), core.PointerFromQObject(parent)))
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
func callbackQAudioInputNotify(ptrName *C.char) {
	defer qt.Recovering("callback QAudioInput::notify")

	if signal := qt.GetSignal(C.GoString(ptrName), "notify"); signal != nil {
		signal.(func())()
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
func callbackQAudioInputStateChanged(ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QAudioInput::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QAudio__State))(QAudio__State(state))
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
func callbackQAudioInputTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAudioInput::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQAudioInputChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAudioInput::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQAudioInputCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAudioInput::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
