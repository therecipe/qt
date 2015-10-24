#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QRotationSensor_HasZ(QtObjectPtr ptr);
QtObjectPtr QRotationSensor_Reading(QtObjectPtr ptr);
QtObjectPtr QRotationSensor_NewQRotationSensor(QtObjectPtr parent);
void QRotationSensor_ConnectHasZChanged(QtObjectPtr ptr);
void QRotationSensor_DisconnectHasZChanged(QtObjectPtr ptr);
void QRotationSensor_SetHasZ(QtObjectPtr ptr, int hasZ);
void QRotationSensor_DestroyQRotationSensor(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif