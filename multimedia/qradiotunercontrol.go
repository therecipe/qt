package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
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
		n.SetObjectName("QRadioTunerControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QRadioTunerControl) QRadioTunerControl_PTR() *QRadioTunerControl {
	return ptr
}

func (ptr *QRadioTunerControl) ConnectAntennaConnectedChanged(f func(connectionStatus bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::antennaConnectedChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectAntennaConnectedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "antennaConnectedChanged", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectAntennaConnectedChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::antennaConnectedChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectAntennaConnectedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "antennaConnectedChanged")
	}
}

//export callbackQRadioTunerControlAntennaConnectedChanged
func callbackQRadioTunerControlAntennaConnectedChanged(ptrName *C.char, connectionStatus C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::antennaConnectedChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "antennaConnectedChanged").(func(bool))(int(connectionStatus) != 0)
}

func (ptr *QRadioTunerControl) Band() QRadioTuner__Band {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::band")
		}
	}()

	if ptr.Pointer() != nil {
		return QRadioTuner__Band(C.QRadioTunerControl_Band(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioTunerControl) ConnectBandChanged(f func(band QRadioTuner__Band)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::bandChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectBandChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bandChanged", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectBandChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::bandChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectBandChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bandChanged")
	}
}

//export callbackQRadioTunerControlBandChanged
func callbackQRadioTunerControlBandChanged(ptrName *C.char, band C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::bandChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "bandChanged").(func(QRadioTuner__Band))(QRadioTuner__Band(band))
}

func (ptr *QRadioTunerControl) CancelSearch() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::cancelSearch")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_CancelSearch(ptr.Pointer())
	}
}

func (ptr *QRadioTunerControl) Error() QRadioTuner__Error {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::error")
		}
	}()

	if ptr.Pointer() != nil {
		return QRadioTuner__Error(C.QRadioTunerControl_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioTunerControl) ErrorString() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::errorString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioTunerControl_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QRadioTunerControl) Frequency() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::frequency")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QRadioTunerControl_Frequency(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioTunerControl) ConnectFrequencyChanged(f func(frequency int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::frequencyChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectFrequencyChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "frequencyChanged", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectFrequencyChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::frequencyChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectFrequencyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "frequencyChanged")
	}
}

//export callbackQRadioTunerControlFrequencyChanged
func callbackQRadioTunerControlFrequencyChanged(ptrName *C.char, frequency C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::frequencyChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "frequencyChanged").(func(int))(int(frequency))
}

func (ptr *QRadioTunerControl) FrequencyStep(band QRadioTuner__Band) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::frequencyStep")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QRadioTunerControl_FrequencyStep(ptr.Pointer(), C.int(band)))
	}
	return 0
}

func (ptr *QRadioTunerControl) IsAntennaConnected() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::isAntennaConnected")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QRadioTunerControl_IsAntennaConnected(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRadioTunerControl) IsBandSupported(band QRadioTuner__Band) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::isBandSupported")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QRadioTunerControl_IsBandSupported(ptr.Pointer(), C.int(band)) != 0
	}
	return false
}

func (ptr *QRadioTunerControl) IsMuted() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::isMuted")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QRadioTunerControl_IsMuted(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRadioTunerControl) IsSearching() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::isSearching")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QRadioTunerControl_IsSearching(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRadioTunerControl) IsStereo() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::isStereo")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QRadioTunerControl_IsStereo(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRadioTunerControl) ConnectMutedChanged(f func(muted bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::mutedChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectMutedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mutedChanged", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectMutedChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::mutedChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectMutedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mutedChanged")
	}
}

//export callbackQRadioTunerControlMutedChanged
func callbackQRadioTunerControlMutedChanged(ptrName *C.char, muted C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::mutedChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "mutedChanged").(func(bool))(int(muted) != 0)
}

func (ptr *QRadioTunerControl) SearchAllStations(searchMode QRadioTuner__SearchMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::searchAllStations")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_SearchAllStations(ptr.Pointer(), C.int(searchMode))
	}
}

func (ptr *QRadioTunerControl) SearchBackward() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::searchBackward")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_SearchBackward(ptr.Pointer())
	}
}

func (ptr *QRadioTunerControl) SearchForward() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::searchForward")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_SearchForward(ptr.Pointer())
	}
}

