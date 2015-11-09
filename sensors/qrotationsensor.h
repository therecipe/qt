#ifdef __cplusplus
extern "C" {
#endif

int QRotationSensor_HasZ(void* ptr);
void* QRotationSensor_Reading(void* ptr);
void* QRotationSensor_NewQRotationSensor(void* parent);
void QRotationSensor_ConnectHasZChanged(void* ptr);
void QRotationSensor_DisconnectHasZChanged(void* ptr);
void QRotationSensor_SetHasZ(void* ptr, int hasZ);
void QRotationSensor_DestroyQRotationSensor(void* ptr);

#ifdef __cplusplus
}
#endif