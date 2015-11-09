#ifdef __cplusplus
extern "C" {
#endif

void* QNmeaPositionInfoSource_NewQNmeaPositionInfoSource(int updateMode, void* parent);
void* QNmeaPositionInfoSource_Device(void* ptr);
int QNmeaPositionInfoSource_Error(void* ptr);
int QNmeaPositionInfoSource_MinimumUpdateInterval(void* ptr);
void QNmeaPositionInfoSource_RequestUpdate(void* ptr, int msec);
void QNmeaPositionInfoSource_SetDevice(void* ptr, void* device);
void QNmeaPositionInfoSource_SetUpdateInterval(void* ptr, int msec);
void QNmeaPositionInfoSource_StartUpdates(void* ptr);
void QNmeaPositionInfoSource_StopUpdates(void* ptr);
int QNmeaPositionInfoSource_SupportedPositioningMethods(void* ptr);
int QNmeaPositionInfoSource_UpdateMode(void* ptr);
void QNmeaPositionInfoSource_DestroyQNmeaPositionInfoSource(void* ptr);

#ifdef __cplusplus
}
#endif