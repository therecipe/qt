#ifdef __cplusplus
extern "C" {
#endif

void* QGeoRouteRequest_NewQGeoRouteRequest2(void* origin, void* destination);
void* QGeoRouteRequest_NewQGeoRouteRequest3(void* other);
int QGeoRouteRequest_FeatureWeight(void* ptr, int featureType);
int QGeoRouteRequest_ManeuverDetail(void* ptr);
int QGeoRouteRequest_NumberAlternativeRoutes(void* ptr);
int QGeoRouteRequest_RouteOptimization(void* ptr);
int QGeoRouteRequest_SegmentDetail(void* ptr);
void QGeoRouteRequest_SetFeatureWeight(void* ptr, int featureType, int featureWeight);
void QGeoRouteRequest_SetManeuverDetail(void* ptr, int maneuverDetail);
void QGeoRouteRequest_SetNumberAlternativeRoutes(void* ptr, int alternatives);
void QGeoRouteRequest_SetRouteOptimization(void* ptr, int optimization);
void QGeoRouteRequest_SetSegmentDetail(void* ptr, int segmentDetail);
void QGeoRouteRequest_SetTravelModes(void* ptr, int travelModes);
int QGeoRouteRequest_TravelModes(void* ptr);
void QGeoRouteRequest_DestroyQGeoRouteRequest(void* ptr);

#ifdef __cplusplus
}
#endif