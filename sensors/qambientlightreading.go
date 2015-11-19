package sensors

//#include "qambientlightreading.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QAmbientLightReading struct {
	QSensorReading
}

type QAmbientLightReading_ITF interface {
	QSensorReading_ITF
	QAmbientLightReading_PTR() *QAmbientLightReading
}

func PointerFromQAmbientLightReading(ptr QAmbientLightReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientLightReading_PTR().Pointer()
	}
	return nil
}

func NewQAmbientLightReadingFromPointer(ptr unsafe.Pointer) *QAmbientLightReading {
	var n = new(QAmbientLightReading)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAmbientLightReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAmbientLightReading) QAmbientLightReading_PTR() *QAmbientLightReading {
	return ptr
}

//QAmbientLightReading::LightLevel
type QAmbientLightReading__LightLevel int64

const (
	QAmbientLightReading__Undefined = QAmbientLightReading__LightLevel(0)
	QAmbientLightReading__Dark      = QAmbientLightReading__LightLevel(1)
	QAmbientLightReading__Twilight  = QAmbientLightReading__LightLevel(2)
	QAmbientLightReading__Light     = QAmbientLightReading__LightLevel(3)
	QAmbientLightReading__Bright    = QAmbientLightReading__LightLevel(4)
	QAmbientLightReading__Sunny     = QAmbientLightReading__LightLevel(5)
)

func (ptr *QAmbientLightReading) LightLevel() QAmbientLightReading__LightLevel {
	if ptr.Pointer() != nil {
		return QAmbientLightReading__LightLevel(C.QAmbientLightReading_LightLevel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAmbientLightReading) SetLightLevel(lightLevel QAmbientLightReading__LightLevel) {
	if ptr.Pointer() != nil {
		C.QAmbientLightReading_SetLightLevel(ptr.Pointer(), C.int(lightLevel))
	}
}
