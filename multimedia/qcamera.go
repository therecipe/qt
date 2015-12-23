package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QCamera struct {
	QMediaObject
}

type QCamera_ITF interface {
	QMediaObject_ITF
	QCamera_PTR() *QCamera
}

func PointerFromQCamera(ptr QCamera_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCamera_PTR().Pointer()
	}
	return nil
}

func NewQCameraFromPointer(ptr unsafe.Pointer) *QCamera {
	var n = new(QCamera)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCamera_") {
		n.SetObjectName("QCamera_" + qt.Identifier())
	}
	return n
}

func (ptr *QCamera) QCamera_PTR() *QCamera {
	return ptr
}

//QCamera::CaptureMode
type QCamera__CaptureMode int64

const (
	QCamera__CaptureViewfinder = QCamera__CaptureMode(0)
	QCamera__CaptureStillImage = QCamera__CaptureMode(0x01)
	QCamera__CaptureVideo      = QCamera__CaptureMode(0x02)
)

//QCamera::Error
type QCamera__Error int64

const (
	QCamera__NoError                  = QCamera__Error(0)
	QCamera__CameraError              = QCamera__Error(1)
	QCamera__InvalidRequestError      = QCamera__Error(2)
	QCamera__ServiceMissingError      = QCamera__Error(3)
	QCamera__NotSupportedFeatureError = QCamera__Error(4)
)

//QCamera::LockChangeReason
type QCamera__LockChangeReason int64

const (
	QCamera__UserRequest       = QCamera__LockChangeReason(0)
	QCamera__LockAcquired      = QCamera__LockChangeReason(1)
	QCamera__LockFailed        = QCamera__LockChangeReason(2)
	QCamera__LockLost          = QCamera__LockChangeReason(3)
	QCamera__LockTemporaryLost = QCamera__LockChangeReason(4)
)

//QCamera::LockStatus
type QCamera__LockStatus int64

const (
	QCamera__Unlocked  = QCamera__LockStatus(0)
	QCamera__Searching = QCamera__LockStatus(1)
	QCamera__Locked    = QCamera__LockStatus(2)
)

//QCamera::LockType
type QCamera__LockType int64

const (
	QCamera__NoLock           = QCamera__LockType(0)
	QCamera__LockExposure     = QCamera__LockType(0x01)
	QCamera__LockWhiteBalance = QCamera__LockType(0x02)
	QCamera__LockFocus        = QCamera__LockType(0x04)
)

//QCamera::Position
type QCamera__Position int64

const (
	QCamera__UnspecifiedPosition = QCamera__Position(0)
	QCamera__BackFace            = QCamera__Position(1)
	QCamera__FrontFace           = QCamera__Position(2)
)

//QCamera::State
type QCamera__State int64

const (
	QCamera__UnloadedState = QCamera__State(0)
	QCamera__LoadedState   = QCamera__State(1)
	QCamera__ActiveState   = QCamera__State(2)
)

//QCamera::Status
type QCamera__Status int64

const (
	QCamera__UnavailableStatus = QCamera__Status(0)
	QCamera__UnloadedStatus    = QCamera__Status(1)
	QCamera__LoadingStatus     = QCamera__Status(2)
	QCamera__UnloadingStatus   = QCamera__Status(3)
	QCamera__LoadedStatus      = QCamera__Status(4)
	QCamera__StandbyStatus     = QCamera__Status(5)
	QCamera__StartingStatus    = QCamera__Status(6)
	QCamera__StoppingStatus    = QCamera__Status(7)
	QCamera__ActiveStatus      = QCamera__Status(8)
)

