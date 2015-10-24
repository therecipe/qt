#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QAbstractSocket_NewQAbstractSocket(int socketType, QtObjectPtr parent);
void QAbstractSocket_Abort(QtObjectPtr ptr);
int QAbstractSocket_AtEnd(QtObjectPtr ptr);
int QAbstractSocket_CanReadLine(QtObjectPtr ptr);
void QAbstractSocket_Close(QtObjectPtr ptr);
void QAbstractSocket_ConnectConnected(QtObjectPtr ptr);
void QAbstractSocket_DisconnectConnected(QtObjectPtr ptr);
void QAbstractSocket_DisconnectFromHost(QtObjectPtr ptr);
void QAbstractSocket_ConnectDisconnected(QtObjectPtr ptr);
void QAbstractSocket_DisconnectDisconnected(QtObjectPtr ptr);
int QAbstractSocket_Error(QtObjectPtr ptr);
int QAbstractSocket_Flush(QtObjectPtr ptr);
void QAbstractSocket_ConnectHostFound(QtObjectPtr ptr);
void QAbstractSocket_DisconnectHostFound(QtObjectPtr ptr);
int QAbstractSocket_IsSequential(QtObjectPtr ptr);
int QAbstractSocket_IsValid(QtObjectPtr ptr);
int QAbstractSocket_PauseMode(QtObjectPtr ptr);
char* QAbstractSocket_PeerName(QtObjectPtr ptr);
void QAbstractSocket_Resume(QtObjectPtr ptr);
void QAbstractSocket_SetPauseMode(QtObjectPtr ptr, int pauseMode);
void QAbstractSocket_SetProxy(QtObjectPtr ptr, QtObjectPtr networkProxy);
void QAbstractSocket_SetSocketOption(QtObjectPtr ptr, int option, char* value);
char* QAbstractSocket_SocketOption(QtObjectPtr ptr, int option);
int QAbstractSocket_SocketType(QtObjectPtr ptr);
void QAbstractSocket_ConnectStateChanged(QtObjectPtr ptr);
void QAbstractSocket_DisconnectStateChanged(QtObjectPtr ptr);
int QAbstractSocket_WaitForBytesWritten(QtObjectPtr ptr, int msecs);
int QAbstractSocket_WaitForConnected(QtObjectPtr ptr, int msecs);
int QAbstractSocket_WaitForDisconnected(QtObjectPtr ptr, int msecs);
int QAbstractSocket_WaitForReadyRead(QtObjectPtr ptr, int msecs);
void QAbstractSocket_DestroyQAbstractSocket(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif