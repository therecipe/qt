package multimedia

//#include "qradiotuner.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QRadioTuner struct {
	QMediaObject
}

type QRadioTunerITF interface {
	QMediaObjectITF
	QRadioTunerPTR() *QRadioTuner
}

func PointerFromQRadioTuner(ptr QRadioTunerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRadioTunerPTR().Pointer()
	}
	return nil
}

func QRadioTunerFromPointer(ptr unsafe.Pointer) *QRadioTuner {
	var n = new(QRadioTuner)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QRadioTuner_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QRadioTuner) QRadioTunerPTR() *QRadioTuner {
	return ptr
}

//QRadioTuner::Band
type QRadioTuner__Band int

var (
	QRadioTuner__AM  = QRadioTuner__Band(0)
	QRadioTuner__FM  = QRadioTuner__Band(1)
	QRadioTuner__SW  = QRadioTuner__Band(2)
	QRadioTuner__LW  = QRadioTuner__Band(3)
	QRadioTuner__FM2 = QRadioTuner__Band(4)
)

//QRadioTuner::Error
type QRadioTuner__Error int

var (
	QRadioTuner__NoError         = QRadioTuner__Error(0)
	QRadioTuner__ResourceError   = QRadioTuner__Error(1)
	QRadioTuner__OpenError       = QRadioTuner__Error(2)
	QRadioTuner__OutOfRangeError = QRadioTuner__Error(3)
)

//QRadioTuner::SearchMode
type QRadioTuner__SearchMode int

var (
	QRadioTuner__SearchFast         = QRadioTuner__SearchMode(0)
	QRadioTuner__SearchGetStationId = QRadioTuner__SearchMode(1)
)

//QRadioTuner::State
type QRadioTuner__State int

var (
	QRadioTuner__ActiveState  = QRadioTuner__State(0)
	QRadioTuner__StoppedState = QRadioTuner__State(1)
)

//QRadioTuner::StereoMode
type QRadioTuner__StereoMode int

var (
	QRadioTuner__ForceStereo = QRadioTuner__StereoMode(0)
	QRadioTuner__ForceMono   = QRadioTuner__StereoMode(1)
	QRadioTuner__Auto        = QRadioTuner__StereoMode(2)
)

func (ptr *QRadioTuner) Band() QRadioTuner__Band {
	if ptr.Pointer() != nil {
		return QRadioTuner__Band(C.QRadioTuner_Band(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRadioTuner) Frequency() int {
	if ptr.Pointer() != nil {
		return int(C.QRadioTuner_Frequency(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRadioTuner) IsAntennaConnected() bool {
	if ptr.Pointer() != nil {
		return C.QRadioTuner_IsAntennaConnected(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRadioTuner) IsMuted() bool {
	if ptr.Pointer() != nil {
		return C.QRadioTuner_IsMuted(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRadioTuner) IsSearching() bool {
	if ptr.Pointer() != nil {
		return C.QRadioTuner_IsSearching(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRadioTuner) IsStereo() bool {
	if ptr.Pointer() != nil {
		return C.QRadioTuner_IsStereo(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRadioTuner) RadioData() *QRadioData {
	if ptr.Pointer() != nil {
		return QRadioDataFromPointer(unsafe.Pointer(C.QRadioTuner_RadioData(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QRadioTuner) SetMuted(muted bool) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_SetMuted(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(muted)))
	}
}

func (ptr *QRadioTuner) SetStereoMode(mode QRadioTuner__StereoMode) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_SetStereoMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QRadioTuner) SetVolume(volume int) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_SetVolume(C.QtObjectPtr(ptr.Pointer()), C.int(volume))
	}
}

func (ptr *QRadioTuner) SignalStrength() int {
	if ptr.Pointer() != nil {
		return int(C.QRadioTuner_SignalStrength(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRadioTuner) StereoMode() QRadioTuner__StereoMode {
	if ptr.Pointer() != nil {
		return QRadioTuner__StereoMode(C.QRadioTuner_StereoMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRadioTuner) Volume() int {
	if ptr.Pointer() != nil {
		return int(C.QRadioTuner_Volume(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQRadioTuner(parent core.QObjectITF) *QRadioTuner {
	return QRadioTunerFromPointer(unsafe.Pointer(C.QRadioTuner_NewQRadioTuner(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QRadioTuner) ConnectAntennaConnectedChanged(f func(connectionStatus bool)) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectAntennaConnectedChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "antennaConnectedChanged", f)
	}
}

func (ptr *QRadioTuner) DisconnectAntennaConnectedChanged() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectAntennaConnectedChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "antennaConnectedChanged")
	}
}

//export callbackQRadioTunerAntennaConnectedChanged
func callbackQRadioTunerAntennaConnectedChanged(ptrName *C.char, connectionStatus C.int) {
	qt.GetSignal(C.GoString(ptrName), "antennaConnectedChanged").(func(bool))(int(connectionStatus) != 0)
}

func (ptr *QRadioTuner) ConnectBandChanged(f func(band QRadioTuner__Band)) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectBandChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "bandChanged", f)
	}
}

func (ptr *QRadioTuner) DisconnectBandChanged() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectBandChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "bandChanged")
	}
}

//export callbackQRadioTunerBandChanged
func callbackQRadioTunerBandChanged(ptrName *C.char, band C.int) {
	qt.GetSignal(C.GoString(ptrName), "bandChanged").(func(QRadioTuner__Band))(QRadioTuner__Band(band))
}

func (ptr *QRadioTuner) CancelSearch() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_CancelSearch(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QRadioTuner) Error() QRadioTuner__Error {
	if ptr.Pointer() != nil {
		return QRadioTuner__Error(C.QRadioTuner_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRadioTuner) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioTuner_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QRadioTuner) ConnectFrequencyChanged(f func(frequency int)) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectFrequencyChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "frequencyChanged", f)
	}
}

func (ptr *QRadioTuner) DisconnectFrequencyChanged() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectFrequencyChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "frequencyChanged")
	}
}

//export callbackQRadioTunerFrequencyChanged
func callbackQRadioTunerFrequencyChanged(ptrName *C.char, frequency C.int) {
	qt.GetSignal(C.GoString(ptrName), "frequencyChanged").(func(int))(int(frequency))
}

func (ptr *QRadioTuner) FrequencyStep(band QRadioTuner__Band) int {
	if ptr.Pointer() != nil {
		return int(C.QRadioTuner_FrequencyStep(C.QtObjectPtr(ptr.Pointer()), C.int(band)))
	}
	return 0
}

func (ptr *QRadioTuner) IsBandSupported(band QRadioTuner__Band) bool {
	if ptr.Pointer() != nil {
		return C.QRadioTuner_IsBandSupported(C.QtObjectPtr(ptr.Pointer()), C.int(band)) != 0
	}
	return false
}

func (ptr *QRadioTuner) ConnectMutedChanged(f func(muted bool)) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectMutedChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "mutedChanged", f)
	}
}

func (ptr *QRadioTuner) DisconnectMutedChanged() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectMutedChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "mutedChanged")
	}
}

//export callbackQRadioTunerMutedChanged
func callbackQRadioTunerMutedChanged(ptrName *C.char, muted C.int) {
	qt.GetSignal(C.GoString(ptrName), "mutedChanged").(func(bool))(int(muted) != 0)
}

func (ptr *QRadioTuner) SearchAllStations(searchMode QRadioTuner__SearchMode) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_SearchAllStations(C.QtObjectPtr(ptr.Pointer()), C.int(searchMode))
	}
}

func (ptr *QRadioTuner) SearchBackward() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_SearchBackward(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QRadioTuner) SearchForward() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_SearchForward(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QRadioTuner) ConnectSearchingChanged(f func(searching bool)) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectSearchingChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "searchingChanged", f)
	}
}

func (ptr *QRadioTuner) DisconnectSearchingChanged() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectSearchingChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "searchingChanged")
	}
}

//export callbackQRadioTunerSearchingChanged
func callbackQRadioTunerSearchingChanged(ptrName *C.char, searching C.int) {
	qt.GetSignal(C.GoString(ptrName), "searchingChanged").(func(bool))(int(searching) != 0)
}

func (ptr *QRadioTuner) SetBand(band QRadioTuner__Band) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_SetBand(C.QtObjectPtr(ptr.Pointer()), C.int(band))
	}
}

func (ptr *QRadioTuner) SetFrequency(frequency int) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_SetFrequency(C.QtObjectPtr(ptr.Pointer()), C.int(frequency))
	}
}

func (ptr *QRadioTuner) ConnectSignalStrengthChanged(f func(strength int)) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectSignalStrengthChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "signalStrengthChanged", f)
	}
}

func (ptr *QRadioTuner) DisconnectSignalStrengthChanged() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectSignalStrengthChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "signalStrengthChanged")
	}
}

//export callbackQRadioTunerSignalStrengthChanged
func callbackQRadioTunerSignalStrengthChanged(ptrName *C.char, strength C.int) {
	qt.GetSignal(C.GoString(ptrName), "signalStrengthChanged").(func(int))(int(strength))
}

func (ptr *QRadioTuner) Start() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_Start(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QRadioTuner) ConnectStateChanged(f func(state QRadioTuner__State)) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QRadioTuner) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQRadioTunerStateChanged
func callbackQRadioTunerStateChanged(ptrName *C.char, state C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QRadioTuner__State))(QRadioTuner__State(state))
}

func (ptr *QRadioTuner) ConnectStationFound(f func(frequency int, stationId string)) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectStationFound(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "stationFound", f)
	}
}

func (ptr *QRadioTuner) DisconnectStationFound() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectStationFound(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "stationFound")
	}
}

//export callbackQRadioTunerStationFound
func callbackQRadioTunerStationFound(ptrName *C.char, frequency C.int, stationId *C.char) {
	qt.GetSignal(C.GoString(ptrName), "stationFound").(func(int, string))(int(frequency), C.GoString(stationId))
}

func (ptr *QRadioTuner) ConnectStereoStatusChanged(f func(stereo bool)) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectStereoStatusChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "stereoStatusChanged", f)
	}
}

func (ptr *QRadioTuner) DisconnectStereoStatusChanged() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectStereoStatusChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "stereoStatusChanged")
	}
}

//export callbackQRadioTunerStereoStatusChanged
func callbackQRadioTunerStereoStatusChanged(ptrName *C.char, stereo C.int) {
	qt.GetSignal(C.GoString(ptrName), "stereoStatusChanged").(func(bool))(int(stereo) != 0)
}

func (ptr *QRadioTuner) Stop() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_Stop(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QRadioTuner) ConnectVolumeChanged(f func(volume int)) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectVolumeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "volumeChanged", f)
	}
}

func (ptr *QRadioTuner) DisconnectVolumeChanged() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectVolumeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "volumeChanged")
	}
}

//export callbackQRadioTunerVolumeChanged
func callbackQRadioTunerVolumeChanged(ptrName *C.char, volume C.int) {
	qt.GetSignal(C.GoString(ptrName), "volumeChanged").(func(int))(int(volume))
}

func (ptr *QRadioTuner) DestroyQRadioTuner() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_DestroyQRadioTuner(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
