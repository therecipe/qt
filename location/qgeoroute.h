#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QGeoRoute_NewQGeoRoute();
QtObjectPtr QGeoRoute_NewQGeoRoute2(QtObjectPtr other);
char* QGeoRoute_RouteId(QtObjectPtr ptr);
void QGeoRoute_SetBounds(QtObjectPtr ptr, QtObjectPtr bounds);
void QGeoRoute_SetFirstRouteSegment(QtObjectPtr ptr, QtObjectPtr routeSegment);
void QGeoRoute_SetRequest(QtObjectPtr ptr, QtObjectPtr request);
void QGeoRoute_SetRouteId(QtObjectPtr ptr, char* id);
void QGeoRoute_SetTravelMode(QtObjectPtr ptr, int mode);
void QGeoRoute_SetTravelTime(QtObjectPtr ptr, int secs);
int QGeoRoute_TravelMode(QtObjectPtr ptr);
int QGeoRoute_TravelTime(QtObjectPtr ptr);
void QGeoRoute_DestroyQGeoRoute(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif