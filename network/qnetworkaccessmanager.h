#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QNetworkAccessManager_ProxyFactory(QtObjectPtr ptr);
QtObjectPtr QNetworkAccessManager_NewQNetworkAccessManager(QtObjectPtr parent);
void QNetworkAccessManager_ConnectAuthenticationRequired(QtObjectPtr ptr);
void QNetworkAccessManager_DisconnectAuthenticationRequired(QtObjectPtr ptr);
QtObjectPtr QNetworkAccessManager_Cache(QtObjectPtr ptr);
void QNetworkAccessManager_ClearAccessCache(QtObjectPtr ptr);
QtObjectPtr QNetworkAccessManager_CookieJar(QtObjectPtr ptr);
QtObjectPtr QNetworkAccessManager_DeleteResource(QtObjectPtr ptr, QtObjectPtr request);
void QNetworkAccessManager_ConnectEncrypted(QtObjectPtr ptr);
void QNetworkAccessManager_DisconnectEncrypted(QtObjectPtr ptr);
void QNetworkAccessManager_ConnectFinished(QtObjectPtr ptr);
void QNetworkAccessManager_DisconnectFinished(QtObjectPtr ptr);
QtObjectPtr QNetworkAccessManager_Get(QtObjectPtr ptr, QtObjectPtr request);
QtObjectPtr QNetworkAccessManager_Head(QtObjectPtr ptr, QtObjectPtr request);
int QNetworkAccessManager_NetworkAccessible(QtObjectPtr ptr);
void QNetworkAccessManager_ConnectNetworkAccessibleChanged(QtObjectPtr ptr);
void QNetworkAccessManager_DisconnectNetworkAccessibleChanged(QtObjectPtr ptr);
QtObjectPtr QNetworkAccessManager_Post3(QtObjectPtr ptr, QtObjectPtr request, QtObjectPtr multiPart);
QtObjectPtr QNetworkAccessManager_Post(QtObjectPtr ptr, QtObjectPtr request, QtObjectPtr data);
QtObjectPtr QNetworkAccessManager_Post2(QtObjectPtr ptr, QtObjectPtr request, QtObjectPtr data);
void QNetworkAccessManager_ConnectPreSharedKeyAuthenticationRequired(QtObjectPtr ptr);
void QNetworkAccessManager_DisconnectPreSharedKeyAuthenticationRequired(QtObjectPtr ptr);
QtObjectPtr QNetworkAccessManager_Put2(QtObjectPtr ptr, QtObjectPtr request, QtObjectPtr multiPart);
QtObjectPtr QNetworkAccessManager_Put(QtObjectPtr ptr, QtObjectPtr request, QtObjectPtr data);
QtObjectPtr QNetworkAccessManager_Put3(QtObjectPtr ptr, QtObjectPtr request, QtObjectPtr data);
QtObjectPtr QNetworkAccessManager_SendCustomRequest(QtObjectPtr ptr, QtObjectPtr request, QtObjectPtr verb, QtObjectPtr data);
void QNetworkAccessManager_SetCache(QtObjectPtr ptr, QtObjectPtr cache);
void QNetworkAccessManager_SetConfiguration(QtObjectPtr ptr, QtObjectPtr config);
void QNetworkAccessManager_SetCookieJar(QtObjectPtr ptr, QtObjectPtr cookieJar);
void QNetworkAccessManager_SetNetworkAccessible(QtObjectPtr ptr, int accessible);
void QNetworkAccessManager_SetProxy(QtObjectPtr ptr, QtObjectPtr proxy);
void QNetworkAccessManager_SetProxyFactory(QtObjectPtr ptr, QtObjectPtr factory);
char* QNetworkAccessManager_SupportedSchemes(QtObjectPtr ptr);
void QNetworkAccessManager_DestroyQNetworkAccessManager(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif