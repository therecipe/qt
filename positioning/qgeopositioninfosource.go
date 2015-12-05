package positioning

//#include "positioning.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"strings"
	"unsafe"
)

type QGeoPositionInfoSource struct {
	core.QObject
}

type QGeoPositionInfoSource_ITF interface {
	core.QObject_ITF
	QGeoPositionInfoSource_PTR() *QGeoPositionInfoSource
}

func PointerFromQGeoPositionInfoSource(ptr QGeoPositionInfoSource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoPositionInfoSource_PTR().Pointer()
	}
	return nil
}

func NewQGeoPositionInfoSourceFromPointer(ptr unsafe.Pointer) *QGeoPositionInfoSource {
	var n = new(QGeoPositionInfoSource)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGeoPositionInfoSource_") {
		n.SetObjectName("QGeoPositionInfoSource_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGeoPositionInfoSource) QGeoPositionInfoSource_PTR() *QGeoPositionInfoSource {
	return ptr
}

//QGeoPositionInfoSource::Error
type QGeoPositionInfoSource__Error int64

const (
	QGeoPositionInfoSource__AccessError        = QGeoPositionInfoSource__Error(0)
	QGeoPositionInfoSource__ClosedError        = QGeoPositionInfoSource__Error(1)
	QGeoPositionInfoSource__UnknownSourceError = QGeoPositionInfoSource__Error(2)
	QGeoPositionInfoSource__NoError            = QGeoPositionInfoSource__Error(3)
)

//QGeoPositionInfoSource::PositioningMethod
type QGeoPositionInfoSource__PositioningMethod int64

const (
	QGeoPositionInfoSource__NoPositioningMethods           = QGeoPositionInfoSource__PositioningMethod(0x00000000)
	QGeoPositionInfoSource__SatellitePositioningMethods    = QGeoPositionInfoSource__PositioningMethod(0x000000ff)
	QGeoPositionInfoSource__NonSatellitePositioningMethods = QGeoPositionInfoSource__PositioningMethod(0xffffff00)
	QGeoPositionInfoSource__AllPositioningMethods          = QGeoPositionInfoSource__PositioningMethod(0xffffffff)
)

func (ptr *QGeoPositionInfoSource) SetUpdateInterval(msec int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoPositionInfoSource::setUpdateInterval")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_SetUpdateInterval(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QGeoPositionInfoSource) SourceName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoPositionInfoSource::sourceName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoPositionInfoSource_SourceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoPositionInfoSource) UpdateInterval() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoPositionInfoSource::updateInterval")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QGeoPositionInfoSource_UpdateInterval(ptr.Pointer()))
	}
	return 0
}

func QGeoPositionInfoSource_AvailableSources() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoPositionInfoSource::availableSources")
		}
	}()

	return strings.Split(C.GoString(C.QGeoPositionInfoSource_QGeoPositionInfoSource_AvailableSources()), ",,,")
}

func QGeoPositionInfoSource_CreateDefaultSource(parent core.QObject_ITF) *QGeoPositionInfoSource {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoPositionInfoSource::createDefaultSource")
		}
	}()

	return NewQGeoPositionInfoSourceFromPointer(C.QGeoPositionInfoSource_QGeoPositionInfoSource_CreateDefaultSource(core.PointerFromQObject(parent)))
}

func QGeoPositionInfoSource_CreateSource(sourceName string, parent core.QObject_ITF) *QGeoPositionInfoSource {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoPositionInfoSource::createSource")
		}
	}()

	return NewQGeoPositionInfoSourceFromPointer(C.QGeoPositionInfoSource_QGeoPositionInfoSource_CreateSource(C.CString(sourceName), core.PointerFromQObject(parent)))
}

func (ptr *QGeoPositionInfoSource) Error() QGeoPositionInfoSource__Error {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoPositionInfoSource::error")
		}
	}()

	if ptr.Pointer() != nil {
		return QGeoPositionInfoSource__Error(C.QGeoPositionInfoSource_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoPositionInfoSource) MinimumUpdateInterval() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoPositionInfoSource::minimumUpdateInterval")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QGeoPositionInfoSource_MinimumUpdateInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoPositionInfoSource) PreferredPositioningMethods() QGeoPositionInfoSource__PositioningMethod {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoPositionInfoSource::preferredPositioningMethods")
		}
	}()

	if ptr.Pointer() != nil {
		return QGeoPositionInfoSource__PositioningMethod(C.QGeoPositionInfoSource_PreferredPositioningMethods(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoPositionInfoSource) RequestUpdate(timeout int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoPositionInfoSource::requestUpdate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_RequestUpdate(ptr.Pointer(), C.int(timeout))
	}
}

func (ptr *QGeoPositionInfoSource) SetPreferredPositioningMethods(methods QGeoPositionInfoSource__PositioningMethod) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoPositionInfoSource::setPreferredPositioningMethods")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_SetPreferredPositioningMethods(ptr.Pointer(), C.int(methods))
	}
}

func (ptr *QGeoPositionInfoSource) StartUpdates() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoPositionInfoSource::startUpdates")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_StartUpdates(ptr.Pointer())
	}
}

func (ptr *QGeoPositionInfoSource) StopUpdates() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoPositionInfoSource::stopUpdates")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_StopUpdates(ptr.Pointer())
	}
}

func (ptr *QGeoPositionInfoSource) SupportedPositioningMethods() QGeoPositionInfoSource__PositioningMethod {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoPositionInfoSource::supportedPositioningMethods")
		}
	}()

	if ptr.Pointer() != nil {
		return QGeoPositionInfoSource__PositioningMethod(C.QGeoPositionInfoSource_SupportedPositioningMethods(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoPositionInfoSource) ConnectUpdateTimeout(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoPositionInfoSource::updateTimeout")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_ConnectUpdateTimeout(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "updateTimeout", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectUpdateTimeout() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoPositionInfoSource::updateTimeout")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_DisconnectUpdateTimeout(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "updateTimeout")
	}
}

//export callbackQGeoPositionInfoSourceUpdateTimeout
func callbackQGeoPositionInfoSourceUpdateTimeout(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoPositionInfoSource::updateTimeout")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "updateTimeout").(func())()
}

func (ptr *QGeoPositionInfoSource) DestroyQGeoPositionInfoSource() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoPositionInfoSource::~QGeoPositionInfoSource")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_DestroyQGeoPositionInfoSource(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
