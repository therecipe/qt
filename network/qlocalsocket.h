#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QLocalSocket_Open(QtObjectPtr ptr, int openMode);
QtObjectPtr QLocalSocket_NewQLocalSocket(QtObjectPtr parent);
void QLocalSocket_ConnectToServer2(QtObjectPtr ptr, char* name, int openMode);
void QLocalSocket_ConnectConnected(QtObjectPtr ptr);
void QLocalSocket_DisconnectConnected(QtObjectPtr ptr);
void QLocalSocket_ConnectDisconnected(QtObjectPtr ptr);
void QLocalSocket_DisconnectDisconnected(QtObjectPtr ptr);
char* QLocalSocket_FullServerName(QtObjectPtr ptr);
int QLocalSocket_IsSequential(QtObjectPtr ptr);
char* QLocalSocket_ServerName(QtObjectPtr ptr);
void QLocalSocket_SetServerName(QtObjectPtr ptr, char* name);
void QLocalSocket_ConnectStateChanged(QtObjectPtr ptr);
void QLocalSocket_DisconnectStateChanged(QtObjectPtr ptr);
void QLocalSocket_DestroyQLocalSocket(QtObjectPtr ptr);
void QLocalSocket_Abort(QtObjectPtr ptr);
int QLocalSocket_CanReadLine(QtObjectPtr ptr);
void QLocalSocket_Close(QtObjectPtr ptr);
void QLocalSocket_ConnectToServer(QtObjectPtr ptr, int openMode);
void QLocalSocket_DisconnectFromServer(QtObjectPtr ptr);
int QLocalSocket_Error(QtObjectPtr ptr);
int QLocalSocket_Flush(QtObjectPtr ptr);
int QLocalSocket_IsValid(QtObjectPtr ptr);
int QLocalSocket_WaitForBytesWritten(QtObjectPtr ptr, int msecs);
int QLocalSocket_WaitForConnected(QtObjectPtr ptr, int msecs);
int QLocalSocket_WaitForDisconnected(QtObjectPtr ptr, int msecs);
int QLocalSocket_WaitForReadyRead(QtObjectPtr ptr, int msecs);

#ifdef __cplusplus
}
#endif