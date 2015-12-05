package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QDial struct {
	QAbstractSlider
}

type QDial_ITF interface {
	QAbstractSlider_ITF
	QDial_PTR() *QDial
}

func PointerFromQDial(ptr QDial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDial_PTR().Pointer()
	}
	return nil
}

func NewQDialFromPointer(ptr unsafe.Pointer) *QDial {
	var n = new(QDial)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDial_") {
		n.SetObjectName("QDial_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDial) QDial_PTR() *QDial {
	return ptr
}

func (ptr *QDial) NotchSize() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDial::notchSize")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QDial_NotchSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDial) NotchTarget() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDial::notchTarget")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QDial_NotchTarget(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDial) NotchesVisible() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDial::notchesVisible")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QDial_NotchesVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDial) SetNotchesVisible(visible bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDial::setNotchesVisible")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDial_SetNotchesVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QDial) SetWrapping(on bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDial::setWrapping")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDial_SetWrapping(ptr.Pointer(), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QDial) Wrapping() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDial::wrapping")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QDial_Wrapping(ptr.Pointer()) != 0
	}
	return false
}

func NewQDial(parent QWidget_ITF) *QDial {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDial::QDial")
		}
	}()

	return NewQDialFromPointer(C.QDial_NewQDial(PointerFromQWidget(parent)))
}

func (ptr *QDial) DestroyQDial() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDial::~QDial")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDial_DestroyQDial(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
