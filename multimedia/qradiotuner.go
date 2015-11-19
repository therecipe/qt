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
	if n.ObjectName() == "" {
		n.SetObjectName("QRadioTuner_" + qt.RandomIdentifier())
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
	if ptr.Pointer() != nil {
		return QRadioTuner__Band(C.QRadioTuner_Band(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioTuner) Frequency() int {
	if ptr.Pointer() != nil {
		return int(C.QRadioTuner_Frequency(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioTuner) IsAntennaConnected() bool {
	if ptr.Pointer() != nil {
		return C.QRadioTuner_IsAntennaConnected(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRadioTuner) IsMuted() bool {
	if ptr.Pointer() != nil {
		return C.QRadioTuner_IsMuted(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRadioTuner) IsSearching() bool {
	if ptr.Pointer() != nil {
		return C.QRadioTuner_IsSearching(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRadioTuner) IsStereo() bool {
	if ptr.Pointer() != nil {
		return C.QRadioTuner_IsStereo(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRadioTuner) RadioData() *QRadioData {
	if ptr.Pointer() != nil {
		return NewQRadioDataFromPointer(C.QRadioTuner_RadioData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QRadioTuner) SetMuted(muted bool) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_SetMuted(ptr.Pointer(), C.int(qt.GoBoolToInt(muted)))
	}
}

func (ptr *QRadioTuner) SetStereoMode(mode QRadioTuner__StereoMode) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_SetStereoMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QRadioTuner) SetVolume(volume int) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_SetVolume(ptr.Pointer(), C.int(volume))
	}
}

func (ptr *QRadioTuner) SignalStrength() int {
	if ptr.Pointer() != nil {
		return int(C.QRadioTuner_SignalStrength(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioTuner) StereoMode() QRadioTuner__StereoMode {
	if ptr.Pointer() != nil {
		return QRadioTuner__StereoMode(C.QRadioTuner_StereoMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioTuner) Volume() int {
	if ptr.Pointer() != nil {
		return int(C.QRadioTuner_Volume(ptr.Pointer()))
	}
	return 0
}

func NewQRadioTuner(parent core.QObject_ITF) *QRadioTuner {
	return NewQRadioTunerFromPointer(C.QRadioTuner_NewQRadioTuner(core.PointerFromQObject(parent)))
}

func (ptr *QRadioTuner) ConnectAntennaConnectedChanged(f func(connectionStatus bool)) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectAntennaConnectedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "antennaConnectedChanged", f)
	}
}

func (ptr *QRadioTuner) DisconnectAntennaConnectedChanged() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectAntennaConnectedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "antennaConnectedChanged")
	}
}

//export callbackQRadioTunerAntennaConnectedChanged
func callbackQRadioTunerAntennaConnectedChanged(ptrName *C.char, connectionStatus C.int) {
	qt.GetSignal(C.GoString(ptrName), "antennaConnectedChanged").(func(bool))(int(connectionStatus) != 0)
}

func (ptr *QRadioTuner) ConnectBandChanged(f func(band QRadioTuner__Band)) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectBandChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bandChanged", f)
	}
}

func (ptr *QRadioTuner) DisconnectBandChanged() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectBandChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bandChanged")
	}
}

//export callbackQRadioTunerBandChanged
func callbackQRadioTunerBandChanged(ptrName *C.char, band C.int) {
	qt.GetSignal(C.GoString(ptrName), "bandChanged").(func(QRadioTuner__Band))(QRadioTuner__Band(band))
}

func (ptr *QRadioTuner) CancelSearch() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_CancelSearch(ptr.Pointer())
	}
}

func (ptr *QRadioTuner) Error() QRadioTuner__Error {
	if ptr.Pointer() != nil {
		return QRadioTuner__Error(C.QRadioTuner_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioTuner) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioTuner_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QRadioTuner) ConnectFrequencyChanged(f func(frequency int)) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectFrequencyChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "frequencyChanged", f)
	}
}

func (ptr *QRadioTuner) DisconnectFrequencyChanged() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectFrequencyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "frequencyChanged")
	}
}

//export callbackQRadioTunerFrequencyChanged
func callbackQRadioTunerFrequencyChanged(ptrName *C.char, frequency C.int) {
	qt.GetSignal(C.GoString(ptrName), "frequencyChanged").(func(int))(int(frequency))
}

func (ptr *QRadioTuner) FrequencyStep(band QRadioTuner__Band) int {
	if ptr.Pointer() != nil {
		return int(C.QRadioTuner_FrequencyStep(ptr.Pointer(), C.int(band)))
	}
	return 0
}

func (ptr *QRadioTuner) IsBandSupported(band QRadioTuner__Band) bool {
	if ptr.Pointer() != nil {
		return C.QRadioTuner_IsBandSupported(ptr.Pointer(), C.int(band)) != 0
	}
	return false
}

func (ptr *QRadioTuner) ConnectMutedChanged(f func(muted bool)) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectMutedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mutedChanged", f)
	}
}

func (ptr *QRadioTuner) DisconnectMutedChanged() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectMutedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mutedChanged")
	}
}

//export callbackQRadioTunerMutedChanged
func callbackQRadioTunerMutedChanged(ptrName *C.char, muted C.int) {
	qt.GetSignal(C.GoString(ptrName), "mutedChanged").(func(bool))(int(muted) != 0)
}

func (ptr *QRadioTuner) SearchAllStations(searchMode QRadioTuner__SearchMode) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_SearchAllStations(ptr.Pointer(), C.int(searchMode))
	}
}

func (ptr *QRadioTuner) SearchBackward() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_SearchBackward(ptr.Pointer())
	}
}

func (ptr *QRadioTuner) SearchForward() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_SearchForward(ptr.Pointer())
	}
}

func (ptr *QRadioTuner) ConnectSearchingChanged(f func(searching bool)) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectSearchingChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "searchingChanged", f)
	}
}

func (ptr *QRadioTuner) DisconnectSearchingChanged() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectSearchingChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "searchingChanged")
	}
}

//export callbackQRadioTunerSearchingChanged
func callbackQRadioTunerSearchingChanged(ptrName *C.char, searching C.int) {
	qt.GetSignal(C.GoString(ptrName), "searchingChanged").(func(bool))(int(searching) != 0)
}

func (ptr *QRadioTuner) SetBand(band QRadioTuner__Band) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_SetBand(ptr.Pointer(), C.int(band))
	}
}

func (ptr *QRadioTuner) SetFrequency(frequency int) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_SetFrequency(ptr.Pointer(), C.int(frequency))
	}
}

func (ptr *QRadioTuner) ConnectSignalStrengthChanged(f func(strength int)) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectSignalStrengthChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "signalStrengthChanged", f)
	}
}

func (ptr *QRadioTuner) DisconnectSignalStrengthChanged() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectSignalStrengthChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "signalStrengthChanged")
	}
}

//export callbackQRadioTunerSignalStrengthChanged
func callbackQRadioTunerSignalStrengthChanged(ptrName *C.char, strength C.int) {
	qt.GetSignal(C.GoString(ptrName), "signalStrengthChanged").(func(int))(int(strength))
}

func (ptr *QRadioTuner) Start() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_Start(ptr.Pointer())
	}
}

func (ptr *QRadioTuner) ConnectStateChanged(f func(state QRadioTuner__State)) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QRadioTuner) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQRadioTunerStateChanged
func callbackQRadioTunerStateChanged(ptrName *C.char, state C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QRadioTuner__State))(QRadioTuner__State(state))
}

func (ptr *QRadioTuner) ConnectStationFound(f func(frequency int, stationId string)) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectStationFound(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stationFound", f)
	}
}

func (ptr *QRadioTuner) DisconnectStationFound() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectStationFound(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stationFound")
	}
}

//export callbackQRadioTunerStationFound
func callbackQRadioTunerStationFound(ptrName *C.char, frequency C.int, stationId *C.char) {
	qt.GetSignal(C.GoString(ptrName), "stationFound").(func(int, string))(int(frequency), C.GoString(stationId))
}

func (ptr *QRadioTuner) ConnectStereoStatusChanged(f func(stereo bool)) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectStereoStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stereoStatusChanged", f)
	}
}

func (ptr *QRadioTuner) DisconnectStereoStatusChanged() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectStereoStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stereoStatusChanged")
	}
}

//export callbackQRadioTunerStereoStatusChanged
func callbackQRadioTunerStereoStatusChanged(ptrName *C.char, stereo C.int) {
	qt.GetSignal(C.GoString(ptrName), "stereoStatusChanged").(func(bool))(int(stereo) != 0)
}

func (ptr *QRadioTuner) Stop() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_Stop(ptr.Pointer())
	}
}

func (ptr *QRadioTuner) ConnectVolumeChanged(f func(volume int)) {
	if ptr.Pointer() != nil {
		C.QRadioTuner_ConnectVolumeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "volumeChanged", f)
	}
}

func (ptr *QRadioTuner) DisconnectVolumeChanged() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_DisconnectVolumeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "volumeChanged")
	}
}

//export callbackQRadioTunerVolumeChanged
func callbackQRadioTunerVolumeChanged(ptrName *C.char, volume C.int) {
	qt.GetSignal(C.GoString(ptrName), "volumeChanged").(func(int))(int(volume))
}

func (ptr *QRadioTuner) DestroyQRadioTuner() {
	if ptr.Pointer() != nil {
		C.QRadioTuner_DestroyQRadioTuner(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
