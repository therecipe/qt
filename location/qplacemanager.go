package location

//#include "qplacemanager.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPlaceManager struct {
	core.QObject
}

type QPlaceManager_ITF interface {
	core.QObject_ITF
	QPlaceManager_PTR() *QPlaceManager
}

func PointerFromQPlaceManager(ptr QPlaceManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceManager_PTR().Pointer()
	}
	return nil
}

func NewQPlaceManagerFromPointer(ptr unsafe.Pointer) *QPlaceManager {
	var n = new(QPlaceManager)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QPlaceManager_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPlaceManager) QPlaceManager_PTR() *QPlaceManager {
	return ptr
}
