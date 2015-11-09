#ifdef __cplusplus
extern "C" {
#endif

void* QTiltSensor_NewQTiltSensor(void* parent);
void* QTiltSensor_Reading(void* ptr);
void QTiltSensor_DestroyQTiltSensor(void* ptr);
void QTiltSensor_Calibrate(void* ptr);

#ifdef __cplusplus
}
#endif