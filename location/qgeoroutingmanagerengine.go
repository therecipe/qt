package location

//#include "location.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/positioning"
	"unsafe"
)

type QGeoRoutingManagerEngine struct {
	core.QObject
}

type QGeoRoutingManagerEngine_ITF interface {
	core.QObject_ITF
	QGeoRoutingManagerEngine_PTR() *QGeoRoutingManagerEngine
}

func PointerFromQGeoRoutingManagerEngine(ptr QGeoRoutingManagerEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoRoutingManagerEngine_PTR().Pointer()
	}
	return nil
}

func NewQGeoRoutingManagerEngineFromPointer(ptr unsafe.Pointer) *QGeoRoutingManagerEngine {
	var n = new(QGeoRoutingManagerEngine)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGeoRoutingManagerEngine_") {
		n.SetObjectName("QGeoRoutingManagerEngine_" + qt.Identifier())
	}
	return n
}

func (ptr *QGeoRoutingManagerEngine) QGeoRoutingManagerEngine_PTR() *QGeoRoutingManagerEngine {
	return ptr
}

func (ptr *QGeoRoutingManagerEngine) CalculateRoute(request QGeoRouteRequest_ITF) *QGeoRouteReply {
	defer qt.Recovering("QGeoRoutingManagerEngine::calculateRoute")

	if ptr.Pointer() != nil {
		return NewQGeoRouteReplyFromPointer(C.QGeoRoutingManagerEngine_CalculateRoute(ptr.Pointer(), PointerFromQGeoRouteRequest(request)))
	}
	return nil
}

func (ptr *QGeoRoutingManagerEngine) ConnectError(f func(reply *QGeoRouteReply, error QGeoRouteReply__Error, errorString string)) {
	defer qt.Recovering("connect QGeoRoutingManagerEngine::error")

	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QGeoRoutingManagerEngine) DisconnectError() {
	defer qt.Recovering("disconnect QGeoRoutingManagerEngine::error")

	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQGeoRoutingManagerEngineError
func callbackQGeoRoutingManagerEngineError(ptr unsafe.Pointer, ptrName *C.char, reply unsafe.Pointer, error C.int, errorString *C.char) {
	defer qt.Recovering("callback QGeoRoutingManagerEngine::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error"); signal != nil {
		signal.(func(*QGeoRouteReply, QGeoRouteReply__Error, string))(NewQGeoRouteReplyFromPointer(reply), QGeoRouteReply__Error(error), C.GoString(errorString))
	}

}

func (ptr *QGeoRoutingManagerEngine) Error(reply QGeoRouteReply_ITF, error QGeoRouteReply__Error, errorString string) {
	defer qt.Recovering("QGeoRoutingManagerEngine::error")

	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_Error(ptr.Pointer(), PointerFromQGeoRouteReply(reply), C.int(error), C.CString(errorString))
	}
}

func (ptr *QGeoRoutingManagerEngine) ConnectFinished(f func(reply *QGeoRouteReply)) {
	defer qt.Recovering("connect QGeoRoutingManagerEngine::finished")

	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QGeoRoutingManagerEngine) DisconnectFinished() {
	defer qt.Recovering("disconnect QGeoRoutingManagerEngine::finished")

	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQGeoRoutingManagerEngineFinished
func callbackQGeoRoutingManagerEngineFinished(ptr unsafe.Pointer, ptrName *C.char, reply unsafe.Pointer) {
	defer qt.Recovering("callback QGeoRoutingManagerEngine::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func(*QGeoRouteReply))(NewQGeoRouteReplyFromPointer(reply))
	}

}

func (ptr *QGeoRoutingManagerEngine) Finished(reply QGeoRouteReply_ITF) {
	defer qt.Recovering("QGeoRoutingManagerEngine::finished")

	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_Finished(ptr.Pointer(), PointerFromQGeoRouteReply(reply))
	}
}

