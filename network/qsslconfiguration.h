#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSslConfiguration_NewQSslConfiguration();
QtObjectPtr QSslConfiguration_NewQSslConfiguration2(QtObjectPtr other);
int QSslConfiguration_IsNull(QtObjectPtr ptr);
int QSslConfiguration_NextProtocolNegotiationStatus(QtObjectPtr ptr);
int QSslConfiguration_PeerVerifyDepth(QtObjectPtr ptr);
int QSslConfiguration_PeerVerifyMode(QtObjectPtr ptr);
int QSslConfiguration_SessionTicketLifeTimeHint(QtObjectPtr ptr);
void QSslConfiguration_QSslConfiguration_SetDefaultConfiguration(QtObjectPtr configuration);
void QSslConfiguration_SetLocalCertificate(QtObjectPtr ptr, QtObjectPtr certificate);
void QSslConfiguration_SetPeerVerifyDepth(QtObjectPtr ptr, int depth);
void QSslConfiguration_SetPeerVerifyMode(QtObjectPtr ptr, int mode);
void QSslConfiguration_SetPrivateKey(QtObjectPtr ptr, QtObjectPtr key);
void QSslConfiguration_SetSessionTicket(QtObjectPtr ptr, QtObjectPtr sessionTicket);
void QSslConfiguration_Swap(QtObjectPtr ptr, QtObjectPtr other);
void QSslConfiguration_DestroyQSslConfiguration(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif