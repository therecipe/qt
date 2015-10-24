package positioning

//#include "qgeoareamonitorsource.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QGeoAreaMonitorSource struct {
	core.QObject
}

type QGeoAreaMonitorSourceITF interface {
	core.QObjectITF
	QGeoAreaMonitorSourcePTR() *QGeoAreaMonitorSource
}

func PointerFromQGeoAreaMonitorSource(ptr QGeoAreaMonitorSourceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoAreaMonitorSourcePTR().Pointer()
	}
	return nil
}

func QGeoAreaMonitorSourceFromPointer(ptr unsafe.Pointer) *QGeoAreaMonitorSource {
	var n = new(QGeoAreaMonitorSource)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGeoAreaMonitorSource_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGeoAreaMonitorSource) QGeoAreaMonitorSourcePTR() *QGeoAreaMonitorSource {
	return ptr
}

//QGeoAreaMonitorSource::AreaMonitorFeature
type QGeoAreaMonitorSource__AreaMonitorFeature int

var (
	QGeoAreaMonitorSource__PersistentAreaMonitorFeature = QGeoAreaMonitorSource__AreaMonitorFeature(0x00000001)
	QGeoAreaMonitorSource__AnyAreaMonitorFeature        = QGeoAreaMonitorSource__AreaMonitorFeature(0xffffffff)
)

//QGeoAreaMonitorSource::Error
type QGeoAreaMonitorSource__Error int

var (
	QGeoAreaMonitorSource__AccessError              = QGeoAreaMonitorSource__Error(0)
	QGeoAreaMonitorSource__InsufficientPositionInfo = QGeoAreaMonitorSource__Error(1)
	QGeoAreaMonitorSource__UnknownSourceError       = QGeoAreaMonitorSource__Error(2)
	QGeoAreaMonitorSource__NoError                  = QGeoAreaMonitorSource__Error(3)
)

func QGeoAreaMonitorSource_AvailableSources() []string {
	return strings.Split(C.GoString(C.QGeoAreaMonitorSource_QGeoAreaMonitorSource_AvailableSources()), "|")
}

func QGeoAreaMonitorSource_CreateDefaultSource(parent core.QObjectITF) *QGeoAreaMonitorSource {
	return QGeoAreaMonitorSourceFromPointer(unsafe.Pointer(C.QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateDefaultSource(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func QGeoAreaMonitorSource_CreateSource(sourceName string, parent core.QObjectITF) *QGeoAreaMonitorSource {
	return QGeoAreaMonitorSourceFromPointer(unsafe.Pointer(C.QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateSource(C.CString(sourceName), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QGeoAreaMonitorSource) Error() QGeoAreaMonitorSource__Error {
	if ptr.Pointer() != nil {
		return QGeoAreaMonitorSource__Error(C.QGeoAreaMonitorSource_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoAreaMonitorSource) PositionInfoSource() *QGeoPositionInfoSource {
	if ptr.Pointer() != nil {
		return QGeoPositionInfoSourceFromPointer(unsafe.Pointer(C.QGeoAreaMonitorSource_PositionInfoSource(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QGeoAreaMonitorSource) RequestUpdate(monitor QGeoAreaMonitorInfoITF, signal string) bool {
	if ptr.Pointer() != nil {
		return C.QGeoAreaMonitorSource_RequestUpdate(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGeoAreaMonitorInfo(monitor)), C.CString(signal)) != 0
	}
	return false
}

func (ptr *QGeoAreaMonitorSource) SetPositionInfoSource(newSource QGeoPositionInfoSourceITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_SetPositionInfoSource(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGeoPositionInfoSource(newSource)))
	}
}

func (ptr *QGeoAreaMonitorSource) SourceName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAreaMonitorSource_SourceName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QGeoAreaMonitorSource) StartMonitoring(monitor QGeoAreaMonitorInfoITF) bool {
	if ptr.Pointer() != nil {
		return C.QGeoAreaMonitorSource_StartMonitoring(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGeoAreaMonitorInfo(monitor))) != 0
	}
	return false
}

func (ptr *QGeoAreaMonitorSource) StopMonitoring(monitor QGeoAreaMonitorInfoITF) bool {
	if ptr.Pointer() != nil {
		return C.QGeoAreaMonitorSource_StopMonitoring(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGeoAreaMonitorInfo(monitor))) != 0
	}
	return false
}

func (ptr *QGeoAreaMonitorSource) SupportedAreaMonitorFeatures() QGeoAreaMonitorSource__AreaMonitorFeature {
	if ptr.Pointer() != nil {
		return QGeoAreaMonitorSource__AreaMonitorFeature(C.QGeoAreaMonitorSource_SupportedAreaMonitorFeatures(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoAreaMonitorSource) DestroyQGeoAreaMonitorSource() {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_DestroyQGeoAreaMonitorSource(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
