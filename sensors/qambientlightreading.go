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

type QAmbientLightReadingITF interface {
	QSensorReadingITF
	QAmbientLightReadingPTR() *QAmbientLightReading
}

func PointerFromQAmbientLightReading(ptr QAmbientLightReadingITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientLightReadingPTR().Pointer()
	}
	return nil
}

func QAmbientLightReadingFromPointer(ptr unsafe.Pointer) *QAmbientLightReading {
	var n = new(QAmbientLightReading)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAmbientLightReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAmbientLightReading) QAmbientLightReadingPTR() *QAmbientLightReading {
	return ptr
}

//QAmbientLightReading::LightLevel
type QAmbientLightReading__LightLevel int

var (
	QAmbientLightReading__Undefined = QAmbientLightReading__LightLevel(0)
	QAmbientLightReading__Dark      = QAmbientLightReading__LightLevel(1)
	QAmbientLightReading__Twilight  = QAmbientLightReading__LightLevel(2)
	QAmbientLightReading__Light     = QAmbientLightReading__LightLevel(3)
	QAmbientLightReading__Bright    = QAmbientLightReading__LightLevel(4)
	QAmbientLightReading__Sunny     = QAmbientLightReading__LightLevel(5)
)

func (ptr *QAmbientLightReading) LightLevel() QAmbientLightReading__LightLevel {
	if ptr.Pointer() != nil {
		return QAmbientLightReading__LightLevel(C.QAmbientLightReading_LightLevel(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAmbientLightReading) SetLightLevel(lightLevel QAmbientLightReading__LightLevel) {
	if ptr.Pointer() != nil {
		C.QAmbientLightReading_SetLightLevel(C.QtObjectPtr(ptr.Pointer()), C.int(lightLevel))
	}
}
