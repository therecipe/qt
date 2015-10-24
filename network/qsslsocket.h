#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSslSocket_NewQSslSocket(QtObjectPtr parent);
void QSslSocket_Abort(QtObjectPtr ptr);
void QSslSocket_AddCaCertificate(QtObjectPtr ptr, QtObjectPtr certificate);
void QSslSocket_QSslSocket_AddDefaultCaCertificate(QtObjectPtr certificate);
int QSslSocket_AtEnd(QtObjectPtr ptr);
int QSslSocket_CanReadLine(QtObjectPtr ptr);
void QSslSocket_Close(QtObjectPtr ptr);
void QSslSocket_ConnectEncrypted(QtObjectPtr ptr);
void QSslSocket_DisconnectEncrypted(QtObjectPtr ptr);
int QSslSocket_Flush(QtObjectPtr ptr);
void QSslSocket_IgnoreSslErrors(QtObjectPtr ptr);
int QSslSocket_IsEncrypted(QtObjectPtr ptr);
int QSslSocket_Mode(QtObjectPtr ptr);
void QSslSocket_ConnectModeChanged(QtObjectPtr ptr);
void QSslSocket_DisconnectModeChanged(QtObjectPtr ptr);
int QSslSocket_PeerVerifyDepth(QtObjectPtr ptr);
int QSslSocket_PeerVerifyMode(QtObjectPtr ptr);
char* QSslSocket_PeerVerifyName(QtObjectPtr ptr);
void QSslSocket_ConnectPreSharedKeyAuthenticationRequired(QtObjectPtr ptr);
void QSslSocket_DisconnectPreSharedKeyAuthenticationRequired(QtObjectPtr ptr);
void QSslSocket_Resume(QtObjectPtr ptr);
void QSslSocket_SetLocalCertificate(QtObjectPtr ptr, QtObjectPtr certificate);
void QSslSocket_SetPeerVerifyDepth(QtObjectPtr ptr, int depth);
void QSslSocket_SetPeerVerifyMode(QtObjectPtr ptr, int mode);
void QSslSocket_SetPeerVerifyName(QtObjectPtr ptr, char* hostName);
void QSslSocket_SetPrivateKey(QtObjectPtr ptr, QtObjectPtr key);
void QSslSocket_SetSocketOption(QtObjectPtr ptr, int option, char* value);
void QSslSocket_SetSslConfiguration(QtObjectPtr ptr, QtObjectPtr configuration);
char* QSslSocket_SocketOption(QtObjectPtr ptr, int option);
char* QSslSocket_QSslSocket_SslLibraryBuildVersionString();
char* QSslSocket_QSslSocket_SslLibraryVersionString();
void QSslSocket_StartClientEncryption(QtObjectPtr ptr);
void QSslSocket_StartServerEncryption(QtObjectPtr ptr);
int QSslSocket_QSslSocket_SupportsSsl();
int QSslSocket_WaitForBytesWritten(QtObjectPtr ptr, int msecs);
int QSslSocket_WaitForConnected(QtObjectPtr ptr, int msecs);
int QSslSocket_WaitForDisconnected(QtObjectPtr ptr, int msecs);
int QSslSocket_WaitForEncrypted(QtObjectPtr ptr, int msecs);
int QSslSocket_WaitForReadyRead(QtObjectPtr ptr, int msecs);
void QSslSocket_DestroyQSslSocket(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif