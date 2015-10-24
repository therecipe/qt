#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QNmeaPositionInfoSource_NewQNmeaPositionInfoSource(int updateMode, QtObjectPtr parent);
QtObjectPtr QNmeaPositionInfoSource_Device(QtObjectPtr ptr);
int QNmeaPositionInfoSource_Error(QtObjectPtr ptr);
int QNmeaPositionInfoSource_MinimumUpdateInterval(QtObjectPtr ptr);
void QNmeaPositionInfoSource_RequestUpdate(QtObjectPtr ptr, int msec);
void QNmeaPositionInfoSource_SetDevice(QtObjectPtr ptr, QtObjectPtr device);
void QNmeaPositionInfoSource_SetUpdateInterval(QtObjectPtr ptr, int msec);
void QNmeaPositionInfoSource_StartUpdates(QtObjectPtr ptr);
void QNmeaPositionInfoSource_StopUpdates(QtObjectPtr ptr);
int QNmeaPositionInfoSource_SupportedPositioningMethods(QtObjectPtr ptr);
int QNmeaPositionInfoSource_UpdateMode(QtObjectPtr ptr);
void QNmeaPositionInfoSource_DestroyQNmeaPositionInfoSource(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif