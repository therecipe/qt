#ifdef __cplusplus
extern "C" {
#endif

void QBasicTimer_Start(void* ptr, int msec, void* object);
void* QBasicTimer_NewQBasicTimer();
int QBasicTimer_IsActive(void* ptr);
void QBasicTimer_Start2(void* ptr, int msec, int timerType, void* obj);
void QBasicTimer_Stop(void* ptr);
int QBasicTimer_TimerId(void* ptr);
void QBasicTimer_DestroyQBasicTimer(void* ptr);

#ifdef __cplusplus
}
#endif