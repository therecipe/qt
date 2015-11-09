#ifdef __cplusplus
extern "C" {
#endif

int QAccelerometer_AccelerationMode(void* ptr);
void* QAccelerometer_Reading(void* ptr);
void* QAccelerometer_NewQAccelerometer(void* parent);
void QAccelerometer_ConnectAccelerationModeChanged(void* ptr);
void QAccelerometer_DisconnectAccelerationModeChanged(void* ptr);
void QAccelerometer_SetAccelerationMode(void* ptr, int accelerationMode);
void QAccelerometer_DestroyQAccelerometer(void* ptr);

#ifdef __cplusplus
}
#endif