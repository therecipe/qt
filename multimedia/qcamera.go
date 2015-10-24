package multimedia

//#include "qcamera.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QCamera struct {
	QMediaObject
}

type QCameraITF interface {
	QMediaObjectITF
	QCameraPTR() *QCamera
}

func PointerFromQCamera(ptr QCameraITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraPTR().Pointer()
	}
	return nil
}

func QCameraFromPointer(ptr unsafe.Pointer) *QCamera {
	var n = new(QCamera)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QCamera_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCamera) QCameraPTR() *QCamera {
	return ptr
}

//QCamera::CaptureMode
type QCamera__CaptureMode int

var (
	QCamera__CaptureViewfinder = QCamera__CaptureMode(0)
	QCamera__CaptureStillImage = QCamera__CaptureMode(0x01)
	QCamera__CaptureVideo      = QCamera__CaptureMode(0x02)
)

//QCamera::Error
type QCamera__Error int

var (
	QCamera__NoError                  = QCamera__Error(0)
	QCamera__CameraError              = QCamera__Error(1)
	QCamera__InvalidRequestError      = QCamera__Error(2)
	QCamera__ServiceMissingError      = QCamera__Error(3)
	QCamera__NotSupportedFeatureError = QCamera__Error(4)
)

//QCamera::LockChangeReason
type QCamera__LockChangeReason int

var (
	QCamera__UserRequest       = QCamera__LockChangeReason(0)
	QCamera__LockAcquired      = QCamera__LockChangeReason(1)
	QCamera__LockFailed        = QCamera__LockChangeReason(2)
	QCamera__LockLost          = QCamera__LockChangeReason(3)
	QCamera__LockTemporaryLost = QCamera__LockChangeReason(4)
)

//QCamera::LockStatus
type QCamera__LockStatus int

var (
	QCamera__Unlocked  = QCamera__LockStatus(0)
	QCamera__Searching = QCamera__LockStatus(1)
	QCamera__Locked    = QCamera__LockStatus(2)
)

//QCamera::LockType
type QCamera__LockType int

var (
	QCamera__NoLock           = QCamera__LockType(0)
	QCamera__LockExposure     = QCamera__LockType(0x01)
	QCamera__LockWhiteBalance = QCamera__LockType(0x02)
	QCamera__LockFocus        = QCamera__LockType(0x04)
)

//QCamera::Position
type QCamera__Position int

var (
	QCamera__UnspecifiedPosition = QCamera__Position(0)
	QCamera__BackFace            = QCamera__Position(1)
	QCamera__FrontFace           = QCamera__Position(2)
)

//QCamera::State
type QCamera__State int

var (
	QCamera__UnloadedState = QCamera__State(0)
	QCamera__LoadedState   = QCamera__State(1)
	QCamera__ActiveState   = QCamera__State(2)
)

//QCamera::Status
type QCamera__Status int

