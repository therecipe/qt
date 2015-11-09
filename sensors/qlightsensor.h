#ifdef __cplusplus
extern "C" {
#endif

double QLightSensor_FieldOfView(void* ptr);
void* QLightSensor_Reading(void* ptr);
void* QLightSensor_NewQLightSensor(void* parent);
void QLightSensor_SetFieldOfView(void* ptr, double fieldOfView);
void QLightSensor_DestroyQLightSensor(void* ptr);

#ifdef __cplusplus
}
#endif