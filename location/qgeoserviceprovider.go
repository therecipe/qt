package location

//#include "qgeoserviceprovider.h"
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

type QGeoServiceProviderITF interface {
	core.QObjectITF
	QGeoServiceProviderPTR() *QGeoServiceProvider
}

func PointerFromQGeoServiceProvider(ptr QGeoServiceProviderITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoServiceProviderPTR().Pointer()
	}
	return nil
}

func QGeoServiceProviderFromPointer(ptr unsafe.Pointer) *QGeoServiceProvider {
	var n = new(QGeoServiceProvider)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGeoServiceProvider_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGeoServiceProvider) QGeoServiceProviderPTR() *QGeoServiceProvider {
	return ptr
}

//QGeoServiceProvider::Error
type QGeoServiceProvider__Error int

var (
	QGeoServiceProvider__NoError                       = QGeoServiceProvider__Error(0)
	QGeoServiceProvider__NotSupportedError             = QGeoServiceProvider__Error(1)
	QGeoServiceProvider__UnknownParameterError         = QGeoServiceProvider__Error(2)
	QGeoServiceProvider__MissingRequiredParameterError = QGeoServiceProvider__Error(3)
	QGeoServiceProvider__ConnectionError               = QGeoServiceProvider__Error(4)
)

//QGeoServiceProvider::GeocodingFeature
type QGeoServiceProvider__GeocodingFeature int

var (
	QGeoServiceProvider__NoGeocodingFeatures       = QGeoServiceProvider__GeocodingFeature(0)
	QGeoServiceProvider__OnlineGeocodingFeature    = QGeoServiceProvider__GeocodingFeature(C.QGeoServiceProvider_OnlineGeocodingFeature_Type())
	QGeoServiceProvider__OfflineGeocodingFeature   = QGeoServiceProvider__GeocodingFeature(C.QGeoServiceProvider_OfflineGeocodingFeature_Type())
	QGeoServiceProvider__ReverseGeocodingFeature   = QGeoServiceProvider__GeocodingFeature(C.QGeoServiceProvider_ReverseGeocodingFeature_Type())
	QGeoServiceProvider__LocalizedGeocodingFeature = QGeoServiceProvider__GeocodingFeature(C.QGeoServiceProvider_LocalizedGeocodingFeature_Type())
	QGeoServiceProvider__AnyGeocodingFeatures      = QGeoServiceProvider__GeocodingFeature(C.QGeoServiceProvider_AnyGeocodingFeatures_Type())
)

//QGeoServiceProvider::MappingFeature
type QGeoServiceProvider__MappingFeature int

var (
	QGeoServiceProvider__NoMappingFeatures       = QGeoServiceProvider__MappingFeature(0)
	QGeoServiceProvider__OnlineMappingFeature    = QGeoServiceProvider__MappingFeature(C.QGeoServiceProvider_OnlineMappingFeature_Type())
	QGeoServiceProvider__OfflineMappingFeature   = QGeoServiceProvider__MappingFeature(C.QGeoServiceProvider_OfflineMappingFeature_Type())
	QGeoServiceProvider__LocalizedMappingFeature = QGeoServiceProvider__MappingFeature(C.QGeoServiceProvider_LocalizedMappingFeature_Type())
	QGeoServiceProvider__AnyMappingFeatures      = QGeoServiceProvider__MappingFeature(C.QGeoServiceProvider_AnyMappingFeatures_Type())
)

//QGeoServiceProvider::PlacesFeature
type QGeoServiceProvider__PlacesFeature int

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
type QGeoServiceProvider__RoutingFeature int

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
	return strings.Split(C.GoString(C.QGeoServiceProvider_QGeoServiceProvider_AvailableServiceProviders()), "|")
}

func (ptr *QGeoServiceProvider) Error() QGeoServiceProvider__Error {
	if ptr.Pointer() != nil {
		return QGeoServiceProvider__Error(C.QGeoServiceProvider_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoServiceProvider) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoServiceProvider_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QGeoServiceProvider) GeocodingFeatures() QGeoServiceProvider__GeocodingFeature {
	if ptr.Pointer() != nil {
		return QGeoServiceProvider__GeocodingFeature(C.QGeoServiceProvider_GeocodingFeatures(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoServiceProvider) GeocodingManager() *QGeoCodingManager {
	if ptr.Pointer() != nil {
		return QGeoCodingManagerFromPointer(unsafe.Pointer(C.QGeoServiceProvider_GeocodingManager(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QGeoServiceProvider) MappingFeatures() QGeoServiceProvider__MappingFeature {
	if ptr.Pointer() != nil {
		return QGeoServiceProvider__MappingFeature(C.QGeoServiceProvider_MappingFeatures(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoServiceProvider) PlaceManager() *QPlaceManager {
	if ptr.Pointer() != nil {
		return QPlaceManagerFromPointer(unsafe.Pointer(C.QGeoServiceProvider_PlaceManager(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QGeoServiceProvider) PlacesFeatures() QGeoServiceProvider__PlacesFeature {
	if ptr.Pointer() != nil {
		return QGeoServiceProvider__PlacesFeature(C.QGeoServiceProvider_PlacesFeatures(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoServiceProvider) RoutingFeatures() QGeoServiceProvider__RoutingFeature {
	if ptr.Pointer() != nil {
		return QGeoServiceProvider__RoutingFeature(C.QGeoServiceProvider_RoutingFeatures(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoServiceProvider) RoutingManager() *QGeoRoutingManager {
	if ptr.Pointer() != nil {
		return QGeoRoutingManagerFromPointer(unsafe.Pointer(C.QGeoServiceProvider_RoutingManager(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QGeoServiceProvider) SetAllowExperimental(allow bool) {
	if ptr.Pointer() != nil {
		C.QGeoServiceProvider_SetAllowExperimental(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(allow)))
	}
}

func (ptr *QGeoServiceProvider) SetLocale(locale core.QLocaleITF) {
	if ptr.Pointer() != nil {
		C.QGeoServiceProvider_SetLocale(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQLocale(locale)))
	}
}

func (ptr *QGeoServiceProvider) DestroyQGeoServiceProvider() {
	if ptr.Pointer() != nil {
		C.QGeoServiceProvider_DestroyQGeoServiceProvider(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