func (ptr *QRadioTunerControl) ConnectSearchingChanged(f func(searching bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::searchingChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectSearchingChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "searchingChanged", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectSearchingChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::searchingChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectSearchingChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "searchingChanged")
	}
}

//export callbackQRadioTunerControlSearchingChanged
func callbackQRadioTunerControlSearchingChanged(ptrName *C.char, searching C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::searchingChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "searchingChanged").(func(bool))(int(searching) != 0)
}

func (ptr *QRadioTunerControl) SetBand(band QRadioTuner__Band) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::setBand")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_SetBand(ptr.Pointer(), C.int(band))
	}
}

func (ptr *QRadioTunerControl) SetFrequency(frequency int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::setFrequency")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_SetFrequency(ptr.Pointer(), C.int(frequency))
	}
}

func (ptr *QRadioTunerControl) SetMuted(muted bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::setMuted")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_SetMuted(ptr.Pointer(), C.int(qt.GoBoolToInt(muted)))
	}
}

func (ptr *QRadioTunerControl) SetStereoMode(mode QRadioTuner__StereoMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::setStereoMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_SetStereoMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QRadioTunerControl) SetVolume(volume int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::setVolume")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_SetVolume(ptr.Pointer(), C.int(volume))
	}
}

func (ptr *QRadioTunerControl) SignalStrength() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::signalStrength")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QRadioTunerControl_SignalStrength(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioTunerControl) ConnectSignalStrengthChanged(f func(strength int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::signalStrengthChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectSignalStrengthChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "signalStrengthChanged", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectSignalStrengthChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::signalStrengthChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectSignalStrengthChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "signalStrengthChanged")
	}
}

//export callbackQRadioTunerControlSignalStrengthChanged
func callbackQRadioTunerControlSignalStrengthChanged(ptrName *C.char, strength C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::signalStrengthChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "signalStrengthChanged").(func(int))(int(strength))
}

func (ptr *QRadioTunerControl) Start() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::start")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_Start(ptr.Pointer())
	}
}

func (ptr *QRadioTunerControl) ConnectStateChanged(f func(state QRadioTuner__State)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::stateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectStateChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::stateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQRadioTunerControlStateChanged
func callbackQRadioTunerControlStateChanged(ptrName *C.char, state C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::stateChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QRadioTuner__State))(QRadioTuner__State(state))
}

func (ptr *QRadioTunerControl) ConnectStationFound(f func(frequency int, stationId string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::stationFound")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectStationFound(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stationFound", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectStationFound() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::stationFound")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectStationFound(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stationFound")
	}
}

//export callbackQRadioTunerControlStationFound
func callbackQRadioTunerControlStationFound(ptrName *C.char, frequency C.int, stationId *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::stationFound")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "stationFound").(func(int, string))(int(frequency), C.GoString(stationId))
}

func (ptr *QRadioTunerControl) StereoMode() QRadioTuner__StereoMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::stereoMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QRadioTuner__StereoMode(C.QRadioTunerControl_StereoMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioTunerControl) ConnectStereoStatusChanged(f func(stereo bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::stereoStatusChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectStereoStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stereoStatusChanged", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectStereoStatusChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::stereoStatusChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectStereoStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stereoStatusChanged")
	}
}

//export callbackQRadioTunerControlStereoStatusChanged
func callbackQRadioTunerControlStereoStatusChanged(ptrName *C.char, stereo C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::stereoStatusChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "stereoStatusChanged").(func(bool))(int(stereo) != 0)
}

func (ptr *QRadioTunerControl) Stop() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::stop")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_Stop(ptr.Pointer())
	}
}

func (ptr *QRadioTunerControl) Volume() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::volume")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QRadioTunerControl_Volume(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioTunerControl) ConnectVolumeChanged(f func(volume int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::volumeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_ConnectVolumeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "volumeChanged", f)
	}
}

func (ptr *QRadioTunerControl) DisconnectVolumeChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::volumeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DisconnectVolumeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "volumeChanged")
	}
}

//export callbackQRadioTunerControlVolumeChanged
func callbackQRadioTunerControlVolumeChanged(ptrName *C.char, volume C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::volumeChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "volumeChanged").(func(int))(int(volume))
}

func (ptr *QRadioTunerControl) DestroyQRadioTunerControl() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRadioTunerControl::~QRadioTunerControl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRadioTunerControl_DestroyQRadioTunerControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
