package positioning

//#include "qgeosatelliteinfosource.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QGeoSatelliteInfoSource struct {
	core.QObject
}

type QGeoSatelliteInfoSourceITF interface {
	core.QObjectITF
	QGeoSatelliteInfoSourcePTR() *QGeoSatelliteInfoSource
}

func PointerFromQGeoSatelliteInfoSource(ptr QGeoSatelliteInfoSourceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoSatelliteInfoSourcePTR().Pointer()
	}
	return nil
}

func QGeoSatelliteInfoSourceFromPointer(ptr unsafe.Pointer) *QGeoSatelliteInfoSource {
	var n = new(QGeoSatelliteInfoSource)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGeoSatelliteInfoSource_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGeoSatelliteInfoSource) QGeoSatelliteInfoSourcePTR() *QGeoSatelliteInfoSource {
	return ptr
}

//QGeoSatelliteInfoSource::Error
type QGeoSatelliteInfoSource__Error int

var (
	QGeoSatelliteInfoSource__AccessError        = QGeoSatelliteInfoSource__Error(0)
	QGeoSatelliteInfoSource__ClosedError        = QGeoSatelliteInfoSource__Error(1)
	QGeoSatelliteInfoSource__NoError            = QGeoSatelliteInfoSource__Error(2)
	QGeoSatelliteInfoSource__UnknownSourceError = QGeoSatelliteInfoSource__Error(-1)
)

func (ptr *QGeoSatelliteInfoSource) SetUpdateInterval(msec int) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_SetUpdateInterval(C.QtObjectPtr(ptr.Pointer()), C.int(msec))
	}
}

func (ptr *QGeoSatelliteInfoSource) UpdateInterval() int {
	if ptr.Pointer() != nil {
		return int(C.QGeoSatelliteInfoSource_UpdateInterval(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func QGeoSatelliteInfoSource_AvailableSources() []string {
	return strings.Split(C.GoString(C.QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_AvailableSources()), "|")
}

func QGeoSatelliteInfoSource_CreateDefaultSource(parent core.QObjectITF) *QGeoSatelliteInfoSource {
	return QGeoSatelliteInfoSourceFromPointer(unsafe.Pointer(C.QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateDefaultSource(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func QGeoSatelliteInfoSource_CreateSource(sourceName string, parent core.QObjectITF) *QGeoSatelliteInfoSource {
	return QGeoSatelliteInfoSourceFromPointer(unsafe.Pointer(C.QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateSource(C.CString(sourceName), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QGeoSatelliteInfoSource) Error() QGeoSatelliteInfoSource__Error {
	if ptr.Pointer() != nil {
		return QGeoSatelliteInfoSource__Error(C.QGeoSatelliteInfoSource_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoSatelliteInfoSource) MinimumUpdateInterval() int {
	if ptr.Pointer() != nil {
		return int(C.QGeoSatelliteInfoSource_MinimumUpdateInterval(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoSatelliteInfoSource) ConnectRequestTimeout(f func()) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_ConnectRequestTimeout(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "requestTimeout", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectRequestTimeout() {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_DisconnectRequestTimeout(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "requestTimeout")
	}
}

//export callbackQGeoSatelliteInfoSourceRequestTimeout
func callbackQGeoSatelliteInfoSourceRequestTimeout(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "requestTimeout").(func())()
}

func (ptr *QGeoSatelliteInfoSource) RequestUpdate(timeout int) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_RequestUpdate(C.QtObjectPtr(ptr.Pointer()), C.int(timeout))
	}
}

func (ptr *QGeoSatelliteInfoSource) SourceName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoSatelliteInfoSource_SourceName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QGeoSatelliteInfoSource) StartUpdates() {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_StartUpdates(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QGeoSatelliteInfoSource) StopUpdates() {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_StopUpdates(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QGeoSatelliteInfoSource) DestroyQGeoSatelliteInfoSource() {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_DestroyQGeoSatelliteInfoSource(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
