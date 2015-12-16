package multimedia

//#include "multimedia.h"
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

	return NewQRadioTunerFromPointer(C.QRadioTuner_NewQRadioTuner(core.PointerFromQObject(parent)))
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
func callbackQRadioTunerAntennaConnectedChanged(ptrName *C.char, connectionStatus C.int) {
	defer qt.Recovering("callback QRadioTuner::antennaConnectedChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "antennaConnectedChanged")
	if signal != nil {
		signal.(func(bool))(int(connectionStatus) != 0)
	}

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
func callbackQRadioTunerBandChanged(ptrName *C.char, band C.int) {
	defer qt.Recovering("callback QRadioTuner::bandChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "bandChanged")
	if signal != nil {
		signal.(func(QRadioTuner__Band))(QRadioTuner__Band(band))
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
func callbackQRadioTunerError2(ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QRadioTuner::error")

	var signal = qt.GetSignal(C.GoString(ptrName), "error2")
	if signal != nil {
		signal.(func(QRadioTuner__Error))(QRadioTuner__Error(error))
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
func callbackQRadioTunerFrequencyChanged(ptrName *C.char, frequency C.int) {
	defer qt.Recovering("callback QRadioTuner::frequencyChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "frequencyChanged")
	if signal != nil {
		signal.(func(int))(int(frequency))
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
func callbackQRadioTunerMutedChanged(ptrName *C.char, muted C.int) {
	defer qt.Recovering("callback QRadioTuner::mutedChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "mutedChanged")
	if signal != nil {
		signal.(func(bool))(int(muted) != 0)
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
func callbackQRadioTunerSearchingChanged(ptrName *C.char, searching C.int) {
	defer qt.Recovering("callback QRadioTuner::searchingChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "searchingChanged")
	if signal != nil {
		signal.(func(bool))(int(searching) != 0)
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
func callbackQRadioTunerSignalStrengthChanged(ptrName *C.char, strength C.int) {
	defer qt.Recovering("callback QRadioTuner::signalStrengthChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "signalStrengthChanged")
	if signal != nil {
		signal.(func(int))(int(strength))
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
func callbackQRadioTunerStateChanged(ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QRadioTuner::stateChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "stateChanged")
	if signal != nil {
		signal.(func(QRadioTuner__State))(QRadioTuner__State(state))
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
func callbackQRadioTunerStationFound(ptrName *C.char, frequency C.int, stationId *C.char) {
	defer qt.Recovering("callback QRadioTuner::stationFound")

	var signal = qt.GetSignal(C.GoString(ptrName), "stationFound")
	if signal != nil {
		signal.(func(int, string))(int(frequency), C.GoString(stationId))
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
func callbackQRadioTunerStereoStatusChanged(ptrName *C.char, stereo C.int) {
	defer qt.Recovering("callback QRadioTuner::stereoStatusChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "stereoStatusChanged")
	if signal != nil {
		signal.(func(bool))(int(stereo) != 0)
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
func callbackQRadioTunerVolumeChanged(ptrName *C.char, volume C.int) {
	defer qt.Recovering("callback QRadioTuner::volumeChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "volumeChanged")
	if signal != nil {
		signal.(func(int))(int(volume))
	}

}

func (ptr *QRadioTuner) DestroyQRadioTuner() {
	defer qt.Recovering("QRadioTuner::~QRadioTuner")

	if ptr.Pointer() != nil {
		C.QRadioTuner_DestroyQRadioTuner(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
