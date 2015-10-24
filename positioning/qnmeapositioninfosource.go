package positioning

//#include "qnmeapositioninfosource.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QNmeaPositionInfoSource struct {
	QGeoPositionInfoSource
}

type QNmeaPositionInfoSourceITF interface {
	QGeoPositionInfoSourceITF
	QNmeaPositionInfoSourcePTR() *QNmeaPositionInfoSource
}

func PointerFromQNmeaPositionInfoSource(ptr QNmeaPositionInfoSourceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNmeaPositionInfoSourcePTR().Pointer()
	}
	return nil
}

func QNmeaPositionInfoSourceFromPointer(ptr unsafe.Pointer) *QNmeaPositionInfoSource {
	var n = new(QNmeaPositionInfoSource)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QNmeaPositionInfoSource_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QNmeaPositionInfoSource) QNmeaPositionInfoSourcePTR() *QNmeaPositionInfoSource {
	return ptr
}

//QNmeaPositionInfoSource::UpdateMode
type QNmeaPositionInfoSource__UpdateMode int

var (
	QNmeaPositionInfoSource__RealTimeMode   = QNmeaPositionInfoSource__UpdateMode(1)
	QNmeaPositionInfoSource__SimulationMode = QNmeaPositionInfoSource__UpdateMode(2)
)

func NewQNmeaPositionInfoSource(updateMode QNmeaPositionInfoSource__UpdateMode, parent core.QObjectITF) *QNmeaPositionInfoSource {
	return QNmeaPositionInfoSourceFromPointer(unsafe.Pointer(C.QNmeaPositionInfoSource_NewQNmeaPositionInfoSource(C.int(updateMode), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QNmeaPositionInfoSource) Device() *core.QIODevice {
	if ptr.Pointer() != nil {
		return core.QIODeviceFromPointer(unsafe.Pointer(C.QNmeaPositionInfoSource_Device(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QNmeaPositionInfoSource) Error() QGeoPositionInfoSource__Error {
	if ptr.Pointer() != nil {
		return QGeoPositionInfoSource__Error(C.QNmeaPositionInfoSource_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNmeaPositionInfoSource) MinimumUpdateInterval() int {
	if ptr.Pointer() != nil {
		return int(C.QNmeaPositionInfoSource_MinimumUpdateInterval(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNmeaPositionInfoSource) RequestUpdate(msec int) {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_RequestUpdate(C.QtObjectPtr(ptr.Pointer()), C.int(msec))
	}
}

func (ptr *QNmeaPositionInfoSource) SetDevice(device core.QIODeviceITF) {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_SetDevice(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQIODevice(device)))
	}
}

func (ptr *QNmeaPositionInfoSource) SetUpdateInterval(msec int) {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_SetUpdateInterval(C.QtObjectPtr(ptr.Pointer()), C.int(msec))
	}
}

func (ptr *QNmeaPositionInfoSource) StartUpdates() {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_StartUpdates(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QNmeaPositionInfoSource) StopUpdates() {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_StopUpdates(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QNmeaPositionInfoSource) SupportedPositioningMethods() QGeoPositionInfoSource__PositioningMethod {
	if ptr.Pointer() != nil {
		return QGeoPositionInfoSource__PositioningMethod(C.QNmeaPositionInfoSource_SupportedPositioningMethods(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNmeaPositionInfoSource) UpdateMode() QNmeaPositionInfoSource__UpdateMode {
	if ptr.Pointer() != nil {
		return QNmeaPositionInfoSource__UpdateMode(C.QNmeaPositionInfoSource_UpdateMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNmeaPositionInfoSource) DestroyQNmeaPositionInfoSource() {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_DestroyQNmeaPositionInfoSource(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
