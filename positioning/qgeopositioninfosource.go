package positioning

//#include "qgeopositioninfosource.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QGeoPositionInfoSource struct {
	core.QObject
}

type QGeoPositionInfoSourceITF interface {
	core.QObjectITF
	QGeoPositionInfoSourcePTR() *QGeoPositionInfoSource
}

func PointerFromQGeoPositionInfoSource(ptr QGeoPositionInfoSourceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoPositionInfoSourcePTR().Pointer()
	}
	return nil
}

func QGeoPositionInfoSourceFromPointer(ptr unsafe.Pointer) *QGeoPositionInfoSource {
	var n = new(QGeoPositionInfoSource)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGeoPositionInfoSource_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGeoPositionInfoSource) QGeoPositionInfoSourcePTR() *QGeoPositionInfoSource {
	return ptr
}

//QGeoPositionInfoSource::Error
type QGeoPositionInfoSource__Error int

var (
	QGeoPositionInfoSource__AccessError        = QGeoPositionInfoSource__Error(0)
	QGeoPositionInfoSource__ClosedError        = QGeoPositionInfoSource__Error(1)
	QGeoPositionInfoSource__UnknownSourceError = QGeoPositionInfoSource__Error(2)
	QGeoPositionInfoSource__NoError            = QGeoPositionInfoSource__Error(3)
)

//QGeoPositionInfoSource::PositioningMethod
type QGeoPositionInfoSource__PositioningMethod int

var (
	QGeoPositionInfoSource__NoPositioningMethods           = QGeoPositionInfoSource__PositioningMethod(0x00000000)
	QGeoPositionInfoSource__SatellitePositioningMethods    = QGeoPositionInfoSource__PositioningMethod(0x000000ff)
	QGeoPositionInfoSource__NonSatellitePositioningMethods = QGeoPositionInfoSource__PositioningMethod(0xffffff00)
	QGeoPositionInfoSource__AllPositioningMethods          = QGeoPositionInfoSource__PositioningMethod(0xffffffff)
)

func (ptr *QGeoPositionInfoSource) SetUpdateInterval(msec int) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_SetUpdateInterval(C.QtObjectPtr(ptr.Pointer()), C.int(msec))
	}
}

func (ptr *QGeoPositionInfoSource) SourceName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoPositionInfoSource_SourceName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QGeoPositionInfoSource) UpdateInterval() int {
	if ptr.Pointer() != nil {
		return int(C.QGeoPositionInfoSource_UpdateInterval(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func QGeoPositionInfoSource_AvailableSources() []string {
	return strings.Split(C.GoString(C.QGeoPositionInfoSource_QGeoPositionInfoSource_AvailableSources()), "|")
}

func QGeoPositionInfoSource_CreateDefaultSource(parent core.QObjectITF) *QGeoPositionInfoSource {
	return QGeoPositionInfoSourceFromPointer(unsafe.Pointer(C.QGeoPositionInfoSource_QGeoPositionInfoSource_CreateDefaultSource(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func QGeoPositionInfoSource_CreateSource(sourceName string, parent core.QObjectITF) *QGeoPositionInfoSource {
	return QGeoPositionInfoSourceFromPointer(unsafe.Pointer(C.QGeoPositionInfoSource_QGeoPositionInfoSource_CreateSource(C.CString(sourceName), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QGeoPositionInfoSource) Error() QGeoPositionInfoSource__Error {
	if ptr.Pointer() != nil {
		return QGeoPositionInfoSource__Error(C.QGeoPositionInfoSource_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoPositionInfoSource) MinimumUpdateInterval() int {
	if ptr.Pointer() != nil {
		return int(C.QGeoPositionInfoSource_MinimumUpdateInterval(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoPositionInfoSource) PreferredPositioningMethods() QGeoPositionInfoSource__PositioningMethod {
	if ptr.Pointer() != nil {
		return QGeoPositionInfoSource__PositioningMethod(C.QGeoPositionInfoSource_PreferredPositioningMethods(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoPositionInfoSource) RequestUpdate(timeout int) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_RequestUpdate(C.QtObjectPtr(ptr.Pointer()), C.int(timeout))
	}
}

func (ptr *QGeoPositionInfoSource) SetPreferredPositioningMethods(methods QGeoPositionInfoSource__PositioningMethod) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_SetPreferredPositioningMethods(C.QtObjectPtr(ptr.Pointer()), C.int(methods))
	}
}

func (ptr *QGeoPositionInfoSource) StartUpdates() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_StartUpdates(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QGeoPositionInfoSource) StopUpdates() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_StopUpdates(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QGeoPositionInfoSource) SupportedPositioningMethods() QGeoPositionInfoSource__PositioningMethod {
	if ptr.Pointer() != nil {
		return QGeoPositionInfoSource__PositioningMethod(C.QGeoPositionInfoSource_SupportedPositioningMethods(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoPositionInfoSource) ConnectUpdateTimeout(f func()) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_ConnectUpdateTimeout(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "updateTimeout", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectUpdateTimeout() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_DisconnectUpdateTimeout(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "updateTimeout")
	}
}

//export callbackQGeoPositionInfoSourceUpdateTimeout
func callbackQGeoPositionInfoSourceUpdateTimeout(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "updateTimeout").(func())()
}

func (ptr *QGeoPositionInfoSource) DestroyQGeoPositionInfoSource() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_DestroyQGeoPositionInfoSource(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
