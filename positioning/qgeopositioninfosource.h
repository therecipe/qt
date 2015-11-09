#ifdef __cplusplus
extern "C" {
#endif

void QGeoPositionInfoSource_SetUpdateInterval(void* ptr, int msec);
char* QGeoPositionInfoSource_SourceName(void* ptr);
int QGeoPositionInfoSource_UpdateInterval(void* ptr);
char* QGeoPositionInfoSource_QGeoPositionInfoSource_AvailableSources();
void* QGeoPositionInfoSource_QGeoPositionInfoSource_CreateDefaultSource(void* parent);
void* QGeoPositionInfoSource_QGeoPositionInfoSource_CreateSource(char* sourceName, void* parent);
int QGeoPositionInfoSource_Error(void* ptr);
int QGeoPositionInfoSource_MinimumUpdateInterval(void* ptr);
int QGeoPositionInfoSource_PreferredPositioningMethods(void* ptr);
void QGeoPositionInfoSource_RequestUpdate(void* ptr, int timeout);
void QGeoPositionInfoSource_SetPreferredPositioningMethods(void* ptr, int methods);
void QGeoPositionInfoSource_StartUpdates(void* ptr);
void QGeoPositionInfoSource_StopUpdates(void* ptr);
int QGeoPositionInfoSource_SupportedPositioningMethods(void* ptr);
void QGeoPositionInfoSource_ConnectUpdateTimeout(void* ptr);
void QGeoPositionInfoSource_DisconnectUpdateTimeout(void* ptr);
void QGeoPositionInfoSource_DestroyQGeoPositionInfoSource(void* ptr);

#ifdef __cplusplus
}
#endif