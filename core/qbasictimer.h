#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QBasicTimer_Start(QtObjectPtr ptr, int msec, QtObjectPtr object);
QtObjectPtr QBasicTimer_NewQBasicTimer();
int QBasicTimer_IsActive(QtObjectPtr ptr);
void QBasicTimer_Start2(QtObjectPtr ptr, int msec, int timerType, QtObjectPtr obj);
void QBasicTimer_Stop(QtObjectPtr ptr);
int QBasicTimer_TimerId(QtObjectPtr ptr);
void QBasicTimer_DestroyQBasicTimer(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif