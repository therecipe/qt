#ifdef __cplusplus
extern "C" {
#endif

void QSensorBackend_AddDataRate(void* ptr, double min, double max);
int QSensorBackend_IsFeatureSupported(void* ptr, int feature);
void QSensorBackend_SensorBusy(void* ptr);
void QSensorBackend_SensorError(void* ptr, int error);
void QSensorBackend_AddOutputRange(void* ptr, double min, double max, double accuracy);
void QSensorBackend_NewReadingAvailable(void* ptr);
void* QSensorBackend_Reading(void* ptr);
void* QSensorBackend_Sensor(void* ptr);
void QSensorBackend_SensorStopped(void* ptr);
void QSensorBackend_SetDataRates(void* ptr, void* otherSensor);
void QSensorBackend_SetDescription(void* ptr, char* description);
void QSensorBackend_Start(void* ptr);
void QSensorBackend_Stop(void* ptr);

#ifdef __cplusplus
}
#endif