#ifdef __cplusplus
extern "C" {
#endif

void* QDBusServer_NewQDBusServer2(void* parent);
void* QDBusServer_NewQDBusServer(char* address, void* parent);
char* QDBusServer_Address(void* ptr);
int QDBusServer_IsAnonymousAuthenticationAllowed(void* ptr);
int QDBusServer_IsConnected(void* ptr);
void QDBusServer_SetAnonymousAuthenticationAllowed(void* ptr, int value);
void QDBusServer_DestroyQDBusServer(void* ptr);

#ifdef __cplusplus
}
#endif