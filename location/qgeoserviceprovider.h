#ifdef __cplusplus
extern "C" {
#endif

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
int QGeoServiceProvider_Error(void* ptr);
char* QGeoServiceProvider_ErrorString(void* ptr);
int QGeoServiceProvider_GeocodingFeatures(void* ptr);
void* QGeoServiceProvider_GeocodingManager(void* ptr);
int QGeoServiceProvider_MappingFeatures(void* ptr);
void* QGeoServiceProvider_PlaceManager(void* ptr);
int QGeoServiceProvider_PlacesFeatures(void* ptr);
int QGeoServiceProvider_RoutingFeatures(void* ptr);
void* QGeoServiceProvider_RoutingManager(void* ptr);
void QGeoServiceProvider_SetAllowExperimental(void* ptr, int allow);
void QGeoServiceProvider_SetLocale(void* ptr, void* locale);
void QGeoServiceProvider_DestroyQGeoServiceProvider(void* ptr);

#ifdef __cplusplus
}
#endif