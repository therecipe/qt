package multimedia

//#include "qradiotunercontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QRadioTunerControl struct {
	QMediaControl
}

type QRadioTunerControlITF interface {
	QMediaControlITF
	QRadioTunerControlPTR() *QRadioTunerControl
}

func PointerFromQRadioTunerControl(ptr QRadioTunerControlITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRadioTunerControlPTR().Pointer()
	}
	return nil
}

func QRadioTunerControlFromPointer(ptr unsafe.Pointer) *QRadioTunerControl {
	var n = new(QRadioTunerControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QRadioTunerControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QRadioTunerControl) QRadioTunerControlPTR() *QRadioTunerControl {
	return ptr
}

func (ptr *QRadioTunerControl) ConnectAntennaConnectedChanged(f func(connectionStatus bool)) {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectAntennaConnectedChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "antennaConnectedChanged", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectAntennaConnectedChanged() {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectAntennaConnectedChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "antennaConnectedChanged")
	}
}

//export callbackQRadioTunerControlAntennaConnectedChanged
func callbackQRadioTunerControlAntennaConnectedChanged(ptrName *C.char, connectionStatus C.int) {
	qt.GetSignal(C.GoString(ptrName), "antennaConnectedChanged").(func(bool))(int(connectionStatus) != 0)
}

func (ptr *QRadioTunerControl) Band() QRadioTuner__Band {
	if ptr.Pointer() != nil {
		return QRadioTuner__Band(C.QRadioTunerControl_Band(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRadioTunerControl) ConnectBandChanged(f func(band QRadioTuner__Band)) {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectBandChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "bandChanged", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectBandChanged() {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectBandChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "bandChanged")
	}
}

//export callbackQRadioTunerControlBandChanged
func callbackQRadioTunerControlBandChanged(ptrName *C.char, band C.int) {
	qt.GetSignal(C.GoString(ptrName), "bandChanged").(func(QRadioTuner__Band))(QRadioTuner__Band(band))
}

func (ptr *QRadioTunerControl) CancelSearch() {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_CancelSearch(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QRadioTunerControl) Error() QRadioTuner__Error {
	if ptr.Pointer() != nil {
		return QRadioTuner__Error(C.QRadioTunerControl_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRadioTunerControl) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioTunerControl_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QRadioTunerControl) Frequency() int {
	if ptr.Pointer() != nil {
		return int(C.QRadioTunerControl_Frequency(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRadioTunerControl) ConnectFrequencyChanged(f func(frequency int)) {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectFrequencyChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "frequencyChanged", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectFrequencyChanged() {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectFrequencyChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "frequencyChanged")
	}
}

//export callbackQRadioTunerControlFrequencyChanged
func callbackQRadioTunerControlFrequencyChanged(ptrName *C.char, frequency C.int) {
	qt.GetSignal(C.GoString(ptrName), "frequencyChanged").(func(int))(int(frequency))
}

func (ptr *QRadioTunerControl) FrequencyStep(band QRadioTuner__Band) int {
	if ptr.Pointer() != nil {
		return int(C.QRadioTunerControl_FrequencyStep(C.QtObjectPtr(ptr.Pointer()), C.int(band)))
	}
	return 0
}

func (ptr *QRadioTunerControl) IsAntennaConnected() bool {
	if ptr.Pointer() != nil {
		return C.QRadioTunerControl_IsAntennaConnected(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRadioTunerControl) IsBandSupported(band QRadioTuner__Band) bool {
	if ptr.Pointer() != nil {
		return C.QRadioTunerControl_IsBandSupported(C.QtObjectPtr(ptr.Pointer()), C.int(band)) != 0
	}
	return false
}

func (ptr *QRadioTunerControl) IsMuted() bool {
	if ptr.Pointer() != nil {
		return C.QRadioTunerControl_IsMuted(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRadioTunerControl) IsSearching() bool {
	if ptr.Pointer() != nil {
		return C.QRadioTunerControl_IsSearching(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRadioTunerControl) IsStereo() bool {
	if ptr.Pointer() != nil {
		return C.QRadioTunerControl_IsStereo(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRadioTunerControl) ConnectMutedChanged(f func(muted bool)) {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectMutedChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "mutedChanged", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectMutedChanged() {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectMutedChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "mutedChanged")
	}
}

//export callbackQRadioTunerControlMutedChanged
func callbackQRadioTunerControlMutedChanged(ptrName *C.char, muted C.int) {
	qt.GetSignal(C.GoString(ptrName), "mutedChanged").(func(bool))(int(muted) != 0)
}

func (ptr *QRadioTunerControl) SearchAllStations(searchMode QRadioTuner__SearchMode) {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_SearchAllStations(C.QtObjectPtr(ptr.Pointer()), C.int(searchMode))
	}
}

func (ptr *QRadioTunerControl) SearchBackward() {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_SearchBackward(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QRadioTunerControl) SearchForward() {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_SearchForward(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QRadioTunerControl) ConnectSearchingChanged(f func(searching bool)) {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectSearchingChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "searchingChanged", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectSearchingChanged() {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectSearchingChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "searchingChanged")
	}
}

//export callbackQRadioTunerControlSearchingChanged
func callbackQRadioTunerControlSearchingChanged(ptrName *C.char, searching C.int) {
	qt.GetSignal(C.GoString(ptrName), "searchingChanged").(func(bool))(int(searching) != 0)
}

func (ptr *QRadioTunerControl) SetBand(band QRadioTuner__Band) {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_SetBand(C.QtObjectPtr(ptr.Pointer()), C.int(band))
	}
}

func (ptr *QRadioTunerControl) SetFrequency(frequency int) {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_SetFrequency(C.QtObjectPtr(ptr.Pointer()), C.int(frequency))
	}
}

func (ptr *QRadioTunerControl) SetMuted(muted bool) {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_SetMuted(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(muted)))
	}
}

func (ptr *QRadioTunerControl) SetStereoMode(mode QRadioTuner__StereoMode) {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_SetStereoMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QRadioTunerControl) SetVolume(volume int) {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_SetVolume(C.QtObjectPtr(ptr.Pointer()), C.int(volume))
	}
}

func (ptr *QRadioTunerControl) SignalStrength() int {
	if ptr.Pointer() != nil {
		return int(C.QRadioTunerControl_SignalStrength(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRadioTunerControl) ConnectSignalStrengthChanged(f func(strength int)) {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectSignalStrengthChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "signalStrengthChanged", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectSignalStrengthChanged() {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectSignalStrengthChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "signalStrengthChanged")
	}
}

//export callbackQRadioTunerControlSignalStrengthChanged
func callbackQRadioTunerControlSignalStrengthChanged(ptrName *C.char, strength C.int) {
	qt.GetSignal(C.GoString(ptrName), "signalStrengthChanged").(func(int))(int(strength))
}

func (ptr *QRadioTunerControl) Start() {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_Start(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QRadioTunerControl) ConnectStateChanged(f func(state QRadioTuner__State)) {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQRadioTunerControlStateChanged
func callbackQRadioTunerControlStateChanged(ptrName *C.char, state C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QRadioTuner__State))(QRadioTuner__State(state))
}

func (ptr *QRadioTunerControl) ConnectStationFound(f func(frequency int, stationId string)) {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectStationFound(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "stationFound", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectStationFound() {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectStationFound(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "stationFound")
	}
}

//export callbackQRadioTunerControlStationFound
func callbackQRadioTunerControlStationFound(ptrName *C.char, frequency C.int, stationId *C.char) {
	qt.GetSignal(C.GoString(ptrName), "stationFound").(func(int, string))(int(frequency), C.GoString(stationId))
}

func (ptr *QRadioTunerControl) StereoMode() QRadioTuner__StereoMode {
	if ptr.Pointer() != nil {
		return QRadioTuner__StereoMode(C.QRadioTunerControl_StereoMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRadioTunerControl) ConnectStereoStatusChanged(f func(stereo bool)) {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectStereoStatusChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "stereoStatusChanged", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectStereoStatusChanged() {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectStereoStatusChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "stereoStatusChanged")
	}
}

//export callbackQRadioTunerControlStereoStatusChanged
func callbackQRadioTunerControlStereoStatusChanged(ptrName *C.char, stereo C.int) {
	qt.GetSignal(C.GoString(ptrName), "stereoStatusChanged").(func(bool))(int(stereo) != 0)
}

func (ptr *QRadioTunerControl) Stop() {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_Stop(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QRadioTunerControl) Volume() int {
	if ptr.Pointer() != nil {
		return int(C.QRadioTunerControl_Volume(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRadioTunerControl) ConnectVolumeChanged(f func(volume int)) {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectVolumeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "volumeChanged", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectVolumeChanged() {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectVolumeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "volumeChanged")
	}
}

//export callbackQRadioTunerControlVolumeChanged
func callbackQRadioTunerControlVolumeChanged(ptrName *C.char, volume C.int) {
	qt.GetSignal(C.GoString(ptrName), "volumeChanged").(func(int))(int(volume))
}

func (ptr *QRadioTunerControl) DestroyQRadioTunerControl() {
	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DestroyQRadioTunerControl(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
