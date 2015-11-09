#ifdef __cplusplus
extern "C" {
#endif

void* QNetworkProxy_NewQNetworkProxy();
void* QNetworkProxy_NewQNetworkProxy3(void* other);
int QNetworkProxy_Capabilities(void* ptr);
int QNetworkProxy_HasRawHeader(void* ptr, void* headerName);
void* QNetworkProxy_Header(void* ptr, int header);
char* QNetworkProxy_HostName(void* ptr);
int QNetworkProxy_IsCachingProxy(void* ptr);
int QNetworkProxy_IsTransparentProxy(void* ptr);
char* QNetworkProxy_Password(void* ptr);
void* QNetworkProxy_RawHeader(void* ptr, void* headerName);
void QNetworkProxy_QNetworkProxy_SetApplicationProxy(void* networkProxy);
void QNetworkProxy_SetCapabilities(void* ptr, int capabilities);
void QNetworkProxy_SetHeader(void* ptr, int header, void* value);
void QNetworkProxy_SetHostName(void* ptr, char* hostName);
void QNetworkProxy_SetPassword(void* ptr, char* password);
void QNetworkProxy_SetRawHeader(void* ptr, void* headerName, void* headerValue);
void QNetworkProxy_SetType(void* ptr, int ty);
void QNetworkProxy_SetUser(void* ptr, char* user);
void QNetworkProxy_Swap(void* ptr, void* other);
int QNetworkProxy_Type(void* ptr);
char* QNetworkProxy_User(void* ptr);
void QNetworkProxy_DestroyQNetworkProxy(void* ptr);

#ifdef __cplusplus
}
#endif