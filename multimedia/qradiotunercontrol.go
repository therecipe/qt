package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

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
func callbackQRadioTunerControlAntennaConnectedChanged(ptrName *C.char, connectionStatus C.int) {
	defer qt.Recovering("callback QRadioTunerControl::antennaConnectedChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "antennaConnectedChanged")
	if signal != nil {
		signal.(func(bool))(int(connectionStatus) != 0)
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
func callbackQRadioTunerControlBandChanged(ptrName *C.char, band C.int) {
	defer qt.Recovering("callback QRadioTunerControl::bandChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "bandChanged")
	if signal != nil {
		signal.(func(QRadioTuner__Band))(QRadioTuner__Band(band))
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
func callbackQRadioTunerControlError2(ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QRadioTunerControl::error")

	var signal = qt.GetSignal(C.GoString(ptrName), "error2")
	if signal != nil {
		signal.(func(QRadioTuner__Error))(QRadioTuner__Error(error))
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
func callbackQRadioTunerControlFrequencyChanged(ptrName *C.char, frequency C.int) {
	defer qt.Recovering("callback QRadioTunerControl::frequencyChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "frequencyChanged")
	if signal != nil {
		signal.(func(int))(int(frequency))
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
func callbackQRadioTunerControlMutedChanged(ptrName *C.char, muted C.int) {
	defer qt.Recovering("callback QRadioTunerControl::mutedChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "mutedChanged")
	if signal != nil {
		signal.(func(bool))(int(muted) != 0)
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
func callbackQRadioTunerControlSearchingChanged(ptrName *C.char, searching C.int) {
	defer qt.Recovering("callback QRadioTunerControl::searchingChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "searchingChanged")
	if signal != nil {
		signal.(func(bool))(int(searching) != 0)
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
func callbackQRadioTunerControlSignalStrengthChanged(ptrName *C.char, strength C.int) {
	defer qt.Recovering("callback QRadioTunerControl::signalStrengthChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "signalStrengthChanged")
	if signal != nil {
		signal.(func(int))(int(strength))
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
func callbackQRadioTunerControlStateChanged(ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QRadioTunerControl::stateChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "stateChanged")
	if signal != nil {
		signal.(func(QRadioTuner__State))(QRadioTuner__State(state))
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
func callbackQRadioTunerControlStationFound(ptrName *C.char, frequency C.int, stationId *C.char) {
	defer qt.Recovering("callback QRadioTunerControl::stationFound")

	var signal = qt.GetSignal(C.GoString(ptrName), "stationFound")
	if signal != nil {
		signal.(func(int, string))(int(frequency), C.GoString(stationId))
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
func callbackQRadioTunerControlStereoStatusChanged(ptrName *C.char, stereo C.int) {
	defer qt.Recovering("callback QRadioTunerControl::stereoStatusChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "stereoStatusChanged")
	if signal != nil {
		signal.(func(bool))(int(stereo) != 0)
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
func callbackQRadioTunerControlVolumeChanged(ptrName *C.char, volume C.int) {
	defer qt.Recovering("callback QRadioTunerControl::volumeChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "volumeChanged")
	if signal != nil {
		signal.(func(int))(int(volume))
	}

}

func (ptr *QRadioTunerControl) DestroyQRadioTunerControl() {
	defer qt.Recovering("QRadioTunerControl::~QRadioTunerControl")

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DestroyQRadioTunerControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
