#ifdef __cplusplus
extern "C" {
#endif

int QSessionManager_RestartHint(void* ptr);
char* QSessionManager_SessionKey(void* ptr);
int QSessionManager_AllowsErrorInteraction(void* ptr);
int QSessionManager_AllowsInteraction(void* ptr);
void QSessionManager_Cancel(void* ptr);
char* QSessionManager_DiscardCommand(void* ptr);
int QSessionManager_IsPhase2(void* ptr);
void QSessionManager_Release(void* ptr);
void QSessionManager_RequestPhase2(void* ptr);
char* QSessionManager_RestartCommand(void* ptr);
char* QSessionManager_SessionId(void* ptr);
void QSessionManager_SetDiscardCommand(void* ptr, char* command);
void QSessionManager_SetManagerProperty2(void* ptr, char* name, char* value);
void QSessionManager_SetManagerProperty(void* ptr, char* name, char* value);
void QSessionManager_SetRestartCommand(void* ptr, char* command);
void QSessionManager_SetRestartHint(void* ptr, int hint);

#ifdef __cplusplus
}
#endif