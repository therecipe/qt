#ifdef __cplusplus
extern "C" {
#endif

char* QGeoAreaMonitorSource_QGeoAreaMonitorSource_AvailableSources();
void* QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateDefaultSource(void* parent);
void* QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateSource(char* sourceName, void* parent);
int QGeoAreaMonitorSource_Error(void* ptr);
void* QGeoAreaMonitorSource_PositionInfoSource(void* ptr);
int QGeoAreaMonitorSource_RequestUpdate(void* ptr, void* monitor, char* signal);
void QGeoAreaMonitorSource_SetPositionInfoSource(void* ptr, void* newSource);
char* QGeoAreaMonitorSource_SourceName(void* ptr);
int QGeoAreaMonitorSource_StartMonitoring(void* ptr, void* monitor);
int QGeoAreaMonitorSource_StopMonitoring(void* ptr, void* monitor);
int QGeoAreaMonitorSource_SupportedAreaMonitorFeatures(void* ptr);
void QGeoAreaMonitorSource_DestroyQGeoAreaMonitorSource(void* ptr);

#ifdef __cplusplus
}
#endif