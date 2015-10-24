package sensors

//#include "qcompass.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QCompass struct {
	QSensor
}

type QCompassITF interface {
	QSensorITF
	QCompassPTR() *QCompass
}

func PointerFromQCompass(ptr QCompassITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCompassPTR().Pointer()
	}
	return nil
}

func QCompassFromPointer(ptr unsafe.Pointer) *QCompass {
	var n = new(QCompass)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QCompass_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCompass) QCompassPTR() *QCompass {
	return ptr
}

func (ptr *QCompass) Reading() *QCompassReading {
	if ptr.Pointer() != nil {
		return QCompassReadingFromPointer(unsafe.Pointer(C.QCompass_Reading(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func NewQCompass(parent core.QObjectITF) *QCompass {
	return QCompassFromPointer(unsafe.Pointer(C.QCompass_NewQCompass(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QCompass) DestroyQCompass() {
	if ptr.Pointer() != nil {
		C.QCompass_DestroyQCompass(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
