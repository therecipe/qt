package location

//#include "qplacemanagerengine.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPlaceManagerEngine struct {
	core.QObject
}

type QPlaceManagerEngineITF interface {
	core.QObjectITF
	QPlaceManagerEnginePTR() *QPlaceManagerEngine
}

func PointerFromQPlaceManagerEngine(ptr QPlaceManagerEngineITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceManagerEnginePTR().Pointer()
	}
	return nil
}

func QPlaceManagerEngineFromPointer(ptr unsafe.Pointer) *QPlaceManagerEngine {
	var n = new(QPlaceManagerEngine)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QPlaceManagerEngine_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPlaceManagerEngine) QPlaceManagerEnginePTR() *QPlaceManagerEngine {
	return ptr
}