func (ptr *QCamera) CaptureMode() QCamera__CaptureMode {
	defer qt.Recovering("QCamera::captureMode")

	if ptr.Pointer() != nil {
		return QCamera__CaptureMode(C.QCamera_CaptureMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCamera) SearchAndLock2(locks QCamera__LockType) {
	defer qt.Recovering("QCamera::searchAndLock")

	if ptr.Pointer() != nil {
		C.QCamera_SearchAndLock2(ptr.Pointer(), C.int(locks))
	}
}

func (ptr *QCamera) SetCaptureMode(mode QCamera__CaptureMode) {
	defer qt.Recovering("QCamera::setCaptureMode")

	if ptr.Pointer() != nil {
		C.QCamera_SetCaptureMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCamera) State() QCamera__State {
	defer qt.Recovering("QCamera::state")

	if ptr.Pointer() != nil {
		return QCamera__State(C.QCamera_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCamera) Status() QCamera__Status {
	defer qt.Recovering("QCamera::status")

	if ptr.Pointer() != nil {
		return QCamera__Status(C.QCamera_Status(ptr.Pointer()))
	}
	return 0
}

func NewQCamera4(position QCamera__Position, parent core.QObject_ITF) *QCamera {
	defer qt.Recovering("QCamera::QCamera")

	return NewQCameraFromPointer(C.QCamera_NewQCamera4(C.int(position), core.PointerFromQObject(parent)))
}

func NewQCamera(parent core.QObject_ITF) *QCamera {
	defer qt.Recovering("QCamera::QCamera")

	return NewQCameraFromPointer(C.QCamera_NewQCamera(core.PointerFromQObject(parent)))
}

func NewQCamera2(deviceName core.QByteArray_ITF, parent core.QObject_ITF) *QCamera {
	defer qt.Recovering("QCamera::QCamera")

	return NewQCameraFromPointer(C.QCamera_NewQCamera2(core.PointerFromQByteArray(deviceName), core.PointerFromQObject(parent)))
}

func NewQCamera3(cameraInfo QCameraInfo_ITF, parent core.QObject_ITF) *QCamera {
	defer qt.Recovering("QCamera::QCamera")

	return NewQCameraFromPointer(C.QCamera_NewQCamera3(PointerFromQCameraInfo(cameraInfo), core.PointerFromQObject(parent)))
}

func (ptr *QCamera) Availability() QMultimedia__AvailabilityStatus {
	defer qt.Recovering("QCamera::availability")

	if ptr.Pointer() != nil {
		return QMultimedia__AvailabilityStatus(C.QCamera_Availability(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCamera) ConnectCaptureModeChanged(f func(mode QCamera__CaptureMode)) {
	defer qt.Recovering("connect QCamera::captureModeChanged")

	if ptr.Pointer() != nil {
		C.QCamera_ConnectCaptureModeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "captureModeChanged", f)
	}
}

func (ptr *QCamera) DisconnectCaptureModeChanged() {
	defer qt.Recovering("disconnect QCamera::captureModeChanged")

	if ptr.Pointer() != nil {
		C.QCamera_DisconnectCaptureModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "captureModeChanged")
	}
}

//export callbackQCameraCaptureModeChanged
func callbackQCameraCaptureModeChanged(ptrName *C.char, mode C.int) {
	defer qt.Recovering("callback QCamera::captureModeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "captureModeChanged"); signal != nil {
		signal.(func(QCamera__CaptureMode))(QCamera__CaptureMode(mode))
	}

}

func (ptr *QCamera) ConnectError2(f func(value QCamera__Error)) {
	defer qt.Recovering("connect QCamera::error")

	if ptr.Pointer() != nil {
		C.QCamera_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QCamera) DisconnectError2() {
	defer qt.Recovering("disconnect QCamera::error")

	if ptr.Pointer() != nil {
		C.QCamera_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQCameraError2
func callbackQCameraError2(ptrName *C.char, value C.int) {
	defer qt.Recovering("callback QCamera::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QCamera__Error))(QCamera__Error(value))
	}

}

func (ptr *QCamera) Error() QCamera__Error {
	defer qt.Recovering("QCamera::error")

	if ptr.Pointer() != nil {
		return QCamera__Error(C.QCamera_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCamera) ErrorString() string {
	defer qt.Recovering("QCamera::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QCamera_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QCamera) Exposure() *QCameraExposure {
	defer qt.Recovering("QCamera::exposure")

	if ptr.Pointer() != nil {
		return NewQCameraExposureFromPointer(C.QCamera_Exposure(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCamera) Focus() *QCameraFocus {
	defer qt.Recovering("QCamera::focus")

	if ptr.Pointer() != nil {
		return NewQCameraFocusFromPointer(C.QCamera_Focus(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCamera) ImageProcessing() *QCameraImageProcessing {
	defer qt.Recovering("QCamera::imageProcessing")

	if ptr.Pointer() != nil {
		return NewQCameraImageProcessingFromPointer(C.QCamera_ImageProcessing(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCamera) IsCaptureModeSupported(mode QCamera__CaptureMode) bool {
	defer qt.Recovering("QCamera::isCaptureModeSupported")

	if ptr.Pointer() != nil {
		return C.QCamera_IsCaptureModeSupported(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCamera) Load() {
	defer qt.Recovering("QCamera::load")

	if ptr.Pointer() != nil {
		C.QCamera_Load(ptr.Pointer())
	}
}

func (ptr *QCamera) ConnectLockFailed(f func()) {
	defer qt.Recovering("connect QCamera::lockFailed")

	if ptr.Pointer() != nil {
		C.QCamera_ConnectLockFailed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "lockFailed", f)
	}
}

func (ptr *QCamera) DisconnectLockFailed() {
	defer qt.Recovering("disconnect QCamera::lockFailed")

	if ptr.Pointer() != nil {
		C.QCamera_DisconnectLockFailed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "lockFailed")
	}
}

//export callbackQCameraLockFailed
func callbackQCameraLockFailed(ptrName *C.char) {
	defer qt.Recovering("callback QCamera::lockFailed")

	if signal := qt.GetSignal(C.GoString(ptrName), "lockFailed"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCamera) LockStatus() QCamera__LockStatus {
	defer qt.Recovering("QCamera::lockStatus")

	if ptr.Pointer() != nil {
		return QCamera__LockStatus(C.QCamera_LockStatus(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCamera) LockStatus2(lockType QCamera__LockType) QCamera__LockStatus {
	defer qt.Recovering("QCamera::lockStatus")

	if ptr.Pointer() != nil {
		return QCamera__LockStatus(C.QCamera_LockStatus2(ptr.Pointer(), C.int(lockType)))
	}
	return 0
}

func (ptr *QCamera) ConnectLockStatusChanged(f func(status QCamera__LockStatus, reason QCamera__LockChangeReason)) {
	defer qt.Recovering("connect QCamera::lockStatusChanged")

	if ptr.Pointer() != nil {
		C.QCamera_ConnectLockStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "lockStatusChanged", f)
	}
}

func (ptr *QCamera) DisconnectLockStatusChanged() {
	defer qt.Recovering("disconnect QCamera::lockStatusChanged")

	if ptr.Pointer() != nil {
		C.QCamera_DisconnectLockStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "lockStatusChanged")
	}
}

//export callbackQCameraLockStatusChanged
func callbackQCameraLockStatusChanged(ptrName *C.char, status C.int, reason C.int) {
	defer qt.Recovering("callback QCamera::lockStatusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "lockStatusChanged"); signal != nil {
		signal.(func(QCamera__LockStatus, QCamera__LockChangeReason))(QCamera__LockStatus(status), QCamera__LockChangeReason(reason))
	}

}

func (ptr *QCamera) ConnectLockStatusChanged2(f func(lock QCamera__LockType, status QCamera__LockStatus, reason QCamera__LockChangeReason)) {
	defer qt.Recovering("connect QCamera::lockStatusChanged")

	if ptr.Pointer() != nil {
		C.QCamera_ConnectLockStatusChanged2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "lockStatusChanged2", f)
	}
}

func (ptr *QCamera) DisconnectLockStatusChanged2() {
	defer qt.Recovering("disconnect QCamera::lockStatusChanged")

	if ptr.Pointer() != nil {
		C.QCamera_DisconnectLockStatusChanged2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "lockStatusChanged2")
	}
}

//export callbackQCameraLockStatusChanged2
func callbackQCameraLockStatusChanged2(ptrName *C.char, lock C.int, status C.int, reason C.int) {
	defer qt.Recovering("callback QCamera::lockStatusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "lockStatusChanged2"); signal != nil {
		signal.(func(QCamera__LockType, QCamera__LockStatus, QCamera__LockChangeReason))(QCamera__LockType(lock), QCamera__LockStatus(status), QCamera__LockChangeReason(reason))
	}

}

func (ptr *QCamera) ConnectLocked(f func()) {
	defer qt.Recovering("connect QCamera::locked")

	if ptr.Pointer() != nil {
		C.QCamera_ConnectLocked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "locked", f)
	}
}

func (ptr *QCamera) DisconnectLocked() {
	defer qt.Recovering("disconnect QCamera::locked")

	if ptr.Pointer() != nil {
		C.QCamera_DisconnectLocked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "locked")
	}
}

//export callbackQCameraLocked
func callbackQCameraLocked(ptrName *C.char) {
	defer qt.Recovering("callback QCamera::locked")

	if signal := qt.GetSignal(C.GoString(ptrName), "locked"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCamera) RequestedLocks() QCamera__LockType {
	defer qt.Recovering("QCamera::requestedLocks")

	if ptr.Pointer() != nil {
		return QCamera__LockType(C.QCamera_RequestedLocks(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCamera) SearchAndLock() {
	defer qt.Recovering("QCamera::searchAndLock")

	if ptr.Pointer() != nil {
		C.QCamera_SearchAndLock(ptr.Pointer())
	}
}

func (ptr *QCamera) SetViewfinder3(surface QAbstractVideoSurface_ITF) {
	defer qt.Recovering("QCamera::setViewfinder")

	if ptr.Pointer() != nil {
		C.QCamera_SetViewfinder3(ptr.Pointer(), PointerFromQAbstractVideoSurface(surface))
	}
}

func (ptr *QCamera) SetViewfinder2(viewfinder unsafe.Pointer) {
	defer qt.Recovering("QCamera::setViewfinder")

	if ptr.Pointer() != nil {
		C.QCamera_SetViewfinder2(ptr.Pointer(), viewfinder)
	}
}

func (ptr *QCamera) SetViewfinder(viewfinder unsafe.Pointer) {
	defer qt.Recovering("QCamera::setViewfinder")

	if ptr.Pointer() != nil {
		C.QCamera_SetViewfinder(ptr.Pointer(), viewfinder)
	}
}

func (ptr *QCamera) SetViewfinderSettings(settings QCameraViewfinderSettings_ITF) {
	defer qt.Recovering("QCamera::setViewfinderSettings")

	if ptr.Pointer() != nil {
		C.QCamera_SetViewfinderSettings(ptr.Pointer(), PointerFromQCameraViewfinderSettings(settings))
	}
}

func (ptr *QCamera) Start() {
	defer qt.Recovering("QCamera::start")

	if ptr.Pointer() != nil {
		C.QCamera_Start(ptr.Pointer())
	}
}

func (ptr *QCamera) ConnectStateChanged(f func(state QCamera__State)) {
	defer qt.Recovering("connect QCamera::stateChanged")

	if ptr.Pointer() != nil {
		C.QCamera_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QCamera) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QCamera::stateChanged")

	if ptr.Pointer() != nil {
		C.QCamera_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQCameraStateChanged
func callbackQCameraStateChanged(ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QCamera::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QCamera__State))(QCamera__State(state))
	}

}

func (ptr *QCamera) ConnectStatusChanged(f func(status QCamera__Status)) {
	defer qt.Recovering("connect QCamera::statusChanged")

	if ptr.Pointer() != nil {
		C.QCamera_ConnectStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "statusChanged", f)
	}
}

func (ptr *QCamera) DisconnectStatusChanged() {
	defer qt.Recovering("disconnect QCamera::statusChanged")

	if ptr.Pointer() != nil {
		C.QCamera_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "statusChanged")
	}
}

//export callbackQCameraStatusChanged
func callbackQCameraStatusChanged(ptrName *C.char, status C.int) {
	defer qt.Recovering("callback QCamera::statusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "statusChanged"); signal != nil {
		signal.(func(QCamera__Status))(QCamera__Status(status))
	}

}

func (ptr *QCamera) Stop() {
	defer qt.Recovering("QCamera::stop")

	if ptr.Pointer() != nil {
		C.QCamera_Stop(ptr.Pointer())
	}
}

func (ptr *QCamera) SupportedLocks() QCamera__LockType {
	defer qt.Recovering("QCamera::supportedLocks")

	if ptr.Pointer() != nil {
		return QCamera__LockType(C.QCamera_SupportedLocks(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCamera) Unload() {
	defer qt.Recovering("QCamera::unload")

	if ptr.Pointer() != nil {
		C.QCamera_Unload(ptr.Pointer())
	}
}

func (ptr *QCamera) Unlock() {
	defer qt.Recovering("QCamera::unlock")

	if ptr.Pointer() != nil {
		C.QCamera_Unlock(ptr.Pointer())
	}
}

func (ptr *QCamera) Unlock2(locks QCamera__LockType) {
	defer qt.Recovering("QCamera::unlock")

	if ptr.Pointer() != nil {
		C.QCamera_Unlock2(ptr.Pointer(), C.int(locks))
	}
}

func (ptr *QCamera) DestroyQCamera() {
	defer qt.Recovering("QCamera::~QCamera")

	if ptr.Pointer() != nil {
		C.QCamera_DestroyQCamera(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCamera) ConnectUnbind(f func(object *core.QObject)) {
	defer qt.Recovering("connect QCamera::unbind")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "unbind", f)
	}
}

func (ptr *QCamera) DisconnectUnbind() {
	defer qt.Recovering("disconnect QCamera::unbind")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "unbind")
	}
}

//export callbackQCameraUnbind
func callbackQCameraUnbind(ptrName *C.char, object unsafe.Pointer) bool {
	defer qt.Recovering("callback QCamera::unbind")

	if signal := qt.GetSignal(C.GoString(ptrName), "unbind"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(object))
		return true
	}
	return false

}

func (ptr *QCamera) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCamera::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCamera) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCamera::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraTimerEvent
func callbackQCameraTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCamera::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCamera) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCamera::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCamera) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCamera::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraChildEvent
func callbackQCameraChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCamera::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCamera) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCamera::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCamera) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCamera::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraCustomEvent
func callbackQCameraCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCamera::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
