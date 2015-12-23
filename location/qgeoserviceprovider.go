package location

//#include "location.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QGeoServiceProvider struct {
	core.QObject
}

type QGeoServiceProvider_ITF interface {
	core.QObject_ITF
	QGeoServiceProvider_PTR() *QGeoServiceProvider
}

func PointerFromQGeoServiceProvider(ptr QGeoServiceProvider_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoServiceProvider_PTR().Pointer()
	}
	return nil
}

func NewQGeoServiceProviderFromPointer(ptr unsafe.Pointer) *QGeoServiceProvider {
	var n = new(QGeoServiceProvider)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGeoServiceProvider_") {
		n.SetObjectName("QGeoServiceProvider_" + qt.Identifier())
	}
	return n
}

func (ptr *QGeoServiceProvider) QGeoServiceProvider_PTR() *QGeoServiceProvider {
	return ptr
}

//QGeoServiceProvider::Error
type QGeoServiceProvider__Error int64

const (
	QGeoServiceProvider__NoError                       = QGeoServiceProvider__Error(0)
	QGeoServiceProvider__NotSupportedError             = QGeoServiceProvider__Error(1)
	QGeoServiceProvider__UnknownParameterError         = QGeoServiceProvider__Error(2)
	QGeoServiceProvider__MissingRequiredParameterError = QGeoServiceProvider__Error(3)
	QGeoServiceProvider__ConnectionError               = QGeoServiceProvider__Error(4)
)

//QGeoServiceProvider::GeocodingFeature
type QGeoServiceProvider__GeocodingFeature int64

var (
	QGeoServiceProvider__NoGeocodingFeatures       = QGeoServiceProvider__GeocodingFeature(0)
	QGeoServiceProvider__OnlineGeocodingFeature    = QGeoServiceProvider__GeocodingFeature(C.QGeoServiceProvider_OnlineGeocodingFeature_Type())
	QGeoServiceProvider__OfflineGeocodingFeature   = QGeoServiceProvider__GeocodingFeature(C.QGeoServiceProvider_OfflineGeocodingFeature_Type())
	QGeoServiceProvider__ReverseGeocodingFeature   = QGeoServiceProvider__GeocodingFeature(C.QGeoServiceProvider_ReverseGeocodingFeature_Type())
	QGeoServiceProvider__LocalizedGeocodingFeature = QGeoServiceProvider__GeocodingFeature(C.QGeoServiceProvider_LocalizedGeocodingFeature_Type())
	QGeoServiceProvider__AnyGeocodingFeatures      = QGeoServiceProvider__GeocodingFeature(C.QGeoServiceProvider_AnyGeocodingFeatures_Type())
)

//QGeoServiceProvider::MappingFeature
type QGeoServiceProvider__MappingFeature int64

var (
	QGeoServiceProvider__NoMappingFeatures       = QGeoServiceProvider__MappingFeature(0)
	QGeoServiceProvider__OnlineMappingFeature    = QGeoServiceProvider__MappingFeature(C.QGeoServiceProvider_OnlineMappingFeature_Type())
	QGeoServiceProvider__OfflineMappingFeature   = QGeoServiceProvider__MappingFeature(C.QGeoServiceProvider_OfflineMappingFeature_Type())
	QGeoServiceProvider__LocalizedMappingFeature = QGeoServiceProvider__MappingFeature(C.QGeoServiceProvider_LocalizedMappingFeature_Type())
	QGeoServiceProvider__AnyMappingFeatures      = QGeoServiceProvider__MappingFeature(C.QGeoServiceProvider_AnyMappingFeatures_Type())
)

//QGeoServiceProvider::PlacesFeature
type QGeoServiceProvider__PlacesFeature int64

var (
	QGeoServiceProvider__NoPlacesFeatures            = QGeoServiceProvider__PlacesFeature(0)
	QGeoServiceProvider__OnlinePlacesFeature         = QGeoServiceProvider__PlacesFeature(C.QGeoServiceProvider_OnlinePlacesFeature_Type())
	QGeoServiceProvider__OfflinePlacesFeature        = QGeoServiceProvider__PlacesFeature(C.QGeoServiceProvider_OfflinePlacesFeature_Type())
	QGeoServiceProvider__SavePlaceFeature            = QGeoServiceProvider__PlacesFeature(C.QGeoServiceProvider_SavePlaceFeature_Type())
	QGeoServiceProvider__RemovePlaceFeature          = QGeoServiceProvider__PlacesFeature(C.QGeoServiceProvider_RemovePlaceFeature_Type())
	QGeoServiceProvider__SaveCategoryFeature         = QGeoServiceProvider__PlacesFeature(C.QGeoServiceProvider_SaveCategoryFeature_Type())
	QGeoServiceProvider__RemoveCategoryFeature       = QGeoServiceProvider__PlacesFeature(C.QGeoServiceProvider_RemoveCategoryFeature_Type())
	QGeoServiceProvider__PlaceRecommendationsFeature = QGeoServiceProvider__PlacesFeature(C.QGeoServiceProvider_PlaceRecommendationsFeature_Type())
	QGeoServiceProvider__SearchSuggestionsFeature    = QGeoServiceProvider__PlacesFeature(C.QGeoServiceProvider_SearchSuggestionsFeature_Type())
	QGeoServiceProvider__LocalizedPlacesFeature      = QGeoServiceProvider__PlacesFeature(C.QGeoServiceProvider_LocalizedPlacesFeature_Type())
	QGeoServiceProvider__NotificationsFeature        = QGeoServiceProvider__PlacesFeature(C.QGeoServiceProvider_NotificationsFeature_Type())
	QGeoServiceProvider__PlaceMatchingFeature        = QGeoServiceProvider__PlacesFeature(C.QGeoServiceProvider_PlaceMatchingFeature_Type())
	QGeoServiceProvider__AnyPlacesFeatures           = QGeoServiceProvider__PlacesFeature(C.QGeoServiceProvider_AnyPlacesFeatures_Type())
)

//QGeoServiceProvider::RoutingFeature
type QGeoServiceProvider__RoutingFeature int64

var (
	QGeoServiceProvider__NoRoutingFeatures          = QGeoServiceProvider__RoutingFeature(0)
	QGeoServiceProvider__OnlineRoutingFeature       = QGeoServiceProvider__RoutingFeature(C.QGeoServiceProvider_OnlineRoutingFeature_Type())
	QGeoServiceProvider__OfflineRoutingFeature      = QGeoServiceProvider__RoutingFeature(C.QGeoServiceProvider_OfflineRoutingFeature_Type())
	QGeoServiceProvider__LocalizedRoutingFeature    = QGeoServiceProvider__RoutingFeature(C.QGeoServiceProvider_LocalizedRoutingFeature_Type())
	QGeoServiceProvider__RouteUpdatesFeature        = QGeoServiceProvider__RoutingFeature(C.QGeoServiceProvider_RouteUpdatesFeature_Type())
	QGeoServiceProvider__AlternativeRoutesFeature   = QGeoServiceProvider__RoutingFeature(C.QGeoServiceProvider_AlternativeRoutesFeature_Type())
	QGeoServiceProvider__ExcludeAreasRoutingFeature = QGeoServiceProvider__RoutingFeature(C.QGeoServiceProvider_ExcludeAreasRoutingFeature_Type())
	QGeoServiceProvider__AnyRoutingFeatures         = QGeoServiceProvider__RoutingFeature(C.QGeoServiceProvider_AnyRoutingFeatures_Type())
)

func QGeoServiceProvider_AvailableServiceProviders() []string {
	defer qt.Recovering("QGeoServiceProvider::availableServiceProviders")

	return strings.Split(C.GoString(C.QGeoServiceProvider_QGeoServiceProvider_AvailableServiceProviders()), ",,,")
}

func (ptr *QGeoServiceProvider) Error() QGeoServiceProvider__Error {
	defer qt.Recovering("QGeoServiceProvider::error")

	if ptr.Pointer() != nil {
		return QGeoServiceProvider__Error(C.QGeoServiceProvider_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoServiceProvider) ErrorString() string {
	defer qt.Recovering("QGeoServiceProvider::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoServiceProvider_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoServiceProvider) GeocodingFeatures() QGeoServiceProvider__GeocodingFeature {
	defer qt.Recovering("QGeoServiceProvider::geocodingFeatures")

	if ptr.Pointer() != nil {
		return QGeoServiceProvider__GeocodingFeature(C.QGeoServiceProvider_GeocodingFeatures(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoServiceProvider) GeocodingManager() *QGeoCodingManager {
	defer qt.Recovering("QGeoServiceProvider::geocodingManager")

	if ptr.Pointer() != nil {
		return NewQGeoCodingManagerFromPointer(C.QGeoServiceProvider_GeocodingManager(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGeoServiceProvider) MappingFeatures() QGeoServiceProvider__MappingFeature {
	defer qt.Recovering("QGeoServiceProvider::mappingFeatures")

	if ptr.Pointer() != nil {
		return QGeoServiceProvider__MappingFeature(C.QGeoServiceProvider_MappingFeatures(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoServiceProvider) PlaceManager() *QPlaceManager {
	defer qt.Recovering("QGeoServiceProvider::placeManager")

	if ptr.Pointer() != nil {
		return NewQPlaceManagerFromPointer(C.QGeoServiceProvider_PlaceManager(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGeoServiceProvider) PlacesFeatures() QGeoServiceProvider__PlacesFeature {
	defer qt.Recovering("QGeoServiceProvider::placesFeatures")

	if ptr.Pointer() != nil {
		return QGeoServiceProvider__PlacesFeature(C.QGeoServiceProvider_PlacesFeatures(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoServiceProvider) RoutingFeatures() QGeoServiceProvider__RoutingFeature {
	defer qt.Recovering("QGeoServiceProvider::routingFeatures")

	if ptr.Pointer() != nil {
		return QGeoServiceProvider__RoutingFeature(C.QGeoServiceProvider_RoutingFeatures(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoServiceProvider) RoutingManager() *QGeoRoutingManager {
	defer qt.Recovering("QGeoServiceProvider::routingManager")

	if ptr.Pointer() != nil {
		return NewQGeoRoutingManagerFromPointer(C.QGeoServiceProvider_RoutingManager(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGeoServiceProvider) SetAllowExperimental(allow bool) {
	defer qt.Recovering("QGeoServiceProvider::setAllowExperimental")

	if ptr.Pointer() != nil {
		C.QGeoServiceProvider_SetAllowExperimental(ptr.Pointer(), C.int(qt.GoBoolToInt(allow)))
	}
}

func (ptr *QGeoServiceProvider) SetLocale(locale core.QLocale_ITF) {
	defer qt.Recovering("QGeoServiceProvider::setLocale")

	if ptr.Pointer() != nil {
		C.QGeoServiceProvider_SetLocale(ptr.Pointer(), core.PointerFromQLocale(locale))
	}
}

func (ptr *QGeoServiceProvider) DestroyQGeoServiceProvider() {
	defer qt.Recovering("QGeoServiceProvider::~QGeoServiceProvider")

	if ptr.Pointer() != nil {
		C.QGeoServiceProvider_DestroyQGeoServiceProvider(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoServiceProvider) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGeoServiceProvider::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGeoServiceProvider) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGeoServiceProvider::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGeoServiceProviderTimerEvent
func callbackQGeoServiceProviderTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGeoServiceProvider::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGeoServiceProvider) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGeoServiceProvider::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGeoServiceProvider) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGeoServiceProvider::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGeoServiceProviderChildEvent
func callbackQGeoServiceProviderChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGeoServiceProvider::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGeoServiceProvider) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGeoServiceProvider::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGeoServiceProvider) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGeoServiceProvider::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGeoServiceProviderCustomEvent
func callbackQGeoServiceProviderCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGeoServiceProvider::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
