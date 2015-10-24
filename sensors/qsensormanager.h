#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSensorManager_QSensorManager_CreateBackend(QtObjectPtr sensor);
int QSensorManager_QSensorManager_IsBackendRegistered(QtObjectPtr ty, QtObjectPtr identifier);
void QSensorManager_QSensorManager_RegisterBackend(QtObjectPtr ty, QtObjectPtr identifier, QtObjectPtr factory);
void QSensorManager_QSensorManager_SetDefaultBackend(QtObjectPtr ty, QtObjectPtr identifier);
void QSensorManager_QSensorManager_UnregisterBackend(QtObjectPtr ty, QtObjectPtr identifier);

#ifdef __cplusplus
}
#endif