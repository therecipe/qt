#ifdef __cplusplus
extern "C" {
#endif

void* QStateMachine_NewQStateMachine(void* parent);
void* QStateMachine_NewQStateMachine2(int childMode, void* parent);
void QStateMachine_AddDefaultAnimation(void* ptr, void* animation);
void QStateMachine_AddState(void* ptr, void* state);
void QStateMachine_ClearError(void* ptr);
int QStateMachine_CancelDelayedEvent(void* ptr, int id);
int QStateMachine_Error(void* ptr);
char* QStateMachine_ErrorString(void* ptr);
int QStateMachine_EventFilter(void* ptr, void* watched, void* event);
int QStateMachine_GlobalRestorePolicy(void* ptr);
int QStateMachine_IsAnimated(void* ptr);
int QStateMachine_IsRunning(void* ptr);
int QStateMachine_PostDelayedEvent(void* ptr, void* event, int delay);
void QStateMachine_PostEvent(void* ptr, void* event, int priority);
void QStateMachine_RemoveDefaultAnimation(void* ptr, void* animation);
void QStateMachine_RemoveState(void* ptr, void* state);
void QStateMachine_ConnectRunningChanged(void* ptr);
void QStateMachine_DisconnectRunningChanged(void* ptr);
void QStateMachine_SetAnimated(void* ptr, int enabled);
void QStateMachine_SetGlobalRestorePolicy(void* ptr, int restorePolicy);
void QStateMachine_SetRunning(void* ptr, int running);
void QStateMachine_Start(void* ptr);
void QStateMachine_ConnectStarted(void* ptr);
void QStateMachine_DisconnectStarted(void* ptr);
void QStateMachine_Stop(void* ptr);
void QStateMachine_ConnectStopped(void* ptr);
void QStateMachine_DisconnectStopped(void* ptr);
void QStateMachine_DestroyQStateMachine(void* ptr);

#ifdef __cplusplus
}
#endif