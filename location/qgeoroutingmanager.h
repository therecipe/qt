#ifdef __cplusplus
extern "C" {
#endif

void* QGeoRoutingManager_CalculateRoute(void* ptr, void* request);
void QGeoRoutingManager_ConnectError(void* ptr);
void QGeoRoutingManager_DisconnectError(void* ptr);
void QGeoRoutingManager_ConnectFinished(void* ptr);
void QGeoRoutingManager_DisconnectFinished(void* ptr);
char* QGeoRoutingManager_ManagerName(void* ptr);
int QGeoRoutingManager_ManagerVersion(void* ptr);
int QGeoRoutingManager_MeasurementSystem(void* ptr);
void QGeoRoutingManager_SetLocale(void* ptr, void* locale);
void QGeoRoutingManager_SetMeasurementSystem(void* ptr, int system);
int QGeoRoutingManager_SupportedFeatureTypes(void* ptr);
int QGeoRoutingManager_SupportedFeatureWeights(void* ptr);
int QGeoRoutingManager_SupportedManeuverDetails(void* ptr);
int QGeoRoutingManager_SupportedRouteOptimizations(void* ptr);
int QGeoRoutingManager_SupportedSegmentDetails(void* ptr);
int QGeoRoutingManager_SupportedTravelModes(void* ptr);
void* QGeoRoutingManager_UpdateRoute(void* ptr, void* route, void* position);
void QGeoRoutingManager_DestroyQGeoRoutingManager(void* ptr);

#ifdef __cplusplus
}
#endif