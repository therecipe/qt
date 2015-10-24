package location

//#include "qgeocodingmanager.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGeoCodingManager struct {
	core.QObject
}

type QGeoCodingManagerITF interface {
	core.QObjectITF
	QGeoCodingManagerPTR() *QGeoCodingManager
}

func PointerFromQGeoCodingManager(ptr QGeoCodingManagerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoCodingManagerPTR().Pointer()
	}
	return nil
}

func QGeoCodingManagerFromPointer(ptr unsafe.Pointer) *QGeoCodingManager {
	var n = new(QGeoCodingManager)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGeoCodingManager_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGeoCodingManager) QGeoCodingManagerPTR() *QGeoCodingManager {
	return ptr
}
