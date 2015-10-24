package multimedia

//#include "qcamerafeedbackcontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QCameraFeedbackControl struct {
	QMediaControl
}

type QCameraFeedbackControlITF interface {
	QMediaControlITF
	QCameraFeedbackControlPTR() *QCameraFeedbackControl
}

func PointerFromQCameraFeedbackControl(ptr QCameraFeedbackControlITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraFeedbackControlPTR().Pointer()
	}
	return nil
}

func QCameraFeedbackControlFromPointer(ptr unsafe.Pointer) *QCameraFeedbackControl {
	var n = new(QCameraFeedbackControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QCameraFeedbackControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCameraFeedbackControl) QCameraFeedbackControlPTR() *QCameraFeedbackControl {
	return ptr
}

//QCameraFeedbackControl::EventType
type QCameraFeedbackControl__EventType int

var (
	QCameraFeedbackControl__ViewfinderStarted   = QCameraFeedbackControl__EventType(1)
	QCameraFeedbackControl__ViewfinderStopped   = QCameraFeedbackControl__EventType(2)
	QCameraFeedbackControl__ImageCaptured       = QCameraFeedbackControl__EventType(3)
	QCameraFeedbackControl__ImageSaved          = QCameraFeedbackControl__EventType(4)
	QCameraFeedbackControl__ImageError          = QCameraFeedbackControl__EventType(5)
	QCameraFeedbackControl__RecordingStarted    = QCameraFeedbackControl__EventType(6)
	QCameraFeedbackControl__RecordingInProgress = QCameraFeedbackControl__EventType(7)
	QCameraFeedbackControl__RecordingStopped    = QCameraFeedbackControl__EventType(8)
	QCameraFeedbackControl__AutoFocusInProgress = QCameraFeedbackControl__EventType(9)
	QCameraFeedbackControl__AutoFocusLocked     = QCameraFeedbackControl__EventType(10)
	QCameraFeedbackControl__AutoFocusFailed     = QCameraFeedbackControl__EventType(11)
)

func (ptr *QCameraFeedbackControl) IsEventFeedbackEnabled(event QCameraFeedbackControl__EventType) bool {
	if ptr.Pointer() != nil {
		return C.QCameraFeedbackControl_IsEventFeedbackEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(event)) != 0
	}
	return false
}

func (ptr *QCameraFeedbackControl) IsEventFeedbackLocked(event QCameraFeedbackControl__EventType) bool {
	if ptr.Pointer() != nil {
		return C.QCameraFeedbackControl_IsEventFeedbackLocked(C.QtObjectPtr(ptr.Pointer()), C.int(event)) != 0
	}
	return false
}

func (ptr *QCameraFeedbackControl) ResetEventFeedback(event QCameraFeedbackControl__EventType) {
	if ptr.Pointer() != nil {
		C.QCameraFeedbackControl_ResetEventFeedback(C.QtObjectPtr(ptr.Pointer()), C.int(event))
	}
}

func (ptr *QCameraFeedbackControl) SetEventFeedbackEnabled(event QCameraFeedbackControl__EventType, enabled bool) bool {
	if ptr.Pointer() != nil {
		return C.QCameraFeedbackControl_SetEventFeedbackEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(event), C.int(qt.GoBoolToInt(enabled))) != 0
	}
	return false
}

func (ptr *QCameraFeedbackControl) SetEventFeedbackSound(event QCameraFeedbackControl__EventType, filePath string) bool {
	if ptr.Pointer() != nil {
		return C.QCameraFeedbackControl_SetEventFeedbackSound(C.QtObjectPtr(ptr.Pointer()), C.int(event), C.CString(filePath)) != 0
	}
	return false
}

func (ptr *QCameraFeedbackControl) DestroyQCameraFeedbackControl() {
	if ptr.Pointer() != nil {
		C.QCameraFeedbackControl_DestroyQCameraFeedbackControl(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
