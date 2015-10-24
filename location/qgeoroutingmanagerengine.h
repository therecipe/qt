#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QGeoRoutingManagerEngine_CalculateRoute(QtObjectPtr ptr, QtObjectPtr request);
void QGeoRoutingManagerEngine_ConnectError(QtObjectPtr ptr);
void QGeoRoutingManagerEngine_DisconnectError(QtObjectPtr ptr);
void QGeoRoutingManagerEngine_ConnectFinished(QtObjectPtr ptr);
void QGeoRoutingManagerEngine_DisconnectFinished(QtObjectPtr ptr);
char* QGeoRoutingManagerEngine_ManagerName(QtObjectPtr ptr);
int QGeoRoutingManagerEngine_ManagerVersion(QtObjectPtr ptr);
int QGeoRoutingManagerEngine_MeasurementSystem(QtObjectPtr ptr);
void QGeoRoutingManagerEngine_SetLocale(QtObjectPtr ptr, QtObjectPtr locale);
void QGeoRoutingManagerEngine_SetMeasurementSystem(QtObjectPtr ptr, int system);
int QGeoRoutingManagerEngine_SupportedFeatureTypes(QtObjectPtr ptr);
int QGeoRoutingManagerEngine_SupportedFeatureWeights(QtObjectPtr ptr);
int QGeoRoutingManagerEngine_SupportedManeuverDetails(QtObjectPtr ptr);
int QGeoRoutingManagerEngine_SupportedRouteOptimizations(QtObjectPtr ptr);
int QGeoRoutingManagerEngine_SupportedSegmentDetails(QtObjectPtr ptr);
int QGeoRoutingManagerEngine_SupportedTravelModes(QtObjectPtr ptr);
QtObjectPtr QGeoRoutingManagerEngine_UpdateRoute(QtObjectPtr ptr, QtObjectPtr route, QtObjectPtr position);
void QGeoRoutingManagerEngine_DestroyQGeoRoutingManagerEngine(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif