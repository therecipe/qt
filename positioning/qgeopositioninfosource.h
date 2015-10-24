#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QGeoPositionInfoSource_SetUpdateInterval(QtObjectPtr ptr, int msec);
char* QGeoPositionInfoSource_SourceName(QtObjectPtr ptr);
int QGeoPositionInfoSource_UpdateInterval(QtObjectPtr ptr);
char* QGeoPositionInfoSource_QGeoPositionInfoSource_AvailableSources();
QtObjectPtr QGeoPositionInfoSource_QGeoPositionInfoSource_CreateDefaultSource(QtObjectPtr parent);
QtObjectPtr QGeoPositionInfoSource_QGeoPositionInfoSource_CreateSource(char* sourceName, QtObjectPtr parent);
int QGeoPositionInfoSource_Error(QtObjectPtr ptr);
int QGeoPositionInfoSource_MinimumUpdateInterval(QtObjectPtr ptr);
int QGeoPositionInfoSource_PreferredPositioningMethods(QtObjectPtr ptr);
void QGeoPositionInfoSource_RequestUpdate(QtObjectPtr ptr, int timeout);
void QGeoPositionInfoSource_SetPreferredPositioningMethods(QtObjectPtr ptr, int methods);
void QGeoPositionInfoSource_StartUpdates(QtObjectPtr ptr);
void QGeoPositionInfoSource_StopUpdates(QtObjectPtr ptr);
int QGeoPositionInfoSource_SupportedPositioningMethods(QtObjectPtr ptr);
void QGeoPositionInfoSource_ConnectUpdateTimeout(QtObjectPtr ptr);
void QGeoPositionInfoSource_DisconnectUpdateTimeout(QtObjectPtr ptr);
void QGeoPositionInfoSource_DestroyQGeoPositionInfoSource(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif