#ifdef __cplusplus
extern "C" {
#endif

void* QSensorGestureManager_NewQSensorGestureManager(void* parent);
char* QSensorGestureManager_GestureIds(void* ptr);
void QSensorGestureManager_ConnectNewSensorGestureAvailable(void* ptr);
void QSensorGestureManager_DisconnectNewSensorGestureAvailable(void* ptr);
char* QSensorGestureManager_RecognizerSignals(void* ptr, char* gestureId);
int QSensorGestureManager_RegisterSensorGestureRecognizer(void* ptr, void* recognizer);
void* QSensorGestureManager_QSensorGestureManager_SensorGestureRecognizer(char* id);
void QSensorGestureManager_DestroyQSensorGestureManager(void* ptr);

#ifdef __cplusplus
}
#endif