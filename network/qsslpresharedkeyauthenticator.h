#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSslPreSharedKeyAuthenticator_NewQSslPreSharedKeyAuthenticator();
QtObjectPtr QSslPreSharedKeyAuthenticator_NewQSslPreSharedKeyAuthenticator2(QtObjectPtr authenticator);
int QSslPreSharedKeyAuthenticator_MaximumIdentityLength(QtObjectPtr ptr);
int QSslPreSharedKeyAuthenticator_MaximumPreSharedKeyLength(QtObjectPtr ptr);
void QSslPreSharedKeyAuthenticator_SetIdentity(QtObjectPtr ptr, QtObjectPtr identity);
void QSslPreSharedKeyAuthenticator_SetPreSharedKey(QtObjectPtr ptr, QtObjectPtr preSharedKey);
void QSslPreSharedKeyAuthenticator_Swap(QtObjectPtr ptr, QtObjectPtr authenticator);
void QSslPreSharedKeyAuthenticator_DestroyQSslPreSharedKeyAuthenticator(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif