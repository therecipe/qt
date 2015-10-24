#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QGeoAreaMonitorSource_QGeoAreaMonitorSource_AvailableSources();
QtObjectPtr QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateDefaultSource(QtObjectPtr parent);
QtObjectPtr QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateSource(char* sourceName, QtObjectPtr parent);
int QGeoAreaMonitorSource_Error(QtObjectPtr ptr);
QtObjectPtr QGeoAreaMonitorSource_PositionInfoSource(QtObjectPtr ptr);
int QGeoAreaMonitorSource_RequestUpdate(QtObjectPtr ptr, QtObjectPtr monitor, char* signal);
void QGeoAreaMonitorSource_SetPositionInfoSource(QtObjectPtr ptr, QtObjectPtr newSource);
char* QGeoAreaMonitorSource_SourceName(QtObjectPtr ptr);
int QGeoAreaMonitorSource_StartMonitoring(QtObjectPtr ptr, QtObjectPtr monitor);
int QGeoAreaMonitorSource_StopMonitoring(QtObjectPtr ptr, QtObjectPtr monitor);
int QGeoAreaMonitorSource_SupportedAreaMonitorFeatures(QtObjectPtr ptr);
void QGeoAreaMonitorSource_DestroyQGeoAreaMonitorSource(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif