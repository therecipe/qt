#ifdef __cplusplus
extern "C" {
#endif

void QThread_SetPriority(void* ptr, int priority);
void* QThread_NewQThread(void* parent);
void* QThread_QThread_CurrentThread();
int QThread_Event(void* ptr, void* event);
void* QThread_EventDispatcher(void* ptr);
void QThread_Exit(void* ptr, int returnCode);
void QThread_ConnectFinished(void* ptr);
void QThread_DisconnectFinished(void* ptr);
int QThread_IsFinished(void* ptr);
int QThread_IsInterruptionRequested(void* ptr);
int QThread_IsRunning(void* ptr);
int QThread_LoopLevel(void* ptr);
int QThread_Priority(void* ptr);
void QThread_Quit(void* ptr);
void QThread_RequestInterruption(void* ptr);
void QThread_SetEventDispatcher(void* ptr, void* eventDispatcher);
void QThread_ConnectStarted(void* ptr);
void QThread_DisconnectStarted(void* ptr);
void QThread_DestroyQThread(void* ptr);
int QThread_QThread_IdealThreadCount();
void QThread_Start(void* ptr, int priority);
void QThread_Terminate(void* ptr);
void QThread_QThread_YieldCurrentThread();

#ifdef __cplusplus
}
#endif