package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

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

	return NewQMediaRecorderFromPointer(C.QMediaRecorder_NewQMediaRecorder(PointerFromQMediaObject(mediaObject), core.PointerFromQObject(parent)))
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
func callbackQMediaRecorderActualLocationChanged(ptrName *C.char, location unsafe.Pointer) {
	defer qt.Recovering("callback QMediaRecorder::actualLocationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "actualLocationChanged"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(location))
	}

}

func (ptr *QMediaRecorder) AudioCodecDescription(codec string) string {
	defer qt.Recovering("QMediaRecorder::audioCodecDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaRecorder_AudioCodecDescription(ptr.Pointer(), C.CString(codec)))
	}
	return ""
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
func callbackQMediaRecorderAvailabilityChanged2(ptrName *C.char, availability C.int) {
	defer qt.Recovering("callback QMediaRecorder::availabilityChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "availabilityChanged2"); signal != nil {
		signal.(func(QMultimedia__AvailabilityStatus))(QMultimedia__AvailabilityStatus(availability))
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
func callbackQMediaRecorderAvailabilityChanged(ptrName *C.char, available C.int) {
	defer qt.Recovering("callback QMediaRecorder::availabilityChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "availabilityChanged"); signal != nil {
		signal.(func(bool))(int(available) != 0)
	}

}

func (ptr *QMediaRecorder) AvailableMetaData() []string {
	defer qt.Recovering("QMediaRecorder::availableMetaData")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMediaRecorder_AvailableMetaData(ptr.Pointer())), ",,,")
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
func callbackQMediaRecorderDurationChanged(ptrName *C.char, duration C.longlong) {
	defer qt.Recovering("callback QMediaRecorder::durationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "durationChanged"); signal != nil {
		signal.(func(int64))(int64(duration))
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
func callbackQMediaRecorderError2(ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QMediaRecorder::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QMediaRecorder__Error))(QMediaRecorder__Error(error))
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
func callbackQMediaRecorderMetaDataAvailableChanged(ptrName *C.char, available C.int) {
	defer qt.Recovering("callback QMediaRecorder::metaDataAvailableChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaDataAvailableChanged"); signal != nil {
		signal.(func(bool))(int(available) != 0)
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
func callbackQMediaRecorderMetaDataChanged(ptrName *C.char) {
	defer qt.Recovering("callback QMediaRecorder::metaDataChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaDataChanged"); signal != nil {
		signal.(func())()
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
func callbackQMediaRecorderMetaDataChanged2(ptrName *C.char, key *C.char, value unsafe.Pointer) {
	defer qt.Recovering("callback QMediaRecorder::metaDataChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaDataChanged2"); signal != nil {
		signal.(func(string, *core.QVariant))(C.GoString(key), core.NewQVariantFromPointer(value))
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
func callbackQMediaRecorderMetaDataWritableChanged(ptrName *C.char, writable C.int) {
	defer qt.Recovering("callback QMediaRecorder::metaDataWritableChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaDataWritableChanged"); signal != nil {
		signal.(func(bool))(int(writable) != 0)
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
func callbackQMediaRecorderMutedChanged(ptrName *C.char, muted C.int) {
	defer qt.Recovering("callback QMediaRecorder::mutedChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "mutedChanged"); signal != nil {
		signal.(func(bool))(int(muted) != 0)
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
func callbackQMediaRecorderStateChanged(ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QMediaRecorder::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QMediaRecorder__State))(QMediaRecorder__State(state))
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
func callbackQMediaRecorderStatusChanged(ptrName *C.char, status C.int) {
	defer qt.Recovering("callback QMediaRecorder::statusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "statusChanged"); signal != nil {
		signal.(func(QMediaRecorder__Status))(QMediaRecorder__Status(status))
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
		return strings.Split(C.GoString(C.QMediaRecorder_SupportedAudioCodecs(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QMediaRecorder) SupportedContainers() []string {
	defer qt.Recovering("QMediaRecorder::supportedContainers")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMediaRecorder_SupportedContainers(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QMediaRecorder) SupportedVideoCodecs() []string {
	defer qt.Recovering("QMediaRecorder::supportedVideoCodecs")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMediaRecorder_SupportedVideoCodecs(ptr.Pointer())), ",,,")
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
func callbackQMediaRecorderVolumeChanged(ptrName *C.char, volume C.double) {
	defer qt.Recovering("callback QMediaRecorder::volumeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "volumeChanged"); signal != nil {
		signal.(func(float64))(float64(volume))
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
func callbackQMediaRecorderTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMediaRecorder::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQMediaRecorderChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMediaRecorder::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQMediaRecorderCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMediaRecorder::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
