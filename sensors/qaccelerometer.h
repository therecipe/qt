#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QAccelerometer_AccelerationMode(QtObjectPtr ptr);
QtObjectPtr QAccelerometer_Reading(QtObjectPtr ptr);
QtObjectPtr QAccelerometer_NewQAccelerometer(QtObjectPtr parent);
void QAccelerometer_ConnectAccelerationModeChanged(QtObjectPtr ptr);
void QAccelerometer_DisconnectAccelerationModeChanged(QtObjectPtr ptr);
void QAccelerometer_SetAccelerationMode(QtObjectPtr ptr, int accelerationMode);
void QAccelerometer_DestroyQAccelerometer(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif