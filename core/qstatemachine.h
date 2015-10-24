#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QStateMachine_NewQStateMachine(QtObjectPtr parent);
QtObjectPtr QStateMachine_NewQStateMachine2(int childMode, QtObjectPtr parent);
void QStateMachine_AddDefaultAnimation(QtObjectPtr ptr, QtObjectPtr animation);
void QStateMachine_AddState(QtObjectPtr ptr, QtObjectPtr state);
void QStateMachine_ClearError(QtObjectPtr ptr);
int QStateMachine_CancelDelayedEvent(QtObjectPtr ptr, int id);
int QStateMachine_Error(QtObjectPtr ptr);
char* QStateMachine_ErrorString(QtObjectPtr ptr);
int QStateMachine_EventFilter(QtObjectPtr ptr, QtObjectPtr watched, QtObjectPtr event);
int QStateMachine_GlobalRestorePolicy(QtObjectPtr ptr);
int QStateMachine_IsAnimated(QtObjectPtr ptr);
int QStateMachine_IsRunning(QtObjectPtr ptr);
int QStateMachine_PostDelayedEvent(QtObjectPtr ptr, QtObjectPtr event, int delay);
void QStateMachine_PostEvent(QtObjectPtr ptr, QtObjectPtr event, int priority);
void QStateMachine_RemoveDefaultAnimation(QtObjectPtr ptr, QtObjectPtr animation);
void QStateMachine_RemoveState(QtObjectPtr ptr, QtObjectPtr state);
void QStateMachine_ConnectRunningChanged(QtObjectPtr ptr);
void QStateMachine_DisconnectRunningChanged(QtObjectPtr ptr);
void QStateMachine_SetAnimated(QtObjectPtr ptr, int enabled);
void QStateMachine_SetGlobalRestorePolicy(QtObjectPtr ptr, int restorePolicy);
void QStateMachine_SetRunning(QtObjectPtr ptr, int running);
void QStateMachine_Start(QtObjectPtr ptr);
void QStateMachine_ConnectStarted(QtObjectPtr ptr);
void QStateMachine_DisconnectStarted(QtObjectPtr ptr);
void QStateMachine_Stop(QtObjectPtr ptr);
void QStateMachine_ConnectStopped(QtObjectPtr ptr);
void QStateMachine_DisconnectStopped(QtObjectPtr ptr);
void QStateMachine_DestroyQStateMachine(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif