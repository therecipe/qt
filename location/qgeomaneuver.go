package location

//#include "qgeomaneuver.h"
import "C"
import (
	"github.com/therecipe/qt/positioning"
	"unsafe"
)

type QGeoManeuver struct {
	ptr unsafe.Pointer
}

type QGeoManeuverITF interface {
	QGeoManeuverPTR() *QGeoManeuver
}

func (p *QGeoManeuver) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoManeuver) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoManeuver(ptr QGeoManeuverITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoManeuverPTR().Pointer()
	}
	return nil
}

func QGeoManeuverFromPointer(ptr unsafe.Pointer) *QGeoManeuver {
	var n = new(QGeoManeuver)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGeoManeuver) QGeoManeuverPTR() *QGeoManeuver {
	return ptr
}

//QGeoManeuver::InstructionDirection
type QGeoManeuver__InstructionDirection int

var (
	QGeoManeuver__NoDirection         = QGeoManeuver__InstructionDirection(0)
	QGeoManeuver__DirectionForward    = QGeoManeuver__InstructionDirection(1)
	QGeoManeuver__DirectionBearRight  = QGeoManeuver__InstructionDirection(2)
	QGeoManeuver__DirectionLightRight = QGeoManeuver__InstructionDirection(3)
	QGeoManeuver__DirectionRight      = QGeoManeuver__InstructionDirection(4)
	QGeoManeuver__DirectionHardRight  = QGeoManeuver__InstructionDirection(5)
	QGeoManeuver__DirectionUTurnRight = QGeoManeuver__InstructionDirection(6)
	QGeoManeuver__DirectionUTurnLeft  = QGeoManeuver__InstructionDirection(7)
	QGeoManeuver__DirectionHardLeft   = QGeoManeuver__InstructionDirection(8)
	QGeoManeuver__DirectionLeft       = QGeoManeuver__InstructionDirection(9)
	QGeoManeuver__DirectionLightLeft  = QGeoManeuver__InstructionDirection(10)
	QGeoManeuver__DirectionBearLeft   = QGeoManeuver__InstructionDirection(11)
)

func NewQGeoManeuver() *QGeoManeuver {
	return QGeoManeuverFromPointer(unsafe.Pointer(C.QGeoManeuver_NewQGeoManeuver()))
}

func NewQGeoManeuver2(other QGeoManeuverITF) *QGeoManeuver {
	return QGeoManeuverFromPointer(unsafe.Pointer(C.QGeoManeuver_NewQGeoManeuver2(C.QtObjectPtr(PointerFromQGeoManeuver(other)))))
}

func (ptr *QGeoManeuver) Direction() QGeoManeuver__InstructionDirection {
	if ptr.Pointer() != nil {
		return QGeoManeuver__InstructionDirection(C.QGeoManeuver_Direction(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoManeuver) InstructionText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoManeuver_InstructionText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QGeoManeuver) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QGeoManeuver_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGeoManeuver) SetDirection(direction QGeoManeuver__InstructionDirection) {
	if ptr.Pointer() != nil {
		C.QGeoManeuver_SetDirection(C.QtObjectPtr(ptr.Pointer()), C.int(direction))
	}
}

func (ptr *QGeoManeuver) SetInstructionText(instructionText string) {
	if ptr.Pointer() != nil {
		C.QGeoManeuver_SetInstructionText(C.QtObjectPtr(ptr.Pointer()), C.CString(instructionText))
	}
}

func (ptr *QGeoManeuver) SetPosition(position positioning.QGeoCoordinateITF) {
	if ptr.Pointer() != nil {
		C.QGeoManeuver_SetPosition(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(positioning.PointerFromQGeoCoordinate(position)))
	}
}

func (ptr *QGeoManeuver) SetTimeToNextInstruction(secs int) {
	if ptr.Pointer() != nil {
		C.QGeoManeuver_SetTimeToNextInstruction(C.QtObjectPtr(ptr.Pointer()), C.int(secs))
	}
}

func (ptr *QGeoManeuver) SetWaypoint(coordinate positioning.QGeoCoordinateITF) {
	if ptr.Pointer() != nil {
		C.QGeoManeuver_SetWaypoint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(positioning.PointerFromQGeoCoordinate(coordinate)))
	}
}

func (ptr *QGeoManeuver) TimeToNextInstruction() int {
	if ptr.Pointer() != nil {
		return int(C.QGeoManeuver_TimeToNextInstruction(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoManeuver) DestroyQGeoManeuver() {
	if ptr.Pointer() != nil {
		C.QGeoManeuver_DestroyQGeoManeuver(C.QtObjectPtr(ptr.Pointer()))
	}
}
