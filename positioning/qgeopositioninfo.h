#ifdef __cplusplus
extern "C" {
#endif

void* QGeoPositionInfo_NewQGeoPositionInfo();
void* QGeoPositionInfo_NewQGeoPositionInfo2(void* coordinate, void* timestamp);
void* QGeoPositionInfo_NewQGeoPositionInfo3(void* other);
double QGeoPositionInfo_Attribute(void* ptr, int attribute);
int QGeoPositionInfo_HasAttribute(void* ptr, int attribute);
int QGeoPositionInfo_IsValid(void* ptr);
void QGeoPositionInfo_RemoveAttribute(void* ptr, int attribute);
void QGeoPositionInfo_SetAttribute(void* ptr, int attribute, double value);
void QGeoPositionInfo_SetCoordinate(void* ptr, void* coordinate);
void QGeoPositionInfo_SetTimestamp(void* ptr, void* timestamp);
void* QGeoPositionInfo_Timestamp(void* ptr);
void QGeoPositionInfo_DestroyQGeoPositionInfo(void* ptr);

#ifdef __cplusplus
}
#endif