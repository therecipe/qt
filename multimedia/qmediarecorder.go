package multimedia

//#include "qmediarecorder.h"
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
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QMediaRecorder_" + qt.RandomIdentifier())
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

func (ptr *QMediaRecorder) IsMetaDataAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QMediaRecorder_IsMetaDataAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaRecorder) IsMetaDataWritable() bool {
	if ptr.Pointer() != nil {
		return C.QMediaRecorder_IsMetaDataWritable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaRecorder) IsMuted() bool {
	if ptr.Pointer() != nil {
		return C.QMediaRecorder_IsMuted(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaRecorder) SetMuted(muted bool) {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_SetMuted(ptr.Pointer(), C.int(qt.GoBoolToInt(muted)))
	}
}

func (ptr *QMediaRecorder) SetOutputLocation(location core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QMediaRecorder_SetOutputLocation(ptr.Pointer(), core.PointerFromQUrl(location)) != 0
	}
	return false
}

func (ptr *QMediaRecorder) SetVolume(volume float64) {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_SetVolume(ptr.Pointer(), C.double(volume))
	}
}

func (ptr *QMediaRecorder) Volume() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QMediaRecorder_Volume(ptr.Pointer()))
	}
	return 0
}

func NewQMediaRecorder(mediaObject QMediaObject_ITF, parent core.QObject_ITF) *QMediaRecorder {
	return NewQMediaRecorderFromPointer(C.QMediaRecorder_NewQMediaRecorder(PointerFromQMediaObject(mediaObject), core.PointerFromQObject(parent)))
}

func (ptr *QMediaRecorder) AudioCodecDescription(codec string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaRecorder_AudioCodecDescription(ptr.Pointer(), C.CString(codec)))
	}
	return ""
}

func (ptr *QMediaRecorder) ConnectAvailabilityChanged(f func(available bool)) {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_ConnectAvailabilityChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "availabilityChanged", f)
	}
}

func (ptr *QMediaRecorder) DisconnectAvailabilityChanged() {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_DisconnectAvailabilityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "availabilityChanged")
	}
}

//export callbackQMediaRecorderAvailabilityChanged
func callbackQMediaRecorderAvailabilityChanged(ptrName *C.char, available C.int) {
	qt.GetSignal(C.GoString(ptrName), "availabilityChanged").(func(bool))(int(available) != 0)
}

func (ptr *QMediaRecorder) AvailableMetaData() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMediaRecorder_AvailableMetaData(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QMediaRecorder) ContainerDescription(format string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaRecorder_ContainerDescription(ptr.Pointer(), C.CString(format)))
	}
	return ""
}

func (ptr *QMediaRecorder) ContainerFormat() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaRecorder_ContainerFormat(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaRecorder) Error() QMediaRecorder__Error {
	if ptr.Pointer() != nil {
		return QMediaRecorder__Error(C.QMediaRecorder_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaRecorder) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaRecorder_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaRecorder) IsAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QMediaRecorder_IsAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaRecorder) MediaObject() *QMediaObject {
	if ptr.Pointer() != nil {
		return NewQMediaObjectFromPointer(C.QMediaRecorder_MediaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaRecorder) MetaData(key string) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QMediaRecorder_MetaData(ptr.Pointer(), C.CString(key)))
	}
	return nil
}

func (ptr *QMediaRecorder) ConnectMetaDataAvailableChanged(f func(available bool)) {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_ConnectMetaDataAvailableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataAvailableChanged", f)
	}
}

func (ptr *QMediaRecorder) DisconnectMetaDataAvailableChanged() {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_DisconnectMetaDataAvailableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataAvailableChanged")
	}
}

//export callbackQMediaRecorderMetaDataAvailableChanged
func callbackQMediaRecorderMetaDataAvailableChanged(ptrName *C.char, available C.int) {
	qt.GetSignal(C.GoString(ptrName), "metaDataAvailableChanged").(func(bool))(int(available) != 0)
}

func (ptr *QMediaRecorder) ConnectMetaDataChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_ConnectMetaDataChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataChanged", f)
	}
}

func (ptr *QMediaRecorder) DisconnectMetaDataChanged() {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_DisconnectMetaDataChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataChanged")
	}
}

//export callbackQMediaRecorderMetaDataChanged
func callbackQMediaRecorderMetaDataChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "metaDataChanged").(func())()
}

