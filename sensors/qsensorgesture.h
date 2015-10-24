#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSensorGesture_NewQSensorGesture(char* ids, QtObjectPtr parent);
char* QSensorGesture_GestureSignals(QtObjectPtr ptr);
char* QSensorGesture_InvalidIds(QtObjectPtr ptr);
int QSensorGesture_IsActive(QtObjectPtr ptr);
void QSensorGesture_StartDetection(QtObjectPtr ptr);
void QSensorGesture_StopDetection(QtObjectPtr ptr);
char* QSensorGesture_ValidIds(QtObjectPtr ptr);
void QSensorGesture_DestroyQSensorGesture(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif