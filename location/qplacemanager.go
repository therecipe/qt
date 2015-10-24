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

type QPlaceManagerITF interface {
	core.QObjectITF
	QPlaceManagerPTR() *QPlaceManager
}

func PointerFromQPlaceManager(ptr QPlaceManagerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceManagerPTR().Pointer()
	}
	return nil
}

func QPlaceManagerFromPointer(ptr unsafe.Pointer) *QPlaceManager {
	var n = new(QPlaceManager)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QPlaceManager_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPlaceManager) QPlaceManagerPTR() *QPlaceManager {
	return ptr
}
