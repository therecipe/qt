#ifdef __cplusplus
extern "C" {
#endif

void* QWebSocketServer_NewQWebSocketServer(char* serverName, int secureMode, void* parent);
void QWebSocketServer_ConnectAcceptError(void* ptr);
void QWebSocketServer_DisconnectAcceptError(void* ptr);
void QWebSocketServer_Close(void* ptr);
void QWebSocketServer_ConnectClosed(void* ptr);
void QWebSocketServer_DisconnectClosed(void* ptr);
char* QWebSocketServer_ErrorString(void* ptr);
int QWebSocketServer_HasPendingConnections(void* ptr);
int QWebSocketServer_IsListening(void* ptr);
int QWebSocketServer_MaxPendingConnections(void* ptr);
void QWebSocketServer_ConnectNewConnection(void* ptr);
void QWebSocketServer_DisconnectNewConnection(void* ptr);
void* QWebSocketServer_NextPendingConnection(void* ptr);
void QWebSocketServer_ConnectOriginAuthenticationRequired(void* ptr);
void QWebSocketServer_DisconnectOriginAuthenticationRequired(void* ptr);
void QWebSocketServer_PauseAccepting(void* ptr);
void QWebSocketServer_ResumeAccepting(void* ptr);
int QWebSocketServer_SecureMode(void* ptr);
char* QWebSocketServer_ServerName(void* ptr);
void QWebSocketServer_SetMaxPendingConnections(void* ptr, int numConnections);
void QWebSocketServer_SetProxy(void* ptr, void* networkProxy);
void QWebSocketServer_SetServerName(void* ptr, char* serverName);
int QWebSocketServer_SetSocketDescriptor(void* ptr, int socketDescriptor);
void QWebSocketServer_SetSslConfiguration(void* ptr, void* sslConfiguration);
int QWebSocketServer_SocketDescriptor(void* ptr);
void QWebSocketServer_DestroyQWebSocketServer(void* ptr);

#ifdef __cplusplus
}
#endif