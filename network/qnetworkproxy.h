#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QNetworkProxy_NewQNetworkProxy();
QtObjectPtr QNetworkProxy_NewQNetworkProxy3(QtObjectPtr other);
int QNetworkProxy_Capabilities(QtObjectPtr ptr);
int QNetworkProxy_HasRawHeader(QtObjectPtr ptr, QtObjectPtr headerName);
char* QNetworkProxy_Header(QtObjectPtr ptr, int header);
char* QNetworkProxy_HostName(QtObjectPtr ptr);
int QNetworkProxy_IsCachingProxy(QtObjectPtr ptr);
int QNetworkProxy_IsTransparentProxy(QtObjectPtr ptr);
char* QNetworkProxy_Password(QtObjectPtr ptr);
void QNetworkProxy_QNetworkProxy_SetApplicationProxy(QtObjectPtr networkProxy);
void QNetworkProxy_SetCapabilities(QtObjectPtr ptr, int capabilities);
void QNetworkProxy_SetHeader(QtObjectPtr ptr, int header, char* value);
void QNetworkProxy_SetHostName(QtObjectPtr ptr, char* hostName);
void QNetworkProxy_SetPassword(QtObjectPtr ptr, char* password);
void QNetworkProxy_SetRawHeader(QtObjectPtr ptr, QtObjectPtr headerName, QtObjectPtr headerValue);
void QNetworkProxy_SetType(QtObjectPtr ptr, int ty);
void QNetworkProxy_SetUser(QtObjectPtr ptr, char* user);
void QNetworkProxy_Swap(QtObjectPtr ptr, QtObjectPtr other);
int QNetworkProxy_Type(QtObjectPtr ptr);
char* QNetworkProxy_User(QtObjectPtr ptr);
void QNetworkProxy_DestroyQNetworkProxy(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif