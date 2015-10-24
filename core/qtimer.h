#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QTimer_RemainingTime(QtObjectPtr ptr);
void QTimer_SetInterval(QtObjectPtr ptr, int msec);
QtObjectPtr QTimer_NewQTimer(QtObjectPtr parent);
int QTimer_Interval(QtObjectPtr ptr);
int QTimer_IsActive(QtObjectPtr ptr);
int QTimer_IsSingleShot(QtObjectPtr ptr);
void QTimer_SetSingleShot(QtObjectPtr ptr, int singleShot);
void QTimer_SetTimerType(QtObjectPtr ptr, int atype);
void QTimer_QTimer_SingleShot2(int msec, int timerType, QtObjectPtr receiver, char* member);
void QTimer_QTimer_SingleShot(int msec, QtObjectPtr receiver, char* member);
void QTimer_Start2(QtObjectPtr ptr);
void QTimer_Start(QtObjectPtr ptr, int msec);
void QTimer_Stop(QtObjectPtr ptr);
void QTimer_ConnectTimeout(QtObjectPtr ptr);
void QTimer_DisconnectTimeout(QtObjectPtr ptr);
int QTimer_TimerId(QtObjectPtr ptr);
int QTimer_TimerType(QtObjectPtr ptr);
void QTimer_DestroyQTimer(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif