#ifdef __cplusplus
extern "C" {
#endif

int QAbstractAnimation_CurrentLoop(void* ptr);
int QAbstractAnimation_CurrentTime(void* ptr);
int QAbstractAnimation_Direction(void* ptr);
int QAbstractAnimation_LoopCount(void* ptr);
void QAbstractAnimation_SetCurrentTime(void* ptr, int msecs);
void QAbstractAnimation_SetDirection(void* ptr, int direction);
void QAbstractAnimation_SetLoopCount(void* ptr, int loopCount);
void QAbstractAnimation_ConnectCurrentLoopChanged(void* ptr);
void QAbstractAnimation_DisconnectCurrentLoopChanged(void* ptr);
int QAbstractAnimation_CurrentLoopTime(void* ptr);
void QAbstractAnimation_ConnectDirectionChanged(void* ptr);
void QAbstractAnimation_DisconnectDirectionChanged(void* ptr);
int QAbstractAnimation_Duration(void* ptr);
void QAbstractAnimation_ConnectFinished(void* ptr);
void QAbstractAnimation_DisconnectFinished(void* ptr);
void* QAbstractAnimation_Group(void* ptr);
void QAbstractAnimation_Pause(void* ptr);
void QAbstractAnimation_Resume(void* ptr);
void QAbstractAnimation_SetPaused(void* ptr, int paused);
void QAbstractAnimation_Start(void* ptr, int policy);
void QAbstractAnimation_ConnectStateChanged(void* ptr);
void QAbstractAnimation_DisconnectStateChanged(void* ptr);
void QAbstractAnimation_Stop(void* ptr);
int QAbstractAnimation_TotalDuration(void* ptr);
void QAbstractAnimation_DestroyQAbstractAnimation(void* ptr);

#ifdef __cplusplus
}
#endif