func (ptr *QGeoRoutingManagerEngine) ManagerName() string {
	defer qt.Recovering("QGeoRoutingManagerEngine::managerName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoRoutingManagerEngine_ManagerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoRoutingManagerEngine) ManagerVersion() int {
	defer qt.Recovering("QGeoRoutingManagerEngine::managerVersion")

	if ptr.Pointer() != nil {
		return int(C.QGeoRoutingManagerEngine_ManagerVersion(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoutingManagerEngine) MeasurementSystem() core.QLocale__MeasurementSystem {
	defer qt.Recovering("QGeoRoutingManagerEngine::measurementSystem")

	if ptr.Pointer() != nil {
		return core.QLocale__MeasurementSystem(C.QGeoRoutingManagerEngine_MeasurementSystem(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoutingManagerEngine) SetLocale(locale core.QLocale_ITF) {
	defer qt.Recovering("QGeoRoutingManagerEngine::setLocale")

	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_SetLocale(ptr.Pointer(), core.PointerFromQLocale(locale))
	}
}

func (ptr *QGeoRoutingManagerEngine) SetMeasurementSystem(system core.QLocale__MeasurementSystem) {
	defer qt.Recovering("QGeoRoutingManagerEngine::setMeasurementSystem")

	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_SetMeasurementSystem(ptr.Pointer(), C.int(system))
	}
}

func (ptr *QGeoRoutingManagerEngine) SupportedFeatureTypes() QGeoRouteRequest__FeatureType {
	defer qt.Recovering("QGeoRoutingManagerEngine::supportedFeatureTypes")

	if ptr.Pointer() != nil {
		return QGeoRouteRequest__FeatureType(C.QGeoRoutingManagerEngine_SupportedFeatureTypes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoutingManagerEngine) SupportedFeatureWeights() QGeoRouteRequest__FeatureWeight {
	defer qt.Recovering("QGeoRoutingManagerEngine::supportedFeatureWeights")

	if ptr.Pointer() != nil {
		return QGeoRouteRequest__FeatureWeight(C.QGeoRoutingManagerEngine_SupportedFeatureWeights(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoutingManagerEngine) SupportedManeuverDetails() QGeoRouteRequest__ManeuverDetail {
	defer qt.Recovering("QGeoRoutingManagerEngine::supportedManeuverDetails")

	if ptr.Pointer() != nil {
		return QGeoRouteRequest__ManeuverDetail(C.QGeoRoutingManagerEngine_SupportedManeuverDetails(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoutingManagerEngine) SupportedRouteOptimizations() QGeoRouteRequest__RouteOptimization {
	defer qt.Recovering("QGeoRoutingManagerEngine::supportedRouteOptimizations")

	if ptr.Pointer() != nil {
		return QGeoRouteRequest__RouteOptimization(C.QGeoRoutingManagerEngine_SupportedRouteOptimizations(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoutingManagerEngine) SupportedSegmentDetails() QGeoRouteRequest__SegmentDetail {
	defer qt.Recovering("QGeoRoutingManagerEngine::supportedSegmentDetails")

	if ptr.Pointer() != nil {
		return QGeoRouteRequest__SegmentDetail(C.QGeoRoutingManagerEngine_SupportedSegmentDetails(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoutingManagerEngine) SupportedTravelModes() QGeoRouteRequest__TravelMode {
	defer qt.Recovering("QGeoRoutingManagerEngine::supportedTravelModes")

	if ptr.Pointer() != nil {
		return QGeoRouteRequest__TravelMode(C.QGeoRoutingManagerEngine_SupportedTravelModes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoutingManagerEngine) UpdateRoute(route QGeoRoute_ITF, position positioning.QGeoCoordinate_ITF) *QGeoRouteReply {
	defer qt.Recovering("QGeoRoutingManagerEngine::updateRoute")

	if ptr.Pointer() != nil {
		return NewQGeoRouteReplyFromPointer(C.QGeoRoutingManagerEngine_UpdateRoute(ptr.Pointer(), PointerFromQGeoRoute(route), positioning.PointerFromQGeoCoordinate(position)))
	}
	return nil
}

func (ptr *QGeoRoutingManagerEngine) DestroyQGeoRoutingManagerEngine() {
	defer qt.Recovering("QGeoRoutingManagerEngine::~QGeoRoutingManagerEngine")

	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_DestroyQGeoRoutingManagerEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoRoutingManagerEngine) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGeoRoutingManagerEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGeoRoutingManagerEngine) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGeoRoutingManagerEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGeoRoutingManagerEngineTimerEvent
func callbackQGeoRoutingManagerEngineTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGeoRoutingManagerEngine::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGeoRoutingManagerEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGeoRoutingManagerEngine) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGeoRoutingManagerEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGeoRoutingManagerEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGeoRoutingManagerEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGeoRoutingManagerEngine) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGeoRoutingManagerEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGeoRoutingManagerEngine) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGeoRoutingManagerEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGeoRoutingManagerEngineChildEvent
func callbackQGeoRoutingManagerEngineChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGeoRoutingManagerEngine::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGeoRoutingManagerEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGeoRoutingManagerEngine) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGeoRoutingManagerEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGeoRoutingManagerEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGeoRoutingManagerEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGeoRoutingManagerEngine) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGeoRoutingManagerEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGeoRoutingManagerEngine) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGeoRoutingManagerEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGeoRoutingManagerEngineCustomEvent
func callbackQGeoRoutingManagerEngineCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGeoRoutingManagerEngine::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGeoRoutingManagerEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGeoRoutingManagerEngine) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGeoRoutingManagerEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGeoRoutingManagerEngine) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGeoRoutingManagerEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
