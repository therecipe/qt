#ifdef __cplusplus
extern "C" {
#endif

void QGeoSatelliteInfoSource_SetUpdateInterval(void* ptr, int msec);
int QGeoSatelliteInfoSource_UpdateInterval(void* ptr);
char* QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_AvailableSources();
void* QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateDefaultSource(void* parent);
void* QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateSource(char* sourceName, void* parent);
int QGeoSatelliteInfoSource_Error(void* ptr);
int QGeoSatelliteInfoSource_MinimumUpdateInterval(void* ptr);
void QGeoSatelliteInfoSource_ConnectRequestTimeout(void* ptr);
void QGeoSatelliteInfoSource_DisconnectRequestTimeout(void* ptr);
void QGeoSatelliteInfoSource_RequestUpdate(void* ptr, int timeout);
char* QGeoSatelliteInfoSource_SourceName(void* ptr);
void QGeoSatelliteInfoSource_StartUpdates(void* ptr);
void QGeoSatelliteInfoSource_StopUpdates(void* ptr);
void QGeoSatelliteInfoSource_DestroyQGeoSatelliteInfoSource(void* ptr);

#ifdef __cplusplus
}
#endif