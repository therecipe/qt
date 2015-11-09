#ifdef __cplusplus
extern "C" {
#endif

void* QNetworkRequest_NewQNetworkRequest2(void* other);
void* QNetworkRequest_NewQNetworkRequest(void* url);
void* QNetworkRequest_Attribute(void* ptr, int code, void* defaultValue);
int QNetworkRequest_HasRawHeader(void* ptr, void* headerName);
void* QNetworkRequest_Header(void* ptr, int header);
void* QNetworkRequest_OriginatingObject(void* ptr);
int QNetworkRequest_Priority(void* ptr);
void* QNetworkRequest_RawHeader(void* ptr, void* headerName);
void QNetworkRequest_SetAttribute(void* ptr, int code, void* value);
void QNetworkRequest_SetHeader(void* ptr, int header, void* value);
void QNetworkRequest_SetOriginatingObject(void* ptr, void* object);
void QNetworkRequest_SetPriority(void* ptr, int priority);
void QNetworkRequest_SetRawHeader(void* ptr, void* headerName, void* headerValue);
void QNetworkRequest_SetSslConfiguration(void* ptr, void* config);
void QNetworkRequest_SetUrl(void* ptr, void* url);
void QNetworkRequest_Swap(void* ptr, void* other);
void QNetworkRequest_DestroyQNetworkRequest(void* ptr);

#ifdef __cplusplus
}
#endif