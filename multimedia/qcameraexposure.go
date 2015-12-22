package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QCameraExposure struct {
	core.QObject
}

type QCameraExposure_ITF interface {
	core.QObject_ITF
	QCameraExposure_PTR() *QCameraExposure
}

func PointerFromQCameraExposure(ptr QCameraExposure_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraExposure_PTR().Pointer()
	}
	return nil
}

func NewQCameraExposureFromPointer(ptr unsafe.Pointer) *QCameraExposure {
	var n = new(QCameraExposure)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCameraExposure_") {
		n.SetObjectName("QCameraExposure_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraExposure) QCameraExposure_PTR() *QCameraExposure {
	return ptr
}

//QCameraExposure::ExposureMode
type QCameraExposure__ExposureMode int64

const (
	QCameraExposure__ExposureAuto          = QCameraExposure__ExposureMode(0)
	QCameraExposure__ExposureManual        = QCameraExposure__ExposureMode(1)
	QCameraExposure__ExposurePortrait      = QCameraExposure__ExposureMode(2)
	QCameraExposure__ExposureNight         = QCameraExposure__ExposureMode(3)
	QCameraExposure__ExposureBacklight     = QCameraExposure__ExposureMode(4)
	QCameraExposure__ExposureSpotlight     = QCameraExposure__ExposureMode(5)
	QCameraExposure__ExposureSports        = QCameraExposure__ExposureMode(6)
	QCameraExposure__ExposureSnow          = QCameraExposure__ExposureMode(7)
	QCameraExposure__ExposureBeach         = QCameraExposure__ExposureMode(8)
	QCameraExposure__ExposureLargeAperture = QCameraExposure__ExposureMode(9)
	QCameraExposure__ExposureSmallAperture = QCameraExposure__ExposureMode(10)
	QCameraExposure__ExposureAction        = QCameraExposure__ExposureMode(11)
	QCameraExposure__ExposureLandscape     = QCameraExposure__ExposureMode(12)
	QCameraExposure__ExposureNightPortrait = QCameraExposure__ExposureMode(13)
	QCameraExposure__ExposureTheatre       = QCameraExposure__ExposureMode(14)
	QCameraExposure__ExposureSunset        = QCameraExposure__ExposureMode(15)
	QCameraExposure__ExposureSteadyPhoto   = QCameraExposure__ExposureMode(16)
	QCameraExposure__ExposureFireworks     = QCameraExposure__ExposureMode(17)
	QCameraExposure__ExposureParty         = QCameraExposure__ExposureMode(18)
	QCameraExposure__ExposureCandlelight   = QCameraExposure__ExposureMode(19)
	QCameraExposure__ExposureBarcode       = QCameraExposure__ExposureMode(20)
	QCameraExposure__ExposureModeVendor    = QCameraExposure__ExposureMode(1000)
)

//QCameraExposure::FlashMode
type QCameraExposure__FlashMode int64

const (
	QCameraExposure__FlashAuto                 = QCameraExposure__FlashMode(0x1)
	QCameraExposure__FlashOff                  = QCameraExposure__FlashMode(0x2)
	QCameraExposure__FlashOn                   = QCameraExposure__FlashMode(0x4)
	QCameraExposure__FlashRedEyeReduction      = QCameraExposure__FlashMode(0x8)
	QCameraExposure__FlashFill                 = QCameraExposure__FlashMode(0x10)
	QCameraExposure__FlashTorch                = QCameraExposure__FlashMode(0x20)
	QCameraExposure__FlashVideoLight           = QCameraExposure__FlashMode(0x40)
	QCameraExposure__FlashSlowSyncFrontCurtain = QCameraExposure__FlashMode(0x80)
	QCameraExposure__FlashSlowSyncRearCurtain  = QCameraExposure__FlashMode(0x100)
	QCameraExposure__FlashManual               = QCameraExposure__FlashMode(0x200)
)

//QCameraExposure::MeteringMode
type QCameraExposure__MeteringMode int64

const (
	QCameraExposure__MeteringMatrix  = QCameraExposure__MeteringMode(1)
	QCameraExposure__MeteringAverage = QCameraExposure__MeteringMode(2)
	QCameraExposure__MeteringSpot    = QCameraExposure__MeteringMode(3)
)

func (ptr *QCameraExposure) Aperture() float64 {
	defer qt.Recovering("QCameraExposure::aperture")

	if ptr.Pointer() != nil {
		return float64(C.QCameraExposure_Aperture(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraExposure) ExposureCompensation() float64 {
	defer qt.Recovering("QCameraExposure::exposureCompensation")

	if ptr.Pointer() != nil {
		return float64(C.QCameraExposure_ExposureCompensation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraExposure) ExposureMode() QCameraExposure__ExposureMode {
	defer qt.Recovering("QCameraExposure::exposureMode")

	if ptr.Pointer() != nil {
		return QCameraExposure__ExposureMode(C.QCameraExposure_ExposureMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraExposure) FlashMode() QCameraExposure__FlashMode {
	defer qt.Recovering("QCameraExposure::flashMode")

	if ptr.Pointer() != nil {
		return QCameraExposure__FlashMode(C.QCameraExposure_FlashMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraExposure) IsoSensitivity() int {
	defer qt.Recovering("QCameraExposure::isoSensitivity")

	if ptr.Pointer() != nil {
		return int(C.QCameraExposure_IsoSensitivity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraExposure) MeteringMode() QCameraExposure__MeteringMode {
	defer qt.Recovering("QCameraExposure::meteringMode")

	if ptr.Pointer() != nil {
		return QCameraExposure__MeteringMode(C.QCameraExposure_MeteringMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraExposure) SetAutoAperture() {
	defer qt.Recovering("QCameraExposure::setAutoAperture")

	if ptr.Pointer() != nil {
		C.QCameraExposure_SetAutoAperture(ptr.Pointer())
	}
}

func (ptr *QCameraExposure) SetAutoIsoSensitivity() {
	defer qt.Recovering("QCameraExposure::setAutoIsoSensitivity")

	if ptr.Pointer() != nil {
		C.QCameraExposure_SetAutoIsoSensitivity(ptr.Pointer())
	}
}

func (ptr *QCameraExposure) SetExposureCompensation(ev float64) {
	defer qt.Recovering("QCameraExposure::setExposureCompensation")

	if ptr.Pointer() != nil {
		C.QCameraExposure_SetExposureCompensation(ptr.Pointer(), C.double(ev))
	}
}

func (ptr *QCameraExposure) SetExposureMode(mode QCameraExposure__ExposureMode) {
	defer qt.Recovering("QCameraExposure::setExposureMode")

	if ptr.Pointer() != nil {
		C.QCameraExposure_SetExposureMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraExposure) SetFlashMode(mode QCameraExposure__FlashMode) {
	defer qt.Recovering("QCameraExposure::setFlashMode")

	if ptr.Pointer() != nil {
		C.QCameraExposure_SetFlashMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraExposure) SetManualAperture(aperture float64) {
	defer qt.Recovering("QCameraExposure::setManualAperture")

	if ptr.Pointer() != nil {
		C.QCameraExposure_SetManualAperture(ptr.Pointer(), C.double(aperture))
	}
}

func (ptr *QCameraExposure) SetManualIsoSensitivity(iso int) {
	defer qt.Recovering("QCameraExposure::setManualIsoSensitivity")

	if ptr.Pointer() != nil {
		C.QCameraExposure_SetManualIsoSensitivity(ptr.Pointer(), C.int(iso))
	}
}

func (ptr *QCameraExposure) SetMeteringMode(mode QCameraExposure__MeteringMode) {
	defer qt.Recovering("QCameraExposure::setMeteringMode")

	if ptr.Pointer() != nil {
		C.QCameraExposure_SetMeteringMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraExposure) SetSpotMeteringPoint(point core.QPointF_ITF) {
	defer qt.Recovering("QCameraExposure::setSpotMeteringPoint")

	if ptr.Pointer() != nil {
		C.QCameraExposure_SetSpotMeteringPoint(ptr.Pointer(), core.PointerFromQPointF(point))
	}
}

func (ptr *QCameraExposure) ConnectApertureChanged(f func(value float64)) {
	defer qt.Recovering("connect QCameraExposure::apertureChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposure_ConnectApertureChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "apertureChanged", f)
	}
}

func (ptr *QCameraExposure) DisconnectApertureChanged() {
	defer qt.Recovering("disconnect QCameraExposure::apertureChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposure_DisconnectApertureChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "apertureChanged")
	}
}

//export callbackQCameraExposureApertureChanged
func callbackQCameraExposureApertureChanged(ptrName *C.char, value C.double) {
	defer qt.Recovering("callback QCameraExposure::apertureChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "apertureChanged"); signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QCameraExposure) ConnectApertureRangeChanged(f func()) {
	defer qt.Recovering("connect QCameraExposure::apertureRangeChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposure_ConnectApertureRangeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "apertureRangeChanged", f)
	}
}

func (ptr *QCameraExposure) DisconnectApertureRangeChanged() {
	defer qt.Recovering("disconnect QCameraExposure::apertureRangeChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposure_DisconnectApertureRangeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "apertureRangeChanged")
	}
}

//export callbackQCameraExposureApertureRangeChanged
func callbackQCameraExposureApertureRangeChanged(ptrName *C.char) {
	defer qt.Recovering("callback QCameraExposure::apertureRangeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "apertureRangeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCameraExposure) ConnectExposureCompensationChanged(f func(value float64)) {
	defer qt.Recovering("connect QCameraExposure::exposureCompensationChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposure_ConnectExposureCompensationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "exposureCompensationChanged", f)
	}
}

func (ptr *QCameraExposure) DisconnectExposureCompensationChanged() {
	defer qt.Recovering("disconnect QCameraExposure::exposureCompensationChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposure_DisconnectExposureCompensationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "exposureCompensationChanged")
	}
}

//export callbackQCameraExposureExposureCompensationChanged
func callbackQCameraExposureExposureCompensationChanged(ptrName *C.char, value C.double) {
	defer qt.Recovering("callback QCameraExposure::exposureCompensationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "exposureCompensationChanged"); signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QCameraExposure) ConnectFlashReady(f func(ready bool)) {
	defer qt.Recovering("connect QCameraExposure::flashReady")

	if ptr.Pointer() != nil {
		C.QCameraExposure_ConnectFlashReady(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "flashReady", f)
	}
}

func (ptr *QCameraExposure) DisconnectFlashReady() {
	defer qt.Recovering("disconnect QCameraExposure::flashReady")

	if ptr.Pointer() != nil {
		C.QCameraExposure_DisconnectFlashReady(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "flashReady")
	}
}

//export callbackQCameraExposureFlashReady
func callbackQCameraExposureFlashReady(ptrName *C.char, ready C.int) {
	defer qt.Recovering("callback QCameraExposure::flashReady")

	if signal := qt.GetSignal(C.GoString(ptrName), "flashReady"); signal != nil {
		signal.(func(bool))(int(ready) != 0)
	}

}

func (ptr *QCameraExposure) IsAvailable() bool {
	defer qt.Recovering("QCameraExposure::isAvailable")

	if ptr.Pointer() != nil {
		return C.QCameraExposure_IsAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCameraExposure) IsExposureModeSupported(mode QCameraExposure__ExposureMode) bool {
	defer qt.Recovering("QCameraExposure::isExposureModeSupported")

	if ptr.Pointer() != nil {
		return C.QCameraExposure_IsExposureModeSupported(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraExposure) IsFlashModeSupported(mode QCameraExposure__FlashMode) bool {
	defer qt.Recovering("QCameraExposure::isFlashModeSupported")

	if ptr.Pointer() != nil {
		return C.QCameraExposure_IsFlashModeSupported(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraExposure) IsFlashReady() bool {
	defer qt.Recovering("QCameraExposure::isFlashReady")

	if ptr.Pointer() != nil {
		return C.QCameraExposure_IsFlashReady(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCameraExposure) IsMeteringModeSupported(mode QCameraExposure__MeteringMode) bool {
	defer qt.Recovering("QCameraExposure::isMeteringModeSupported")

	if ptr.Pointer() != nil {
		return C.QCameraExposure_IsMeteringModeSupported(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraExposure) ConnectIsoSensitivityChanged(f func(value int)) {
	defer qt.Recovering("connect QCameraExposure::isoSensitivityChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposure_ConnectIsoSensitivityChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "isoSensitivityChanged", f)
	}
}

func (ptr *QCameraExposure) DisconnectIsoSensitivityChanged() {
	defer qt.Recovering("disconnect QCameraExposure::isoSensitivityChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposure_DisconnectIsoSensitivityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "isoSensitivityChanged")
	}
}

//export callbackQCameraExposureIsoSensitivityChanged
func callbackQCameraExposureIsoSensitivityChanged(ptrName *C.char, value C.int) {
	defer qt.Recovering("callback QCameraExposure::isoSensitivityChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "isoSensitivityChanged"); signal != nil {
		signal.(func(int))(int(value))
	}

}

func (ptr *QCameraExposure) RequestedAperture() float64 {
	defer qt.Recovering("QCameraExposure::requestedAperture")

	if ptr.Pointer() != nil {
		return float64(C.QCameraExposure_RequestedAperture(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraExposure) RequestedIsoSensitivity() int {
	defer qt.Recovering("QCameraExposure::requestedIsoSensitivity")

	if ptr.Pointer() != nil {
		return int(C.QCameraExposure_RequestedIsoSensitivity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraExposure) RequestedShutterSpeed() float64 {
	defer qt.Recovering("QCameraExposure::requestedShutterSpeed")

	if ptr.Pointer() != nil {
		return float64(C.QCameraExposure_RequestedShutterSpeed(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraExposure) SetAutoShutterSpeed() {
	defer qt.Recovering("QCameraExposure::setAutoShutterSpeed")

	if ptr.Pointer() != nil {
		C.QCameraExposure_SetAutoShutterSpeed(ptr.Pointer())
	}
}

func (ptr *QCameraExposure) SetManualShutterSpeed(seconds float64) {
	defer qt.Recovering("QCameraExposure::setManualShutterSpeed")

	if ptr.Pointer() != nil {
		C.QCameraExposure_SetManualShutterSpeed(ptr.Pointer(), C.double(seconds))
	}
}

func (ptr *QCameraExposure) ShutterSpeed() float64 {
	defer qt.Recovering("QCameraExposure::shutterSpeed")

	if ptr.Pointer() != nil {
		return float64(C.QCameraExposure_ShutterSpeed(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraExposure) ConnectShutterSpeedChanged(f func(speed float64)) {
	defer qt.Recovering("connect QCameraExposure::shutterSpeedChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposure_ConnectShutterSpeedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "shutterSpeedChanged", f)
	}
}

func (ptr *QCameraExposure) DisconnectShutterSpeedChanged() {
	defer qt.Recovering("disconnect QCameraExposure::shutterSpeedChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposure_DisconnectShutterSpeedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "shutterSpeedChanged")
	}
}

//export callbackQCameraExposureShutterSpeedChanged
func callbackQCameraExposureShutterSpeedChanged(ptrName *C.char, speed C.double) {
	defer qt.Recovering("callback QCameraExposure::shutterSpeedChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "shutterSpeedChanged"); signal != nil {
		signal.(func(float64))(float64(speed))
	}

}

func (ptr *QCameraExposure) ConnectShutterSpeedRangeChanged(f func()) {
	defer qt.Recovering("connect QCameraExposure::shutterSpeedRangeChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposure_ConnectShutterSpeedRangeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "shutterSpeedRangeChanged", f)
	}
}

func (ptr *QCameraExposure) DisconnectShutterSpeedRangeChanged() {
	defer qt.Recovering("disconnect QCameraExposure::shutterSpeedRangeChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposure_DisconnectShutterSpeedRangeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "shutterSpeedRangeChanged")
	}
}

//export callbackQCameraExposureShutterSpeedRangeChanged
func callbackQCameraExposureShutterSpeedRangeChanged(ptrName *C.char) {
	defer qt.Recovering("callback QCameraExposure::shutterSpeedRangeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "shutterSpeedRangeChanged"); signal != nil {
		signal.(func())()
	}

}
