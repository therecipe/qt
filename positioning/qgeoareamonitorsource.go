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

type QGeoAreaMonitorSource struct {
	core.QObject
}

type QGeoAreaMonitorSource_ITF interface {
	core.QObject_ITF
	QGeoAreaMonitorSource_PTR() *QGeoAreaMonitorSource
}

func PointerFromQGeoAreaMonitorSource(ptr QGeoAreaMonitorSource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoAreaMonitorSource_PTR().Pointer()
	}
	return nil
}

func NewQGeoAreaMonitorSourceFromPointer(ptr unsafe.Pointer) *QGeoAreaMonitorSource {
	var n = new(QGeoAreaMonitorSource)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGeoAreaMonitorSource_") {
		n.SetObjectName("QGeoAreaMonitorSource_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGeoAreaMonitorSource) QGeoAreaMonitorSource_PTR() *QGeoAreaMonitorSource {
	return ptr
}

//QGeoAreaMonitorSource::AreaMonitorFeature
type QGeoAreaMonitorSource__AreaMonitorFeature int64

const (
	QGeoAreaMonitorSource__PersistentAreaMonitorFeature = QGeoAreaMonitorSource__AreaMonitorFeature(0x00000001)
	QGeoAreaMonitorSource__AnyAreaMonitorFeature        = QGeoAreaMonitorSource__AreaMonitorFeature(0xffffffff)
)

//QGeoAreaMonitorSource::Error
type QGeoAreaMonitorSource__Error int64

const (
	QGeoAreaMonitorSource__AccessError              = QGeoAreaMonitorSource__Error(0)
	QGeoAreaMonitorSource__InsufficientPositionInfo = QGeoAreaMonitorSource__Error(1)
	QGeoAreaMonitorSource__UnknownSourceError       = QGeoAreaMonitorSource__Error(2)
	QGeoAreaMonitorSource__NoError                  = QGeoAreaMonitorSource__Error(3)
)

func QGeoAreaMonitorSource_AvailableSources() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAreaMonitorSource::availableSources")
		}
	}()

	return strings.Split(C.GoString(C.QGeoAreaMonitorSource_QGeoAreaMonitorSource_AvailableSources()), ",,,")
}

func QGeoAreaMonitorSource_CreateDefaultSource(parent core.QObject_ITF) *QGeoAreaMonitorSource {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAreaMonitorSource::createDefaultSource")
		}
	}()

	return NewQGeoAreaMonitorSourceFromPointer(C.QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateDefaultSource(core.PointerFromQObject(parent)))
}

func QGeoAreaMonitorSource_CreateSource(sourceName string, parent core.QObject_ITF) *QGeoAreaMonitorSource {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAreaMonitorSource::createSource")
		}
	}()

	return NewQGeoAreaMonitorSourceFromPointer(C.QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateSource(C.CString(sourceName), core.PointerFromQObject(parent)))
}

func (ptr *QGeoAreaMonitorSource) Error() QGeoAreaMonitorSource__Error {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAreaMonitorSource::error")
		}
	}()

	if ptr.Pointer() != nil {
		return QGeoAreaMonitorSource__Error(C.QGeoAreaMonitorSource_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoAreaMonitorSource) PositionInfoSource() *QGeoPositionInfoSource {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAreaMonitorSource::positionInfoSource")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQGeoPositionInfoSourceFromPointer(C.QGeoAreaMonitorSource_PositionInfoSource(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGeoAreaMonitorSource) RequestUpdate(monitor QGeoAreaMonitorInfo_ITF, signal string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAreaMonitorSource::requestUpdate")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QGeoAreaMonitorSource_RequestUpdate(ptr.Pointer(), PointerFromQGeoAreaMonitorInfo(monitor), C.CString(signal)) != 0
	}
	return false
}

func (ptr *QGeoAreaMonitorSource) SetPositionInfoSource(newSource QGeoPositionInfoSource_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAreaMonitorSource::setPositionInfoSource")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_SetPositionInfoSource(ptr.Pointer(), PointerFromQGeoPositionInfoSource(newSource))
	}
}

func (ptr *QGeoAreaMonitorSource) SourceName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAreaMonitorSource::sourceName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAreaMonitorSource_SourceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAreaMonitorSource) StartMonitoring(monitor QGeoAreaMonitorInfo_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAreaMonitorSource::startMonitoring")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QGeoAreaMonitorSource_StartMonitoring(ptr.Pointer(), PointerFromQGeoAreaMonitorInfo(monitor)) != 0
	}
	return false
}

func (ptr *QGeoAreaMonitorSource) StopMonitoring(monitor QGeoAreaMonitorInfo_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAreaMonitorSource::stopMonitoring")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QGeoAreaMonitorSource_StopMonitoring(ptr.Pointer(), PointerFromQGeoAreaMonitorInfo(monitor)) != 0
	}
	return false
}

func (ptr *QGeoAreaMonitorSource) SupportedAreaMonitorFeatures() QGeoAreaMonitorSource__AreaMonitorFeature {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAreaMonitorSource::supportedAreaMonitorFeatures")
		}
	}()

	if ptr.Pointer() != nil {
		return QGeoAreaMonitorSource__AreaMonitorFeature(C.QGeoAreaMonitorSource_SupportedAreaMonitorFeatures(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoAreaMonitorSource) DestroyQGeoAreaMonitorSource() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAreaMonitorSource::~QGeoAreaMonitorSource")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_DestroyQGeoAreaMonitorSource(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
