#ifdef __cplusplus
extern "C" {
#endif

void QLocalServer_SetSocketOptions(void* ptr, int options);
void* QLocalServer_NewQLocalServer(void* parent);
void QLocalServer_Close(void* ptr);
char* QLocalServer_ErrorString(void* ptr);
char* QLocalServer_FullServerName(void* ptr);
int QLocalServer_HasPendingConnections(void* ptr);
int QLocalServer_IsListening(void* ptr);
int QLocalServer_Listen(void* ptr, char* name);
int QLocalServer_MaxPendingConnections(void* ptr);
void QLocalServer_ConnectNewConnection(void* ptr);
void QLocalServer_DisconnectNewConnection(void* ptr);
void* QLocalServer_NextPendingConnection(void* ptr);
int QLocalServer_QLocalServer_RemoveServer(char* name);
int QLocalServer_ServerError(void* ptr);
char* QLocalServer_ServerName(void* ptr);
void QLocalServer_SetMaxPendingConnections(void* ptr, int numConnections);
int QLocalServer_SocketOptions(void* ptr);
int QLocalServer_WaitForNewConnection(void* ptr, int msec, int timedOut);
void QLocalServer_DestroyQLocalServer(void* ptr);

#ifdef __cplusplus
}
#endif