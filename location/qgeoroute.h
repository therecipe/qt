#ifdef __cplusplus
extern "C" {
#endif

void* QGeoRoute_NewQGeoRoute();
void* QGeoRoute_NewQGeoRoute2(void* other);
double QGeoRoute_Distance(void* ptr);
char* QGeoRoute_RouteId(void* ptr);
void QGeoRoute_SetBounds(void* ptr, void* bounds);
void QGeoRoute_SetDistance(void* ptr, double distance);
void QGeoRoute_SetFirstRouteSegment(void* ptr, void* routeSegment);
void QGeoRoute_SetRequest(void* ptr, void* request);
void QGeoRoute_SetRouteId(void* ptr, char* id);
void QGeoRoute_SetTravelMode(void* ptr, int mode);
void QGeoRoute_SetTravelTime(void* ptr, int secs);
int QGeoRoute_TravelMode(void* ptr);
int QGeoRoute_TravelTime(void* ptr);
void QGeoRoute_DestroyQGeoRoute(void* ptr);

#ifdef __cplusplus
}
#endif