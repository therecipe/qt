#ifdef __cplusplus
extern "C" {
#endif

void* QGeoRouteSegment_NewQGeoRouteSegment();
void* QGeoRouteSegment_NewQGeoRouteSegment2(void* other);
double QGeoRouteSegment_Distance(void* ptr);
int QGeoRouteSegment_IsValid(void* ptr);
void QGeoRouteSegment_SetDistance(void* ptr, double distance);
void QGeoRouteSegment_SetManeuver(void* ptr, void* maneuver);
void QGeoRouteSegment_SetNextRouteSegment(void* ptr, void* routeSegment);
void QGeoRouteSegment_SetTravelTime(void* ptr, int secs);
int QGeoRouteSegment_TravelTime(void* ptr);
void QGeoRouteSegment_DestroyQGeoRouteSegment(void* ptr);

#ifdef __cplusplus
}
#endif