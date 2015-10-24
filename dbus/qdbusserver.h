#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDBusServer_NewQDBusServer2(QtObjectPtr parent);
QtObjectPtr QDBusServer_NewQDBusServer(char* address, QtObjectPtr parent);
char* QDBusServer_Address(QtObjectPtr ptr);
int QDBusServer_IsAnonymousAuthenticationAllowed(QtObjectPtr ptr);
int QDBusServer_IsConnected(QtObjectPtr ptr);
void QDBusServer_SetAnonymousAuthenticationAllowed(QtObjectPtr ptr, int value);
void QDBusServer_DestroyQDBusServer(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif