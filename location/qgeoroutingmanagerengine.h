#ifdef __cplusplus
extern "C" {
#endif

void* QGeoRoutingManagerEngine_CalculateRoute(void* ptr, void* request);
void QGeoRoutingManagerEngine_ConnectError(void* ptr);
void QGeoRoutingManagerEngine_DisconnectError(void* ptr);
void QGeoRoutingManagerEngine_ConnectFinished(void* ptr);
void QGeoRoutingManagerEngine_DisconnectFinished(void* ptr);
char* QGeoRoutingManagerEngine_ManagerName(void* ptr);
int QGeoRoutingManagerEngine_ManagerVersion(void* ptr);
int QGeoRoutingManagerEngine_MeasurementSystem(void* ptr);
void QGeoRoutingManagerEngine_SetLocale(void* ptr, void* locale);
void QGeoRoutingManagerEngine_SetMeasurementSystem(void* ptr, int system);
int QGeoRoutingManagerEngine_SupportedFeatureTypes(void* ptr);
int QGeoRoutingManagerEngine_SupportedFeatureWeights(void* ptr);
int QGeoRoutingManagerEngine_SupportedManeuverDetails(void* ptr);
int QGeoRoutingManagerEngine_SupportedRouteOptimizations(void* ptr);
int QGeoRoutingManagerEngine_SupportedSegmentDetails(void* ptr);
int QGeoRoutingManagerEngine_SupportedTravelModes(void* ptr);
void* QGeoRoutingManagerEngine_UpdateRoute(void* ptr, void* route, void* position);
void QGeoRoutingManagerEngine_DestroyQGeoRoutingManagerEngine(void* ptr);

#ifdef __cplusplus
}
#endif