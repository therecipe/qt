#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QGeoAreaMonitorInfo_NewQGeoAreaMonitorInfo2(QtObjectPtr other);
QtObjectPtr QGeoAreaMonitorInfo_NewQGeoAreaMonitorInfo(char* name);
char* QGeoAreaMonitorInfo_Identifier(QtObjectPtr ptr);
int QGeoAreaMonitorInfo_IsPersistent(QtObjectPtr ptr);
int QGeoAreaMonitorInfo_IsValid(QtObjectPtr ptr);
char* QGeoAreaMonitorInfo_Name(QtObjectPtr ptr);
void QGeoAreaMonitorInfo_SetArea(QtObjectPtr ptr, QtObjectPtr newShape);
void QGeoAreaMonitorInfo_SetExpiration(QtObjectPtr ptr, QtObjectPtr expiry);
void QGeoAreaMonitorInfo_SetName(QtObjectPtr ptr, char* name);
void QGeoAreaMonitorInfo_SetPersistent(QtObjectPtr ptr, int isPersistent);
void QGeoAreaMonitorInfo_DestroyQGeoAreaMonitorInfo(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif