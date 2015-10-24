#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QThread_SetPriority(QtObjectPtr ptr, int priority);
QtObjectPtr QThread_NewQThread(QtObjectPtr parent);
QtObjectPtr QThread_QThread_CurrentThread();
int QThread_Event(QtObjectPtr ptr, QtObjectPtr event);
QtObjectPtr QThread_EventDispatcher(QtObjectPtr ptr);
void QThread_Exit(QtObjectPtr ptr, int returnCode);
void QThread_ConnectFinished(QtObjectPtr ptr);
void QThread_DisconnectFinished(QtObjectPtr ptr);
int QThread_IsFinished(QtObjectPtr ptr);
int QThread_IsInterruptionRequested(QtObjectPtr ptr);
int QThread_IsRunning(QtObjectPtr ptr);
int QThread_LoopLevel(QtObjectPtr ptr);
int QThread_Priority(QtObjectPtr ptr);
void QThread_Quit(QtObjectPtr ptr);
void QThread_RequestInterruption(QtObjectPtr ptr);
void QThread_SetEventDispatcher(QtObjectPtr ptr, QtObjectPtr eventDispatcher);
void QThread_ConnectStarted(QtObjectPtr ptr);
void QThread_DisconnectStarted(QtObjectPtr ptr);
void QThread_DestroyQThread(QtObjectPtr ptr);
int QThread_QThread_IdealThreadCount();
void QThread_Start(QtObjectPtr ptr, int priority);
void QThread_Terminate(QtObjectPtr ptr);
void QThread_QThread_YieldCurrentThread();

#ifdef __cplusplus
}
#endif