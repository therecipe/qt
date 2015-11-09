#ifdef __cplusplus
extern "C" {
#endif

int QLocalSocket_Open(void* ptr, int openMode);
void* QLocalSocket_NewQLocalSocket(void* parent);
void QLocalSocket_ConnectToServer2(void* ptr, char* name, int openMode);
void QLocalSocket_ConnectConnected(void* ptr);
void QLocalSocket_DisconnectConnected(void* ptr);
void QLocalSocket_ConnectDisconnected(void* ptr);
void QLocalSocket_DisconnectDisconnected(void* ptr);
char* QLocalSocket_FullServerName(void* ptr);
int QLocalSocket_IsSequential(void* ptr);
char* QLocalSocket_ServerName(void* ptr);
void QLocalSocket_SetServerName(void* ptr, char* name);
void QLocalSocket_ConnectStateChanged(void* ptr);
void QLocalSocket_DisconnectStateChanged(void* ptr);
void QLocalSocket_DestroyQLocalSocket(void* ptr);
void QLocalSocket_Abort(void* ptr);
int QLocalSocket_CanReadLine(void* ptr);
void QLocalSocket_Close(void* ptr);
void QLocalSocket_ConnectToServer(void* ptr, int openMode);
void QLocalSocket_DisconnectFromServer(void* ptr);
int QLocalSocket_Error(void* ptr);
int QLocalSocket_Flush(void* ptr);
int QLocalSocket_IsValid(void* ptr);
int QLocalSocket_WaitForBytesWritten(void* ptr, int msecs);
int QLocalSocket_WaitForConnected(void* ptr, int msecs);
int QLocalSocket_WaitForDisconnected(void* ptr, int msecs);
int QLocalSocket_WaitForReadyRead(void* ptr, int msecs);

#ifdef __cplusplus
}
#endif