package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QCameraFeedbackControl struct {
	QMediaControl
}

type QCameraFeedbackControl_ITF interface {
	QMediaControl_ITF
	QCameraFeedbackControl_PTR() *QCameraFeedbackControl
}

func PointerFromQCameraFeedbackControl(ptr QCameraFeedbackControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraFeedbackControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraFeedbackControlFromPointer(ptr unsafe.Pointer) *QCameraFeedbackControl {
	var n = new(QCameraFeedbackControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCameraFeedbackControl_") {
		n.SetObjectName("QCameraFeedbackControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraFeedbackControl) QCameraFeedbackControl_PTR() *QCameraFeedbackControl {
	return ptr
}

//QCameraFeedbackControl::EventType
type QCameraFeedbackControl__EventType int64

const (
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
	defer qt.Recovering("QCameraFeedbackControl::isEventFeedbackEnabled")

	if ptr.Pointer() != nil {
		return C.QCameraFeedbackControl_IsEventFeedbackEnabled(ptr.Pointer(), C.int(event)) != 0
	}
	return false
}

func (ptr *QCameraFeedbackControl) IsEventFeedbackLocked(event QCameraFeedbackControl__EventType) bool {
	defer qt.Recovering("QCameraFeedbackControl::isEventFeedbackLocked")

	if ptr.Pointer() != nil {
		return C.QCameraFeedbackControl_IsEventFeedbackLocked(ptr.Pointer(), C.int(event)) != 0
	}
	return false
}

func (ptr *QCameraFeedbackControl) ResetEventFeedback(event QCameraFeedbackControl__EventType) {
	defer qt.Recovering("QCameraFeedbackControl::resetEventFeedback")

	if ptr.Pointer() != nil {
		C.QCameraFeedbackControl_ResetEventFeedback(ptr.Pointer(), C.int(event))
	}
}

func (ptr *QCameraFeedbackControl) SetEventFeedbackEnabled(event QCameraFeedbackControl__EventType, enabled bool) bool {
	defer qt.Recovering("QCameraFeedbackControl::setEventFeedbackEnabled")

	if ptr.Pointer() != nil {
		return C.QCameraFeedbackControl_SetEventFeedbackEnabled(ptr.Pointer(), C.int(event), C.int(qt.GoBoolToInt(enabled))) != 0
	}
	return false
}

func (ptr *QCameraFeedbackControl) SetEventFeedbackSound(event QCameraFeedbackControl__EventType, filePath string) bool {
	defer qt.Recovering("QCameraFeedbackControl::setEventFeedbackSound")

	if ptr.Pointer() != nil {
		return C.QCameraFeedbackControl_SetEventFeedbackSound(ptr.Pointer(), C.int(event), C.CString(filePath)) != 0
	}
	return false
}

func (ptr *QCameraFeedbackControl) DestroyQCameraFeedbackControl() {
	defer qt.Recovering("QCameraFeedbackControl::~QCameraFeedbackControl")

	if ptr.Pointer() != nil {
		C.QCameraFeedbackControl_DestroyQCameraFeedbackControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