var (
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
	if ptr.Pointer() != nil {
		return QCamera__CaptureMode(C.QCamera_CaptureMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCamera) SearchAndLock2(locks QCamera__LockType) {
	if ptr.Pointer() != nil {
		C.QCamera_SearchAndLock2(C.QtObjectPtr(ptr.Pointer()), C.int(locks))
	}
}

func (ptr *QCamera) SetCaptureMode(mode QCamera__CaptureMode) {
	if ptr.Pointer() != nil {
		C.QCamera_SetCaptureMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QCamera) Status() QCamera__Status {
	if ptr.Pointer() != nil {
		return QCamera__Status(C.QCamera_Status(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQCamera4(position QCamera__Position, parent core.QObjectITF) *QCamera {
	return QCameraFromPointer(unsafe.Pointer(C.QCamera_NewQCamera4(C.int(position), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQCamera(parent core.QObjectITF) *QCamera {
	return QCameraFromPointer(unsafe.Pointer(C.QCamera_NewQCamera(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQCamera2(deviceName core.QByteArrayITF, parent core.QObjectITF) *QCamera {
	return QCameraFromPointer(unsafe.Pointer(C.QCamera_NewQCamera2(C.QtObjectPtr(core.PointerFromQByteArray(deviceName)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQCamera3(cameraInfo QCameraInfoITF, parent core.QObjectITF) *QCamera {
	return QCameraFromPointer(unsafe.Pointer(C.QCamera_NewQCamera3(C.QtObjectPtr(PointerFromQCameraInfo(cameraInfo)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QCamera) ConnectCaptureModeChanged(f func(mode QCamera__CaptureMode)) {
	if ptr.Pointer() != nil {
		C.QCamera_ConnectCaptureModeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "captureModeChanged", f)
	}
}

func (ptr *QCamera) DisconnectCaptureModeChanged() {
	if ptr.Pointer() != nil {
		C.QCamera_DisconnectCaptureModeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "captureModeChanged")
	}
}

//export callbackQCameraCaptureModeChanged
func callbackQCameraCaptureModeChanged(ptrName *C.char, mode C.int) {
	qt.GetSignal(C.GoString(ptrName), "captureModeChanged").(func(QCamera__CaptureMode))(QCamera__CaptureMode(mode))
}

func (ptr *QCamera) Error() QCamera__Error {
	if ptr.Pointer() != nil {
		return QCamera__Error(C.QCamera_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCamera) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QCamera_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QCamera) Exposure() *QCameraExposure {
	if ptr.Pointer() != nil {
		return QCameraExposureFromPointer(unsafe.Pointer(C.QCamera_Exposure(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QCamera) Focus() *QCameraFocus {
	if ptr.Pointer() != nil {
		return QCameraFocusFromPointer(unsafe.Pointer(C.QCamera_Focus(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QCamera) ImageProcessing() *QCameraImageProcessing {
	if ptr.Pointer() != nil {
		return QCameraImageProcessingFromPointer(unsafe.Pointer(C.QCamera_ImageProcessing(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QCamera) IsCaptureModeSupported(mode QCamera__CaptureMode) bool {
	if ptr.Pointer() != nil {
		return C.QCamera_IsCaptureModeSupported(C.QtObjectPtr(ptr.Pointer()), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCamera) Load() {
	if ptr.Pointer() != nil {
		C.QCamera_Load(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QCamera) ConnectLockFailed(f func()) {
	if ptr.Pointer() != nil {
		C.QCamera_ConnectLockFailed(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "lockFailed", f)
	}
}

func (ptr *QCamera) DisconnectLockFailed() {
	if ptr.Pointer() != nil {
		C.QCamera_DisconnectLockFailed(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "lockFailed")
	}
}

//export callbackQCameraLockFailed
func callbackQCameraLockFailed(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "lockFailed").(func())()
}

func (ptr *QCamera) LockStatus() QCamera__LockStatus {
	if ptr.Pointer() != nil {
		return QCamera__LockStatus(C.QCamera_LockStatus(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCamera) LockStatus2(lockType QCamera__LockType) QCamera__LockStatus {
	if ptr.Pointer() != nil {
		return QCamera__LockStatus(C.QCamera_LockStatus2(C.QtObjectPtr(ptr.Pointer()), C.int(lockType)))
	}
	return 0
}

func (ptr *QCamera) ConnectLockStatusChanged(f func(status QCamera__LockStatus, reason QCamera__LockChangeReason)) {
	if ptr.Pointer() != nil {
		C.QCamera_ConnectLockStatusChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "lockStatusChanged", f)
	}
}

func (ptr *QCamera) DisconnectLockStatusChanged() {
	if ptr.Pointer() != nil {
		C.QCamera_DisconnectLockStatusChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "lockStatusChanged")
	}
}

//export callbackQCameraLockStatusChanged
func callbackQCameraLockStatusChanged(ptrName *C.char, status C.int, reason C.int) {
	qt.GetSignal(C.GoString(ptrName), "lockStatusChanged").(func(QCamera__LockStatus, QCamera__LockChangeReason))(QCamera__LockStatus(status), QCamera__LockChangeReason(reason))
}

func (ptr *QCamera) ConnectLocked(f func()) {
	if ptr.Pointer() != nil {
		C.QCamera_ConnectLocked(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "locked", f)
	}
}

func (ptr *QCamera) DisconnectLocked() {
	if ptr.Pointer() != nil {
		C.QCamera_DisconnectLocked(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "locked")
	}
}

//export callbackQCameraLocked
func callbackQCameraLocked(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "locked").(func())()
}

func (ptr *QCamera) RequestedLocks() QCamera__LockType {
	if ptr.Pointer() != nil {
		return QCamera__LockType(C.QCamera_RequestedLocks(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCamera) SearchAndLock() {
	if ptr.Pointer() != nil {
		C.QCamera_SearchAndLock(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QCamera) SetViewfinder3(surface QAbstractVideoSurfaceITF) {
	if ptr.Pointer() != nil {
		C.QCamera_SetViewfinder3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractVideoSurface(surface)))
	}
}

func (ptr *QCamera) SetViewfinderSettings(settings QCameraViewfinderSettingsITF) {
	if ptr.Pointer() != nil {
		C.QCamera_SetViewfinderSettings(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQCameraViewfinderSettings(settings)))
	}
}

func (ptr *QCamera) Start() {
	if ptr.Pointer() != nil {
		C.QCamera_Start(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QCamera) ConnectStateChanged(f func(state QCamera__State)) {
	if ptr.Pointer() != nil {
		C.QCamera_ConnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QCamera) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QCamera_DisconnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQCameraStateChanged
func callbackQCameraStateChanged(ptrName *C.char, state C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QCamera__State))(QCamera__State(state))
}

func (ptr *QCamera) ConnectStatusChanged(f func(status QCamera__Status)) {
	if ptr.Pointer() != nil {
		C.QCamera_ConnectStatusChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "statusChanged", f)
	}
}

func (ptr *QCamera) DisconnectStatusChanged() {
	if ptr.Pointer() != nil {
		C.QCamera_DisconnectStatusChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "statusChanged")
	}
}

//export callbackQCameraStatusChanged
func callbackQCameraStatusChanged(ptrName *C.char, status C.int) {
	qt.GetSignal(C.GoString(ptrName), "statusChanged").(func(QCamera__Status))(QCamera__Status(status))
}

func (ptr *QCamera) Stop() {
	if ptr.Pointer() != nil {
		C.QCamera_Stop(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QCamera) SupportedLocks() QCamera__LockType {
	if ptr.Pointer() != nil {
		return QCamera__LockType(C.QCamera_SupportedLocks(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCamera) Unload() {
	if ptr.Pointer() != nil {
		C.QCamera_Unload(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QCamera) Unlock() {
	if ptr.Pointer() != nil {
		C.QCamera_Unlock(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QCamera) Unlock2(locks QCamera__LockType) {
	if ptr.Pointer() != nil {
		C.QCamera_Unlock2(C.QtObjectPtr(ptr.Pointer()), C.int(locks))
	}
}

func (ptr *QCamera) DestroyQCamera() {
	if ptr.Pointer() != nil {
		C.QCamera_DestroyQCamera(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