func (ptr *QMediaRecorder) ConnectMetaDataWritableChanged(f func(writable bool)) {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_ConnectMetaDataWritableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataWritableChanged", f)
	}
}

func (ptr *QMediaRecorder) DisconnectMetaDataWritableChanged() {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_DisconnectMetaDataWritableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataWritableChanged")
	}
}

//export callbackQMediaRecorderMetaDataWritableChanged
func callbackQMediaRecorderMetaDataWritableChanged(ptrName *C.char, writable C.int) {
	qt.GetSignal(C.GoString(ptrName), "metaDataWritableChanged").(func(bool))(int(writable) != 0)
}

func (ptr *QMediaRecorder) ConnectMutedChanged(f func(muted bool)) {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_ConnectMutedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mutedChanged", f)
	}
}

func (ptr *QMediaRecorder) DisconnectMutedChanged() {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_DisconnectMutedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mutedChanged")
	}
}

//export callbackQMediaRecorderMutedChanged
func callbackQMediaRecorderMutedChanged(ptrName *C.char, muted C.int) {
	qt.GetSignal(C.GoString(ptrName), "mutedChanged").(func(bool))(int(muted) != 0)
}

func (ptr *QMediaRecorder) Pause() {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_Pause(ptr.Pointer())
	}
}

func (ptr *QMediaRecorder) Record() {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_Record(ptr.Pointer())
	}
}

func (ptr *QMediaRecorder) SetAudioSettings(settings QAudioEncoderSettings_ITF) {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_SetAudioSettings(ptr.Pointer(), PointerFromQAudioEncoderSettings(settings))
	}
}

func (ptr *QMediaRecorder) SetContainerFormat(container string) {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_SetContainerFormat(ptr.Pointer(), C.CString(container))
	}
}

func (ptr *QMediaRecorder) SetEncodingSettings(audio QAudioEncoderSettings_ITF, video QVideoEncoderSettings_ITF, container string) {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_SetEncodingSettings(ptr.Pointer(), PointerFromQAudioEncoderSettings(audio), PointerFromQVideoEncoderSettings(video), C.CString(container))
	}
}

func (ptr *QMediaRecorder) SetMetaData(key string, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_SetMetaData(ptr.Pointer(), C.CString(key), core.PointerFromQVariant(value))
	}
}

func (ptr *QMediaRecorder) SetVideoSettings(settings QVideoEncoderSettings_ITF) {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_SetVideoSettings(ptr.Pointer(), PointerFromQVideoEncoderSettings(settings))
	}
}

func (ptr *QMediaRecorder) ConnectStateChanged(f func(state QMediaRecorder__State)) {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QMediaRecorder) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQMediaRecorderStateChanged
func callbackQMediaRecorderStateChanged(ptrName *C.char, state C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QMediaRecorder__State))(QMediaRecorder__State(state))
}

func (ptr *QMediaRecorder) Status() QMediaRecorder__Status {
	if ptr.Pointer() != nil {
		return QMediaRecorder__Status(C.QMediaRecorder_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaRecorder) ConnectStatusChanged(f func(status QMediaRecorder__Status)) {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_ConnectStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "statusChanged", f)
	}
}

func (ptr *QMediaRecorder) DisconnectStatusChanged() {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "statusChanged")
	}
}

//export callbackQMediaRecorderStatusChanged
func callbackQMediaRecorderStatusChanged(ptrName *C.char, status C.int) {
	qt.GetSignal(C.GoString(ptrName), "statusChanged").(func(QMediaRecorder__Status))(QMediaRecorder__Status(status))
}

func (ptr *QMediaRecorder) Stop() {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_Stop(ptr.Pointer())
	}
}

func (ptr *QMediaRecorder) SupportedAudioCodecs() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMediaRecorder_SupportedAudioCodecs(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QMediaRecorder) SupportedContainers() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMediaRecorder_SupportedContainers(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QMediaRecorder) SupportedVideoCodecs() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMediaRecorder_SupportedVideoCodecs(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QMediaRecorder) VideoCodecDescription(codec string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaRecorder_VideoCodecDescription(ptr.Pointer(), C.CString(codec)))
	}
	return ""
}

func (ptr *QMediaRecorder) DestroyQMediaRecorder() {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_DestroyQMediaRecorder(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
