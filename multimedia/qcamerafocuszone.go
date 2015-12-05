package multimedia

//#include "multimedia.h"
import "C"
import (
	"log"
	"unsafe"
)

type QCameraFocusZone struct {
	ptr unsafe.Pointer
}

type QCameraFocusZone_ITF interface {
	QCameraFocusZone_PTR() *QCameraFocusZone
}

func (p *QCameraFocusZone) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QCameraFocusZone) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQCameraFocusZone(ptr QCameraFocusZone_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraFocusZone_PTR().Pointer()
	}
	return nil
}

func NewQCameraFocusZoneFromPointer(ptr unsafe.Pointer) *QCameraFocusZone {
	var n = new(QCameraFocusZone)
	n.SetPointer(ptr)
	return n
}

func (ptr *QCameraFocusZone) QCameraFocusZone_PTR() *QCameraFocusZone {
	return ptr
}

//QCameraFocusZone::FocusZoneStatus
type QCameraFocusZone__FocusZoneStatus int64

const (
	QCameraFocusZone__Invalid  = QCameraFocusZone__FocusZoneStatus(0)
	QCameraFocusZone__Unused   = QCameraFocusZone__FocusZoneStatus(1)
	QCameraFocusZone__Selected = QCameraFocusZone__FocusZoneStatus(2)
	QCameraFocusZone__Focused  = QCameraFocusZone__FocusZoneStatus(3)
)

func NewQCameraFocusZone(other QCameraFocusZone_ITF) *QCameraFocusZone {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocusZone::QCameraFocusZone")
		}
	}()

	return NewQCameraFocusZoneFromPointer(C.QCameraFocusZone_NewQCameraFocusZone(PointerFromQCameraFocusZone(other)))
}

func (ptr *QCameraFocusZone) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocusZone::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QCameraFocusZone_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCameraFocusZone) Status() QCameraFocusZone__FocusZoneStatus {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocusZone::status")
		}
	}()

	if ptr.Pointer() != nil {
		return QCameraFocusZone__FocusZoneStatus(C.QCameraFocusZone_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFocusZone) DestroyQCameraFocusZone() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocusZone::~QCameraFocusZone")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraFocusZone_DestroyQCameraFocusZone(ptr.Pointer())
	}
}
