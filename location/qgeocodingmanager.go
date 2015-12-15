package location

//#include "location.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGeoCodingManager struct {
	core.QObject
}

type QGeoCodingManager_ITF interface {
	core.QObject_ITF
	QGeoCodingManager_PTR() *QGeoCodingManager
}

func PointerFromQGeoCodingManager(ptr QGeoCodingManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoCodingManager_PTR().Pointer()
	}
	return nil
}

func NewQGeoCodingManagerFromPointer(ptr unsafe.Pointer) *QGeoCodingManager {
	var n = new(QGeoCodingManager)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGeoCodingManager_") {
		n.SetObjectName("QGeoCodingManager_" + qt.Identifier())
	}
	return n
}

func (ptr *QGeoCodingManager) QGeoCodingManager_PTR() *QGeoCodingManager {
	return ptr
}
