#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QSensorGestureRecognizer_CreateBackend(QtObjectPtr ptr);
char* QSensorGestureRecognizer_GestureSignals(QtObjectPtr ptr);
char* QSensorGestureRecognizer_Id(QtObjectPtr ptr);
int QSensorGestureRecognizer_IsActive(QtObjectPtr ptr);
void QSensorGestureRecognizer_StartBackend(QtObjectPtr ptr);
void QSensorGestureRecognizer_StopBackend(QtObjectPtr ptr);
void QSensorGestureRecognizer_DestroyQSensorGestureRecognizer(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif