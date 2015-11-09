#ifdef __cplusplus
extern "C" {
#endif

void* QSensorGesture_NewQSensorGesture(char* ids, void* parent);
char* QSensorGesture_GestureSignals(void* ptr);
char* QSensorGesture_InvalidIds(void* ptr);
int QSensorGesture_IsActive(void* ptr);
void QSensorGesture_StartDetection(void* ptr);
void QSensorGesture_StopDetection(void* ptr);
char* QSensorGesture_ValidIds(void* ptr);
void QSensorGesture_DestroyQSensorGesture(void* ptr);

#ifdef __cplusplus
}
#endif