#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QWebSocketServer_NewQWebSocketServer(char* serverName, int secureMode, QtObjectPtr parent);
void QWebSocketServer_ConnectAcceptError(QtObjectPtr ptr);
void QWebSocketServer_DisconnectAcceptError(QtObjectPtr ptr);
void QWebSocketServer_Close(QtObjectPtr ptr);
void QWebSocketServer_ConnectClosed(QtObjectPtr ptr);
void QWebSocketServer_DisconnectClosed(QtObjectPtr ptr);
char* QWebSocketServer_ErrorString(QtObjectPtr ptr);
int QWebSocketServer_HasPendingConnections(QtObjectPtr ptr);
int QWebSocketServer_IsListening(QtObjectPtr ptr);
int QWebSocketServer_MaxPendingConnections(QtObjectPtr ptr);
void QWebSocketServer_ConnectNewConnection(QtObjectPtr ptr);
void QWebSocketServer_DisconnectNewConnection(QtObjectPtr ptr);
QtObjectPtr QWebSocketServer_NextPendingConnection(QtObjectPtr ptr);
void QWebSocketServer_ConnectOriginAuthenticationRequired(QtObjectPtr ptr);
void QWebSocketServer_DisconnectOriginAuthenticationRequired(QtObjectPtr ptr);
void QWebSocketServer_PauseAccepting(QtObjectPtr ptr);
void QWebSocketServer_ResumeAccepting(QtObjectPtr ptr);
int QWebSocketServer_SecureMode(QtObjectPtr ptr);
char* QWebSocketServer_ServerName(QtObjectPtr ptr);
char* QWebSocketServer_ServerUrl(QtObjectPtr ptr);
void QWebSocketServer_SetMaxPendingConnections(QtObjectPtr ptr, int numConnections);
void QWebSocketServer_SetProxy(QtObjectPtr ptr, QtObjectPtr networkProxy);
void QWebSocketServer_SetServerName(QtObjectPtr ptr, char* serverName);
int QWebSocketServer_SetSocketDescriptor(QtObjectPtr ptr, int socketDescriptor);
void QWebSocketServer_SetSslConfiguration(QtObjectPtr ptr, QtObjectPtr sslConfiguration);
int QWebSocketServer_SocketDescriptor(QtObjectPtr ptr);
void QWebSocketServer_DestroyQWebSocketServer(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif