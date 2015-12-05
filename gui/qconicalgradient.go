package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QConicalGradient struct {
	QGradient
}

type QConicalGradient_ITF interface {
	QGradient_ITF
	QConicalGradient_PTR() *QConicalGradient
}

func PointerFromQConicalGradient(ptr QConicalGradient_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QConicalGradient_PTR().Pointer()
	}
	return nil
}

func NewQConicalGradientFromPointer(ptr unsafe.Pointer) *QConicalGradient {
	var n = new(QConicalGradient)
	n.SetPointer(ptr)
	return n
}

func (ptr *QConicalGradient) QConicalGradient_PTR() *QConicalGradient {
	return ptr
}

func NewQConicalGradient() *QConicalGradient {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QConicalGradient::QConicalGradient")
		}
	}()

	return NewQConicalGradientFromPointer(C.QConicalGradient_NewQConicalGradient())
}

func NewQConicalGradient2(center core.QPointF_ITF, angle float64) *QConicalGradient {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QConicalGradient::QConicalGradient")
		}
	}()

	return NewQConicalGradientFromPointer(C.QConicalGradient_NewQConicalGradient2(core.PointerFromQPointF(center), C.double(angle)))
}

func NewQConicalGradient3(cx float64, cy float64, angle float64) *QConicalGradient {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QConicalGradient::QConicalGradient")
		}
	}()

	return NewQConicalGradientFromPointer(C.QConicalGradient_NewQConicalGradient3(C.double(cx), C.double(cy), C.double(angle)))
}

func (ptr *QConicalGradient) Angle() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QConicalGradient::angle")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QConicalGradient_Angle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QConicalGradient) SetAngle(angle float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QConicalGradient::setAngle")
		}
	}()

	if ptr.Pointer() != nil {
		C.QConicalGradient_SetAngle(ptr.Pointer(), C.double(angle))
	}
}

func (ptr *QConicalGradient) SetCenter(center core.QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QConicalGradient::setCenter")
		}
	}()

	if ptr.Pointer() != nil {
		C.QConicalGradient_SetCenter(ptr.Pointer(), core.PointerFromQPointF(center))
	}
}

func (ptr *QConicalGradient) SetCenter2(x float64, y float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QConicalGradient::setCenter")
		}
	}()

	if ptr.Pointer() != nil {
		C.QConicalGradient_SetCenter2(ptr.Pointer(), C.double(x), C.double(y))
	}
}
