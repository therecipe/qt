package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QAltimeter struct {
	QSensor
}

type QAltimeter_ITF interface {
	QSensor_ITF
	QAltimeter_PTR() *QAltimeter
}

func PointerFromQAltimeter(ptr QAltimeter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAltimeter_PTR().Pointer()
	}
	return nil
}

func NewQAltimeterFromPointer(ptr unsafe.Pointer) *QAltimeter {
	var n = new(QAltimeter)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAltimeter_") {
		n.SetObjectName("QAltimeter_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAltimeter) QAltimeter_PTR() *QAltimeter {
	return ptr
}

func (ptr *QAltimeter) Reading() *QAltimeterReading {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAltimeter::reading")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAltimeterReadingFromPointer(C.QAltimeter_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQAltimeter(parent core.QObject_ITF) *QAltimeter {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAltimeter::QAltimeter")
		}
	}()

	return NewQAltimeterFromPointer(C.QAltimeter_NewQAltimeter(core.PointerFromQObject(parent)))
}

func (ptr *QAltimeter) DestroyQAltimeter() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAltimeter::~QAltimeter")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAltimeter_DestroyQAltimeter(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
