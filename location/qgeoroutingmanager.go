package location

//#include "qgeoroutingmanager.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/positioning"
	"unsafe"
)

type QGeoRoutingManager struct {
	core.QObject
}

type QGeoRoutingManagerITF interface {
	core.QObjectITF
	QGeoRoutingManagerPTR() *QGeoRoutingManager
}

func PointerFromQGeoRoutingManager(ptr QGeoRoutingManagerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoRoutingManagerPTR().Pointer()
	}
	return nil
}

func QGeoRoutingManagerFromPointer(ptr unsafe.Pointer) *QGeoRoutingManager {
	var n = new(QGeoRoutingManager)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGeoRoutingManager_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGeoRoutingManager) QGeoRoutingManagerPTR() *QGeoRoutingManager {
	return ptr
}

func (ptr *QGeoRoutingManager) CalculateRoute(request QGeoRouteRequestITF) *QGeoRouteReply {
	if ptr.Pointer() != nil {
		return QGeoRouteReplyFromPointer(unsafe.Pointer(C.QGeoRoutingManager_CalculateRoute(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGeoRouteRequest(request)))))
	}
	return nil
}

func (ptr *QGeoRoutingManager) ConnectError(f func(reply QGeoRouteReplyITF, error QGeoRouteReply__Error, errorString string)) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManager_ConnectError(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QGeoRoutingManager) DisconnectError() {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManager_DisconnectError(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQGeoRoutingManagerError
func callbackQGeoRoutingManagerError(ptrName *C.char, reply unsafe.Pointer, error C.int, errorString *C.char) {
	qt.GetSignal(C.GoString(ptrName), "error").(func(*QGeoRouteReply, QGeoRouteReply__Error, string))(QGeoRouteReplyFromPointer(reply), QGeoRouteReply__Error(error), C.GoString(errorString))
}

func (ptr *QGeoRoutingManager) ConnectFinished(f func(reply QGeoRouteReplyITF)) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManager_ConnectFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QGeoRoutingManager) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManager_DisconnectFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQGeoRoutingManagerFinished
func callbackQGeoRoutingManagerFinished(ptrName *C.char, reply unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "finished").(func(*QGeoRouteReply))(QGeoRouteReplyFromPointer(reply))
}

func (ptr *QGeoRoutingManager) ManagerName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoRoutingManager_ManagerName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QGeoRoutingManager) ManagerVersion() int {
	if ptr.Pointer() != nil {
		return int(C.QGeoRoutingManager_ManagerVersion(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoRoutingManager) MeasurementSystem() core.QLocale__MeasurementSystem {
	if ptr.Pointer() != nil {
		return core.QLocale__MeasurementSystem(C.QGeoRoutingManager_MeasurementSystem(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoRoutingManager) SetLocale(locale core.QLocaleITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManager_SetLocale(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQLocale(locale)))
	}
}

func (ptr *QGeoRoutingManager) SetMeasurementSystem(system core.QLocale__MeasurementSystem) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManager_SetMeasurementSystem(C.QtObjectPtr(ptr.Pointer()), C.int(system))
	}
}

func (ptr *QGeoRoutingManager) SupportedFeatureTypes() QGeoRouteRequest__FeatureType {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__FeatureType(C.QGeoRoutingManager_SupportedFeatureTypes(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoRoutingManager) SupportedFeatureWeights() QGeoRouteRequest__FeatureWeight {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__FeatureWeight(C.QGeoRoutingManager_SupportedFeatureWeights(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoRoutingManager) SupportedManeuverDetails() QGeoRouteRequest__ManeuverDetail {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__ManeuverDetail(C.QGeoRoutingManager_SupportedManeuverDetails(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoRoutingManager) SupportedRouteOptimizations() QGeoRouteRequest__RouteOptimization {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__RouteOptimization(C.QGeoRoutingManager_SupportedRouteOptimizations(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoRoutingManager) SupportedSegmentDetails() QGeoRouteRequest__SegmentDetail {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__SegmentDetail(C.QGeoRoutingManager_SupportedSegmentDetails(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoRoutingManager) SupportedTravelModes() QGeoRouteRequest__TravelMode {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__TravelMode(C.QGeoRoutingManager_SupportedTravelModes(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoRoutingManager) UpdateRoute(route QGeoRouteITF, position positioning.QGeoCoordinateITF) *QGeoRouteReply {
	if ptr.Pointer() != nil {
		return QGeoRouteReplyFromPointer(unsafe.Pointer(C.QGeoRoutingManager_UpdateRoute(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGeoRoute(route)), C.QtObjectPtr(positioning.PointerFromQGeoCoordinate(position)))))
	}
	return nil
}

func (ptr *QGeoRoutingManager) DestroyQGeoRoutingManager() {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManager_DestroyQGeoRoutingManager(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
