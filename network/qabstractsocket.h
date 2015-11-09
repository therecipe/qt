#ifdef __cplusplus
extern "C" {
#endif

void* QAbstractSocket_NewQAbstractSocket(int socketType, void* parent);
void QAbstractSocket_Abort(void* ptr);
int QAbstractSocket_AtEnd(void* ptr);
int QAbstractSocket_CanReadLine(void* ptr);
void QAbstractSocket_Close(void* ptr);
void QAbstractSocket_ConnectConnected(void* ptr);
void QAbstractSocket_DisconnectConnected(void* ptr);
void QAbstractSocket_DisconnectFromHost(void* ptr);
void QAbstractSocket_ConnectDisconnected(void* ptr);
void QAbstractSocket_DisconnectDisconnected(void* ptr);
int QAbstractSocket_Error(void* ptr);
int QAbstractSocket_Flush(void* ptr);
void QAbstractSocket_ConnectHostFound(void* ptr);
void QAbstractSocket_DisconnectHostFound(void* ptr);
int QAbstractSocket_IsSequential(void* ptr);
int QAbstractSocket_IsValid(void* ptr);
int QAbstractSocket_PauseMode(void* ptr);
char* QAbstractSocket_PeerName(void* ptr);
void QAbstractSocket_Resume(void* ptr);
void QAbstractSocket_SetPauseMode(void* ptr, int pauseMode);
void QAbstractSocket_SetProxy(void* ptr, void* networkProxy);
void QAbstractSocket_SetSocketOption(void* ptr, int option, void* value);
void* QAbstractSocket_SocketOption(void* ptr, int option);
int QAbstractSocket_SocketType(void* ptr);
void QAbstractSocket_ConnectStateChanged(void* ptr);
void QAbstractSocket_DisconnectStateChanged(void* ptr);
int QAbstractSocket_WaitForBytesWritten(void* ptr, int msecs);
int QAbstractSocket_WaitForConnected(void* ptr, int msecs);
int QAbstractSocket_WaitForDisconnected(void* ptr, int msecs);
int QAbstractSocket_WaitForReadyRead(void* ptr, int msecs);
void QAbstractSocket_DestroyQAbstractSocket(void* ptr);

#ifdef __cplusplus
}
#endif