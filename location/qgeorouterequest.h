#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QGeoRouteRequest_NewQGeoRouteRequest2(QtObjectPtr origin, QtObjectPtr destination);
QtObjectPtr QGeoRouteRequest_NewQGeoRouteRequest3(QtObjectPtr other);
int QGeoRouteRequest_FeatureWeight(QtObjectPtr ptr, int featureType);
int QGeoRouteRequest_ManeuverDetail(QtObjectPtr ptr);
int QGeoRouteRequest_NumberAlternativeRoutes(QtObjectPtr ptr);
int QGeoRouteRequest_RouteOptimization(QtObjectPtr ptr);
int QGeoRouteRequest_SegmentDetail(QtObjectPtr ptr);
void QGeoRouteRequest_SetFeatureWeight(QtObjectPtr ptr, int featureType, int featureWeight);
void QGeoRouteRequest_SetManeuverDetail(QtObjectPtr ptr, int maneuverDetail);
void QGeoRouteRequest_SetNumberAlternativeRoutes(QtObjectPtr ptr, int alternatives);
void QGeoRouteRequest_SetRouteOptimization(QtObjectPtr ptr, int optimization);
void QGeoRouteRequest_SetSegmentDetail(QtObjectPtr ptr, int segmentDetail);
void QGeoRouteRequest_SetTravelModes(QtObjectPtr ptr, int travelModes);
int QGeoRouteRequest_TravelModes(QtObjectPtr ptr);
void QGeoRouteRequest_DestroyQGeoRouteRequest(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif