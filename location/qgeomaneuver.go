package location

//#include "location.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/positioning"
	"unsafe"
)

type QGeoManeuver struct {
	ptr unsafe.Pointer
}

type QGeoManeuver_ITF interface {
	QGeoManeuver_PTR() *QGeoManeuver
}

func (p *QGeoManeuver) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoManeuver) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoManeuver(ptr QGeoManeuver_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoManeuver_PTR().Pointer()
	}
	return nil
}

func NewQGeoManeuverFromPointer(ptr unsafe.Pointer) *QGeoManeuver {
	var n = new(QGeoManeuver)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGeoManeuver) QGeoManeuver_PTR() *QGeoManeuver {
	return ptr
}

//QGeoManeuver::InstructionDirection
type QGeoManeuver__InstructionDirection int64

const (
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
	defer qt.Recovering("QGeoManeuver::QGeoManeuver")

	return NewQGeoManeuverFromPointer(C.QGeoManeuver_NewQGeoManeuver())
}

func NewQGeoManeuver2(other QGeoManeuver_ITF) *QGeoManeuver {
	defer qt.Recovering("QGeoManeuver::QGeoManeuver")

	return NewQGeoManeuverFromPointer(C.QGeoManeuver_NewQGeoManeuver2(PointerFromQGeoManeuver(other)))
}

func (ptr *QGeoManeuver) Direction() QGeoManeuver__InstructionDirection {
	defer qt.Recovering("QGeoManeuver::direction")

	if ptr.Pointer() != nil {
		return QGeoManeuver__InstructionDirection(C.QGeoManeuver_Direction(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoManeuver) DistanceToNextInstruction() float64 {
	defer qt.Recovering("QGeoManeuver::distanceToNextInstruction")

	if ptr.Pointer() != nil {
		return float64(C.QGeoManeuver_DistanceToNextInstruction(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoManeuver) InstructionText() string {
	defer qt.Recovering("QGeoManeuver::instructionText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoManeuver_InstructionText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoManeuver) IsValid() bool {
	defer qt.Recovering("QGeoManeuver::isValid")

	if ptr.Pointer() != nil {
		return C.QGeoManeuver_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGeoManeuver) SetDirection(direction QGeoManeuver__InstructionDirection) {
	defer qt.Recovering("QGeoManeuver::setDirection")

	if ptr.Pointer() != nil {
		C.QGeoManeuver_SetDirection(ptr.Pointer(), C.int(direction))
	}
}

func (ptr *QGeoManeuver) SetDistanceToNextInstruction(distance float64) {
	defer qt.Recovering("QGeoManeuver::setDistanceToNextInstruction")

	if ptr.Pointer() != nil {
		C.QGeoManeuver_SetDistanceToNextInstruction(ptr.Pointer(), C.double(distance))
	}
}

func (ptr *QGeoManeuver) SetInstructionText(instructionText string) {
	defer qt.Recovering("QGeoManeuver::setInstructionText")

	if ptr.Pointer() != nil {
		C.QGeoManeuver_SetInstructionText(ptr.Pointer(), C.CString(instructionText))
	}
}

func (ptr *QGeoManeuver) SetPosition(position positioning.QGeoCoordinate_ITF) {
	defer qt.Recovering("QGeoManeuver::setPosition")

	if ptr.Pointer() != nil {
		C.QGeoManeuver_SetPosition(ptr.Pointer(), positioning.PointerFromQGeoCoordinate(position))
	}
}

func (ptr *QGeoManeuver) SetTimeToNextInstruction(secs int) {
	defer qt.Recovering("QGeoManeuver::setTimeToNextInstruction")

	if ptr.Pointer() != nil {
		C.QGeoManeuver_SetTimeToNextInstruction(ptr.Pointer(), C.int(secs))
	}
}

func (ptr *QGeoManeuver) SetWaypoint(coordinate positioning.QGeoCoordinate_ITF) {
	defer qt.Recovering("QGeoManeuver::setWaypoint")

	if ptr.Pointer() != nil {
		C.QGeoManeuver_SetWaypoint(ptr.Pointer(), positioning.PointerFromQGeoCoordinate(coordinate))
	}
}

func (ptr *QGeoManeuver) TimeToNextInstruction() int {
	defer qt.Recovering("QGeoManeuver::timeToNextInstruction")

	if ptr.Pointer() != nil {
		return int(C.QGeoManeuver_TimeToNextInstruction(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoManeuver) DestroyQGeoManeuver() {
	defer qt.Recovering("QGeoManeuver::~QGeoManeuver")

	if ptr.Pointer() != nil {
		C.QGeoManeuver_DestroyQGeoManeuver(ptr.Pointer())
	}
}
