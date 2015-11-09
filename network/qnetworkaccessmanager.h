#ifdef __cplusplus
extern "C" {
#endif

void* QNetworkAccessManager_ProxyFactory(void* ptr);
void* QNetworkAccessManager_NewQNetworkAccessManager(void* parent);
void QNetworkAccessManager_ConnectAuthenticationRequired(void* ptr);
void QNetworkAccessManager_DisconnectAuthenticationRequired(void* ptr);
void* QNetworkAccessManager_Cache(void* ptr);
void QNetworkAccessManager_ClearAccessCache(void* ptr);
void* QNetworkAccessManager_CookieJar(void* ptr);
void* QNetworkAccessManager_DeleteResource(void* ptr, void* request);
void QNetworkAccessManager_ConnectEncrypted(void* ptr);
void QNetworkAccessManager_DisconnectEncrypted(void* ptr);
void QNetworkAccessManager_ConnectFinished(void* ptr);
void QNetworkAccessManager_DisconnectFinished(void* ptr);
void* QNetworkAccessManager_Get(void* ptr, void* request);
void* QNetworkAccessManager_Head(void* ptr, void* request);
int QNetworkAccessManager_NetworkAccessible(void* ptr);
void QNetworkAccessManager_ConnectNetworkAccessibleChanged(void* ptr);
void QNetworkAccessManager_DisconnectNetworkAccessibleChanged(void* ptr);
void* QNetworkAccessManager_Post3(void* ptr, void* request, void* multiPart);
void* QNetworkAccessManager_Post(void* ptr, void* request, void* data);
void* QNetworkAccessManager_Post2(void* ptr, void* request, void* data);
void QNetworkAccessManager_ConnectPreSharedKeyAuthenticationRequired(void* ptr);
void QNetworkAccessManager_DisconnectPreSharedKeyAuthenticationRequired(void* ptr);
void* QNetworkAccessManager_Put2(void* ptr, void* request, void* multiPart);
void* QNetworkAccessManager_Put(void* ptr, void* request, void* data);
void* QNetworkAccessManager_Put3(void* ptr, void* request, void* data);
void* QNetworkAccessManager_SendCustomRequest(void* ptr, void* request, void* verb, void* data);
void QNetworkAccessManager_SetCache(void* ptr, void* cache);
void QNetworkAccessManager_SetConfiguration(void* ptr, void* config);
void QNetworkAccessManager_SetCookieJar(void* ptr, void* cookieJar);
void QNetworkAccessManager_SetNetworkAccessible(void* ptr, int accessible);
void QNetworkAccessManager_SetProxy(void* ptr, void* proxy);
void QNetworkAccessManager_SetProxyFactory(void* ptr, void* factory);
char* QNetworkAccessManager_SupportedSchemes(void* ptr);
void QNetworkAccessManager_DestroyQNetworkAccessManager(void* ptr);

#ifdef __cplusplus
}
#endif