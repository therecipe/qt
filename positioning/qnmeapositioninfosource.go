package positioning

//#include "positioning.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QNmeaPositionInfoSource struct {
	QGeoPositionInfoSource
}

type QNmeaPositionInfoSource_ITF interface {
	QGeoPositionInfoSource_ITF
	QNmeaPositionInfoSource_PTR() *QNmeaPositionInfoSource
}

func PointerFromQNmeaPositionInfoSource(ptr QNmeaPositionInfoSource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNmeaPositionInfoSource_PTR().Pointer()
	}
	return nil
}

func NewQNmeaPositionInfoSourceFromPointer(ptr unsafe.Pointer) *QNmeaPositionInfoSource {
	var n = new(QNmeaPositionInfoSource)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QNmeaPositionInfoSource_") {
		n.SetObjectName("QNmeaPositionInfoSource_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QNmeaPositionInfoSource) QNmeaPositionInfoSource_PTR() *QNmeaPositionInfoSource {
	return ptr
}

//QNmeaPositionInfoSource::UpdateMode
type QNmeaPositionInfoSource__UpdateMode int64

const (
	QNmeaPositionInfoSource__RealTimeMode   = QNmeaPositionInfoSource__UpdateMode(1)
	QNmeaPositionInfoSource__SimulationMode = QNmeaPositionInfoSource__UpdateMode(2)
)

func NewQNmeaPositionInfoSource(updateMode QNmeaPositionInfoSource__UpdateMode, parent core.QObject_ITF) *QNmeaPositionInfoSource {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNmeaPositionInfoSource::QNmeaPositionInfoSource")
		}
	}()

	return NewQNmeaPositionInfoSourceFromPointer(C.QNmeaPositionInfoSource_NewQNmeaPositionInfoSource(C.int(updateMode), core.PointerFromQObject(parent)))
}

func (ptr *QNmeaPositionInfoSource) Device() *core.QIODevice {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNmeaPositionInfoSource::device")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QNmeaPositionInfoSource_Device(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNmeaPositionInfoSource) Error() QGeoPositionInfoSource__Error {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNmeaPositionInfoSource::error")
		}
	}()

	if ptr.Pointer() != nil {
		return QGeoPositionInfoSource__Error(C.QNmeaPositionInfoSource_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNmeaPositionInfoSource) MinimumUpdateInterval() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNmeaPositionInfoSource::minimumUpdateInterval")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QNmeaPositionInfoSource_MinimumUpdateInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNmeaPositionInfoSource) RequestUpdate(msec int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNmeaPositionInfoSource::requestUpdate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_RequestUpdate(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QNmeaPositionInfoSource) SetDevice(device core.QIODevice_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNmeaPositionInfoSource::setDevice")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_SetDevice(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

func (ptr *QNmeaPositionInfoSource) SetUpdateInterval(msec int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNmeaPositionInfoSource::setUpdateInterval")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_SetUpdateInterval(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QNmeaPositionInfoSource) StartUpdates() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNmeaPositionInfoSource::startUpdates")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_StartUpdates(ptr.Pointer())
	}
}

func (ptr *QNmeaPositionInfoSource) StopUpdates() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNmeaPositionInfoSource::stopUpdates")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_StopUpdates(ptr.Pointer())
	}
}

func (ptr *QNmeaPositionInfoSource) SupportedPositioningMethods() QGeoPositionInfoSource__PositioningMethod {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNmeaPositionInfoSource::supportedPositioningMethods")
		}
	}()

	if ptr.Pointer() != nil {
		return QGeoPositionInfoSource__PositioningMethod(C.QNmeaPositionInfoSource_SupportedPositioningMethods(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNmeaPositionInfoSource) UpdateMode() QNmeaPositionInfoSource__UpdateMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNmeaPositionInfoSource::updateMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QNmeaPositionInfoSource__UpdateMode(C.QNmeaPositionInfoSource_UpdateMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNmeaPositionInfoSource) DestroyQNmeaPositionInfoSource() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNmeaPositionInfoSource::~QNmeaPositionInfoSource")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_DestroyQNmeaPositionInfoSource(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
