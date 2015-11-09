#ifdef __cplusplus
extern "C" {
#endif

void* QSslConfiguration_NewQSslConfiguration();
void* QSslConfiguration_NewQSslConfiguration2(void* other);
int QSslConfiguration_IsNull(void* ptr);
void* QSslConfiguration_NextNegotiatedProtocol(void* ptr);
int QSslConfiguration_NextProtocolNegotiationStatus(void* ptr);
int QSslConfiguration_PeerVerifyDepth(void* ptr);
int QSslConfiguration_PeerVerifyMode(void* ptr);
void* QSslConfiguration_SessionTicket(void* ptr);
int QSslConfiguration_SessionTicketLifeTimeHint(void* ptr);
void QSslConfiguration_QSslConfiguration_SetDefaultConfiguration(void* configuration);
void QSslConfiguration_SetLocalCertificate(void* ptr, void* certificate);
void QSslConfiguration_SetPeerVerifyDepth(void* ptr, int depth);
void QSslConfiguration_SetPeerVerifyMode(void* ptr, int mode);
void QSslConfiguration_SetPrivateKey(void* ptr, void* key);
void QSslConfiguration_SetSessionTicket(void* ptr, void* sessionTicket);
void QSslConfiguration_Swap(void* ptr, void* other);
void QSslConfiguration_DestroyQSslConfiguration(void* ptr);

#ifdef __cplusplus
}
#endif