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

type QMediaRecorderITF interface {
	core.QObjectITF
	QMediaBindableInterfaceITF
	QMediaRecorderPTR() *QMediaRecorder
}

func (p *QMediaRecorder) Pointer() unsafe.Pointer {
	return p.QObjectPTR().Pointer()
}

func (p *QMediaRecorder) SetPointer(ptr unsafe.Pointer) {
	p.QObjectPTR().SetPointer(ptr)
	p.QMediaBindableInterfacePTR().SetPointer(ptr)
}

func PointerFromQMediaRecorder(ptr QMediaRecorderITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaRecorderPTR().Pointer()
	}
	return nil
}

func QMediaRecorderFromPointer(ptr unsafe.Pointer) *QMediaRecorder {
	var n = new(QMediaRecorder)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMediaRecorder_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaRecorder) QMediaRecorderPTR() *QMediaRecorder {
	return ptr
}

//QMediaRecorder::Error
type QMediaRecorder__Error int

var (
	QMediaRecorder__NoError         = QMediaRecorder__Error(0)
	QMediaRecorder__ResourceError   = QMediaRecorder__Error(1)
	QMediaRecorder__FormatError     = QMediaRecorder__Error(2)
	QMediaRecorder__OutOfSpaceError = QMediaRecorder__Error(3)
)

//QMediaRecorder::State
type QMediaRecorder__State int

var (
	QMediaRecorder__StoppedState   = QMediaRecorder__State(0)
	QMediaRecorder__RecordingState = QMediaRecorder__State(1)
	QMediaRecorder__PausedState    = QMediaRecorder__State(2)
)

//QMediaRecorder::Status
type QMediaRecorder__Status int

var (
	QMediaRecorder__UnavailableStatus = QMediaRecorder__Status(0)
	QMediaRecorder__UnloadedStatus    = QMediaRecorder__Status(1)
	QMediaRecorder__LoadingStatus     = QMediaRecorder__Status(2)
	QMediaRecorder__LoadedStatus      = QMediaRecorder__Status(3)
	QMediaRecorder__StartingStatus    = QMediaRecorder__Status(4)
	QMediaRecorder__RecordingStatus   = QMediaRecorder__Status(5)
	QMediaRecorder__PausedStatus      = QMediaRecorder__Status(6)
	QMediaRecorder__FinalizingStatus  = QMediaRecorder__Status(7)
)

func (ptr *QMediaRecorder) ActualLocation() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaRecorder_ActualLocation(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QMediaRecorder) IsMetaDataAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QMediaRecorder_IsMetaDataAvailable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMediaRecorder) IsMetaDataWritable() bool {
	if ptr.Pointer() != nil {
		return C.QMediaRecorder_IsMetaDataWritable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMediaRecorder) IsMuted() bool {
	if ptr.Pointer() != nil {
		return C.QMediaRecorder_IsMuted(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMediaRecorder) OutputLocation() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaRecorder_OutputLocation(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QMediaRecorder) SetMuted(muted bool) {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_SetMuted(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(muted)))
	}
}

func (ptr *QMediaRecorder) SetOutputLocation(location string) bool {
	if ptr.Pointer() != nil {
		return C.QMediaRecorder_SetOutputLocation(C.QtObjectPtr(ptr.Pointer()), C.CString(location)) != 0
	}
	return false
}

func NewQMediaRecorder(mediaObject QMediaObjectITF, parent core.QObjectITF) *QMediaRecorder {
	return QMediaRecorderFromPointer(unsafe.Pointer(C.QMediaRecorder_NewQMediaRecorder(C.QtObjectPtr(PointerFromQMediaObject(mediaObject)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QMediaRecorder) ConnectActualLocationChanged(f func(location string)) {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_ConnectActualLocationChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "actualLocationChanged", f)
	}
}

func (ptr *QMediaRecorder) DisconnectActualLocationChanged() {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_DisconnectActualLocationChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "actualLocationChanged")
	}
}

//export callbackQMediaRecorderActualLocationChanged
func callbackQMediaRecorderActualLocationChanged(ptrName *C.char, location *C.char) {
	qt.GetSignal(C.GoString(ptrName), "actualLocationChanged").(func(string))(C.GoString(location))
}

func (ptr *QMediaRecorder) AudioCodecDescription(codec string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaRecorder_AudioCodecDescription(C.QtObjectPtr(ptr.Pointer()), C.CString(codec)))
	}
	return ""
}

func (ptr *QMediaRecorder) ConnectAvailabilityChanged(f func(available bool)) {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_ConnectAvailabilityChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "availabilityChanged", f)
	}
}

func (ptr *QMediaRecorder) DisconnectAvailabilityChanged() {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_DisconnectAvailabilityChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "availabilityChanged")
	}
}

//export callbackQMediaRecorderAvailabilityChanged
func callbackQMediaRecorderAvailabilityChanged(ptrName *C.char, available C.int) {
	qt.GetSignal(C.GoString(ptrName), "availabilityChanged").(func(bool))(int(available) != 0)
}

func (ptr *QMediaRecorder) AvailableMetaData() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMediaRecorder_AvailableMetaData(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QMediaRecorder) ContainerDescription(format string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaRecorder_ContainerDescription(C.QtObjectPtr(ptr.Pointer()), C.CString(format)))
	}
	return ""
}

func (ptr *QMediaRecorder) ContainerFormat() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaRecorder_ContainerFormat(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QMediaRecorder) Error() QMediaRecorder__Error {
	if ptr.Pointer() != nil {
		return QMediaRecorder__Error(C.QMediaRecorder_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMediaRecorder) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaRecorder_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QMediaRecorder) IsAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QMediaRecorder_IsAvailable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMediaRecorder) MediaObject() *QMediaObject {
	if ptr.Pointer() != nil {
		return QMediaObjectFromPointer(unsafe.Pointer(C.QMediaRecorder_MediaObject(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QMediaRecorder) MetaData(key string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaRecorder_MetaData(C.QtObjectPtr(ptr.Pointer()), C.CString(key)))
	}
	return ""
}

func (ptr *QMediaRecorder) ConnectMetaDataAvailableChanged(f func(available bool)) {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_ConnectMetaDataAvailableChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "metaDataAvailableChanged", f)
	}
}

func (ptr *QMediaRecorder) DisconnectMetaDataAvailableChanged() {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_DisconnectMetaDataAvailableChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataAvailableChanged")
	}
}

//export callbackQMediaRecorderMetaDataAvailableChanged
func callbackQMediaRecorderMetaDataAvailableChanged(ptrName *C.char, available C.int) {
	qt.GetSignal(C.GoString(ptrName), "metaDataAvailableChanged").(func(bool))(int(available) != 0)
}

func (ptr *QMediaRecorder) ConnectMetaDataChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_ConnectMetaDataChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "metaDataChanged", f)
	}
}

func (ptr *QMediaRecorder) DisconnectMetaDataChanged() {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_DisconnectMetaDataChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataChanged")
	}
}

//export callbackQMediaRecorderMetaDataChanged
func callbackQMediaRecorderMetaDataChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "metaDataChanged").(func())()
}

func (ptr *QMediaRecorder) ConnectMetaDataWritableChanged(f func(writable bool)) {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_ConnectMetaDataWritableChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "metaDataWritableChanged", f)
	}
}

func (ptr *QMediaRecorder) DisconnectMetaDataWritableChanged() {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_DisconnectMetaDataWritableChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataWritableChanged")
	}
}

//export callbackQMediaRecorderMetaDataWritableChanged
func callbackQMediaRecorderMetaDataWritableChanged(ptrName *C.char, writable C.int) {
	qt.GetSignal(C.GoString(ptrName), "metaDataWritableChanged").(func(bool))(int(writable) != 0)
}

func (ptr *QMediaRecorder) ConnectMutedChanged(f func(muted bool)) {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_ConnectMutedChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "mutedChanged", f)
	}
}

func (ptr *QMediaRecorder) DisconnectMutedChanged() {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_DisconnectMutedChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "mutedChanged")
	}
}

//export callbackQMediaRecorderMutedChanged
func callbackQMediaRecorderMutedChanged(ptrName *C.char, muted C.int) {
	qt.GetSignal(C.GoString(ptrName), "mutedChanged").(func(bool))(int(muted) != 0)
}

func (ptr *QMediaRecorder) Pause() {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_Pause(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMediaRecorder) Record() {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_Record(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMediaRecorder) SetAudioSettings(settings QAudioEncoderSettingsITF) {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_SetAudioSettings(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAudioEncoderSettings(settings)))
	}
}

func (ptr *QMediaRecorder) SetContainerFormat(container string) {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_SetContainerFormat(C.QtObjectPtr(ptr.Pointer()), C.CString(container))
	}
}

func (ptr *QMediaRecorder) SetEncodingSettings(audio QAudioEncoderSettingsITF, video QVideoEncoderSettingsITF, container string) {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_SetEncodingSettings(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAudioEncoderSettings(audio)), C.QtObjectPtr(PointerFromQVideoEncoderSettings(video)), C.CString(container))
	}
}

func (ptr *QMediaRecorder) SetMetaData(key string, value string) {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_SetMetaData(C.QtObjectPtr(ptr.Pointer()), C.CString(key), C.CString(value))
	}
}

func (ptr *QMediaRecorder) SetVideoSettings(settings QVideoEncoderSettingsITF) {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_SetVideoSettings(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQVideoEncoderSettings(settings)))
	}
}

func (ptr *QMediaRecorder) ConnectStateChanged(f func(state QMediaRecorder__State)) {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_ConnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QMediaRecorder) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_DisconnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQMediaRecorderStateChanged
func callbackQMediaRecorderStateChanged(ptrName *C.char, state C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QMediaRecorder__State))(QMediaRecorder__State(state))
}

func (ptr *QMediaRecorder) Status() QMediaRecorder__Status {
	if ptr.Pointer() != nil {
		return QMediaRecorder__Status(C.QMediaRecorder_Status(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMediaRecorder) ConnectStatusChanged(f func(status QMediaRecorder__Status)) {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_ConnectStatusChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "statusChanged", f)
	}
}

func (ptr *QMediaRecorder) DisconnectStatusChanged() {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_DisconnectStatusChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "statusChanged")
	}
}

//export callbackQMediaRecorderStatusChanged
func callbackQMediaRecorderStatusChanged(ptrName *C.char, status C.int) {
	qt.GetSignal(C.GoString(ptrName), "statusChanged").(func(QMediaRecorder__Status))(QMediaRecorder__Status(status))
}

func (ptr *QMediaRecorder) Stop() {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_Stop(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMediaRecorder) SupportedAudioCodecs() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMediaRecorder_SupportedAudioCodecs(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QMediaRecorder) SupportedContainers() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMediaRecorder_SupportedContainers(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QMediaRecorder) SupportedVideoCodecs() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMediaRecorder_SupportedVideoCodecs(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QMediaRecorder) VideoCodecDescription(codec string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaRecorder_VideoCodecDescription(C.QtObjectPtr(ptr.Pointer()), C.CString(codec)))
	}
	return ""
}

func (ptr *QMediaRecorder) DestroyQMediaRecorder() {
	if ptr.Pointer() != nil {
		C.QMediaRecorder_DestroyQMediaRecorder(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
