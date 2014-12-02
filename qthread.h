#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Types
int QThread_IdlePriority();
int QThread_LowestPriority();
int QThread_LowPriority();
int QThread_NormalPriority();
int QThread_HighPriority();
int QThread_HighestPriority();
int QThread_TimeCriticalPriority();
int QThread_InheritPriority();
//Public Functions
QtObjectPtr QThread_New_QObject(QtObjectPtr parent);
void QThread_Destroy(QtObjectPtr ptr);
void QThread_Exit_Int(QtObjectPtr ptr, int returnCode);
int QThread_IsFinished(QtObjectPtr ptr);
int QThread_IsInterruptionRequested(QtObjectPtr ptr);
int QThread_IsRunning(QtObjectPtr ptr);
int QThread_Priority(QtObjectPtr ptr);
void QThread_RequestInterruption(QtObjectPtr ptr);
void QThread_SetPriority_Priority(QtObjectPtr ptr, int priority);
//Public Slots
void QThread_ConnectSlotQuit(QtObjectPtr ptr);
void QThread_DisconnectSlotQuit(QtObjectPtr ptr);
void QThread_Quit(QtObjectPtr ptr);
void QThread_ConnectSlotTerminate(QtObjectPtr ptr);
void QThread_DisconnectSlotTerminate(QtObjectPtr ptr);
void QThread_Terminate(QtObjectPtr ptr);
//Signals
void QThread_ConnectSignalFinished(QtObjectPtr ptr);
void QThread_DisconnectSignalFinished(QtObjectPtr ptr);
void QThread_ConnectSignalStarted(QtObjectPtr ptr);
void QThread_DisconnectSignalStarted(QtObjectPtr ptr);
//Static Public Members
QtObjectPtr QThread_CurrentThread();
int QThread_IdealThreadCount();

#ifdef __cplusplus
}
#endif
