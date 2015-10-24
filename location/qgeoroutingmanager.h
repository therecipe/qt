#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QGeoRoutingManager_CalculateRoute(QtObjectPtr ptr, QtObjectPtr request);
void QGeoRoutingManager_ConnectError(QtObjectPtr ptr);
void QGeoRoutingManager_DisconnectError(QtObjectPtr ptr);
void QGeoRoutingManager_ConnectFinished(QtObjectPtr ptr);
void QGeoRoutingManager_DisconnectFinished(QtObjectPtr ptr);
char* QGeoRoutingManager_ManagerName(QtObjectPtr ptr);
int QGeoRoutingManager_ManagerVersion(QtObjectPtr ptr);
int QGeoRoutingManager_MeasurementSystem(QtObjectPtr ptr);
void QGeoRoutingManager_SetLocale(QtObjectPtr ptr, QtObjectPtr locale);
void QGeoRoutingManager_SetMeasurementSystem(QtObjectPtr ptr, int system);
int QGeoRoutingManager_SupportedFeatureTypes(QtObjectPtr ptr);
int QGeoRoutingManager_SupportedFeatureWeights(QtObjectPtr ptr);
int QGeoRoutingManager_SupportedManeuverDetails(QtObjectPtr ptr);
int QGeoRoutingManager_SupportedRouteOptimizations(QtObjectPtr ptr);
int QGeoRoutingManager_SupportedSegmentDetails(QtObjectPtr ptr);
int QGeoRoutingManager_SupportedTravelModes(QtObjectPtr ptr);
QtObjectPtr QGeoRoutingManager_UpdateRoute(QtObjectPtr ptr, QtObjectPtr route, QtObjectPtr position);
void QGeoRoutingManager_DestroyQGeoRoutingManager(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif