#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QGeoSatelliteInfoSource_SetUpdateInterval(QtObjectPtr ptr, int msec);
int QGeoSatelliteInfoSource_UpdateInterval(QtObjectPtr ptr);
char* QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_AvailableSources();
QtObjectPtr QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateDefaultSource(QtObjectPtr parent);
QtObjectPtr QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateSource(char* sourceName, QtObjectPtr parent);
int QGeoSatelliteInfoSource_Error(QtObjectPtr ptr);
int QGeoSatelliteInfoSource_MinimumUpdateInterval(QtObjectPtr ptr);
void QGeoSatelliteInfoSource_ConnectRequestTimeout(QtObjectPtr ptr);
void QGeoSatelliteInfoSource_DisconnectRequestTimeout(QtObjectPtr ptr);
void QGeoSatelliteInfoSource_RequestUpdate(QtObjectPtr ptr, int timeout);
char* QGeoSatelliteInfoSource_SourceName(QtObjectPtr ptr);
void QGeoSatelliteInfoSource_StartUpdates(QtObjectPtr ptr);
void QGeoSatelliteInfoSource_StopUpdates(QtObjectPtr ptr);
void QGeoSatelliteInfoSource_DestroyQGeoSatelliteInfoSource(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif