#ifdef __cplusplus
extern "C" {
#endif

void* QGeoAreaMonitorInfo_NewQGeoAreaMonitorInfo2(void* other);
void* QGeoAreaMonitorInfo_NewQGeoAreaMonitorInfo(char* name);
void* QGeoAreaMonitorInfo_Expiration(void* ptr);
char* QGeoAreaMonitorInfo_Identifier(void* ptr);
int QGeoAreaMonitorInfo_IsPersistent(void* ptr);
int QGeoAreaMonitorInfo_IsValid(void* ptr);
char* QGeoAreaMonitorInfo_Name(void* ptr);
void QGeoAreaMonitorInfo_SetArea(void* ptr, void* newShape);
void QGeoAreaMonitorInfo_SetExpiration(void* ptr, void* expiry);
void QGeoAreaMonitorInfo_SetName(void* ptr, char* name);
void QGeoAreaMonitorInfo_SetPersistent(void* ptr, int isPersistent);
void QGeoAreaMonitorInfo_DestroyQGeoAreaMonitorInfo(void* ptr);

#ifdef __cplusplus
}
#endif