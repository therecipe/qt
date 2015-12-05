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

type QGeoSatelliteInfoSource struct {
	core.QObject
}

type QGeoSatelliteInfoSource_ITF interface {
	core.QObject_ITF
	QGeoSatelliteInfoSource_PTR() *QGeoSatelliteInfoSource
}

func PointerFromQGeoSatelliteInfoSource(ptr QGeoSatelliteInfoSource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoSatelliteInfoSource_PTR().Pointer()
	}
	return nil
}

func NewQGeoSatelliteInfoSourceFromPointer(ptr unsafe.Pointer) *QGeoSatelliteInfoSource {
	var n = new(QGeoSatelliteInfoSource)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGeoSatelliteInfoSource_") {
		n.SetObjectName("QGeoSatelliteInfoSource_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGeoSatelliteInfoSource) QGeoSatelliteInfoSource_PTR() *QGeoSatelliteInfoSource {
	return ptr
}

//QGeoSatelliteInfoSource::Error
type QGeoSatelliteInfoSource__Error int64

const (
	QGeoSatelliteInfoSource__AccessError        = QGeoSatelliteInfoSource__Error(0)
	QGeoSatelliteInfoSource__ClosedError        = QGeoSatelliteInfoSource__Error(1)
	QGeoSatelliteInfoSource__NoError            = QGeoSatelliteInfoSource__Error(2)
	QGeoSatelliteInfoSource__UnknownSourceError = QGeoSatelliteInfoSource__Error(-1)
)

func (ptr *QGeoSatelliteInfoSource) SetUpdateInterval(msec int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoSatelliteInfoSource::setUpdateInterval")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_SetUpdateInterval(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QGeoSatelliteInfoSource) UpdateInterval() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoSatelliteInfoSource::updateInterval")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QGeoSatelliteInfoSource_UpdateInterval(ptr.Pointer()))
	}
	return 0
}

func QGeoSatelliteInfoSource_AvailableSources() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoSatelliteInfoSource::availableSources")
		}
	}()

	return strings.Split(C.GoString(C.QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_AvailableSources()), ",,,")
}

func QGeoSatelliteInfoSource_CreateDefaultSource(parent core.QObject_ITF) *QGeoSatelliteInfoSource {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoSatelliteInfoSource::createDefaultSource")
		}
	}()

	return NewQGeoSatelliteInfoSourceFromPointer(C.QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateDefaultSource(core.PointerFromQObject(parent)))
}

func QGeoSatelliteInfoSource_CreateSource(sourceName string, parent core.QObject_ITF) *QGeoSatelliteInfoSource {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoSatelliteInfoSource::createSource")
		}
	}()

	return NewQGeoSatelliteInfoSourceFromPointer(C.QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateSource(C.CString(sourceName), core.PointerFromQObject(parent)))
}

func (ptr *QGeoSatelliteInfoSource) Error() QGeoSatelliteInfoSource__Error {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoSatelliteInfoSource::error")
		}
	}()

	if ptr.Pointer() != nil {
		return QGeoSatelliteInfoSource__Error(C.QGeoSatelliteInfoSource_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoSatelliteInfoSource) MinimumUpdateInterval() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoSatelliteInfoSource::minimumUpdateInterval")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QGeoSatelliteInfoSource_MinimumUpdateInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoSatelliteInfoSource) ConnectRequestTimeout(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoSatelliteInfoSource::requestTimeout")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_ConnectRequestTimeout(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "requestTimeout", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectRequestTimeout() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoSatelliteInfoSource::requestTimeout")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_DisconnectRequestTimeout(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "requestTimeout")
	}
}

//export callbackQGeoSatelliteInfoSourceRequestTimeout
func callbackQGeoSatelliteInfoSourceRequestTimeout(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoSatelliteInfoSource::requestTimeout")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "requestTimeout").(func())()
}

func (ptr *QGeoSatelliteInfoSource) RequestUpdate(timeout int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoSatelliteInfoSource::requestUpdate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_RequestUpdate(ptr.Pointer(), C.int(timeout))
	}
}

func (ptr *QGeoSatelliteInfoSource) SourceName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoSatelliteInfoSource::sourceName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoSatelliteInfoSource_SourceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoSatelliteInfoSource) StartUpdates() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoSatelliteInfoSource::startUpdates")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_StartUpdates(ptr.Pointer())
	}
}

func (ptr *QGeoSatelliteInfoSource) StopUpdates() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoSatelliteInfoSource::stopUpdates")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_StopUpdates(ptr.Pointer())
	}
}

func (ptr *QGeoSatelliteInfoSource) DestroyQGeoSatelliteInfoSource() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoSatelliteInfoSource::~QGeoSatelliteInfoSource")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_DestroyQGeoSatelliteInfoSource(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
