#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QSessionManager_RestartHint(QtObjectPtr ptr);
char* QSessionManager_SessionKey(QtObjectPtr ptr);
int QSessionManager_AllowsErrorInteraction(QtObjectPtr ptr);
int QSessionManager_AllowsInteraction(QtObjectPtr ptr);
void QSessionManager_Cancel(QtObjectPtr ptr);
char* QSessionManager_DiscardCommand(QtObjectPtr ptr);
int QSessionManager_IsPhase2(QtObjectPtr ptr);
void QSessionManager_Release(QtObjectPtr ptr);
void QSessionManager_RequestPhase2(QtObjectPtr ptr);
char* QSessionManager_RestartCommand(QtObjectPtr ptr);
char* QSessionManager_SessionId(QtObjectPtr ptr);
void QSessionManager_SetDiscardCommand(QtObjectPtr ptr, char* command);
void QSessionManager_SetManagerProperty2(QtObjectPtr ptr, char* name, char* value);
void QSessionManager_SetManagerProperty(QtObjectPtr ptr, char* name, char* value);
void QSessionManager_SetRestartCommand(QtObjectPtr ptr, char* command);
void QSessionManager_SetRestartHint(QtObjectPtr ptr, int hint);

#ifdef __cplusplus
}
#endif