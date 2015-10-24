package core

//#include "qiodevice.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QIODevice struct {
	QObject
}

type QIODeviceITF interface {
	QObjectITF
	QIODevicePTR() *QIODevice
}

func PointerFromQIODevice(ptr QIODeviceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIODevicePTR().Pointer()
	}
	return nil
}

func QIODeviceFromPointer(ptr unsafe.Pointer) *QIODevice {
	var n = new(QIODevice)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QIODevice_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QIODevice) QIODevicePTR() *QIODevice {
	return ptr
}

//QIODevice::OpenModeFlag
type QIODevice__OpenModeFlag int

var (
	QIODevice__NotOpen    = QIODevice__OpenModeFlag(0x0000)
	QIODevice__ReadOnly   = QIODevice__OpenModeFlag(0x0001)
	QIODevice__WriteOnly  = QIODevice__OpenModeFlag(0x0002)
	QIODevice__ReadWrite  = QIODevice__OpenModeFlag(QIODevice__ReadOnly | QIODevice__WriteOnly)
	QIODevice__Append     = QIODevice__OpenModeFlag(0x0004)
	QIODevice__Truncate   = QIODevice__OpenModeFlag(0x0008)
	QIODevice__Text       = QIODevice__OpenModeFlag(0x0010)
	QIODevice__Unbuffered = QIODevice__OpenModeFlag(0x0020)
)
