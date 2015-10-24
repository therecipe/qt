#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSensorGestureManager_NewQSensorGestureManager(QtObjectPtr parent);
char* QSensorGestureManager_GestureIds(QtObjectPtr ptr);
void QSensorGestureManager_ConnectNewSensorGestureAvailable(QtObjectPtr ptr);
void QSensorGestureManager_DisconnectNewSensorGestureAvailable(QtObjectPtr ptr);
char* QSensorGestureManager_RecognizerSignals(QtObjectPtr ptr, char* gestureId);
int QSensorGestureManager_RegisterSensorGestureRecognizer(QtObjectPtr ptr, QtObjectPtr recognizer);
QtObjectPtr QSensorGestureManager_QSensorGestureManager_SensorGestureRecognizer(char* id);
void QSensorGestureManager_DestroyQSensorGestureManager(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif