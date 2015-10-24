#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator3(QtObjectPtr other);
QtObjectPtr QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator(char* origin);
QtObjectPtr QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator2(QtObjectPtr other);
int QWebSocketCorsAuthenticator_Allowed(QtObjectPtr ptr);
char* QWebSocketCorsAuthenticator_Origin(QtObjectPtr ptr);
void QWebSocketCorsAuthenticator_SetAllowed(QtObjectPtr ptr, int allowed);
void QWebSocketCorsAuthenticator_Swap(QtObjectPtr ptr, QtObjectPtr other);
void QWebSocketCorsAuthenticator_DestroyQWebSocketCorsAuthenticator(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif