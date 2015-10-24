#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QGeoPositionInfo_NewQGeoPositionInfo();
QtObjectPtr QGeoPositionInfo_NewQGeoPositionInfo2(QtObjectPtr coordinate, QtObjectPtr timestamp);
QtObjectPtr QGeoPositionInfo_NewQGeoPositionInfo3(QtObjectPtr other);
int QGeoPositionInfo_HasAttribute(QtObjectPtr ptr, int attribute);
int QGeoPositionInfo_IsValid(QtObjectPtr ptr);
void QGeoPositionInfo_RemoveAttribute(QtObjectPtr ptr, int attribute);
void QGeoPositionInfo_SetCoordinate(QtObjectPtr ptr, QtObjectPtr coordinate);
void QGeoPositionInfo_SetTimestamp(QtObjectPtr ptr, QtObjectPtr timestamp);
void QGeoPositionInfo_DestroyQGeoPositionInfo(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif