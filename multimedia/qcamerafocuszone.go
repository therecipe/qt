package multimedia

//#include "qcamerafocuszone.h"
import "C"
import (
	"unsafe"
)

type QCameraFocusZone struct {
	ptr unsafe.Pointer
}

type QCameraFocusZoneITF interface {
	QCameraFocusZonePTR() *QCameraFocusZone
}

func (p *QCameraFocusZone) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QCameraFocusZone) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQCameraFocusZone(ptr QCameraFocusZoneITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraFocusZonePTR().Pointer()
	}
	return nil
}

func QCameraFocusZoneFromPointer(ptr unsafe.Pointer) *QCameraFocusZone {
	var n = new(QCameraFocusZone)
	n.SetPointer(ptr)
	return n
}

func (ptr *QCameraFocusZone) QCameraFocusZonePTR() *QCameraFocusZone {
	return ptr
}

//QCameraFocusZone::FocusZoneStatus
type QCameraFocusZone__FocusZoneStatus int

var (
	QCameraFocusZone__Invalid  = QCameraFocusZone__FocusZoneStatus(0)
	QCameraFocusZone__Unused   = QCameraFocusZone__FocusZoneStatus(1)
	QCameraFocusZone__Selected = QCameraFocusZone__FocusZoneStatus(2)
	QCameraFocusZone__Focused  = QCameraFocusZone__FocusZoneStatus(3)
)

func NewQCameraFocusZone(other QCameraFocusZoneITF) *QCameraFocusZone {
	return QCameraFocusZoneFromPointer(unsafe.Pointer(C.QCameraFocusZone_NewQCameraFocusZone(C.QtObjectPtr(PointerFromQCameraFocusZone(other)))))
}

func (ptr *QCameraFocusZone) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QCameraFocusZone_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QCameraFocusZone) Status() QCameraFocusZone__FocusZoneStatus {
	if ptr.Pointer() != nil {
		return QCameraFocusZone__FocusZoneStatus(C.QCameraFocusZone_Status(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCameraFocusZone) DestroyQCameraFocusZone() {
	if ptr.Pointer() != nil {
		C.QCameraFocusZone_DestroyQCameraFocusZone(C.QtObjectPtr(ptr.Pointer()))
	}
}
