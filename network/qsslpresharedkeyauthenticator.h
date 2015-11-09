#ifdef __cplusplus
extern "C" {
#endif

void* QSslPreSharedKeyAuthenticator_NewQSslPreSharedKeyAuthenticator();
void* QSslPreSharedKeyAuthenticator_NewQSslPreSharedKeyAuthenticator2(void* authenticator);
void* QSslPreSharedKeyAuthenticator_Identity(void* ptr);
void* QSslPreSharedKeyAuthenticator_IdentityHint(void* ptr);
int QSslPreSharedKeyAuthenticator_MaximumIdentityLength(void* ptr);
int QSslPreSharedKeyAuthenticator_MaximumPreSharedKeyLength(void* ptr);
void* QSslPreSharedKeyAuthenticator_PreSharedKey(void* ptr);
void QSslPreSharedKeyAuthenticator_SetIdentity(void* ptr, void* identity);
void QSslPreSharedKeyAuthenticator_SetPreSharedKey(void* ptr, void* preSharedKey);
void QSslPreSharedKeyAuthenticator_Swap(void* ptr, void* authenticator);
void QSslPreSharedKeyAuthenticator_DestroyQSslPreSharedKeyAuthenticator(void* ptr);

#ifdef __cplusplus
}
#endif