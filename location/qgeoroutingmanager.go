package location

//#include "location.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/positioning"
	"log"
	"unsafe"
)

type QGeoRoutingManager struct {
	core.QObject
}

type QGeoRoutingManager_ITF interface {
	core.QObject_ITF
	QGeoRoutingManager_PTR() *QGeoRoutingManager
}

func PointerFromQGeoRoutingManager(ptr QGeoRoutingManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoRoutingManager_PTR().Pointer()
	}
	return nil
}

func NewQGeoRoutingManagerFromPointer(ptr unsafe.Pointer) *QGeoRoutingManager {
	var n = new(QGeoRoutingManager)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGeoRoutingManager_") {
		n.SetObjectName("QGeoRoutingManager_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGeoRoutingManager) QGeoRoutingManager_PTR() *QGeoRoutingManager {
	return ptr
}

func (ptr *QGeoRoutingManager) CalculateRoute(request QGeoRouteRequest_ITF) *QGeoRouteReply {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRoutingManager::calculateRoute")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQGeoRouteReplyFromPointer(C.QGeoRoutingManager_CalculateRoute(ptr.Pointer(), PointerFromQGeoRouteRequest(request)))
	}
	return nil
}

func (ptr *QGeoRoutingManager) ConnectError(f func(reply *QGeoRouteReply, error QGeoRouteReply__Error, errorString string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRoutingManager::error")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoRoutingManager_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QGeoRoutingManager) DisconnectError() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRoutingManager::error")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoRoutingManager_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQGeoRoutingManagerError
func callbackQGeoRoutingManagerError(ptrName *C.char, reply unsafe.Pointer, error C.int, errorString *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRoutingManager::error")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "error").(func(*QGeoRouteReply, QGeoRouteReply__Error, string))(NewQGeoRouteReplyFromPointer(reply), QGeoRouteReply__Error(error), C.GoString(errorString))
}

func (ptr *QGeoRoutingManager) ConnectFinished(f func(reply *QGeoRouteReply)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRoutingManager::finished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoRoutingManager_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QGeoRoutingManager) DisconnectFinished() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRoutingManager::finished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoRoutingManager_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQGeoRoutingManagerFinished
func callbackQGeoRoutingManagerFinished(ptrName *C.char, reply unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRoutingManager::finished")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "finished").(func(*QGeoRouteReply))(NewQGeoRouteReplyFromPointer(reply))
}

func (ptr *QGeoRoutingManager) ManagerName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRoutingManager::managerName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoRoutingManager_ManagerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoRoutingManager) ManagerVersion() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRoutingManager::managerVersion")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QGeoRoutingManager_ManagerVersion(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoutingManager) MeasurementSystem() core.QLocale__MeasurementSystem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRoutingManager::measurementSystem")
		}
	}()

	if ptr.Pointer() != nil {
		return core.QLocale__MeasurementSystem(C.QGeoRoutingManager_MeasurementSystem(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoutingManager) SetLocale(locale core.QLocale_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRoutingManager::setLocale")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoRoutingManager_SetLocale(ptr.Pointer(), core.PointerFromQLocale(locale))
	}
}

func (ptr *QGeoRoutingManager) SetMeasurementSystem(system core.QLocale__MeasurementSystem) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRoutingManager::setMeasurementSystem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoRoutingManager_SetMeasurementSystem(ptr.Pointer(), C.int(system))
	}
}

func (ptr *QGeoRoutingManager) SupportedFeatureTypes() QGeoRouteRequest__FeatureType {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRoutingManager::supportedFeatureTypes")
		}
	}()

	if ptr.Pointer() != nil {
		return QGeoRouteRequest__FeatureType(C.QGeoRoutingManager_SupportedFeatureTypes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoutingManager) SupportedFeatureWeights() QGeoRouteRequest__FeatureWeight {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRoutingManager::supportedFeatureWeights")
		}
	}()

	if ptr.Pointer() != nil {
		return QGeoRouteRequest__FeatureWeight(C.QGeoRoutingManager_SupportedFeatureWeights(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoutingManager) SupportedManeuverDetails() QGeoRouteRequest__ManeuverDetail {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRoutingManager::supportedManeuverDetails")
		}
	}()

	if ptr.Pointer() != nil {
		return QGeoRouteRequest__ManeuverDetail(C.QGeoRoutingManager_SupportedManeuverDetails(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoutingManager) SupportedRouteOptimizations() QGeoRouteRequest__RouteOptimization {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRoutingManager::supportedRouteOptimizations")
		}
	}()

	if ptr.Pointer() != nil {
		return QGeoRouteRequest__RouteOptimization(C.QGeoRoutingManager_SupportedRouteOptimizations(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoutingManager) SupportedSegmentDetails() QGeoRouteRequest__SegmentDetail {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRoutingManager::supportedSegmentDetails")
		}
	}()

	if ptr.Pointer() != nil {
		return QGeoRouteRequest__SegmentDetail(C.QGeoRoutingManager_SupportedSegmentDetails(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoutingManager) SupportedTravelModes() QGeoRouteRequest__TravelMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRoutingManager::supportedTravelModes")
		}
	}()

	if ptr.Pointer() != nil {
		return QGeoRouteRequest__TravelMode(C.QGeoRoutingManager_SupportedTravelModes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoutingManager) UpdateRoute(route QGeoRoute_ITF, position positioning.QGeoCoordinate_ITF) *QGeoRouteReply {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRoutingManager::updateRoute")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQGeoRouteReplyFromPointer(C.QGeoRoutingManager_UpdateRoute(ptr.Pointer(), PointerFromQGeoRoute(route), positioning.PointerFromQGeoCoordinate(position)))
	}
	return nil
}

func (ptr *QGeoRoutingManager) DestroyQGeoRoutingManager() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRoutingManager::~QGeoRoutingManager")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoRoutingManager_DestroyQGeoRoutingManager(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
