#ifdef __cplusplus
extern "C" {
#endif

void* QTcpServer_NewQTcpServer(void* parent);
void QTcpServer_ConnectAcceptError(void* ptr);
void QTcpServer_DisconnectAcceptError(void* ptr);
void QTcpServer_Close(void* ptr);
char* QTcpServer_ErrorString(void* ptr);
int QTcpServer_HasPendingConnections(void* ptr);
int QTcpServer_IsListening(void* ptr);
int QTcpServer_MaxPendingConnections(void* ptr);
void QTcpServer_ConnectNewConnection(void* ptr);
void QTcpServer_DisconnectNewConnection(void* ptr);
void* QTcpServer_NextPendingConnection(void* ptr);
void QTcpServer_PauseAccepting(void* ptr);
void QTcpServer_ResumeAccepting(void* ptr);
int QTcpServer_ServerError(void* ptr);
void QTcpServer_SetMaxPendingConnections(void* ptr, int numConnections);
void QTcpServer_SetProxy(void* ptr, void* networkProxy);
int QTcpServer_WaitForNewConnection(void* ptr, int msec, int timedOut);
void QTcpServer_DestroyQTcpServer(void* ptr);

#ifdef __cplusplus
}
#endif