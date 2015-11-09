#ifdef __cplusplus
extern "C" {
#endif

int QTimer_RemainingTime(void* ptr);
void QTimer_SetInterval(void* ptr, int msec);
void* QTimer_NewQTimer(void* parent);
int QTimer_Interval(void* ptr);
int QTimer_IsActive(void* ptr);
int QTimer_IsSingleShot(void* ptr);
void QTimer_SetSingleShot(void* ptr, int singleShot);
void QTimer_SetTimerType(void* ptr, int atype);
void QTimer_QTimer_SingleShot2(int msec, int timerType, void* receiver, char* member);
void QTimer_QTimer_SingleShot(int msec, void* receiver, char* member);
void QTimer_Start2(void* ptr);
void QTimer_Start(void* ptr, int msec);
void QTimer_Stop(void* ptr);
void QTimer_ConnectTimeout(void* ptr);
void QTimer_DisconnectTimeout(void* ptr);
int QTimer_TimerId(void* ptr);
int QTimer_TimerType(void* ptr);
void QTimer_DestroyQTimer(void* ptr);

#ifdef __cplusplus
}
#endif