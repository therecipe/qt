package location

//#include "qgeocodingmanagerengine.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGeoCodingManagerEngine struct {
	core.QObject
}

type QGeoCodingManagerEngineITF interface {
	core.QObjectITF
	QGeoCodingManagerEnginePTR() *QGeoCodingManagerEngine
}

func PointerFromQGeoCodingManagerEngine(ptr QGeoCodingManagerEngineITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoCodingManagerEnginePTR().Pointer()
	}
	return nil
}

func QGeoCodingManagerEngineFromPointer(ptr unsafe.Pointer) *QGeoCodingManagerEngine {
	var n = new(QGeoCodingManagerEngine)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGeoCodingManagerEngine_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGeoCodingManagerEngine) QGeoCodingManagerEnginePTR() *QGeoCodingManagerEngine {
	return ptr
}
