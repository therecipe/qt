#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QLocalServer_SetSocketOptions(QtObjectPtr ptr, int options);
QtObjectPtr QLocalServer_NewQLocalServer(QtObjectPtr parent);
void QLocalServer_Close(QtObjectPtr ptr);
char* QLocalServer_ErrorString(QtObjectPtr ptr);
char* QLocalServer_FullServerName(QtObjectPtr ptr);
int QLocalServer_HasPendingConnections(QtObjectPtr ptr);
int QLocalServer_IsListening(QtObjectPtr ptr);
int QLocalServer_Listen(QtObjectPtr ptr, char* name);
int QLocalServer_MaxPendingConnections(QtObjectPtr ptr);
void QLocalServer_ConnectNewConnection(QtObjectPtr ptr);
void QLocalServer_DisconnectNewConnection(QtObjectPtr ptr);
QtObjectPtr QLocalServer_NextPendingConnection(QtObjectPtr ptr);
int QLocalServer_QLocalServer_RemoveServer(char* name);
int QLocalServer_ServerError(QtObjectPtr ptr);
char* QLocalServer_ServerName(QtObjectPtr ptr);
void QLocalServer_SetMaxPendingConnections(QtObjectPtr ptr, int numConnections);
int QLocalServer_SocketOptions(QtObjectPtr ptr);
int QLocalServer_WaitForNewConnection(QtObjectPtr ptr, int msec, int timedOut);
void QLocalServer_DestroyQLocalServer(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif