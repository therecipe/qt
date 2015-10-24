#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QTcpServer_NewQTcpServer(QtObjectPtr parent);
void QTcpServer_ConnectAcceptError(QtObjectPtr ptr);
void QTcpServer_DisconnectAcceptError(QtObjectPtr ptr);
void QTcpServer_Close(QtObjectPtr ptr);
char* QTcpServer_ErrorString(QtObjectPtr ptr);
int QTcpServer_HasPendingConnections(QtObjectPtr ptr);
int QTcpServer_IsListening(QtObjectPtr ptr);
int QTcpServer_MaxPendingConnections(QtObjectPtr ptr);
void QTcpServer_ConnectNewConnection(QtObjectPtr ptr);
void QTcpServer_DisconnectNewConnection(QtObjectPtr ptr);
QtObjectPtr QTcpServer_NextPendingConnection(QtObjectPtr ptr);
void QTcpServer_PauseAccepting(QtObjectPtr ptr);
void QTcpServer_ResumeAccepting(QtObjectPtr ptr);
int QTcpServer_ServerError(QtObjectPtr ptr);
void QTcpServer_SetMaxPendingConnections(QtObjectPtr ptr, int numConnections);
void QTcpServer_SetProxy(QtObjectPtr ptr, QtObjectPtr networkProxy);
int QTcpServer_WaitForNewConnection(QtObjectPtr ptr, int msec, int timedOut);
void QTcpServer_DestroyQTcpServer(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif