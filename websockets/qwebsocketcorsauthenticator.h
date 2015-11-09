#ifdef __cplusplus
extern "C" {
#endif

void* QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator3(void* other);
void* QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator(char* origin);
void* QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator2(void* other);
int QWebSocketCorsAuthenticator_Allowed(void* ptr);
char* QWebSocketCorsAuthenticator_Origin(void* ptr);
void QWebSocketCorsAuthenticator_SetAllowed(void* ptr, int allowed);
void QWebSocketCorsAuthenticator_Swap(void* ptr, void* other);
void QWebSocketCorsAuthenticator_DestroyQWebSocketCorsAuthenticator(void* ptr);

#ifdef __cplusplus
}
#endif