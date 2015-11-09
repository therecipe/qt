#ifdef __cplusplus
extern "C" {
#endif

void QSensorGestureRecognizer_CreateBackend(void* ptr);
char* QSensorGestureRecognizer_GestureSignals(void* ptr);
char* QSensorGestureRecognizer_Id(void* ptr);
int QSensorGestureRecognizer_IsActive(void* ptr);
void QSensorGestureRecognizer_StartBackend(void* ptr);
void QSensorGestureRecognizer_StopBackend(void* ptr);
void QSensorGestureRecognizer_DestroyQSensorGestureRecognizer(void* ptr);

#ifdef __cplusplus
}
#endif