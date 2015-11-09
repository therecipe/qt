#ifdef __cplusplus
extern "C" {
#endif

void* QSensorManager_QSensorManager_CreateBackend(void* sensor);
int QSensorManager_QSensorManager_IsBackendRegistered(void* ty, void* identifier);
void QSensorManager_QSensorManager_RegisterBackend(void* ty, void* identifier, void* factory);
void QSensorManager_QSensorManager_SetDefaultBackend(void* ty, void* identifier);
void QSensorManager_QSensorManager_UnregisterBackend(void* ty, void* identifier);

#ifdef __cplusplus
}
#endif