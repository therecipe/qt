#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QGeoRouteSegment_NewQGeoRouteSegment();
QtObjectPtr QGeoRouteSegment_NewQGeoRouteSegment2(QtObjectPtr other);
int QGeoRouteSegment_IsValid(QtObjectPtr ptr);
void QGeoRouteSegment_SetManeuver(QtObjectPtr ptr, QtObjectPtr maneuver);
void QGeoRouteSegment_SetNextRouteSegment(QtObjectPtr ptr, QtObjectPtr routeSegment);
void QGeoRouteSegment_SetTravelTime(QtObjectPtr ptr, int secs);
int QGeoRouteSegment_TravelTime(QtObjectPtr ptr);
void QGeoRouteSegment_DestroyQGeoRouteSegment(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif