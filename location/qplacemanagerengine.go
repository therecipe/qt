package location

//#include "location.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPlaceManagerEngine struct {
	core.QObject
}

type QPlaceManagerEngine_ITF interface {
	core.QObject_ITF
	QPlaceManagerEngine_PTR() *QPlaceManagerEngine
}

func PointerFromQPlaceManagerEngine(ptr QPlaceManagerEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceManagerEngine_PTR().Pointer()
	}
	return nil
}

func NewQPlaceManagerEngineFromPointer(ptr unsafe.Pointer) *QPlaceManagerEngine {
	var n = new(QPlaceManagerEngine)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QPlaceManagerEngine_") {
		n.SetObjectName("QPlaceManagerEngine_" + qt.Identifier())
	}
	return n
}

func (ptr *QPlaceManagerEngine) QPlaceManagerEngine_PTR() *QPlaceManagerEngine {
	return ptr
}
