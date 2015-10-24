#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QSensorBackend_IsFeatureSupported(QtObjectPtr ptr, int feature);
void QSensorBackend_SensorBusy(QtObjectPtr ptr);
void QSensorBackend_SensorError(QtObjectPtr ptr, int error);
void QSensorBackend_NewReadingAvailable(QtObjectPtr ptr);
QtObjectPtr QSensorBackend_Reading(QtObjectPtr ptr);
QtObjectPtr QSensorBackend_Sensor(QtObjectPtr ptr);
void QSensorBackend_SensorStopped(QtObjectPtr ptr);
void QSensorBackend_SetDataRates(QtObjectPtr ptr, QtObjectPtr otherSensor);
void QSensorBackend_SetDescription(QtObjectPtr ptr, char* description);
void QSensorBackend_Start(QtObjectPtr ptr);
void QSensorBackend_Stop(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif