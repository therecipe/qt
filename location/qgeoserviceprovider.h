#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QGeoServiceProvider_OnlineGeocodingFeature_Type();
int QGeoServiceProvider_OfflineGeocodingFeature_Type();
int QGeoServiceProvider_ReverseGeocodingFeature_Type();
int QGeoServiceProvider_LocalizedGeocodingFeature_Type();
int QGeoServiceProvider_AnyGeocodingFeatures_Type();
int QGeoServiceProvider_OnlineMappingFeature_Type();
int QGeoServiceProvider_OfflineMappingFeature_Type();
int QGeoServiceProvider_LocalizedMappingFeature_Type();
int QGeoServiceProvider_AnyMappingFeatures_Type();
int QGeoServiceProvider_OnlinePlacesFeature_Type();
int QGeoServiceProvider_OfflinePlacesFeature_Type();
int QGeoServiceProvider_SavePlaceFeature_Type();
int QGeoServiceProvider_RemovePlaceFeature_Type();
int QGeoServiceProvider_SaveCategoryFeature_Type();
int QGeoServiceProvider_RemoveCategoryFeature_Type();
int QGeoServiceProvider_PlaceRecommendationsFeature_Type();
int QGeoServiceProvider_SearchSuggestionsFeature_Type();
int QGeoServiceProvider_LocalizedPlacesFeature_Type();
int QGeoServiceProvider_NotificationsFeature_Type();
int QGeoServiceProvider_PlaceMatchingFeature_Type();
int QGeoServiceProvider_AnyPlacesFeatures_Type();
int QGeoServiceProvider_OnlineRoutingFeature_Type();
int QGeoServiceProvider_OfflineRoutingFeature_Type();
int QGeoServiceProvider_LocalizedRoutingFeature_Type();
int QGeoServiceProvider_RouteUpdatesFeature_Type();
int QGeoServiceProvider_AlternativeRoutesFeature_Type();
int QGeoServiceProvider_ExcludeAreasRoutingFeature_Type();
int QGeoServiceProvider_AnyRoutingFeatures_Type();
char* QGeoServiceProvider_QGeoServiceProvider_AvailableServiceProviders();
int QGeoServiceProvider_Error(QtObjectPtr ptr);
char* QGeoServiceProvider_ErrorString(QtObjectPtr ptr);
int QGeoServiceProvider_GeocodingFeatures(QtObjectPtr ptr);
QtObjectPtr QGeoServiceProvider_GeocodingManager(QtObjectPtr ptr);
int QGeoServiceProvider_MappingFeatures(QtObjectPtr ptr);
QtObjectPtr QGeoServiceProvider_PlaceManager(QtObjectPtr ptr);
int QGeoServiceProvider_PlacesFeatures(QtObjectPtr ptr);
int QGeoServiceProvider_RoutingFeatures(QtObjectPtr ptr);
QtObjectPtr QGeoServiceProvider_RoutingManager(QtObjectPtr ptr);
void QGeoServiceProvider_SetAllowExperimental(QtObjectPtr ptr, int allow);
void QGeoServiceProvider_SetLocale(QtObjectPtr ptr, QtObjectPtr locale);
void QGeoServiceProvider_DestroyQGeoServiceProvider(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif