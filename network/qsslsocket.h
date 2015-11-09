#ifdef __cplusplus
extern "C" {
#endif

void* QSslSocket_NewQSslSocket(void* parent);
void QSslSocket_Abort(void* ptr);
void QSslSocket_AddCaCertificate(void* ptr, void* certificate);
void QSslSocket_QSslSocket_AddDefaultCaCertificate(void* certificate);
int QSslSocket_AtEnd(void* ptr);
int QSslSocket_CanReadLine(void* ptr);
void QSslSocket_Close(void* ptr);
void QSslSocket_ConnectEncrypted(void* ptr);
void QSslSocket_DisconnectEncrypted(void* ptr);
int QSslSocket_Flush(void* ptr);
void QSslSocket_IgnoreSslErrors(void* ptr);
int QSslSocket_IsEncrypted(void* ptr);
int QSslSocket_Mode(void* ptr);
void QSslSocket_ConnectModeChanged(void* ptr);
void QSslSocket_DisconnectModeChanged(void* ptr);
int QSslSocket_PeerVerifyDepth(void* ptr);
int QSslSocket_PeerVerifyMode(void* ptr);
char* QSslSocket_PeerVerifyName(void* ptr);
void QSslSocket_ConnectPreSharedKeyAuthenticationRequired(void* ptr);
void QSslSocket_DisconnectPreSharedKeyAuthenticationRequired(void* ptr);
void QSslSocket_Resume(void* ptr);
void QSslSocket_SetLocalCertificate(void* ptr, void* certificate);
void QSslSocket_SetPeerVerifyDepth(void* ptr, int depth);
void QSslSocket_SetPeerVerifyMode(void* ptr, int mode);
void QSslSocket_SetPeerVerifyName(void* ptr, char* hostName);
void QSslSocket_SetPrivateKey(void* ptr, void* key);
void QSslSocket_SetSocketOption(void* ptr, int option, void* value);
void QSslSocket_SetSslConfiguration(void* ptr, void* configuration);
void* QSslSocket_SocketOption(void* ptr, int option);
char* QSslSocket_QSslSocket_SslLibraryBuildVersionString();
char* QSslSocket_QSslSocket_SslLibraryVersionString();
void QSslSocket_StartClientEncryption(void* ptr);
void QSslSocket_StartServerEncryption(void* ptr);
int QSslSocket_QSslSocket_SupportsSsl();
int QSslSocket_WaitForBytesWritten(void* ptr, int msecs);
int QSslSocket_WaitForConnected(void* ptr, int msecs);
int QSslSocket_WaitForDisconnected(void* ptr, int msecs);
int QSslSocket_WaitForEncrypted(void* ptr, int msecs);
int QSslSocket_WaitForReadyRead(void* ptr, int msecs);
void QSslSocket_DestroyQSslSocket(void* ptr);

#ifdef __cplusplus
}
#endif